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

type KmsKeyRingIamPolicy struct {
	pulumi.CustomResourceState

	Etag                  pulumi.StringOutput `pulumi:"etag"`
	KeyRingId             pulumi.StringOutput `pulumi:"keyRingId"`
	KmsKeyRingIamPolicyId pulumi.StringOutput `pulumi:"kmsKeyRingIamPolicyId"`
	PolicyData            pulumi.StringOutput `pulumi:"policyData"`
}

// NewKmsKeyRingIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewKmsKeyRingIamPolicy(ctx *pulumi.Context,
	name string, args *KmsKeyRingIamPolicyArgs, opts ...pulumi.ResourceOption) (*KmsKeyRingIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyRingId == nil {
		return nil, errors.New("invalid value for required argument 'KeyRingId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource KmsKeyRingIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/kmsKeyRingIamPolicy:KmsKeyRingIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsKeyRingIamPolicy gets an existing KmsKeyRingIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsKeyRingIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsKeyRingIamPolicyState, opts ...pulumi.ResourceOption) (*KmsKeyRingIamPolicy, error) {
	var resource KmsKeyRingIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/kmsKeyRingIamPolicy:KmsKeyRingIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsKeyRingIamPolicy resources.
type kmsKeyRingIamPolicyState struct {
	Etag                  *string `pulumi:"etag"`
	KeyRingId             *string `pulumi:"keyRingId"`
	KmsKeyRingIamPolicyId *string `pulumi:"kmsKeyRingIamPolicyId"`
	PolicyData            *string `pulumi:"policyData"`
}

type KmsKeyRingIamPolicyState struct {
	Etag                  pulumi.StringPtrInput
	KeyRingId             pulumi.StringPtrInput
	KmsKeyRingIamPolicyId pulumi.StringPtrInput
	PolicyData            pulumi.StringPtrInput
}

func (KmsKeyRingIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsKeyRingIamPolicyState)(nil)).Elem()
}

type kmsKeyRingIamPolicyArgs struct {
	KeyRingId             string  `pulumi:"keyRingId"`
	KmsKeyRingIamPolicyId *string `pulumi:"kmsKeyRingIamPolicyId"`
	PolicyData            string  `pulumi:"policyData"`
}

// The set of arguments for constructing a KmsKeyRingIamPolicy resource.
type KmsKeyRingIamPolicyArgs struct {
	KeyRingId             pulumi.StringInput
	KmsKeyRingIamPolicyId pulumi.StringPtrInput
	PolicyData            pulumi.StringInput
}

func (KmsKeyRingIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsKeyRingIamPolicyArgs)(nil)).Elem()
}

type KmsKeyRingIamPolicyInput interface {
	pulumi.Input

	ToKmsKeyRingIamPolicyOutput() KmsKeyRingIamPolicyOutput
	ToKmsKeyRingIamPolicyOutputWithContext(ctx context.Context) KmsKeyRingIamPolicyOutput
}

func (*KmsKeyRingIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsKeyRingIamPolicy)(nil)).Elem()
}

func (i *KmsKeyRingIamPolicy) ToKmsKeyRingIamPolicyOutput() KmsKeyRingIamPolicyOutput {
	return i.ToKmsKeyRingIamPolicyOutputWithContext(context.Background())
}

func (i *KmsKeyRingIamPolicy) ToKmsKeyRingIamPolicyOutputWithContext(ctx context.Context) KmsKeyRingIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsKeyRingIamPolicyOutput)
}

type KmsKeyRingIamPolicyOutput struct{ *pulumi.OutputState }

func (KmsKeyRingIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsKeyRingIamPolicy)(nil)).Elem()
}

func (o KmsKeyRingIamPolicyOutput) ToKmsKeyRingIamPolicyOutput() KmsKeyRingIamPolicyOutput {
	return o
}

func (o KmsKeyRingIamPolicyOutput) ToKmsKeyRingIamPolicyOutputWithContext(ctx context.Context) KmsKeyRingIamPolicyOutput {
	return o
}

func (o KmsKeyRingIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o KmsKeyRingIamPolicyOutput) KeyRingId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingIamPolicy) pulumi.StringOutput { return v.KeyRingId }).(pulumi.StringOutput)
}

func (o KmsKeyRingIamPolicyOutput) KmsKeyRingIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingIamPolicy) pulumi.StringOutput { return v.KmsKeyRingIamPolicyId }).(pulumi.StringOutput)
}

func (o KmsKeyRingIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsKeyRingIamPolicyInput)(nil)).Elem(), &KmsKeyRingIamPolicy{})
	pulumi.RegisterOutputType(KmsKeyRingIamPolicyOutput{})
}
