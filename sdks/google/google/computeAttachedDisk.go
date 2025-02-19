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

type ComputeAttachedDisk struct {
	pulumi.CustomResourceState

	ComputeAttachedDiskId pulumi.StringOutput `pulumi:"computeAttachedDiskId"`
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
	// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
	// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
	// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	DeviceName pulumi.StringOutput `pulumi:"deviceName"`
	// name or self_link of the disk that will be attached.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
	// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
	// on the resource or provider.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
	// please don't specify this field without advice from Google.)
	Interface pulumi.StringPtrOutput `pulumi:"interface"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
	// disk in READ_WRITE mode.
	Mode pulumi.StringPtrOutput `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
	// defined in the link will take precedence.
	Project  pulumi.StringOutput                  `pulumi:"project"`
	Timeouts ComputeAttachedDiskTimeoutsPtrOutput `pulumi:"timeouts"`
	// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
	// defined in the link will take precedence.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeAttachedDisk registers a new resource with the given unique name, arguments, and options.
func NewComputeAttachedDisk(ctx *pulumi.Context,
	name string, args *ComputeAttachedDiskArgs, opts ...pulumi.ResourceOption) (*ComputeAttachedDisk, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeAttachedDisk
	err = ctx.RegisterPackageResource("google:index/computeAttachedDisk:ComputeAttachedDisk", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeAttachedDisk gets an existing ComputeAttachedDisk resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeAttachedDisk(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeAttachedDiskState, opts ...pulumi.ResourceOption) (*ComputeAttachedDisk, error) {
	var resource ComputeAttachedDisk
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeAttachedDisk:ComputeAttachedDisk", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeAttachedDisk resources.
type computeAttachedDiskState struct {
	ComputeAttachedDiskId *string `pulumi:"computeAttachedDiskId"`
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
	// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
	// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
	// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	DeviceName *string `pulumi:"deviceName"`
	// name or self_link of the disk that will be attached.
	Disk *string `pulumi:"disk"`
	// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
	// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
	// on the resource or provider.
	Instance *string `pulumi:"instance"`
	// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
	// please don't specify this field without advice from Google.)
	Interface *string `pulumi:"interface"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
	// disk in READ_WRITE mode.
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
	// defined in the link will take precedence.
	Project  *string                      `pulumi:"project"`
	Timeouts *ComputeAttachedDiskTimeouts `pulumi:"timeouts"`
	// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
	// defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

type ComputeAttachedDiskState struct {
	ComputeAttachedDiskId pulumi.StringPtrInput
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
	// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
	// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
	// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	DeviceName pulumi.StringPtrInput
	// name or self_link of the disk that will be attached.
	Disk pulumi.StringPtrInput
	// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
	// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
	// on the resource or provider.
	Instance pulumi.StringPtrInput
	// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
	// please don't specify this field without advice from Google.)
	Interface pulumi.StringPtrInput
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
	// disk in READ_WRITE mode.
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
	// defined in the link will take precedence.
	Project  pulumi.StringPtrInput
	Timeouts ComputeAttachedDiskTimeoutsPtrInput
	// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
	// defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (ComputeAttachedDiskState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeAttachedDiskState)(nil)).Elem()
}

type computeAttachedDiskArgs struct {
	ComputeAttachedDiskId *string `pulumi:"computeAttachedDiskId"`
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
	// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
	// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
	// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	DeviceName *string `pulumi:"deviceName"`
	// name or self_link of the disk that will be attached.
	Disk string `pulumi:"disk"`
	// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
	// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
	// on the resource or provider.
	Instance string `pulumi:"instance"`
	// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
	// please don't specify this field without advice from Google.)
	Interface *string `pulumi:"interface"`
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
	// disk in READ_WRITE mode.
	Mode *string `pulumi:"mode"`
	// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
	// defined in the link will take precedence.
	Project  *string                      `pulumi:"project"`
	Timeouts *ComputeAttachedDiskTimeouts `pulumi:"timeouts"`
	// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
	// defined in the link will take precedence.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeAttachedDisk resource.
type ComputeAttachedDiskArgs struct {
	ComputeAttachedDiskId pulumi.StringPtrInput
	// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
	// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
	// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
	// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
	DeviceName pulumi.StringPtrInput
	// name or self_link of the disk that will be attached.
	Disk pulumi.StringInput
	// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
	// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
	// on the resource or provider.
	Instance pulumi.StringInput
	// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
	// please don't specify this field without advice from Google.)
	Interface pulumi.StringPtrInput
	// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
	// disk in READ_WRITE mode.
	Mode pulumi.StringPtrInput
	// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
	// defined in the link will take precedence.
	Project  pulumi.StringPtrInput
	Timeouts ComputeAttachedDiskTimeoutsPtrInput
	// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
	// defined in the link will take precedence.
	Zone pulumi.StringPtrInput
}

func (ComputeAttachedDiskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeAttachedDiskArgs)(nil)).Elem()
}

type ComputeAttachedDiskInput interface {
	pulumi.Input

	ToComputeAttachedDiskOutput() ComputeAttachedDiskOutput
	ToComputeAttachedDiskOutputWithContext(ctx context.Context) ComputeAttachedDiskOutput
}

func (*ComputeAttachedDisk) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeAttachedDisk)(nil)).Elem()
}

func (i *ComputeAttachedDisk) ToComputeAttachedDiskOutput() ComputeAttachedDiskOutput {
	return i.ToComputeAttachedDiskOutputWithContext(context.Background())
}

func (i *ComputeAttachedDisk) ToComputeAttachedDiskOutputWithContext(ctx context.Context) ComputeAttachedDiskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeAttachedDiskOutput)
}

type ComputeAttachedDiskOutput struct{ *pulumi.OutputState }

func (ComputeAttachedDiskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeAttachedDisk)(nil)).Elem()
}

func (o ComputeAttachedDiskOutput) ToComputeAttachedDiskOutput() ComputeAttachedDiskOutput {
	return o
}

func (o ComputeAttachedDiskOutput) ToComputeAttachedDiskOutputWithContext(ctx context.Context) ComputeAttachedDiskOutput {
	return o
}

func (o ComputeAttachedDiskOutput) ComputeAttachedDiskId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.ComputeAttachedDiskId }).(pulumi.StringOutput)
}

// Specifies a unique device name of your choice that is reflected into the /dev/disk/by-id/google-* tree of a Linux
// operating system running within the instance. This name can be used to reference the device for mounting, resizing, and
// so on, from within the instance. If not specified, the server chooses a default device name to apply to this disk, in
// the form persistent-disks-x, where x is a number assigned by Google Compute Engine.
func (o ComputeAttachedDiskOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.DeviceName }).(pulumi.StringOutput)
}

// name or self_link of the disk that will be attached.
func (o ComputeAttachedDiskOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// name or self_link of the compute instance that the disk will be attached to. If the self_link is provided then zone and
// project are extracted from the self link. If only the name is used then zone and project must be defined as properties
// on the resource or provider.
func (o ComputeAttachedDiskOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The disk interface used for attaching this disk. One of SCSI or NVME. (This field is only used for specific cases,
// please don't specify this field without advice from Google.)
func (o ComputeAttachedDiskOutput) Interface() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringPtrOutput { return v.Interface }).(pulumi.StringPtrOutput)
}

// The mode in which to attach this disk, either READ_WRITE or READ_ONLY. If not specified, the default is to attach the
// disk in READ_WRITE mode.
func (o ComputeAttachedDiskOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringPtrOutput { return v.Mode }).(pulumi.StringPtrOutput)
}

// The project that the referenced compute instance is a part of. If instance is referenced by its self_link the project
// defined in the link will take precedence.
func (o ComputeAttachedDiskOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeAttachedDiskOutput) Timeouts() ComputeAttachedDiskTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) ComputeAttachedDiskTimeoutsPtrOutput { return v.Timeouts }).(ComputeAttachedDiskTimeoutsPtrOutput)
}

// The zone that the referenced compute instance is located within. If instance is referenced by its self_link the zone
// defined in the link will take precedence.
func (o ComputeAttachedDiskOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeAttachedDisk) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeAttachedDiskInput)(nil)).Elem(), &ComputeAttachedDisk{})
	pulumi.RegisterOutputType(ComputeAttachedDiskOutput{})
}
