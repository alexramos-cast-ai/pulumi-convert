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

type ComputeDiskResourcePolicyAttachment struct {
	pulumi.CustomResourceState

	ComputeDiskResourcePolicyAttachmentId pulumi.StringOutput `pulumi:"computeDiskResourcePolicyAttachmentId"`
	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringOutput `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name     pulumi.StringOutput                                  `pulumi:"name"`
	Project  pulumi.StringOutput                                  `pulumi:"project"`
	Timeouts ComputeDiskResourcePolicyAttachmentTimeoutsPtrOutput `pulumi:"timeouts"`
	// A reference to the zone where the disk resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputeDiskResourcePolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewComputeDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, args *ComputeDiskResourcePolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*ComputeDiskResourcePolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Disk == nil {
		return nil, errors.New("invalid value for required argument 'Disk'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeDiskResourcePolicyAttachment
	err = ctx.RegisterPackageResource("google:index/computeDiskResourcePolicyAttachment:ComputeDiskResourcePolicyAttachment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeDiskResourcePolicyAttachment gets an existing ComputeDiskResourcePolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeDiskResourcePolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeDiskResourcePolicyAttachmentState, opts ...pulumi.ResourceOption) (*ComputeDiskResourcePolicyAttachment, error) {
	var resource ComputeDiskResourcePolicyAttachment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeDiskResourcePolicyAttachment:ComputeDiskResourcePolicyAttachment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeDiskResourcePolicyAttachment resources.
type computeDiskResourcePolicyAttachmentState struct {
	ComputeDiskResourcePolicyAttachmentId *string `pulumi:"computeDiskResourcePolicyAttachmentId"`
	// The name of the disk in which the resource policies are attached to.
	Disk *string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name     *string                                      `pulumi:"name"`
	Project  *string                                      `pulumi:"project"`
	Timeouts *ComputeDiskResourcePolicyAttachmentTimeouts `pulumi:"timeouts"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

type ComputeDiskResourcePolicyAttachmentState struct {
	ComputeDiskResourcePolicyAttachmentId pulumi.StringPtrInput
	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringPtrInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeDiskResourcePolicyAttachmentTimeoutsPtrInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (ComputeDiskResourcePolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeDiskResourcePolicyAttachmentState)(nil)).Elem()
}

type computeDiskResourcePolicyAttachmentArgs struct {
	ComputeDiskResourcePolicyAttachmentId *string `pulumi:"computeDiskResourcePolicyAttachmentId"`
	// The name of the disk in which the resource policies are attached to.
	Disk string `pulumi:"disk"`
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name     *string                                      `pulumi:"name"`
	Project  *string                                      `pulumi:"project"`
	Timeouts *ComputeDiskResourcePolicyAttachmentTimeouts `pulumi:"timeouts"`
	// A reference to the zone where the disk resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeDiskResourcePolicyAttachment resource.
type ComputeDiskResourcePolicyAttachmentArgs struct {
	ComputeDiskResourcePolicyAttachmentId pulumi.StringPtrInput
	// The name of the disk in which the resource policies are attached to.
	Disk pulumi.StringInput
	// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeDiskResourcePolicyAttachmentTimeoutsPtrInput
	// A reference to the zone where the disk resides.
	Zone pulumi.StringPtrInput
}

func (ComputeDiskResourcePolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeDiskResourcePolicyAttachmentArgs)(nil)).Elem()
}

type ComputeDiskResourcePolicyAttachmentInput interface {
	pulumi.Input

	ToComputeDiskResourcePolicyAttachmentOutput() ComputeDiskResourcePolicyAttachmentOutput
	ToComputeDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeDiskResourcePolicyAttachmentOutput
}

func (*ComputeDiskResourcePolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeDiskResourcePolicyAttachment)(nil)).Elem()
}

func (i *ComputeDiskResourcePolicyAttachment) ToComputeDiskResourcePolicyAttachmentOutput() ComputeDiskResourcePolicyAttachmentOutput {
	return i.ToComputeDiskResourcePolicyAttachmentOutputWithContext(context.Background())
}

func (i *ComputeDiskResourcePolicyAttachment) ToComputeDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeDiskResourcePolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeDiskResourcePolicyAttachmentOutput)
}

type ComputeDiskResourcePolicyAttachmentOutput struct{ *pulumi.OutputState }

func (ComputeDiskResourcePolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeDiskResourcePolicyAttachment)(nil)).Elem()
}

func (o ComputeDiskResourcePolicyAttachmentOutput) ToComputeDiskResourcePolicyAttachmentOutput() ComputeDiskResourcePolicyAttachmentOutput {
	return o
}

func (o ComputeDiskResourcePolicyAttachmentOutput) ToComputeDiskResourcePolicyAttachmentOutputWithContext(ctx context.Context) ComputeDiskResourcePolicyAttachmentOutput {
	return o
}

func (o ComputeDiskResourcePolicyAttachmentOutput) ComputeDiskResourcePolicyAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) pulumi.StringOutput {
		return v.ComputeDiskResourcePolicyAttachmentId
	}).(pulumi.StringOutput)
}

// The name of the disk in which the resource policies are attached to.
func (o ComputeDiskResourcePolicyAttachmentOutput) Disk() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Disk }).(pulumi.StringOutput)
}

// The resource policy to be attached to the disk for scheduling snapshot creation. Do not specify the self link.
func (o ComputeDiskResourcePolicyAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeDiskResourcePolicyAttachmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeDiskResourcePolicyAttachmentOutput) Timeouts() ComputeDiskResourcePolicyAttachmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) ComputeDiskResourcePolicyAttachmentTimeoutsPtrOutput {
		return v.Timeouts
	}).(ComputeDiskResourcePolicyAttachmentTimeoutsPtrOutput)
}

// A reference to the zone where the disk resides.
func (o ComputeDiskResourcePolicyAttachmentOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeDiskResourcePolicyAttachment) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeDiskResourcePolicyAttachmentInput)(nil)).Elem(), &ComputeDiskResourcePolicyAttachment{})
	pulumi.RegisterOutputType(ComputeDiskResourcePolicyAttachmentOutput{})
}
