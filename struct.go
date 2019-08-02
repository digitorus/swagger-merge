package main

import (
	"encoding/json"
)

// SwaggerObject is the root document object
// http://swagger.io/specification/#swaggerObject
type SwaggerObject struct {
	Swagger             string                      `json:"swagger"`
	Info                *json.RawMessage            `json:"info"`
	Host                string                      `json:"host,omitempty"`
	BasePath            string                      `json:"basePath,omitempty"`
	Schemes             []string                    `json:"schemes"`
	Consumes            []string                    `json:"consumes"`
	Produces            []string                    `json:"produces"`
	Paths               map[string]json.RawMessage `json:"paths"`
	Definitions         map[string]json.RawMessage            `json:"definitions"`
	StreamDefinitions   map[string]json.RawMessage            `json:"x-stream-definitions,omitempty"`
	SecurityDefinitions map[string]json.RawMessage            `json:"securityDefinitions,omitempty"`
	Security            *[]json.RawMessage            `json:"security,omitempty"`
	ExternalDocs        *json.RawMessage            `json:"externalDocs,omitempty"`
}
