/*
Readarr

Readarr API docs

API version: v0.3.18.2411
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// SeriesResource struct for SeriesResource
type SeriesResource struct {
	Id *int32 `json:"id,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Links []*SeriesBookLinkResource `json:"links,omitempty"`
}

// NewSeriesResource instantiates a new SeriesResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesResource() *SeriesResource {
	this := SeriesResource{}
	return &this
}

// NewSeriesResourceWithDefaults instantiates a new SeriesResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesResourceWithDefaults() *SeriesResource {
	this := SeriesResource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SeriesResource) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesResource) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SeriesResource) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SeriesResource) SetId(v int32) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesResource) GetTitle() string {
	if o == nil || isNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesResource) GetTitleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *SeriesResource) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *SeriesResource) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *SeriesResource) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *SeriesResource) UnsetTitle() {
	o.Title.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesResource) GetDescription() string {
	if o == nil || isNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesResource) GetDescriptionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *SeriesResource) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *SeriesResource) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *SeriesResource) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *SeriesResource) UnsetDescription() {
	o.Description.Unset()
}

// GetLinks returns the Links field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesResource) GetLinks() []*SeriesBookLinkResource {
	if o == nil {
		var ret []*SeriesBookLinkResource
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesResource) GetLinksOk() ([]*SeriesBookLinkResource, bool) {
	if o == nil || isNil(o.Links) {
    return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SeriesResource) HasLinks() bool {
	if o != nil && isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given []SeriesBookLinkResource and assigns it to the Links field.
func (o *SeriesResource) SetLinks(v []*SeriesBookLinkResource) {
	o.Links = v
}

func (o SeriesResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableSeriesResource struct {
	value *SeriesResource
	isSet bool
}

func (v NullableSeriesResource) Get() *SeriesResource {
	return v.value
}

func (v *NullableSeriesResource) Set(val *SeriesResource) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesResource) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesResource(val *SeriesResource) *NullableSeriesResource {
	return &NullableSeriesResource{value: val, isSet: true}
}

func (v NullableSeriesResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


