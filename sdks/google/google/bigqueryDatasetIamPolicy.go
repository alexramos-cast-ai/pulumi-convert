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

type BigqueryDatasetIamPolicy struct {
	pulumi.CustomResourceState

	BigqueryDatasetIamPolicyId pulumi.StringOutput `pulumi:"bigqueryDatasetIamPolicyId"`
	DatasetId                  pulumi.StringOutput `pulumi:"datasetId"`
	Etag                       pulumi.StringOutput `pulumi:"etag"`
	PolicyData                 pulumi.StringOutput `pulumi:"policyData"`
	Project                    pulumi.StringOutput `pulumi:"project"`
}

// NewBigqueryDatasetIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewBigqueryDatasetIamPolicy(ctx *pulumi.Context,
	name string, args *BigqueryDatasetIamPolicyArgs, opts ...pulumi.ResourceOption) (*BigqueryDatasetIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryDatasetIamPolicy
	err = ctx.RegisterPackageResource("google:index/bigqueryDatasetIamPolicy:BigqueryDatasetIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryDatasetIamPolicy gets an existing BigqueryDatasetIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryDatasetIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryDatasetIamPolicyState, opts ...pulumi.ResourceOption) (*BigqueryDatasetIamPolicy, error) {
	var resource BigqueryDatasetIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigqueryDatasetIamPolicy:BigqueryDatasetIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryDatasetIamPolicy resources.
type bigqueryDatasetIamPolicyState struct {
	BigqueryDatasetIamPolicyId *string `pulumi:"bigqueryDatasetIamPolicyId"`
	DatasetId                  *string `pulumi:"datasetId"`
	Etag                       *string `pulumi:"etag"`
	PolicyData                 *string `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
}

type BigqueryDatasetIamPolicyState struct {
	BigqueryDatasetIamPolicyId pulumi.StringPtrInput
	DatasetId                  pulumi.StringPtrInput
	Etag                       pulumi.StringPtrInput
	PolicyData                 pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
}

func (BigqueryDatasetIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatasetIamPolicyState)(nil)).Elem()
}

type bigqueryDatasetIamPolicyArgs struct {
	BigqueryDatasetIamPolicyId *string `pulumi:"bigqueryDatasetIamPolicyId"`
	DatasetId                  string  `pulumi:"datasetId"`
	PolicyData                 string  `pulumi:"policyData"`
	Project                    *string `pulumi:"project"`
}

// The set of arguments for constructing a BigqueryDatasetIamPolicy resource.
type BigqueryDatasetIamPolicyArgs struct {
	BigqueryDatasetIamPolicyId pulumi.StringPtrInput
	DatasetId                  pulumi.StringInput
	PolicyData                 pulumi.StringInput
	Project                    pulumi.StringPtrInput
}

func (BigqueryDatasetIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatasetIamPolicyArgs)(nil)).Elem()
}

type BigqueryDatasetIamPolicyInput interface {
	pulumi.Input

	ToBigqueryDatasetIamPolicyOutput() BigqueryDatasetIamPolicyOutput
	ToBigqueryDatasetIamPolicyOutputWithContext(ctx context.Context) BigqueryDatasetIamPolicyOutput
}

func (*BigqueryDatasetIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatasetIamPolicy)(nil)).Elem()
}

func (i *BigqueryDatasetIamPolicy) ToBigqueryDatasetIamPolicyOutput() BigqueryDatasetIamPolicyOutput {
	return i.ToBigqueryDatasetIamPolicyOutputWithContext(context.Background())
}

func (i *BigqueryDatasetIamPolicy) ToBigqueryDatasetIamPolicyOutputWithContext(ctx context.Context) BigqueryDatasetIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryDatasetIamPolicyOutput)
}

type BigqueryDatasetIamPolicyOutput struct{ *pulumi.OutputState }

func (BigqueryDatasetIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatasetIamPolicy)(nil)).Elem()
}

func (o BigqueryDatasetIamPolicyOutput) ToBigqueryDatasetIamPolicyOutput() BigqueryDatasetIamPolicyOutput {
	return o
}

func (o BigqueryDatasetIamPolicyOutput) ToBigqueryDatasetIamPolicyOutputWithContext(ctx context.Context) BigqueryDatasetIamPolicyOutput {
	return o
}

func (o BigqueryDatasetIamPolicyOutput) BigqueryDatasetIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetIamPolicy) pulumi.StringOutput { return v.BigqueryDatasetIamPolicyId }).(pulumi.StringOutput)
}

func (o BigqueryDatasetIamPolicyOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetIamPolicy) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

func (o BigqueryDatasetIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o BigqueryDatasetIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o BigqueryDatasetIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryDatasetIamPolicyInput)(nil)).Elem(), &BigqueryDatasetIamPolicy{})
	pulumi.RegisterOutputType(BigqueryDatasetIamPolicyOutput{})
}
