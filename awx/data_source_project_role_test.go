package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceProjectRolesRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Missing parameter",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProjectRole().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "project_id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
			},
		},
		{
			name: "No project",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProjectRole().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Project",
				Detail:   "Fail to find the project got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{Name: "orga"}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "One project",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProjectRole().Schema, resourceDataMapProject),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(
					&awx.Project{
						ID:   1,
						Name: "orga",
						SummaryFields: &awx.Summary{
							ObjectRoles: &awx.ObjectRoles{
								UseRole: &awx.ApplyRole{ID: 1, Description: "ApplyRole", Name: "ApplyRole"},
							},
						}}, nil)
			},
			newData: map[string]interface{}{
				//TODO: schema
				"name": "ApplyRole",
			},
		},
		{
			name: "One project match with name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProjectRole().Schema, resourceDataMapProject),
			},
			want: nil,
			id:   "3",
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(
					&awx.Project{
						ID:   1,
						Name: "orga",
						SummaryFields: &awx.Summary{
							ObjectRoles: &awx.ObjectRoles{
								UseRole: &awx.ApplyRole{ID: 3, Description: "foo", Name: "foo"},
							},
						}}, nil)
			},
			newData: map[string]interface{}{
				//TODO: schema
				"name": "foo",
			},
		},
		{
			name: "One project match without name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProjectRole().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Failed to fetch project role - Not Found",
				Detail:   "The project role was not found",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(
					&awx.Project{
						ID:   1,
						Name: "orga",
						SummaryFields: &awx.Summary{
							ObjectRoles: &awx.ObjectRoles{
								UseRole: &awx.ApplyRole{ID: 3, Description: "ApplyRole", Name: "ApplyRole"},
							},
						}}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceProjectRolesRead)

		})
	}
}
