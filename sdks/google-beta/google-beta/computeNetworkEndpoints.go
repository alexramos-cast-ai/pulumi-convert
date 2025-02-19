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

type ComputeNetworkEndpoints struct {
	pulumi.CustomResourceState

	ComputeNetworkEndpointsId pulumi.StringOutput `pulumi:"computeNetworkEndpointsId"`
	// The network endpoint group these endpoints are part of.
	NetworkEndpointGroup pulumi.StringOutput `pulumi:"networkEndpointGroup"`
	// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
	// and port, along with additional information depending on the NEG type.
	NetworkEndpoints ComputeNetworkEndpointsNetworkEndpointArrayOutput `pulumi:"networkEndpoints"`
	Project          pulumi.StringOutput                               `pulumi:"project"`
	Timeouts         ComputeNetworkEndpointsTimeoutsPtrOutput          `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeNetworkEndpoints registers a new resource with the given unique name, arguments, and options.
func NewComputeNetworkEndpoints(ctx *pulumi.Context,
	name string, args *ComputeNetworkEndpointsArgs, opts ...pulumi.ResourceOption) (*ComputeNetworkEndpoints, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'NetworkEndpointGroup'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeNetworkEndpoints
	err = ctx.RegisterPackageResource("google-beta:index/computeNetworkEndpoints:ComputeNetworkEndpoints", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeNetworkEndpoints gets an existing ComputeNetworkEndpoints resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeNetworkEndpoints(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeNetworkEndpointsState, opts ...pulumi.ResourceOption) (*ComputeNetworkEndpoints, error) {
	var resource ComputeNetworkEndpoints
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeNetworkEndpoints:ComputeNetworkEndpoints", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeNetworkEndpoints resources.
type computeNetworkEndpointsState struct {
	ComputeNetworkEndpointsId *string `pulumi:"computeNetworkEndpointsId"`
	// The network endpoint group these endpoints are part of.
	NetworkEndpointGroup *string `pulumi:"networkEndpointGroup"`
	// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
	// and port, along with additional information depending on the NEG type.
	NetworkEndpoints []ComputeNetworkEndpointsNetworkEndpoint `pulumi:"networkEndpoints"`
	Project          *string                                  `pulumi:"project"`
	Timeouts         *ComputeNetworkEndpointsTimeouts         `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

type ComputeNetworkEndpointsState struct {
	ComputeNetworkEndpointsId pulumi.StringPtrInput
	// The network endpoint group these endpoints are part of.
	NetworkEndpointGroup pulumi.StringPtrInput
	// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
	// and port, along with additional information depending on the NEG type.
	NetworkEndpoints ComputeNetworkEndpointsNetworkEndpointArrayInput
	Project          pulumi.StringPtrInput
	Timeouts         ComputeNetworkEndpointsTimeoutsPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (ComputeNetworkEndpointsState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkEndpointsState)(nil)).Elem()
}

type computeNetworkEndpointsArgs struct {
	ComputeNetworkEndpointsId *string `pulumi:"computeNetworkEndpointsId"`
	// The network endpoint group these endpoints are part of.
	NetworkEndpointGroup string `pulumi:"networkEndpointGroup"`
	// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
	// and port, along with additional information depending on the NEG type.
	NetworkEndpoints []ComputeNetworkEndpointsNetworkEndpoint `pulumi:"networkEndpoints"`
	Project          *string                                  `pulumi:"project"`
	Timeouts         *ComputeNetworkEndpointsTimeouts         `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeNetworkEndpoints resource.
type ComputeNetworkEndpointsArgs struct {
	ComputeNetworkEndpointsId pulumi.StringPtrInput
	// The network endpoint group these endpoints are part of.
	NetworkEndpointGroup pulumi.StringInput
	// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
	// and port, along with additional information depending on the NEG type.
	NetworkEndpoints ComputeNetworkEndpointsNetworkEndpointArrayInput
	Project          pulumi.StringPtrInput
	Timeouts         ComputeNetworkEndpointsTimeoutsPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (ComputeNetworkEndpointsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkEndpointsArgs)(nil)).Elem()
}

type ComputeNetworkEndpointsInput interface {
	pulumi.Input

	ToComputeNetworkEndpointsOutput() ComputeNetworkEndpointsOutput
	ToComputeNetworkEndpointsOutputWithContext(ctx context.Context) ComputeNetworkEndpointsOutput
}

func (*ComputeNetworkEndpoints) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkEndpoints)(nil)).Elem()
}

func (i *ComputeNetworkEndpoints) ToComputeNetworkEndpointsOutput() ComputeNetworkEndpointsOutput {
	return i.ToComputeNetworkEndpointsOutputWithContext(context.Background())
}

func (i *ComputeNetworkEndpoints) ToComputeNetworkEndpointsOutputWithContext(ctx context.Context) ComputeNetworkEndpointsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeNetworkEndpointsOutput)
}

type ComputeNetworkEndpointsOutput struct{ *pulumi.OutputState }

func (ComputeNetworkEndpointsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkEndpoints)(nil)).Elem()
}

func (o ComputeNetworkEndpointsOutput) ToComputeNetworkEndpointsOutput() ComputeNetworkEndpointsOutput {
	return o
}

func (o ComputeNetworkEndpointsOutput) ToComputeNetworkEndpointsOutputWithContext(ctx context.Context) ComputeNetworkEndpointsOutput {
	return o
}

func (o ComputeNetworkEndpointsOutput) ComputeNetworkEndpointsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) pulumi.StringOutput { return v.ComputeNetworkEndpointsId }).(pulumi.StringOutput)
}

// The network endpoint group these endpoints are part of.
func (o ComputeNetworkEndpointsOutput) NetworkEndpointGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) pulumi.StringOutput { return v.NetworkEndpointGroup }).(pulumi.StringOutput)
}

// The network endpoints to be added to the enclosing network endpoint group (NEG). Each endpoint specifies an IP address
// and port, along with additional information depending on the NEG type.
func (o ComputeNetworkEndpointsOutput) NetworkEndpoints() ComputeNetworkEndpointsNetworkEndpointArrayOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) ComputeNetworkEndpointsNetworkEndpointArrayOutput {
		return v.NetworkEndpoints
	}).(ComputeNetworkEndpointsNetworkEndpointArrayOutput)
}

func (o ComputeNetworkEndpointsOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeNetworkEndpointsOutput) Timeouts() ComputeNetworkEndpointsTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) ComputeNetworkEndpointsTimeoutsPtrOutput { return v.Timeouts }).(ComputeNetworkEndpointsTimeoutsPtrOutput)
}

// Zone where the containing network endpoint group is located.
func (o ComputeNetworkEndpointsOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoints) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeNetworkEndpointsInput)(nil)).Elem(), &ComputeNetworkEndpoints{})
	pulumi.RegisterOutputType(ComputeNetworkEndpointsOutput{})
}
