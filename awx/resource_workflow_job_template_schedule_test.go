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

func Test_resourceWorkflowJobTemplateScheduleCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "WorkflowJobTemplateSchedule not created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateSchedule().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Schedule",
				Detail:   "Schedule failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplateSchedule := &awx.Schedule{}
				mockAWX.On("CreateWorkflowJobTemplateSchedule",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(workflowJobTemplateSchedule, fmt.Errorf("nothing"))
			},
		},
		{
			name: "WorkflowJobTemplateSchedule created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceWorkflowJobTemplateSchedule().Schema, resourceDataMapWorkflowJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				workflowJobTemplateSchedule := &awx.Schedule{}
				mockAWX.On("CreateWorkflowJobTemplateSchedule",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						workflowJobTemplateSchedule.ID = 2
					}).
					Return(workflowJobTemplateSchedule, nil)
				mockAWX.On("GetScheduleByID", mock.Anything, mock.Anything).Return(workflowJobTemplateSchedule, nil)
			},
			id: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceWorkflowJobTemplateScheduleCreate)
		})
	}
}
