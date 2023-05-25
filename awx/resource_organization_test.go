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

func Test_resourceOrganizationCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Organizations not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Organizations",
				Detail:   "Organizations with name foo, failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				organizations := &awx.Organization{}
				mockAWX.On("CreateOrganization",
					mock.Anything,
					mock.Anything).
					Return(organizations, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Organizations created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				organizations := &awx.Organization{}
				mockAWX.On("CreateOrganization",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						organizations.ID = 2
						organizations.Description = data["description"].(string) + "_created"
						organizations.Name = data["name"].(string) + "_created"
					}).
					Return(organizations, nil)
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(organizations, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceOrganizationsCreate)
		})
	}
}

func Test_resourceOrganizationDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Organizations not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Organization delete failed",
				Detail:   "Fail to delete Organization, OrganizationID 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteOrganization", mock.Anything, mock.Anything).Return(&awx.Organization{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Organizations deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteOrganization", mock.Anything, mock.Anything).Return(&awx.Organization{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceOrganizationsDelete)
		})
	}
}

func Test_resourceOrganizationRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Organization not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Organization",
				Detail:   "Unable to load Organization with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				organizations := &awx.Organization{}
				mockAWX.On("GetOrganizationsByID",
					mock.Anything,
					mock.Anything).
					Return(organizations, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Organizations found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				organizations := &awx.Organization{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetOrganizationsByID",
					mock.Anything,
					mock.Anything).
					Return(organizations, nil)
			},
			newData: map[string]interface{}{
				"description": "data_read",
				"name":        "toto",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceOrganizationsRead)
		})
	}
}

func Test_resourceOrganizationUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Organization not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Organizations",
				Detail:   "Unable to load Organizations with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "Organization cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Organizations",
				Detail:   "Organizations with name foo failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{}, nil)
				mockAWX.On("UpdateOrganization", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Organization{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Organization updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceOrganization().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{}, nil)
				mockAWX.On("UpdateOrganization", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Organization{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceOrganizationsUpdate)
		})
	}
}
