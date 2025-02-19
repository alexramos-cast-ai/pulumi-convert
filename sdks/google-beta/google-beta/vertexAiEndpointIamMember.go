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

type VertexAiEndpointIamMember struct {
	pulumi.CustomResourceState

	Condition                   VertexAiEndpointIamMemberConditionPtrOutput `pulumi:"condition"`
	Endpoint                    pulumi.StringOutput                         `pulumi:"endpoint"`
	Etag                        pulumi.StringOutput                         `pulumi:"etag"`
	Location                    pulumi.StringOutput                         `pulumi:"location"`
	Member                      pulumi.StringOutput                         `pulumi:"member"`
	Project                     pulumi.StringOutput                         `pulumi:"project"`
	Role                        pulumi.StringOutput                         `pulumi:"role"`
	VertexAiEndpointIamMemberId pulumi.StringOutput                         `pulumi:"vertexAiEndpointIamMemberId"`
}

// NewVertexAiEndpointIamMember registers a new resource with the given unique name, arguments, and options.
func NewVertexAiEndpointIamMember(ctx *pulumi.Context,
	name string, args *VertexAiEndpointIamMemberArgs, opts ...pulumi.ResourceOption) (*VertexAiEndpointIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Endpoint == nil {
		return nil, errors.New("invalid value for required argument 'Endpoint'")
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
	var resource VertexAiEndpointIamMember
	err = ctx.RegisterPackageResource("google-beta:index/vertexAiEndpointIamMember:VertexAiEndpointIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVertexAiEndpointIamMember gets an existing VertexAiEndpointIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVertexAiEndpointIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VertexAiEndpointIamMemberState, opts ...pulumi.ResourceOption) (*VertexAiEndpointIamMember, error) {
	var resource VertexAiEndpointIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/vertexAiEndpointIamMember:VertexAiEndpointIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VertexAiEndpointIamMember resources.
type vertexAiEndpointIamMemberState struct {
	Condition                   *VertexAiEndpointIamMemberCondition `pulumi:"condition"`
	Endpoint                    *string                             `pulumi:"endpoint"`
	Etag                        *string                             `pulumi:"etag"`
	Location                    *string                             `pulumi:"location"`
	Member                      *string                             `pulumi:"member"`
	Project                     *string                             `pulumi:"project"`
	Role                        *string                             `pulumi:"role"`
	VertexAiEndpointIamMemberId *string                             `pulumi:"vertexAiEndpointIamMemberId"`
}

type VertexAiEndpointIamMemberState struct {
	Condition                   VertexAiEndpointIamMemberConditionPtrInput
	Endpoint                    pulumi.StringPtrInput
	Etag                        pulumi.StringPtrInput
	Location                    pulumi.StringPtrInput
	Member                      pulumi.StringPtrInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringPtrInput
	VertexAiEndpointIamMemberId pulumi.StringPtrInput
}

func (VertexAiEndpointIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiEndpointIamMemberState)(nil)).Elem()
}

type vertexAiEndpointIamMemberArgs struct {
	Condition                   *VertexAiEndpointIamMemberCondition `pulumi:"condition"`
	Endpoint                    string                              `pulumi:"endpoint"`
	Location                    *string                             `pulumi:"location"`
	Member                      string                              `pulumi:"member"`
	Project                     *string                             `pulumi:"project"`
	Role                        string                              `pulumi:"role"`
	VertexAiEndpointIamMemberId *string                             `pulumi:"vertexAiEndpointIamMemberId"`
}

// The set of arguments for constructing a VertexAiEndpointIamMember resource.
type VertexAiEndpointIamMemberArgs struct {
	Condition                   VertexAiEndpointIamMemberConditionPtrInput
	Endpoint                    pulumi.StringInput
	Location                    pulumi.StringPtrInput
	Member                      pulumi.StringInput
	Project                     pulumi.StringPtrInput
	Role                        pulumi.StringInput
	VertexAiEndpointIamMemberId pulumi.StringPtrInput
}

func (VertexAiEndpointIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiEndpointIamMemberArgs)(nil)).Elem()
}

type VertexAiEndpointIamMemberInput interface {
	pulumi.Input

	ToVertexAiEndpointIamMemberOutput() VertexAiEndpointIamMemberOutput
	ToVertexAiEndpointIamMemberOutputWithContext(ctx context.Context) VertexAiEndpointIamMemberOutput
}

func (*VertexAiEndpointIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiEndpointIamMember)(nil)).Elem()
}

func (i *VertexAiEndpointIamMember) ToVertexAiEndpointIamMemberOutput() VertexAiEndpointIamMemberOutput {
	return i.ToVertexAiEndpointIamMemberOutputWithContext(context.Background())
}

func (i *VertexAiEndpointIamMember) ToVertexAiEndpointIamMemberOutputWithContext(ctx context.Context) VertexAiEndpointIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VertexAiEndpointIamMemberOutput)
}

type VertexAiEndpointIamMemberOutput struct{ *pulumi.OutputState }

func (VertexAiEndpointIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiEndpointIamMember)(nil)).Elem()
}

func (o VertexAiEndpointIamMemberOutput) ToVertexAiEndpointIamMemberOutput() VertexAiEndpointIamMemberOutput {
	return o
}

func (o VertexAiEndpointIamMemberOutput) ToVertexAiEndpointIamMemberOutputWithContext(ctx context.Context) VertexAiEndpointIamMemberOutput {
	return o
}

func (o VertexAiEndpointIamMemberOutput) Condition() VertexAiEndpointIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) VertexAiEndpointIamMemberConditionPtrOutput { return v.Condition }).(VertexAiEndpointIamMemberConditionPtrOutput)
}

func (o VertexAiEndpointIamMemberOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o VertexAiEndpointIamMemberOutput) VertexAiEndpointIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiEndpointIamMember) pulumi.StringOutput { return v.VertexAiEndpointIamMemberId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VertexAiEndpointIamMemberInput)(nil)).Elem(), &VertexAiEndpointIamMember{})
	pulumi.RegisterOutputType(VertexAiEndpointIamMemberOutput{})
}
