package awx

import (
	awx "github.com/adeo-opensource/goawx/client"
)

var resourceDataMapTeam = map[string]interface{}{
	"name":             "foo",
	"role_entitlement": []interface{}{"1", "2"},
}

var resourceDataMapTeamWithoutRoleEntitlment = map[string]interface{}{
	"name": "foo",
}

func (mockAwx MockAWX) ListTeams(params map[string]string) ([]*awx.Team, *awx.ListTeamsResponse, error) {
	args := mockAwx.Called(params)
	return args.Get(0).([]*awx.Team), args.Get(1).(*awx.ListTeamsResponse), args.Error(2)
}

func (mockAwx MockAWX) CreateTeam(data map[string]interface{}, params map[string]string) (*awx.Team, error) {
	args := mockAwx.Called(data, params)
	return args.Get(0).(*awx.Team), args.Error(1)
}

func (mockAwx MockAWX) UpdateTeam(id int, data map[string]interface{}, params map[string]string) (*awx.Team, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Team), args.Error(1)
}

func (mockAwx MockAWX) DeleteTeam(id int) (*awx.Team, error) {
	args := mockAwx.Called(id)
	return args.Get(0).(*awx.Team), args.Error(1)
}

func (mockAwx MockAWX) GetTeamByID(id int, params map[string]string) (*awx.Team, error) {
	args := mockAwx.Called(id, params)
	return args.Get(0).(*awx.Team), args.Error(1)
}
func (mockAwx MockAWX) UpdateTeamRoleEntitlement(id int, data map[string]interface{}, params map[string]string) (interface{}, error) {
	args := mockAwx.Called(id, data, params)
	return args.Get(0).(*awx.Team), args.Error(1)
}
