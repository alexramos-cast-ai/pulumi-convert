// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeRegionSslCertificate(ctx *pulumi.Context, args *LookupComputeRegionSslCertificateArgs, opts ...pulumi.InvokeOption) (*LookupComputeRegionSslCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeRegionSslCertificateResult
	err = ctx.InvokePackage("google:index/getComputeRegionSslCertificate:getComputeRegionSslCertificate", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeRegionSslCertificate.
type LookupComputeRegionSslCertificateArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getComputeRegionSslCertificate.
type LookupComputeRegionSslCertificateResult struct {
	Certificate       string  `pulumi:"certificate"`
	CertificateId     float64 `pulumi:"certificateId"`
	CreationTimestamp string  `pulumi:"creationTimestamp"`
	Description       string  `pulumi:"description"`
	ExpireTime        string  `pulumi:"expireTime"`
	Id                string  `pulumi:"id"`
	Name              string  `pulumi:"name"`
	NamePrefix        string  `pulumi:"namePrefix"`
	PrivateKey        string  `pulumi:"privateKey"`
	Project           *string `pulumi:"project"`
	Region            *string `pulumi:"region"`
	SelfLink          string  `pulumi:"selfLink"`
}

func LookupComputeRegionSslCertificateOutput(ctx *pulumi.Context, args LookupComputeRegionSslCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupComputeRegionSslCertificateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeRegionSslCertificateResultOutput, error) {
			args := v.(LookupComputeRegionSslCertificateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeRegionSslCertificateResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getComputeRegionSslCertificate:getComputeRegionSslCertificate", args, LookupComputeRegionSslCertificateResultOutput{}, options).(LookupComputeRegionSslCertificateResultOutput), nil
		}).(LookupComputeRegionSslCertificateResultOutput)
}

// A collection of arguments for invoking getComputeRegionSslCertificate.
type LookupComputeRegionSslCertificateOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupComputeRegionSslCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRegionSslCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getComputeRegionSslCertificate.
type LookupComputeRegionSslCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupComputeRegionSslCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeRegionSslCertificateResult)(nil)).Elem()
}

func (o LookupComputeRegionSslCertificateResultOutput) ToLookupComputeRegionSslCertificateResultOutput() LookupComputeRegionSslCertificateResultOutput {
	return o
}

func (o LookupComputeRegionSslCertificateResultOutput) ToLookupComputeRegionSslCertificateResultOutputWithContext(ctx context.Context) LookupComputeRegionSslCertificateResultOutput {
	return o
}

func (o LookupComputeRegionSslCertificateResultOutput) Certificate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.Certificate }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) CertificateId() pulumi.Float64Output {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) float64 { return v.CertificateId }).(pulumi.Float64Output)
}

func (o LookupComputeRegionSslCertificateResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.ExpireTime }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.NamePrefix }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) PrivateKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.PrivateKey }).(pulumi.StringOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupComputeRegionSslCertificateResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeRegionSslCertificateResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeRegionSslCertificateResultOutput{})
}
