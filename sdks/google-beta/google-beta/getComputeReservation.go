// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupComputeReservation(ctx *pulumi.Context, args *LookupComputeReservationArgs, opts ...pulumi.InvokeOption) (*LookupComputeReservationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupComputeReservationResult
	err = ctx.InvokePackage("google-beta:index/getComputeReservation:getComputeReservation", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeReservation.
type LookupComputeReservationArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Zone    string  `pulumi:"zone"`
}

// A collection of values returned by getComputeReservation.
type LookupComputeReservationResult struct {
	Commitment                  string                                     `pulumi:"commitment"`
	CreationTimestamp           string                                     `pulumi:"creationTimestamp"`
	Description                 string                                     `pulumi:"description"`
	Id                          string                                     `pulumi:"id"`
	Name                        string                                     `pulumi:"name"`
	Project                     *string                                    `pulumi:"project"`
	SelfLink                    string                                     `pulumi:"selfLink"`
	ShareSettings               []GetComputeReservationShareSetting        `pulumi:"shareSettings"`
	SpecificReservationRequired bool                                       `pulumi:"specificReservationRequired"`
	SpecificReservations        []GetComputeReservationSpecificReservation `pulumi:"specificReservations"`
	Status                      string                                     `pulumi:"status"`
	Zone                        string                                     `pulumi:"zone"`
}

func LookupComputeReservationOutput(ctx *pulumi.Context, args LookupComputeReservationOutputArgs, opts ...pulumi.InvokeOption) LookupComputeReservationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupComputeReservationResultOutput, error) {
			args := v.(LookupComputeReservationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupComputeReservationResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google-beta:index/getComputeReservation:getComputeReservation", args, LookupComputeReservationResultOutput{}, options).(LookupComputeReservationResultOutput), nil
		}).(LookupComputeReservationResultOutput)
}

// A collection of arguments for invoking getComputeReservation.
type LookupComputeReservationOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Zone    pulumi.StringInput    `pulumi:"zone"`
}

func (LookupComputeReservationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeReservationArgs)(nil)).Elem()
}

// A collection of values returned by getComputeReservation.
type LookupComputeReservationResultOutput struct{ *pulumi.OutputState }

func (LookupComputeReservationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeReservationResult)(nil)).Elem()
}

func (o LookupComputeReservationResultOutput) ToLookupComputeReservationResultOutput() LookupComputeReservationResultOutput {
	return o
}

func (o LookupComputeReservationResultOutput) ToLookupComputeReservationResultOutputWithContext(ctx context.Context) LookupComputeReservationResultOutput {
	return o
}

func (o LookupComputeReservationResultOutput) Commitment() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Commitment }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.CreationTimestamp }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupComputeReservationResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) ShareSettings() GetComputeReservationShareSettingArrayOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) []GetComputeReservationShareSetting { return v.ShareSettings }).(GetComputeReservationShareSettingArrayOutput)
}

func (o LookupComputeReservationResultOutput) SpecificReservationRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) bool { return v.SpecificReservationRequired }).(pulumi.BoolOutput)
}

func (o LookupComputeReservationResultOutput) SpecificReservations() GetComputeReservationSpecificReservationArrayOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) []GetComputeReservationSpecificReservation {
		return v.SpecificReservations
	}).(GetComputeReservationSpecificReservationArrayOutput)
}

func (o LookupComputeReservationResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupComputeReservationResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeReservationResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeReservationResultOutput{})
}
