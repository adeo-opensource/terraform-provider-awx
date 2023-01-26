package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapExecutionEnvironment = map[string]interface{}{
	"name":         "foo",
	"organization": 1,
	"credential":   1,
	"image":        "dockerhub.io/small:tag",
	"description":  "data",
}

func (mockAwx MockAWX) ListExecutionEnvironments(params map[string]string) ([]*awx.ExecutionEnvironment, *awx.ListExecutionEnvironmentsResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.ExecutionEnvironment), args.Get(1).(*awx.ListExecutionEnvironmentsResponse), args.Error(2)
}

func (mockAwx MockAWX) GetExecutionEnvironmentByID(id int, params map[string]string) (*awx.ExecutionEnvironment, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.ExecutionEnvironment), args.Error(1)
}

func (mockAwx MockAWX) CreateExecutionEnvironment(data map[string]interface{}, params map[string]string) (*awx.ExecutionEnvironment, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.ExecutionEnvironment), args.Error(1)
}

func (mockAwx MockAWX) UpdateExecutionEnvironment(id int, data map[string]interface{}, params map[string]string) (*awx.ExecutionEnvironment, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.ExecutionEnvironment), args.Error(1)
}

func (mockAwx MockAWX) DeleteExecutionEnvironment(id int) (*awx.ExecutionEnvironment, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.ExecutionEnvironment), args.Error(1)
}
