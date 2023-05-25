package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceProjectCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Error on list project",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: Fail to find Project",
				Detail:   "Fail to find Project foo Organization ID 1, nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListProjects", mock.Anything).Return([]*awx.Project{}, &awx.ListProjectsResponse{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Project cannot be created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: Project not created",
				Detail:   "Project with name foo  in the Organization ID 1 not created, nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				project := &awx.Project{}
				mockAWX.On("CreateProject",
					mock.Anything,
					mock.Anything).
					Return(project, fmt.Errorf("nothing"))
				mockAWX.On("ListProjects",
					mock.Anything).
					Return([]*awx.Project{project}, &awx.ListProjectsResponse{}, nil)
			},
		},
		{
			name: "Project created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			mock: func(mockAWX *MockAWX) {
				project := &awx.Project{}
				mockAWX.On("CreateProject",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						project.ID = 2
						project.Description = data["description"].(string) + "_created"
						project.Name = data["name"].(string) + "_created"
					}).
					Return(project, nil)
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(project, nil)
				mockAWX.On("ListProjects",
					mock.Anything).
					Return([]*awx.Project{project}, &awx.ListProjectsResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceProjectCreate)
		})
	}
}

func Test_resourceProjectDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Project not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch project",
				Detail:   "Unable to load project with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Project update cannot be canceled",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Delete: Fail to cancel Job",
				Detail:   "Fail to cancel the Job 1 for Project with ID 1, got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, nil)
				mockAWX.On("ProjectUpdateCancel", mock.Anything).Return(&awx.ProjectUpdateCancel{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Project not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Project delete failed",
				Detail:   "Fail to delete Project, ProjectID 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, nil)
				mockAWX.On("ProjectUpdateCancel", mock.Anything, mock.Anything).Return(&awx.ProjectUpdateCancel{}, nil)
				mockAWX.On("ProjectUpdateGet", mock.Anything).Return(&awx.Job{Finished: time.Now()}, nil)
				mockAWX.On("DeleteProject", mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, fmt.Errorf("nothing"))

			},
		},
		{
			name: "Project deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, nil)
				mockAWX.On("ProjectUpdateCancel", mock.Anything, mock.Anything).Return(&awx.ProjectUpdateCancel{}, nil)
				mockAWX.On("ProjectUpdateGet", mock.Anything).Return(&awx.Job{Finished: time.Now()}, nil)
				mockAWX.On("DeleteProject", mock.Anything).Return(&awx.Project{ID: 1, SummaryFields: &awx.Summary{CurrentJob: map[string]interface{}{"id": float64(1)}}}, nil)

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceProjectDelete)
		})
	}
}

func Test_resourceProjectRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Project not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch project",
				Detail:   "Unable to load project with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				project := &awx.Project{}
				mockAWX.On("GetProjectByID",
					mock.Anything,
					mock.Anything).
					Return(project, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Project found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			mock: func(mockAWX *MockAWX) {
				project := &awx.Project{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetProjectByID",
					mock.Anything,
					mock.Anything).
					Return(project, nil)
			},
			newData: map[string]interface{}{
				"description": "data_read",
				"name":        "toto",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceProjectRead)
		})
	}
}

func Test_resourceProjectUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Project cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: Fail To Update Project",
				Detail:   "Fail to get Project with ID 1, got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{}, nil)
				mockAWX.On("UpdateProject", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Project{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Project updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceProject().Schema, resourceDataMapProject),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetProjectByID", mock.Anything, mock.Anything).Return(&awx.Project{}, nil)
				mockAWX.On("UpdateProject", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Project{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceProjectUpdate)
		})
	}
}
