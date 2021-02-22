// Package server provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package server

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture   string          `json:"architecture"`
	ImageType      string          `json:"image_type"`
	UploadRequests []UploadRequest `json:"upload_requests"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Version string `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options AWSUploadRequestOptions `json:"options"`
	Type    UploadTypes             `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Status string      `json:"status"`
	Type   UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
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

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
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

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
