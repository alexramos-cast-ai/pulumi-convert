// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupServiceAccount(ctx *pulumi.Context, args *LookupServiceAccountArgs, opts ...pulumi.InvokeOption) (*LookupServiceAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupServiceAccountResult
	err = ctx.InvokePackage("google-beta:index/getServiceAccount:getServiceAccount", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServiceAccount.
type LookupServiceAccountArgs struct {
	AccountId string  `pulumi:"accountId"`
	Id        *string `pulumi:"id"`
	Project   *string `pulumi:"project"`
}

// A collection of values returned by getServiceAccount.
type LookupServiceAccountResult struct {
	AccountId   string  `pulumi:"accountId"`
	Disabled    bool    `pulumi:"disabled"`
	DisplayName string  `pulumi:"displayName"`
	Email       string  `pulumi:"email"`
	Id          string  `pulumi:"id"`
	Member      string  `pulumi:"member"`
	Name        string  `pulumi:"name"`
	Project     *string `pulumi:"project"`
	UniqueId    string  `pulumi:"uniqueId"`
}

func LookupServiceAccountOutput(ctx *pulumi.Context, args LookupServiceAccountOutputArgs, opts ...pulumi.InvokeOption) LookupServiceAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupServiceAccountResultOutput, error) {
			args := v.(LookupServiceAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupServiceAccountResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getServiceAccount:getServiceAccount", args, LookupServiceAccountResultOutput{}, options).(LookupServiceAccountResultOutput), nil
		}).(LookupServiceAccountResultOutput)
}

// A collection of arguments for invoking getServiceAccount.
type LookupServiceAccountOutputArgs struct {
	AccountId pulumi.StringInput    `pulumi:"accountId"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupServiceAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountArgs)(nil)).Elem()
}

// A collection of values returned by getServiceAccount.
type LookupServiceAccountResultOutput struct{ *pulumi.OutputState }

func (LookupServiceAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServiceAccountResult)(nil)).Elem()
}

func (o LookupServiceAccountResultOutput) ToLookupServiceAccountResultOutput() LookupServiceAccountResultOutput {
	return o
}

func (o LookupServiceAccountResultOutput) ToLookupServiceAccountResultOutputWithContext(ctx context.Context) LookupServiceAccountResultOutput {
	return o
}

func (o LookupServiceAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupServiceAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Member() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Member }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupServiceAccountResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupServiceAccountResultOutput) UniqueId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServiceAccountResult) string { return v.UniqueId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServiceAccountResultOutput{})
}
