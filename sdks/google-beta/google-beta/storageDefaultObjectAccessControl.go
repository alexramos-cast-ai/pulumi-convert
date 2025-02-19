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

type StorageDefaultObjectAccessControl struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain pulumi.StringOutput `pulumi:"domain"`
	// The email address associated with the entity.
	Email pulumi.StringOutput `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringOutput `pulumi:"entity"`
	// The ID for the entity
	EntityId pulumi.StringOutput `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation pulumi.Float64Output `pulumi:"generation"`
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrOutput `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeams StorageDefaultObjectAccessControlProjectTeamArrayOutput `pulumi:"projectTeams"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	Role                                pulumi.StringOutput                                `pulumi:"role"`
	StorageDefaultObjectAccessControlId pulumi.StringOutput                                `pulumi:"storageDefaultObjectAccessControlId"`
	Timeouts                            StorageDefaultObjectAccessControlTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewStorageDefaultObjectAccessControl registers a new resource with the given unique name, arguments, and options.
func NewStorageDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, args *StorageDefaultObjectAccessControlArgs, opts ...pulumi.ResourceOption) (*StorageDefaultObjectAccessControl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Entity == nil {
		return nil, errors.New("invalid value for required argument 'Entity'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageDefaultObjectAccessControl
	err = ctx.RegisterPackageResource("google-beta:index/storageDefaultObjectAccessControl:StorageDefaultObjectAccessControl", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageDefaultObjectAccessControl gets an existing StorageDefaultObjectAccessControl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageDefaultObjectAccessControl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageDefaultObjectAccessControlState, opts ...pulumi.ResourceOption) (*StorageDefaultObjectAccessControl, error) {
	var resource StorageDefaultObjectAccessControl
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/storageDefaultObjectAccessControl:StorageDefaultObjectAccessControl", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageDefaultObjectAccessControl resources.
type storageDefaultObjectAccessControlState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// The domain associated with the entity.
	Domain *string `pulumi:"domain"`
	// The email address associated with the entity.
	Email *string `pulumi:"email"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity *string `pulumi:"entity"`
	// The ID for the entity
	EntityId *string `pulumi:"entityId"`
	// The content generation of the object, if applied to an object.
	Generation *float64 `pulumi:"generation"`
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The project team associated with the entity
	ProjectTeams []StorageDefaultObjectAccessControlProjectTeam `pulumi:"projectTeams"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	Role                                *string                                    `pulumi:"role"`
	StorageDefaultObjectAccessControlId *string                                    `pulumi:"storageDefaultObjectAccessControlId"`
	Timeouts                            *StorageDefaultObjectAccessControlTimeouts `pulumi:"timeouts"`
}

type StorageDefaultObjectAccessControlState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// The domain associated with the entity.
	Domain pulumi.StringPtrInput
	// The email address associated with the entity.
	Email pulumi.StringPtrInput
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringPtrInput
	// The ID for the entity
	EntityId pulumi.StringPtrInput
	// The content generation of the object, if applied to an object.
	Generation pulumi.Float64PtrInput
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The project team associated with the entity
	ProjectTeams StorageDefaultObjectAccessControlProjectTeamArrayInput
	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	Role                                pulumi.StringPtrInput
	StorageDefaultObjectAccessControlId pulumi.StringPtrInput
	Timeouts                            StorageDefaultObjectAccessControlTimeoutsPtrInput
}

func (StorageDefaultObjectAccessControlState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDefaultObjectAccessControlState)(nil)).Elem()
}

type storageDefaultObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity string `pulumi:"entity"`
	// The name of the object, if applied to an object.
	Object *string `pulumi:"object"`
	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	Role                                string                                     `pulumi:"role"`
	StorageDefaultObjectAccessControlId *string                                    `pulumi:"storageDefaultObjectAccessControlId"`
	Timeouts                            *StorageDefaultObjectAccessControlTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a StorageDefaultObjectAccessControl resource.
type StorageDefaultObjectAccessControlArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
	// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
	// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
	Entity pulumi.StringInput
	// The name of the object, if applied to an object.
	Object pulumi.StringPtrInput
	// The access permission for the entity. Possible values: ["OWNER", "READER"]
	Role                                pulumi.StringInput
	StorageDefaultObjectAccessControlId pulumi.StringPtrInput
	Timeouts                            StorageDefaultObjectAccessControlTimeoutsPtrInput
}

func (StorageDefaultObjectAccessControlArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDefaultObjectAccessControlArgs)(nil)).Elem()
}

type StorageDefaultObjectAccessControlInput interface {
	pulumi.Input

	ToStorageDefaultObjectAccessControlOutput() StorageDefaultObjectAccessControlOutput
	ToStorageDefaultObjectAccessControlOutputWithContext(ctx context.Context) StorageDefaultObjectAccessControlOutput
}

func (*StorageDefaultObjectAccessControl) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDefaultObjectAccessControl)(nil)).Elem()
}

func (i *StorageDefaultObjectAccessControl) ToStorageDefaultObjectAccessControlOutput() StorageDefaultObjectAccessControlOutput {
	return i.ToStorageDefaultObjectAccessControlOutputWithContext(context.Background())
}

func (i *StorageDefaultObjectAccessControl) ToStorageDefaultObjectAccessControlOutputWithContext(ctx context.Context) StorageDefaultObjectAccessControlOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageDefaultObjectAccessControlOutput)
}

type StorageDefaultObjectAccessControlOutput struct{ *pulumi.OutputState }

func (StorageDefaultObjectAccessControlOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageDefaultObjectAccessControl)(nil)).Elem()
}

func (o StorageDefaultObjectAccessControlOutput) ToStorageDefaultObjectAccessControlOutput() StorageDefaultObjectAccessControlOutput {
	return o
}

func (o StorageDefaultObjectAccessControlOutput) ToStorageDefaultObjectAccessControlOutputWithContext(ctx context.Context) StorageDefaultObjectAccessControlOutput {
	return o
}

// The name of the bucket.
func (o StorageDefaultObjectAccessControlOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// The domain associated with the entity.
func (o StorageDefaultObjectAccessControlOutput) Domain() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.Domain }).(pulumi.StringOutput)
}

// The email address associated with the entity.
func (o StorageDefaultObjectAccessControlOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

// The entity holding the permission, in one of the following forms: * user-{{userId}} * user-{{email}} (such as
// "user-liz@example.com") * group-{{groupId}} * group-{{email}} (such as "group-example@googlegroups.com") *
// domain-{{domain}} (such as "domain-example.com") * project-team-{{projectId}} * allUsers * allAuthenticatedUsers
func (o StorageDefaultObjectAccessControlOutput) Entity() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.Entity }).(pulumi.StringOutput)
}

// The ID for the entity
func (o StorageDefaultObjectAccessControlOutput) EntityId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.EntityId }).(pulumi.StringOutput)
}

// The content generation of the object, if applied to an object.
func (o StorageDefaultObjectAccessControlOutput) Generation() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.Float64Output { return v.Generation }).(pulumi.Float64Output)
}

// The name of the object, if applied to an object.
func (o StorageDefaultObjectAccessControlOutput) Object() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringPtrOutput { return v.Object }).(pulumi.StringPtrOutput)
}

// The project team associated with the entity
func (o StorageDefaultObjectAccessControlOutput) ProjectTeams() StorageDefaultObjectAccessControlProjectTeamArrayOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) StorageDefaultObjectAccessControlProjectTeamArrayOutput {
		return v.ProjectTeams
	}).(StorageDefaultObjectAccessControlProjectTeamArrayOutput)
}

// The access permission for the entity. Possible values: ["OWNER", "READER"]
func (o StorageDefaultObjectAccessControlOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o StorageDefaultObjectAccessControlOutput) StorageDefaultObjectAccessControlId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) pulumi.StringOutput {
		return v.StorageDefaultObjectAccessControlId
	}).(pulumi.StringOutput)
}

func (o StorageDefaultObjectAccessControlOutput) Timeouts() StorageDefaultObjectAccessControlTimeoutsPtrOutput {
	return o.ApplyT(func(v *StorageDefaultObjectAccessControl) StorageDefaultObjectAccessControlTimeoutsPtrOutput {
		return v.Timeouts
	}).(StorageDefaultObjectAccessControlTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageDefaultObjectAccessControlInput)(nil)).Elem(), &StorageDefaultObjectAccessControl{})
	pulumi.RegisterOutputType(StorageDefaultObjectAccessControlOutput{})
}
