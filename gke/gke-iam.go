package gke

import (
	"fmt"

	"crypto/sha1"
	"encoding/hex"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// func notImplemented(message string) pulumi.AnyOutput {
//   panic(message)
// }

type GkeIamArgs struct {
	ProjectId      pulumi.StringInput
	GkeClusterName pulumi.StringInput
	// ServiceAccountsUniqueIds []pulumi.StringInput
	// ComputeManagerProjectIds []pulumi.StringInput
	// CreateServiceAccount              pulumi.Bool
	// SetupCloudProxyWorkloadIdentity   pulumi.BoolInput
	// WorkloadIdentityNamespace         pulumi.StringInput
	// CloudProxyServiceAccountNamespace pulumi.StringInput
	// CloudProxyServiceAccountName      pulumi.StringInput
	CastaiRolePermissions     []pulumi.StringInput
	ComputeManagerPermissions []pulumi.StringInput
}

type GkeIam struct {
	pulumi.ResourceState
	PrivateKey                       pulumi.AnyOutput
	ServiceAccountId                 pulumi.AnyOutput
	ServiceAccountEmail              pulumi.AnyOutput
	DefaultComputeManagerPermissions pulumi.AnyOutput
	DefaultCastaiRolePermissions     pulumi.AnyOutput
}

var DefaultComputeManagerPermissions = []string{
	"container.clusters.get",
	"container.clusters.update",
	"container.certificateSigningRequests.approve",
	"compute.instances.get",
	"compute.instances.list",
	"compute.instances.create",
	"compute.instances.start",
	"compute.instances.stop",
	"compute.instances.delete",
	"compute.instances.setLabels",
	"compute.instances.setServiceAccount",
	"compute.instances.setMetadata",
	"compute.instances.setTags",
	"compute.instanceGroupManagers.get",
	"compute.instanceGroupManagers.update",
	"compute.instanceGroups.get",
	"compute.networks.use",
	"compute.networks.useExternalIp",
	"compute.subnetworks.get",
	"compute.subnetworks.use",
	"compute.subnetworks.useExternalIp",
	"compute.addresses.use",
	"compute.disks.use",
	"compute.disks.create",
	"compute.disks.setLabels",
	"compute.images.get",
	"compute.images.useReadOnly",
	"compute.instanceTemplates.get",
	"compute.instanceTemplates.list",
	"compute.instanceTemplates.create",
	"compute.instanceTemplates.delete",
	"compute.regionOperations.get",
	"compute.zoneOperations.get",
	"compute.zones.list",
	"compute.zones.get",
	"serviceusage.services.list",
	"resourcemanager.projects.getIamPolicy",
	"compute.targetPools.get",
	"compute.targetPools.addInstance",
	"compute.targetPools.removeInstance",
	"compute.instances.use",
}
var DefaultCastaiRolePermissions = []string{
	"container.clusters.get",
	"container.clusters.update",
	"container.certificateSigningRequests.approve",
	"compute.instances.get",
	"compute.instances.list",
	"compute.instances.create",
	"compute.instances.start",
	"compute.instances.stop",
	"compute.instances.delete",
	"compute.instances.setLabels",
	"compute.instances.setServiceAccount",
	"compute.instances.setMetadata",
	"compute.instances.setTags",
	"compute.instanceGroupManagers.get",
	"compute.instanceGroupManagers.update",
	"compute.instanceGroups.get",
	"compute.networks.use",
	"compute.networks.useExternalIp",
	"compute.subnetworks.get",
	"compute.subnetworks.use",
	"compute.subnetworks.useExternalIp",
	"compute.addresses.use",
	"compute.disks.use",
	"compute.disks.create",
	"compute.disks.setLabels",
	"compute.images.get",
	"compute.images.useReadOnly",
	"compute.instanceTemplates.get",
	"compute.instanceTemplates.list",
	"compute.instanceTemplates.create",
	"compute.instanceTemplates.delete",
	"compute.regionOperations.get",
	"compute.zoneOperations.get",
	"compute.zones.list",
	"compute.zones.get",
	"serviceusage.services.list",
	"resourcemanager.projects.getIamPolicy",
	"compute.targetPools.get",
	"compute.targetPools.addInstance",
	"compute.targetPools.removeInstance",
	"compute.instances.use",
}

func NewGkeIam(ctx *pulumi.Context, name string, args *GkeIamArgs, opts ...pulumi.ResourceOption) (*GkeIam, error) {
	var componentResource GkeIam
	err := ctx.RegisterComponentResource("components:index:GkeIam", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}

	// GKE IAM ROLE PREPERATION
	// Create clustername hash
	clusterNameHash := sha1.New()
	clusterNameHash.Write([]byte(fmt.Sprintf("%s", args.GkeClusterName)))
	clusterHashString := hex.EncodeToString(clusterNameHash.Sum(nil))

	// Set Custom Role ID
	customRoleId := pulumi.String(fmt.Sprintf("castai.gkeAccess.%v.tf", clusterHashString[:8]))

	// conditionExpression := pulumi.String(strings.Join([]string{"compact([\"resource.type==\\\"gke.io/cluster\\\"\","}, "||"))

	// defaultPermissions := []string{
	// 	fmt.Sprintf("roles/iam.serviceAccountUser"),
	// 	fmt.Sprintf("projects/%v/roles/%v", args.ProjectId, customRoleId),
	// }
	// scopedPermissions := []string{
	// 	fmt.Sprintf("projects/%v/roles/%v", args.ProjectId, customRoleId),
	// }

	// computeManagerProjectIds := fmt.Sprintf("%s", args.ComputeManagerProjectIds)
	// computeManagerProjectIdRes, err := castai.GetGkeUserPolicies(ctx, &castai.GetGkeUserPoliciesArgs{
	// 	&computeManagerProjectIds,
	// })
	// if err != nil {
	// 	return nil, err
	// }

	var castaiRolePermissions pulumi.StringArray
	if float64(len(args.CastaiRolePermissions)) > 0 {
		castaiRolePermissions = args.CastaiRolePermissions
	} else {
		castaiRolePermissions = pulumi.ToStringArray(DefaultCastaiRolePermissions)
	}

	// CREATING THE CLUSTER IAM ROLES
	_, err = projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s-castai_role", name), &projects.IAMCustomRoleArgs{
		RoleId:      pulumi.String(customRoleId),
		Title:       pulumi.String("Role to manage GKE cluster via CAST AI"),
		Description: pulumi.String("Role to manage GKE cluster via CAST AI"),
		Permissions: castaiRolePermissions,
		Project:     args.ProjectId,
		Stage:       pulumi.String("GA"),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// TODO: Allow support for supplied permissions
	// var computeManagerPermissions pulumi.StringArray
	// if float64(len(args.ComputeManagerPermissions)) > 0 {
	// 	computeManagerPermissions = args.ComputeManagerPermissions
	// } else {
	// 	computeManagerPermissions = pulumi.StringArray{
	// 		pulumi.String("iam.roles.list"),
	// 		pulumi.String("iam.roles.list"),
	// 		pulumi.String("iam.roles.list"),
	// 	}
	// }

	// var computeManagerRole []*projects.IAMCustomRole
	// for _, values := range computeManagerProjectIdRes {
	// 	__res, err := projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s-compute_manager_role-%v", name, values), &projects.IAMCustomRoleArgs{
	// 		Project:     pulumi.String(computeManagerProjectIdRes.Id),
	// 		RoleId:      pulumi.String(computeManagerProjectIdRes.Id),
	// 		Title:       pulumi.String("Role to manage GKE compute resources via CAST AI"),
	// 		Description: pulumi.String("Role to manage GKE compute resources via CAST AI"),
	// 		Permissions: computeManagerPermissions,
	// 		Stage:       pulumi.String("GA"),
	// 	}, pulumi.Parent(&componentResource))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	computeManagerRole = append(computeManagerRole, __res)
	// }
	var castaiComputeRolePermission pulumi.StringArray
	if float64(len(args.ComputeManagerPermissions)) > 0 {
		castaiComputeRolePermission = args.ComputeManagerPermissions
	} else {
		castaiComputeRolePermission = pulumi.ToStringArray(DefaultCastaiRolePermissions)
	}

	computeRoleId := pulumi.String(fmt.Sprintf("castai.gkeAccess.%v.tf", clusterHashString[:8]))
	computeManagerRole, err := projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s-compute_manager_role-%v", name, "1"), &projects.IAMCustomRoleArgs{
		Project:     args.ProjectId,
		RoleId:      pulumi.String(computeRoleId),
		Title:       pulumi.String("Role to manage GKE compute resources via CAST AI"),
		Description: pulumi.String("Role to manage GKE compute resources via CAST AI"),
		Permissions: castaiComputeRolePermission,
		Stage:       pulumi.String("GA"),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// BINDING THE GKE IAM ROLES
	// TODO: Allow support for supplied role
	// var computeManagerBinding []*projects.IAMBinding
	// for _, values := range computeManagerRole {
	// 	__res, err := projects.NewIAMBinding(ctx, fmt.Sprintf("%s-compute_manager_binding-%v", name, values), &projects.IAMBindingArgs{
	// 		Project: values.Project,
	// 		Role:    values.RoleId,
	// 		Members: pulumi.StringArray{
	// 			pulumi.String("ServiceAccountEmail"), // TODO: Replace with the output of service account email
	// 		},
	// 	}, pulumi.Parent(&componentResource))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	computeManagerBinding = append(computeManagerBinding, __res)
	// }

	_, err = projects.NewIAMBinding(ctx, fmt.Sprintf("%s-compute_manager_binding", name), &projects.IAMBindingArgs{
		Project: args.ProjectId,
		Role:    computeManagerRole.Name,
		Members: pulumi.StringArray{
			pulumi.String("ServiceAccountEmail"), // TODO: Replace with the output of service account email
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
	// 	"privateKey":          gkeIam.PrivateKey,
	// 	"serviceAccountId":    gkeIam.AccountId,
	// 	"serviceAccountEmail": gkeIam.ServiceAccountEmail,
	// 	"defaultComputeManagerPermissions": []string{
	// 		"container.clusters.get",
	// 		"container.clusters.update",
	// 		"container.certificateSigningRequests.approve",
	// 		"compute.instances.get",
	// 		"compute.instances.list",
	// 		"compute.instances.create",
	// 		"compute.instances.start",
	// 		"compute.instances.stop",
	// 		"compute.instances.delete",
	// 		"compute.instances.setLabels",
	// 		"compute.instances.setServiceAccount",
	// 		"compute.instances.setMetadata",
	// 		"compute.instances.setTags",
	// 		"compute.instanceGroupManagers.get",
	// 		"compute.instanceGroupManagers.update",
	// 		"compute.instanceGroups.get",
	// 		"compute.networks.use",
	// 		"compute.networks.useExternalIp",
	// 		"compute.subnetworks.get",
	// 		"compute.subnetworks.use",
	// 		"compute.subnetworks.useExternalIp",
	// 		"compute.addresses.use",
	// 		"compute.disks.use",
	// 		"compute.disks.create",
	// 		"compute.disks.setLabels",
	// 		"compute.images.get",
	// 		"compute.images.useReadOnly",
	// 		"compute.instanceTemplates.get",
	// 		"compute.instanceTemplates.list",
	// 		"compute.instanceTemplates.create",
	// 		"compute.instanceTemplates.delete",
	// 		"compute.regionOperations.get",
	// 		"compute.zoneOperations.get",
	// 		"compute.zones.list",
	// 		"compute.zones.get",
	// 		"serviceusage.services.list",
	// 		"resourcemanager.projects.getIamPolicy",
	// 		"compute.targetPools.get",
	// 		"compute.targetPools.addInstance",
	// 		"compute.targetPools.removeInstance",
	// 		"compute.instances.use",
	// 	},
	// 	"defaultCastaiRolePermissions": []string{
	// 		"container.clusters.get",
	// 		"container.clusters.update",
	// 		"container.certificateSigningRequests.approve",
	// 		"compute.instances.get",
	// 		"compute.instances.list",
	// 		"compute.instances.create",
	// 		"compute.instances.start",
	// 		"compute.instances.stop",
	// 		"compute.instances.delete",
	// 		"compute.instances.setLabels",
	// 		"compute.instances.setServiceAccount",
	// 		"compute.instances.setMetadata",
	// 		"compute.instances.setTags",
	// 		"compute.instanceGroupManagers.get",
	// 		"compute.instanceGroupManagers.update",
	// 		"compute.instanceGroups.get",
	// 		"compute.networks.use",
	// 		"compute.networks.useExternalIp",
	// 		"compute.subnetworks.get",
	// 		"compute.subnetworks.use",
	// 		"compute.subnetworks.useExternalIp",
	// 		"compute.addresses.use",
	// 		"compute.disks.use",
	// 		"compute.disks.create",
	// 		"compute.disks.setLabels",
	// 		"compute.images.get",
	// 		"compute.images.useReadOnly",
	// 		"compute.instanceTemplates.get",
	// 		"compute.instanceTemplates.list",
	// 		"compute.instanceTemplates.create",
	// 		"compute.instanceTemplates.delete",
	// 		"compute.regionOperations.get",
	// 		"compute.zoneOperations.get",
	// 		"compute.zones.list",
	// 		"compute.zones.get",
	// 		"serviceusage.services.list",
	// 		"resourcemanager.projects.getIamPolicy",
	// 		"compute.targetPools.get",
	// 		"compute.targetPools.addInstance",
	// 		"compute.targetPools.removeInstance",
	// 		"compute.instances.use",
	// 	},
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// componentResource.PrivateKey = gkeClientConfig.AccessToken
	// componentResource.ServiceAccountId = gkeClientConfig.AccountId
	// componentResource.ServiceAccountEmail = gkeClientCOnfig.ServiceAccountEmail
	// componentResource.DefaultComputeManagerPermissions = []string{
	// 	"container.clusters.get",
	// 	"container.clusters.update",
	// 	"container.certificateSigningRequests.approve",
	// 	"compute.instances.get",
	// 	"compute.instances.list",
	// 	"compute.instances.create",
	// 	"compute.instances.start",
	// 	"compute.instances.stop",
	// 	"compute.instances.delete",
	// 	"compute.instances.setLabels",
	// 	"compute.instances.setServiceAccount",
	// 	"compute.instances.setMetadata",
	// 	"compute.instances.setTags",
	// 	"compute.instanceGroupManagers.get",
	// 	"compute.instanceGroupManagers.update",
	// 	"compute.instanceGroups.get",
	// 	"compute.networks.use",
	// 	"compute.networks.useExternalIp",
	// 	"compute.subnetworks.get",
	// 	"compute.subnetworks.use",
	// 	"compute.subnetworks.useExternalIp",
	// 	"compute.addresses.use",
	// 	"compute.disks.use",
	// 	"compute.disks.create",
	// 	"compute.disks.setLabels",
	// 	"compute.images.get",
	// 	"compute.images.useReadOnly",
	// 	"compute.instanceTemplates.get",
	// 	"compute.instanceTemplates.list",
	// 	"compute.instanceTemplates.create",
	// 	"compute.instanceTemplates.delete",
	// 	"compute.regionOperations.get",
	// 	"compute.zoneOperations.get",
	// 	"compute.zones.list",
	// 	"compute.zones.get",
	// 	"serviceusage.services.list",
	// 	"resourcemanager.projects.getIamPolicy",
	// 	"compute.targetPools.get",
	// 	"compute.targetPools.addInstance",
	// 	"compute.targetPools.removeInstance",
	// 	"compute.instances.use",
	// }
	// componentResource.DefaultCastaiRolePermissions = []string{
	// 	"container.clusters.get",
	// 	"container.clusters.update",
	// 	"container.certificateSigningRequests.approve",
	// 	"compute.instances.get",
	// 	"compute.instances.list",
	// 	"compute.instances.create",
	// 	"compute.instances.start",
	// 	"compute.instances.stop",
	// 	"compute.instances.delete",
	// 	"compute.instances.setLabels",
	// 	"compute.instances.setServiceAccount",
	// 	"compute.instances.setMetadata",
	// 	"compute.instances.setTags",
	// 	"compute.instanceGroupManagers.get",
	// 	"compute.instanceGroupManagers.update",
	// 	"compute.instanceGroups.get",
	// 	"compute.networks.use",
	// 	"compute.networks.useExternalIp",
	// 	"compute.subnetworks.get",
	// 	"compute.subnetworks.use",
	// 	"compute.subnetworks.useExternalIp",
	// 	"compute.addresses.use",
	// 	"compute.disks.use",
	// 	"compute.disks.create",
	// 	"compute.disks.setLabels",
	// 	"compute.images.get",
	// 	"compute.images.useReadOnly",
	// 	"compute.instanceTemplates.get",
	// 	"compute.instanceTemplates.list",
	// 	"compute.instanceTemplates.create",
	// 	"compute.instanceTemplates.delete",
	// 	"compute.regionOperations.get",
	// 	"compute.zoneOperations.get",
	// 	"compute.zones.list",
	// 	"compute.zones.get",
	// 	"serviceusage.services.list",
	// 	"resourcemanager.projects.getIamPolicy",
	// 	"compute.targetPools.get",
	// 	"compute.targetPools.addInstance",
	// 	"compute.targetPools.removeInstance",
	// 	"compute.instances.use",
	// }
	// return &componentResource, nil

	return &componentResource, nil
}
