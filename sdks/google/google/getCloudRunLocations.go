// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetCloudRunLocations(ctx *pulumi.Context, args *GetCloudRunLocationsArgs, opts ...pulumi.InvokeOption) (*GetCloudRunLocationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetCloudRunLocationsResult
	err = ctx.InvokePackage("google:index/getCloudRunLocations:getCloudRunLocations", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudRunLocations.
type GetCloudRunLocationsArgs struct {
	Id      *string `pulumi:"id"`
	Project *string `pulumi:"project"`
}

// A collection of values returned by getCloudRunLocations.
type GetCloudRunLocationsResult struct {
	Id        string   `pulumi:"id"`
	Locations []string `pulumi:"locations"`
	Project   string   `pulumi:"project"`
}

func GetCloudRunLocationsOutput(ctx *pulumi.Context, args GetCloudRunLocationsOutputArgs, opts ...pulumi.InvokeOption) GetCloudRunLocationsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetCloudRunLocationsResultOutput, error) {
			args := v.(GetCloudRunLocationsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetCloudRunLocationsResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getCloudRunLocations:getCloudRunLocations", args, GetCloudRunLocationsResultOutput{}, options).(GetCloudRunLocationsResultOutput), nil
		}).(GetCloudRunLocationsResultOutput)
}

// A collection of arguments for invoking getCloudRunLocations.
type GetCloudRunLocationsOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (GetCloudRunLocationsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudRunLocationsArgs)(nil)).Elem()
}

// A collection of values returned by getCloudRunLocations.
type GetCloudRunLocationsResultOutput struct{ *pulumi.OutputState }

func (GetCloudRunLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetCloudRunLocationsResult)(nil)).Elem()
}

func (o GetCloudRunLocationsResultOutput) ToGetCloudRunLocationsResultOutput() GetCloudRunLocationsResultOutput {
	return o
}

func (o GetCloudRunLocationsResultOutput) ToGetCloudRunLocationsResultOutputWithContext(ctx context.Context) GetCloudRunLocationsResultOutput {
	return o
}

func (o GetCloudRunLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudRunLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetCloudRunLocationsResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetCloudRunLocationsResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o GetCloudRunLocationsResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v GetCloudRunLocationsResult) string { return v.Project }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetCloudRunLocationsResultOutput{})
}
