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

func Test_resourceInventorySourceCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "InventorySource not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create Inventory Source",
				Detail:   "Unable to create Inventory Source got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				inventorySource := &awx.InventorySource{}
				mockAWX.On("CreateInventorySource",
					mock.Anything,
					mock.Anything).
					Return(inventorySource, fmt.Errorf("nothing"))
			},
		},
		{
			name: "InventorySource created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			mock: func(mockAWX *MockAWX) {
				inventorySource := &awx.InventorySource{}
				mockAWX.On("CreateInventorySource",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						inventorySource.ID = 2
						inventorySource.Description = data["description"].(string) + "_created"
						inventorySource.Name = data["name"].(string) + "_created"
						inventorySource.Inventory = data["inventory"].(int)
						inventorySource.EnabledVar = data["enabled_var"].(string) + "_created"
						inventorySource.EnabledValue = data["enabled_value"].(string) + "_created"
						inventorySource.Overwrite = data["overwrite"].(bool)
						inventorySource.OverwriteVars = data["overwrite_vars"].(bool)
						inventorySource.UpdateOnLaunch = data["update_on_launch"].(bool)
						inventorySource.Inventory = data["inventory"].(int)
						inventorySource.Credential = data["credential"].(int)
						inventorySource.Source = data["source"].(string) + "_created"
						inventorySource.SourceVars = data["source_vars"].(string) + "_created"
						inventorySource.HostFilter = data["host_filter"].(string) + "_created"
						inventorySource.UpdateCacheTimeout = data["update_cache_timeout"].(int)
						inventorySource.Verbosity = data["verbosity"].(int)

					}).
					Return(inventorySource, nil)
				mockAWX.On("GetInventorySourceByID", mock.Anything, mock.Anything).Return(inventorySource, nil)
			},
			newData: map[string]interface{}{
				"name":                 "foo_created",
				"description":          "data_created",
				"enabled_var":          "yes_created",
				"enabled_value":        "maybe_created",
				"overwrite":            true,
				"overwrite_vars":       true,
				"update_on_launch":     true,
				"inventory_id":         1,
				"credential_id":        1,
				"source":               "louise_created",
				"source_vars":          "amand_created\n",
				"host_filter":          "*.localhost_created",
				"update_cache_timeout": 10,
				"verbosity":            3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceInventorySourceCreate)
		})
	}
}

func Test_resourceInventorySourceDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Source not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Inventory Source delete failed",
				Detail:   "Fail to delete Inventory Source, Inventory Source 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteInventorySource", mock.Anything).Return(&awx.InventorySource{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Source deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteInventorySource", mock.Anything).Return(&awx.InventorySource{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceInventorySourceDelete)
		})
	}
}

func Test_resourceInventorySourceRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Inventory Source not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch Inventory Source",
				Detail:   "Unable to load Inventory Source with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				source := &awx.InventorySource{}
				mockAWX.On("GetInventorySourceByID",
					mock.Anything,
					mock.Anything).
					Return(source, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Source found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			mock: func(mockAWX *MockAWX) {
				source := &awx.InventorySource{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetInventorySourceByID",
					mock.Anything,
					mock.Anything).
					Return(source, nil)
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
			runTestCase(t, tt, resourceInventorySourceRead)
		})
	}
}

func Test_resourceInventorySourceUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Source not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update Inventory Source",
				Detail:   "Unable to update Inventory Source with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				source := &awx.InventorySource{}
				mockAWX.On("UpdateInventorySource",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(source, fmt.Errorf("nothing"))
				mockAWX.On("GetInventorySourceByID", mock.Anything, mock.Anything).Return(source, nil)

			},
		},
		{
			name: "Source updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceInventorySource().Schema, resourceDataMapInventorySource),
			},
			mock: func(mockAWX *MockAWX) {
				source := &awx.InventorySource{}
				mockAWX.On("UpdateInventorySource",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(1).(map[string]interface{})
						source.ID = 2
						source.Description = data["description"].(string) + "_updated"
						source.Name = data["name"].(string) + "_updated"
					}).
					Return(source, nil)
				mockAWX.On("GetInventorySourceByID", mock.Anything, mock.Anything).Return(source, nil)

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
			runTestCase(t, tt, resourceInventorySourceUpdate)
		})
	}
}
