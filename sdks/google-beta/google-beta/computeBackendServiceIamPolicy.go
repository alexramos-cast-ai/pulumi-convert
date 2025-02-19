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

type ComputeBackendServiceIamPolicy struct {
	pulumi.CustomResourceState

	ComputeBackendServiceIamPolicyId pulumi.StringOutput `pulumi:"computeBackendServiceIamPolicyId"`
	Etag                             pulumi.StringOutput `pulumi:"etag"`
	Name                             pulumi.StringOutput `pulumi:"name"`
	PolicyData                       pulumi.StringOutput `pulumi:"policyData"`
	Project                          pulumi.StringOutput `pulumi:"project"`
}

// NewComputeBackendServiceIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewComputeBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, args *ComputeBackendServiceIamPolicyArgs, opts ...pulumi.ResourceOption) (*ComputeBackendServiceIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeBackendServiceIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/computeBackendServiceIamPolicy:ComputeBackendServiceIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeBackendServiceIamPolicy gets an existing ComputeBackendServiceIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeBackendServiceIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeBackendServiceIamPolicyState, opts ...pulumi.ResourceOption) (*ComputeBackendServiceIamPolicy, error) {
	var resource ComputeBackendServiceIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeBackendServiceIamPolicy:ComputeBackendServiceIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeBackendServiceIamPolicy resources.
type computeBackendServiceIamPolicyState struct {
	ComputeBackendServiceIamPolicyId *string `pulumi:"computeBackendServiceIamPolicyId"`
	Etag                             *string `pulumi:"etag"`
	Name                             *string `pulumi:"name"`
	PolicyData                       *string `pulumi:"policyData"`
	Project                          *string `pulumi:"project"`
}

type ComputeBackendServiceIamPolicyState struct {
	ComputeBackendServiceIamPolicyId pulumi.StringPtrInput
	Etag                             pulumi.StringPtrInput
	Name                             pulumi.StringPtrInput
	PolicyData                       pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
}

func (ComputeBackendServiceIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeBackendServiceIamPolicyState)(nil)).Elem()
}

type computeBackendServiceIamPolicyArgs struct {
	ComputeBackendServiceIamPolicyId *string `pulumi:"computeBackendServiceIamPolicyId"`
	Name                             *string `pulumi:"name"`
	PolicyData                       string  `pulumi:"policyData"`
	Project                          *string `pulumi:"project"`
}

// The set of arguments for constructing a ComputeBackendServiceIamPolicy resource.
type ComputeBackendServiceIamPolicyArgs struct {
	ComputeBackendServiceIamPolicyId pulumi.StringPtrInput
	Name                             pulumi.StringPtrInput
	PolicyData                       pulumi.StringInput
	Project                          pulumi.StringPtrInput
}

func (ComputeBackendServiceIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeBackendServiceIamPolicyArgs)(nil)).Elem()
}

type ComputeBackendServiceIamPolicyInput interface {
	pulumi.Input

	ToComputeBackendServiceIamPolicyOutput() ComputeBackendServiceIamPolicyOutput
	ToComputeBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeBackendServiceIamPolicyOutput
}

func (*ComputeBackendServiceIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBackendServiceIamPolicy)(nil)).Elem()
}

func (i *ComputeBackendServiceIamPolicy) ToComputeBackendServiceIamPolicyOutput() ComputeBackendServiceIamPolicyOutput {
	return i.ToComputeBackendServiceIamPolicyOutputWithContext(context.Background())
}

func (i *ComputeBackendServiceIamPolicy) ToComputeBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeBackendServiceIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeBackendServiceIamPolicyOutput)
}

type ComputeBackendServiceIamPolicyOutput struct{ *pulumi.OutputState }

func (ComputeBackendServiceIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBackendServiceIamPolicy)(nil)).Elem()
}

func (o ComputeBackendServiceIamPolicyOutput) ToComputeBackendServiceIamPolicyOutput() ComputeBackendServiceIamPolicyOutput {
	return o
}

func (o ComputeBackendServiceIamPolicyOutput) ToComputeBackendServiceIamPolicyOutputWithContext(ctx context.Context) ComputeBackendServiceIamPolicyOutput {
	return o
}

func (o ComputeBackendServiceIamPolicyOutput) ComputeBackendServiceIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendServiceIamPolicy) pulumi.StringOutput { return v.ComputeBackendServiceIamPolicyId }).(pulumi.StringOutput)
}

func (o ComputeBackendServiceIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendServiceIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ComputeBackendServiceIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendServiceIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeBackendServiceIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendServiceIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o ComputeBackendServiceIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendServiceIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeBackendServiceIamPolicyInput)(nil)).Elem(), &ComputeBackendServiceIamPolicy{})
	pulumi.RegisterOutputType(ComputeBackendServiceIamPolicyOutput{})
}
