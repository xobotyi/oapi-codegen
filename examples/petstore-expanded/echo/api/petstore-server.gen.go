// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/xobotyi/oapi-codegen version (devel) DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	. "github.com/xobotyi/oapi-codegen/examples/petstore-expanded/echo/api/models"
	"github.com/xobotyi/oapi-codegen/pkg/runtime"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Returns all pets
	// (GET /pets)
	FindPets(ctx echo.Context, params FindPetsParams) error
	// Creates a new pet
	// (POST /pets)
	AddPet(ctx echo.Context) error
	// Deletes a pet by ID
	// (DELETE /pets/{id})
	DeletePet(ctx echo.Context, id int64) error
	// Returns a pet by ID
	// (GET /pets/{id})
	FindPetByID(ctx echo.Context, id int64) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// FindPets converts echo context to params.
func (w *ServerInterfaceWrapper) FindPets(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params FindPetsParams
	// ------------- Optional query parameter "tags" -------------

	err = runtime.BindQueryParameter("form", true, false, "tags", ctx.QueryParams(), &params.Tags)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter tags: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPets(ctx, params)
	return err
}

// AddPet converts echo context to params.
func (w *ServerInterfaceWrapper) AddPet(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddPet(ctx)
	return err
}

// DeletePet converts echo context to params.
func (w *ServerInterfaceWrapper) DeletePet(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeletePet(ctx, id)
	return err
}

// FindPetByID converts echo context to params.
func (w *ServerInterfaceWrapper) FindPetByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int64

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, ctx.Param("id"), &id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.FindPetByID(ctx, id)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/pets", wrapper.FindPets)
	router.POST(baseURL+"/pets", wrapper.AddPet)
	router.DELETE(baseURL+"/pets/:id", wrapper.DeletePet)
	router.GET(baseURL+"/pets/:id", wrapper.FindPetByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"3d3v/hYAAP//wO3O5VcSAAA=",
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
