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
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(&awx.Credential{}, fmt.Errorf("error"))
			},
		},
		{
			name: "Credentials found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialByID().Schema, resourceDataMapCredential),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(&awx.Credential{ID: 1, Kind: "toto", Inputs: map[string]interface{}{"username": "borto"}}, nil)
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
