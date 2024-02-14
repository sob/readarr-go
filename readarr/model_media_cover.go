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

// MediaCover struct for MediaCover
type MediaCover struct {
	Url NullableString `json:"url,omitempty"`
	CoverType *MediaCoverTypes `json:"coverType,omitempty"`
	Extension NullableString `json:"extension,omitempty"`
}

// NewMediaCover instantiates a new MediaCover object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaCover() *MediaCover {
	this := MediaCover{}
	return &this
}

// NewMediaCoverWithDefaults instantiates a new MediaCover object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaCoverWithDefaults() *MediaCover {
	this := MediaCover{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaCover) GetUrl() string {
	if o == nil || isNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaCover) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MediaCover) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MediaCover) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MediaCover) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MediaCover) UnsetUrl() {
	o.Url.Unset()
}

// GetCoverType returns the CoverType field value if set, zero value otherwise.
func (o *MediaCover) GetCoverType() MediaCoverTypes {
	if o == nil || isNil(o.CoverType) {
		var ret MediaCoverTypes
		return ret
	}
	return *o.CoverType
}

// GetCoverTypeOk returns a tuple with the CoverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaCover) GetCoverTypeOk() (*MediaCoverTypes, bool) {
	if o == nil || isNil(o.CoverType) {
    return nil, false
	}
	return o.CoverType, true
}

// HasCoverType returns a boolean if a field has been set.
func (o *MediaCover) HasCoverType() bool {
	if o != nil && !isNil(o.CoverType) {
		return true
	}

	return false
}

// SetCoverType gets a reference to the given MediaCoverTypes and assigns it to the CoverType field.
func (o *MediaCover) SetCoverType(v MediaCoverTypes) {
	o.CoverType = &v
}

// GetExtension returns the Extension field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MediaCover) GetExtension() string {
	if o == nil || isNil(o.Extension.Get()) {
		var ret string
		return ret
	}
	return *o.Extension.Get()
}

// GetExtensionOk returns a tuple with the Extension field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MediaCover) GetExtensionOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Extension.Get(), o.Extension.IsSet()
}

// HasExtension returns a boolean if a field has been set.
func (o *MediaCover) HasExtension() bool {
	if o != nil && o.Extension.IsSet() {
		return true
	}

	return false
}

// SetExtension gets a reference to the given NullableString and assigns it to the Extension field.
func (o *MediaCover) SetExtension(v string) {
	o.Extension.Set(&v)
}
// SetExtensionNil sets the value for Extension to be an explicit nil
func (o *MediaCover) SetExtensionNil() {
	o.Extension.Set(nil)
}

// UnsetExtension ensures that no value is present for Extension, not even an explicit nil
func (o *MediaCover) UnsetExtension() {
	o.Extension.Unset()
}

func (o MediaCover) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if !isNil(o.CoverType) {
		toSerialize["coverType"] = o.CoverType
	}
	if o.Extension.IsSet() {
		toSerialize["extension"] = o.Extension.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMediaCover struct {
	value *MediaCover
	isSet bool
}

func (v NullableMediaCover) Get() *MediaCover {
	return v.value
}

func (v *NullableMediaCover) Set(val *MediaCover) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaCover) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaCover) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaCover(val *MediaCover) *NullableMediaCover {
	return &NullableMediaCover{value: val, isSet: true}
}

func (v NullableMediaCover) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaCover) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


