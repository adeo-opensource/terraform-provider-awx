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

func Test_dataSourceInventoriesRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Error on list inventory",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventory().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Inventory Group",
				Detail:   "Fail to find the group got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListInventories", mock.Anything).Return([]*awx.Inventory{}, &awx.ListInventoriesResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Two inventory",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventory().Schema, resourceDataMapInventory),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Group, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListInventories", mock.Anything).Return([]*awx.Inventory{{}, {}}, &awx.ListInventoriesResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One inventory",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventory().Schema, resourceDataMapInventory),
			},
			want: nil,
			id:   "4",
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListInventories", mock.Anything).Return([]*awx.Inventory{{
					Kind:        "invent",
					HostFilter:  "filter",
					Description: "an inventory",
					Variables:   "toto:toto",
					ID:          4,
				}}, &awx.ListInventoriesResponse{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceInventoriesRead)
		})
	}
}
