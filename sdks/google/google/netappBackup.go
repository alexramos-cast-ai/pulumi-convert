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

type NetappBackup struct {
	pulumi.CustomResourceState

	// Type of backup, manually created or created by a backup policy. Possible Values : [TYPE_UNSPECIFIED, MANUAL, SCHEDULED]
	BackupType pulumi.StringOutput `pulumi:"backupType"`
	// Backups of a volume build incrementally on top of each other. They form a "backup chain". Total size of all backups in a
	// chain in bytes = baseline backup size + sum(incremental backup size)
	ChainStorageBytes pulumi.StringOutput `pulumi:"chainStorageBytes"`
	// Create time of the backup. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location of the backup.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the backup. Needs to be unique per location.
	Name           pulumi.StringOutput `pulumi:"name"`
	NetappBackupId pulumi.StringOutput `pulumi:"netappBackupId"`
	Project        pulumi.StringOutput `pulumi:"project"`
	// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
	// initiate the backup creation. Format:
	// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	SourceSnapshot pulumi.StringPtrOutput `pulumi:"sourceSnapshot"`
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''
	SourceVolume pulumi.StringPtrOutput `pulumi:"sourceVolume"`
	// The state of the Backup Vault. Possible Values : [STATE_UNSPECIFIED, CREATING, UPLOADING, READY, DELETING, ERROR,
	// UPDATING]
	State pulumi.StringOutput `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput        `pulumi:"terraformLabels"`
	Timeouts        NetappBackupTimeoutsPtrOutput `pulumi:"timeouts"`
	// Name of the backup vault to store the backup in.
	VaultName pulumi.StringOutput `pulumi:"vaultName"`
	// Size of the file system when the backup was created. When creating a new volume from the backup, the volume capacity
	// will have to be at least as big.
	VolumeUsageBytes pulumi.StringOutput `pulumi:"volumeUsageBytes"`
}

// NewNetappBackup registers a new resource with the given unique name, arguments, and options.
func NewNetappBackup(ctx *pulumi.Context,
	name string, args *NetappBackupArgs, opts ...pulumi.ResourceOption) (*NetappBackup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.VaultName == nil {
		return nil, errors.New("invalid value for required argument 'VaultName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource NetappBackup
	err = ctx.RegisterPackageResource("google:index/netappBackup:NetappBackup", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetappBackup gets an existing NetappBackup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetappBackup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetappBackupState, opts ...pulumi.ResourceOption) (*NetappBackup, error) {
	var resource NetappBackup
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/netappBackup:NetappBackup", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetappBackup resources.
type netappBackupState struct {
	// Type of backup, manually created or created by a backup policy. Possible Values : [TYPE_UNSPECIFIED, MANUAL, SCHEDULED]
	BackupType *string `pulumi:"backupType"`
	// Backups of a volume build incrementally on top of each other. They form a "backup chain". Total size of all backups in a
	// chain in bytes = baseline backup size + sum(incremental backup size)
	ChainStorageBytes *string `pulumi:"chainStorageBytes"`
	// Create time of the backup. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime *string `pulumi:"createTime"`
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location of the backup.
	Location *string `pulumi:"location"`
	// The resource name of the backup. Needs to be unique per location.
	Name           *string `pulumi:"name"`
	NetappBackupId *string `pulumi:"netappBackupId"`
	Project        *string `pulumi:"project"`
	// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
	// initiate the backup creation. Format:
	// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''
	SourceVolume *string `pulumi:"sourceVolume"`
	// The state of the Backup Vault. Possible Values : [STATE_UNSPECIFIED, CREATING, UPLOADING, READY, DELETING, ERROR,
	// UPDATING]
	State *string `pulumi:"state"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string     `pulumi:"terraformLabels"`
	Timeouts        *NetappBackupTimeouts `pulumi:"timeouts"`
	// Name of the backup vault to store the backup in.
	VaultName *string `pulumi:"vaultName"`
	// Size of the file system when the backup was created. When creating a new volume from the backup, the volume capacity
	// will have to be at least as big.
	VolumeUsageBytes *string `pulumi:"volumeUsageBytes"`
}

type NetappBackupState struct {
	// Type of backup, manually created or created by a backup policy. Possible Values : [TYPE_UNSPECIFIED, MANUAL, SCHEDULED]
	BackupType pulumi.StringPtrInput
	// Backups of a volume build incrementally on top of each other. They form a "backup chain". Total size of all backups in a
	// chain in bytes = baseline backup size + sum(incremental backup size)
	ChainStorageBytes pulumi.StringPtrInput
	// Create time of the backup. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
	CreateTime pulumi.StringPtrInput
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location of the backup.
	Location pulumi.StringPtrInput
	// The resource name of the backup. Needs to be unique per location.
	Name           pulumi.StringPtrInput
	NetappBackupId pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
	// initiate the backup creation. Format:
	// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	SourceSnapshot pulumi.StringPtrInput
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''
	SourceVolume pulumi.StringPtrInput
	// The state of the Backup Vault. Possible Values : [STATE_UNSPECIFIED, CREATING, UPLOADING, READY, DELETING, ERROR,
	// UPDATING]
	State pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        NetappBackupTimeoutsPtrInput
	// Name of the backup vault to store the backup in.
	VaultName pulumi.StringPtrInput
	// Size of the file system when the backup was created. When creating a new volume from the backup, the volume capacity
	// will have to be at least as big.
	VolumeUsageBytes pulumi.StringPtrInput
}

func (NetappBackupState) ElementType() reflect.Type {
	return reflect.TypeOf((*netappBackupState)(nil)).Elem()
}

type netappBackupArgs struct {
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description *string `pulumi:"description"`
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Location of the backup.
	Location string `pulumi:"location"`
	// The resource name of the backup. Needs to be unique per location.
	Name           *string `pulumi:"name"`
	NetappBackupId *string `pulumi:"netappBackupId"`
	Project        *string `pulumi:"project"`
	// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
	// initiate the backup creation. Format:
	// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	SourceSnapshot *string `pulumi:"sourceSnapshot"`
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''
	SourceVolume *string               `pulumi:"sourceVolume"`
	Timeouts     *NetappBackupTimeouts `pulumi:"timeouts"`
	// Name of the backup vault to store the backup in.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a NetappBackup resource.
type NetappBackupArgs struct {
	// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
	Description pulumi.StringPtrInput
	// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
	// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
	// 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Location of the backup.
	Location pulumi.StringInput
	// The resource name of the backup. Needs to be unique per location.
	Name           pulumi.StringPtrInput
	NetappBackupId pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
	// initiate the backup creation. Format:
	// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}''
	SourceSnapshot pulumi.StringPtrInput
	// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}''
	SourceVolume pulumi.StringPtrInput
	Timeouts     NetappBackupTimeoutsPtrInput
	// Name of the backup vault to store the backup in.
	VaultName pulumi.StringInput
}

func (NetappBackupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*netappBackupArgs)(nil)).Elem()
}

type NetappBackupInput interface {
	pulumi.Input

	ToNetappBackupOutput() NetappBackupOutput
	ToNetappBackupOutputWithContext(ctx context.Context) NetappBackupOutput
}

func (*NetappBackup) ElementType() reflect.Type {
	return reflect.TypeOf((**NetappBackup)(nil)).Elem()
}

func (i *NetappBackup) ToNetappBackupOutput() NetappBackupOutput {
	return i.ToNetappBackupOutputWithContext(context.Background())
}

func (i *NetappBackup) ToNetappBackupOutputWithContext(ctx context.Context) NetappBackupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetappBackupOutput)
}

type NetappBackupOutput struct{ *pulumi.OutputState }

func (NetappBackupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetappBackup)(nil)).Elem()
}

func (o NetappBackupOutput) ToNetappBackupOutput() NetappBackupOutput {
	return o
}

func (o NetappBackupOutput) ToNetappBackupOutputWithContext(ctx context.Context) NetappBackupOutput {
	return o
}

// Type of backup, manually created or created by a backup policy. Possible Values : [TYPE_UNSPECIFIED, MANUAL, SCHEDULED]
func (o NetappBackupOutput) BackupType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.BackupType }).(pulumi.StringOutput)
}

// Backups of a volume build incrementally on top of each other. They form a "backup chain". Total size of all backups in a
// chain in bytes = baseline backup size + sum(incremental backup size)
func (o NetappBackupOutput) ChainStorageBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.ChainStorageBytes }).(pulumi.StringOutput)
}

// Create time of the backup. A timestamp in RFC3339 UTC "Zulu" format. Examples: "2023-06-22T09:13:01.617Z".
func (o NetappBackupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
func (o NetappBackupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetappBackupOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Labels as key value pairs. Example: '{ "owner": "Bob", "department": "finance", "purpose": "testing" }'. **Note**: This
// field is non-authoritative, and will only manage the labels present in your configuration. Please refer to the field
// 'effective_labels' for all of the labels present on the resource.
func (o NetappBackupOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location of the backup.
func (o NetappBackupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the backup. Needs to be unique per location.
func (o NetappBackupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NetappBackupOutput) NetappBackupId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.NetappBackupId }).(pulumi.StringOutput)
}

func (o NetappBackupOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// If specified, backup will be created from the given snapshot. If not specified, there will be a new snapshot taken to
// initiate the backup creation. Format:
// 'projects/{{projectId}}/locations/{{location}}/volumes/{{volumename}}/snapshots/{{snapshotname}}”
func (o NetappBackupOutput) SourceSnapshot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringPtrOutput { return v.SourceSnapshot }).(pulumi.StringPtrOutput)
}

// ID of volumes this backup belongs to. Format: 'projects/{{projects_id}}/locations/{{location}}/volumes/{{name}}”
func (o NetappBackupOutput) SourceVolume() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringPtrOutput { return v.SourceVolume }).(pulumi.StringPtrOutput)
}

// The state of the Backup Vault. Possible Values : [STATE_UNSPECIFIED, CREATING, UPLOADING, READY, DELETING, ERROR,
// UPDATING]
func (o NetappBackupOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o NetappBackupOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o NetappBackupOutput) Timeouts() NetappBackupTimeoutsPtrOutput {
	return o.ApplyT(func(v *NetappBackup) NetappBackupTimeoutsPtrOutput { return v.Timeouts }).(NetappBackupTimeoutsPtrOutput)
}

// Name of the backup vault to store the backup in.
func (o NetappBackupOutput) VaultName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.VaultName }).(pulumi.StringOutput)
}

// Size of the file system when the backup was created. When creating a new volume from the backup, the volume capacity
// will have to be at least as big.
func (o NetappBackupOutput) VolumeUsageBytes() pulumi.StringOutput {
	return o.ApplyT(func(v *NetappBackup) pulumi.StringOutput { return v.VolumeUsageBytes }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetappBackupInput)(nil)).Elem(), &NetappBackup{})
	pulumi.RegisterOutputType(NetappBackupOutput{})
}
