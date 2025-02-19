// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package castai

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/castai/v7/castai/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RoleBindings struct {
	pulumi.CustomResourceState

	// Description of the role binding.
	Description pulumi.StringOutput `pulumi:"description"`
	// Name of role binding.
	Name pulumi.StringOutput `pulumi:"name"`
	// CAST AI organization ID.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	RoleBindingsId pulumi.StringOutput `pulumi:"roleBindingsId"`
	// ID of role from role binding.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// Scope of the role binding.
	Scope    RoleBindingsScopeOutput        `pulumi:"scope"`
	Subjects RoleBindingsSubjectArrayOutput `pulumi:"subjects"`
	Timeouts RoleBindingsTimeoutsPtrOutput  `pulumi:"timeouts"`
}

// NewRoleBindings registers a new resource with the given unique name, arguments, and options.
func NewRoleBindings(ctx *pulumi.Context,
	name string, args *RoleBindingsArgs, opts ...pulumi.ResourceOption) (*RoleBindings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Subjects == nil {
		return nil, errors.New("invalid value for required argument 'Subjects'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource RoleBindings
	err = ctx.RegisterPackageResource("castai:index/roleBindings:RoleBindings", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoleBindings gets an existing RoleBindings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoleBindings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleBindingsState, opts ...pulumi.ResourceOption) (*RoleBindings, error) {
	var resource RoleBindings
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("castai:index/roleBindings:RoleBindings", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RoleBindings resources.
type roleBindingsState struct {
	// Description of the role binding.
	Description *string `pulumi:"description"`
	// Name of role binding.
	Name *string `pulumi:"name"`
	// CAST AI organization ID.
	OrganizationId *string `pulumi:"organizationId"`
	RoleBindingsId *string `pulumi:"roleBindingsId"`
	// ID of role from role binding.
	RoleId *string `pulumi:"roleId"`
	// Scope of the role binding.
	Scope    *RoleBindingsScope    `pulumi:"scope"`
	Subjects []RoleBindingsSubject `pulumi:"subjects"`
	Timeouts *RoleBindingsTimeouts `pulumi:"timeouts"`
}

type RoleBindingsState struct {
	// Description of the role binding.
	Description pulumi.StringPtrInput
	// Name of role binding.
	Name pulumi.StringPtrInput
	// CAST AI organization ID.
	OrganizationId pulumi.StringPtrInput
	RoleBindingsId pulumi.StringPtrInput
	// ID of role from role binding.
	RoleId pulumi.StringPtrInput
	// Scope of the role binding.
	Scope    RoleBindingsScopePtrInput
	Subjects RoleBindingsSubjectArrayInput
	Timeouts RoleBindingsTimeoutsPtrInput
}

func (RoleBindingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingsState)(nil)).Elem()
}

type roleBindingsArgs struct {
	// Description of the role binding.
	Description *string `pulumi:"description"`
	// Name of role binding.
	Name *string `pulumi:"name"`
	// CAST AI organization ID.
	OrganizationId string  `pulumi:"organizationId"`
	RoleBindingsId *string `pulumi:"roleBindingsId"`
	// ID of role from role binding.
	RoleId string `pulumi:"roleId"`
	// Scope of the role binding.
	Scope    RoleBindingsScope     `pulumi:"scope"`
	Subjects []RoleBindingsSubject `pulumi:"subjects"`
	Timeouts *RoleBindingsTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a RoleBindings resource.
type RoleBindingsArgs struct {
	// Description of the role binding.
	Description pulumi.StringPtrInput
	// Name of role binding.
	Name pulumi.StringPtrInput
	// CAST AI organization ID.
	OrganizationId pulumi.StringInput
	RoleBindingsId pulumi.StringPtrInput
	// ID of role from role binding.
	RoleId pulumi.StringInput
	// Scope of the role binding.
	Scope    RoleBindingsScopeInput
	Subjects RoleBindingsSubjectArrayInput
	Timeouts RoleBindingsTimeoutsPtrInput
}

func (RoleBindingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleBindingsArgs)(nil)).Elem()
}

type RoleBindingsInput interface {
	pulumi.Input

	ToRoleBindingsOutput() RoleBindingsOutput
	ToRoleBindingsOutputWithContext(ctx context.Context) RoleBindingsOutput
}

func (*RoleBindings) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindings)(nil)).Elem()
}

func (i *RoleBindings) ToRoleBindingsOutput() RoleBindingsOutput {
	return i.ToRoleBindingsOutputWithContext(context.Background())
}

func (i *RoleBindings) ToRoleBindingsOutputWithContext(ctx context.Context) RoleBindingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleBindingsOutput)
}

type RoleBindingsOutput struct{ *pulumi.OutputState }

func (RoleBindingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RoleBindings)(nil)).Elem()
}

func (o RoleBindingsOutput) ToRoleBindingsOutput() RoleBindingsOutput {
	return o
}

func (o RoleBindingsOutput) ToRoleBindingsOutputWithContext(ctx context.Context) RoleBindingsOutput {
	return o
}

// Description of the role binding.
func (o RoleBindingsOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBindings) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Name of role binding.
func (o RoleBindingsOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBindings) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// CAST AI organization ID.
func (o RoleBindingsOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBindings) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o RoleBindingsOutput) RoleBindingsId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBindings) pulumi.StringOutput { return v.RoleBindingsId }).(pulumi.StringOutput)
}

// ID of role from role binding.
func (o RoleBindingsOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RoleBindings) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// Scope of the role binding.
func (o RoleBindingsOutput) Scope() RoleBindingsScopeOutput {
	return o.ApplyT(func(v *RoleBindings) RoleBindingsScopeOutput { return v.Scope }).(RoleBindingsScopeOutput)
}

func (o RoleBindingsOutput) Subjects() RoleBindingsSubjectArrayOutput {
	return o.ApplyT(func(v *RoleBindings) RoleBindingsSubjectArrayOutput { return v.Subjects }).(RoleBindingsSubjectArrayOutput)
}

func (o RoleBindingsOutput) Timeouts() RoleBindingsTimeoutsPtrOutput {
	return o.ApplyT(func(v *RoleBindings) RoleBindingsTimeoutsPtrOutput { return v.Timeouts }).(RoleBindingsTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleBindingsInput)(nil)).Elem(), &RoleBindings{})
	pulumi.RegisterOutputType(RoleBindingsOutput{})
}
