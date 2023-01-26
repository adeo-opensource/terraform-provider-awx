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

func Test_dataSourceSchedulesRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Error on list schedule",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceSchedule().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Schedule Group",
				Detail:   "Fail to find the group got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListSchedule", mock.Anything).Return([]*awx.Schedule{}, &awx.ListSchedulesResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Two schedule",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Group, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListSchedule", mock.Anything).Return([]*awx.Schedule{{}, {}}, &awx.ListSchedulesResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One schedule",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceSchedule().Schema, resourceDataMapSchedule),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListSchedule", mock.Anything).Return([]*awx.Schedule{{
					Description: "a schedule",
					Name:        "the schedule",
				}}, &awx.ListSchedulesResponse{}, nil)
			},
			newData: map[string]interface{}{
				//TODO: schema
				"name": "the schedule",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceSchedulesRead)

		})
	}
}
