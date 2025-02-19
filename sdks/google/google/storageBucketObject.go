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

type StorageBucketObject struct {
	pulumi.CustomResourceState

	// The name of the containing bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
	// users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrOutput `pulumi:"cacheControl"`
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
	// the raw contents of the object, please define an output.
	Content pulumi.StringOutput `pulumi:"content"`
	// Content-Disposition of the object data.
	ContentDisposition pulumi.StringPtrOutput `pulumi:"contentDisposition"`
	// Content-Encoding of the object data.
	ContentEncoding pulumi.StringPtrOutput `pulumi:"contentEncoding"`
	// Content-Language of the object data.
	ContentLanguage pulumi.StringPtrOutput `pulumi:"contentLanguage"`
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringOutput `pulumi:"contentType"`
	// Base 64 CRC32 hash of the uploaded data.
	Crc32c pulumi.StringOutput `pulumi:"crc32c"`
	// Encryption key; encoded using base64.
	CustomerEncryption StorageBucketObjectCustomerEncryptionPtrOutput `pulumi:"customerEncryption"`
	DetectMd5hash      pulumi.StringPtrOutput                         `pulumi:"detectMd5hash"`
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
	// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
	// will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrOutput `pulumi:"eventBasedHold"`
	// The content generation of this object. Used for object versioning and soft delete.
	Generation pulumi.Float64Output `pulumi:"generation"`
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
	// value, if any.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringOutput `pulumi:"md5hash"`
	// A url reference to download this object.
	MediaLink pulumi.StringOutput `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the object. Use this field in interpolations with google.StorageObjectAcl to recreate
	// google.StorageObjectAcl resources when your google.StorageBucketObject is recreated.
	OutputName pulumi.StringOutput `pulumi:"outputName"`
	// Object level retention configuration.
	Retention StorageBucketObjectRetentionPtrOutput `pulumi:"retention"`
	// A url reference to this object.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined if content is not.
	Source                pulumi.StringPtrOutput `pulumi:"source"`
	StorageBucketObjectId pulumi.StringOutput    `pulumi:"storageBucketObjectId"`
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
	// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	StorageClass pulumi.StringOutput `pulumi:"storageClass"`
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold pulumi.BoolPtrOutput                 `pulumi:"temporaryHold"`
	Timeouts      StorageBucketObjectTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewStorageBucketObject registers a new resource with the given unique name, arguments, and options.
func NewStorageBucketObject(ctx *pulumi.Context,
	name string, args *StorageBucketObjectArgs, opts ...pulumi.ResourceOption) (*StorageBucketObject, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Content != nil {
		args.Content = pulumi.ToSecret(args.Content).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"content",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource StorageBucketObject
	err = ctx.RegisterPackageResource("google:index/storageBucketObject:StorageBucketObject", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageBucketObject gets an existing StorageBucketObject resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageBucketObject(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageBucketObjectState, opts ...pulumi.ResourceOption) (*StorageBucketObject, error) {
	var resource StorageBucketObject
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/storageBucketObject:StorageBucketObject", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageBucketObject resources.
type storageBucketObjectState struct {
	// The name of the containing bucket.
	Bucket *string `pulumi:"bucket"`
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
	// users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
	// the raw contents of the object, please define an output.
	Content *string `pulumi:"content"`
	// Content-Disposition of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Content-Encoding of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// Content-Language of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `pulumi:"contentType"`
	// Base 64 CRC32 hash of the uploaded data.
	Crc32c *string `pulumi:"crc32c"`
	// Encryption key; encoded using base64.
	CustomerEncryption *StorageBucketObjectCustomerEncryption `pulumi:"customerEncryption"`
	DetectMd5hash      *string                                `pulumi:"detectMd5hash"`
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
	// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
	// will be subject to bucket-level retention (if any).
	EventBasedHold *bool `pulumi:"eventBasedHold"`
	// The content generation of this object. Used for object versioning and soft delete.
	Generation *float64 `pulumi:"generation"`
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
	// value, if any.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Base 64 MD5 hash of the uploaded data.
	Md5hash *string `pulumi:"md5hash"`
	// A url reference to download this object.
	MediaLink *string `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name *string `pulumi:"name"`
	// The name of the object. Use this field in interpolations with google.StorageObjectAcl to recreate
	// google.StorageObjectAcl resources when your google.StorageBucketObject is recreated.
	OutputName *string `pulumi:"outputName"`
	// Object level retention configuration.
	Retention *StorageBucketObjectRetention `pulumi:"retention"`
	// A url reference to this object.
	SelfLink *string `pulumi:"selfLink"`
	// A path to the data you want to upload. Must be defined if content is not.
	Source                *string `pulumi:"source"`
	StorageBucketObjectId *string `pulumi:"storageBucketObjectId"`
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
	// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	StorageClass *string `pulumi:"storageClass"`
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold *bool                        `pulumi:"temporaryHold"`
	Timeouts      *StorageBucketObjectTimeouts `pulumi:"timeouts"`
}

type StorageBucketObjectState struct {
	// The name of the containing bucket.
	Bucket pulumi.StringPtrInput
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
	// users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
	// the raw contents of the object, please define an output.
	Content pulumi.StringPtrInput
	// Content-Disposition of the object data.
	ContentDisposition pulumi.StringPtrInput
	// Content-Encoding of the object data.
	ContentEncoding pulumi.StringPtrInput
	// Content-Language of the object data.
	ContentLanguage pulumi.StringPtrInput
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringPtrInput
	// Base 64 CRC32 hash of the uploaded data.
	Crc32c pulumi.StringPtrInput
	// Encryption key; encoded using base64.
	CustomerEncryption StorageBucketObjectCustomerEncryptionPtrInput
	DetectMd5hash      pulumi.StringPtrInput
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
	// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
	// will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrInput
	// The content generation of this object. Used for object versioning and soft delete.
	Generation pulumi.Float64PtrInput
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
	// value, if any.
	KmsKeyName pulumi.StringPtrInput
	// Base 64 MD5 hash of the uploaded data.
	Md5hash pulumi.StringPtrInput
	// A url reference to download this object.
	MediaLink pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name pulumi.StringPtrInput
	// The name of the object. Use this field in interpolations with google.StorageObjectAcl to recreate
	// google.StorageObjectAcl resources when your google.StorageBucketObject is recreated.
	OutputName pulumi.StringPtrInput
	// Object level retention configuration.
	Retention StorageBucketObjectRetentionPtrInput
	// A url reference to this object.
	SelfLink pulumi.StringPtrInput
	// A path to the data you want to upload. Must be defined if content is not.
	Source                pulumi.StringPtrInput
	StorageBucketObjectId pulumi.StringPtrInput
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
	// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	StorageClass pulumi.StringPtrInput
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold pulumi.BoolPtrInput
	Timeouts      StorageBucketObjectTimeoutsPtrInput
}

func (StorageBucketObjectState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketObjectState)(nil)).Elem()
}

type storageBucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket string `pulumi:"bucket"`
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
	// users, the default will be public, max-age=3600
	CacheControl *string `pulumi:"cacheControl"`
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
	// the raw contents of the object, please define an output.
	Content *string `pulumi:"content"`
	// Content-Disposition of the object data.
	ContentDisposition *string `pulumi:"contentDisposition"`
	// Content-Encoding of the object data.
	ContentEncoding *string `pulumi:"contentEncoding"`
	// Content-Language of the object data.
	ContentLanguage *string `pulumi:"contentLanguage"`
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType *string `pulumi:"contentType"`
	// Encryption key; encoded using base64.
	CustomerEncryption *StorageBucketObjectCustomerEncryption `pulumi:"customerEncryption"`
	DetectMd5hash      *string                                `pulumi:"detectMd5hash"`
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
	// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
	// will be subject to bucket-level retention (if any).
	EventBasedHold *bool `pulumi:"eventBasedHold"`
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
	// value, if any.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name *string `pulumi:"name"`
	// Object level retention configuration.
	Retention *StorageBucketObjectRetention `pulumi:"retention"`
	// A path to the data you want to upload. Must be defined if content is not.
	Source                *string `pulumi:"source"`
	StorageBucketObjectId *string `pulumi:"storageBucketObjectId"`
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
	// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	StorageClass *string `pulumi:"storageClass"`
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold *bool                        `pulumi:"temporaryHold"`
	Timeouts      *StorageBucketObjectTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a StorageBucketObject resource.
type StorageBucketObjectArgs struct {
	// The name of the containing bucket.
	Bucket pulumi.StringInput
	// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
	// users, the default will be public, max-age=3600
	CacheControl pulumi.StringPtrInput
	// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
	// the raw contents of the object, please define an output.
	Content pulumi.StringPtrInput
	// Content-Disposition of the object data.
	ContentDisposition pulumi.StringPtrInput
	// Content-Encoding of the object data.
	ContentEncoding pulumi.StringPtrInput
	// Content-Language of the object data.
	ContentLanguage pulumi.StringPtrInput
	// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
	ContentType pulumi.StringPtrInput
	// Encryption key; encoded using base64.
	CustomerEncryption StorageBucketObjectCustomerEncryptionPtrInput
	DetectMd5hash      pulumi.StringPtrInput
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
	// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
	// will be subject to bucket-level retention (if any).
	EventBasedHold pulumi.BoolPtrInput
	// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
	// value, if any.
	KmsKeyName pulumi.StringPtrInput
	// User-provided metadata, in key/value pairs.
	Metadata pulumi.StringMapInput
	// The name of the object. If you're interpolating the name of this object, see output_name instead.
	Name pulumi.StringPtrInput
	// Object level retention configuration.
	Retention StorageBucketObjectRetentionPtrInput
	// A path to the data you want to upload. Must be defined if content is not.
	Source                pulumi.StringPtrInput
	StorageBucketObjectId pulumi.StringPtrInput
	// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
	// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
	StorageClass pulumi.StringPtrInput
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
	// overwrites.
	TemporaryHold pulumi.BoolPtrInput
	Timeouts      StorageBucketObjectTimeoutsPtrInput
}

func (StorageBucketObjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageBucketObjectArgs)(nil)).Elem()
}

type StorageBucketObjectInput interface {
	pulumi.Input

	ToStorageBucketObjectOutput() StorageBucketObjectOutput
	ToStorageBucketObjectOutputWithContext(ctx context.Context) StorageBucketObjectOutput
}

func (*StorageBucketObject) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketObject)(nil)).Elem()
}

func (i *StorageBucketObject) ToStorageBucketObjectOutput() StorageBucketObjectOutput {
	return i.ToStorageBucketObjectOutputWithContext(context.Background())
}

func (i *StorageBucketObject) ToStorageBucketObjectOutputWithContext(ctx context.Context) StorageBucketObjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBucketObjectOutput)
}

type StorageBucketObjectOutput struct{ *pulumi.OutputState }

func (StorageBucketObjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBucketObject)(nil)).Elem()
}

func (o StorageBucketObjectOutput) ToStorageBucketObjectOutput() StorageBucketObjectOutput {
	return o
}

func (o StorageBucketObjectOutput) ToStorageBucketObjectOutputWithContext(ctx context.Context) StorageBucketObjectOutput {
	return o
}

// The name of the containing bucket.
func (o StorageBucketObjectOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// Cache-Control directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous
// users, the default will be public, max-age=3600
func (o StorageBucketObjectOutput) CacheControl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.CacheControl }).(pulumi.StringPtrOutput)
}

// Data as string to be uploaded. Must be defined if source is not. Note: The content field is marked as sensitive. To view
// the raw contents of the object, please define an output.
func (o StorageBucketObjectOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.Content }).(pulumi.StringOutput)
}

// Content-Disposition of the object data.
func (o StorageBucketObjectOutput) ContentDisposition() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.ContentDisposition }).(pulumi.StringPtrOutput)
}

// Content-Encoding of the object data.
func (o StorageBucketObjectOutput) ContentEncoding() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.ContentEncoding }).(pulumi.StringPtrOutput)
}

// Content-Language of the object data.
func (o StorageBucketObjectOutput) ContentLanguage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.ContentLanguage }).(pulumi.StringPtrOutput)
}

// Content-Type of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
func (o StorageBucketObjectOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.ContentType }).(pulumi.StringOutput)
}

// Base 64 CRC32 hash of the uploaded data.
func (o StorageBucketObjectOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.Crc32c }).(pulumi.StringOutput)
}

// Encryption key; encoded using base64.
func (o StorageBucketObjectOutput) CustomerEncryption() StorageBucketObjectCustomerEncryptionPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) StorageBucketObjectCustomerEncryptionPtrOutput {
		return v.CustomerEncryption
	}).(StorageBucketObjectCustomerEncryptionPtrOutput)
}

func (o StorageBucketObjectOutput) DetectMd5hash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.DetectMd5hash }).(pulumi.StringPtrOutput)
}

// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is
// signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects
// will be subject to bucket-level retention (if any).
func (o StorageBucketObjectOutput) EventBasedHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.BoolPtrOutput { return v.EventBasedHold }).(pulumi.BoolPtrOutput)
}

// The content generation of this object. Used for object versioning and soft delete.
func (o StorageBucketObjectOutput) Generation() pulumi.Float64Output {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.Float64Output { return v.Generation }).(pulumi.Float64Output)
}

// Resource name of the Cloud KMS key that will be used to encrypt the object. Overrides the object metadata's kmsKeyName
// value, if any.
func (o StorageBucketObjectOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.KmsKeyName }).(pulumi.StringOutput)
}

// Base 64 MD5 hash of the uploaded data.
func (o StorageBucketObjectOutput) Md5hash() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.Md5hash }).(pulumi.StringOutput)
}

// A url reference to download this object.
func (o StorageBucketObjectOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.MediaLink }).(pulumi.StringOutput)
}

// User-provided metadata, in key/value pairs.
func (o StorageBucketObjectOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The name of the object. If you're interpolating the name of this object, see output_name instead.
func (o StorageBucketObjectOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the object. Use this field in interpolations with google.StorageObjectAcl to recreate
// google.StorageObjectAcl resources when your google.StorageBucketObject is recreated.
func (o StorageBucketObjectOutput) OutputName() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.OutputName }).(pulumi.StringOutput)
}

// Object level retention configuration.
func (o StorageBucketObjectOutput) Retention() StorageBucketObjectRetentionPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) StorageBucketObjectRetentionPtrOutput { return v.Retention }).(StorageBucketObjectRetentionPtrOutput)
}

// A url reference to this object.
func (o StorageBucketObjectOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// A path to the data you want to upload. Must be defined if content is not.
func (o StorageBucketObjectOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringPtrOutput { return v.Source }).(pulumi.StringPtrOutput)
}

func (o StorageBucketObjectOutput) StorageBucketObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.StorageBucketObjectId }).(pulumi.StringOutput)
}

// The StorageClass of the new bucket object. Supported values include: MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE,
// ARCHIVE. If not provided, this defaults to the bucket's default storage class or to a standard class.
func (o StorageBucketObjectOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.StringOutput { return v.StorageClass }).(pulumi.StringOutput)
}

// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and
// overwrites.
func (o StorageBucketObjectOutput) TemporaryHold() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) pulumi.BoolPtrOutput { return v.TemporaryHold }).(pulumi.BoolPtrOutput)
}

func (o StorageBucketObjectOutput) Timeouts() StorageBucketObjectTimeoutsPtrOutput {
	return o.ApplyT(func(v *StorageBucketObject) StorageBucketObjectTimeoutsPtrOutput { return v.Timeouts }).(StorageBucketObjectTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StorageBucketObjectInput)(nil)).Elem(), &StorageBucketObject{})
	pulumi.RegisterOutputType(StorageBucketObjectOutput{})
}
