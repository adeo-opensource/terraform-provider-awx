package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
)

var resourceDataMapInventory = map[string]interface{}{
	"description":     "data",
	"name":            "foo",
	"inventory_id":    1,
	"id":              1,
	"organization_id": 1,
	"kind":            "inv",
	"host_filter":     "*.localhost",
	"variables":       "toto:toto",
}

func (mockAwx MockAWX) GetInventoryByID(id int, params map[string]string) (*awx.Inventory, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Inventory), args.Error(1)
}

func (mockAwx MockAWX) ListInventories(params map[string]string) ([]*awx.Inventory, *awx.ListInventoriesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Inventory), args.Get(1).(*awx.ListInventoriesResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateInventory(data map[string]interface{}, params map[string]string) (*awx.Inventory, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Inventory), args.Error(1)
}

func (mockAwx MockAWX) UpdateInventory(id int, data map[string]interface{}, params map[string]string) (*awx.Inventory, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Inventory), args.Error(1)
}

func (mockAwx MockAWX) GetInventory(id int, params map[string]string) (*awx.Inventory, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Inventory), args.Error(1)
}

func (mockAwx MockAWX) DeleteInventory(id int) (*awx.Inventory, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Inventory), args.Error(1)
}
