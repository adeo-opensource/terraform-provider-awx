package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapWorkflowJobTemplate = map[string]interface{}{
	"name":                          "foo",
	"id":                            1,
	"project_id":                    1,
	"workflow_job_template_node_id": 1,
	"description":                   "toto",
}

func (mockAwx MockAWX) ListWorkflowJobTemplates(params map[string]string) ([]*awx.WorkflowJobTemplate, *awx.ListWorkflowJobTemplatesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.WorkflowJobTemplate), args.Get(1).(*awx.ListWorkflowJobTemplatesResponse), args.Error(2)
}

func (mockAwx MockAWX) ListWorkflowJobTemplateSuccessNodeSteps(id int, params map[string]string) ([]*awx.WorkflowJobTemplateNode, *awx.ListWorkflowJobTemplateNodesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.WorkflowJobTemplateNode), args.Get(1).(*awx.ListWorkflowJobTemplateNodesResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateWorkflowJobTemplateSuccessNodeStep(id int, data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) ListWorkflowJobTemplateFailureNodeSteps(id int, params map[string]string) ([]*awx.WorkflowJobTemplateNode, *awx.ListWorkflowJobTemplateNodesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.WorkflowJobTemplateNode), args.Get(1).(*awx.ListWorkflowJobTemplateNodesResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateWorkflowJobTemplateFailureNodeStep(id int, data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) ListWorkflowJobTemplateAlwaysNodeSteps(id int, params map[string]string) ([]*awx.WorkflowJobTemplateNode, *awx.ListWorkflowJobTemplateNodesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.WorkflowJobTemplateNode), args.Get(1).(*awx.ListWorkflowJobTemplateNodesResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateWorkflowJobTemplateAlwaysNodeStep(id int, data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) GetWorkflowJobTemplateNodeByID(id int, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) ListWorkflowJobTemplateNodes(params map[string]string) ([]*awx.WorkflowJobTemplateNode, *awx.ListWorkflowJobTemplateNodesResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.WorkflowJobTemplateNode), args.Get(1).(*awx.ListWorkflowJobTemplateNodesResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateWorkflowJobTemplateNode(data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) UpdateWorkflowJobTemplateNode(id int, data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) DeleteWorkflowJobTemplateNode(id int) (*awx.WorkflowJobTemplateNode, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.WorkflowJobTemplateNode), args.Error(1)
}

func (mockAwx MockAWX) GetWorkflowJobTemplateByID(id int, params map[string]string) (*awx.WorkflowJobTemplate, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.WorkflowJobTemplate), args.Error(1)
}

func (mockAwx MockAWX) CreateWorkflowJobTemplate(data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplate, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.WorkflowJobTemplate), args.Error(1)
}

func (mockAwx MockAWX) UpdateWorkflowJobTemplate(id int, data map[string]interface{}, params map[string]string) (*awx.WorkflowJobTemplate, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.WorkflowJobTemplate), args.Error(1)
}

func (mockAwx MockAWX) DeleteWorkflowJobTemplate(id int) (*awx.WorkflowJobTemplate, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.WorkflowJobTemplate), args.Error(1)
}

func (mockAwx MockAWX) LaunchWorkflow(id int, data map[string]interface{}, params map[string]string) (*awx.JobLaunch, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.JobLaunch), args.Error(1)
}

func (mockAwx MockAWX) ListWorkflowJobTemplateSchedules(id int, params map[string]string) ([]*awx.Schedule, *awx.ListSchedulesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.Schedule), args.Get(1).(*awx.ListSchedulesResponse), args.Error(2)
}
func (mockAwx MockAWX) CreateWorkflowJobTemplateSchedule(id int, data map[string]interface{}, params map[string]string) (*awx.Schedule, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Schedule), args.Error(1)
}
