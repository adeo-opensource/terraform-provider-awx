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

func Test_resourceScheduleCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Schedule not created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Schedule",
				Detail:   "Schedule failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				Schedule := &awx.Schedule{}
				mockAWX.On("CreateSchedule",
					mock.Anything,
					mock.Anything).
					Return(Schedule, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Schedule created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			mock: func(mockAWX *MockAWX) {
				Schedule := &awx.Schedule{}
				mockAWX.On("CreateSchedule",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						Schedule.ID = 2
						Schedule.Description = data["description"].(string) + "_created"
						Schedule.Name = data["name"].(string) + "_created"
					}).
					Return(Schedule, nil)
				mockAWX.On("GetScheduleByID", mock.Anything, mock.Anything).Return(Schedule, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceScheduleCreate)
		})
	}
}

func Test_resourceScheduleDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Schedule not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Schedule delete failed",
				Detail:   "Fail to delete Schedule, id 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteSchedule", mock.Anything, mock.Anything).Return(&awx.Schedule{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Schedule deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteSchedule", mock.Anything, mock.Anything).Return(&awx.Schedule{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceScheduleDelete)
		})
	}
}

func Test_resourceScheduleRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Schedule not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch schedule",
				Detail:   "Unable to load schedule with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				Schedule := &awx.Schedule{}
				mockAWX.On("GetScheduleByID",
					mock.Anything,
					mock.Anything).
					Return(Schedule, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Schedule found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			mock: func(mockAWX *MockAWX) {
				Schedule := &awx.Schedule{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetScheduleByID",
					mock.Anything,
					mock.Anything).
					Return(Schedule, nil)
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
			runTestCase(t, tt, resourceScheduleRead)
		})
	}
}

func Test_resourceScheduleUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Schedule not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch schedule",
				Detail:   "Unable to load schedule with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetScheduleByID", mock.Anything, mock.Anything).Return(&awx.Schedule{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "Schedule cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Schedule",
				Detail:   "Schedule with name foo failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetScheduleByID", mock.Anything, mock.Anything).Return(&awx.Schedule{}, nil)
				mockAWX.On("UpdateSchedule", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Schedule{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Schedule updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSchedule().Schema, resourceDataMapSchedule),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetScheduleByID", mock.Anything, mock.Anything).Return(&awx.Schedule{}, nil)
				mockAWX.On("UpdateSchedule", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Schedule{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceScheduleUpdate)
		})
	}
}
