// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDataCatalogTagTemplateIamPolicy(ctx *pulumi.Context, args *LookupDataCatalogTagTemplateIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupDataCatalogTagTemplateIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupDataCatalogTagTemplateIamPolicyResult
	err = ctx.InvokePackage("google:index/getDataCatalogTagTemplateIamPolicy:getDataCatalogTagTemplateIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataCatalogTagTemplateIamPolicy.
type LookupDataCatalogTagTemplateIamPolicyArgs struct {
	Id          *string `pulumi:"id"`
	Project     *string `pulumi:"project"`
	Region      *string `pulumi:"region"`
	TagTemplate string  `pulumi:"tagTemplate"`
}

// A collection of values returned by getDataCatalogTagTemplateIamPolicy.
type LookupDataCatalogTagTemplateIamPolicyResult struct {
	Etag        string `pulumi:"etag"`
	Id          string `pulumi:"id"`
	PolicyData  string `pulumi:"policyData"`
	Project     string `pulumi:"project"`
	Region      string `pulumi:"region"`
	TagTemplate string `pulumi:"tagTemplate"`
}

func LookupDataCatalogTagTemplateIamPolicyOutput(ctx *pulumi.Context, args LookupDataCatalogTagTemplateIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupDataCatalogTagTemplateIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDataCatalogTagTemplateIamPolicyResultOutput, error) {
			args := v.(LookupDataCatalogTagTemplateIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupDataCatalogTagTemplateIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getDataCatalogTagTemplateIamPolicy:getDataCatalogTagTemplateIamPolicy", args, LookupDataCatalogTagTemplateIamPolicyResultOutput{}, options).(LookupDataCatalogTagTemplateIamPolicyResultOutput), nil
		}).(LookupDataCatalogTagTemplateIamPolicyResultOutput)
}

// A collection of arguments for invoking getDataCatalogTagTemplateIamPolicy.
type LookupDataCatalogTagTemplateIamPolicyOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	Region      pulumi.StringPtrInput `pulumi:"region"`
	TagTemplate pulumi.StringInput    `pulumi:"tagTemplate"`
}

func (LookupDataCatalogTagTemplateIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCatalogTagTemplateIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getDataCatalogTagTemplateIamPolicy.
type LookupDataCatalogTagTemplateIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupDataCatalogTagTemplateIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataCatalogTagTemplateIamPolicyResult)(nil)).Elem()
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) ToLookupDataCatalogTagTemplateIamPolicyResultOutput() LookupDataCatalogTagTemplateIamPolicyResultOutput {
	return o
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) ToLookupDataCatalogTagTemplateIamPolicyResultOutputWithContext(ctx context.Context) LookupDataCatalogTagTemplateIamPolicyResultOutput {
	return o
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o LookupDataCatalogTagTemplateIamPolicyResultOutput) TagTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataCatalogTagTemplateIamPolicyResult) string { return v.TagTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataCatalogTagTemplateIamPolicyResultOutput{})
}
