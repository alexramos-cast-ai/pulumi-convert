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

type KmsEkmConnection struct {
	pulumi.CustomResourceState

	// Output only. The time at which the EkmConnection was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
	// KeyManagementMode is CLOUD_KMS.
	CryptoSpacePath pulumi.StringOutput `pulumi:"cryptoSpacePath"`
	// Optional. Etag of the currently stored EkmConnection.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
	// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	KeyManagementMode  pulumi.StringPtrOutput `pulumi:"keyManagementMode"`
	KmsEkmConnectionId pulumi.StringOutput    `pulumi:"kmsEkmConnectionId"`
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name for the EkmConnection.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
	// only a single ServiceResolver is supported
	ServiceResolvers KmsEkmConnectionServiceResolverArrayOutput `pulumi:"serviceResolvers"`
	Timeouts         KmsEkmConnectionTimeoutsPtrOutput          `pulumi:"timeouts"`
}

// NewKmsEkmConnection registers a new resource with the given unique name, arguments, and options.
func NewKmsEkmConnection(ctx *pulumi.Context,
	name string, args *KmsEkmConnectionArgs, opts ...pulumi.ResourceOption) (*KmsEkmConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.ServiceResolvers == nil {
		return nil, errors.New("invalid value for required argument 'ServiceResolvers'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource KmsEkmConnection
	err = ctx.RegisterPackageResource("google-beta:index/kmsEkmConnection:KmsEkmConnection", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsEkmConnection gets an existing KmsEkmConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsEkmConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsEkmConnectionState, opts ...pulumi.ResourceOption) (*KmsEkmConnection, error) {
	var resource KmsEkmConnection
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/kmsEkmConnection:KmsEkmConnection", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsEkmConnection resources.
type kmsEkmConnectionState struct {
	// Output only. The time at which the EkmConnection was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
	// KeyManagementMode is CLOUD_KMS.
	CryptoSpacePath *string `pulumi:"cryptoSpacePath"`
	// Optional. Etag of the currently stored EkmConnection.
	Etag *string `pulumi:"etag"`
	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
	// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	KeyManagementMode  *string `pulumi:"keyManagementMode"`
	KmsEkmConnectionId *string `pulumi:"kmsEkmConnectionId"`
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location *string `pulumi:"location"`
	// The resource name for the EkmConnection.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
	// only a single ServiceResolver is supported
	ServiceResolvers []KmsEkmConnectionServiceResolver `pulumi:"serviceResolvers"`
	Timeouts         *KmsEkmConnectionTimeouts         `pulumi:"timeouts"`
}

type KmsEkmConnectionState struct {
	// Output only. The time at which the EkmConnection was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
	// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
	// KeyManagementMode is CLOUD_KMS.
	CryptoSpacePath pulumi.StringPtrInput
	// Optional. Etag of the currently stored EkmConnection.
	Etag pulumi.StringPtrInput
	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
	// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	KeyManagementMode  pulumi.StringPtrInput
	KmsEkmConnectionId pulumi.StringPtrInput
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location pulumi.StringPtrInput
	// The resource name for the EkmConnection.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
	// only a single ServiceResolver is supported
	ServiceResolvers KmsEkmConnectionServiceResolverArrayInput
	Timeouts         KmsEkmConnectionTimeoutsPtrInput
}

func (KmsEkmConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsEkmConnectionState)(nil)).Elem()
}

type kmsEkmConnectionArgs struct {
	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
	// KeyManagementMode is CLOUD_KMS.
	CryptoSpacePath *string `pulumi:"cryptoSpacePath"`
	// Optional. Etag of the currently stored EkmConnection.
	Etag *string `pulumi:"etag"`
	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
	// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	KeyManagementMode  *string `pulumi:"keyManagementMode"`
	KmsEkmConnectionId *string `pulumi:"kmsEkmConnectionId"`
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location string `pulumi:"location"`
	// The resource name for the EkmConnection.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
	// only a single ServiceResolver is supported
	ServiceResolvers []KmsEkmConnectionServiceResolver `pulumi:"serviceResolvers"`
	Timeouts         *KmsEkmConnectionTimeouts         `pulumi:"timeouts"`
}

// The set of arguments for constructing a KmsEkmConnection resource.
type KmsEkmConnectionArgs struct {
	// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
	// KeyManagementMode is CLOUD_KMS.
	CryptoSpacePath pulumi.StringPtrInput
	// Optional. Etag of the currently stored EkmConnection.
	Etag pulumi.StringPtrInput
	// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
	// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
	KeyManagementMode  pulumi.StringPtrInput
	KmsEkmConnectionId pulumi.StringPtrInput
	// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
	Location pulumi.StringInput
	// The resource name for the EkmConnection.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
	// only a single ServiceResolver is supported
	ServiceResolvers KmsEkmConnectionServiceResolverArrayInput
	Timeouts         KmsEkmConnectionTimeoutsPtrInput
}

func (KmsEkmConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsEkmConnectionArgs)(nil)).Elem()
}

type KmsEkmConnectionInput interface {
	pulumi.Input

	ToKmsEkmConnectionOutput() KmsEkmConnectionOutput
	ToKmsEkmConnectionOutputWithContext(ctx context.Context) KmsEkmConnectionOutput
}

func (*KmsEkmConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsEkmConnection)(nil)).Elem()
}

func (i *KmsEkmConnection) ToKmsEkmConnectionOutput() KmsEkmConnectionOutput {
	return i.ToKmsEkmConnectionOutputWithContext(context.Background())
}

func (i *KmsEkmConnection) ToKmsEkmConnectionOutputWithContext(ctx context.Context) KmsEkmConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsEkmConnectionOutput)
}

type KmsEkmConnectionOutput struct{ *pulumi.OutputState }

func (KmsEkmConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsEkmConnection)(nil)).Elem()
}

func (o KmsEkmConnectionOutput) ToKmsEkmConnectionOutput() KmsEkmConnectionOutput {
	return o
}

func (o KmsEkmConnectionOutput) ToKmsEkmConnectionOutputWithContext(ctx context.Context) KmsEkmConnectionOutput {
	return o
}

// Output only. The time at which the EkmConnection was created. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
// resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o KmsEkmConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Identifies the EKM Crypto Space that this EkmConnection maps to. Note: This field is required if
// KeyManagementMode is CLOUD_KMS.
func (o KmsEkmConnectionOutput) CryptoSpacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.CryptoSpacePath }).(pulumi.StringOutput)
}

// Optional. Etag of the currently stored EkmConnection.
func (o KmsEkmConnectionOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. Describes who can perform control plane operations on the EKM. If unset, this defaults to MANUAL Default
// value: "MANUAL" Possible values: ["MANUAL", "CLOUD_KMS"]
func (o KmsEkmConnectionOutput) KeyManagementMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringPtrOutput { return v.KeyManagementMode }).(pulumi.StringPtrOutput)
}

func (o KmsEkmConnectionOutput) KmsEkmConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.KmsEkmConnectionId }).(pulumi.StringOutput)
}

// The location for the EkmConnection. A full list of valid locations can be found by running 'gcloud kms locations list'.
func (o KmsEkmConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name for the EkmConnection.
func (o KmsEkmConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o KmsEkmConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsEkmConnection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A list of ServiceResolvers where the EKM can be reached. There should be one ServiceResolver per EKM replica. Currently,
// only a single ServiceResolver is supported
func (o KmsEkmConnectionOutput) ServiceResolvers() KmsEkmConnectionServiceResolverArrayOutput {
	return o.ApplyT(func(v *KmsEkmConnection) KmsEkmConnectionServiceResolverArrayOutput { return v.ServiceResolvers }).(KmsEkmConnectionServiceResolverArrayOutput)
}

func (o KmsEkmConnectionOutput) Timeouts() KmsEkmConnectionTimeoutsPtrOutput {
	return o.ApplyT(func(v *KmsEkmConnection) KmsEkmConnectionTimeoutsPtrOutput { return v.Timeouts }).(KmsEkmConnectionTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsEkmConnectionInput)(nil)).Elem(), &KmsEkmConnection{})
	pulumi.RegisterOutputType(KmsEkmConnectionOutput{})
}
