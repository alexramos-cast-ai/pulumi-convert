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

type NetworkSecurityMirroringEndpointGroup struct {
	pulumi.CustomResourceState

	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
	MirroringDeploymentGroup pulumi.StringOutput `pulumi:"mirroringDeploymentGroup"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_endpoint_group_id from the method_signature of Create RPC
	MirroringEndpointGroupId pulumi.StringOutput `pulumi:"mirroringEndpointGroupId"`
	// Immutable. Identifier. The name of the MirroringEndpointGroup.
	Name                                    pulumi.StringOutput `pulumi:"name"`
	NetworkSecurityMirroringEndpointGroupId pulumi.StringOutput `pulumi:"networkSecurityMirroringEndpointGroupId"`
	Project                                 pulumi.StringOutput `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. Current state of the endpoint group. Possible values: STATE_UNSPECIFIED ACTIVE CLOSED CREATING DELETING
	// OUT_OF_SYNC
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                 `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityMirroringEndpointGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityMirroringEndpointGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityMirroringEndpointGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityMirroringEndpointGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityMirroringEndpointGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MirroringDeploymentGroup == nil {
		return nil, errors.New("invalid value for required argument 'MirroringDeploymentGroup'")
	}
	if args.MirroringEndpointGroupId == nil {
		return nil, errors.New("invalid value for required argument 'MirroringEndpointGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkSecurityMirroringEndpointGroup
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityMirroringEndpointGroup:NetworkSecurityMirroringEndpointGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityMirroringEndpointGroup gets an existing NetworkSecurityMirroringEndpointGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityMirroringEndpointGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityMirroringEndpointGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityMirroringEndpointGroup, error) {
	var resource NetworkSecurityMirroringEndpointGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityMirroringEndpointGroup:NetworkSecurityMirroringEndpointGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityMirroringEndpointGroup resources.
type networkSecurityMirroringEndpointGroupState struct {
	// Output only. [Output only] Create time stamp
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
	Location *string `pulumi:"location"`
	// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
	MirroringDeploymentGroup *string `pulumi:"mirroringDeploymentGroup"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_endpoint_group_id from the method_signature of Create RPC
	MirroringEndpointGroupId *string `pulumi:"mirroringEndpointGroupId"`
	// Immutable. Identifier. The name of the MirroringEndpointGroup.
	Name                                    *string `pulumi:"name"`
	NetworkSecurityMirroringEndpointGroupId *string `pulumi:"networkSecurityMirroringEndpointGroupId"`
	Project                                 *string `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. Current state of the endpoint group. Possible values: STATE_UNSPECIFIED ACTIVE CLOSED CREATING DELETING
	// OUT_OF_SYNC
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                              `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityMirroringEndpointGroupTimeouts `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityMirroringEndpointGroupState struct {
	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
	Location pulumi.StringPtrInput
	// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
	MirroringDeploymentGroup pulumi.StringPtrInput
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_endpoint_group_id from the method_signature of Create RPC
	MirroringEndpointGroupId pulumi.StringPtrInput
	// Immutable. Identifier. The name of the MirroringEndpointGroup.
	Name                                    pulumi.StringPtrInput
	NetworkSecurityMirroringEndpointGroupId pulumi.StringPtrInput
	Project                                 pulumi.StringPtrInput
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolPtrInput
	// Output only. Current state of the endpoint group. Possible values: STATE_UNSPECIFIED ACTIVE CLOSED CREATING DELETING
	// OUT_OF_SYNC
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityMirroringEndpointGroupTimeoutsPtrInput
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityMirroringEndpointGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityMirroringEndpointGroupState)(nil)).Elem()
}

type networkSecurityMirroringEndpointGroupArgs struct {
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
	Location string `pulumi:"location"`
	// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
	MirroringDeploymentGroup string `pulumi:"mirroringDeploymentGroup"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_endpoint_group_id from the method_signature of Create RPC
	MirroringEndpointGroupId                string                                         `pulumi:"mirroringEndpointGroupId"`
	NetworkSecurityMirroringEndpointGroupId *string                                        `pulumi:"networkSecurityMirroringEndpointGroupId"`
	Project                                 *string                                        `pulumi:"project"`
	Timeouts                                *NetworkSecurityMirroringEndpointGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityMirroringEndpointGroup resource.
type NetworkSecurityMirroringEndpointGroupArgs struct {
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
	Location pulumi.StringInput
	// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
	// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
	MirroringDeploymentGroup pulumi.StringInput
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// mirroring_endpoint_group_id from the method_signature of Create RPC
	MirroringEndpointGroupId                pulumi.StringInput
	NetworkSecurityMirroringEndpointGroupId pulumi.StringPtrInput
	Project                                 pulumi.StringPtrInput
	Timeouts                                NetworkSecurityMirroringEndpointGroupTimeoutsPtrInput
}

func (NetworkSecurityMirroringEndpointGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityMirroringEndpointGroupArgs)(nil)).Elem()
}

type NetworkSecurityMirroringEndpointGroupInput interface {
	pulumi.Input

	ToNetworkSecurityMirroringEndpointGroupOutput() NetworkSecurityMirroringEndpointGroupOutput
	ToNetworkSecurityMirroringEndpointGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringEndpointGroupOutput
}

func (*NetworkSecurityMirroringEndpointGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityMirroringEndpointGroup)(nil)).Elem()
}

func (i *NetworkSecurityMirroringEndpointGroup) ToNetworkSecurityMirroringEndpointGroupOutput() NetworkSecurityMirroringEndpointGroupOutput {
	return i.ToNetworkSecurityMirroringEndpointGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityMirroringEndpointGroup) ToNetworkSecurityMirroringEndpointGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringEndpointGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityMirroringEndpointGroupOutput)
}

type NetworkSecurityMirroringEndpointGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityMirroringEndpointGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityMirroringEndpointGroup)(nil)).Elem()
}

func (o NetworkSecurityMirroringEndpointGroupOutput) ToNetworkSecurityMirroringEndpointGroupOutput() NetworkSecurityMirroringEndpointGroupOutput {
	return o
}

func (o NetworkSecurityMirroringEndpointGroupOutput) ToNetworkSecurityMirroringEndpointGroupOutputWithContext(ctx context.Context) NetworkSecurityMirroringEndpointGroupOutput {
	return o
}

// Output only. [Output only] Create time stamp
func (o NetworkSecurityMirroringEndpointGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringEndpointGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o NetworkSecurityMirroringEndpointGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122. See documentation for resource type 'networksecurity.googleapis.com/MirroringEndpointGroup'.
func (o NetworkSecurityMirroringEndpointGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Required. Immutable. The Mirroring Deployment Group that this resource is connected to. Format is:
// 'projects/{project}/locations/global/mirroringDeploymentGroups/{mirroringDeploymentGroup}'
func (o NetworkSecurityMirroringEndpointGroupOutput) MirroringDeploymentGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.MirroringDeploymentGroup }).(pulumi.StringOutput)
}

// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
// mirroring_endpoint_group_id from the method_signature of Create RPC
func (o NetworkSecurityMirroringEndpointGroupOutput) MirroringEndpointGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.MirroringEndpointGroupId }).(pulumi.StringOutput)
}

// Immutable. Identifier. The name of the MirroringEndpointGroup.
func (o NetworkSecurityMirroringEndpointGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringEndpointGroupOutput) NetworkSecurityMirroringEndpointGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput {
		return v.NetworkSecurityMirroringEndpointGroupId
	}).(pulumi.StringOutput)
}

func (o NetworkSecurityMirroringEndpointGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
func (o NetworkSecurityMirroringEndpointGroupOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. Current state of the endpoint group. Possible values: STATE_UNSPECIFIED ACTIVE CLOSED CREATING DELETING
// OUT_OF_SYNC
func (o NetworkSecurityMirroringEndpointGroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityMirroringEndpointGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityMirroringEndpointGroupOutput) Timeouts() NetworkSecurityMirroringEndpointGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) NetworkSecurityMirroringEndpointGroupTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityMirroringEndpointGroupTimeoutsPtrOutput)
}

// Output only. [Output only] Update time stamp
func (o NetworkSecurityMirroringEndpointGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityMirroringEndpointGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityMirroringEndpointGroupInput)(nil)).Elem(), &NetworkSecurityMirroringEndpointGroup{})
	pulumi.RegisterOutputType(NetworkSecurityMirroringEndpointGroupOutput{})
}
