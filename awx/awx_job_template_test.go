package awx

import awx "github.com/adeo-opensource/goawx/client"

var resourceDataMapJobTemplate = map[string]interface{}{
	"job_template_id":          4,
	"name":                     "foo",
	"id":                       1,
	"project_id":               10293,
	"organization_id":          1,
	"notification_template_id": 1,
}

func (mockAwx MockAWX) ListJobTemplates(params map[string]string) ([]*awx.JobTemplate, *awx.ListJobTemplatesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.JobTemplate), args.Get(1).(*awx.ListJobTemplatesResponse), args.Error(2)
}

func (mockAwx MockAWX) GetJobTemplateByID(id int, params map[string]string) (*awx.JobTemplate, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}

func (mockAwx MockAWX) LaunchJob(id int, data map[string]interface{}, params map[string]string) (*awx.JobLaunch, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobLaunch), args.Error(1)
}

func (mockAwx MockAWX) CreateJobTemplate(data map[string]interface{}, params map[string]string) (*awx.JobTemplate, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}

func (mockAwx MockAWX) UpdateJobTemplate(id int, data map[string]interface{}, params map[string]string) (*awx.JobTemplate, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}
func (mockAwx MockAWX) DeleteJobTemplate(id int) (*awx.JobTemplate, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}
func (mockAwx MockAWX) DisAssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*awx.JobTemplate, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}
func (mockAwx MockAWX) AssociateCredentials(id int, data map[string]interface{}, params map[string]string) (*awx.JobTemplate, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobTemplate), args.Error(1)
}
