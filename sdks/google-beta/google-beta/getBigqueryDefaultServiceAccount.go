// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBigqueryDefaultServiceAccount(ctx *pulumi.Context, args *GetBigqueryDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetBigqueryDefaultServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetBigqueryDefaultServiceAccountResult
	err = ctx.InvokePackage("google-beta:index/getBigqueryDefaultServiceAccount:getBigqueryDefaultServiceAccount", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBigqueryDefaultServiceAccount.
type GetBigqueryDefaultServiceAccountArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getBigqueryDefaultServiceAccount.
type GetBigqueryDefaultServiceAccountResult struct {
	Email   string `pulumi:"email"`
	Id      string `pulumi:"id"`
	Member  string `pulumi:"member"`
	Project string `pulumi:"project"`
}

func GetBigqueryDefaultServiceAccountOutput(ctx *pulumi.Context, args GetBigqueryDefaultServiceAccountOutputArgs, opts ...pulumi.InvokeOption) GetBigqueryDefaultServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBigqueryDefaultServiceAccountResultOutput, error) {
			args := v.(GetBigqueryDefaultServiceAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetBigqueryDefaultServiceAccountResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getBigqueryDefaultServiceAccount:getBigqueryDefaultServiceAccount", args, GetBigqueryDefaultServiceAccountResultOutput{}, options).(GetBigqueryDefaultServiceAccountResultOutput), nil
		}).(GetBigqueryDefaultServiceAccountResultOutput)
}

// A collection of arguments for invoking getBigqueryDefaultServiceAccount.
type GetBigqueryDefaultServiceAccountOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetBigqueryDefaultServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBigqueryDefaultServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getBigqueryDefaultServiceAccount.
type GetBigqueryDefaultServiceAccountResultOutput struct{ *pulumi.OutputState }

func (GetBigqueryDefaultServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBigqueryDefaultServiceAccountResult)(nil)).Elem()
}

func (o GetBigqueryDefaultServiceAccountResultOutput) ToGetBigqueryDefaultServiceAccountResultOutput() GetBigqueryDefaultServiceAccountResultOutput {
	return o
}

func (o GetBigqueryDefaultServiceAccountResultOutput) ToGetBigqueryDefaultServiceAccountResultOutputWithContext(ctx context.Context) GetBigqueryDefaultServiceAccountResultOutput {
	return o
}

func (o GetBigqueryDefaultServiceAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigqueryDefaultServiceAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o GetBigqueryDefaultServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigqueryDefaultServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBigqueryDefaultServiceAccountResultOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigqueryDefaultServiceAccountResult) string { return v.Member }).(pulumi.StringOutput)
}

func (o GetBigqueryDefaultServiceAccountResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetBigqueryDefaultServiceAccountResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBigqueryDefaultServiceAccountResultOutput{})
}
