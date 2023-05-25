package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
)

var resourceDataMapOrganization = map[string]interface{}{
	"name":            "foo",
	"organization_id": 1,
	"id":              1,
	"project_id":      1,
}

func (mockAwx MockAWX) GetOrganizationsByID(id int, params map[string]string) (*awx.Organization, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Organization), args.Error(1)
}

func (mockAwx MockAWX) ListOrganizations(params map[string]string) ([]*awx.Organization, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Organization), args.Error(1)
}
func (mockAwx MockAWX) CreateOrganization(data map[string]interface{}, params map[string]string) (*awx.Organization, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Organization), args.Error(1)
}
func (mockAwx MockAWX) UpdateOrganization(id int, data map[string]interface{}, params map[string]string) (*awx.Organization, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Organization), args.Error(1)
}
func (mockAwx MockAWX) DeleteOrganization(id int) (*awx.Organization, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Organization), args.Error(1)
}
