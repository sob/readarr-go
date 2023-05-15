/*
Readarr

Readarr API docs

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"encoding/json"
)

// CustomFormat struct for CustomFormat
type CustomFormat struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	IncludeCustomFormatWhenRenaming *bool `json:"includeCustomFormatWhenRenaming,omitempty"`
	Specifications []*ICustomFormatSpecification `json:"specifications,omitempty"`
}

// NewCustomFormat instantiates a new CustomFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomFormat() *CustomFormat {
	this := CustomFormat{}
	return &this
}

// NewCustomFormatWithDefaults instantiates a new CustomFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomFormatWithDefaults() *CustomFormat {
	this := CustomFormat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomFormat) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormat) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomFormat) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CustomFormat) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormat) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormat) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CustomFormat) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CustomFormat) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CustomFormat) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CustomFormat) UnsetName() {
	o.Name.Unset()
}

// GetIncludeCustomFormatWhenRenaming returns the IncludeCustomFormatWhenRenaming field value if set, zero value otherwise.
func (o *CustomFormat) GetIncludeCustomFormatWhenRenaming() bool {
	if o == nil || isNil(o.IncludeCustomFormatWhenRenaming) {
		var ret bool
		return ret
	}
	return *o.IncludeCustomFormatWhenRenaming
}

// GetIncludeCustomFormatWhenRenamingOk returns a tuple with the IncludeCustomFormatWhenRenaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomFormat) GetIncludeCustomFormatWhenRenamingOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeCustomFormatWhenRenaming) {
    return nil, false
	}
	return o.IncludeCustomFormatWhenRenaming, true
}

// HasIncludeCustomFormatWhenRenaming returns a boolean if a field has been set.
func (o *CustomFormat) HasIncludeCustomFormatWhenRenaming() bool {
	if o != nil && !isNil(o.IncludeCustomFormatWhenRenaming) {
		return true
	}

	return false
}

// SetIncludeCustomFormatWhenRenaming gets a reference to the given bool and assigns it to the IncludeCustomFormatWhenRenaming field.
func (o *CustomFormat) SetIncludeCustomFormatWhenRenaming(v bool) {
	o.IncludeCustomFormatWhenRenaming = &v
}

// GetSpecifications returns the Specifications field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CustomFormat) GetSpecifications() []*ICustomFormatSpecification {
	if o == nil {
		var ret []*ICustomFormatSpecification
		return ret
	}
	return o.Specifications
}

// GetSpecificationsOk returns a tuple with the Specifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CustomFormat) GetSpecificationsOk() ([]*ICustomFormatSpecification, bool) {
	if o == nil || isNil(o.Specifications) {
    return nil, false
	}
	return o.Specifications, true
}

// HasSpecifications returns a boolean if a field has been set.
func (o *CustomFormat) HasSpecifications() bool {
	if o != nil && isNil(o.Specifications) {
		return true
	}

	return false
}

// SetSpecifications gets a reference to the given []ICustomFormatSpecification and assigns it to the Specifications field.
func (o *CustomFormat) SetSpecifications(v []*ICustomFormatSpecification) {
	o.Specifications = v
}

func (o CustomFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.IncludeCustomFormatWhenRenaming) {
		toSerialize["includeCustomFormatWhenRenaming"] = o.IncludeCustomFormatWhenRenaming
	}
	if o.Specifications != nil {
		toSerialize["specifications"] = o.Specifications
	}
	return json.Marshal(toSerialize)
}

type NullableCustomFormat struct {
	value *CustomFormat
	isSet bool
}

func (v NullableCustomFormat) Get() *CustomFormat {
	return v.value
}

func (v *NullableCustomFormat) Set(val *CustomFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomFormat(val *CustomFormat) *NullableCustomFormat {
	return &NullableCustomFormat{value: val, isSet: true}
}

func (v NullableCustomFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


