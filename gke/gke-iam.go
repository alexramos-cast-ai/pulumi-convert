package gke

import (
	"fmt"

	"crypto/sha1"
	"encoding/hex"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// func notImplemented(message string) pulumi.AnyOutput {
//   panic(message)
// }

type GkeIamArgs struct {
	ProjectId                pulumi.StringInput
	GkeClusterName           pulumi.StringInput
	ServiceAccountsUniqueIds pulumi.StringInput
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
	PrivateKey                       pulumi.StringOutput
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
	fmt.Println("Beginning GKE IAM")

	// GKE IAM ROLE PREPERATION
	// Create clustername hash
	clusterNameHash := sha1.New()
	clusterNameHash.Write([]byte(fmt.Sprintf("%s", args.GkeClusterName)))
	clusterHashString := hex.EncodeToString(clusterNameHash.Sum(nil))

	// Set Custom Role ID
	clusterCustomRoleId := pulumi.String(fmt.Sprintf("cluster.castai.gkeAccess.%v.tf", clusterHashString[:8]))

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
	fmt.Println("Creating IAM ROLE")
	_, err = projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s_castai_role", name), &projects.IAMCustomRoleArgs{
		RoleId:      pulumi.String(clusterCustomRoleId),
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

	computeRoleId := pulumi.String(fmt.Sprintf("compute.castai.gkeAccess.%v.tf", clusterHashString[:8]))
	computeManagerRole, err := projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s_compute_manager_role", name), &projects.IAMCustomRoleArgs{
		Project:     args.ProjectId,
		RoleId:      computeRoleId,
		Title:       pulumi.String("Role to manage GKE compute resources via CAST AI"),
		Description: pulumi.String("Role to manage GKE compute resources via CAST AI"),
		Permissions: castaiComputeRolePermission,
		Stage:       pulumi.String("GA"),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	// ÇREATE SERVICE ACCOUNT
	serviceAccountHash := sha1.New()
	serviceAccountHash.Write([]byte(fmt.Sprintf("%s", args.GkeClusterName)))
	serviceAccountHashString := hex.EncodeToString(serviceAccountHash.Sum(nil))

	serviceAccountId := pulumi.String(fmt.Sprintf("castai-gke-%v", serviceAccountHashString[:8]))
	serviceAccountEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountId, args.ProjectId)

	serviceAccountRes, err := serviceaccount.NewAccount(ctx, fmt.Sprintf("%s-castai_service_account", name), &serviceaccount.AccountArgs{
		AccountId:   pulumi.String(serviceAccountId),
		DisplayName: pulumi.Sprintf("Service account to manage %v cluster via CAST", args.GkeClusterName),
		Project:     args.ProjectId,
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	fmt.Println("Creating ServiceAccount key")
	serviceAccountKeyRes, err := serviceaccount.NewKey(ctx, fmt.Sprintf("%s-castai_key", name), &serviceaccount.KeyArgs{
		// ServiceAccountId: pulumi.Sprintf("projects/%s/serviceAccounts/%s", args.ProjectId, serviceAccountRes.AccountId),
		ServiceAccountId: serviceAccountRes.Name,
		PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
	}, pulumi.Parent(&componentResource), pulumi.DependsOn([]pulumi.Resource{
		serviceAccountRes,
	}))
	if err != nil {
		return nil, err
	}

	decoded := std.Base64decodeOutput(ctx, std.Base64decodeOutputArgs{
		Input: serviceAccountKeyRes.PrivateKey,
	}, nil).ApplyT(func(invoke std.Base64decodeResult) (string, error) {
		return invoke.Result, nil
	}).(pulumi.StringOutput)

	// TODO: Customer provided service account?
	// var tmp4 map[string]string
	// if args.CreateServiceAccount && len(args.ServiceAccountsUniqueIds) == 0 {
	// 	tmp4 = map[string]string("TODO: For expression")
	// } else {
	// 	tmp4 = map[string]interface{}{}
	// }
	// var project []*projects.IAMMember
	// for key0, _ := range tmp4 {
	// 	__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-project-%v", name, key0), &projects.IAMMemberArgs{
	// 		Project: args.ProjectId,
	// 		Role:    pulumi.String(key0),
	// 		Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
	// 	}, pulumi.Parent(&componentResource))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	project = append(project, __res)
	// }
	// var tmp5 map[string]string
	// if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
	// 	tmp5 = map[string]string("TODO: For expression")
	// } else {
	// 	tmp5 = map[string]interface{}{}
	// }
	// var scopedProject []*projects.IAMMember
	// for key0, _ := range tmp5 {

	_, err = projects.NewIAMBinding(ctx, fmt.Sprintf("%s-scoped_project", name), &projects.IAMBindingArgs{
		Project: args.ProjectId,
		Role:    pulumi.Sprintf("projects/%s/roles/%s", args.ProjectId, clusterCustomRoleId),
		Members: pulumi.StringArray{
			pulumi.Sprintf("serviceAccount:%s", serviceAccountEmail),
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	// scopedProject = append(scopedProject, __res)
	// var tmp6 float64
	// if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
	// 	tmp6 = 1
	// } else {
	// 	tmp6 = 0
	// }
	// var scopedServiceAccountUser []*projects.IAMMember
	// for index := 0; index < tmp6; index++ {
	// 	key0 := index
	// 	_ := index
	_, err = projects.NewIAMMember(ctx, fmt.Sprintf("%s-scoped_service_account_user", name), &projects.IAMMemberArgs{
		Project: args.ProjectId,
		Role:    pulumi.String("roles/iam.serviceAccountUser"),
		Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
		Condition: &projects.IAMMemberConditionArgs{
			Title:       pulumi.String("iam_condition"),
			Description: pulumi.String("IAM condition with limited scope"),
			Expression:  pulumi.Sprintf("resource.name.startsWith(\"projects/-/serviceAccounts/%s\")", args.ServiceAccountsUniqueIds),
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	// scopedServiceAccountUser = append(scopedServiceAccountUser, __res)

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

	// fmt.Println("Binding role")
	_, err = projects.NewIAMBinding(ctx, fmt.Sprintf("%s-compute_manager_binding", name), &projects.IAMBindingArgs{
		Project: args.ProjectId,
		Role:    computeManagerRole.Name,
		Members: pulumi.StringArray{
			pulumi.Sprintf("serviceAccount:%s", serviceAccountEmail), // TODO: Replace with the output of service account email
		},
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}

	componentResource.PrivateKey = decoded

	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"privateKey":          decoded,
		"serviceAccountId":    serviceAccountRes.AccountId,
		"serviceAccountEmail": serviceAccountRes.Email,
	})
	if err != nil {
		return nil, err
	}
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
