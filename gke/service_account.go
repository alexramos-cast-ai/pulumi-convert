// package gke

// import {}

// 	// ÇREATE SERVICE ACCOUNT
// 	serviceAccountHash := sha1.New()
// 	serviceAccountHash.Write([]byte(fmt.Sprintf("%s", args.GkeClusterName)))
// 	serviceAccountHashString := hex.EncodeToString(serviceAccountHash.Sum(nil))

// 	serviceAccountId := pulumi.String(fmt.Sprintf("castai.gkeAccess.%v.tf", serviceAccountHashString[:8]))
// 	serviceAccountEmail := fmt.Sprintf("%v@%v.iam.gserviceaccount.com", serviceAccountId, args.ProjectId)
// 	var tmp2 float64
// 	if args.CreateServiceAccount {
// 		tmp2 = 1
// 	} else {
// 		tmp2 = 0
// 	}
// 	var castaiServiceAccount []*serviceaccount.Account
// 	for _, values := range castaiServiceAccount {
// 		__res, err := serviceaccount.NewAccount(ctx, fmt.Sprintf("%s-castai_service_account-%v", name, key0), &serviceaccount.AccountArgs{
// 			AccountId:   pulumi.String(serviceAccountId),
// 			DisplayName: pulumi.Sprintf("Service account to manage %v cluster via CAST", args.GkeClusterName),
// 			Project:     args.ProjectId,
// 		}, pulumi.Parent(&componentResource))
// 		if err != nil {
// 			return nil, err
// 		}
// 		castaiServiceAccount = append(castaiServiceAccount, __res)
// 	}
// 	var tmp3 float64
// 	if args.CreateServiceAccount {
// 		tmp3 = 1
// 	} else {
// 		tmp3 = 0
// 	}
// 	var castaiKey []*serviceaccount.Key
// 	for index := 0; index < tmp3; index++ {
// 		key0 := index
// 		_ := index
// 		__res, err := serviceaccount.NewKey(ctx, fmt.Sprintf("%s-castai_key-%v", name, key0), &serviceaccount.KeyArgs{
// 			ServiceAccountId: castaiServiceAccount[0].Name,
// 			PublicKeyType:    pulumi.String("TYPE_X509_PEM_FILE"),
// 		}, pulumi.Parent(&componentResource))
// 		if err != nil {
// 			return nil, err
// 		}
// 		castaiKey = append(castaiKey, __res)
// 	}
// 	var tmp4 map[string]string
// 	if args.CreateServiceAccount && len(args.ServiceAccountsUniqueIds) == 0 {
// 		tmp4 = map[string]string("TODO: For expression")
// 	} else {
// 		tmp4 = map[string]interface{}{}
// 	}
// 	var project []*projects.IAMMember
// 	for key0, _ := range tmp4 {
// 		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-project-%v", name, key0), &projects.IAMMemberArgs{
// 			Project: args.ProjectId,
// 			Role:    pulumi.String(key0),
// 			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
// 		}, pulumi.Parent(&componentResource))
// 		if err != nil {
// 			return nil, err
// 		}
// 		project = append(project, __res)
// 	}
// 	var tmp5 map[string]string
// 	if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
// 		tmp5 = map[string]string("TODO: For expression")
// 	} else {
// 		tmp5 = map[string]interface{}{}
// 	}
// 	var scopedProject []*projects.IAMMember
// 	for key0, _ := range tmp5 {
// 		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-scoped_project-%v", name, key0), &projects.IAMMemberArgs{
// 			Project: args.ProjectId,
// 			Role:    pulumi.String(key0),
// 			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
// 		}, pulumi.Parent(&componentResource))
// 		if err != nil {
// 			return nil, err
// 		}
// 		scopedProject = append(scopedProject, __res)
// 	}
// 	var tmp6 float64
// 	if args.CreateServiceAccount && float64(len(args.ServiceAccountsUniqueIds)) > 0 {
// 		tmp6 = 1
// 	} else {
// 		tmp6 = 0
// 	}
// 	var scopedServiceAccountUser []*projects.IAMMember
// 	for index := 0; index < tmp6; index++ {
// 		key0 := index
// 		_ := index
// 		__res, err := projects.NewIAMMember(ctx, fmt.Sprintf("%s-scoped_service_account_user-%v", name, key0), &projects.IAMMemberArgs{
// 			Project: args.ProjectId,
// 			Role:    pulumi.String("roles/iam.serviceAccountUser"),
// 			Member:  pulumi.Sprintf("serviceAccount:%v", serviceAccountEmail),
// 			Condition: &projects.IAMMemberConditionArgs{
// 				Title:       pulumi.String("iam_condition"),
// 				Description: pulumi.String("IAM condition with limited scope"),
// 				Expression:  pulumi.String(conditionExpression),
// 			},
// 		}, pulumi.Parent(&componentResource))
// 		if err != nil {
// 			return nil, err
// 		}
// 		scopedServiceAccountUser = append(scopedServiceAccountUser, __res)
// 	}