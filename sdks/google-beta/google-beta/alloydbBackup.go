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

type AlloydbBackup struct {
	pulumi.CustomResourceState

	AlloydbBackupId pulumi.StringOutput `pulumi:"alloydbBackupId"`
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// The ID of the alloydb backup.
	BackupId pulumi.StringOutput `pulumi:"backupId"`
	// The full resource name of the backup source cluster (e.g.,
	// projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName pulumi.StringOutput `pulumi:"clusterName"`
	// Output only. The system-generated UID of the cluster which was used to create this resource.
	ClusterUid pulumi.StringOutput `pulumi:"clusterUid"`
	// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// User-provided description of the backup.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User-settable and human-readable display name for the Backup.
	DisplayName          pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	EffectiveLabels      pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
	// (customer-managed encryption key).
	EncryptionConfig AlloydbBackupEncryptionConfigPtrOutput `pulumi:"encryptionConfig"`
	// EncryptionInfo describes the encryption information of a cluster or a backup.
	EncryptionInfos AlloydbBackupEncryptionInfoArrayOutput `pulumi:"encryptionInfos"`
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy. Once the expiry quantity
	// is over retention, the backup is eligible to be garbage collected.
	ExpiryQuantities AlloydbBackupExpiryQuantityArrayOutput `pulumi:"expiryQuantities"`
	// Output only. The time at which after the backup is eligible to be garbage collected. It is the duration specified by the
	// backup's retention policy, added to the backup's createTime.
	ExpiryTime pulumi.StringOutput `pulumi:"expiryTime"`
	// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
	// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the alloydb backup should reside.
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively
	// updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. The size of the backup in bytes.
	SizeBytes pulumi.StringOutput `pulumi:"sizeBytes"`
	// Output only. The current state of the backup.
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput         `pulumi:"terraformLabels"`
	Timeouts        AlloydbBackupTimeoutsPtrOutput `pulumi:"timeouts"`
	// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
	// "AUTOMATED", "CONTINUOUS"]
	Type pulumi.StringOutput `pulumi:"type"`
	// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is
	// retained until it is deleted.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewAlloydbBackup registers a new resource with the given unique name, arguments, and options.
func NewAlloydbBackup(ctx *pulumi.Context,
	name string, args *AlloydbBackupArgs, opts ...pulumi.ResourceOption) (*AlloydbBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BackupId == nil {
		return nil, errors.New("invalid value for required argument 'BackupId'")
	}
	if args.ClusterName == nil {
		return nil, errors.New("invalid value for required argument 'ClusterName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AlloydbBackup
	err = ctx.RegisterPackageResource("google-beta:index/alloydbBackup:AlloydbBackup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlloydbBackup gets an existing AlloydbBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlloydbBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlloydbBackupState, opts ...pulumi.ResourceOption) (*AlloydbBackup, error) {
	var resource AlloydbBackup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/alloydbBackup:AlloydbBackup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AlloydbBackup resources.
type alloydbBackupState struct {
	AlloydbBackupId *string `pulumi:"alloydbBackupId"`
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// The ID of the alloydb backup.
	BackupId *string `pulumi:"backupId"`
	// The full resource name of the backup source cluster (e.g.,
	// projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName *string `pulumi:"clusterName"`
	// Output only. The system-generated UID of the cluster which was used to create this resource.
	ClusterUid *string `pulumi:"clusterUid"`
	// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime *string `pulumi:"deleteTime"`
	// User-provided description of the backup.
	Description *string `pulumi:"description"`
	// User-settable and human-readable display name for the Backup.
	DisplayName          *string           `pulumi:"displayName"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	EffectiveLabels      map[string]string `pulumi:"effectiveLabels"`
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
	// (customer-managed encryption key).
	EncryptionConfig *AlloydbBackupEncryptionConfig `pulumi:"encryptionConfig"`
	// EncryptionInfo describes the encryption information of a cluster or a backup.
	EncryptionInfos []AlloydbBackupEncryptionInfo `pulumi:"encryptionInfos"`
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag *string `pulumi:"etag"`
	// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy. Once the expiry quantity
	// is over retention, the backup is eligible to be garbage collected.
	ExpiryQuantities []AlloydbBackupExpiryQuantity `pulumi:"expiryQuantities"`
	// Output only. The time at which after the backup is eligible to be garbage collected. It is the duration specified by the
	// backup's retention policy, added to the backup's createTime.
	ExpiryTime *string `pulumi:"expiryTime"`
	// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
	// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the alloydb backup should reside.
	Location *string `pulumi:"location"`
	// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively
	// updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. The size of the backup in bytes.
	SizeBytes *string `pulumi:"sizeBytes"`
	// Output only. The current state of the backup.
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string      `pulumi:"terraformLabels"`
	Timeouts        *AlloydbBackupTimeouts `pulumi:"timeouts"`
	// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
	// "AUTOMATED", "CONTINUOUS"]
	Type *string `pulumi:"type"`
	// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is
	// retained until it is deleted.
	Uid *string `pulumi:"uid"`
	// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type AlloydbBackupState struct {
	AlloydbBackupId pulumi.StringPtrInput
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// The ID of the alloydb backup.
	BackupId pulumi.StringPtrInput
	// The full resource name of the backup source cluster (e.g.,
	// projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName pulumi.StringPtrInput
	// Output only. The system-generated UID of the cluster which was used to create this resource.
	ClusterUid pulumi.StringPtrInput
	// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	DeleteTime pulumi.StringPtrInput
	// User-provided description of the backup.
	Description pulumi.StringPtrInput
	// User-settable and human-readable display name for the Backup.
	DisplayName          pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	EffectiveLabels      pulumi.StringMapInput
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
	// (customer-managed encryption key).
	EncryptionConfig AlloydbBackupEncryptionConfigPtrInput
	// EncryptionInfo describes the encryption information of a cluster or a backup.
	EncryptionInfos AlloydbBackupEncryptionInfoArrayInput
	// For Resource freshness validation (https://google.aip.dev/154)
	Etag pulumi.StringPtrInput
	// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy. Once the expiry quantity
	// is over retention, the backup is eligible to be garbage collected.
	ExpiryQuantities AlloydbBackupExpiryQuantityArrayInput
	// Output only. The time at which after the backup is eligible to be garbage collected. It is the duration specified by the
	// backup's retention policy, added to the backup's createTime.
	ExpiryTime pulumi.StringPtrInput
	// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
	// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the alloydb backup should reside.
	Location pulumi.StringPtrInput
	// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively
	// updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
	Reconciling pulumi.BoolPtrInput
	// Output only. The size of the backup in bytes.
	SizeBytes pulumi.StringPtrInput
	// Output only. The current state of the backup.
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        AlloydbBackupTimeoutsPtrInput
	// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
	// "AUTOMATED", "CONTINUOUS"]
	Type pulumi.StringPtrInput
	// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is
	// retained until it is deleted.
	Uid pulumi.StringPtrInput
	// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (AlloydbBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*alloydbBackupState)(nil)).Elem()
}

type alloydbBackupArgs struct {
	AlloydbBackupId *string `pulumi:"alloydbBackupId"`
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// The ID of the alloydb backup.
	BackupId string `pulumi:"backupId"`
	// The full resource name of the backup source cluster (e.g.,
	// projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName string `pulumi:"clusterName"`
	// User-provided description of the backup.
	Description *string `pulumi:"description"`
	// User-settable and human-readable display name for the Backup.
	DisplayName *string `pulumi:"displayName"`
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
	// (customer-managed encryption key).
	EncryptionConfig *AlloydbBackupEncryptionConfig `pulumi:"encryptionConfig"`
	// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
	// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the alloydb backup should reside.
	Location string                 `pulumi:"location"`
	Project  *string                `pulumi:"project"`
	Timeouts *AlloydbBackupTimeouts `pulumi:"timeouts"`
	// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
	// "AUTOMATED", "CONTINUOUS"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AlloydbBackup resource.
type AlloydbBackupArgs struct {
	AlloydbBackupId pulumi.StringPtrInput
	// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
	// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
	// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
	// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// The ID of the alloydb backup.
	BackupId pulumi.StringInput
	// The full resource name of the backup source cluster (e.g.,
	// projects/{project}/locations/{location}/clusters/{clusterId}).
	ClusterName pulumi.StringInput
	// User-provided description of the backup.
	Description pulumi.StringPtrInput
	// User-settable and human-readable display name for the Backup.
	DisplayName pulumi.StringPtrInput
	// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
	// (customer-managed encryption key).
	EncryptionConfig AlloydbBackupEncryptionConfigPtrInput
	// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
	// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
	// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the alloydb backup should reside.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	Timeouts AlloydbBackupTimeoutsPtrInput
	// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
	// "AUTOMATED", "CONTINUOUS"]
	Type pulumi.StringPtrInput
}

func (AlloydbBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alloydbBackupArgs)(nil)).Elem()
}

type AlloydbBackupInput interface {
	pulumi.Input

	ToAlloydbBackupOutput() AlloydbBackupOutput
	ToAlloydbBackupOutputWithContext(ctx context.Context) AlloydbBackupOutput
}

func (*AlloydbBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**AlloydbBackup)(nil)).Elem()
}

func (i *AlloydbBackup) ToAlloydbBackupOutput() AlloydbBackupOutput {
	return i.ToAlloydbBackupOutputWithContext(context.Background())
}

func (i *AlloydbBackup) ToAlloydbBackupOutputWithContext(ctx context.Context) AlloydbBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlloydbBackupOutput)
}

type AlloydbBackupOutput struct{ *pulumi.OutputState }

func (AlloydbBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlloydbBackup)(nil)).Elem()
}

func (o AlloydbBackupOutput) ToAlloydbBackupOutput() AlloydbBackupOutput {
	return o
}

func (o AlloydbBackupOutput) ToAlloydbBackupOutputWithContext(ctx context.Context) AlloydbBackupOutput {
	return o
}

func (o AlloydbBackupOutput) AlloydbBackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.AlloydbBackupId }).(pulumi.StringOutput)
}

// Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels.
// https://google.aip.dev/128 An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
// "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the annotations present in your
// configuration. Please refer to the field 'effective_annotations' for all of the annotations present on the resource.
func (o AlloydbBackupOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// The ID of the alloydb backup.
func (o AlloydbBackupOutput) BackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.BackupId }).(pulumi.StringOutput)
}

// The full resource name of the backup source cluster (e.g.,
// projects/{project}/locations/{location}/clusters/{clusterId}).
func (o AlloydbBackupOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.ClusterName }).(pulumi.StringOutput)
}

// Output only. The system-generated UID of the cluster which was used to create this resource.
func (o AlloydbBackupOutput) ClusterUid() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.ClusterUid }).(pulumi.StringOutput)
}

// Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o AlloydbBackupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o AlloydbBackupOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// User-provided description of the backup.
func (o AlloydbBackupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User-settable and human-readable display name for the Backup.
func (o AlloydbBackupOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o AlloydbBackupOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

func (o AlloydbBackupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK
// (customer-managed encryption key).
func (o AlloydbBackupOutput) EncryptionConfig() AlloydbBackupEncryptionConfigPtrOutput {
	return o.ApplyT(func(v *AlloydbBackup) AlloydbBackupEncryptionConfigPtrOutput { return v.EncryptionConfig }).(AlloydbBackupEncryptionConfigPtrOutput)
}

// EncryptionInfo describes the encryption information of a cluster or a backup.
func (o AlloydbBackupOutput) EncryptionInfos() AlloydbBackupEncryptionInfoArrayOutput {
	return o.ApplyT(func(v *AlloydbBackup) AlloydbBackupEncryptionInfoArrayOutput { return v.EncryptionInfos }).(AlloydbBackupEncryptionInfoArrayOutput)
}

// For Resource freshness validation (https://google.aip.dev/154)
func (o AlloydbBackupOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy. Once the expiry quantity
// is over retention, the backup is eligible to be garbage collected.
func (o AlloydbBackupOutput) ExpiryQuantities() AlloydbBackupExpiryQuantityArrayOutput {
	return o.ApplyT(func(v *AlloydbBackup) AlloydbBackupExpiryQuantityArrayOutput { return v.ExpiryQuantities }).(AlloydbBackupExpiryQuantityArrayOutput)
}

// Output only. The time at which after the backup is eligible to be garbage collected. It is the duration specified by the
// backup's retention policy, added to the backup's createTime.
func (o AlloydbBackupOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.ExpiryTime }).(pulumi.StringOutput)
}

// User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name":
// "wrench", "mass": "1.3kg", "count": "3" }. **Note**: This field is non-authoritative, and will only manage the labels
// present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o AlloydbBackupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the alloydb backup should reside.
func (o AlloydbBackupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
func (o AlloydbBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AlloydbBackupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively
// updating the resource. This can happen due to user-triggered updates or system actions like failover or maintenance.
func (o AlloydbBackupOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Output only. The size of the backup in bytes.
func (o AlloydbBackupOutput) SizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.SizeBytes }).(pulumi.StringOutput)
}

// Output only. The current state of the backup.
func (o AlloydbBackupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o AlloydbBackupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o AlloydbBackupOutput) Timeouts() AlloydbBackupTimeoutsPtrOutput {
	return o.ApplyT(func(v *AlloydbBackup) AlloydbBackupTimeoutsPtrOutput { return v.Timeouts }).(AlloydbBackupTimeoutsPtrOutput)
}

// The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND",
// "AUTOMATED", "CONTINUOUS"]
func (o AlloydbBackupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is
// retained until it is deleted.
func (o AlloydbBackupOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o AlloydbBackupOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlloydbBackup) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlloydbBackupInput)(nil)).Elem(), &AlloydbBackup{})
	pulumi.RegisterOutputType(AlloydbBackupOutput{})
}
