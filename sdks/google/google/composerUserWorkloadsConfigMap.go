// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComposerUserWorkloadsConfigMap struct {
	pulumi.CustomResourceState

	ComposerUserWorkloadsConfigMapId pulumi.StringOutput `pulumi:"composerUserWorkloadsConfigMapId"`
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
	// https://kubernetes.io/docs/concepts/configuration/configmap/
	Data pulumi.StringMapOutput `pulumi:"data"`
	// Environment where the Kubernetes ConfigMap will be stored and used.
	Environment pulumi.StringOutput `pulumi:"environment"`
	// Name of the Kubernetes ConfigMap.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region   pulumi.StringOutput                             `pulumi:"region"`
	Timeouts ComposerUserWorkloadsConfigMapTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComposerUserWorkloadsConfigMap registers a new resource with the given unique name, arguments, and options.
func NewComposerUserWorkloadsConfigMap(ctx *pulumi.Context,
	name string, args *ComposerUserWorkloadsConfigMapArgs, opts ...pulumi.ResourceOption) (*ComposerUserWorkloadsConfigMap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Environment == nil {
		return nil, errors.New("invalid value for required argument 'Environment'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComposerUserWorkloadsConfigMap
	err = ctx.RegisterPackageResource("google:index/composerUserWorkloadsConfigMap:ComposerUserWorkloadsConfigMap", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComposerUserWorkloadsConfigMap gets an existing ComposerUserWorkloadsConfigMap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComposerUserWorkloadsConfigMap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComposerUserWorkloadsConfigMapState, opts ...pulumi.ResourceOption) (*ComposerUserWorkloadsConfigMap, error) {
	var resource ComposerUserWorkloadsConfigMap
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/composerUserWorkloadsConfigMap:ComposerUserWorkloadsConfigMap", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComposerUserWorkloadsConfigMap resources.
type composerUserWorkloadsConfigMapState struct {
	ComposerUserWorkloadsConfigMapId *string `pulumi:"composerUserWorkloadsConfigMapId"`
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
	// https://kubernetes.io/docs/concepts/configuration/configmap/
	Data map[string]string `pulumi:"data"`
	// Environment where the Kubernetes ConfigMap will be stored and used.
	Environment *string `pulumi:"environment"`
	// Name of the Kubernetes ConfigMap.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region   *string                                 `pulumi:"region"`
	Timeouts *ComposerUserWorkloadsConfigMapTimeouts `pulumi:"timeouts"`
}

type ComposerUserWorkloadsConfigMapState struct {
	ComposerUserWorkloadsConfigMapId pulumi.StringPtrInput
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
	// https://kubernetes.io/docs/concepts/configuration/configmap/
	Data pulumi.StringMapInput
	// Environment where the Kubernetes ConfigMap will be stored and used.
	Environment pulumi.StringPtrInput
	// Name of the Kubernetes ConfigMap.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region   pulumi.StringPtrInput
	Timeouts ComposerUserWorkloadsConfigMapTimeoutsPtrInput
}

func (ComposerUserWorkloadsConfigMapState) ElementType() reflect.Type {
	return reflect.TypeOf((*composerUserWorkloadsConfigMapState)(nil)).Elem()
}

type composerUserWorkloadsConfigMapArgs struct {
	ComposerUserWorkloadsConfigMapId *string `pulumi:"composerUserWorkloadsConfigMapId"`
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
	// https://kubernetes.io/docs/concepts/configuration/configmap/
	Data map[string]string `pulumi:"data"`
	// Environment where the Kubernetes ConfigMap will be stored and used.
	Environment string `pulumi:"environment"`
	// Name of the Kubernetes ConfigMap.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The location or Compute Engine region for the environment.
	Region   *string                                 `pulumi:"region"`
	Timeouts *ComposerUserWorkloadsConfigMapTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComposerUserWorkloadsConfigMap resource.
type ComposerUserWorkloadsConfigMapArgs struct {
	ComposerUserWorkloadsConfigMapId pulumi.StringPtrInput
	// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
	// https://kubernetes.io/docs/concepts/configuration/configmap/
	Data pulumi.StringMapInput
	// Environment where the Kubernetes ConfigMap will be stored and used.
	Environment pulumi.StringInput
	// Name of the Kubernetes ConfigMap.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The location or Compute Engine region for the environment.
	Region   pulumi.StringPtrInput
	Timeouts ComposerUserWorkloadsConfigMapTimeoutsPtrInput
}

func (ComposerUserWorkloadsConfigMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*composerUserWorkloadsConfigMapArgs)(nil)).Elem()
}

type ComposerUserWorkloadsConfigMapInput interface {
	pulumi.Input

	ToComposerUserWorkloadsConfigMapOutput() ComposerUserWorkloadsConfigMapOutput
	ToComposerUserWorkloadsConfigMapOutputWithContext(ctx context.Context) ComposerUserWorkloadsConfigMapOutput
}

func (*ComposerUserWorkloadsConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((**ComposerUserWorkloadsConfigMap)(nil)).Elem()
}

func (i *ComposerUserWorkloadsConfigMap) ToComposerUserWorkloadsConfigMapOutput() ComposerUserWorkloadsConfigMapOutput {
	return i.ToComposerUserWorkloadsConfigMapOutputWithContext(context.Background())
}

func (i *ComposerUserWorkloadsConfigMap) ToComposerUserWorkloadsConfigMapOutputWithContext(ctx context.Context) ComposerUserWorkloadsConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComposerUserWorkloadsConfigMapOutput)
}

type ComposerUserWorkloadsConfigMapOutput struct{ *pulumi.OutputState }

func (ComposerUserWorkloadsConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComposerUserWorkloadsConfigMap)(nil)).Elem()
}

func (o ComposerUserWorkloadsConfigMapOutput) ToComposerUserWorkloadsConfigMapOutput() ComposerUserWorkloadsConfigMapOutput {
	return o
}

func (o ComposerUserWorkloadsConfigMapOutput) ToComposerUserWorkloadsConfigMapOutputWithContext(ctx context.Context) ComposerUserWorkloadsConfigMapOutput {
	return o
}

func (o ComposerUserWorkloadsConfigMapOutput) ComposerUserWorkloadsConfigMapId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringOutput { return v.ComposerUserWorkloadsConfigMapId }).(pulumi.StringOutput)
}

// The "data" field of Kubernetes ConfigMap, organized in key-value pairs. For details see:
// https://kubernetes.io/docs/concepts/configuration/configmap/
func (o ComposerUserWorkloadsConfigMapOutput) Data() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringMapOutput { return v.Data }).(pulumi.StringMapOutput)
}

// Environment where the Kubernetes ConfigMap will be stored and used.
func (o ComposerUserWorkloadsConfigMapOutput) Environment() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringOutput { return v.Environment }).(pulumi.StringOutput)
}

// Name of the Kubernetes ConfigMap.
func (o ComposerUserWorkloadsConfigMapOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComposerUserWorkloadsConfigMapOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The location or Compute Engine region for the environment.
func (o ComposerUserWorkloadsConfigMapOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o ComposerUserWorkloadsConfigMapOutput) Timeouts() ComposerUserWorkloadsConfigMapTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComposerUserWorkloadsConfigMap) ComposerUserWorkloadsConfigMapTimeoutsPtrOutput {
		return v.Timeouts
	}).(ComposerUserWorkloadsConfigMapTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComposerUserWorkloadsConfigMapInput)(nil)).Elem(), &ComposerUserWorkloadsConfigMap{})
	pulumi.RegisterOutputType(ComposerUserWorkloadsConfigMapOutput{})
}
