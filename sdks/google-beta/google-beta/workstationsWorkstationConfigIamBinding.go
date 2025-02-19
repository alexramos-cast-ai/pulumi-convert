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

type WorkstationsWorkstationConfigIamBinding struct {
	pulumi.CustomResourceState

	Condition                                 WorkstationsWorkstationConfigIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                      pulumi.StringOutput                                       `pulumi:"etag"`
	Location                                  pulumi.StringOutput                                       `pulumi:"location"`
	Members                                   pulumi.StringArrayOutput                                  `pulumi:"members"`
	Project                                   pulumi.StringOutput                                       `pulumi:"project"`
	Role                                      pulumi.StringOutput                                       `pulumi:"role"`
	WorkstationClusterId                      pulumi.StringOutput                                       `pulumi:"workstationClusterId"`
	WorkstationConfigId                       pulumi.StringOutput                                       `pulumi:"workstationConfigId"`
	WorkstationsWorkstationConfigIamBindingId pulumi.StringOutput                                       `pulumi:"workstationsWorkstationConfigIamBindingId"`
}

// NewWorkstationsWorkstationConfigIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWorkstationsWorkstationConfigIamBinding(ctx *pulumi.Context,
	name string, args *WorkstationsWorkstationConfigIamBindingArgs, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationConfigIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
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
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource WorkstationsWorkstationConfigIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/workstationsWorkstationConfigIamBinding:WorkstationsWorkstationConfigIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationsWorkstationConfigIamBinding gets an existing WorkstationsWorkstationConfigIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationsWorkstationConfigIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationsWorkstationConfigIamBindingState, opts ...pulumi.ResourceOption) (*WorkstationsWorkstationConfigIamBinding, error) {
	var resource WorkstationsWorkstationConfigIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/workstationsWorkstationConfigIamBinding:WorkstationsWorkstationConfigIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationsWorkstationConfigIamBinding resources.
type workstationsWorkstationConfigIamBindingState struct {
	Condition                                 *WorkstationsWorkstationConfigIamBindingCondition `pulumi:"condition"`
	Etag                                      *string                                           `pulumi:"etag"`
	Location                                  *string                                           `pulumi:"location"`
	Members                                   []string                                          `pulumi:"members"`
	Project                                   *string                                           `pulumi:"project"`
	Role                                      *string                                           `pulumi:"role"`
	WorkstationClusterId                      *string                                           `pulumi:"workstationClusterId"`
	WorkstationConfigId                       *string                                           `pulumi:"workstationConfigId"`
	WorkstationsWorkstationConfigIamBindingId *string                                           `pulumi:"workstationsWorkstationConfigIamBindingId"`
}

type WorkstationsWorkstationConfigIamBindingState struct {
	Condition                                 WorkstationsWorkstationConfigIamBindingConditionPtrInput
	Etag                                      pulumi.StringPtrInput
	Location                                  pulumi.StringPtrInput
	Members                                   pulumi.StringArrayInput
	Project                                   pulumi.StringPtrInput
	Role                                      pulumi.StringPtrInput
	WorkstationClusterId                      pulumi.StringPtrInput
	WorkstationConfigId                       pulumi.StringPtrInput
	WorkstationsWorkstationConfigIamBindingId pulumi.StringPtrInput
}

func (WorkstationsWorkstationConfigIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationConfigIamBindingState)(nil)).Elem()
}

type workstationsWorkstationConfigIamBindingArgs struct {
	Condition                                 *WorkstationsWorkstationConfigIamBindingCondition `pulumi:"condition"`
	Location                                  *string                                           `pulumi:"location"`
	Members                                   []string                                          `pulumi:"members"`
	Project                                   *string                                           `pulumi:"project"`
	Role                                      string                                            `pulumi:"role"`
	WorkstationClusterId                      string                                            `pulumi:"workstationClusterId"`
	WorkstationConfigId                       string                                            `pulumi:"workstationConfigId"`
	WorkstationsWorkstationConfigIamBindingId *string                                           `pulumi:"workstationsWorkstationConfigIamBindingId"`
}

// The set of arguments for constructing a WorkstationsWorkstationConfigIamBinding resource.
type WorkstationsWorkstationConfigIamBindingArgs struct {
	Condition                                 WorkstationsWorkstationConfigIamBindingConditionPtrInput
	Location                                  pulumi.StringPtrInput
	Members                                   pulumi.StringArrayInput
	Project                                   pulumi.StringPtrInput
	Role                                      pulumi.StringInput
	WorkstationClusterId                      pulumi.StringInput
	WorkstationConfigId                       pulumi.StringInput
	WorkstationsWorkstationConfigIamBindingId pulumi.StringPtrInput
}

func (WorkstationsWorkstationConfigIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationsWorkstationConfigIamBindingArgs)(nil)).Elem()
}

type WorkstationsWorkstationConfigIamBindingInput interface {
	pulumi.Input

	ToWorkstationsWorkstationConfigIamBindingOutput() WorkstationsWorkstationConfigIamBindingOutput
	ToWorkstationsWorkstationConfigIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationConfigIamBindingOutput
}

func (*WorkstationsWorkstationConfigIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationConfigIamBinding)(nil)).Elem()
}

func (i *WorkstationsWorkstationConfigIamBinding) ToWorkstationsWorkstationConfigIamBindingOutput() WorkstationsWorkstationConfigIamBindingOutput {
	return i.ToWorkstationsWorkstationConfigIamBindingOutputWithContext(context.Background())
}

func (i *WorkstationsWorkstationConfigIamBinding) ToWorkstationsWorkstationConfigIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationConfigIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationsWorkstationConfigIamBindingOutput)
}

type WorkstationsWorkstationConfigIamBindingOutput struct{ *pulumi.OutputState }

func (WorkstationsWorkstationConfigIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationsWorkstationConfigIamBinding)(nil)).Elem()
}

func (o WorkstationsWorkstationConfigIamBindingOutput) ToWorkstationsWorkstationConfigIamBindingOutput() WorkstationsWorkstationConfigIamBindingOutput {
	return o
}

func (o WorkstationsWorkstationConfigIamBindingOutput) ToWorkstationsWorkstationConfigIamBindingOutputWithContext(ctx context.Context) WorkstationsWorkstationConfigIamBindingOutput {
	return o
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Condition() WorkstationsWorkstationConfigIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) WorkstationsWorkstationConfigIamBindingConditionPtrOutput {
		return v.Condition
	}).(WorkstationsWorkstationConfigIamBindingConditionPtrOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) WorkstationConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput { return v.WorkstationConfigId }).(pulumi.StringOutput)
}

func (o WorkstationsWorkstationConfigIamBindingOutput) WorkstationsWorkstationConfigIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationsWorkstationConfigIamBinding) pulumi.StringOutput {
		return v.WorkstationsWorkstationConfigIamBindingId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationsWorkstationConfigIamBindingInput)(nil)).Elem(), &WorkstationsWorkstationConfigIamBinding{})
	pulumi.RegisterOutputType(WorkstationsWorkstationConfigIamBindingOutput{})
}
