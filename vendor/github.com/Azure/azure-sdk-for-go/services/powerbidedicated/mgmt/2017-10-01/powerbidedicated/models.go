package powerbidedicated

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
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// The package's fully qualified name.
const fqdn = "github.com/Azure/azure-sdk-for-go/services/powerbidedicated/mgmt/2017-10-01/powerbidedicated"

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Paused ...
	Paused ProvisioningState = "Paused"
	// Pausing ...
	Pausing ProvisioningState = "Pausing"
	// Preparing ...
	Preparing ProvisioningState = "Preparing"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Resuming ...
	Resuming ProvisioningState = "Resuming"
	// Scaling ...
	Scaling ProvisioningState = "Scaling"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
	// Suspended ...
	Suspended ProvisioningState = "Suspended"
	// Suspending ...
	Suspending ProvisioningState = "Suspending"
	// Updating ...
	Updating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Deleting, Failed, Paused, Pausing, Preparing, Provisioning, Resuming, Scaling, Succeeded, Suspended, Suspending, Updating}
}

// SkuTier enumerates the values for sku tier.
type SkuTier string

const (
	// PBIEAzure ...
	PBIEAzure SkuTier = "PBIE_Azure"
)

// PossibleSkuTierValues returns an array of possible values for the SkuTier const type.
func PossibleSkuTierValues() []SkuTier {
	return []SkuTier{PBIEAzure}
}

// State enumerates the values for state.
type State string

const (
	// StateDeleting ...
	StateDeleting State = "Deleting"
	// StateFailed ...
	StateFailed State = "Failed"
	// StatePaused ...
	StatePaused State = "Paused"
	// StatePausing ...
	StatePausing State = "Pausing"
	// StatePreparing ...
	StatePreparing State = "Preparing"
	// StateProvisioning ...
	StateProvisioning State = "Provisioning"
	// StateResuming ...
	StateResuming State = "Resuming"
	// StateScaling ...
	StateScaling State = "Scaling"
	// StateSucceeded ...
	StateSucceeded State = "Succeeded"
	// StateSuspended ...
	StateSuspended State = "Suspended"
	// StateSuspending ...
	StateSuspending State = "Suspending"
	// StateUpdating ...
	StateUpdating State = "Updating"
)

// PossibleStateValues returns an array of possible values for the State const type.
func PossibleStateValues() []State {
	return []State{StateDeleting, StateFailed, StatePaused, StatePausing, StatePreparing, StateProvisioning, StateResuming, StateScaling, StateSucceeded, StateSuspended, StateSuspending, StateUpdating}
}

// CapacitiesCreateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CapacitiesCreateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CapacitiesCreateFuture) Result(client CapacitiesClient) (dc DedicatedCapacity, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesCreateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("powerbidedicated.CapacitiesCreateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if dc.Response.Response, err = future.GetResult(sender); err == nil && dc.Response.Response.StatusCode != http.StatusNoContent {
		dc, err = client.CreateResponder(dc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesCreateFuture", "Result", dc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CapacitiesDeleteFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CapacitiesDeleteFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CapacitiesDeleteFuture) Result(client CapacitiesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesDeleteFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("powerbidedicated.CapacitiesDeleteFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// CapacitiesResumeFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CapacitiesResumeFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CapacitiesResumeFuture) Result(client CapacitiesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesResumeFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("powerbidedicated.CapacitiesResumeFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// CapacitiesSuspendFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CapacitiesSuspendFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CapacitiesSuspendFuture) Result(client CapacitiesClient) (ar autorest.Response, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesSuspendFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("powerbidedicated.CapacitiesSuspendFuture")
		return
	}
	ar.Response = future.Response()
	return
}

// CapacitiesUpdateFuture an abstraction for monitoring and retrieving the results of a long-running
// operation.
type CapacitiesUpdateFuture struct {
	azure.Future
}

// Result returns the result of the asynchronous operation.
// If the operation has not completed it will return an error.
func (future *CapacitiesUpdateFuture) Result(client CapacitiesClient) (dc DedicatedCapacity, err error) {
	var done bool
	done, err = future.Done(client)
	if err != nil {
		err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesUpdateFuture", "Result", future.Response(), "Polling failure")
		return
	}
	if !done {
		err = azure.NewAsyncOpIncompleteError("powerbidedicated.CapacitiesUpdateFuture")
		return
	}
	sender := autorest.DecorateSender(client, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	if dc.Response.Response, err = future.GetResult(sender); err == nil && dc.Response.Response.StatusCode != http.StatusNoContent {
		dc, err = client.UpdateResponder(dc.Response.Response)
		if err != nil {
			err = autorest.NewErrorWithError(err, "powerbidedicated.CapacitiesUpdateFuture", "Result", dc.Response.Response, "Failure responding to request")
		}
	}
	return
}

// CheckCapacityNameAvailabilityParameters details of capacity name request body.
type CheckCapacityNameAvailabilityParameters struct {
	// Name - Name for checking availability.
	Name *string `json:"name,omitempty"`
	// Type - The resource type of powerbi dedicated.
	Type *string `json:"type,omitempty"`
}

// CheckCapacityNameAvailabilityResult the checking result of capacity name availibility.
type CheckCapacityNameAvailabilityResult struct {
	autorest.Response `json:"-"`
	// NameAvailable - Indicator of availablity of the capacity name.
	NameAvailable *bool `json:"nameAvailable,omitempty"`
	// Reason - The reason of unavailability.
	Reason *string `json:"reason,omitempty"`
	// Message - The detailed message of the request unavailability.
	Message *string `json:"message,omitempty"`
}

// DedicatedCapacities an array of Dedicated capacities resources.
type DedicatedCapacities struct {
	autorest.Response `json:"-"`
	// Value - An array of Dedicated capacities resources.
	Value *[]DedicatedCapacity `json:"value,omitempty"`
}

// DedicatedCapacity represents an instance of a Dedicated Capacity resource.
type DedicatedCapacity struct {
	autorest.Response `json:"-"`
	// DedicatedCapacityProperties - Properties of the provision operation request.
	*DedicatedCapacityProperties `json:"properties,omitempty"`
	// ID - An identifier that represents the PowerBI Dedicated resource.
	ID *string `json:"id,omitempty"`
	// Name - The name of the PowerBI Dedicated resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the PowerBI Dedicated resource.
	Type *string `json:"type,omitempty"`
	// Location - Location of the PowerBI Dedicated resource.
	Location *string `json:"location,omitempty"`
	// Sku - The SKU of the PowerBI Dedicated resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for DedicatedCapacity.
func (dc DedicatedCapacity) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dc.DedicatedCapacityProperties != nil {
		objectMap["properties"] = dc.DedicatedCapacityProperties
	}
	if dc.ID != nil {
		objectMap["id"] = dc.ID
	}
	if dc.Name != nil {
		objectMap["name"] = dc.Name
	}
	if dc.Type != nil {
		objectMap["type"] = dc.Type
	}
	if dc.Location != nil {
		objectMap["location"] = dc.Location
	}
	if dc.Sku != nil {
		objectMap["sku"] = dc.Sku
	}
	if dc.Tags != nil {
		objectMap["tags"] = dc.Tags
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DedicatedCapacity struct.
func (dc *DedicatedCapacity) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "properties":
			if v != nil {
				var dedicatedCapacityProperties DedicatedCapacityProperties
				err = json.Unmarshal(*v, &dedicatedCapacityProperties)
				if err != nil {
					return err
				}
				dc.DedicatedCapacityProperties = &dedicatedCapacityProperties
			}
		case "id":
			if v != nil {
				var ID string
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				dc.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				dc.Name = &name
			}
		case "type":
			if v != nil {
				var typeVar string
				err = json.Unmarshal(*v, &typeVar)
				if err != nil {
					return err
				}
				dc.Type = &typeVar
			}
		case "location":
			if v != nil {
				var location string
				err = json.Unmarshal(*v, &location)
				if err != nil {
					return err
				}
				dc.Location = &location
			}
		case "sku":
			if v != nil {
				var sku ResourceSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				dc.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				dc.Tags = tags
			}
		}
	}

	return nil
}

// DedicatedCapacityAdministrators an array of administrator user identities
type DedicatedCapacityAdministrators struct {
	// Members - An array of administrator user identities.
	Members *[]string `json:"members,omitempty"`
}

// DedicatedCapacityMutableProperties an object that represents a set of mutable Dedicated capacity
// resource properties.
type DedicatedCapacityMutableProperties struct {
	// Administration - A collection of Dedicated capacity administrators
	Administration *DedicatedCapacityAdministrators `json:"administration,omitempty"`
}

// DedicatedCapacityProperties properties of Dedicated Capacity resource.
type DedicatedCapacityProperties struct {
	// State - The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning. Possible values include: 'StateDeleting', 'StateSucceeded', 'StateFailed', 'StatePaused', 'StateSuspended', 'StateProvisioning', 'StateUpdating', 'StateSuspending', 'StatePausing', 'StateResuming', 'StatePreparing', 'StateScaling'
	State State `json:"state,omitempty"`
	// ProvisioningState - The current deployment state of PowerBI Dedicatedresource. The provisioningState is to indicate states for resource provisioning. Possible values include: 'Deleting', 'Succeeded', 'Failed', 'Paused', 'Suspended', 'Provisioning', 'Updating', 'Suspending', 'Pausing', 'Resuming', 'Preparing', 'Scaling'
	ProvisioningState ProvisioningState `json:"provisioningState,omitempty"`
	// Administration - A collection of Dedicated capacity administrators
	Administration *DedicatedCapacityAdministrators `json:"administration,omitempty"`
}

// DedicatedCapacityUpdateParameters provision request specification
type DedicatedCapacityUpdateParameters struct {
	// Sku - The SKU of the Dedicated capacity resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional provisioning properties.
	Tags map[string]*string `json:"tags"`
	// DedicatedCapacityMutableProperties - Properties of the provision operation request.
	*DedicatedCapacityMutableProperties `json:"properties,omitempty"`
}

// MarshalJSON is the custom marshaler for DedicatedCapacityUpdateParameters.
func (dcup DedicatedCapacityUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if dcup.Sku != nil {
		objectMap["sku"] = dcup.Sku
	}
	if dcup.Tags != nil {
		objectMap["tags"] = dcup.Tags
	}
	if dcup.DedicatedCapacityMutableProperties != nil {
		objectMap["properties"] = dcup.DedicatedCapacityMutableProperties
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for DedicatedCapacityUpdateParameters struct.
func (dcup *DedicatedCapacityUpdateParameters) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "sku":
			if v != nil {
				var sku ResourceSku
				err = json.Unmarshal(*v, &sku)
				if err != nil {
					return err
				}
				dcup.Sku = &sku
			}
		case "tags":
			if v != nil {
				var tags map[string]*string
				err = json.Unmarshal(*v, &tags)
				if err != nil {
					return err
				}
				dcup.Tags = tags
			}
		case "properties":
			if v != nil {
				var dedicatedCapacityMutableProperties DedicatedCapacityMutableProperties
				err = json.Unmarshal(*v, &dedicatedCapacityMutableProperties)
				if err != nil {
					return err
				}
				dcup.DedicatedCapacityMutableProperties = &dedicatedCapacityMutableProperties
			}
		}
	}

	return nil
}

// ErrorResponse describes the format of Error response.
type ErrorResponse struct {
	// Code - Error code
	Code *string `json:"code,omitempty"`
	// Message - Error message indicating why the operation failed.
	Message *string `json:"message,omitempty"`
}

// Operation capacities REST API operation.
type Operation struct {
	// Name - Operation name: {provider}/{resource}/{operation}.
	Name *string `json:"name,omitempty"`
	// Display - The object that represents the operation.
	Display *OperationDisplay `json:"display,omitempty"`
}

// OperationDisplay the object that represents the operation.
type OperationDisplay struct {
	// Provider - Service provider: Microsoft.PowerBIDedicated.
	Provider *string `json:"provider,omitempty"`
	// Resource - Resource on which the operation is performed: capacity, etc.
	Resource *string `json:"resource,omitempty"`
	// Operation - Operation type: create, update, delete, etc.
	Operation *string `json:"operation,omitempty"`
}

// OperationListResult result listing capacities. It contains a list of operations and a URL link to get
// the next set of results.
type OperationListResult struct {
	autorest.Response `json:"-"`
	// Value - List of capacities supported by the Microsoft.PowerBIDedicated resource provider.
	Value *[]Operation `json:"value,omitempty"`
	// NextLink - URL to get the next set of operation list results if there are any.
	NextLink *string `json:"nextLink,omitempty"`
}

// OperationListResultIterator provides access to a complete listing of Operation values.
type OperationListResultIterator struct {
	i    int
	page OperationListResultPage
}

// NextWithContext advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
func (iter *OperationListResultIterator) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultIterator.NextWithContext")
		defer func() {
			sc := -1
			if iter.Response().Response.Response != nil {
				sc = iter.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	iter.i++
	if iter.i < len(iter.page.Values()) {
		return nil
	}
	err = iter.page.NextWithContext(ctx)
	if err != nil {
		iter.i--
		return err
	}
	iter.i = 0
	return nil
}

// Next advances to the next value.  If there was an error making
// the request the iterator does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (iter *OperationListResultIterator) Next() error {
	return iter.NextWithContext(context.Background())
}

// NotDone returns true if the enumeration should be started or is not yet complete.
func (iter OperationListResultIterator) NotDone() bool {
	return iter.page.NotDone() && iter.i < len(iter.page.Values())
}

// Response returns the raw server response from the last page request.
func (iter OperationListResultIterator) Response() OperationListResult {
	return iter.page.Response()
}

// Value returns the current value or a zero-initialized value if the
// iterator has advanced beyond the end of the collection.
func (iter OperationListResultIterator) Value() Operation {
	if !iter.page.NotDone() {
		return Operation{}
	}
	return iter.page.Values()[iter.i]
}

// IsEmpty returns true if the ListResult contains no values.
func (olr OperationListResult) IsEmpty() bool {
	return olr.Value == nil || len(*olr.Value) == 0
}

// operationListResultPreparer prepares a request to retrieve the next set of results.
// It returns nil if no more results exist.
func (olr OperationListResult) operationListResultPreparer(ctx context.Context) (*http.Request, error) {
	if olr.NextLink == nil || len(to.String(olr.NextLink)) < 1 {
		return nil, nil
	}
	return autorest.Prepare((&http.Request{}).WithContext(ctx),
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(olr.NextLink)))
}

// OperationListResultPage contains a page of Operation values.
type OperationListResultPage struct {
	fn  func(context.Context, OperationListResult) (OperationListResult, error)
	olr OperationListResult
}

// NextWithContext advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
func (page *OperationListResultPage) NextWithContext(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/OperationListResultPage.NextWithContext")
		defer func() {
			sc := -1
			if page.Response().Response.Response != nil {
				sc = page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	next, err := page.fn(ctx, page.olr)
	if err != nil {
		return err
	}
	page.olr = next
	return nil
}

// Next advances to the next page of values.  If there was an error making
// the request the page does not advance and the error is returned.
// Deprecated: Use NextWithContext() instead.
func (page *OperationListResultPage) Next() error {
	return page.NextWithContext(context.Background())
}

// NotDone returns true if the page enumeration should be started or is not yet complete.
func (page OperationListResultPage) NotDone() bool {
	return !page.olr.IsEmpty()
}

// Response returns the raw server response from the last page request.
func (page OperationListResultPage) Response() OperationListResult {
	return page.olr
}

// Values returns the slice of values for the current page or nil if there are no values.
func (page OperationListResultPage) Values() []Operation {
	if page.olr.IsEmpty() {
		return nil
	}
	return *page.olr.Value
}

// Resource represents an instance of an PowerBI Dedicated resource.
type Resource struct {
	// ID - An identifier that represents the PowerBI Dedicated resource.
	ID *string `json:"id,omitempty"`
	// Name - The name of the PowerBI Dedicated resource.
	Name *string `json:"name,omitempty"`
	// Type - The type of the PowerBI Dedicated resource.
	Type *string `json:"type,omitempty"`
	// Location - Location of the PowerBI Dedicated resource.
	Location *string `json:"location,omitempty"`
	// Sku - The SKU of the PowerBI Dedicated resource.
	Sku *ResourceSku `json:"sku,omitempty"`
	// Tags - Key-value pairs of additional resource provisioning properties.
	Tags map[string]*string `json:"tags"`
}

// MarshalJSON is the custom marshaler for Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if r.ID != nil {
		objectMap["id"] = r.ID
	}
	if r.Name != nil {
		objectMap["name"] = r.Name
	}
	if r.Type != nil {
		objectMap["type"] = r.Type
	}
	if r.Location != nil {
		objectMap["location"] = r.Location
	}
	if r.Sku != nil {
		objectMap["sku"] = r.Sku
	}
	if r.Tags != nil {
		objectMap["tags"] = r.Tags
	}
	return json.Marshal(objectMap)
}

// ResourceSku represents the SKU name and Azure pricing tier for PowerBI Dedicated resource.
type ResourceSku struct {
	// Name - Name of the SKU level.
	Name *string `json:"name,omitempty"`
	// Tier - The name of the Azure pricing tier to which the SKU applies. Possible values include: 'PBIEAzure'
	Tier SkuTier `json:"tier,omitempty"`
}

// SkuDetailsForExistingResource an object that represents SKU details for existing resources
type SkuDetailsForExistingResource struct {
	// Sku - The SKU in SKU details for existing resources.
	Sku *ResourceSku `json:"sku,omitempty"`
}

// SkuEnumerationForExistingResourceResult an object that represents enumerating SKUs for existing
// resources
type SkuEnumerationForExistingResourceResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of available SKUs for existing resources
	Value *[]SkuDetailsForExistingResource `json:"value,omitempty"`
}

// SkuEnumerationForNewResourceResult an object that represents enumerating SKUs for new resources
type SkuEnumerationForNewResourceResult struct {
	autorest.Response `json:"-"`
	// Value - The collection of available SKUs for new resources
	Value *[]ResourceSku `json:"value,omitempty"`
}