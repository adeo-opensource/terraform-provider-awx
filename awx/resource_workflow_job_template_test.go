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

func Test_resourceWorkflowJobTemplateCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplate cannot be created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create WorkflowJobTemplate",
				Detail:   "WorkflowJobTemplate with name foo failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplate := &awx.WorkflowJobTemplate{}
				mockAWX.On("CreateWorkflowJobTemplate",
					mock.Anything,
					mock.Anything).
					Return(workflowJobTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplate created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplate := &awx.WorkflowJobTemplate{}
				mockAWX.On("CreateWorkflowJobTemplate",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						workflowJobTemplate.ID = 2
						workflowJobTemplate.Description = data["description"].(string) + "_created"
						workflowJobTemplate.Name = data["name"].(string) + "_created"
					}).
					Return(workflowJobTemplate, nil)
				mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(workflowJobTemplate, nil)
				mockAWX.On("ListWorkflowJobTemplates",
					mock.Anything).
					Return([]*awx.WorkflowJobTemplate{workflowJobTemplate}, &awx.ListWorkflowJobTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceWorkflowJobTemplateCreate)
		})
	}
}

func Test_resourceWorkflowJobTemplateDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplate not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "WorkflowJobTemplate delete failed",
				Detail:   "Fail to delete WorkflowJobTemplate, id 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteWorkflowJobTemplate", mock.Anything).Return(&awx.WorkflowJobTemplate{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "WorkflowJobTemplate deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteWorkflowJobTemplate", mock.Anything).Return(&awx.WorkflowJobTemplate{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, nil)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceWorkflowJobTemplateDelete)
		})
	}
}

func Test_resourceWorkflowJobTemplateRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch workflow job template",
				Detail:   "Unable to load workflow job template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplate := &awx.WorkflowJobTemplate{}
				mockAWX.On("GetWorkflowJobTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(workflowJobTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplate found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplate := &awx.WorkflowJobTemplate{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetWorkflowJobTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(workflowJobTemplate, nil)
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
			runTestCase(t, tt, resourceWorkflowJobTemplateRead)
		})
	}
}

func Test_resourceWorkflowJobTemplateUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplate cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update WorkflowJobTemplate",
				Detail:   "WorkflowJobTemplate with name foo failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				mockAWX.On("UpdateWorkflowJobTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "WorkflowJobTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job Workflow template",
				Detail:   "Unable to load job Workflow template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, fmt.Errorf("nothing"))
				mockAWX.On("UpdateWorkflowJobTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "WorkflowJobTemplate updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				mockAWX.On("UpdateWorkflowJobTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceWorkflowJobTemplateUpdate)
		})
	}
}
