// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetClientConfig(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetClientConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetClientConfigResult
	err = ctx.InvokePackage("google-beta:index/getClientConfig:getClientConfig", nil, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getClientConfig.
type GetClientConfigResult struct {
	AccessToken   string            `pulumi:"accessToken"`
	DefaultLabels map[string]string `pulumi:"defaultLabels"`
	Id            string            `pulumi:"id"`
	Project       string            `pulumi:"project"`
	Region        string            `pulumi:"region"`
	Zone          string            `pulumi:"zone"`
}

func GetClientConfigOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetClientConfigResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetClientConfigResultOutput, error) {
		options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
		ref, err := internal.PkgGetPackageRef(ctx)
		if err != nil {
			return GetClientConfigResultOutput{}, err
		}
		options.PackageRef = ref
		return ctx.InvokeOutput("google-beta:index/getClientConfig:getClientConfig", nil, GetClientConfigResultOutput{}, options).(GetClientConfigResultOutput), nil
	}).(GetClientConfigResultOutput)
}

// A collection of values returned by getClientConfig.
type GetClientConfigResultOutput struct{ *pulumi.OutputState }

func (GetClientConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetClientConfigResult)(nil)).Elem()
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutput() GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) ToGetClientConfigResultOutputWithContext(ctx context.Context) GetClientConfigResultOutput {
	return o
}

func (o GetClientConfigResultOutput) AccessToken() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.AccessToken }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) DefaultLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetClientConfigResult) map[string]string { return v.DefaultLabels }).(pulumi.StringMapOutput)
}

func (o GetClientConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Project }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Region }).(pulumi.StringOutput)
}

func (o GetClientConfigResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetClientConfigResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetClientConfigResultOutput{})
}
