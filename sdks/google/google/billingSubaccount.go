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

type BillingSubaccount struct {
	pulumi.CustomResourceState

	BillingAccountId     pulumi.StringOutput    `pulumi:"billingAccountId"`
	BillingSubaccountId  pulumi.StringOutput    `pulumi:"billingSubaccountId"`
	DeletionPolicy       pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	DisplayName          pulumi.StringOutput    `pulumi:"displayName"`
	MasterBillingAccount pulumi.StringOutput    `pulumi:"masterBillingAccount"`
	Name                 pulumi.StringOutput    `pulumi:"name"`
	Open                 pulumi.BoolOutput      `pulumi:"open"`
}

// NewBillingSubaccount registers a new resource with the given unique name, arguments, and options.
func NewBillingSubaccount(ctx *pulumi.Context,
	name string, args *BillingSubaccountArgs, opts ...pulumi.ResourceOption) (*BillingSubaccount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.MasterBillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'MasterBillingAccount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BillingSubaccount
	err = ctx.RegisterPackageResource("google:index/billingSubaccount:BillingSubaccount", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingSubaccount gets an existing BillingSubaccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingSubaccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingSubaccountState, opts ...pulumi.ResourceOption) (*BillingSubaccount, error) {
	var resource BillingSubaccount
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/billingSubaccount:BillingSubaccount", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingSubaccount resources.
type billingSubaccountState struct {
	BillingAccountId     *string `pulumi:"billingAccountId"`
	BillingSubaccountId  *string `pulumi:"billingSubaccountId"`
	DeletionPolicy       *string `pulumi:"deletionPolicy"`
	DisplayName          *string `pulumi:"displayName"`
	MasterBillingAccount *string `pulumi:"masterBillingAccount"`
	Name                 *string `pulumi:"name"`
	Open                 *bool   `pulumi:"open"`
}

type BillingSubaccountState struct {
	BillingAccountId     pulumi.StringPtrInput
	BillingSubaccountId  pulumi.StringPtrInput
	DeletionPolicy       pulumi.StringPtrInput
	DisplayName          pulumi.StringPtrInput
	MasterBillingAccount pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	Open                 pulumi.BoolPtrInput
}

func (BillingSubaccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingSubaccountState)(nil)).Elem()
}

type billingSubaccountArgs struct {
	BillingSubaccountId  *string `pulumi:"billingSubaccountId"`
	DeletionPolicy       *string `pulumi:"deletionPolicy"`
	DisplayName          string  `pulumi:"displayName"`
	MasterBillingAccount string  `pulumi:"masterBillingAccount"`
}

// The set of arguments for constructing a BillingSubaccount resource.
type BillingSubaccountArgs struct {
	BillingSubaccountId  pulumi.StringPtrInput
	DeletionPolicy       pulumi.StringPtrInput
	DisplayName          pulumi.StringInput
	MasterBillingAccount pulumi.StringInput
}

func (BillingSubaccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingSubaccountArgs)(nil)).Elem()
}

type BillingSubaccountInput interface {
	pulumi.Input

	ToBillingSubaccountOutput() BillingSubaccountOutput
	ToBillingSubaccountOutputWithContext(ctx context.Context) BillingSubaccountOutput
}

func (*BillingSubaccount) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingSubaccount)(nil)).Elem()
}

func (i *BillingSubaccount) ToBillingSubaccountOutput() BillingSubaccountOutput {
	return i.ToBillingSubaccountOutputWithContext(context.Background())
}

func (i *BillingSubaccount) ToBillingSubaccountOutputWithContext(ctx context.Context) BillingSubaccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingSubaccountOutput)
}

type BillingSubaccountOutput struct{ *pulumi.OutputState }

func (BillingSubaccountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingSubaccount)(nil)).Elem()
}

func (o BillingSubaccountOutput) ToBillingSubaccountOutput() BillingSubaccountOutput {
	return o
}

func (o BillingSubaccountOutput) ToBillingSubaccountOutputWithContext(ctx context.Context) BillingSubaccountOutput {
	return o
}

func (o BillingSubaccountOutput) BillingAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringOutput { return v.BillingAccountId }).(pulumi.StringOutput)
}

func (o BillingSubaccountOutput) BillingSubaccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringOutput { return v.BillingSubaccountId }).(pulumi.StringOutput)
}

func (o BillingSubaccountOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

func (o BillingSubaccountOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o BillingSubaccountOutput) MasterBillingAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringOutput { return v.MasterBillingAccount }).(pulumi.StringOutput)
}

func (o BillingSubaccountOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o BillingSubaccountOutput) Open() pulumi.BoolOutput {
	return o.ApplyT(func(v *BillingSubaccount) pulumi.BoolOutput { return v.Open }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingSubaccountInput)(nil)).Elem(), &BillingSubaccount{})
	pulumi.RegisterOutputType(BillingSubaccountOutput{})
}
