package options

import "github.com/getkin/kin-openapi/openapi2"

// more parameter to setup your own options
func SetupOptions(t *openapi2.T) {
	t.Swagger = "2.0"

	t.Schemes = []string{"http", "https"}

	t.BasePath = "/api"

	t.Consumes = []string{"application/json"}

	t.Host = "localhost"
}
