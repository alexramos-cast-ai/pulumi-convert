// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVertexAiFeaturestoreIamPolicy(ctx *pulumi.Context, args *LookupVertexAiFeaturestoreIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupVertexAiFeaturestoreIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupVertexAiFeaturestoreIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getVertexAiFeaturestoreIamPolicy:getVertexAiFeaturestoreIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVertexAiFeaturestoreIamPolicy.
type LookupVertexAiFeaturestoreIamPolicyArgs struct {
	Featurestore string  `pulumi:"featurestore"`
	Id           *string `pulumi:"id"`
	Project      *string `pulumi:"project"`
	Region       *string `pulumi:"region"`
}

// A collection of values returned by getVertexAiFeaturestoreIamPolicy.
type LookupVertexAiFeaturestoreIamPolicyResult struct {
	Etag         string `pulumi:"etag"`
	Featurestore string `pulumi:"featurestore"`
	Id           string `pulumi:"id"`
	PolicyData   string `pulumi:"policyData"`
	Project      string `pulumi:"project"`
	Region       string `pulumi:"region"`
}

func LookupVertexAiFeaturestoreIamPolicyOutput(ctx *pulumi.Context, args LookupVertexAiFeaturestoreIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupVertexAiFeaturestoreIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVertexAiFeaturestoreIamPolicyResultOutput, error) {
			args := v.(LookupVertexAiFeaturestoreIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupVertexAiFeaturestoreIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getVertexAiFeaturestoreIamPolicy:getVertexAiFeaturestoreIamPolicy", args, LookupVertexAiFeaturestoreIamPolicyResultOutput{}, options).(LookupVertexAiFeaturestoreIamPolicyResultOutput), nil
		}).(LookupVertexAiFeaturestoreIamPolicyResultOutput)
}

// A collection of arguments for invoking getVertexAiFeaturestoreIamPolicy.
type LookupVertexAiFeaturestoreIamPolicyOutputArgs struct {
	Featurestore pulumi.StringInput    `pulumi:"featurestore"`
	Id           pulumi.StringPtrInput `pulumi:"id"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	Region       pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupVertexAiFeaturestoreIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVertexAiFeaturestoreIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getVertexAiFeaturestoreIamPolicy.
type LookupVertexAiFeaturestoreIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupVertexAiFeaturestoreIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVertexAiFeaturestoreIamPolicyResult)(nil)).Elem()
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) ToLookupVertexAiFeaturestoreIamPolicyResultOutput() LookupVertexAiFeaturestoreIamPolicyResultOutput {
	return o
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) ToLookupVertexAiFeaturestoreIamPolicyResultOutputWithContext(ctx context.Context) LookupVertexAiFeaturestoreIamPolicyResultOutput {
	return o
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) Featurestore() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.Featurestore }).(pulumi.StringOutput)
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupVertexAiFeaturestoreIamPolicyResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiFeaturestoreIamPolicyResult) string { return v.Region }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVertexAiFeaturestoreIamPolicyResultOutput{})
}
