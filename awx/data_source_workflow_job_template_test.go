package awx

import (
	"context"
	"fmt"
	awx "github.com/adeo-opensource/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func Test_dataSourceWorkflowJobTemplateRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Error on list workflow job template",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceWorkflowJobTemplate().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Workflow Job Template",
				Detail:   "Fail to find the workflow job template got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListWorkflowJobTemplates", mock.Anything).Return([]*awx.WorkflowJobTemplate{}, &awx.ListWorkflowJobTemplatesResponse{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "two workflow job template without the right name and id",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceWorkflowJobTemplate().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Workflow Job Template, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListWorkflowJobTemplates", mock.Anything).Return([]*awx.WorkflowJobTemplate{
					{
						Name:                 "a",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					},
					{
						Name:                 "b",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					}}, &awx.ListWorkflowJobTemplatesResponse{}, nil)
			},
		},
		{
			name: "one workflow job template without the right name and id",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListWorkflowJobTemplates", mock.Anything).Return([]*awx.WorkflowJobTemplate{
					{
						Name:                 "none",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					}}, &awx.ListWorkflowJobTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "none",
				//TODO: schema
				//"description":              "a workflow job template",
				//"organization_id":          2,
				//"inventory_id":             3,
				//"survey_enabled":           true,
				//"allow_simultaneous":       false,
				//"ask_variables_on_launch":  false,
				//"limit":                    1,
				//"scm_branch":               "main",
				//"ask_inventory_on_launch":  "true",
				//"ask_scm_branch_on_launch": "false",
				//"ask_limit_on_launch":      "true",
				//"webhook_service":          "awx",
				//"webhook_credential":       "awx-credentials",
				//"variables":                "toto:toto\n",
			},
		},
		{
			name: "two workflow job template with id and  without the right name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceWorkflowJobTemplate().Schema, resourceDataMapWorkflowJobTemplate),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one Workflow Job Template, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListWorkflowJobTemplates", mock.Anything).Return([]*awx.WorkflowJobTemplate{
					{
						Name:                 "a",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					},
					{
						Name:                 "b",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					}}, &awx.ListWorkflowJobTemplatesResponse{}, nil)
			},
		},
		{
			name: "One workflow job template with the right name",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceWorkflowJobTemplate().Schema, resourceDataMapMissingId),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListWorkflowJobTemplates", mock.Anything).Return([]*awx.WorkflowJobTemplate{
					{
						Name:                 "foo",
						Description:          "a workflow job template",
						Organization:         2,
						Inventory:            3,
						SurveyEnabled:        true,
						AllowSimultaneous:    false,
						AskVariablesOnLaunch: false,
						Limit:                1,
						ScmBranch:            "main",
						AskLimitOnLaunch:     true,
						AskInventoryOnLaunch: true,
						AskScmBranchOnLaunch: false,
						WebhookService:       "awx",
						WebhookCredential:    "awx-credentials",
						ExtraVars:            "toto:toto",
					}}, &awx.ListWorkflowJobTemplatesResponse{}, nil)
			},
			newData: map[string]interface{}{
				"name": "foo",
				//TODO: schema
				//"description":              "a workflow job template",
				//"organization_id":          2,
				//"inventory_id":             3,
				//"survey_enabled":           true,
				//"allow_simultaneous":       false,
				//"ask_variables_on_launch":  false,
				//"limit":                    1,
				//"scm_branch":               "main",
				//"ask_inventory_on_launch":  "true",
				//"ask_scm_branch_on_launch": "false",
				//"ask_limit_on_launch":      "true",
				//"webhook_service":          "awx",
				//"webhook_credential":       "awx-credentials",
				//"variables":                "toto:toto\n",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceWorkflowJobTemplateRead)
		})
	}
}
