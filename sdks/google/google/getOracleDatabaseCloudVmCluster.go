// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOracleDatabaseCloudVmCluster(ctx *pulumi.Context, args *LookupOracleDatabaseCloudVmClusterArgs, opts ...pulumi.InvokeOption) (*LookupOracleDatabaseCloudVmClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupOracleDatabaseCloudVmClusterResult
	err = ctx.InvokePackage("google:index/getOracleDatabaseCloudVmCluster:getOracleDatabaseCloudVmCluster", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOracleDatabaseCloudVmCluster.
type LookupOracleDatabaseCloudVmClusterArgs struct {
	CloudVmClusterId string  `pulumi:"cloudVmClusterId"`
	Id               *string `pulumi:"id"`
	Location         string  `pulumi:"location"`
	Project          *string `pulumi:"project"`
}

// A collection of values returned by getOracleDatabaseCloudVmCluster.
type LookupOracleDatabaseCloudVmClusterResult struct {
	BackupSubnetCidr      string                                    `pulumi:"backupSubnetCidr"`
	Cidr                  string                                    `pulumi:"cidr"`
	CloudVmClusterId      string                                    `pulumi:"cloudVmClusterId"`
	CreateTime            string                                    `pulumi:"createTime"`
	DeletionProtection    bool                                      `pulumi:"deletionProtection"`
	DisplayName           string                                    `pulumi:"displayName"`
	EffectiveLabels       map[string]string                         `pulumi:"effectiveLabels"`
	ExadataInfrastructure string                                    `pulumi:"exadataInfrastructure"`
	GcpOracleZone         string                                    `pulumi:"gcpOracleZone"`
	Id                    string                                    `pulumi:"id"`
	Labels                map[string]string                         `pulumi:"labels"`
	Location              string                                    `pulumi:"location"`
	Name                  string                                    `pulumi:"name"`
	Network               string                                    `pulumi:"network"`
	Project               *string                                   `pulumi:"project"`
	Properties            []GetOracleDatabaseCloudVmClusterProperty `pulumi:"properties"`
	TerraformLabels       map[string]string                         `pulumi:"terraformLabels"`
}

func LookupOracleDatabaseCloudVmClusterOutput(ctx *pulumi.Context, args LookupOracleDatabaseCloudVmClusterOutputArgs, opts ...pulumi.InvokeOption) LookupOracleDatabaseCloudVmClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOracleDatabaseCloudVmClusterResultOutput, error) {
			args := v.(LookupOracleDatabaseCloudVmClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupOracleDatabaseCloudVmClusterResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getOracleDatabaseCloudVmCluster:getOracleDatabaseCloudVmCluster", args, LookupOracleDatabaseCloudVmClusterResultOutput{}, options).(LookupOracleDatabaseCloudVmClusterResultOutput), nil
		}).(LookupOracleDatabaseCloudVmClusterResultOutput)
}

// A collection of arguments for invoking getOracleDatabaseCloudVmCluster.
type LookupOracleDatabaseCloudVmClusterOutputArgs struct {
	CloudVmClusterId pulumi.StringInput    `pulumi:"cloudVmClusterId"`
	Id               pulumi.StringPtrInput `pulumi:"id"`
	Location         pulumi.StringInput    `pulumi:"location"`
	Project          pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupOracleDatabaseCloudVmClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOracleDatabaseCloudVmClusterArgs)(nil)).Elem()
}

// A collection of values returned by getOracleDatabaseCloudVmCluster.
type LookupOracleDatabaseCloudVmClusterResultOutput struct{ *pulumi.OutputState }

func (LookupOracleDatabaseCloudVmClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOracleDatabaseCloudVmClusterResult)(nil)).Elem()
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) ToLookupOracleDatabaseCloudVmClusterResultOutput() LookupOracleDatabaseCloudVmClusterResultOutput {
	return o
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) ToLookupOracleDatabaseCloudVmClusterResultOutputWithContext(ctx context.Context) LookupOracleDatabaseCloudVmClusterResultOutput {
	return o
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) BackupSubnetCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.BackupSubnetCidr }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Cidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.Cidr }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) CloudVmClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.CloudVmClusterId }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) DeletionProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) bool { return v.DeletionProtection }).(pulumi.BoolOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) ExadataInfrastructure() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.ExadataInfrastructure }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) GcpOracleZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.GcpOracleZone }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) Properties() GetOracleDatabaseCloudVmClusterPropertyArrayOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) []GetOracleDatabaseCloudVmClusterProperty {
		return v.Properties
	}).(GetOracleDatabaseCloudVmClusterPropertyArrayOutput)
}

func (o LookupOracleDatabaseCloudVmClusterResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOracleDatabaseCloudVmClusterResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOracleDatabaseCloudVmClusterResultOutput{})
}
