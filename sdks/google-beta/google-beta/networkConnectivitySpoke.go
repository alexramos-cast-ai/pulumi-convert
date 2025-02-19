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

type NetworkConnectivitySpoke struct {
	pulumi.CustomResourceState

	// Output only. The time the spoke was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An optional description of the spoke.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the group that this spoke is associated with.
	Group pulumi.StringOutput `pulumi:"group"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringOutput `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
	// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
	// advertising the same prefixes.
	LinkedInterconnectAttachments NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrOutput `pulumi:"linkedInterconnectAttachments"`
	// Producer VPC network that is associated with the spoke.
	LinkedProducerVpcNetwork NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrOutput `pulumi:"linkedProducerVpcNetwork"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrOutput `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork NetworkConnectivitySpokeLinkedVpcNetworkPtrOutput `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels NetworkConnectivitySpokeLinkedVpnTunnelsPtrOutput `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name                       pulumi.StringOutput `pulumi:"name"`
	NetworkConnectivitySpokeId pulumi.StringOutput `pulumi:"networkConnectivitySpokeId"`
	Project                    pulumi.StringOutput `pulumi:"project"`
	// Output only. The current lifecycle state of this spoke.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                    `pulumi:"terraformLabels"`
	Timeouts        NetworkConnectivitySpokeTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is
	// deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// Output only. The time the spoke was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkConnectivitySpoke registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnectivitySpoke(ctx *pulumi.Context,
	name string, args *NetworkConnectivitySpokeArgs, opts ...pulumi.ResourceOption) (*NetworkConnectivitySpoke, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hub == nil {
		return nil, errors.New("invalid value for required argument 'Hub'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkConnectivitySpoke
	err = ctx.RegisterPackageResource("google-beta:index/networkConnectivitySpoke:NetworkConnectivitySpoke", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkConnectivitySpoke gets an existing NetworkConnectivitySpoke resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkConnectivitySpoke(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectivitySpokeState, opts ...pulumi.ResourceOption) (*NetworkConnectivitySpoke, error) {
	var resource NetworkConnectivitySpoke
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkConnectivitySpoke:NetworkConnectivitySpoke", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkConnectivitySpoke resources.
type networkConnectivitySpokeState struct {
	// Output only. The time the spoke was created.
	CreateTime *string `pulumi:"createTime"`
	// An optional description of the spoke.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the group that this spoke is associated with.
	Group *string `pulumi:"group"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub *string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
	// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
	// advertising the same prefixes.
	LinkedInterconnectAttachments *NetworkConnectivitySpokeLinkedInterconnectAttachments `pulumi:"linkedInterconnectAttachments"`
	// Producer VPC network that is associated with the spoke.
	LinkedProducerVpcNetwork *NetworkConnectivitySpokeLinkedProducerVpcNetwork `pulumi:"linkedProducerVpcNetwork"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances *NetworkConnectivitySpokeLinkedRouterApplianceInstances `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork *NetworkConnectivitySpokeLinkedVpcNetwork `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels *NetworkConnectivitySpokeLinkedVpnTunnels `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name                       *string `pulumi:"name"`
	NetworkConnectivitySpokeId *string `pulumi:"networkConnectivitySpokeId"`
	Project                    *string `pulumi:"project"`
	// Output only. The current lifecycle state of this spoke.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                 `pulumi:"terraformLabels"`
	Timeouts        *NetworkConnectivitySpokeTimeouts `pulumi:"timeouts"`
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is
	// deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId *string `pulumi:"uniqueId"`
	// Output only. The time the spoke was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkConnectivitySpokeState struct {
	// Output only. The time the spoke was created.
	CreateTime pulumi.StringPtrInput
	// An optional description of the spoke.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The name of the group that this spoke is associated with.
	Group pulumi.StringPtrInput
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringPtrInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
	// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
	// advertising the same prefixes.
	LinkedInterconnectAttachments NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrInput
	// Producer VPC network that is associated with the spoke.
	LinkedProducerVpcNetwork NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrInput
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork NetworkConnectivitySpokeLinkedVpcNetworkPtrInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels NetworkConnectivitySpokeLinkedVpnTunnelsPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name                       pulumi.StringPtrInput
	NetworkConnectivitySpokeId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	// Output only. The current lifecycle state of this spoke.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkConnectivitySpokeTimeoutsPtrInput
	// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is
	// deleted and another with the same name is created, the new spoke is assigned a different unique_id.
	UniqueId pulumi.StringPtrInput
	// Output only. The time the spoke was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkConnectivitySpokeState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectivitySpokeState)(nil)).Elem()
}

type networkConnectivitySpokeArgs struct {
	// An optional description of the spoke.
	Description *string `pulumi:"description"`
	// The name of the group that this spoke is associated with.
	Group *string `pulumi:"group"`
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
	// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
	// advertising the same prefixes.
	LinkedInterconnectAttachments *NetworkConnectivitySpokeLinkedInterconnectAttachments `pulumi:"linkedInterconnectAttachments"`
	// Producer VPC network that is associated with the spoke.
	LinkedProducerVpcNetwork *NetworkConnectivitySpokeLinkedProducerVpcNetwork `pulumi:"linkedProducerVpcNetwork"`
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances *NetworkConnectivitySpokeLinkedRouterApplianceInstances `pulumi:"linkedRouterApplianceInstances"`
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork *NetworkConnectivitySpokeLinkedVpcNetwork `pulumi:"linkedVpcNetwork"`
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels *NetworkConnectivitySpokeLinkedVpnTunnels `pulumi:"linkedVpnTunnels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name                       *string                           `pulumi:"name"`
	NetworkConnectivitySpokeId *string                           `pulumi:"networkConnectivitySpokeId"`
	Project                    *string                           `pulumi:"project"`
	Timeouts                   *NetworkConnectivitySpokeTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkConnectivitySpoke resource.
type NetworkConnectivitySpokeArgs struct {
	// An optional description of the spoke.
	Description pulumi.StringPtrInput
	// The name of the group that this spoke is associated with.
	Group pulumi.StringPtrInput
	// Immutable. The URI of the hub that this spoke is attached to.
	Hub pulumi.StringInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
	// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
	// advertising the same prefixes.
	LinkedInterconnectAttachments NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrInput
	// Producer VPC network that is associated with the spoke.
	LinkedProducerVpcNetwork NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrInput
	// The URIs of linked Router appliance resources
	LinkedRouterApplianceInstances NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrInput
	// VPC network that is associated with the spoke.
	LinkedVpcNetwork NetworkConnectivitySpokeLinkedVpcNetworkPtrInput
	// The URIs of linked VPN tunnel resources
	LinkedVpnTunnels NetworkConnectivitySpokeLinkedVpnTunnelsPtrInput
	// The location for the resource
	Location pulumi.StringInput
	// Immutable. The name of the spoke. Spoke names must be unique.
	Name                       pulumi.StringPtrInput
	NetworkConnectivitySpokeId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	Timeouts                   NetworkConnectivitySpokeTimeoutsPtrInput
}

func (NetworkConnectivitySpokeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectivitySpokeArgs)(nil)).Elem()
}

type NetworkConnectivitySpokeInput interface {
	pulumi.Input

	ToNetworkConnectivitySpokeOutput() NetworkConnectivitySpokeOutput
	ToNetworkConnectivitySpokeOutputWithContext(ctx context.Context) NetworkConnectivitySpokeOutput
}

func (*NetworkConnectivitySpoke) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectivitySpoke)(nil)).Elem()
}

func (i *NetworkConnectivitySpoke) ToNetworkConnectivitySpokeOutput() NetworkConnectivitySpokeOutput {
	return i.ToNetworkConnectivitySpokeOutputWithContext(context.Background())
}

func (i *NetworkConnectivitySpoke) ToNetworkConnectivitySpokeOutputWithContext(ctx context.Context) NetworkConnectivitySpokeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectivitySpokeOutput)
}

type NetworkConnectivitySpokeOutput struct{ *pulumi.OutputState }

func (NetworkConnectivitySpokeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectivitySpoke)(nil)).Elem()
}

func (o NetworkConnectivitySpokeOutput) ToNetworkConnectivitySpokeOutput() NetworkConnectivitySpokeOutput {
	return o
}

func (o NetworkConnectivitySpokeOutput) ToNetworkConnectivitySpokeOutputWithContext(ctx context.Context) NetworkConnectivitySpokeOutput {
	return o
}

// Output only. The time the spoke was created.
func (o NetworkConnectivitySpokeOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// An optional description of the spoke.
func (o NetworkConnectivitySpokeOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectivitySpokeOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the group that this spoke is associated with.
func (o NetworkConnectivitySpokeOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// Immutable. The URI of the hub that this spoke is attached to.
func (o NetworkConnectivitySpokeOutput) Hub() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.Hub }).(pulumi.StringOutput)
}

// Optional labels in key:value format. For more information about labels, see [Requirements for
// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o NetworkConnectivitySpokeOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// A collection of VLAN attachment resources. These resources should be redundant attachments that all advertise the same
// prefixes to Google Cloud. Alternatively, in active/passive configurations, all attachments should be capable of
// advertising the same prefixes.
func (o NetworkConnectivitySpokeOutput) LinkedInterconnectAttachments() NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrOutput {
		return v.LinkedInterconnectAttachments
	}).(NetworkConnectivitySpokeLinkedInterconnectAttachmentsPtrOutput)
}

// Producer VPC network that is associated with the spoke.
func (o NetworkConnectivitySpokeOutput) LinkedProducerVpcNetwork() NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrOutput {
		return v.LinkedProducerVpcNetwork
	}).(NetworkConnectivitySpokeLinkedProducerVpcNetworkPtrOutput)
}

// The URIs of linked Router appliance resources
func (o NetworkConnectivitySpokeOutput) LinkedRouterApplianceInstances() NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrOutput {
		return v.LinkedRouterApplianceInstances
	}).(NetworkConnectivitySpokeLinkedRouterApplianceInstancesPtrOutput)
}

// VPC network that is associated with the spoke.
func (o NetworkConnectivitySpokeOutput) LinkedVpcNetwork() NetworkConnectivitySpokeLinkedVpcNetworkPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeLinkedVpcNetworkPtrOutput {
		return v.LinkedVpcNetwork
	}).(NetworkConnectivitySpokeLinkedVpcNetworkPtrOutput)
}

// The URIs of linked VPN tunnel resources
func (o NetworkConnectivitySpokeOutput) LinkedVpnTunnels() NetworkConnectivitySpokeLinkedVpnTunnelsPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeLinkedVpnTunnelsPtrOutput {
		return v.LinkedVpnTunnels
	}).(NetworkConnectivitySpokeLinkedVpnTunnelsPtrOutput)
}

// The location for the resource
func (o NetworkConnectivitySpokeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Immutable. The name of the spoke. Spoke names must be unique.
func (o NetworkConnectivitySpokeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkConnectivitySpokeOutput) NetworkConnectivitySpokeId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.NetworkConnectivitySpokeId }).(pulumi.StringOutput)
}

func (o NetworkConnectivitySpokeOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The current lifecycle state of this spoke.
func (o NetworkConnectivitySpokeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkConnectivitySpokeOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkConnectivitySpokeOutput) Timeouts() NetworkConnectivitySpokeTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) NetworkConnectivitySpokeTimeoutsPtrOutput { return v.Timeouts }).(NetworkConnectivitySpokeTimeoutsPtrOutput)
}

// Output only. The Google-generated UUID for the spoke. This value is unique across all spoke resources. If a spoke is
// deleted and another with the same name is created, the new spoke is assigned a different unique_id.
func (o NetworkConnectivitySpokeOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.UniqueId }).(pulumi.StringOutput)
}

// Output only. The time the spoke was last updated.
func (o NetworkConnectivitySpokeOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivitySpoke) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConnectivitySpokeInput)(nil)).Elem(), &NetworkConnectivitySpoke{})
	pulumi.RegisterOutputType(NetworkConnectivitySpokeOutput{})
}
