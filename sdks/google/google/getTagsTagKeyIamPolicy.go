// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagsTagKeyIamPolicy(ctx *pulumi.Context, args *LookupTagsTagKeyIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupTagsTagKeyIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupTagsTagKeyIamPolicyResult
	err = ctx.InvokePackage("google:index/getTagsTagKeyIamPolicy:getTagsTagKeyIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagsTagKeyIamPolicy.
type LookupTagsTagKeyIamPolicyArgs struct {
	Id     *string `pulumi:"id"`
	TagKey string  `pulumi:"tagKey"`
}

// A collection of values returned by getTagsTagKeyIamPolicy.
type LookupTagsTagKeyIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	PolicyData string `pulumi:"policyData"`
	TagKey     string `pulumi:"tagKey"`
}

func LookupTagsTagKeyIamPolicyOutput(ctx *pulumi.Context, args LookupTagsTagKeyIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupTagsTagKeyIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTagsTagKeyIamPolicyResultOutput, error) {
			args := v.(LookupTagsTagKeyIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupTagsTagKeyIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getTagsTagKeyIamPolicy:getTagsTagKeyIamPolicy", args, LookupTagsTagKeyIamPolicyResultOutput{}, options).(LookupTagsTagKeyIamPolicyResultOutput), nil
		}).(LookupTagsTagKeyIamPolicyResultOutput)
}

// A collection of arguments for invoking getTagsTagKeyIamPolicy.
type LookupTagsTagKeyIamPolicyOutputArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	TagKey pulumi.StringInput    `pulumi:"tagKey"`
}

func (LookupTagsTagKeyIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagKeyIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getTagsTagKeyIamPolicy.
type LookupTagsTagKeyIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupTagsTagKeyIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagKeyIamPolicyResult)(nil)).Elem()
}

func (o LookupTagsTagKeyIamPolicyResultOutput) ToLookupTagsTagKeyIamPolicyResultOutput() LookupTagsTagKeyIamPolicyResultOutput {
	return o
}

func (o LookupTagsTagKeyIamPolicyResultOutput) ToLookupTagsTagKeyIamPolicyResultOutputWithContext(ctx context.Context) LookupTagsTagKeyIamPolicyResultOutput {
	return o
}

func (o LookupTagsTagKeyIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyIamPolicyResultOutput) TagKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyIamPolicyResult) string { return v.TagKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagsTagKeyIamPolicyResultOutput{})
}
