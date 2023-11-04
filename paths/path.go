package paths

import (
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"
)

// more parameter to setup your own paths
func SetUpPaths(t *openapi2.T) {
	t.Paths = map[string]*openapi2.PathItem{
		"/user/login": {
			Post: &openapi2.Operation{
				Summary:     "login",
				Description: "login to server",
				Tags:        []string{"Account"},
				OperationID: "Login",
				Parameters: openapi2.Parameters{
					{
						In:       "body",
						Name:     "body",
						Required: true,

						Schema: &openapi3.SchemaRef{
							Ref: "#/definitions/LoginRequest",
						},
					},
				},
				Responses: map[string]*openapi2.Response{
					"200": {
						Description: "OK",
						Schema: &openapi3.SchemaRef{
							Ref: "#/definitions/LoginResponse",
						},
					},
				},
			},
		},
		"/user/info/{id}": {
			Get: &openapi2.Operation{
				Summary:     "get user info",
				Description: "get user",
				OperationID: "ListUserInfo",
				Tags: []string{
					"Account",
				},
				Parameters: openapi2.Parameters{
					{
						In:       "path",
						Name:     "id",
						Required: true,
						Type:     "string",
					},
					{
						In:       "body",
						Name:     "body",
						Required: true,
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Properties: openapi3.Schemas{
									"Data": &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type: "object",
										},
									},
								},
							},
						},
					},
				},
				Responses: map[string]*openapi2.Response{
					"200": {
						Description: "OK",
						Schema: &openapi3.SchemaRef{
							Value: &openapi3.Schema{
								Properties: openapi3.Schemas{
									"Data": &openapi3.SchemaRef{
										Value: &openapi3.Schema{
											Type: "array",
											Items: &openapi3.SchemaRef{
												Value: &openapi3.Schema{
													Properties: openapi3.Schemas{
														"Username": &openapi3.SchemaRef{
															Value: &openapi3.Schema{
																Type:        "string",
																Description: "name of the user",
															},
														},
														"Age": &openapi3.SchemaRef{
															Value: &openapi3.Schema{
																Type:        "integer",
																Description: "user age",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
