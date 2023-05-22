//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package v20220901privatepreview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// ResourceGroupsClient contains the methods for the ResourceGroups group.
// Don't use this type directly, use NewResourceGroupsClient() instead.
type ResourceGroupsClient struct {
	host string
	pl runtime.Pipeline
}

// NewResourceGroupsClient creates a new instance of ResourceGroupsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewResourceGroupsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceGroupsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceGroupsClient{
		host: ep,
pl: pl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a ResourceGroupResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The plane type.
// planeName - The name of the plane
// resourceGroupName - UCP resourcegroup name
// resource - Resource create parameters.
// options - ResourceGroupsClientCreateOrUpdateOptions contains the optional parameters for the ResourceGroupsClient.CreateOrUpdate
// method.
func (client *ResourceGroupsClient) CreateOrUpdate(ctx context.Context, planeType string, planeName string, resourceGroupName string, resource ResourceGroupResource, options *ResourceGroupsClientCreateOrUpdateOptions) (ResourceGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, planeType, planeName, resourceGroupName, resource, options)
	if err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ResourceGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ResourceGroupsClient) createOrUpdateCreateRequest(ctx context.Context, planeType string, planeName string, resourceGroupName string, resource ResourceGroupResource, options *ResourceGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/resourcegroups/{resourceGroupName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", planeName)
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ResourceGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ResourceGroupsClientCreateOrUpdateResponse, error) {
	result := ResourceGroupsClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return ResourceGroupsClientCreateOrUpdateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroupResource); err != nil {
		return ResourceGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an existing ResourceGroupResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The plane type.
// planeName - The name of the plane
// resourceGroupName - UCP resourcegroup name
// options - ResourceGroupsClientDeleteOptions contains the optional parameters for the ResourceGroupsClient.Delete method.
func (client *ResourceGroupsClient) Delete(ctx context.Context, planeType string, planeName string, resourceGroupName string, options *ResourceGroupsClientDeleteOptions) (ResourceGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, planeType, planeName, resourceGroupName, options)
	if err != nil {
		return ResourceGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return ResourceGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return client.deleteHandleResponse(resp)
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceGroupsClient) deleteCreateRequest(ctx context.Context, planeType string, planeName string, resourceGroupName string, options *ResourceGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/resourcegroups/{resourceGroupName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", planeName)
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ResourceGroupsClient) deleteHandleResponse(resp *http.Response) (ResourceGroupsClientDeleteResponse, error) {
	result := ResourceGroupsClientDeleteResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return ResourceGroupsClientDeleteResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	return result, nil
}

// Get - Retrieves information about a ResourceGroupResource
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-09-01-privatepreview
// planeType - The plane type.
// planeName - The name of the plane
// resourceGroupName - UCP resourcegroup name
// options - ResourceGroupsClientGetOptions contains the optional parameters for the ResourceGroupsClient.Get method.
func (client *ResourceGroupsClient) Get(ctx context.Context, planeType string, planeName string, resourceGroupName string, options *ResourceGroupsClientGetOptions) (ResourceGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, planeType, planeName, resourceGroupName, options)
	if err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ResourceGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ResourceGroupsClient) getCreateRequest(ctx context.Context, planeType string, planeName string, resourceGroupName string, options *ResourceGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/resourcegroups/{resourceGroupName}"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", planeName)
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourceGroupsClient) getHandleResponse(resp *http.Response) (ResourceGroupsClientGetResponse, error) {
	result := ResourceGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroupResource); err != nil {
		return ResourceGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRootScopePager - Lists information about all ResourceGroupResource
// Generated from API version 2022-09-01-privatepreview
// planeType - The plane type.
// planeName - The name of the plane
// options - ResourceGroupsClientListByRootScopeOptions contains the optional parameters for the ResourceGroupsClient.ListByRootScope
// method.
func (client *ResourceGroupsClient) NewListByRootScopePager(planeType string, planeName string, options *ResourceGroupsClientListByRootScopeOptions) (*runtime.Pager[ResourceGroupsClientListByRootScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[ResourceGroupsClientListByRootScopeResponse]{
		More: func(page ResourceGroupsClientListByRootScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceGroupsClientListByRootScopeResponse) (ResourceGroupsClientListByRootScopeResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByRootScopeCreateRequest(ctx, planeType, planeName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ResourceGroupsClientListByRootScopeResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ResourceGroupsClientListByRootScopeResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ResourceGroupsClientListByRootScopeResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRootScopeHandleResponse(resp)
		},
	})
}

// listByRootScopeCreateRequest creates the ListByRootScope request.
func (client *ResourceGroupsClient) listByRootScopeCreateRequest(ctx context.Context, planeType string, planeName string, options *ResourceGroupsClientListByRootScopeOptions) (*policy.Request, error) {
	urlPath := "/planes/{planeType}/{planeName}/resourcegroups"
	if planeType == "" {
		return nil, errors.New("parameter planeType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeType}", url.PathEscape(planeType))
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", planeName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRootScopeHandleResponse handles the ListByRootScope response.
func (client *ResourceGroupsClient) listByRootScopeHandleResponse(resp *http.Response) (ResourceGroupsClientListByRootScopeResponse, error) {
	result := ResourceGroupsClientListByRootScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGroupResourceListResult); err != nil {
		return ResourceGroupsClientListByRootScopeResponse{}, err
	}
	return result, nil
}

