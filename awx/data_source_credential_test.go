package awx

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_dataSourceCredentialByIDRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Missing parameters",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialByID().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
			},
		},
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialByID().Schema, resourceDataMapCredential),
			},
			want: diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to fetch credential",
				Detail:   "The given credential ID is invalid or malformed",
			}},
			mock: func(mockAWX *MockAWX) {
				mockCredentialService := mockAWX.CredentialService.(mockGeneric[awx.Credential])
				mockCredentialService.On("GetByID", mock.Anything, mock.Anything).Return(&awx.Credential{}, fmt.Errorf("error"))
				mockAWX.CredentialService = mockCredentialService
			},
		},
		{
			name: "Credential found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialByID().Schema, resourceDataMapCredential),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockCredentialService := mockAWX.CredentialService.(mockGeneric[awx.Credential])
				mockCredentialService.On("GetByID", mock.Anything, mock.Anything).Return(&awx.Credential{ID: 1, Kind: "toto", Inputs: map[string]interface{}{"username": "borto"}}, nil)
				mockAWX.CredentialService = mockCredentialService
			},
			newData: map[string]interface{}{
				"username": "borto",
				"kind":     "toto",
				"tower_id": 1,
			},
			id: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceCredentialByIDRead)
		})
	}
}
