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

func Test_resourceSettingsLDAPTeamMapCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Can't find settings",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: failed to fetch settings",
				Detail:   "Failed to fetch any ldap setting, got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{}
				mockAWX.On("GetSettingsBySlug",
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "SettingsLDAPTeamMap created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"fii\": {}}")}
				settingsLDAPTeamMapAfter := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {\"organization\": \"3\"}}")}
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMapAfter, nil)
			},
			newData: map[string]interface{}{
				"organization": "3",
			},
		},
		{
			name: "SettingsLDAPTeamMap can't be created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"fii\": {}}")}
				settingsLDAPTeamMapAfter := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {\"organization\": \"3\"}}")}
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMapAfter, nil)
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: team map not created",
				Detail:   "failed to save team map data, got: nothing",
			}},
		},
		{
			name: "team already exist",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {}}")}
				settingsLDAPTeamMapAfter := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {\"organization\": \"3\"}}")}
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMapAfter, nil)
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: team map already exists",
				Detail:   "Map for ldap to team map  already exists",
			}},
		},
		{
			name: "Unparsable auth ldap team map",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: failed to parse AUTH_LDAP_TEAM_MAP setting",
				Detail:   "Failed to parse AUTH_LDAP_TEAM_MAP setting, got: invalid character ']' after top-level value with input []]",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("[]]")}
				settingsLDAPTeamMapAfter := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {\"organization\": \"3\"}}")}
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMapAfter, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceSettingsLDAPTeamMapCreate)
		})
	}
}

func Test_resourceSettingsLDAPTeamMapDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "No settings found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Delete: Unable to fetch settings",
				Detail:   "Unable to load settings with slug ldap: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"fii\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Auth ldap team map unparsable",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Delete: failed to parse AUTH_LDAP_TEAM_MAP setting",
				Detail:   "Failed to parse AUTH_LDAP_TEAM_MAP setting, got: unexpected end of JSON input",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"fii\": {}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "SettingsLDAPTeamMap not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Delete: team map not updated",
				Detail:   "failed to save team map data, got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"fii\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "SettingsLDAPTeamMap deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"1\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings",
					"ldap",
					mock.MatchedBy(func(data map[string]interface{}) bool {
						if data["AUTH_LDAP_TEAM_MAP"] != nil {
							teammap, _ := data["AUTH_LDAP_TEAM_MAP"].(teammap)
							for _, _ = range teammap {
								return false
							}
							return true
						}
						return false
					}),
					mock.Anything).
					Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings",
					"ldap",
					mock.Anything,
					mock.Anything).
					Return(settingsLDAPTeamMap, fmt.Errorf("should not match"))
				mockAWX.On("DeleteSettingsLDAPTeamMap", mock.Anything, mock.Anything).Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceSettingsLDAPTeamMapDelete)
		})
	}
}

func Test_resourceSettingsLDAPTeamMapRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Unable to fetch settings",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch settings",
				Detail:   "Unable to load settings with slug ldap: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"1\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Auth ldap team map not parsable",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"1\"}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to parse AUTH_LDAP_TEAM_MAP",
				Detail:   "Unable to parse AUTH_LDAP_TEAM_MAP, got: invalid character '}' after object key",
			}},
		},
		{
			name: "Ldap team not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch ldap team map",
				Detail:   "Unable to load ldap team map 1: not found",
			}},
		},
		{
			name: "SettingsLDAPTeamMap found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"1\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
			},
			newData: map[string]interface{}{
				"name": "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceSettingsLDAPTeamMapRead)
		})
	}
}

func Test_resourceSettingsLDAPTeamMapUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Settings not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: Unable to fetch settings",
				Detail:   "Unable to load settings with slug ldap: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Auth ldap team map not parsable",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: failed to parse AUTH_LDAP_TEAM_MAP setting",
				Detail:   "Failed to parse AUTH_LDAP_TEAM_MAP setting, got: invalid character '}' after object key",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\"}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)

			},
		},
		{
			name: "SettingsLDAPTeamMap cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: team map not created",
				Detail:   "failed to save team map data, got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Setting{}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "SettingsLDAPTeamMap updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceSettingsLDAPTeamMap().Schema, resourceDataMapSettingsLDAPTeamMap),
			},
			mock: func(mockAWX *MockAWX) {
				settingsLDAPTeamMap := &awx.Setting{"AUTH_LDAP_TEAM_MAP": []byte("{\"foo\": {}}")}
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)
				mockAWX.On("UpdateSettings", mock.Anything, mock.Anything, mock.Anything).Return(settingsLDAPTeamMap, nil)
				mockAWX.On("GetSettingsBySlug", mock.Anything, mock.Anything).Once().Return(settingsLDAPTeamMap, nil)

			},
			newData: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceSettingsLDAPTeamMapUpdate)
		})
	}
}
