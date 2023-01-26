package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapCredentialType = map[string]interface{}{
	"id":          1,
	"inputs":      "{ \"data\":\"data\"}",
	"injectors":   "{ \"data\":\"data\"}",
	"description": "data",
	"name":        "foo",
}

func (mockAwx MockAWX) GetCredentialTypeByID(id int, params map[string]string) (*awx.CredentialType, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.CredentialType), args.Error(1)
}

func (mockAwx MockAWX) CreateCredentialType(data map[string]interface{}, params map[string]string) (*awx.CredentialType, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.CredentialType), args.Error(1)
}

func (mockAwx MockAWX) UpdateCredentialTypeByID(id int, data map[string]interface{}, params map[string]string) (*awx.CredentialType, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.CredentialType), args.Error(1)
}

func (mockAwx MockAWX) DeleteCredentialTypeByID(id int, params map[string]string) error {
	args := mockAwx.Called(id, params)
	return args.Error(0)
}
