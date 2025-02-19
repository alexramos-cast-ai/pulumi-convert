// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHealthcareConsentStoreIamPolicy(ctx *pulumi.Context, args *LookupHealthcareConsentStoreIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupHealthcareConsentStoreIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupHealthcareConsentStoreIamPolicyResult
	err = ctx.InvokePackage("google:index/getHealthcareConsentStoreIamPolicy:getHealthcareConsentStoreIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getHealthcareConsentStoreIamPolicy.
type LookupHealthcareConsentStoreIamPolicyArgs struct {
	ConsentStoreId string  `pulumi:"consentStoreId"`
	Dataset        string  `pulumi:"dataset"`
	Id             *string `pulumi:"id"`
}

// A collection of values returned by getHealthcareConsentStoreIamPolicy.
type LookupHealthcareConsentStoreIamPolicyResult struct {
	ConsentStoreId string `pulumi:"consentStoreId"`
	Dataset        string `pulumi:"dataset"`
	Etag           string `pulumi:"etag"`
	Id             string `pulumi:"id"`
	PolicyData     string `pulumi:"policyData"`
}

func LookupHealthcareConsentStoreIamPolicyOutput(ctx *pulumi.Context, args LookupHealthcareConsentStoreIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupHealthcareConsentStoreIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupHealthcareConsentStoreIamPolicyResultOutput, error) {
			args := v.(LookupHealthcareConsentStoreIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupHealthcareConsentStoreIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getHealthcareConsentStoreIamPolicy:getHealthcareConsentStoreIamPolicy", args, LookupHealthcareConsentStoreIamPolicyResultOutput{}, options).(LookupHealthcareConsentStoreIamPolicyResultOutput), nil
		}).(LookupHealthcareConsentStoreIamPolicyResultOutput)
}

// A collection of arguments for invoking getHealthcareConsentStoreIamPolicy.
type LookupHealthcareConsentStoreIamPolicyOutputArgs struct {
	ConsentStoreId pulumi.StringInput    `pulumi:"consentStoreId"`
	Dataset        pulumi.StringInput    `pulumi:"dataset"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupHealthcareConsentStoreIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthcareConsentStoreIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getHealthcareConsentStoreIamPolicy.
type LookupHealthcareConsentStoreIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupHealthcareConsentStoreIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupHealthcareConsentStoreIamPolicyResult)(nil)).Elem()
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) ToLookupHealthcareConsentStoreIamPolicyResultOutput() LookupHealthcareConsentStoreIamPolicyResultOutput {
	return o
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) ToLookupHealthcareConsentStoreIamPolicyResultOutputWithContext(ctx context.Context) LookupHealthcareConsentStoreIamPolicyResultOutput {
	return o
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) ConsentStoreId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthcareConsentStoreIamPolicyResult) string { return v.ConsentStoreId }).(pulumi.StringOutput)
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) Dataset() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthcareConsentStoreIamPolicyResult) string { return v.Dataset }).(pulumi.StringOutput)
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthcareConsentStoreIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthcareConsentStoreIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupHealthcareConsentStoreIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupHealthcareConsentStoreIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupHealthcareConsentStoreIamPolicyResultOutput{})
}
