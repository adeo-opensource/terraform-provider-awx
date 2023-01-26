package awx

import (
	"context"
	"fmt"
	awx "github.com/denouche/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceCredentialInputSourceCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create new credentials",
				Detail:   "Unable to create new credentials: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialInputSource{}
				mockAWX.On("CreateCredentialInputSource",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Credentials created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialInputSource{}
				mockAWX.On("CreateCredentialInputSource",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						credentials.ID = 2
						credentials.Description = data["description"].(string) + "_created"
						credentials.SourceCredential = data["source_credential"].(int)
						credentials.TargetCredential = data["target_credential"].(int)
						credentials.InputFieldName = data["input_field_name"].(string) + "_created"
						credentials.Metadata = data["metadata"].(map[string]interface{})
					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialInputSourceByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"description":      "data_created",
				"input_field_name": "input_field_created",
				"target":           3,
				"source":           42,
				"metadata":         map[string]interface{}{"key": "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialInputSourceCreate)
		})
	}
}

func Test_resourceCredentialInputSourceDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to delete existing credentials",
				Detail:   "Unable to delete existing credentials with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteCredentialInputSourceByID", mock.Anything, mock.Anything).Return(fmt.Errorf("nothing"))
			},
		},
		{
			name: "Credentials deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteCredentialInputSourceByID", mock.Anything, mock.Anything).Return(nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialInputSourceDelete)
		})
	}
}

func Test_resourceCredentialInputSourceRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch credentials",
				Detail:   "Unable to credentials with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialInputSource{}
				mockAWX.On("GetCredentialInputSourceByID",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialInputSourceRead)
		})
	}
}

func Test_resourceCredentialInputSourceUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update existing credentials",
				Detail:   "Unable to update existing credentials with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialInputSource{}
				mockAWX.On("UpdateCredentialInputSourceByID",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Credentials updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialInputSource().Schema, resourceDataMapCredentialInput),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialInputSource{}
				mockAWX.On("UpdateCredentialInputSourceByID",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						credentials.ID = 2
						credentials.Description = data["description"].(string) + "_updated"
						credentials.SourceCredential = data["source_credential"].(int)
						credentials.TargetCredential = data["target_credential"].(int)
						credentials.InputFieldName = data["input_field_name"].(string) + "_updated"
						credentials.Metadata = data["metadata"].(map[string]interface{})

					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialInputSourceByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"description":      "data_updated",
				"input_field_name": "input_field_updated",
				"target":           3,
				"source":           42,
				"metadata":         map[string]interface{}{"key": "value"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialInputSourceUpdate)
		})
	}
}
