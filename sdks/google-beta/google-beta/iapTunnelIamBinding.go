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

type IapTunnelIamBinding struct {
	pulumi.CustomResourceState

	Condition             IapTunnelIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                  pulumi.StringOutput                   `pulumi:"etag"`
	IapTunnelIamBindingId pulumi.StringOutput                   `pulumi:"iapTunnelIamBindingId"`
	Members               pulumi.StringArrayOutput              `pulumi:"members"`
	Project               pulumi.StringOutput                   `pulumi:"project"`
	Role                  pulumi.StringOutput                   `pulumi:"role"`
}

// NewIapTunnelIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIapTunnelIamBinding(ctx *pulumi.Context,
	name string, args *IapTunnelIamBindingArgs, opts ...pulumi.ResourceOption) (*IapTunnelIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IapTunnelIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/iapTunnelIamBinding:IapTunnelIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIapTunnelIamBinding gets an existing IapTunnelIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIapTunnelIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IapTunnelIamBindingState, opts ...pulumi.ResourceOption) (*IapTunnelIamBinding, error) {
	var resource IapTunnelIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/iapTunnelIamBinding:IapTunnelIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IapTunnelIamBinding resources.
type iapTunnelIamBindingState struct {
	Condition             *IapTunnelIamBindingCondition `pulumi:"condition"`
	Etag                  *string                       `pulumi:"etag"`
	IapTunnelIamBindingId *string                       `pulumi:"iapTunnelIamBindingId"`
	Members               []string                      `pulumi:"members"`
	Project               *string                       `pulumi:"project"`
	Role                  *string                       `pulumi:"role"`
}

type IapTunnelIamBindingState struct {
	Condition             IapTunnelIamBindingConditionPtrInput
	Etag                  pulumi.StringPtrInput
	IapTunnelIamBindingId pulumi.StringPtrInput
	Members               pulumi.StringArrayInput
	Project               pulumi.StringPtrInput
	Role                  pulumi.StringPtrInput
}

func (IapTunnelIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iapTunnelIamBindingState)(nil)).Elem()
}

type iapTunnelIamBindingArgs struct {
	Condition             *IapTunnelIamBindingCondition `pulumi:"condition"`
	IapTunnelIamBindingId *string                       `pulumi:"iapTunnelIamBindingId"`
	Members               []string                      `pulumi:"members"`
	Project               *string                       `pulumi:"project"`
	Role                  string                        `pulumi:"role"`
}

// The set of arguments for constructing a IapTunnelIamBinding resource.
type IapTunnelIamBindingArgs struct {
	Condition             IapTunnelIamBindingConditionPtrInput
	IapTunnelIamBindingId pulumi.StringPtrInput
	Members               pulumi.StringArrayInput
	Project               pulumi.StringPtrInput
	Role                  pulumi.StringInput
}

func (IapTunnelIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iapTunnelIamBindingArgs)(nil)).Elem()
}

type IapTunnelIamBindingInput interface {
	pulumi.Input

	ToIapTunnelIamBindingOutput() IapTunnelIamBindingOutput
	ToIapTunnelIamBindingOutputWithContext(ctx context.Context) IapTunnelIamBindingOutput
}

func (*IapTunnelIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**IapTunnelIamBinding)(nil)).Elem()
}

func (i *IapTunnelIamBinding) ToIapTunnelIamBindingOutput() IapTunnelIamBindingOutput {
	return i.ToIapTunnelIamBindingOutputWithContext(context.Background())
}

func (i *IapTunnelIamBinding) ToIapTunnelIamBindingOutputWithContext(ctx context.Context) IapTunnelIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IapTunnelIamBindingOutput)
}

type IapTunnelIamBindingOutput struct{ *pulumi.OutputState }

func (IapTunnelIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IapTunnelIamBinding)(nil)).Elem()
}

func (o IapTunnelIamBindingOutput) ToIapTunnelIamBindingOutput() IapTunnelIamBindingOutput {
	return o
}

func (o IapTunnelIamBindingOutput) ToIapTunnelIamBindingOutputWithContext(ctx context.Context) IapTunnelIamBindingOutput {
	return o
}

func (o IapTunnelIamBindingOutput) Condition() IapTunnelIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) IapTunnelIamBindingConditionPtrOutput { return v.Condition }).(IapTunnelIamBindingConditionPtrOutput)
}

func (o IapTunnelIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o IapTunnelIamBindingOutput) IapTunnelIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) pulumi.StringOutput { return v.IapTunnelIamBindingId }).(pulumi.StringOutput)
}

func (o IapTunnelIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o IapTunnelIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o IapTunnelIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *IapTunnelIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IapTunnelIamBindingInput)(nil)).Elem(), &IapTunnelIamBinding{})
	pulumi.RegisterOutputType(IapTunnelIamBindingOutput{})
}
