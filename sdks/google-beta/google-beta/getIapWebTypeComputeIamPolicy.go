// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIapWebTypeComputeIamPolicy(ctx *pulumi.Context, args *LookupIapWebTypeComputeIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupIapWebTypeComputeIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIapWebTypeComputeIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getIapWebTypeComputeIamPolicy:getIapWebTypeComputeIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIapWebTypeComputeIamPolicy.
type LookupIapWebTypeComputeIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getIapWebTypeComputeIamPolicy.
type LookupIapWebTypeComputeIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupIapWebTypeComputeIamPolicyOutput(ctx *pulumi.Context, args LookupIapWebTypeComputeIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupIapWebTypeComputeIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIapWebTypeComputeIamPolicyResultOutput, error) {
			args := v.(LookupIapWebTypeComputeIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupIapWebTypeComputeIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getIapWebTypeComputeIamPolicy:getIapWebTypeComputeIamPolicy", args, LookupIapWebTypeComputeIamPolicyResultOutput{}, options).(LookupIapWebTypeComputeIamPolicyResultOutput), nil
		}).(LookupIapWebTypeComputeIamPolicyResultOutput)
}

// A collection of arguments for invoking getIapWebTypeComputeIamPolicy.
type LookupIapWebTypeComputeIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupIapWebTypeComputeIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapWebTypeComputeIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getIapWebTypeComputeIamPolicy.
type LookupIapWebTypeComputeIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupIapWebTypeComputeIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIapWebTypeComputeIamPolicyResult)(nil)).Elem()
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) ToLookupIapWebTypeComputeIamPolicyResultOutput() LookupIapWebTypeComputeIamPolicyResultOutput {
	return o
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) ToLookupIapWebTypeComputeIamPolicyResultOutputWithContext(ctx context.Context) LookupIapWebTypeComputeIamPolicyResultOutput {
	return o
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapWebTypeComputeIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapWebTypeComputeIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapWebTypeComputeIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupIapWebTypeComputeIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIapWebTypeComputeIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIapWebTypeComputeIamPolicyResultOutput{})
}
