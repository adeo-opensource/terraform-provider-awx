package awx

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"

	awx "github.com/adeo-opensource/goawx/client"
)

func Test_getCreateWorkflowJobTemplateNodeStepFuncForType(t *testing.T) {
	type args struct {
		client awx.WorkflowJobTemplateNodeStepService
		typ    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Success",
			args: args{
				typ:    "success",
				client: MockAWX{},
			},
			want: "CreateWorkflowJobTemplateSuccessNodeStep-fm",
		},
		{
			name: "failure",
			args: args{
				typ:    "failure",
				client: MockAWX{},
			},
			want: "CreateWorkflowJobTemplateFailureNodeStep-fm",
		},
		{
			name: "always",
			args: args{
				typ:    "always",
				client: MockAWX{},
			},
			want: "CreateWorkflowJobTemplateAlwaysNodeStep-fm",
		},
		{
			name: "Default",
			args: args{
				typ:    "default",
				client: MockAWX{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCreateWorkflowJobTemplateNodeStepFuncForType(tt.args.client, tt.args.typ)
			if getFunctionName(got) != tt.want {
				t.Errorf("getResourceJobTemplateNotificationTemplateAssociateFuncForType() = %v, want %v", getFunctionName(got), tt.want)
			}
		})
	}
}

func Test_createNodeForWorkflowJob(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "create node for success workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeSuccess().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateSuccessNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
		{
			typ: "failure",
			commonTestCase: commonTestCase{
				name: "create node for failure workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeFailure().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateFailureNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
		{
			typ: "always",
			commonTestCase: commonTestCase{
				name: "create node for always workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeAlways().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateAlwaysNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "can't create node",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeSuccess().Schema, resourceDataMapWorkflowJobTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Unable to create WorkflowJobTemplateNodeSuccess",
					Detail:   "WorkflowJobTemplateNodeSuccess with JobTemplateID 0 failed to create nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateSuccessNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, fmt.Errorf("nothing"))
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
				return createNodeForWorkflowJob(m.(awx.WorkflowJobTemplateNodeStepService), tt.typ, ctx, d, m)
			})
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeAlwaysCreate(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "always",
			commonTestCase: commonTestCase{
				name: "create node for always workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeAlways().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateAlwaysNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceWorkflowJobTemplateNodeAlwaysCreate)
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeSuccessCreate(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "create node for always workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeSuccess().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateSuccessNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceWorkflowJobTemplateNodeSuccessCreate)
		})
	}
}

func Test_resourceWorkflowJobTemplateNodeFailureCreate(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "failure",
			commonTestCase: commonTestCase{
				name: "create node for failure workflow",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNodeFailure().Schema, resourceDataMapWorkflowJobTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("CreateWorkflowJobTemplateFailureNodeStep", mock.Anything, mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateNodeByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplateNode{ID: 4}, nil)

				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceWorkflowJobTemplateNodeFailureCreate)
		})
	}
}
