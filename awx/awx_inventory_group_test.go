package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapInventoryGroup = map[string]interface{}{
	"name":         "foo",
	"inventory_id": 1,
	"description":  "data",
	"variables":    "toto:toto",
}

func (mockAwx MockAWX) ListInventoryGroups(id int, params map[string]string) ([]*awx.Group, *awx.ListGroupsResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.Group), args.Get(1).(*awx.ListGroupsResponse), args.Error(2)
}
func (mockAwx MockAWX) CreateGroup(data map[string]interface{}, params map[string]string) (*awx.Group, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Group), args.Error(1)
}
func (mockAwx MockAWX) UpdateGroup(id int, data map[string]interface{}, params map[string]string) (*awx.Group, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Group), args.Error(1)
}
func (mockAwx MockAWX) DeleteGroup(id int) (*awx.Group, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Group), args.Error(1)
}
func (mockAwx MockAWX) GetGroupByID(id int, params map[string]string) (*awx.Group, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Group), args.Error(1)
}
