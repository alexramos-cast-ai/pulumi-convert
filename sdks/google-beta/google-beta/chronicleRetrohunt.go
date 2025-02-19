// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ChronicleRetrohunt struct {
	pulumi.CustomResourceState

	ChronicleRetrohuntId pulumi.StringOutput `pulumi:"chronicleRetrohuntId"`
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ExecutionIntervals ChronicleRetrohuntExecutionIntervalArrayOutput `pulumi:"executionIntervals"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the retrohunt. Retrohunt is the child of a rule revision. {rule} in the format below is structured
	// as {rule_id@revision_id}. Format:
	// projects/{project}/locations/{location}/instances/{instance}/rules/{rule}/retrohunts/{retrohunt}
	Name pulumi.StringOutput `pulumi:"name"`
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ProcessInterval ChronicleRetrohuntProcessIntervalOutput `pulumi:"processInterval"`
	// Output only. Percent progress of the retrohunt towards completion, from 0.00 to 100.00.
	ProgressPercentage pulumi.Float64Output `pulumi:"progressPercentage"`
	Project            pulumi.StringOutput  `pulumi:"project"`
	// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
	Retrohunt pulumi.StringOutput `pulumi:"retrohunt"`
	// The Rule ID of the rule.
	Rule pulumi.StringOutput `pulumi:"rule"`
	// Output only. The state of the retrohunt. Possible values: RUNNING DONE CANCELLED FAILED
	State    pulumi.StringOutput                 `pulumi:"state"`
	Timeouts ChronicleRetrohuntTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewChronicleRetrohunt registers a new resource with the given unique name, arguments, and options.
func NewChronicleRetrohunt(ctx *pulumi.Context,
	name string, args *ChronicleRetrohuntArgs, opts ...pulumi.ResourceOption) (*ChronicleRetrohunt, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ProcessInterval == nil {
		return nil, errors.New("invalid value for required argument 'ProcessInterval'")
	}
	if args.Rule == nil {
		return nil, errors.New("invalid value for required argument 'Rule'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ChronicleRetrohunt
	err = ctx.RegisterPackageResource("google-beta:index/chronicleRetrohunt:ChronicleRetrohunt", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChronicleRetrohunt gets an existing ChronicleRetrohunt resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChronicleRetrohunt(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChronicleRetrohuntState, opts ...pulumi.ResourceOption) (*ChronicleRetrohunt, error) {
	var resource ChronicleRetrohunt
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/chronicleRetrohunt:ChronicleRetrohunt", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChronicleRetrohunt resources.
type chronicleRetrohuntState struct {
	ChronicleRetrohuntId *string `pulumi:"chronicleRetrohuntId"`
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ExecutionIntervals []ChronicleRetrohuntExecutionInterval `pulumi:"executionIntervals"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance *string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location *string `pulumi:"location"`
	// The resource name of the retrohunt. Retrohunt is the child of a rule revision. {rule} in the format below is structured
	// as {rule_id@revision_id}. Format:
	// projects/{project}/locations/{location}/instances/{instance}/rules/{rule}/retrohunts/{retrohunt}
	Name *string `pulumi:"name"`
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ProcessInterval *ChronicleRetrohuntProcessInterval `pulumi:"processInterval"`
	// Output only. Percent progress of the retrohunt towards completion, from 0.00 to 100.00.
	ProgressPercentage *float64 `pulumi:"progressPercentage"`
	Project            *string  `pulumi:"project"`
	// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
	Retrohunt *string `pulumi:"retrohunt"`
	// The Rule ID of the rule.
	Rule *string `pulumi:"rule"`
	// Output only. The state of the retrohunt. Possible values: RUNNING DONE CANCELLED FAILED
	State    *string                     `pulumi:"state"`
	Timeouts *ChronicleRetrohuntTimeouts `pulumi:"timeouts"`
}

type ChronicleRetrohuntState struct {
	ChronicleRetrohuntId pulumi.StringPtrInput
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ExecutionIntervals ChronicleRetrohuntExecutionIntervalArrayInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringPtrInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringPtrInput
	// The resource name of the retrohunt. Retrohunt is the child of a rule revision. {rule} in the format below is structured
	// as {rule_id@revision_id}. Format:
	// projects/{project}/locations/{location}/instances/{instance}/rules/{rule}/retrohunts/{retrohunt}
	Name pulumi.StringPtrInput
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ProcessInterval ChronicleRetrohuntProcessIntervalPtrInput
	// Output only. Percent progress of the retrohunt towards completion, from 0.00 to 100.00.
	ProgressPercentage pulumi.Float64PtrInput
	Project            pulumi.StringPtrInput
	// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
	Retrohunt pulumi.StringPtrInput
	// The Rule ID of the rule.
	Rule pulumi.StringPtrInput
	// Output only. The state of the retrohunt. Possible values: RUNNING DONE CANCELLED FAILED
	State    pulumi.StringPtrInput
	Timeouts ChronicleRetrohuntTimeoutsPtrInput
}

func (ChronicleRetrohuntState) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleRetrohuntState)(nil)).Elem()
}

type chronicleRetrohuntArgs struct {
	ChronicleRetrohuntId *string `pulumi:"chronicleRetrohuntId"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location string `pulumi:"location"`
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ProcessInterval ChronicleRetrohuntProcessInterval `pulumi:"processInterval"`
	Project         *string                           `pulumi:"project"`
	// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
	Retrohunt *string `pulumi:"retrohunt"`
	// The Rule ID of the rule.
	Rule     string                      `pulumi:"rule"`
	Timeouts *ChronicleRetrohuntTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ChronicleRetrohunt resource.
type ChronicleRetrohuntArgs struct {
	ChronicleRetrohuntId pulumi.StringPtrInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringInput
	// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
	// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
	// and end are unspecified, the interval matches any time.
	ProcessInterval ChronicleRetrohuntProcessIntervalInput
	Project         pulumi.StringPtrInput
	// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
	Retrohunt pulumi.StringPtrInput
	// The Rule ID of the rule.
	Rule     pulumi.StringInput
	Timeouts ChronicleRetrohuntTimeoutsPtrInput
}

func (ChronicleRetrohuntArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleRetrohuntArgs)(nil)).Elem()
}

type ChronicleRetrohuntInput interface {
	pulumi.Input

	ToChronicleRetrohuntOutput() ChronicleRetrohuntOutput
	ToChronicleRetrohuntOutputWithContext(ctx context.Context) ChronicleRetrohuntOutput
}

func (*ChronicleRetrohunt) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleRetrohunt)(nil)).Elem()
}

func (i *ChronicleRetrohunt) ToChronicleRetrohuntOutput() ChronicleRetrohuntOutput {
	return i.ToChronicleRetrohuntOutputWithContext(context.Background())
}

func (i *ChronicleRetrohunt) ToChronicleRetrohuntOutputWithContext(ctx context.Context) ChronicleRetrohuntOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChronicleRetrohuntOutput)
}

type ChronicleRetrohuntOutput struct{ *pulumi.OutputState }

func (ChronicleRetrohuntOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleRetrohunt)(nil)).Elem()
}

func (o ChronicleRetrohuntOutput) ToChronicleRetrohuntOutput() ChronicleRetrohuntOutput {
	return o
}

func (o ChronicleRetrohuntOutput) ToChronicleRetrohuntOutputWithContext(ctx context.Context) ChronicleRetrohuntOutput {
	return o
}

func (o ChronicleRetrohuntOutput) ChronicleRetrohuntId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.ChronicleRetrohuntId }).(pulumi.StringOutput)
}

// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
// and end are unspecified, the interval matches any time.
func (o ChronicleRetrohuntOutput) ExecutionIntervals() ChronicleRetrohuntExecutionIntervalArrayOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) ChronicleRetrohuntExecutionIntervalArrayOutput {
		return v.ExecutionIntervals
	}).(ChronicleRetrohuntExecutionIntervalArrayOutput)
}

// The unique identifier for the Chronicle instance, which is the same as the customer ID.
func (o ChronicleRetrohuntOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
// "europe-west2".
func (o ChronicleRetrohuntOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the retrohunt. Retrohunt is the child of a rule revision. {rule} in the format below is structured
// as {rule_id@revision_id}. Format:
// projects/{project}/locations/{location}/instances/{instance}/rules/{rule}/retrohunts/{retrohunt}
func (o ChronicleRetrohuntOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Represents a time interval, encoded as a Timestamp start (inclusive) and a Timestamp end (exclusive). The start must be
// less than or equal to the end. When the start equals the end, the interval is empty (matches no time). When both start
// and end are unspecified, the interval matches any time.
func (o ChronicleRetrohuntOutput) ProcessInterval() ChronicleRetrohuntProcessIntervalOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) ChronicleRetrohuntProcessIntervalOutput { return v.ProcessInterval }).(ChronicleRetrohuntProcessIntervalOutput)
}

// Output only. Percent progress of the retrohunt towards completion, from 0.00 to 100.00.
func (o ChronicleRetrohuntOutput) ProgressPercentage() pulumi.Float64Output {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.Float64Output { return v.ProgressPercentage }).(pulumi.Float64Output)
}

func (o ChronicleRetrohuntOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The retrohunt ID of the Retrohunt. A retrohunt is an execution of a Rule over a time range in the past.
func (o ChronicleRetrohuntOutput) Retrohunt() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Retrohunt }).(pulumi.StringOutput)
}

// The Rule ID of the rule.
func (o ChronicleRetrohuntOutput) Rule() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.Rule }).(pulumi.StringOutput)
}

// Output only. The state of the retrohunt. Possible values: RUNNING DONE CANCELLED FAILED
func (o ChronicleRetrohuntOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ChronicleRetrohuntOutput) Timeouts() ChronicleRetrohuntTimeoutsPtrOutput {
	return o.ApplyT(func(v *ChronicleRetrohunt) ChronicleRetrohuntTimeoutsPtrOutput { return v.Timeouts }).(ChronicleRetrohuntTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChronicleRetrohuntInput)(nil)).Elem(), &ChronicleRetrohunt{})
	pulumi.RegisterOutputType(ChronicleRetrohuntOutput{})
}
