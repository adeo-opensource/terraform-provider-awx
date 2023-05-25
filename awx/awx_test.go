package awx

import (
	"context"
	awx "github.com/adeo-opensource/goawx/client"
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
)

func runTestCase(t *testing.T, tt commonTestCase, callback func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics) {
	mockAWX := MockAWX{
		Mock: mock.Mock{},
		AWX: awx.AWX{
			CredentialService:           mockGeneric[awx.Credential]{},
			CredentialTypeService:       mockGeneric[awx.CredentialType]{},
			ExecutionEnvironmentService: mockGeneric[awx.ExecutionEnvironment]{},
			//GroupService:                mockGeneric[awx.GroupService]{},
			HostService:      mockGeneric[awx.HostService]{},
			InventoryService: mockGeneric[awx.InventoryService]{},
		},
	}
	tt.mock(&mockAWX)
	if got := callback(tt.args.ctx, tt.args.d, mockAWX.AWX); !reflect.DeepEqual(got, tt.want) {
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
