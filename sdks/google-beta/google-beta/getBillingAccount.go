// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetBillingAccount(ctx *pulumi.Context, args *GetBillingAccountArgs, opts ...pulumi.InvokeOption) (*GetBillingAccountResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetBillingAccountResult
	err = ctx.InvokePackage("google-beta:index/getBillingAccount:getBillingAccount", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getBillingAccount.
type GetBillingAccountArgs struct {
	BillingAccount *string `pulumi:"billingAccount"`
	DisplayName    *string `pulumi:"displayName"`
	Id             *string `pulumi:"id"`
	LookupProjects *bool   `pulumi:"lookupProjects"`
	Open           *bool   `pulumi:"open"`
}

// A collection of values returned by getBillingAccount.
type GetBillingAccountResult struct {
	BillingAccount *string  `pulumi:"billingAccount"`
	DisplayName    string   `pulumi:"displayName"`
	Id             string   `pulumi:"id"`
	LookupProjects *bool    `pulumi:"lookupProjects"`
	Name           string   `pulumi:"name"`
	Open           bool     `pulumi:"open"`
	ProjectIds     []string `pulumi:"projectIds"`
}

func GetBillingAccountOutput(ctx *pulumi.Context, args GetBillingAccountOutputArgs, opts ...pulumi.InvokeOption) GetBillingAccountResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetBillingAccountResultOutput, error) {
			args := v.(GetBillingAccountArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetBillingAccountResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getBillingAccount:getBillingAccount", args, GetBillingAccountResultOutput{}, options).(GetBillingAccountResultOutput), nil
		}).(GetBillingAccountResultOutput)
}

// A collection of arguments for invoking getBillingAccount.
type GetBillingAccountOutputArgs struct {
	BillingAccount pulumi.StringPtrInput `pulumi:"billingAccount"`
	DisplayName    pulumi.StringPtrInput `pulumi:"displayName"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	LookupProjects pulumi.BoolPtrInput   `pulumi:"lookupProjects"`
	Open           pulumi.BoolPtrInput   `pulumi:"open"`
}

func (GetBillingAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingAccountArgs)(nil)).Elem()
}

// A collection of values returned by getBillingAccount.
type GetBillingAccountResultOutput struct{ *pulumi.OutputState }

func (GetBillingAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetBillingAccountResult)(nil)).Elem()
}

func (o GetBillingAccountResultOutput) ToGetBillingAccountResultOutput() GetBillingAccountResultOutput {
	return o
}

func (o GetBillingAccountResultOutput) ToGetBillingAccountResultOutputWithContext(ctx context.Context) GetBillingAccountResultOutput {
	return o
}

func (o GetBillingAccountResultOutput) BillingAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetBillingAccountResult) *string { return v.BillingAccount }).(pulumi.StringPtrOutput)
}

func (o GetBillingAccountResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingAccountResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GetBillingAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetBillingAccountResultOutput) LookupProjects() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetBillingAccountResult) *bool { return v.LookupProjects }).(pulumi.BoolPtrOutput)
}

func (o GetBillingAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetBillingAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetBillingAccountResultOutput) Open() pulumi.BoolOutput {
	return o.ApplyT(func(v GetBillingAccountResult) bool { return v.Open }).(pulumi.BoolOutput)
}

func (o GetBillingAccountResultOutput) ProjectIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetBillingAccountResult) []string { return v.ProjectIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetBillingAccountResultOutput{})
}
