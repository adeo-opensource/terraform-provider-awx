package awx

import (
	"context"
	"reflect"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCredentialTypeServiceDeleteByID(t *testing.T) {
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
			if got := CredentialTypeServiceDeleteByID(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CredentialTypeServiceDeleteByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentialsServiceDeleteByID(t *testing.T) {
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
			if got := CredentialsServiceDeleteByID(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CredentialsServiceDeleteByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagCreateFail(t *testing.T) {
	type args struct {
		tfMethode string
		err       error
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
			if got := buildDiagCreateFail(tt.args.tfMethode, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDiagCreateFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagDeleteFail(t *testing.T) {
	type args struct {
		tfMethode string
		details   string
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
			if got := buildDiagDeleteFail(tt.args.tfMethode, tt.args.details); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDiagDeleteFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagDeleteFailDetails(t *testing.T) {
	type args struct {
		tfMethode     string
		detailsString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildDiagDeleteFailDetails(tt.args.tfMethode, tt.args.detailsString); got != tt.want {
				t.Errorf("buildDiagDeleteFailDetails() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagDeleteFailSummary(t *testing.T) {
	type args struct {
		tfMethode string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildDiagDeleteFailSummary(tt.args.tfMethode); got != tt.want {
				t.Errorf("buildDiagDeleteFailSummary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagNotFoundFail(t *testing.T) {
	type args struct {
		tfMethode string
		id        int
		err       error
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
			if got := buildDiagNotFoundFail(tt.args.tfMethode, tt.args.id, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDiagNotFoundFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagUpdateFail(t *testing.T) {
	type args struct {
		tfMethode string
		id        int
		err       error
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
			if got := buildDiagUpdateFail(tt.args.tfMethode, tt.args.id, tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDiagUpdateFail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildDiagnosticsMessage(t *testing.T) {
	type args struct {
		diagSummary string
		diagDetails string
		detailsVars []interface{}
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
			if got := buildDiagnosticsMessage(tt.args.diagSummary, tt.args.diagDetails, tt.args.detailsVars...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildDiagnosticsMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStateIDToNummeric(t *testing.T) {
	type args struct {
		tfElement string
		d         *schema.ResourceData
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 diag.Diagnostics
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := convertStateIDToNummeric(tt.args.tfElement, tt.args.d)
			if got != tt.want {
				t.Errorf("convertStateIDToNummeric() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("convertStateIDToNummeric() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_normalizeJsonOk(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := normalizeJsonOk(tt.args.s)
			if got != tt.want {
				t.Errorf("normalizeJsonOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("normalizeJsonOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_normalizeJsonYaml(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalizeJsonYaml(tt.args.s); got != tt.want {
				t.Errorf("normalizeJsonYaml() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeYamlOk(t *testing.T) {
	type args struct {
		s interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add runTestCase cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := normalizeYamlOk(tt.args.s)
			if got != tt.want {
				t.Errorf("normalizeYamlOk() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("normalizeYamlOk() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_resourceJobTemplateDelete(t *testing.T) {
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
			if got := resourceJobTemplateDelete(tt.args.ctx, tt.args.d, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("resourceJobTemplateDelete() = %v, want %v", got, tt.want)
			}
		})
	}
}
