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

// checks if the SeriesBookLinkListLazyLoaded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesBookLinkListLazyLoaded{}

// SeriesBookLinkListLazyLoaded struct for SeriesBookLinkListLazyLoaded
type SeriesBookLinkListLazyLoaded struct {
	Value []SeriesBookLink `json:"value,omitempty"`
	IsLoaded *bool `json:"isLoaded,omitempty"`
}

// NewSeriesBookLinkListLazyLoaded instantiates a new SeriesBookLinkListLazyLoaded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesBookLinkListLazyLoaded() *SeriesBookLinkListLazyLoaded {
	this := SeriesBookLinkListLazyLoaded{}
	return &this
}

// NewSeriesBookLinkListLazyLoadedWithDefaults instantiates a new SeriesBookLinkListLazyLoaded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesBookLinkListLazyLoadedWithDefaults() *SeriesBookLinkListLazyLoaded {
	this := SeriesBookLinkListLazyLoaded{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SeriesBookLinkListLazyLoaded) GetValue() []SeriesBookLink {
	if o == nil {
		var ret []SeriesBookLink
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SeriesBookLinkListLazyLoaded) GetValueOk() ([]SeriesBookLink, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SeriesBookLinkListLazyLoaded) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []SeriesBookLink and assigns it to the Value field.
func (o *SeriesBookLinkListLazyLoaded) SetValue(v []SeriesBookLink) {
	o.Value = v
}

// GetIsLoaded returns the IsLoaded field value if set, zero value otherwise.
func (o *SeriesBookLinkListLazyLoaded) GetIsLoaded() bool {
	if o == nil || IsNil(o.IsLoaded) {
		var ret bool
		return ret
	}
	return *o.IsLoaded
}

// GetIsLoadedOk returns a tuple with the IsLoaded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesBookLinkListLazyLoaded) GetIsLoadedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsLoaded) {
		return nil, false
	}
	return o.IsLoaded, true
}

// HasIsLoaded returns a boolean if a field has been set.
func (o *SeriesBookLinkListLazyLoaded) HasIsLoaded() bool {
	if o != nil && !IsNil(o.IsLoaded) {
		return true
	}

	return false
}

// SetIsLoaded gets a reference to the given bool and assigns it to the IsLoaded field.
func (o *SeriesBookLinkListLazyLoaded) SetIsLoaded(v bool) {
	o.IsLoaded = &v
}

func (o SeriesBookLinkListLazyLoaded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesBookLinkListLazyLoaded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.IsLoaded) {
		toSerialize["isLoaded"] = o.IsLoaded
	}
	return toSerialize, nil
}

type NullableSeriesBookLinkListLazyLoaded struct {
	value *SeriesBookLinkListLazyLoaded
	isSet bool
}

func (v NullableSeriesBookLinkListLazyLoaded) Get() *SeriesBookLinkListLazyLoaded {
	return v.value
}

func (v *NullableSeriesBookLinkListLazyLoaded) Set(val *SeriesBookLinkListLazyLoaded) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesBookLinkListLazyLoaded) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesBookLinkListLazyLoaded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesBookLinkListLazyLoaded(val *SeriesBookLinkListLazyLoaded) *NullableSeriesBookLinkListLazyLoaded {
	return &NullableSeriesBookLinkListLazyLoaded{value: val, isSet: true}
}

func (v NullableSeriesBookLinkListLazyLoaded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesBookLinkListLazyLoaded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


