package main

import (
	"github.com/castai/pulumi-convert/castai"
	"github.com/castai/pulumi-convert/gke"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		conf := config.New(ctx, "")
		apiKey := conf.Require("apiToken")
		// TODO: source config via YAML/JSON file perhaps?
		subnets := pulumi.StringArray{
			pulumi.String("projects/success-team-dev/regions/asia-southeast1/subnetworks/default"),
		}

		gkeArgs := &gke.GkeIamArgs{
			ProjectId:                pulumi.String("<GOOGLE_PROJECT_ID"),
			GkeClusterName:           pulumi.String("<CLUSTER_NAME>"),
			ServiceAccountsUniqueIds: pulumi.String("<DEFAULT_COMPUTE_ENGINE_SERVICEACCOUNT_UNIQUEID>"), // TODO: How to get this dynamically
			// ComputeManagerProjectIds:          pulumi.StringArray{},
			// CreateServiceAccount:              pulumi.Bool(false),
			// SetupCloudProxyWorkloadIdentity:   pulumi.Bool(false),
			// WorkloadIdentityNamespace:         pulumi.String(""),
			// CloudProxyServiceAccountNamespace: pulumi.String(""),
			// CloudProxyServiceAccountName:      pulumi.String(""),
			CastaiRolePermissions:     pulumi.StringArray{},
			ComputeManagerPermissions: pulumi.StringArray{},
		}
		clusterRegion := "asia-southeast1"
		tags := pulumi.StringMap{}

		gkeClusterArgs := &castai.GkeClusterArgs{
			ApiUrl:                    pulumi.String("https://api.cast.ai"),
			CastaiApiToken:            pulumi.String(apiKey),
			GrpcUrl:                   pulumi.String("grpc.cast.ai:443"),
			ApiGrpcAddr:               pulumi.String("grpc.cast.ai:443"),
			KvisorControllerExtraArgs: pulumi.StringMap{},
			ProjectId:                 gkeArgs.ProjectId,
			GkeClusterName:            gkeArgs.GkeClusterName,
			AutoscalerPoliciesJson:    pulumi.String(""),
			AutoscalerSettings:        pulumi.StringMap{},
			DeleteNodesOnDisconnect:   pulumi.Bool(true),
			GkeClusterLocation:        pulumi.String("asia-southeast1-a"),
			GkeCredentials:            pulumi.String(""),
			// CastaiComponentsLabels:       pulumi.StringArray{},
			Subnets:                      subnets,
			NodeConfigurations:           pulumi.Map{},
			DefaultNodeConfiguration:     pulumi.String("default"),
			DefaultNodeConfigurationName: pulumi.String("default"),
			NodeTemplates:                pulumi.StringArray{},
			WorkloadScalingPolicies:      pulumi.StringArray{},
			InstallSecurityAgent:         pulumi.Bool(true),
			AgentVersion:                 pulumi.String("0.97.0"),
			ClusterControllerVersion:     pulumi.String("0.80.0"),
			EvictorVersion:               pulumi.String("0.31.42"),
			EvictorExtVersion:            pulumi.String(""),
			PodPinnerVersion:             pulumi.String("1.0.1"),
			SpotHandlerVersion:           pulumi.String("0.26.0"),
			KvisorVersion:                pulumi.String("1.0.71"),
			AgentValues:                  pulumi.StringArray{},
			SpotHandlerValues:            pulumi.StringArray{},
			ClusterControllerValues:      pulumi.StringArray{},
			EvictorValues:                pulumi.StringArray{},
			EvictorExtValues:             pulumi.StringArray{},
			PodPinnerValues:              pulumi.StringArray{},
			KvisorValues:                 pulumi.StringArray{},
			SelfManaged:                  pulumi.Bool(true),
			WaitForClusterReady:          pulumi.Bool(false),
			InstallWorkloadAutoscaler:    pulumi.Bool(false),
			WorkloadAutoscalerVersion:    pulumi.String(""),
			WorkloadAutoscalerValues:     pulumi.StringArray{},
			InstallCloudProxy:            pulumi.Bool(true),
			CloudProxyVersion:            pulumi.String(""),
			CloudProxyValues:             pulumi.StringArray{},
			CloudProxyGrpcUrlOverride:    pulumi.String(""),
		}

		// Get the GCP client config
		_, err := organizations.GetClientConfig(ctx, nil, nil)
		if err != nil {
			return err
		}

		_ = container.LookupClusterOutput(ctx, container.LookupClusterOutputArgs{
			Name:     gkeArgs.GkeClusterName,
			Location: pulumi.String(clusterRegion),
			Project:  gkeArgs.ProjectId,
		}, nil)

		gkeIamRes, err := gke.NewGkeIam(ctx, "castai-gke-iam", &gke.GkeIamArgs{
			ProjectId:                gkeArgs.ProjectId,
			GkeClusterName:           gkeArgs.GkeClusterName,
			ServiceAccountsUniqueIds: gkeArgs.ServiceAccountsUniqueIds,
		})
		if err != nil {
			return err
		}

		_, err = castai.NewGkeCluster(ctx, "castai-gke-cluster", &castai.GkeClusterArgs{
			ApiUrl:                    gkeClusterArgs.ApiUrl,
			CastaiApiToken:            pulumi.String(apiKey),
			GrpcUrl:                   gkeClusterArgs.GrpcUrl,
			WaitForClusterReady:       gkeClusterArgs.WaitForClusterReady,
			ProjectId:                 gkeClusterArgs.ProjectId,
			GkeClusterName:            gkeClusterArgs.GkeClusterName,
			GkeClusterLocation:        gkeClusterArgs.GkeClusterLocation,
			GkeCredentials:            gkeIamRes.PrivateKey,
			DeleteNodesOnDisconnect:   gkeClusterArgs.DeleteNodesOnDisconnect,
			InstallWorkloadAutoscaler: gkeClusterArgs.InstallWorkloadAutoscaler,
			AgentVersion:              gkeClusterArgs.AgentVersion,
			ClusterControllerVersion:  gkeClusterArgs.ClusterControllerVersion,
			EvictorVersion:            gkeClusterArgs.EvictorExtVersion,
			Subnets:                   gkeClusterArgs.Subnets,
			SelfManaged:               gkeClusterArgs.SelfManaged,
			NodeConfigurations: pulumi.Map{
				"default": pulumi.Map{
					"minDiskSize":  pulumi.String("100"),
					"diskCpuRatio": pulumi.String("0"),
					"subnets":      subnets,
					"tags":         tags,
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			gkeIamRes,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
