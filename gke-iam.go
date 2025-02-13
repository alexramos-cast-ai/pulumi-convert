package main

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/projects"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi-std/sdk/go/std"
	castai "github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// func notImplemented(message string) pulumi.AnyOutput {
//   panic(message)
// }

type GkeIamArgs struct {
	ProjectId                         pulumi.StringInput
	GkeClusterName                    pulumi.StringInput
	ServiceAccountsUniqueIds          []pulumi.StringInput
	ComputeManagerProjectIds          []pulumi.StringInput
	CreateServiceAccount              pulumi.BoolInput
	SetupCloudProxyWorkloadIdentity   pulumi.BoolInput
	WorkloadIdentityNamespace         pulumi.StringInput
	CloudProxyServiceAccountNamespace pulumi.StringInput
	CloudProxyServiceAccountName      pulumi.StringInput
	CastaiRolePermissions             []pulumi.StringInput
	ComputeManagerPermissions         []pulumi.StringInput
}

type GkeIam struct {
	pulumi.ResourceState
	PrivateKey                       pulumi.AnyOutput
	ServiceAccountId                 pulumi.AnyOutput
	ServiceAccountEmail              pulumi.AnyOutput
	DefaultComputeManagerPermissions pulumi.AnyOutput
	DefaultCastaiRolePermissions     pulumi.AnyOutput
}

func NewGkeIam(ctx *pulumi.Context, name string, args *GkeIamArgs, opts ...pulumi.ResourceOption) (*GkeIam, error) {
	var componentResource GkeIam
	err := ctx.RegisterComponentResource("components:index:GkeIam", name, &componentResource, opts...)
	if err != nil {
		return nil, err
	}
	customRoleId := std.SubstrOutput(ctx, std.SubstrOutputArgs{
		Input: std.Sha1Output(ctx, std.Sha1OutputArgs{
			Input: args.GkeClusterName,
		}, nil).ApplyT(func(invoke std.Sha1Result) (*string, error) {
			return invoke.Result, nil
		}).(pulumi.StringPtrOutput),
		Length: pulumi.Int(0),
		Offset: pulumi.Int(8),
	}, nil).ApplyT(func(invoke std.SubstrResult) (string, error) {
		return fmt.Sprintf("castai.gkeAccess.%v.tf", invoke.Result), nil
	}).(pulumi.StringOutput)
	conditionExpression := std.JoinOutput(ctx, std.JoinOutputArgs{
		Separator: pulumi.String("||"),
		Input: []pulumi.String(
			("formatlist(\"resource.name.startsWith(\\\"projects/-/serviceAccounts/%s\\\")\",var.service_accounts_unique_ids)")),
	}, nil).ApplyT(func(invoke std.JoinResult) (*string, error) {
		return invoke.Result, nil
	}).(pulumi.StringPtrOutput)
	defaultPermissions := []interface{}{
		"roles/iam.serviceAccountUser",
		fmt.Sprintf("projects/%v/roles/%v", args.ProjectId, customRoleId),
	}
	scopedPermissions := pulumi.StringArray{
		fmt.Sprintf("projects/%v/roles/%v", args.ProjectId, customRoleId),
	}
	_ := args.ComputeManagerProjectIds
	_, err = castai.GkeUserPolicies(ctx, map[string]interface{}{}, nil)
	if err != nil {
		return err
	}
	var tmp0 pulumi.StringArray
	if float64(len(args.CastaiRolePermissions)) > 0 {
		tmp0 = args.CastaiRolePermissions
	} else {
		tmp0 = []pulumi.String(notImplemented("toset(data.castai_gke_user_policies.gke.policy)"))
	}
	_, err = projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s-castai_role", name), &projects.IAMCustomRoleArgs{
		RoleId:      pulumi.String(customRoleId),
		Title:       pulumi.String("Role to manage GKE cluster via CAST AI"),
		Description: pulumi.String("Role to manage GKE cluster via CAST AI"),
		Permissions: tmp0,
		Project:     args.ProjectId,
		Stage:       pulumi.String("GA"),
	}, pulumi.Parent(&componentResource))
	if err != nil {
		return nil, err
	}
	var tmp1 pulumi.StringArray
	if float64(len(args.ComputeManagerPermissions)) > 0 {
		tmp1 = args.ComputeManagerPermissions
	} else {
		tmp1 = []pulumi.String(notImplemented("toset(data.castai_gke_user_policies.gke.policy)"))
	}
	var computeManagerRole []*projects.IAMCustomRole
	for index := 0; index < notImplemented("toset(local.compute_manager_project_ids)"); index++ {
		key0 := index
		_ := index
		__res, err := projects.NewIAMCustomRole(ctx, fmt.Sprintf("%s-compute_manager_role-%v", name, key0), &projects.IAMCustomRoleArgs{
			Project: pulumi.Any(key0),
			RoleId: std.SubstrOutput(ctx, std.SubstrOutputArgs{
				Input: std.Sha1Output(ctx, std.Sha1OutputArgs{
					Input: pulumi.Any(key0),
				}, nil).ApplyT(func(invoke std.Sha1Result) (*string, error) {
					return invoke.Result, nil
				}).(pulumi.StringPtrOutput),
				Length: pulumi.Int(0),
				Offset: pulumi.Int(8),
			}, nil).ApplyT(func(invoke std.SubstrResult) (string, error) {
				return fmt.Sprintf("castai.gkeAccess.%v.tf", invoke.Result), nil
			}).(pulumi.StringOutput),
			Title:       pulumi.String("Role to manage GKE compute resources via CAST AI"),
			Description: pulumi.String("Role to manage GKE compute resources via CAST AI"),
			Permissions: tmp1,
			Stage:       pulumi.String("GA"),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		computeManagerRole = append(computeManagerRole, __res)
	}
	var computeManagerBinding []*projects.IAMBinding
	for index := 0; index < notImplemented("toset(local.compute_manager_project_ids)"); index++ {
		key0 := index
		_ := index
		__res, err := projects.NewIAMBinding(ctx, fmt.Sprintf("%s-compute_manager_binding-%v", name, key0), &projects.IAMBindingArgs{
			Project: pulumi.Any(key0),
			Role: std.SubstrOutput(ctx, std.SubstrOutputArgs{
				Input: std.Sha1Output(ctx, std.Sha1OutputArgs{
					Input: pulumi.Any(key0),
				}, nil).ApplyT(func(invoke std.Sha1Result) (*string, error) {
					return invoke.Result, nil
				}).(pulumi.StringPtrOutput),
				Length: pulumi.Int(0),
				Offset: pulumi.Int(8),
			}, nil).ApplyT(func(invoke std.SubstrResult) (string, error) {
				return fmt.Sprintf("projects/%v/roles/castai.gkeAccess.%v.tf", key0, invoke.Result), nil
			}).(pulumi.StringOutput),
			Members: []pulumi.String(notImplemented("compact([\"serviceAccount:${local.service_account_email}\",var.setup_cloud_proxy_workload_identity?local.workload_identity_sa:null])")),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		computeManagerBinding = append(computeManagerBinding, __res)
	}
	serviceAccountId := std.SubstrOutput(ctx, std.SubstrOutputArgs{
		Input: std.Sha1Output(ctx, std.Sha1OutputArgs{
			Input: args.GkeClusterName,
		}, nil).ApplyT(func(invoke std.Sha1Result) (*string, error) {
			return invoke.Result, nil
		}).(pulumi.StringPtrOutput),
		Length: pulumi.Int(0),
		Offset: pulumi.Int(8),
	}, nil).ApplyT(func(invoke std.SubstrResult) (string, error) {
		return fmt.Sprintf("castai-gke-tf-%v", invoke.Result), nil
	}).(pulumi.StringOutput)
	serviceAccountEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountId, args.ProjectId)
	var tmp2 float64
	if args.CreateServiceAccount {
		tmp2 = 1
	} else {
		tmp2 = 0
	}
	var castaiServiceAccount []*serviceaccount.Account
	for index := 0; index < tmp2; index++ {
		key0 := index
		_ := index
		__res, err := serviceaccount.NewAccount(ctx, fmt.Sprintf("%s-castai_service_account-%v", name, key0), &serviceaccount.AccountArgs{
			AccountId:   pulumi.String(serviceAccountId),
			DisplayName: pulumi.Sprintf("Service account to manage %v cluster via CAST", args.GkeClusterName),
			Project:     args.ProjectId,
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		castaiServiceAccount = append(castaiServiceAccount, __res)
	}
	var tmp3 float64
	if args.CreateServiceAccount {
		tmp3 = 1
	} else {
		tmp3 = 0
	}
	var castaiKey []*serviceaccount.Key
	for index := 0; index < tmp3; index++ {
		key0 := index
		_ := index
		__res, err := serviceaccount.NewKey(ctx, fmt.Sprintf("%s-castai_key-%v", name, key0), &serviceaccount.KeyArgs{
			ServiceAccountId: castaiServiceAccount[0].Name,
			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		castaiKey = append(castaiKey, __res)
	}
	var tmp4 map[string]string
	if args.CreateServiceAccount && len(args.ServiceAccountsUniqueIds) == 0 {
		tmp4 = map[string]string("TODO: For expression")
	} else {
		tmp4 = map[string]interface{}{}
	}
	var project []*projects.IAMMember
	for key0, _ := range tmp4 {
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-project-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String(key0),
			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		project = append(project, __res)
	}
	var tmp5 map[string]string
	if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
		tmp5 = map[string]string("TODO: For expression")
	} else {
		tmp5 = map[string]interface{}{}
	}
	var scopedProject []*projects.IAMMember
	for key0, _ := range tmp5 {
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-scoped_project-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String(key0),
			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		scopedProject = append(scopedProject, __res)
	}
	var tmp6 float64
	if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
		tmp6 = 1
	} else {
		tmp6 = 0
	}
	var scopedServiceAccountUser []*projects.IAMMember
	for index := 0; index < tmp6; index++ {
		key0 := index
		_ := index
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-scoped_service_account_user-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String("roles/iam.serviceAccountUser"),
			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
			Condition: &projects.IAMMemberConditionArgs{
				Title:       pulumi.String("iam_condition"),
				Description: pulumi.String("IAM condition with limited scope"),
				Expression:  pulumi.String(conditionExpression),
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		scopedServiceAccountUser = append(scopedServiceAccountUser, __res)
	}
	var tmp7 string
	if args.WorkloadIdentityNamespace != "" {
		tmp7 = args.WorkloadIdentityNamespace
	} else {
		tmp7 = fmt.Sprintf("%v.svc.id.goog", args.ProjectId)
	}
	myWorkloadIdentityNamespace := tmp7
	workloadIdentitySa := fmt.Sprintf("serviceAccount:%v[%v/%v]", myWorkloadIdentityNamespace, args.CloudProxyServiceAccountNamespace, args.CloudProxyServiceAccountName)
	var tmp8 map[string]string
	if args.SetupCloudProxyWorkloadIdentity && len(args.ServiceAccountsUniqueIds) == 0 {
		tmp8 = map[string]string("TODO: For expression")
	} else {
		tmp8 = map[string]interface{}{}
	}
	var workloadIdentityProject []*projects.IAMMember
	for key0, _ := range tmp8 {
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_project-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String(key0),
			Member:  pulumi.String(workloadIdentitySa),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		workloadIdentityProject = append(workloadIdentityProject, __res)
	}
	var tmp9 map[string]string
	if args.SetupCloudProxyWorkloadIdentity && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
		tmp9 = map[string]string("TODO: For expression")
	} else {
		tmp9 = map[string]interface{}{}
	}
	var workloadIdentityScopedProject []*projects.IAMMember
	for key0, _ := range tmp9 {
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_scoped_project-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String(key0),
			Member:  pulumi.String(workloadIdentitySa),
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		workloadIdentityScopedProject = append(workloadIdentityScopedProject, __res)
	}
	var tmp10 float64
	if args.SetupCloudProxyWorkloadIdentity && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
		tmp10 = 1
	} else {
		tmp10 = 0
	}
	var workloadIdentityScopedServiceAccountUser []*projects.IAMMember
	for index := 0; index < tmp10; index++ {
		key0 := index
		_ := index
		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_scoped_service_account_user-%v", name, key0), &projects.IAMMemberArgs{
			Project: args.ProjectId,
			Role:    pulumi.String("roles/iam.serviceAccountUser"),
			Member:  pulumi.String(workloadIdentitySa),
			Condition: &projects.IAMMemberConditionArgs{
				Title:       pulumi.String("iam_condition"),
				Description: pulumi.String("IAM condition with limited scope"),
				Expression:  pulumi.String(conditionExpression),
			},
		}, pulumi.Parent(&componentResource))
		if err != nil {
			return nil, err
		}
		workloadIdentityScopedServiceAccountUser = append(workloadIdentityScopedServiceAccountUser, __res)
	}
	err = ctx.RegisterResourceOutputs(&componentResource, pulumi.Map{
		"privateKey":          gkeClientConfig.AccessToken,
		"serviceAccountId":    gkeClientConfig.AccountId,
		"serviceAccountEmail": gkeClientConfig.ServiceAccountEmail,
		"defaultComputeManagerPermissions": []string{
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
		},
		"defaultCastaiRolePermissions": []string{
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
		},
	})
	if err != nil {
		return nil, err
	}
	componentResource.PrivateKey = gkeClientConfig.AccessToken
	componentResource.ServiceAccountId = gkeClientConfig.AccountId
	componentResource.ServiceAccountEmail = gkeClientCOnfig.ServiceAccountEmail
	componentResource.DefaultComputeManagerPermissions = []string{
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
	componentResource.DefaultCastaiRolePermissions = []string{
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
	return &componentResource, nil
}
