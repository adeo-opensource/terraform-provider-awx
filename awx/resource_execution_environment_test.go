package awx

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/mock"
	"testing"

	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceExecutionEnvironmentsCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create ExecutionEnvironments",
				Detail:   "ExecutionEnvironments with name, failed to create foo",
			}},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{}
				mockAWX.On("CreateExecutionEnvironment",
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{}
				mockAWX.On("CreateExecutionEnvironment",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						executionEnvironment.ID = 2
						executionEnvironment.Description = data["description"].(string) + "_created"
						executionEnvironment.Image = data["image"].(string) + "_created"
						executionEnvironment.Organization = *(data["organization"].(*int))
						executionEnvironment.Credential = *(data["credential"].(*int))
					}).
					Return(executionEnvironment, nil)
				mockAWX.On("GetExecutionEnvironmentByID", mock.Anything, mock.Anything).Return(executionEnvironment, nil)

			},
			newData: map[string]interface{}{
				"description": "data_created",
				"image":       "dockerhub.io/small:tag_created",
				//TODO: schema
				//"organization": 1,
				//"credential":   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceExecutionEnvironmentsCreate)
		})
	}
}

func Test_resourceExecutionEnvironmentsDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "ExecutionEnvironment delete failed",
				Detail:   "Fail to delete ExecutionEnvironment, ExecutionEnvironmentID 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteExecutionEnvironment", mock.Anything).Return(&awx.ExecutionEnvironment{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteExecutionEnvironment", mock.Anything).Return(&awx.ExecutionEnvironment{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceExecutionEnvironmentsDelete)
		})
	}
}

func Test_resourceExecutionEnvironmentsRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch ExecutionEnvironment",
				Detail:   "Unable to load ExecutionEnvironment with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{}
				mockAWX.On("GetExecutionEnvironmentByID",
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{
					Name:         "toto",
					Image:        "dockerhub.io/execution",
					Description:  "data_read",
					Organization: 1,
					Credential:   1,
				}
				mockAWX.On("GetExecutionEnvironmentByID",
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, nil)
			},
			newData: map[string]interface{}{
				"description": "data_read",
				"image":       "dockerhub.io/execution",
				"name":        "toto",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceExecutionEnvironmentsRead)
		})
	}

}

func Test_resourceExecutionEnvironmentsUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update ExecutionEnvironments",
				Detail:   "ExecutionEnvironments with name foo failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{}
				mockAWX.On("UpdateExecutionEnvironment",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, fmt.Errorf("nothing"))
				mockAWX.On("GetExecutionEnvironmentByID", mock.Anything, mock.Anything).Return(executionEnvironment, nil)

			},
		},
		{
			name: "Execution environment updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceExecutionEnvironment().Schema, resourceDataMapExecutionEnvironment),
			},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.ExecutionEnvironment{}
				mockAWX.On("UpdateExecutionEnvironment",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						executionEnvironment.ID = 2
						executionEnvironment.Description = data["description"].(string) + "_updated"
						executionEnvironment.Name = data["name"].(string) + "_updated"
						executionEnvironment.Image = data["image"].(string) + "_updated"
						executionEnvironment.Organization = *(data["organization"].(*int))
						executionEnvironment.Credential = *(data["credential"].(*int))
					}).
					Return(executionEnvironment, nil)
				mockAWX.On("GetExecutionEnvironmentByID", mock.Anything, mock.Anything).Return(executionEnvironment, nil)

			},
			newData: map[string]interface{}{
				"description": "data_updated",
				"image":       "dockerhub.io/small:tag_updated",
				"name":        "foo_updated",
				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceExecutionEnvironmentsUpdate)
		})
	}
}
