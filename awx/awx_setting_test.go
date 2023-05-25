package awx

import awx "github.com/adeo-opensource/goawx/client"

var resourceDataMapSetting = map[string]interface{}{
	"name":  "foo",
	"value": "[\"1\"]",
}
var resourceDataMapSettingMapValue = map[string]interface{}{
	"name":  "foo",
	"value": "{\"toto\": \"foo\"}",
}
var resourceDataMapSettingClassicValue = map[string]interface{}{
	"name":  "foo",
	"value": "1",
}
var resourceDataMapSettingsLDAPTeamMap = map[string]interface{}{
	"name":  "foo",
	"value": "1",
}

func (mockAwx MockAWX) ListSettings(params map[string]string) ([]*awx.SettingSummary, *awx.ListSettingsResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.SettingSummary), args.Get(1).(*awx.ListSettingsResponse), args.Error(2)
}
func (mockAwx MockAWX) GetSettingsBySlug(slug string, params map[string]string) (*awx.Setting, error) {
	args := mockAwx.Called(slug, params)
	return args.Get(0).(*awx.Setting), args.Error(1)
}
func (mockAwx MockAWX) UpdateSettings(slug string, data map[string]interface{}, params map[string]string) (*awx.Setting, error) {
	args := mockAwx.Called(slug, data, params)
	return args.Get(0).(*awx.Setting), args.Error(1)
}
func (mockAwx MockAWX) DeleteSettings(slug string) (*awx.Setting, error) {
	args := mockAwx.Called(slug)
	return args.Get(0).(*awx.Setting), args.Error(1)
}
