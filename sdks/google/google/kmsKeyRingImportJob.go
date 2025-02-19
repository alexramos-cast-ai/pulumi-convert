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

type KmsKeyRingImportJob struct {
	pulumi.CustomResourceState

	// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
	// statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
	// ImportMethod is one with a protection level of HSM.
	Attestations KmsKeyRingImportJobAttestationArrayOutput `pulumi:"attestations"`
	// The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
	ExpireTime pulumi.StringOutput `pulumi:"expireTime"`
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
	ImportJobId pulumi.StringOutput `pulumi:"importJobId"`
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
	// "RSA_OAEP_4096_SHA1_AES_256"]
	ImportMethod pulumi.StringOutput `pulumi:"importMethod"`
	// The KeyRing that this import job belongs to. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing               pulumi.StringOutput `pulumi:"keyRing"`
	KmsKeyRingImportJobId pulumi.StringOutput `pulumi:"kmsKeyRingImportJobId"`
	// The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
	Name pulumi.StringOutput `pulumi:"name"`
	// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
	// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	ProtectionLevel pulumi.StringOutput `pulumi:"protectionLevel"`
	// The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
	PublicKeys KmsKeyRingImportJobPublicKeyArrayOutput `pulumi:"publicKeys"`
	// The current state of the ImportJob, indicating if it can be used.
	State    pulumi.StringOutput                  `pulumi:"state"`
	Timeouts KmsKeyRingImportJobTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewKmsKeyRingImportJob registers a new resource with the given unique name, arguments, and options.
func NewKmsKeyRingImportJob(ctx *pulumi.Context,
	name string, args *KmsKeyRingImportJobArgs, opts ...pulumi.ResourceOption) (*KmsKeyRingImportJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImportJobId == nil {
		return nil, errors.New("invalid value for required argument 'ImportJobId'")
	}
	if args.ImportMethod == nil {
		return nil, errors.New("invalid value for required argument 'ImportMethod'")
	}
	if args.KeyRing == nil {
		return nil, errors.New("invalid value for required argument 'KeyRing'")
	}
	if args.ProtectionLevel == nil {
		return nil, errors.New("invalid value for required argument 'ProtectionLevel'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource KmsKeyRingImportJob
	err = ctx.RegisterPackageResource("google:index/kmsKeyRingImportJob:KmsKeyRingImportJob", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKmsKeyRingImportJob gets an existing KmsKeyRingImportJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKmsKeyRingImportJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KmsKeyRingImportJobState, opts ...pulumi.ResourceOption) (*KmsKeyRingImportJob, error) {
	var resource KmsKeyRingImportJob
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/kmsKeyRingImportJob:KmsKeyRingImportJob", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KmsKeyRingImportJob resources.
type kmsKeyRingImportJobState struct {
	// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
	// statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
	// ImportMethod is one with a protection level of HSM.
	Attestations []KmsKeyRingImportJobAttestation `pulumi:"attestations"`
	// The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
	ExpireTime *string `pulumi:"expireTime"`
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
	ImportJobId *string `pulumi:"importJobId"`
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
	// "RSA_OAEP_4096_SHA1_AES_256"]
	ImportMethod *string `pulumi:"importMethod"`
	// The KeyRing that this import job belongs to. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing               *string `pulumi:"keyRing"`
	KmsKeyRingImportJobId *string `pulumi:"kmsKeyRingImportJobId"`
	// The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
	Name *string `pulumi:"name"`
	// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
	// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	ProtectionLevel *string `pulumi:"protectionLevel"`
	// The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
	PublicKeys []KmsKeyRingImportJobPublicKey `pulumi:"publicKeys"`
	// The current state of the ImportJob, indicating if it can be used.
	State    *string                      `pulumi:"state"`
	Timeouts *KmsKeyRingImportJobTimeouts `pulumi:"timeouts"`
}

type KmsKeyRingImportJobState struct {
	// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
	// statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
	// ImportMethod is one with a protection level of HSM.
	Attestations KmsKeyRingImportJobAttestationArrayInput
	// The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
	ExpireTime pulumi.StringPtrInput
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
	ImportJobId pulumi.StringPtrInput
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
	// "RSA_OAEP_4096_SHA1_AES_256"]
	ImportMethod pulumi.StringPtrInput
	// The KeyRing that this import job belongs to. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing               pulumi.StringPtrInput
	KmsKeyRingImportJobId pulumi.StringPtrInput
	// The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
	Name pulumi.StringPtrInput
	// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
	// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	ProtectionLevel pulumi.StringPtrInput
	// The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
	PublicKeys KmsKeyRingImportJobPublicKeyArrayInput
	// The current state of the ImportJob, indicating if it can be used.
	State    pulumi.StringPtrInput
	Timeouts KmsKeyRingImportJobTimeoutsPtrInput
}

func (KmsKeyRingImportJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsKeyRingImportJobState)(nil)).Elem()
}

type kmsKeyRingImportJobArgs struct {
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
	ImportJobId string `pulumi:"importJobId"`
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
	// "RSA_OAEP_4096_SHA1_AES_256"]
	ImportMethod string `pulumi:"importMethod"`
	// The KeyRing that this import job belongs to. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing               string  `pulumi:"keyRing"`
	KmsKeyRingImportJobId *string `pulumi:"kmsKeyRingImportJobId"`
	// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
	// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	ProtectionLevel string                       `pulumi:"protectionLevel"`
	Timeouts        *KmsKeyRingImportJobTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a KmsKeyRingImportJob resource.
type KmsKeyRingImportJobArgs struct {
	// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
	ImportJobId pulumi.StringInput
	// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
	// "RSA_OAEP_4096_SHA1_AES_256"]
	ImportMethod pulumi.StringInput
	// The KeyRing that this import job belongs to. Format:
	// ''projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}''.
	KeyRing               pulumi.StringInput
	KmsKeyRingImportJobId pulumi.StringPtrInput
	// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
	// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
	ProtectionLevel pulumi.StringInput
	Timeouts        KmsKeyRingImportJobTimeoutsPtrInput
}

func (KmsKeyRingImportJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kmsKeyRingImportJobArgs)(nil)).Elem()
}

type KmsKeyRingImportJobInput interface {
	pulumi.Input

	ToKmsKeyRingImportJobOutput() KmsKeyRingImportJobOutput
	ToKmsKeyRingImportJobOutputWithContext(ctx context.Context) KmsKeyRingImportJobOutput
}

func (*KmsKeyRingImportJob) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsKeyRingImportJob)(nil)).Elem()
}

func (i *KmsKeyRingImportJob) ToKmsKeyRingImportJobOutput() KmsKeyRingImportJobOutput {
	return i.ToKmsKeyRingImportJobOutputWithContext(context.Background())
}

func (i *KmsKeyRingImportJob) ToKmsKeyRingImportJobOutputWithContext(ctx context.Context) KmsKeyRingImportJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KmsKeyRingImportJobOutput)
}

type KmsKeyRingImportJobOutput struct{ *pulumi.OutputState }

func (KmsKeyRingImportJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KmsKeyRingImportJob)(nil)).Elem()
}

func (o KmsKeyRingImportJobOutput) ToKmsKeyRingImportJobOutput() KmsKeyRingImportJobOutput {
	return o
}

func (o KmsKeyRingImportJobOutput) ToKmsKeyRingImportJobOutputWithContext(ctx context.Context) KmsKeyRingImportJobOutput {
	return o
}

// Statement that was generated and signed by the key creator (for example, an HSM) at key creation time. Use this
// statement to verify attributes of the key as stored on the HSM, independently of Google. Only present if the chosen
// ImportMethod is one with a protection level of HSM.
func (o KmsKeyRingImportJobOutput) Attestations() KmsKeyRingImportJobAttestationArrayOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) KmsKeyRingImportJobAttestationArrayOutput { return v.Attestations }).(KmsKeyRingImportJobAttestationArrayOutput)
}

// The time at which this resource is scheduled for expiration and can no longer be used. This is in RFC3339 text format.
func (o KmsKeyRingImportJobOutput) ExpireTime() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.ExpireTime }).(pulumi.StringOutput)
}

// It must be unique within a KeyRing and match the regular expression [a-zA-Z0-9_-]{1,63}
func (o KmsKeyRingImportJobOutput) ImportJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.ImportJobId }).(pulumi.StringOutput)
}

// The wrapping method to be used for incoming key material. Possible values: ["RSA_OAEP_3072_SHA1_AES_256",
// "RSA_OAEP_4096_SHA1_AES_256"]
func (o KmsKeyRingImportJobOutput) ImportMethod() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.ImportMethod }).(pulumi.StringOutput)
}

// The KeyRing that this import job belongs to. Format:
// ”projects/{{project}}/locations/{{location}}/keyRings/{{keyRing}}”.
func (o KmsKeyRingImportJobOutput) KeyRing() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.KeyRing }).(pulumi.StringOutput)
}

func (o KmsKeyRingImportJobOutput) KmsKeyRingImportJobId() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.KmsKeyRingImportJobId }).(pulumi.StringOutput)
}

// The resource name for this ImportJob in the format projects/*/locations/*/keyRings/*/importJobs/*.
func (o KmsKeyRingImportJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The protection level of the ImportJob. This must match the protectionLevel of the versionTemplate on the CryptoKey you
// attempt to import into. Possible values: ["SOFTWARE", "HSM", "EXTERNAL"]
func (o KmsKeyRingImportJobOutput) ProtectionLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.ProtectionLevel }).(pulumi.StringOutput)
}

// The public key with which to wrap key material prior to import. Only returned if state is 'ACTIVE'.
func (o KmsKeyRingImportJobOutput) PublicKeys() KmsKeyRingImportJobPublicKeyArrayOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) KmsKeyRingImportJobPublicKeyArrayOutput { return v.PublicKeys }).(KmsKeyRingImportJobPublicKeyArrayOutput)
}

// The current state of the ImportJob, indicating if it can be used.
func (o KmsKeyRingImportJobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o KmsKeyRingImportJobOutput) Timeouts() KmsKeyRingImportJobTimeoutsPtrOutput {
	return o.ApplyT(func(v *KmsKeyRingImportJob) KmsKeyRingImportJobTimeoutsPtrOutput { return v.Timeouts }).(KmsKeyRingImportJobTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KmsKeyRingImportJobInput)(nil)).Elem(), &KmsKeyRingImportJob{})
	pulumi.RegisterOutputType(KmsKeyRingImportJobOutput{})
}
