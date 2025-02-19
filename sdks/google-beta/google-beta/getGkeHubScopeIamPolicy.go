// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGkeHubScopeIamPolicy(ctx *pulumi.Context, args *LookupGkeHubScopeIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupGkeHubScopeIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupGkeHubScopeIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getGkeHubScopeIamPolicy:getGkeHubScopeIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGkeHubScopeIamPolicy.
type LookupGkeHubScopeIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
	ScopeId string  `pulumi:"scopeId"`
}

// A collection of values returned by getGkeHubScopeIamPolicy.
type LookupGkeHubScopeIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
	ScopeId    string `pulumi:"scopeId"`
}

func LookupGkeHubScopeIamPolicyOutput(ctx *pulumi.Context, args LookupGkeHubScopeIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupGkeHubScopeIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupGkeHubScopeIamPolicyResultOutput, error) {
			args := v.(LookupGkeHubScopeIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupGkeHubScopeIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getGkeHubScopeIamPolicy:getGkeHubScopeIamPolicy", args, LookupGkeHubScopeIamPolicyResultOutput{}, options).(LookupGkeHubScopeIamPolicyResultOutput), nil
		}).(LookupGkeHubScopeIamPolicyResultOutput)
}

// A collection of arguments for invoking getGkeHubScopeIamPolicy.
type LookupGkeHubScopeIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	ScopeId pulumi.StringInput    `pulumi:"scopeId"`
}

func (LookupGkeHubScopeIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubScopeIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getGkeHubScopeIamPolicy.
type LookupGkeHubScopeIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupGkeHubScopeIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGkeHubScopeIamPolicyResult)(nil)).Elem()
}

func (o LookupGkeHubScopeIamPolicyResultOutput) ToLookupGkeHubScopeIamPolicyResultOutput() LookupGkeHubScopeIamPolicyResultOutput {
	return o
}

func (o LookupGkeHubScopeIamPolicyResultOutput) ToLookupGkeHubScopeIamPolicyResultOutputWithContext(ctx context.Context) LookupGkeHubScopeIamPolicyResultOutput {
	return o
}

func (o LookupGkeHubScopeIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubScopeIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupGkeHubScopeIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubScopeIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGkeHubScopeIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubScopeIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupGkeHubScopeIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubScopeIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupGkeHubScopeIamPolicyResultOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGkeHubScopeIamPolicyResult) string { return v.ScopeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGkeHubScopeIamPolicyResultOutput{})
}
