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

type ClouddeployDeliveryPipelineIamPolicy struct {
	pulumi.CustomResourceState

	ClouddeployDeliveryPipelineIamPolicyId pulumi.StringOutput `pulumi:"clouddeployDeliveryPipelineIamPolicyId"`
	Etag                                   pulumi.StringOutput `pulumi:"etag"`
	Location                               pulumi.StringOutput `pulumi:"location"`
	Name                                   pulumi.StringOutput `pulumi:"name"`
	PolicyData                             pulumi.StringOutput `pulumi:"policyData"`
	Project                                pulumi.StringOutput `pulumi:"project"`
}

// NewClouddeployDeliveryPipelineIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewClouddeployDeliveryPipelineIamPolicy(ctx *pulumi.Context,
	name string, args *ClouddeployDeliveryPipelineIamPolicyArgs, opts ...pulumi.ResourceOption) (*ClouddeployDeliveryPipelineIamPolicy, error) {
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
	var resource ClouddeployDeliveryPipelineIamPolicy
	err = ctx.RegisterPackageResource("google:index/clouddeployDeliveryPipelineIamPolicy:ClouddeployDeliveryPipelineIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClouddeployDeliveryPipelineIamPolicy gets an existing ClouddeployDeliveryPipelineIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClouddeployDeliveryPipelineIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClouddeployDeliveryPipelineIamPolicyState, opts ...pulumi.ResourceOption) (*ClouddeployDeliveryPipelineIamPolicy, error) {
	var resource ClouddeployDeliveryPipelineIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/clouddeployDeliveryPipelineIamPolicy:ClouddeployDeliveryPipelineIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClouddeployDeliveryPipelineIamPolicy resources.
type clouddeployDeliveryPipelineIamPolicyState struct {
	ClouddeployDeliveryPipelineIamPolicyId *string `pulumi:"clouddeployDeliveryPipelineIamPolicyId"`
	Etag                                   *string `pulumi:"etag"`
	Location                               *string `pulumi:"location"`
	Name                                   *string `pulumi:"name"`
	PolicyData                             *string `pulumi:"policyData"`
	Project                                *string `pulumi:"project"`
}

type ClouddeployDeliveryPipelineIamPolicyState struct {
	ClouddeployDeliveryPipelineIamPolicyId pulumi.StringPtrInput
	Etag                                   pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	Name                                   pulumi.StringPtrInput
	PolicyData                             pulumi.StringPtrInput
	Project                                pulumi.StringPtrInput
}

func (ClouddeployDeliveryPipelineIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployDeliveryPipelineIamPolicyState)(nil)).Elem()
}

type clouddeployDeliveryPipelineIamPolicyArgs struct {
	ClouddeployDeliveryPipelineIamPolicyId *string `pulumi:"clouddeployDeliveryPipelineIamPolicyId"`
	Location                               *string `pulumi:"location"`
	Name                                   *string `pulumi:"name"`
	PolicyData                             string  `pulumi:"policyData"`
	Project                                *string `pulumi:"project"`
}

// The set of arguments for constructing a ClouddeployDeliveryPipelineIamPolicy resource.
type ClouddeployDeliveryPipelineIamPolicyArgs struct {
	ClouddeployDeliveryPipelineIamPolicyId pulumi.StringPtrInput
	Location                               pulumi.StringPtrInput
	Name                                   pulumi.StringPtrInput
	PolicyData                             pulumi.StringInput
	Project                                pulumi.StringPtrInput
}

func (ClouddeployDeliveryPipelineIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clouddeployDeliveryPipelineIamPolicyArgs)(nil)).Elem()
}

type ClouddeployDeliveryPipelineIamPolicyInput interface {
	pulumi.Input

	ToClouddeployDeliveryPipelineIamPolicyOutput() ClouddeployDeliveryPipelineIamPolicyOutput
	ToClouddeployDeliveryPipelineIamPolicyOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamPolicyOutput
}

func (*ClouddeployDeliveryPipelineIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployDeliveryPipelineIamPolicy)(nil)).Elem()
}

func (i *ClouddeployDeliveryPipelineIamPolicy) ToClouddeployDeliveryPipelineIamPolicyOutput() ClouddeployDeliveryPipelineIamPolicyOutput {
	return i.ToClouddeployDeliveryPipelineIamPolicyOutputWithContext(context.Background())
}

func (i *ClouddeployDeliveryPipelineIamPolicy) ToClouddeployDeliveryPipelineIamPolicyOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClouddeployDeliveryPipelineIamPolicyOutput)
}

type ClouddeployDeliveryPipelineIamPolicyOutput struct{ *pulumi.OutputState }

func (ClouddeployDeliveryPipelineIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClouddeployDeliveryPipelineIamPolicy)(nil)).Elem()
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) ToClouddeployDeliveryPipelineIamPolicyOutput() ClouddeployDeliveryPipelineIamPolicyOutput {
	return o
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) ToClouddeployDeliveryPipelineIamPolicyOutputWithContext(ctx context.Context) ClouddeployDeliveryPipelineIamPolicyOutput {
	return o
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) ClouddeployDeliveryPipelineIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput {
		return v.ClouddeployDeliveryPipelineIamPolicyId
	}).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o ClouddeployDeliveryPipelineIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ClouddeployDeliveryPipelineIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClouddeployDeliveryPipelineIamPolicyInput)(nil)).Elem(), &ClouddeployDeliveryPipelineIamPolicy{})
	pulumi.RegisterOutputType(ClouddeployDeliveryPipelineIamPolicyOutput{})
}
