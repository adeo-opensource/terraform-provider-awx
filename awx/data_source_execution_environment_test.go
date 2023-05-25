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

func Test_dataSourceExecutionEnvironmentsRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Error on list execution environment",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceExecutionEnvironment().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch execution environment",
				Detail:   "Fail to find the execution environment got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockEEService := mockAWX.ExecutionEnvironmentService.(mockGeneric[awx.ExecutionEnvironment])
				mockEEService.On("List", mock.Anything).Return([]*awx.ExecutionEnvironment{}, &awx.ListExecutionEnvironmentsResponse{}, fmt.Errorf("nothing"))
				mockAWX.ExecutionEnvironmentService = mockEEService
			},
		},

		{
			name: "Two execution environments",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one element",
				Detail:   "The query returns more than one execution environment, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockEEService := mockAWX.ExecutionEnvironmentService.(mockGeneric[awx.ExecutionEnvironment])
				mockEEService.On("List", mock.Anything).Return([]*awx.ExecutionEnvironment{
					{
						ID:           1,
						Name:         "ExecutionEnvironment1",
						Image:        "open/awx-ee:latest",
						Description:  "An execution environment",
						Organization: 1,
						Credential:   2,
					},
					{
						ID:           2,
						Name:         "ExecutionEnvironment2",
						Image:        "open/awx-ee:latest",
						Description:  "An execution environment",
						Organization: 1,
						Credential:   2,
					},
				}, &awx.ListExecutionEnvironmentsResponse{}, nil)
				mockAWX.ExecutionEnvironmentService = mockEEService
			},
			newData: nil,
		},
		//TODO
		{
			name: "One execution environment",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockEEService := mockAWX.ExecutionEnvironmentService.(mockGeneric[awx.ExecutionEnvironment])
				mockEEService.On("List", mock.Anything).Return([]*awx.ExecutionEnvironment{{
					ID:           3,
					Name:         "ExecutionEnvironment",
					Image:        "open/awx-ee:latest",
					Description:  "An execution environment",
					Organization: 1,
					Credential:   2,
				}}, &awx.ListExecutionEnvironmentsResponse{}, nil)
				mockAWX.ExecutionEnvironmentService = mockEEService
			},
			id: "3",
			newData: map[string]interface{}{
				"name": "ExecutionEnvironment",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceExecutionEnvironmentsRead)

		})
	}
}
