package castai

import (
	"fmt"

	// "github.com/pulumi/pulumi-castai/sdk/go/castai"
	// "github.com/pulumi/pulumi-helm/sdk/go/helm"

	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
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
	NodeConfigurations           pulumi.Map
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
	Subnets                      pulumi.StringArrayInput
}

type GkeCluster struct {
	pulumi.ResourceState
	ClusterId                pulumi.AnyOutput
	CastaiNodeConfigurations pulumi.AnyOutput
	CastaiNodeTemplates      pulumi.AnyOutput
}

func NewGkeCluster(ctx *pulumi.Context, name string, args *GkeClusterArgs, opts ...pulumi.ResourceOption) (*GkeCluster, error) {
	var componentResource GkeCluster
	err := ctx.RegisterComponentResource("components:index:GkeCluster", name, &componentResource, opts...)
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

	// Helm release

	agentValues := pulumi.Map{
		"provider": pulumi.String("gke"),
		"additionalEnv": pulumi.Map{
			"STATIC_CLUSTER_ID": castaiCluster.GkeClusterId,
		},
		"apiKey": pulumi.StringInput(args.CastaiApiToken),
	}

	castaiAgent, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_agent", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-agent"),
		Chart:           pulumi.String("castai-agent"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(true),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.AgentVersion,
		Values:          agentValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// TODO: Set condition if evictor is self managed and if enabled
	// var tmp0 float64
	// if args.SelfManaged {
	// 	tmp0 = 0
	// } else {
	// 	tmp0 = 1
	// }
	evictorValues := pulumi.Map{}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_evictor", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-evictor"),
		Chart:           pulumi.String("castai-evictor"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.EvictorVersion,
		Values:          evictorValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}

	// TODO: Support condition if Pod pinner is enabled
	// var tmp1 float64
	// if args.SelfManaged {
	// 	tmp1 = 0
	// } else {
	// 	tmp1 = 1
	// }

	podPinnerValues := pulumi.Map{
		"castai": pulumi.Map{
			"apiKey":    pulumi.StringInput(castaiCluster.ClusterToken),
			"clusterID": pulumi.StringInput(castaiCluster.GkeClusterId),
		},
	}

	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_pod_pinner", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-pod-pinner"),
		Chart:           pulumi.String("castai-pod-pinner"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.PodPinnerVersion,
		Values:          podPinnerValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}

	// TODO: Support condition for sef-manage component
	// var tmp2 float64
	// if args.SelfManaged {
	// 	tmp2 = 0
	// } else {
	// 	tmp2 = 1
	// }
	controllerValues := pulumi.Map{
		"castai": pulumi.Map{
			"apiKey":    pulumi.StringInput(castaiCluster.ClusterToken),
			"clusterID": pulumi.StringInput(castaiCluster.GkeClusterId),
		},
	}

	castaiClusterController, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_cluster_controller", name), &helm.ReleaseArgs{
		Name:            pulumi.String("cluster-controller"),
		Chart:           pulumi.String("castai-cluster-controller"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.ClusterControllerVersion,
		Values:          controllerValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}

	// TODO: Check this NewCommand which I think only waits for cluster to be ready
	// 	_, err = local.NewCommand(ctx, fmt.Sprintf("%s-waitForClusterProvisioner0", name), &local.CommandArgs{
	// 		Create: pulumi.Sprintf(`RETRY_COUNT=20
	// POOLING_INTERVAL=30

	// for i in $(seq 1 $RETRY_COUNT); do
	//     sleep $POOLING_INTERVAL
	//     curl -s %v/v1/kubernetes/external-clusters/%v -H \"x-api-key: $API_KEY\" | grep '\"status\"\\s*:\\s*\"ready\"' && exit 0doneecho \"Cluster is not ready after 10 minutes\"exit 1
	// `, args.ApiUrl, castaiCluster.Id),
	// 		Interpreter: pulumi.StringArray{
	// 			pulumi.String("bash"),
	// 			pulumi.String("-c"),
	// 		},
	// 		Environment: pulumi.StringMap{
	// 			"API_KEY": args.CastaiApiToken,
	// 		},
	// 	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
	// 		waitForCluster,
	// 	}))
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// TODO: Support condition if enabled
	// var tmp5 float64
	// if args.SelfManaged {
	// 	tmp5 = 1
	// } else {
	// 	tmp5 = 0
	// }

	// TODO: Whats evictor-ext
	// _, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_evictor_ext", name), &helm.ReleaseArgs{
	// 	Name:            "castai-evictor-ext",
	// 	Repository:      "https://castai.github.io/helm-charts",
	// 	Chart:           "castai-evictor-ext",
	// 	Namespace:       "castai-agent",
	// 	CreateNamespace: false,
	// 	CleanupOnFail:   true,
	// 	Wait:            true,
	// 	Version:         args.EvictorExtVersion,
	// 	Values:          args.EvictorExtValues,
	// }, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
	// 	castaiEvictor,
	// }))
	// if err != nil {
	// 	return nil, err
	// }

	spotHandlerValues := pulumi.Map{
		"castai": pulumi.Map{
			"provider":  pulumi.String("gcp"),
			"clusterID": pulumi.StringInput(castaiCluster.GkeClusterId),
		},
	}

	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_spot_handler", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-spot-handler"),
		Chart:           pulumi.String("castai-spot-handler"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.SpotHandlerVersion,
		Values:          spotHandlerValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}

	// TODO: Support Condition for installing kvisor
	// var tmp7 float64
	// if args.InstallSecurityAgent && !args.SelfManaged {
	// 	tmp7 = 1
	// } else {
	// 	tmp7 = 0
	// }

	kvisorValues := pulumi.Map{
		"castai": pulumi.Map{
			"clusterID": castaiCluster.GkeClusterId,
			"apiKey":    pulumi.StringInput(args.CastaiApiToken),
		},
		"controller": pulumi.Map{
			"extraArgs": pulumi.Map{
				"kube-bench-cloud-provider": pulumi.String("gke"),
			},
		},
	}
	_, err = helm.NewRelease(ctx, fmt.Sprintf("%s-castai_kvisor", name), &helm.ReleaseArgs{
		Name:            pulumi.String("castai-kvisor"),
		Chart:           pulumi.String("castai-kvisor"),
		Namespace:       pulumi.String("castai-agent"),
		CreateNamespace: pulumi.Bool(false),
		CleanupOnFail:   pulumi.Bool(true),
		Version:         args.KvisorVersion,
		Values:          kvisorValues,
		RepositoryOpts: &helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://castai.github.io/helm-charts"),
		},
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
	}))
	if err != nil {
		return nil, err
	}

	// TODO: Add support for CloudProxy
	// var tmp8 float64
	// if args.InstallCloudProxy {
	// 	tmp8 = 1
	// } else {
	// 	tmp8 = 0
	// }
	// var castaiCloudProxy []*helm.Release
	// for index := 0; index < tmp8; index++ {
	// 	key0 := index
	// 	_ := index
	// 	__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_cloud_proxy-%v", name, key0), &helm.ReleaseArgs{
	// 		Name:            "castai-cloud-proxy",
	// 		Repository:      "https://castai.github.io/helm-charts",
	// 		Chart:           "castai-cloud-proxy",
	// 		Version:         args.CloudProxyVersion,
	// 		Namespace:       "castai-agent",
	// 		CreateNamespace: true,
	// 		CleanupOnFail:   true,
	// 		Wait:            true,
	// 		Values:          args.CloudProxyValues,
	// 		Set: []map[string]interface{}{
	// 			map[string]interface{}{
	// 				"name":  "castai.clusterID",
	// 				"value": castaiCluster.Id,
	// 			},
	// 			map[string]interface{}{
	// 				"name":  "castai.grpcURL",
	// 				"value": notImplemented("coalesce(var.cloud_proxy_grpc_url_override,var.grpc_url)"),
	// 			},
	// 		},
	// 		SetSensitive: []map[string]interface{}{
	// 			map[string]interface{}{
	// 				"name":  "castai.apiKey",
	// 				"value": castaiCluster.ClusterToken,
	// 			},
	// 		},
	// 	}, pulumi.Parent(&componentResource))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	castaiCloudProxy = append(castaiCloudProxy, __res)
	// }
	// var tmp9 float64
	// if args.InstallSecurityAgent && args.SelfManaged {
	// 	tmp9 = 1
	// } else {
	// 	tmp9 = 0
	// }

	// NODE CONFIGURATION
	// TODO: Support custom node configuration

	var networkTags pulumi.StringArray

	nodeConfig := &castai.NodeConfigurationGkeArgs{
		// Loadbalancers:               &castai.NodeConfigurationGkeLoadbalancerArgs{},
		MaxPodsPerNode:              pulumi.Float64(100),
		NetworkTags:                 networkTags,
		DiskType:                    pulumi.String("pd-standard"),
		UseEphemeralStorageLocalSsd: pulumi.Bool(false),
	}

	// var this []*castai.NodeConfiguration
	// for _, value := range args.NodeConfigurations {
	newNodeConfigurationRes, err := castai.NewNodeConfiguration(ctx, fmt.Sprintf("%s-this-", name), &castai.NodeConfigurationArgs{
		Gke:             nodeConfig,
		ClusterId:       castaiCluster.GkeClusterId,
		Name:            pulumi.String("default"),
		DiskCpuRatio:    pulumi.Float64(0),
		DrainTimeoutSec: pulumi.Float64(100),
		MinDiskSize:     pulumi.Float64(100),
		Subnets:         args.Subnets,
		// SshPublicKey:    pulumi.String("EMPTY_SSH"),
		// Image:           pulumi.String("EMPTY_IMAGE"),
		Tags: pulumi.StringMap{
			"name": pulumi.String("test"),
		},
		// InitScript: pulumi.String("EMPTY_INIT_SCRIPT"),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		castaiAgent,
		castaiClusterController,
	}))
	if err != nil {
		return nil, err
	}
	// this = append(this, __res)
	// }

	// Default node configuration
	_, err = castai.NewNodeConfigurationDefault(ctx, fmt.Sprintf("%s-this", name), &castai.NodeConfigurationDefaultArgs{
		ClusterId:       castaiCluster.GkeClusterId,
		ConfigurationId: newNodeConfigurationRes.ID(),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		newNodeConfigurationRes,
	}))
	if err != nil {
		return nil, err
	}

	_, err = castai.NewAutoscaler(ctx, fmt.Sprintf("%s-castai_autoscaler_policies", name), &castai.AutoscalerArgs{
		AutoscalerSettings: &castai.AutoscalerAutoscalerSettingsArgs{
			Enabled: pulumi.Bool(true),
		},
		ClusterId:              castaiCluster.GkeClusterId,
		AutoscalerPoliciesJson: args.AutoscalerPoliciesJson,
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		newNodeConfigurationRes,
	}))
	if err != nil {
		return nil, err
	}

	// NODE TEMPLATE
	nodeTemplate, err := castai.NewNodeTemplate(ctx, fmt.Sprintf("%s-this", name), &castai.NodeTemplateArgs{
		CustomTaints:                             &castai.NodeTemplateCustomTaintArray{},
		Constraints:                              &castai.NodeTemplateConstraintsArgs{},
		ClusterId:                                castaiCluster.GkeClusterId,
		Name:                                     pulumi.Sprintf("%s_nodetemplate", name),
		ConfigurationId:                          newNodeConfigurationRes.ID(),
		IsDefault:                                pulumi.Bool(false),
		IsEnabled:                                pulumi.Bool(true),
		ShouldTaint:                              pulumi.Bool(false),
		CustomInstancesEnabled:                   pulumi.Bool(false),
		CustomInstancesWithExtendedMemoryEnabled: pulumi.Bool(false),
		// CustomLabels:                             pulumi.StringMapInput(pulumi.ToStringMap{}),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		newNodeConfigurationRes,
	}))
	if err != nil {
		return nil, err
	}

	// var thisWorkloadScalingPolicy []*castai.WorkloadScalingPolicy
	// for key0, _ := range "TODO: For expression" {
	// 	__res, err := castai.NewWorkloadScalingPolicy(ctx, fmt.Sprintf("%s-this-%v", name, key0), &castai.WorkloadScalingPolicyArgs{
	// 		Name:             notImplemented("try(each.value.name,each.key)"),
	// 		ClusterId:        castaiCluster.Id,
	// 		ApplyType:        notImplemented("try(each.value.apply_type,\"DEFERRED\")"),
	// 		ManagementOption: notImplemented("try(each.value.management_option,\"READ_ONLY\")"),
	// 		Cpu: []map[string]interface{}{
	// 			map[string]interface{}{
	// 				"function":       notImplemented("try(each.value.cpu.function,\"QUANTILE\")"),
	// 				"overhead":       notImplemented("try(each.value.cpu.overhead,0)"),
	// 				"applyThreshold": notImplemented("try(each.value.cpu.apply_threshold,0.1)"),
	// 				"args":           notImplemented("try(each.value.cpu.args,[\"0.8\"])"),
	// 			},
	// 		},
	// 		Memory: []map[string]interface{}{
	// 			map[string]interface{}{
	// 				"function":       notImplemented("try(each.value.memory.function,\"MAX\")"),
	// 				"overhead":       notImplemented("try(each.value.memory.overhead,0.1)"),
	// 				"applyThreshold": notImplemented("try(each.value.memory.apply_threshold,0.1)"),
	// 			},
	// 		},
	// 	}, pulumi.Parent(&componentResource))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	thisWorkloadScalingPolicy = append(thisWorkloadScalingPolicy, __res)
	// }

	//TODO: Add support for workload autoscaler helmrelease
	//---------------------------------------------------#
	// CAST.AI Workload Autoscaler configuration         #
	//---------------------------------------------------#
	// var tmp10 float64
	// if args.InstallWorkloadAutoscaler && !args.SelfManaged {
	// 	tmp10 = 1
	// } else {
	// 	tmp10 = 0
	// }
	// var castaiWorkloadAutoscaler []*helm.Release
	// for index := 0; index < tmp10; index++ {
	// 	key0 := index
	// 	_ := index
	// 	__res, err := helm.NewRelease(ctx, fmt.Sprintf("%s-castai_workload_autoscaler-%v", name, key0), &helm.ReleaseArgs{
	// 		Name:            "castai-workload-autoscaler",
	// 		Repository:      "https://castai.github.io/helm-charts",
	// 		Chart:           "castai-workload-autoscaler",
	// 		Namespace:       "castai-agent",
	// 		CreateNamespace: true,
	// 		CleanupOnFail:   true,
	// 		Wait:            true,
	// 		Version:         args.WorkloadAutoscalerVersion,
	// 		Values:          args.WorkloadAutoscalerValues,
	// 		Set: []map[string]interface{}{
	// 			map[string]interface{}{
	// 				"name":  "castai.apiKeySecretRef",
	// 				"value": "castai-cluster-controller",
	// 			},
	// 			map[string]interface{}{
	// 				"name":  "castai.configMapRef",
	// 				"value": "castai-cluster-controller",
	// 			},
	// 		},
	// 	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
	// 		castaiAgent,
	// 		castaiClusterController,
	// 	}))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	castaiWorkloadAutoscaler = append(castaiWorkloadAutoscaler, __res)
	// }

	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"clusterId":                castaiCluster.GkeClusterId,
		"castaiNodeConfigurations": newNodeConfigurationRes,
		"castaiNodeTemplates":      nodeTemplate,
	})
	if err != nil {
		return nil, err
	}
	componentResource.ClusterId = pulumi.AnyOutput(castaiCluster.GkeClusterId)
	// componentResource.CastaiNodeConfigurations = nodeConfigurationRes.ConfigurationId
	// componentResource.CastaiNodeTemplates = "TODO: For expression"
	return &componentResource, nil
}
