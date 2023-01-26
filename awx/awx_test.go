package awx

import (
	"context"
	awx "github.com/denouche/goawx/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

type MockAWX struct {
	mock.Mock
	awx.AWX
}

type args struct {
	ctx context.Context
	d   *schema.ResourceData
}

type commonTestCase struct {
	name    string
	args    args
	want    diag.Diagnostics
	newData map[string]interface{}
	mock    func(*MockAWX)
	id      string
}

var (
	resourceDataMapMissingId = map[string]interface{}{
		"name":          "foo",
		"credential_id": 0,
		"project_id":    0,
		"id":            0,
	}
	resourceDataMap = map[string]interface{}{
		"name":                "foo",
		"credential_id":       1,
		"credential":          1,
		"secret":              "terces",
		"inventory_id":        1,
		"inventory":           1,
		"project_id":          1,
		"id":                  1,
		"organization_id":     1,
		"organization":        1,
		"tenant":              "awx",
		"url":                 "http://localhost",
		"description":         "data",
		"client":              "terraform",
		"username":            "user",
		"password":            "pass",
		"project":             "proj",
		"ssh_key_data":        "a ssh key",
		"ssh_public_key_data": "a public ssh key",
		"ssh_key_unlock":      "alomora",
		"become_method":       "su",
		"become_username":     "root",
		"become_password":     "toor",
		"inputs":              "{ \"data\":\"data\"}",
		"injectors":           "{ \"data\":\"data\"}",
		"input_field_name":    "input_field",
		"target":              3,
		"source":              42,
		"metadata":            map[string]interface{}{"key": "value"},
		"image":               "dockerhub.io/small:tag",
		"instance_id":         1,
		"variables":           "toto:toto",
		"group_ids":           []interface{}{1},
	}
)

func runTestCase(t *testing.T, tt commonTestCase, callback func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics) {
	mockAWX := MockAWX{}
	tt.mock(&mockAWX)
	if got := callback(tt.args.ctx, tt.args.d, mockAWX); !reflect.DeepEqual(got, tt.want) {
		t.Errorf("%s = %v, want %v", tt.name, got, tt.want)
	}
	if tt.id != "" && tt.args.d.Id() != tt.id {
		t.Errorf("Id = %v, want %v", tt.args.d.Id(), tt.id)
	}
	for key, value := range tt.newData {
		if !reflect.DeepEqual(tt.args.d.Get(key), value) {
			t.Errorf("data.%s = %v, want %v", key, tt.args.d.Get(key), value)

		}
	}
}
