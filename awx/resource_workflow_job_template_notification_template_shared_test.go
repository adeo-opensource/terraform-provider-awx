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

func Test_getResourceWorkflowJobTemplateNotificationTemplateAssociateFuncForType(t *testing.T) {
	type args struct {
		client awx.WorkflowJobTemplateNotificationTemplatesService
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
			want: "AssociateWorkflowJobTemplateNotificationTemplatesSuccess-fm",
		},
		{
			name: "error",
			args: args{
				typ:    "error",
				client: MockAWX{},
			},
			want: "AssociateWorkflowJobTemplateNotificationTemplatesError-fm",
		},
		{
			name: "started",
			args: args{
				typ:    "started",
				client: MockAWX{},
			},
			want: "AssociateWorkflowJobTemplateNotificationTemplatesStarted-fm",
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
			got := getResourceWorkflowJobTemplateNotificationTemplateAssociateFuncForType(tt.args.client, tt.args.typ)
			if getFunctionName(got) != tt.want {
				t.Errorf("getResourceWorkflowJobTemplateNotificationTemplateAssociateFuncForType() = %v, want %v", getFunctionName(got), tt.want)
			}
		})
	}
}

func Test_getResourceWorkflowJobTemplateNotificationTemplateDisassociateFuncForType(t *testing.T) {
	type args struct {
		client awx.WorkflowJobTemplateNotificationTemplatesService
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
			want: "DisassociateWorkflowJobTemplateNotificationTemplatesSuccess-fm",
		},
		{
			name: "error",
			args: args{
				typ:    "error",
				client: MockAWX{},
			},
			want: "DisassociateWorkflowJobTemplateNotificationTemplatesError-fm",
		},
		{
			name: "started",
			args: args{
				typ:    "started",
				client: MockAWX{},
			},
			want: "DisassociateWorkflowJobTemplateNotificationTemplatesStarted-fm",
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
			got := getResourceWorkflowJobTemplateNotificationTemplateDisassociateFuncForType(tt.args.client, tt.args.typ)
			if getFunctionName(got) != tt.want {
				t.Errorf("getResourceWorkflowJobTemplateNotificationTemplateDisassociateFuncForType() = %v, want %v", getFunctionName(got), tt.want)
			}
		})
	}
}

func Test_resourceWorkflowJobTemplateNotificationTemplateCreateForType(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "association between template and job",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "association can't find job template",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Unable to fetch workflow job template",
					Detail:   "Unable to load workflow job template with id 0: got nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, fmt.Errorf("nothing"))
				},
			},
		},
		{
			typ: "fail",
			commonTestCase: commonTestCase{
				name: "can't find association function",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateStarted().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: WorkflowJobTemplate not AssociateWorkflowJobTemplateNotificationTemplates",
					Detail:   "Fail to find association function for notification_template type fail",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "can't associate",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateError().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: WorkflowJobTemplate not AssociateWorkflowJobTemplateNotificationTemplates",
					Detail:   "Fail to associate notification_template credentials with ID 1, for workflow_job_template ID 0, got error: nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, fmt.Errorf("nothing"))
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceWorkflowJobTemplateNotificationTemplateCreateForType(tt.typ))
		})
	}
}

func Test_resourceWorkflowJobTemplateNotificationTemplateDeleteForType(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "disassociation between template and job",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				id: "",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "disassociation can't find job template",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Unable to fetch workflow job template",
					Detail:   "Unable to load workflow job template with id 0: got nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, fmt.Errorf("nothing"))
				},
			},
		},
		{
			typ: "fail",
			commonTestCase: commonTestCase{
				name: "can't find disassociation function",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateStarted().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: WorkflowJobTemplate not DisassociateWorkflowJobTemplateNotificationTemplates",
					Detail:   "Fail to find disassociation function for notification_template type fail",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "can't disassociate",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateNotificationTemplateError().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: WorkflowJobTemplate not DisassociateWorkflowJobTemplateNotificationTemplates",
					Detail:   "Fail to disassociate notification_template credentials with ID 1, for job_template ID 0, got error: nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateWorkflowJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, fmt.Errorf("nothing"))
					mockAWX.On("GetWorkflowJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.WorkflowJobTemplate{}, nil)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceWorkflowJobTemplateNotificationTemplateDeleteForType(tt.typ))
		})
	}
}
