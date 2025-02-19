// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetOracleDatabaseAutonomousDatabases(ctx *pulumi.Context, args *GetOracleDatabaseAutonomousDatabasesArgs, opts ...pulumi.InvokeOption) (*GetOracleDatabaseAutonomousDatabasesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv GetOracleDatabaseAutonomousDatabasesResult
	err = ctx.InvokePackage("google-beta:index/getOracleDatabaseAutonomousDatabases:getOracleDatabaseAutonomousDatabases", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOracleDatabaseAutonomousDatabases.
type GetOracleDatabaseAutonomousDatabasesArgs struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

// A collection of values returned by getOracleDatabaseAutonomousDatabases.
type GetOracleDatabaseAutonomousDatabasesResult struct {
	AutonomousDatabases []GetOracleDatabaseAutonomousDatabasesAutonomousDatabase `pulumi:"autonomousDatabases"`
	Id                  string                                                   `pulumi:"id"`
	Location            string                                                   `pulumi:"location"`
	Project             *string                                                  `pulumi:"project"`
}

func GetOracleDatabaseAutonomousDatabasesOutput(ctx *pulumi.Context, args GetOracleDatabaseAutonomousDatabasesOutputArgs, opts ...pulumi.InvokeOption) GetOracleDatabaseAutonomousDatabasesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetOracleDatabaseAutonomousDatabasesResultOutput, error) {
			args := v.(GetOracleDatabaseAutonomousDatabasesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return GetOracleDatabaseAutonomousDatabasesResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getOracleDatabaseAutonomousDatabases:getOracleDatabaseAutonomousDatabases", args, GetOracleDatabaseAutonomousDatabasesResultOutput{}, options).(GetOracleDatabaseAutonomousDatabasesResultOutput), nil
		}).(GetOracleDatabaseAutonomousDatabasesResultOutput)
}

// A collection of arguments for invoking getOracleDatabaseAutonomousDatabases.
type GetOracleDatabaseAutonomousDatabasesOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (GetOracleDatabaseAutonomousDatabasesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseAutonomousDatabasesArgs)(nil)).Elem()
}

// A collection of values returned by getOracleDatabaseAutonomousDatabases.
type GetOracleDatabaseAutonomousDatabasesResultOutput struct{ *pulumi.OutputState }

func (GetOracleDatabaseAutonomousDatabasesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetOracleDatabaseAutonomousDatabasesResult)(nil)).Elem()
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) ToGetOracleDatabaseAutonomousDatabasesResultOutput() GetOracleDatabaseAutonomousDatabasesResultOutput {
	return o
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) ToGetOracleDatabaseAutonomousDatabasesResultOutputWithContext(ctx context.Context) GetOracleDatabaseAutonomousDatabasesResultOutput {
	return o
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) AutonomousDatabases() GetOracleDatabaseAutonomousDatabasesAutonomousDatabaseArrayOutput {
	return o.ApplyT(func(v GetOracleDatabaseAutonomousDatabasesResult) []GetOracleDatabaseAutonomousDatabasesAutonomousDatabase {
		return v.AutonomousDatabases
	}).(GetOracleDatabaseAutonomousDatabasesAutonomousDatabaseArrayOutput)
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseAutonomousDatabasesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetOracleDatabaseAutonomousDatabasesResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetOracleDatabaseAutonomousDatabasesResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetOracleDatabaseAutonomousDatabasesResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetOracleDatabaseAutonomousDatabasesResultOutput{})
}
