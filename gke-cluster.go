package main

import (
	"fmt"

	// "github.com/pulumi/pulumi-castai/sdk/go/castai"
	// "github.com/pulumi/pulumi-helm/sdk/go/helm"
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	"github.com/pulumi/pulumi-null/sdk/go/null"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// func notImplemented(message string) pulumi.AnyOutput {
// 	panic(message)
// }

type GkeClusterArgs struct {
	ApiUrl                       pulumi.StringInput
	CastaiApiToken               pulumi.StringInput
	GrpcUrl                      pulumi.StringInput
	ApiGrpcAddr                  pulumi.StringInput
	KvisorControllerExtraArgs    map[string]pulumi.StringInput
	ProjectId                    pulumi.StringInput
	GkeClusterName               pulumi.StringInput
	AutoscalerPoliciesJson       pulumi.StringInput
	AutoscalerSettings           interface{}
	DeleteNodesOnDisconnect      pulumi.BoolInput
	GkeClusterLocation           pulumi.StringInput
	GkeCredentials               pulumi.StringInput
	CastaiComponentsLabels       map[string]interface{}
	NodeConfigurations           interface{}
	DefaultNodeConfiguration     pulumi.StringInput
	DefaultNodeConfigurationName pulumi.StringInput
	NodeTemplates                interface{}
	WorkloadScalingPolicies      interface{}
	InstallSecurityAgent         pulumi.BoolInput
	AgentVersion                 pulumi.StringInput
	ClusterControllerVersion     pulumi.StringInput
	EvictorVersion               pulumi.StringInput
	EvictorExtVersion            pulumi.StringInput
	PodPinnerVersion             pulumi.StringInput
	SpotHandlerVersion           pulumi.StringInput
	KvisorVersion                pulumi.StringInput
	AgentValues                  []pulumi.StringInput
	SpotHandlerValues            []pulumi.StringInput
	ClusterControllerValues      []pulumi.StringInput
	EvictorValues                []pulumi.StringInput
	EvictorExtValues             []pulumi.StringInput
	PodPinnerValues              []pulumi.StringInput
	KvisorValues                 []pulumi.StringInput
	SelfManaged                  pulumi.BoolInput
	WaitForClusterReady          pulumi.BoolInput
	InstallWorkloadAutoscaler    pulumi.BoolInput
	WorkloadAutoscalerVersion    pulumi.StringInput
	WorkloadAutoscalerValues     []pulumi.StringInput
	InstallCloudProxy            pulumi.BoolInput
	CloudProxyVersion            pulumi.StringInput
	CloudProxyValues             []pulumi.StringInput
	CloudProxyGrpcUrlOverride    pulumi.StringInput
}

type GkeCluster struct {
	pulumi.ResourceState
	ClusterId                pulumi.AnyOutput
	CastaiNodeConfigurations pulumi.AnyOutput
	CastaiNodeTemplates      pulumi.AnyOutput
}

func NewGkeCluster(
	ctx *pulumi.Context,
	name string,
	args *GkeClusterArgs,
	opts ...pulumi.ResourceOption,
) (*GkeCluster, error) {
	var componentResource GkeCluster
	err := ctx.RegisterComponentResource("components:index:GkeCluster7120", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	castaiCluster, err := castai.NewGkeCluster(ctx, fmt.Sprintf("%s-castai_cluster", name), &castai.GkeClusterArgs{
		ProjectId:               args.ProjectId,
		Location:                args.GkeClusterLocation,
		Name:                    args.GkeClusterName,
		DeleteNodesOnDisconnect: args.DeleteNodesOnDisconnect,
		CredentialsJson:         args.GkeCredentials,
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	var this []*castai.NodeConfiguration
	for key0, _ := range "TODO: For expression" {
		__res, err := castai.NewNodeConfiguration(ctx, fmt.Sprintf("%s-this-%v", name, key0), &castai.NodeConfigurationArgs{
			Gke: []map[string]interface{}{
				map[string]interface{}{
					"loadbalancers":               "TODO: For expression",
					"maxPodsPerNode":              notImplemented("try(each.value.max_pods_per_node,110)"),
					"networkTags":                 notImplemented("try(each.value.network_tags,null)"),
					"diskType":                    notImplemented("try(each.value.disk_type,null)"),
					"useEphemeralStorageLocalSsd": notImplemented("try(each.value.use_ephemeral_storage_local_ssd,null)"),
				},
			},
			ClusterId:       castaiCluster.Id,
			Name:            notImplemented("try(each.value.name,each.key)"),
			DiskCpuRatio:    notImplemented("try(each.value.disk_cpu_ratio,0)"),
			DrainTimeoutSec: notImplemented("try(each.value.drain_timeout_sec,0)"),
			MinDiskSize:     notImplemented("try(each.value.min_disk_size,100)"),
			Subnets:         notImplemented("try(each.value.subnets,null)"),
			SshPublicKey:    notImplemented("try(each.value.ssh_public_key,null)"),
			Image:           notImplemented("try(each.value.image,null)"),
			Tags:            notImplemented("try(each.value.tags,{})"),
			InitScript:      notImplemented("try(each.value.init_script,null)"),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		this = append(this, __res)
	}
	_, err = castai.NewNodeConfigurationDefault(ctx, fmt.Sprintf("%s-this", name), &castai.NodeConfigurationDefaultArgs{
		ClusterId: castaiCluster.Id,
		// ConfigurationId: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered conditional expression @ main.pp:43,21-122),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	castaiAgent, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_agent", name), &helm.ReleaseArgs{
		Set: []interface{}{
			map[string]interface{}{
				"name":  "replicaCount",
				"value": "2",
			},
			map[string]interface{}{
				"name":  "provider",
				"value": "gke",
			},
			map[string]interface{}{
				"name":  "additionalEnv.STATIC_CLUSTER_ID",
				"value": castaiCluster.Id,
			},
			map[string]interface{}{
				"name":  "createNamespace",
				"value": "false",
			},
		},
		Name:            "castai-agent",
		Repository:      "https://castai.github.io/helm-charts",
		Chart:           "castai-agent",
		Namespace:       "castai-agent",
		CreateNamespace: true,
		CleanupOnFail:   true,
		Wait:            true,
		Version:         args.AgentVersion,
		Values:          args.AgentValues,
		SetSensitive: []map[string]interface{}{
			map[string]interface{}{
				"name":  "apiKey",
				"value": castaiCluster.ClusterToken,
			},
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	var tmp0 float64
	if args.SelfManaged {
		tmp0 = 0
	} else {
		tmp0 = 1
	}
	var castaiEvictor []*helm.Release
	for index := 0; index < tmp0; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_evictor-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "replicaCount",
					"value": "0",
				},
				map[string]interface{}{
					"name":  "castai-evictor-ext.enabled",
					"value": "false",
				},
			},
			Name:            "castai-evictor",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-evictor",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.EvictorVersion,
			Values:          args.EvictorValues,
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiEvictor = append(castaiEvictor, __res)
	}
	var tmp1 float64
	if args.SelfManaged {
		tmp1 = 0
	} else {
		tmp1 = 1
	}
	var castaiPodPinner []*helm.Release
	for index := 0; index < tmp1; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_pod_pinner-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
				map[string]interface{}{
					"name":  "replicaCount",
					"value": "0",
				},
			},
			Name:            "castai-pod-pinner",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-pod-pinner",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.PodPinnerVersion,
			Values:          args.PodPinnerValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiPodPinner = append(castaiPodPinner, __res)
	}
	castaiAutoscalerPolicies, err := castai.NewAutoscaler(ctx, fmt.Sprintf("%s-castai_autoscaler_policies", name), &castai.AutoscalerArgs{
		AutoscalerSettings:     "TODO: For expression",
		ClusterId:              castaiCluster.Id,
		AutoscalerPoliciesJson: args.AutoscalerPoliciesJson,
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
		castaiEvictor,
		castaiPodPinner,
	}))
	if err != nil {
		return nil, err
	}
	var thisNodeTemplate []*castai.NodeTemplate
	for key0, _ := range "TODO: For expression" {
		__res, err := castai.NewNodeTemplate(ctx, fmt.Sprintf("%s-this-%v", name, key0), &castai.NodeTemplateArgs{
			CustomTaints: "TODO: For expression",
			Constraints:  "TODO: For expression",
			ClusterId:    castaiCluster.Id,
			Name:         notImplemented("try(each.value.name,each.key)"),
			// ConfigurationId: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered conditional expression @ main.pp:114,46-210),
			IsDefault:                                notImplemented("try(each.value.is_default,false)"),
			IsEnabled:                                notImplemented("try(each.value.is_enabled,true)"),
			ShouldTaint:                              notImplemented("try(each.value.should_taint,true)"),
			CustomInstancesEnabled:                   notImplemented("try(each.value.custom_instances_enabled,false)"),
			CustomInstancesWithExtendedMemoryEnabled: notImplemented("try(each.value.custom_instances_with_extended_memory_enabled,false)"),
			CustomLabels:                             notImplemented("try(each.value.custom_labels,{})"),
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAutoscalerPolicies,
		}))
		if err != nil {
			return nil, err
		}
		thisNodeTemplate = append(thisNodeTemplate, __res)
	}
	var thisWorkloadScalingPolicy []*castai.WorkloadScalingPolicy
	for key0, _ := range "TODO: For expression" {
		__res, err := castai.NewWorkloadScalingPolicy(ctx, fmt.Sprintf("%s-this-%v", name, key0), &castai.WorkloadScalingPolicyArgs{
			Name:             notImplemented("try(each.value.name,each.key)"),
			ClusterId:        castaiCluster.Id,
			ApplyType:        notImplemented("try(each.value.apply_type,\"DEFERRED\")"),
			ManagementOption: notImplemented("try(each.value.management_option,\"READ_ONLY\")"),
			Cpu: []map[string]interface{}{
				map[string]interface{}{
					"function":       notImplemented("try(each.value.cpu.function,\"QUANTILE\")"),
					"overhead":       notImplemented("try(each.value.cpu.overhead,0)"),
					"applyThreshold": notImplemented("try(each.value.cpu.apply_threshold,0.1)"),
					"args":           notImplemented("try(each.value.cpu.args,[\"0.8\"])"),
				},
			},
			Memory: []map[string]interface{}{
				map[string]interface{}{
					"function":       notImplemented("try(each.value.memory.function,\"MAX\")"),
					"overhead":       notImplemented("try(each.value.memory.overhead,0.1)"),
					"applyThreshold": notImplemented("try(each.value.memory.apply_threshold,0.1)"),
				},
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		thisWorkloadScalingPolicy = append(thisWorkloadScalingPolicy, __res)
	}
	var tmp2 float64
	if args.SelfManaged {
		tmp2 = 0
	} else {
		tmp2 = 1
	}
	var castaiClusterController []*helm.Release
	for index := 0; index < tmp2; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_cluster_controller-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
			},
			Name:            "cluster-controller",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-cluster-controller",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.ClusterControllerVersion,
			Values:          args.ClusterControllerValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiClusterController = append(castaiClusterController, __res)
	}
	var tmp3 float64
	if args.SelfManaged {
		tmp3 = 1
	} else {
		tmp3 = 0
	}
	var castaiClusterControllerSelfManaged []*helm.Release
	for index := 0; index < tmp3; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_cluster_controller_self_managed-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
			},
			Name:            "cluster-controller",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-cluster-controller",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.ClusterControllerVersion,
			Values:          args.ClusterControllerValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiClusterControllerSelfManaged = append(castaiClusterControllerSelfManaged, __res)
	}
	var tmp4 float64
	if args.WaitForClusterReady {
		tmp4 = 1
	} else {
		tmp4 = 0
	}
	var waitForCluster []*null.Resource
	for index := 0; index < tmp4; index++ {
		key0 := index
		_ := index
		__res, err := null.NewResource(ctx, fmt.Sprintf("%s-wait_for_cluster-%v", name, key0), nil, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiClusterController,
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		waitForCluster = append(waitForCluster, __res)
	}
	_, err = local.NewCommand(ctx, fmt.Sprintf("%s-waitForClusterProvisioner0", name), &local.CommandArgs{
		Create: pulumi.Sprintf(`RETRY_COUNT=20
POOLING_INTERVAL=30

for i in $(seq 1 $RETRY_COUNT); do
    sleep $POOLING_INTERVAL
    curl -s %v/v1/kubernetes/external-clusters/%v -H \"x-api-key: $API_KEY\" | grep '\"status\"\\s*:\\s*\"ready\"' && exit 0doneecho \"Cluster is not ready after 10 minutes\"exit 1
`, args.ApiUrl, castaiCluster.Id),
		Interpreter: pulumi.StringArray{
			pulumi.String("bash"),
			pulumi.String("-c"),
		},
		Environment: pulumi.StringMap{
			"API_KEY": args.CastaiApiToken,
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		waitForCluster,
	}))
	if err != nil {
		return nil, err
	}
	var tmp5 float64
	if args.SelfManaged {
		tmp5 = 1
	} else {
		tmp5 = 0
	}
	var castaiEvictorSelfManaged []*helm.Release
	for index := 0; index < tmp5; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_evictor_self_managed-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai-evictor-ext.enabled",
					"value": "false",
				},
				map[string]interface{}{
					"name":  "managedByCASTAI",
					"value": "false",
				},
			},
			Name:            "castai-evictor",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-evictor",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.EvictorVersion,
			Values:          args.EvictorValues,
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiEvictorSelfManaged = append(castaiEvictorSelfManaged, __res)
	}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_evictor_ext", name), &helm.ReleaseArgs{
		Name:            "castai-evictor-ext",
		Repository:      "https://castai.github.io/helm-charts",
		Chart:           "castai-evictor-ext",
		Namespace:       "castai-agent",
		CreateNamespace: false,
		CleanupOnFail:   true,
		Wait:            true,
		Version:         args.EvictorExtVersion,
		Values:          args.EvictorExtValues,
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiEvictor,
	}))
	if err != nil {
		return nil, err
	}
	var tmp6 float64
	if args.SelfManaged {
		tmp6 = 1
	} else {
		tmp6 = 0
	}
	var castaiPodPinnerSelfManaged []*helm.Release
	for index := 0; index < tmp6; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_pod_pinner_self_managed-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
				map[string]interface{}{
					"name":  "managedByCASTAI",
					"value": "false",
				},
			},
			Name:            "castai-pod-pinner",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-pod-pinner",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.PodPinnerVersion,
			Values:          args.PodPinnerValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
		}))
		if err != nil {
			return nil, err
		}
		castaiPodPinnerSelfManaged = append(castaiPodPinnerSelfManaged, __res)
	}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_spot_handler", name), &helm.ReleaseArgs{
		Set: []interface{}{
			map[string]interface{}{
				"name":  "castai.provider",
				"value": "gcp",
			},
			map[string]interface{}{
				"name":  "createNamespace",
				"value": "false",
			},
			map[string]interface{}{
				"name":  "castai.clusterID",
				"value": castaiCluster.Id,
			},
		},
		Name:            "castai-spot-handler",
		Repository:      "https://castai.github.io/helm-charts",
		Chart:           "castai-spot-handler",
		Namespace:       "castai-agent",
		CreateNamespace: true,
		CleanupOnFail:   true,
		Wait:            true,
		Version:         args.SpotHandlerVersion,
		Values:          args.SpotHandlerValues,
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}
	var tmp7 float64
	if args.InstallSecurityAgent && !args.SelfManaged {
		tmp7 = 1
	} else {
		tmp7 = 0
	}
	var castaiKvisor []*helm.Release
	for index := 0; index < tmp7; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_kvisor-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
				map[string]interface{}{
					"name":  "castai.grpcAddr",
					"value": args.ApiGrpcAddr,
				},
				map[string]interface{}{
					"name":  "controller.extraArgs.kube-bench-cloud-provider",
					"value": "gke",
				},
			},
			Name:            "castai-kvisor",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-kvisor",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Version:         args.KvisorVersion,
			Values:          args.KvisorValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		castaiKvisor = append(castaiKvisor, __res)
	}
	var tmp8 float64
	if args.InstallCloudProxy {
		tmp8 = 1
	} else {
		tmp8 = 0
	}
	var castaiCloudProxy []*helm.Release
	for index := 0; index < tmp8; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_cloud_proxy-%v", name, key0), &helm.ReleaseArgs{
			Name:            "castai-cloud-proxy",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-cloud-proxy",
			Version:         args.CloudProxyVersion,
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Values:          args.CloudProxyValues,
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
				map[string]interface{}{
					"name":  "castai.grpcURL",
					"value": notImplemented("coalesce(var.cloud_proxy_grpc_url_override,var.grpc_url)"),
				},
			},
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		castaiCloudProxy = append(castaiCloudProxy, __res)
	}
	var tmp9 float64
	if args.InstallSecurityAgent && args.SelfManaged {
		tmp9 = 1
	} else {
		tmp9 = 0
	}
	var castaiKvisorSelfManaged []*helm.Release
	for index := 0; index < tmp9; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_kvisor_self_managed-%v", name, key0), &helm.ReleaseArgs{
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.clusterID",
					"value": castaiCluster.Id,
				},
				map[string]interface{}{
					"name":  "castai.grpcAddr",
					"value": args.ApiGrpcAddr,
				},
				map[string]interface{}{
					"name":  "controller.extraArgs.kube-bench-cloud-provider",
					"value": "gke",
				},
			},
			Name:            "castai-kvisor",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-kvisor",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Version:         args.KvisorVersion,
			Values:          args.KvisorValues,
			SetSensitive: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKey",
					"value": castaiCluster.ClusterToken,
				},
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		castaiKvisorSelfManaged = append(castaiKvisorSelfManaged, __res)
	}
	//---------------------------------------------------#
	// CAST.AI Workload Autoscaler configuration         #
	//---------------------------------------------------#
	var tmp10 float64
	if args.InstallWorkloadAutoscaler && !args.SelfManaged {
		tmp10 = 1
	} else {
		tmp10 = 0
	}
	var castaiWorkloadAutoscaler []*helm.Release
	for index := 0; index < tmp10; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_workload_autoscaler-%v", name, key0), &helm.ReleaseArgs{
			Name:            "castai-workload-autoscaler",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-workload-autoscaler",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.WorkloadAutoscalerVersion,
			Values:          args.WorkloadAutoscalerValues,
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKeySecretRef",
					"value": "castai-cluster-controller",
				},
				map[string]interface{}{
					"name":  "castai.configMapRef",
					"value": "castai-cluster-controller",
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
			castaiClusterController,
		}))
		if err != nil {
			return nil, err
		}
		castaiWorkloadAutoscaler = append(castaiWorkloadAutoscaler, __res)
	}
	var tmp11 float64
	if args.InstallWorkloadAutoscaler && args.SelfManaged {
		tmp11 = 1
	} else {
		tmp11 = 0
	}
	var castaiWorkloadAutoscalerSelfManaged []*helm.Release
	for index := 0; index < tmp11; index++ {
		key0 := index
		_ := index
		__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_workload_autoscaler_self_managed-%v", name, key0), &helm.ReleaseArgs{
			Name:            "castai-workload-autoscaler",
			Repository:      "https://castai.github.io/helm-charts",
			Chart:           "castai-workload-autoscaler",
			Namespace:       "castai-agent",
			CreateNamespace: true,
			CleanupOnFail:   true,
			Wait:            true,
			Version:         args.WorkloadAutoscalerVersion,
			Values:          args.WorkloadAutoscalerValues,
			Set: []map[string]interface{}{
				map[string]interface{}{
					"name":  "castai.apiKeySecretRef",
					"value": "castai-cluster-controller",
				},
				map[string]interface{}{
					"name":  "castai.configMapRef",
					"value": "castai-cluster-controller",
				},
			},
		}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
			castaiAgent,
			castaiClusterController,
		}))
		if err != nil {
			return nil, err
		}
		castaiWorkloadAutoscalerSelfManaged = append(castaiWorkloadAutoscalerSelfManaged, __res)
	}
	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"clusterId":                castaiCluster.Id,
		"castaiNodeConfigurations": "TODO: For expression",
		"castaiNodeTemplates":      "TODO: For expression",
	})
	if err != nil {
		return nil, err
	}
	componentResource.ClusterId = castaiCluster.Id
	componentResource.CastaiNodeConfigurations = "TODO: For expression"
	componentResource.CastaiNodeTemplates = "TODO: For expression"
	return &componentResource, nil
}
