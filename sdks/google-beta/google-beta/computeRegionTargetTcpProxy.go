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

type ComputeRegionTargetTcpProxy struct {
	pulumi.CustomResourceState

	// A reference to the BackendService resource.
	BackendService                pulumi.StringOutput `pulumi:"backendService"`
	ComputeRegionTargetTcpProxyId pulumi.StringOutput `pulumi:"computeRegionTargetTcpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrOutput `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId pulumi.Float64Output `pulumi:"proxyId"`
	// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringOutput                          `pulumi:"region"`
	SelfLink pulumi.StringOutput                          `pulumi:"selfLink"`
	Timeouts ComputeRegionTargetTcpProxyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeRegionTargetTcpProxy registers a new resource with the given unique name, arguments, and options.
func NewComputeRegionTargetTcpProxy(ctx *pulumi.Context,
	name string, args *ComputeRegionTargetTcpProxyArgs, opts ...pulumi.ResourceOption) (*ComputeRegionTargetTcpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendService == nil {
		return nil, errors.New("invalid value for required argument 'BackendService'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeRegionTargetTcpProxy
	err = ctx.RegisterPackageResource("google-beta:index/computeRegionTargetTcpProxy:ComputeRegionTargetTcpProxy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeRegionTargetTcpProxy gets an existing ComputeRegionTargetTcpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeRegionTargetTcpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeRegionTargetTcpProxyState, opts ...pulumi.ResourceOption) (*ComputeRegionTargetTcpProxy, error) {
	var resource ComputeRegionTargetTcpProxy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeRegionTargetTcpProxy:ComputeRegionTargetTcpProxy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeRegionTargetTcpProxy resources.
type computeRegionTargetTcpProxyState struct {
	// A reference to the BackendService resource.
	BackendService                *string `pulumi:"backendService"`
	ComputeRegionTargetTcpProxyId *string `pulumi:"computeRegionTargetTcpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The unique identifier for the resource.
	ProxyId *float64 `pulumi:"proxyId"`
	// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
	Region   *string                              `pulumi:"region"`
	SelfLink *string                              `pulumi:"selfLink"`
	Timeouts *ComputeRegionTargetTcpProxyTimeouts `pulumi:"timeouts"`
}

type ComputeRegionTargetTcpProxyState struct {
	// A reference to the BackendService resource.
	BackendService                pulumi.StringPtrInput
	ComputeRegionTargetTcpProxyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrInput
	// The unique identifier for the resource.
	ProxyId pulumi.Float64PtrInput
	// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringPtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeRegionTargetTcpProxyTimeoutsPtrInput
}

func (ComputeRegionTargetTcpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionTargetTcpProxyState)(nil)).Elem()
}

type computeRegionTargetTcpProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService                string  `pulumi:"backendService"`
	ComputeRegionTargetTcpProxyId *string `pulumi:"computeRegionTargetTcpProxyId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader *string `pulumi:"proxyHeader"`
	// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
	Region   *string                              `pulumi:"region"`
	Timeouts *ComputeRegionTargetTcpProxyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeRegionTargetTcpProxy resource.
type ComputeRegionTargetTcpProxyArgs struct {
	// A reference to the BackendService resource.
	BackendService                pulumi.StringInput
	ComputeRegionTargetTcpProxyId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
	// ["NONE", "PROXY_V1"]
	ProxyHeader pulumi.StringPtrInput
	// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
	Region   pulumi.StringPtrInput
	Timeouts ComputeRegionTargetTcpProxyTimeoutsPtrInput
}

func (ComputeRegionTargetTcpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeRegionTargetTcpProxyArgs)(nil)).Elem()
}

type ComputeRegionTargetTcpProxyInput interface {
	pulumi.Input

	ToComputeRegionTargetTcpProxyOutput() ComputeRegionTargetTcpProxyOutput
	ToComputeRegionTargetTcpProxyOutputWithContext(ctx context.Context) ComputeRegionTargetTcpProxyOutput
}

func (*ComputeRegionTargetTcpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionTargetTcpProxy)(nil)).Elem()
}

func (i *ComputeRegionTargetTcpProxy) ToComputeRegionTargetTcpProxyOutput() ComputeRegionTargetTcpProxyOutput {
	return i.ToComputeRegionTargetTcpProxyOutputWithContext(context.Background())
}

func (i *ComputeRegionTargetTcpProxy) ToComputeRegionTargetTcpProxyOutputWithContext(ctx context.Context) ComputeRegionTargetTcpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeRegionTargetTcpProxyOutput)
}

type ComputeRegionTargetTcpProxyOutput struct{ *pulumi.OutputState }

func (ComputeRegionTargetTcpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeRegionTargetTcpProxy)(nil)).Elem()
}

func (o ComputeRegionTargetTcpProxyOutput) ToComputeRegionTargetTcpProxyOutput() ComputeRegionTargetTcpProxyOutput {
	return o
}

func (o ComputeRegionTargetTcpProxyOutput) ToComputeRegionTargetTcpProxyOutputWithContext(ctx context.Context) ComputeRegionTargetTcpProxyOutput {
	return o
}

// A reference to the BackendService resource.
func (o ComputeRegionTargetTcpProxyOutput) BackendService() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.BackendService }).(pulumi.StringOutput)
}

func (o ComputeRegionTargetTcpProxyOutput) ComputeRegionTargetTcpProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.ComputeRegionTargetTcpProxyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeRegionTargetTcpProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeRegionTargetTcpProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeRegionTargetTcpProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeRegionTargetTcpProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
// INTERNAL_SELF_MANAGED.
func (o ComputeRegionTargetTcpProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// Specifies the type of proxy header to append before sending data to the backend. Default value: "NONE" Possible values:
// ["NONE", "PROXY_V1"]
func (o ComputeRegionTargetTcpProxyOutput) ProxyHeader() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringPtrOutput { return v.ProxyHeader }).(pulumi.StringPtrOutput)
}

// The unique identifier for the resource.
func (o ComputeRegionTargetTcpProxyOutput) ProxyId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.Float64Output { return v.ProxyId }).(pulumi.Float64Output)
}

// The Region in which the created target TCP proxy should reside. If it is not provided, the provider region is used.
func (o ComputeRegionTargetTcpProxyOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComputeRegionTargetTcpProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeRegionTargetTcpProxyOutput) Timeouts() ComputeRegionTargetTcpProxyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeRegionTargetTcpProxy) ComputeRegionTargetTcpProxyTimeoutsPtrOutput { return v.Timeouts }).(ComputeRegionTargetTcpProxyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeRegionTargetTcpProxyInput)(nil)).Elem(), &ComputeRegionTargetTcpProxy{})
	pulumi.RegisterOutputType(ComputeRegionTargetTcpProxyOutput{})
}
