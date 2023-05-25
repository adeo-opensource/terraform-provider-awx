package awx

import (
	"context"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceProjectsRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Two project",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Project, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListProjects", mock.Anything).Return([]*awx.Project{{}, {}}, &awx.ListProjectsResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One project",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceProject().Schema, resourceDataMapProject),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListProjects", mock.Anything).Return([]*awx.Project{{
					Description: "a project",
					Name:        "the project",
					Credential:  "1",
				}}, &awx.ListProjectsResponse{}, nil)
			},
			newData: map[string]interface{}{
				//TODO: schema
				"name": "the project",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceProjectsRead)

		})
	}
}
