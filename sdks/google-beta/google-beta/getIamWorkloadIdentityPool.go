// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIamWorkloadIdentityPool(ctx *pulumi.Context, args *LookupIamWorkloadIdentityPoolArgs, opts ...pulumi.InvokeOption) (*LookupIamWorkloadIdentityPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupIamWorkloadIdentityPoolResult
	err = ctx.InvokePackage("google-beta:index/getIamWorkloadIdentityPool:getIamWorkloadIdentityPool", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamWorkloadIdentityPool.
type LookupIamWorkloadIdentityPoolArgs struct {
	Id                     *string `pulumi:"id"`
	Project                *string `pulumi:"project"`
	WorkloadIdentityPoolId string  `pulumi:"workloadIdentityPoolId"`
}

// A collection of values returned by getIamWorkloadIdentityPool.
type LookupIamWorkloadIdentityPoolResult struct {
	Description            string  `pulumi:"description"`
	Disabled               bool    `pulumi:"disabled"`
	DisplayName            string  `pulumi:"displayName"`
	Id                     string  `pulumi:"id"`
	Name                   string  `pulumi:"name"`
	Project                *string `pulumi:"project"`
	State                  string  `pulumi:"state"`
	WorkloadIdentityPoolId string  `pulumi:"workloadIdentityPoolId"`
}

func LookupIamWorkloadIdentityPoolOutput(ctx *pulumi.Context, args LookupIamWorkloadIdentityPoolOutputArgs, opts ...pulumi.InvokeOption) LookupIamWorkloadIdentityPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIamWorkloadIdentityPoolResultOutput, error) {
			args := v.(LookupIamWorkloadIdentityPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupIamWorkloadIdentityPoolResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getIamWorkloadIdentityPool:getIamWorkloadIdentityPool", args, LookupIamWorkloadIdentityPoolResultOutput{}, options).(LookupIamWorkloadIdentityPoolResultOutput), nil
		}).(LookupIamWorkloadIdentityPoolResultOutput)
}

// A collection of arguments for invoking getIamWorkloadIdentityPool.
type LookupIamWorkloadIdentityPoolOutputArgs struct {
	Id                     pulumi.StringPtrInput `pulumi:"id"`
	Project                pulumi.StringPtrInput `pulumi:"project"`
	WorkloadIdentityPoolId pulumi.StringInput    `pulumi:"workloadIdentityPoolId"`
}

func (LookupIamWorkloadIdentityPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamWorkloadIdentityPoolArgs)(nil)).Elem()
}

// A collection of values returned by getIamWorkloadIdentityPool.
type LookupIamWorkloadIdentityPoolResultOutput struct{ *pulumi.OutputState }

func (LookupIamWorkloadIdentityPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIamWorkloadIdentityPoolResult)(nil)).Elem()
}

func (o LookupIamWorkloadIdentityPoolResultOutput) ToLookupIamWorkloadIdentityPoolResultOutput() LookupIamWorkloadIdentityPoolResultOutput {
	return o
}

func (o LookupIamWorkloadIdentityPoolResultOutput) ToLookupIamWorkloadIdentityPoolResultOutputWithContext(ctx context.Context) LookupIamWorkloadIdentityPoolResultOutput {
	return o
}

func (o LookupIamWorkloadIdentityPoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupIamWorkloadIdentityPoolResultOutput) WorkloadIdentityPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIamWorkloadIdentityPoolResult) string { return v.WorkloadIdentityPoolId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIamWorkloadIdentityPoolResultOutput{})
}
