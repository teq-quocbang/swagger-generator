package paths

import (
	"fmt"
	"reflect"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/teq-quocbang/swagger-generator/models"
)

// more parameter to setup your own paths
func SetUpPaths(paths models.Paths, t *openapi2.T) error {
	result := map[string]*openapi2.PathItem{}
	for apiPath, path := range paths {
		switch path.Method {
		case "get": // get
			result[apiPath] = &openapi2.PathItem{
				Get: &openapi2.Operation{
					Summary:     path.Summary,
					Consumes:    path.Consumes,
					OperationID: path.OperationID,
					Tags:        path.Tag,
					Parameters:  parseParameterSchema(path.Request, t, path.OperationID),
					Responses:   parseResponseSchema(path.Response),
				},
			}
		case "post": // post
			result[apiPath] = &openapi2.PathItem{
				Post: &openapi2.Operation{
					Summary:     path.Summary,
					Consumes:    path.Consumes,
					OperationID: path.OperationID,
					Tags:        path.Tag,
					Parameters:  parseParameterSchema(path.Request, t, path.OperationID),
					Responses:   parseResponseSchema(path.Response),
				},
			}
		case "put": // put
			result[apiPath] = &openapi2.PathItem{
				Put: &openapi2.Operation{
					Summary:     path.Summary,
					Consumes:    path.Consumes,
					OperationID: path.OperationID,
					Tags:        path.Tag,
					Parameters:  parseParameterSchema(path.Request, t, path.OperationID),
					Responses:   parseResponseSchema(path.Response),
				},
			}
		case "patch": // patch
			result[apiPath] = &openapi2.PathItem{
				Patch: &openapi2.Operation{
					Summary:     path.Summary,
					Consumes:    path.Consumes,
					OperationID: path.OperationID,
					Tags:        path.Tag,
					Parameters:  parseParameterSchema(path.Request, t, path.OperationID),
					Responses:   parseResponseSchema(path.Response),
				},
			}
		case "delete": // delete
			result[apiPath] = &openapi2.PathItem{
				Delete: &openapi2.Operation{
					Summary:     path.Summary,
					Consumes:    path.Consumes,
					OperationID: path.OperationID,
					Tags:        path.Tag,
					Parameters:  parseParameterSchema(path.Request, t, path.OperationID),
					Responses:   parseResponseSchema(path.Response),
				},
			}
		default:
			return fmt.Errorf("unexpected method, your method is: %s", path.Method)
		}
	}

	t.Paths = result
	return nil
}

func parseResponseSchema(response []models.Response) map[string]*openapi2.Response {
	result := map[string]*openapi2.Response{}
	for _, r := range response {
		strCode := fmt.Sprintf("%d", r.Code)
		result[strCode] = &openapi2.Response{
			Description: r.Message,
			Schema: &openapi3.SchemaRef{
				Value: parseSchema(r.Data),
			},
		}
	}

	return result
}

func parseParameterSchema(request []models.Request, t *openapi2.T, operationID string) openapi2.Parameters {
	result := make(openapi2.Parameters, len(request))

	for i, r := range request {
		parameter := &openapi2.Parameter{
			In:       r.In,
			Name:     r.Name,
			Required: r.Required,
			Type:     r.Type,
		}
		if r.In == "body" {
			// move schema to definitions
			if r.IsToModel {
				definitionKey := fmt.Sprintf("%sRequest", operationID)
				t.Definitions = map[string]*openapi3.SchemaRef{
					definitionKey: {
						Value: parseSchema(r.Schema),
					},
				}
				parameter.Schema = &openapi3.SchemaRef{
					Ref: fmt.Sprintf("#/definitions/%s", definitionKey),
				}
			} else {
				parameter.Schema = &openapi3.SchemaRef{
					Value: parseSchema(r.Schema),
				}
			}
		}
		result[i] = parameter
	}

	return result
}

func parseSchema(schema interface{}) *openapi3.Schema {
	typeOf := reflect.TypeOf(schema)
	result := &openapi3.Schema{}

	switch typeOf.Kind() {
	case reflect.Struct:
		fmt.Println(schema)
	case reflect.Array:
		fmt.Println(schema)
	case reflect.Slice:
		// this is special case
		// with slice case the only check first element type all remaining type is same
		value := reflect.ValueOf(schema).Index(1)
		types := getType(value.Elem().Type())

		switch types {
		case "slice":
			for i := 0; i < reflect.ValueOf(schema).Len(); i++ {
				values := reflect.ValueOf(schema).Index(i)
				result = &openapi3.Schema{
					Type: types,
					Items: &openapi3.SchemaRef{
						Value: parseSchema(values.Interface()),
					},
				}
			}
		case "object":
			for i := 0; i < reflect.ValueOf(schema).Len(); i++ {
				values := reflect.ValueOf(schema).Index(i)
				result = parseSchema(values.Interface())
			}
		default:
			result = &openapi3.Schema{
				Type:    types,
				Example: schema,
			}
		}
	case reflect.Map:
		// sey type = object
		result.Type = "object"

		// set schema
		value := reflect.ValueOf(schema)
		iter := value.MapRange()
		schemas := openapi3.Schemas{}
		//	{
		//		"access_token": "abc",
		//		"refresh_token": "xyz"
		//	}
		for iter.Next() {
			key := iter.Key()      // access_token
			values := iter.Value() // "abc"

			// to swagger result is
			//
			// access_token:
			// 	 type: "string"
			// 	 example: "abc"
			// 	 description: ""
			types := getType(values.Elem().Type())
			if types == "object" || types == "slice" { // with map or slice
				switch types {
				case "object": // golang type is map
					schemas[key.String()] = &openapi3.SchemaRef{
						Value: parseSchema(values.Interface()),
					}
				case "slice":
					schemas[key.String()] = &openapi3.SchemaRef{
						Value: &openapi3.Schema{
							Type: "array",
							Items: &openapi3.SchemaRef{
								Value: parseSchema(values.Interface()),
							},
						},
					}
				}

			} else { // with another types
				schemas[key.String()] = &openapi3.SchemaRef{
					Value: &openapi3.Schema{
						Type:    types,              // string
						Example: values.Interface(), // abc
						// TODO: missing description
					},
				}
			}
		}
		result.Properties = schemas
	default:
		fmt.Println(schema)
	}

	return result
}

func getType(r reflect.Type) string {
	switch r.Kind() {
	case reflect.String:
		return "string"
	case reflect.Slice:
		return "slice"
	case reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64, reflect.Uint, reflect.Uint8,
		reflect.Uint16, reflect.Uint32,
		reflect.Uint64, reflect.Float64:
		return "integer"
	case reflect.Map:
		return "object"
	default:
		return ""
	}
}
