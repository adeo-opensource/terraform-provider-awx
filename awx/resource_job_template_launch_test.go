package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceJobDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Source not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateLaunch().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job",
				Detail:   "Unable to load job with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJob", mock.Anything, mock.Anything).Return(&awx.Job{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Source deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateLaunch().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJob", mock.Anything, mock.Anything).Return(&awx.Job{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceJobDelete)
		})
	}
}

func Test_resourceJobTemplateLaunch(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "JobTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateLaunch().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to create JobTemplate",
				Detail:   "JobTemplate with id 4 failed to create nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobLaunch{}
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				mockAWX.On("LaunchJob",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Return(jobTemplate, fmt.Errorf("nothing"))
			},
		},
		{
			name: "JobTemplate created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateLaunch().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobLaunch{}
				mockAWX.On("LaunchJob",
					mock.Anything,
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						jobTemplate.ID = 2

					}).
					Return(jobTemplate, nil)
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
			},
			id: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceJobTemplateLaunchCreate)
		})
	}
}

func Test_resourceJobTemplateLaunchCreate(t *testing.T) {
	type args struct {
		ctx context.Context
		d   *schema.ResourceData
		m   interface{}
	}
	tests := []struct {
		name string
		args args
		want diag.Diagnostics
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := resourceJobTemplateLaunchCreate(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceJobTemplateLaunchCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}
