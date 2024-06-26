// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cdnv3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opBlockDomainUrlPost = "BlockDomainUrl"

// BlockDomainUrlPostRequest generates a "ksc/request.Request" representing the
// client's request for the BlockDomainUrlPost operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See BlockDomainUrlPost for more information on using the BlockDomainUrlPost
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//	// Example sending a request using the BlockDomainUrlPostRequest method.
//	req, resp := client.BlockDomainUrlPostRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2019-06-01/BlockDomainUrlPost
func (c *Cdnv3) BlockDomainUrlPostRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opBlockDomainUrlPost,
		HTTPMethod: "POST",
		HTTPPath:   "/2019-06-01/content/BlockDomainUrl",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// BlockDomainUrlPost API operation for cdnv3.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for cdnv3's
// API operation BlockDomainUrlPost for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/cdn-2019-06-01/BlockDomainUrlPost
func (c *Cdnv3) BlockDomainUrlPost(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.BlockDomainUrlPostRequest(input)
	return out, req.Send()
}

// BlockDomainUrlPostWithContext is the same as BlockDomainUrlPost with the addition of
// the ability to pass a context and additional request options.
//
// See BlockDomainUrlPost for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Cdnv3) BlockDomainUrlPostWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.BlockDomainUrlPostRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}
