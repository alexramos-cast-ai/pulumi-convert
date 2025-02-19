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

type TpuV2Vm struct {
	pulumi.CustomResourceState

	// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
	// neither is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorConfig TpuV2VmAcceleratorConfigPtrOutput `pulumi:"acceleratorConfig"`
	// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
	// is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorType pulumi.StringOutput `pulumi:"acceleratorType"`
	// The API version that created this Node.
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringOutput `pulumi:"cidrBlock"`
	// The additional data disks for the Node.
	DataDisks TpuV2VmDataDiskArrayOutput `pulumi:"dataDisks"`
	// Text description of the TPU.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The health status of the TPU node.
	Health pulumi.StringOutput `pulumi:"health"`
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription pulumi.StringOutput `pulumi:"healthDescription"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Whether the Node belongs to a Multislice group.
	MultisliceNode pulumi.BoolOutput `pulumi:"multisliceNode"`
	// The immutable name of the TPU.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configurations for the TPU node.
	NetworkConfig TpuV2VmNetworkConfigPtrOutput `pulumi:"networkConfig"`
	// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
	// node.
	NetworkConfigs TpuV2VmNetworkConfigArrayOutput `pulumi:"networkConfigs"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the
	// node reach out to the 0th entry in this map first.
	NetworkEndpoints TpuV2VmNetworkEndpointArrayOutput `pulumi:"networkEndpoints"`
	Project          pulumi.StringOutput               `pulumi:"project"`
	// The qualified name of the QueuedResource that requested this Node.
	QueuedResource pulumi.StringOutput `pulumi:"queuedResource"`
	// Runtime version for the TPU.
	RuntimeVersion pulumi.StringOutput `pulumi:"runtimeVersion"`
	// The scheduling options for this node.
	SchedulingConfig TpuV2VmSchedulingConfigPtrOutput `pulumi:"schedulingConfig"`
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
	// service account will be used.
	ServiceAccount TpuV2VmServiceAccountPtrOutput `pulumi:"serviceAccount"`
	// Shielded Instance options.
	ShieldedInstanceConfig TpuV2VmShieldedInstanceConfigPtrOutput `pulumi:"shieldedInstanceConfig"`
	// The current state for the TPU Node.
	State pulumi.StringOutput `pulumi:"state"`
	// The Symptoms that have occurred to the TPU Node.
	Symptoms TpuV2VmSymptomArrayOutput `pulumi:"symptoms"`
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput   `pulumi:"terraformLabels"`
	Timeouts        TpuV2VmTimeoutsPtrOutput `pulumi:"timeouts"`
	TpuV2VmId       pulumi.StringOutput      `pulumi:"tpuV2VmId"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTpuV2Vm registers a new resource with the given unique name, arguments, and options.
func NewTpuV2Vm(ctx *pulumi.Context,
	name string, args *TpuV2VmArgs, opts ...pulumi.ResourceOption) (*TpuV2Vm, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RuntimeVersion == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource TpuV2Vm
	err = ctx.RegisterPackageResource("google-beta:index/tpuV2Vm:TpuV2Vm", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTpuV2Vm gets an existing TpuV2Vm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTpuV2Vm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TpuV2VmState, opts ...pulumi.ResourceOption) (*TpuV2Vm, error) {
	var resource TpuV2Vm
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/tpuV2Vm:TpuV2Vm", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TpuV2Vm resources.
type tpuV2VmState struct {
	// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
	// neither is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorConfig *TpuV2VmAcceleratorConfig `pulumi:"acceleratorConfig"`
	// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
	// is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The API version that created this Node.
	ApiVersion *string `pulumi:"apiVersion"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The additional data disks for the Node.
	DataDisks []TpuV2VmDataDisk `pulumi:"dataDisks"`
	// Text description of the TPU.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The health status of the TPU node.
	Health *string `pulumi:"health"`
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription *string `pulumi:"healthDescription"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata map[string]string `pulumi:"metadata"`
	// Whether the Node belongs to a Multislice group.
	MultisliceNode *bool `pulumi:"multisliceNode"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// Network configurations for the TPU node.
	NetworkConfig *TpuV2VmNetworkConfig `pulumi:"networkConfig"`
	// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
	// node.
	NetworkConfigs []TpuV2VmNetworkConfig `pulumi:"networkConfigs"`
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the
	// node reach out to the 0th entry in this map first.
	NetworkEndpoints []TpuV2VmNetworkEndpoint `pulumi:"networkEndpoints"`
	Project          *string                  `pulumi:"project"`
	// The qualified name of the QueuedResource that requested this Node.
	QueuedResource *string `pulumi:"queuedResource"`
	// Runtime version for the TPU.
	RuntimeVersion *string `pulumi:"runtimeVersion"`
	// The scheduling options for this node.
	SchedulingConfig *TpuV2VmSchedulingConfig `pulumi:"schedulingConfig"`
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
	// service account will be used.
	ServiceAccount *TpuV2VmServiceAccount `pulumi:"serviceAccount"`
	// Shielded Instance options.
	ShieldedInstanceConfig *TpuV2VmShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// The current state for the TPU Node.
	State *string `pulumi:"state"`
	// The Symptoms that have occurred to the TPU Node.
	Symptoms []TpuV2VmSymptom `pulumi:"symptoms"`
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags []string `pulumi:"tags"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	Timeouts        *TpuV2VmTimeouts  `pulumi:"timeouts"`
	TpuV2VmId       *string           `pulumi:"tpuV2VmId"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type TpuV2VmState struct {
	// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
	// neither is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorConfig TpuV2VmAcceleratorConfigPtrInput
	// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
	// is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorType pulumi.StringPtrInput
	// The API version that created this Node.
	ApiVersion pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The additional data disks for the Node.
	DataDisks TpuV2VmDataDiskArrayInput
	// Text description of the TPU.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// The health status of the TPU node.
	Health pulumi.StringPtrInput
	// If this field is populated, it contains a description of why the TPU Node is unhealthy.
	HealthDescription pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata pulumi.StringMapInput
	// Whether the Node belongs to a Multislice group.
	MultisliceNode pulumi.BoolPtrInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// Network configurations for the TPU node.
	NetworkConfig TpuV2VmNetworkConfigPtrInput
	// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
	// node.
	NetworkConfigs TpuV2VmNetworkConfigArrayInput
	// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the
	// node reach out to the 0th entry in this map first.
	NetworkEndpoints TpuV2VmNetworkEndpointArrayInput
	Project          pulumi.StringPtrInput
	// The qualified name of the QueuedResource that requested this Node.
	QueuedResource pulumi.StringPtrInput
	// Runtime version for the TPU.
	RuntimeVersion pulumi.StringPtrInput
	// The scheduling options for this node.
	SchedulingConfig TpuV2VmSchedulingConfigPtrInput
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
	// service account will be used.
	ServiceAccount TpuV2VmServiceAccountPtrInput
	// Shielded Instance options.
	ShieldedInstanceConfig TpuV2VmShieldedInstanceConfigPtrInput
	// The current state for the TPU Node.
	State pulumi.StringPtrInput
	// The Symptoms that have occurred to the TPU Node.
	Symptoms TpuV2VmSymptomArrayInput
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags pulumi.StringArrayInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        TpuV2VmTimeoutsPtrInput
	TpuV2VmId       pulumi.StringPtrInput
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuV2VmState) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuV2VmState)(nil)).Elem()
}

type tpuV2VmArgs struct {
	// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
	// neither is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorConfig *TpuV2VmAcceleratorConfig `pulumi:"acceleratorConfig"`
	// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
	// is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorType *string `pulumi:"acceleratorType"`
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock *string `pulumi:"cidrBlock"`
	// The additional data disks for the Node.
	DataDisks []TpuV2VmDataDisk `pulumi:"dataDisks"`
	// Text description of the TPU.
	Description *string `pulumi:"description"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata map[string]string `pulumi:"metadata"`
	// The immutable name of the TPU.
	Name *string `pulumi:"name"`
	// Network configurations for the TPU node.
	NetworkConfig *TpuV2VmNetworkConfig `pulumi:"networkConfig"`
	// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
	// node.
	NetworkConfigs []TpuV2VmNetworkConfig `pulumi:"networkConfigs"`
	Project        *string                `pulumi:"project"`
	// Runtime version for the TPU.
	RuntimeVersion string `pulumi:"runtimeVersion"`
	// The scheduling options for this node.
	SchedulingConfig *TpuV2VmSchedulingConfig `pulumi:"schedulingConfig"`
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
	// service account will be used.
	ServiceAccount *TpuV2VmServiceAccount `pulumi:"serviceAccount"`
	// Shielded Instance options.
	ShieldedInstanceConfig *TpuV2VmShieldedInstanceConfig `pulumi:"shieldedInstanceConfig"`
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags      []string         `pulumi:"tags"`
	Timeouts  *TpuV2VmTimeouts `pulumi:"timeouts"`
	TpuV2VmId *string          `pulumi:"tpuV2VmId"`
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TpuV2Vm resource.
type TpuV2VmArgs struct {
	// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
	// neither is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorConfig TpuV2VmAcceleratorConfigPtrInput
	// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
	// is specified, 'accelerator_type' defaults to 'v2-8'.
	AcceleratorType pulumi.StringPtrInput
	// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
	// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
	// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
	// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
	// that is using that CIDR block.
	CidrBlock pulumi.StringPtrInput
	// The additional data disks for the Node.
	DataDisks TpuV2VmDataDiskArrayInput
	// Text description of the TPU.
	Description pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
	Metadata pulumi.StringMapInput
	// The immutable name of the TPU.
	Name pulumi.StringPtrInput
	// Network configurations for the TPU node.
	NetworkConfig TpuV2VmNetworkConfigPtrInput
	// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
	// node.
	NetworkConfigs TpuV2VmNetworkConfigArrayInput
	Project        pulumi.StringPtrInput
	// Runtime version for the TPU.
	RuntimeVersion pulumi.StringInput
	// The scheduling options for this node.
	SchedulingConfig TpuV2VmSchedulingConfigPtrInput
	// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
	// service account will be used.
	ServiceAccount TpuV2VmServiceAccountPtrInput
	// Shielded Instance options.
	ShieldedInstanceConfig TpuV2VmShieldedInstanceConfigPtrInput
	// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
	Tags      pulumi.StringArrayInput
	Timeouts  TpuV2VmTimeoutsPtrInput
	TpuV2VmId pulumi.StringPtrInput
	// The GCP location for the TPU. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuV2VmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuV2VmArgs)(nil)).Elem()
}

type TpuV2VmInput interface {
	pulumi.Input

	ToTpuV2VmOutput() TpuV2VmOutput
	ToTpuV2VmOutputWithContext(ctx context.Context) TpuV2VmOutput
}

func (*TpuV2Vm) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuV2Vm)(nil)).Elem()
}

func (i *TpuV2Vm) ToTpuV2VmOutput() TpuV2VmOutput {
	return i.ToTpuV2VmOutputWithContext(context.Background())
}

func (i *TpuV2Vm) ToTpuV2VmOutputWithContext(ctx context.Context) TpuV2VmOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TpuV2VmOutput)
}

type TpuV2VmOutput struct{ *pulumi.OutputState }

func (TpuV2VmOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuV2Vm)(nil)).Elem()
}

func (o TpuV2VmOutput) ToTpuV2VmOutput() TpuV2VmOutput {
	return o
}

func (o TpuV2VmOutput) ToTpuV2VmOutputWithContext(ctx context.Context) TpuV2VmOutput {
	return o
}

// The AccleratorConfig for the TPU Node. 'accelerator_config' cannot be used at the same time as 'accelerator_type'. If
// neither is specified, 'accelerator_type' defaults to 'v2-8'.
func (o TpuV2VmOutput) AcceleratorConfig() TpuV2VmAcceleratorConfigPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmAcceleratorConfigPtrOutput { return v.AcceleratorConfig }).(TpuV2VmAcceleratorConfigPtrOutput)
}

// TPU accelerator type for the TPU. 'accelerator_type' cannot be used at the same time as 'accelerator_config'. If neither
// is specified, 'accelerator_type' defaults to 'v2-8'.
func (o TpuV2VmOutput) AcceleratorType() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.AcceleratorType }).(pulumi.StringOutput)
}

// The API version that created this Node.
func (o TpuV2VmOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// The CIDR block that the TPU node will use when selecting an IP address. This CIDR block must be a /29 block; the Compute
// Engine networks API forbids a smaller block, and using a larger block would be wasteful (a node can only consume one IP
// address). Errors will occur if the CIDR block has already been used for a currently existing TPU node, the CIDR block
// conflicts with any subnetworks in the user's provided network, or the provided network is peered with another network
// that is using that CIDR block.
func (o TpuV2VmOutput) CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.CidrBlock }).(pulumi.StringOutput)
}

// The additional data disks for the Node.
func (o TpuV2VmOutput) DataDisks() TpuV2VmDataDiskArrayOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmDataDiskArrayOutput { return v.DataDisks }).(TpuV2VmDataDiskArrayOutput)
}

// Text description of the TPU.
func (o TpuV2VmOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o TpuV2VmOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The health status of the TPU node.
func (o TpuV2VmOutput) Health() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.Health }).(pulumi.StringOutput)
}

// If this field is populated, it contains a description of why the TPU Node is unhealthy.
func (o TpuV2VmOutput) HealthDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.HealthDescription }).(pulumi.StringOutput)
}

// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o TpuV2VmOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Custom metadata to apply to the TPU Node. Can set startup-script and shutdown-script.
func (o TpuV2VmOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Whether the Node belongs to a Multislice group.
func (o TpuV2VmOutput) MultisliceNode() pulumi.BoolOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.BoolOutput { return v.MultisliceNode }).(pulumi.BoolOutput)
}

// The immutable name of the TPU.
func (o TpuV2VmOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network configurations for the TPU node.
func (o TpuV2VmOutput) NetworkConfig() TpuV2VmNetworkConfigPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmNetworkConfigPtrOutput { return v.NetworkConfig }).(TpuV2VmNetworkConfigPtrOutput)
}

// Repeated network configurations for the TPU node. This field is used to specify multiple network configs for the TPU
// node.
func (o TpuV2VmOutput) NetworkConfigs() TpuV2VmNetworkConfigArrayOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmNetworkConfigArrayOutput { return v.NetworkConfigs }).(TpuV2VmNetworkConfigArrayOutput)
}

// The network endpoints where TPU workers can be accessed and sent work. It is recommended that runtime clients of the
// node reach out to the 0th entry in this map first.
func (o TpuV2VmOutput) NetworkEndpoints() TpuV2VmNetworkEndpointArrayOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmNetworkEndpointArrayOutput { return v.NetworkEndpoints }).(TpuV2VmNetworkEndpointArrayOutput)
}

func (o TpuV2VmOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The qualified name of the QueuedResource that requested this Node.
func (o TpuV2VmOutput) QueuedResource() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.QueuedResource }).(pulumi.StringOutput)
}

// Runtime version for the TPU.
func (o TpuV2VmOutput) RuntimeVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.RuntimeVersion }).(pulumi.StringOutput)
}

// The scheduling options for this node.
func (o TpuV2VmOutput) SchedulingConfig() TpuV2VmSchedulingConfigPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmSchedulingConfigPtrOutput { return v.SchedulingConfig }).(TpuV2VmSchedulingConfigPtrOutput)
}

// The Google Cloud Platform Service Account to be used by the TPU node VMs. If None is specified, the default compute
// service account will be used.
func (o TpuV2VmOutput) ServiceAccount() TpuV2VmServiceAccountPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmServiceAccountPtrOutput { return v.ServiceAccount }).(TpuV2VmServiceAccountPtrOutput)
}

// Shielded Instance options.
func (o TpuV2VmOutput) ShieldedInstanceConfig() TpuV2VmShieldedInstanceConfigPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmShieldedInstanceConfigPtrOutput { return v.ShieldedInstanceConfig }).(TpuV2VmShieldedInstanceConfigPtrOutput)
}

// The current state for the TPU Node.
func (o TpuV2VmOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The Symptoms that have occurred to the TPU Node.
func (o TpuV2VmOutput) Symptoms() TpuV2VmSymptomArrayOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmSymptomArrayOutput { return v.Symptoms }).(TpuV2VmSymptomArrayOutput)
}

// Tags to apply to the TPU Node. Tags are used to identify valid sources or targets for network firewalls.
func (o TpuV2VmOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o TpuV2VmOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o TpuV2VmOutput) Timeouts() TpuV2VmTimeoutsPtrOutput {
	return o.ApplyT(func(v *TpuV2Vm) TpuV2VmTimeoutsPtrOutput { return v.Timeouts }).(TpuV2VmTimeoutsPtrOutput)
}

func (o TpuV2VmOutput) TpuV2VmId() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.TpuV2VmId }).(pulumi.StringOutput)
}

// The GCP location for the TPU. If it is not provided, the provider zone is used.
func (o TpuV2VmOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2Vm) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TpuV2VmInput)(nil)).Elem(), &TpuV2Vm{})
	pulumi.RegisterOutputType(TpuV2VmOutput{})
}
