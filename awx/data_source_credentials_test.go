package awx

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	awx "github.com/denouche/goawx/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_dataSourceCredentialsRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Error on list credentials",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentials().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch credentials",
				Detail:   "Unable to fetch credentials from AWX API",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListCredentials", mock.Anything).Return([]*awx.Credential{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "One credential",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceCredentials().Schema, resourceDataMapCredential),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListCredentials", mock.Anything).Return([]*awx.Credential{{ID: 1, Kind: "toto", Inputs: map[string]interface{}{"username": "borto"}}}, nil)
			},
			newData: map[string]interface{}{
				"credentials": []interface{}{
					map[string]interface{}{
						"id":       1,
						"kind":     "toto",
						"username": "borto",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceCredentialsRead)
		})
	}
}