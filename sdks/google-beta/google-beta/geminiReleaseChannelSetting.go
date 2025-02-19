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

type GeminiReleaseChannelSetting struct {
	pulumi.CustomResourceState

	// Output only. [Output only] Create time stamp.
	CreateTime                    pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels               pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	GeminiReleaseChannelSettingId pulumi.StringOutput    `pulumi:"geminiReleaseChannelSettingId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122.
	Location pulumi.StringOutput `pulumi:"location"`
	// Identifier. Name of the resource.
	// Format:projects/{project}/locations/{location}/releaseChannelSettings/{releaseChannelSetting}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
	ReleaseChannel pulumi.StringPtrOutput `pulumi:"releaseChannel"`
	// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
	// release_channel_setting_id from the method_signature of Create RPC
	ReleaseChannelSettingId pulumi.StringOutput `pulumi:"releaseChannelSettingId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                       `pulumi:"terraformLabels"`
	Timeouts        GeminiReleaseChannelSettingTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGeminiReleaseChannelSetting registers a new resource with the given unique name, arguments, and options.
func NewGeminiReleaseChannelSetting(ctx *pulumi.Context,
	name string, args *GeminiReleaseChannelSettingArgs, opts ...pulumi.ResourceOption) (*GeminiReleaseChannelSetting, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ReleaseChannelSettingId == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseChannelSettingId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GeminiReleaseChannelSetting
	err = ctx.RegisterPackageResource("google-beta:index/geminiReleaseChannelSetting:GeminiReleaseChannelSetting", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGeminiReleaseChannelSetting gets an existing GeminiReleaseChannelSetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGeminiReleaseChannelSetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeminiReleaseChannelSettingState, opts ...pulumi.ResourceOption) (*GeminiReleaseChannelSetting, error) {
	var resource GeminiReleaseChannelSetting
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/geminiReleaseChannelSetting:GeminiReleaseChannelSetting", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GeminiReleaseChannelSetting resources.
type geminiReleaseChannelSettingState struct {
	// Output only. [Output only] Create time stamp.
	CreateTime                    *string           `pulumi:"createTime"`
	EffectiveLabels               map[string]string `pulumi:"effectiveLabels"`
	GeminiReleaseChannelSettingId *string           `pulumi:"geminiReleaseChannelSettingId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122.
	Location *string `pulumi:"location"`
	// Identifier. Name of the resource.
	// Format:projects/{project}/locations/{location}/releaseChannelSettings/{releaseChannelSetting}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
	// release_channel_setting_id from the method_signature of Create RPC
	ReleaseChannelSettingId *string `pulumi:"releaseChannelSettingId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                    `pulumi:"terraformLabels"`
	Timeouts        *GeminiReleaseChannelSettingTimeouts `pulumi:"timeouts"`
	// Output only. [Output only] Update time stamp.
	UpdateTime *string `pulumi:"updateTime"`
}

type GeminiReleaseChannelSettingState struct {
	// Output only. [Output only] Create time stamp.
	CreateTime                    pulumi.StringPtrInput
	EffectiveLabels               pulumi.StringMapInput
	GeminiReleaseChannelSettingId pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122.
	Location pulumi.StringPtrInput
	// Identifier. Name of the resource.
	// Format:projects/{project}/locations/{location}/releaseChannelSettings/{releaseChannelSetting}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
	ReleaseChannel pulumi.StringPtrInput
	// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
	// release_channel_setting_id from the method_signature of Create RPC
	ReleaseChannelSettingId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        GeminiReleaseChannelSettingTimeoutsPtrInput
	// Output only. [Output only] Update time stamp.
	UpdateTime pulumi.StringPtrInput
}

func (GeminiReleaseChannelSettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*geminiReleaseChannelSettingState)(nil)).Elem()
}

type geminiReleaseChannelSettingArgs struct {
	GeminiReleaseChannelSettingId *string `pulumi:"geminiReleaseChannelSettingId"`
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122.
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
	// release_channel_setting_id from the method_signature of Create RPC
	ReleaseChannelSettingId string                               `pulumi:"releaseChannelSettingId"`
	Timeouts                *GeminiReleaseChannelSettingTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a GeminiReleaseChannelSetting resource.
type GeminiReleaseChannelSettingArgs struct {
	GeminiReleaseChannelSettingId pulumi.StringPtrInput
	// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
	// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
	// https://google.aip.dev/122.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
	ReleaseChannel pulumi.StringPtrInput
	// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
	// release_channel_setting_id from the method_signature of Create RPC
	ReleaseChannelSettingId pulumi.StringInput
	Timeouts                GeminiReleaseChannelSettingTimeoutsPtrInput
}

func (GeminiReleaseChannelSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*geminiReleaseChannelSettingArgs)(nil)).Elem()
}

type GeminiReleaseChannelSettingInput interface {
	pulumi.Input

	ToGeminiReleaseChannelSettingOutput() GeminiReleaseChannelSettingOutput
	ToGeminiReleaseChannelSettingOutputWithContext(ctx context.Context) GeminiReleaseChannelSettingOutput
}

func (*GeminiReleaseChannelSetting) ElementType() reflect.Type {
	return reflect.TypeOf((**GeminiReleaseChannelSetting)(nil)).Elem()
}

func (i *GeminiReleaseChannelSetting) ToGeminiReleaseChannelSettingOutput() GeminiReleaseChannelSettingOutput {
	return i.ToGeminiReleaseChannelSettingOutputWithContext(context.Background())
}

func (i *GeminiReleaseChannelSetting) ToGeminiReleaseChannelSettingOutputWithContext(ctx context.Context) GeminiReleaseChannelSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeminiReleaseChannelSettingOutput)
}

type GeminiReleaseChannelSettingOutput struct{ *pulumi.OutputState }

func (GeminiReleaseChannelSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GeminiReleaseChannelSetting)(nil)).Elem()
}

func (o GeminiReleaseChannelSettingOutput) ToGeminiReleaseChannelSettingOutput() GeminiReleaseChannelSettingOutput {
	return o
}

func (o GeminiReleaseChannelSettingOutput) ToGeminiReleaseChannelSettingOutputWithContext(ctx context.Context) GeminiReleaseChannelSettingOutput {
	return o
}

// Output only. [Output only] Create time stamp.
func (o GeminiReleaseChannelSettingOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o GeminiReleaseChannelSettingOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o GeminiReleaseChannelSettingOutput) GeminiReleaseChannelSettingId() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.GeminiReleaseChannelSettingId }).(pulumi.StringOutput)
}

// Optional. Labels as key value pairs. **Note**: This field is non-authoritative, and will only manage the labels present
// in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o GeminiReleaseChannelSettingOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Resource ID segment making up resource 'name'. It identifies the resource within its parent collection as described in
// https://google.aip.dev/122.
func (o GeminiReleaseChannelSettingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Identifier. Name of the resource.
// Format:projects/{project}/locations/{location}/releaseChannelSettings/{releaseChannelSetting}
func (o GeminiReleaseChannelSettingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GeminiReleaseChannelSettingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. Release channel to be used. Possible values: STABLE EXPERIMENTAL
func (o GeminiReleaseChannelSettingOutput) ReleaseChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringPtrOutput { return v.ReleaseChannel }).(pulumi.StringPtrOutput)
}

// Required. Id of the requesting object. If auto-generating Id server-side, remove this field and
// release_channel_setting_id from the method_signature of Create RPC
func (o GeminiReleaseChannelSettingOutput) ReleaseChannelSettingId() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.ReleaseChannelSettingId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o GeminiReleaseChannelSettingOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o GeminiReleaseChannelSettingOutput) Timeouts() GeminiReleaseChannelSettingTimeoutsPtrOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) GeminiReleaseChannelSettingTimeoutsPtrOutput { return v.Timeouts }).(GeminiReleaseChannelSettingTimeoutsPtrOutput)
}

// Output only. [Output only] Update time stamp.
func (o GeminiReleaseChannelSettingOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GeminiReleaseChannelSetting) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeminiReleaseChannelSettingInput)(nil)).Elem(), &GeminiReleaseChannelSetting{})
	pulumi.RegisterOutputType(GeminiReleaseChannelSettingOutput{})
}
