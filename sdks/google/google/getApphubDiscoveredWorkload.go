// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetApphubDiscoveredWorkload(ctx *pulumi.Context, args *GetApphubDiscoveredWorkloadArgs, opts ...pulumi.InvokeOption) (*GetApphubDiscoveredWorkloadResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetApphubDiscoveredWorkloadResult
	err = ctx.InvokePackage("google:index/getApphubDiscoveredWorkload:getApphubDiscoveredWorkload", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApphubDiscoveredWorkload.
type GetApphubDiscoveredWorkloadArgs struct {
	Id          *string `pulumi:"id"`
	Location    string  `pulumi:"location"`
	Project     *string `pulumi:"project"`
	WorkloadUri string  `pulumi:"workloadUri"`
}

// A collection of values returned by getApphubDiscoveredWorkload.
type GetApphubDiscoveredWorkloadResult struct {
	Id                 string                                         `pulumi:"id"`
	Location           string                                         `pulumi:"location"`
	Name               string                                         `pulumi:"name"`
	Project            *string                                        `pulumi:"project"`
	WorkloadProperties []GetApphubDiscoveredWorkloadWorkloadProperty  `pulumi:"workloadProperties"`
	WorkloadReferences []GetApphubDiscoveredWorkloadWorkloadReference `pulumi:"workloadReferences"`
	WorkloadUri        string                                         `pulumi:"workloadUri"`
}

func GetApphubDiscoveredWorkloadOutput(ctx *pulumi.Context, args GetApphubDiscoveredWorkloadOutputArgs, opts ...pulumi.InvokeOption) GetApphubDiscoveredWorkloadResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetApphubDiscoveredWorkloadResultOutput, error) {
			args := v.(GetApphubDiscoveredWorkloadArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetApphubDiscoveredWorkloadResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getApphubDiscoveredWorkload:getApphubDiscoveredWorkload", args, GetApphubDiscoveredWorkloadResultOutput{}, options).(GetApphubDiscoveredWorkloadResultOutput), nil
		}).(GetApphubDiscoveredWorkloadResultOutput)
}

// A collection of arguments for invoking getApphubDiscoveredWorkload.
type GetApphubDiscoveredWorkloadOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	Location    pulumi.StringInput    `pulumi:"location"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	WorkloadUri pulumi.StringInput    `pulumi:"workloadUri"`
}

func (GetApphubDiscoveredWorkloadOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApphubDiscoveredWorkloadArgs)(nil)).Elem()
}

// A collection of values returned by getApphubDiscoveredWorkload.
type GetApphubDiscoveredWorkloadResultOutput struct{ *pulumi.OutputState }

func (GetApphubDiscoveredWorkloadResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetApphubDiscoveredWorkloadResult)(nil)).Elem()
}

func (o GetApphubDiscoveredWorkloadResultOutput) ToGetApphubDiscoveredWorkloadResultOutput() GetApphubDiscoveredWorkloadResultOutput {
	return o
}

func (o GetApphubDiscoveredWorkloadResultOutput) ToGetApphubDiscoveredWorkloadResultOutputWithContext(ctx context.Context) GetApphubDiscoveredWorkloadResultOutput {
	return o
}

func (o GetApphubDiscoveredWorkloadResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) WorkloadProperties() GetApphubDiscoveredWorkloadWorkloadPropertyArrayOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) []GetApphubDiscoveredWorkloadWorkloadProperty {
		return v.WorkloadProperties
	}).(GetApphubDiscoveredWorkloadWorkloadPropertyArrayOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) WorkloadReferences() GetApphubDiscoveredWorkloadWorkloadReferenceArrayOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) []GetApphubDiscoveredWorkloadWorkloadReference {
		return v.WorkloadReferences
	}).(GetApphubDiscoveredWorkloadWorkloadReferenceArrayOutput)
}

func (o GetApphubDiscoveredWorkloadResultOutput) WorkloadUri() pulumi.StringOutput {
	return o.ApplyT(func(v GetApphubDiscoveredWorkloadResult) string { return v.WorkloadUri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetApphubDiscoveredWorkloadResultOutput{})
}
