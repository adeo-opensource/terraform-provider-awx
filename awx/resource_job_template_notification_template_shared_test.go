package awx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"reflect"
	"runtime"
	"strings"
	"testing"

	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type forTypeTestCase struct {
	typ string
	commonTestCase
}

func getFunctionName(temp interface{}) string {
	strs := strings.Split(runtime.FuncForPC(reflect.ValueOf(temp).Pointer()).Name(), ".")
	return strs[len(strs)-1]
}

func Test_getResourceJobTemplateNotificationTemplateAssociateFuncForType(t *testing.T) {
	type args struct {
		client awx.JobTemplateNotificationTemplatesService
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
			want: "AssociateJobTemplateNotificationTemplatesSuccess-fm",
		},
		{
			name: "Error",
			args: args{
				typ:    "error",
				client: MockAWX{},
			},
			want: "AssociateJobTemplateNotificationTemplatesError-fm",
		},
		{
			name: "Started",
			args: args{
				typ:    "started",
				client: MockAWX{},
			},
			want: "AssociateJobTemplateNotificationTemplatesStarted-fm",
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
			got := getResourceJobTemplateNotificationTemplateAssociateFuncForType(tt.args.client, tt.args.typ)
			if getFunctionName(got) != tt.want {
				t.Errorf("getResourceJobTemplateNotificationTemplateAssociateFuncForType() = %v, want %v", getFunctionName(got), tt.want)
			}
		})
	}
}

func Test_getResourceJobTemplateNotificationTemplateDisassociateFuncForType(t *testing.T) {
	type args struct {
		client awx.JobTemplateNotificationTemplatesService
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
			want: "DisassociateJobTemplateNotificationTemplatesSuccess-fm",
		},
		{
			name: "Error",
			args: args{
				typ:    "error",
				client: MockAWX{},
			},
			want: "DisassociateJobTemplateNotificationTemplatesError-fm",
		},
		{
			name: "Started",
			args: args{
				typ:    "started",
				client: MockAWX{},
			},
			want: "DisassociateJobTemplateNotificationTemplatesStarted-fm",
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
			got := getResourceJobTemplateNotificationTemplateDisassociateFuncForType(tt.args.client, tt.args.typ)
			if getFunctionName(got) != tt.want {
				t.Errorf("getResourceJobTemplateNotificationTemplateAssociateFuncForType() = %v, want %v", getFunctionName(got), tt.want)
			}
		})
	}

}

func Test_resourceJobTemplateNotificationTemplateCreateForType(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "association between template and job",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				id: "4",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "association can't find job template",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Unable to fetch job template",
					Detail:   "Unable to load job template with id 0: got nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))
				},
			},
		},
		{
			typ: "fail",
			commonTestCase: commonTestCase{
				name: "can't find association function",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateStarted().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: JobTemplate not AssociateJobTemplateNotificationTemplates",
					Detail:   "Fail to find association function for notification_template type fail",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "can't associate",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateError().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: JobTemplate not AssociateJobTemplateNotificationTemplates",
					Detail:   "Fail to associate notification_template credentials with ID 1, for job_template ID 0, got error: nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("AssociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, fmt.Errorf("nothing"))
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceJobTemplateNotificationTemplateCreateForType(tt.typ))
		})
	}
}

func Test_resourceJobTemplateNotificationTemplateDeleteForType(t *testing.T) {
	tests := []forTypeTestCase{
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "disassociation between template and job",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				id: "",
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "disassociation can't find job template",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateSuccess().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Unable to fetch job template",
					Detail:   "Unable to load job template with id 0: got nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))
				},
			},
		},
		{
			typ: "fail",
			commonTestCase: commonTestCase{
				name: "can't find disassociation function",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateStarted().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: JobTemplate not DisassociateJobTemplateNotificationTemplates",
					Detail:   "Fail to find disassociation function for notification_template type fail",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, nil)
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
		{
			typ: "success",
			commonTestCase: commonTestCase{
				name: "can't disassociate",
				args: args{
					ctx: context.Background(),
					d:   schema.TestResourceDataRaw(t, resourceJobTemplateNotificationTemplateError().Schema, resourceDataMapNotificationTemplate),
				},
				want: diag.Diagnostics{{
					Severity: diag.Error,
					Summary:  "Create: JobTemplate not DisassociateJobTemplateNotificationTemplates",
					Detail:   "Fail to disassociate notification_template credentials with ID 1, for job_template ID 0, got error: nothing",
				}},
				mock: func(mockAWX *MockAWX) {
					mockAWX.On("DisassociateJobTemplateNotificationTemplatesSuccess", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{ID: 4}, fmt.Errorf("nothing"))
					mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt.commonTestCase, resourceJobTemplateNotificationTemplateDeleteForType(tt.typ))
		})
	}
}
