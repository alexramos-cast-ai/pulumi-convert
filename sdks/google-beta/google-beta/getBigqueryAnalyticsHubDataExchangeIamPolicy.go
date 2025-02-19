// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupBigqueryAnalyticsHubDataExchangeIamPolicy(ctx *pulumi.Context, args *LookupBigqueryAnalyticsHubDataExchangeIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getBigqueryAnalyticsHubDataExchangeIamPolicy:getBigqueryAnalyticsHubDataExchangeIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBigqueryAnalyticsHubDataExchangeIamPolicy.
type LookupBigqueryAnalyticsHubDataExchangeIamPolicyArgs struct {
	DataExchangeId string  `pulumi:"dataExchangeId"`
	Id             *string `pulumi:"id"`
	Location       *string `pulumi:"location"`
	Project        *string `pulumi:"project"`
}

// A collection of values returned by getBigqueryAnalyticsHubDataExchangeIamPolicy.
type LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult struct {
	DataExchangeId string `pulumi:"dataExchangeId"`
	Etag           string `pulumi:"etag"`
	Id             string `pulumi:"id"`
	Location       string `pulumi:"location"`
	PolicyData     string `pulumi:"policyData"`
	Project        string `pulumi:"project"`
}

func LookupBigqueryAnalyticsHubDataExchangeIamPolicyOutput(ctx *pulumi.Context, args LookupBigqueryAnalyticsHubDataExchangeIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput, error) {
			args := v.(LookupBigqueryAnalyticsHubDataExchangeIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getBigqueryAnalyticsHubDataExchangeIamPolicy:getBigqueryAnalyticsHubDataExchangeIamPolicy", args, LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput{}, options).(LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput), nil
		}).(LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput)
}

// A collection of arguments for invoking getBigqueryAnalyticsHubDataExchangeIamPolicy.
type LookupBigqueryAnalyticsHubDataExchangeIamPolicyOutputArgs struct {
	DataExchangeId pulumi.StringInput    `pulumi:"dataExchangeId"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	Location       pulumi.StringPtrInput `pulumi:"location"`
	Project        pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBigqueryAnalyticsHubDataExchangeIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigqueryAnalyticsHubDataExchangeIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getBigqueryAnalyticsHubDataExchangeIamPolicy.
type LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult)(nil)).Elem()
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) ToLookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput() LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput {
	return o
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) ToLookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutputWithContext(ctx context.Context) LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput {
	return o
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) DataExchangeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.DataExchangeId }).(pulumi.StringOutput)
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBigqueryAnalyticsHubDataExchangeIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBigqueryAnalyticsHubDataExchangeIamPolicyResultOutput{})
}
