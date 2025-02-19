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

type NetworkServicesTlsRoute struct {
	pulumi.CustomResourceState

	// Time the TlsRoute was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayOutput `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
	// Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayOutput `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name                      pulumi.StringOutput `pulumi:"name"`
	NetworkServicesTlsRouteId pulumi.StringOutput `pulumi:"networkServicesTlsRouteId"`
	Project                   pulumi.StringOutput `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules NetworkServicesTlsRouteRuleArrayOutput `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput                      `pulumi:"selfLink"`
	Timeouts NetworkServicesTlsRouteTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the TlsRoute was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkServicesTlsRoute registers a new resource with the given unique name, arguments, and options.
func NewNetworkServicesTlsRoute(ctx *pulumi.Context,
	name string, args *NetworkServicesTlsRouteArgs, opts ...pulumi.ResourceOption) (*NetworkServicesTlsRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkServicesTlsRoute
	err = ctx.RegisterPackageResource("google-beta:index/networkServicesTlsRoute:NetworkServicesTlsRoute", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkServicesTlsRoute gets an existing NetworkServicesTlsRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkServicesTlsRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkServicesTlsRouteState, opts ...pulumi.ResourceOption) (*NetworkServicesTlsRoute, error) {
	var resource NetworkServicesTlsRoute
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkServicesTlsRoute:NetworkServicesTlsRoute", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkServicesTlsRoute resources.
type networkServicesTlsRouteState struct {
	// Time the TlsRoute was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
	// Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name                      *string `pulumi:"name"`
	NetworkServicesTlsRouteId *string `pulumi:"networkServicesTlsRouteId"`
	Project                   *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules []NetworkServicesTlsRouteRule `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink *string                          `pulumi:"selfLink"`
	Timeouts *NetworkServicesTlsRouteTimeouts `pulumi:"timeouts"`
	// Time the TlsRoute was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkServicesTlsRouteState struct {
	// Time the TlsRoute was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
	// Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayInput
	// Name of the TlsRoute resource.
	Name                      pulumi.StringPtrInput
	NetworkServicesTlsRouteId pulumi.StringPtrInput
	Project                   pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	Rules NetworkServicesTlsRouteRuleArrayInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	Timeouts NetworkServicesTlsRouteTimeoutsPtrInput
	// Time the TlsRoute was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkServicesTlsRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesTlsRouteState)(nil)).Elem()
}

type networkServicesTlsRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
	// Mesh should be of a type SIDECAR
	Meshes []string `pulumi:"meshes"`
	// Name of the TlsRoute resource.
	Name                      *string `pulumi:"name"`
	NetworkServicesTlsRouteId *string `pulumi:"networkServicesTlsRouteId"`
	Project                   *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules    []NetworkServicesTlsRouteRule    `pulumi:"rules"`
	Timeouts *NetworkServicesTlsRouteTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkServicesTlsRoute resource.
type NetworkServicesTlsRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
	// Mesh should be of a type SIDECAR
	Meshes pulumi.StringArrayInput
	// Name of the TlsRoute resource.
	Name                      pulumi.StringPtrInput
	NetworkServicesTlsRouteId pulumi.StringPtrInput
	Project                   pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	Rules    NetworkServicesTlsRouteRuleArrayInput
	Timeouts NetworkServicesTlsRouteTimeoutsPtrInput
}

func (NetworkServicesTlsRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesTlsRouteArgs)(nil)).Elem()
}

type NetworkServicesTlsRouteInput interface {
	pulumi.Input

	ToNetworkServicesTlsRouteOutput() NetworkServicesTlsRouteOutput
	ToNetworkServicesTlsRouteOutputWithContext(ctx context.Context) NetworkServicesTlsRouteOutput
}

func (*NetworkServicesTlsRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesTlsRoute)(nil)).Elem()
}

func (i *NetworkServicesTlsRoute) ToNetworkServicesTlsRouteOutput() NetworkServicesTlsRouteOutput {
	return i.ToNetworkServicesTlsRouteOutputWithContext(context.Background())
}

func (i *NetworkServicesTlsRoute) ToNetworkServicesTlsRouteOutputWithContext(ctx context.Context) NetworkServicesTlsRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkServicesTlsRouteOutput)
}

type NetworkServicesTlsRouteOutput struct{ *pulumi.OutputState }

func (NetworkServicesTlsRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesTlsRoute)(nil)).Elem()
}

func (o NetworkServicesTlsRouteOutput) ToNetworkServicesTlsRouteOutput() NetworkServicesTlsRouteOutput {
	return o
}

func (o NetworkServicesTlsRouteOutput) ToNetworkServicesTlsRouteOutputWithContext(ctx context.Context) NetworkServicesTlsRouteOutput {
	return o
}

// Time the TlsRoute was created in UTC.
func (o NetworkServicesTlsRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o NetworkServicesTlsRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Gateways defines a list of gateways this TlsRoute is attached to, as one of the routing rules to route the requests
// served by the gateway. Each gateway reference should match the pattern:
// projects/*/locations/global/gateways/<gateway_name>
func (o NetworkServicesTlsRouteOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringArrayOutput { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Meshes defines a list of meshes this TlsRoute is attached to, as one of the routing rules to route the requests served
// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name> The attached
// Mesh should be of a type SIDECAR
func (o NetworkServicesTlsRouteOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringArrayOutput { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the TlsRoute resource.
func (o NetworkServicesTlsRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkServicesTlsRouteOutput) NetworkServicesTlsRouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.NetworkServicesTlsRouteId }).(pulumi.StringOutput)
}

func (o NetworkServicesTlsRouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Rules that define how traffic is routed and handled.
func (o NetworkServicesTlsRouteOutput) Rules() NetworkServicesTlsRouteRuleArrayOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) NetworkServicesTlsRouteRuleArrayOutput { return v.Rules }).(NetworkServicesTlsRouteRuleArrayOutput)
}

// Server-defined URL of this resource.
func (o NetworkServicesTlsRouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

func (o NetworkServicesTlsRouteOutput) Timeouts() NetworkServicesTlsRouteTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) NetworkServicesTlsRouteTimeoutsPtrOutput { return v.Timeouts }).(NetworkServicesTlsRouteTimeoutsPtrOutput)
}

// Time the TlsRoute was updated in UTC.
func (o NetworkServicesTlsRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesTlsRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkServicesTlsRouteInput)(nil)).Elem(), &NetworkServicesTlsRoute{})
	pulumi.RegisterOutputType(NetworkServicesTlsRouteOutput{})
}
