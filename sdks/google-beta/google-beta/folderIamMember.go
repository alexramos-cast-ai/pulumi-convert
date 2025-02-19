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

type FolderIamMember struct {
	pulumi.CustomResourceState

	Condition         FolderIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag              pulumi.StringOutput               `pulumi:"etag"`
	Folder            pulumi.StringOutput               `pulumi:"folder"`
	FolderIamMemberId pulumi.StringOutput               `pulumi:"folderIamMemberId"`
	Member            pulumi.StringOutput               `pulumi:"member"`
	Role              pulumi.StringOutput               `pulumi:"role"`
}

// NewFolderIamMember registers a new resource with the given unique name, arguments, and options.
func NewFolderIamMember(ctx *pulumi.Context,
	name string, args *FolderIamMemberArgs, opts ...pulumi.ResourceOption) (*FolderIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FolderIamMember
	err = ctx.RegisterPackageResource("google-beta:index/folderIamMember:FolderIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderIamMember gets an existing FolderIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderIamMemberState, opts ...pulumi.ResourceOption) (*FolderIamMember, error) {
	var resource FolderIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/folderIamMember:FolderIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderIamMember resources.
type folderIamMemberState struct {
	Condition         *FolderIamMemberCondition `pulumi:"condition"`
	Etag              *string                   `pulumi:"etag"`
	Folder            *string                   `pulumi:"folder"`
	FolderIamMemberId *string                   `pulumi:"folderIamMemberId"`
	Member            *string                   `pulumi:"member"`
	Role              *string                   `pulumi:"role"`
}

type FolderIamMemberState struct {
	Condition         FolderIamMemberConditionPtrInput
	Etag              pulumi.StringPtrInput
	Folder            pulumi.StringPtrInput
	FolderIamMemberId pulumi.StringPtrInput
	Member            pulumi.StringPtrInput
	Role              pulumi.StringPtrInput
}

func (FolderIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamMemberState)(nil)).Elem()
}

type folderIamMemberArgs struct {
	Condition         *FolderIamMemberCondition `pulumi:"condition"`
	Folder            string                    `pulumi:"folder"`
	FolderIamMemberId *string                   `pulumi:"folderIamMemberId"`
	Member            string                    `pulumi:"member"`
	Role              string                    `pulumi:"role"`
}

// The set of arguments for constructing a FolderIamMember resource.
type FolderIamMemberArgs struct {
	Condition         FolderIamMemberConditionPtrInput
	Folder            pulumi.StringInput
	FolderIamMemberId pulumi.StringPtrInput
	Member            pulumi.StringInput
	Role              pulumi.StringInput
}

func (FolderIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamMemberArgs)(nil)).Elem()
}

type FolderIamMemberInput interface {
	pulumi.Input

	ToFolderIamMemberOutput() FolderIamMemberOutput
	ToFolderIamMemberOutputWithContext(ctx context.Context) FolderIamMemberOutput
}

func (*FolderIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamMember)(nil)).Elem()
}

func (i *FolderIamMember) ToFolderIamMemberOutput() FolderIamMemberOutput {
	return i.ToFolderIamMemberOutputWithContext(context.Background())
}

func (i *FolderIamMember) ToFolderIamMemberOutputWithContext(ctx context.Context) FolderIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderIamMemberOutput)
}

type FolderIamMemberOutput struct{ *pulumi.OutputState }

func (FolderIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamMember)(nil)).Elem()
}

func (o FolderIamMemberOutput) ToFolderIamMemberOutput() FolderIamMemberOutput {
	return o
}

func (o FolderIamMemberOutput) ToFolderIamMemberOutputWithContext(ctx context.Context) FolderIamMemberOutput {
	return o
}

func (o FolderIamMemberOutput) Condition() FolderIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *FolderIamMember) FolderIamMemberConditionPtrOutput { return v.Condition }).(FolderIamMemberConditionPtrOutput)
}

func (o FolderIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FolderIamMemberOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamMember) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

func (o FolderIamMemberOutput) FolderIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamMember) pulumi.StringOutput { return v.FolderIamMemberId }).(pulumi.StringOutput)
}

func (o FolderIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o FolderIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderIamMemberInput)(nil)).Elem(), &FolderIamMember{})
	pulumi.RegisterOutputType(FolderIamMemberOutput{})
}
