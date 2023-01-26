package awx

import (
	"context"
	"fmt"
	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceOrganizationsRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "No organization",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizations().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch organizations",
				Detail:   "Unable to fetch organizations from AWX API",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListOrganizations", mock.Anything).Return([]*awx.Organization{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "One organization",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceOrganizations().Schema, resourceDataMapMissingId),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListOrganizations", mock.Anything).Return([]*awx.Organization{{ID: 1, Name: "orga"}}, nil)
			},
			newData: map[string]interface{}{
				"organizations": []interface{}{
					map[string]interface{}{
						"id":   1,
						"name": "orga",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceOrganizationsRead)
		})
	}
}
