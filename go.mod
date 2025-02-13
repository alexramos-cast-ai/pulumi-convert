module gke_cluster_existing

go 1.22

toolchain go1.23.4

require (
	github.com/pulumi/pulumi-gcp/sdk/v8 v8.18.0
	github.com/pulumi/pulumi/sdk/v3 v3.147.0
)

require github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7 v7.0.0-00010101000000-000000000000 // indirect

replace github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7 => ./sdks/castai

replace github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6 => ./sdks/google

replace github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6 => ./sdks/google-beta

replace github.com/pulumi/pulumi-terraform-provider/sdks/go/helm/v2 => ./sdks/helm
