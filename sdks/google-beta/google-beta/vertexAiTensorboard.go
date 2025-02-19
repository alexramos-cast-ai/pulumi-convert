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

type VertexAiTensorboard struct {
	pulumi.CustomResourceState

	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not
	// end with a '/'.
	BlobStoragePathPrefix pulumi.StringOutput `pulumi:"blobStoragePathPrefix"`
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of this Tensorboard.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User provided name of this Tensorboard.
	DisplayName     pulumi.StringOutput    `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
	// Tensorboard will be secured by this key.
	EncryptionSpec VertexAiTensorboardEncryptionSpecPtrOutput `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the Tensorboard.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the tensorboard. eg us-central1
	Region pulumi.StringOutput `pulumi:"region"`
	// The number of Runs stored in this Tensorboard.
	RunCount pulumi.StringOutput `pulumi:"runCount"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput               `pulumi:"terraformLabels"`
	Timeouts        VertexAiTensorboardTimeoutsPtrOutput `pulumi:"timeouts"`
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime            pulumi.StringOutput `pulumi:"updateTime"`
	VertexAiTensorboardId pulumi.StringOutput `pulumi:"vertexAiTensorboardId"`
}

// NewVertexAiTensorboard registers a new resource with the given unique name, arguments, and options.
func NewVertexAiTensorboard(ctx *pulumi.Context,
	name string, args *VertexAiTensorboardArgs, opts ...pulumi.ResourceOption) (*VertexAiTensorboard, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource VertexAiTensorboard
	err = ctx.RegisterPackageResource("google-beta:index/vertexAiTensorboard:VertexAiTensorboard", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVertexAiTensorboard gets an existing VertexAiTensorboard resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVertexAiTensorboard(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VertexAiTensorboardState, opts ...pulumi.ResourceOption) (*VertexAiTensorboard, error) {
	var resource VertexAiTensorboard
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/vertexAiTensorboard:VertexAiTensorboard", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VertexAiTensorboard resources.
type vertexAiTensorboardState struct {
	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not
	// end with a '/'.
	BlobStoragePathPrefix *string `pulumi:"blobStoragePathPrefix"`
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// Description of this Tensorboard.
	Description *string `pulumi:"description"`
	// User provided name of this Tensorboard.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
	// Tensorboard will be secured by this key.
	EncryptionSpec *VertexAiTensorboardEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the Tensorboard.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The region of the tensorboard. eg us-central1
	Region *string `pulumi:"region"`
	// The number of Runs stored in this Tensorboard.
	RunCount *string `pulumi:"runCount"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string            `pulumi:"terraformLabels"`
	Timeouts        *VertexAiTensorboardTimeouts `pulumi:"timeouts"`
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime            *string `pulumi:"updateTime"`
	VertexAiTensorboardId *string `pulumi:"vertexAiTensorboardId"`
}

type VertexAiTensorboardState struct {
	// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not
	// end with a '/'.
	BlobStoragePathPrefix pulumi.StringPtrInput
	// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// Description of this Tensorboard.
	Description pulumi.StringPtrInput
	// User provided name of this Tensorboard.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
	// Tensorboard will be secured by this key.
	EncryptionSpec VertexAiTensorboardEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Name of the Tensorboard.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The region of the tensorboard. eg us-central1
	Region pulumi.StringPtrInput
	// The number of Runs stored in this Tensorboard.
	RunCount pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        VertexAiTensorboardTimeoutsPtrInput
	// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
	// to nine fractional digits.
	UpdateTime            pulumi.StringPtrInput
	VertexAiTensorboardId pulumi.StringPtrInput
}

func (VertexAiTensorboardState) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiTensorboardState)(nil)).Elem()
}

type vertexAiTensorboardArgs struct {
	// Description of this Tensorboard.
	Description *string `pulumi:"description"`
	// User provided name of this Tensorboard.
	DisplayName string `pulumi:"displayName"`
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
	// Tensorboard will be secured by this key.
	EncryptionSpec *VertexAiTensorboardEncryptionSpec `pulumi:"encryptionSpec"`
	// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels  map[string]string `pulumi:"labels"`
	Project *string           `pulumi:"project"`
	// The region of the tensorboard. eg us-central1
	Region                *string                      `pulumi:"region"`
	Timeouts              *VertexAiTensorboardTimeouts `pulumi:"timeouts"`
	VertexAiTensorboardId *string                      `pulumi:"vertexAiTensorboardId"`
}

// The set of arguments for constructing a VertexAiTensorboard resource.
type VertexAiTensorboardArgs struct {
	// Description of this Tensorboard.
	Description pulumi.StringPtrInput
	// User provided name of this Tensorboard.
	DisplayName pulumi.StringInput
	// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
	// Tensorboard will be secured by this key.
	EncryptionSpec VertexAiTensorboardEncryptionSpecPtrInput
	// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
	// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels  pulumi.StringMapInput
	Project pulumi.StringPtrInput
	// The region of the tensorboard. eg us-central1
	Region                pulumi.StringPtrInput
	Timeouts              VertexAiTensorboardTimeoutsPtrInput
	VertexAiTensorboardId pulumi.StringPtrInput
}

func (VertexAiTensorboardArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vertexAiTensorboardArgs)(nil)).Elem()
}

type VertexAiTensorboardInput interface {
	pulumi.Input

	ToVertexAiTensorboardOutput() VertexAiTensorboardOutput
	ToVertexAiTensorboardOutputWithContext(ctx context.Context) VertexAiTensorboardOutput
}

func (*VertexAiTensorboard) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiTensorboard)(nil)).Elem()
}

func (i *VertexAiTensorboard) ToVertexAiTensorboardOutput() VertexAiTensorboardOutput {
	return i.ToVertexAiTensorboardOutputWithContext(context.Background())
}

func (i *VertexAiTensorboard) ToVertexAiTensorboardOutputWithContext(ctx context.Context) VertexAiTensorboardOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VertexAiTensorboardOutput)
}

type VertexAiTensorboardOutput struct{ *pulumi.OutputState }

func (VertexAiTensorboardOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VertexAiTensorboard)(nil)).Elem()
}

func (o VertexAiTensorboardOutput) ToVertexAiTensorboardOutput() VertexAiTensorboardOutput {
	return o
}

func (o VertexAiTensorboardOutput) ToVertexAiTensorboardOutputWithContext(ctx context.Context) VertexAiTensorboardOutput {
	return o
}

// Consumer project Cloud Storage path prefix used to store blob data, which can either be a bucket or directory. Does not
// end with a '/'.
func (o VertexAiTensorboardOutput) BlobStoragePathPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.BlobStoragePathPrefix }).(pulumi.StringOutput)
}

// The timestamp of when the Tensorboard was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits.
func (o VertexAiTensorboardOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of this Tensorboard.
func (o VertexAiTensorboardOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User provided name of this Tensorboard.
func (o VertexAiTensorboardOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o VertexAiTensorboardOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Customer-managed encryption key spec for a Tensorboard. If set, this Tensorboard and all sub-resources of this
// Tensorboard will be secured by this key.
func (o VertexAiTensorboardOutput) EncryptionSpec() VertexAiTensorboardEncryptionSpecPtrOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) VertexAiTensorboardEncryptionSpecPtrOutput { return v.EncryptionSpec }).(VertexAiTensorboardEncryptionSpecPtrOutput)
}

// The labels with user-defined metadata to organize your Tensorboards. **Note**: This field is non-authoritative, and will
// only manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o VertexAiTensorboardOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the Tensorboard.
func (o VertexAiTensorboardOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VertexAiTensorboardOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the tensorboard. eg us-central1
func (o VertexAiTensorboardOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The number of Runs stored in this Tensorboard.
func (o VertexAiTensorboardOutput) RunCount() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.RunCount }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o VertexAiTensorboardOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o VertexAiTensorboardOutput) Timeouts() VertexAiTensorboardTimeoutsPtrOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) VertexAiTensorboardTimeoutsPtrOutput { return v.Timeouts }).(VertexAiTensorboardTimeoutsPtrOutput)
}

// The timestamp of when the Tensorboard was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
// to nine fractional digits.
func (o VertexAiTensorboardOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func (o VertexAiTensorboardOutput) VertexAiTensorboardId() pulumi.StringOutput {
	return o.ApplyT(func(v *VertexAiTensorboard) pulumi.StringOutput { return v.VertexAiTensorboardId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VertexAiTensorboardInput)(nil)).Elem(), &VertexAiTensorboard{})
	pulumi.RegisterOutputType(VertexAiTensorboardOutput{})
}
