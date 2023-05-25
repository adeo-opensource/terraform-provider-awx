package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceCredentialSCMCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialSCM().Schema, resourceDataMapCredential),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create new credentials",
				Detail:   "Unable to create new credentials: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.Credential{}
				mockAWX.On("CreateCredentials",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Credentials created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialSCM().Schema, resourceDataMapCredential),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.Credential{}
				mockAWX.On("CreateCredentials",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						credentials.ID = 2
						credentials.Name = data["name"].(string) + "_created"
						credentials.Description = data["description"].(string) + "_created"
						credentials.OrganizationID = data["organization"].(int)
						credentials.Inputs = data["inputs"].(map[string]interface{})
						credentials.CredentialTypeID = data["credential_type"].(int)
					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"name":            "foo_created",
				"description":     "data_created",
				"organization_id": 1,
				"username":        "user",
				"password":        "pass",
				"ssh_key_data":    "a ssh key",
				"ssh_key_unlock":  "alomora",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialSCMCreate)
		})
	}
}

func Test_resourceCredentialSCMRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialSCM().Schema, resourceDataMapCredential),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch credentials",
				Detail:   "Unable to credentials with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.Credential{}
				mockAWX.On("GetCredentialsByID",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialSCMRead)
		})
	}
}

func Test_resourceCredentialSCMUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialSCM().Schema, resourceDataMapCredential),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update existing credentials",
				Detail:   "Unable to update existing credentials with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.Credential{}
				mockAWX.On("UpdateCredentialsByID",
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
				d:   schema.TestResourceDataRaw(t, resourceCredentialSCM().Schema, resourceDataMapCredential),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.Credential{}
				mockAWX.On("UpdateCredentialsByID",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						credentials.ID = 2
						credentials.Name = data["name"].(string) + "_updated"
						credentials.Description = data["description"].(string) + "_updated"
						credentials.OrganizationID = data["organization"].(int)
						credentials.Inputs = data["inputs"].(map[string]interface{})
						credentials.CredentialTypeID = data["credential_type"].(int)
					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialsByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"name":            "foo_updated",
				"description":     "data_updated",
				"organization_id": 1,
				"username":        "user",
				"password":        "pass",
				"ssh_key_data":    "a ssh key",
				"ssh_key_unlock":  "alomora",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialSCMUpdate)
		})
	}
}
