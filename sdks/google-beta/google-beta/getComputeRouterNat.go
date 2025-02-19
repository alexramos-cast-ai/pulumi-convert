// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeRouterNat(ctx *pulumi.Context, args *LookupComputeRouterNatArgs, opts ...pulumi.InvokeOption) (*LookupComputeRouterNatResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeRouterNatResult
	err = ctx.InvokePackage("google-beta:index/getComputeRouterNat:getComputeRouterNat", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeRouterNat.
type LookupComputeRouterNatArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	Router  string  `pulumi:"router"`
}

// A collection of values returned by getComputeRouterNat.
type LookupComputeRouterNatResult struct {
	AutoNetworkTier                  string                          `pulumi:"autoNetworkTier"`
	DrainNatIps                      []string                        `pulumi:"drainNatIps"`
	EnableDynamicPortAllocation      bool                            `pulumi:"enableDynamicPortAllocation"`
	EnableEndpointIndependentMapping bool                            `pulumi:"enableEndpointIndependentMapping"`
	EndpointTypes                    []string                        `pulumi:"endpointTypes"`
	IcmpIdleTimeoutSec               float64                         `pulumi:"icmpIdleTimeoutSec"`
	Id                               string                          `pulumi:"id"`
	InitialNatIps                    []string                        `pulumi:"initialNatIps"`
	LogConfigs                       []GetComputeRouterNatLogConfig  `pulumi:"logConfigs"`
	MaxPortsPerVm                    float64                         `pulumi:"maxPortsPerVm"`
	MinPortsPerVm                    float64                         `pulumi:"minPortsPerVm"`
	Name                             string                          `pulumi:"name"`
	NatIpAllocateOption              string                          `pulumi:"natIpAllocateOption"`
	NatIps                           []string                        `pulumi:"natIps"`
	Project                          *string                         `pulumi:"project"`
	Region                           *string                         `pulumi:"region"`
	Router                           string                          `pulumi:"router"`
	Rules                            []GetComputeRouterNatRule       `pulumi:"rules"`
	SourceSubnetworkIpRangesToNat    string                          `pulumi:"sourceSubnetworkIpRangesToNat"`
	Subnetworks                      []GetComputeRouterNatSubnetwork `pulumi:"subnetworks"`
	TcpEstablishedIdleTimeoutSec     float64                         `pulumi:"tcpEstablishedIdleTimeoutSec"`
	TcpTimeWaitTimeoutSec            float64                         `pulumi:"tcpTimeWaitTimeoutSec"`
	TcpTransitoryIdleTimeoutSec      float64                         `pulumi:"tcpTransitoryIdleTimeoutSec"`
	Type                             string                          `pulumi:"type"`
	UdpIdleTimeoutSec                float64                         `pulumi:"udpIdleTimeoutSec"`
}

func LookupComputeRouterNatOutput(ctx *pulumi.Context, args LookupComputeRouterNatOutputArgs, opts ...pulumi.InvokeOption) LookupComputeRouterNatResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeRouterNatResultOutput, error) {
			args := v.(LookupComputeRouterNatArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeRouterNatResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeRouterNat:getComputeRouterNat", args, LookupComputeRouterNatResultOutput{}, options).(LookupComputeRouterNatResultOutput), nil
		}).(LookupComputeRouterNatResultOutput)
}

// A collection of arguments for invoking getComputeRouterNat.
type LookupComputeRouterNatOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
	Router  pulumi.StringInput    `pulumi:"router"`
}

func (LookupComputeRouterNatOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRouterNatArgs)(nil)).Elem()
}

// A collection of values returned by getComputeRouterNat.
type LookupComputeRouterNatResultOutput struct{ *pulumi.OutputState }

func (LookupComputeRouterNatResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRouterNatResult)(nil)).Elem()
}

func (o LookupComputeRouterNatResultOutput) ToLookupComputeRouterNatResultOutput() LookupComputeRouterNatResultOutput {
	return o
}

func (o LookupComputeRouterNatResultOutput) ToLookupComputeRouterNatResultOutputWithContext(ctx context.Context) LookupComputeRouterNatResultOutput {
	return o
}

func (o LookupComputeRouterNatResultOutput) AutoNetworkTier() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.AutoNetworkTier }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) DrainNatIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []string { return v.DrainNatIps }).(pulumi.StringArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) EnableDynamicPortAllocation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) bool { return v.EnableDynamicPortAllocation }).(pulumi.BoolOutput)
}

func (o LookupComputeRouterNatResultOutput) EnableEndpointIndependentMapping() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) bool { return v.EnableEndpointIndependentMapping }).(pulumi.BoolOutput)
}

func (o LookupComputeRouterNatResultOutput) EndpointTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []string { return v.EndpointTypes }).(pulumi.StringArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) IcmpIdleTimeoutSec() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.IcmpIdleTimeoutSec }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) InitialNatIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []string { return v.InitialNatIps }).(pulumi.StringArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) LogConfigs() GetComputeRouterNatLogConfigArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []GetComputeRouterNatLogConfig { return v.LogConfigs }).(GetComputeRouterNatLogConfigArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) MaxPortsPerVm() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.MaxPortsPerVm }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) MinPortsPerVm() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.MinPortsPerVm }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) NatIpAllocateOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.NatIpAllocateOption }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) NatIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []string { return v.NatIps }).(pulumi.StringArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRouterNatResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRouterNatResultOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.Router }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) Rules() GetComputeRouterNatRuleArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []GetComputeRouterNatRule { return v.Rules }).(GetComputeRouterNatRuleArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) SourceSubnetworkIpRangesToNat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.SourceSubnetworkIpRangesToNat }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) Subnetworks() GetComputeRouterNatSubnetworkArrayOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) []GetComputeRouterNatSubnetwork { return v.Subnetworks }).(GetComputeRouterNatSubnetworkArrayOutput)
}

func (o LookupComputeRouterNatResultOutput) TcpEstablishedIdleTimeoutSec() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.TcpEstablishedIdleTimeoutSec }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) TcpTimeWaitTimeoutSec() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.TcpTimeWaitTimeoutSec }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) TcpTransitoryIdleTimeoutSec() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.TcpTransitoryIdleTimeoutSec }).(pulumi.Float64Output)
}

func (o LookupComputeRouterNatResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRouterNatResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupComputeRouterNatResultOutput) UdpIdleTimeoutSec() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRouterNatResult) float64 { return v.UdpIdleTimeoutSec }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeRouterNatResultOutput{})
}
