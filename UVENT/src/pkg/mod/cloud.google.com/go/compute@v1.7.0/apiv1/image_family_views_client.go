// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package compute

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	httptransport "google.golang.org/api/transport/http"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newImageFamilyViewsClientHook clientHook

// ImageFamilyViewsCallOptions contains the retry settings for each method of ImageFamilyViewsClient.
type ImageFamilyViewsCallOptions struct {
	Get []gax.CallOption
}

func defaultImageFamilyViewsRESTCallOptions() *ImageFamilyViewsCallOptions {
	return &ImageFamilyViewsCallOptions{
		Get: []gax.CallOption{},
	}
}

// internalImageFamilyViewsClient is an interface that defines the methods available from Google Compute Engine API.
type internalImageFamilyViewsClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Get(context.Context, *computepb.GetImageFamilyViewRequest, ...gax.CallOption) (*computepb.ImageFamilyView, error)
}

// ImageFamilyViewsClient is a client for interacting with Google Compute Engine API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The ImageFamilyViews API.
type ImageFamilyViewsClient struct {
	// The internal transport-dependent client.
	internalClient internalImageFamilyViewsClient

	// The call options for this service.
	CallOptions *ImageFamilyViewsCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageFamilyViewsClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageFamilyViewsClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *ImageFamilyViewsClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Get returns the latest image that is part of an image family, is not deprecated and is rolled out in the specified zone.
func (c *ImageFamilyViewsClient) Get(ctx context.Context, req *computepb.GetImageFamilyViewRequest, opts ...gax.CallOption) (*computepb.ImageFamilyView, error) {
	return c.internalClient.Get(ctx, req, opts...)
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageFamilyViewsRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ImageFamilyViewsClient
	CallOptions **ImageFamilyViewsCallOptions
}

// NewImageFamilyViewsRESTClient creates a new image family views rest client.
//
// The ImageFamilyViews API.
func NewImageFamilyViewsRESTClient(ctx context.Context, opts ...option.ClientOption) (*ImageFamilyViewsClient, error) {
	clientOpts := append(defaultImageFamilyViewsRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultImageFamilyViewsRESTCallOptions()
	c := &imageFamilyViewsRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &ImageFamilyViewsClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultImageFamilyViewsRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://compute.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://compute.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://compute.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageFamilyViewsRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageFamilyViewsRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *imageFamilyViewsRESTClient) Connection() *grpc.ClientConn {
	return nil
}

// Get returns the latest image that is part of an image family, is not deprecated and is rolled out in the specified zone.
func (c *imageFamilyViewsRESTClient) Get(ctx context.Context, req *computepb.GetImageFamilyViewRequest, opts ...gax.CallOption) (*computepb.ImageFamilyView, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/compute/v1/projects/%v/zones/%v/imageFamilyViews/%v", req.GetProject(), req.GetZone(), req.GetFamily())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v&%s=%v", "project", url.QueryEscape(req.GetProject()), "zone", url.QueryEscape(req.GetZone()), "family", url.QueryEscape(req.GetFamily())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).Get[0:len((*c.CallOptions).Get):len((*c.CallOptions).Get)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &computepb.ImageFamilyView{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}