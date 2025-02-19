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

type MonitoringDashboard struct {
	pulumi.CustomResourceState

	// The JSON representation of a dashboard, following the format at
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	DashboardJson         pulumi.StringOutput `pulumi:"dashboardJson"`
	MonitoringDashboardId pulumi.StringOutput `pulumi:"monitoringDashboardId"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project  pulumi.StringOutput                  `pulumi:"project"`
	Timeouts MonitoringDashboardTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewMonitoringDashboard registers a new resource with the given unique name, arguments, and options.
func NewMonitoringDashboard(ctx *pulumi.Context,
	name string, args *MonitoringDashboardArgs, opts ...pulumi.ResourceOption) (*MonitoringDashboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DashboardJson == nil {
		return nil, errors.New("invalid value for required argument 'DashboardJson'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource MonitoringDashboard
	err = ctx.RegisterPackageResource("google-beta:index/monitoringDashboard:MonitoringDashboard", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringDashboard gets an existing MonitoringDashboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringDashboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringDashboardState, opts ...pulumi.ResourceOption) (*MonitoringDashboard, error) {
	var resource MonitoringDashboard
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/monitoringDashboard:MonitoringDashboard", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringDashboard resources.
type monitoringDashboardState struct {
	// The JSON representation of a dashboard, following the format at
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	DashboardJson         *string `pulumi:"dashboardJson"`
	MonitoringDashboardId *string `pulumi:"monitoringDashboardId"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project  *string                      `pulumi:"project"`
	Timeouts *MonitoringDashboardTimeouts `pulumi:"timeouts"`
}

type MonitoringDashboardState struct {
	// The JSON representation of a dashboard, following the format at
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	DashboardJson         pulumi.StringPtrInput
	MonitoringDashboardId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project  pulumi.StringPtrInput
	Timeouts MonitoringDashboardTimeoutsPtrInput
}

func (MonitoringDashboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringDashboardState)(nil)).Elem()
}

type monitoringDashboardArgs struct {
	// The JSON representation of a dashboard, following the format at
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	DashboardJson         string  `pulumi:"dashboardJson"`
	MonitoringDashboardId *string `pulumi:"monitoringDashboardId"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project  *string                      `pulumi:"project"`
	Timeouts *MonitoringDashboardTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a MonitoringDashboard resource.
type MonitoringDashboardArgs struct {
	// The JSON representation of a dashboard, following the format at
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
	DashboardJson         pulumi.StringInput
	MonitoringDashboardId pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project  pulumi.StringPtrInput
	Timeouts MonitoringDashboardTimeoutsPtrInput
}

func (MonitoringDashboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringDashboardArgs)(nil)).Elem()
}

type MonitoringDashboardInput interface {
	pulumi.Input

	ToMonitoringDashboardOutput() MonitoringDashboardOutput
	ToMonitoringDashboardOutputWithContext(ctx context.Context) MonitoringDashboardOutput
}

func (*MonitoringDashboard) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringDashboard)(nil)).Elem()
}

func (i *MonitoringDashboard) ToMonitoringDashboardOutput() MonitoringDashboardOutput {
	return i.ToMonitoringDashboardOutputWithContext(context.Background())
}

func (i *MonitoringDashboard) ToMonitoringDashboardOutputWithContext(ctx context.Context) MonitoringDashboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringDashboardOutput)
}

type MonitoringDashboardOutput struct{ *pulumi.OutputState }

func (MonitoringDashboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringDashboard)(nil)).Elem()
}

func (o MonitoringDashboardOutput) ToMonitoringDashboardOutput() MonitoringDashboardOutput {
	return o
}

func (o MonitoringDashboardOutput) ToMonitoringDashboardOutputWithContext(ctx context.Context) MonitoringDashboardOutput {
	return o
}

// The JSON representation of a dashboard, following the format at
// https://cloud.google.com/monitoring/api/ref_v3/rest/v1/projects.dashboards.
func (o MonitoringDashboardOutput) DashboardJson() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringDashboard) pulumi.StringOutput { return v.DashboardJson }).(pulumi.StringOutput)
}

func (o MonitoringDashboardOutput) MonitoringDashboardId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringDashboard) pulumi.StringOutput { return v.MonitoringDashboardId }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
func (o MonitoringDashboardOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringDashboard) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o MonitoringDashboardOutput) Timeouts() MonitoringDashboardTimeoutsPtrOutput {
	return o.ApplyT(func(v *MonitoringDashboard) MonitoringDashboardTimeoutsPtrOutput { return v.Timeouts }).(MonitoringDashboardTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringDashboardInput)(nil)).Elem(), &MonitoringDashboard{})
	pulumi.RegisterOutputType(MonitoringDashboardOutput{})
}
