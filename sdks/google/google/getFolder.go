// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFolder(ctx *pulumi.Context, args *LookupFolderArgs, opts ...pulumi.InvokeOption) (*LookupFolderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupFolderResult
	err = ctx.InvokePackage("google:index/getFolder:getFolder", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFolder.
type LookupFolderArgs struct {
	Folder             string  `pulumi:"folder"`
	Id                 *string `pulumi:"id"`
	LookupOrganization *bool   `pulumi:"lookupOrganization"`
}

// A collection of values returned by getFolder.
type LookupFolderResult struct {
	CreateTime         string `pulumi:"createTime"`
	DeletionProtection bool   `pulumi:"deletionProtection"`
	DisplayName        string `pulumi:"displayName"`
	Folder             string `pulumi:"folder"`
	FolderId           string `pulumi:"folderId"`
	Id                 string `pulumi:"id"`
	LifecycleState     string `pulumi:"lifecycleState"`
	LookupOrganization *bool  `pulumi:"lookupOrganization"`
	Name               string `pulumi:"name"`
	Organization       string `pulumi:"organization"`
	Parent             string `pulumi:"parent"`
}

func LookupFolderOutput(ctx *pulumi.Context, args LookupFolderOutputArgs, opts ...pulumi.InvokeOption) LookupFolderResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupFolderResultOutput, error) {
			args := v.(LookupFolderArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupFolderResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getFolder:getFolder", args, LookupFolderResultOutput{}, options).(LookupFolderResultOutput), nil
		}).(LookupFolderResultOutput)
}

// A collection of arguments for invoking getFolder.
type LookupFolderOutputArgs struct {
	Folder             pulumi.StringInput    `pulumi:"folder"`
	Id                 pulumi.StringPtrInput `pulumi:"id"`
	LookupOrganization pulumi.BoolPtrInput   `pulumi:"lookupOrganization"`
}

func (LookupFolderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderArgs)(nil)).Elem()
}

// A collection of values returned by getFolder.
type LookupFolderResultOutput struct{ *pulumi.OutputState }

func (LookupFolderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFolderResult)(nil)).Elem()
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutput() LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) ToLookupFolderResultOutputWithContext(ctx context.Context) LookupFolderResultOutput {
	return o
}

func (o LookupFolderResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFolderResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupFolderResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Folder }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) FolderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.FolderId }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.LifecycleState }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) LookupOrganization() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupFolderResult) *bool { return v.LookupOrganization }).(pulumi.BoolPtrOutput)
}

func (o LookupFolderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Organization }).(pulumi.StringOutput)
}

func (o LookupFolderResultOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFolderResult) string { return v.Parent }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFolderResultOutput{})
}
