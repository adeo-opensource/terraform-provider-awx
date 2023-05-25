package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceInventoryRoleRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Missing parameter",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "inventory_id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Inventory role not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Inventory",
				Detail:   "Fail to find the inventory, got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Inventory role found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapInventory),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{
					Name:          "inventRole",
					ID:            1,
					SummaryFields: &awx.Summary{ObjectRoles: &awx.ObjectRoles{AdhocRole: &awx.ApplyRole{ID: 1, Description: "ApplyRole", Name: "ApplyRole"}}},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "ApplyRole",
			},
		},
		{
			name: "Second Organization role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 2, Description: "ApplyRole", Name: "ApplyRole"},
							ReadRole:  &awx.ApplyRole{ID: 1, Description: "ReadRole", Name: "ReadRole"},
						},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "ReadRole",
			},
		},
		{
			name: " Organization role by name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapInventory),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 2, Description: "ApplyRole", Name: "ApplyRole"},
							ReadRole:  &awx.ApplyRole{ID: 9, Description: "ReadRole", Name: "ReadRole"},
							AdminRole: &awx.ApplyRole{ID: 4, Description: "foo", Name: "foo"},
						},
					},
				}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo",
			},
		},
		{
			name: " Organization role not matching",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryRole().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Failed to fetch inventory role - Not Found",
				Detail:   "The project role was not found",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetInventoryByID", mock.Anything, mock.Anything).Return(&awx.Inventory{
					Description: "an organization",
					SummaryFields: &awx.Summary{
						ObjectRoles: &awx.ObjectRoles{
							AdhocRole: &awx.ApplyRole{ID: 2, Description: "ApplyRole", Name: "ApplyRole"},
							ReadRole:  &awx.ApplyRole{ID: 9, Description: "ReadRole", Name: "ReadRole"},
							AdminRole: &awx.ApplyRole{ID: 4, Description: "AdminRole", Name: "AdminRole"},
						},
					},
				}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceInventoryRoleRead)
		})
	}
}
