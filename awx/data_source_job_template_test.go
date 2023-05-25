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

func Test_dataSourceJobTemplateRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Error on list job template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceJobTemplate().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Job templates",
				Detail:   "Fail to find the job template got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListJobTemplates", mock.Anything).Return([]*awx.JobTemplate{}, &awx.ListJobTemplatesResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Two job template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Job template, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListJobTemplates", mock.Anything).Return([]*awx.JobTemplate{{}, {}}, &awx.ListJobTemplatesResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One job template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListJobTemplates", mock.Anything).Return([]*awx.JobTemplate{{
					Description: "a jobtemplate",
					Name:        "bar",
				}}, &awx.ListJobTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "bar",

				//TODO: update schema
			},
		},
		{
			name: "One job template match name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListJobTemplates", mock.Anything).Return([]*awx.JobTemplate{{
					Description: "a jobtemplate",
					Name:        "foo",
				}}, &awx.ListJobTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo",
			},
		},
		{
			name: "One job template not matching",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceJobTemplate().Schema, resourceDataMapMissingId),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListJobTemplates", mock.Anything).Return([]*awx.JobTemplate{{
					Description: "a jobtemplate",
				}}, &awx.ListJobTemplatesResponse{}, nil)
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Job template, 1",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceJobTemplateRead)
		})
	}
}
