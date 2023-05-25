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

func Test_dataSourceOrganizationRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "No organization",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganization().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch organization",
				Detail:   "Fail to find the organization got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListOrganizations", mock.Anything).Return([]*awx.Organization{}, fmt.Errorf("nothing"))
			},
		},
		//TODO: schema
		//{
		//	name: "One organization",
		//	args: args{
		//		ctx: context.Background(),
		//		d:   schema.TestResourceDataRaw(t, dataSourceOrganization().Schema, resourceDataMapMissingId),
		//	},
		//	want: nil,
		//	mockawx: func(mockAWX *MockAWX) {
		//		mockAWX.On("ListOrganizations", mockawx.Anything).Return([]*awx.Organization{{Name: "orga"}}, nil)
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceOrganizationRead)
		})
	}
}
