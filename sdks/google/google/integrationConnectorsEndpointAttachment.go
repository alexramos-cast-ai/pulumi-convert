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

type IntegrationConnectorsEndpointAttachment struct {
	pulumi.CustomResourceState

	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Enable global access for endpoint attachment.
	EndpointGlobalAccess pulumi.BoolPtrOutput `pulumi:"endpointGlobalAccess"`
	// The Private Service Connect connection endpoint ip.
	EndpointIp                                pulumi.StringOutput `pulumi:"endpointIp"`
	IntegrationConnectorsEndpointAttachmentId pulumi.StringOutput `pulumi:"integrationConnectorsEndpointAttachmentId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location in which Endpoint Attachment needs to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of Endpoint Attachment needs to be created.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The path of the service attachment.
	ServiceAttachment pulumi.StringOutput `pulumi:"serviceAttachment"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                                   `pulumi:"terraformLabels"`
	Timeouts        IntegrationConnectorsEndpointAttachmentTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewIntegrationConnectorsEndpointAttachment registers a new resource with the given unique name, arguments, and options.
func NewIntegrationConnectorsEndpointAttachment(ctx *pulumi.Context,
	name string, args *IntegrationConnectorsEndpointAttachmentArgs, opts ...pulumi.ResourceOption) (*IntegrationConnectorsEndpointAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ServiceAttachment == nil {
		return nil, errors.New("invalid value for required argument 'ServiceAttachment'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IntegrationConnectorsEndpointAttachment
	err = ctx.RegisterPackageResource("google:index/integrationConnectorsEndpointAttachment:IntegrationConnectorsEndpointAttachment", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationConnectorsEndpointAttachment gets an existing IntegrationConnectorsEndpointAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationConnectorsEndpointAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationConnectorsEndpointAttachmentState, opts ...pulumi.ResourceOption) (*IntegrationConnectorsEndpointAttachment, error) {
	var resource IntegrationConnectorsEndpointAttachment
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/integrationConnectorsEndpointAttachment:IntegrationConnectorsEndpointAttachment", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationConnectorsEndpointAttachment resources.
type integrationConnectorsEndpointAttachmentState struct {
	// Time the Namespace was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Description of the resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Enable global access for endpoint attachment.
	EndpointGlobalAccess *bool `pulumi:"endpointGlobalAccess"`
	// The Private Service Connect connection endpoint ip.
	EndpointIp                                *string `pulumi:"endpointIp"`
	IntegrationConnectorsEndpointAttachmentId *string `pulumi:"integrationConnectorsEndpointAttachmentId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Location in which Endpoint Attachment needs to be created.
	Location *string `pulumi:"location"`
	// Name of Endpoint Attachment needs to be created.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The path of the service attachment.
	ServiceAttachment *string `pulumi:"serviceAttachment"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                                `pulumi:"terraformLabels"`
	Timeouts        *IntegrationConnectorsEndpointAttachmentTimeouts `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type IntegrationConnectorsEndpointAttachmentState struct {
	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Description of the resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Enable global access for endpoint attachment.
	EndpointGlobalAccess pulumi.BoolPtrInput
	// The Private Service Connect connection endpoint ip.
	EndpointIp                                pulumi.StringPtrInput
	IntegrationConnectorsEndpointAttachmentId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Location in which Endpoint Attachment needs to be created.
	Location pulumi.StringPtrInput
	// Name of Endpoint Attachment needs to be created.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The path of the service attachment.
	ServiceAttachment pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        IntegrationConnectorsEndpointAttachmentTimeoutsPtrInput
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (IntegrationConnectorsEndpointAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsEndpointAttachmentState)(nil)).Elem()
}

type integrationConnectorsEndpointAttachmentArgs struct {
	// Description of the resource.
	Description *string `pulumi:"description"`
	// Enable global access for endpoint attachment.
	EndpointGlobalAccess                      *bool   `pulumi:"endpointGlobalAccess"`
	IntegrationConnectorsEndpointAttachmentId *string `pulumi:"integrationConnectorsEndpointAttachmentId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Location in which Endpoint Attachment needs to be created.
	Location string `pulumi:"location"`
	// Name of Endpoint Attachment needs to be created.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The path of the service attachment.
	ServiceAttachment string                                           `pulumi:"serviceAttachment"`
	Timeouts          *IntegrationConnectorsEndpointAttachmentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IntegrationConnectorsEndpointAttachment resource.
type IntegrationConnectorsEndpointAttachmentArgs struct {
	// Description of the resource.
	Description pulumi.StringPtrInput
	// Enable global access for endpoint attachment.
	EndpointGlobalAccess                      pulumi.BoolPtrInput
	IntegrationConnectorsEndpointAttachmentId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Location in which Endpoint Attachment needs to be created.
	Location pulumi.StringInput
	// Name of Endpoint Attachment needs to be created.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The path of the service attachment.
	ServiceAttachment pulumi.StringInput
	Timeouts          IntegrationConnectorsEndpointAttachmentTimeoutsPtrInput
}

func (IntegrationConnectorsEndpointAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsEndpointAttachmentArgs)(nil)).Elem()
}

type IntegrationConnectorsEndpointAttachmentInput interface {
	pulumi.Input

	ToIntegrationConnectorsEndpointAttachmentOutput() IntegrationConnectorsEndpointAttachmentOutput
	ToIntegrationConnectorsEndpointAttachmentOutputWithContext(ctx context.Context) IntegrationConnectorsEndpointAttachmentOutput
}

func (*IntegrationConnectorsEndpointAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsEndpointAttachment)(nil)).Elem()
}

func (i *IntegrationConnectorsEndpointAttachment) ToIntegrationConnectorsEndpointAttachmentOutput() IntegrationConnectorsEndpointAttachmentOutput {
	return i.ToIntegrationConnectorsEndpointAttachmentOutputWithContext(context.Background())
}

func (i *IntegrationConnectorsEndpointAttachment) ToIntegrationConnectorsEndpointAttachmentOutputWithContext(ctx context.Context) IntegrationConnectorsEndpointAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationConnectorsEndpointAttachmentOutput)
}

type IntegrationConnectorsEndpointAttachmentOutput struct{ *pulumi.OutputState }

func (IntegrationConnectorsEndpointAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsEndpointAttachment)(nil)).Elem()
}

func (o IntegrationConnectorsEndpointAttachmentOutput) ToIntegrationConnectorsEndpointAttachmentOutput() IntegrationConnectorsEndpointAttachmentOutput {
	return o
}

func (o IntegrationConnectorsEndpointAttachmentOutput) ToIntegrationConnectorsEndpointAttachmentOutputWithContext(ctx context.Context) IntegrationConnectorsEndpointAttachmentOutput {
	return o
}

// Time the Namespace was created in UTC.
func (o IntegrationConnectorsEndpointAttachmentOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the resource.
func (o IntegrationConnectorsEndpointAttachmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IntegrationConnectorsEndpointAttachmentOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Enable global access for endpoint attachment.
func (o IntegrationConnectorsEndpointAttachmentOutput) EndpointGlobalAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.BoolPtrOutput { return v.EndpointGlobalAccess }).(pulumi.BoolPtrOutput)
}

// The Private Service Connect connection endpoint ip.
func (o IntegrationConnectorsEndpointAttachmentOutput) EndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.EndpointIp }).(pulumi.StringOutput)
}

func (o IntegrationConnectorsEndpointAttachmentOutput) IntegrationConnectorsEndpointAttachmentId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput {
		return v.IntegrationConnectorsEndpointAttachmentId
	}).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o IntegrationConnectorsEndpointAttachmentOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location in which Endpoint Attachment needs to be created.
func (o IntegrationConnectorsEndpointAttachmentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Name of Endpoint Attachment needs to be created.
func (o IntegrationConnectorsEndpointAttachmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o IntegrationConnectorsEndpointAttachmentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The path of the service attachment.
func (o IntegrationConnectorsEndpointAttachmentOutput) ServiceAttachment() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.ServiceAttachment }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o IntegrationConnectorsEndpointAttachmentOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o IntegrationConnectorsEndpointAttachmentOutput) Timeouts() IntegrationConnectorsEndpointAttachmentTimeoutsPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) IntegrationConnectorsEndpointAttachmentTimeoutsPtrOutput {
		return v.Timeouts
	}).(IntegrationConnectorsEndpointAttachmentTimeoutsPtrOutput)
}

// Time the Namespace was updated in UTC.
func (o IntegrationConnectorsEndpointAttachmentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsEndpointAttachment) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationConnectorsEndpointAttachmentInput)(nil)).Elem(), &IntegrationConnectorsEndpointAttachment{})
	pulumi.RegisterOutputType(IntegrationConnectorsEndpointAttachmentOutput{})
}
