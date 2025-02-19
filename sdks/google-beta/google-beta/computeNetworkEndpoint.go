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

type ComputeNetworkEndpoint struct {
	pulumi.CustomResourceState

	ComputeNetworkEndpointId pulumi.StringOutput `pulumi:"computeNetworkEndpointId"`
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringPtrOutput `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringOutput `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
	// of 'GCE_VM_IP'
	Port     pulumi.Float64PtrOutput                 `pulumi:"port"`
	Project  pulumi.StringOutput                     `pulumi:"project"`
	Timeouts ComputeNetworkEndpointTimeoutsPtrOutput `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeNetworkEndpoint registers a new resource with the given unique name, arguments, and options.
func NewComputeNetworkEndpoint(ctx *pulumi.Context,
	name string, args *ComputeNetworkEndpointArgs, opts ...pulumi.ResourceOption) (*ComputeNetworkEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	if args.NetworkEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'NetworkEndpointGroup'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeNetworkEndpoint
	err = ctx.RegisterPackageResource("google-beta:index/computeNetworkEndpoint:ComputeNetworkEndpoint", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeNetworkEndpoint gets an existing ComputeNetworkEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeNetworkEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeNetworkEndpointState, opts ...pulumi.ResourceOption) (*ComputeNetworkEndpoint, error) {
	var resource ComputeNetworkEndpoint
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeNetworkEndpoint:ComputeNetworkEndpoint", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeNetworkEndpoint resources.
type computeNetworkEndpointState struct {
	ComputeNetworkEndpointId *string `pulumi:"computeNetworkEndpointId"`
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance *string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
	IpAddress *string `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup *string `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
	// of 'GCE_VM_IP'
	Port     *float64                        `pulumi:"port"`
	Project  *string                         `pulumi:"project"`
	Timeouts *ComputeNetworkEndpointTimeouts `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

type ComputeNetworkEndpointState struct {
	ComputeNetworkEndpointId pulumi.StringPtrInput
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringPtrInput
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
	IpAddress pulumi.StringPtrInput
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringPtrInput
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
	// of 'GCE_VM_IP'
	Port     pulumi.Float64PtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeNetworkEndpointTimeoutsPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (ComputeNetworkEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkEndpointState)(nil)).Elem()
}

type computeNetworkEndpointArgs struct {
	ComputeNetworkEndpointId *string `pulumi:"computeNetworkEndpointId"`
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance *string `pulumi:"instance"`
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
	IpAddress string `pulumi:"ipAddress"`
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup string `pulumi:"networkEndpointGroup"`
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
	// of 'GCE_VM_IP'
	Port     *float64                        `pulumi:"port"`
	Project  *string                         `pulumi:"project"`
	Timeouts *ComputeNetworkEndpointTimeouts `pulumi:"timeouts"`
	// Zone where the containing network endpoint group is located.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeNetworkEndpoint resource.
type ComputeNetworkEndpointArgs struct {
	ComputeNetworkEndpointId pulumi.StringPtrInput
	// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
	// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
	Instance pulumi.StringPtrInput
	// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
	// aliased IP range).
	IpAddress pulumi.StringInput
	// The network endpoint group this endpoint is part of.
	NetworkEndpointGroup pulumi.StringInput
	// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
	// of 'GCE_VM_IP'
	Port     pulumi.Float64PtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeNetworkEndpointTimeoutsPtrInput
	// Zone where the containing network endpoint group is located.
	Zone pulumi.StringPtrInput
}

func (ComputeNetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeNetworkEndpointArgs)(nil)).Elem()
}

type ComputeNetworkEndpointInput interface {
	pulumi.Input

	ToComputeNetworkEndpointOutput() ComputeNetworkEndpointOutput
	ToComputeNetworkEndpointOutputWithContext(ctx context.Context) ComputeNetworkEndpointOutput
}

func (*ComputeNetworkEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkEndpoint)(nil)).Elem()
}

func (i *ComputeNetworkEndpoint) ToComputeNetworkEndpointOutput() ComputeNetworkEndpointOutput {
	return i.ToComputeNetworkEndpointOutputWithContext(context.Background())
}

func (i *ComputeNetworkEndpoint) ToComputeNetworkEndpointOutputWithContext(ctx context.Context) ComputeNetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeNetworkEndpointOutput)
}

type ComputeNetworkEndpointOutput struct{ *pulumi.OutputState }

func (ComputeNetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeNetworkEndpoint)(nil)).Elem()
}

func (o ComputeNetworkEndpointOutput) ToComputeNetworkEndpointOutput() ComputeNetworkEndpointOutput {
	return o
}

func (o ComputeNetworkEndpointOutput) ToComputeNetworkEndpointOutputWithContext(ctx context.Context) ComputeNetworkEndpointOutput {
	return o
}

func (o ComputeNetworkEndpointOutput) ComputeNetworkEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringOutput { return v.ComputeNetworkEndpointId }).(pulumi.StringOutput)
}

// The name for a specific VM instance that the IP address belongs to. This is required for network endpoints of type
// GCE_VM_IP_PORT. The instance must be in the same zone of network endpoint group.
func (o ComputeNetworkEndpointOutput) Instance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringPtrOutput { return v.Instance }).(pulumi.StringPtrOutput)
}

// IPv4 address of network endpoint. The IP address must belong to a VM in GCE (either the primary IP or as part of an
// aliased IP range).
func (o ComputeNetworkEndpointOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The network endpoint group this endpoint is part of.
func (o ComputeNetworkEndpointOutput) NetworkEndpointGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringOutput { return v.NetworkEndpointGroup }).(pulumi.StringOutput)
}

// Port number of network endpoint. **Note** 'port' is required unless the Network Endpoint Group is created with the type
// of 'GCE_VM_IP'
func (o ComputeNetworkEndpointOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.Float64PtrOutput { return v.Port }).(pulumi.Float64PtrOutput)
}

func (o ComputeNetworkEndpointOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeNetworkEndpointOutput) Timeouts() ComputeNetworkEndpointTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) ComputeNetworkEndpointTimeoutsPtrOutput { return v.Timeouts }).(ComputeNetworkEndpointTimeoutsPtrOutput)
}

// Zone where the containing network endpoint group is located.
func (o ComputeNetworkEndpointOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeNetworkEndpoint) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeNetworkEndpointInput)(nil)).Elem(), &ComputeNetworkEndpoint{})
	pulumi.RegisterOutputType(ComputeNetworkEndpointOutput{})
}
