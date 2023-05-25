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

func Test_dataSourceInventoryGroupRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Missing parameter inventory_id",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryGroup().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Missing Parameters",
				Detail:   "inventory_id parameter is required.",
			}},
			mock: func(mockAWX *MockAWX) {
				mockGroupService := mockAWX.GroupService.(mockGeneric[awx.Group])
				mockGroupService.On("List", mock.Anything).Return([]*awx.Group{}, nil, fmt.Errorf("nothing"))
				mockAWX.GroupService = mockGroupService
			},
		},
		{
			name: "No inventory group",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Inventory Group",
				Detail:   "Fail to find the group got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListInventoryGroups", mock.Anything, mock.Anything).Return([]*awx.Group{}, &awx.ListGroupsResponse{}, fmt.Errorf("nothing"))
			},
			newData: nil,
		},
		{
			name: "Two inventory group",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Group, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockGroupService := mockAWX.GroupService.(mockGeneric[awx.Group])
				mockGroupService.On("List", mock.Anything).Return([]*awx.Group{{}, {}}, &awx.ListGroupsResponse{}, nil)
				mockAWX.GroupService = mockGroupService
			},
			newData: nil,
		},
		{
			name: "One Inventory groupe",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceInventoryGroup().Schema, resourceDataMapInventoryGroup),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListInventoryGroups", mock.Anything, mock.Anything).Return([]*awx.Group{{
					Name:        "inv",
					ID:          1,
					Inventory:   2,
					Description: "an inventory",
					Variables:   "toto:toto",
				}}, &awx.ListGroupsResponse{}, nil)
			},
			newData: map[string]interface{}{
				//TODO: Schema update
				//"description":  "an inventory",
				//"variables":    "toto:toto\n",
				"inventory_id": 2,
				"name":         "inv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceInventoryGroupRead)
		})
	}
}
