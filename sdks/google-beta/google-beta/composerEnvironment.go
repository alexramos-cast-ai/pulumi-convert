// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComposerEnvironment struct {
	pulumi.CustomResourceState

	ComposerEnvironmentId pulumi.StringOutput `pulumi:"composerEnvironmentId"`
	// Configuration parameters for this environment.
	Config          ComposerEnvironmentConfigPtrOutput `pulumi:"config"`
	EffectiveLabels pulumi.StringMapOutput             `pulumi:"effectiveLabels"`
	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
	// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
	// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region pulumi.StringOutput `pulumi:"region"`
	// Configuration options for storage used by Composer environment.
	StorageConfig ComposerEnvironmentStorageConfigPtrOutput `pulumi:"storageConfig"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput               `pulumi:"terraformLabels"`
	Timeouts        ComposerEnvironmentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComposerEnvironment registers a new resource with the given unique name, arguments, and options.
func NewComposerEnvironment(ctx *pulumi.Context,
	name string, args *ComposerEnvironmentArgs, opts ...pulumi.ResourceOption) (*ComposerEnvironment, error) {
	if args == nil {
		args = &ComposerEnvironmentArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComposerEnvironment
	err = ctx.RegisterPackageResource("google-beta:index/composerEnvironment:ComposerEnvironment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComposerEnvironment gets an existing ComposerEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComposerEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComposerEnvironmentState, opts ...pulumi.ResourceOption) (*ComposerEnvironment, error) {
	var resource ComposerEnvironment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/composerEnvironment:ComposerEnvironment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComposerEnvironment resources.
type composerEnvironmentState struct {
	ComposerEnvironmentId *string `pulumi:"composerEnvironmentId"`
	// Configuration parameters for this environment.
	Config          *ComposerEnvironmentConfig `pulumi:"config"`
	EffectiveLabels map[string]string          `pulumi:"effectiveLabels"`
	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
	// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
	// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the environment.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region *string `pulumi:"region"`
	// Configuration options for storage used by Composer environment.
	StorageConfig *ComposerEnvironmentStorageConfig `pulumi:"storageConfig"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string            `pulumi:"terraformLabels"`
	Timeouts        *ComposerEnvironmentTimeouts `pulumi:"timeouts"`
}

type ComposerEnvironmentState struct {
	ComposerEnvironmentId pulumi.StringPtrInput
	// Configuration parameters for this environment.
	Config          ComposerEnvironmentConfigPtrInput
	EffectiveLabels pulumi.StringMapInput
	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
	// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
	// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of the environment.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region pulumi.StringPtrInput
	// Configuration options for storage used by Composer environment.
	StorageConfig ComposerEnvironmentStorageConfigPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComposerEnvironmentTimeoutsPtrInput
}

func (ComposerEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*composerEnvironmentState)(nil)).Elem()
}

type composerEnvironmentArgs struct {
	ComposerEnvironmentId *string `pulumi:"composerEnvironmentId"`
	// Configuration parameters for this environment.
	Config *ComposerEnvironmentConfig `pulumi:"config"`
	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
	// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
	// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the environment.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region *string `pulumi:"region"`
	// Configuration options for storage used by Composer environment.
	StorageConfig *ComposerEnvironmentStorageConfig `pulumi:"storageConfig"`
	Timeouts      *ComposerEnvironmentTimeouts      `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComposerEnvironment resource.
type ComposerEnvironmentArgs struct {
	ComposerEnvironmentId pulumi.StringPtrInput
	// Configuration parameters for this environment.
	Config ComposerEnvironmentConfigPtrInput
	// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
	// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
	// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
	// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
	// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Name of the environment.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region pulumi.StringPtrInput
	// Configuration options for storage used by Composer environment.
	StorageConfig ComposerEnvironmentStorageConfigPtrInput
	Timeouts      ComposerEnvironmentTimeoutsPtrInput
}

func (ComposerEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*composerEnvironmentArgs)(nil)).Elem()
}

type ComposerEnvironmentInput interface {
	pulumi.Input

	ToComposerEnvironmentOutput() ComposerEnvironmentOutput
	ToComposerEnvironmentOutputWithContext(ctx context.Context) ComposerEnvironmentOutput
}

func (*ComposerEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**ComposerEnvironment)(nil)).Elem()
}

func (i *ComposerEnvironment) ToComposerEnvironmentOutput() ComposerEnvironmentOutput {
	return i.ToComposerEnvironmentOutputWithContext(context.Background())
}

func (i *ComposerEnvironment) ToComposerEnvironmentOutputWithContext(ctx context.Context) ComposerEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComposerEnvironmentOutput)
}

type ComposerEnvironmentOutput struct{ *pulumi.OutputState }

func (ComposerEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComposerEnvironment)(nil)).Elem()
}

func (o ComposerEnvironmentOutput) ToComposerEnvironmentOutput() ComposerEnvironmentOutput {
	return o
}

func (o ComposerEnvironmentOutput) ToComposerEnvironmentOutputWithContext(ctx context.Context) ComposerEnvironmentOutput {
	return o
}

func (o ComposerEnvironmentOutput) ComposerEnvironmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringOutput { return v.ComposerEnvironmentId }).(pulumi.StringOutput)
}

// Configuration parameters for this environment.
func (o ComposerEnvironmentOutput) Config() ComposerEnvironmentConfigPtrOutput {
	return o.ApplyT(func(v *ComposerEnvironment) ComposerEnvironmentConfigPtrOutput { return v.Config }).(ComposerEnvironmentConfigPtrOutput)
}

func (o ComposerEnvironmentOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// User-defined labels for this environment. The labels map can contain no more than 64 entries. Entries of the labels map
// are UTF8 strings that comply with the following restrictions: Label keys must be between 1 and 63 characters long and
// must conform to the following regular expression: a-z?. Label values must be between 0 and 63 characters long and must
// conform to the regular expression (a-z?)?. No more than 64 labels can be associated with a given environment. Both keys
// and values must be <= 128 bytes in size. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o ComposerEnvironmentOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the environment.
func (o ComposerEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
func (o ComposerEnvironmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The location or Compute Engine region for the environment.
func (o ComposerEnvironmentOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// Configuration options for storage used by Composer environment.
func (o ComposerEnvironmentOutput) StorageConfig() ComposerEnvironmentStorageConfigPtrOutput {
	return o.ApplyT(func(v *ComposerEnvironment) ComposerEnvironmentStorageConfigPtrOutput { return v.StorageConfig }).(ComposerEnvironmentStorageConfigPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComposerEnvironmentOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComposerEnvironment) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComposerEnvironmentOutput) Timeouts() ComposerEnvironmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComposerEnvironment) ComposerEnvironmentTimeoutsPtrOutput { return v.Timeouts }).(ComposerEnvironmentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComposerEnvironmentInput)(nil)).Elem(), &ComposerEnvironment{})
	pulumi.RegisterOutputType(ComposerEnvironmentOutput{})
}
