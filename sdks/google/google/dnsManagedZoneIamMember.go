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

type DnsManagedZoneIamMember struct {
	pulumi.CustomResourceState

	Condition                 DnsManagedZoneIamMemberConditionPtrOutput `pulumi:"condition"`
	DnsManagedZoneIamMemberId pulumi.StringOutput                       `pulumi:"dnsManagedZoneIamMemberId"`
	Etag                      pulumi.StringOutput                       `pulumi:"etag"`
	ManagedZone               pulumi.StringOutput                       `pulumi:"managedZone"`
	Member                    pulumi.StringOutput                       `pulumi:"member"`
	Project                   pulumi.StringOutput                       `pulumi:"project"`
	Role                      pulumi.StringOutput                       `pulumi:"role"`
}

// NewDnsManagedZoneIamMember registers a new resource with the given unique name, arguments, and options.
func NewDnsManagedZoneIamMember(ctx *pulumi.Context,
	name string, args *DnsManagedZoneIamMemberArgs, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedZone == nil {
		return nil, errors.New("invalid value for required argument 'ManagedZone'")
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
	var resource DnsManagedZoneIamMember
	err = ctx.RegisterPackageResource("google:index/dnsManagedZoneIamMember:DnsManagedZoneIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsManagedZoneIamMember gets an existing DnsManagedZoneIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsManagedZoneIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsManagedZoneIamMemberState, opts ...pulumi.ResourceOption) (*DnsManagedZoneIamMember, error) {
	var resource DnsManagedZoneIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dnsManagedZoneIamMember:DnsManagedZoneIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsManagedZoneIamMember resources.
type dnsManagedZoneIamMemberState struct {
	Condition                 *DnsManagedZoneIamMemberCondition `pulumi:"condition"`
	DnsManagedZoneIamMemberId *string                           `pulumi:"dnsManagedZoneIamMemberId"`
	Etag                      *string                           `pulumi:"etag"`
	ManagedZone               *string                           `pulumi:"managedZone"`
	Member                    *string                           `pulumi:"member"`
	Project                   *string                           `pulumi:"project"`
	Role                      *string                           `pulumi:"role"`
}

type DnsManagedZoneIamMemberState struct {
	Condition                 DnsManagedZoneIamMemberConditionPtrInput
	DnsManagedZoneIamMemberId pulumi.StringPtrInput
	Etag                      pulumi.StringPtrInput
	ManagedZone               pulumi.StringPtrInput
	Member                    pulumi.StringPtrInput
	Project                   pulumi.StringPtrInput
	Role                      pulumi.StringPtrInput
}

func (DnsManagedZoneIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamMemberState)(nil)).Elem()
}

type dnsManagedZoneIamMemberArgs struct {
	Condition                 *DnsManagedZoneIamMemberCondition `pulumi:"condition"`
	DnsManagedZoneIamMemberId *string                           `pulumi:"dnsManagedZoneIamMemberId"`
	ManagedZone               string                            `pulumi:"managedZone"`
	Member                    string                            `pulumi:"member"`
	Project                   *string                           `pulumi:"project"`
	Role                      string                            `pulumi:"role"`
}

// The set of arguments for constructing a DnsManagedZoneIamMember resource.
type DnsManagedZoneIamMemberArgs struct {
	Condition                 DnsManagedZoneIamMemberConditionPtrInput
	DnsManagedZoneIamMemberId pulumi.StringPtrInput
	ManagedZone               pulumi.StringInput
	Member                    pulumi.StringInput
	Project                   pulumi.StringPtrInput
	Role                      pulumi.StringInput
}

func (DnsManagedZoneIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsManagedZoneIamMemberArgs)(nil)).Elem()
}

type DnsManagedZoneIamMemberInput interface {
	pulumi.Input

	ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput
	ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput
}

func (*DnsManagedZoneIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamMember)(nil)).Elem()
}

func (i *DnsManagedZoneIamMember) ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput {
	return i.ToDnsManagedZoneIamMemberOutputWithContext(context.Background())
}

func (i *DnsManagedZoneIamMember) ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsManagedZoneIamMemberOutput)
}

type DnsManagedZoneIamMemberOutput struct{ *pulumi.OutputState }

func (DnsManagedZoneIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsManagedZoneIamMember)(nil)).Elem()
}

func (o DnsManagedZoneIamMemberOutput) ToDnsManagedZoneIamMemberOutput() DnsManagedZoneIamMemberOutput {
	return o
}

func (o DnsManagedZoneIamMemberOutput) ToDnsManagedZoneIamMemberOutputWithContext(ctx context.Context) DnsManagedZoneIamMemberOutput {
	return o
}

func (o DnsManagedZoneIamMemberOutput) Condition() DnsManagedZoneIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) DnsManagedZoneIamMemberConditionPtrOutput { return v.Condition }).(DnsManagedZoneIamMemberConditionPtrOutput)
}

func (o DnsManagedZoneIamMemberOutput) DnsManagedZoneIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.DnsManagedZoneIamMemberId }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) ManagedZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.ManagedZone }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DnsManagedZoneIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DnsManagedZoneIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsManagedZoneIamMemberInput)(nil)).Elem(), &DnsManagedZoneIamMember{})
	pulumi.RegisterOutputType(DnsManagedZoneIamMemberOutput{})
}
