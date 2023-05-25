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

func Test_resourceHostCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Host",
				Detail:   "Unable to create Host got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.Host{}
				mockAWX.On("CreateHost",
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{}
				mockAWX.On("CreateHost",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						host.ID = 2
						host.Description = data["description"].(string) + "_created"
					}).
					Return(host, nil)
				mockAWX.On("GetHostByID", mock.Anything, mock.Anything).Return(host, nil)
				mockAWX.On("AssociateGroup", mock.Anything, mock.Anything, mock.Anything).Return(host, nil)

			},
			newData: map[string]interface{}{
				"description": "data_created",
				//TODO: schema
				//"organization": 1,
				//"credential":   1,
			},
		},
		{
			name: "Execution environment can't associate",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Host",
				Detail:   "Assign Group Id 1 to hostid 2 fail, got  nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{}
				mockAWX.On("CreateHost",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						host.ID = 2
						host.Description = data["description"].(string) + "_created"
					}).
					Return(host, nil)
				mockAWX.On("GetHostByID", mock.Anything, mock.Anything).Return(host, nil)
				mockAWX.On("AssociateGroup", mock.Anything, mock.Anything, mock.Anything).Return(host, fmt.Errorf("nothing"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			tt.args.d.MarkNewResource()
			runTestCase(t, tt, resourceHostCreate)
		})
	}
}

func Test_resourceHostDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Host delete failed",
				Detail:   "Fail to delete Host, id 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteHost", mock.Anything).Return(&awx.Host{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteHost", mock.Anything).Return(&awx.Host{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceHostDelete)
		})
	}
}

func Test_resourceHostRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Host",
				Detail:   "Unable to load Host with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				executionEnvironment := &awx.Host{}
				mockAWX.On("GetHostByID",
					mock.Anything,
					mock.Anything).
					Return(executionEnvironment, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Execution environment found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetHostByID",
					mock.Anything,
					mock.Anything).
					Return(host, nil)
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
			runTestCase(t, tt, resourceHostRead)
		})
	}
}

func Test_resourceHostUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Execution environment not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Host",
				Detail:   "Unable to update Host with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{}
				mockAWX.On("UpdateHost",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(host, fmt.Errorf("nothing"))
				mockAWX.On("GetHostByID", mock.Anything, mock.Anything).Return(host, nil)

			},
		},
		{
			name: "Execution environment updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{}
				mockAWX.On("UpdateHost",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						host.ID = 2
						host.Description = data["description"].(string) + "_updated"
						host.Name = data["name"].(string) + "_updated"
						host.Inventory = data["inventory"].(int)
						host.Enabled = data["enabled"].(bool)
						host.InstanceID = data["instance_id"].(string) + "_updated"
						host.Variables = data["variables"].(string) + "_updated"
					}).
					Return(host, nil)
				mockAWX.On("GetHostByID", mock.Anything, mock.Anything).Return(host, nil)
				mockAWX.On("AssociateGroup", mock.Anything, mock.Anything, mock.Anything).Return(host, nil)
			},
			newData: map[string]interface{}{
				"description":  "data_updated",
				"name":         "foo_updated",
				"inventory_id": 1,
				"enabled":      true,
				"instance_id":  "1_updated",
				"variables":    "toto:toto_updated\n",
				//TODO: schema
			},
		},
		{
			name: "Execution environment can't associate",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceHost().Schema, resourceDataMapHost),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Host",
				Detail:   "Assign Group Id 1 to hostid 1 fail, got  nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				host := &awx.Host{}
				mockAWX.On("UpdateHost",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						host.ID = 2
						host.Description = data["description"].(string) + "_updated"
						host.Name = data["name"].(string) + "_updated"
						host.Inventory = data["inventory"].(int)
						host.Enabled = data["enabled"].(bool)
						host.InstanceID = data["instance_id"].(string) + "_updated"
						host.Variables = data["variables"].(string) + "_updated"
					}).
					Return(host, nil)
				mockAWX.On("GetHostByID", mock.Anything, mock.Anything).Return(host, nil)
				mockAWX.On("AssociateGroup", mock.Anything, mock.Anything, mock.Anything).Return(host, fmt.Errorf("nothing"))

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceHostUpdate)
		})
	}
}
