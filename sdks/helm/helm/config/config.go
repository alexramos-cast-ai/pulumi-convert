// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/helm/v2/helm/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

// Helm burst limit. Increase this if you have a cluster with many CRDs
func GetBurstLimit(ctx *pulumi.Context) float64 {
	return config.GetFloat64(ctx, "helm:burstLimit")
}

// Debug indicates whether or not Helm is running in Debug mode.
func GetDebug(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "helm:debug")
}

// Enable and disable experimental features.
func GetExperiments(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:experiments")
}

// The backend storage driver. Values are: configmap, secret, memory, sql
func GetHelmDriver(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:helmDriver")
}

// Kubernetes configuration.
func GetKubernetes(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:kubernetes")
}

// The path to the helm plugins directory
func GetPluginsPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:pluginsPath")
}

// RegistryClient configuration.
func GetRegistries(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:registries")
}

// The path to the registry config file
func GetRegistryConfigPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:registryConfigPath")
}

// The path to the file containing cached repository indexes
func GetRepositoryCache(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:repositoryCache")
}

// The path to the file containing repository names and URLs
func GetRepositoryConfigPath(ctx *pulumi.Context) string {
	return config.Get(ctx, "helm:repositoryConfigPath")
}
