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

type ComputeRouterRoutePolicy struct {
	pulumi.CustomResourceState

	ComputeRouterRoutePolicyId pulumi.StringOutput `pulumi:"computeRouterRoutePolicyId"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
	// by the Router
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the router and NAT reside.
	Region pulumi.StringOutput `pulumi:"region"`
	// The name of the Cloud Router in which this route policy will be configured.
	Router pulumi.StringOutput `pulumi:"router"`
	// List of terms (the order in the list is not important, they are evaluated in order of priority).
	Terms    ComputeRouterRoutePolicyTermArrayOutput   `pulumi:"terms"`
	Timeouts ComputeRouterRoutePolicyTimeoutsPtrOutput `pulumi:"timeouts"`
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
	// "ROUTE_POLICY_TYPE_EXPORT"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewComputeRouterRoutePolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeRouterRoutePolicy(ctx *pulumi.Context,
	name string, args *ComputeRouterRoutePolicyArgs, opts ...pulumi.ResourceOption) (*ComputeRouterRoutePolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Router == nil {
		return nil, errors.New("invalid value for required argument 'Router'")
	}
	if args.Terms == nil {
		return nil, errors.New("invalid value for required argument 'Terms'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRouterRoutePolicy
	err = ctx.RegisterPackageResource("google-beta:index/computeRouterRoutePolicy:ComputeRouterRoutePolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRouterRoutePolicy gets an existing ComputeRouterRoutePolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRouterRoutePolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRouterRoutePolicyState, opts ...pulumi.ResourceOption) (*ComputeRouterRoutePolicy, error) {
	var resource ComputeRouterRoutePolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeRouterRoutePolicy:ComputeRouterRoutePolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRouterRoutePolicy resources.
type computeRouterRoutePolicyState struct {
	ComputeRouterRoutePolicyId *string `pulumi:"computeRouterRoutePolicyId"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	Fingerprint *string `pulumi:"fingerprint"`
	// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
	// by the Router
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Region where the router and NAT reside.
	Region *string `pulumi:"region"`
	// The name of the Cloud Router in which this route policy will be configured.
	Router *string `pulumi:"router"`
	// List of terms (the order in the list is not important, they are evaluated in order of priority).
	Terms    []ComputeRouterRoutePolicyTerm    `pulumi:"terms"`
	Timeouts *ComputeRouterRoutePolicyTimeouts `pulumi:"timeouts"`
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
	// "ROUTE_POLICY_TYPE_EXPORT"]
	Type *string `pulumi:"type"`
}

type ComputeRouterRoutePolicyState struct {
	ComputeRouterRoutePolicyId pulumi.StringPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	Fingerprint pulumi.StringPtrInput
	// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
	// by the Router
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Region where the router and NAT reside.
	Region pulumi.StringPtrInput
	// The name of the Cloud Router in which this route policy will be configured.
	Router pulumi.StringPtrInput
	// List of terms (the order in the list is not important, they are evaluated in order of priority).
	Terms    ComputeRouterRoutePolicyTermArrayInput
	Timeouts ComputeRouterRoutePolicyTimeoutsPtrInput
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
	// "ROUTE_POLICY_TYPE_EXPORT"]
	Type pulumi.StringPtrInput
}

func (ComputeRouterRoutePolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRouterRoutePolicyState)(nil)).Elem()
}

type computeRouterRoutePolicyArgs struct {
	ComputeRouterRoutePolicyId *string `pulumi:"computeRouterRoutePolicyId"`
	// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
	// by the Router
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Region where the router and NAT reside.
	Region *string `pulumi:"region"`
	// The name of the Cloud Router in which this route policy will be configured.
	Router string `pulumi:"router"`
	// List of terms (the order in the list is not important, they are evaluated in order of priority).
	Terms    []ComputeRouterRoutePolicyTerm    `pulumi:"terms"`
	Timeouts *ComputeRouterRoutePolicyTimeouts `pulumi:"timeouts"`
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
	// "ROUTE_POLICY_TYPE_EXPORT"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a ComputeRouterRoutePolicy resource.
type ComputeRouterRoutePolicyArgs struct {
	ComputeRouterRoutePolicyId pulumi.StringPtrInput
	// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
	// by the Router
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Region where the router and NAT reside.
	Region pulumi.StringPtrInput
	// The name of the Cloud Router in which this route policy will be configured.
	Router pulumi.StringInput
	// List of terms (the order in the list is not important, they are evaluated in order of priority).
	Terms    ComputeRouterRoutePolicyTermArrayInput
	Timeouts ComputeRouterRoutePolicyTimeoutsPtrInput
	// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
	// "ROUTE_POLICY_TYPE_EXPORT"]
	Type pulumi.StringPtrInput
}

func (ComputeRouterRoutePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRouterRoutePolicyArgs)(nil)).Elem()
}

type ComputeRouterRoutePolicyInput interface {
	pulumi.Input

	ToComputeRouterRoutePolicyOutput() ComputeRouterRoutePolicyOutput
	ToComputeRouterRoutePolicyOutputWithContext(ctx context.Context) ComputeRouterRoutePolicyOutput
}

func (*ComputeRouterRoutePolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRouterRoutePolicy)(nil)).Elem()
}

func (i *ComputeRouterRoutePolicy) ToComputeRouterRoutePolicyOutput() ComputeRouterRoutePolicyOutput {
	return i.ToComputeRouterRoutePolicyOutputWithContext(context.Background())
}

func (i *ComputeRouterRoutePolicy) ToComputeRouterRoutePolicyOutputWithContext(ctx context.Context) ComputeRouterRoutePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRouterRoutePolicyOutput)
}

type ComputeRouterRoutePolicyOutput struct{ *pulumi.OutputState }

func (ComputeRouterRoutePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRouterRoutePolicy)(nil)).Elem()
}

func (o ComputeRouterRoutePolicyOutput) ToComputeRouterRoutePolicyOutput() ComputeRouterRoutePolicyOutput {
	return o
}

func (o ComputeRouterRoutePolicyOutput) ToComputeRouterRoutePolicyOutputWithContext(ctx context.Context) ComputeRouterRoutePolicyOutput {
	return o
}

func (o ComputeRouterRoutePolicyOutput) ComputeRouterRoutePolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.ComputeRouterRoutePolicyId }).(pulumi.StringOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeRouterRoutePolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// Name of the route policy. This policy's name, which must be a resource ID segment and unique within all policies owned
// by the Router
func (o ComputeRouterRoutePolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRouterRoutePolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Region where the router and NAT reside.
func (o ComputeRouterRoutePolicyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The name of the Cloud Router in which this route policy will be configured.
func (o ComputeRouterRoutePolicyOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringOutput { return v.Router }).(pulumi.StringOutput)
}

// List of terms (the order in the list is not important, they are evaluated in order of priority).
func (o ComputeRouterRoutePolicyOutput) Terms() ComputeRouterRoutePolicyTermArrayOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) ComputeRouterRoutePolicyTermArrayOutput { return v.Terms }).(ComputeRouterRoutePolicyTermArrayOutput)
}

func (o ComputeRouterRoutePolicyOutput) Timeouts() ComputeRouterRoutePolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) ComputeRouterRoutePolicyTimeoutsPtrOutput { return v.Timeouts }).(ComputeRouterRoutePolicyTimeoutsPtrOutput)
}

// This is policy's type, which is one of IMPORT or EXPORT Possible values: ["ROUTE_POLICY_TYPE_IMPORT",
// "ROUTE_POLICY_TYPE_EXPORT"]
func (o ComputeRouterRoutePolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRouterRoutePolicy) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRouterRoutePolicyInput)(nil)).Elem(), &ComputeRouterRoutePolicy{})
	pulumi.RegisterOutputType(ComputeRouterRoutePolicyOutput{})
}
