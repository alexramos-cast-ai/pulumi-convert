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

type ComputeTargetHttpProxy struct {
	pulumi.CustomResourceState

	ComputeTargetHttpProxyId pulumi.StringOutput `pulumi:"computeTargetHttpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
	// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
	// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
	// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
	// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
	// option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.Float64PtrOutput `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolOutput `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId  pulumi.Float64Output                    `pulumi:"proxyId"`
	SelfLink pulumi.StringOutput                     `pulumi:"selfLink"`
	Timeouts ComputeTargetHttpProxyTimeoutsPtrOutput `pulumi:"timeouts"`
	// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringOutput `pulumi:"urlMap"`
}

// NewComputeTargetHttpProxy registers a new resource with the given unique name, arguments, and options.
func NewComputeTargetHttpProxy(ctx *pulumi.Context,
	name string, args *ComputeTargetHttpProxyArgs, opts ...pulumi.ResourceOption) (*ComputeTargetHttpProxy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UrlMap == nil {
		return nil, errors.New("invalid value for required argument 'UrlMap'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeTargetHttpProxy
	err = ctx.RegisterPackageResource("google-beta:index/computeTargetHttpProxy:ComputeTargetHttpProxy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeTargetHttpProxy gets an existing ComputeTargetHttpProxy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeTargetHttpProxy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeTargetHttpProxyState, opts ...pulumi.ResourceOption) (*ComputeTargetHttpProxy, error) {
	var resource ComputeTargetHttpProxy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeTargetHttpProxy:ComputeTargetHttpProxy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeTargetHttpProxy resources.
type computeTargetHttpProxyState struct {
	ComputeTargetHttpProxyId *string `pulumi:"computeTargetHttpProxyId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
	// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
	// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
	// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
	// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
	// option is not available publicly.
	HttpKeepAliveTimeoutSec *float64 `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool `pulumi:"proxyBind"`
	// The unique identifier for the resource.
	ProxyId  *float64                        `pulumi:"proxyId"`
	SelfLink *string                         `pulumi:"selfLink"`
	Timeouts *ComputeTargetHttpProxyTimeouts `pulumi:"timeouts"`
	// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap *string `pulumi:"urlMap"`
}

type ComputeTargetHttpProxyState struct {
	ComputeTargetHttpProxyId pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
	// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
	// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
	// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
	// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
	// option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.Float64PtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	// The unique identifier for the resource.
	ProxyId  pulumi.Float64PtrInput
	SelfLink pulumi.StringPtrInput
	Timeouts ComputeTargetHttpProxyTimeoutsPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringPtrInput
}

func (ComputeTargetHttpProxyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeTargetHttpProxyState)(nil)).Elem()
}

type computeTargetHttpProxyArgs struct {
	ComputeTargetHttpProxyId *string `pulumi:"computeTargetHttpProxyId"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
	// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
	// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
	// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
	// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
	// option is not available publicly.
	HttpKeepAliveTimeoutSec *float64 `pulumi:"httpKeepAliveTimeoutSec"`
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind *bool                           `pulumi:"proxyBind"`
	Timeouts  *ComputeTargetHttpProxyTimeouts `pulumi:"timeouts"`
	// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap string `pulumi:"urlMap"`
}

// The set of arguments for constructing a ComputeTargetHttpProxy resource.
type ComputeTargetHttpProxyArgs struct {
	ComputeTargetHttpProxyId pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
	// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
	// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
	// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
	// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
	// option is not available publicly.
	HttpKeepAliveTimeoutSec pulumi.Float64PtrInput
	// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
	// INTERNAL_SELF_MANAGED.
	ProxyBind pulumi.BoolPtrInput
	Timeouts  ComputeTargetHttpProxyTimeoutsPtrInput
	// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
	UrlMap pulumi.StringInput
}

func (ComputeTargetHttpProxyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeTargetHttpProxyArgs)(nil)).Elem()
}

type ComputeTargetHttpProxyInput interface {
	pulumi.Input

	ToComputeTargetHttpProxyOutput() ComputeTargetHttpProxyOutput
	ToComputeTargetHttpProxyOutputWithContext(ctx context.Context) ComputeTargetHttpProxyOutput
}

func (*ComputeTargetHttpProxy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeTargetHttpProxy)(nil)).Elem()
}

func (i *ComputeTargetHttpProxy) ToComputeTargetHttpProxyOutput() ComputeTargetHttpProxyOutput {
	return i.ToComputeTargetHttpProxyOutputWithContext(context.Background())
}

func (i *ComputeTargetHttpProxy) ToComputeTargetHttpProxyOutputWithContext(ctx context.Context) ComputeTargetHttpProxyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeTargetHttpProxyOutput)
}

type ComputeTargetHttpProxyOutput struct{ *pulumi.OutputState }

func (ComputeTargetHttpProxyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeTargetHttpProxy)(nil)).Elem()
}

func (o ComputeTargetHttpProxyOutput) ToComputeTargetHttpProxyOutput() ComputeTargetHttpProxyOutput {
	return o
}

func (o ComputeTargetHttpProxyOutput) ToComputeTargetHttpProxyOutputWithContext(ctx context.Context) ComputeTargetHttpProxyOutput {
	return o
}

func (o ComputeTargetHttpProxyOutput) ComputeTargetHttpProxyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.ComputeTargetHttpProxyId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeTargetHttpProxyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o ComputeTargetHttpProxyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies how long to keep a connection open, after completing a response, while there is no matching traffic (in
// seconds). If an HTTP keepalive is not specified, a default value will be used. For Global external HTTP(S) load
// balancer, the default value is 610 seconds, the minimum allowed value is 5 seconds and the maximum allowed value is 1200
// seconds. For cross-region internal HTTP(S) load balancer, the default value is 600 seconds, the minimum allowed value is
// 5 seconds, and the maximum allowed value is 600 seconds. For Global external HTTP(S) load balancer (classic), this
// option is not available publicly.
func (o ComputeTargetHttpProxyOutput) HttpKeepAliveTimeoutSec() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.Float64PtrOutput { return v.HttpKeepAliveTimeoutSec }).(pulumi.Float64PtrOutput)
}

// Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeTargetHttpProxyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeTargetHttpProxyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to
// INTERNAL_SELF_MANAGED.
func (o ComputeTargetHttpProxyOutput) ProxyBind() pulumi.BoolOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.BoolOutput { return v.ProxyBind }).(pulumi.BoolOutput)
}

// The unique identifier for the resource.
func (o ComputeTargetHttpProxyOutput) ProxyId() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.Float64Output { return v.ProxyId }).(pulumi.Float64Output)
}

func (o ComputeTargetHttpProxyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o ComputeTargetHttpProxyOutput) Timeouts() ComputeTargetHttpProxyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) ComputeTargetHttpProxyTimeoutsPtrOutput { return v.Timeouts }).(ComputeTargetHttpProxyTimeoutsPtrOutput)
}

// A reference to the UrlMap resource that defines the mapping from URL to the BackendService.
func (o ComputeTargetHttpProxyOutput) UrlMap() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeTargetHttpProxy) pulumi.StringOutput { return v.UrlMap }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeTargetHttpProxyInput)(nil)).Elem(), &ComputeTargetHttpProxy{})
	pulumi.RegisterOutputType(ComputeTargetHttpProxyOutput{})
}
