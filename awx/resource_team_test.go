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

func Test_resourceTeamCreate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Team not created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: Team not created",
				Detail:   "Team with name foo  in the Organization ID 0 not created, nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{}
				mockAWX.On("CreateTeam",
					mock.Anything,
					mock.Anything).
					Return(team, fmt.Errorf("nothing"))
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
					ID:          2,
				}}, &awx.ListTeamsResponse{}, nil)
			},
		},
		{
			name: "Team cannot be listed",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Create: Fail to find Team",
				Detail:   "Fail to find Team foo Organization ID 0, nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
					ID:          2,
				}}, &awx.ListTeamsResponse{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Team created can't find team role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeamWithoutRoleEntitlment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch team roles",
				Detail:   "Unable to load team roles with id 2: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{}
				mockAWX.On("CreateTeam",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						team.ID = 2
						team.Description = data["description"].(string) + "_created"
						team.Name = data["name"].(string) + "_created"
					}).
					Return(team, nil)
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
					ID:          2,
				}}, &awx.ListTeamsResponse{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, fmt.Errorf("nothing"))
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(team, nil)
			},
		},
		{
			name: "Team created",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{}
				mockAWX.On("CreateTeam",
					mock.Anything,
					mock.Anything).
					Run(func(args mock.Arguments) {
						data := args.Get(0).(map[string]interface{})
						team.ID = 2
						team.Description = data["description"].(string) + "_created"
						team.Name = data["name"].(string) + "_created"
					}).
					Return(team, nil)
				mockAWX.On("ListTeams", mock.Anything).Return([]*awx.Team{{
					Description: "a team",
					Name:        "the team",
					ID:          2,
				}}, &awx.ListTeamsResponse{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(team, nil)
				mockAWX.On("UpdateTeamRoleEntitlement", mock.Anything, mock.Anything, mock.Anything).Return(team, nil)
			},
			newData: map[string]interface{}{
				"name": "foo_created",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runTestCase(t, tt, resourceTeamCreate)
		})
	}
}

func Test_resourceTeamDelete(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Team not deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Team delete failed",
				Detail:   "Fail to delete Team, TeamID 1, got nothing ",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteTeam", mock.Anything, mock.Anything).Return(&awx.Team{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Team deleted",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("DeleteTeam", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceTeamDelete)
		})
	}
}

func Test_resourceTeamRead(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Team not found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch team",
				Detail:   "Unable to load team with id 1: got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{}
				mockAWX.On("GetTeamByID",
					mock.Anything,
					mock.Anything).
					Return(team, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Team found can't read role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetTeamByID",
					mock.Anything,
					mock.Anything).
					Return(team, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, fmt.Errorf("nothing"))
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch team roles",
				Detail:   "Unable to load team roles with id 1: got nothing",
			}},
		},
		{
			name: "Team found",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				team := &awx.Team{
					Name:        "toto",
					Description: "data_read",
				}
				mockAWX.On("GetTeamByID",
					mock.Anything,
					mock.Anything).
					Return(team, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
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
			runTestCase(t, tt, resourceTeamRead)
		})
	}
}

func Test_resourceTeamUpdate(t *testing.T) {
	tests := []commonTestCase{
		{
			name: "Team not updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeamWithoutRoleEntitlment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: Failed To Update Team",
				Detail:   "Fail to get Team with ID 1, got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, fmt.Errorf("nothing"))
			},
		},
		{
			name: "Team cannot be updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeamWithoutRoleEntitlment),
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: Failed To Update Team",
				Detail:   "Fail to get Team with ID 1, got nothing",
			}},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, fmt.Errorf("nothing"))

			},
			newData: map[string]interface{}{
				//TODO: schema
			},
		},
		{
			name: "Team updated can't found role entitlements",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeamWithoutRoleEntitlment),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, fmt.Errorf("nothing"))
			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Unable to fetch team roles",
				Detail:   "Unable to load team roles with id 1: got nothing",
			}},
		},
		{
			name: "Team updated",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeamWithoutRoleEntitlment),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Team updated with role entitlement",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
				mockAWX.On("UpdateTeamRoleEntitlement", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, nil)

			},
			newData: map[string]interface{}{

				//TODO: schema
			},
		},
		{
			name: "Team updated with role entitlement but cannot update team role",
			args: args{
				ctx: context.Background(),
				d:   schema.TestResourceDataRaw(t, resourceTeam().Schema, resourceDataMapTeam),
			},
			mock: func(mockAWX *MockAWX) {
				mockAWX.On("GetTeamByID", mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("UpdateTeam", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, nil)
				mockAWX.On("ListTeamRoleEntitlements", mock.Anything, mock.Anything).Return([]*awx.ApplyRole{{
					Description: "a team",
					Name:        "ApplyRole",
				}}, &awx.ListTeamRolesResponse{}, nil)
				mockAWX.On("UpdateTeamRoleEntitlement", mock.Anything, mock.Anything, mock.Anything).Return(&awx.Team{}, fmt.Errorf("nothing"))

			},
			want: diag.Diagnostics{{
				Severity: diag.Error,
				Summary:  "Update: Failed To Update Team Role Entitlement",
				Detail:   "Failed to add team role entitlement: got nothing",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.d.SetId("1")
			runTestCase(t, tt, resourceTeamUpdate)
		})
	}
}
