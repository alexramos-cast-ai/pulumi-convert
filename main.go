package main

import (
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/container"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		// GKE cluster name in GCP project.
		clusterName := cfg.Require("clusterName")
		// The region to create the cluster.
		clusterRegion := cfg.Require("clusterRegion")
		// The zones to create the cluster.
		clusterZones := cfg.Require("clusterZones")
		// GCP project ID in which GKE cluster would be created.
		projectId := cfg.Require("projectId")
		// URL of alternative CAST AI API to be used during development or testing
		castaiApiUrl := "https://api.cast.ai"
		if param := cfg.Get("castaiApiUrl"); param != "" {
			castaiApiUrl = param
		}
		// CAST AI API token created in console.cast.ai API Access keys section.
		castaiApiToken := cfg.Require("castaiApiToken")
		// CAST AI gRPC URL
		castaiGrpcUrl := "grpc.cast.ai:443"
		if param := cfg.Get("castaiGrpcUrl"); param != "" {
			castaiGrpcUrl = param
		}
		// Optional parameter, if set to true - CAST AI provisioned nodes will be deleted from cloud on cluster disconnection. For production use it is recommended to set it to false.
		deleteNodesOnDisconnect := true
		if param := cfg.GetBool("deleteNodesOnDisconnect"); param {
			deleteNodesOnDisconnect = param
		}
		// Optional tags for new cluster nodes. This parameter applies only to new nodes - tags for old nodes are not reconciled.
		tags := map[string]interface{}{}
		if param := cfg.GetObject("tags"); param != nil {
			tags = param
		}
		// Cluster subnets
		subnets := cfg.Require("subnets")
		// Configure Data sources and providers required for CAST AI connection.
		_, err := organizations.GetClientConfig(ctx, map[string]interface{}{}, nil)
		if err != nil {
			return err
		}
		_ = container.LookupClusterOutput(ctx, container.GetClusterOutputArgs{
			Name:     pulumi.String(clusterName),
			Location: pulumi.String(clusterRegion),
			Project:  pulumi.String(projectId),
		}, nil)
		_, err = NewGkeIam050(ctx, "castai-gke-iam", &GkeIam050Args{
			ProjectId:      projectId,
			GkeClusterName: clusterName,
		})
		if err != nil {
			return err
		}
		_, err = NewGkeCluster7120(ctx, "castai-gke-cluster", &GkeCluster7120Args{
			ApiUrl:                    castaiApiUrl,
			CastaiApiToken:            castaiApiToken,
			GrpcUrl:                   castaiGrpcUrl,
			WaitForClusterReady:       true,
			ProjectId:                 projectId,
			GkeClusterName:            clusterName,
			GkeClusterLocation:        clusterRegion,
			GkeCredentials:            castai_gke_iam.PrivateKey,
			DeleteNodesOnDisconnect:   deleteNodesOnDisconnect,
			InstallWorkloadAutoscaler: true,
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
