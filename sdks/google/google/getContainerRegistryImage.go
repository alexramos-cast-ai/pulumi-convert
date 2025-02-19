// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetContainerRegistryImage(ctx *pulumi.Context, args *GetContainerRegistryImageArgs, opts ...pulumi.InvokeOption) (*GetContainerRegistryImageResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetContainerRegistryImageResult
	err = ctx.InvokePackage("google:index/getContainerRegistryImage:getContainerRegistryImage", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getContainerRegistryImage.
type GetContainerRegistryImageArgs struct {
	Digest  *string `pulumi:"digest"`
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	Tag     *string `pulumi:"tag"`
}

// A collection of values returned by getContainerRegistryImage.
type GetContainerRegistryImageResult struct {
	Digest   *string `pulumi:"digest"`
	Id       string  `pulumi:"id"`
	ImageUrl string  `pulumi:"imageUrl"`
	Name     string  `pulumi:"name"`
	Project  string  `pulumi:"project"`
	Region   *string `pulumi:"region"`
	Tag      *string `pulumi:"tag"`
}

func GetContainerRegistryImageOutput(ctx *pulumi.Context, args GetContainerRegistryImageOutputArgs, opts ...pulumi.InvokeOption) GetContainerRegistryImageResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetContainerRegistryImageResultOutput, error) {
			args := v.(GetContainerRegistryImageArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetContainerRegistryImageResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getContainerRegistryImage:getContainerRegistryImage", args, GetContainerRegistryImageResultOutput{}, options).(GetContainerRegistryImageResultOutput), nil
		}).(GetContainerRegistryImageResultOutput)
}

// A collection of arguments for invoking getContainerRegistryImage.
type GetContainerRegistryImageOutputArgs struct {
	Digest  pulumi.StringPtrInput `pulumi:"digest"`
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
	Tag     pulumi.StringPtrInput `pulumi:"tag"`
}

func (GetContainerRegistryImageOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRegistryImageArgs)(nil)).Elem()
}

// A collection of values returned by getContainerRegistryImage.
type GetContainerRegistryImageResultOutput struct{ *pulumi.OutputState }

func (GetContainerRegistryImageResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRegistryImageResult)(nil)).Elem()
}

func (o GetContainerRegistryImageResultOutput) ToGetContainerRegistryImageResultOutput() GetContainerRegistryImageResultOutput {
	return o
}

func (o GetContainerRegistryImageResultOutput) ToGetContainerRegistryImageResultOutputWithContext(ctx context.Context) GetContainerRegistryImageResultOutput {
	return o
}

func (o GetContainerRegistryImageResultOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o GetContainerRegistryImageResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetContainerRegistryImageResultOutput) ImageUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) string { return v.ImageUrl }).(pulumi.StringOutput)
}

func (o GetContainerRegistryImageResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetContainerRegistryImageResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetContainerRegistryImageResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o GetContainerRegistryImageResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetContainerRegistryImageResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerRegistryImageResultOutput{})
}
