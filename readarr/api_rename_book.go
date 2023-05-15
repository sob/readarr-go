/*
Readarr

Readarr API docs

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package readarr

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// RenameBookAPIService RenameBookAPI service
type RenameBookAPIService service
type ApiListRenameRequest struct {
	ctx context.Context
	ApiService *RenameBookAPIService
	authorId *int32
	bookId *int32
}

func (r ApiListRenameRequest) AuthorId(authorId int32) ApiListRenameRequest {
	r.authorId = &authorId
	return r
}

func (r ApiListRenameRequest) BookId(bookId int32) ApiListRenameRequest {
	r.bookId = &bookId
	return r
}

func (r ApiListRenameRequest) Execute() ([]*RenameBookResource, *http.Response, error) {
	return r.ApiService.ListRenameExecute(r)
}

/*
ListRename Method for ListRename

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRenameRequest
*/
func (a *RenameBookAPIService) ListRename(ctx context.Context) ApiListRenameRequest {
	return ApiListRenameRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []RenameBookResource
func (a *RenameBookAPIService) ListRenameExecute(r ApiListRenameRequest) ([]*RenameBookResource, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []*RenameBookResource
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RenameBookAPIService.ListRename")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/rename"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.authorId != nil {
		localVarQueryParams.Add("authorId", parameterToString(*r.authorId, ""))
	}
	if r.bookId != nil {
		localVarQueryParams.Add("bookId", parameterToString(*r.bookId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
