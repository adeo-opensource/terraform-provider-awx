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

func Test_resourceJobTemplateCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "JobTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create JobTemplate",
				Detail:   "JobTemplate with name foo in the project id 10293, failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{}
				mockAWX.On("CreateJobTemplate",
					mock.Anything,
					mock.Anything).
					Return(jobTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "JobTemplate created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{}
				mockAWX.On("CreateJobTemplate",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						jobTemplate.ID = 2
						jobTemplate.Description = data["description"].(string) + "_created"
						jobTemplate.Name = data["name"].(string) + "_created"
						jobTemplate.Verbosity = data["verbosity"].(int)

					}).
					Return(jobTemplate, nil)
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(jobTemplate, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceJobTemplateCreate)
		})
	}
}

func Test_resourceJobTemplateRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Job Template not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job template",
				Detail:   "Unable to load job template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{}
				mockAWX.On("GetJobTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(jobTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "JobTemplate found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetJobTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(jobTemplate, nil)
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
			runTestCase(t, tt, resourceJobTemplateRead)
		})
	}
}

func Test_resourceJobTemplateUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Job template not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job template",
				Detail:   "Unable to load job template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "Job template cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update JobTemplate",
				Detail:   "JobTemplate with name foo in the project id 10293 failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				mockAWX.On("UpdateJobTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Job template updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplate().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				mockAWX.On("UpdateJobTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceJobTemplateUpdate)
		})
	}
}
