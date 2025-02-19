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

type CloudRunV2ServiceIamMember struct {
	pulumi.CustomResourceState

	CloudRunV2ServiceIamMemberId pulumi.StringOutput                          `pulumi:"cloudRunV2ServiceIamMemberId"`
	Condition                    CloudRunV2ServiceIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                         pulumi.StringOutput                          `pulumi:"etag"`
	Location                     pulumi.StringOutput                          `pulumi:"location"`
	Member                       pulumi.StringOutput                          `pulumi:"member"`
	Name                         pulumi.StringOutput                          `pulumi:"name"`
	Project                      pulumi.StringOutput                          `pulumi:"project"`
	Role                         pulumi.StringOutput                          `pulumi:"role"`
}

// NewCloudRunV2ServiceIamMember registers a new resource with the given unique name, arguments, and options.
func NewCloudRunV2ServiceIamMember(ctx *pulumi.Context,
	name string, args *CloudRunV2ServiceIamMemberArgs, opts ...pulumi.ResourceOption) (*CloudRunV2ServiceIamMember, error) {
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
	var resource CloudRunV2ServiceIamMember
	err = ctx.RegisterPackageResource("google:index/cloudRunV2ServiceIamMember:CloudRunV2ServiceIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudRunV2ServiceIamMember gets an existing CloudRunV2ServiceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudRunV2ServiceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudRunV2ServiceIamMemberState, opts ...pulumi.ResourceOption) (*CloudRunV2ServiceIamMember, error) {
	var resource CloudRunV2ServiceIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/cloudRunV2ServiceIamMember:CloudRunV2ServiceIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudRunV2ServiceIamMember resources.
type cloudRunV2ServiceIamMemberState struct {
	CloudRunV2ServiceIamMemberId *string                              `pulumi:"cloudRunV2ServiceIamMemberId"`
	Condition                    *CloudRunV2ServiceIamMemberCondition `pulumi:"condition"`
	Etag                         *string                              `pulumi:"etag"`
	Location                     *string                              `pulumi:"location"`
	Member                       *string                              `pulumi:"member"`
	Name                         *string                              `pulumi:"name"`
	Project                      *string                              `pulumi:"project"`
	Role                         *string                              `pulumi:"role"`
}

type CloudRunV2ServiceIamMemberState struct {
	CloudRunV2ServiceIamMemberId pulumi.StringPtrInput
	Condition                    CloudRunV2ServiceIamMemberConditionPtrInput
	Etag                         pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	Member                       pulumi.StringPtrInput
	Name                         pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringPtrInput
}

func (CloudRunV2ServiceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunV2ServiceIamMemberState)(nil)).Elem()
}

type cloudRunV2ServiceIamMemberArgs struct {
	CloudRunV2ServiceIamMemberId *string                              `pulumi:"cloudRunV2ServiceIamMemberId"`
	Condition                    *CloudRunV2ServiceIamMemberCondition `pulumi:"condition"`
	Location                     *string                              `pulumi:"location"`
	Member                       string                               `pulumi:"member"`
	Name                         *string                              `pulumi:"name"`
	Project                      *string                              `pulumi:"project"`
	Role                         string                               `pulumi:"role"`
}

// The set of arguments for constructing a CloudRunV2ServiceIamMember resource.
type CloudRunV2ServiceIamMemberArgs struct {
	CloudRunV2ServiceIamMemberId pulumi.StringPtrInput
	Condition                    CloudRunV2ServiceIamMemberConditionPtrInput
	Location                     pulumi.StringPtrInput
	Member                       pulumi.StringInput
	Name                         pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Role                         pulumi.StringInput
}

func (CloudRunV2ServiceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunV2ServiceIamMemberArgs)(nil)).Elem()
}

type CloudRunV2ServiceIamMemberInput interface {
	pulumi.Input

	ToCloudRunV2ServiceIamMemberOutput() CloudRunV2ServiceIamMemberOutput
	ToCloudRunV2ServiceIamMemberOutputWithContext(ctx context.Context) CloudRunV2ServiceIamMemberOutput
}

func (*CloudRunV2ServiceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunV2ServiceIamMember)(nil)).Elem()
}

func (i *CloudRunV2ServiceIamMember) ToCloudRunV2ServiceIamMemberOutput() CloudRunV2ServiceIamMemberOutput {
	return i.ToCloudRunV2ServiceIamMemberOutputWithContext(context.Background())
}

func (i *CloudRunV2ServiceIamMember) ToCloudRunV2ServiceIamMemberOutputWithContext(ctx context.Context) CloudRunV2ServiceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRunV2ServiceIamMemberOutput)
}

type CloudRunV2ServiceIamMemberOutput struct{ *pulumi.OutputState }

func (CloudRunV2ServiceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunV2ServiceIamMember)(nil)).Elem()
}

func (o CloudRunV2ServiceIamMemberOutput) ToCloudRunV2ServiceIamMemberOutput() CloudRunV2ServiceIamMemberOutput {
	return o
}

func (o CloudRunV2ServiceIamMemberOutput) ToCloudRunV2ServiceIamMemberOutputWithContext(ctx context.Context) CloudRunV2ServiceIamMemberOutput {
	return o
}

func (o CloudRunV2ServiceIamMemberOutput) CloudRunV2ServiceIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.CloudRunV2ServiceIamMemberId }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Condition() CloudRunV2ServiceIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) CloudRunV2ServiceIamMemberConditionPtrOutput { return v.Condition }).(CloudRunV2ServiceIamMemberConditionPtrOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRunV2ServiceIamMemberInput)(nil)).Elem(), &CloudRunV2ServiceIamMember{})
	pulumi.RegisterOutputType(CloudRunV2ServiceIamMemberOutput{})
}
