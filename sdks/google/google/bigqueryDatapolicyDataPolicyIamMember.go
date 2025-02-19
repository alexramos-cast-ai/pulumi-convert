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

type BigqueryDatapolicyDataPolicyIamMember struct {
	pulumi.CustomResourceState

	BigqueryDatapolicyDataPolicyIamMemberId pulumi.StringOutput                                     `pulumi:"bigqueryDatapolicyDataPolicyIamMemberId"`
	Condition                               BigqueryDatapolicyDataPolicyIamMemberConditionPtrOutput `pulumi:"condition"`
	DataPolicyId                            pulumi.StringOutput                                     `pulumi:"dataPolicyId"`
	Etag                                    pulumi.StringOutput                                     `pulumi:"etag"`
	Location                                pulumi.StringOutput                                     `pulumi:"location"`
	Member                                  pulumi.StringOutput                                     `pulumi:"member"`
	Project                                 pulumi.StringOutput                                     `pulumi:"project"`
	Role                                    pulumi.StringOutput                                     `pulumi:"role"`
}

// NewBigqueryDatapolicyDataPolicyIamMember registers a new resource with the given unique name, arguments, and options.
func NewBigqueryDatapolicyDataPolicyIamMember(ctx *pulumi.Context,
	name string, args *BigqueryDatapolicyDataPolicyIamMemberArgs, opts ...pulumi.ResourceOption) (*BigqueryDatapolicyDataPolicyIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataPolicyId == nil {
		return nil, errors.New("invalid value for required argument 'DataPolicyId'")
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
	var resource BigqueryDatapolicyDataPolicyIamMember
	err = ctx.RegisterPackageResource("google:index/bigqueryDatapolicyDataPolicyIamMember:BigqueryDatapolicyDataPolicyIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryDatapolicyDataPolicyIamMember gets an existing BigqueryDatapolicyDataPolicyIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryDatapolicyDataPolicyIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryDatapolicyDataPolicyIamMemberState, opts ...pulumi.ResourceOption) (*BigqueryDatapolicyDataPolicyIamMember, error) {
	var resource BigqueryDatapolicyDataPolicyIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigqueryDatapolicyDataPolicyIamMember:BigqueryDatapolicyDataPolicyIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryDatapolicyDataPolicyIamMember resources.
type bigqueryDatapolicyDataPolicyIamMemberState struct {
	BigqueryDatapolicyDataPolicyIamMemberId *string                                         `pulumi:"bigqueryDatapolicyDataPolicyIamMemberId"`
	Condition                               *BigqueryDatapolicyDataPolicyIamMemberCondition `pulumi:"condition"`
	DataPolicyId                            *string                                         `pulumi:"dataPolicyId"`
	Etag                                    *string                                         `pulumi:"etag"`
	Location                                *string                                         `pulumi:"location"`
	Member                                  *string                                         `pulumi:"member"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    *string                                         `pulumi:"role"`
}

type BigqueryDatapolicyDataPolicyIamMemberState struct {
	BigqueryDatapolicyDataPolicyIamMemberId pulumi.StringPtrInput
	Condition                               BigqueryDatapolicyDataPolicyIamMemberConditionPtrInput
	DataPolicyId                            pulumi.StringPtrInput
	Etag                                    pulumi.StringPtrInput
	Location                                pulumi.StringPtrInput
	Member                                  pulumi.StringPtrInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringPtrInput
}

func (BigqueryDatapolicyDataPolicyIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatapolicyDataPolicyIamMemberState)(nil)).Elem()
}

type bigqueryDatapolicyDataPolicyIamMemberArgs struct {
	BigqueryDatapolicyDataPolicyIamMemberId *string                                         `pulumi:"bigqueryDatapolicyDataPolicyIamMemberId"`
	Condition                               *BigqueryDatapolicyDataPolicyIamMemberCondition `pulumi:"condition"`
	DataPolicyId                            string                                          `pulumi:"dataPolicyId"`
	Location                                *string                                         `pulumi:"location"`
	Member                                  string                                          `pulumi:"member"`
	Project                                 *string                                         `pulumi:"project"`
	Role                                    string                                          `pulumi:"role"`
}

// The set of arguments for constructing a BigqueryDatapolicyDataPolicyIamMember resource.
type BigqueryDatapolicyDataPolicyIamMemberArgs struct {
	BigqueryDatapolicyDataPolicyIamMemberId pulumi.StringPtrInput
	Condition                               BigqueryDatapolicyDataPolicyIamMemberConditionPtrInput
	DataPolicyId                            pulumi.StringInput
	Location                                pulumi.StringPtrInput
	Member                                  pulumi.StringInput
	Project                                 pulumi.StringPtrInput
	Role                                    pulumi.StringInput
}

func (BigqueryDatapolicyDataPolicyIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatapolicyDataPolicyIamMemberArgs)(nil)).Elem()
}

type BigqueryDatapolicyDataPolicyIamMemberInput interface {
	pulumi.Input

	ToBigqueryDatapolicyDataPolicyIamMemberOutput() BigqueryDatapolicyDataPolicyIamMemberOutput
	ToBigqueryDatapolicyDataPolicyIamMemberOutputWithContext(ctx context.Context) BigqueryDatapolicyDataPolicyIamMemberOutput
}

func (*BigqueryDatapolicyDataPolicyIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatapolicyDataPolicyIamMember)(nil)).Elem()
}

func (i *BigqueryDatapolicyDataPolicyIamMember) ToBigqueryDatapolicyDataPolicyIamMemberOutput() BigqueryDatapolicyDataPolicyIamMemberOutput {
	return i.ToBigqueryDatapolicyDataPolicyIamMemberOutputWithContext(context.Background())
}

func (i *BigqueryDatapolicyDataPolicyIamMember) ToBigqueryDatapolicyDataPolicyIamMemberOutputWithContext(ctx context.Context) BigqueryDatapolicyDataPolicyIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryDatapolicyDataPolicyIamMemberOutput)
}

type BigqueryDatapolicyDataPolicyIamMemberOutput struct{ *pulumi.OutputState }

func (BigqueryDatapolicyDataPolicyIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatapolicyDataPolicyIamMember)(nil)).Elem()
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) ToBigqueryDatapolicyDataPolicyIamMemberOutput() BigqueryDatapolicyDataPolicyIamMemberOutput {
	return o
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) ToBigqueryDatapolicyDataPolicyIamMemberOutputWithContext(ctx context.Context) BigqueryDatapolicyDataPolicyIamMemberOutput {
	return o
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) BigqueryDatapolicyDataPolicyIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput {
		return v.BigqueryDatapolicyDataPolicyIamMemberId
	}).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Condition() BigqueryDatapolicyDataPolicyIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) BigqueryDatapolicyDataPolicyIamMemberConditionPtrOutput {
		return v.Condition
	}).(BigqueryDatapolicyDataPolicyIamMemberConditionPtrOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) DataPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.DataPolicyId }).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o BigqueryDatapolicyDataPolicyIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatapolicyDataPolicyIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryDatapolicyDataPolicyIamMemberInput)(nil)).Elem(), &BigqueryDatapolicyDataPolicyIamMember{})
	pulumi.RegisterOutputType(BigqueryDatapolicyDataPolicyIamMemberOutput{})
}
