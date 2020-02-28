package insights

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// AnalyticsItemsClient is the composite Swagger for Application Insights Management Client
type AnalyticsItemsClient struct {
	BaseClient
}

// NewAnalyticsItemsClient creates an instance of the AnalyticsItemsClient client.
func NewAnalyticsItemsClient(subscriptionID string) AnalyticsItemsClient {
	return NewAnalyticsItemsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewAnalyticsItemsClientWithBaseURI creates an instance of the AnalyticsItemsClient client.
func NewAnalyticsItemsClientWithBaseURI(baseURI string, subscriptionID string) AnalyticsItemsClient {
	return AnalyticsItemsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Delete deletes a specific Analytics Items defined within an Application Insights component.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// scopePath - enum indicating if this item definition is owned by a specific user or is shared between all
// users with access to the Application Insights component.
// ID - the Id of a specific item defined in the Application Insights component
// name - the name of a specific item defined in the Application Insights component
func (client AnalyticsItemsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, ID string, name string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AnalyticsItemsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.AnalyticsItemsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName, scopePath, ID, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client AnalyticsItemsClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, ID string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"scopePath":         autorest.Encode("path", scopePath),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(ID) > 0 {
		queryParameters["id"] = autorest.Encode("query", ID)
	}
	if len(name) > 0 {
		queryParameters["name"] = autorest.Encode("query", name)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client AnalyticsItemsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client AnalyticsItemsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets a specific Analytics Items defined within an Application Insights component.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// scopePath - enum indicating if this item definition is owned by a specific user or is shared between all
// users with access to the Application Insights component.
// ID - the Id of a specific item defined in the Application Insights component
// name - the name of a specific item defined in the Application Insights component
func (client AnalyticsItemsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, ID string, name string) (result ApplicationInsightsComponentAnalyticsItem, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AnalyticsItemsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.AnalyticsItemsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName, scopePath, ID, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client AnalyticsItemsClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, ID string, name string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"scopePath":         autorest.Encode("path", scopePath),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(ID) > 0 {
		queryParameters["id"] = autorest.Encode("query", ID)
	}
	if len(name) > 0 {
		queryParameters["name"] = autorest.Encode("query", name)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client AnalyticsItemsClient) GetSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client AnalyticsItemsClient) GetResponder(resp *http.Response) (result ApplicationInsightsComponentAnalyticsItem, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of Analytics Items defined within an Application Insights component.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// scopePath - enum indicating if this item definition is owned by a specific user or is shared between all
// users with access to the Application Insights component.
// scope - enum indicating if this item definition is owned by a specific user or is shared between all users
// with access to the Application Insights component.
// typeParameter - enum indicating the type of the Analytics item.
// includeContent - flag indicating whether or not to return the content of each applicable item. If false,
// only return the item information.
func (client AnalyticsItemsClient) List(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, scope ItemScope, typeParameter ItemTypeParameter, includeContent *bool) (result ListApplicationInsightsComponentAnalyticsItem, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AnalyticsItemsClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.AnalyticsItemsClient", "List", err.Error())
	}

	req, err := client.ListPreparer(ctx, resourceGroupName, resourceName, scopePath, scope, typeParameter, includeContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AnalyticsItemsClient) ListPreparer(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, scope ItemScope, typeParameter ItemTypeParameter, includeContent *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"scopePath":         autorest.Encode("path", scopePath),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(string(scope)) > 0 {
		queryParameters["scope"] = autorest.Encode("query", scope)
	} else {
		queryParameters["scope"] = autorest.Encode("query", "shared")
	}
	if len(string(typeParameter)) > 0 {
		queryParameters["type"] = autorest.Encode("query", typeParameter)
	} else {
		queryParameters["type"] = autorest.Encode("query", "none")
	}
	if includeContent != nil {
		queryParameters["includeContent"] = autorest.Encode("query", *includeContent)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AnalyticsItemsClient) ListSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AnalyticsItemsClient) ListResponder(resp *http.Response) (result ListApplicationInsightsComponentAnalyticsItem, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Put adds or Updates a specific Analytics Item within an Application Insights component.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// scopePath - enum indicating if this item definition is owned by a specific user or is shared between all
// users with access to the Application Insights component.
// itemProperties - properties that need to be specified to create a new item and add it to an Application
// Insights component.
// overrideItem - flag indicating whether or not to force save an item. This allows overriding an item if it
// already exists.
func (client AnalyticsItemsClient) Put(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, itemProperties ApplicationInsightsComponentAnalyticsItem, overrideItem *bool) (result ApplicationInsightsComponentAnalyticsItem, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/AnalyticsItemsClient.Put")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.AnalyticsItemsClient", "Put", err.Error())
	}

	req, err := client.PutPreparer(ctx, resourceGroupName, resourceName, scopePath, itemProperties, overrideItem)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Put", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Put", resp, "Failure sending request")
		return
	}

	result, err = client.PutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.AnalyticsItemsClient", "Put", resp, "Failure responding to request")
	}

	return
}

// PutPreparer prepares the Put request.
func (client AnalyticsItemsClient) PutPreparer(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, itemProperties ApplicationInsightsComponentAnalyticsItem, overrideItem *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"scopePath":         autorest.Encode("path", scopePath),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if overrideItem != nil {
		queryParameters["overrideItem"] = autorest.Encode("query", *overrideItem)
	}

	itemProperties.Version = nil
	itemProperties.TimeCreated = nil
	itemProperties.TimeModified = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item", pathParameters),
		autorest.WithJSON(itemProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSender sends the Put request. The method will close the
// http.Response Body if it receives an error.
func (client AnalyticsItemsClient) PutSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), azure.DoRetryWithRegistration(client.Client))
	return autorest.SendWithSender(client, req, sd...)
}

// PutResponder handles the response to the Put request. The method always
// closes the http.Response Body.
func (client AnalyticsItemsClient) PutResponder(resp *http.Response) (result ApplicationInsightsComponentAnalyticsItem, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}