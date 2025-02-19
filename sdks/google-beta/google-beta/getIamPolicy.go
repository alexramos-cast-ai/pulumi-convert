// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIamPolicy(ctx *pulumi.Context, args *GetIamPolicyArgs, opts ...pulumi.InvokeOption) (*GetIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetIamPolicyResult
	err = ctx.InvokePackage("google-beta:index/getIamPolicy:getIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIamPolicy.
type GetIamPolicyArgs struct {
	AuditConfigs []GetIamPolicyAuditConfig `pulumi:"auditConfigs"`
	Bindings     []GetIamPolicyBinding     `pulumi:"bindings"`
	Id           *string                   `pulumi:"id"`
}

// A collection of values returned by getIamPolicy.
type GetIamPolicyResult struct {
	AuditConfigs []GetIamPolicyAuditConfig `pulumi:"auditConfigs"`
	Bindings     []GetIamPolicyBinding     `pulumi:"bindings"`
	Id           string                    `pulumi:"id"`
	PolicyData   string                    `pulumi:"policyData"`
}

func GetIamPolicyOutput(ctx *pulumi.Context, args GetIamPolicyOutputArgs, opts ...pulumi.InvokeOption) GetIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetIamPolicyResultOutput, error) {
			args := v.(GetIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getIamPolicy:getIamPolicy", args, GetIamPolicyResultOutput{}, options).(GetIamPolicyResultOutput), nil
		}).(GetIamPolicyResultOutput)
}

// A collection of arguments for invoking getIamPolicy.
type GetIamPolicyOutputArgs struct {
	AuditConfigs GetIamPolicyAuditConfigArrayInput `pulumi:"auditConfigs"`
	Bindings     GetIamPolicyBindingArrayInput     `pulumi:"bindings"`
	Id           pulumi.StringPtrInput             `pulumi:"id"`
}

func (GetIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getIamPolicy.
type GetIamPolicyResultOutput struct{ *pulumi.OutputState }

func (GetIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetIamPolicyResult)(nil)).Elem()
}

func (o GetIamPolicyResultOutput) ToGetIamPolicyResultOutput() GetIamPolicyResultOutput {
	return o
}

func (o GetIamPolicyResultOutput) ToGetIamPolicyResultOutputWithContext(ctx context.Context) GetIamPolicyResultOutput {
	return o
}

func (o GetIamPolicyResultOutput) AuditConfigs() GetIamPolicyAuditConfigArrayOutput {
	return o.ApplyT(func(v GetIamPolicyResult) []GetIamPolicyAuditConfig { return v.AuditConfigs }).(GetIamPolicyAuditConfigArrayOutput)
}

func (o GetIamPolicyResultOutput) Bindings() GetIamPolicyBindingArrayOutput {
	return o.ApplyT(func(v GetIamPolicyResult) []GetIamPolicyBinding { return v.Bindings }).(GetIamPolicyBindingArrayOutput)
}

func (o GetIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v GetIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetIamPolicyResultOutput{})
}
