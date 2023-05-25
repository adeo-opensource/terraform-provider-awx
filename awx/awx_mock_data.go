package awx

var resourceDataMapCredential = map[string]interface{}{
	"credential_id":       1,
	"secret":              "terces",
	"id":                  1,
	"tenant":              "awx",
	"url":                 "http://localhost",
	"description":         "data",
	"client":              "terraform",
	"organization_id":     1,
	"name":                "foo",
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
}

var resourceDataMapExecutionEnvironment = map[string]interface{}{
	"name":         "foo",
	"organization": 1,
	"credential":   1,
	"image":        "dockerhub.io/small:tag",
	"description":  "data",
}

var resourceDataMapCredentialType = map[string]interface{}{
	"id":          1,
	"inputs":      "{ \"data\":\"data\"}",
	"injectors":   "{ \"data\":\"data\"}",
	"description": "data",
	"name":        "foo",
}
