// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagsTagKey(ctx *pulumi.Context, args *LookupTagsTagKeyArgs, opts ...pulumi.InvokeOption) (*LookupTagsTagKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupTagsTagKeyResult
	err = ctx.InvokePackage("google-beta:index/getTagsTagKey:getTagsTagKey", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagsTagKey.
type LookupTagsTagKeyArgs struct {
	Parent    string `pulumi:"parent"`
	ShortName string `pulumi:"shortName"`
}

// A collection of values returned by getTagsTagKey.
type LookupTagsTagKeyResult struct {
	CreateTime     string `pulumi:"createTime"`
	Description    string `pulumi:"description"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	NamespacedName string `pulumi:"namespacedName"`
	Parent         string `pulumi:"parent"`
	ShortName      string `pulumi:"shortName"`
	UpdateTime     string `pulumi:"updateTime"`
}

func LookupTagsTagKeyOutput(ctx *pulumi.Context, args LookupTagsTagKeyOutputArgs, opts ...pulumi.InvokeOption) LookupTagsTagKeyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTagsTagKeyResultOutput, error) {
			args := v.(LookupTagsTagKeyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupTagsTagKeyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getTagsTagKey:getTagsTagKey", args, LookupTagsTagKeyResultOutput{}, options).(LookupTagsTagKeyResultOutput), nil
		}).(LookupTagsTagKeyResultOutput)
}

// A collection of arguments for invoking getTagsTagKey.
type LookupTagsTagKeyOutputArgs struct {
	Parent    pulumi.StringInput `pulumi:"parent"`
	ShortName pulumi.StringInput `pulumi:"shortName"`
}

func (LookupTagsTagKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagKeyArgs)(nil)).Elem()
}

// A collection of values returned by getTagsTagKey.
type LookupTagsTagKeyResultOutput struct{ *pulumi.OutputState }

func (LookupTagsTagKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagKeyResult)(nil)).Elem()
}

func (o LookupTagsTagKeyResultOutput) ToLookupTagsTagKeyResultOutput() LookupTagsTagKeyResultOutput {
	return o
}

func (o LookupTagsTagKeyResultOutput) ToLookupTagsTagKeyResultOutputWithContext(ctx context.Context) LookupTagsTagKeyResultOutput {
	return o
}

func (o LookupTagsTagKeyResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) NamespacedName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.NamespacedName }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.ShortName }).(pulumi.StringOutput)
}

func (o LookupTagsTagKeyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagKeyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagsTagKeyResultOutput{})
}
