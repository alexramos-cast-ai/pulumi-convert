package main

import (
	"github.com/castai/pulumi-convert/castai"
	"github.com/castai/pulumi-convert/gke"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		gkeArgs := gke.GkeIamArgs{
			ProjectId:      pulumi.String("success-team-dev"),
			GkeClusterName: pulumi.String("gke-021701-acr"),
			// ServiceAccountsUniqueIds:          pulumi.StringArray{},
			// ComputeManagerProjectIds:          pulumi.StringArray{},
			// CreateServiceAccount:              pulumi.Bool(false),
			// SetupCloudProxyWorkloadIdentity:   pulumi.Bool(false),
			// WorkloadIdentityNamespace:         pulumi.String(""),
			// CloudProxyServiceAccountNamespace: pulumi.String(""),
			// CloudProxyServiceAccountName:      pulumi.String(""),
			CastaiRolePermissions:     pulumi.StringArray{},
			ComputeManagerPermissions: pulumi.StringArray{},
		}
		clusterRegion := "us-central1"
		subnets := pulumi.StringArray{}
		tags := pulumi.StringMap{}

		gkeClusterArgs := castai.GkeClusterArgs{
			ApiUrl:                    pulumi.String("https://api.cast.ai"),
			CastaiApiToken:            pulumi.String("3a9a4ebf136f2abdc3dc4484e117678dd5edf07c4d63b1d3bbb58bd4b5a08625"),
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
			NodeConfigurations:           pulumi.Map{},
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

		gkeIamRes, err := gke.NewGkeIam(ctx, "castai-gke-iam", &gke.GkeIamArgs{
			ProjectId:      gkeArgs.ProjectId,
			GkeClusterName: gkeArgs.GkeClusterName,
		})
		if err != nil {
			return err
		}

		key := gkeIamRes.PrivateKey

		_, err = castai.NewGkeCluster(ctx, "castai-gke-cluster", &castai.GkeClusterArgs{
			ApiUrl:                    gkeClusterArgs.ApiUrl,
			CastaiApiToken:            gkeClusterArgs.CastaiApiToken,
			GrpcUrl:                   gkeClusterArgs.GrpcUrl,
			WaitForClusterReady:       gkeClusterArgs.WaitForClusterReady,
			ProjectId:                 gkeClusterArgs.ProjectId,
			GkeClusterName:            gkeClusterArgs.GkeClusterName,
			GkeClusterLocation:        gkeClusterArgs.GkeClusterLocation,
			GkeCredentials:            pulumi.StringInput(key),
			DeleteNodesOnDisconnect:   gkeClusterArgs.DeleteNodesOnDisconnect,
			InstallWorkloadAutoscaler: gkeClusterArgs.InstallWorkloadAutoscaler,
			NodeConfigurations: pulumi.Map{
				"default": pulumi.Map{
					"minDiskSize":  pulumi.String("100"),
					"diskCpuRatio": pulumi.String("0"),
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
