// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sqlserver

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

const opCloneSecurityGroup = "CloneSecurityGroup"

// CloneSecurityGroupRequest generates a "ksc/request.Request" representing the
// client's request for the CloneSecurityGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CloneSecurityGroup for more information on using the CloneSecurityGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CloneSecurityGroupRequest method.
//    req, resp := client.CloneSecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CloneSecurityGroup
func (c *Sqlserver) CloneSecurityGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCloneSecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CloneSecurityGroup API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation CloneSecurityGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CloneSecurityGroup
func (c *Sqlserver) CloneSecurityGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CloneSecurityGroupRequest(input)
	return out, req.Send()
}

// CloneSecurityGroupWithContext is the same as CloneSecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CloneSecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) CloneSecurityGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CloneSecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBInstance = "CreateDBInstance"

// CreateDBInstanceRequest generates a "ksc/request.Request" representing the
// client's request for the CreateDBInstance operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateDBInstance for more information on using the CreateDBInstance
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateDBInstanceRequest method.
//    req, resp := client.CreateDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CreateDBInstance
func (c *Sqlserver) CreateDBInstanceRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateDBInstance API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation CreateDBInstance for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CreateDBInstance
func (c *Sqlserver) CreateDBInstance(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceRequest(input)
	return out, req.Send()
}

// CreateDBInstanceWithContext is the same as CreateDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) CreateDBInstanceWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateSecurityGroup = "CreateSecurityGroup"

// CreateSecurityGroupRequest generates a "ksc/request.Request" representing the
// client's request for the CreateSecurityGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateSecurityGroup for more information on using the CreateSecurityGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateSecurityGroupRequest method.
//    req, resp := client.CreateSecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CreateSecurityGroup
func (c *Sqlserver) CreateSecurityGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateSecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// CreateSecurityGroup API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation CreateSecurityGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/CreateSecurityGroup
func (c *Sqlserver) CreateSecurityGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateSecurityGroupRequest(input)
	return out, req.Send()
}

// CreateSecurityGroupWithContext is the same as CreateSecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateSecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) CreateSecurityGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateSecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDBInstance = "DeleteDBInstance"

// DeleteDBInstanceRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteDBInstance operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteDBInstance for more information on using the DeleteDBInstance
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteDBInstanceRequest method.
//    req, resp := client.DeleteDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DeleteDBInstance
func (c *Sqlserver) DeleteDBInstanceRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDBInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteDBInstance API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation DeleteDBInstance for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DeleteDBInstance
func (c *Sqlserver) DeleteDBInstance(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDBInstanceRequest(input)
	return out, req.Send()
}

// DeleteDBInstanceWithContext is the same as DeleteDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) DeleteDBInstanceWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteSecurityGroup = "DeleteSecurityGroup"

// DeleteSecurityGroupRequest generates a "ksc/request.Request" representing the
// client's request for the DeleteSecurityGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteSecurityGroup for more information on using the DeleteSecurityGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteSecurityGroupRequest method.
//    req, resp := client.DeleteSecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DeleteSecurityGroup
func (c *Sqlserver) DeleteSecurityGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteSecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteSecurityGroup API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation DeleteSecurityGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DeleteSecurityGroup
func (c *Sqlserver) DeleteSecurityGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteSecurityGroupRequest(input)
	return out, req.Send()
}

// DeleteSecurityGroupWithContext is the same as DeleteSecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) DeleteSecurityGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteSecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstances = "DescribeDBInstances"

// DescribeDBInstancesRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeDBInstances operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeDBInstances for more information on using the DescribeDBInstances
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeDBInstancesRequest method.
//    req, resp := client.DescribeDBInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DescribeDBInstances
func (c *Sqlserver) DescribeDBInstancesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDBInstances API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation DescribeDBInstances for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DescribeDBInstances
func (c *Sqlserver) DescribeDBInstances(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	return out, req.Send()
}

// DescribeDBInstancesWithContext is the same as DescribeDBInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) DescribeDBInstancesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSecurityGroup = "DescribeSecurityGroup"

// DescribeSecurityGroupRequest generates a "ksc/request.Request" representing the
// client's request for the DescribeSecurityGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DescribeSecurityGroup for more information on using the DescribeSecurityGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DescribeSecurityGroupRequest method.
//    req, resp := client.DescribeSecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DescribeSecurityGroup
func (c *Sqlserver) DescribeSecurityGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSecurityGroup API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation DescribeSecurityGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/DescribeSecurityGroup
func (c *Sqlserver) DescribeSecurityGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSecurityGroupRequest(input)
	return out, req.Send()
}

// DescribeSecurityGroupWithContext is the same as DescribeSecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) DescribeSecurityGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstance = "ModifyDBInstance"

// ModifyDBInstanceRequest generates a "ksc/request.Request" representing the
// client's request for the ModifyDBInstance operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifyDBInstance for more information on using the ModifyDBInstance
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifyDBInstanceRequest method.
//    req, resp := client.ModifyDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifyDBInstance
func (c *Sqlserver) ModifyDBInstanceRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstance,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyDBInstance API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation ModifyDBInstance for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifyDBInstance
func (c *Sqlserver) ModifyDBInstance(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceWithContext is the same as ModifyDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) ModifyDBInstanceWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySecurityGroup = "ModifySecurityGroup"

// ModifySecurityGroupRequest generates a "ksc/request.Request" representing the
// client's request for the ModifySecurityGroup operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifySecurityGroup for more information on using the ModifySecurityGroup
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifySecurityGroupRequest method.
//    req, resp := client.ModifySecurityGroupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroup
func (c *Sqlserver) ModifySecurityGroupRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySecurityGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroup API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation ModifySecurityGroup for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroup
func (c *Sqlserver) ModifySecurityGroup(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupWithContext is the same as ModifySecurityGroup with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) ModifySecurityGroupWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySecurityGroupRule = "ModifySecurityGroupRule"

// ModifySecurityGroupRuleRequest generates a "ksc/request.Request" representing the
// client's request for the ModifySecurityGroupRule operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifySecurityGroupRule for more information on using the ModifySecurityGroupRule
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifySecurityGroupRuleRequest method.
//    req, resp := client.ModifySecurityGroupRuleRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroupRule
func (c *Sqlserver) ModifySecurityGroupRuleRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySecurityGroupRule,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroupRule API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation ModifySecurityGroupRule for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroupRule
func (c *Sqlserver) ModifySecurityGroupRule(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupRuleWithContext is the same as ModifySecurityGroupRule with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupRule for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) ModifySecurityGroupRuleWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySecurityGroupRuleName = "ModifySecurityGroupRuleName"

// ModifySecurityGroupRuleNameRequest generates a "ksc/request.Request" representing the
// client's request for the ModifySecurityGroupRuleName operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ModifySecurityGroupRuleName for more information on using the ModifySecurityGroupRuleName
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ModifySecurityGroupRuleNameRequest method.
//    req, resp := client.ModifySecurityGroupRuleNameRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroupRuleName
func (c *Sqlserver) ModifySecurityGroupRuleNameRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySecurityGroupRuleName,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySecurityGroupRuleName API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation ModifySecurityGroupRuleName for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/ModifySecurityGroupRuleName
func (c *Sqlserver) ModifySecurityGroupRuleName(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleNameRequest(input)
	return out, req.Send()
}

// ModifySecurityGroupRuleNameWithContext is the same as ModifySecurityGroupRuleName with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySecurityGroupRuleName for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) ModifySecurityGroupRuleNameWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySecurityGroupRuleNameRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSecurityGroupRelation = "SecurityGroupRelation"

// SecurityGroupRelationRequest generates a "ksc/request.Request" representing the
// client's request for the SecurityGroupRelation operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SecurityGroupRelation for more information on using the SecurityGroupRelation
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SecurityGroupRelationRequest method.
//    req, resp := client.SecurityGroupRelationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/SecurityGroupRelation
func (c *Sqlserver) SecurityGroupRelationRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSecurityGroupRelation,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// SecurityGroupRelation API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation SecurityGroupRelation for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/SecurityGroupRelation
func (c *Sqlserver) SecurityGroupRelation(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SecurityGroupRelationRequest(input)
	return out, req.Send()
}

// SecurityGroupRelationWithContext is the same as SecurityGroupRelation with the addition of
// the ability to pass a context and additional request options.
//
// See SecurityGroupRelation for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) SecurityGroupRelationWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SecurityGroupRelationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStatisticDBInstances = "StatisticDBInstances"

// StatisticDBInstancesRequest generates a "ksc/request.Request" representing the
// client's request for the StatisticDBInstances operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See StatisticDBInstances for more information on using the StatisticDBInstances
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the StatisticDBInstancesRequest method.
//    req, resp := client.StatisticDBInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/StatisticDBInstances
func (c *Sqlserver) StatisticDBInstancesRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStatisticDBInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// StatisticDBInstances API operation for sqlserver.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the KSC API reference guide for sqlserver's
// API operation StatisticDBInstances for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqlserver-2019-04-25/StatisticDBInstances
func (c *Sqlserver) StatisticDBInstances(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StatisticDBInstancesRequest(input)
	return out, req.Send()
}

// StatisticDBInstancesWithContext is the same as StatisticDBInstances with the addition of
// the ability to pass a context and additional request options.
//
// See StatisticDBInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Sqlserver) StatisticDBInstancesWithContext(ctx aws.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StatisticDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}