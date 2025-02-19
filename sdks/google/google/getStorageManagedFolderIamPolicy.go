// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupStorageManagedFolderIamPolicy(ctx *pulumi.Context, args *LookupStorageManagedFolderIamPolicyArgs, opts ...pulumi.InvokeOption) (*LookupStorageManagedFolderIamPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupStorageManagedFolderIamPolicyResult
	err = ctx.InvokePackage("google:index/getStorageManagedFolderIamPolicy:getStorageManagedFolderIamPolicy", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageManagedFolderIamPolicy.
type LookupStorageManagedFolderIamPolicyArgs struct {
	Bucket        string  `pulumi:"bucket"`
	Id            *string `pulumi:"id"`
	ManagedFolder string  `pulumi:"managedFolder"`
}

// A collection of values returned by getStorageManagedFolderIamPolicy.
type LookupStorageManagedFolderIamPolicyResult struct {
	Bucket        string `pulumi:"bucket"`
	Etag          string `pulumi:"etag"`
	Id            string `pulumi:"id"`
	ManagedFolder string `pulumi:"managedFolder"`
	PolicyData    string `pulumi:"policyData"`
}

func LookupStorageManagedFolderIamPolicyOutput(ctx *pulumi.Context, args LookupStorageManagedFolderIamPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupStorageManagedFolderIamPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStorageManagedFolderIamPolicyResultOutput, error) {
			args := v.(LookupStorageManagedFolderIamPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupStorageManagedFolderIamPolicyResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getStorageManagedFolderIamPolicy:getStorageManagedFolderIamPolicy", args, LookupStorageManagedFolderIamPolicyResultOutput{}, options).(LookupStorageManagedFolderIamPolicyResultOutput), nil
		}).(LookupStorageManagedFolderIamPolicyResultOutput)
}

// A collection of arguments for invoking getStorageManagedFolderIamPolicy.
type LookupStorageManagedFolderIamPolicyOutputArgs struct {
	Bucket        pulumi.StringInput    `pulumi:"bucket"`
	Id            pulumi.StringPtrInput `pulumi:"id"`
	ManagedFolder pulumi.StringInput    `pulumi:"managedFolder"`
}

func (LookupStorageManagedFolderIamPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageManagedFolderIamPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getStorageManagedFolderIamPolicy.
type LookupStorageManagedFolderIamPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupStorageManagedFolderIamPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageManagedFolderIamPolicyResult)(nil)).Elem()
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) ToLookupStorageManagedFolderIamPolicyResultOutput() LookupStorageManagedFolderIamPolicyResultOutput {
	return o
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) ToLookupStorageManagedFolderIamPolicyResultOutputWithContext(ctx context.Context) LookupStorageManagedFolderIamPolicyResultOutput {
	return o
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageManagedFolderIamPolicyResult) string { return v.Bucket }).(pulumi.StringOutput)
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageManagedFolderIamPolicyResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageManagedFolderIamPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) ManagedFolder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageManagedFolderIamPolicyResult) string { return v.ManagedFolder }).(pulumi.StringOutput)
}

func (o LookupStorageManagedFolderIamPolicyResultOutput) PolicyData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageManagedFolderIamPolicyResult) string { return v.PolicyData }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageManagedFolderIamPolicyResultOutput{})
}
