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

type DataplexLakeIamPolicy struct {
	pulumi.CustomResourceState

	DataplexLakeIamPolicyId pulumi.StringOutput `pulumi:"dataplexLakeIamPolicyId"`
	Etag                    pulumi.StringOutput `pulumi:"etag"`
	Lake                    pulumi.StringOutput `pulumi:"lake"`
	Location                pulumi.StringOutput `pulumi:"location"`
	PolicyData              pulumi.StringOutput `pulumi:"policyData"`
	Project                 pulumi.StringOutput `pulumi:"project"`
}

// NewDataplexLakeIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewDataplexLakeIamPolicy(ctx *pulumi.Context,
	name string, args *DataplexLakeIamPolicyArgs, opts ...pulumi.ResourceOption) (*DataplexLakeIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexLakeIamPolicy
	err = ctx.RegisterPackageResource("google-beta:index/dataplexLakeIamPolicy:DataplexLakeIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexLakeIamPolicy gets an existing DataplexLakeIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexLakeIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexLakeIamPolicyState, opts ...pulumi.ResourceOption) (*DataplexLakeIamPolicy, error) {
	var resource DataplexLakeIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataplexLakeIamPolicy:DataplexLakeIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexLakeIamPolicy resources.
type dataplexLakeIamPolicyState struct {
	DataplexLakeIamPolicyId *string `pulumi:"dataplexLakeIamPolicyId"`
	Etag                    *string `pulumi:"etag"`
	Lake                    *string `pulumi:"lake"`
	Location                *string `pulumi:"location"`
	PolicyData              *string `pulumi:"policyData"`
	Project                 *string `pulumi:"project"`
}

type DataplexLakeIamPolicyState struct {
	DataplexLakeIamPolicyId pulumi.StringPtrInput
	Etag                    pulumi.StringPtrInput
	Lake                    pulumi.StringPtrInput
	Location                pulumi.StringPtrInput
	PolicyData              pulumi.StringPtrInput
	Project                 pulumi.StringPtrInput
}

func (DataplexLakeIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexLakeIamPolicyState)(nil)).Elem()
}

type dataplexLakeIamPolicyArgs struct {
	DataplexLakeIamPolicyId *string `pulumi:"dataplexLakeIamPolicyId"`
	Lake                    string  `pulumi:"lake"`
	Location                *string `pulumi:"location"`
	PolicyData              string  `pulumi:"policyData"`
	Project                 *string `pulumi:"project"`
}

// The set of arguments for constructing a DataplexLakeIamPolicy resource.
type DataplexLakeIamPolicyArgs struct {
	DataplexLakeIamPolicyId pulumi.StringPtrInput
	Lake                    pulumi.StringInput
	Location                pulumi.StringPtrInput
	PolicyData              pulumi.StringInput
	Project                 pulumi.StringPtrInput
}

func (DataplexLakeIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexLakeIamPolicyArgs)(nil)).Elem()
}

type DataplexLakeIamPolicyInput interface {
	pulumi.Input

	ToDataplexLakeIamPolicyOutput() DataplexLakeIamPolicyOutput
	ToDataplexLakeIamPolicyOutputWithContext(ctx context.Context) DataplexLakeIamPolicyOutput
}

func (*DataplexLakeIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexLakeIamPolicy)(nil)).Elem()
}

func (i *DataplexLakeIamPolicy) ToDataplexLakeIamPolicyOutput() DataplexLakeIamPolicyOutput {
	return i.ToDataplexLakeIamPolicyOutputWithContext(context.Background())
}

func (i *DataplexLakeIamPolicy) ToDataplexLakeIamPolicyOutputWithContext(ctx context.Context) DataplexLakeIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexLakeIamPolicyOutput)
}

type DataplexLakeIamPolicyOutput struct{ *pulumi.OutputState }

func (DataplexLakeIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexLakeIamPolicy)(nil)).Elem()
}

func (o DataplexLakeIamPolicyOutput) ToDataplexLakeIamPolicyOutput() DataplexLakeIamPolicyOutput {
	return o
}

func (o DataplexLakeIamPolicyOutput) ToDataplexLakeIamPolicyOutputWithContext(ctx context.Context) DataplexLakeIamPolicyOutput {
	return o
}

func (o DataplexLakeIamPolicyOutput) DataplexLakeIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.DataplexLakeIamPolicyId }).(pulumi.StringOutput)
}

func (o DataplexLakeIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexLakeIamPolicyOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

func (o DataplexLakeIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexLakeIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o DataplexLakeIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexLakeIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexLakeIamPolicyInput)(nil)).Elem(), &DataplexLakeIamPolicy{})
	pulumi.RegisterOutputType(DataplexLakeIamPolicyOutput{})
}
