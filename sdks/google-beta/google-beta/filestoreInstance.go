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

type FilestoreInstance struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Indicates whether the instance is protected against deletion.
	DeletionProtectionEnabled pulumi.BoolPtrOutput `pulumi:"deletionProtectionEnabled"`
	// The reason for enabling deletion protection.
	DeletionProtectionReason pulumi.StringPtrOutput `pulumi:"deletionProtectionReason"`
	// A description of the instance.
	Description     pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Output only fields for replication configuration.
	EffectiveReplications FilestoreInstanceEffectiveReplicationArrayOutput `pulumi:"effectiveReplications"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares          FilestoreInstanceFileSharesOutput `pulumi:"fileShares"`
	FilestoreInstanceId pulumi.StringOutput               `pulumi:"filestoreInstanceId"`
	// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
	// instance only, indicating the active as the peer_instance
	InitialReplication FilestoreInstanceInitialReplicationPtrOutput `pulumi:"initialReplication"`
	// KMS key name used for data encryption.
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the instance.
	Name pulumi.StringOutput `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks FilestoreInstanceNetworkArrayOutput `pulumi:"networks"`
	// Performance configuration for the instance. If not provided, the default performance settings will be used.
	PerformanceConfig FilestoreInstancePerformanceConfigPtrOutput `pulumi:"performanceConfig"`
	Project           pulumi.StringOutput                         `pulumi:"project"`
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
	// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
	// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
	// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
	// 'google_tags_tag_value' resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput `pulumi:"terraformLabels"`
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
	// ZONAL, REGIONAL and ENTERPRISE
	Tier     pulumi.StringOutput                `pulumi:"tier"`
	Timeouts FilestoreInstanceTimeoutsPtrOutput `pulumi:"timeouts"`
	// The name of the Filestore zone of the instance.
	//
	// Deprecated: Deprecated
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewFilestoreInstance registers a new resource with the given unique name, arguments, and options.
func NewFilestoreInstance(ctx *pulumi.Context,
	name string, args *FilestoreInstanceArgs, opts ...pulumi.ResourceOption) (*FilestoreInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FileShares == nil {
		return nil, errors.New("invalid value for required argument 'FileShares'")
	}
	if args.Networks == nil {
		return nil, errors.New("invalid value for required argument 'Networks'")
	}
	if args.Tier == nil {
		return nil, errors.New("invalid value for required argument 'Tier'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FilestoreInstance
	err = ctx.RegisterPackageResource("google-beta:index/filestoreInstance:FilestoreInstance", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFilestoreInstance gets an existing FilestoreInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFilestoreInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FilestoreInstanceState, opts ...pulumi.ResourceOption) (*FilestoreInstance, error) {
	var resource FilestoreInstance
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/filestoreInstance:FilestoreInstance", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FilestoreInstance resources.
type filestoreInstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// Indicates whether the instance is protected against deletion.
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason for enabling deletion protection.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// A description of the instance.
	Description     *string           `pulumi:"description"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Output only fields for replication configuration.
	EffectiveReplications []FilestoreInstanceEffectiveReplication `pulumi:"effectiveReplications"`
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag *string `pulumi:"etag"`
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares          *FilestoreInstanceFileShares `pulumi:"fileShares"`
	FilestoreInstanceId *string                      `pulumi:"filestoreInstanceId"`
	// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
	// instance only, indicating the active as the peer_instance
	InitialReplication *FilestoreInstanceInitialReplication `pulumi:"initialReplication"`
	// KMS key name used for data encryption.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location *string `pulumi:"location"`
	// The resource name of the instance.
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks []FilestoreInstanceNetwork `pulumi:"networks"`
	// Performance configuration for the instance. If not provided, the default performance settings will be used.
	PerformanceConfig *FilestoreInstancePerformanceConfig `pulumi:"performanceConfig"`
	Project           *string                             `pulumi:"project"`
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
	// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
	// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
	Protocol *string `pulumi:"protocol"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
	// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
	// 'google_tags_tag_value' resource.
	Tags map[string]string `pulumi:"tags"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
	// ZONAL, REGIONAL and ENTERPRISE
	Tier     *string                    `pulumi:"tier"`
	Timeouts *FilestoreInstanceTimeouts `pulumi:"timeouts"`
	// The name of the Filestore zone of the instance.
	//
	// Deprecated: Deprecated
	Zone *string `pulumi:"zone"`
}

type FilestoreInstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// Indicates whether the instance is protected against deletion.
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason for enabling deletion protection.
	DeletionProtectionReason pulumi.StringPtrInput
	// A description of the instance.
	Description     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Output only fields for replication configuration.
	EffectiveReplications FilestoreInstanceEffectiveReplicationArrayInput
	// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
	Etag pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares          FilestoreInstanceFileSharesPtrInput
	FilestoreInstanceId pulumi.StringPtrInput
	// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
	// instance only, indicating the active as the peer_instance
	InitialReplication FilestoreInstanceInitialReplicationPtrInput
	// KMS key name used for data encryption.
	KmsKeyName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringPtrInput
	// The resource name of the instance.
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks FilestoreInstanceNetworkArrayInput
	// Performance configuration for the instance. If not provided, the default performance settings will be used.
	PerformanceConfig FilestoreInstancePerformanceConfigPtrInput
	Project           pulumi.StringPtrInput
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
	// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
	// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
	Protocol pulumi.StringPtrInput
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
	// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
	// 'google_tags_tag_value' resource.
	Tags pulumi.StringMapInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
	// ZONAL, REGIONAL and ENTERPRISE
	Tier     pulumi.StringPtrInput
	Timeouts FilestoreInstanceTimeoutsPtrInput
	// The name of the Filestore zone of the instance.
	//
	// Deprecated: Deprecated
	Zone pulumi.StringPtrInput
}

func (FilestoreInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*filestoreInstanceState)(nil)).Elem()
}

type filestoreInstanceArgs struct {
	// Indicates whether the instance is protected against deletion.
	DeletionProtectionEnabled *bool `pulumi:"deletionProtectionEnabled"`
	// The reason for enabling deletion protection.
	DeletionProtectionReason *string `pulumi:"deletionProtectionReason"`
	// A description of the instance.
	Description *string `pulumi:"description"`
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares          FilestoreInstanceFileShares `pulumi:"fileShares"`
	FilestoreInstanceId *string                     `pulumi:"filestoreInstanceId"`
	// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
	// instance only, indicating the active as the peer_instance
	InitialReplication *FilestoreInstanceInitialReplication `pulumi:"initialReplication"`
	// KMS key name used for data encryption.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location *string `pulumi:"location"`
	// The resource name of the instance.
	Name *string `pulumi:"name"`
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks []FilestoreInstanceNetwork `pulumi:"networks"`
	// Performance configuration for the instance. If not provided, the default performance settings will be used.
	PerformanceConfig *FilestoreInstancePerformanceConfig `pulumi:"performanceConfig"`
	Project           *string                             `pulumi:"project"`
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
	// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
	// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
	Protocol *string `pulumi:"protocol"`
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
	// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
	// 'google_tags_tag_value' resource.
	Tags map[string]string `pulumi:"tags"`
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
	// ZONAL, REGIONAL and ENTERPRISE
	Tier     string                     `pulumi:"tier"`
	Timeouts *FilestoreInstanceTimeouts `pulumi:"timeouts"`
	// The name of the Filestore zone of the instance.
	//
	// Deprecated: Deprecated
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a FilestoreInstance resource.
type FilestoreInstanceArgs struct {
	// Indicates whether the instance is protected against deletion.
	DeletionProtectionEnabled pulumi.BoolPtrInput
	// The reason for enabling deletion protection.
	DeletionProtectionReason pulumi.StringPtrInput
	// A description of the instance.
	Description pulumi.StringPtrInput
	// File system shares on the instance. For this version, only a single file share is supported.
	FileShares          FilestoreInstanceFileSharesInput
	FilestoreInstanceId pulumi.StringPtrInput
	// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
	// instance only, indicating the active as the peer_instance
	InitialReplication FilestoreInstanceInitialReplicationPtrInput
	// KMS key name used for data encryption.
	KmsKeyName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
	Location pulumi.StringPtrInput
	// The resource name of the instance.
	Name pulumi.StringPtrInput
	// VPC networks to which the instance is connected. For this version, only a single network is supported.
	Networks FilestoreInstanceNetworkArrayInput
	// Performance configuration for the instance. If not provided, the default performance settings will be used.
	PerformanceConfig FilestoreInstancePerformanceConfigPtrInput
	Project           pulumi.StringPtrInput
	// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
	// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
	// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
	Protocol pulumi.StringPtrInput
	// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
	// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
	// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
	// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
	// 'google_tags_tag_value' resource.
	Tags pulumi.StringMapInput
	// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
	// ZONAL, REGIONAL and ENTERPRISE
	Tier     pulumi.StringInput
	Timeouts FilestoreInstanceTimeoutsPtrInput
	// The name of the Filestore zone of the instance.
	//
	// Deprecated: Deprecated
	Zone pulumi.StringPtrInput
}

func (FilestoreInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*filestoreInstanceArgs)(nil)).Elem()
}

type FilestoreInstanceInput interface {
	pulumi.Input

	ToFilestoreInstanceOutput() FilestoreInstanceOutput
	ToFilestoreInstanceOutputWithContext(ctx context.Context) FilestoreInstanceOutput
}

func (*FilestoreInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**FilestoreInstance)(nil)).Elem()
}

func (i *FilestoreInstance) ToFilestoreInstanceOutput() FilestoreInstanceOutput {
	return i.ToFilestoreInstanceOutputWithContext(context.Background())
}

func (i *FilestoreInstance) ToFilestoreInstanceOutputWithContext(ctx context.Context) FilestoreInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FilestoreInstanceOutput)
}

type FilestoreInstanceOutput struct{ *pulumi.OutputState }

func (FilestoreInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FilestoreInstance)(nil)).Elem()
}

func (o FilestoreInstanceOutput) ToFilestoreInstanceOutput() FilestoreInstanceOutput {
	return o
}

func (o FilestoreInstanceOutput) ToFilestoreInstanceOutputWithContext(ctx context.Context) FilestoreInstanceOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o FilestoreInstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Indicates whether the instance is protected against deletion.
func (o FilestoreInstanceOutput) DeletionProtectionEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.BoolPtrOutput { return v.DeletionProtectionEnabled }).(pulumi.BoolPtrOutput)
}

// The reason for enabling deletion protection.
func (o FilestoreInstanceOutput) DeletionProtectionReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringPtrOutput { return v.DeletionProtectionReason }).(pulumi.StringPtrOutput)
}

// A description of the instance.
func (o FilestoreInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o FilestoreInstanceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Output only fields for replication configuration.
func (o FilestoreInstanceOutput) EffectiveReplications() FilestoreInstanceEffectiveReplicationArrayOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstanceEffectiveReplicationArrayOutput {
		return v.EffectiveReplications
	}).(FilestoreInstanceEffectiveReplicationArrayOutput)
}

// Server-specified ETag for the instance resource to prevent simultaneous updates from overwriting each other.
func (o FilestoreInstanceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// File system shares on the instance. For this version, only a single file share is supported.
func (o FilestoreInstanceOutput) FileShares() FilestoreInstanceFileSharesOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstanceFileSharesOutput { return v.FileShares }).(FilestoreInstanceFileSharesOutput)
}

func (o FilestoreInstanceOutput) FilestoreInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.FilestoreInstanceId }).(pulumi.StringOutput)
}

// Replication configuration, once set, this cannot be updated. Addtionally this should be specified on the replica
// instance only, indicating the active as the peer_instance
func (o FilestoreInstanceOutput) InitialReplication() FilestoreInstanceInitialReplicationPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstanceInitialReplicationPtrOutput { return v.InitialReplication }).(FilestoreInstanceInitialReplicationPtrOutput)
}

// KMS key name used for data encryption.
func (o FilestoreInstanceOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringPtrOutput { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o FilestoreInstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
func (o FilestoreInstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the instance.
func (o FilestoreInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// VPC networks to which the instance is connected. For this version, only a single network is supported.
func (o FilestoreInstanceOutput) Networks() FilestoreInstanceNetworkArrayOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstanceNetworkArrayOutput { return v.Networks }).(FilestoreInstanceNetworkArrayOutput)
}

// Performance configuration for the instance. If not provided, the default performance settings will be used.
func (o FilestoreInstanceOutput) PerformanceConfig() FilestoreInstancePerformanceConfigPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstancePerformanceConfigPtrOutput { return v.PerformanceConfig }).(FilestoreInstancePerformanceConfigPtrOutput)
}

func (o FilestoreInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Either NFSv3, for using NFS version 3 as file sharing protocol, or NFSv4.1, for using NFS version 4.1 as file sharing
// protocol. NFSv4.1 can be used with HIGH_SCALE_SSD, ZONAL, REGIONAL and ENTERPRISE. The default is NFSv3. Default value:
// "NFS_V3" Possible values: ["NFS_V3", "NFS_V4_1"]
func (o FilestoreInstanceOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

// A map of resource manager tags. Resource manager tag keys and values have the same definition as resource manager tags.
// Keys must be in the format tagKeys/{tag_key_id}, and values are in the format tagValues/456. The field is ignored when
// empty. The field is immutable and causes resource replacement when mutated. This field is only set at create time and
// modifying this field after creation will trigger recreation. To apply tags to an existing resource, see the
// 'google_tags_tag_value' resource.
func (o FilestoreInstanceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o FilestoreInstanceOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

// The service tier of the instance. Possible values include: STANDARD, PREMIUM, BASIC_HDD, BASIC_SSD, HIGH_SCALE_SSD,
// ZONAL, REGIONAL and ENTERPRISE
func (o FilestoreInstanceOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

func (o FilestoreInstanceOutput) Timeouts() FilestoreInstanceTimeoutsPtrOutput {
	return o.ApplyT(func(v *FilestoreInstance) FilestoreInstanceTimeoutsPtrOutput { return v.Timeouts }).(FilestoreInstanceTimeoutsPtrOutput)
}

// The name of the Filestore zone of the instance.
//
// Deprecated: Deprecated
func (o FilestoreInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *FilestoreInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FilestoreInstanceInput)(nil)).Elem(), &FilestoreInstance{})
	pulumi.RegisterOutputType(FilestoreInstanceOutput{})
}
