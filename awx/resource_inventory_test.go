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

func Test_resourceInventoryCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Inventory",
				Detail:   "Unable to create Inventory got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{}
				mockAWX.On("CreateInventory",
					mock.Anything,
					mock.Anything).
					Return(inventory, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Inventory created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{}
				mockAWX.On("CreateInventory",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						inventory.ID = 2
						inventory.Description = data["description"].(string) + "_created"
						inventory.Organization = data["organization"].(int)
						inventory.Name = data["name"].(string) + "_created"
						inventory.Kind = data["kind"].(string) + "_created"
						inventory.HostFilter = data["host_filter"].(string) + "_created"
						inventory.Variables = data["variables"].(string) + "_created"

					}).
					Return(inventory, nil)
				mockAWX.On("GetInventory", mock.Anything, mock.Anything).Return(inventory, nil)
			},
			newData: map[string]interface{}{
				"description":     "data_created",
				"name":            "foo_created",
				"organization_id": 1,
				"kind":            "inv_created",
				"host_filter":     "*.localhost_created",
				"variables":       "toto:toto_created\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceInventoryCreate)
		})
	}
}

func Test_resourceInventoryDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Inventory delete failed",
				Detail:   "Fail to delete Inventory, Inventory 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteInventory", mock.Anything).Return(&awx.Inventory{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Inventory deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteInventory", mock.Anything).Return(&awx.Inventory{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceInventoryDelete)
		})
	}
}

func Test_resourceInventoryRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Inventory",
				Detail:   "Unable to load Inventory with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{}
				mockAWX.On("GetInventory",
					mock.Anything,
					mock.Anything).
					Return(inventory, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Inventory found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{
					Name:         "toto",
					Description:  "data_read",
					Organization: 1,
				}
				mockAWX.On("GetInventory",
					mock.Anything,
					mock.Anything).
					Return(inventory, nil)
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
			runTestCase(t, tt, resourceInventoryRead)
		})
	}
}

func Test_resourceInventoryUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Inventory",
				Detail:   "Unable to update Inventory with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{}
				mockAWX.On("UpdateInventory",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(inventory, fmt.Errorf("nothing"))
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(inventory, nil)

			},
		},
		{
			name: "Inventory updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventory().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				inventory := &awx.Inventory{}
				mockAWX.On("UpdateInventory",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						inventory.ID = 2
						inventory.Description = data["description"].(string) + "_updated"
						inventory.Name = data["name"].(string) + "_updated"
						inventory.Organization = data["organization"].(int)
					}).
					Return(inventory, nil)
				mockAWX.On("GetInventory", mock.Anything, mock.Anything).Return(inventory, nil)

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
			runTestCase(t, tt, resourceInventoryUpdate)
		})
	}
}
