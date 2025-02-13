package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		gkeArgs := GkeIamArgs{
			ProjectId:                         pulumi.String("my-project"),
			GkeClusterName:                    pulumi.String("my-cluster"),
			ServiceAccountsUniqueIds:          pulumi.StringArray{},
			ComputeManagerProjectIds:          pulumi.StringArray{},
			CreateServiceAccount:              pulumi.Bool(false),
			SetupCloudProxyWorkloadIdentity:   pulumi.Bool(false),
			WorkloadIdentityNamespace:         pulumi.String(""),
			CloudProxyServiceAccountNamespace: pulumi.String(""),
			CloudProxyServiceAccountName:      pulumi.String(""),
			CastaiRolePermissions:             pulumi.StringArray{},
			ComputeManagerPermissions:         pulumi.StringArray{},
		}
		clusterRegion := "us-central1"
		subnets := pulumi.StringArray{}
		tags := pulumi.StringMap{}

		gkeClusterArgs := GkeClusterArgs{
			ApiUrl:                    pulumi.String("https://api.cast.ai"),
			CastaiApiToken:            pulumi.String("my-token"),
			GrpcUrl:                   pulumi.String("grpc.cast.ai:443"),
			ApiGrpcAddr:               pulumi.String("grpc.cast.ai:443"),
			KvisorControllerExtraArgs: pulumi.StringMap{},
			ProjectId:                 pulumi.String("my-project"),
			GkeClusterName:            pulumi.String("my-cluster"),
			AutoscalerPoliciesJson:    pulumi.String(""),
			AutoscalerSettings:        pulumi.StringMap{},
			DeleteNodesOnDisconnect:   pulumi.Bool(true),
			GkeClusterLocation:        pulumi.String("us-central1"),
			GkeCredentials:            pulumi.String(""),
			// CastaiComponentsLabels:       pulumi.StringArray{},
			NodeConfigurations:           pulumi.StringMap{},
			DefaultNodeConfiguration:     pulumi.String("default"),
			DefaultNodeConfigurationName: pulumi.String("default"),
			NodeTemplates:                pulumi.StringArray{},
			WorkloadScalingPolicies:      pulumi.StringArray{},
			InstallSecurityAgent:         pulumi.Bool(true),
			AgentVersion:                 pulumi.String(""),
			ClusterControllerVersion:     pulumi.String(""),
			EvictorVersion:               pulumi.String(""),
			EvictorExtVersion:            pulumi.String(""),
			PodPinnerVersion:             pulumi.String(""),
			SpotHandlerVersion:           pulumi.String(""),
			KvisorVersion:                pulumi.String(""),
			AgentValues:                  pulumi.StringArray{},
			SpotHandlerValues:            pulumi.StringArray{},
			ClusterControllerValues:      pulumi.StringArray{},
			EvictorValues:                pulumi.StringArray{},
			EvictorExtValues:             pulumi.StringArray{},
			PodPinnerValues:              pulumi.StringArray{},
			KvisorValues:                 pulumi.StringArray{},
			SelfManaged:                  pulumi.Bool(false),
			WaitForClusterReady:          pulumi.Bool(true),
			InstallWorkloadAutoscaler:    pulumi.Bool(true),
			WorkloadAutoscalerVersion:    pulumi.String(""),
			WorkloadAutoscalerValues:     pulumi.StringArray{},
			InstallCloudProxy:            pulumi.Bool(true),
			CloudProxyVersion:            pulumi.String(""),
			CloudProxyValues:             pulumi.StringArray{},
			CloudProxyGrpcUrlOverride:    pulumi.String(""),
		}

		// Get the GCP client config
		gkeClientConfig, err := organizations.GetClientConfig(ctx, nil, nil)
		if err != nil {
			return err
		}
		token := gkeClientConfig.AccessToken

		token = token

		_ = container.LookupClusterOutput(ctx, container.LookupClusterOutputArgs{
			Name:     gkeArgs.GkeClusterName,
			Location: pulumi.String(clusterRegion),
			Project:  gkeArgs.ProjectId,
		}, nil)

		_, err = NewGkeIam(ctx, "castai-gke-iam", &GkeIamArgs{
			ProjectId:      gkeArgs.ProjectId,
			GkeClusterName: gkeArgs.GkeClusterName,
		})
		if err != nil {
			return err
		}

		_, err = NewGkeCluster(ctx, "castai-gke-cluster", &GkeClusterArgs{
			ApiUrl:                    gkeClusterArgs.ApiUrl,
			CastaiApiToken:            gkeClusterArgs.CastaiApiToken,
			GrpcUrl:                   gkeClusterArgs.GrpcUrl,
			WaitForClusterReady:       gkeClusterArgs.WaitForClusterReady,
			ProjectId:                 gkeClusterArgs.ProjectId,
			GkeClusterName:            gkeClusterArgs.GkeClusterName,
			GkeClusterLocation:        gkeClusterArgs.GkeClusterLocation,
			GkeCredentials:            gkeClusterArgs.GkeCredentials,
			DeleteNodesOnDisconnect:   gkeClusterArgs.DeleteNodesOnDisconnect,
			InstallWorkloadAutoscaler: gkeClusterArgs.InstallWorkloadAutoscaler,
			NodeConfigurations: map[string]interface{}{
				"default": map[string]interface{}{
					"minDiskSize":  100,
					"diskCpuRatio": 0,
					"subnets":      subnets,
					"tags":         tags,
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
