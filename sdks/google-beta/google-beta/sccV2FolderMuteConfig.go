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

type SccV2FolderMuteConfig struct {
	pulumi.CustomResourceState

	// The time at which the mute config was created. This field is set by the server and will be ignored if provided on config
	// creation.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the mute config.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
	// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
	// created under the project = Y scope, it might not match any findings.
	Filter pulumi.StringOutput `pulumi:"filter"`
	// The folder whose Cloud Security Command Center the Mute Config lives in.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// location Id is provided by folder. If not provided, Use global as default.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Email address of the user who last edited the mute config. This field is set by the server and will be ignored if
	// provided on config creation or update.
	MostRecentEditor pulumi.StringOutput `pulumi:"mostRecentEditor"`
	// Unique identifier provided by the client within the parent scope.
	MuteConfigId pulumi.StringOutput `pulumi:"muteConfigId"`
	// Name of the mute config. Its format is organizations/{organization}/locations/global/muteConfigs/{configId},
	// folders/{folder}/locations/global/muteConfigs/{configId}, or projects/{project}/locations/global/muteConfigs/{configId}
	Name                    pulumi.StringOutput                    `pulumi:"name"`
	SccV2FolderMuteConfigId pulumi.StringOutput                    `pulumi:"sccV2FolderMuteConfigId"`
	Timeouts                SccV2FolderMuteConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// The type of the mute config.
	Type pulumi.StringOutput `pulumi:"type"`
	// Output only. The most recent time at which the mute config was updated. This field is set by the server and will be
	// ignored if provided on config creation or update.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccV2FolderMuteConfig registers a new resource with the given unique name, arguments, and options.
func NewSccV2FolderMuteConfig(ctx *pulumi.Context,
	name string, args *SccV2FolderMuteConfigArgs, opts ...pulumi.ResourceOption) (*SccV2FolderMuteConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Filter == nil {
		return nil, errors.New("invalid value for required argument 'Filter'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	if args.MuteConfigId == nil {
		return nil, errors.New("invalid value for required argument 'MuteConfigId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccV2FolderMuteConfig
	err = ctx.RegisterPackageResource("google-beta:index/sccV2FolderMuteConfig:SccV2FolderMuteConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccV2FolderMuteConfig gets an existing SccV2FolderMuteConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccV2FolderMuteConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccV2FolderMuteConfigState, opts ...pulumi.ResourceOption) (*SccV2FolderMuteConfig, error) {
	var resource SccV2FolderMuteConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/sccV2FolderMuteConfig:SccV2FolderMuteConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccV2FolderMuteConfig resources.
type sccV2FolderMuteConfigState struct {
	// The time at which the mute config was created. This field is set by the server and will be ignored if provided on config
	// creation.
	CreateTime *string `pulumi:"createTime"`
	// A description of the mute config.
	Description *string `pulumi:"description"`
	// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
	// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
	// created under the project = Y scope, it might not match any findings.
	Filter *string `pulumi:"filter"`
	// The folder whose Cloud Security Command Center the Mute Config lives in.
	Folder *string `pulumi:"folder"`
	// location Id is provided by folder. If not provided, Use global as default.
	Location *string `pulumi:"location"`
	// Email address of the user who last edited the mute config. This field is set by the server and will be ignored if
	// provided on config creation or update.
	MostRecentEditor *string `pulumi:"mostRecentEditor"`
	// Unique identifier provided by the client within the parent scope.
	MuteConfigId *string `pulumi:"muteConfigId"`
	// Name of the mute config. Its format is organizations/{organization}/locations/global/muteConfigs/{configId},
	// folders/{folder}/locations/global/muteConfigs/{configId}, or projects/{project}/locations/global/muteConfigs/{configId}
	Name                    *string                        `pulumi:"name"`
	SccV2FolderMuteConfigId *string                        `pulumi:"sccV2FolderMuteConfigId"`
	Timeouts                *SccV2FolderMuteConfigTimeouts `pulumi:"timeouts"`
	// The type of the mute config.
	Type *string `pulumi:"type"`
	// Output only. The most recent time at which the mute config was updated. This field is set by the server and will be
	// ignored if provided on config creation or update.
	UpdateTime *string `pulumi:"updateTime"`
}

type SccV2FolderMuteConfigState struct {
	// The time at which the mute config was created. This field is set by the server and will be ignored if provided on config
	// creation.
	CreateTime pulumi.StringPtrInput
	// A description of the mute config.
	Description pulumi.StringPtrInput
	// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
	// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
	// created under the project = Y scope, it might not match any findings.
	Filter pulumi.StringPtrInput
	// The folder whose Cloud Security Command Center the Mute Config lives in.
	Folder pulumi.StringPtrInput
	// location Id is provided by folder. If not provided, Use global as default.
	Location pulumi.StringPtrInput
	// Email address of the user who last edited the mute config. This field is set by the server and will be ignored if
	// provided on config creation or update.
	MostRecentEditor pulumi.StringPtrInput
	// Unique identifier provided by the client within the parent scope.
	MuteConfigId pulumi.StringPtrInput
	// Name of the mute config. Its format is organizations/{organization}/locations/global/muteConfigs/{configId},
	// folders/{folder}/locations/global/muteConfigs/{configId}, or projects/{project}/locations/global/muteConfigs/{configId}
	Name                    pulumi.StringPtrInput
	SccV2FolderMuteConfigId pulumi.StringPtrInput
	Timeouts                SccV2FolderMuteConfigTimeoutsPtrInput
	// The type of the mute config.
	Type pulumi.StringPtrInput
	// Output only. The most recent time at which the mute config was updated. This field is set by the server and will be
	// ignored if provided on config creation or update.
	UpdateTime pulumi.StringPtrInput
}

func (SccV2FolderMuteConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2FolderMuteConfigState)(nil)).Elem()
}

type sccV2FolderMuteConfigArgs struct {
	// A description of the mute config.
	Description *string `pulumi:"description"`
	// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
	// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
	// created under the project = Y scope, it might not match any findings.
	Filter string `pulumi:"filter"`
	// The folder whose Cloud Security Command Center the Mute Config lives in.
	Folder string `pulumi:"folder"`
	// location Id is provided by folder. If not provided, Use global as default.
	Location *string `pulumi:"location"`
	// Unique identifier provided by the client within the parent scope.
	MuteConfigId            string                         `pulumi:"muteConfigId"`
	SccV2FolderMuteConfigId *string                        `pulumi:"sccV2FolderMuteConfigId"`
	Timeouts                *SccV2FolderMuteConfigTimeouts `pulumi:"timeouts"`
	// The type of the mute config.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a SccV2FolderMuteConfig resource.
type SccV2FolderMuteConfigArgs struct {
	// A description of the mute config.
	Description pulumi.StringPtrInput
	// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
	// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
	// created under the project = Y scope, it might not match any findings.
	Filter pulumi.StringInput
	// The folder whose Cloud Security Command Center the Mute Config lives in.
	Folder pulumi.StringInput
	// location Id is provided by folder. If not provided, Use global as default.
	Location pulumi.StringPtrInput
	// Unique identifier provided by the client within the parent scope.
	MuteConfigId            pulumi.StringInput
	SccV2FolderMuteConfigId pulumi.StringPtrInput
	Timeouts                SccV2FolderMuteConfigTimeoutsPtrInput
	// The type of the mute config.
	Type pulumi.StringInput
}

func (SccV2FolderMuteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2FolderMuteConfigArgs)(nil)).Elem()
}

type SccV2FolderMuteConfigInput interface {
	pulumi.Input

	ToSccV2FolderMuteConfigOutput() SccV2FolderMuteConfigOutput
	ToSccV2FolderMuteConfigOutputWithContext(ctx context.Context) SccV2FolderMuteConfigOutput
}

func (*SccV2FolderMuteConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2FolderMuteConfig)(nil)).Elem()
}

func (i *SccV2FolderMuteConfig) ToSccV2FolderMuteConfigOutput() SccV2FolderMuteConfigOutput {
	return i.ToSccV2FolderMuteConfigOutputWithContext(context.Background())
}

func (i *SccV2FolderMuteConfig) ToSccV2FolderMuteConfigOutputWithContext(ctx context.Context) SccV2FolderMuteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccV2FolderMuteConfigOutput)
}

type SccV2FolderMuteConfigOutput struct{ *pulumi.OutputState }

func (SccV2FolderMuteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2FolderMuteConfig)(nil)).Elem()
}

func (o SccV2FolderMuteConfigOutput) ToSccV2FolderMuteConfigOutput() SccV2FolderMuteConfigOutput {
	return o
}

func (o SccV2FolderMuteConfigOutput) ToSccV2FolderMuteConfigOutputWithContext(ctx context.Context) SccV2FolderMuteConfigOutput {
	return o
}

// The time at which the mute config was created. This field is set by the server and will be ignored if provided on config
// creation.
func (o SccV2FolderMuteConfigOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of the mute config.
func (o SccV2FolderMuteConfigOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An expression that defines the filter to apply across create/update events of findings. While creating a filter string,
// be mindful of the scope in which the mute configuration is being created. E.g., If a filter contains project = X but is
// created under the project = Y scope, it might not match any findings.
func (o SccV2FolderMuteConfigOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.Filter }).(pulumi.StringOutput)
}

// The folder whose Cloud Security Command Center the Mute Config lives in.
func (o SccV2FolderMuteConfigOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// location Id is provided by folder. If not provided, Use global as default.
func (o SccV2FolderMuteConfigOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Email address of the user who last edited the mute config. This field is set by the server and will be ignored if
// provided on config creation or update.
func (o SccV2FolderMuteConfigOutput) MostRecentEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.MostRecentEditor }).(pulumi.StringOutput)
}

// Unique identifier provided by the client within the parent scope.
func (o SccV2FolderMuteConfigOutput) MuteConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.MuteConfigId }).(pulumi.StringOutput)
}

// Name of the mute config. Its format is organizations/{organization}/locations/global/muteConfigs/{configId},
// folders/{folder}/locations/global/muteConfigs/{configId}, or projects/{project}/locations/global/muteConfigs/{configId}
func (o SccV2FolderMuteConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SccV2FolderMuteConfigOutput) SccV2FolderMuteConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.SccV2FolderMuteConfigId }).(pulumi.StringOutput)
}

func (o SccV2FolderMuteConfigOutput) Timeouts() SccV2FolderMuteConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) SccV2FolderMuteConfigTimeoutsPtrOutput { return v.Timeouts }).(SccV2FolderMuteConfigTimeoutsPtrOutput)
}

// The type of the mute config.
func (o SccV2FolderMuteConfigOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Output only. The most recent time at which the mute config was updated. This field is set by the server and will be
// ignored if provided on config creation or update.
func (o SccV2FolderMuteConfigOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderMuteConfig) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccV2FolderMuteConfigInput)(nil)).Elem(), &SccV2FolderMuteConfig{})
	pulumi.RegisterOutputType(SccV2FolderMuteConfigOutput{})
}
