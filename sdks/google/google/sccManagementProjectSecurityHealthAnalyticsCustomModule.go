// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SccManagementProjectSecurityHealthAnalyticsCustomModule struct {
	pulumi.CustomResourceState

	// If empty, indicates that the custom module was created in the organization,folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule pulumi.StringOutput `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrOutput `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrOutput `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor pulumi.StringOutput `pulumi:"lastEditor"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The resource name of the custom module. Its format is
	// "projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securityHealthAnalyticsCustomModule}".
	// The id {securityHealthAnalyticsCustomModule} is server-generated and is not user settable. It will be a numeric id
	// containing 1-20 digits.
	Name                                                      pulumi.StringOutput                                                      `pulumi:"name"`
	Project                                                   pulumi.StringOutput                                                      `pulumi:"project"`
	SccManagementProjectSecurityHealthAnalyticsCustomModuleId pulumi.StringOutput                                                      `pulumi:"sccManagementProjectSecurityHealthAnalyticsCustomModuleId"`
	Timeouts                                                  SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrOutput `pulumi:"timeouts"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccManagementProjectSecurityHealthAnalyticsCustomModule registers a new resource with the given unique name, arguments, and options.
func NewSccManagementProjectSecurityHealthAnalyticsCustomModule(ctx *pulumi.Context,
	name string, args *SccManagementProjectSecurityHealthAnalyticsCustomModuleArgs, opts ...pulumi.ResourceOption) (*SccManagementProjectSecurityHealthAnalyticsCustomModule, error) {
	if args == nil {
		args = &SccManagementProjectSecurityHealthAnalyticsCustomModuleArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccManagementProjectSecurityHealthAnalyticsCustomModule
	err = ctx.RegisterPackageResource("google:index/sccManagementProjectSecurityHealthAnalyticsCustomModule:SccManagementProjectSecurityHealthAnalyticsCustomModule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccManagementProjectSecurityHealthAnalyticsCustomModule gets an existing SccManagementProjectSecurityHealthAnalyticsCustomModule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccManagementProjectSecurityHealthAnalyticsCustomModule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccManagementProjectSecurityHealthAnalyticsCustomModuleState, opts ...pulumi.ResourceOption) (*SccManagementProjectSecurityHealthAnalyticsCustomModule, error) {
	var resource SccManagementProjectSecurityHealthAnalyticsCustomModule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sccManagementProjectSecurityHealthAnalyticsCustomModule:SccManagementProjectSecurityHealthAnalyticsCustomModule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccManagementProjectSecurityHealthAnalyticsCustomModule resources.
type sccManagementProjectSecurityHealthAnalyticsCustomModuleState struct {
	// If empty, indicates that the custom module was created in the organization,folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule *string `pulumi:"ancestorModule"`
	// The user specified custom configuration for the module.
	CustomConfig *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName *string `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// The editor that last updated the custom module.
	LastEditor *string `pulumi:"lastEditor"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	Location *string `pulumi:"location"`
	// The resource name of the custom module. Its format is
	// "projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securityHealthAnalyticsCustomModule}".
	// The id {securityHealthAnalyticsCustomModule} is server-generated and is not user settable. It will be a numeric id
	// containing 1-20 digits.
	Name                                                      *string                                                          `pulumi:"name"`
	Project                                                   *string                                                          `pulumi:"project"`
	SccManagementProjectSecurityHealthAnalyticsCustomModuleId *string                                                          `pulumi:"sccManagementProjectSecurityHealthAnalyticsCustomModuleId"`
	Timeouts                                                  *SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeouts `pulumi:"timeouts"`
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccManagementProjectSecurityHealthAnalyticsCustomModuleState struct {
	// If empty, indicates that the custom module was created in the organization,folder, or project in which you are viewing
	// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
	// inherited.
	AncestorModule pulumi.StringPtrInput
	// The user specified custom configuration for the module.
	CustomConfig SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrInput
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// The editor that last updated the custom module.
	LastEditor pulumi.StringPtrInput
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	Location pulumi.StringPtrInput
	// The resource name of the custom module. Its format is
	// "projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securityHealthAnalyticsCustomModule}".
	// The id {securityHealthAnalyticsCustomModule} is server-generated and is not user settable. It will be a numeric id
	// containing 1-20 digits.
	Name                                                      pulumi.StringPtrInput
	Project                                                   pulumi.StringPtrInput
	SccManagementProjectSecurityHealthAnalyticsCustomModuleId pulumi.StringPtrInput
	Timeouts                                                  SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrInput
	// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccManagementProjectSecurityHealthAnalyticsCustomModuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccManagementProjectSecurityHealthAnalyticsCustomModuleState)(nil)).Elem()
}

type sccManagementProjectSecurityHealthAnalyticsCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig *SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfig `pulumi:"customConfig"`
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName *string `pulumi:"displayName"`
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState *string `pulumi:"enablementState"`
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	Location                                                  *string                                                          `pulumi:"location"`
	Project                                                   *string                                                          `pulumi:"project"`
	SccManagementProjectSecurityHealthAnalyticsCustomModuleId *string                                                          `pulumi:"sccManagementProjectSecurityHealthAnalyticsCustomModuleId"`
	Timeouts                                                  *SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccManagementProjectSecurityHealthAnalyticsCustomModule resource.
type SccManagementProjectSecurityHealthAnalyticsCustomModuleArgs struct {
	// The user specified custom configuration for the module.
	CustomConfig SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrInput
	// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
	// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
	// lowercase letter, and contain alphanumeric characters or underscores only.
	DisplayName pulumi.StringPtrInput
	// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
	EnablementState pulumi.StringPtrInput
	// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
	Location                                                  pulumi.StringPtrInput
	Project                                                   pulumi.StringPtrInput
	SccManagementProjectSecurityHealthAnalyticsCustomModuleId pulumi.StringPtrInput
	Timeouts                                                  SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrInput
}

func (SccManagementProjectSecurityHealthAnalyticsCustomModuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccManagementProjectSecurityHealthAnalyticsCustomModuleArgs)(nil)).Elem()
}

type SccManagementProjectSecurityHealthAnalyticsCustomModuleInput interface {
	pulumi.Input

	ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutput() SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput
	ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutputWithContext(ctx context.Context) SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput
}

func (*SccManagementProjectSecurityHealthAnalyticsCustomModule) ElementType() reflect.Type {
	return reflect.TypeOf((**SccManagementProjectSecurityHealthAnalyticsCustomModule)(nil)).Elem()
}

func (i *SccManagementProjectSecurityHealthAnalyticsCustomModule) ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutput() SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput {
	return i.ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutputWithContext(context.Background())
}

func (i *SccManagementProjectSecurityHealthAnalyticsCustomModule) ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutputWithContext(ctx context.Context) SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput)
}

type SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput struct{ *pulumi.OutputState }

func (SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccManagementProjectSecurityHealthAnalyticsCustomModule)(nil)).Elem()
}

func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutput() SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput {
	return o
}

func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) ToSccManagementProjectSecurityHealthAnalyticsCustomModuleOutputWithContext(ctx context.Context) SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput {
	return o
}

// If empty, indicates that the custom module was created in the organization,folder, or project in which you are viewing
// the custom module. Otherwise, ancestor_module specifies the organization or folder from which the custom module is
// inherited.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) AncestorModule() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput {
		return v.AncestorModule
	}).(pulumi.StringOutput)
}

// The user specified custom configuration for the module.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) CustomConfig() SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrOutput {
		return v.CustomConfig
	}).(SccManagementProjectSecurityHealthAnalyticsCustomModuleCustomConfigPtrOutput)
}

// The display name of the Security Health Analytics custom module. This display name becomes the finding category for all
// findings that are returned by this custom module. The display name must be between 1 and 128 characters, start with a
// lowercase letter, and contain alphanumeric characters or underscores only.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringPtrOutput {
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// The enablement state of the custom module. Possible values: ["ENABLED", "DISABLED"]
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) EnablementState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringPtrOutput {
		return v.EnablementState
	}).(pulumi.StringPtrOutput)
}

// The editor that last updated the custom module.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) LastEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput {
		return v.LastEditor
	}).(pulumi.StringOutput)
}

// Location ID of the parent organization. If not provided, 'global' will be used as the default location.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringPtrOutput {
		return v.Location
	}).(pulumi.StringPtrOutput)
}

// The resource name of the custom module. Its format is
// "projects/{project}/locations/{location}/securityHealthAnalyticsCustomModules/{securityHealthAnalyticsCustomModule}".
// The id {securityHealthAnalyticsCustomModule} is server-generated and is not user settable. It will be a numeric id
// containing 1-20 digits.
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) SccManagementProjectSecurityHealthAnalyticsCustomModuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput {
		return v.SccManagementProjectSecurityHealthAnalyticsCustomModuleId
	}).(pulumi.StringOutput)
}

func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) Timeouts() SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrOutput {
		return v.Timeouts
	}).(SccManagementProjectSecurityHealthAnalyticsCustomModuleTimeoutsPtrOutput)
}

// The time at which the custom module was last updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccManagementProjectSecurityHealthAnalyticsCustomModule) pulumi.StringOutput {
		return v.UpdateTime
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccManagementProjectSecurityHealthAnalyticsCustomModuleInput)(nil)).Elem(), &SccManagementProjectSecurityHealthAnalyticsCustomModule{})
	pulumi.RegisterOutputType(SccManagementProjectSecurityHealthAnalyticsCustomModuleOutput{})
}
