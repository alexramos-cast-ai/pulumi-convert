// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataprocAutoscalingPolicyIamPolicy(ctx *pulumi.Context, args *LookupDataprocAutoscalingPolicyIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataprocAutoscalingPolicyIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupDataprocAutoscalingPolicyIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getDataprocAutoscalingPolicyIamPolicy:getDataprocAutoscalingPolicyIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataprocAutoscalingPolicyIamPolicy.
type LookupDataprocAutoscalingPolicyIamPolicyArgs struct {
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	PolicyId string  `pulumi:"policyId"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getDataprocAutoscalingPolicyIamPolicy.
type LookupDataprocAutoscalingPolicyIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Location   string `pulumi:"location"`
	PolicyData string `pulumi:"policyData"`
	PolicyId   string `pulumi:"policyId"`
	Project    string `pulumi:"project"`
}

func LookupDataprocAutoscalingPolicyIamPolicyOutput(ctx *pulumi.Context, args LookupDataprocAutoscalingPolicyIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDataprocAutoscalingPolicyIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataprocAutoscalingPolicyIamPolicyResultOutput, error) {
			args := v.(LookupDataprocAutoscalingPolicyIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupDataprocAutoscalingPolicyIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getDataprocAutoscalingPolicyIamPolicy:getDataprocAutoscalingPolicyIamPolicy", args, LookupDataprocAutoscalingPolicyIamPolicyResultOutput{}, options).(LookupDataprocAutoscalingPolicyIamPolicyResultOutput), nil
		}).(LookupDataprocAutoscalingPolicyIamPolicyResultOutput)
}

// A collection of arguments for invoking getDataprocAutoscalingPolicyIamPolicy.
type LookupDataprocAutoscalingPolicyIamPolicyOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	PolicyId pulumi.StringInput    `pulumi:"policyId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDataprocAutoscalingPolicyIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocAutoscalingPolicyIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getDataprocAutoscalingPolicyIamPolicy.
type LookupDataprocAutoscalingPolicyIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDataprocAutoscalingPolicyIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataprocAutoscalingPolicyIamPolicyResult)(nil)).Elem()
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) ToLookupDataprocAutoscalingPolicyIamPolicyResultOutput() LookupDataprocAutoscalingPolicyIamPolicyResultOutput {
	return o
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) ToLookupDataprocAutoscalingPolicyIamPolicyResultOutputWithContext(ctx context.Context) LookupDataprocAutoscalingPolicyIamPolicyResultOutput {
	return o
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.PolicyId }).(pulumi.StringOutput)
}

func (o LookupDataprocAutoscalingPolicyIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataprocAutoscalingPolicyIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataprocAutoscalingPolicyIamPolicyResultOutput{})
}
