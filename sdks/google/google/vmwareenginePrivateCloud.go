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

type VmwareenginePrivateCloud struct {
	pulumi.CustomResourceState

	// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
	// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	DeletionDelayHours pulumi.Float64PtrOutput `pulumi:"deletionDelayHours"`
	// User-provided description for this private cloud.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Details about a HCX Cloud Manager appliance.
	Hcxes VmwareenginePrivateCloudHcxArrayOutput `pulumi:"hcxes"`
	// The location where the PrivateCloud should reside.
	Location pulumi.StringOutput `pulumi:"location"`
	// The management cluster for this private cloud. This used for creating and managing the default cluster.
	ManagementCluster VmwareenginePrivateCloudManagementClusterOutput `pulumi:"managementCluster"`
	// The ID of the PrivateCloud.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration in the consumer project with which the peering has to be done.
	NetworkConfig VmwareenginePrivateCloudNetworkConfigOutput `pulumi:"networkConfig"`
	// Details about a NSX Manager appliance.
	Nsxes   VmwareenginePrivateCloudNsxArrayOutput `pulumi:"nsxes"`
	Project pulumi.StringOutput                    `pulumi:"project"`
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
	// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
	// deletion_delay_hours.
	SendDeletionDelayHoursIfZero pulumi.BoolPtrOutput `pulumi:"sendDeletionDelayHoursIfZero"`
	// State of the resource. New values may be added to this enum when appropriate.
	State    pulumi.StringOutput                       `pulumi:"state"`
	Timeouts VmwareenginePrivateCloudTimeoutsPtrOutput `pulumi:"timeouts"`
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// System-generated unique identifier for the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Details about a vCenter Server management appliance.
	Vcenters                   VmwareenginePrivateCloudVcenterArrayOutput `pulumi:"vcenters"`
	VmwareenginePrivateCloudId pulumi.StringOutput                        `pulumi:"vmwareenginePrivateCloudId"`
}

// NewVmwareenginePrivateCloud registers a new resource with the given unique name, arguments, and options.
func NewVmwareenginePrivateCloud(ctx *pulumi.Context,
	name string, args *VmwareenginePrivateCloudArgs, opts ...pulumi.ResourceOption) (*VmwareenginePrivateCloud, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ManagementCluster == nil {
		return nil, errors.New("invalid value for required argument 'ManagementCluster'")
	}
	if args.NetworkConfig == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfig'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VmwareenginePrivateCloud
	err = ctx.RegisterPackageResource("google:index/vmwareenginePrivateCloud:VmwareenginePrivateCloud", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmwareenginePrivateCloud gets an existing VmwareenginePrivateCloud resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmwareenginePrivateCloud(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmwareenginePrivateCloudState, opts ...pulumi.ResourceOption) (*VmwareenginePrivateCloud, error) {
	var resource VmwareenginePrivateCloud
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/vmwareenginePrivateCloud:VmwareenginePrivateCloud", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmwareenginePrivateCloud resources.
type vmwareenginePrivateCloudState struct {
	// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
	// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	DeletionDelayHours *float64 `pulumi:"deletionDelayHours"`
	// User-provided description for this private cloud.
	Description *string `pulumi:"description"`
	// Details about a HCX Cloud Manager appliance.
	Hcxes []VmwareenginePrivateCloudHcx `pulumi:"hcxes"`
	// The location where the PrivateCloud should reside.
	Location *string `pulumi:"location"`
	// The management cluster for this private cloud. This used for creating and managing the default cluster.
	ManagementCluster *VmwareenginePrivateCloudManagementCluster `pulumi:"managementCluster"`
	// The ID of the PrivateCloud.
	Name *string `pulumi:"name"`
	// Network configuration in the consumer project with which the peering has to be done.
	NetworkConfig *VmwareenginePrivateCloudNetworkConfig `pulumi:"networkConfig"`
	// Details about a NSX Manager appliance.
	Nsxes   []VmwareenginePrivateCloudNsx `pulumi:"nsxes"`
	Project *string                       `pulumi:"project"`
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
	// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
	// deletion_delay_hours.
	SendDeletionDelayHoursIfZero *bool `pulumi:"sendDeletionDelayHoursIfZero"`
	// State of the resource. New values may be added to this enum when appropriate.
	State    *string                           `pulumi:"state"`
	Timeouts *VmwareenginePrivateCloudTimeouts `pulumi:"timeouts"`
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
	Type *string `pulumi:"type"`
	// System-generated unique identifier for the resource.
	Uid *string `pulumi:"uid"`
	// Details about a vCenter Server management appliance.
	Vcenters                   []VmwareenginePrivateCloudVcenter `pulumi:"vcenters"`
	VmwareenginePrivateCloudId *string                           `pulumi:"vmwareenginePrivateCloudId"`
}

type VmwareenginePrivateCloudState struct {
	// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
	// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	DeletionDelayHours pulumi.Float64PtrInput
	// User-provided description for this private cloud.
	Description pulumi.StringPtrInput
	// Details about a HCX Cloud Manager appliance.
	Hcxes VmwareenginePrivateCloudHcxArrayInput
	// The location where the PrivateCloud should reside.
	Location pulumi.StringPtrInput
	// The management cluster for this private cloud. This used for creating and managing the default cluster.
	ManagementCluster VmwareenginePrivateCloudManagementClusterPtrInput
	// The ID of the PrivateCloud.
	Name pulumi.StringPtrInput
	// Network configuration in the consumer project with which the peering has to be done.
	NetworkConfig VmwareenginePrivateCloudNetworkConfigPtrInput
	// Details about a NSX Manager appliance.
	Nsxes   VmwareenginePrivateCloudNsxArrayInput
	Project pulumi.StringPtrInput
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
	// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
	// deletion_delay_hours.
	SendDeletionDelayHoursIfZero pulumi.BoolPtrInput
	// State of the resource. New values may be added to this enum when appropriate.
	State    pulumi.StringPtrInput
	Timeouts VmwareenginePrivateCloudTimeoutsPtrInput
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
	Type pulumi.StringPtrInput
	// System-generated unique identifier for the resource.
	Uid pulumi.StringPtrInput
	// Details about a vCenter Server management appliance.
	Vcenters                   VmwareenginePrivateCloudVcenterArrayInput
	VmwareenginePrivateCloudId pulumi.StringPtrInput
}

func (VmwareenginePrivateCloudState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareenginePrivateCloudState)(nil)).Elem()
}

type vmwareenginePrivateCloudArgs struct {
	// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
	// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	DeletionDelayHours *float64 `pulumi:"deletionDelayHours"`
	// User-provided description for this private cloud.
	Description *string `pulumi:"description"`
	// The location where the PrivateCloud should reside.
	Location string `pulumi:"location"`
	// The management cluster for this private cloud. This used for creating and managing the default cluster.
	ManagementCluster VmwareenginePrivateCloudManagementCluster `pulumi:"managementCluster"`
	// The ID of the PrivateCloud.
	Name *string `pulumi:"name"`
	// Network configuration in the consumer project with which the peering has to be done.
	NetworkConfig VmwareenginePrivateCloudNetworkConfig `pulumi:"networkConfig"`
	Project       *string                               `pulumi:"project"`
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
	// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
	// deletion_delay_hours.
	SendDeletionDelayHoursIfZero *bool                             `pulumi:"sendDeletionDelayHoursIfZero"`
	Timeouts                     *VmwareenginePrivateCloudTimeouts `pulumi:"timeouts"`
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
	Type                       *string `pulumi:"type"`
	VmwareenginePrivateCloudId *string `pulumi:"vmwareenginePrivateCloudId"`
}

// The set of arguments for constructing a VmwareenginePrivateCloud resource.
type VmwareenginePrivateCloudArgs struct {
	// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
	// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
	DeletionDelayHours pulumi.Float64PtrInput
	// User-provided description for this private cloud.
	Description pulumi.StringPtrInput
	// The location where the PrivateCloud should reside.
	Location pulumi.StringInput
	// The management cluster for this private cloud. This used for creating and managing the default cluster.
	ManagementCluster VmwareenginePrivateCloudManagementClusterInput
	// The ID of the PrivateCloud.
	Name pulumi.StringPtrInput
	// Network configuration in the consumer project with which the peering has to be done.
	NetworkConfig VmwareenginePrivateCloudNetworkConfigInput
	Project       pulumi.StringPtrInput
	// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
	// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
	// deletion_delay_hours.
	SendDeletionDelayHoursIfZero pulumi.BoolPtrInput
	Timeouts                     VmwareenginePrivateCloudTimeoutsPtrInput
	// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
	Type                       pulumi.StringPtrInput
	VmwareenginePrivateCloudId pulumi.StringPtrInput
}

func (VmwareenginePrivateCloudArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareenginePrivateCloudArgs)(nil)).Elem()
}

type VmwareenginePrivateCloudInput interface {
	pulumi.Input

	ToVmwareenginePrivateCloudOutput() VmwareenginePrivateCloudOutput
	ToVmwareenginePrivateCloudOutputWithContext(ctx context.Context) VmwareenginePrivateCloudOutput
}

func (*VmwareenginePrivateCloud) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareenginePrivateCloud)(nil)).Elem()
}

func (i *VmwareenginePrivateCloud) ToVmwareenginePrivateCloudOutput() VmwareenginePrivateCloudOutput {
	return i.ToVmwareenginePrivateCloudOutputWithContext(context.Background())
}

func (i *VmwareenginePrivateCloud) ToVmwareenginePrivateCloudOutputWithContext(ctx context.Context) VmwareenginePrivateCloudOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmwareenginePrivateCloudOutput)
}

type VmwareenginePrivateCloudOutput struct{ *pulumi.OutputState }

func (VmwareenginePrivateCloudOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmwareenginePrivateCloud)(nil)).Elem()
}

func (o VmwareenginePrivateCloudOutput) ToVmwareenginePrivateCloudOutput() VmwareenginePrivateCloudOutput {
	return o
}

func (o VmwareenginePrivateCloudOutput) ToVmwareenginePrivateCloudOutputWithContext(ctx context.Context) VmwareenginePrivateCloudOutput {
	return o
}

// The number of hours to delay this request. You can set this value to an hour between 0 to 8, where setting it to 0
// starts the deletion request immediately. If no value is set, a default value is set at the API Level.
func (o VmwareenginePrivateCloudOutput) DeletionDelayHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.Float64PtrOutput { return v.DeletionDelayHours }).(pulumi.Float64PtrOutput)
}

// User-provided description for this private cloud.
func (o VmwareenginePrivateCloudOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Details about a HCX Cloud Manager appliance.
func (o VmwareenginePrivateCloudOutput) Hcxes() VmwareenginePrivateCloudHcxArrayOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudHcxArrayOutput { return v.Hcxes }).(VmwareenginePrivateCloudHcxArrayOutput)
}

// The location where the PrivateCloud should reside.
func (o VmwareenginePrivateCloudOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The management cluster for this private cloud. This used for creating and managing the default cluster.
func (o VmwareenginePrivateCloudOutput) ManagementCluster() VmwareenginePrivateCloudManagementClusterOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudManagementClusterOutput {
		return v.ManagementCluster
	}).(VmwareenginePrivateCloudManagementClusterOutput)
}

// The ID of the PrivateCloud.
func (o VmwareenginePrivateCloudOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network configuration in the consumer project with which the peering has to be done.
func (o VmwareenginePrivateCloudOutput) NetworkConfig() VmwareenginePrivateCloudNetworkConfigOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudNetworkConfigOutput { return v.NetworkConfig }).(VmwareenginePrivateCloudNetworkConfigOutput)
}

// Details about a NSX Manager appliance.
func (o VmwareenginePrivateCloudOutput) Nsxes() VmwareenginePrivateCloudNsxArrayOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudNsxArrayOutput { return v.Nsxes }).(VmwareenginePrivateCloudNsxArrayOutput)
}

func (o VmwareenginePrivateCloudOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// While set true, deletion_delay_hours value will be sent in the request even for zero value of the field. This field is
// only useful for setting 0 value to the deletion_delay_hours field. It can be used both alone and together with
// deletion_delay_hours.
func (o VmwareenginePrivateCloudOutput) SendDeletionDelayHoursIfZero() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.BoolPtrOutput { return v.SendDeletionDelayHoursIfZero }).(pulumi.BoolPtrOutput)
}

// State of the resource. New values may be added to this enum when appropriate.
func (o VmwareenginePrivateCloudOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o VmwareenginePrivateCloudOutput) Timeouts() VmwareenginePrivateCloudTimeoutsPtrOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudTimeoutsPtrOutput { return v.Timeouts }).(VmwareenginePrivateCloudTimeoutsPtrOutput)
}

// Initial type of the private cloud. Possible values: ["STANDARD", "TIME_LIMITED", "STRETCHED"]
func (o VmwareenginePrivateCloudOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// System-generated unique identifier for the resource.
func (o VmwareenginePrivateCloudOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Details about a vCenter Server management appliance.
func (o VmwareenginePrivateCloudOutput) Vcenters() VmwareenginePrivateCloudVcenterArrayOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) VmwareenginePrivateCloudVcenterArrayOutput { return v.Vcenters }).(VmwareenginePrivateCloudVcenterArrayOutput)
}

func (o VmwareenginePrivateCloudOutput) VmwareenginePrivateCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v *VmwareenginePrivateCloud) pulumi.StringOutput { return v.VmwareenginePrivateCloudId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmwareenginePrivateCloudInput)(nil)).Elem(), &VmwareenginePrivateCloud{})
	pulumi.RegisterOutputType(VmwareenginePrivateCloudOutput{})
}
