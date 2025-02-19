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

type DataprocAutoscalingPolicyIamPolicy struct {
	pulumi.CustomResourceState

	DataprocAutoscalingPolicyIamPolicyId pulumi.StringOutput `pulumi:"dataprocAutoscalingPolicyIamPolicyId"`
	Etag                                 pulumi.StringOutput `pulumi:"etag"`
	Location                             pulumi.StringOutput `pulumi:"location"`
	PolicyData                           pulumi.StringOutput `pulumi:"policyData"`
	PolicyId                             pulumi.StringOutput `pulumi:"policyId"`
	Project                              pulumi.StringOutput `pulumi:"project"`
}

// NewDataprocAutoscalingPolicyIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataprocAutoscalingPolicyIamPolicy(ctx *pulumi.Context,
	name string, args *DataprocAutoscalingPolicyIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicyIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocAutoscalingPolicyIamPolicy
	err = ctx.RegisterPackageResource("google:index/dataprocAutoscalingPolicyIamPolicy:DataprocAutoscalingPolicyIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocAutoscalingPolicyIamPolicy gets an existing DataprocAutoscalingPolicyIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocAutoscalingPolicyIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocAutoscalingPolicyIamPolicyState, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicyIamPolicy, error) {
	var resource DataprocAutoscalingPolicyIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataprocAutoscalingPolicyIamPolicy:DataprocAutoscalingPolicyIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocAutoscalingPolicyIamPolicy resources.
type dataprocAutoscalingPolicyIamPolicyState struct {
	DataprocAutoscalingPolicyIamPolicyId *string `pulumi:"dataprocAutoscalingPolicyIamPolicyId"`
	Etag                                 *string `pulumi:"etag"`
	Location                             *string `pulumi:"location"`
	PolicyData                           *string `pulumi:"policyData"`
	PolicyId                             *string `pulumi:"policyId"`
	Project                              *string `pulumi:"project"`
}

type DataprocAutoscalingPolicyIamPolicyState struct {
	DataprocAutoscalingPolicyIamPolicyId pulumi.StringPtrInput
	Etag                                 pulumi.StringPtrInput
	Location                             pulumi.StringPtrInput
	PolicyData                           pulumi.StringPtrInput
	PolicyId                             pulumi.StringPtrInput
	Project                              pulumi.StringPtrInput
}

func (DataprocAutoscalingPolicyIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyIamPolicyState)(nil)).Elem()
}

type dataprocAutoscalingPolicyIamPolicyArgs struct {
	DataprocAutoscalingPolicyIamPolicyId *string `pulumi:"dataprocAutoscalingPolicyIamPolicyId"`
	Location                             *string `pulumi:"location"`
	PolicyData                           string  `pulumi:"policyData"`
	PolicyId                             string  `pulumi:"policyId"`
	Project                              *string `pulumi:"project"`
}

// The set of arguments for constructing a DataprocAutoscalingPolicyIamPolicy resource.
type DataprocAutoscalingPolicyIamPolicyArgs struct {
	DataprocAutoscalingPolicyIamPolicyId pulumi.StringPtrInput
	Location                             pulumi.StringPtrInput
	PolicyData                           pulumi.StringInput
	PolicyId                             pulumi.StringInput
	Project                              pulumi.StringPtrInput
}

func (DataprocAutoscalingPolicyIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyIamPolicyArgs)(nil)).Elem()
}

type DataprocAutoscalingPolicyIamPolicyInput interface {
	pulumi.Input

	ToDataprocAutoscalingPolicyIamPolicyOutput() DataprocAutoscalingPolicyIamPolicyOutput
	ToDataprocAutoscalingPolicyIamPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamPolicyOutput
}

func (*DataprocAutoscalingPolicyIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicyIamPolicy)(nil)).Elem()
}

func (i *DataprocAutoscalingPolicyIamPolicy) ToDataprocAutoscalingPolicyIamPolicyOutput() DataprocAutoscalingPolicyIamPolicyOutput {
	return i.ToDataprocAutoscalingPolicyIamPolicyOutputWithContext(context.Background())
}

func (i *DataprocAutoscalingPolicyIamPolicy) ToDataprocAutoscalingPolicyIamPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocAutoscalingPolicyIamPolicyOutput)
}

type DataprocAutoscalingPolicyIamPolicyOutput struct{ *pulumi.OutputState }

func (DataprocAutoscalingPolicyIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicyIamPolicy)(nil)).Elem()
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) ToDataprocAutoscalingPolicyIamPolicyOutput() DataprocAutoscalingPolicyIamPolicyOutput {
	return o
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) ToDataprocAutoscalingPolicyIamPolicyOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamPolicyOutput {
	return o
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) DataprocAutoscalingPolicyIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput {
		return v.DataprocAutoscalingPolicyIamPolicyId
	}).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocAutoscalingPolicyIamPolicyInput)(nil)).Elem(), &DataprocAutoscalingPolicyIamPolicy{})
	pulumi.RegisterOutputType(DataprocAutoscalingPolicyIamPolicyOutput{})
}
