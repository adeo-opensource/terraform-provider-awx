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

func Test_resourceCredentialTypeCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialType().Schema, resourceDataMapCredentialType),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create new credential type",
				Detail:   "Unable to create new credential type: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialType{}
				mockAWX.On("CreateCredentialType",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Credentials created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialType().Schema, resourceDataMapCredentialType),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialType{}
				mockAWX.On("CreateCredentialType",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						credentials.ID = 2
						credentials.Name = data["name"].(string) + "_created"
						credentials.Description = data["description"].(string) + "_created"
					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialTypeByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"name":        "foo_created",
				"description": "data_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialTypeCreate)
		})
	}
}

func Test_resourceCredentialTypeRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialType().Schema, resourceDataMapCredentialType),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch credential type",
				Detail:   "Unable to credential type with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialType{}
				mockAWX.On("GetCredentialTypeByID",
					mock.Anything,
					mock.Anything).
					Return(credentials, fmt.Errorf("nothing"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialTypeRead)
		})
	}
}

func Test_resourceCredentialTypeUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Credentials not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceCredentialType().Schema, resourceDataMapCredentialType),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update existing credential type",
				Detail:   "Unable to update existing credential type with id 0: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialType{}
				mockAWX.On("UpdateCredentialTypeByID",
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
				d:   schema.TestResourceDataRaw(t, resourceCredentialType().Schema, resourceDataMapCredentialType),
			},
			mock: func(mockAWX *MockAWX) {
				credentials := &awx.CredentialType{}
				mockAWX.On("UpdateCredentialTypeByID",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						credentials.ID = 2
						credentials.Name = data["name"].(string) + "_updated"
						credentials.Description = data["description"].(string) + "_updated"
						//TODO: validate input
					}).
					Return(credentials, nil)
				mockAWX.On("GetCredentialTypeByID", mock.Anything, mock.Anything).Return(credentials, nil)

			},
			newData: map[string]interface{}{
				"name":        "foo_updated",
				"description": "data_updated",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceCredentialTypeUpdate)
		})
	}
}
