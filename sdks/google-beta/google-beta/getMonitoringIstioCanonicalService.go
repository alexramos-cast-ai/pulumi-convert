// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitoringIstioCanonicalService(ctx *pulumi.Context, args *GetMonitoringIstioCanonicalServiceArgs, opts ...pulumi.InvokeOption) (*GetMonitoringIstioCanonicalServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetMonitoringIstioCanonicalServiceResult
	err = ctx.InvokePackage("google-beta:index/getMonitoringIstioCanonicalService:getMonitoringIstioCanonicalService", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMonitoringIstioCanonicalService.
type GetMonitoringIstioCanonicalServiceArgs struct {
	CanonicalService          string  `pulumi:"canonicalService"`
	CanonicalServiceNamespace string  `pulumi:"canonicalServiceNamespace"`
	Id                        *string `pulumi:"id"`
	MeshUid                   string  `pulumi:"meshUid"`
	Project                   *string `pulumi:"project"`
}

// A collection of values returned by getMonitoringIstioCanonicalService.
type GetMonitoringIstioCanonicalServiceResult struct {
	CanonicalService          string                                        `pulumi:"canonicalService"`
	CanonicalServiceNamespace string                                        `pulumi:"canonicalServiceNamespace"`
	DisplayName               string                                        `pulumi:"displayName"`
	Id                        string                                        `pulumi:"id"`
	MeshUid                   string                                        `pulumi:"meshUid"`
	Name                      string                                        `pulumi:"name"`
	Project                   *string                                       `pulumi:"project"`
	ServiceId                 string                                        `pulumi:"serviceId"`
	Telemetries               []GetMonitoringIstioCanonicalServiceTelemetry `pulumi:"telemetries"`
	UserLabels                map[string]string                             `pulumi:"userLabels"`
}

func GetMonitoringIstioCanonicalServiceOutput(ctx *pulumi.Context, args GetMonitoringIstioCanonicalServiceOutputArgs, opts ...pulumi.InvokeOption) GetMonitoringIstioCanonicalServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetMonitoringIstioCanonicalServiceResultOutput, error) {
			args := v.(GetMonitoringIstioCanonicalServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetMonitoringIstioCanonicalServiceResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getMonitoringIstioCanonicalService:getMonitoringIstioCanonicalService", args, GetMonitoringIstioCanonicalServiceResultOutput{}, options).(GetMonitoringIstioCanonicalServiceResultOutput), nil
		}).(GetMonitoringIstioCanonicalServiceResultOutput)
}

// A collection of arguments for invoking getMonitoringIstioCanonicalService.
type GetMonitoringIstioCanonicalServiceOutputArgs struct {
	CanonicalService          pulumi.StringInput    `pulumi:"canonicalService"`
	CanonicalServiceNamespace pulumi.StringInput    `pulumi:"canonicalServiceNamespace"`
	Id                        pulumi.StringPtrInput `pulumi:"id"`
	MeshUid                   pulumi.StringInput    `pulumi:"meshUid"`
	Project                   pulumi.StringPtrInput `pulumi:"project"`
}

func (GetMonitoringIstioCanonicalServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringIstioCanonicalServiceArgs)(nil)).Elem()
}

// A collection of values returned by getMonitoringIstioCanonicalService.
type GetMonitoringIstioCanonicalServiceResultOutput struct{ *pulumi.OutputState }

func (GetMonitoringIstioCanonicalServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitoringIstioCanonicalServiceResult)(nil)).Elem()
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) ToGetMonitoringIstioCanonicalServiceResultOutput() GetMonitoringIstioCanonicalServiceResultOutput {
	return o
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) ToGetMonitoringIstioCanonicalServiceResultOutputWithContext(ctx context.Context) GetMonitoringIstioCanonicalServiceResultOutput {
	return o
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) CanonicalService() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.CanonicalService }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) CanonicalServiceNamespace() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.CanonicalServiceNamespace }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) MeshUid() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.MeshUid }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) Telemetries() GetMonitoringIstioCanonicalServiceTelemetryArrayOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) []GetMonitoringIstioCanonicalServiceTelemetry {
		return v.Telemetries
	}).(GetMonitoringIstioCanonicalServiceTelemetryArrayOutput)
}

func (o GetMonitoringIstioCanonicalServiceResultOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetMonitoringIstioCanonicalServiceResult) map[string]string { return v.UserLabels }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitoringIstioCanonicalServiceResultOutput{})
}
