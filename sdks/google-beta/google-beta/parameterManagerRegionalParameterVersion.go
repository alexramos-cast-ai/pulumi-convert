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

type ParameterManagerRegionalParameterVersion struct {
	pulumi.CustomResourceState

	// The time at which the Regional Parameter Version was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Location of Parameter Manager Regional parameter resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the Regional Parameter Version. Format:
	// 'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameter Manager Regional Parameter resource.
	Parameter pulumi.StringOutput `pulumi:"parameter"`
	// The Regional Parameter data.
	ParameterData                              pulumi.StringOutput `pulumi:"parameterData"`
	ParameterManagerRegionalParameterVersionId pulumi.StringOutput `pulumi:"parameterManagerRegionalParameterVersionId"`
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	ParameterVersionId pulumi.StringOutput                                       `pulumi:"parameterVersionId"`
	Timeouts           ParameterManagerRegionalParameterVersionTimeoutsPtrOutput `pulumi:"timeouts"`
	// The time at which the Regional Parameter Version was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewParameterManagerRegionalParameterVersion registers a new resource with the given unique name, arguments, and options.
func NewParameterManagerRegionalParameterVersion(ctx *pulumi.Context,
	name string, args *ParameterManagerRegionalParameterVersionArgs, opts ...pulumi.ResourceOption) (*ParameterManagerRegionalParameterVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parameter == nil {
		return nil, errors.New("invalid value for required argument 'Parameter'")
	}
	if args.ParameterData == nil {
		return nil, errors.New("invalid value for required argument 'ParameterData'")
	}
	if args.ParameterVersionId == nil {
		return nil, errors.New("invalid value for required argument 'ParameterVersionId'")
	}
	if args.ParameterData != nil {
		args.ParameterData = pulumi.ToSecret(args.ParameterData).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"parameterData",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ParameterManagerRegionalParameterVersion
	err = ctx.RegisterPackageResource("google-beta:index/parameterManagerRegionalParameterVersion:ParameterManagerRegionalParameterVersion", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterManagerRegionalParameterVersion gets an existing ParameterManagerRegionalParameterVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterManagerRegionalParameterVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterManagerRegionalParameterVersionState, opts ...pulumi.ResourceOption) (*ParameterManagerRegionalParameterVersion, error) {
	var resource ParameterManagerRegionalParameterVersion
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/parameterManagerRegionalParameterVersion:ParameterManagerRegionalParameterVersion", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterManagerRegionalParameterVersion resources.
type parameterManagerRegionalParameterVersionState struct {
	// The time at which the Regional Parameter Version was created.
	CreateTime *string `pulumi:"createTime"`
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	Disabled *bool `pulumi:"disabled"`
	// Location of Parameter Manager Regional parameter resource.
	Location *string `pulumi:"location"`
	// The resource name of the Regional Parameter Version. Format:
	// 'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'
	Name *string `pulumi:"name"`
	// Parameter Manager Regional Parameter resource.
	Parameter *string `pulumi:"parameter"`
	// The Regional Parameter data.
	ParameterData                              *string `pulumi:"parameterData"`
	ParameterManagerRegionalParameterVersionId *string `pulumi:"parameterManagerRegionalParameterVersionId"`
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	ParameterVersionId *string                                           `pulumi:"parameterVersionId"`
	Timeouts           *ParameterManagerRegionalParameterVersionTimeouts `pulumi:"timeouts"`
	// The time at which the Regional Parameter Version was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ParameterManagerRegionalParameterVersionState struct {
	// The time at which the Regional Parameter Version was created.
	CreateTime pulumi.StringPtrInput
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	Disabled pulumi.BoolPtrInput
	// Location of Parameter Manager Regional parameter resource.
	Location pulumi.StringPtrInput
	// The resource name of the Regional Parameter Version. Format:
	// 'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'
	Name pulumi.StringPtrInput
	// Parameter Manager Regional Parameter resource.
	Parameter pulumi.StringPtrInput
	// The Regional Parameter data.
	ParameterData                              pulumi.StringPtrInput
	ParameterManagerRegionalParameterVersionId pulumi.StringPtrInput
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	ParameterVersionId pulumi.StringPtrInput
	Timeouts           ParameterManagerRegionalParameterVersionTimeoutsPtrInput
	// The time at which the Regional Parameter Version was updated.
	UpdateTime pulumi.StringPtrInput
}

func (ParameterManagerRegionalParameterVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterManagerRegionalParameterVersionState)(nil)).Elem()
}

type parameterManagerRegionalParameterVersionArgs struct {
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	Disabled *bool `pulumi:"disabled"`
	// Parameter Manager Regional Parameter resource.
	Parameter string `pulumi:"parameter"`
	// The Regional Parameter data.
	ParameterData                              string  `pulumi:"parameterData"`
	ParameterManagerRegionalParameterVersionId *string `pulumi:"parameterManagerRegionalParameterVersionId"`
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	ParameterVersionId string                                            `pulumi:"parameterVersionId"`
	Timeouts           *ParameterManagerRegionalParameterVersionTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ParameterManagerRegionalParameterVersion resource.
type ParameterManagerRegionalParameterVersionArgs struct {
	// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
	Disabled pulumi.BoolPtrInput
	// Parameter Manager Regional Parameter resource.
	Parameter pulumi.StringInput
	// The Regional Parameter data.
	ParameterData                              pulumi.StringInput
	ParameterManagerRegionalParameterVersionId pulumi.StringPtrInput
	// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
	ParameterVersionId pulumi.StringInput
	Timeouts           ParameterManagerRegionalParameterVersionTimeoutsPtrInput
}

func (ParameterManagerRegionalParameterVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterManagerRegionalParameterVersionArgs)(nil)).Elem()
}

type ParameterManagerRegionalParameterVersionInput interface {
	pulumi.Input

	ToParameterManagerRegionalParameterVersionOutput() ParameterManagerRegionalParameterVersionOutput
	ToParameterManagerRegionalParameterVersionOutputWithContext(ctx context.Context) ParameterManagerRegionalParameterVersionOutput
}

func (*ParameterManagerRegionalParameterVersion) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterManagerRegionalParameterVersion)(nil)).Elem()
}

func (i *ParameterManagerRegionalParameterVersion) ToParameterManagerRegionalParameterVersionOutput() ParameterManagerRegionalParameterVersionOutput {
	return i.ToParameterManagerRegionalParameterVersionOutputWithContext(context.Background())
}

func (i *ParameterManagerRegionalParameterVersion) ToParameterManagerRegionalParameterVersionOutputWithContext(ctx context.Context) ParameterManagerRegionalParameterVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterManagerRegionalParameterVersionOutput)
}

type ParameterManagerRegionalParameterVersionOutput struct{ *pulumi.OutputState }

func (ParameterManagerRegionalParameterVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterManagerRegionalParameterVersion)(nil)).Elem()
}

func (o ParameterManagerRegionalParameterVersionOutput) ToParameterManagerRegionalParameterVersionOutput() ParameterManagerRegionalParameterVersionOutput {
	return o
}

func (o ParameterManagerRegionalParameterVersionOutput) ToParameterManagerRegionalParameterVersionOutputWithContext(ctx context.Context) ParameterManagerRegionalParameterVersionOutput {
	return o
}

// The time at which the Regional Parameter Version was created.
func (o ParameterManagerRegionalParameterVersionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The current state of Regional Parameter Version. This field is only applicable for updating Regional Parameter Version.
func (o ParameterManagerRegionalParameterVersionOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Location of Parameter Manager Regional parameter resource.
func (o ParameterManagerRegionalParameterVersionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the Regional Parameter Version. Format:
// 'projects/{{project}}/locations/{{location}}/parameters/{{parameter_id}}/versions/{{parameter_version_id}}'
func (o ParameterManagerRegionalParameterVersionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Parameter Manager Regional Parameter resource.
func (o ParameterManagerRegionalParameterVersionOutput) Parameter() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.Parameter }).(pulumi.StringOutput)
}

// The Regional Parameter data.
func (o ParameterManagerRegionalParameterVersionOutput) ParameterData() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.ParameterData }).(pulumi.StringOutput)
}

func (o ParameterManagerRegionalParameterVersionOutput) ParameterManagerRegionalParameterVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput {
		return v.ParameterManagerRegionalParameterVersionId
	}).(pulumi.StringOutput)
}

// Version ID of the Regional Parameter Version Resource. This must be unique within the Regional Parameter.
func (o ParameterManagerRegionalParameterVersionOutput) ParameterVersionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.ParameterVersionId }).(pulumi.StringOutput)
}

func (o ParameterManagerRegionalParameterVersionOutput) Timeouts() ParameterManagerRegionalParameterVersionTimeoutsPtrOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) ParameterManagerRegionalParameterVersionTimeoutsPtrOutput {
		return v.Timeouts
	}).(ParameterManagerRegionalParameterVersionTimeoutsPtrOutput)
}

// The time at which the Regional Parameter Version was updated.
func (o ParameterManagerRegionalParameterVersionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterManagerRegionalParameterVersion) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterManagerRegionalParameterVersionInput)(nil)).Elem(), &ParameterManagerRegionalParameterVersion{})
	pulumi.RegisterOutputType(ParameterManagerRegionalParameterVersionOutput{})
}
