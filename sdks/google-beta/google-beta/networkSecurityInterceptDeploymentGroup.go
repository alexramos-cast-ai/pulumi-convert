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

type NetworkSecurityInterceptDeploymentGroup struct {
	pulumi.CustomResourceState

	// Output only. The list of Intercept Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroupArrayOutput `pulumi:"connectedEndpointGroups"`
	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// intercept_deployment_group_id from the method_signature of Create RPC
	InterceptDeploymentGroupId pulumi.StringOutput `pulumi:"interceptDeploymentGroupId"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. Identifier. Then name of the InterceptDeploymentGroup.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringOutput `pulumi:"network"`
	NetworkSecurityInterceptDeploymentGroupId pulumi.StringOutput `pulumi:"networkSecurityInterceptDeploymentGroupId"`
	Project                                   pulumi.StringOutput `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                   `pulumi:"terraformLabels"`
	Timeouts        NetworkSecurityInterceptDeploymentGroupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewNetworkSecurityInterceptDeploymentGroup registers a new resource with the given unique name, arguments, and options.
func NewNetworkSecurityInterceptDeploymentGroup(ctx *pulumi.Context,
	name string, args *NetworkSecurityInterceptDeploymentGroupArgs, opts ...pulumi.ResourceOption) (*NetworkSecurityInterceptDeploymentGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InterceptDeploymentGroupId == nil {
		return nil, errors.New("invalid value for required argument 'InterceptDeploymentGroupId'")
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
	var resource NetworkSecurityInterceptDeploymentGroup
	err = ctx.RegisterPackageResource("google-beta:index/networkSecurityInterceptDeploymentGroup:NetworkSecurityInterceptDeploymentGroup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkSecurityInterceptDeploymentGroup gets an existing NetworkSecurityInterceptDeploymentGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkSecurityInterceptDeploymentGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkSecurityInterceptDeploymentGroupState, opts ...pulumi.ResourceOption) (*NetworkSecurityInterceptDeploymentGroup, error) {
	var resource NetworkSecurityInterceptDeploymentGroup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkSecurityInterceptDeploymentGroup:NetworkSecurityInterceptDeploymentGroup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkSecurityInterceptDeploymentGroup resources.
type networkSecurityInterceptDeploymentGroupState struct {
	// Output only. The list of Intercept Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups []NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroup `pulumi:"connectedEndpointGroups"`
	// Output only. [Output only] Create time stamp
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// intercept_deployment_group_id from the method_signature of Create RPC
	InterceptDeploymentGroupId *string `pulumi:"interceptDeploymentGroupId"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
	Location *string `pulumi:"location"`
	// Output only. Identifier. Then name of the InterceptDeploymentGroup.
	Name *string `pulumi:"name"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   *string `pulumi:"network"`
	NetworkSecurityInterceptDeploymentGroupId *string `pulumi:"networkSecurityInterceptDeploymentGroupId"`
	Project                                   *string `pulumi:"project"`
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                `pulumi:"terraformLabels"`
	Timeouts        *NetworkSecurityInterceptDeploymentGroupTimeouts `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp
	UpdateTime *string `pulumi:"updateTime"`
}

type NetworkSecurityInterceptDeploymentGroupState struct {
	// Output only. The list of Intercept Endpoint Groups that are connected to this resource.
	ConnectedEndpointGroups NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroupArrayInput
	// Output only. [Output only] Create time stamp
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// intercept_deployment_group_id from the method_signature of Create RPC
	InterceptDeploymentGroupId pulumi.StringPtrInput
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
	Location pulumi.StringPtrInput
	// Output only. Identifier. Then name of the InterceptDeploymentGroup.
	Name pulumi.StringPtrInput
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringPtrInput
	NetworkSecurityInterceptDeploymentGroupId pulumi.StringPtrInput
	Project                                   pulumi.StringPtrInput
	// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
	Reconciling pulumi.BoolPtrInput
	// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkSecurityInterceptDeploymentGroupTimeoutsPtrInput
	// Output only. [Output only] Update time stamp
	UpdateTime pulumi.StringPtrInput
}

func (NetworkSecurityInterceptDeploymentGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityInterceptDeploymentGroupState)(nil)).Elem()
}

type networkSecurityInterceptDeploymentGroupArgs struct {
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// intercept_deployment_group_id from the method_signature of Create RPC
	InterceptDeploymentGroupId string `pulumi:"interceptDeploymentGroupId"`
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
	Location string `pulumi:"location"`
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   string                                           `pulumi:"network"`
	NetworkSecurityInterceptDeploymentGroupId *string                                          `pulumi:"networkSecurityInterceptDeploymentGroupId"`
	Project                                   *string                                          `pulumi:"project"`
	Timeouts                                  *NetworkSecurityInterceptDeploymentGroupTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a NetworkSecurityInterceptDeploymentGroup resource.
type NetworkSecurityInterceptDeploymentGroupArgs struct {
	// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
	// intercept_deployment_group_id from the method_signature of Create RPC
	InterceptDeploymentGroupId pulumi.StringInput
	// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122. See documentation for resource type
	// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
	Location pulumi.StringInput
	// Required. Immutable. The network that is being used for the deployment. Format is:
	// projects/{project}/global/networks/{network}.
	Network                                   pulumi.StringInput
	NetworkSecurityInterceptDeploymentGroupId pulumi.StringPtrInput
	Project                                   pulumi.StringPtrInput
	Timeouts                                  NetworkSecurityInterceptDeploymentGroupTimeoutsPtrInput
}

func (NetworkSecurityInterceptDeploymentGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkSecurityInterceptDeploymentGroupArgs)(nil)).Elem()
}

type NetworkSecurityInterceptDeploymentGroupInput interface {
	pulumi.Input

	ToNetworkSecurityInterceptDeploymentGroupOutput() NetworkSecurityInterceptDeploymentGroupOutput
	ToNetworkSecurityInterceptDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityInterceptDeploymentGroupOutput
}

func (*NetworkSecurityInterceptDeploymentGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityInterceptDeploymentGroup)(nil)).Elem()
}

func (i *NetworkSecurityInterceptDeploymentGroup) ToNetworkSecurityInterceptDeploymentGroupOutput() NetworkSecurityInterceptDeploymentGroupOutput {
	return i.ToNetworkSecurityInterceptDeploymentGroupOutputWithContext(context.Background())
}

func (i *NetworkSecurityInterceptDeploymentGroup) ToNetworkSecurityInterceptDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityInterceptDeploymentGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkSecurityInterceptDeploymentGroupOutput)
}

type NetworkSecurityInterceptDeploymentGroupOutput struct{ *pulumi.OutputState }

func (NetworkSecurityInterceptDeploymentGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkSecurityInterceptDeploymentGroup)(nil)).Elem()
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) ToNetworkSecurityInterceptDeploymentGroupOutput() NetworkSecurityInterceptDeploymentGroupOutput {
	return o
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) ToNetworkSecurityInterceptDeploymentGroupOutputWithContext(ctx context.Context) NetworkSecurityInterceptDeploymentGroupOutput {
	return o
}

// Output only. The list of Intercept Endpoint Groups that are connected to this resource.
func (o NetworkSecurityInterceptDeploymentGroupOutput) ConnectedEndpointGroups() NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroupArrayOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroupArrayOutput {
		return v.ConnectedEndpointGroups
	}).(NetworkSecurityInterceptDeploymentGroupConnectedEndpointGroupArrayOutput)
}

// Output only. [Output only] Create time stamp
func (o NetworkSecurityInterceptDeploymentGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Required. Id of the requesting object If auto-generating Id server-side, remove this field and
// intercept_deployment_group_id from the method_signature of Create RPC
func (o NetworkSecurityInterceptDeploymentGroupOutput) InterceptDeploymentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput {
		return v.InterceptDeploymentGroupId
	}).(pulumi.StringOutput)
}

// Optional. Labels as key value pairs **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o NetworkSecurityInterceptDeploymentGroupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122. See documentation for resource type
// 'networksecurity.googleapis.com/InterceptDeploymentGroup'.
func (o NetworkSecurityInterceptDeploymentGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. Identifier. Then name of the InterceptDeploymentGroup.
func (o NetworkSecurityInterceptDeploymentGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Required. Immutable. The network that is being used for the deployment. Format is:
// projects/{project}/global/networks/{network}.
func (o NetworkSecurityInterceptDeploymentGroupOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) NetworkSecurityInterceptDeploymentGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput {
		return v.NetworkSecurityInterceptDeploymentGroupId
	}).(pulumi.StringOutput)
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Whether reconciling is in progress, recommended per https://google.aip.dev/128.
func (o NetworkSecurityInterceptDeploymentGroupOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. Current state of the deployment group. Possible values: STATE_UNSPECIFIED ACTIVE CREATING DELETING
func (o NetworkSecurityInterceptDeploymentGroupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkSecurityInterceptDeploymentGroupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkSecurityInterceptDeploymentGroupOutput) Timeouts() NetworkSecurityInterceptDeploymentGroupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) NetworkSecurityInterceptDeploymentGroupTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkSecurityInterceptDeploymentGroupTimeoutsPtrOutput)
}

// Output only. [Output only] Update time stamp
func (o NetworkSecurityInterceptDeploymentGroupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkSecurityInterceptDeploymentGroup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkSecurityInterceptDeploymentGroupInput)(nil)).Elem(), &NetworkSecurityInterceptDeploymentGroup{})
	pulumi.RegisterOutputType(NetworkSecurityInterceptDeploymentGroupOutput{})
}
