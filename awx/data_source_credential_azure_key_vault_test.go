package awx

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	awx "github.com/denouche/goawx/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_dataSourceCredentialAzureRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Missing parameters",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialAzure().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "credential_id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
			},
		},
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialAzure().Schema, resourceDataMapCredential),
			},
			want: diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to fetch credentials",
				Detail:   fmt.Sprintf("Unable to credentials with credentialId %d: %s", 1, "error"),
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(&awx.Credential{}, fmt.Errorf("error"))
			},
		},
		{
			name: "Credentials founds",
			args: args{ctx: context.Background(), d: schema.TestResourceDataRaw(t, dataSourceCredentialAzure().Schema, resourceDataMapCredential)},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(&awx.Credential{
					ID:             1,
					Name:           "foo",
					Description:    "a description",
					Kind:           "toto",
					OrganizationID: 1,
					Inputs: map[string]interface{}{
						"username": "borto",
						"url":      "awx.url.foo",
						"client":   "test",
						"tenant":   "adeo-oss"},
				}, nil)
			},
			newData: map[string]interface{}{
				"name":            "foo",
				"description":     "a description",
				"organization_id": 1,
				"url":             "awx.url.foo",
				"client":          "test",
				"secret":          "terces",
				"tenant":          "adeo-oss",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceCredentialAzureRead)
		})
	}
}
