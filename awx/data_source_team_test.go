package awx

import (
	"context"
	"fmt"
	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"testing"
)

func (mockAwx MockAWX) ListTeamRoleEntitlements(id int, params map[string]string) ([]*awx.ApplyRole, *awx.ListTeamRolesResponse, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).([]*awx.ApplyRole), args.Get(1).(*awx.ListTeamRolesResponse), args.Error(2)

}

func Test_dataSourceTeamsRead(t *testing.T) {

	tests := []commonTestCase{
		{
			name: "Error on list team",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceTeam().Schema, resourceDataMapMissingId),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Fail to fetch Team",
				Detail:   "Fail to find the team got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{}, &awx.ListTeamsResponse{}, fmt.Errorf("nothing"))
			},
		},

		{
			name: "Two team",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: find more than one Element",
				Detail:   "The Query Returns more than one team, 2",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{}, {}}, &awx.ListTeamsResponse{}, nil)
			},
			newData: nil,
		},
		{
			name: "One team with role entitlements",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceTeam().Schema, resourceDataMapTeam),
			},
			want: nil,
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
					ID:          2,
				}}, &awx.ListTeamsResponse{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
			},
			id: "2",
			newData: map[string]interface{}{
				//TODO: schema
				"name": "the team",
			},
		},
		{
			name: "One team without role entitlements",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, dataSourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Get: Failed to fetch team role entitlements",
				Detail:   "Fail to retrieve team role entitlements got: nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
				}}, &awx.ListTeamsResponse{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{}}, &awx.ListTeamRolesResponse{}, fmt.Errorf("nothing"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, dataSourceTeamsRead)
		})
	}
}
