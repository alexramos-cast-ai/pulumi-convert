// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeRegionResizeRequest struct {
	pulumi.CustomResourceState

	ComputeRegionResizeRequestId pulumi.StringOutput `pulumi:"computeRegionResizeRequestId"`
	// The creation timestamp for this resize request in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resize-request.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The reference of the regional instance group manager this ResizeRequest is a part of.
	InstanceGroupManager pulumi.StringOutput `pulumi:"instanceGroupManager"`
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
	// be deleted.
	RequestedRunDuration ComputeRegionResizeRequestRequestedRunDurationPtrOutput `pulumi:"requestedRunDuration"`
	// The number of instances to be created by this resize request. The group's target size will be increased by this number.
	ResizeBy pulumi.Float64Output `pulumi:"resizeBy"`
	// Current state of the request.
	State pulumi.StringOutput `pulumi:"state"`
	// Status of the request.
	Statuses ComputeRegionResizeRequestStatusArrayOutput `pulumi:"statuses"`
	Timeouts ComputeRegionResizeRequestTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeRegionResizeRequest registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionResizeRequest(ctx *pulumi.Context,
	name string, args *ComputeRegionResizeRequestArgs, opts ...pulumi.ResourceOption) (*ComputeRegionResizeRequest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceGroupManager == nil {
		return nil, errors.New("invalid value for required argument 'InstanceGroupManager'")
	}
	if args.ResizeBy == nil {
		return nil, errors.New("invalid value for required argument 'ResizeBy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRegionResizeRequest
	err = ctx.RegisterPackageResource("google-beta:index/computeRegionResizeRequest:ComputeRegionResizeRequest", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionResizeRequest gets an existing ComputeRegionResizeRequest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionResizeRequest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionResizeRequestState, opts ...pulumi.ResourceOption) (*ComputeRegionResizeRequest, error) {
	var resource ComputeRegionResizeRequest
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeRegionResizeRequest:ComputeRegionResizeRequest", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionResizeRequest resources.
type computeRegionResizeRequestState struct {
	ComputeRegionResizeRequestId *string `pulumi:"computeRegionResizeRequestId"`
	// The creation timestamp for this resize request in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resize-request.
	Description *string `pulumi:"description"`
	// The reference of the regional instance group manager this ResizeRequest is a part of.
	InstanceGroupManager *string `pulumi:"instanceGroupManager"`
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
	// be deleted.
	RequestedRunDuration *ComputeRegionResizeRequestRequestedRunDuration `pulumi:"requestedRunDuration"`
	// The number of instances to be created by this resize request. The group's target size will be increased by this number.
	ResizeBy *float64 `pulumi:"resizeBy"`
	// Current state of the request.
	State *string `pulumi:"state"`
	// Status of the request.
	Statuses []ComputeRegionResizeRequestStatus  `pulumi:"statuses"`
	Timeouts *ComputeRegionResizeRequestTimeouts `pulumi:"timeouts"`
}

type ComputeRegionResizeRequestState struct {
	ComputeRegionResizeRequestId pulumi.StringPtrInput
	// The creation timestamp for this resize request in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resize-request.
	Description pulumi.StringPtrInput
	// The reference of the regional instance group manager this ResizeRequest is a part of.
	InstanceGroupManager pulumi.StringPtrInput
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
	// be deleted.
	RequestedRunDuration ComputeRegionResizeRequestRequestedRunDurationPtrInput
	// The number of instances to be created by this resize request. The group's target size will be increased by this number.
	ResizeBy pulumi.Float64PtrInput
	// Current state of the request.
	State pulumi.StringPtrInput
	// Status of the request.
	Statuses ComputeRegionResizeRequestStatusArrayInput
	Timeouts ComputeRegionResizeRequestTimeoutsPtrInput
}

func (ComputeRegionResizeRequestState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionResizeRequestState)(nil)).Elem()
}

type computeRegionResizeRequestArgs struct {
	ComputeRegionResizeRequestId *string `pulumi:"computeRegionResizeRequestId"`
	// An optional description of this resize-request.
	Description *string `pulumi:"description"`
	// The reference of the regional instance group manager this ResizeRequest is a part of.
	InstanceGroupManager string `pulumi:"instanceGroupManager"`
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
	// be deleted.
	RequestedRunDuration *ComputeRegionResizeRequestRequestedRunDuration `pulumi:"requestedRunDuration"`
	// The number of instances to be created by this resize request. The group's target size will be increased by this number.
	ResizeBy float64                             `pulumi:"resizeBy"`
	Timeouts *ComputeRegionResizeRequestTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeRegionResizeRequest resource.
type ComputeRegionResizeRequestArgs struct {
	ComputeRegionResizeRequestId pulumi.StringPtrInput
	// An optional description of this resize-request.
	Description pulumi.StringPtrInput
	// The reference of the regional instance group manager this ResizeRequest is a part of.
	InstanceGroupManager pulumi.StringInput
	// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
	// be deleted.
	RequestedRunDuration ComputeRegionResizeRequestRequestedRunDurationPtrInput
	// The number of instances to be created by this resize request. The group's target size will be increased by this number.
	ResizeBy pulumi.Float64Input
	Timeouts ComputeRegionResizeRequestTimeoutsPtrInput
}

func (ComputeRegionResizeRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionResizeRequestArgs)(nil)).Elem()
}

type ComputeRegionResizeRequestInput interface {
	pulumi.Input

	ToComputeRegionResizeRequestOutput() ComputeRegionResizeRequestOutput
	ToComputeRegionResizeRequestOutputWithContext(ctx context.Context) ComputeRegionResizeRequestOutput
}

func (*ComputeRegionResizeRequest) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionResizeRequest)(nil)).Elem()
}

func (i *ComputeRegionResizeRequest) ToComputeRegionResizeRequestOutput() ComputeRegionResizeRequestOutput {
	return i.ToComputeRegionResizeRequestOutputWithContext(context.Background())
}

func (i *ComputeRegionResizeRequest) ToComputeRegionResizeRequestOutputWithContext(ctx context.Context) ComputeRegionResizeRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionResizeRequestOutput)
}

type ComputeRegionResizeRequestOutput struct{ *pulumi.OutputState }

func (ComputeRegionResizeRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionResizeRequest)(nil)).Elem()
}

func (o ComputeRegionResizeRequestOutput) ToComputeRegionResizeRequestOutput() ComputeRegionResizeRequestOutput {
	return o
}

func (o ComputeRegionResizeRequestOutput) ToComputeRegionResizeRequestOutputWithContext(ctx context.Context) ComputeRegionResizeRequestOutput {
	return o
}

func (o ComputeRegionResizeRequestOutput) ComputeRegionResizeRequestId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.ComputeRegionResizeRequestId }).(pulumi.StringOutput)
}

// The creation timestamp for this resize request in RFC3339 text format.
func (o ComputeRegionResizeRequestOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resize-request.
func (o ComputeRegionResizeRequestOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The reference of the regional instance group manager this ResizeRequest is a part of.
func (o ComputeRegionResizeRequestOutput) InstanceGroupManager() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.InstanceGroupManager }).(pulumi.StringOutput)
}

// The name of this resize request. The name must be 1-63 characters long, and comply with RFC1035.
func (o ComputeRegionResizeRequestOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRegionResizeRequestOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The reference of the compute region scoping this request. If it is not provided, the provider region is used.
func (o ComputeRegionResizeRequestOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Requested run duration for instances that will be created by this request. At the end of the run duration instances will
// be deleted.
func (o ComputeRegionResizeRequestOutput) RequestedRunDuration() ComputeRegionResizeRequestRequestedRunDurationPtrOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) ComputeRegionResizeRequestRequestedRunDurationPtrOutput {
		return v.RequestedRunDuration
	}).(ComputeRegionResizeRequestRequestedRunDurationPtrOutput)
}

// The number of instances to be created by this resize request. The group's target size will be increased by this number.
func (o ComputeRegionResizeRequestOutput) ResizeBy() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.Float64Output { return v.ResizeBy }).(pulumi.Float64Output)
}

// Current state of the request.
func (o ComputeRegionResizeRequestOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Status of the request.
func (o ComputeRegionResizeRequestOutput) Statuses() ComputeRegionResizeRequestStatusArrayOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) ComputeRegionResizeRequestStatusArrayOutput { return v.Statuses }).(ComputeRegionResizeRequestStatusArrayOutput)
}

func (o ComputeRegionResizeRequestOutput) Timeouts() ComputeRegionResizeRequestTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRegionResizeRequest) ComputeRegionResizeRequestTimeoutsPtrOutput { return v.Timeouts }).(ComputeRegionResizeRequestTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionResizeRequestInput)(nil)).Elem(), &ComputeRegionResizeRequest{})
	pulumi.RegisterOutputType(ComputeRegionResizeRequestOutput{})
}
