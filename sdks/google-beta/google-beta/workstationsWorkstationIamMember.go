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

type WorkstationsWorkstationIamMember struct {
	pulumi.CustomResourceState

	Condition                          WorkstationsWorkstationIamMemberConditionPtrOutput `pulumi:"condition"`
	Etag                               pulumi.StringOutput                                `pulumi:"etag"`
	Location                           pulumi.StringOutput                                `pulumi:"location"`
	Member                             pulumi.StringOutput                                `pulumi:"member"`
	Project                            pulumi.StringOutput                                `pulumi:"project"`
	Role                               pulumi.StringOutput                                `pulumi:"role"`
	WorkstationClusterId               pulumi.StringOutput                                `pulumi:"workstationClusterId"`
	WorkstationConfigId                pulumi.StringOutput                                `pulumi:"workstationConfigId"`
	WorkstationId                      pulumi.StringOutput                                `pulumi:"workstationId"`
	WorkstationsWorkstationIamMemberId pulumi.StringOutput                                `pulumi:"workstationsWorkstationIamMemberId"`
}

// NewWorkstationsWorkstationIamMember registers a new resource with the given unique name, arguments, and options.
func NewWorkstationsWorkstationIamMember(ctx *pulumi.Context,
	name string, args *WorkstationsWorkstationIamMemberArgs, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.WorkstationClusterId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationClusterId'")
	}
	if args.WorkstationConfigId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationConfigId'")
	}
	if args.WorkstationId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource WorkstationsWorkstationIamMember
	err = ctx.RegisterPackageResource("google-beta:index/workstationsWorkstationIamMember:WorkstationsWorkstationIamMember", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationsWorkstationIamMember gets an existing WorkstationsWorkstationIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationsWorkstationIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationsWorkstationIamMemberState, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationIamMember, error) {
	var resource WorkstationsWorkstationIamMember
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/workstationsWorkstationIamMember:WorkstationsWorkstationIamMember", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationsWorkstationIamMember resources.
type workstationsWorkstationIamMemberState struct {
	Condition                          *WorkstationsWorkstationIamMemberCondition `pulumi:"condition"`
	Etag                               *string                                    `pulumi:"etag"`
	Location                           *string                                    `pulumi:"location"`
	Member                             *string                                    `pulumi:"member"`
	Project                            *string                                    `pulumi:"project"`
	Role                               *string                                    `pulumi:"role"`
	WorkstationClusterId               *string                                    `pulumi:"workstationClusterId"`
	WorkstationConfigId                *string                                    `pulumi:"workstationConfigId"`
	WorkstationId                      *string                                    `pulumi:"workstationId"`
	WorkstationsWorkstationIamMemberId *string                                    `pulumi:"workstationsWorkstationIamMemberId"`
}

type WorkstationsWorkstationIamMemberState struct {
	Condition                          WorkstationsWorkstationIamMemberConditionPtrInput
	Etag                               pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	Member                             pulumi.StringPtrInput
	Project                            pulumi.StringPtrInput
	Role                               pulumi.StringPtrInput
	WorkstationClusterId               pulumi.StringPtrInput
	WorkstationConfigId                pulumi.StringPtrInput
	WorkstationId                      pulumi.StringPtrInput
	WorkstationsWorkstationIamMemberId pulumi.StringPtrInput
}

func (WorkstationsWorkstationIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationIamMemberState)(nil)).Elem()
}

type workstationsWorkstationIamMemberArgs struct {
	Condition                          *WorkstationsWorkstationIamMemberCondition `pulumi:"condition"`
	Location                           *string                                    `pulumi:"location"`
	Member                             string                                     `pulumi:"member"`
	Project                            *string                                    `pulumi:"project"`
	Role                               string                                     `pulumi:"role"`
	WorkstationClusterId               string                                     `pulumi:"workstationClusterId"`
	WorkstationConfigId                string                                     `pulumi:"workstationConfigId"`
	WorkstationId                      string                                     `pulumi:"workstationId"`
	WorkstationsWorkstationIamMemberId *string                                    `pulumi:"workstationsWorkstationIamMemberId"`
}

// The set of arguments for constructing a WorkstationsWorkstationIamMember resource.
type WorkstationsWorkstationIamMemberArgs struct {
	Condition                          WorkstationsWorkstationIamMemberConditionPtrInput
	Location                           pulumi.StringPtrInput
	Member                             pulumi.StringInput
	Project                            pulumi.StringPtrInput
	Role                               pulumi.StringInput
	WorkstationClusterId               pulumi.StringInput
	WorkstationConfigId                pulumi.StringInput
	WorkstationId                      pulumi.StringInput
	WorkstationsWorkstationIamMemberId pulumi.StringPtrInput
}

func (WorkstationsWorkstationIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationIamMemberArgs)(nil)).Elem()
}

type WorkstationsWorkstationIamMemberInput interface {
	pulumi.Input

	ToWorkstationsWorkstationIamMemberOutput() WorkstationsWorkstationIamMemberOutput
	ToWorkstationsWorkstationIamMemberOutputWithContext(ctx context.Context) WorkstationsWorkstationIamMemberOutput
}

func (*WorkstationsWorkstationIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationIamMember)(nil)).Elem()
}

func (i *WorkstationsWorkstationIamMember) ToWorkstationsWorkstationIamMemberOutput() WorkstationsWorkstationIamMemberOutput {
	return i.ToWorkstationsWorkstationIamMemberOutputWithContext(context.Background())
}

func (i *WorkstationsWorkstationIamMember) ToWorkstationsWorkstationIamMemberOutputWithContext(ctx context.Context) WorkstationsWorkstationIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationsWorkstationIamMemberOutput)
}

type WorkstationsWorkstationIamMemberOutput struct{ *pulumi.OutputState }

func (WorkstationsWorkstationIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationIamMember)(nil)).Elem()
}

func (o WorkstationsWorkstationIamMemberOutput) ToWorkstationsWorkstationIamMemberOutput() WorkstationsWorkstationIamMemberOutput {
	return o
}

func (o WorkstationsWorkstationIamMemberOutput) ToWorkstationsWorkstationIamMemberOutputWithContext(ctx context.Context) WorkstationsWorkstationIamMemberOutput {
	return o
}

func (o WorkstationsWorkstationIamMemberOutput) Condition() WorkstationsWorkstationIamMemberConditionPtrOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) WorkstationsWorkstationIamMemberConditionPtrOutput {
		return v.Condition
	}).(WorkstationsWorkstationIamMemberConditionPtrOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.Member }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) WorkstationId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput { return v.WorkstationId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationIamMemberOutput) WorkstationsWorkstationIamMemberId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationIamMember) pulumi.StringOutput {
		return v.WorkstationsWorkstationIamMemberId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationsWorkstationIamMemberInput)(nil)).Elem(), &WorkstationsWorkstationIamMember{})
	pulumi.RegisterOutputType(WorkstationsWorkstationIamMemberOutput{})
}
