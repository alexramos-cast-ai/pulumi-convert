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

type NetworkSecurityMirroringDeploymentGroup struct {
	pulumi.CustomResourceState

	// Output only. The list of Mirroring Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupArrayOutput `pulumi:"connectedEndpointGroups"`
	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_deployment_group_id from the method_signature of Create RPC
	MirroringDeploymentGroupId pulumi.StringOutput `pulumi:"mirroringDeploymentGroupId"`
	// Immutable. Identifier. Then name of the MirroringDeploymentGroup.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringOutput `pulumi:"network"`
	NetworkSecurityMirroringDeploymentGroupId pulumi.StringOutput `pulumi:"networkSecurityMirroringDeploymentGroupId"`
	Project                                   pulumi.StringOutput `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                   `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityMirroringDeploymentGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityMirroringDeploymentGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityMirroringDeploymentGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityMirroringDeploymentGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityMirroringDeploymentGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MirroringDeploymentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'MirroringDeploymentGroupId'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityMirroringDeploymentGroup
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityMirroringDeploymentGroup:NetworkSecurityMirroringDeploymentGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityMirroringDeploymentGroup gets an existing NetworkSecurityMirroringDeploymentGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityMirroringDeploymentGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityMirroringDeploymentGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityMirroringDeploymentGroup, error) {
	var resource NetworkSecurityMirroringDeploymentGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityMirroringDeploymentGroup:NetworkSecurityMirroringDeploymentGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityMirroringDeploymentGroup resources.
type networkSecurityMirroringDeploymentGroupState struct {
	// Output only. The list of Mirroring Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups []NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroup `pulumi:"connectedEndpointGroups"`
	// Output only. [Output only] Create time stamp
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	Location *string `pulumi:"location"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_deployment_group_id from the method_signature of Create RPC
	MirroringDeploymentGroupId *string `pulumi:"mirroringDeploymentGroupId"`
	// Immutable. Identifier. Then name of the MirroringDeploymentGroup.
	Name *string `pulumi:"name"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   *string `pulumi:"network"`
	NetworkSecurityMirroringDeploymentGroupId *string `pulumi:"networkSecurityMirroringDeploymentGroupId"`
	Project                                   *string `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityMirroringDeploymentGroupTimeouts `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityMirroringDeploymentGroupState struct {
	// Output only. The list of Mirroring Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupArrayInput
	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	Location pulumi.StringPtrInput
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_deployment_group_id from the method_signature of Create RPC
	MirroringDeploymentGroupId pulumi.StringPtrInput
	// Immutable. Identifier. Then name of the MirroringDeploymentGroup.
	Name pulumi.StringPtrInput
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringPtrInput
	NetworkSecurityMirroringDeploymentGroupId pulumi.StringPtrInput
	Project                                   pulumi.StringPtrInput
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolPtrInput
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityMirroringDeploymentGroupTimeoutsPtrInput
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityMirroringDeploymentGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityMirroringDeploymentGroupState)(nil)).Elem()
}

type networkSecurityMirroringDeploymentGroupArgs struct {
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	Location string `pulumi:"location"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_deployment_group_id from the method_signature of Create RPC
	MirroringDeploymentGroupId string `pulumi:"mirroringDeploymentGroupId"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   string                                           `pulumi:"network"`
	NetworkSecurityMirroringDeploymentGroupId *string                                          `pulumi:"networkSecurityMirroringDeploymentGroupId"`
	Project                                   *string                                          `pulumi:"project"`
	Timeouts                                  *NetworkSecurityMirroringDeploymentGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityMirroringDeploymentGroup resource.
type NetworkSecurityMirroringDeploymentGroupArgs struct {
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
	Location pulumi.StringInput
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_deployment_group_id from the method_signature of Create RPC
	MirroringDeploymentGroupId pulumi.StringInput
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringInput
	NetworkSecurityMirroringDeploymentGroupId pulumi.StringPtrInput
	Project                                   pulumi.StringPtrInput
	Timeouts                                  NetworkSecurityMirroringDeploymentGroupTimeoutsPtrInput
}

func (NetworkSecurityMirroringDeploymentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityMirroringDeploymentGroupArgs)(nil)).Elem()
}

type NetworkSecurityMirroringDeploymentGroupInput interface {
	pulumi.Input

	ToNetworkSecurityMirroringDeploymentGroupOutput() NetworkSecurityMirroringDeploymentGroupOutput
	ToNetworkSecurityMirroringDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringDeploymentGroupOutput
}

func (*NetworkSecurityMirroringDeploymentGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityMirroringDeploymentGroup)(nil)).Elem()
}

func (i *NetworkSecurityMirroringDeploymentGroup) ToNetworkSecurityMirroringDeploymentGroupOutput() NetworkSecurityMirroringDeploymentGroupOutput {
	return i.ToNetworkSecurityMirroringDeploymentGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityMirroringDeploymentGroup) ToNetworkSecurityMirroringDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringDeploymentGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityMirroringDeploymentGroupOutput)
}

type NetworkSecurityMirroringDeploymentGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityMirroringDeploymentGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityMirroringDeploymentGroup)(nil)).Elem()
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) ToNetworkSecurityMirroringDeploymentGroupOutput() NetworkSecurityMirroringDeploymentGroupOutput {
	return o
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) ToNetworkSecurityMirroringDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringDeploymentGroupOutput {
	return o
}

// Output only. The list of Mirroring Endpoint Groups that are connected to this resource.
func (o NetworkSecurityMirroringDeploymentGroupOutput) ConnectedEndpointGroups() NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupArrayOutput {
		return v.ConnectedEndpointGroups
	}).(NetworkSecurityMirroringDeploymentGroupConnectedEndpointGroupArrayOutput)
}

// Output only. [Output only] Create time stamp
func (o NetworkSecurityMirroringDeploymentGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o NetworkSecurityMirroringDeploymentGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122. See documentation for resource type
// 'networksecurity.googleapis.com/MirroringDeploymentGroup'.
func (o NetworkSecurityMirroringDeploymentGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
// mirroring_deployment_group_id from the method_signature of Create RPC
func (o NetworkSecurityMirroringDeploymentGroupOutput) MirroringDeploymentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput {
		return v.MirroringDeploymentGroupId
	}).(pulumi.StringOutput)
}

// Immutable. Identifier. Then name of the MirroringDeploymentGroup.
func (o NetworkSecurityMirroringDeploymentGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required. Immutable. The network that is being used for the deployment. Format is:
// projects/{project}/global/networks/{network}.
func (o NetworkSecurityMirroringDeploymentGroupOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) NetworkSecurityMirroringDeploymentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput {
		return v.NetworkSecurityMirroringDeploymentGroupId
	}).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
func (o NetworkSecurityMirroringDeploymentGroupOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
func (o NetworkSecurityMirroringDeploymentGroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityMirroringDeploymentGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityMirroringDeploymentGroupOutput) Timeouts() NetworkSecurityMirroringDeploymentGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) NetworkSecurityMirroringDeploymentGroupTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityMirroringDeploymentGroupTimeoutsPtrOutput)
}

// Output only. [Output only] Update time stamp
func (o NetworkSecurityMirroringDeploymentGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringDeploymentGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityMirroringDeploymentGroupInput)(nil)).Elem(), &NetworkSecurityMirroringDeploymentGroup{})
	pulumi.RegisterOutputType(NetworkSecurityMirroringDeploymentGroupOutput{})
}
