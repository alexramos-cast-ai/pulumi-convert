// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TpuV2QueuedResource struct {
	pulumi.CustomResourceState

	// The immutable name of the Queued Resource.
	Name     pulumi.StringOutput                  `pulumi:"name"`
	Project  pulumi.StringOutput                  `pulumi:"project"`
	Timeouts TpuV2QueuedResourceTimeoutsPtrOutput `pulumi:"timeouts"`
	// Defines a TPU resource.
	Tpu                   TpuV2QueuedResourceTpuPtrOutput `pulumi:"tpu"`
	TpuV2QueuedResourceId pulumi.StringOutput             `pulumi:"tpuV2QueuedResourceId"`
	// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTpuV2QueuedResource registers a new resource with the given unique name, arguments, and options.
func NewTpuV2QueuedResource(ctx *pulumi.Context,
	name string, args *TpuV2QueuedResourceArgs, opts ...pulumi.ResourceOption) (*TpuV2QueuedResource, error) {
	if args == nil {
		args = &TpuV2QueuedResourceArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource TpuV2QueuedResource
	err = ctx.RegisterPackageResource("google-beta:index/tpuV2QueuedResource:TpuV2QueuedResource", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTpuV2QueuedResource gets an existing TpuV2QueuedResource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTpuV2QueuedResource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TpuV2QueuedResourceState, opts ...pulumi.ResourceOption) (*TpuV2QueuedResource, error) {
	var resource TpuV2QueuedResource
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/tpuV2QueuedResource:TpuV2QueuedResource", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TpuV2QueuedResource resources.
type tpuV2QueuedResourceState struct {
	// The immutable name of the Queued Resource.
	Name     *string                      `pulumi:"name"`
	Project  *string                      `pulumi:"project"`
	Timeouts *TpuV2QueuedResourceTimeouts `pulumi:"timeouts"`
	// Defines a TPU resource.
	Tpu                   *TpuV2QueuedResourceTpu `pulumi:"tpu"`
	TpuV2QueuedResourceId *string                 `pulumi:"tpuV2QueuedResourceId"`
	// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type TpuV2QueuedResourceState struct {
	// The immutable name of the Queued Resource.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts TpuV2QueuedResourceTimeoutsPtrInput
	// Defines a TPU resource.
	Tpu                   TpuV2QueuedResourceTpuPtrInput
	TpuV2QueuedResourceId pulumi.StringPtrInput
	// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuV2QueuedResourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuV2QueuedResourceState)(nil)).Elem()
}

type tpuV2QueuedResourceArgs struct {
	// The immutable name of the Queued Resource.
	Name     *string                      `pulumi:"name"`
	Project  *string                      `pulumi:"project"`
	Timeouts *TpuV2QueuedResourceTimeouts `pulumi:"timeouts"`
	// Defines a TPU resource.
	Tpu                   *TpuV2QueuedResourceTpu `pulumi:"tpu"`
	TpuV2QueuedResourceId *string                 `pulumi:"tpuV2QueuedResourceId"`
	// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TpuV2QueuedResource resource.
type TpuV2QueuedResourceArgs struct {
	// The immutable name of the Queued Resource.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts TpuV2QueuedResourceTimeoutsPtrInput
	// Defines a TPU resource.
	Tpu                   TpuV2QueuedResourceTpuPtrInput
	TpuV2QueuedResourceId pulumi.StringPtrInput
	// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (TpuV2QueuedResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tpuV2QueuedResourceArgs)(nil)).Elem()
}

type TpuV2QueuedResourceInput interface {
	pulumi.Input

	ToTpuV2QueuedResourceOutput() TpuV2QueuedResourceOutput
	ToTpuV2QueuedResourceOutputWithContext(ctx context.Context) TpuV2QueuedResourceOutput
}

func (*TpuV2QueuedResource) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuV2QueuedResource)(nil)).Elem()
}

func (i *TpuV2QueuedResource) ToTpuV2QueuedResourceOutput() TpuV2QueuedResourceOutput {
	return i.ToTpuV2QueuedResourceOutputWithContext(context.Background())
}

func (i *TpuV2QueuedResource) ToTpuV2QueuedResourceOutputWithContext(ctx context.Context) TpuV2QueuedResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TpuV2QueuedResourceOutput)
}

type TpuV2QueuedResourceOutput struct{ *pulumi.OutputState }

func (TpuV2QueuedResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TpuV2QueuedResource)(nil)).Elem()
}

func (o TpuV2QueuedResourceOutput) ToTpuV2QueuedResourceOutput() TpuV2QueuedResourceOutput {
	return o
}

func (o TpuV2QueuedResourceOutput) ToTpuV2QueuedResourceOutputWithContext(ctx context.Context) TpuV2QueuedResourceOutput {
	return o
}

// The immutable name of the Queued Resource.
func (o TpuV2QueuedResourceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TpuV2QueuedResourceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o TpuV2QueuedResourceOutput) Timeouts() TpuV2QueuedResourceTimeoutsPtrOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) TpuV2QueuedResourceTimeoutsPtrOutput { return v.Timeouts }).(TpuV2QueuedResourceTimeoutsPtrOutput)
}

// Defines a TPU resource.
func (o TpuV2QueuedResourceOutput) Tpu() TpuV2QueuedResourceTpuPtrOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) TpuV2QueuedResourceTpuPtrOutput { return v.Tpu }).(TpuV2QueuedResourceTpuPtrOutput)
}

func (o TpuV2QueuedResourceOutput) TpuV2QueuedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) pulumi.StringOutput { return v.TpuV2QueuedResourceId }).(pulumi.StringOutput)
}

// The GCP location for the Queued Resource. If it is not provided, the provider zone is used.
func (o TpuV2QueuedResourceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TpuV2QueuedResource) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TpuV2QueuedResourceInput)(nil)).Elem(), &TpuV2QueuedResource{})
	pulumi.RegisterOutputType(TpuV2QueuedResourceOutput{})
}
