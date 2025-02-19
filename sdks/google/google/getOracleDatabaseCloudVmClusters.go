// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOracleDatabaseCloudVmClusters(ctx *pulumi.Context, args *GetOracleDatabaseCloudVmClustersArgs, opts ...pulumi.InvokeOption) (*GetOracleDatabaseCloudVmClustersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOracleDatabaseCloudVmClustersResult
	err = ctx.InvokePackage("google:index/getOracleDatabaseCloudVmClusters:getOracleDatabaseCloudVmClusters", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOracleDatabaseCloudVmClusters.
type GetOracleDatabaseCloudVmClustersArgs struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getOracleDatabaseCloudVmClusters.
type GetOracleDatabaseCloudVmClustersResult struct {
	CloudVmClusters []GetOracleDatabaseCloudVmClustersCloudVmCluster `pulumi:"cloudVmClusters"`
	Id              string                                           `pulumi:"id"`
	Location        string                                           `pulumi:"location"`
	Project         *string                                          `pulumi:"project"`
}

func GetOracleDatabaseCloudVmClustersOutput(ctx *pulumi.Context, args GetOracleDatabaseCloudVmClustersOutputArgs, opts ...pulumi.InvokeOption) GetOracleDatabaseCloudVmClustersResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOracleDatabaseCloudVmClustersResultOutput, error) {
			args := v.(GetOracleDatabaseCloudVmClustersArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOracleDatabaseCloudVmClustersResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getOracleDatabaseCloudVmClusters:getOracleDatabaseCloudVmClusters", args, GetOracleDatabaseCloudVmClustersResultOutput{}, options).(GetOracleDatabaseCloudVmClustersResultOutput), nil
		}).(GetOracleDatabaseCloudVmClustersResultOutput)
}

// A collection of arguments for invoking getOracleDatabaseCloudVmClusters.
type GetOracleDatabaseCloudVmClustersOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (GetOracleDatabaseCloudVmClustersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseCloudVmClustersArgs)(nil)).Elem()
}

// A collection of values returned by getOracleDatabaseCloudVmClusters.
type GetOracleDatabaseCloudVmClustersResultOutput struct{ *pulumi.OutputState }

func (GetOracleDatabaseCloudVmClustersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseCloudVmClustersResult)(nil)).Elem()
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) ToGetOracleDatabaseCloudVmClustersResultOutput() GetOracleDatabaseCloudVmClustersResultOutput {
	return o
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) ToGetOracleDatabaseCloudVmClustersResultOutputWithContext(ctx context.Context) GetOracleDatabaseCloudVmClustersResultOutput {
	return o
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) CloudVmClusters() GetOracleDatabaseCloudVmClustersCloudVmClusterArrayOutput {
	return o.ApplyT(func(v GetOracleDatabaseCloudVmClustersResult) []GetOracleDatabaseCloudVmClustersCloudVmCluster {
		return v.CloudVmClusters
	}).(GetOracleDatabaseCloudVmClustersCloudVmClusterArrayOutput)
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseCloudVmClustersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseCloudVmClustersResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseCloudVmClustersResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOracleDatabaseCloudVmClustersResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOracleDatabaseCloudVmClustersResultOutput{})
}
