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

func Test_dataSourceCredentialTypeByIDRead(t *testing.T) {
	// some assertions here

	tests := []commonTestCase{
		{
			name: "Missing parameters",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialTypeByID().Schema, resourceDataMapMissingId),
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
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialTypeByID().Schema, resourceDataMapCredentialType),
			},
			want: diag.Diagnostics{diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to fetch credential type",
				Detail:   "Unable to fetch credential type with ID: 1. Error: error",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetCredentialTypeByID", mock.Anything, mock.Anything).Return(&awx.CredentialType{}, fmt.Errorf("error"))
			},
		},
		{
			name: "Credentials found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentialTypeByID().Schema, resourceDataMapCredentialType),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetCredentialTypeByID", mock.Anything, mock.Anything).Return(&awx.CredentialType{ID: 1, Injectors: "inject", Name: "credType", Description: "cred description", Kind: "toto", Inputs: "runTestCase"}, nil)
			},
			newData: map[string]interface{}{
				"kind":        "toto",
				"name":        "credType",
				"description": "cred description",
				"injectors":   "inject",
				"inputs":      "runTestCase",
			},
			id: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceCredentialTypeByIDRead)
		})
	}
}
