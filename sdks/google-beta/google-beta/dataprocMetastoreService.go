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

type DataprocMetastoreService struct {
	pulumi.CustomResourceState

	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri pulumi.StringOutput `pulumi:"artifactGcsUri"`
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
	// "SPANNER"]
	DatabaseType               pulumi.StringPtrOutput `pulumi:"databaseType"`
	DataprocMetastoreServiceId pulumi.StringOutput    `pulumi:"dataprocMetastoreServiceId"`
	// Indicates if the dataproc metastore should be protected against accidental deletions.
	DeletionProtection pulumi.BoolPtrOutput   `pulumi:"deletionProtection"`
	EffectiveLabels    pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
	EncryptionConfig DataprocMetastoreServiceEncryptionConfigPtrOutput `pulumi:"encryptionConfig"`
	// The URI of the endpoint used to access the metastore service.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig DataprocMetastoreServiceHiveMetastoreConfigPtrOutput `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the metastore service should reside. The default value is 'global'.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
	// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
	MaintenanceWindow DataprocMetastoreServiceMaintenanceWindowPtrOutput `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration DataprocMetastoreServiceMetadataIntegrationPtrOutput `pulumi:"metadataIntegration"`
	// The relative resource name of the metastore service.
	Name pulumi.StringOutput `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
	// form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringOutput `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig DataprocMetastoreServiceNetworkConfigPtrOutput `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.Float64Output `pulumi:"port"`
	Project pulumi.StringOutput  `pulumi:"project"`
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
	// ["CANARY", "STABLE"]
	ReleaseChannel pulumi.StringPtrOutput `pulumi:"releaseChannel"`
	// Represents the scaling configuration of a metastore service.
	ScalingConfig DataprocMetastoreServiceScalingConfigPtrOutput `pulumi:"scalingConfig"`
	// The configuration of scheduled backup for the metastore service.
	ScheduledBackup DataprocMetastoreServiceScheduledBackupPtrOutput `pulumi:"scheduledBackup"`
	// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The current state of the metastore service.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of the metastore service, if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig DataprocMetastoreServiceTelemetryConfigPtrOutput `pulumi:"telemetryConfig"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput `pulumi:"terraformLabels"`
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
	Tier     pulumi.StringOutput                       `pulumi:"tier"`
	Timeouts DataprocMetastoreServiceTimeoutsPtrOutput `pulumi:"timeouts"`
	// The globally unique resource identifier of the metastore service.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewDataprocMetastoreService registers a new resource with the given unique name, arguments, and options.
func NewDataprocMetastoreService(ctx *pulumi.Context,
	name string, args *DataprocMetastoreServiceArgs, opts ...pulumi.ResourceOption) (*DataprocMetastoreService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocMetastoreService
	err = ctx.RegisterPackageResource("google-beta:index/dataprocMetastoreService:DataprocMetastoreService", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocMetastoreService gets an existing DataprocMetastoreService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocMetastoreService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocMetastoreServiceState, opts ...pulumi.ResourceOption) (*DataprocMetastoreService, error) {
	var resource DataprocMetastoreService
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataprocMetastoreService:DataprocMetastoreService", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocMetastoreService resources.
type dataprocMetastoreServiceState struct {
	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri *string `pulumi:"artifactGcsUri"`
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
	// "SPANNER"]
	DatabaseType               *string `pulumi:"databaseType"`
	DataprocMetastoreServiceId *string `pulumi:"dataprocMetastoreServiceId"`
	// Indicates if the dataproc metastore should be protected against accidental deletions.
	DeletionProtection *bool             `pulumi:"deletionProtection"`
	EffectiveLabels    map[string]string `pulumi:"effectiveLabels"`
	// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
	EncryptionConfig *DataprocMetastoreServiceEncryptionConfig `pulumi:"encryptionConfig"`
	// The URI of the endpoint used to access the metastore service.
	EndpointUri *string `pulumi:"endpointUri"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig *DataprocMetastoreServiceHiveMetastoreConfig `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore service should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
	// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
	MaintenanceWindow *DataprocMetastoreServiceMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration *DataprocMetastoreServiceMetadataIntegration `pulumi:"metadataIntegration"`
	// The relative resource name of the metastore service.
	Name *string `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
	// form: "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig *DataprocMetastoreServiceNetworkConfig `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    *float64 `pulumi:"port"`
	Project *string  `pulumi:"project"`
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
	// ["CANARY", "STABLE"]
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Represents the scaling configuration of a metastore service.
	ScalingConfig *DataprocMetastoreServiceScalingConfig `pulumi:"scalingConfig"`
	// The configuration of scheduled backup for the metastore service.
	ScheduledBackup *DataprocMetastoreServiceScheduledBackup `pulumi:"scheduledBackup"`
	// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	ServiceId *string `pulumi:"serviceId"`
	// The current state of the metastore service.
	State *string `pulumi:"state"`
	// Additional information about the current state of the metastore service, if available.
	StateMessage *string `pulumi:"stateMessage"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig *DataprocMetastoreServiceTelemetryConfig `pulumi:"telemetryConfig"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string `pulumi:"terraformLabels"`
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
	Tier     *string                           `pulumi:"tier"`
	Timeouts *DataprocMetastoreServiceTimeouts `pulumi:"timeouts"`
	// The globally unique resource identifier of the metastore service.
	Uid *string `pulumi:"uid"`
}

type DataprocMetastoreServiceState struct {
	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri pulumi.StringPtrInput
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
	// "SPANNER"]
	DatabaseType               pulumi.StringPtrInput
	DataprocMetastoreServiceId pulumi.StringPtrInput
	// Indicates if the dataproc metastore should be protected against accidental deletions.
	DeletionProtection pulumi.BoolPtrInput
	EffectiveLabels    pulumi.StringMapInput
	// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
	EncryptionConfig DataprocMetastoreServiceEncryptionConfigPtrInput
	// The URI of the endpoint used to access the metastore service.
	EndpointUri pulumi.StringPtrInput
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig DataprocMetastoreServiceHiveMetastoreConfigPtrInput
	// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the metastore service should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
	// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
	MaintenanceWindow DataprocMetastoreServiceMaintenanceWindowPtrInput
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration DataprocMetastoreServiceMetadataIntegrationPtrInput
	// The relative resource name of the metastore service.
	Name pulumi.StringPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
	// form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringPtrInput
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig DataprocMetastoreServiceNetworkConfigPtrInput
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.Float64PtrInput
	Project pulumi.StringPtrInput
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
	// ["CANARY", "STABLE"]
	ReleaseChannel pulumi.StringPtrInput
	// Represents the scaling configuration of a metastore service.
	ScalingConfig DataprocMetastoreServiceScalingConfigPtrInput
	// The configuration of scheduled backup for the metastore service.
	ScheduledBackup DataprocMetastoreServiceScheduledBackupPtrInput
	// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	ServiceId pulumi.StringPtrInput
	// The current state of the metastore service.
	State pulumi.StringPtrInput
	// Additional information about the current state of the metastore service, if available.
	StateMessage pulumi.StringPtrInput
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig DataprocMetastoreServiceTelemetryConfigPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
	Tier     pulumi.StringPtrInput
	Timeouts DataprocMetastoreServiceTimeoutsPtrInput
	// The globally unique resource identifier of the metastore service.
	Uid pulumi.StringPtrInput
}

func (DataprocMetastoreServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocMetastoreServiceState)(nil)).Elem()
}

type dataprocMetastoreServiceArgs struct {
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
	// "SPANNER"]
	DatabaseType               *string `pulumi:"databaseType"`
	DataprocMetastoreServiceId *string `pulumi:"dataprocMetastoreServiceId"`
	// Indicates if the dataproc metastore should be protected against accidental deletions.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
	EncryptionConfig *DataprocMetastoreServiceEncryptionConfig `pulumi:"encryptionConfig"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig *DataprocMetastoreServiceHiveMetastoreConfig `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// The location where the metastore service should reside. The default value is 'global'.
	Location *string `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
	// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
	MaintenanceWindow *DataprocMetastoreServiceMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration *DataprocMetastoreServiceMetadataIntegration `pulumi:"metadataIntegration"`
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
	// form: "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig *DataprocMetastoreServiceNetworkConfig `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    *float64 `pulumi:"port"`
	Project *string  `pulumi:"project"`
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
	// ["CANARY", "STABLE"]
	ReleaseChannel *string `pulumi:"releaseChannel"`
	// Represents the scaling configuration of a metastore service.
	ScalingConfig *DataprocMetastoreServiceScalingConfig `pulumi:"scalingConfig"`
	// The configuration of scheduled backup for the metastore service.
	ScheduledBackup *DataprocMetastoreServiceScheduledBackup `pulumi:"scheduledBackup"`
	// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	ServiceId string `pulumi:"serviceId"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig *DataprocMetastoreServiceTelemetryConfig `pulumi:"telemetryConfig"`
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
	Tier     *string                           `pulumi:"tier"`
	Timeouts *DataprocMetastoreServiceTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DataprocMetastoreService resource.
type DataprocMetastoreServiceArgs struct {
	// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
	// "SPANNER"]
	DatabaseType               pulumi.StringPtrInput
	DataprocMetastoreServiceId pulumi.StringPtrInput
	// Indicates if the dataproc metastore should be protected against accidental deletions.
	DeletionProtection pulumi.BoolPtrInput
	// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
	EncryptionConfig DataprocMetastoreServiceEncryptionConfigPtrInput
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig DataprocMetastoreServiceHiveMetastoreConfigPtrInput
	// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// The location where the metastore service should reside. The default value is 'global'.
	Location pulumi.StringPtrInput
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
	// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
	MaintenanceWindow DataprocMetastoreServiceMaintenanceWindowPtrInput
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration DataprocMetastoreServiceMetadataIntegrationPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
	// form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringPtrInput
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig DataprocMetastoreServiceNetworkConfigPtrInput
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.Float64PtrInput
	Project pulumi.StringPtrInput
	// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
	// ["CANARY", "STABLE"]
	ReleaseChannel pulumi.StringPtrInput
	// Represents the scaling configuration of a metastore service.
	ScalingConfig DataprocMetastoreServiceScalingConfigPtrInput
	// The configuration of scheduled backup for the metastore service.
	ScheduledBackup DataprocMetastoreServiceScheduledBackupPtrInput
	// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
	// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
	ServiceId pulumi.StringInput
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig DataprocMetastoreServiceTelemetryConfigPtrInput
	// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
	Tier     pulumi.StringPtrInput
	Timeouts DataprocMetastoreServiceTimeoutsPtrInput
}

func (DataprocMetastoreServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocMetastoreServiceArgs)(nil)).Elem()
}

type DataprocMetastoreServiceInput interface {
	pulumi.Input

	ToDataprocMetastoreServiceOutput() DataprocMetastoreServiceOutput
	ToDataprocMetastoreServiceOutputWithContext(ctx context.Context) DataprocMetastoreServiceOutput
}

func (*DataprocMetastoreService) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocMetastoreService)(nil)).Elem()
}

func (i *DataprocMetastoreService) ToDataprocMetastoreServiceOutput() DataprocMetastoreServiceOutput {
	return i.ToDataprocMetastoreServiceOutputWithContext(context.Background())
}

func (i *DataprocMetastoreService) ToDataprocMetastoreServiceOutputWithContext(ctx context.Context) DataprocMetastoreServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocMetastoreServiceOutput)
}

type DataprocMetastoreServiceOutput struct{ *pulumi.OutputState }

func (DataprocMetastoreServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocMetastoreService)(nil)).Elem()
}

func (o DataprocMetastoreServiceOutput) ToDataprocMetastoreServiceOutput() DataprocMetastoreServiceOutput {
	return o
}

func (o DataprocMetastoreServiceOutput) ToDataprocMetastoreServiceOutputWithContext(ctx context.Context) DataprocMetastoreServiceOutput {
	return o
}

// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
func (o DataprocMetastoreServiceOutput) ArtifactGcsUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.ArtifactGcsUri }).(pulumi.StringOutput)
}

// The database type that the Metastore service stores its data. Default value: "MYSQL" Possible values: ["MYSQL",
// "SPANNER"]
func (o DataprocMetastoreServiceOutput) DatabaseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringPtrOutput { return v.DatabaseType }).(pulumi.StringPtrOutput)
}

func (o DataprocMetastoreServiceOutput) DataprocMetastoreServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.DataprocMetastoreServiceId }).(pulumi.StringOutput)
}

// Indicates if the dataproc metastore should be protected against accidental deletions.
func (o DataprocMetastoreServiceOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o DataprocMetastoreServiceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Information used to configure the Dataproc Metastore service to encrypt customer data at rest.
func (o DataprocMetastoreServiceOutput) EncryptionConfig() DataprocMetastoreServiceEncryptionConfigPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceEncryptionConfigPtrOutput {
		return v.EncryptionConfig
	}).(DataprocMetastoreServiceEncryptionConfigPtrOutput)
}

// The URI of the endpoint used to access the metastore service.
func (o DataprocMetastoreServiceOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

// Configuration information specific to running Hive metastore software as the metastore service.
func (o DataprocMetastoreServiceOutput) HiveMetastoreConfig() DataprocMetastoreServiceHiveMetastoreConfigPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceHiveMetastoreConfigPtrOutput {
		return v.HiveMetastoreConfig
	}).(DataprocMetastoreServiceHiveMetastoreConfigPtrOutput)
}

// User-defined labels for the metastore service. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o DataprocMetastoreServiceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the metastore service should reside. The default value is 'global'.
func (o DataprocMetastoreServiceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for
// maintenance purposes in UTC time. Maintenance window is not needed for services with the 'SPANNER' database type.
func (o DataprocMetastoreServiceOutput) MaintenanceWindow() DataprocMetastoreServiceMaintenanceWindowPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceMaintenanceWindowPtrOutput {
		return v.MaintenanceWindow
	}).(DataprocMetastoreServiceMaintenanceWindowPtrOutput)
}

// The setting that defines how metastore metadata should be integrated with external services and systems.
func (o DataprocMetastoreServiceOutput) MetadataIntegration() DataprocMetastoreServiceMetadataIntegrationPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceMetadataIntegrationPtrOutput {
		return v.MetadataIntegration
	}).(DataprocMetastoreServiceMetadataIntegrationPtrOutput)
}

// The relative resource name of the metastore service.
func (o DataprocMetastoreServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following
// form: "projects/{projectNumber}/global/networks/{network_id}".
func (o DataprocMetastoreServiceOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The configuration specifying the network settings for the Dataproc Metastore service.
func (o DataprocMetastoreServiceOutput) NetworkConfig() DataprocMetastoreServiceNetworkConfigPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceNetworkConfigPtrOutput {
		return v.NetworkConfig
	}).(DataprocMetastoreServiceNetworkConfigPtrOutput)
}

// The TCP port at which the metastore service is reached. Default: 9083.
func (o DataprocMetastoreServiceOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.Float64Output { return v.Port }).(pulumi.Float64Output)
}

func (o DataprocMetastoreServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The release channel of the service. If unspecified, defaults to 'STABLE'. Default value: "STABLE" Possible values:
// ["CANARY", "STABLE"]
func (o DataprocMetastoreServiceOutput) ReleaseChannel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringPtrOutput { return v.ReleaseChannel }).(pulumi.StringPtrOutput)
}

// Represents the scaling configuration of a metastore service.
func (o DataprocMetastoreServiceOutput) ScalingConfig() DataprocMetastoreServiceScalingConfigPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceScalingConfigPtrOutput {
		return v.ScalingConfig
	}).(DataprocMetastoreServiceScalingConfigPtrOutput)
}

// The configuration of scheduled backup for the metastore service.
func (o DataprocMetastoreServiceOutput) ScheduledBackup() DataprocMetastoreServiceScheduledBackupPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceScheduledBackupPtrOutput {
		return v.ScheduledBackup
	}).(DataprocMetastoreServiceScheduledBackupPtrOutput)
}

// The ID of the metastore service. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and
// hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 63 characters.
func (o DataprocMetastoreServiceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The current state of the metastore service.
func (o DataprocMetastoreServiceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of the metastore service, if available.
func (o DataprocMetastoreServiceOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.StateMessage }).(pulumi.StringOutput)
}

// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
func (o DataprocMetastoreServiceOutput) TelemetryConfig() DataprocMetastoreServiceTelemetryConfigPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceTelemetryConfigPtrOutput {
		return v.TelemetryConfig
	}).(DataprocMetastoreServiceTelemetryConfigPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataprocMetastoreServiceOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

// The tier of the service. Possible values: ["DEVELOPER", "ENTERPRISE"]
func (o DataprocMetastoreServiceOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

func (o DataprocMetastoreServiceOutput) Timeouts() DataprocMetastoreServiceTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) DataprocMetastoreServiceTimeoutsPtrOutput { return v.Timeouts }).(DataprocMetastoreServiceTimeoutsPtrOutput)
}

// The globally unique resource identifier of the metastore service.
func (o DataprocMetastoreServiceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocMetastoreService) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocMetastoreServiceInput)(nil)).Elem(), &DataprocMetastoreService{})
	pulumi.RegisterOutputType(DataprocMetastoreServiceOutput{})
}
