/**
 * Teleport
 * Copyright (C) 2025 Gravitational, Inc.
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

import { ResourceKind, RoleVersion } from 'teleport/services/resources';

import {
  defaultRoleVersion,
  kubernetesResourceKindOptionsMap,
  kubernetesVerbOptionsMap,
  newKubernetesResourceModel,
  ResourceAccess,
  roleToRoleEditorModel,
  RuleModel,
} from './standardmodel';
import {
  KubernetesAccessValidationResult,
  validateAdminRule,
  validateResourceAccess,
  validateRoleEditorModel,
} from './validation';
import { withDefaults } from './withDefaults';

const minimalRoleModel = () =>
  roleToRoleEditorModel(
    withDefaults({
      metadata: { name: 'role-name' },
      version: defaultRoleVersion,
    })
  );

const validity = (arr: { valid: boolean }[]) => arr.map(it => it.valid);

describe('validateRoleEditorModel', () => {
  test('valid minimal model', () => {
    const result = validateRoleEditorModel(
      minimalRoleModel(),
      undefined,
      undefined
    );
    expect(result.metadata.valid).toBe(true);
    expect(result.resources).toEqual([]);
    expect(result.rules).toEqual([]);
    expect(result.isValid).toBe(true);
  });

  test('valid complex model', () => {
    const model = minimalRoleModel();
    model.metadata.labels = [{ name: 'foo', value: 'bar' }];
    model.resources = [
      {
        kind: 'kube_cluster',
        labels: [{ name: 'foo', value: 'bar' }],
        groups: [],
        users: [],
        resources: [
          {
            id: 'dummy-id',
            kind: { label: 'pod', value: 'pod' },
            name: 'res-name',
            namespace: 'dummy-namespace',
            verbs: [],
            roleVersion: defaultRoleVersion,
          },
        ],
        roleVersion: defaultRoleVersion,
        hideValidationErrors: false,
      },
      {
        kind: 'node',
        labels: [{ name: 'foo', value: 'bar' }],
        logins: [{ label: 'root', value: 'root' }],
        hideValidationErrors: false,
      },
      {
        kind: 'app',
        labels: [{ name: 'foo', value: 'bar' }],
        awsRoleARNs: ['some-arn'],
        azureIdentities: ['some-azure-id'],
        gcpServiceAccounts: ['some-gcp-acct'],
        hideValidationErrors: false,
      },
      {
        kind: 'db',
        labels: [{ name: 'foo', value: 'bar' }],
        roles: [{ label: 'some-role', value: 'some-role' }],
        names: [],
        users: [],
        dbServiceLabels: [{ name: 'asdf', value: 'qwer' }],
        hideValidationErrors: false,
      },
      {
        kind: 'windows_desktop',
        labels: [{ name: 'foo', value: 'bar' }],
        logins: [],
        hideValidationErrors: false,
      },
    ];
    model.rules = [
      {
        id: 'dummy-id-1',
        resources: [{ label: ResourceKind.Node, value: ResourceKind.Node }],
        allVerbs: true,
        verbs: [
          { verb: 'read', checked: true },
          { verb: 'list', checked: true },
          { verb: 'create', checked: true },
          { verb: 'update', checked: true },
          { verb: 'delete', checked: true },
          { verb: '*', checked: true },
        ],
        where: '',
        hideValidationErrors: false,
      },
      {
        id: 'dummy-id-2',
        resources: [{ label: ResourceKind.Node, value: ResourceKind.Node }],
        allVerbs: false,
        verbs: [
          { verb: 'read', checked: false },
          { verb: 'list', checked: false },
          { verb: 'create', checked: true },
          { verb: 'update', checked: false },
          { verb: 'delete', checked: true },
          { verb: '*', checked: false },
        ],
        where: '',
        hideValidationErrors: false,
      },
    ];
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(result.metadata.valid).toBe(true);
    expect(validity(result.resources)).toEqual([true, true, true, true, true]);
    expect(validity(result.rules)).toEqual([true, true]);
    expect(result.isValid).toBe(true);
  });

  test('invalid role name', () => {
    const model = minimalRoleModel();
    model.metadata.name = '';
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(result.metadata.valid).toBe(false);
    expect(result.isValid).toBe(false);
  });

  test('conflicting role name', () => {
    const model = minimalRoleModel();
    model.metadata.nameCollision = true;
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(result.metadata.valid).toBe(false);
    expect(result.isValid).toBe(false);
  });

  test('invalid resources', () => {
    const model = minimalRoleModel();
    model.resources = [
      {
        kind: 'node',
        labels: [{ name: 'foo', value: '' }],
        logins: [],
        hideValidationErrors: false,
      },
      {
        kind: 'node',
        labels: [],
        logins: [],
        hideValidationErrors: false,
      },
      {
        kind: 'kube_cluster',
        groups: [],
        labels: [],
        resources: [],
        users: [],
        hideValidationErrors: false,
        roleVersion: RoleVersion.V7,
      },
      {
        kind: 'db',
        labels: [],
        names: [],
        users: [],
        roles: [],
        dbServiceLabels: [],
        hideValidationErrors: false,
      },
      {
        kind: 'app',
        labels: [],
        awsRoleARNs: [],
        azureIdentities: [],
        gcpServiceAccounts: [],
        hideValidationErrors: false,
      },
      {
        kind: 'windows_desktop',
        labels: [],
        logins: [],
        hideValidationErrors: false,
      },
    ];
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(validity(result.resources)).toEqual([
      false,
      false,
      false,
      false,
      false,
      false,
    ]);
    expect(result.isValid).toBe(false);
  });

  test('forbids mixing "*" and other Kubernetes verbs', () => {
    const model = minimalRoleModel();
    model.resources = [
      {
        kind: 'kube_cluster',
        groups: [],
        labels: [],
        users: [],
        resources: [
          {
            ...newKubernetesResourceModel(defaultRoleVersion),
            verbs: [
              kubernetesVerbOptionsMap.get('*'),
              kubernetesVerbOptionsMap.get('get'),
            ],
          },
        ],
        roleVersion: defaultRoleVersion,
        hideValidationErrors: false,
      },
    ];
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(validity(result.resources)).toEqual([false]);
  });

  test.each`
    roleVersion | results
    ${'v3'}     | ${[false, true, false]}
    ${'v4'}     | ${[false, true, false]}
    ${'v5'}     | ${[false, true, false]}
    ${'v6'}     | ${[false, true, false]}
    ${'v7'}     | ${[true, true, true]}
  `(
    'correct types of resources allowed for $roleVersion',
    ({ roleVersion, results }) => {
      const model = minimalRoleModel();
      model.resources = [
        {
          kind: 'kube_cluster',
          groups: [],
          labels: [],
          users: [],
          roleVersion,
          resources: [
            {
              ...newKubernetesResourceModel(defaultRoleVersion),
              kind: kubernetesResourceKindOptionsMap.get('job'),
              roleVersion,
            },
            {
              ...newKubernetesResourceModel(defaultRoleVersion),
              kind: kubernetesResourceKindOptionsMap.get('pod'),
              roleVersion,
            },
            {
              ...newKubernetesResourceModel(defaultRoleVersion),
              kind: kubernetesResourceKindOptionsMap.get('service'),
              roleVersion,
            },
          ],
          hideValidationErrors: false,
        },
      ];
      const result = validateRoleEditorModel(model, undefined, undefined);
      const resourceResult = result
        .resources[0] as KubernetesAccessValidationResult;
      expect(validity(resourceResult.fields.resources.results)).toEqual(
        results
      );
    }
  );

  test('invalid Admin Rules', () => {
    const model = minimalRoleModel();
    model.rules = [
      {
        id: 'dummy-id-1',
        // No resources
        resources: [],
        allVerbs: false,
        verbs: [
          { verb: 'read', checked: false },
          { verb: 'list', checked: false },
          { verb: 'create', checked: true },
          { verb: 'update', checked: false },
          { verb: 'delete', checked: true },
        ],
        where: '',
        hideValidationErrors: false,
      },
      {
        id: 'dummy-id-2',
        resources: [{ label: ResourceKind.Node, value: ResourceKind.Node }],
        allVerbs: false,
        // No verbs
        verbs: [
          { verb: 'read', checked: false },
          { verb: 'list', checked: false },
          { verb: 'create', checked: false },
          { verb: 'update', checked: false },
          { verb: 'delete', checked: false },
        ],
        where: '',
        hideValidationErrors: false,
      },
    ];
    const result = validateRoleEditorModel(model, undefined, undefined);
    expect(validity(result.rules)).toEqual([false, false]);
    expect(result.isValid).toBe(false);
  });

  it('reuses previously computed section results', () => {
    const model = minimalRoleModel();
    const result1 = validateRoleEditorModel(model, undefined, undefined);
    const result2 = validateRoleEditorModel(model, model, result1);
    expect(result2.metadata).toBe(result1.metadata);
    expect(result2.resources).toBe(result1.resources);
    expect(result2.rules).toBe(result1.rules);
    expect(result2.isValid).toBe(result1.isValid);
  });
});

describe('validateResourceAccess', () => {
  it('reuses previously computed results', () => {
    const resource: ResourceAccess = {
      kind: 'node',
      labels: [],
      logins: [],
      hideValidationErrors: false,
    };
    const result1 = validateResourceAccess(resource, undefined, undefined);
    const result2 = validateResourceAccess(resource, resource, result1);
    expect(result2).toBe(result1);
  });
});

describe('validateAdminRule', () => {
  it('reuses previously computed results', () => {
    const rule: RuleModel = {
      id: 'some-id',
      resources: [],
      allVerbs: false,
      verbs: [],
      where: '',
      hideValidationErrors: false,
    };
    const result1 = validateAdminRule(rule, undefined, undefined);
    const result2 = validateAdminRule(rule, rule, result1);
    expect(result2).toBe(result1);
  });
});
