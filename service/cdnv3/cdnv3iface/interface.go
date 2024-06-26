// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cdnv3iface provides an interface to enable mocking the cdnv3 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cdnv3iface

import (
	"github.com/KscSDK/ksc-sdk-go/service/cdnv3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// Cdnv3API provides an interface to enable mocking the
// cdnv3.Cdnv3 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// cdnv3.
//	func myFunc(svc cdnv3iface.Cdnv3API) bool {
//	    // Make svc.BlockDomainUrlPost request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cdnv3.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCdnv3Client struct {
//	    cdnv3iface.Cdnv3API
//	}
//	func (m *mockCdnv3Client) BlockDomainUrlPost(input *map[string]interface{}) (*map[string]interface{}, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCdnv3Client{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Cdnv3API interface {
	BlockDomainUrlPost(*map[string]interface{}) (*map[string]interface{}, error)
	BlockDomainUrlPostWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	BlockDomainUrlPostRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ Cdnv3API = (*cdnv3.Cdnv3)(nil)
