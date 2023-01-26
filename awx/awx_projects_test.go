package awx

import (
	awx "github.com/denouche/goawx/client"
)

var resourceDataMapProject = map[string]interface{}{
	"name":            "foo",
	"organization_id": 1,
	"project_id":      1,
	"id":              1,
}

func (mockAwx MockAWX) GetProjectByID(id int, params map[string]string) (*awx.Project, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Project), args.Error(1)
}

func (mockAwx MockAWX) ListProjects(params map[string]string) ([]*awx.Project, *awx.ListProjectsResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Project), args.Get(1).(*awx.ListProjectsResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateProject(data map[string]interface{}, params map[string]string) (*awx.Project, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Project), args.Error(1)
}

func (mockAwx MockAWX) UpdateProject(id int, data map[string]interface{}, params map[string]string) (*awx.Project, error) {
	args := mockAwx.Called(id, data)
	return args.Get(0).(*awx.Project), args.Error(1)
}

func (mockAwx MockAWX) DeleteProject(id int) (*awx.Project, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Project), args.Error(1)
}

func (mockAwx MockAWX) ProjectUpdateCancel(id int) (*awx.ProjectUpdateCancel, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.ProjectUpdateCancel), args.Error(1)
}
func (mockAwx MockAWX) ProjectUpdateGet(id int) (*awx.Job, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Job), args.Error(1)
}
