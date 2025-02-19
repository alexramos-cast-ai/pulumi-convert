// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagsTagValue(ctx *pulumi.Context, args *LookupTagsTagValueArgs, opts ...pulumi.InvokeOption) (*LookupTagsTagValueResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupTagsTagValueResult
	err = ctx.InvokePackage("google:index/getTagsTagValue:getTagsTagValue", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTagsTagValue.
type LookupTagsTagValueArgs struct {
	Parent    string `pulumi:"parent"`
	ShortName string `pulumi:"shortName"`
}

// A collection of values returned by getTagsTagValue.
type LookupTagsTagValueResult struct {
	CreateTime     string `pulumi:"createTime"`
	Description    string `pulumi:"description"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	NamespacedName string `pulumi:"namespacedName"`
	Parent         string `pulumi:"parent"`
	ShortName      string `pulumi:"shortName"`
	UpdateTime     string `pulumi:"updateTime"`
}

func LookupTagsTagValueOutput(ctx *pulumi.Context, args LookupTagsTagValueOutputArgs, opts ...pulumi.InvokeOption) LookupTagsTagValueResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupTagsTagValueResultOutput, error) {
			args := v.(LookupTagsTagValueArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupTagsTagValueResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getTagsTagValue:getTagsTagValue", args, LookupTagsTagValueResultOutput{}, options).(LookupTagsTagValueResultOutput), nil
		}).(LookupTagsTagValueResultOutput)
}

// A collection of arguments for invoking getTagsTagValue.
type LookupTagsTagValueOutputArgs struct {
	Parent    pulumi.StringInput `pulumi:"parent"`
	ShortName pulumi.StringInput `pulumi:"shortName"`
}

func (LookupTagsTagValueOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagValueArgs)(nil)).Elem()
}

// A collection of values returned by getTagsTagValue.
type LookupTagsTagValueResultOutput struct{ *pulumi.OutputState }

func (LookupTagsTagValueResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagsTagValueResult)(nil)).Elem()
}

func (o LookupTagsTagValueResultOutput) ToLookupTagsTagValueResultOutput() LookupTagsTagValueResultOutput {
	return o
}

func (o LookupTagsTagValueResultOutput) ToLookupTagsTagValueResultOutputWithContext(ctx context.Context) LookupTagsTagValueResultOutput {
	return o
}

func (o LookupTagsTagValueResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) NamespacedName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.NamespacedName }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.Parent }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.ShortName }).(pulumi.StringOutput)
}

func (o LookupTagsTagValueResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagsTagValueResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagsTagValueResultOutput{})
}
