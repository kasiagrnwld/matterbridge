// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

//
type WorkbookFunctionsPermutationaRequestBuilder struct{ BaseRequestBuilder }

// Permutationa action undocumented
func (b *WorkbookFunctionsRequestBuilder) Permutationa(reqObj *WorkbookFunctionsPermutationaRequestParameter) *WorkbookFunctionsPermutationaRequestBuilder {
	bb := &WorkbookFunctionsPermutationaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/permutationa"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFunctionsPermutationaRequest struct{ BaseRequest }

//
func (b *WorkbookFunctionsPermutationaRequestBuilder) Request() *WorkbookFunctionsPermutationaRequest {
	return &WorkbookFunctionsPermutationaRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFunctionsPermutationaRequest) Post(ctx context.Context) (resObj *WorkbookFunctionResult, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}