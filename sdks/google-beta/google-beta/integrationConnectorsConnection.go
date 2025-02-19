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

type IntegrationConnectorsConnection struct {
	pulumi.CustomResourceState

	// authConfig for the connection.
	AuthConfig IntegrationConnectorsConnectionAuthConfigPtrOutput `pulumi:"authConfig"`
	// Config Variables for the connection.
	ConfigVariables IntegrationConnectorsConnectionConfigVariableArrayOutput `pulumi:"configVariables"`
	// Connection revision. This field is only updated when the connection is created or updated by User.
	ConnectionRevision pulumi.StringOutput `pulumi:"connectionRevision"`
	// connectorVersion of the Connector.
	ConnectorVersion pulumi.StringOutput `pulumi:"connectorVersion"`
	// This configuration provides infra configs like rate limit threshold which need to be configurable for every connector
	// version.
	ConnectorVersionInfraConfigs IntegrationConnectorsConnectionConnectorVersionInfraConfigArrayOutput `pulumi:"connectorVersionInfraConfigs"`
	// Flag to mark the version indicating the launch stage.
	ConnectorVersionLaunchStage pulumi.StringOutput `pulumi:"connectorVersionLaunchStage"`
	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// An arbitrary description for the Connection.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Define the Connectors target endpoint.
	DestinationConfigs IntegrationConnectorsConnectionDestinationConfigArrayOutput `pulumi:"destinationConfigs"`
	EffectiveLabels    pulumi.StringMapOutput                                      `pulumi:"effectiveLabels"`
	// Eventing Configuration of a connection
	EventingConfig IntegrationConnectorsConnectionEventingConfigPtrOutput `pulumi:"eventingConfig"`
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
	// "ONLY_EVENTING"]
	EventingEnablementType pulumi.StringPtrOutput `pulumi:"eventingEnablementType"`
	// Eventing Runtime Data.
	EventingRuntimeDatas              IntegrationConnectorsConnectionEventingRuntimeDataArrayOutput `pulumi:"eventingRuntimeDatas"`
	IntegrationConnectorsConnectionId pulumi.StringOutput                                           `pulumi:"integrationConnectorsConnectionId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Location in which Connection needs to be created.
	Location pulumi.StringOutput `pulumi:"location"`
	// Determines whether or no a connection is locked. If locked, a reason must be specified.
	LockConfig IntegrationConnectorsConnectionLockConfigPtrOutput `pulumi:"lockConfig"`
	// Log configuration for the connection.
	LogConfig IntegrationConnectorsConnectionLogConfigPtrOutput `pulumi:"logConfig"`
	// Name of Connection needs to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// Node configuration for the connection.
	NodeConfig IntegrationConnectorsConnectionNodeConfigPtrOutput `pulumi:"nodeConfig"`
	Project    pulumi.StringOutput                                `pulumi:"project"`
	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g.
	// "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	ServiceDirectory pulumi.StringOutput `pulumi:"serviceDirectory"`
	// SSL Configuration of a connection
	SslConfig IntegrationConnectorsConnectionSslConfigPtrOutput `pulumi:"sslConfig"`
	// Status of the Integration Connector.
	Statuses IntegrationConnectorsConnectionStatusArrayOutput `pulumi:"statuses"`
	// This subscription type enum states the subscription type of the project.
	SubscriptionType pulumi.StringOutput `pulumi:"subscriptionType"`
	// Suspended indicates if a user has suspended a connection or not.
	Suspended pulumi.BoolPtrOutput `pulumi:"suspended"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                           `pulumi:"terraformLabels"`
	Timeouts        IntegrationConnectorsConnectionTimeoutsPtrOutput `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewIntegrationConnectorsConnection registers a new resource with the given unique name, arguments, and options.
func NewIntegrationConnectorsConnection(ctx *pulumi.Context,
	name string, args *IntegrationConnectorsConnectionArgs, opts ...pulumi.ResourceOption) (*IntegrationConnectorsConnection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectorVersion == nil {
		return nil, errors.New("invalid value for required argument 'ConnectorVersion'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource IntegrationConnectorsConnection
	err = ctx.RegisterPackageResource("google-beta:index/integrationConnectorsConnection:IntegrationConnectorsConnection", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationConnectorsConnection gets an existing IntegrationConnectorsConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationConnectorsConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationConnectorsConnectionState, opts ...pulumi.ResourceOption) (*IntegrationConnectorsConnection, error) {
	var resource IntegrationConnectorsConnection
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/integrationConnectorsConnection:IntegrationConnectorsConnection", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationConnectorsConnection resources.
type integrationConnectorsConnectionState struct {
	// authConfig for the connection.
	AuthConfig *IntegrationConnectorsConnectionAuthConfig `pulumi:"authConfig"`
	// Config Variables for the connection.
	ConfigVariables []IntegrationConnectorsConnectionConfigVariable `pulumi:"configVariables"`
	// Connection revision. This field is only updated when the connection is created or updated by User.
	ConnectionRevision *string `pulumi:"connectionRevision"`
	// connectorVersion of the Connector.
	ConnectorVersion *string `pulumi:"connectorVersion"`
	// This configuration provides infra configs like rate limit threshold which need to be configurable for every connector
	// version.
	ConnectorVersionInfraConfigs []IntegrationConnectorsConnectionConnectorVersionInfraConfig `pulumi:"connectorVersionInfraConfigs"`
	// Flag to mark the version indicating the launch stage.
	ConnectorVersionLaunchStage *string `pulumi:"connectorVersionLaunchStage"`
	// Time the Namespace was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// An arbitrary description for the Connection.
	Description *string `pulumi:"description"`
	// Define the Connectors target endpoint.
	DestinationConfigs []IntegrationConnectorsConnectionDestinationConfig `pulumi:"destinationConfigs"`
	EffectiveLabels    map[string]string                                  `pulumi:"effectiveLabels"`
	// Eventing Configuration of a connection
	EventingConfig *IntegrationConnectorsConnectionEventingConfig `pulumi:"eventingConfig"`
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
	// "ONLY_EVENTING"]
	EventingEnablementType *string `pulumi:"eventingEnablementType"`
	// Eventing Runtime Data.
	EventingRuntimeDatas              []IntegrationConnectorsConnectionEventingRuntimeData `pulumi:"eventingRuntimeDatas"`
	IntegrationConnectorsConnectionId *string                                              `pulumi:"integrationConnectorsConnectionId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Location in which Connection needs to be created.
	Location *string `pulumi:"location"`
	// Determines whether or no a connection is locked. If locked, a reason must be specified.
	LockConfig *IntegrationConnectorsConnectionLockConfig `pulumi:"lockConfig"`
	// Log configuration for the connection.
	LogConfig *IntegrationConnectorsConnectionLogConfig `pulumi:"logConfig"`
	// Name of Connection needs to be created.
	Name *string `pulumi:"name"`
	// Node configuration for the connection.
	NodeConfig *IntegrationConnectorsConnectionNodeConfig `pulumi:"nodeConfig"`
	Project    *string                                    `pulumi:"project"`
	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g.
	// "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	ServiceDirectory *string `pulumi:"serviceDirectory"`
	// SSL Configuration of a connection
	SslConfig *IntegrationConnectorsConnectionSslConfig `pulumi:"sslConfig"`
	// Status of the Integration Connector.
	Statuses []IntegrationConnectorsConnectionStatus `pulumi:"statuses"`
	// This subscription type enum states the subscription type of the project.
	SubscriptionType *string `pulumi:"subscriptionType"`
	// Suspended indicates if a user has suspended a connection or not.
	Suspended *bool `pulumi:"suspended"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                        `pulumi:"terraformLabels"`
	Timeouts        *IntegrationConnectorsConnectionTimeouts `pulumi:"timeouts"`
	// Time the Namespace was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type IntegrationConnectorsConnectionState struct {
	// authConfig for the connection.
	AuthConfig IntegrationConnectorsConnectionAuthConfigPtrInput
	// Config Variables for the connection.
	ConfigVariables IntegrationConnectorsConnectionConfigVariableArrayInput
	// Connection revision. This field is only updated when the connection is created or updated by User.
	ConnectionRevision pulumi.StringPtrInput
	// connectorVersion of the Connector.
	ConnectorVersion pulumi.StringPtrInput
	// This configuration provides infra configs like rate limit threshold which need to be configurable for every connector
	// version.
	ConnectorVersionInfraConfigs IntegrationConnectorsConnectionConnectorVersionInfraConfigArrayInput
	// Flag to mark the version indicating the launch stage.
	ConnectorVersionLaunchStage pulumi.StringPtrInput
	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringPtrInput
	// An arbitrary description for the Connection.
	Description pulumi.StringPtrInput
	// Define the Connectors target endpoint.
	DestinationConfigs IntegrationConnectorsConnectionDestinationConfigArrayInput
	EffectiveLabels    pulumi.StringMapInput
	// Eventing Configuration of a connection
	EventingConfig IntegrationConnectorsConnectionEventingConfigPtrInput
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
	// "ONLY_EVENTING"]
	EventingEnablementType pulumi.StringPtrInput
	// Eventing Runtime Data.
	EventingRuntimeDatas              IntegrationConnectorsConnectionEventingRuntimeDataArrayInput
	IntegrationConnectorsConnectionId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Location in which Connection needs to be created.
	Location pulumi.StringPtrInput
	// Determines whether or no a connection is locked. If locked, a reason must be specified.
	LockConfig IntegrationConnectorsConnectionLockConfigPtrInput
	// Log configuration for the connection.
	LogConfig IntegrationConnectorsConnectionLogConfigPtrInput
	// Name of Connection needs to be created.
	Name pulumi.StringPtrInput
	// Node configuration for the connection.
	NodeConfig IntegrationConnectorsConnectionNodeConfigPtrInput
	Project    pulumi.StringPtrInput
	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount pulumi.StringPtrInput
	// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g.
	// "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
	ServiceDirectory pulumi.StringPtrInput
	// SSL Configuration of a connection
	SslConfig IntegrationConnectorsConnectionSslConfigPtrInput
	// Status of the Integration Connector.
	Statuses IntegrationConnectorsConnectionStatusArrayInput
	// This subscription type enum states the subscription type of the project.
	SubscriptionType pulumi.StringPtrInput
	// Suspended indicates if a user has suspended a connection or not.
	Suspended pulumi.BoolPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        IntegrationConnectorsConnectionTimeoutsPtrInput
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (IntegrationConnectorsConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsConnectionState)(nil)).Elem()
}

type integrationConnectorsConnectionArgs struct {
	// authConfig for the connection.
	AuthConfig *IntegrationConnectorsConnectionAuthConfig `pulumi:"authConfig"`
	// Config Variables for the connection.
	ConfigVariables []IntegrationConnectorsConnectionConfigVariable `pulumi:"configVariables"`
	// connectorVersion of the Connector.
	ConnectorVersion string `pulumi:"connectorVersion"`
	// An arbitrary description for the Connection.
	Description *string `pulumi:"description"`
	// Define the Connectors target endpoint.
	DestinationConfigs []IntegrationConnectorsConnectionDestinationConfig `pulumi:"destinationConfigs"`
	// Eventing Configuration of a connection
	EventingConfig *IntegrationConnectorsConnectionEventingConfig `pulumi:"eventingConfig"`
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
	// "ONLY_EVENTING"]
	EventingEnablementType            *string `pulumi:"eventingEnablementType"`
	IntegrationConnectorsConnectionId *string `pulumi:"integrationConnectorsConnectionId"`
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Location in which Connection needs to be created.
	Location string `pulumi:"location"`
	// Determines whether or no a connection is locked. If locked, a reason must be specified.
	LockConfig *IntegrationConnectorsConnectionLockConfig `pulumi:"lockConfig"`
	// Log configuration for the connection.
	LogConfig *IntegrationConnectorsConnectionLogConfig `pulumi:"logConfig"`
	// Name of Connection needs to be created.
	Name *string `pulumi:"name"`
	// Node configuration for the connection.
	NodeConfig *IntegrationConnectorsConnectionNodeConfig `pulumi:"nodeConfig"`
	Project    *string                                    `pulumi:"project"`
	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount *string `pulumi:"serviceAccount"`
	// SSL Configuration of a connection
	SslConfig *IntegrationConnectorsConnectionSslConfig `pulumi:"sslConfig"`
	// Suspended indicates if a user has suspended a connection or not.
	Suspended *bool                                    `pulumi:"suspended"`
	Timeouts  *IntegrationConnectorsConnectionTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IntegrationConnectorsConnection resource.
type IntegrationConnectorsConnectionArgs struct {
	// authConfig for the connection.
	AuthConfig IntegrationConnectorsConnectionAuthConfigPtrInput
	// Config Variables for the connection.
	ConfigVariables IntegrationConnectorsConnectionConfigVariableArrayInput
	// connectorVersion of the Connector.
	ConnectorVersion pulumi.StringInput
	// An arbitrary description for the Connection.
	Description pulumi.StringPtrInput
	// Define the Connectors target endpoint.
	DestinationConfigs IntegrationConnectorsConnectionDestinationConfigArrayInput
	// Eventing Configuration of a connection
	EventingConfig IntegrationConnectorsConnectionEventingConfigPtrInput
	// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
	// "ONLY_EVENTING"]
	EventingEnablementType            pulumi.StringPtrInput
	IntegrationConnectorsConnectionId pulumi.StringPtrInput
	// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Location in which Connection needs to be created.
	Location pulumi.StringInput
	// Determines whether or no a connection is locked. If locked, a reason must be specified.
	LockConfig IntegrationConnectorsConnectionLockConfigPtrInput
	// Log configuration for the connection.
	LogConfig IntegrationConnectorsConnectionLogConfigPtrInput
	// Name of Connection needs to be created.
	Name pulumi.StringPtrInput
	// Node configuration for the connection.
	NodeConfig IntegrationConnectorsConnectionNodeConfigPtrInput
	Project    pulumi.StringPtrInput
	// Service account needed for runtime plane to access Google Cloud resources.
	ServiceAccount pulumi.StringPtrInput
	// SSL Configuration of a connection
	SslConfig IntegrationConnectorsConnectionSslConfigPtrInput
	// Suspended indicates if a user has suspended a connection or not.
	Suspended pulumi.BoolPtrInput
	Timeouts  IntegrationConnectorsConnectionTimeoutsPtrInput
}

func (IntegrationConnectorsConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationConnectorsConnectionArgs)(nil)).Elem()
}

type IntegrationConnectorsConnectionInput interface {
	pulumi.Input

	ToIntegrationConnectorsConnectionOutput() IntegrationConnectorsConnectionOutput
	ToIntegrationConnectorsConnectionOutputWithContext(ctx context.Context) IntegrationConnectorsConnectionOutput
}

func (*IntegrationConnectorsConnection) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsConnection)(nil)).Elem()
}

func (i *IntegrationConnectorsConnection) ToIntegrationConnectorsConnectionOutput() IntegrationConnectorsConnectionOutput {
	return i.ToIntegrationConnectorsConnectionOutputWithContext(context.Background())
}

func (i *IntegrationConnectorsConnection) ToIntegrationConnectorsConnectionOutputWithContext(ctx context.Context) IntegrationConnectorsConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationConnectorsConnectionOutput)
}

type IntegrationConnectorsConnectionOutput struct{ *pulumi.OutputState }

func (IntegrationConnectorsConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IntegrationConnectorsConnection)(nil)).Elem()
}

func (o IntegrationConnectorsConnectionOutput) ToIntegrationConnectorsConnectionOutput() IntegrationConnectorsConnectionOutput {
	return o
}

func (o IntegrationConnectorsConnectionOutput) ToIntegrationConnectorsConnectionOutputWithContext(ctx context.Context) IntegrationConnectorsConnectionOutput {
	return o
}

// authConfig for the connection.
func (o IntegrationConnectorsConnectionOutput) AuthConfig() IntegrationConnectorsConnectionAuthConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionAuthConfigPtrOutput {
		return v.AuthConfig
	}).(IntegrationConnectorsConnectionAuthConfigPtrOutput)
}

// Config Variables for the connection.
func (o IntegrationConnectorsConnectionOutput) ConfigVariables() IntegrationConnectorsConnectionConfigVariableArrayOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionConfigVariableArrayOutput {
		return v.ConfigVariables
	}).(IntegrationConnectorsConnectionConfigVariableArrayOutput)
}

// Connection revision. This field is only updated when the connection is created or updated by User.
func (o IntegrationConnectorsConnectionOutput) ConnectionRevision() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.ConnectionRevision }).(pulumi.StringOutput)
}

// connectorVersion of the Connector.
func (o IntegrationConnectorsConnectionOutput) ConnectorVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.ConnectorVersion }).(pulumi.StringOutput)
}

// This configuration provides infra configs like rate limit threshold which need to be configurable for every connector
// version.
func (o IntegrationConnectorsConnectionOutput) ConnectorVersionInfraConfigs() IntegrationConnectorsConnectionConnectorVersionInfraConfigArrayOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionConnectorVersionInfraConfigArrayOutput {
		return v.ConnectorVersionInfraConfigs
	}).(IntegrationConnectorsConnectionConnectorVersionInfraConfigArrayOutput)
}

// Flag to mark the version indicating the launch stage.
func (o IntegrationConnectorsConnectionOutput) ConnectorVersionLaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.ConnectorVersionLaunchStage }).(pulumi.StringOutput)
}

// Time the Namespace was created in UTC.
func (o IntegrationConnectorsConnectionOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// An arbitrary description for the Connection.
func (o IntegrationConnectorsConnectionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Define the Connectors target endpoint.
func (o IntegrationConnectorsConnectionOutput) DestinationConfigs() IntegrationConnectorsConnectionDestinationConfigArrayOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionDestinationConfigArrayOutput {
		return v.DestinationConfigs
	}).(IntegrationConnectorsConnectionDestinationConfigArrayOutput)
}

func (o IntegrationConnectorsConnectionOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Eventing Configuration of a connection
func (o IntegrationConnectorsConnectionOutput) EventingConfig() IntegrationConnectorsConnectionEventingConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionEventingConfigPtrOutput {
		return v.EventingConfig
	}).(IntegrationConnectorsConnectionEventingConfigPtrOutput)
}

// Eventing enablement type. Will be nil if eventing is not enabled. Possible values: ["EVENTING_AND_CONNECTION",
// "ONLY_EVENTING"]
func (o IntegrationConnectorsConnectionOutput) EventingEnablementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringPtrOutput { return v.EventingEnablementType }).(pulumi.StringPtrOutput)
}

// Eventing Runtime Data.
func (o IntegrationConnectorsConnectionOutput) EventingRuntimeDatas() IntegrationConnectorsConnectionEventingRuntimeDataArrayOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionEventingRuntimeDataArrayOutput {
		return v.EventingRuntimeDatas
	}).(IntegrationConnectorsConnectionEventingRuntimeDataArrayOutput)
}

func (o IntegrationConnectorsConnectionOutput) IntegrationConnectorsConnectionId() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput {
		return v.IntegrationConnectorsConnectionId
	}).(pulumi.StringOutput)
}

// Resource labels to represent user provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o IntegrationConnectorsConnectionOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Location in which Connection needs to be created.
func (o IntegrationConnectorsConnectionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Determines whether or no a connection is locked. If locked, a reason must be specified.
func (o IntegrationConnectorsConnectionOutput) LockConfig() IntegrationConnectorsConnectionLockConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionLockConfigPtrOutput {
		return v.LockConfig
	}).(IntegrationConnectorsConnectionLockConfigPtrOutput)
}

// Log configuration for the connection.
func (o IntegrationConnectorsConnectionOutput) LogConfig() IntegrationConnectorsConnectionLogConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionLogConfigPtrOutput {
		return v.LogConfig
	}).(IntegrationConnectorsConnectionLogConfigPtrOutput)
}

// Name of Connection needs to be created.
func (o IntegrationConnectorsConnectionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Node configuration for the connection.
func (o IntegrationConnectorsConnectionOutput) NodeConfig() IntegrationConnectorsConnectionNodeConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionNodeConfigPtrOutput {
		return v.NodeConfig
	}).(IntegrationConnectorsConnectionNodeConfigPtrOutput)
}

func (o IntegrationConnectorsConnectionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Service account needed for runtime plane to access Google Cloud resources.
func (o IntegrationConnectorsConnectionOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// The name of the Service Directory service name. Used for Private Harpoon to resolve the ILB address. e.g.
// "projects/cloud-connectors-e2e-testing/locations/us-central1/namespaces/istio-system/services/istio-ingressgateway-connectors"
func (o IntegrationConnectorsConnectionOutput) ServiceDirectory() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.ServiceDirectory }).(pulumi.StringOutput)
}

// SSL Configuration of a connection
func (o IntegrationConnectorsConnectionOutput) SslConfig() IntegrationConnectorsConnectionSslConfigPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionSslConfigPtrOutput {
		return v.SslConfig
	}).(IntegrationConnectorsConnectionSslConfigPtrOutput)
}

// Status of the Integration Connector.
func (o IntegrationConnectorsConnectionOutput) Statuses() IntegrationConnectorsConnectionStatusArrayOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionStatusArrayOutput {
		return v.Statuses
	}).(IntegrationConnectorsConnectionStatusArrayOutput)
}

// This subscription type enum states the subscription type of the project.
func (o IntegrationConnectorsConnectionOutput) SubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.SubscriptionType }).(pulumi.StringOutput)
}

// Suspended indicates if a user has suspended a connection or not.
func (o IntegrationConnectorsConnectionOutput) Suspended() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.BoolPtrOutput { return v.Suspended }).(pulumi.BoolPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o IntegrationConnectorsConnectionOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o IntegrationConnectorsConnectionOutput) Timeouts() IntegrationConnectorsConnectionTimeoutsPtrOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) IntegrationConnectorsConnectionTimeoutsPtrOutput {
		return v.Timeouts
	}).(IntegrationConnectorsConnectionTimeoutsPtrOutput)
}

// Time the Namespace was updated in UTC.
func (o IntegrationConnectorsConnectionOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *IntegrationConnectorsConnection) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntegrationConnectorsConnectionInput)(nil)).Elem(), &IntegrationConnectorsConnection{})
	pulumi.RegisterOutputType(IntegrationConnectorsConnectionOutput{})
}
