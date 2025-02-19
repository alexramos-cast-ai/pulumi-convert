// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupParameterManagerParameterVersion(ctx *pulumi.Context, args *LookupParameterManagerParameterVersionArgs, opts ...pulumi.InvokeOption) (*LookupParameterManagerParameterVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupParameterManagerParameterVersionResult
	err = ctx.InvokePackage("google-beta:index/getParameterManagerParameterVersion:getParameterManagerParameterVersion", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getParameterManagerParameterVersion.
type LookupParameterManagerParameterVersionArgs struct {
	Id                 *string `pulumi:"id"`
	Parameter          string  `pulumi:"parameter"`
	ParameterVersionId string  `pulumi:"parameterVersionId"`
	Project            *string `pulumi:"project"`
}

// A collection of values returned by getParameterManagerParameterVersion.
type LookupParameterManagerParameterVersionResult struct {
	CreateTime         string `pulumi:"createTime"`
	Disabled           bool   `pulumi:"disabled"`
	Id                 string `pulumi:"id"`
	Name               string `pulumi:"name"`
	Parameter          string `pulumi:"parameter"`
	ParameterData      string `pulumi:"parameterData"`
	ParameterVersionId string `pulumi:"parameterVersionId"`
	Project            string `pulumi:"project"`
	UpdateTime         string `pulumi:"updateTime"`
}

func LookupParameterManagerParameterVersionOutput(ctx *pulumi.Context, args LookupParameterManagerParameterVersionOutputArgs, opts ...pulumi.InvokeOption) LookupParameterManagerParameterVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupParameterManagerParameterVersionResultOutput, error) {
			args := v.(LookupParameterManagerParameterVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupParameterManagerParameterVersionResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getParameterManagerParameterVersion:getParameterManagerParameterVersion", args, LookupParameterManagerParameterVersionResultOutput{}, options).(LookupParameterManagerParameterVersionResultOutput), nil
		}).(LookupParameterManagerParameterVersionResultOutput)
}

// A collection of arguments for invoking getParameterManagerParameterVersion.
type LookupParameterManagerParameterVersionOutputArgs struct {
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	Parameter          pulumi.StringInput    `pulumi:"parameter"`
	ParameterVersionId pulumi.StringInput    `pulumi:"parameterVersionId"`
	Project            pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupParameterManagerParameterVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterManagerParameterVersionArgs)(nil)).Elem()
}

// A collection of values returned by getParameterManagerParameterVersion.
type LookupParameterManagerParameterVersionResultOutput struct{ *pulumi.OutputState }

func (LookupParameterManagerParameterVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupParameterManagerParameterVersionResult)(nil)).Elem()
}

func (o LookupParameterManagerParameterVersionResultOutput) ToLookupParameterManagerParameterVersionResultOutput() LookupParameterManagerParameterVersionResultOutput {
	return o
}

func (o LookupParameterManagerParameterVersionResultOutput) ToLookupParameterManagerParameterVersionResultOutputWithContext(ctx context.Context) LookupParameterManagerParameterVersionResultOutput {
	return o
}

func (o LookupParameterManagerParameterVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.Parameter }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) ParameterData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.ParameterData }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) ParameterVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.ParameterVersionId }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o LookupParameterManagerParameterVersionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupParameterManagerParameterVersionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupParameterManagerParameterVersionResultOutput{})
}
