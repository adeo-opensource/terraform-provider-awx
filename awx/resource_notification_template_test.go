package awx

import (
	"context"
	"fmt"
	awx "github.com/denouche/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceNotificationTemplateCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "NotificationTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create NotificationTemplate",
				Detail:   "NotificationTemplate failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				notificationTemplate := &awx.NotificationTemplate{}
				mockAWX.On("CreateNotificationTemplate",
					mock.Anything,
					mock.Anything).
					Return(notificationTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "NotificationTemplate created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				notificationTemplate := &awx.NotificationTemplate{}
				mockAWX.On("CreateNotificationTemplate",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						notificationTemplate.ID = 2
						notificationTemplate.Description = data["description"].(string) + "_created"
						notificationTemplate.Name = data["name"].(string) + "_created"
					}).
					Return(notificationTemplate, nil)
				mockAWX.On("GetNotificationTemplateByID", mock.Anything, mock.Anything).Return(notificationTemplate, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceNotificationTemplateCreate)
		})
	}
}

func Test_resourceNotificationTemplateDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "NotificationTemplate not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Notification Template delete failed",
				Detail:   "Fail to delete Notification Template, id 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteNotificationTemplate", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "NotificationTemplate deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteNotificationTemplate", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceNotificationTemplateDelete)
		})
	}
}

func Test_resourceNotificationTemplateRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Notification Template not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch notification_template",
				Detail:   "Unable to load notification_template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				notificationTemplate := &awx.NotificationTemplate{}
				mockAWX.On("GetNotificationTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(notificationTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "NotificationTemplate found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				notificationTemplate := &awx.NotificationTemplate{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetNotificationTemplateByID",
					mock.Anything,
					mock.Anything).
					Return(notificationTemplate, nil)
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
			runTestCase(t, tt, resourceNotificationTemplateRead)
		})
	}
}

func Test_resourceNotificationTemplateUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Notification Template not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch notification_template",
				Detail:   "Unable to load notification_template with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetNotificationTemplateByID", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "Notification Template cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to update NotificationTemplate",
				Detail:   "notification_template with name foo failed to update nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetNotificationTemplateByID", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, nil)
				mockAWX.On("UpdateNotificationTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Notification Template updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceNotificationTemplate().Schema, resourceDataMapNotificationTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetNotificationTemplateByID", mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, nil)
				mockAWX.On("UpdateNotificationTemplate", mock.Anything, mock.Anything, mock.Anything).Return(&awx.NotificationTemplate{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceNotificationTemplateUpdate)
		})
	}
}
