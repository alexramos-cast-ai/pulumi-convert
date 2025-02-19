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

type LoggingLogViewIamPolicy struct {
	pulumi.CustomResourceState

	Bucket                    pulumi.StringOutput `pulumi:"bucket"`
	Etag                      pulumi.StringOutput `pulumi:"etag"`
	Location                  pulumi.StringOutput `pulumi:"location"`
	LoggingLogViewIamPolicyId pulumi.StringOutput `pulumi:"loggingLogViewIamPolicyId"`
	Name                      pulumi.StringOutput `pulumi:"name"`
	Parent                    pulumi.StringOutput `pulumi:"parent"`
	PolicyData                pulumi.StringOutput `pulumi:"policyData"`
}

// NewLoggingLogViewIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewLoggingLogViewIamPolicy(ctx *pulumi.Context,
	name string, args *LoggingLogViewIamPolicyArgs, opts ...pulumi.ResourceOption) (*LoggingLogViewIamPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.PolicyData == nil {
		return nil, errors.New("invalid value for required argument 'PolicyData'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingLogViewIamPolicy
	err = ctx.RegisterPackageResource("google:index/loggingLogViewIamPolicy:LoggingLogViewIamPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingLogViewIamPolicy gets an existing LoggingLogViewIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingLogViewIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingLogViewIamPolicyState, opts ...pulumi.ResourceOption) (*LoggingLogViewIamPolicy, error) {
	var resource LoggingLogViewIamPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/loggingLogViewIamPolicy:LoggingLogViewIamPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingLogViewIamPolicy resources.
type loggingLogViewIamPolicyState struct {
	Bucket                    *string `pulumi:"bucket"`
	Etag                      *string `pulumi:"etag"`
	Location                  *string `pulumi:"location"`
	LoggingLogViewIamPolicyId *string `pulumi:"loggingLogViewIamPolicyId"`
	Name                      *string `pulumi:"name"`
	Parent                    *string `pulumi:"parent"`
	PolicyData                *string `pulumi:"policyData"`
}

type LoggingLogViewIamPolicyState struct {
	Bucket                    pulumi.StringPtrInput
	Etag                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	LoggingLogViewIamPolicyId pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Parent                    pulumi.StringPtrInput
	PolicyData                pulumi.StringPtrInput
}

func (LoggingLogViewIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingLogViewIamPolicyState)(nil)).Elem()
}

type loggingLogViewIamPolicyArgs struct {
	Bucket                    string  `pulumi:"bucket"`
	Location                  *string `pulumi:"location"`
	LoggingLogViewIamPolicyId *string `pulumi:"loggingLogViewIamPolicyId"`
	Name                      *string `pulumi:"name"`
	Parent                    string  `pulumi:"parent"`
	PolicyData                string  `pulumi:"policyData"`
}

// The set of arguments for constructing a LoggingLogViewIamPolicy resource.
type LoggingLogViewIamPolicyArgs struct {
	Bucket                    pulumi.StringInput
	Location                  pulumi.StringPtrInput
	LoggingLogViewIamPolicyId pulumi.StringPtrInput
	Name                      pulumi.StringPtrInput
	Parent                    pulumi.StringInput
	PolicyData                pulumi.StringInput
}

func (LoggingLogViewIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingLogViewIamPolicyArgs)(nil)).Elem()
}

type LoggingLogViewIamPolicyInput interface {
	pulumi.Input

	ToLoggingLogViewIamPolicyOutput() LoggingLogViewIamPolicyOutput
	ToLoggingLogViewIamPolicyOutputWithContext(ctx context.Context) LoggingLogViewIamPolicyOutput
}

func (*LoggingLogViewIamPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingLogViewIamPolicy)(nil)).Elem()
}

func (i *LoggingLogViewIamPolicy) ToLoggingLogViewIamPolicyOutput() LoggingLogViewIamPolicyOutput {
	return i.ToLoggingLogViewIamPolicyOutputWithContext(context.Background())
}

func (i *LoggingLogViewIamPolicy) ToLoggingLogViewIamPolicyOutputWithContext(ctx context.Context) LoggingLogViewIamPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingLogViewIamPolicyOutput)
}

type LoggingLogViewIamPolicyOutput struct{ *pulumi.OutputState }

func (LoggingLogViewIamPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingLogViewIamPolicy)(nil)).Elem()
}

func (o LoggingLogViewIamPolicyOutput) ToLoggingLogViewIamPolicyOutput() LoggingLogViewIamPolicyOutput {
	return o
}

func (o LoggingLogViewIamPolicyOutput) ToLoggingLogViewIamPolicyOutputWithContext(ctx context.Context) LoggingLogViewIamPolicyOutput {
	return o
}

func (o LoggingLogViewIamPolicyOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) LoggingLogViewIamPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.LoggingLogViewIamPolicyId }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

func (o LoggingLogViewIamPolicyOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingLogViewIamPolicy) pulumi.StringOutput { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingLogViewIamPolicyInput)(nil)).Elem(), &LoggingLogViewIamPolicy{})
	pulumi.RegisterOutputType(LoggingLogViewIamPolicyOutput{})
}
