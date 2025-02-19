// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupProjectOrganizationPolicy(ctx *pulumi.Context, args *LookupProjectOrganizationPolicyArgs, opts ...pulumi.InvokeOption) (*LookupProjectOrganizationPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupProjectOrganizationPolicyResult
	err = ctx.InvokePackage("google:index/getProjectOrganizationPolicy:getProjectOrganizationPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProjectOrganizationPolicy.
type LookupProjectOrganizationPolicyArgs struct {
	Constraint string  `pulumi:"constraint"`
	Id         *string `pulumi:"id"`
	Project    string  `pulumi:"project"`
}

// A collection of values returned by getProjectOrganizationPolicy.
type LookupProjectOrganizationPolicyResult struct {
	BooleanPolicies []GetProjectOrganizationPolicyBooleanPolicy `pulumi:"booleanPolicies"`
	Constraint      string                                      `pulumi:"constraint"`
	Etag            string                                      `pulumi:"etag"`
	Id              string                                      `pulumi:"id"`
	ListPolicies    []GetProjectOrganizationPolicyListPolicy    `pulumi:"listPolicies"`
	Project         string                                      `pulumi:"project"`
	RestorePolicies []GetProjectOrganizationPolicyRestorePolicy `pulumi:"restorePolicies"`
	UpdateTime      string                                      `pulumi:"updateTime"`
	Version         float64                                     `pulumi:"version"`
}

func LookupProjectOrganizationPolicyOutput(ctx *pulumi.Context, args LookupProjectOrganizationPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupProjectOrganizationPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupProjectOrganizationPolicyResultOutput, error) {
			args := v.(LookupProjectOrganizationPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupProjectOrganizationPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getProjectOrganizationPolicy:getProjectOrganizationPolicy", args, LookupProjectOrganizationPolicyResultOutput{}, options).(LookupProjectOrganizationPolicyResultOutput), nil
		}).(LookupProjectOrganizationPolicyResultOutput)
}

// A collection of arguments for invoking getProjectOrganizationPolicy.
type LookupProjectOrganizationPolicyOutputArgs struct {
	Constraint pulumi.StringInput    `pulumi:"constraint"`
	Id         pulumi.StringPtrInput `pulumi:"id"`
	Project    pulumi.StringInput    `pulumi:"project"`
}

func (LookupProjectOrganizationPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectOrganizationPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getProjectOrganizationPolicy.
type LookupProjectOrganizationPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupProjectOrganizationPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupProjectOrganizationPolicyResult)(nil)).Elem()
}

func (o LookupProjectOrganizationPolicyResultOutput) ToLookupProjectOrganizationPolicyResultOutput() LookupProjectOrganizationPolicyResultOutput {
	return o
}

func (o LookupProjectOrganizationPolicyResultOutput) ToLookupProjectOrganizationPolicyResultOutputWithContext(ctx context.Context) LookupProjectOrganizationPolicyResultOutput {
	return o
}

func (o LookupProjectOrganizationPolicyResultOutput) BooleanPolicies() GetProjectOrganizationPolicyBooleanPolicyArrayOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) []GetProjectOrganizationPolicyBooleanPolicy {
		return v.BooleanPolicies
	}).(GetProjectOrganizationPolicyBooleanPolicyArrayOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) Constraint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) string { return v.Constraint }).(pulumi.StringOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) ListPolicies() GetProjectOrganizationPolicyListPolicyArrayOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) []GetProjectOrganizationPolicyListPolicy {
		return v.ListPolicies
	}).(GetProjectOrganizationPolicyListPolicyArrayOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) RestorePolicies() GetProjectOrganizationPolicyRestorePolicyArrayOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) []GetProjectOrganizationPolicyRestorePolicy {
		return v.RestorePolicies
	}).(GetProjectOrganizationPolicyRestorePolicyArrayOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o LookupProjectOrganizationPolicyResultOutput) Version() pulumi.Float64Output {
	return o.ApplyT(func(v LookupProjectOrganizationPolicyResult) float64 { return v.Version }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupProjectOrganizationPolicyResultOutput{})
}
