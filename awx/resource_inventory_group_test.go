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

func Test_resourceInventoryGroupCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Group not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Inventory Group",
				Detail:   "Unable to create Inventory Group got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{}
				mockAWX.On("CreateGroup",
					mock.Anything,
					mock.Anything).
					Return(group, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Group created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{}
				mockAWX.On("CreateGroup",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						group.ID = 2
						group.Description = data["description"].(string) + "_created"
						group.Name = data["name"].(string) + "_created"
						group.Variables = data["variables"].(string) + "_created"
						group.Inventory = data["inventory"].(int)
					}).
					Return(group, nil)
				mockAWX.On("GetGroupByID", mock.Anything, mock.Anything).Return(group, nil)
			},
			newData: map[string]interface{}{
				"description":  "data_created",
				"name":         "foo_created",
				"variables":    "toto:toto_created\n",
				"inventory_id": 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceInventoryGroupCreate)
		})
	}
}

func Test_resourceInventoryGroupDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Group not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Inventory Group delete failed",
				Detail:   "Fail to delete Inventory Group, ID: 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteGroup", mock.Anything).Return(&awx.Group{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Group deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteGroup", mock.Anything).Return(&awx.Group{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceInventoryGroupDelete)
		})
	}

}

func Test_resourceInventoryGroupRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Inventory Group",
				Detail:   "Unable to load Inventory Group with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{}
				mockAWX.On("GetGroupByID",
					mock.Anything,
					mock.Anything).
					Return(group, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Group found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetGroupByID",
					mock.Anything,
					mock.Anything).
					Return(group, nil)
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
			runTestCase(t, tt, resourceInventoryGroupRead)
		})
	}
}

func Test_resourceInventoryGroupUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Group not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Inventory Group",
				Detail:   "Unable to update Inventory Group with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{}
				mockAWX.On("UpdateGroup",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(group, fmt.Errorf("nothing"))
				mockAWX.On("GetGroupByID", mock.Anything, mock.Anything).Return(group, nil)

			},
		},
		{
			name: "Group updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			mock: func(mockAWX *MockAWX) {
				group := &awx.Group{}
				mockAWX.On("UpdateGroup",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						group.ID = 2
						group.Description = data["description"].(string) + "_updated"
						group.Name = data["name"].(string) + "_updated"
					}).
					Return(group, nil)
				mockAWX.On("GetGroupByID", mock.Anything, mock.Anything).Return(group, nil)

			},
			newData: map[string]interface{}{
				"description": "data_updated",
				"name":        "foo_updated",
				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceInventoryGroupUpdate)
		})
	}
}
