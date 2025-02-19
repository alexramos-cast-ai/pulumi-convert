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

type ComputeBackendBucketSignedUrlKey struct {
	pulumi.CustomResourceState

	// The backend bucket this signed URL key belongs.
	BackendBucket                      pulumi.StringOutput `pulumi:"backendBucket"`
	ComputeBackendBucketSignedUrlKeyId pulumi.StringOutput `pulumi:"computeBackendBucketSignedUrlKeyId"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringOutput `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name     pulumi.StringOutput                               `pulumi:"name"`
	Project  pulumi.StringOutput                               `pulumi:"project"`
	Timeouts ComputeBackendBucketSignedUrlKeyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeBackendBucketSignedUrlKey registers a new resource with the given unique name, arguments, and options.
func NewComputeBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, args *ComputeBackendBucketSignedUrlKeyArgs, opts ...pulumi.ResourceOption) (*ComputeBackendBucketSignedUrlKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackendBucket == nil {
		return nil, errors.New("invalid value for required argument 'BackendBucket'")
	}
	if args.KeyValue == nil {
		return nil, errors.New("invalid value for required argument 'KeyValue'")
	}
	if args.KeyValue != nil {
		args.KeyValue = pulumi.ToSecret(args.KeyValue).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"keyValue",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeBackendBucketSignedUrlKey
	err = ctx.RegisterPackageResource("google:index/computeBackendBucketSignedUrlKey:ComputeBackendBucketSignedUrlKey", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeBackendBucketSignedUrlKey gets an existing ComputeBackendBucketSignedUrlKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeBackendBucketSignedUrlKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeBackendBucketSignedUrlKeyState, opts ...pulumi.ResourceOption) (*ComputeBackendBucketSignedUrlKey, error) {
	var resource ComputeBackendBucketSignedUrlKey
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeBackendBucketSignedUrlKey:ComputeBackendBucketSignedUrlKey", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeBackendBucketSignedUrlKey resources.
type computeBackendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket                      *string `pulumi:"backendBucket"`
	ComputeBackendBucketSignedUrlKeyId *string `pulumi:"computeBackendBucketSignedUrlKeyId"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue *string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name     *string                                   `pulumi:"name"`
	Project  *string                                   `pulumi:"project"`
	Timeouts *ComputeBackendBucketSignedUrlKeyTimeouts `pulumi:"timeouts"`
}

type ComputeBackendBucketSignedUrlKeyState struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket                      pulumi.StringPtrInput
	ComputeBackendBucketSignedUrlKeyId pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringPtrInput
	// Name of the signed URL key.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeBackendBucketSignedUrlKeyTimeoutsPtrInput
}

func (ComputeBackendBucketSignedUrlKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeBackendBucketSignedUrlKeyState)(nil)).Elem()
}

type computeBackendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket                      string  `pulumi:"backendBucket"`
	ComputeBackendBucketSignedUrlKeyId *string `pulumi:"computeBackendBucketSignedUrlKeyId"`
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue string `pulumi:"keyValue"`
	// Name of the signed URL key.
	Name     *string                                   `pulumi:"name"`
	Project  *string                                   `pulumi:"project"`
	Timeouts *ComputeBackendBucketSignedUrlKeyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeBackendBucketSignedUrlKey resource.
type ComputeBackendBucketSignedUrlKeyArgs struct {
	// The backend bucket this signed URL key belongs.
	BackendBucket                      pulumi.StringInput
	ComputeBackendBucketSignedUrlKeyId pulumi.StringPtrInput
	// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
	KeyValue pulumi.StringInput
	// Name of the signed URL key.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts ComputeBackendBucketSignedUrlKeyTimeoutsPtrInput
}

func (ComputeBackendBucketSignedUrlKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeBackendBucketSignedUrlKeyArgs)(nil)).Elem()
}

type ComputeBackendBucketSignedUrlKeyInput interface {
	pulumi.Input

	ToComputeBackendBucketSignedUrlKeyOutput() ComputeBackendBucketSignedUrlKeyOutput
	ToComputeBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) ComputeBackendBucketSignedUrlKeyOutput
}

func (*ComputeBackendBucketSignedUrlKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBackendBucketSignedUrlKey)(nil)).Elem()
}

func (i *ComputeBackendBucketSignedUrlKey) ToComputeBackendBucketSignedUrlKeyOutput() ComputeBackendBucketSignedUrlKeyOutput {
	return i.ToComputeBackendBucketSignedUrlKeyOutputWithContext(context.Background())
}

func (i *ComputeBackendBucketSignedUrlKey) ToComputeBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) ComputeBackendBucketSignedUrlKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeBackendBucketSignedUrlKeyOutput)
}

type ComputeBackendBucketSignedUrlKeyOutput struct{ *pulumi.OutputState }

func (ComputeBackendBucketSignedUrlKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeBackendBucketSignedUrlKey)(nil)).Elem()
}

func (o ComputeBackendBucketSignedUrlKeyOutput) ToComputeBackendBucketSignedUrlKeyOutput() ComputeBackendBucketSignedUrlKeyOutput {
	return o
}

func (o ComputeBackendBucketSignedUrlKeyOutput) ToComputeBackendBucketSignedUrlKeyOutputWithContext(ctx context.Context) ComputeBackendBucketSignedUrlKeyOutput {
	return o
}

// The backend bucket this signed URL key belongs.
func (o ComputeBackendBucketSignedUrlKeyOutput) BackendBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) pulumi.StringOutput { return v.BackendBucket }).(pulumi.StringOutput)
}

func (o ComputeBackendBucketSignedUrlKeyOutput) ComputeBackendBucketSignedUrlKeyId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) pulumi.StringOutput {
		return v.ComputeBackendBucketSignedUrlKeyId
	}).(pulumi.StringOutput)
}

// 128-bit key value used for signing the URL. The key value must be a valid RFC 4648 Section 5 base64url encoded string.
func (o ComputeBackendBucketSignedUrlKeyOutput) KeyValue() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) pulumi.StringOutput { return v.KeyValue }).(pulumi.StringOutput)
}

// Name of the signed URL key.
func (o ComputeBackendBucketSignedUrlKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeBackendBucketSignedUrlKeyOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o ComputeBackendBucketSignedUrlKeyOutput) Timeouts() ComputeBackendBucketSignedUrlKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeBackendBucketSignedUrlKey) ComputeBackendBucketSignedUrlKeyTimeoutsPtrOutput {
		return v.Timeouts
	}).(ComputeBackendBucketSignedUrlKeyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeBackendBucketSignedUrlKeyInput)(nil)).Elem(), &ComputeBackendBucketSignedUrlKey{})
	pulumi.RegisterOutputType(ComputeBackendBucketSignedUrlKeyOutput{})
}
