// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetStorageObjectSignedUrl(ctx *pulumi.Context, args *GetStorageObjectSignedUrlArgs, opts ...pulumi.InvokeOption) (*GetStorageObjectSignedUrlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetStorageObjectSignedUrlResult
	err = ctx.InvokePackage("google:index/getStorageObjectSignedUrl:getStorageObjectSignedUrl", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageObjectSignedUrl.
type GetStorageObjectSignedUrlArgs struct {
	Bucket           string            `pulumi:"bucket"`
	ContentMd5       *string           `pulumi:"contentMd5"`
	ContentType      *string           `pulumi:"contentType"`
	Credentials      *string           `pulumi:"credentials"`
	Duration         *string           `pulumi:"duration"`
	ExtensionHeaders map[string]string `pulumi:"extensionHeaders"`
	HttpMethod       *string           `pulumi:"httpMethod"`
	Id               *string           `pulumi:"id"`
	Path             string            `pulumi:"path"`
}

// A collection of values returned by getStorageObjectSignedUrl.
type GetStorageObjectSignedUrlResult struct {
	Bucket           string            `pulumi:"bucket"`
	ContentMd5       *string           `pulumi:"contentMd5"`
	ContentType      *string           `pulumi:"contentType"`
	Credentials      *string           `pulumi:"credentials"`
	Duration         *string           `pulumi:"duration"`
	ExtensionHeaders map[string]string `pulumi:"extensionHeaders"`
	HttpMethod       *string           `pulumi:"httpMethod"`
	Id               string            `pulumi:"id"`
	Path             string            `pulumi:"path"`
	SignedUrl        string            `pulumi:"signedUrl"`
}

func GetStorageObjectSignedUrlOutput(ctx *pulumi.Context, args GetStorageObjectSignedUrlOutputArgs, opts ...pulumi.InvokeOption) GetStorageObjectSignedUrlResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetStorageObjectSignedUrlResultOutput, error) {
			args := v.(GetStorageObjectSignedUrlArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetStorageObjectSignedUrlResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getStorageObjectSignedUrl:getStorageObjectSignedUrl", args, GetStorageObjectSignedUrlResultOutput{}, options).(GetStorageObjectSignedUrlResultOutput), nil
		}).(GetStorageObjectSignedUrlResultOutput)
}

// A collection of arguments for invoking getStorageObjectSignedUrl.
type GetStorageObjectSignedUrlOutputArgs struct {
	Bucket           pulumi.StringInput    `pulumi:"bucket"`
	ContentMd5       pulumi.StringPtrInput `pulumi:"contentMd5"`
	ContentType      pulumi.StringPtrInput `pulumi:"contentType"`
	Credentials      pulumi.StringPtrInput `pulumi:"credentials"`
	Duration         pulumi.StringPtrInput `pulumi:"duration"`
	ExtensionHeaders pulumi.StringMapInput `pulumi:"extensionHeaders"`
	HttpMethod       pulumi.StringPtrInput `pulumi:"httpMethod"`
	Id               pulumi.StringPtrInput `pulumi:"id"`
	Path             pulumi.StringInput    `pulumi:"path"`
}

func (GetStorageObjectSignedUrlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageObjectSignedUrlArgs)(nil)).Elem()
}

// A collection of values returned by getStorageObjectSignedUrl.
type GetStorageObjectSignedUrlResultOutput struct{ *pulumi.OutputState }

func (GetStorageObjectSignedUrlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetStorageObjectSignedUrlResult)(nil)).Elem()
}

func (o GetStorageObjectSignedUrlResultOutput) ToGetStorageObjectSignedUrlResultOutput() GetStorageObjectSignedUrlResultOutput {
	return o
}

func (o GetStorageObjectSignedUrlResultOutput) ToGetStorageObjectSignedUrlResultOutputWithContext(ctx context.Context) GetStorageObjectSignedUrlResultOutput {
	return o
}

func (o GetStorageObjectSignedUrlResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) ContentMd5() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) *string { return v.ContentMd5 }).(pulumi.StringPtrOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) Credentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) *string { return v.Credentials }).(pulumi.StringPtrOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) ExtensionHeaders() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) map[string]string { return v.ExtensionHeaders }).(pulumi.StringMapOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) *string { return v.HttpMethod }).(pulumi.StringPtrOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetStorageObjectSignedUrlResultOutput) SignedUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetStorageObjectSignedUrlResult) string { return v.SignedUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetStorageObjectSignedUrlResultOutput{})
}
