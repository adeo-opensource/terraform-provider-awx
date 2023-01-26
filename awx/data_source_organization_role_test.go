package awx

import (
	"context"
	"fmt"
	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceOrganizationRolesRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "No organization_id",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "organization_id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return([]*awx.NotificationTemplate{}, &awx.ListNotificationTemplatesResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Error on get organization",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch organization role",
				Detail:   "Fail to find the organization role got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{}, fmt.Errorf("nothing"))
			},
			newData: nil,
		},
		{
			name: "No Organization role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapOrganization),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Failed to fetch organization role - Not Found",
				Detail:   "The organization role was not found",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{
					Description: "a notification template",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				//TODO: update schema
			},
		},
		{
			name: "First Organization role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 1, Description: "ApplyRole", Name: "ApplyRole"},
						},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "ApplyRole",
			},
		},
		{
			name: "Second Organization role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 2, Description: "ApplyRole", Name: "ApplyRole"},
							ReadRole:  &awx.ApplyRole{ID: 1, Description: "ReadRole", Name: "ReadRole"},
						},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "ReadRole",
			},
		},
		{
			name: " Organization role by name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizationRole().Schema, resourceDataMapOrganization),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetOrganizationsByID", mock.Anything, mock.Anything).Return(&awx.Organization{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 2, Description: "ApplyRole", Name: "ApplyRole"},
							ReadRole:  &awx.ApplyRole{ID: 9, Description: "ReadRole", Name: "ReadRole"},
							AdminRole: &awx.ApplyRole{ID: 4, Description: "foo", Name: "foo"},
						},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceOrganizationRolesRead)
		})
	}
}
