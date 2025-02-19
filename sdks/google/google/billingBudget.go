// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BillingBudget struct {
	pulumi.CustomResourceState

	// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	AllUpdatesRule BillingBudgetAllUpdatesRulePtrOutput `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	Amount BillingBudgetAmountOutput `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount  pulumi.StringOutput `pulumi:"billingAccount"`
	BillingBudgetId pulumi.StringOutput `pulumi:"billingBudgetId"`
	// Filters that define which resources are used to compute the actual spend against the budget.
	BudgetFilter BillingBudgetBudgetFilterPtrOutput `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
	// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	OwnershipScope pulumi.StringPtrOutput `pulumi:"ownershipScope"`
	// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
	// the budget.
	ThresholdRules BillingBudgetThresholdRuleArrayOutput `pulumi:"thresholdRules"`
	Timeouts       BillingBudgetTimeoutsPtrOutput        `pulumi:"timeouts"`
}

// NewBillingBudget registers a new resource with the given unique name, arguments, and options.
func NewBillingBudget(ctx *pulumi.Context,
	name string, args *BillingBudgetArgs, opts ...pulumi.ResourceOption) (*BillingBudget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Amount == nil {
		return nil, errors.New("invalid value for required argument 'Amount'")
	}
	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BillingBudget
	err = ctx.RegisterPackageResource("google:index/billingBudget:BillingBudget", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingBudget gets an existing BillingBudget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingBudgetState, opts ...pulumi.ResourceOption) (*BillingBudget, error) {
	var resource BillingBudget
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/billingBudget:BillingBudget", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingBudget resources.
type billingBudgetState struct {
	// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	AllUpdatesRule *BillingBudgetAllUpdatesRule `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	Amount *BillingBudgetAmount `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount  *string `pulumi:"billingAccount"`
	BillingBudgetId *string `pulumi:"billingBudgetId"`
	// Filters that define which resources are used to compute the actual spend against the budget.
	BudgetFilter *BillingBudgetBudgetFilter `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name *string `pulumi:"name"`
	// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
	// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	OwnershipScope *string `pulumi:"ownershipScope"`
	// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
	// the budget.
	ThresholdRules []BillingBudgetThresholdRule `pulumi:"thresholdRules"`
	Timeouts       *BillingBudgetTimeouts       `pulumi:"timeouts"`
}

type BillingBudgetState struct {
	// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	AllUpdatesRule BillingBudgetAllUpdatesRulePtrInput
	// The budgeted amount for each usage period.
	Amount BillingBudgetAmountPtrInput
	// ID of the billing account to set a budget on.
	BillingAccount  pulumi.StringPtrInput
	BillingBudgetId pulumi.StringPtrInput
	// Filters that define which resources are used to compute the actual spend against the budget.
	BudgetFilter BillingBudgetBudgetFilterPtrInput
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
	// billingAccounts/{billingAccountId}/budgets/{budgetId}.
	Name pulumi.StringPtrInput
	// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
	// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	OwnershipScope pulumi.StringPtrInput
	// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
	// the budget.
	ThresholdRules BillingBudgetThresholdRuleArrayInput
	Timeouts       BillingBudgetTimeoutsPtrInput
}

func (BillingBudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingBudgetState)(nil)).Elem()
}

type billingBudgetArgs struct {
	// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	AllUpdatesRule *BillingBudgetAllUpdatesRule `pulumi:"allUpdatesRule"`
	// The budgeted amount for each usage period.
	Amount BillingBudgetAmount `pulumi:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount  string  `pulumi:"billingAccount"`
	BillingBudgetId *string `pulumi:"billingBudgetId"`
	// Filters that define which resources are used to compute the actual spend against the budget.
	BudgetFilter *BillingBudgetBudgetFilter `pulumi:"budgetFilter"`
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
	// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	OwnershipScope *string `pulumi:"ownershipScope"`
	// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
	// the budget.
	ThresholdRules []BillingBudgetThresholdRule `pulumi:"thresholdRules"`
	Timeouts       *BillingBudgetTimeouts       `pulumi:"timeouts"`
}

// The set of arguments for constructing a BillingBudget resource.
type BillingBudgetArgs struct {
	// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
	// using threshold rules.
	AllUpdatesRule BillingBudgetAllUpdatesRulePtrInput
	// The budgeted amount for each usage period.
	Amount BillingBudgetAmountInput
	// ID of the billing account to set a budget on.
	BillingAccount  pulumi.StringInput
	BillingBudgetId pulumi.StringPtrInput
	// Filters that define which resources are used to compute the actual spend against the budget.
	BudgetFilter BillingBudgetBudgetFilterPtrInput
	// User data for display name in UI. Must be <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
	// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
	OwnershipScope pulumi.StringPtrInput
	// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
	// the budget.
	ThresholdRules BillingBudgetThresholdRuleArrayInput
	Timeouts       BillingBudgetTimeoutsPtrInput
}

func (BillingBudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingBudgetArgs)(nil)).Elem()
}

type BillingBudgetInput interface {
	pulumi.Input

	ToBillingBudgetOutput() BillingBudgetOutput
	ToBillingBudgetOutputWithContext(ctx context.Context) BillingBudgetOutput
}

func (*BillingBudget) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingBudget)(nil)).Elem()
}

func (i *BillingBudget) ToBillingBudgetOutput() BillingBudgetOutput {
	return i.ToBillingBudgetOutputWithContext(context.Background())
}

func (i *BillingBudget) ToBillingBudgetOutputWithContext(ctx context.Context) BillingBudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingBudgetOutput)
}

type BillingBudgetOutput struct{ *pulumi.OutputState }

func (BillingBudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingBudget)(nil)).Elem()
}

func (o BillingBudgetOutput) ToBillingBudgetOutput() BillingBudgetOutput {
	return o
}

func (o BillingBudgetOutput) ToBillingBudgetOutputWithContext(ctx context.Context) BillingBudgetOutput {
	return o
}

// Defines notifications that are sent on every update to the billing account's spend, regardless of the thresholds defined
// using threshold rules.
func (o BillingBudgetOutput) AllUpdatesRule() BillingBudgetAllUpdatesRulePtrOutput {
	return o.ApplyT(func(v *BillingBudget) BillingBudgetAllUpdatesRulePtrOutput { return v.AllUpdatesRule }).(BillingBudgetAllUpdatesRulePtrOutput)
}

// The budgeted amount for each usage period.
func (o BillingBudgetOutput) Amount() BillingBudgetAmountOutput {
	return o.ApplyT(func(v *BillingBudget) BillingBudgetAmountOutput { return v.Amount }).(BillingBudgetAmountOutput)
}

// ID of the billing account to set a budget on.
func (o BillingBudgetOutput) BillingAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingBudget) pulumi.StringOutput { return v.BillingAccount }).(pulumi.StringOutput)
}

func (o BillingBudgetOutput) BillingBudgetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingBudget) pulumi.StringOutput { return v.BillingBudgetId }).(pulumi.StringOutput)
}

// Filters that define which resources are used to compute the actual spend against the budget.
func (o BillingBudgetOutput) BudgetFilter() BillingBudgetBudgetFilterPtrOutput {
	return o.ApplyT(func(v *BillingBudget) BillingBudgetBudgetFilterPtrOutput { return v.BudgetFilter }).(BillingBudgetBudgetFilterPtrOutput)
}

// User data for display name in UI. Must be <= 60 chars.
func (o BillingBudgetOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingBudget) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form
// billingAccounts/{billingAccountId}/budgets/{budgetId}.
func (o BillingBudgetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingBudget) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ownership scope of the budget. The ownership scope and users' IAM permissions determine who has full access to the
// budget's data. Possible values: ["OWNERSHIP_SCOPE_UNSPECIFIED", "ALL_USERS", "BILLING_ACCOUNT"]
func (o BillingBudgetOutput) OwnershipScope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingBudget) pulumi.StringPtrOutput { return v.OwnershipScope }).(pulumi.StringPtrOutput)
}

// Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of
// the budget.
func (o BillingBudgetOutput) ThresholdRules() BillingBudgetThresholdRuleArrayOutput {
	return o.ApplyT(func(v *BillingBudget) BillingBudgetThresholdRuleArrayOutput { return v.ThresholdRules }).(BillingBudgetThresholdRuleArrayOutput)
}

func (o BillingBudgetOutput) Timeouts() BillingBudgetTimeoutsPtrOutput {
	return o.ApplyT(func(v *BillingBudget) BillingBudgetTimeoutsPtrOutput { return v.Timeouts }).(BillingBudgetTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingBudgetInput)(nil)).Elem(), &BillingBudget{})
	pulumi.RegisterOutputType(BillingBudgetOutput{})
}
