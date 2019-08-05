package main

import (
	"encoding/json"
)

// SwaggerObject is the root document object
// http://swagger.io/specification/#swaggerObject
type SwaggerObject struct {
	Swagger             string                     `json:"swagger"`
	Info                *InfoObject                `json:"info"`
	Host                string                     `json:"host,omitempty"`
	BasePath            string                     `json:"basePath,omitempty"`
	Schemes             []string                   `json:"schemes"`
	Consumes            []string                   `json:"consumes"`
	Produces            []string                   `json:"produces"`
	Paths               map[string]json.RawMessage `json:"paths"`
	Definitions         map[string]json.RawMessage `json:"definitions"`
	StreamDefinitions   map[string]json.RawMessage `json:"x-stream-definitions,omitempty"`
	SecurityDefinitions map[string]json.RawMessage `json:"securityDefinitions,omitempty"`
	Security            *[]json.RawMessage         `json:"security,omitempty"`
	ExternalDocs        *json.RawMessage           `json:"externalDocs,omitempty"`
}

// InfoObject contains basic information about the API as specified at
// http://swagger.io/specification/#infoObject
type InfoObject struct {
	Title          string `json:"title"`
	Description    string `json:"description,omitempty"`
	TermsOfService string `json:"termsOfService,omitempty"`
	Version        string `json:"version"`

	Contact *ContactObject `json:"contact,omitempty"`
	License *LicenseObject `json:"license,omitempty"`
}

// ContactObject contains a contact as specified at
// http://swagger.io/specification/#contactObject
type ContactObject struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

// LicenseObject contains a license as specified at
// http://swagger.io/specification/#licenseObject
type LicenseObject struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}
