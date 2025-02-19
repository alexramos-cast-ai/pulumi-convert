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

type RuntimeconfigConfigIamPolicy struct {
	pulumi.CustomResourceState

	Config                         pulumi.StringOutput `pulumi:"config"`
	Etag                           pulumi.StringOutput `pulumi:"etag"`
	PolicyData                     pulumi.StringOutput `pulumi:"policyData"`
	Project                        pulumi.StringOutput `pulumi:"project"`
	RuntimeconfigConfigIamPolicyId pulumi.StringOutput `pulumi:"runtimeconfigConfigIamPolicyId"`
}

// NewRuntimeconfigConfigIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewRuntimeconfigConfigIamPolicy(ctx *pulumi.Context,
	name string, args *RuntimeconfigConfigIamPolicyArgs, opts ...pulumi.ResourceOption) (*RuntimeconfigConfigIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RuntimeconfigConfigIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/runtimeconfigConfigIamPolicy:RuntimeconfigConfigIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuntimeconfigConfigIamPolicy gets an existing RuntimeconfigConfigIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuntimeconfigConfigIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuntimeconfigConfigIamPolicyState, opts ...pulumi.ResourceOption) (*RuntimeconfigConfigIamPolicy, error) {
	var resource RuntimeconfigConfigIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/runtimeconfigConfigIamPolicy:RuntimeconfigConfigIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuntimeconfigConfigIamPolicy resources.
type runtimeconfigConfigIamPolicyState struct {
	Config                         *string `pulumi:"config"`
	Etag                           *string `pulumi:"etag"`
	PolicyData                     *string `pulumi:"policyData"`
	Project                        *string `pulumi:"project"`
	RuntimeconfigConfigIamPolicyId *string `pulumi:"runtimeconfigConfigIamPolicyId"`
}

type RuntimeconfigConfigIamPolicyState struct {
	Config                         pulumi.StringPtrInput
	Etag                           pulumi.StringPtrInput
	PolicyData                     pulumi.StringPtrInput
	Project                        pulumi.StringPtrInput
	RuntimeconfigConfigIamPolicyId pulumi.StringPtrInput
}

func (RuntimeconfigConfigIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeconfigConfigIamPolicyState)(nil)).Elem()
}

type runtimeconfigConfigIamPolicyArgs struct {
	Config                         string  `pulumi:"config"`
	PolicyData                     string  `pulumi:"policyData"`
	Project                        *string `pulumi:"project"`
	RuntimeconfigConfigIamPolicyId *string `pulumi:"runtimeconfigConfigIamPolicyId"`
}

// The set of arguments for constructing a RuntimeconfigConfigIamPolicy resource.
type RuntimeconfigConfigIamPolicyArgs struct {
	Config                         pulumi.StringInput
	PolicyData                     pulumi.StringInput
	Project                        pulumi.StringPtrInput
	RuntimeconfigConfigIamPolicyId pulumi.StringPtrInput
}

func (RuntimeconfigConfigIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*runtimeconfigConfigIamPolicyArgs)(nil)).Elem()
}

type RuntimeconfigConfigIamPolicyInput interface {
	pulumi.Input

	ToRuntimeconfigConfigIamPolicyOutput() RuntimeconfigConfigIamPolicyOutput
	ToRuntimeconfigConfigIamPolicyOutputWithContext(ctx context.Context) RuntimeconfigConfigIamPolicyOutput
}

func (*RuntimeconfigConfigIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeconfigConfigIamPolicy)(nil)).Elem()
}

func (i *RuntimeconfigConfigIamPolicy) ToRuntimeconfigConfigIamPolicyOutput() RuntimeconfigConfigIamPolicyOutput {
	return i.ToRuntimeconfigConfigIamPolicyOutputWithContext(context.Background())
}

func (i *RuntimeconfigConfigIamPolicy) ToRuntimeconfigConfigIamPolicyOutputWithContext(ctx context.Context) RuntimeconfigConfigIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuntimeconfigConfigIamPolicyOutput)
}

type RuntimeconfigConfigIamPolicyOutput struct{ *pulumi.OutputState }

func (RuntimeconfigConfigIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuntimeconfigConfigIamPolicy)(nil)).Elem()
}

func (o RuntimeconfigConfigIamPolicyOutput) ToRuntimeconfigConfigIamPolicyOutput() RuntimeconfigConfigIamPolicyOutput {
	return o
}

func (o RuntimeconfigConfigIamPolicyOutput) ToRuntimeconfigConfigIamPolicyOutputWithContext(ctx context.Context) RuntimeconfigConfigIamPolicyOutput {
	return o
}

func (o RuntimeconfigConfigIamPolicyOutput) Config() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfigIamPolicy) pulumi.StringOutput { return v.Config }).(pulumi.StringOutput)
}

func (o RuntimeconfigConfigIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfigIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o RuntimeconfigConfigIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfigIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o RuntimeconfigConfigIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfigIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o RuntimeconfigConfigIamPolicyOutput) RuntimeconfigConfigIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuntimeconfigConfigIamPolicy) pulumi.StringOutput { return v.RuntimeconfigConfigIamPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuntimeconfigConfigIamPolicyInput)(nil)).Elem(), &RuntimeconfigConfigIamPolicy{})
	pulumi.RegisterOutputType(RuntimeconfigConfigIamPolicyOutput{})
}
