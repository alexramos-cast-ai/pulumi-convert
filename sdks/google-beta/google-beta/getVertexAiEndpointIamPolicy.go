// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVertexAiEndpointIamPolicy(ctx *pulumi.Context, args *LookupVertexAiEndpointIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupVertexAiEndpointIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupVertexAiEndpointIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getVertexAiEndpointIamPolicy:getVertexAiEndpointIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVertexAiEndpointIamPolicy.
type LookupVertexAiEndpointIamPolicyArgs struct {
	Endpoint string  `pulumi:"endpoint"`
	Id       *string `pulumi:"id"`
	Location *string `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getVertexAiEndpointIamPolicy.
type LookupVertexAiEndpointIamPolicyResult struct {
	Endpoint   string `pulumi:"endpoint"`
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Location   string `pulumi:"location"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupVertexAiEndpointIamPolicyOutput(ctx *pulumi.Context, args LookupVertexAiEndpointIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupVertexAiEndpointIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVertexAiEndpointIamPolicyResultOutput, error) {
			args := v.(LookupVertexAiEndpointIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupVertexAiEndpointIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getVertexAiEndpointIamPolicy:getVertexAiEndpointIamPolicy", args, LookupVertexAiEndpointIamPolicyResultOutput{}, options).(LookupVertexAiEndpointIamPolicyResultOutput), nil
		}).(LookupVertexAiEndpointIamPolicyResultOutput)
}

// A collection of arguments for invoking getVertexAiEndpointIamPolicy.
type LookupVertexAiEndpointIamPolicyOutputArgs struct {
	Endpoint pulumi.StringInput    `pulumi:"endpoint"`
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringPtrInput `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupVertexAiEndpointIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVertexAiEndpointIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getVertexAiEndpointIamPolicy.
type LookupVertexAiEndpointIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupVertexAiEndpointIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVertexAiEndpointIamPolicyResult)(nil)).Elem()
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) ToLookupVertexAiEndpointIamPolicyResultOutput() LookupVertexAiEndpointIamPolicyResultOutput {
	return o
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) ToLookupVertexAiEndpointIamPolicyResultOutputWithContext(ctx context.Context) LookupVertexAiEndpointIamPolicyResultOutput {
	return o
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupVertexAiEndpointIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVertexAiEndpointIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVertexAiEndpointIamPolicyResultOutput{})
}
