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

// QualityProfileQualityItem struct for QualityProfileQualityItem
type QualityProfileQualityItem struct {
	Id *int32 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Quality *Quality `json:"quality,omitempty"`
	Items []*QualityProfileQualityItem `json:"items,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewQualityProfileQualityItem instantiates a new QualityProfileQualityItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualityProfileQualityItem() *QualityProfileQualityItem {
	this := QualityProfileQualityItem{}
	return &this
}

// NewQualityProfileQualityItemWithDefaults instantiates a new QualityProfileQualityItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualityProfileQualityItemWithDefaults() *QualityProfileQualityItem {
	this := QualityProfileQualityItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *QualityProfileQualityItem) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItem) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *QualityProfileQualityItem) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *QualityProfileQualityItem) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileQualityItem) GetName() string {
	if o == nil || isNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileQualityItem) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *QualityProfileQualityItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *QualityProfileQualityItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *QualityProfileQualityItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *QualityProfileQualityItem) UnsetName() {
	o.Name.Unset()
}

// GetQuality returns the Quality field value if set, zero value otherwise.
func (o *QualityProfileQualityItem) GetQuality() Quality {
	if o == nil || isNil(o.Quality) {
		var ret Quality
		return ret
	}
	return *o.Quality
}

// GetQualityOk returns a tuple with the Quality field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItem) GetQualityOk() (*Quality, bool) {
	if o == nil || isNil(o.Quality) {
    return nil, false
	}
	return o.Quality, true
}

// HasQuality returns a boolean if a field has been set.
func (o *QualityProfileQualityItem) HasQuality() bool {
	if o != nil && !isNil(o.Quality) {
		return true
	}

	return false
}

// SetQuality gets a reference to the given Quality and assigns it to the Quality field.
func (o *QualityProfileQualityItem) SetQuality(v Quality) {
	o.Quality = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *QualityProfileQualityItem) GetItems() []*QualityProfileQualityItem {
	if o == nil {
		var ret []*QualityProfileQualityItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *QualityProfileQualityItem) GetItemsOk() ([]*QualityProfileQualityItem, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *QualityProfileQualityItem) HasItems() bool {
	if o != nil && isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []QualityProfileQualityItem and assigns it to the Items field.
func (o *QualityProfileQualityItem) SetItems(v []*QualityProfileQualityItem) {
	o.Items = v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *QualityProfileQualityItem) GetAllowed() bool {
	if o == nil || isNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QualityProfileQualityItem) GetAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.Allowed) {
    return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *QualityProfileQualityItem) HasAllowed() bool {
	if o != nil && !isNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *QualityProfileQualityItem) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o QualityProfileQualityItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !isNil(o.Quality) {
		toSerialize["quality"] = o.Quality
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return json.Marshal(toSerialize)
}

type NullableQualityProfileQualityItem struct {
	value *QualityProfileQualityItem
	isSet bool
}

func (v NullableQualityProfileQualityItem) Get() *QualityProfileQualityItem {
	return v.value
}

func (v *NullableQualityProfileQualityItem) Set(val *QualityProfileQualityItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQualityProfileQualityItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQualityProfileQualityItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualityProfileQualityItem(val *QualityProfileQualityItem) *NullableQualityProfileQualityItem {
	return &NullableQualityProfileQualityItem{value: val, isSet: true}
}

func (v NullableQualityProfileQualityItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualityProfileQualityItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


