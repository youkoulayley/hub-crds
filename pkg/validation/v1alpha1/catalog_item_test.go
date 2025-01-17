/*
Copyright (C) 2022-2025 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package v1alpha1_test

import (
	"testing"

	"k8s.io/apimachinery/pkg/util/validation/field"
)

func TestAPICatalogItem_Validation(t *testing.T) {
	t.Parallel()

	tests := []validationTestCase{
		{
			desc: "missing resource namespace",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: "my-catalog-items"
`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.namespace", BadValue: ""}},
		},
		{
			desc: "valid: minimal",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default`),
		},
		{
			desc: "valid: full",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  groups:
    - my-group
  apis:
    - name: my-api
  apiSelector:
    matchLabels:
      key: value
  operationFilter:
    include:
      - my-filter`),
		},
		{
			desc: "invalid resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: .non-dns-compliant-catalog-items
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: ".non-dns-compliant-catalog-items", Detail: "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character (e.g. 'my-name',  or '123-abc', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?')"}},
		},
		{
			desc: "missing resource name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: ""
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "metadata.name", BadValue: "", Detail: "name or generateName is required"}},
		},
		{
			desc: "resource name is too long",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: catalog-items-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name
  namespace: default`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "metadata.name", BadValue: "catalog-items-with-a-way-toooooooooooooooooooooooooooooooooooooo-long-name", Detail: "must be no more than 63 characters"}},
		},
		{
			desc: "duplicated APIs",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "duplicated API: implicit default",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  apis:
    - name: my-api
    - name: my-api`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec.apis", BadValue: "array", Detail: "duplicated apis"}},
		},
		{
			desc: "invalid API selector",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  apiSelector:
    matchExpressions:
      - key: value`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiSelector.matchExpressions[0].operator", BadValue: ""}},
		},
		{
			desc: "everyone and groups both set",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan:
    name: my-plan
  everyone: true
  groups:
    - my-group`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeInvalid, Field: "spec", BadValue: "object", Detail: "groups and everyone are mutually exclusive"}},
		},
		{
			desc: "missing apiPlan name",
			manifest: []byte(`
apiVersion: hub.traefik.io/v1alpha1
kind: APICatalogItem
metadata:
  name: my-catalog-items
  namespace: default
spec:
  apiPlan: {}`),
			wantErrs: field.ErrorList{{Type: field.ErrorTypeRequired, Field: "spec.apiPlan.name", BadValue: ""}},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			checkValidation(t, test)
		})
	}
}
