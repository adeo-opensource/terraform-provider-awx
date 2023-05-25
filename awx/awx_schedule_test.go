package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
)

var resourceDataMapSchedule = map[string]interface{}{
	"name": "foo",
}

func (mockAwx MockAWX) ListSchedule(params map[string]string) ([]*awx.Schedule, *awx.ListSchedulesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Schedule), args.Get(1).(*awx.ListSchedulesResponse), args.Error(2)
}
func (mockAwx MockAWX) GetScheduleByID(id int, params map[string]string) (*awx.Schedule, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Schedule), args.Error(1)
}
func (mockAwx MockAWX) CreateSchedule(data map[string]interface{}, params map[string]string) (*awx.Schedule, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Schedule), args.Error(1)
}
func (mockAwx MockAWX) UpdateSchedule(id int, data map[string]interface{}, params map[string]string) (*awx.Schedule, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Schedule), args.Error(1)
}
func (mockAwx MockAWX) DeleteSchedule(id int) (*awx.Schedule, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Schedule), args.Error(1)
}
