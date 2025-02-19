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

type ColabRuntimeTemplateIamBinding struct {
	pulumi.CustomResourceState

	ColabRuntimeTemplateIamBindingId pulumi.StringOutput                              `pulumi:"colabRuntimeTemplateIamBindingId"`
	Condition                        ColabRuntimeTemplateIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                             pulumi.StringOutput                              `pulumi:"etag"`
	Location                         pulumi.StringOutput                              `pulumi:"location"`
	Members                          pulumi.StringArrayOutput                         `pulumi:"members"`
	Project                          pulumi.StringOutput                              `pulumi:"project"`
	Role                             pulumi.StringOutput                              `pulumi:"role"`
	RuntimeTemplate                  pulumi.StringOutput                              `pulumi:"runtimeTemplate"`
}

// NewColabRuntimeTemplateIamBinding registers a new resource with the given unique name, arguments, and options.
func NewColabRuntimeTemplateIamBinding(ctx *pulumi.Context,
	name string, args *ColabRuntimeTemplateIamBindingArgs, opts ...pulumi.ResourceOption) (*ColabRuntimeTemplateIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.RuntimeTemplate == nil {
		return nil, errors.New("invalid value for required argument 'RuntimeTemplate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ColabRuntimeTemplateIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/colabRuntimeTemplateIamBinding:ColabRuntimeTemplateIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetColabRuntimeTemplateIamBinding gets an existing ColabRuntimeTemplateIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetColabRuntimeTemplateIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ColabRuntimeTemplateIamBindingState, opts ...pulumi.ResourceOption) (*ColabRuntimeTemplateIamBinding, error) {
	var resource ColabRuntimeTemplateIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/colabRuntimeTemplateIamBinding:ColabRuntimeTemplateIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ColabRuntimeTemplateIamBinding resources.
type colabRuntimeTemplateIamBindingState struct {
	ColabRuntimeTemplateIamBindingId *string                                  `pulumi:"colabRuntimeTemplateIamBindingId"`
	Condition                        *ColabRuntimeTemplateIamBindingCondition `pulumi:"condition"`
	Etag                             *string                                  `pulumi:"etag"`
	Location                         *string                                  `pulumi:"location"`
	Members                          []string                                 `pulumi:"members"`
	Project                          *string                                  `pulumi:"project"`
	Role                             *string                                  `pulumi:"role"`
	RuntimeTemplate                  *string                                  `pulumi:"runtimeTemplate"`
}

type ColabRuntimeTemplateIamBindingState struct {
	ColabRuntimeTemplateIamBindingId pulumi.StringPtrInput
	Condition                        ColabRuntimeTemplateIamBindingConditionPtrInput
	Etag                             pulumi.StringPtrInput
	Location                         pulumi.StringPtrInput
	Members                          pulumi.StringArrayInput
	Project                          pulumi.StringPtrInput
	Role                             pulumi.StringPtrInput
	RuntimeTemplate                  pulumi.StringPtrInput
}

func (ColabRuntimeTemplateIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*colabRuntimeTemplateIamBindingState)(nil)).Elem()
}

type colabRuntimeTemplateIamBindingArgs struct {
	ColabRuntimeTemplateIamBindingId *string                                  `pulumi:"colabRuntimeTemplateIamBindingId"`
	Condition                        *ColabRuntimeTemplateIamBindingCondition `pulumi:"condition"`
	Location                         *string                                  `pulumi:"location"`
	Members                          []string                                 `pulumi:"members"`
	Project                          *string                                  `pulumi:"project"`
	Role                             string                                   `pulumi:"role"`
	RuntimeTemplate                  string                                   `pulumi:"runtimeTemplate"`
}

// The set of arguments for constructing a ColabRuntimeTemplateIamBinding resource.
type ColabRuntimeTemplateIamBindingArgs struct {
	ColabRuntimeTemplateIamBindingId pulumi.StringPtrInput
	Condition                        ColabRuntimeTemplateIamBindingConditionPtrInput
	Location                         pulumi.StringPtrInput
	Members                          pulumi.StringArrayInput
	Project                          pulumi.StringPtrInput
	Role                             pulumi.StringInput
	RuntimeTemplate                  pulumi.StringInput
}

func (ColabRuntimeTemplateIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*colabRuntimeTemplateIamBindingArgs)(nil)).Elem()
}

type ColabRuntimeTemplateIamBindingInput interface {
	pulumi.Input

	ToColabRuntimeTemplateIamBindingOutput() ColabRuntimeTemplateIamBindingOutput
	ToColabRuntimeTemplateIamBindingOutputWithContext(ctx context.Context) ColabRuntimeTemplateIamBindingOutput
}

func (*ColabRuntimeTemplateIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**ColabRuntimeTemplateIamBinding)(nil)).Elem()
}

func (i *ColabRuntimeTemplateIamBinding) ToColabRuntimeTemplateIamBindingOutput() ColabRuntimeTemplateIamBindingOutput {
	return i.ToColabRuntimeTemplateIamBindingOutputWithContext(context.Background())
}

func (i *ColabRuntimeTemplateIamBinding) ToColabRuntimeTemplateIamBindingOutputWithContext(ctx context.Context) ColabRuntimeTemplateIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColabRuntimeTemplateIamBindingOutput)
}

type ColabRuntimeTemplateIamBindingOutput struct{ *pulumi.OutputState }

func (ColabRuntimeTemplateIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ColabRuntimeTemplateIamBinding)(nil)).Elem()
}

func (o ColabRuntimeTemplateIamBindingOutput) ToColabRuntimeTemplateIamBindingOutput() ColabRuntimeTemplateIamBindingOutput {
	return o
}

func (o ColabRuntimeTemplateIamBindingOutput) ToColabRuntimeTemplateIamBindingOutputWithContext(ctx context.Context) ColabRuntimeTemplateIamBindingOutput {
	return o
}

func (o ColabRuntimeTemplateIamBindingOutput) ColabRuntimeTemplateIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.ColabRuntimeTemplateIamBindingId }).(pulumi.StringOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Condition() ColabRuntimeTemplateIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) ColabRuntimeTemplateIamBindingConditionPtrOutput {
		return v.Condition
	}).(ColabRuntimeTemplateIamBindingConditionPtrOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o ColabRuntimeTemplateIamBindingOutput) RuntimeTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *ColabRuntimeTemplateIamBinding) pulumi.StringOutput { return v.RuntimeTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ColabRuntimeTemplateIamBindingInput)(nil)).Elem(), &ColabRuntimeTemplateIamBinding{})
	pulumi.RegisterOutputType(ColabRuntimeTemplateIamBindingOutput{})
}
