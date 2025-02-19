// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupKmsAutokeyConfig(ctx *pulumi.Context, args *LookupKmsAutokeyConfigArgs, opts ...pulumi.InvokeOption) (*LookupKmsAutokeyConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupKmsAutokeyConfigResult
	err = ctx.InvokePackage("google-beta:index/getKmsAutokeyConfig:getKmsAutokeyConfig", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKmsAutokeyConfig.
type LookupKmsAutokeyConfigArgs struct {
	Folder string  `pulumi:"folder"`
	Id     *string `pulumi:"id"`
}

// A collection of values returned by getKmsAutokeyConfig.
type LookupKmsAutokeyConfigResult struct {
	Folder     string `pulumi:"folder"`
	Id         string `pulumi:"id"`
	KeyProject string `pulumi:"keyProject"`
}

func LookupKmsAutokeyConfigOutput(ctx *pulumi.Context, args LookupKmsAutokeyConfigOutputArgs, opts ...pulumi.InvokeOption) LookupKmsAutokeyConfigResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupKmsAutokeyConfigResultOutput, error) {
			args := v.(LookupKmsAutokeyConfigArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupKmsAutokeyConfigResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getKmsAutokeyConfig:getKmsAutokeyConfig", args, LookupKmsAutokeyConfigResultOutput{}, options).(LookupKmsAutokeyConfigResultOutput), nil
		}).(LookupKmsAutokeyConfigResultOutput)
}

// A collection of arguments for invoking getKmsAutokeyConfig.
type LookupKmsAutokeyConfigOutputArgs struct {
	Folder pulumi.StringInput    `pulumi:"folder"`
	Id     pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupKmsAutokeyConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKmsAutokeyConfigArgs)(nil)).Elem()
}

// A collection of values returned by getKmsAutokeyConfig.
type LookupKmsAutokeyConfigResultOutput struct{ *pulumi.OutputState }

func (LookupKmsAutokeyConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupKmsAutokeyConfigResult)(nil)).Elem()
}

func (o LookupKmsAutokeyConfigResultOutput) ToLookupKmsAutokeyConfigResultOutput() LookupKmsAutokeyConfigResultOutput {
	return o
}

func (o LookupKmsAutokeyConfigResultOutput) ToLookupKmsAutokeyConfigResultOutputWithContext(ctx context.Context) LookupKmsAutokeyConfigResultOutput {
	return o
}

func (o LookupKmsAutokeyConfigResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsAutokeyConfigResult) string { return v.Folder }).(pulumi.StringOutput)
}

func (o LookupKmsAutokeyConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsAutokeyConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupKmsAutokeyConfigResultOutput) KeyProject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupKmsAutokeyConfigResult) string { return v.KeyProject }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupKmsAutokeyConfigResultOutput{})
}
