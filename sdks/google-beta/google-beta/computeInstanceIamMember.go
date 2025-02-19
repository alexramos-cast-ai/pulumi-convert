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

type ComputeInstanceIamMember struct {
	pulumi.CustomResourceState

	ComputeInstanceIamMemberId pulumi.StringOutput                        `pulumi:"computeInstanceIamMemberId"`
	Condition                  ComputeInstanceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                       pulumi.StringOutput                        `pulumi:"etag"`
	InstanceName               pulumi.StringOutput                        `pulumi:"instanceName"`
	Member                     pulumi.StringOutput                        `pulumi:"member"`
	Project                    pulumi.StringOutput                        `pulumi:"project"`
	Role                       pulumi.StringOutput                        `pulumi:"role"`
	Zone                       pulumi.StringOutput                        `pulumi:"zone"`
}

// NewComputeInstanceIamMember registers a new resource with the given unique name, arguments, and options.
func NewComputeInstanceIamMember(ctx *pulumi.Context,
	name string, args *ComputeInstanceIamMemberArgs, opts ...pulumi.ResourceOption) (*ComputeInstanceIamMember, error) {
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
	var resource ComputeInstanceIamMember
	err = ctx.RegisterPackageResource("google-beta:index/computeInstanceIamMember:ComputeInstanceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeInstanceIamMember gets an existing ComputeInstanceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeInstanceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeInstanceIamMemberState, opts ...pulumi.ResourceOption) (*ComputeInstanceIamMember, error) {
	var resource ComputeInstanceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeInstanceIamMember:ComputeInstanceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeInstanceIamMember resources.
type computeInstanceIamMemberState struct {
	ComputeInstanceIamMemberId *string                            `pulumi:"computeInstanceIamMemberId"`
	Condition                  *ComputeInstanceIamMemberCondition `pulumi:"condition"`
	Etag                       *string                            `pulumi:"etag"`
	InstanceName               *string                            `pulumi:"instanceName"`
	Member                     *string                            `pulumi:"member"`
	Project                    *string                            `pulumi:"project"`
	Role                       *string                            `pulumi:"role"`
	Zone                       *string                            `pulumi:"zone"`
}

type ComputeInstanceIamMemberState struct {
	ComputeInstanceIamMemberId pulumi.StringPtrInput
	Condition                  ComputeInstanceIamMemberConditionPtrInput
	Etag                       pulumi.StringPtrInput
	InstanceName               pulumi.StringPtrInput
	Member                     pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	Role                       pulumi.StringPtrInput
	Zone                       pulumi.StringPtrInput
}

func (ComputeInstanceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceIamMemberState)(nil)).Elem()
}

type computeInstanceIamMemberArgs struct {
	ComputeInstanceIamMemberId *string                            `pulumi:"computeInstanceIamMemberId"`
	Condition                  *ComputeInstanceIamMemberCondition `pulumi:"condition"`
	InstanceName               string                             `pulumi:"instanceName"`
	Member                     string                             `pulumi:"member"`
	Project                    *string                            `pulumi:"project"`
	Role                       string                             `pulumi:"role"`
	Zone                       *string                            `pulumi:"zone"`
}

// The set of arguments for constructing a ComputeInstanceIamMember resource.
type ComputeInstanceIamMemberArgs struct {
	ComputeInstanceIamMemberId pulumi.StringPtrInput
	Condition                  ComputeInstanceIamMemberConditionPtrInput
	InstanceName               pulumi.StringInput
	Member                     pulumi.StringInput
	Project                    pulumi.StringPtrInput
	Role                       pulumi.StringInput
	Zone                       pulumi.StringPtrInput
}

func (ComputeInstanceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeInstanceIamMemberArgs)(nil)).Elem()
}

type ComputeInstanceIamMemberInput interface {
	pulumi.Input

	ToComputeInstanceIamMemberOutput() ComputeInstanceIamMemberOutput
	ToComputeInstanceIamMemberOutputWithContext(ctx context.Context) ComputeInstanceIamMemberOutput
}

func (*ComputeInstanceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceIamMember)(nil)).Elem()
}

func (i *ComputeInstanceIamMember) ToComputeInstanceIamMemberOutput() ComputeInstanceIamMemberOutput {
	return i.ToComputeInstanceIamMemberOutputWithContext(context.Background())
}

func (i *ComputeInstanceIamMember) ToComputeInstanceIamMemberOutputWithContext(ctx context.Context) ComputeInstanceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeInstanceIamMemberOutput)
}

type ComputeInstanceIamMemberOutput struct{ *pulumi.OutputState }

func (ComputeInstanceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeInstanceIamMember)(nil)).Elem()
}

func (o ComputeInstanceIamMemberOutput) ToComputeInstanceIamMemberOutput() ComputeInstanceIamMemberOutput {
	return o
}

func (o ComputeInstanceIamMemberOutput) ToComputeInstanceIamMemberOutputWithContext(ctx context.Context) ComputeInstanceIamMemberOutput {
	return o
}

func (o ComputeInstanceIamMemberOutput) ComputeInstanceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.ComputeInstanceIamMemberId }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) Condition() ComputeInstanceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) ComputeInstanceIamMemberConditionPtrOutput { return v.Condition }).(ComputeInstanceIamMemberConditionPtrOutput)
}

func (o ComputeInstanceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ComputeInstanceIamMemberOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeInstanceIamMember) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeInstanceIamMemberInput)(nil)).Elem(), &ComputeInstanceIamMember{})
	pulumi.RegisterOutputType(ComputeInstanceIamMemberOutput{})
}
