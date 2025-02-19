// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeSnapshotIamPolicy(ctx *pulumi.Context, args *LookupComputeSnapshotIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupComputeSnapshotIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeSnapshotIamPolicyResult
	err = ctx.InvokePackage("google:index/getComputeSnapshotIamPolicy:getComputeSnapshotIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeSnapshotIamPolicy.
type LookupComputeSnapshotIamPolicyArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getComputeSnapshotIamPolicy.
type LookupComputeSnapshotIamPolicyResult struct {
	Etag       string `pulumi:"etag"`
	Id         string `pulumi:"id"`
	Name       string `pulumi:"name"`
	PolicyData string `pulumi:"policyData"`
	Project    string `pulumi:"project"`
}

func LookupComputeSnapshotIamPolicyOutput(ctx *pulumi.Context, args LookupComputeSnapshotIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupComputeSnapshotIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeSnapshotIamPolicyResultOutput, error) {
			args := v.(LookupComputeSnapshotIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeSnapshotIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComputeSnapshotIamPolicy:getComputeSnapshotIamPolicy", args, LookupComputeSnapshotIamPolicyResultOutput{}, options).(LookupComputeSnapshotIamPolicyResultOutput), nil
		}).(LookupComputeSnapshotIamPolicyResultOutput)
}

// A collection of arguments for invoking getComputeSnapshotIamPolicy.
type LookupComputeSnapshotIamPolicyOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupComputeSnapshotIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSnapshotIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getComputeSnapshotIamPolicy.
type LookupComputeSnapshotIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupComputeSnapshotIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeSnapshotIamPolicyResult)(nil)).Elem()
}

func (o LookupComputeSnapshotIamPolicyResultOutput) ToLookupComputeSnapshotIamPolicyResultOutput() LookupComputeSnapshotIamPolicyResultOutput {
	return o
}

func (o LookupComputeSnapshotIamPolicyResultOutput) ToLookupComputeSnapshotIamPolicyResultOutputWithContext(ctx context.Context) LookupComputeSnapshotIamPolicyResultOutput {
	return o
}

func (o LookupComputeSnapshotIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupComputeSnapshotIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeSnapshotIamPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotIamPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeSnapshotIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func (o LookupComputeSnapshotIamPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeSnapshotIamPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeSnapshotIamPolicyResultOutput{})
}
