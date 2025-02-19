// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package googlebeta

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google-beta/v6/google-beta/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GkeHubFeatureMembership struct {
	pulumi.CustomResourceState

	// Config Management-specific spec.
	Configmanagement GkeHubFeatureMembershipConfigmanagementPtrOutput `pulumi:"configmanagement"`
	// The name of the feature
	Feature                   pulumi.StringOutput `pulumi:"feature"`
	GkeHubFeatureMembershipId pulumi.StringOutput `pulumi:"gkeHubFeatureMembershipId"`
	// The location of the feature
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the membership
	Membership pulumi.StringOutput `pulumi:"membership"`
	// The location of the membership
	MembershipLocation pulumi.StringPtrOutput `pulumi:"membershipLocation"`
	// Manage Mesh Features
	Mesh GkeHubFeatureMembershipMeshPtrOutput `pulumi:"mesh"`
	// Policy Controller-specific spec.
	Policycontroller GkeHubFeatureMembershipPolicycontrollerPtrOutput `pulumi:"policycontroller"`
	// The project of the feature
	Project  pulumi.StringOutput                      `pulumi:"project"`
	Timeouts GkeHubFeatureMembershipTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewGkeHubFeatureMembership registers a new resource with the given unique name, arguments, and options.
func NewGkeHubFeatureMembership(ctx *pulumi.Context,
	name string, args *GkeHubFeatureMembershipArgs, opts ...pulumi.ResourceOption) (*GkeHubFeatureMembership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Feature == nil {
		return nil, errors.New("invalid value for required argument 'Feature'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Membership == nil {
		return nil, errors.New("invalid value for required argument 'Membership'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeHubFeatureMembership
	err = ctx.RegisterPackageResource("google-beta:index/gkeHubFeatureMembership:GkeHubFeatureMembership", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubFeatureMembership gets an existing GkeHubFeatureMembership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubFeatureMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubFeatureMembershipState, opts ...pulumi.ResourceOption) (*GkeHubFeatureMembership, error) {
	var resource GkeHubFeatureMembership
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/gkeHubFeatureMembership:GkeHubFeatureMembership", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubFeatureMembership resources.
type gkeHubFeatureMembershipState struct {
	// Config Management-specific spec.
	Configmanagement *GkeHubFeatureMembershipConfigmanagement `pulumi:"configmanagement"`
	// The name of the feature
	Feature                   *string `pulumi:"feature"`
	GkeHubFeatureMembershipId *string `pulumi:"gkeHubFeatureMembershipId"`
	// The location of the feature
	Location *string `pulumi:"location"`
	// The name of the membership
	Membership *string `pulumi:"membership"`
	// The location of the membership
	MembershipLocation *string `pulumi:"membershipLocation"`
	// Manage Mesh Features
	Mesh *GkeHubFeatureMembershipMesh `pulumi:"mesh"`
	// Policy Controller-specific spec.
	Policycontroller *GkeHubFeatureMembershipPolicycontroller `pulumi:"policycontroller"`
	// The project of the feature
	Project  *string                          `pulumi:"project"`
	Timeouts *GkeHubFeatureMembershipTimeouts `pulumi:"timeouts"`
}

type GkeHubFeatureMembershipState struct {
	// Config Management-specific spec.
	Configmanagement GkeHubFeatureMembershipConfigmanagementPtrInput
	// The name of the feature
	Feature                   pulumi.StringPtrInput
	GkeHubFeatureMembershipId pulumi.StringPtrInput
	// The location of the feature
	Location pulumi.StringPtrInput
	// The name of the membership
	Membership pulumi.StringPtrInput
	// The location of the membership
	MembershipLocation pulumi.StringPtrInput
	// Manage Mesh Features
	Mesh GkeHubFeatureMembershipMeshPtrInput
	// Policy Controller-specific spec.
	Policycontroller GkeHubFeatureMembershipPolicycontrollerPtrInput
	// The project of the feature
	Project  pulumi.StringPtrInput
	Timeouts GkeHubFeatureMembershipTimeoutsPtrInput
}

func (GkeHubFeatureMembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubFeatureMembershipState)(nil)).Elem()
}

type gkeHubFeatureMembershipArgs struct {
	// Config Management-specific spec.
	Configmanagement *GkeHubFeatureMembershipConfigmanagement `pulumi:"configmanagement"`
	// The name of the feature
	Feature                   string  `pulumi:"feature"`
	GkeHubFeatureMembershipId *string `pulumi:"gkeHubFeatureMembershipId"`
	// The location of the feature
	Location string `pulumi:"location"`
	// The name of the membership
	Membership string `pulumi:"membership"`
	// The location of the membership
	MembershipLocation *string `pulumi:"membershipLocation"`
	// Manage Mesh Features
	Mesh *GkeHubFeatureMembershipMesh `pulumi:"mesh"`
	// Policy Controller-specific spec.
	Policycontroller *GkeHubFeatureMembershipPolicycontroller `pulumi:"policycontroller"`
	// The project of the feature
	Project  *string                          `pulumi:"project"`
	Timeouts *GkeHubFeatureMembershipTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a GkeHubFeatureMembership resource.
type GkeHubFeatureMembershipArgs struct {
	// Config Management-specific spec.
	Configmanagement GkeHubFeatureMembershipConfigmanagementPtrInput
	// The name of the feature
	Feature                   pulumi.StringInput
	GkeHubFeatureMembershipId pulumi.StringPtrInput
	// The location of the feature
	Location pulumi.StringInput
	// The name of the membership
	Membership pulumi.StringInput
	// The location of the membership
	MembershipLocation pulumi.StringPtrInput
	// Manage Mesh Features
	Mesh GkeHubFeatureMembershipMeshPtrInput
	// Policy Controller-specific spec.
	Policycontroller GkeHubFeatureMembershipPolicycontrollerPtrInput
	// The project of the feature
	Project  pulumi.StringPtrInput
	Timeouts GkeHubFeatureMembershipTimeoutsPtrInput
}

func (GkeHubFeatureMembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubFeatureMembershipArgs)(nil)).Elem()
}

type GkeHubFeatureMembershipInput interface {
	pulumi.Input

	ToGkeHubFeatureMembershipOutput() GkeHubFeatureMembershipOutput
	ToGkeHubFeatureMembershipOutputWithContext(ctx context.Context) GkeHubFeatureMembershipOutput
}

func (*GkeHubFeatureMembership) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubFeatureMembership)(nil)).Elem()
}

func (i *GkeHubFeatureMembership) ToGkeHubFeatureMembershipOutput() GkeHubFeatureMembershipOutput {
	return i.ToGkeHubFeatureMembershipOutputWithContext(context.Background())
}

func (i *GkeHubFeatureMembership) ToGkeHubFeatureMembershipOutputWithContext(ctx context.Context) GkeHubFeatureMembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubFeatureMembershipOutput)
}

type GkeHubFeatureMembershipOutput struct{ *pulumi.OutputState }

func (GkeHubFeatureMembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubFeatureMembership)(nil)).Elem()
}

func (o GkeHubFeatureMembershipOutput) ToGkeHubFeatureMembershipOutput() GkeHubFeatureMembershipOutput {
	return o
}

func (o GkeHubFeatureMembershipOutput) ToGkeHubFeatureMembershipOutputWithContext(ctx context.Context) GkeHubFeatureMembershipOutput {
	return o
}

// Config Management-specific spec.
func (o GkeHubFeatureMembershipOutput) Configmanagement() GkeHubFeatureMembershipConfigmanagementPtrOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) GkeHubFeatureMembershipConfigmanagementPtrOutput {
		return v.Configmanagement
	}).(GkeHubFeatureMembershipConfigmanagementPtrOutput)
}

// The name of the feature
func (o GkeHubFeatureMembershipOutput) Feature() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringOutput { return v.Feature }).(pulumi.StringOutput)
}

func (o GkeHubFeatureMembershipOutput) GkeHubFeatureMembershipId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringOutput { return v.GkeHubFeatureMembershipId }).(pulumi.StringOutput)
}

// The location of the feature
func (o GkeHubFeatureMembershipOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the membership
func (o GkeHubFeatureMembershipOutput) Membership() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringOutput { return v.Membership }).(pulumi.StringOutput)
}

// The location of the membership
func (o GkeHubFeatureMembershipOutput) MembershipLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringPtrOutput { return v.MembershipLocation }).(pulumi.StringPtrOutput)
}

// Manage Mesh Features
func (o GkeHubFeatureMembershipOutput) Mesh() GkeHubFeatureMembershipMeshPtrOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) GkeHubFeatureMembershipMeshPtrOutput { return v.Mesh }).(GkeHubFeatureMembershipMeshPtrOutput)
}

// Policy Controller-specific spec.
func (o GkeHubFeatureMembershipOutput) Policycontroller() GkeHubFeatureMembershipPolicycontrollerPtrOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) GkeHubFeatureMembershipPolicycontrollerPtrOutput {
		return v.Policycontroller
	}).(GkeHubFeatureMembershipPolicycontrollerPtrOutput)
}

// The project of the feature
func (o GkeHubFeatureMembershipOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o GkeHubFeatureMembershipOutput) Timeouts() GkeHubFeatureMembershipTimeoutsPtrOutput {
	return o.ApplyT(func(v *GkeHubFeatureMembership) GkeHubFeatureMembershipTimeoutsPtrOutput { return v.Timeouts }).(GkeHubFeatureMembershipTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubFeatureMembershipInput)(nil)).Elem(), &GkeHubFeatureMembership{})
	pulumi.RegisterOutputType(GkeHubFeatureMembershipOutput{})
}
