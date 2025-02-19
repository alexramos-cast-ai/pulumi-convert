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

type PrivatecaCertificateTemplateIamBinding struct {
	pulumi.CustomResourceState

	CertificateTemplate                      pulumi.StringOutput                                      `pulumi:"certificateTemplate"`
	Condition                                PrivatecaCertificateTemplateIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                                     pulumi.StringOutput                                      `pulumi:"etag"`
	Location                                 pulumi.StringOutput                                      `pulumi:"location"`
	Members                                  pulumi.StringArrayOutput                                 `pulumi:"members"`
	PrivatecaCertificateTemplateIamBindingId pulumi.StringOutput                                      `pulumi:"privatecaCertificateTemplateIamBindingId"`
	Project                                  pulumi.StringOutput                                      `pulumi:"project"`
	Role                                     pulumi.StringOutput                                      `pulumi:"role"`
}

// NewPrivatecaCertificateTemplateIamBinding registers a new resource with the given unique name, arguments, and options.
func NewPrivatecaCertificateTemplateIamBinding(ctx *pulumi.Context,
	name string, args *PrivatecaCertificateTemplateIamBindingArgs, opts ...pulumi.ResourceOption) (*PrivatecaCertificateTemplateIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CertificateTemplate == nil {
		return nil, errors.New("invalid value for required argument 'CertificateTemplate'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource PrivatecaCertificateTemplateIamBinding
	err = ctx.RegisterPackageResource("google:index/privatecaCertificateTemplateIamBinding:PrivatecaCertificateTemplateIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivatecaCertificateTemplateIamBinding gets an existing PrivatecaCertificateTemplateIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivatecaCertificateTemplateIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivatecaCertificateTemplateIamBindingState, opts ...pulumi.ResourceOption) (*PrivatecaCertificateTemplateIamBinding, error) {
	var resource PrivatecaCertificateTemplateIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/privatecaCertificateTemplateIamBinding:PrivatecaCertificateTemplateIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivatecaCertificateTemplateIamBinding resources.
type privatecaCertificateTemplateIamBindingState struct {
	CertificateTemplate                      *string                                          `pulumi:"certificateTemplate"`
	Condition                                *PrivatecaCertificateTemplateIamBindingCondition `pulumi:"condition"`
	Etag                                     *string                                          `pulumi:"etag"`
	Location                                 *string                                          `pulumi:"location"`
	Members                                  []string                                         `pulumi:"members"`
	PrivatecaCertificateTemplateIamBindingId *string                                          `pulumi:"privatecaCertificateTemplateIamBindingId"`
	Project                                  *string                                          `pulumi:"project"`
	Role                                     *string                                          `pulumi:"role"`
}

type PrivatecaCertificateTemplateIamBindingState struct {
	CertificateTemplate                      pulumi.StringPtrInput
	Condition                                PrivatecaCertificateTemplateIamBindingConditionPtrInput
	Etag                                     pulumi.StringPtrInput
	Location                                 pulumi.StringPtrInput
	Members                                  pulumi.StringArrayInput
	PrivatecaCertificateTemplateIamBindingId pulumi.StringPtrInput
	Project                                  pulumi.StringPtrInput
	Role                                     pulumi.StringPtrInput
}

func (PrivatecaCertificateTemplateIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*privatecaCertificateTemplateIamBindingState)(nil)).Elem()
}

type privatecaCertificateTemplateIamBindingArgs struct {
	CertificateTemplate                      string                                           `pulumi:"certificateTemplate"`
	Condition                                *PrivatecaCertificateTemplateIamBindingCondition `pulumi:"condition"`
	Location                                 *string                                          `pulumi:"location"`
	Members                                  []string                                         `pulumi:"members"`
	PrivatecaCertificateTemplateIamBindingId *string                                          `pulumi:"privatecaCertificateTemplateIamBindingId"`
	Project                                  *string                                          `pulumi:"project"`
	Role                                     string                                           `pulumi:"role"`
}

// The set of arguments for constructing a PrivatecaCertificateTemplateIamBinding resource.
type PrivatecaCertificateTemplateIamBindingArgs struct {
	CertificateTemplate                      pulumi.StringInput
	Condition                                PrivatecaCertificateTemplateIamBindingConditionPtrInput
	Location                                 pulumi.StringPtrInput
	Members                                  pulumi.StringArrayInput
	PrivatecaCertificateTemplateIamBindingId pulumi.StringPtrInput
	Project                                  pulumi.StringPtrInput
	Role                                     pulumi.StringInput
}

func (PrivatecaCertificateTemplateIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privatecaCertificateTemplateIamBindingArgs)(nil)).Elem()
}

type PrivatecaCertificateTemplateIamBindingInput interface {
	pulumi.Input

	ToPrivatecaCertificateTemplateIamBindingOutput() PrivatecaCertificateTemplateIamBindingOutput
	ToPrivatecaCertificateTemplateIamBindingOutputWithContext(ctx context.Context) PrivatecaCertificateTemplateIamBindingOutput
}

func (*PrivatecaCertificateTemplateIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatecaCertificateTemplateIamBinding)(nil)).Elem()
}

func (i *PrivatecaCertificateTemplateIamBinding) ToPrivatecaCertificateTemplateIamBindingOutput() PrivatecaCertificateTemplateIamBindingOutput {
	return i.ToPrivatecaCertificateTemplateIamBindingOutputWithContext(context.Background())
}

func (i *PrivatecaCertificateTemplateIamBinding) ToPrivatecaCertificateTemplateIamBindingOutputWithContext(ctx context.Context) PrivatecaCertificateTemplateIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivatecaCertificateTemplateIamBindingOutput)
}

type PrivatecaCertificateTemplateIamBindingOutput struct{ *pulumi.OutputState }

func (PrivatecaCertificateTemplateIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivatecaCertificateTemplateIamBinding)(nil)).Elem()
}

func (o PrivatecaCertificateTemplateIamBindingOutput) ToPrivatecaCertificateTemplateIamBindingOutput() PrivatecaCertificateTemplateIamBindingOutput {
	return o
}

func (o PrivatecaCertificateTemplateIamBindingOutput) ToPrivatecaCertificateTemplateIamBindingOutputWithContext(ctx context.Context) PrivatecaCertificateTemplateIamBindingOutput {
	return o
}

func (o PrivatecaCertificateTemplateIamBindingOutput) CertificateTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput { return v.CertificateTemplate }).(pulumi.StringOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Condition() PrivatecaCertificateTemplateIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) PrivatecaCertificateTemplateIamBindingConditionPtrOutput {
		return v.Condition
	}).(PrivatecaCertificateTemplateIamBindingConditionPtrOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) PrivatecaCertificateTemplateIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput {
		return v.PrivatecaCertificateTemplateIamBindingId
	}).(pulumi.StringOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o PrivatecaCertificateTemplateIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *PrivatecaCertificateTemplateIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PrivatecaCertificateTemplateIamBindingInput)(nil)).Elem(), &PrivatecaCertificateTemplateIamBinding{})
	pulumi.RegisterOutputType(PrivatecaCertificateTemplateIamBindingOutput{})
}
