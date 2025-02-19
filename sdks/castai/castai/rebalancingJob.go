// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RebalancingJob struct {
	pulumi.CustomResourceState

	// CAST AI cluster id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The job will only be executed if it's enabled.
	Enabled          pulumi.BoolPtrOutput `pulumi:"enabled"`
	RebalancingJobId pulumi.StringOutput  `pulumi:"rebalancingJobId"`
	// Rebalancing schedule of this job.
	RebalancingScheduleId pulumi.StringOutput             `pulumi:"rebalancingScheduleId"`
	Timeouts              RebalancingJobTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewRebalancingJob registers a new resource with the given unique name, arguments, and options.
func NewRebalancingJob(ctx *pulumi.Context,
	name string, args *RebalancingJobArgs, opts ...pulumi.ResourceOption) (*RebalancingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.RebalancingScheduleId == nil {
		return nil, errors.New("invalid value for required argument 'RebalancingScheduleId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RebalancingJob
	err = ctx.RegisterPackageResource("castai:index/rebalancingJob:RebalancingJob", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRebalancingJob gets an existing RebalancingJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRebalancingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RebalancingJobState, opts ...pulumi.ResourceOption) (*RebalancingJob, error) {
	var resource RebalancingJob
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/rebalancingJob:RebalancingJob", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RebalancingJob resources.
type rebalancingJobState struct {
	// CAST AI cluster id.
	ClusterId *string `pulumi:"clusterId"`
	// The job will only be executed if it's enabled.
	Enabled          *bool   `pulumi:"enabled"`
	RebalancingJobId *string `pulumi:"rebalancingJobId"`
	// Rebalancing schedule of this job.
	RebalancingScheduleId *string                 `pulumi:"rebalancingScheduleId"`
	Timeouts              *RebalancingJobTimeouts `pulumi:"timeouts"`
}

type RebalancingJobState struct {
	// CAST AI cluster id.
	ClusterId pulumi.StringPtrInput
	// The job will only be executed if it's enabled.
	Enabled          pulumi.BoolPtrInput
	RebalancingJobId pulumi.StringPtrInput
	// Rebalancing schedule of this job.
	RebalancingScheduleId pulumi.StringPtrInput
	Timeouts              RebalancingJobTimeoutsPtrInput
}

func (RebalancingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*rebalancingJobState)(nil)).Elem()
}

type rebalancingJobArgs struct {
	// CAST AI cluster id.
	ClusterId string `pulumi:"clusterId"`
	// The job will only be executed if it's enabled.
	Enabled          *bool   `pulumi:"enabled"`
	RebalancingJobId *string `pulumi:"rebalancingJobId"`
	// Rebalancing schedule of this job.
	RebalancingScheduleId string                  `pulumi:"rebalancingScheduleId"`
	Timeouts              *RebalancingJobTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a RebalancingJob resource.
type RebalancingJobArgs struct {
	// CAST AI cluster id.
	ClusterId pulumi.StringInput
	// The job will only be executed if it's enabled.
	Enabled          pulumi.BoolPtrInput
	RebalancingJobId pulumi.StringPtrInput
	// Rebalancing schedule of this job.
	RebalancingScheduleId pulumi.StringInput
	Timeouts              RebalancingJobTimeoutsPtrInput
}

func (RebalancingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rebalancingJobArgs)(nil)).Elem()
}

type RebalancingJobInput interface {
	pulumi.Input

	ToRebalancingJobOutput() RebalancingJobOutput
	ToRebalancingJobOutputWithContext(ctx context.Context) RebalancingJobOutput
}

func (*RebalancingJob) ElementType() reflect.Type {
	return reflect.TypeOf((**RebalancingJob)(nil)).Elem()
}

func (i *RebalancingJob) ToRebalancingJobOutput() RebalancingJobOutput {
	return i.ToRebalancingJobOutputWithContext(context.Background())
}

func (i *RebalancingJob) ToRebalancingJobOutputWithContext(ctx context.Context) RebalancingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RebalancingJobOutput)
}

type RebalancingJobOutput struct{ *pulumi.OutputState }

func (RebalancingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RebalancingJob)(nil)).Elem()
}

func (o RebalancingJobOutput) ToRebalancingJobOutput() RebalancingJobOutput {
	return o
}

func (o RebalancingJobOutput) ToRebalancingJobOutputWithContext(ctx context.Context) RebalancingJobOutput {
	return o
}

// CAST AI cluster id.
func (o RebalancingJobOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RebalancingJob) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The job will only be executed if it's enabled.
func (o RebalancingJobOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RebalancingJob) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o RebalancingJobOutput) RebalancingJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *RebalancingJob) pulumi.StringOutput { return v.RebalancingJobId }).(pulumi.StringOutput)
}

// Rebalancing schedule of this job.
func (o RebalancingJobOutput) RebalancingScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RebalancingJob) pulumi.StringOutput { return v.RebalancingScheduleId }).(pulumi.StringOutput)
}

func (o RebalancingJobOutput) Timeouts() RebalancingJobTimeoutsPtrOutput {
	return o.ApplyT(func(v *RebalancingJob) RebalancingJobTimeoutsPtrOutput { return v.Timeouts }).(RebalancingJobTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RebalancingJobInput)(nil)).Elem(), &RebalancingJob{})
	pulumi.RegisterOutputType(RebalancingJobOutput{})
}
