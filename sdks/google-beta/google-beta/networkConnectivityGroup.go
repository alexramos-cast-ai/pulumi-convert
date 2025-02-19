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

type NetworkConnectivityGroup struct {
	pulumi.CustomResourceState

	// Optional. The auto-accept setting for this group.
	AutoAccept NetworkConnectivityGroupAutoAcceptPtrOutput `pulumi:"autoAccept"`
	// Output only. The time the hub was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An optional description of the group.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the hub. Hub names must be unique. They use the following form:
	// projects/{projectNumber}/locations/global/hubs/{hubId}
	Hub pulumi.StringOutput `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
	Name                       pulumi.StringOutput `pulumi:"name"`
	NetworkConnectivityGroupId pulumi.StringOutput `pulumi:"networkConnectivityGroupId"`
	Project                    pulumi.StringOutput `pulumi:"project"`
	// Output only. The name of the route table that corresponds to this group. They use the following form:
	// 'projects/{projectNumber}/locations/global/hubs/{hubId}/routeTables/{route_table_id}'
	RouteTable pulumi.StringOutput `pulumi:"routeTable"`
	// Output only. The current lifecycle state of this hub.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                    `pulumi:"terraformLabels"`
	Timeouts        NetworkConnectivityGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is
	// deleted and another with the same name is created, the new route table is assigned a different uniqueId.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time the hub was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkConnectivityGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkConnectivityGroup(ctx *pulumi.Context,
	name string, args *NetworkConnectivityGroupArgs, opts ...pulumi.ResourceOption) (*NetworkConnectivityGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hub == nil {
		return nil, errors.New("invalid value for required argument 'Hub'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkConnectivityGroup
	err = ctx.RegisterPackageResource("google-beta:index/networkConnectivityGroup:NetworkConnectivityGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkConnectivityGroup gets an existing NetworkConnectivityGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkConnectivityGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkConnectivityGroupState, opts ...pulumi.ResourceOption) (*NetworkConnectivityGroup, error) {
	var resource NetworkConnectivityGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkConnectivityGroup:NetworkConnectivityGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkConnectivityGroup resources.
type networkConnectivityGroupState struct {
	// Optional. The auto-accept setting for this group.
	AutoAccept *NetworkConnectivityGroupAutoAccept `pulumi:"autoAccept"`
	// Output only. The time the hub was created.
	CreateTime *string `pulumi:"createTime"`
	// An optional description of the group.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the hub. Hub names must be unique. They use the following form:
	// projects/{projectNumber}/locations/global/hubs/{hubId}
	Hub *string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
	Name                       *string `pulumi:"name"`
	NetworkConnectivityGroupId *string `pulumi:"networkConnectivityGroupId"`
	Project                    *string `pulumi:"project"`
	// Output only. The name of the route table that corresponds to this group. They use the following form:
	// 'projects/{projectNumber}/locations/global/hubs/{hubId}/routeTables/{route_table_id}'
	RouteTable *string `pulumi:"routeTable"`
	// Output only. The current lifecycle state of this hub.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                 `pulumi:"terraformLabels"`
	Timeouts        *NetworkConnectivityGroupTimeouts `pulumi:"timeouts"`
	// Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is
	// deleted and another with the same name is created, the new route table is assigned a different uniqueId.
	Uid *string `pulumi:"uid"`
	// Output only. The time the hub was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkConnectivityGroupState struct {
	// Optional. The auto-accept setting for this group.
	AutoAccept NetworkConnectivityGroupAutoAcceptPtrInput
	// Output only. The time the hub was created.
	CreateTime pulumi.StringPtrInput
	// An optional description of the group.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The name of the hub. Hub names must be unique. They use the following form:
	// projects/{projectNumber}/locations/global/hubs/{hubId}
	Hub pulumi.StringPtrInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
	Name                       pulumi.StringPtrInput
	NetworkConnectivityGroupId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	// Output only. The name of the route table that corresponds to this group. They use the following form:
	// 'projects/{projectNumber}/locations/global/hubs/{hubId}/routeTables/{route_table_id}'
	RouteTable pulumi.StringPtrInput
	// Output only. The current lifecycle state of this hub.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkConnectivityGroupTimeoutsPtrInput
	// Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is
	// deleted and another with the same name is created, the new route table is assigned a different uniqueId.
	Uid pulumi.StringPtrInput
	// Output only. The time the hub was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (NetworkConnectivityGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectivityGroupState)(nil)).Elem()
}

type networkConnectivityGroupArgs struct {
	// Optional. The auto-accept setting for this group.
	AutoAccept *NetworkConnectivityGroupAutoAccept `pulumi:"autoAccept"`
	// An optional description of the group.
	Description *string `pulumi:"description"`
	// The name of the hub. Hub names must be unique. They use the following form:
	// projects/{projectNumber}/locations/global/hubs/{hubId}
	Hub string `pulumi:"hub"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
	Name                       *string                           `pulumi:"name"`
	NetworkConnectivityGroupId *string                           `pulumi:"networkConnectivityGroupId"`
	Project                    *string                           `pulumi:"project"`
	Timeouts                   *NetworkConnectivityGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkConnectivityGroup resource.
type NetworkConnectivityGroupArgs struct {
	// Optional. The auto-accept setting for this group.
	AutoAccept NetworkConnectivityGroupAutoAcceptPtrInput
	// An optional description of the group.
	Description pulumi.StringPtrInput
	// The name of the hub. Hub names must be unique. They use the following form:
	// projects/{projectNumber}/locations/global/hubs/{hubId}
	Hub pulumi.StringInput
	// Optional labels in key:value format. For more information about labels, see [Requirements for
	// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
	// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
	Name                       pulumi.StringPtrInput
	NetworkConnectivityGroupId pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	Timeouts                   NetworkConnectivityGroupTimeoutsPtrInput
}

func (NetworkConnectivityGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkConnectivityGroupArgs)(nil)).Elem()
}

type NetworkConnectivityGroupInput interface {
	pulumi.Input

	ToNetworkConnectivityGroupOutput() NetworkConnectivityGroupOutput
	ToNetworkConnectivityGroupOutputWithContext(ctx context.Context) NetworkConnectivityGroupOutput
}

func (*NetworkConnectivityGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectivityGroup)(nil)).Elem()
}

func (i *NetworkConnectivityGroup) ToNetworkConnectivityGroupOutput() NetworkConnectivityGroupOutput {
	return i.ToNetworkConnectivityGroupOutputWithContext(context.Background())
}

func (i *NetworkConnectivityGroup) ToNetworkConnectivityGroupOutputWithContext(ctx context.Context) NetworkConnectivityGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkConnectivityGroupOutput)
}

type NetworkConnectivityGroupOutput struct{ *pulumi.OutputState }

func (NetworkConnectivityGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkConnectivityGroup)(nil)).Elem()
}

func (o NetworkConnectivityGroupOutput) ToNetworkConnectivityGroupOutput() NetworkConnectivityGroupOutput {
	return o
}

func (o NetworkConnectivityGroupOutput) ToNetworkConnectivityGroupOutputWithContext(ctx context.Context) NetworkConnectivityGroupOutput {
	return o
}

// Optional. The auto-accept setting for this group.
func (o NetworkConnectivityGroupOutput) AutoAccept() NetworkConnectivityGroupAutoAcceptPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) NetworkConnectivityGroupAutoAcceptPtrOutput { return v.AutoAccept }).(NetworkConnectivityGroupAutoAcceptPtrOutput)
}

// Output only. The time the hub was created.
func (o NetworkConnectivityGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// An optional description of the group.
func (o NetworkConnectivityGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkConnectivityGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the hub. Hub names must be unique. They use the following form:
// projects/{projectNumber}/locations/global/hubs/{hubId}
func (o NetworkConnectivityGroupOutput) Hub() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.Hub }).(pulumi.StringOutput)
}

// Optional labels in key:value format. For more information about labels, see [Requirements for
// labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements). **Note**: This field is
// non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o NetworkConnectivityGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the group. Group names must be unique. Possible values: ["default", "center", "edge"]
func (o NetworkConnectivityGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkConnectivityGroupOutput) NetworkConnectivityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.NetworkConnectivityGroupId }).(pulumi.StringOutput)
}

func (o NetworkConnectivityGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The name of the route table that corresponds to this group. They use the following form:
// 'projects/{projectNumber}/locations/global/hubs/{hubId}/routeTables/{route_table_id}'
func (o NetworkConnectivityGroupOutput) RouteTable() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.RouteTable }).(pulumi.StringOutput)
}

// Output only. The current lifecycle state of this hub.
func (o NetworkConnectivityGroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkConnectivityGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkConnectivityGroupOutput) Timeouts() NetworkConnectivityGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) NetworkConnectivityGroupTimeoutsPtrOutput { return v.Timeouts }).(NetworkConnectivityGroupTimeoutsPtrOutput)
}

// Output only. The Google-generated UUID for the group. This value is unique across all group resources. If a group is
// deleted and another with the same name is created, the new route table is assigned a different uniqueId.
func (o NetworkConnectivityGroupOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time the hub was last updated.
func (o NetworkConnectivityGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkConnectivityGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkConnectivityGroupInput)(nil)).Elem(), &NetworkConnectivityGroup{})
	pulumi.RegisterOutputType(NetworkConnectivityGroupOutput{})
}
