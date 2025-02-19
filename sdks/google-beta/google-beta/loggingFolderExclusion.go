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

type LoggingFolderExclusion struct {
	pulumi.CustomResourceState

	// A human-readable description.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                   pulumi.StringOutput `pulumi:"filter"`
	Folder                   pulumi.StringOutput `pulumi:"folder"`
	LoggingFolderExclusionId pulumi.StringOutput `pulumi:"loggingFolderExclusionId"`
	// The name of the logging exclusion.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewLoggingFolderExclusion registers a new resource with the given unique name, arguments, and options.
func NewLoggingFolderExclusion(ctx *pulumi.Context,
	name string, args *LoggingFolderExclusionArgs, opts ...pulumi.ResourceOption) (*LoggingFolderExclusion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource LoggingFolderExclusion
	err = ctx.RegisterPackageResource("google-beta:index/loggingFolderExclusion:LoggingFolderExclusion", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoggingFolderExclusion gets an existing LoggingFolderExclusion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoggingFolderExclusion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoggingFolderExclusionState, opts ...pulumi.ResourceOption) (*LoggingFolderExclusion, error) {
	var resource LoggingFolderExclusion
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/loggingFolderExclusion:LoggingFolderExclusion", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoggingFolderExclusion resources.
type loggingFolderExclusionState struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                   *string `pulumi:"filter"`
	Folder                   *string `pulumi:"folder"`
	LoggingFolderExclusionId *string `pulumi:"loggingFolderExclusionId"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

type LoggingFolderExclusionState struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                   pulumi.StringPtrInput
	Folder                   pulumi.StringPtrInput
	LoggingFolderExclusionId pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (LoggingFolderExclusionState) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderExclusionState)(nil)).Elem()
}

type loggingFolderExclusionArgs struct {
	// A human-readable description.
	Description *string `pulumi:"description"`
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled *bool `pulumi:"disabled"`
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                   string  `pulumi:"filter"`
	Folder                   string  `pulumi:"folder"`
	LoggingFolderExclusionId *string `pulumi:"loggingFolderExclusionId"`
	// The name of the logging exclusion.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a LoggingFolderExclusion resource.
type LoggingFolderExclusionArgs struct {
	// A human-readable description.
	Description pulumi.StringPtrInput
	// Whether this exclusion rule should be disabled or not. This defaults to false.
	Disabled pulumi.BoolPtrInput
	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	Filter                   pulumi.StringInput
	Folder                   pulumi.StringInput
	LoggingFolderExclusionId pulumi.StringPtrInput
	// The name of the logging exclusion.
	Name pulumi.StringPtrInput
}

func (LoggingFolderExclusionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loggingFolderExclusionArgs)(nil)).Elem()
}

type LoggingFolderExclusionInput interface {
	pulumi.Input

	ToLoggingFolderExclusionOutput() LoggingFolderExclusionOutput
	ToLoggingFolderExclusionOutputWithContext(ctx context.Context) LoggingFolderExclusionOutput
}

func (*LoggingFolderExclusion) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderExclusion)(nil)).Elem()
}

func (i *LoggingFolderExclusion) ToLoggingFolderExclusionOutput() LoggingFolderExclusionOutput {
	return i.ToLoggingFolderExclusionOutputWithContext(context.Background())
}

func (i *LoggingFolderExclusion) ToLoggingFolderExclusionOutputWithContext(ctx context.Context) LoggingFolderExclusionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingFolderExclusionOutput)
}

type LoggingFolderExclusionOutput struct{ *pulumi.OutputState }

func (LoggingFolderExclusionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingFolderExclusion)(nil)).Elem()
}

func (o LoggingFolderExclusionOutput) ToLoggingFolderExclusionOutput() LoggingFolderExclusionOutput {
	return o
}

func (o LoggingFolderExclusionOutput) ToLoggingFolderExclusionOutputWithContext(ctx context.Context) LoggingFolderExclusionOutput {
	return o
}

// A human-readable description.
func (o LoggingFolderExclusionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether this exclusion rule should be disabled or not. This defaults to false.
func (o LoggingFolderExclusionOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
func (o LoggingFolderExclusionOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

func (o LoggingFolderExclusionOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

func (o LoggingFolderExclusionOutput) LoggingFolderExclusionId() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.StringOutput { return v.LoggingFolderExclusionId }).(pulumi.StringOutput)
}

// The name of the logging exclusion.
func (o LoggingFolderExclusionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LoggingFolderExclusion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoggingFolderExclusionInput)(nil)).Elem(), &LoggingFolderExclusion{})
	pulumi.RegisterOutputType(LoggingFolderExclusionOutput{})
}
