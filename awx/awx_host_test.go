package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapHost = map[string]interface{}{
	"name":         "foo",
	"description":  "data",
	"group_ids":    []interface{}{1},
	"instance_id":  1,
	"variables":    "toto:toto",
	"inventory_id": 1,
}

func (mockAwx MockAWX) GetHostByID(id int, params map[string]string) (*awx.Host, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Host), args.Error(1)
}

func (mockAwx MockAWX) CreateHost(data map[string]interface{}, params map[string]string) (*awx.Host, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Host), args.Error(1)
}

func (mockAwx MockAWX) UpdateHost(id int, data map[string]interface{}, params map[string]string) (*awx.Host, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Host), args.Error(1)
}

func (mockAwx MockAWX) AssociateGroup(id int, data map[string]interface{}, params map[string]string) (*awx.Host, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Host), args.Error(1)
}

func (mockAwx MockAWX) DisAssociateGroup(id int, data map[string]interface{}, params map[string]string) (*awx.Host, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Host), args.Error(1)
}

func (mockAwx MockAWX) DeleteHost(id int) (*awx.Host, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Host), args.Error(1)
}
