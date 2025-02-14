package gke

// import {

// }

// TODO: Support Workload identity
// var tmp7 string
// if args.WorkloadIdentityNamespace != "" {
// 	tmp7 = args.WorkloadIdentityNamespace
// } else {
// 	tmp7 = fmt.Sprintf("%v.svc.id.goog", args.ProjectId)
// }
// myWorkloadIdentityNamespace := tmp7
// workloadIdentitySa := fmt.Sprintf("serviceAccount:%v[%v/%v]", myWorkloadIdentityNamespace, args.CloudProxyServiceAccountNamespace, args.CloudProxyServiceAccountName)
// var tmp8 map[string]string
// if args.SetupCloudProxyWorkloadIdentity && len(args.ServiceAccountsUniqueIds) == 0 {
// 	tmp8 = map[string]string("TODO: For expression")
// } else {
// 	tmp8 = map[string]interface{}{}
// }
// var workloadIdentityProject []*projects.IAMMember
// for key0, _ := range tmp8 {
// 	__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_project-%v", name, key0), &projects.IAMMemberArgs{
// 		Project: args.ProjectId,
// 		Role:    pulumi.String(key0),
// 		Member:  pulumi.String(workloadIdentitySa),
// 	}, pulumi.Parent(&componentResource))
// 	if err != nil {
// 		return nil, err
// 	}
// 	workloadIdentityProject = append(workloadIdentityProject, __res)
// }
// var tmp9 map[string]string
// if args.SetupCloudProxyWorkloadIdentity && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
// 	tmp9 = map[string]string("TODO: For expression")
// } else {
// 	tmp9 = map[string]interface{}{}
// }
// var workloadIdentityScopedProject []*projects.IAMMember
// for key0, _ := range tmp9 {
// 	__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_scoped_project-%v", name, key0), &projects.IAMMemberArgs{
// 		Project: args.ProjectId,
// 		Role:    pulumi.String(key0),
// 		Member:  pulumi.String(workloadIdentitySa),
// 	}, pulumi.Parent(&componentResource))
// 	if err != nil {
// 		return nil, err
// 	}
// 	workloadIdentityScopedProject = append(workloadIdentityScopedProject, __res)
// }
// var tmp10 float64
// if args.SetupCloudProxyWorkloadIdentity && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
// 	tmp10 = 1
// } else {
// 	tmp10 = 0
// }
// var workloadIdentityScopedServiceAccountUser []*projects.IAMMember
// for index := 0; index < tmp10; index++ {
// 	key0 := index
// 	_ := index
// 	__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-workload_identity_scoped_service_account_user-%v", name, key0), &projects.IAMMemberArgs{
// 		Project: args.ProjectId,
// 		Role:    pulumi.String("roles/iam.serviceAccountUser"),
// 		Member:  pulumi.String(workloadIdentitySa),
// 		Condition: &projects.IAMMemberConditionArgs{
// 			Title:       pulumi.String("iam_condition"),
// 			Description: pulumi.String("IAM condition with limited scope"),
// 			Expression:  pulumi.String(conditionExpression),
// 		},
// 	}, pulumi.Parent(&componentResource))
// 	if err != nil {
// 		return nil, err
// 	}
// 	workloadIdentityScopedServiceAccountUser = append(workloadIdentityScopedServiceAccountUser, __res)
// }
