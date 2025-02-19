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

type SecretManagerRegionalSecretIamMember struct {
	pulumi.CustomResourceState

	Condition                              SecretManagerRegionalSecretIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                                   pulumi.StringOutput                                    `pulumi:"etag"`
	Location                               pulumi.StringOutput                                    `pulumi:"location"`
	Member                                 pulumi.StringOutput                                    `pulumi:"member"`
	Project                                pulumi.StringOutput                                    `pulumi:"project"`
	Role                                   pulumi.StringOutput                                    `pulumi:"role"`
	SecretId                               pulumi.StringOutput                                    `pulumi:"secretId"`
	SecretManagerRegionalSecretIamMemberId pulumi.StringOutput                                    `pulumi:"secretManagerRegionalSecretIamMemberId"`
}

// NewSecretManagerRegionalSecretIamMember registers a new resource with the given unique name, arguments, and options.
func NewSecretManagerRegionalSecretIamMember(ctx *pulumi.Context,
	name string, args *SecretManagerRegionalSecretIamMemberArgs, opts ...pulumi.ResourceOption) (*SecretManagerRegionalSecretIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.SecretId == nil {
		return nil, errors.New("invalid value for required argument 'SecretId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SecretManagerRegionalSecretIamMember
	err = ctx.RegisterPackageResource("google:index/secretManagerRegionalSecretIamMember:SecretManagerRegionalSecretIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretManagerRegionalSecretIamMember gets an existing SecretManagerRegionalSecretIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretManagerRegionalSecretIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretManagerRegionalSecretIamMemberState, opts ...pulumi.ResourceOption) (*SecretManagerRegionalSecretIamMember, error) {
	var resource SecretManagerRegionalSecretIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/secretManagerRegionalSecretIamMember:SecretManagerRegionalSecretIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretManagerRegionalSecretIamMember resources.
type secretManagerRegionalSecretIamMemberState struct {
	Condition                              *SecretManagerRegionalSecretIamMemberCondition `pulumi:"condition"`
	Etag                                   *string                                        `pulumi:"etag"`
	Location                               *string                                        `pulumi:"location"`
	Member                                 *string                                        `pulumi:"member"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   *string                                        `pulumi:"role"`
	SecretId                               *string                                        `pulumi:"secretId"`
	SecretManagerRegionalSecretIamMemberId *string                                        `pulumi:"secretManagerRegionalSecretIamMemberId"`
}

type SecretManagerRegionalSecretIamMemberState struct {
	Condition                              SecretManagerRegionalSecretIamMemberConditionPtrInput
	Etag                                   pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	Member                                 pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringPtrInput
	SecretId                               pulumi.StringPtrInput
	SecretManagerRegionalSecretIamMemberId pulumi.StringPtrInput
}

func (SecretManagerRegionalSecretIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerRegionalSecretIamMemberState)(nil)).Elem()
}

type secretManagerRegionalSecretIamMemberArgs struct {
	Condition                              *SecretManagerRegionalSecretIamMemberCondition `pulumi:"condition"`
	Location                               *string                                        `pulumi:"location"`
	Member                                 string                                         `pulumi:"member"`
	Project                                *string                                        `pulumi:"project"`
	Role                                   string                                         `pulumi:"role"`
	SecretId                               string                                         `pulumi:"secretId"`
	SecretManagerRegionalSecretIamMemberId *string                                        `pulumi:"secretManagerRegionalSecretIamMemberId"`
}

// The set of arguments for constructing a SecretManagerRegionalSecretIamMember resource.
type SecretManagerRegionalSecretIamMemberArgs struct {
	Condition                              SecretManagerRegionalSecretIamMemberConditionPtrInput
	Location                               pulumi.StringPtrInput
	Member                                 pulumi.StringInput
	Project                                pulumi.StringPtrInput
	Role                                   pulumi.StringInput
	SecretId                               pulumi.StringInput
	SecretManagerRegionalSecretIamMemberId pulumi.StringPtrInput
}

func (SecretManagerRegionalSecretIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretManagerRegionalSecretIamMemberArgs)(nil)).Elem()
}

type SecretManagerRegionalSecretIamMemberInput interface {
	pulumi.Input

	ToSecretManagerRegionalSecretIamMemberOutput() SecretManagerRegionalSecretIamMemberOutput
	ToSecretManagerRegionalSecretIamMemberOutputWithContext(ctx context.Context) SecretManagerRegionalSecretIamMemberOutput
}

func (*SecretManagerRegionalSecretIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerRegionalSecretIamMember)(nil)).Elem()
}

func (i *SecretManagerRegionalSecretIamMember) ToSecretManagerRegionalSecretIamMemberOutput() SecretManagerRegionalSecretIamMemberOutput {
	return i.ToSecretManagerRegionalSecretIamMemberOutputWithContext(context.Background())
}

func (i *SecretManagerRegionalSecretIamMember) ToSecretManagerRegionalSecretIamMemberOutputWithContext(ctx context.Context) SecretManagerRegionalSecretIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretManagerRegionalSecretIamMemberOutput)
}

type SecretManagerRegionalSecretIamMemberOutput struct{ *pulumi.OutputState }

func (SecretManagerRegionalSecretIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretManagerRegionalSecretIamMember)(nil)).Elem()
}

func (o SecretManagerRegionalSecretIamMemberOutput) ToSecretManagerRegionalSecretIamMemberOutput() SecretManagerRegionalSecretIamMemberOutput {
	return o
}

func (o SecretManagerRegionalSecretIamMemberOutput) ToSecretManagerRegionalSecretIamMemberOutputWithContext(ctx context.Context) SecretManagerRegionalSecretIamMemberOutput {
	return o
}

func (o SecretManagerRegionalSecretIamMemberOutput) Condition() SecretManagerRegionalSecretIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) SecretManagerRegionalSecretIamMemberConditionPtrOutput {
		return v.Condition
	}).(SecretManagerRegionalSecretIamMemberConditionPtrOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) SecretId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput { return v.SecretId }).(pulumi.StringOutput)
}

func (o SecretManagerRegionalSecretIamMemberOutput) SecretManagerRegionalSecretIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretManagerRegionalSecretIamMember) pulumi.StringOutput {
		return v.SecretManagerRegionalSecretIamMemberId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretManagerRegionalSecretIamMemberInput)(nil)).Elem(), &SecretManagerRegionalSecretIamMember{})
	pulumi.RegisterOutputType(SecretManagerRegionalSecretIamMemberOutput{})
}
