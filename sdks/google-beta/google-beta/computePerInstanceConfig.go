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

type ComputePerInstanceConfig struct {
	pulumi.CustomResourceState

	ComputePerInstanceConfigId pulumi.StringOutput `pulumi:"computePerInstanceConfigId"`
	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringOutput `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
	// RESTART * REFRESH * NONE
	MinimalAction pulumi.StringPtrOutput `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
	// REPLACE * RESTART * REFRESH * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrOutput `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// The preserved state for this instance.
	PreservedState ComputePerInstanceConfigPreservedStatePtrOutput `pulumi:"preservedState"`
	Project        pulumi.StringOutput                             `pulumi:"project"`
	// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
	// use the behavior as determined by remove_instance_on_destroy.
	RemoveInstanceOnDestroy pulumi.BoolPtrOutput `pulumi:"removeInstanceOnDestroy"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
	// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
	// next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrOutput                      `pulumi:"removeInstanceStateOnDestroy"`
	Timeouts                     ComputePerInstanceConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Zone where the containing instance group manager is located
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewComputePerInstanceConfig registers a new resource with the given unique name, arguments, and options.
func NewComputePerInstanceConfig(ctx *pulumi.Context,
	name string, args *ComputePerInstanceConfigArgs, opts ...pulumi.ResourceOption) (*ComputePerInstanceConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceGroupManager == nil {
		return nil, errors.New("invalid value for required argument 'InstanceGroupManager'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputePerInstanceConfig
	err = ctx.RegisterPackageResource("google-beta:index/computePerInstanceConfig:ComputePerInstanceConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputePerInstanceConfig gets an existing ComputePerInstanceConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputePerInstanceConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputePerInstanceConfigState, opts ...pulumi.ResourceOption) (*ComputePerInstanceConfig, error) {
	var resource ComputePerInstanceConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computePerInstanceConfig:ComputePerInstanceConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputePerInstanceConfig resources.
type computePerInstanceConfigState struct {
	ComputePerInstanceConfigId *string `pulumi:"computePerInstanceConfigId"`
	// The instance group manager this instance config is part of.
	InstanceGroupManager *string `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
	// RESTART * REFRESH * NONE
	MinimalAction *string `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
	// REPLACE * RESTART * REFRESH * NONE
	MostDisruptiveAllowedAction *string `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name *string `pulumi:"name"`
	// The preserved state for this instance.
	PreservedState *ComputePerInstanceConfigPreservedState `pulumi:"preservedState"`
	Project        *string                                 `pulumi:"project"`
	// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
	// use the behavior as determined by remove_instance_on_destroy.
	RemoveInstanceOnDestroy *bool `pulumi:"removeInstanceOnDestroy"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
	// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
	// next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool                             `pulumi:"removeInstanceStateOnDestroy"`
	Timeouts                     *ComputePerInstanceConfigTimeouts `pulumi:"timeouts"`
	// Zone where the containing instance group manager is located
	Zone *string `pulumi:"zone"`
}

type ComputePerInstanceConfigState struct {
	ComputePerInstanceConfigId pulumi.StringPtrInput
	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringPtrInput
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
	// RESTART * REFRESH * NONE
	MinimalAction pulumi.StringPtrInput
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
	// REPLACE * RESTART * REFRESH * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrInput
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringPtrInput
	// The preserved state for this instance.
	PreservedState ComputePerInstanceConfigPreservedStatePtrInput
	Project        pulumi.StringPtrInput
	// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
	// use the behavior as determined by remove_instance_on_destroy.
	RemoveInstanceOnDestroy pulumi.BoolPtrInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
	// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
	// next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
	Timeouts                     ComputePerInstanceConfigTimeoutsPtrInput
	// Zone where the containing instance group manager is located
	Zone pulumi.StringPtrInput
}

func (ComputePerInstanceConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*computePerInstanceConfigState)(nil)).Elem()
}

type computePerInstanceConfigArgs struct {
	ComputePerInstanceConfigId *string `pulumi:"computePerInstanceConfigId"`
	// The instance group manager this instance config is part of.
	InstanceGroupManager string `pulumi:"instanceGroupManager"`
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
	// RESTART * REFRESH * NONE
	MinimalAction *string `pulumi:"minimalAction"`
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
	// REPLACE * RESTART * REFRESH * NONE
	MostDisruptiveAllowedAction *string `pulumi:"mostDisruptiveAllowedAction"`
	// The name for this per-instance config and its corresponding instance.
	Name *string `pulumi:"name"`
	// The preserved state for this instance.
	PreservedState *ComputePerInstanceConfigPreservedState `pulumi:"preservedState"`
	Project        *string                                 `pulumi:"project"`
	// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
	// use the behavior as determined by remove_instance_on_destroy.
	RemoveInstanceOnDestroy *bool `pulumi:"removeInstanceOnDestroy"`
	// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
	// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
	// next instance recreation or update.
	RemoveInstanceStateOnDestroy *bool                             `pulumi:"removeInstanceStateOnDestroy"`
	Timeouts                     *ComputePerInstanceConfigTimeouts `pulumi:"timeouts"`
	// Zone where the containing instance group manager is located
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a ComputePerInstanceConfig resource.
type ComputePerInstanceConfigArgs struct {
	ComputePerInstanceConfigId pulumi.StringPtrInput
	// The instance group manager this instance config is part of.
	InstanceGroupManager pulumi.StringInput
	// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
	// RESTART * REFRESH * NONE
	MinimalAction pulumi.StringPtrInput
	// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
	// REPLACE * RESTART * REFRESH * NONE
	MostDisruptiveAllowedAction pulumi.StringPtrInput
	// The name for this per-instance config and its corresponding instance.
	Name pulumi.StringPtrInput
	// The preserved state for this instance.
	PreservedState ComputePerInstanceConfigPreservedStatePtrInput
	Project        pulumi.StringPtrInput
	// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
	// use the behavior as determined by remove_instance_on_destroy.
	RemoveInstanceOnDestroy pulumi.BoolPtrInput
	// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
	// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
	// next instance recreation or update.
	RemoveInstanceStateOnDestroy pulumi.BoolPtrInput
	Timeouts                     ComputePerInstanceConfigTimeoutsPtrInput
	// Zone where the containing instance group manager is located
	Zone pulumi.StringPtrInput
}

func (ComputePerInstanceConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computePerInstanceConfigArgs)(nil)).Elem()
}

type ComputePerInstanceConfigInput interface {
	pulumi.Input

	ToComputePerInstanceConfigOutput() ComputePerInstanceConfigOutput
	ToComputePerInstanceConfigOutputWithContext(ctx context.Context) ComputePerInstanceConfigOutput
}

func (*ComputePerInstanceConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputePerInstanceConfig)(nil)).Elem()
}

func (i *ComputePerInstanceConfig) ToComputePerInstanceConfigOutput() ComputePerInstanceConfigOutput {
	return i.ToComputePerInstanceConfigOutputWithContext(context.Background())
}

func (i *ComputePerInstanceConfig) ToComputePerInstanceConfigOutputWithContext(ctx context.Context) ComputePerInstanceConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputePerInstanceConfigOutput)
}

type ComputePerInstanceConfigOutput struct{ *pulumi.OutputState }

func (ComputePerInstanceConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputePerInstanceConfig)(nil)).Elem()
}

func (o ComputePerInstanceConfigOutput) ToComputePerInstanceConfigOutput() ComputePerInstanceConfigOutput {
	return o
}

func (o ComputePerInstanceConfigOutput) ToComputePerInstanceConfigOutputWithContext(ctx context.Context) ComputePerInstanceConfigOutput {
	return o
}

func (o ComputePerInstanceConfigOutput) ComputePerInstanceConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringOutput { return v.ComputePerInstanceConfigId }).(pulumi.StringOutput)
}

// The instance group manager this instance config is part of.
func (o ComputePerInstanceConfigOutput) InstanceGroupManager() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringOutput { return v.InstanceGroupManager }).(pulumi.StringOutput)
}

// The minimal action to perform on the instance during an update. Default is 'NONE'. Possible values are: * REPLACE *
// RESTART * REFRESH * NONE
func (o ComputePerInstanceConfigOutput) MinimalAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringPtrOutput { return v.MinimalAction }).(pulumi.StringPtrOutput)
}

// The most disruptive action to perform on the instance during an update. Default is 'REPLACE'. Possible values are: *
// REPLACE * RESTART * REFRESH * NONE
func (o ComputePerInstanceConfigOutput) MostDisruptiveAllowedAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringPtrOutput { return v.MostDisruptiveAllowedAction }).(pulumi.StringPtrOutput)
}

// The name for this per-instance config and its corresponding instance.
func (o ComputePerInstanceConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The preserved state for this instance.
func (o ComputePerInstanceConfigOutput) PreservedState() ComputePerInstanceConfigPreservedStatePtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) ComputePerInstanceConfigPreservedStatePtrOutput {
		return v.PreservedState
	}).(ComputePerInstanceConfigPreservedStatePtrOutput)
}

func (o ComputePerInstanceConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// When true, deleting this config will immediately remove the underlying instance. When false, deleting this config will
// use the behavior as determined by remove_instance_on_destroy.
func (o ComputePerInstanceConfigOutput) RemoveInstanceOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.BoolPtrOutput { return v.RemoveInstanceOnDestroy }).(pulumi.BoolPtrOutput)
}

// When true, deleting this config will immediately remove any specified state from the underlying instance. When false,
// deleting this config will *not* immediately remove any state from the underlying instance. State will be removed on the
// next instance recreation or update.
func (o ComputePerInstanceConfigOutput) RemoveInstanceStateOnDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.BoolPtrOutput { return v.RemoveInstanceStateOnDestroy }).(pulumi.BoolPtrOutput)
}

func (o ComputePerInstanceConfigOutput) Timeouts() ComputePerInstanceConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) ComputePerInstanceConfigTimeoutsPtrOutput { return v.Timeouts }).(ComputePerInstanceConfigTimeoutsPtrOutput)
}

// Zone where the containing instance group manager is located
func (o ComputePerInstanceConfigOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputePerInstanceConfig) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputePerInstanceConfigInput)(nil)).Elem(), &ComputePerInstanceConfig{})
	pulumi.RegisterOutputType(ComputePerInstanceConfigOutput{})
}
