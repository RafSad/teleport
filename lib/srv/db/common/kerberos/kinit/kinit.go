/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

// Package kinit provides utilities for interacting with a KDC (Key Distribution Center) for Kerberos5
package kinit

import (
	"bytes"
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
	"time"

	"github.com/gravitational/trace"
	"github.com/jcmturner/gokrb5/v8/config"
	"github.com/jcmturner/gokrb5/v8/credentials"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/auth/windows"
)

const (
	// krb5ConfigEnv sets the location from which kinit will attempt to read a configuration value
	krb5ConfigEnv = "KRB5_CONFIG"
	// kinitBinary is the binary Name for the kinit executable
	kinitBinary = "kinit"
	// krb5ConfigTemplate is a configuration template suitable for x509 configuration; it is read by the kinit binary
	krb5ConfigTemplate = `[libdefaults]
 default_realm = {{ .RealmName }}
 rdns = false


[realms]
 {{ .RealmName }} = {
  kdc = {{ .AdminServerName }}
  admin_server = {{ .AdminServerName }}
  pkinit_eku_checking = kpServerAuth
  pkinit_kdc_hostname = {{ .KDCHostName }}
 }`
	// certTTL is the certificate time to live; 1 hour
	certTTL = time.Minute * 60
)

// Provider is a kinit provider capable of producing a credentials cacheData for kerberos
type Provider interface {
	// UseOrCreateCredentials uses or updates an existing cacheData or creates a new one
	UseOrCreateCredentials(ctx context.Context) (cache *credentials.CCache, conf *config.Config, err error)
}

// PKInit is a structure used for initializing a kerberos context
type PKInit struct {
	provider Provider
}

// UseOrCreateCredentialsCache uses or creates a credentials cacheData.
func (k *PKInit) UseOrCreateCredentialsCache(ctx context.Context) (*credentials.CCache, *config.Config, error) {
	return k.provider.UseOrCreateCredentials(ctx)
}

// New returns a new PKInit initializer
func New(provider Provider) *PKInit {
	return &PKInit{provider: provider}
}

// CommandConfig is used to configure a kinit binary execution
type CommandConfig struct {
	// AuthClient is a subset of the auth interface
	AuthClient windows.AuthInterface
	// User is the username of the database/AD user
	User string
	// Realm is the domain name
	Realm string
	// KDCHost is the key distribution center hostname (usually AD server)
	KDCHost string
	// AdminServer is the administration server hostname (usually AD server)
	AdminServer string
	// DataDir is the Teleport Data Directory
	DataDir string
	// LDAPCAPEM contains the same certificate as LDAPCA but in PEM format. It
	// can be used to embed the LDAPCA into files without needing to convert
	// it.
	LDAPCAPEM string
	// Command is a command generator that generates an executable command
	Command CommandGenerator
	// CertGetter is a Teleport Certificate getter that prepares an x509 certificate
	// for use with windows AD
	CertGetter CertGetter
}

// NewCommandLineInitializer returns a new command line initializer using a preinstalled `kinit` binary
func NewCommandLineInitializer(config CommandConfig) *CommandLineInitializer {
	cmd := &CommandLineInitializer{
		auth:               config.AuthClient,
		userName:           config.User,
		cacheName:          fmt.Sprintf("%s@%s", config.User, config.Realm),
		RealmName:          config.Realm,
		KDCHostName:        config.KDCHost,
		AdminServerName:    config.AdminServer,
		dataDir:            config.DataDir,
		certPath:           fmt.Sprintf("%s.pem", config.User),
		keyPath:            fmt.Sprintf("%s-key.pem", config.User),
		binary:             kinitBinary,
		command:            config.Command,
		certGetter:         config.CertGetter,
		ldapCertificatePEM: config.LDAPCAPEM,
		logger:             slog.Default(),
	}
	if cmd.command == nil {
		cmd.command = &execCmd{}
	}
	return cmd
}

// CommandGenerator is a small interface for wrapping *exec.Cmd
type CommandGenerator interface {
	// CommandContext is a wrapper for creating a command
	CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd
}

// execCmd is a small wrapper around exec.Cmd
type execCmd struct {
}

// CommandContext returns exec.CommandContext
func (e *execCmd) CommandContext(ctx context.Context, name string, args ...string) *exec.Cmd {
	return exec.CommandContext(ctx, name, args...)
}

// CommandLineInitializer uses a command line `kinit` binary to provide a kerberos CCache
type CommandLineInitializer struct {
	auth windows.AuthInterface

	// RealmName is the kerberos realm Name (domain Name, like `example.com`
	RealmName string
	// KDCHostName is the key distribution center host Name (usually AD host, like ad.example.com)
	KDCHostName string
	// AdminServerName is the admin server Name (usually AD host)
	AdminServerName string

	dataDir   string
	userName  string
	cacheName string

	certPath string
	keyPath  string
	binary   string

	command    CommandGenerator
	certGetter CertGetter

	ldapCertificatePEM string
	logger             *slog.Logger
}

// CertGetter is an interface for getting a new cert/key pair along with a CA cert
type CertGetter interface {
	// GetCertificateBytes returns a new cert/key pair along with a CA for use with x509 Auth
	GetCertificateBytes(ctx context.Context) (*WindowsCAAndKeyPair, error)
}

// DBCertGetter obtains a new cert/key pair along with the Teleport database CA
type DBCertGetter struct {
	// Auth is the auth client
	Auth windows.AuthInterface
	// Logger is the logger to use.
	Logger *slog.Logger
	// ADConfig is the database-level Active Directory config
	ADConfig types.AD
	// UserName is the database username
	UserName string
}

// WindowsCAAndKeyPair is a wrapper around PEM bytes for Windows authentication
type WindowsCAAndKeyPair struct {
	certPEM []byte
	keyPEM  []byte
	caCert  []byte

	sidLookupError error
}

// GetCertificateBytes returns a new cert/key pem and the DB CA bytes
func (d *DBCertGetter) GetCertificateBytes(ctx context.Context) (*WindowsCAAndKeyPair, error) {
	clusterName, err := d.Auth.GetClusterName(ctx)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	sid, sidLookupError := d.GetActiveDirectorySID(ctx, clusterName.GetClusterName())
	if sidLookupError != nil {
		d.Logger.WarnContext(ctx, "Failed to get SID from ActiveDirectory; PKINIT flow is likely to fail.", "error", sidLookupError)
	}

	req := &windows.GenerateCredentialsRequest{
		CAType:             types.DatabaseClientCA,
		Username:           d.UserName,
		Domain:             d.ADConfig.Domain,
		TTL:                certTTL,
		ClusterName:        clusterName.GetClusterName(),
		AuthClient:         d.Auth,
		ActiveDirectorySID: sid,
		OmitCDP:            true,
	}

	certPEM, keyPEM, caCerts, err := windows.DatabaseCredentials(ctx, req)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	return &WindowsCAAndKeyPair{
		certPEM:        certPEM,
		keyPEM:         keyPEM,
		caCert:         bytes.Join(caCerts, []byte("\n")),
		sidLookupError: sidLookupError,
	}, nil
}

func (d *DBCertGetter) getLDAPCACert() (*x509.Certificate, error) {
	ldapPem, _ := pem.Decode([]byte(d.ADConfig.LDAPCert))
	if ldapPem == nil {
		return nil, trace.BadParameter("cannot find valid LDAP certificate block in AD configuration")
	}
	cert, err := x509.ParseCertificate(ldapPem.Bytes)
	if err != nil {
		return nil, trace.Wrap(err)
	}
	return cert, nil
}

func (d *DBCertGetter) GetActiveDirectorySID(ctx context.Context, clusterName string) (string, error) {
	if d.ADConfig.LDAPServiceAccountName == "" || d.ADConfig.LDAPServiceAccountSID == "" {
		return "", trace.BadParameter("cannot query AD: missing LDAP service principal name or SID")
	}

	ldapCert, err := d.getLDAPCACert()
	if err != nil {
		return "", trace.Wrap(err)
	}

	conn := newLDAPConnector(ldapConnectorConfig{
		logger:      d.Logger,
		authClient:  d.Auth,
		clusterName: clusterName,

		ldapConfig: ldapConfig{
			Address:           d.ADConfig.KDCHostName,
			TLSServerName:     d.ADConfig.KDCHostName,
			Domain:            d.ADConfig.Domain,
			ServiceAccount:    d.ADConfig.LDAPServiceAccountName,
			ServiceAccountSID: d.ADConfig.LDAPServiceAccountSID,
			TLSCACert:         ldapCert,
		},
	})

	sid, err := conn.GetActiveDirectorySID(ctx, d.UserName)

	return sid, trace.Wrap(err)
}

// UseOrCreateCredentials uses an existing cacheData or creates a new one
func (k *CommandLineInitializer) UseOrCreateCredentials(ctx context.Context) (*credentials.CCache, *config.Config, error) {
	tmp, err := os.MkdirTemp("", "kinit")
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	defer func() {
		err = os.RemoveAll(tmp)
		if err != nil {
			k.logger.ErrorContext(ctx, "failed removing temporary kinit directory", "error", err)
		}
	}()

	certPath := filepath.Join(tmp, fmt.Sprintf("%s.pem", k.userName))
	keyPath := filepath.Join(tmp, fmt.Sprintf("%s-key.pem", k.userName))
	userCAPath := filepath.Join(tmp, "userca.pem")

	cacheDir := filepath.Join(k.dataDir, "krb5_cache")
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	cachePath := filepath.Join(cacheDir, k.cacheName)

	wca, err := k.certGetter.GetCertificateBytes(ctx)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	// store files in temp dir
	err = os.WriteFile(certPath, wca.certPEM, 0644)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	err = os.WriteFile(keyPath, wca.keyPEM, 0644)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	err = os.WriteFile(userCAPath, k.buildAnchorsFileContents(wca.caCert), 0644)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	krbConfPath := filepath.Join(tmp, fmt.Sprintf("krb_%s", k.userName))
	err = k.WriteKRB5Config(krbConfPath)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	conf, err := config.Load(krbConfPath)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	cmd := k.command.CommandContext(ctx,
		k.binary,
		"-X", fmt.Sprintf("X509_anchors=FILE:%s", userCAPath),
		"-X", fmt.Sprintf("X509_user_identity=FILE:%s,%s", certPath, keyPath), k.userName,
		"-c", cachePath)

	if cmd.Err != nil {
		return nil, nil, trace.Wrap(cmd.Err)
	}

	cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", krb5ConfigEnv, krbConfPath))
	cmd.Env = append(cmd.Env, "KRB5_TRACE=/dev/stdout")

	k.logger.DebugContext(ctx, "running kinit command", "cmd", cmd, "env", cmd.Env)

	kinitOutput, err := cmd.CombinedOutput()
	if err != nil {
		k.logger.ErrorContext(ctx, "Failed to authenticate with KDC", "cmd_output", string(kinitOutput), "error", err)
		if wca.sidLookupError != nil {
			k.logger.WarnContext(ctx, "The failed request was made with an empty SID due to an LDAP lookup failure. AD servers are likely to reject such requests. The lookup error may be due to a non-existent user or invalid configuration.", "sid_lookup_error", wca.sidLookupError)
		}

		return nil, nil, trace.AccessDenied("authentication failed")
	}
	ccache, err := credentials.LoadCCache(cachePath)
	if err != nil {
		return nil, nil, trace.Wrap(err)
	}

	return ccache, conf, nil
}

// krb5ConfigString returns a config suitable for a kdc
func (k *CommandLineInitializer) krb5ConfigString() (string, error) {
	t, err := template.New("krb_conf").Parse(krb5ConfigTemplate)

	if err != nil {
		return "", trace.Wrap(err)
	}
	b := bytes.NewBuffer([]byte{})
	err = t.Execute(b, k)
	if err != nil {
		return "", trace.Wrap(err)
	}

	return b.String(), nil
}

// WriteKRB5Config writes a krb configuration to path
func (k *CommandLineInitializer) WriteKRB5Config(path string) error {
	s, err := k.krb5ConfigString()
	if err != nil {
		return trace.Wrap(err)
	}

	return trace.Wrap(os.WriteFile(path, []byte(s), 0644))
}

// buildAnchorsFileContents generates the contents of the anchors file (pkinit).
// The file must contain the Teleport DB CA and the KDB/LDAP CA, otherwise the
// connections will fail.
func (k *CommandLineInitializer) buildAnchorsFileContents(caBytes []byte) []byte {
	return append(caBytes, []byte(k.ldapCertificatePEM)...)
}
