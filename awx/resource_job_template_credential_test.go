package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/stretchr/testify/mock"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Test_resourceJobTemplateCredentialsCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "JobTemplate not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job template",
				Detail:   "Unable to load job template with id 4: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{}
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(jobTemplate, fmt.Errorf("nothing"))
				mockAWX.On("AssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(jobTemplate, nil)

			},
		},
		{
			name: "JobTemplate can't be associate",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: JobTemplate not AssociateCredentials",
				Detail:   "Fail to add credentials with Id 0, for Template ID 4, got error: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{}
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(jobTemplate, nil)
				mockAWX.On("AssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(jobTemplate, fmt.Errorf("nothing"))

			},
		},
		{
			name: "JobTemplate created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				jobTemplate := &awx.JobTemplate{ID: 10}
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(jobTemplate, nil)
				mockAWX.On("AssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(jobTemplate, nil)

			},
			id: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceJobTemplateCredentialsCreate)
		})
	}

}

func Test_resourceJobTemplateCredentialsDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "JobTemplate not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "JobTemplate DisAssociateCredentials delete failed",
				Detail:   "Fail to delete JobTemplate DisAssociateCredentials, DisAssociateCredentials 0, from JobTemplateID 4 got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				mockAWX.On("DisAssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "JobTemplate not deleted can't fetch job template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch job template",
				Detail:   "Unable to load job template with id 4: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))
				mockAWX.On("DisAssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "JobTemplate deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceJobTemplateCredentials().Schema, resourceDataMapJobTemplate),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetJobTemplateByID", mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
				mockAWX.On("DisAssociateCredentials", mock.Anything, mock.Anything, mock.Anything).Return(&awx.JobTemplate{}, nil)
			},
			id: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceJobTemplateCredentialsDelete)
		})
	}

}
