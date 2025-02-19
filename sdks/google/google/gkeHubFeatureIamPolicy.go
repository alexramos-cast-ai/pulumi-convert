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

type GkeHubFeatureIamPolicy struct {
	pulumi.CustomResourceState

	Etag                     pulumi.StringOutput `pulumi:"etag"`
	GkeHubFeatureIamPolicyId pulumi.StringOutput `pulumi:"gkeHubFeatureIamPolicyId"`
	Location                 pulumi.StringOutput `pulumi:"location"`
	Name                     pulumi.StringOutput `pulumi:"name"`
	PolicyData               pulumi.StringOutput `pulumi:"policyData"`
	Project                  pulumi.StringOutput `pulumi:"project"`
}

// NewGkeHubFeatureIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewGkeHubFeatureIamPolicy(ctx *pulumi.Context,
	name string, args *GkeHubFeatureIamPolicyArgs, opts ...pulumi.ResourceOption) (*GkeHubFeatureIamPolicy, error) {
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
	var resource GkeHubFeatureIamPolicy
	err = ctx.RegisterPackageResource("google:index/gkeHubFeatureIamPolicy:GkeHubFeatureIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubFeatureIamPolicy gets an existing GkeHubFeatureIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubFeatureIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubFeatureIamPolicyState, opts ...pulumi.ResourceOption) (*GkeHubFeatureIamPolicy, error) {
	var resource GkeHubFeatureIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/gkeHubFeatureIamPolicy:GkeHubFeatureIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubFeatureIamPolicy resources.
type gkeHubFeatureIamPolicyState struct {
	Etag                     *string `pulumi:"etag"`
	GkeHubFeatureIamPolicyId *string `pulumi:"gkeHubFeatureIamPolicyId"`
	Location                 *string `pulumi:"location"`
	Name                     *string `pulumi:"name"`
	PolicyData               *string `pulumi:"policyData"`
	Project                  *string `pulumi:"project"`
}

type GkeHubFeatureIamPolicyState struct {
	Etag                     pulumi.StringPtrInput
	GkeHubFeatureIamPolicyId pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	PolicyData               pulumi.StringPtrInput
	Project                  pulumi.StringPtrInput
}

func (GkeHubFeatureIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubFeatureIamPolicyState)(nil)).Elem()
}

type gkeHubFeatureIamPolicyArgs struct {
	GkeHubFeatureIamPolicyId *string `pulumi:"gkeHubFeatureIamPolicyId"`
	Location                 *string `pulumi:"location"`
	Name                     *string `pulumi:"name"`
	PolicyData               string  `pulumi:"policyData"`
	Project                  *string `pulumi:"project"`
}

// The set of arguments for constructing a GkeHubFeatureIamPolicy resource.
type GkeHubFeatureIamPolicyArgs struct {
	GkeHubFeatureIamPolicyId pulumi.StringPtrInput
	Location                 pulumi.StringPtrInput
	Name                     pulumi.StringPtrInput
	PolicyData               pulumi.StringInput
	Project                  pulumi.StringPtrInput
}

func (GkeHubFeatureIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubFeatureIamPolicyArgs)(nil)).Elem()
}

type GkeHubFeatureIamPolicyInput interface {
	pulumi.Input

	ToGkeHubFeatureIamPolicyOutput() GkeHubFeatureIamPolicyOutput
	ToGkeHubFeatureIamPolicyOutputWithContext(ctx context.Context) GkeHubFeatureIamPolicyOutput
}

func (*GkeHubFeatureIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubFeatureIamPolicy)(nil)).Elem()
}

func (i *GkeHubFeatureIamPolicy) ToGkeHubFeatureIamPolicyOutput() GkeHubFeatureIamPolicyOutput {
	return i.ToGkeHubFeatureIamPolicyOutputWithContext(context.Background())
}

func (i *GkeHubFeatureIamPolicy) ToGkeHubFeatureIamPolicyOutputWithContext(ctx context.Context) GkeHubFeatureIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubFeatureIamPolicyOutput)
}

type GkeHubFeatureIamPolicyOutput struct{ *pulumi.OutputState }

func (GkeHubFeatureIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubFeatureIamPolicy)(nil)).Elem()
}

func (o GkeHubFeatureIamPolicyOutput) ToGkeHubFeatureIamPolicyOutput() GkeHubFeatureIamPolicyOutput {
	return o
}

func (o GkeHubFeatureIamPolicyOutput) ToGkeHubFeatureIamPolicyOutputWithContext(ctx context.Context) GkeHubFeatureIamPolicyOutput {
	return o
}

func (o GkeHubFeatureIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o GkeHubFeatureIamPolicyOutput) GkeHubFeatureIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.GkeHubFeatureIamPolicyId }).(pulumi.StringOutput)
}

func (o GkeHubFeatureIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o GkeHubFeatureIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GkeHubFeatureIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o GkeHubFeatureIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubFeatureIamPolicyInput)(nil)).Elem(), &GkeHubFeatureIamPolicy{})
	pulumi.RegisterOutputType(GkeHubFeatureIamPolicyOutput{})
}
