// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VertexAiFeaturestore struct {
	pulumi.CustomResourceState

	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime      pulumi.StringOutput    `pulumi:"createTime"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec VertexAiFeaturestoreEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
	// character cannot be a number.
	Name pulumi.StringOutput `pulumi:"name"`
	// Config for online serving resources.
	OnlineServingConfig VertexAiFeaturestoreOnlineServingConfigPtrOutput `pulumi:"onlineServingConfig"`
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
	// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
	// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays pulumi.Float64PtrOutput `pulumi:"onlineStorageTtlDays"`
	Project              pulumi.StringOutput     `pulumi:"project"`
	// The region of the dataset. eg us-central1
	Region pulumi.StringOutput `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                `pulumi:"terraformLabels"`
	Timeouts        VertexAiFeaturestoreTimeoutsPtrOutput `pulumi:"timeouts"`
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime             pulumi.StringOutput `pulumi:"updateTime"`
	VertexAiFeaturestoreId pulumi.StringOutput `pulumi:"vertexAiFeaturestoreId"`
}

// NewVertexAiFeaturestore registers a new resource with the given unique name, arguments, and options.
func NewVertexAiFeaturestore(ctx *pulumi.Context,
	name string, args *VertexAiFeaturestoreArgs, opts ...pulumi.ResourceOption) (*VertexAiFeaturestore, error) {
	if args == nil {
		args = &VertexAiFeaturestoreArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VertexAiFeaturestore
	err = ctx.RegisterPackageResource("google-beta:index/vertexAiFeaturestore:VertexAiFeaturestore", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVertexAiFeaturestore gets an existing VertexAiFeaturestore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVertexAiFeaturestore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VertexAiFeaturestoreState, opts ...pulumi.ResourceOption) (*VertexAiFeaturestore, error) {
	var resource VertexAiFeaturestore
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/vertexAiFeaturestore:VertexAiFeaturestore", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VertexAiFeaturestore resources.
type vertexAiFeaturestoreState struct {
	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime      *string           `pulumi:"createTime"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec *VertexAiFeaturestoreEncryptionSpec `pulumi:"encryptionSpec"`
	// Used to perform consistent read-modify-write updates.
	Etag *string `pulumi:"etag"`
	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
	// character cannot be a number.
	Name *string `pulumi:"name"`
	// Config for online serving resources.
	OnlineServingConfig *VertexAiFeaturestoreOnlineServingConfig `pulumi:"onlineServingConfig"`
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
	// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
	// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays *float64 `pulumi:"onlineStorageTtlDays"`
	Project              *string  `pulumi:"project"`
	// The region of the dataset. eg us-central1
	Region *string `pulumi:"region"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string             `pulumi:"terraformLabels"`
	Timeouts        *VertexAiFeaturestoreTimeouts `pulumi:"timeouts"`
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime             *string `pulumi:"updateTime"`
	VertexAiFeaturestoreId *string `pulumi:"vertexAiFeaturestoreId"`
}

type VertexAiFeaturestoreState struct {
	// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime      pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec VertexAiFeaturestoreEncryptionSpecPtrInput
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringPtrInput
	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy pulumi.BoolPtrInput
	// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
	// character cannot be a number.
	Name pulumi.StringPtrInput
	// Config for online serving resources.
	OnlineServingConfig VertexAiFeaturestoreOnlineServingConfigPtrInput
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
	// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
	// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays pulumi.Float64PtrInput
	Project              pulumi.StringPtrInput
	// The region of the dataset. eg us-central1
	Region pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        VertexAiFeaturestoreTimeoutsPtrInput
	// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime             pulumi.StringPtrInput
	VertexAiFeaturestoreId pulumi.StringPtrInput
}

func (VertexAiFeaturestoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiFeaturestoreState)(nil)).Elem()
}

type vertexAiFeaturestoreArgs struct {
	// If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec *VertexAiFeaturestoreEncryptionSpec `pulumi:"encryptionSpec"`
	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
	// character cannot be a number.
	Name *string `pulumi:"name"`
	// Config for online serving resources.
	OnlineServingConfig *VertexAiFeaturestoreOnlineServingConfig `pulumi:"onlineServingConfig"`
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
	// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
	// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays *float64 `pulumi:"onlineStorageTtlDays"`
	Project              *string  `pulumi:"project"`
	// The region of the dataset. eg us-central1
	Region                 *string                       `pulumi:"region"`
	Timeouts               *VertexAiFeaturestoreTimeouts `pulumi:"timeouts"`
	VertexAiFeaturestoreId *string                       `pulumi:"vertexAiFeaturestoreId"`
}

// The set of arguments for constructing a VertexAiFeaturestore resource.
type VertexAiFeaturestoreArgs struct {
	// If set, both of the online and offline data storage will be secured by this key.
	EncryptionSpec VertexAiFeaturestoreEncryptionSpecPtrInput
	// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
	ForceDestroy pulumi.BoolPtrInput
	// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
	// character cannot be a number.
	Name pulumi.StringPtrInput
	// Config for online serving resources.
	OnlineServingConfig VertexAiFeaturestoreOnlineServingConfigPtrInput
	// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
	// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
	// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
	// featurestore. If not set, default to 4000 days
	OnlineStorageTtlDays pulumi.Float64PtrInput
	Project              pulumi.StringPtrInput
	// The region of the dataset. eg us-central1
	Region                 pulumi.StringPtrInput
	Timeouts               VertexAiFeaturestoreTimeoutsPtrInput
	VertexAiFeaturestoreId pulumi.StringPtrInput
}

func (VertexAiFeaturestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiFeaturestoreArgs)(nil)).Elem()
}

type VertexAiFeaturestoreInput interface {
	pulumi.Input

	ToVertexAiFeaturestoreOutput() VertexAiFeaturestoreOutput
	ToVertexAiFeaturestoreOutputWithContext(ctx context.Context) VertexAiFeaturestoreOutput
}

func (*VertexAiFeaturestore) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiFeaturestore)(nil)).Elem()
}

func (i *VertexAiFeaturestore) ToVertexAiFeaturestoreOutput() VertexAiFeaturestoreOutput {
	return i.ToVertexAiFeaturestoreOutputWithContext(context.Background())
}

func (i *VertexAiFeaturestore) ToVertexAiFeaturestoreOutputWithContext(ctx context.Context) VertexAiFeaturestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VertexAiFeaturestoreOutput)
}

type VertexAiFeaturestoreOutput struct{ *pulumi.OutputState }

func (VertexAiFeaturestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiFeaturestore)(nil)).Elem()
}

func (o VertexAiFeaturestoreOutput) ToVertexAiFeaturestoreOutput() VertexAiFeaturestoreOutput {
	return o
}

func (o VertexAiFeaturestoreOutput) ToVertexAiFeaturestoreOutputWithContext(ctx context.Context) VertexAiFeaturestoreOutput {
	return o
}

// The timestamp of when the featurestore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits.
func (o VertexAiFeaturestoreOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o VertexAiFeaturestoreOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// If set, both of the online and offline data storage will be secured by this key.
func (o VertexAiFeaturestoreOutput) EncryptionSpec() VertexAiFeaturestoreEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) VertexAiFeaturestoreEncryptionSpecPtrOutput { return v.EncryptionSpec }).(VertexAiFeaturestoreEncryptionSpecPtrOutput)
}

// Used to perform consistent read-modify-write updates.
func (o VertexAiFeaturestoreOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// If set to true, any EntityTypes and Features for this Featurestore will also be deleted
func (o VertexAiFeaturestoreOutput) ForceDestroy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.BoolPtrOutput { return v.ForceDestroy }).(pulumi.BoolPtrOutput)
}

// A set of key/value label pairs to assign to this Featurestore. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o VertexAiFeaturestoreOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the Featurestore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first
// character cannot be a number.
func (o VertexAiFeaturestoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Config for online serving resources.
func (o VertexAiFeaturestoreOutput) OnlineServingConfig() VertexAiFeaturestoreOnlineServingConfigPtrOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) VertexAiFeaturestoreOnlineServingConfigPtrOutput {
		return v.OnlineServingConfig
	}).(VertexAiFeaturestoreOnlineServingConfigPtrOutput)
}

// TTL in days for feature values that will be stored in online serving storage. The Feature Store online storage
// periodically removes obsolete feature values older than onlineStorageTtlDays since the feature generation time. Note
// that onlineStorageTtlDays should be less than or equal to offlineStorageTtlDays for each EntityType under a
// featurestore. If not set, default to 4000 days
func (o VertexAiFeaturestoreOutput) OnlineStorageTtlDays() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.Float64PtrOutput { return v.OnlineStorageTtlDays }).(pulumi.Float64PtrOutput)
}

func (o VertexAiFeaturestoreOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the dataset. eg us-central1
func (o VertexAiFeaturestoreOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o VertexAiFeaturestoreOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o VertexAiFeaturestoreOutput) Timeouts() VertexAiFeaturestoreTimeoutsPtrOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) VertexAiFeaturestoreTimeoutsPtrOutput { return v.Timeouts }).(VertexAiFeaturestoreTimeoutsPtrOutput)
}

// The timestamp of when the featurestore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
// to nine fractional digits.
func (o VertexAiFeaturestoreOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o VertexAiFeaturestoreOutput) VertexAiFeaturestoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiFeaturestore) pulumi.StringOutput { return v.VertexAiFeaturestoreId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VertexAiFeaturestoreInput)(nil)).Elem(), &VertexAiFeaturestore{})
	pulumi.RegisterOutputType(VertexAiFeaturestoreOutput{})
}
