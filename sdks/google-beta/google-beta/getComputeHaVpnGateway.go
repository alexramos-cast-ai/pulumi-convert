// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeHaVpnGateway(ctx *pulumi.Context, args *LookupComputeHaVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupComputeHaVpnGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeHaVpnGatewayResult
	err = ctx.InvokePackage("google-beta:index/getComputeHaVpnGateway:getComputeHaVpnGateway", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeHaVpnGateway.
type LookupComputeHaVpnGatewayArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComputeHaVpnGateway.
type LookupComputeHaVpnGatewayResult struct {
	Description      string                               `pulumi:"description"`
	GatewayIpVersion string                               `pulumi:"gatewayIpVersion"`
	Id               string                               `pulumi:"id"`
	Name             string                               `pulumi:"name"`
	Network          string                               `pulumi:"network"`
	Project          *string                              `pulumi:"project"`
	Region           *string                              `pulumi:"region"`
	SelfLink         string                               `pulumi:"selfLink"`
	StackType        string                               `pulumi:"stackType"`
	VpnInterfaces    []GetComputeHaVpnGatewayVpnInterface `pulumi:"vpnInterfaces"`
}

func LookupComputeHaVpnGatewayOutput(ctx *pulumi.Context, args LookupComputeHaVpnGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupComputeHaVpnGatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeHaVpnGatewayResultOutput, error) {
			args := v.(LookupComputeHaVpnGatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeHaVpnGatewayResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeHaVpnGateway:getComputeHaVpnGateway", args, LookupComputeHaVpnGatewayResultOutput{}, options).(LookupComputeHaVpnGatewayResultOutput), nil
		}).(LookupComputeHaVpnGatewayResultOutput)
}

// A collection of arguments for invoking getComputeHaVpnGateway.
type LookupComputeHaVpnGatewayOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComputeHaVpnGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeHaVpnGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getComputeHaVpnGateway.
type LookupComputeHaVpnGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupComputeHaVpnGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeHaVpnGatewayResult)(nil)).Elem()
}

func (o LookupComputeHaVpnGatewayResultOutput) ToLookupComputeHaVpnGatewayResultOutput() LookupComputeHaVpnGatewayResultOutput {
	return o
}

func (o LookupComputeHaVpnGatewayResultOutput) ToLookupComputeHaVpnGatewayResultOutputWithContext(ctx context.Context) LookupComputeHaVpnGatewayResultOutput {
	return o
}

func (o LookupComputeHaVpnGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) GatewayIpVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.GatewayIpVersion }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.Network }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) StackType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) string { return v.StackType }).(pulumi.StringOutput)
}

func (o LookupComputeHaVpnGatewayResultOutput) VpnInterfaces() GetComputeHaVpnGatewayVpnInterfaceArrayOutput {
	return o.ApplyT(func(v LookupComputeHaVpnGatewayResult) []GetComputeHaVpnGatewayVpnInterface { return v.VpnInterfaces }).(GetComputeHaVpnGatewayVpnInterfaceArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeHaVpnGatewayResultOutput{})
}
