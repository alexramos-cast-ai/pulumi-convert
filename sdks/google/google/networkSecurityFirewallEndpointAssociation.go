// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkSecurityFirewallEndpointAssociation struct {
	pulumi.CustomResourceState

	// Time the firewall endpoint was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
	// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
	// be disabled.
	Disabled        pulumi.BoolPtrOutput   `pulumi:"disabled"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The URL of the firewall endpoint that is being associated.
	FirewallEndpoint pulumi.StringOutput `pulumi:"firewallEndpoint"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location (zone) of the firewall endpoint association.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the firewall endpoint association resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the network that is being associated.
	Network                                      pulumi.StringOutput `pulumi:"network"`
	NetworkSecurityFirewallEndpointAssociationId pulumi.StringOutput `pulumi:"networkSecurityFirewallEndpointAssociationId"`
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The current state of the endpoint.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                      `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityFirewallEndpointAssociationTimeoutsPtrOutput `pulumi:"timeouts"`
	// The URL of the TlsInspectionPolicy that is being associated.
	TlsInspectionPolicy pulumi.StringPtrOutput `pulumi:"tlsInspectionPolicy"`
	// Time the firewall endpoint was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityFirewallEndpointAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityFirewallEndpointAssociation(ctx *pulumi.Context,
	name string, args *NetworkSecurityFirewallEndpointAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityFirewallEndpointAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FirewallEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'FirewallEndpoint'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityFirewallEndpointAssociation
	err = ctx.RegisterPackageResource("google:index/networkSecurityFirewallEndpointAssociation:NetworkSecurityFirewallEndpointAssociation", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityFirewallEndpointAssociation gets an existing NetworkSecurityFirewallEndpointAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityFirewallEndpointAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityFirewallEndpointAssociationState, opts ...pulumi.ResourceOption) (*NetworkSecurityFirewallEndpointAssociation, error) {
	var resource NetworkSecurityFirewallEndpointAssociation
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/networkSecurityFirewallEndpointAssociation:NetworkSecurityFirewallEndpointAssociation", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityFirewallEndpointAssociation resources.
type networkSecurityFirewallEndpointAssociationState struct {
	// Time the firewall endpoint was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
	// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
	// be disabled.
	Disabled        *bool             `pulumi:"disabled"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The URL of the firewall endpoint that is being associated.
	FirewallEndpoint *string `pulumi:"firewallEndpoint"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location (zone) of the firewall endpoint association.
	Location *string `pulumi:"location"`
	// The name of the firewall endpoint association resource.
	Name *string `pulumi:"name"`
	// The URL of the network that is being associated.
	Network                                      *string `pulumi:"network"`
	NetworkSecurityFirewallEndpointAssociationId *string `pulumi:"networkSecurityFirewallEndpointAssociationId"`
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	Parent *string `pulumi:"parent"`
	// Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling *bool `pulumi:"reconciling"`
	// Server-defined URL of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// The current state of the endpoint.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                   `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityFirewallEndpointAssociationTimeouts `pulumi:"timeouts"`
	// The URL of the TlsInspectionPolicy that is being associated.
	TlsInspectionPolicy *string `pulumi:"tlsInspectionPolicy"`
	// Time the firewall endpoint was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityFirewallEndpointAssociationState struct {
	// Time the firewall endpoint was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
	// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
	// be disabled.
	Disabled        pulumi.BoolPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The URL of the firewall endpoint that is being associated.
	FirewallEndpoint pulumi.StringPtrInput
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location (zone) of the firewall endpoint association.
	Location pulumi.StringPtrInput
	// The name of the firewall endpoint association resource.
	Name pulumi.StringPtrInput
	// The URL of the network that is being associated.
	Network                                      pulumi.StringPtrInput
	NetworkSecurityFirewallEndpointAssociationId pulumi.StringPtrInput
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	Parent pulumi.StringPtrInput
	// Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolPtrInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	// The current state of the endpoint.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityFirewallEndpointAssociationTimeoutsPtrInput
	// The URL of the TlsInspectionPolicy that is being associated.
	TlsInspectionPolicy pulumi.StringPtrInput
	// Time the firewall endpoint was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityFirewallEndpointAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityFirewallEndpointAssociationState)(nil)).Elem()
}

type networkSecurityFirewallEndpointAssociationArgs struct {
	// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
	// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
	// be disabled.
	Disabled *bool `pulumi:"disabled"`
	// The URL of the firewall endpoint that is being associated.
	FirewallEndpoint string `pulumi:"firewallEndpoint"`
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location (zone) of the firewall endpoint association.
	Location string `pulumi:"location"`
	// The name of the firewall endpoint association resource.
	Name *string `pulumi:"name"`
	// The URL of the network that is being associated.
	Network                                      string  `pulumi:"network"`
	NetworkSecurityFirewallEndpointAssociationId *string `pulumi:"networkSecurityFirewallEndpointAssociationId"`
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	Parent   *string                                             `pulumi:"parent"`
	Timeouts *NetworkSecurityFirewallEndpointAssociationTimeouts `pulumi:"timeouts"`
	// The URL of the TlsInspectionPolicy that is being associated.
	TlsInspectionPolicy *string `pulumi:"tlsInspectionPolicy"`
}

// The set of arguments for constructing a NetworkSecurityFirewallEndpointAssociation resource.
type NetworkSecurityFirewallEndpointAssociationArgs struct {
	// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
	// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
	// be disabled.
	Disabled pulumi.BoolPtrInput
	// The URL of the firewall endpoint that is being associated.
	FirewallEndpoint pulumi.StringInput
	// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The location (zone) of the firewall endpoint association.
	Location pulumi.StringInput
	// The name of the firewall endpoint association resource.
	Name pulumi.StringPtrInput
	// The URL of the network that is being associated.
	Network                                      pulumi.StringInput
	NetworkSecurityFirewallEndpointAssociationId pulumi.StringPtrInput
	// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
	Parent   pulumi.StringPtrInput
	Timeouts NetworkSecurityFirewallEndpointAssociationTimeoutsPtrInput
	// The URL of the TlsInspectionPolicy that is being associated.
	TlsInspectionPolicy pulumi.StringPtrInput
}

func (NetworkSecurityFirewallEndpointAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityFirewallEndpointAssociationArgs)(nil)).Elem()
}

type NetworkSecurityFirewallEndpointAssociationInput interface {
	pulumi.Input

	ToNetworkSecurityFirewallEndpointAssociationOutput() NetworkSecurityFirewallEndpointAssociationOutput
	ToNetworkSecurityFirewallEndpointAssociationOutputWithContext(ctx context.Context) NetworkSecurityFirewallEndpointAssociationOutput
}

func (*NetworkSecurityFirewallEndpointAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityFirewallEndpointAssociation)(nil)).Elem()
}

func (i *NetworkSecurityFirewallEndpointAssociation) ToNetworkSecurityFirewallEndpointAssociationOutput() NetworkSecurityFirewallEndpointAssociationOutput {
	return i.ToNetworkSecurityFirewallEndpointAssociationOutputWithContext(context.Background())
}

func (i *NetworkSecurityFirewallEndpointAssociation) ToNetworkSecurityFirewallEndpointAssociationOutputWithContext(ctx context.Context) NetworkSecurityFirewallEndpointAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityFirewallEndpointAssociationOutput)
}

type NetworkSecurityFirewallEndpointAssociationOutput struct{ *pulumi.OutputState }

func (NetworkSecurityFirewallEndpointAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityFirewallEndpointAssociation)(nil)).Elem()
}

func (o NetworkSecurityFirewallEndpointAssociationOutput) ToNetworkSecurityFirewallEndpointAssociationOutput() NetworkSecurityFirewallEndpointAssociationOutput {
	return o
}

func (o NetworkSecurityFirewallEndpointAssociationOutput) ToNetworkSecurityFirewallEndpointAssociationOutputWithContext(ctx context.Context) NetworkSecurityFirewallEndpointAssociationOutput {
	return o
}

// Time the firewall endpoint was created in UTC.
func (o NetworkSecurityFirewallEndpointAssociationOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether the association is disabled. True indicates that traffic will not be intercepted. > **Note:** The API will
// reject the request if this value is set to true when creating the resource, otherwise on an update the association can
// be disabled.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

func (o NetworkSecurityFirewallEndpointAssociationOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The URL of the firewall endpoint that is being associated.
func (o NetworkSecurityFirewallEndpointAssociationOutput) FirewallEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.FirewallEndpoint }).(pulumi.StringOutput)
}

// A map of key/value label pairs to assign to the resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location (zone) of the firewall endpoint association.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the firewall endpoint association resource.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of the network that is being associated.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o NetworkSecurityFirewallEndpointAssociationOutput) NetworkSecurityFirewallEndpointAssociationId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput {
		return v.NetworkSecurityFirewallEndpointAssociationId
	}).(pulumi.StringOutput)
}

// The name of the parent this firewall endpoint association belongs to. Format: projects/{project_id}.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Whether reconciling is in progress, recommended per https://google.aip.dev/128.
func (o NetworkSecurityFirewallEndpointAssociationOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Server-defined URL of this resource.
func (o NetworkSecurityFirewallEndpointAssociationOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The current state of the endpoint.
func (o NetworkSecurityFirewallEndpointAssociationOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityFirewallEndpointAssociationOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityFirewallEndpointAssociationOutput) Timeouts() NetworkSecurityFirewallEndpointAssociationTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) NetworkSecurityFirewallEndpointAssociationTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityFirewallEndpointAssociationTimeoutsPtrOutput)
}

// The URL of the TlsInspectionPolicy that is being associated.
func (o NetworkSecurityFirewallEndpointAssociationOutput) TlsInspectionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringPtrOutput {
		return v.TlsInspectionPolicy
	}).(pulumi.StringPtrOutput)
}

// Time the firewall endpoint was updated in UTC.
func (o NetworkSecurityFirewallEndpointAssociationOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityFirewallEndpointAssociation) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityFirewallEndpointAssociationInput)(nil)).Elem(), &NetworkSecurityFirewallEndpointAssociation{})
	pulumi.RegisterOutputType(NetworkSecurityFirewallEndpointAssociationOutput{})
}
