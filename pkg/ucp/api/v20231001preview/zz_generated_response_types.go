//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

// APIVersionsClientCreateOrUpdateResponse contains the response from method APIVersionsClient.BeginCreateOrUpdate.
type APIVersionsClientCreateOrUpdateResponse struct {
	// The resource type for defining an API version of a resource type supported by the containing resource provider.
	APIVersionResource
}

// APIVersionsClientDeleteResponse contains the response from method APIVersionsClient.BeginDelete.
type APIVersionsClientDeleteResponse struct {
	// placeholder for future response values
}

// APIVersionsClientGetResponse contains the response from method APIVersionsClient.Get.
type APIVersionsClientGetResponse struct {
	// The resource type for defining an API version of a resource type supported by the containing resource provider.
	APIVersionResource
}

// APIVersionsClientListResponse contains the response from method APIVersionsClient.NewListPager.
type APIVersionsClientListResponse struct {
	// The response of a ApiVersionResource list operation.
	APIVersionResourceListResult
}

// AwsCredentialsClientCreateOrUpdateResponse contains the response from method AwsCredentialsClient.CreateOrUpdate.
type AwsCredentialsClientCreateOrUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AwsCredentialResource
}

// AwsCredentialsClientDeleteResponse contains the response from method AwsCredentialsClient.Delete.
type AwsCredentialsClientDeleteResponse struct {
	// placeholder for future response values
}

// AwsCredentialsClientGetResponse contains the response from method AwsCredentialsClient.Get.
type AwsCredentialsClientGetResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AwsCredentialResource
}

// AwsCredentialsClientListResponse contains the response from method AwsCredentialsClient.NewListPager.
type AwsCredentialsClientListResponse struct {
	// The response of a AwsCredentialResource list operation.
	AwsCredentialResourceListResult
}

// AwsCredentialsClientUpdateResponse contains the response from method AwsCredentialsClient.Update.
type AwsCredentialsClientUpdateResponse struct {
	// Concrete tracked resource types can be created by aliasing this type using a specific property type.
	AwsCredentialResource
}

// AwsPlanesClientCreateOrUpdateResponse contains the response from method AwsPlanesClient.BeginCreateOrUpdate.
type AwsPlanesClientCreateOrUpdateResponse struct {
	// The AWS plane resource
	AwsPlaneResource
}

// AwsPlanesClientDeleteResponse contains the response from method AwsPlanesClient.BeginDelete.
type AwsPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// AwsPlanesClientGetResponse contains the response from method AwsPlanesClient.Get.
type AwsPlanesClientGetResponse struct {
	// The AWS plane resource
	AwsPlaneResource
}

// AwsPlanesClientListResponse contains the response from method AwsPlanesClient.NewListPager.
type AwsPlanesClientListResponse struct {
	// The response of a AwsPlaneResource list operation.
	AwsPlaneResourceListResult
}

// AwsPlanesClientUpdateResponse contains the response from method AwsPlanesClient.BeginUpdate.
type AwsPlanesClientUpdateResponse struct {
	// The AWS plane resource
	AwsPlaneResource
}

// AzureCredentialsClientCreateOrUpdateResponse contains the response from method AzureCredentialsClient.CreateOrUpdate.
type AzureCredentialsClientCreateOrUpdateResponse struct {
	// Represents Azure Credential Resource
	AzureCredentialResource
}

// AzureCredentialsClientDeleteResponse contains the response from method AzureCredentialsClient.Delete.
type AzureCredentialsClientDeleteResponse struct {
	// placeholder for future response values
}

// AzureCredentialsClientGetResponse contains the response from method AzureCredentialsClient.Get.
type AzureCredentialsClientGetResponse struct {
	// Represents Azure Credential Resource
	AzureCredentialResource
}

// AzureCredentialsClientListResponse contains the response from method AzureCredentialsClient.NewListPager.
type AzureCredentialsClientListResponse struct {
	// The response of a AzureCredentialResource list operation.
	AzureCredentialResourceListResult
}

// AzureCredentialsClientUpdateResponse contains the response from method AzureCredentialsClient.Update.
type AzureCredentialsClientUpdateResponse struct {
	// Represents Azure Credential Resource
	AzureCredentialResource
}

// AzurePlanesClientCreateOrUpdateResponse contains the response from method AzurePlanesClient.BeginCreateOrUpdate.
type AzurePlanesClientCreateOrUpdateResponse struct {
	// The Azure plane resource.
	AzurePlaneResource
}

// AzurePlanesClientDeleteResponse contains the response from method AzurePlanesClient.BeginDelete.
type AzurePlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// AzurePlanesClientGetResponse contains the response from method AzurePlanesClient.Get.
type AzurePlanesClientGetResponse struct {
	// The Azure plane resource.
	AzurePlaneResource
}

// AzurePlanesClientListResponse contains the response from method AzurePlanesClient.NewListPager.
type AzurePlanesClientListResponse struct {
	// The response of a AzurePlaneResource list operation.
	AzurePlaneResourceListResult
}

// AzurePlanesClientUpdateResponse contains the response from method AzurePlanesClient.BeginUpdate.
type AzurePlanesClientUpdateResponse struct {
	// The Azure plane resource.
	AzurePlaneResource
}

// LocationsClientCreateOrUpdateResponse contains the response from method LocationsClient.BeginCreateOrUpdate.
type LocationsClientCreateOrUpdateResponse struct {
	// The resource type for defining a location of the containing resource provider. The location resource represents a logical
// location where the resource provider operates.
	LocationResource
}

// LocationsClientDeleteResponse contains the response from method LocationsClient.BeginDelete.
type LocationsClientDeleteResponse struct {
	// placeholder for future response values
}

// LocationsClientGetResponse contains the response from method LocationsClient.Get.
type LocationsClientGetResponse struct {
	// The resource type for defining a location of the containing resource provider. The location resource represents a logical
// location where the resource provider operates.
	LocationResource
}

// LocationsClientListResponse contains the response from method LocationsClient.NewListPager.
type LocationsClientListResponse struct {
	// The response of a LocationResource list operation.
	LocationResourceListResult
}

// PlanesClientListPlanesResponse contains the response from method PlanesClient.NewListPlanesPager.
type PlanesClientListPlanesResponse struct {
	// The response of a GenericPlaneResource list operation.
	GenericPlaneResourceListResult
}

// RadiusPlanesClientCreateOrUpdateResponse contains the response from method RadiusPlanesClient.BeginCreateOrUpdate.
type RadiusPlanesClientCreateOrUpdateResponse struct {
	// The Radius plane resource.
	RadiusPlaneResource
}

// RadiusPlanesClientDeleteResponse contains the response from method RadiusPlanesClient.BeginDelete.
type RadiusPlanesClientDeleteResponse struct {
	// placeholder for future response values
}

// RadiusPlanesClientGetResponse contains the response from method RadiusPlanesClient.Get.
type RadiusPlanesClientGetResponse struct {
	// The Radius plane resource.
	RadiusPlaneResource
}

// RadiusPlanesClientListResponse contains the response from method RadiusPlanesClient.NewListPager.
type RadiusPlanesClientListResponse struct {
	// The response of a RadiusPlaneResource list operation.
	RadiusPlaneResourceListResult
}

// RadiusPlanesClientUpdateResponse contains the response from method RadiusPlanesClient.BeginUpdate.
type RadiusPlanesClientUpdateResponse struct {
	// The Radius plane resource.
	RadiusPlaneResource
}

// ResourceGroupsClientCreateOrUpdateResponse contains the response from method ResourceGroupsClient.CreateOrUpdate.
type ResourceGroupsClientCreateOrUpdateResponse struct {
	// The resource group resource
	ResourceGroupResource
}

// ResourceGroupsClientDeleteResponse contains the response from method ResourceGroupsClient.Delete.
type ResourceGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceGroupsClientGetResponse contains the response from method ResourceGroupsClient.Get.
type ResourceGroupsClientGetResponse struct {
	// The resource group resource
	ResourceGroupResource
}

// ResourceGroupsClientListResponse contains the response from method ResourceGroupsClient.NewListPager.
type ResourceGroupsClientListResponse struct {
	// The response of a ResourceGroupResource list operation.
	ResourceGroupResourceListResult
}

// ResourceGroupsClientUpdateResponse contains the response from method ResourceGroupsClient.Update.
type ResourceGroupsClientUpdateResponse struct {
	// The resource group resource
	ResourceGroupResource
}

// ResourceProvidersClientCreateOrUpdateResponse contains the response from method ResourceProvidersClient.BeginCreateOrUpdate.
type ResourceProvidersClientCreateOrUpdateResponse struct {
	// The resource type for defining a resource provider.
	ResourceProviderResource
}

// ResourceProvidersClientDeleteResponse contains the response from method ResourceProvidersClient.BeginDelete.
type ResourceProvidersClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceProvidersClientGetProviderSummaryResponse contains the response from method ResourceProvidersClient.GetProviderSummary.
type ResourceProvidersClientGetProviderSummaryResponse struct {
	// The summary of a resource provider configuration. This type is optimized for querying resource providers and supported
// types.
	ResourceProviderSummary
}

// ResourceProvidersClientGetResponse contains the response from method ResourceProvidersClient.Get.
type ResourceProvidersClientGetResponse struct {
	// The resource type for defining a resource provider.
	ResourceProviderResource
}

// ResourceProvidersClientListProviderSummariesResponse contains the response from method ResourceProvidersClient.NewListProviderSummariesPager.
type ResourceProvidersClientListProviderSummariesResponse struct {
	// Paged collection of ResourceProviderSummary items
	PagedResourceProviderSummary
}

// ResourceProvidersClientListResponse contains the response from method ResourceProvidersClient.NewListPager.
type ResourceProvidersClientListResponse struct {
	// The response of a ResourceProviderResource list operation.
	ResourceProviderResourceListResult
}

// ResourceTypesClientCreateOrUpdateResponse contains the response from method ResourceTypesClient.BeginCreateOrUpdate.
type ResourceTypesClientCreateOrUpdateResponse struct {
	// The resource type for defining a resource type supported by the containing resource provider.
	ResourceTypeResource
}

// ResourceTypesClientDeleteResponse contains the response from method ResourceTypesClient.BeginDelete.
type ResourceTypesClientDeleteResponse struct {
	// placeholder for future response values
}

// ResourceTypesClientGetResponse contains the response from method ResourceTypesClient.Get.
type ResourceTypesClientGetResponse struct {
	// The resource type for defining a resource type supported by the containing resource provider.
	ResourceTypeResource
}

// ResourceTypesClientListResponse contains the response from method ResourceTypesClient.NewListPager.
type ResourceTypesClientListResponse struct {
	// The response of a ResourceTypeResource list operation.
	ResourceTypeResourceListResult
}

// ResourcesClientListResponse contains the response from method ResourcesClient.NewListPager.
type ResourcesClientListResponse struct {
	// The response of a GenericResource list operation.
	GenericResourceListResult
}

