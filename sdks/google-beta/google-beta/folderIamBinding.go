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

type FolderIamBinding struct {
	pulumi.CustomResourceState

	Condition          FolderIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag               pulumi.StringOutput                `pulumi:"etag"`
	Folder             pulumi.StringOutput                `pulumi:"folder"`
	FolderIamBindingId pulumi.StringOutput                `pulumi:"folderIamBindingId"`
	Members            pulumi.StringArrayOutput           `pulumi:"members"`
	Role               pulumi.StringOutput                `pulumi:"role"`
}

// NewFolderIamBinding registers a new resource with the given unique name, arguments, and options.
func NewFolderIamBinding(ctx *pulumi.Context,
	name string, args *FolderIamBindingArgs, opts ...pulumi.ResourceOption) (*FolderIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FolderIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/folderIamBinding:FolderIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderIamBinding gets an existing FolderIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderIamBindingState, opts ...pulumi.ResourceOption) (*FolderIamBinding, error) {
	var resource FolderIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/folderIamBinding:FolderIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderIamBinding resources.
type folderIamBindingState struct {
	Condition          *FolderIamBindingCondition `pulumi:"condition"`
	Etag               *string                    `pulumi:"etag"`
	Folder             *string                    `pulumi:"folder"`
	FolderIamBindingId *string                    `pulumi:"folderIamBindingId"`
	Members            []string                   `pulumi:"members"`
	Role               *string                    `pulumi:"role"`
}

type FolderIamBindingState struct {
	Condition          FolderIamBindingConditionPtrInput
	Etag               pulumi.StringPtrInput
	Folder             pulumi.StringPtrInput
	FolderIamBindingId pulumi.StringPtrInput
	Members            pulumi.StringArrayInput
	Role               pulumi.StringPtrInput
}

func (FolderIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamBindingState)(nil)).Elem()
}

type folderIamBindingArgs struct {
	Condition          *FolderIamBindingCondition `pulumi:"condition"`
	Folder             string                     `pulumi:"folder"`
	FolderIamBindingId *string                    `pulumi:"folderIamBindingId"`
	Members            []string                   `pulumi:"members"`
	Role               string                     `pulumi:"role"`
}

// The set of arguments for constructing a FolderIamBinding resource.
type FolderIamBindingArgs struct {
	Condition          FolderIamBindingConditionPtrInput
	Folder             pulumi.StringInput
	FolderIamBindingId pulumi.StringPtrInput
	Members            pulumi.StringArrayInput
	Role               pulumi.StringInput
}

func (FolderIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderIamBindingArgs)(nil)).Elem()
}

type FolderIamBindingInput interface {
	pulumi.Input

	ToFolderIamBindingOutput() FolderIamBindingOutput
	ToFolderIamBindingOutputWithContext(ctx context.Context) FolderIamBindingOutput
}

func (*FolderIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamBinding)(nil)).Elem()
}

func (i *FolderIamBinding) ToFolderIamBindingOutput() FolderIamBindingOutput {
	return i.ToFolderIamBindingOutputWithContext(context.Background())
}

func (i *FolderIamBinding) ToFolderIamBindingOutputWithContext(ctx context.Context) FolderIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderIamBindingOutput)
}

type FolderIamBindingOutput struct{ *pulumi.OutputState }

func (FolderIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderIamBinding)(nil)).Elem()
}

func (o FolderIamBindingOutput) ToFolderIamBindingOutput() FolderIamBindingOutput {
	return o
}

func (o FolderIamBindingOutput) ToFolderIamBindingOutputWithContext(ctx context.Context) FolderIamBindingOutput {
	return o
}

func (o FolderIamBindingOutput) Condition() FolderIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *FolderIamBinding) FolderIamBindingConditionPtrOutput { return v.Condition }).(FolderIamBindingConditionPtrOutput)
}

func (o FolderIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o FolderIamBindingOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamBinding) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

func (o FolderIamBindingOutput) FolderIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamBinding) pulumi.StringOutput { return v.FolderIamBindingId }).(pulumi.StringOutput)
}

func (o FolderIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FolderIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o FolderIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderIamBindingInput)(nil)).Elem(), &FolderIamBinding{})
	pulumi.RegisterOutputType(FolderIamBindingOutput{})
}
