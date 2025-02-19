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

type ComputeProjectCloudArmorTier struct {
	pulumi.CustomResourceState

	// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
	CloudArmorTier                 pulumi.StringOutput                           `pulumi:"cloudArmorTier"`
	ComputeProjectCloudArmorTierId pulumi.StringOutput                           `pulumi:"computeProjectCloudArmorTierId"`
	Project                        pulumi.StringOutput                           `pulumi:"project"`
	Timeouts                       ComputeProjectCloudArmorTierTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeProjectCloudArmorTier registers a new resource with the given unique name, arguments, and options.
func NewComputeProjectCloudArmorTier(ctx *pulumi.Context,
	name string, args *ComputeProjectCloudArmorTierArgs, opts ...pulumi.ResourceOption) (*ComputeProjectCloudArmorTier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudArmorTier == nil {
		return nil, errors.New("invalid value for required argument 'CloudArmorTier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeProjectCloudArmorTier
	err = ctx.RegisterPackageResource("google-beta:index/computeProjectCloudArmorTier:ComputeProjectCloudArmorTier", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeProjectCloudArmorTier gets an existing ComputeProjectCloudArmorTier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeProjectCloudArmorTier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeProjectCloudArmorTierState, opts ...pulumi.ResourceOption) (*ComputeProjectCloudArmorTier, error) {
	var resource ComputeProjectCloudArmorTier
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/computeProjectCloudArmorTier:ComputeProjectCloudArmorTier", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeProjectCloudArmorTier resources.
type computeProjectCloudArmorTierState struct {
	// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
	CloudArmorTier                 *string                               `pulumi:"cloudArmorTier"`
	ComputeProjectCloudArmorTierId *string                               `pulumi:"computeProjectCloudArmorTierId"`
	Project                        *string                               `pulumi:"project"`
	Timeouts                       *ComputeProjectCloudArmorTierTimeouts `pulumi:"timeouts"`
}

type ComputeProjectCloudArmorTierState struct {
	// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
	CloudArmorTier                 pulumi.StringPtrInput
	ComputeProjectCloudArmorTierId pulumi.StringPtrInput
	Project                        pulumi.StringPtrInput
	Timeouts                       ComputeProjectCloudArmorTierTimeoutsPtrInput
}

func (ComputeProjectCloudArmorTierState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeProjectCloudArmorTierState)(nil)).Elem()
}

type computeProjectCloudArmorTierArgs struct {
	// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
	CloudArmorTier                 string                                `pulumi:"cloudArmorTier"`
	ComputeProjectCloudArmorTierId *string                               `pulumi:"computeProjectCloudArmorTierId"`
	Project                        *string                               `pulumi:"project"`
	Timeouts                       *ComputeProjectCloudArmorTierTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeProjectCloudArmorTier resource.
type ComputeProjectCloudArmorTierArgs struct {
	// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
	CloudArmorTier                 pulumi.StringInput
	ComputeProjectCloudArmorTierId pulumi.StringPtrInput
	Project                        pulumi.StringPtrInput
	Timeouts                       ComputeProjectCloudArmorTierTimeoutsPtrInput
}

func (ComputeProjectCloudArmorTierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeProjectCloudArmorTierArgs)(nil)).Elem()
}

type ComputeProjectCloudArmorTierInput interface {
	pulumi.Input

	ToComputeProjectCloudArmorTierOutput() ComputeProjectCloudArmorTierOutput
	ToComputeProjectCloudArmorTierOutputWithContext(ctx context.Context) ComputeProjectCloudArmorTierOutput
}

func (*ComputeProjectCloudArmorTier) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeProjectCloudArmorTier)(nil)).Elem()
}

func (i *ComputeProjectCloudArmorTier) ToComputeProjectCloudArmorTierOutput() ComputeProjectCloudArmorTierOutput {
	return i.ToComputeProjectCloudArmorTierOutputWithContext(context.Background())
}

func (i *ComputeProjectCloudArmorTier) ToComputeProjectCloudArmorTierOutputWithContext(ctx context.Context) ComputeProjectCloudArmorTierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeProjectCloudArmorTierOutput)
}

type ComputeProjectCloudArmorTierOutput struct{ *pulumi.OutputState }

func (ComputeProjectCloudArmorTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeProjectCloudArmorTier)(nil)).Elem()
}

func (o ComputeProjectCloudArmorTierOutput) ToComputeProjectCloudArmorTierOutput() ComputeProjectCloudArmorTierOutput {
	return o
}

func (o ComputeProjectCloudArmorTierOutput) ToComputeProjectCloudArmorTierOutputWithContext(ctx context.Context) ComputeProjectCloudArmorTierOutput {
	return o
}

// Managed protection tier to be set. Possible values: ["CA_STANDARD", "CA_ENTERPRISE_PAYGO", "CA_ENTERPRISE_ANNUAL"]
func (o ComputeProjectCloudArmorTierOutput) CloudArmorTier() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeProjectCloudArmorTier) pulumi.StringOutput { return v.CloudArmorTier }).(pulumi.StringOutput)
}

func (o ComputeProjectCloudArmorTierOutput) ComputeProjectCloudArmorTierId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeProjectCloudArmorTier) pulumi.StringOutput { return v.ComputeProjectCloudArmorTierId }).(pulumi.StringOutput)
}

func (o ComputeProjectCloudArmorTierOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeProjectCloudArmorTier) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeProjectCloudArmorTierOutput) Timeouts() ComputeProjectCloudArmorTierTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeProjectCloudArmorTier) ComputeProjectCloudArmorTierTimeoutsPtrOutput { return v.Timeouts }).(ComputeProjectCloudArmorTierTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeProjectCloudArmorTierInput)(nil)).Elem(), &ComputeProjectCloudArmorTier{})
	pulumi.RegisterOutputType(ComputeProjectCloudArmorTierOutput{})
}
