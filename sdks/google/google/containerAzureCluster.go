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

type ContainerAzureCluster struct {
	pulumi.CustomResourceState

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effective_annotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAzureClusterAuthorizationOutput `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
	// call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringOutput `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication ContainerAzureClusterAzureServicesAuthenticationPtrOutput `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client                  pulumi.StringPtrOutput `pulumi:"client"`
	ContainerAzureClusterId pulumi.StringOutput    `pulumi:"containerAzureClusterId"`
	// Configuration related to the cluster control plane.
	ControlPlane ContainerAzureClusterControlPlaneOutput `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          pulumi.StringPtrOutput `pulumi:"description"`
	EffectiveAnnotations pulumi.StringMapOutput `pulumi:"effectiveAnnotations"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Fleet configuration.
	Fleet ContainerAzureClusterFleetOutput `pulumi:"fleet"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking ContainerAzureClusterNetworkingOutput `pulumi:"networking"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example:
	// `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringOutput `pulumi:"resourceGroupId"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State    pulumi.StringOutput                    `pulumi:"state"`
	Timeouts ContainerAzureClusterTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs ContainerAzureClusterWorkloadIdentityConfigArrayOutput `pulumi:"workloadIdentityConfigs"`
}

// NewContainerAzureCluster registers a new resource with the given unique name, arguments, and options.
func NewContainerAzureCluster(ctx *pulumi.Context,
	name string, args *ContainerAzureClusterArgs, opts ...pulumi.ResourceOption) (*ContainerAzureCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorization == nil {
		return nil, errors.New("invalid value for required argument 'Authorization'")
	}
	if args.AzureRegion == nil {
		return nil, errors.New("invalid value for required argument 'AzureRegion'")
	}
	if args.ControlPlane == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlane'")
	}
	if args.Fleet == nil {
		return nil, errors.New("invalid value for required argument 'Fleet'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Networking == nil {
		return nil, errors.New("invalid value for required argument 'Networking'")
	}
	if args.ResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ContainerAzureCluster
	err = ctx.RegisterPackageResource("google:index/containerAzureCluster:ContainerAzureCluster", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerAzureCluster gets an existing ContainerAzureCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerAzureCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerAzureClusterState, opts ...pulumi.ResourceOption) (*ContainerAzureCluster, error) {
	var resource ContainerAzureCluster
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/containerAzureCluster:ContainerAzureCluster", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerAzureCluster resources.
type containerAzureClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effective_annotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization *ContainerAzureClusterAuthorization `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
	// call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion *string `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication *ContainerAzureClusterAzureServicesAuthentication `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client                  *string `pulumi:"client"`
	ContainerAzureClusterId *string `pulumi:"containerAzureClusterId"`
	// Configuration related to the cluster control plane.
	ControlPlane *ContainerAzureClusterControlPlane `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          *string           `pulumi:"description"`
	EffectiveAnnotations map[string]string `pulumi:"effectiveAnnotations"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint *string `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Fleet configuration.
	Fleet *ContainerAzureClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking *ContainerAzureClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling *bool `pulumi:"reconciling"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example:
	// `/subscriptions/*/resourceGroups/*`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State    *string                        `pulumi:"state"`
	Timeouts *ContainerAzureClusterTimeouts `pulumi:"timeouts"`
	// Output only. A globally unique identifier for the cluster.
	Uid *string `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs []ContainerAzureClusterWorkloadIdentityConfig `pulumi:"workloadIdentityConfigs"`
}

type ContainerAzureClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effective_annotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAzureClusterAuthorizationPtrInput
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
	// call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringPtrInput
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication ContainerAzureClusterAzureServicesAuthenticationPtrInput
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client                  pulumi.StringPtrInput
	ContainerAzureClusterId pulumi.StringPtrInput
	// Configuration related to the cluster control plane.
	ControlPlane ContainerAzureClusterControlPlanePtrInput
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringPtrInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description          pulumi.StringPtrInput
	EffectiveAnnotations pulumi.StringMapInput
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringPtrInput
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Fleet configuration.
	Fleet ContainerAzureClusterFleetPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Cluster-wide networking configuration.
	Networking ContainerAzureClusterNetworkingPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolPtrInput
	// The ARM ID of the resource group where the cluster resources are deployed. For example:
	// `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringPtrInput
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State    pulumi.StringPtrInput
	Timeouts ContainerAzureClusterTimeoutsPtrInput
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringPtrInput
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringPtrInput
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs ContainerAzureClusterWorkloadIdentityConfigArrayInput
}

func (ContainerAzureClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAzureClusterState)(nil)).Elem()
}

type containerAzureClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effective_annotations` for all of the annotations present on the resource.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAzureClusterAuthorization `pulumi:"authorization"`
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
	// call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion string `pulumi:"azureRegion"`
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication *ContainerAzureClusterAzureServicesAuthentication `pulumi:"azureServicesAuthentication"`
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client                  *string `pulumi:"client"`
	ContainerAzureClusterId *string `pulumi:"containerAzureClusterId"`
	// Configuration related to the cluster control plane.
	ControlPlane ContainerAzureClusterControlPlane `pulumi:"controlPlane"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// Fleet configuration.
	Fleet ContainerAzureClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Cluster-wide networking configuration.
	Networking ContainerAzureClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The ARM ID of the resource group where the cluster resources are deployed. For example:
	// `/subscriptions/*/resourceGroups/*`
	ResourceGroupId string                         `pulumi:"resourceGroupId"`
	Timeouts        *ContainerAzureClusterTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ContainerAzureCluster resource.
type ContainerAzureClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
	// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
	// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
	// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
	// `effective_annotations` for all of the annotations present on the resource.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	Authorization ContainerAzureClusterAuthorizationInput
	// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
	// call to list all supported Azure regions within a given Google Cloud region.
	AzureRegion pulumi.StringInput
	// Azure authentication configuration for management of Azure resources
	AzureServicesAuthentication ContainerAzureClusterAzureServicesAuthenticationPtrInput
	// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
	// `AzureCluster`. `AzureClient` names are formatted as
	// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
	// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
	Client                  pulumi.StringPtrInput
	ContainerAzureClusterId pulumi.StringPtrInput
	// Configuration related to the cluster control plane.
	ControlPlane ContainerAzureClusterControlPlaneInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// Fleet configuration.
	Fleet ContainerAzureClusterFleetInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Cluster-wide networking configuration.
	Networking ContainerAzureClusterNetworkingInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The ARM ID of the resource group where the cluster resources are deployed. For example:
	// `/subscriptions/*/resourceGroups/*`
	ResourceGroupId pulumi.StringInput
	Timeouts        ContainerAzureClusterTimeoutsPtrInput
}

func (ContainerAzureClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerAzureClusterArgs)(nil)).Elem()
}

type ContainerAzureClusterInput interface {
	pulumi.Input

	ToContainerAzureClusterOutput() ContainerAzureClusterOutput
	ToContainerAzureClusterOutputWithContext(ctx context.Context) ContainerAzureClusterOutput
}

func (*ContainerAzureCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAzureCluster)(nil)).Elem()
}

func (i *ContainerAzureCluster) ToContainerAzureClusterOutput() ContainerAzureClusterOutput {
	return i.ToContainerAzureClusterOutputWithContext(context.Background())
}

func (i *ContainerAzureCluster) ToContainerAzureClusterOutputWithContext(ctx context.Context) ContainerAzureClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAzureClusterOutput)
}

type ContainerAzureClusterOutput struct{ *pulumi.OutputState }

func (ContainerAzureClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAzureCluster)(nil)).Elem()
}

func (o ContainerAzureClusterOutput) ToContainerAzureClusterOutput() ContainerAzureClusterOutput {
	return o
}

func (o ContainerAzureClusterOutput) ToContainerAzureClusterOutputWithContext(ctx context.Context) ContainerAzureClusterOutput {
	return o
}

// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of
// all keys and values combined is limited to 256k. Keys can have 2 segments: prefix (optional) and name (required),
// separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with
// alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between. **Note**: This field is
// non-authoritative, and will only manage the annotations present in your configuration. Please refer to the field
// `effective_annotations` for all of the annotations present on the resource.
func (o ContainerAzureClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Configuration related to the cluster RBAC settings.
func (o ContainerAzureClusterOutput) Authorization() ContainerAzureClusterAuthorizationOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterAuthorizationOutput { return v.Authorization }).(ContainerAzureClusterAuthorizationOutput)
}

// The Azure region where the cluster runs. Each Google Cloud region supports a subset of nearby Azure regions. You can
// call to list all supported Azure regions within a given Google Cloud region.
func (o ContainerAzureClusterOutput) AzureRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.AzureRegion }).(pulumi.StringOutput)
}

// Azure authentication configuration for management of Azure resources
func (o ContainerAzureClusterOutput) AzureServicesAuthentication() ContainerAzureClusterAzureServicesAuthenticationPtrOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterAzureServicesAuthenticationPtrOutput {
		return v.AzureServicesAuthentication
	}).(ContainerAzureClusterAzureServicesAuthenticationPtrOutput)
}

// Name of the AzureClient. The `AzureClient` resource must reside on the same GCP project and region as the
// `AzureCluster`. `AzureClient` names are formatted as
// `projects/<project-number>/locations/<region>/azureClients/<client-id>`. See Resource Names
// (https:cloud.google.com/apis/design/resource_names) for more details on Google Cloud resource names.
func (o ContainerAzureClusterOutput) Client() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringPtrOutput { return v.Client }).(pulumi.StringPtrOutput)
}

func (o ContainerAzureClusterOutput) ContainerAzureClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.ContainerAzureClusterId }).(pulumi.StringOutput)
}

// Configuration related to the cluster control plane.
func (o ContainerAzureClusterOutput) ControlPlane() ContainerAzureClusterControlPlaneOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterControlPlaneOutput { return v.ControlPlane }).(ContainerAzureClusterControlPlaneOutput)
}

// Output only. The time at which this cluster was created.
func (o ContainerAzureClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
func (o ContainerAzureClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ContainerAzureClusterOutput) EffectiveAnnotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringMapOutput { return v.EffectiveAnnotations }).(pulumi.StringMapOutput)
}

// Output only. The endpoint of the cluster's API server.
func (o ContainerAzureClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
// and delete requests to ensure the client has an up-to-date value before proceeding.
func (o ContainerAzureClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Fleet configuration.
func (o ContainerAzureClusterOutput) Fleet() ContainerAzureClusterFleetOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterFleetOutput { return v.Fleet }).(ContainerAzureClusterFleetOutput)
}

// The location for the resource
func (o ContainerAzureClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of this resource.
func (o ContainerAzureClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Cluster-wide networking configuration.
func (o ContainerAzureClusterOutput) Networking() ContainerAzureClusterNetworkingOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterNetworkingOutput { return v.Networking }).(ContainerAzureClusterNetworkingOutput)
}

// The project for the resource
func (o ContainerAzureClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. If set, there are currently changes in flight to the cluster.
func (o ContainerAzureClusterOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// The ARM ID of the resource group where the cluster resources are deployed. For example:
// `/subscriptions/*/resourceGroups/*`
func (o ContainerAzureClusterOutput) ResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.ResourceGroupId }).(pulumi.StringOutput)
}

// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
// STOPPING, ERROR, DEGRADED
func (o ContainerAzureClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o ContainerAzureClusterOutput) Timeouts() ContainerAzureClusterTimeoutsPtrOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterTimeoutsPtrOutput { return v.Timeouts }).(ContainerAzureClusterTimeoutsPtrOutput)
}

// Output only. A globally unique identifier for the cluster.
func (o ContainerAzureClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Output only. The time at which this cluster was last updated.
func (o ContainerAzureClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Output only. Workload Identity settings.
func (o ContainerAzureClusterOutput) WorkloadIdentityConfigs() ContainerAzureClusterWorkloadIdentityConfigArrayOutput {
	return o.ApplyT(func(v *ContainerAzureCluster) ContainerAzureClusterWorkloadIdentityConfigArrayOutput {
		return v.WorkloadIdentityConfigs
	}).(ContainerAzureClusterWorkloadIdentityConfigArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerAzureClusterInput)(nil)).Elem(), &ContainerAzureCluster{})
	pulumi.RegisterOutputType(ContainerAzureClusterOutput{})
}
