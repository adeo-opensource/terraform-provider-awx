package awx

import awx "github.com/denouche/goawx/client"

var resourceDataMapInventorySource = map[string]interface{}{
	"name":                 "foo",
	"description":          "data",
	"enabled_var":          "yes",
	"enabled_value":        "maybe",
	"overwrite":            true,
	"overwrite_vars":       true,
	"update_on_launch":     true,
	"inventory_id":         1,
	"credential_id":        1,
	"source":               "louise",
	"source_vars":          "amand",
	"host_filter":          "*.localhost",
	"update_cache_timeout": 10,
	"verbosity":            3,
}

func (mockAwx MockAWX) GetInventorySourceByID(id int, params map[string]string) (*awx.InventorySource, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.InventorySource), args.Error(1)
}
func (mockAwx MockAWX) CreateInventorySource(data map[string]interface{}, params map[string]string) (*awx.InventorySource, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.InventorySource), args.Error(1)
}
func (mockAwx MockAWX) UpdateInventorySource(id int, data map[string]interface{}, params map[string]string) (*awx.InventorySource, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.InventorySource), args.Error(1)
}
func (mockAwx MockAWX) DeleteInventorySource(id int) (*awx.InventorySource, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.InventorySource), args.Error(1)
}
