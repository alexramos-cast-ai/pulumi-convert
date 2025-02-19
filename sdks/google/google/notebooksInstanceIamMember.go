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

type NotebooksInstanceIamMember struct {
	pulumi.CustomResourceState

	Condition                    NotebooksInstanceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                         pulumi.StringOutput                          `pulumi:"etag"`
	InstanceName                 pulumi.StringOutput                          `pulumi:"instanceName"`
	Location                     pulumi.StringOutput                          `pulumi:"location"`
	Member                       pulumi.StringOutput                          `pulumi:"member"`
	NotebooksInstanceIamMemberId pulumi.StringOutput                          `pulumi:"notebooksInstanceIamMemberId"`
	Project                      pulumi.StringOutput                          `pulumi:"project"`
	Role                         pulumi.StringOutput                          `pulumi:"role"`
}

// NewNotebooksInstanceIamMember registers a new resource with the given unique name, arguments, and options.
func NewNotebooksInstanceIamMember(ctx *pulumi.Context,
	name string, args *NotebooksInstanceIamMemberArgs, opts ...pulumi.ResourceOption) (*NotebooksInstanceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
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
	var resource NotebooksInstanceIamMember
	err = ctx.RegisterPackageResource("google:index/notebooksInstanceIamMember:NotebooksInstanceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotebooksInstanceIamMember gets an existing NotebooksInstanceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotebooksInstanceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotebooksInstanceIamMemberState, opts ...pulumi.ResourceOption) (*NotebooksInstanceIamMember, error) {
	var resource NotebooksInstanceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/notebooksInstanceIamMember:NotebooksInstanceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NotebooksInstanceIamMember resources.
type notebooksInstanceIamMemberState struct {
	Condition                    *NotebooksInstanceIamMemberCondition `pulumi:"condition"`
	Etag                         *string                              `pulumi:"etag"`
	InstanceName                 *string                              `pulumi:"instanceName"`
	Location                     *string                              `pulumi:"location"`
	Member                       *string                              `pulumi:"member"`
	NotebooksInstanceIamMemberId *string                              `pulumi:"notebooksInstanceIamMemberId"`
	Project                      *string                              `pulumi:"project"`
	Role                         *string                              `pulumi:"role"`
}

type NotebooksInstanceIamMemberState struct {
	Condition                    NotebooksInstanceIamMemberConditionPtrInput
	Etag                         pulumi.StringPtrInput
	InstanceName                 pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Member                       pulumi.StringPtrInput
	NotebooksInstanceIamMemberId pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringPtrInput
}

func (NotebooksInstanceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksInstanceIamMemberState)(nil)).Elem()
}

type notebooksInstanceIamMemberArgs struct {
	Condition                    *NotebooksInstanceIamMemberCondition `pulumi:"condition"`
	InstanceName                 string                               `pulumi:"instanceName"`
	Location                     *string                              `pulumi:"location"`
	Member                       string                               `pulumi:"member"`
	NotebooksInstanceIamMemberId *string                              `pulumi:"notebooksInstanceIamMemberId"`
	Project                      *string                              `pulumi:"project"`
	Role                         string                               `pulumi:"role"`
}

// The set of arguments for constructing a NotebooksInstanceIamMember resource.
type NotebooksInstanceIamMemberArgs struct {
	Condition                    NotebooksInstanceIamMemberConditionPtrInput
	InstanceName                 pulumi.StringInput
	Location                     pulumi.StringPtrInput
	Member                       pulumi.StringInput
	NotebooksInstanceIamMemberId pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringInput
}

func (NotebooksInstanceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notebooksInstanceIamMemberArgs)(nil)).Elem()
}

type NotebooksInstanceIamMemberInput interface {
	pulumi.Input

	ToNotebooksInstanceIamMemberOutput() NotebooksInstanceIamMemberOutput
	ToNotebooksInstanceIamMemberOutputWithContext(ctx context.Context) NotebooksInstanceIamMemberOutput
}

func (*NotebooksInstanceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksInstanceIamMember)(nil)).Elem()
}

func (i *NotebooksInstanceIamMember) ToNotebooksInstanceIamMemberOutput() NotebooksInstanceIamMemberOutput {
	return i.ToNotebooksInstanceIamMemberOutputWithContext(context.Background())
}

func (i *NotebooksInstanceIamMember) ToNotebooksInstanceIamMemberOutputWithContext(ctx context.Context) NotebooksInstanceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotebooksInstanceIamMemberOutput)
}

type NotebooksInstanceIamMemberOutput struct{ *pulumi.OutputState }

func (NotebooksInstanceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotebooksInstanceIamMember)(nil)).Elem()
}

func (o NotebooksInstanceIamMemberOutput) ToNotebooksInstanceIamMemberOutput() NotebooksInstanceIamMemberOutput {
	return o
}

func (o NotebooksInstanceIamMemberOutput) ToNotebooksInstanceIamMemberOutputWithContext(ctx context.Context) NotebooksInstanceIamMemberOutput {
	return o
}

func (o NotebooksInstanceIamMemberOutput) Condition() NotebooksInstanceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) NotebooksInstanceIamMemberConditionPtrOutput { return v.Condition }).(NotebooksInstanceIamMemberConditionPtrOutput)
}

func (o NotebooksInstanceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) NotebooksInstanceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.NotebooksInstanceIamMemberId }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o NotebooksInstanceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *NotebooksInstanceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotebooksInstanceIamMemberInput)(nil)).Elem(), &NotebooksInstanceIamMember{})
	pulumi.RegisterOutputType(NotebooksInstanceIamMemberOutput{})
}
