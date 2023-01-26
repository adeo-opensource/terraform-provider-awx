package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapCredentialInput = map[string]interface{}{
	"name":             "foo",
	"input_field_name": "input_field",
	"target":           3,
	"source":           42,
	"metadata":         map[string]interface{}{"key": "value"},
	"description":      "data",
}

func (mockAwx MockAWX) GetCredentialInputSourceByID(id int, params map[string]string) (*awx.CredentialInputSource, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.CredentialInputSource), args.Error(1)
}

func (mockAwx MockAWX) UpdateCredentialInputSourceByID(id int, data map[string]interface{}, params map[string]string) (*awx.CredentialInputSource, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.CredentialInputSource), args.Error(1)
}

func (mockAwx MockAWX) CreateCredentialInputSource(data map[string]interface{}, params map[string]string) (*awx.CredentialInputSource, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.CredentialInputSource), args.Error(1)
}

func (mockAwx MockAWX) DeleteCredentialInputSourceByID(id int, params map[string]string) error {
	args := mockAwx.Called(id, params)
	return args.Error(0)
}
