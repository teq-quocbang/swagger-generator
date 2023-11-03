package info

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
)

// more parameter to setup your own info
func SetupInfo(t *openapi2.T) {
	t.Info = openapi3.Info{
		Title:          "test Title",
		Description:    "Teqnological asia api",
		TermsOfService: "Golang team",
		Contact: &openapi3.Contact{
			Email: "contact@teqnological.asia",
		},
		License: &openapi3.License{
			Name: "Copy right @ 2023 teqnological",
		},
		Version: "1.1.1",
	}
}
