package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapCredential = map[string]interface{}{
	"credential_id":       1,
	"secret":              "terces",
	"id":                  1,
	"tenant":              "awx",
	"url":                 "http://localhost",
	"description":         "data",
	"client":              "terraform",
	"organization_id":     1,
	"name":                "foo",
	"username":            "user",
	"password":            "pass",
	"project":             "proj",
	"ssh_key_data":        "a ssh key",
	"ssh_public_key_data": "a public ssh key",
	"ssh_key_unlock":      "alomora",
	"become_method":       "su",
	"become_username":     "root",
	"become_password":     "toor",
	"inputs":              "{ \"data\":\"data\"}",
}

func (mockAwx MockAWX) GetCredentialsByID(id int, params map[string]string) (*awx.Credential, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Credential), args.Error(1)
}

func (mockAwx MockAWX) ListCredentials(params map[string]string) ([]*awx.Credential, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Credential), args.Error(1)
}

func (mockAwx MockAWX) CreateCredentials(data map[string]interface{}, params map[string]string) (*awx.Credential, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Credential), args.Error(1)
}

func (mockAwx MockAWX) UpdateCredentialsByID(id int, data map[string]interface{}, params map[string]string) (*awx.Credential, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Credential), args.Error(1)
}
