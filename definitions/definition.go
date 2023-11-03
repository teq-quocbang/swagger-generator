package definitions

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
)

// more parameter to setup your own definitions
func SetUpDefinitions(t *openapi2.T) {
	t.Definitions = map[string]*openapi3.SchemaRef{
		"LoginResponse": {
			Value: openapi3.NewObjectSchema().WithDefault(`{"access_token": "abc", "refresh_token": "abc"}`),
		},
		"LoginRequest": {
			Value: openapi3.NewObjectSchema().WithDefault(`{"username": "", "password": ""}`),
		},
	}
}
