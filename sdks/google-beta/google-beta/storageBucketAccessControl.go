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

type StorageBucketAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
	// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
	// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
	// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	Role                         pulumi.StringPtrOutput                      `pulumi:"role"`
	StorageBucketAccessControlId pulumi.StringOutput                         `pulumi:"storageBucketAccessControlId"`
	Timeouts                     StorageBucketAccessControlTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewStorageBucketAccessControl registers a new resource with the given unique name, arguments, and options.
func NewStorageBucketAccessControl(ctx *pulumi.Context,
	name string, args *StorageBucketAccessControlArgs, opts ...pulumi.ResourceOption) (*StorageBucketAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageBucketAccessControl
	err = ctx.RegisterPackageResource("google-beta:index/storageBucketAccessControl:StorageBucketAccessControl", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucketAccessControl gets an existing StorageBucketAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucketAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketAccessControlState, opts ...pulumi.ResourceOption) (*StorageBucketAccessControl, error) {
	var resource StorageBucketAccessControl
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageBucketAccessControl:StorageBucketAccessControl", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucketAccessControl resources.
type storageBucketAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
	// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
	// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
	// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity *string `pulumi:"entity"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	Role                         *string                             `pulumi:"role"`
	StorageBucketAccessControlId *string                             `pulumi:"storageBucketAccessControlId"`
	Timeouts                     *StorageBucketAccessControlTimeouts `pulumi:"timeouts"`
}

type StorageBucketAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
	// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
	// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
	// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringPtrInput
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	Role                         pulumi.StringPtrInput
	StorageBucketAccessControlId pulumi.StringPtrInput
	Timeouts                     StorageBucketAccessControlTimeoutsPtrInput
}

func (StorageBucketAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketAccessControlState)(nil)).Elem()
}

type storageBucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
	// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
	// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
	// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity string `pulumi:"entity"`
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	Role                         *string                             `pulumi:"role"`
	StorageBucketAccessControlId *string                             `pulumi:"storageBucketAccessControlId"`
	Timeouts                     *StorageBucketAccessControlTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a StorageBucketAccessControl resource.
type StorageBucketAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
	// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
	// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
	// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
	Entity pulumi.StringInput
	// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
	Role                         pulumi.StringPtrInput
	StorageBucketAccessControlId pulumi.StringPtrInput
	Timeouts                     StorageBucketAccessControlTimeoutsPtrInput
}

func (StorageBucketAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketAccessControlArgs)(nil)).Elem()
}

type StorageBucketAccessControlInput interface {
	pulumi.Input

	ToStorageBucketAccessControlOutput() StorageBucketAccessControlOutput
	ToStorageBucketAccessControlOutputWithContext(ctx context.Context) StorageBucketAccessControlOutput
}

func (*StorageBucketAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketAccessControl)(nil)).Elem()
}

func (i *StorageBucketAccessControl) ToStorageBucketAccessControlOutput() StorageBucketAccessControlOutput {
	return i.ToStorageBucketAccessControlOutputWithContext(context.Background())
}

func (i *StorageBucketAccessControl) ToStorageBucketAccessControlOutputWithContext(ctx context.Context) StorageBucketAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketAccessControlOutput)
}

type StorageBucketAccessControlOutput struct{ *pulumi.OutputState }

func (StorageBucketAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketAccessControl)(nil)).Elem()
}

func (o StorageBucketAccessControlOutput) ToStorageBucketAccessControlOutput() StorageBucketAccessControlOutput {
	return o
}

func (o StorageBucketAccessControlOutput) ToStorageBucketAccessControlOutputWithContext(ctx context.Context) StorageBucketAccessControlOutput {
	return o
}

// The name of the bucket.
func (o StorageBucketAccessControlOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The domain associated with the entity.
func (o StorageBucketAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The email address associated with the entity.
func (o StorageBucketAccessControlOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The entity holding the permission, in one of the following forms: user-userId user-email group-groupId group-email
// domain-domain project-team-projectId allUsers allAuthenticatedUsers Examples: The user liz@example.com would be
// user-liz@example.com. The group example@googlegroups.com would be group-example@googlegroups.com. To refer to all
// members of the Google Apps for Business domain example.com, the entity would be domain-example.com.
func (o StorageBucketAccessControlOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// The access permission for the entity. Possible values: ["OWNER", "READER", "WRITER"]
func (o StorageBucketAccessControlOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

func (o StorageBucketAccessControlOutput) StorageBucketAccessControlId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) pulumi.StringOutput { return v.StorageBucketAccessControlId }).(pulumi.StringOutput)
}

func (o StorageBucketAccessControlOutput) Timeouts() StorageBucketAccessControlTimeoutsPtrOutput {
	return o.ApplyT(func(v *StorageBucketAccessControl) StorageBucketAccessControlTimeoutsPtrOutput { return v.Timeouts }).(StorageBucketAccessControlTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketAccessControlInput)(nil)).Elem(), &StorageBucketAccessControl{})
	pulumi.RegisterOutputType(StorageBucketAccessControlOutput{})
}
