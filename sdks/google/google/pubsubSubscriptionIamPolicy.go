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

type PubsubSubscriptionIamPolicy struct {
	pulumi.CustomResourceState

	Etag                          pulumi.StringOutput `pulumi:"etag"`
	PolicyData                    pulumi.StringOutput `pulumi:"policyData"`
	Project                       pulumi.StringOutput `pulumi:"project"`
	PubsubSubscriptionIamPolicyId pulumi.StringOutput `pulumi:"pubsubSubscriptionIamPolicyId"`
	Subscription                  pulumi.StringOutput `pulumi:"subscription"`
}

// NewPubsubSubscriptionIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewPubsubSubscriptionIamPolicy(ctx *pulumi.Context,
	name string, args *PubsubSubscriptionIamPolicyArgs, opts ...pulumi.ResourceOption) (*PubsubSubscriptionIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	if args.Subscription == nil {
		return nil, errors.New("invalid value for required argument 'Subscription'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource PubsubSubscriptionIamPolicy
	err = ctx.RegisterPackageResource("google:index/pubsubSubscriptionIamPolicy:PubsubSubscriptionIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPubsubSubscriptionIamPolicy gets an existing PubsubSubscriptionIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPubsubSubscriptionIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PubsubSubscriptionIamPolicyState, opts ...pulumi.ResourceOption) (*PubsubSubscriptionIamPolicy, error) {
	var resource PubsubSubscriptionIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/pubsubSubscriptionIamPolicy:PubsubSubscriptionIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PubsubSubscriptionIamPolicy resources.
type pubsubSubscriptionIamPolicyState struct {
	Etag                          *string `pulumi:"etag"`
	PolicyData                    *string `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	PubsubSubscriptionIamPolicyId *string `pulumi:"pubsubSubscriptionIamPolicyId"`
	Subscription                  *string `pulumi:"subscription"`
}

type PubsubSubscriptionIamPolicyState struct {
	Etag                          pulumi.StringPtrInput
	PolicyData                    pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	PubsubSubscriptionIamPolicyId pulumi.StringPtrInput
	Subscription                  pulumi.StringPtrInput
}

func (PubsubSubscriptionIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSubscriptionIamPolicyState)(nil)).Elem()
}

type pubsubSubscriptionIamPolicyArgs struct {
	PolicyData                    string  `pulumi:"policyData"`
	Project                       *string `pulumi:"project"`
	PubsubSubscriptionIamPolicyId *string `pulumi:"pubsubSubscriptionIamPolicyId"`
	Subscription                  string  `pulumi:"subscription"`
}

// The set of arguments for constructing a PubsubSubscriptionIamPolicy resource.
type PubsubSubscriptionIamPolicyArgs struct {
	PolicyData                    pulumi.StringInput
	Project                       pulumi.StringPtrInput
	PubsubSubscriptionIamPolicyId pulumi.StringPtrInput
	Subscription                  pulumi.StringInput
}

func (PubsubSubscriptionIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pubsubSubscriptionIamPolicyArgs)(nil)).Elem()
}

type PubsubSubscriptionIamPolicyInput interface {
	pulumi.Input

	ToPubsubSubscriptionIamPolicyOutput() PubsubSubscriptionIamPolicyOutput
	ToPubsubSubscriptionIamPolicyOutputWithContext(ctx context.Context) PubsubSubscriptionIamPolicyOutput
}

func (*PubsubSubscriptionIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSubscriptionIamPolicy)(nil)).Elem()
}

func (i *PubsubSubscriptionIamPolicy) ToPubsubSubscriptionIamPolicyOutput() PubsubSubscriptionIamPolicyOutput {
	return i.ToPubsubSubscriptionIamPolicyOutputWithContext(context.Background())
}

func (i *PubsubSubscriptionIamPolicy) ToPubsubSubscriptionIamPolicyOutputWithContext(ctx context.Context) PubsubSubscriptionIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PubsubSubscriptionIamPolicyOutput)
}

type PubsubSubscriptionIamPolicyOutput struct{ *pulumi.OutputState }

func (PubsubSubscriptionIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PubsubSubscriptionIamPolicy)(nil)).Elem()
}

func (o PubsubSubscriptionIamPolicyOutput) ToPubsubSubscriptionIamPolicyOutput() PubsubSubscriptionIamPolicyOutput {
	return o
}

func (o PubsubSubscriptionIamPolicyOutput) ToPubsubSubscriptionIamPolicyOutputWithContext(ctx context.Context) PubsubSubscriptionIamPolicyOutput {
	return o
}

func (o PubsubSubscriptionIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamPolicyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamPolicy) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamPolicyOutput) PubsubSubscriptionIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamPolicy) pulumi.StringOutput { return v.PubsubSubscriptionIamPolicyId }).(pulumi.StringOutput)
}

func (o PubsubSubscriptionIamPolicyOutput) Subscription() pulumi.StringOutput {
	return o.ApplyT(func(v *PubsubSubscriptionIamPolicy) pulumi.StringOutput { return v.Subscription }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PubsubSubscriptionIamPolicyInput)(nil)).Elem(), &PubsubSubscriptionIamPolicy{})
	pulumi.RegisterOutputType(PubsubSubscriptionIamPolicyOutput{})
}
