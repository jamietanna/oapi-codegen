// Package v2 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version (devel) DO NOT EDIT.
package v2

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Person defines model for Person.
type Person struct {
	FirstName          string `json:"FirstName"`
	GovernmentIDNumber *int64 `json:"GovernmentIDNumber,omitempty"`
	LastName           string `json:"LastName"`
}

// These are fields that specify a person. They are all optional, and
// would be used by an `Edit` style API endpoint, where each is optional.
type PersonProperties struct {
	FirstName          *string `json:"FirstName,omitempty"`
	GovernmentIDNumber *int64  `json:"GovernmentIDNumber,omitempty"`
	LastName           *string `json:"LastName,omitempty"`
}

// PersonWithID defines model for PersonWithID.
type PersonWithID struct {
	FirstName          string `json:"FirstName"`
	GovernmentIDNumber *int64 `json:"GovernmentIDNumber,omitempty"`
	ID                 int64  `json:"ID"`
	LastName           string `json:"LastName"`
}

// This is a person record as returned from a Create endpoint. It contains
// all the fields of a Person, with an additional resource UUID.
type JSONDefault = PersonWithID

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/5SUT2/bOBDFv8qAu0dBTrCLPegWrNtAQJEaaNIc4gAZSyOLKTVkyVEMwfB3L0jJ/+AW",
	"aX0awOTjvDe/0VZVtnOWiSWoYqtC1VKHqVyQD5ZjhcZ8blTxtFV/e2pUof6aHW/Npiuz8fzCW0deNAW1",
	"y7bK0/dee6pV8aQ+ah/kDjtSmfqEU/m8e85UTaHy2omO76n7VgfQARBcksxgo6WFDrlGsX6AJgoBcg0G",
	"gwBjRxmsegGbJNBAOV8y992KfA5JbmN7U8OKwJP0nqmG1QAIL7ckLxBkMAQ3izKHR4KO/JpAWpqeX7I7",
	"eBo7QbbSkocvyTlsWl21YNkM4Lx90zUF2PuGRpOpQ75klSkZHKlC2dUrVaJ2mbqIrNheZEGBAD1NQiAt",
	"CgRHlW6GQ0LRJA3pGBpziCGLGS354L0Pk2+Glw+1PnUOxLWzmiWDTUuegLBq4xD2WqMDd9bqcaDFdm8u",
	"iNe8juZu7Rt57oilnN+lWcRjjfUdiiqUZvnv32MomoXW5OPFAxuXqrtfhviopS3nf0prYvTc1Cjybpu7",
	"7Iztcv47JIOnyvoaMBw5bLztAOF/Tyh0GEMOpUBlWVBzWHKcaiRygsA2gLA4XQ5kwLrWE/6egu19RfDw",
	"UM5/yl7sX3NjU8ZaTPzvnoIEuInxQUosJD2VqTfyYXR0nV/lVzF164jRaVWof/Kr/DqygdKmBGfOYEWt",
	"NfU48jXJJdhf0ei0zgE2yAIoYChus2WCKJVBsCAxwA6/UQSfOmjRuWE0FGeGUaysVaEWJ0/GyQRnOewX",
	"qsHepBZioMSpROeMrpLA7HX6zo1sxOp9cibeUpDnzk7d79LvRwAAAP//lzc18GUFAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
