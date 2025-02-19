// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NetworkServicesEdgeCacheKeyset struct {
	pulumi.CustomResourceState

	// A human-readable description of the resource.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name                             pulumi.StringOutput `pulumi:"name"`
	NetworkServicesEdgeCacheKeysetId pulumi.StringOutput `pulumi:"networkServicesEdgeCacheKeysetId"`
	Project                          pulumi.StringOutput `pulumi:"project"`
	// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
	// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
	// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
	// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
	// keys to a keyset.
	PublicKeys NetworkServicesEdgeCacheKeysetPublicKeyArrayOutput `pulumi:"publicKeys"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                          `pulumi:"terraformLabels"`
	Timeouts        NetworkServicesEdgeCacheKeysetTimeoutsPtrOutput `pulumi:"timeouts"`
	// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
	// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
	// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayOutput `pulumi:"validationSharedKeys"`
}

// NewNetworkServicesEdgeCacheKeyset registers a new resource with the given unique name, arguments, and options.
func NewNetworkServicesEdgeCacheKeyset(ctx *pulumi.Context,
	name string, args *NetworkServicesEdgeCacheKeysetArgs, opts ...pulumi.ResourceOption) (*NetworkServicesEdgeCacheKeyset, error) {
	if args == nil {
		args = &NetworkServicesEdgeCacheKeysetArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetworkServicesEdgeCacheKeyset
	err = ctx.RegisterPackageResource("google-beta:index/networkServicesEdgeCacheKeyset:NetworkServicesEdgeCacheKeyset", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkServicesEdgeCacheKeyset gets an existing NetworkServicesEdgeCacheKeyset resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkServicesEdgeCacheKeyset(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkServicesEdgeCacheKeysetState, opts ...pulumi.ResourceOption) (*NetworkServicesEdgeCacheKeyset, error) {
	var resource NetworkServicesEdgeCacheKeyset
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/networkServicesEdgeCacheKeyset:NetworkServicesEdgeCacheKeyset", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkServicesEdgeCacheKeyset resources.
type networkServicesEdgeCacheKeysetState struct {
	// A human-readable description of the resource.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name                             *string `pulumi:"name"`
	NetworkServicesEdgeCacheKeysetId *string `pulumi:"networkServicesEdgeCacheKeysetId"`
	Project                          *string `pulumi:"project"`
	// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
	// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
	// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
	// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
	// keys to a keyset.
	PublicKeys []NetworkServicesEdgeCacheKeysetPublicKey `pulumi:"publicKeys"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                       `pulumi:"terraformLabels"`
	Timeouts        *NetworkServicesEdgeCacheKeysetTimeouts `pulumi:"timeouts"`
	// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
	// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
	// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys []NetworkServicesEdgeCacheKeysetValidationSharedKey `pulumi:"validationSharedKeys"`
}

type NetworkServicesEdgeCacheKeysetState struct {
	// A human-readable description of the resource.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name                             pulumi.StringPtrInput
	NetworkServicesEdgeCacheKeysetId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
	// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
	// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
	// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
	// keys to a keyset.
	PublicKeys NetworkServicesEdgeCacheKeysetPublicKeyArrayInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetworkServicesEdgeCacheKeysetTimeoutsPtrInput
	// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
	// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
	// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayInput
}

func (NetworkServicesEdgeCacheKeysetState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesEdgeCacheKeysetState)(nil)).Elem()
}

type networkServicesEdgeCacheKeysetArgs struct {
	// A human-readable description of the resource.
	Description *string `pulumi:"description"`
	// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name                             *string `pulumi:"name"`
	NetworkServicesEdgeCacheKeysetId *string `pulumi:"networkServicesEdgeCacheKeysetId"`
	Project                          *string `pulumi:"project"`
	// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
	// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
	// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
	// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
	// keys to a keyset.
	PublicKeys []NetworkServicesEdgeCacheKeysetPublicKey `pulumi:"publicKeys"`
	Timeouts   *NetworkServicesEdgeCacheKeysetTimeouts   `pulumi:"timeouts"`
	// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
	// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
	// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys []NetworkServicesEdgeCacheKeysetValidationSharedKey `pulumi:"validationSharedKeys"`
}

// The set of arguments for constructing a NetworkServicesEdgeCacheKeyset resource.
type NetworkServicesEdgeCacheKeysetArgs struct {
	// A human-readable description of the resource.
	Description pulumi.StringPtrInput
	// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
	// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
	// present on the resource.
	Labels pulumi.StringMapInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
	// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
	// characters must be a dash, underscore, letter or digit.
	Name                             pulumi.StringPtrInput
	NetworkServicesEdgeCacheKeysetId pulumi.StringPtrInput
	Project                          pulumi.StringPtrInput
	// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
	// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
	// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
	// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
	// keys to a keyset.
	PublicKeys NetworkServicesEdgeCacheKeysetPublicKeyArrayInput
	Timeouts   NetworkServicesEdgeCacheKeysetTimeoutsPtrInput
	// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
	// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
	// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
	// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
	ValidationSharedKeys NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayInput
}

func (NetworkServicesEdgeCacheKeysetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkServicesEdgeCacheKeysetArgs)(nil)).Elem()
}

type NetworkServicesEdgeCacheKeysetInput interface {
	pulumi.Input

	ToNetworkServicesEdgeCacheKeysetOutput() NetworkServicesEdgeCacheKeysetOutput
	ToNetworkServicesEdgeCacheKeysetOutputWithContext(ctx context.Context) NetworkServicesEdgeCacheKeysetOutput
}

func (*NetworkServicesEdgeCacheKeyset) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesEdgeCacheKeyset)(nil)).Elem()
}

func (i *NetworkServicesEdgeCacheKeyset) ToNetworkServicesEdgeCacheKeysetOutput() NetworkServicesEdgeCacheKeysetOutput {
	return i.ToNetworkServicesEdgeCacheKeysetOutputWithContext(context.Background())
}

func (i *NetworkServicesEdgeCacheKeyset) ToNetworkServicesEdgeCacheKeysetOutputWithContext(ctx context.Context) NetworkServicesEdgeCacheKeysetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkServicesEdgeCacheKeysetOutput)
}

type NetworkServicesEdgeCacheKeysetOutput struct{ *pulumi.OutputState }

func (NetworkServicesEdgeCacheKeysetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkServicesEdgeCacheKeyset)(nil)).Elem()
}

func (o NetworkServicesEdgeCacheKeysetOutput) ToNetworkServicesEdgeCacheKeysetOutput() NetworkServicesEdgeCacheKeysetOutput {
	return o
}

func (o NetworkServicesEdgeCacheKeysetOutput) ToNetworkServicesEdgeCacheKeysetOutputWithContext(ctx context.Context) NetworkServicesEdgeCacheKeysetOutput {
	return o
}

// A human-readable description of the resource.
func (o NetworkServicesEdgeCacheKeysetOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkServicesEdgeCacheKeysetOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Set of label tags associated with the EdgeCache resource. **Note**: This field is non-authoritative, and will only
// manage the labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels
// present on the resource.
func (o NetworkServicesEdgeCacheKeysetOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-64 characters long, and
// match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter, and all following
// characters must be a dash, underscore, letter or digit.
func (o NetworkServicesEdgeCacheKeysetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetworkServicesEdgeCacheKeysetOutput) NetworkServicesEdgeCacheKeysetId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringOutput { return v.NetworkServicesEdgeCacheKeysetId }).(pulumi.StringOutput)
}

func (o NetworkServicesEdgeCacheKeysetOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// An ordered list of Ed25519 public keys to use for validating signed requests. You must specify 'public_keys' or
// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first. You may specify no more than one
// Google-managed public key. If you specify 'public_keys', you must specify at least one (1) key and may specify up to
// three (3) keys. Ed25519 public keys are not secret, and only allow Google to validate a request was signed by your
// corresponding private key. Ensure that the private key is kept secret, and that only authorized users can add public
// keys to a keyset.
func (o NetworkServicesEdgeCacheKeysetOutput) PublicKeys() NetworkServicesEdgeCacheKeysetPublicKeyArrayOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) NetworkServicesEdgeCacheKeysetPublicKeyArrayOutput {
		return v.PublicKeys
	}).(NetworkServicesEdgeCacheKeysetPublicKeyArrayOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetworkServicesEdgeCacheKeysetOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetworkServicesEdgeCacheKeysetOutput) Timeouts() NetworkServicesEdgeCacheKeysetTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) NetworkServicesEdgeCacheKeysetTimeoutsPtrOutput {
		return v.Timeouts
	}).(NetworkServicesEdgeCacheKeysetTimeoutsPtrOutput)
}

// An ordered list of shared keys to use for validating signed requests. Shared keys are secret. Ensure that only
// authorized users can add 'validation_shared_keys' to a keyset. You can rotate keys by appending (pushing) a new key to
// the list of 'validation_shared_keys' and removing any superseded keys. You must specify 'public_keys' or
// 'validation_shared_keys' (or both). The keys in 'public_keys' are checked first.
func (o NetworkServicesEdgeCacheKeysetOutput) ValidationSharedKeys() NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayOutput {
	return o.ApplyT(func(v *NetworkServicesEdgeCacheKeyset) NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayOutput {
		return v.ValidationSharedKeys
	}).(NetworkServicesEdgeCacheKeysetValidationSharedKeyArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkServicesEdgeCacheKeysetInput)(nil)).Elem(), &NetworkServicesEdgeCacheKeyset{})
	pulumi.RegisterOutputType(NetworkServicesEdgeCacheKeysetOutput{})
}
