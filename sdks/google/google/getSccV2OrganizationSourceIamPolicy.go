// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSccV2OrganizationSourceIamPolicy(ctx *pulumi.Context, args *LookupSccV2OrganizationSourceIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupSccV2OrganizationSourceIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupSccV2OrganizationSourceIamPolicyResult
	err = ctx.InvokePackage("google:index/getSccV2OrganizationSourceIamPolicy:getSccV2OrganizationSourceIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSccV2OrganizationSourceIamPolicy.
type LookupSccV2OrganizationSourceIamPolicyArgs struct {
	Id           *string `pulumi:"id"`
	Organization string  `pulumi:"organization"`
	Source       string  `pulumi:"source"`
}

// A collection of values returned by getSccV2OrganizationSourceIamPolicy.
type LookupSccV2OrganizationSourceIamPolicyResult struct {
	Etag         string `pulumi:"etag"`
	Id           string `pulumi:"id"`
	Organization string `pulumi:"organization"`
	PolicyData   string `pulumi:"policyData"`
	Source       string `pulumi:"source"`
}

func LookupSccV2OrganizationSourceIamPolicyOutput(ctx *pulumi.Context, args LookupSccV2OrganizationSourceIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupSccV2OrganizationSourceIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupSccV2OrganizationSourceIamPolicyResultOutput, error) {
			args := v.(LookupSccV2OrganizationSourceIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupSccV2OrganizationSourceIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getSccV2OrganizationSourceIamPolicy:getSccV2OrganizationSourceIamPolicy", args, LookupSccV2OrganizationSourceIamPolicyResultOutput{}, options).(LookupSccV2OrganizationSourceIamPolicyResultOutput), nil
		}).(LookupSccV2OrganizationSourceIamPolicyResultOutput)
}

// A collection of arguments for invoking getSccV2OrganizationSourceIamPolicy.
type LookupSccV2OrganizationSourceIamPolicyOutputArgs struct {
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Organization pulumi.StringInput    `pulumi:"organization"`
	Source       pulumi.StringInput    `pulumi:"source"`
}

func (LookupSccV2OrganizationSourceIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSccV2OrganizationSourceIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getSccV2OrganizationSourceIamPolicy.
type LookupSccV2OrganizationSourceIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupSccV2OrganizationSourceIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSccV2OrganizationSourceIamPolicyResult)(nil)).Elem()
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) ToLookupSccV2OrganizationSourceIamPolicyResultOutput() LookupSccV2OrganizationSourceIamPolicyResultOutput {
	return o
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) ToLookupSccV2OrganizationSourceIamPolicyResultOutputWithContext(ctx context.Context) LookupSccV2OrganizationSourceIamPolicyResultOutput {
	return o
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSccV2OrganizationSourceIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSccV2OrganizationSourceIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSccV2OrganizationSourceIamPolicyResult) string { return v.Organization }).(pulumi.StringOutput)
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSccV2OrganizationSourceIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupSccV2OrganizationSourceIamPolicyResultOutput) Source() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSccV2OrganizationSourceIamPolicyResult) string { return v.Source }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSccV2OrganizationSourceIamPolicyResultOutput{})
}
