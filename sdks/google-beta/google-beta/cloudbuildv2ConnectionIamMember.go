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

type Cloudbuildv2ConnectionIamMember struct {
	pulumi.CustomResourceState

	Cloudbuildv2ConnectionIamMemberId pulumi.StringOutput                               `pulumi:"cloudbuildv2ConnectionIamMemberId"`
	Condition                         Cloudbuildv2ConnectionIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                              pulumi.StringOutput                               `pulumi:"etag"`
	Location                          pulumi.StringOutput                               `pulumi:"location"`
	Member                            pulumi.StringOutput                               `pulumi:"member"`
	Name                              pulumi.StringOutput                               `pulumi:"name"`
	Project                           pulumi.StringOutput                               `pulumi:"project"`
	Role                              pulumi.StringOutput                               `pulumi:"role"`
}

// NewCloudbuildv2ConnectionIamMember registers a new resource with the given unique name, arguments, and options.
func NewCloudbuildv2ConnectionIamMember(ctx *pulumi.Context,
	name string, args *Cloudbuildv2ConnectionIamMemberArgs, opts ...pulumi.ResourceOption) (*Cloudbuildv2ConnectionIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	var resource Cloudbuildv2ConnectionIamMember
	err = ctx.RegisterPackageResource("google-beta:index/cloudbuildv2ConnectionIamMember:Cloudbuildv2ConnectionIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudbuildv2ConnectionIamMember gets an existing Cloudbuildv2ConnectionIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudbuildv2ConnectionIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *Cloudbuildv2ConnectionIamMemberState, opts ...pulumi.ResourceOption) (*Cloudbuildv2ConnectionIamMember, error) {
	var resource Cloudbuildv2ConnectionIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudbuildv2ConnectionIamMember:Cloudbuildv2ConnectionIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Cloudbuildv2ConnectionIamMember resources.
type cloudbuildv2ConnectionIamMemberState struct {
	Cloudbuildv2ConnectionIamMemberId *string                                   `pulumi:"cloudbuildv2ConnectionIamMemberId"`
	Condition                         *Cloudbuildv2ConnectionIamMemberCondition `pulumi:"condition"`
	Etag                              *string                                   `pulumi:"etag"`
	Location                          *string                                   `pulumi:"location"`
	Member                            *string                                   `pulumi:"member"`
	Name                              *string                                   `pulumi:"name"`
	Project                           *string                                   `pulumi:"project"`
	Role                              *string                                   `pulumi:"role"`
}

type Cloudbuildv2ConnectionIamMemberState struct {
	Cloudbuildv2ConnectionIamMemberId pulumi.StringPtrInput
	Condition                         Cloudbuildv2ConnectionIamMemberConditionPtrInput
	Etag                              pulumi.StringPtrInput
	Location                          pulumi.StringPtrInput
	Member                            pulumi.StringPtrInput
	Name                              pulumi.StringPtrInput
	Project                           pulumi.StringPtrInput
	Role                              pulumi.StringPtrInput
}

func (Cloudbuildv2ConnectionIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudbuildv2ConnectionIamMemberState)(nil)).Elem()
}

type cloudbuildv2ConnectionIamMemberArgs struct {
	Cloudbuildv2ConnectionIamMemberId *string                                   `pulumi:"cloudbuildv2ConnectionIamMemberId"`
	Condition                         *Cloudbuildv2ConnectionIamMemberCondition `pulumi:"condition"`
	Location                          *string                                   `pulumi:"location"`
	Member                            string                                    `pulumi:"member"`
	Name                              *string                                   `pulumi:"name"`
	Project                           *string                                   `pulumi:"project"`
	Role                              string                                    `pulumi:"role"`
}

// The set of arguments for constructing a Cloudbuildv2ConnectionIamMember resource.
type Cloudbuildv2ConnectionIamMemberArgs struct {
	Cloudbuildv2ConnectionIamMemberId pulumi.StringPtrInput
	Condition                         Cloudbuildv2ConnectionIamMemberConditionPtrInput
	Location                          pulumi.StringPtrInput
	Member                            pulumi.StringInput
	Name                              pulumi.StringPtrInput
	Project                           pulumi.StringPtrInput
	Role                              pulumi.StringInput
}

func (Cloudbuildv2ConnectionIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudbuildv2ConnectionIamMemberArgs)(nil)).Elem()
}

type Cloudbuildv2ConnectionIamMemberInput interface {
	pulumi.Input

	ToCloudbuildv2ConnectionIamMemberOutput() Cloudbuildv2ConnectionIamMemberOutput
	ToCloudbuildv2ConnectionIamMemberOutputWithContext(ctx context.Context) Cloudbuildv2ConnectionIamMemberOutput
}

func (*Cloudbuildv2ConnectionIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudbuildv2ConnectionIamMember)(nil)).Elem()
}

func (i *Cloudbuildv2ConnectionIamMember) ToCloudbuildv2ConnectionIamMemberOutput() Cloudbuildv2ConnectionIamMemberOutput {
	return i.ToCloudbuildv2ConnectionIamMemberOutputWithContext(context.Background())
}

func (i *Cloudbuildv2ConnectionIamMember) ToCloudbuildv2ConnectionIamMemberOutputWithContext(ctx context.Context) Cloudbuildv2ConnectionIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Cloudbuildv2ConnectionIamMemberOutput)
}

type Cloudbuildv2ConnectionIamMemberOutput struct{ *pulumi.OutputState }

func (Cloudbuildv2ConnectionIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Cloudbuildv2ConnectionIamMember)(nil)).Elem()
}

func (o Cloudbuildv2ConnectionIamMemberOutput) ToCloudbuildv2ConnectionIamMemberOutput() Cloudbuildv2ConnectionIamMemberOutput {
	return o
}

func (o Cloudbuildv2ConnectionIamMemberOutput) ToCloudbuildv2ConnectionIamMemberOutputWithContext(ctx context.Context) Cloudbuildv2ConnectionIamMemberOutput {
	return o
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Cloudbuildv2ConnectionIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput {
		return v.Cloudbuildv2ConnectionIamMemberId
	}).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Condition() Cloudbuildv2ConnectionIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) Cloudbuildv2ConnectionIamMemberConditionPtrOutput {
		return v.Condition
	}).(Cloudbuildv2ConnectionIamMemberConditionPtrOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o Cloudbuildv2ConnectionIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *Cloudbuildv2ConnectionIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*Cloudbuildv2ConnectionIamMemberInput)(nil)).Elem(), &Cloudbuildv2ConnectionIamMember{})
	pulumi.RegisterOutputType(Cloudbuildv2ConnectionIamMemberOutput{})
}
