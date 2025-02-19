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

type NetworkServicesHttpRoute struct {
	pulumi.CustomResourceState

	// Time the HttpRoute was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayOutput `pulumi:"gateways"`
	// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
	// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
	// Mesh should be of a type SIDECAR.
	Meshes pulumi.StringArrayOutput `pulumi:"meshes"`
	// Name of the HttpRoute resource.
	Name                       pulumi.StringOutput `pulumi:"name"`
	NetworkServicesHttpRouteId pulumi.StringOutput `pulumi:"networkServicesHttpRouteId"`
	Project                    pulumi.StringOutput `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules NetworkServicesHttpRouteRuleArrayOutput `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                    `pulumi:"terraformLabels"`
	Timeouts        NetworkServicesHttpRouteTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the HttpRoute was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkServicesHttpRoute registers a new resource with the given unique name, arguments, and options.
func NewNetworkServicesHttpRoute(ctx *pulumi.Context,
	name string, args *NetworkServicesHttpRouteArgs, opts ...pulumi.ResourceOption) (*NetworkServicesHttpRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostnames == nil {
		return nil, errors.New("invalid value for required argument 'Hostnames'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkServicesHttpRoute
	err = ctx.RegisterPackageResource("google-beta:index/networkServicesHttpRoute:NetworkServicesHttpRoute", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkServicesHttpRoute gets an existing NetworkServicesHttpRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkServicesHttpRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkServicesHttpRouteState, opts ...pulumi.ResourceOption) (*NetworkServicesHttpRoute, error) {
	var resource NetworkServicesHttpRoute
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkServicesHttpRoute:NetworkServicesHttpRoute", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkServicesHttpRoute resources.
type networkServicesHttpRouteState struct {
	// Time the HttpRoute was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
	Hostnames []string `pulumi:"hostnames"`
	// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
	// Mesh should be of a type SIDECAR.
	Meshes []string `pulumi:"meshes"`
	// Name of the HttpRoute resource.
	Name                       *string `pulumi:"name"`
	NetworkServicesHttpRouteId *string `pulumi:"networkServicesHttpRouteId"`
	Project                    *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules []NetworkServicesHttpRouteRule `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                 `pulumi:"terraformLabels"`
	Timeouts        *NetworkServicesHttpRouteTimeouts `pulumi:"timeouts"`
	// Time the HttpRoute was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkServicesHttpRouteState struct {
	// Time the HttpRoute was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
	Hostnames pulumi.StringArrayInput
	// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
	// Mesh should be of a type SIDECAR.
	Meshes pulumi.StringArrayInput
	// Name of the HttpRoute resource.
	Name                       pulumi.StringPtrInput
	NetworkServicesHttpRouteId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	Rules NetworkServicesHttpRouteRuleArrayInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkServicesHttpRouteTimeoutsPtrInput
	// Time the HttpRoute was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkServicesHttpRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesHttpRouteState)(nil)).Elem()
}

type networkServicesHttpRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways []string `pulumi:"gateways"`
	// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
	Hostnames []string `pulumi:"hostnames"`
	// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
	// Mesh should be of a type SIDECAR.
	Meshes []string `pulumi:"meshes"`
	// Name of the HttpRoute resource.
	Name                       *string `pulumi:"name"`
	NetworkServicesHttpRouteId *string `pulumi:"networkServicesHttpRouteId"`
	Project                    *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	Rules    []NetworkServicesHttpRouteRule    `pulumi:"rules"`
	Timeouts *NetworkServicesHttpRouteTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkServicesHttpRoute resource.
type NetworkServicesHttpRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
	// served by the gateway. Each gateway reference should match the pattern:
	// projects/*/locations/global/gateways/<gateway_name>
	Gateways pulumi.StringArrayInput
	// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
	Hostnames pulumi.StringArrayInput
	// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
	// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
	// Mesh should be of a type SIDECAR.
	Meshes pulumi.StringArrayInput
	// Name of the HttpRoute resource.
	Name                       pulumi.StringPtrInput
	NetworkServicesHttpRouteId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	Rules    NetworkServicesHttpRouteRuleArrayInput
	Timeouts NetworkServicesHttpRouteTimeoutsPtrInput
}

func (NetworkServicesHttpRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesHttpRouteArgs)(nil)).Elem()
}

type NetworkServicesHttpRouteInput interface {
	pulumi.Input

	ToNetworkServicesHttpRouteOutput() NetworkServicesHttpRouteOutput
	ToNetworkServicesHttpRouteOutputWithContext(ctx context.Context) NetworkServicesHttpRouteOutput
}

func (*NetworkServicesHttpRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesHttpRoute)(nil)).Elem()
}

func (i *NetworkServicesHttpRoute) ToNetworkServicesHttpRouteOutput() NetworkServicesHttpRouteOutput {
	return i.ToNetworkServicesHttpRouteOutputWithContext(context.Background())
}

func (i *NetworkServicesHttpRoute) ToNetworkServicesHttpRouteOutputWithContext(ctx context.Context) NetworkServicesHttpRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkServicesHttpRouteOutput)
}

type NetworkServicesHttpRouteOutput struct{ *pulumi.OutputState }

func (NetworkServicesHttpRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesHttpRoute)(nil)).Elem()
}

func (o NetworkServicesHttpRouteOutput) ToNetworkServicesHttpRouteOutput() NetworkServicesHttpRouteOutput {
	return o
}

func (o NetworkServicesHttpRouteOutput) ToNetworkServicesHttpRouteOutputWithContext(ctx context.Context) NetworkServicesHttpRouteOutput {
	return o
}

// Time the HttpRoute was created in UTC.
func (o NetworkServicesHttpRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o NetworkServicesHttpRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkServicesHttpRouteOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Gateways defines a list of gateways this HttpRoute is attached to, as one of the routing rules to route the requests
// served by the gateway. Each gateway reference should match the pattern:
// projects/*/locations/global/gateways/<gateway_name>
func (o NetworkServicesHttpRouteOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringArrayOutput { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Set of hosts that should match against the HTTP host header to select a HttpRoute to process the request.
func (o NetworkServicesHttpRouteOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringArrayOutput { return v.Hostnames }).(pulumi.StringArrayOutput)
}

// Set of label tags associated with the HttpRoute resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkServicesHttpRouteOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Meshes defines a list of meshes this HttpRoute is attached to, as one of the routing rules to route the requests served
// by the mesh. Each mesh reference should match the pattern: projects/*/locations/global/meshes/<mesh_name>. The attached
// Mesh should be of a type SIDECAR.
func (o NetworkServicesHttpRouteOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringArrayOutput { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the HttpRoute resource.
func (o NetworkServicesHttpRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkServicesHttpRouteOutput) NetworkServicesHttpRouteId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.NetworkServicesHttpRouteId }).(pulumi.StringOutput)
}

func (o NetworkServicesHttpRouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Rules that define how traffic is routed and handled.
func (o NetworkServicesHttpRouteOutput) Rules() NetworkServicesHttpRouteRuleArrayOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) NetworkServicesHttpRouteRuleArrayOutput { return v.Rules }).(NetworkServicesHttpRouteRuleArrayOutput)
}

// Server-defined URL of this resource.
func (o NetworkServicesHttpRouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkServicesHttpRouteOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkServicesHttpRouteOutput) Timeouts() NetworkServicesHttpRouteTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) NetworkServicesHttpRouteTimeoutsPtrOutput { return v.Timeouts }).(NetworkServicesHttpRouteTimeoutsPtrOutput)
}

// Time the HttpRoute was updated in UTC.
func (o NetworkServicesHttpRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesHttpRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkServicesHttpRouteInput)(nil)).Elem(), &NetworkServicesHttpRoute{})
	pulumi.RegisterOutputType(NetworkServicesHttpRouteOutput{})
}
