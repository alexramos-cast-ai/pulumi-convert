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

type RebalancingSchedule struct {
	pulumi.CustomResourceState

	LaunchConfiguration RebalancingScheduleLaunchConfigurationOutput `pulumi:"launchConfiguration"`
	// Name of the schedule.
	Name                  pulumi.StringOutput                        `pulumi:"name"`
	RebalancingScheduleId pulumi.StringOutput                        `pulumi:"rebalancingScheduleId"`
	Schedule              RebalancingScheduleScheduleOutput          `pulumi:"schedule"`
	Timeouts              RebalancingScheduleTimeoutsPtrOutput       `pulumi:"timeouts"`
	TriggerConditions     RebalancingScheduleTriggerConditionsOutput `pulumi:"triggerConditions"`
}

// NewRebalancingSchedule registers a new resource with the given unique name, arguments, and options.
func NewRebalancingSchedule(ctx *pulumi.Context,
	name string, args *RebalancingScheduleArgs, opts ...pulumi.ResourceOption) (*RebalancingSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LaunchConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'LaunchConfiguration'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	if args.TriggerConditions == nil {
		return nil, errors.New("invalid value for required argument 'TriggerConditions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RebalancingSchedule
	err = ctx.RegisterPackageResource("castai:index/rebalancingSchedule:RebalancingSchedule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRebalancingSchedule gets an existing RebalancingSchedule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRebalancingSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RebalancingScheduleState, opts ...pulumi.ResourceOption) (*RebalancingSchedule, error) {
	var resource RebalancingSchedule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/rebalancingSchedule:RebalancingSchedule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RebalancingSchedule resources.
type rebalancingScheduleState struct {
	LaunchConfiguration *RebalancingScheduleLaunchConfiguration `pulumi:"launchConfiguration"`
	// Name of the schedule.
	Name                  *string                               `pulumi:"name"`
	RebalancingScheduleId *string                               `pulumi:"rebalancingScheduleId"`
	Schedule              *RebalancingScheduleSchedule          `pulumi:"schedule"`
	Timeouts              *RebalancingScheduleTimeouts          `pulumi:"timeouts"`
	TriggerConditions     *RebalancingScheduleTriggerConditions `pulumi:"triggerConditions"`
}

type RebalancingScheduleState struct {
	LaunchConfiguration RebalancingScheduleLaunchConfigurationPtrInput
	// Name of the schedule.
	Name                  pulumi.StringPtrInput
	RebalancingScheduleId pulumi.StringPtrInput
	Schedule              RebalancingScheduleSchedulePtrInput
	Timeouts              RebalancingScheduleTimeoutsPtrInput
	TriggerConditions     RebalancingScheduleTriggerConditionsPtrInput
}

func (RebalancingScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*rebalancingScheduleState)(nil)).Elem()
}

type rebalancingScheduleArgs struct {
	LaunchConfiguration RebalancingScheduleLaunchConfiguration `pulumi:"launchConfiguration"`
	// Name of the schedule.
	Name                  *string                              `pulumi:"name"`
	RebalancingScheduleId *string                              `pulumi:"rebalancingScheduleId"`
	Schedule              RebalancingScheduleSchedule          `pulumi:"schedule"`
	Timeouts              *RebalancingScheduleTimeouts         `pulumi:"timeouts"`
	TriggerConditions     RebalancingScheduleTriggerConditions `pulumi:"triggerConditions"`
}

// The set of arguments for constructing a RebalancingSchedule resource.
type RebalancingScheduleArgs struct {
	LaunchConfiguration RebalancingScheduleLaunchConfigurationInput
	// Name of the schedule.
	Name                  pulumi.StringPtrInput
	RebalancingScheduleId pulumi.StringPtrInput
	Schedule              RebalancingScheduleScheduleInput
	Timeouts              RebalancingScheduleTimeoutsPtrInput
	TriggerConditions     RebalancingScheduleTriggerConditionsInput
}

func (RebalancingScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rebalancingScheduleArgs)(nil)).Elem()
}

type RebalancingScheduleInput interface {
	pulumi.Input

	ToRebalancingScheduleOutput() RebalancingScheduleOutput
	ToRebalancingScheduleOutputWithContext(ctx context.Context) RebalancingScheduleOutput
}

func (*RebalancingSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**RebalancingSchedule)(nil)).Elem()
}

func (i *RebalancingSchedule) ToRebalancingScheduleOutput() RebalancingScheduleOutput {
	return i.ToRebalancingScheduleOutputWithContext(context.Background())
}

func (i *RebalancingSchedule) ToRebalancingScheduleOutputWithContext(ctx context.Context) RebalancingScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RebalancingScheduleOutput)
}

type RebalancingScheduleOutput struct{ *pulumi.OutputState }

func (RebalancingScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RebalancingSchedule)(nil)).Elem()
}

func (o RebalancingScheduleOutput) ToRebalancingScheduleOutput() RebalancingScheduleOutput {
	return o
}

func (o RebalancingScheduleOutput) ToRebalancingScheduleOutputWithContext(ctx context.Context) RebalancingScheduleOutput {
	return o
}

func (o RebalancingScheduleOutput) LaunchConfiguration() RebalancingScheduleLaunchConfigurationOutput {
	return o.ApplyT(func(v *RebalancingSchedule) RebalancingScheduleLaunchConfigurationOutput {
		return v.LaunchConfiguration
	}).(RebalancingScheduleLaunchConfigurationOutput)
}

// Name of the schedule.
func (o RebalancingScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RebalancingSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RebalancingScheduleOutput) RebalancingScheduleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RebalancingSchedule) pulumi.StringOutput { return v.RebalancingScheduleId }).(pulumi.StringOutput)
}

func (o RebalancingScheduleOutput) Schedule() RebalancingScheduleScheduleOutput {
	return o.ApplyT(func(v *RebalancingSchedule) RebalancingScheduleScheduleOutput { return v.Schedule }).(RebalancingScheduleScheduleOutput)
}

func (o RebalancingScheduleOutput) Timeouts() RebalancingScheduleTimeoutsPtrOutput {
	return o.ApplyT(func(v *RebalancingSchedule) RebalancingScheduleTimeoutsPtrOutput { return v.Timeouts }).(RebalancingScheduleTimeoutsPtrOutput)
}

func (o RebalancingScheduleOutput) TriggerConditions() RebalancingScheduleTriggerConditionsOutput {
	return o.ApplyT(func(v *RebalancingSchedule) RebalancingScheduleTriggerConditionsOutput { return v.TriggerConditions }).(RebalancingScheduleTriggerConditionsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RebalancingScheduleInput)(nil)).Elem(), &RebalancingSchedule{})
	pulumi.RegisterOutputType(RebalancingScheduleOutput{})
}
