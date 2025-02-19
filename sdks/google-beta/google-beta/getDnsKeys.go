// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetDnsKeys(ctx *pulumi.Context, args *GetDnsKeysArgs, opts ...pulumi.InvokeOption) (*GetDnsKeysResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetDnsKeysResult
	err = ctx.InvokePackage("google-beta:index/getDnsKeys:getDnsKeys", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsKeys.
type GetDnsKeysArgs struct {
	Id          *string `pulumi:"id"`
	ManagedZone string  `pulumi:"managedZone"`
	Project     *string `pulumi:"project"`
}

// A collection of values returned by getDnsKeys.
type GetDnsKeysResult struct {
	Id              string                     `pulumi:"id"`
	KeySigningKeys  []GetDnsKeysKeySigningKey  `pulumi:"keySigningKeys"`
	ManagedZone     string                     `pulumi:"managedZone"`
	Project         string                     `pulumi:"project"`
	ZoneSigningKeys []GetDnsKeysZoneSigningKey `pulumi:"zoneSigningKeys"`
}

func GetDnsKeysOutput(ctx *pulumi.Context, args GetDnsKeysOutputArgs, opts ...pulumi.InvokeOption) GetDnsKeysResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetDnsKeysResultOutput, error) {
			args := v.(GetDnsKeysArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetDnsKeysResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getDnsKeys:getDnsKeys", args, GetDnsKeysResultOutput{}, options).(GetDnsKeysResultOutput), nil
		}).(GetDnsKeysResultOutput)
}

// A collection of arguments for invoking getDnsKeys.
type GetDnsKeysOutputArgs struct {
	Id          pulumi.StringPtrInput `pulumi:"id"`
	ManagedZone pulumi.StringInput    `pulumi:"managedZone"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
}

func (GetDnsKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsKeysArgs)(nil)).Elem()
}

// A collection of values returned by getDnsKeys.
type GetDnsKeysResultOutput struct{ *pulumi.OutputState }

func (GetDnsKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDnsKeysResult)(nil)).Elem()
}

func (o GetDnsKeysResultOutput) ToGetDnsKeysResultOutput() GetDnsKeysResultOutput {
	return o
}

func (o GetDnsKeysResultOutput) ToGetDnsKeysResultOutputWithContext(ctx context.Context) GetDnsKeysResultOutput {
	return o
}

func (o GetDnsKeysResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsKeysResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDnsKeysResultOutput) KeySigningKeys() GetDnsKeysKeySigningKeyArrayOutput {
	return o.ApplyT(func(v GetDnsKeysResult) []GetDnsKeysKeySigningKey { return v.KeySigningKeys }).(GetDnsKeysKeySigningKeyArrayOutput)
}

func (o GetDnsKeysResultOutput) ManagedZone() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsKeysResult) string { return v.ManagedZone }).(pulumi.StringOutput)
}

func (o GetDnsKeysResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetDnsKeysResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetDnsKeysResultOutput) ZoneSigningKeys() GetDnsKeysZoneSigningKeyArrayOutput {
	return o.ApplyT(func(v GetDnsKeysResult) []GetDnsKeysZoneSigningKey { return v.ZoneSigningKeys }).(GetDnsKeysZoneSigningKeyArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDnsKeysResultOutput{})
}
