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

type SpannerDatabaseIamPolicy struct {
	pulumi.CustomResourceState

	Database                   pulumi.StringOutput `pulumi:"database"`
	Etag                       pulumi.StringOutput `pulumi:"etag"`
	Instance                   pulumi.StringOutput `pulumi:"instance"`
	PolicyData                 pulumi.StringOutput `pulumi:"policyData"`
	Project                    pulumi.StringOutput `pulumi:"project"`
	SpannerDatabaseIamPolicyId pulumi.StringOutput `pulumi:"spannerDatabaseIamPolicyId"`
}

// NewSpannerDatabaseIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewSpannerDatabaseIamPolicy(ctx *pulumi.Context,
	name string, args *SpannerDatabaseIamPolicyArgs, opts ...pulumi.ResourceOption) (*SpannerDatabaseIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Database == nil {
		return nil, errors.New("invalid value for required argument 'Database'")
	}
	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SpannerDatabaseIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/spannerDatabaseIamPolicy:SpannerDatabaseIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpannerDatabaseIamPolicy gets an existing SpannerDatabaseIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpannerDatabaseIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpannerDatabaseIamPolicyState, opts ...pulumi.ResourceOption) (*SpannerDatabaseIamPolicy, error) {
	var resource SpannerDatabaseIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/spannerDatabaseIamPolicy:SpannerDatabaseIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpannerDatabaseIamPolicy resources.
type spannerDatabaseIamPolicyState struct {
	Database                   *string `pulumi:"database"`
	Etag                       *string `pulumi:"etag"`
	Instance                   *string `pulumi:"instance"`
	PolicyData                 *string `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
	SpannerDatabaseIamPolicyId *string `pulumi:"spannerDatabaseIamPolicyId"`
}

type SpannerDatabaseIamPolicyState struct {
	Database                   pulumi.StringPtrInput
	Etag                       pulumi.StringPtrInput
	Instance                   pulumi.StringPtrInput
	PolicyData                 pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	SpannerDatabaseIamPolicyId pulumi.StringPtrInput
}

func (SpannerDatabaseIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*spannerDatabaseIamPolicyState)(nil)).Elem()
}

type spannerDatabaseIamPolicyArgs struct {
	Database                   string  `pulumi:"database"`
	Instance                   string  `pulumi:"instance"`
	PolicyData                 string  `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
	SpannerDatabaseIamPolicyId *string `pulumi:"spannerDatabaseIamPolicyId"`
}

// The set of arguments for constructing a SpannerDatabaseIamPolicy resource.
type SpannerDatabaseIamPolicyArgs struct {
	Database                   pulumi.StringInput
	Instance                   pulumi.StringInput
	PolicyData                 pulumi.StringInput
	Project                    pulumi.StringPtrInput
	SpannerDatabaseIamPolicyId pulumi.StringPtrInput
}

func (SpannerDatabaseIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spannerDatabaseIamPolicyArgs)(nil)).Elem()
}

type SpannerDatabaseIamPolicyInput interface {
	pulumi.Input

	ToSpannerDatabaseIamPolicyOutput() SpannerDatabaseIamPolicyOutput
	ToSpannerDatabaseIamPolicyOutputWithContext(ctx context.Context) SpannerDatabaseIamPolicyOutput
}

func (*SpannerDatabaseIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**SpannerDatabaseIamPolicy)(nil)).Elem()
}

func (i *SpannerDatabaseIamPolicy) ToSpannerDatabaseIamPolicyOutput() SpannerDatabaseIamPolicyOutput {
	return i.ToSpannerDatabaseIamPolicyOutputWithContext(context.Background())
}

func (i *SpannerDatabaseIamPolicy) ToSpannerDatabaseIamPolicyOutputWithContext(ctx context.Context) SpannerDatabaseIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpannerDatabaseIamPolicyOutput)
}

type SpannerDatabaseIamPolicyOutput struct{ *pulumi.OutputState }

func (SpannerDatabaseIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpannerDatabaseIamPolicy)(nil)).Elem()
}

func (o SpannerDatabaseIamPolicyOutput) ToSpannerDatabaseIamPolicyOutput() SpannerDatabaseIamPolicyOutput {
	return o
}

func (o SpannerDatabaseIamPolicyOutput) ToSpannerDatabaseIamPolicyOutputWithContext(ctx context.Context) SpannerDatabaseIamPolicyOutput {
	return o
}

func (o SpannerDatabaseIamPolicyOutput) Database() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.Database }).(pulumi.StringOutput)
}

func (o SpannerDatabaseIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o SpannerDatabaseIamPolicyOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

func (o SpannerDatabaseIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o SpannerDatabaseIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SpannerDatabaseIamPolicyOutput) SpannerDatabaseIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabaseIamPolicy) pulumi.StringOutput { return v.SpannerDatabaseIamPolicyId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpannerDatabaseIamPolicyInput)(nil)).Elem(), &SpannerDatabaseIamPolicy{})
	pulumi.RegisterOutputType(SpannerDatabaseIamPolicyOutput{})
}
