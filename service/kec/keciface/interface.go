// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package keciface provides an interface to enable mocking the kec service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package keciface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/ksc/ksc-sdk-go/service/kec"
)

// KecAPI provides an interface to enable mocking the
// kec.Kec service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // kec.
//    func myFunc(svc keciface.KecAPI) bool {
//        // Make svc.AddVmIntoDataGuard request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kec.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKecClient struct {
//        keciface.KecAPI
//    }
//    func (m *mockKecClient) AddVmIntoDataGuard(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKecClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KecAPI interface {
	AddVmIntoDataGuard(*map[string]interface{}) (*map[string]interface{}, error)
	AddVmIntoDataGuardWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddVmIntoDataGuardRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachKey(*map[string]interface{}) (*map[string]interface{}, error)
	AttachKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AttachNetworkInterface(*map[string]interface{}) (*map[string]interface{}, error)
	AttachNetworkInterfaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AttachNetworkInterfaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CopyImage(*map[string]interface{}) (*map[string]interface{}, error)
	CopyImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CopyImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateDataGuardGroup(*map[string]interface{}) (*map[string]interface{}, error)
	CreateDataGuardGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateDataGuardGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateImage(*map[string]interface{}) (*map[string]interface{}, error)
	CreateImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateLocalVolumeSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	CreateLocalVolumeSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateLocalVolumeSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteDataGuardGroups(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteDataGuardGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteDataGuardGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteLocalVolumeSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteLocalVolumeSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteLocalVolumeSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeAvailabilityZones(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeAvailabilityZonesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeAvailabilityZonesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDataGuardCapacity(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDataGuardCapacityWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDataGuardCapacityRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeDataGuardGroup(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeDataGuardGroupWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeDataGuardGroupRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImageSharePermission(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImageSharePermissionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImageSharePermissionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeImages(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeImagesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeImagesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceFamilys(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceFamilysWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceFamilysRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceTypeConfigs(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceTypeConfigsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceTypeConfigsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstanceVnc(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstanceVncWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstanceVncRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeInstances(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLocalVolumeSnapshots(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLocalVolumeSnapshotsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLocalVolumeSnapshotsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeLocalVolumes(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeLocalVolumesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeLocalVolumesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachKey(*map[string]interface{}) (*map[string]interface{}, error)
	DetachKeyWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachKeyRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DetachNetworkInterface(*map[string]interface{}) (*map[string]interface{}, error)
	DetachNetworkInterfaceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DetachNetworkInterfaceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ImportImage(*map[string]interface{}) (*map[string]interface{}, error)
	ImportImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ImportImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyDataGuardGroups(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyDataGuardGroupsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyDataGuardGroupsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyImageAttribute(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyImageAttributeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyImageAttributeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyImageSharePermission(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyImageSharePermissionWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyImageSharePermissionRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceAttribute(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceAttributeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceAttributeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceImage(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceImageWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceImageRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyInstanceType(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyInstanceTypeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyInstanceTypeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ModifyNetworkInterfaceAttribute(*map[string]interface{}) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ModifyNetworkInterfaceAttributeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RebootInstances(*map[string]interface{}) (*map[string]interface{}, error)
	RebootInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RebootInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveImages(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveImagesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveImagesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RemoveVmFromDataGuard(*map[string]interface{}) (*map[string]interface{}, error)
	RemoveVmFromDataGuardWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RemoveVmFromDataGuardRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RollbackLocalVolume(*map[string]interface{}) (*map[string]interface{}, error)
	RollbackLocalVolumeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RollbackLocalVolumeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RunInstances(*map[string]interface{}) (*map[string]interface{}, error)
	RunInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RunInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StartInstances(*map[string]interface{}) (*map[string]interface{}, error)
	StartInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StartInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	StopInstances(*map[string]interface{}) (*map[string]interface{}, error)
	StopInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	StopInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	TerminateInstances(*map[string]interface{}) (*map[string]interface{}, error)
	TerminateInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	TerminateInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ KecAPI = (*kec.Kec)(nil)