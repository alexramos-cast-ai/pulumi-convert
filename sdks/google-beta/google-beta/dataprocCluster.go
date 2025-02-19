// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataprocCluster struct {
	pulumi.CustomResourceState

	// Allows you to configure various aspects of the cluster.
	ClusterConfig               DataprocClusterClusterConfigPtrOutput `pulumi:"clusterConfig"`
	DataprocClusterId           pulumi.StringOutput                   `pulumi:"dataprocClusterId"`
	EffectiveLabels             pulumi.StringMapOutput                `pulumi:"effectiveLabels"`
	GracefulDecommissionTimeout pulumi.StringPtrOutput                `pulumi:"gracefulDecommissionTimeout"`
	// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	// to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the cluster, unique within the project and zone.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput           `pulumi:"terraformLabels"`
	Timeouts        DataprocClusterTimeoutsPtrOutput `pulumi:"timeouts"`
	// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
	// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
	// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
	VirtualClusterConfig DataprocClusterVirtualClusterConfigPtrOutput `pulumi:"virtualClusterConfig"`
}

// NewDataprocCluster registers a new resource with the given unique name, arguments, and options.
func NewDataprocCluster(ctx *pulumi.Context,
	name string, args *DataprocClusterArgs, opts ...pulumi.ResourceOption) (*DataprocCluster, error) {
	if args == nil {
		args = &DataprocClusterArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocCluster
	err = ctx.RegisterPackageResource("google-beta:index/dataprocCluster:DataprocCluster", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocCluster gets an existing DataprocCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocClusterState, opts ...pulumi.ResourceOption) (*DataprocCluster, error) {
	var resource DataprocCluster
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataprocCluster:DataprocCluster", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocCluster resources.
type dataprocClusterState struct {
	// Allows you to configure various aspects of the cluster.
	ClusterConfig               *DataprocClusterClusterConfig `pulumi:"clusterConfig"`
	DataprocClusterId           *string                       `pulumi:"dataprocClusterId"`
	EffectiveLabels             map[string]string             `pulumi:"effectiveLabels"`
	GracefulDecommissionTimeout *string                       `pulumi:"gracefulDecommissionTimeout"`
	// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	// to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	Region *string `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string        `pulumi:"terraformLabels"`
	Timeouts        *DataprocClusterTimeouts `pulumi:"timeouts"`
	// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
	// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
	// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
	VirtualClusterConfig *DataprocClusterVirtualClusterConfig `pulumi:"virtualClusterConfig"`
}

type DataprocClusterState struct {
	// Allows you to configure various aspects of the cluster.
	ClusterConfig               DataprocClusterClusterConfigPtrInput
	DataprocClusterId           pulumi.StringPtrInput
	EffectiveLabels             pulumi.StringMapInput
	GracefulDecommissionTimeout pulumi.StringPtrInput
	// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	// to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	Region pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataprocClusterTimeoutsPtrInput
	// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
	// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
	// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
	VirtualClusterConfig DataprocClusterVirtualClusterConfigPtrInput
}

func (DataprocClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocClusterState)(nil)).Elem()
}

type dataprocClusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	ClusterConfig               *DataprocClusterClusterConfig `pulumi:"clusterConfig"`
	DataprocClusterId           *string                       `pulumi:"dataprocClusterId"`
	GracefulDecommissionTimeout *string                       `pulumi:"gracefulDecommissionTimeout"`
	// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	// to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the cluster, unique within the project and zone.
	Name *string `pulumi:"name"`
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	Region   *string                  `pulumi:"region"`
	Timeouts *DataprocClusterTimeouts `pulumi:"timeouts"`
	// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
	// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
	// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
	VirtualClusterConfig *DataprocClusterVirtualClusterConfig `pulumi:"virtualClusterConfig"`
}

// The set of arguments for constructing a DataprocCluster resource.
type DataprocClusterArgs struct {
	// Allows you to configure various aspects of the cluster.
	ClusterConfig               DataprocClusterClusterConfigPtrInput
	DataprocClusterId           pulumi.StringPtrInput
	GracefulDecommissionTimeout pulumi.StringPtrInput
	// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
	// to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the cluster, unique within the project and zone.
	Name pulumi.StringPtrInput
	// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the cluster and associated nodes will be created in. Defaults to global.
	Region   pulumi.StringPtrInput
	Timeouts DataprocClusterTimeoutsPtrInput
	// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
	// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
	// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
	VirtualClusterConfig DataprocClusterVirtualClusterConfigPtrInput
}

func (DataprocClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocClusterArgs)(nil)).Elem()
}

type DataprocClusterInput interface {
	pulumi.Input

	ToDataprocClusterOutput() DataprocClusterOutput
	ToDataprocClusterOutputWithContext(ctx context.Context) DataprocClusterOutput
}

func (*DataprocCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocCluster)(nil)).Elem()
}

func (i *DataprocCluster) ToDataprocClusterOutput() DataprocClusterOutput {
	return i.ToDataprocClusterOutputWithContext(context.Background())
}

func (i *DataprocCluster) ToDataprocClusterOutputWithContext(ctx context.Context) DataprocClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocClusterOutput)
}

type DataprocClusterOutput struct{ *pulumi.OutputState }

func (DataprocClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocCluster)(nil)).Elem()
}

func (o DataprocClusterOutput) ToDataprocClusterOutput() DataprocClusterOutput {
	return o
}

func (o DataprocClusterOutput) ToDataprocClusterOutputWithContext(ctx context.Context) DataprocClusterOutput {
	return o
}

// Allows you to configure various aspects of the cluster.
func (o DataprocClusterOutput) ClusterConfig() DataprocClusterClusterConfigPtrOutput {
	return o.ApplyT(func(v *DataprocCluster) DataprocClusterClusterConfigPtrOutput { return v.ClusterConfig }).(DataprocClusterClusterConfigPtrOutput)
}

func (o DataprocClusterOutput) DataprocClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringOutput { return v.DataprocClusterId }).(pulumi.StringOutput)
}

func (o DataprocClusterOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o DataprocClusterOutput) GracefulDecommissionTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringPtrOutput { return v.GracefulDecommissionTimeout }).(pulumi.StringPtrOutput)
}

// The list of the labels (key/value pairs) configured on the resource and to be applied to instances in the cluster.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration. Please refer
// to the field 'effective_labels' for all of the labels present on the resource.
func (o DataprocClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the cluster, unique within the project and zone.
func (o DataprocClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the cluster will exist. If it is not provided, the provider project is used.
func (o DataprocClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region in which the cluster and associated nodes will be created in. Defaults to global.
func (o DataprocClusterOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataprocClusterOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocCluster) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataprocClusterOutput) Timeouts() DataprocClusterTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataprocCluster) DataprocClusterTimeoutsPtrOutput { return v.Timeouts }).(DataprocClusterTimeoutsPtrOutput)
}

// The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying
// compute resources, for example, when creating a Dataproc-on-GKE cluster. Dataproc may set default values, and values may
// change when clusters are updated. Exactly one of config or virtualClusterConfig must be specified.
func (o DataprocClusterOutput) VirtualClusterConfig() DataprocClusterVirtualClusterConfigPtrOutput {
	return o.ApplyT(func(v *DataprocCluster) DataprocClusterVirtualClusterConfigPtrOutput { return v.VirtualClusterConfig }).(DataprocClusterVirtualClusterConfigPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocClusterInput)(nil)).Elem(), &DataprocCluster{})
	pulumi.RegisterOutputType(DataprocClusterOutput{})
}
