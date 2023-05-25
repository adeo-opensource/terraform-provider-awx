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

func Test_dataSourceNotificationTemplatesRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Error on list notification template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceNotificationTemplate().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch NotificationTemplate",
				Detail:   "Fail to find the notification template got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListNotificationTemplates", mock.Anything).Return([]*awx.NotificationTemplate{}, &awx.ListNotificationTemplatesResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Two notification template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one NotificationTemplate, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListNotificationTemplates", mock.Anything).Return([]*awx.NotificationTemplate{{}, {}}, &awx.ListNotificationTemplatesResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One notification template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListNotificationTemplates", mock.Anything).Return([]*awx.NotificationTemplate{{
					Description: "a notification template",
				}}, &awx.ListNotificationTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				//TODO: update schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceNotificationTemplatesRead)
		})
	}
}
