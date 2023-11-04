package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/teq-quocbang/swagger-generator/definitions"
	"github.com/teq-quocbang/swagger-generator/info"
	"github.com/teq-quocbang/swagger-generator/options"
	"github.com/teq-quocbang/swagger-generator/paths"
	"gopkg.in/yaml.v3"
)

func main() {
	t := &openapi2.T{}

	// info
	info.SetupInfo(t)

	// somethings options
	options.SetupOptions(t)

	// definitions
	definitions.SetUpDefinitions(t)

	// paths
	paths.SetUpPaths(t)

	data, err := os.ReadFile("api.json")
	if err != nil {
		log.Fatalf("failed to read api file, error: %v", err)
	}
	var apis APIS
	if err := json.Unmarshal(data, &apis); err != nil {
		log.Fatalf("failed to unmarshal json, error: %v", err)
	}

	for apiURL, api := range apis {
		fmt.Println(apiURL, api)
	}

	swaggerByte, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("failed to marshal swagger, error: %v", err)
	}

	err = os.WriteFile("swagger.yaml", swaggerByte, 0644)
	if err != nil {
		log.Fatalf("failed to write file, error: %v", err)
	}
}
