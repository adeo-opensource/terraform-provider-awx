package awx

import (
	"context"
	"encoding/json"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceSettingRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Setting not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSetting),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch settings",
				Detail:   "Unable to load settings with slug all: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Setting found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSetting),
			},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceSettingRead)
		})
	}
}

func Test_resourceSettingUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Setting not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSetting),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: failed to fetch settings",
				Detail:   "Update to fetch setting, got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Setting cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSetting),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: setting not updated",
				Detail:   "failed to update setting data, got: nothing, [\"1\"]",
			}},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(setting, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Setting updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSetting),
			},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Setting{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Setting updated classic value",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSettingClassicValue),
			},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Setting{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Setting updated map",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSetting().Schema, resourceDataMapSettingMapValue),
			},
			mock: func(mockAWX *MockAWX) {
				setting := &awx.Setting{"toto": json.RawMessage{}}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(setting, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Setting{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceSettingUpdate)
		})
	}
}
