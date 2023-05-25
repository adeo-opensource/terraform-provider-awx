package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
)

func (mockAwx MockAWX) GetJob(id int, params map[string]string) (*awx.Job, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Job), args.Error(1)
}

func (mockAwx MockAWX) CancelJob(id int, data map[string]interface{}, params map[string]string) (*awx.CancelJobResponse, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.CancelJobResponse), args.Error(1)
}

func (mockAwx MockAWX) RelaunchJob(id int, data map[string]interface{}, params map[string]string) (*awx.JobLaunch, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobLaunch), args.Error(1)
}

func (mockAwx MockAWX) GetHostSummaries(id int, params map[string]string) ([]awx.HostSummary, *awx.HostSummariesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]awx.HostSummary), args.Get(0).(*awx.HostSummariesResponse), args.Error(2)
}

func (mockAwx MockAWX) GetJobEvents(id int, params map[string]string) ([]awx.JobEvent, *awx.JobEventsResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]awx.JobEvent), args.Get(1).(*awx.JobEventsResponse), args.Error(2)
}
