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

func Test_resourceWorkflowJobTemplateNodeCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplateNode not created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create WorkflowJobTemplateNode",
				Detail:   "WorkflowJobTemplateNode with JobTemplateID 0 and WorkflowID: 0 failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				WorkflowJobTemplateNode := &awx.WorkflowJobTemplateNode{}
				mockAWX.On("CreateWorkflowJobTemplateNode",
					mock.Anything,
					mock.Anything).
					Return(WorkflowJobTemplateNode, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplateNode created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				WorkflowJobTemplateNode := &awx.WorkflowJobTemplateNode{}
				mockAWX.On("CreateWorkflowJobTemplateNode",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						WorkflowJobTemplateNode.ID = 2
					}).
					Return(WorkflowJobTemplateNode, nil)
				mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(WorkflowJobTemplateNode, nil)
			},
			id: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceWorkflowJobTemplateNodeCreate)
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplateNode not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "WorkflowJobTemplateNode delete failed",
				Detail:   "Fail to delete WorkflowJobTemplateNode, id 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteWorkflowJobTemplateNode", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplateNode deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteWorkflowJobTemplateNode", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceWorkflowJobTemplateNodeDelete)
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplateNode not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch workflow job template node",
				Detail:   "Unable to load workflow job template node with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				WorkflowJobTemplateNode := &awx.WorkflowJobTemplateNode{}
				mockAWX.On("GetWorkflowJobTemplateNodeByID",
					mock.Anything,
					mock.Anything).
					Return(WorkflowJobTemplateNode, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplateNode found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				WorkflowJobTemplateNode := &awx.WorkflowJobTemplateNode{
					ID: 3,
				}
				mockAWX.On("GetWorkflowJobTemplateNodeByID",
					mock.Anything,
					mock.Anything).
					Return(WorkflowJobTemplateNode, nil)
			},
			id: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceWorkflowJobTemplateNodeRead)
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplateNode not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch workflow job template node",
				Detail:   "Unable to load workflow job template node with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "WorkflowJobTemplateNode cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update WorkflowJobTemplateNode",
				Detail:   "WorkflowJobTemplateNode with id 0 failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, nil)
				mockAWX.On("UpdateWorkflowJobTemplateNode", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "WorkflowJobTemplateNode updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNode().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, nil)
				mockAWX.On("UpdateWorkflowJobTemplateNode", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceWorkflowJobTemplateNodeUpdate)
		})
	}
}
