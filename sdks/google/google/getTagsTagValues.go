// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetTagsTagValues(ctx *pulumi.Context, args *GetTagsTagValuesArgs, opts ...pulumi.InvokeOption) (*GetTagsTagValuesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetTagsTagValuesResult
	err = ctx.InvokePackage("google:index/getTagsTagValues:getTagsTagValues", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagsTagValues.
type GetTagsTagValuesArgs struct {
	Id     *string `pulumi:"id"`
	Parent string  `pulumi:"parent"`
}

// A collection of values returned by getTagsTagValues.
type GetTagsTagValuesResult struct {
	Id     string                  `pulumi:"id"`
	Parent string                  `pulumi:"parent"`
	Values []GetTagsTagValuesValue `pulumi:"values"`
}

func GetTagsTagValuesOutput(ctx *pulumi.Context, args GetTagsTagValuesOutputArgs, opts ...pulumi.InvokeOption) GetTagsTagValuesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTagsTagValuesResultOutput, error) {
			args := v.(GetTagsTagValuesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetTagsTagValuesResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getTagsTagValues:getTagsTagValues", args, GetTagsTagValuesResultOutput{}, options).(GetTagsTagValuesResultOutput), nil
		}).(GetTagsTagValuesResultOutput)
}

// A collection of arguments for invoking getTagsTagValues.
type GetTagsTagValuesOutputArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Parent pulumi.StringInput    `pulumi:"parent"`
}

func (GetTagsTagValuesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTagValuesArgs)(nil)).Elem()
}

// A collection of values returned by getTagsTagValues.
type GetTagsTagValuesResultOutput struct{ *pulumi.OutputState }

func (GetTagsTagValuesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTagsTagValuesResult)(nil)).Elem()
}

func (o GetTagsTagValuesResultOutput) ToGetTagsTagValuesResultOutput() GetTagsTagValuesResultOutput {
	return o
}

func (o GetTagsTagValuesResultOutput) ToGetTagsTagValuesResultOutputWithContext(ctx context.Context) GetTagsTagValuesResultOutput {
	return o
}

func (o GetTagsTagValuesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTagValuesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetTagsTagValuesResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v GetTagsTagValuesResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o GetTagsTagValuesResultOutput) Values() GetTagsTagValuesValueArrayOutput {
	return o.ApplyT(func(v GetTagsTagValuesResult) []GetTagsTagValuesValue { return v.Values }).(GetTagsTagValuesValueArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTagsTagValuesResultOutput{})
}
