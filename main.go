package main

import (
	"encoding/json"
	"log"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/getkin/kin-openapi/openapi2"
	"github.com/teq-quocbang/swagger-generator/info"
	"github.com/teq-quocbang/swagger-generator/models"
	"github.com/teq-quocbang/swagger-generator/options"
	"github.com/teq-quocbang/swagger-generator/paths"
)

func main() {
	t := &openapi2.T{}

	// info
	info.SetupInfo(t)

	// somethings options
	options.SetupOptions(t)

	// definitions
	// definitions.SetUpDefinitions(t)

	// paths
	data, err := os.ReadFile("api.json")
	if err != nil {
		log.Fatalf("failed to read api file, error: %v", err)
	}

	var modelPath models.Paths
	if err := json.Unmarshal(data, &modelPath); err != nil {
		log.Fatalf("failed to unmarshal paths, error: %v", err)
	}

	err = paths.SetUpPaths(modelPath, t)
	if err != nil {
		log.Fatal(err)
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
