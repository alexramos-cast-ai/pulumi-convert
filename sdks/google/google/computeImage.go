// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ComputeImage struct {
	pulumi.CustomResourceState

	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.Float64Output `pulumi:"archiveSizeBytes"`
	ComputeImageId   pulumi.StringOutput  `pulumi:"computeImageId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb      pulumi.Float64Output   `pulumi:"diskSizeGb"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
	// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
	// family must comply with RFC1035.
	Family pulumi.StringPtrOutput `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	GuestOsFeatures ComputeImageGuestOsFeatureArrayOutput `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
	// you must provide the same key if you use the image later (e.g. to create a disk from the image)
	ImageEncryptionKey ComputeImageImageEncryptionKeyPtrOutput `pulumi:"imageEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Any applicable license URI.
	Licenses pulumi.StringArrayOutput `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The parameters of the raw disk image.
	RawDisk  ComputeImageRawDiskPtrOutput `pulumi:"rawDisk"`
	SelfLink pulumi.StringOutput          `pulumi:"selfLink"`
	// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
	// not both to create an image.
	SourceDisk pulumi.StringPtrOutput `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
	// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
	SourceImage pulumi.StringPtrOutput `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
	// sourceDisk URL
	SourceSnapshot pulumi.StringPtrOutput `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
	// https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayOutput `pulumi:"storageLocations"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput        `pulumi:"terraformLabels"`
	Timeouts        ComputeImageTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewComputeImage registers a new resource with the given unique name, arguments, and options.
func NewComputeImage(ctx *pulumi.Context,
	name string, args *ComputeImageArgs, opts ...pulumi.ResourceOption) (*ComputeImage, error) {
	if args == nil {
		args = &ComputeImageArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ComputeImage
	err = ctx.RegisterPackageResource("google:index/computeImage:ComputeImage", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComputeImage gets an existing ComputeImage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComputeImage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComputeImageState, opts ...pulumi.ResourceOption) (*ComputeImage, error) {
	var resource ComputeImage
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/computeImage:ComputeImage", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ComputeImage resources.
type computeImageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes *float64 `pulumi:"archiveSizeBytes"`
	ComputeImageId   *string  `pulumi:"computeImageId"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb      *float64          `pulumi:"diskSizeGb"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
	// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
	// family must comply with RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	GuestOsFeatures []ComputeImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
	// you must provide the same key if you use the image later (e.g. to create a disk from the image)
	ImageEncryptionKey *ComputeImageImageEncryptionKey `pulumi:"imageEncryptionKey"`
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint *string `pulumi:"labelFingerprint"`
	// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The parameters of the raw disk image.
	RawDisk  *ComputeImageRawDisk `pulumi:"rawDisk"`
	SelfLink *string              `pulumi:"selfLink"`
	// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
	// not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
	// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
	// sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
	// https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations []string `pulumi:"storageLocations"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string     `pulumi:"terraformLabels"`
	Timeouts        *ComputeImageTimeouts `pulumi:"timeouts"`
}

type ComputeImageState struct {
	// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
	ArchiveSizeBytes pulumi.Float64PtrInput
	ComputeImageId   pulumi.StringPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb      pulumi.Float64PtrInput
	EffectiveLabels pulumi.StringMapInput
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
	// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
	// family must comply with RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	GuestOsFeatures ComputeImageGuestOsFeatureArrayInput
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
	// you must provide the same key if you use the image later (e.g. to create a disk from the image)
	ImageEncryptionKey ComputeImageImageEncryptionKeyPtrInput
	// The fingerprint used for optimistic locking of this resource. Used internally during updates.
	LabelFingerprint pulumi.StringPtrInput
	// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The parameters of the raw disk image.
	RawDisk  ComputeImageRawDiskPtrInput
	SelfLink pulumi.StringPtrInput
	// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
	// not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
	// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
	// sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
	// https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ComputeImageTimeoutsPtrInput
}

func (ComputeImageState) ElementType() reflect.Type {
	return reflect.TypeOf((*computeImageState)(nil)).Elem()
}

type computeImageArgs struct {
	ComputeImageId *string `pulumi:"computeImageId"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb *float64 `pulumi:"diskSizeGb"`
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
	// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
	// family must comply with RFC1035.
	Family *string `pulumi:"family"`
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	GuestOsFeatures []ComputeImageGuestOsFeature `pulumi:"guestOsFeatures"`
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
	// you must provide the same key if you use the image later (e.g. to create a disk from the image)
	ImageEncryptionKey *ComputeImageImageEncryptionKey `pulumi:"imageEncryptionKey"`
	// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Any applicable license URI.
	Licenses []string `pulumi:"licenses"`
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The parameters of the raw disk image.
	RawDisk *ComputeImageRawDisk `pulumi:"rawDisk"`
	// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
	// not both to create an image.
	SourceDisk *string `pulumi:"sourceDisk"`
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
	// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
	SourceImage *string `pulumi:"sourceImage"`
	// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
	// sourceDisk URL
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
	// https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations []string              `pulumi:"storageLocations"`
	Timeouts         *ComputeImageTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ComputeImage resource.
type ComputeImageArgs struct {
	ComputeImageId pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Size of the image when restored onto a persistent disk (in GB).
	DiskSizeGb pulumi.Float64PtrInput
	// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
	// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
	// family must comply with RFC1035.
	Family pulumi.StringPtrInput
	// A list of features to enable on the guest operating system. Applicable only for bootable images.
	GuestOsFeatures ComputeImageGuestOsFeatureArrayInput
	// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
	// you must provide the same key if you use the image later (e.g. to create a disk from the image)
	ImageEncryptionKey ComputeImageImageEncryptionKeyPtrInput
	// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Any applicable license URI.
	Licenses pulumi.StringArrayInput
	// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
	// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
	// digit, except the last character, which cannot be a dash.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The parameters of the raw disk image.
	RawDisk ComputeImageRawDiskPtrInput
	// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
	// not both to create an image.
	SourceDisk pulumi.StringPtrInput
	// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
	// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
	SourceImage pulumi.StringPtrInput
	// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
	// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
	// sourceDisk URL
	SourceSnapshot pulumi.StringPtrInput
	// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
	// https://cloud.google.com/compute/docs/reference/rest/v1/images
	StorageLocations pulumi.StringArrayInput
	Timeouts         ComputeImageTimeoutsPtrInput
}

func (ComputeImageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*computeImageArgs)(nil)).Elem()
}

type ComputeImageInput interface {
	pulumi.Input

	ToComputeImageOutput() ComputeImageOutput
	ToComputeImageOutputWithContext(ctx context.Context) ComputeImageOutput
}

func (*ComputeImage) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeImage)(nil)).Elem()
}

func (i *ComputeImage) ToComputeImageOutput() ComputeImageOutput {
	return i.ToComputeImageOutputWithContext(context.Background())
}

func (i *ComputeImage) ToComputeImageOutputWithContext(ctx context.Context) ComputeImageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComputeImageOutput)
}

type ComputeImageOutput struct{ *pulumi.OutputState }

func (ComputeImageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ComputeImage)(nil)).Elem()
}

func (o ComputeImageOutput) ToComputeImageOutput() ComputeImageOutput {
	return o
}

func (o ComputeImageOutput) ToComputeImageOutputWithContext(ctx context.Context) ComputeImageOutput {
	return o
}

// Size of the image tar.gz archive stored in Google Cloud Storage (in bytes).
func (o ComputeImageOutput) ArchiveSizeBytes() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeImage) pulumi.Float64Output { return v.ArchiveSizeBytes }).(pulumi.Float64Output)
}

func (o ComputeImageOutput) ComputeImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.ComputeImageId }).(pulumi.StringOutput)
}

// Creation timestamp in RFC3339 text format.
func (o ComputeImageOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o ComputeImageOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Size of the image when restored onto a persistent disk (in GB).
func (o ComputeImageOutput) DiskSizeGb() pulumi.Float64Output {
	return o.ApplyT(func(v *ComputeImage) pulumi.Float64Output { return v.DiskSizeGb }).(pulumi.Float64Output)
}

func (o ComputeImageOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the image family to which this image belongs. You can create disks by specifying an image family instead of
// a specific image name. The image family always returns its latest image that is not deprecated. The name of the image
// family must comply with RFC1035.
func (o ComputeImageOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringPtrOutput { return v.Family }).(pulumi.StringPtrOutput)
}

// A list of features to enable on the guest operating system. Applicable only for bootable images.
func (o ComputeImageOutput) GuestOsFeatures() ComputeImageGuestOsFeatureArrayOutput {
	return o.ApplyT(func(v *ComputeImage) ComputeImageGuestOsFeatureArrayOutput { return v.GuestOsFeatures }).(ComputeImageGuestOsFeatureArrayOutput)
}

// Encrypts the image using a customer-supplied encryption key. After you encrypt an image with a customer-supplied key,
// you must provide the same key if you use the image later (e.g. to create a disk from the image)
func (o ComputeImageOutput) ImageEncryptionKey() ComputeImageImageEncryptionKeyPtrOutput {
	return o.ApplyT(func(v *ComputeImage) ComputeImageImageEncryptionKeyPtrOutput { return v.ImageEncryptionKey }).(ComputeImageImageEncryptionKeyPtrOutput)
}

// The fingerprint used for optimistic locking of this resource. Used internally during updates.
func (o ComputeImageOutput) LabelFingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.LabelFingerprint }).(pulumi.StringOutput)
}

// Labels to apply to this Image. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o ComputeImageOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Any applicable license URI.
func (o ComputeImageOutput) Licenses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringArrayOutput { return v.Licenses }).(pulumi.StringArrayOutput)
}

// Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression 'a-z?' which
// means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or
// digit, except the last character, which cannot be a dash.
func (o ComputeImageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ComputeImageOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The parameters of the raw disk image.
func (o ComputeImageOutput) RawDisk() ComputeImageRawDiskPtrOutput {
	return o.ApplyT(func(v *ComputeImage) ComputeImageRawDiskPtrOutput { return v.RawDisk }).(ComputeImageRawDiskPtrOutput)
}

func (o ComputeImageOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The source disk to create this image based on. You must provide either this property or the rawDisk.source property but
// not both to create an image.
func (o ComputeImageOutput) SourceDisk() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringPtrOutput { return v.SourceDisk }).(pulumi.StringPtrOutput)
}

// URL of the source image used to create this image. In order to create an image, you must provide the full or partial URL
// of one of the following: * The selfLink URL * This property * The rawDisk.source URL * The sourceDisk URL
func (o ComputeImageOutput) SourceImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringPtrOutput { return v.SourceImage }).(pulumi.StringPtrOutput)
}

// URL of the source snapshot used to create this image. In order to create an image, you must provide the full or partial
// URL of one of the following: * The selfLink URL * This property * The sourceImage URL * The rawDisk.source URL * The
// sourceDisk URL
func (o ComputeImageOutput) SourceSnapshot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringPtrOutput { return v.SourceSnapshot }).(pulumi.StringPtrOutput)
}

// Cloud Storage bucket storage location of the image (regional or multi-regional). Reference link:
// https://cloud.google.com/compute/docs/reference/rest/v1/images
func (o ComputeImageOutput) StorageLocations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringArrayOutput { return v.StorageLocations }).(pulumi.StringArrayOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ComputeImageOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ComputeImage) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ComputeImageOutput) Timeouts() ComputeImageTimeoutsPtrOutput {
	return o.ApplyT(func(v *ComputeImage) ComputeImageTimeoutsPtrOutput { return v.Timeouts }).(ComputeImageTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComputeImageInput)(nil)).Elem(), &ComputeImage{})
	pulumi.RegisterOutputType(ComputeImageOutput{})
}
