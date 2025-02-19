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

type GkeHubNamespace struct {
	pulumi.CustomResourceState

	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Time the Namespace was deleted in UTC.
	DeleteTime        pulumi.StringOutput    `pulumi:"deleteTime"`
	EffectiveLabels   pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	GkeHubNamespaceId pulumi.StringOutput    `pulumi:"gkeHubNamespaceId"`
	// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name for the namespace
	Name pulumi.StringOutput `pulumi:"name"`
	// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
	// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
	// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
	NamespaceLabels pulumi.StringMapOutput `pulumi:"namespaceLabels"`
	Project         pulumi.StringOutput    `pulumi:"project"`
	// The name of the Scope instance.
	Scope pulumi.StringOutput `pulumi:"scope"`
	// Id of the scope
	ScopeId pulumi.StringOutput `pulumi:"scopeId"`
	// The client-provided identifier of the namespace.
	ScopeNamespaceId pulumi.StringOutput `pulumi:"scopeNamespaceId"`
	// State of the namespace resource.
	States GkeHubNamespaceStateTypeArrayOutput `pulumi:"states"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput           `pulumi:"terraformLabels"`
	Timeouts        GkeHubNamespaceTimeoutsPtrOutput `pulumi:"timeouts"`
	// Google-generated UUID for this resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGkeHubNamespace registers a new resource with the given unique name, arguments, and options.
func NewGkeHubNamespace(ctx *pulumi.Context,
	name string, args *GkeHubNamespaceArgs, opts ...pulumi.ResourceOption) (*GkeHubNamespace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.ScopeId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeId'")
	}
	if args.ScopeNamespaceId == nil {
		return nil, errors.New("invalid value for required argument 'ScopeNamespaceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource GkeHubNamespace
	err = ctx.RegisterPackageResource("google-beta:index/gkeHubNamespace:GkeHubNamespace", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGkeHubNamespace gets an existing GkeHubNamespace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGkeHubNamespace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GkeHubNamespaceState, opts ...pulumi.ResourceOption) (*GkeHubNamespace, error) {
	var resource GkeHubNamespace
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/gkeHubNamespace:GkeHubNamespace", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GkeHubNamespace resources.
type gkeHubNamespaceState struct {
	// Time the Namespace was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// Time the Namespace was deleted in UTC.
	DeleteTime        *string           `pulumi:"deleteTime"`
	EffectiveLabels   map[string]string `pulumi:"effectiveLabels"`
	GkeHubNamespaceId *string           `pulumi:"gkeHubNamespaceId"`
	// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The resource name for the namespace
	Name *string `pulumi:"name"`
	// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
	// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
	// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
	NamespaceLabels map[string]string `pulumi:"namespaceLabels"`
	Project         *string           `pulumi:"project"`
	// The name of the Scope instance.
	Scope *string `pulumi:"scope"`
	// Id of the scope
	ScopeId *string `pulumi:"scopeId"`
	// The client-provided identifier of the namespace.
	ScopeNamespaceId *string `pulumi:"scopeNamespaceId"`
	// State of the namespace resource.
	States []GkeHubNamespaceStateType `pulumi:"states"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string        `pulumi:"terraformLabels"`
	Timeouts        *GkeHubNamespaceTimeouts `pulumi:"timeouts"`
	// Google-generated UUID for this resource.
	Uid *string `pulumi:"uid"`
	// Time the Namespace was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type GkeHubNamespaceState struct {
	// Time the Namespace was created in UTC.
	CreateTime pulumi.StringPtrInput
	// Time the Namespace was deleted in UTC.
	DeleteTime        pulumi.StringPtrInput
	EffectiveLabels   pulumi.StringMapInput
	GkeHubNamespaceId pulumi.StringPtrInput
	// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The resource name for the namespace
	Name pulumi.StringPtrInput
	// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
	// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
	// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
	NamespaceLabels pulumi.StringMapInput
	Project         pulumi.StringPtrInput
	// The name of the Scope instance.
	Scope pulumi.StringPtrInput
	// Id of the scope
	ScopeId pulumi.StringPtrInput
	// The client-provided identifier of the namespace.
	ScopeNamespaceId pulumi.StringPtrInput
	// State of the namespace resource.
	States GkeHubNamespaceStateTypeArrayInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        GkeHubNamespaceTimeoutsPtrInput
	// Google-generated UUID for this resource.
	Uid pulumi.StringPtrInput
	// Time the Namespace was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (GkeHubNamespaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubNamespaceState)(nil)).Elem()
}

type gkeHubNamespaceArgs struct {
	GkeHubNamespaceId *string `pulumi:"gkeHubNamespaceId"`
	// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
	// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
	// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
	NamespaceLabels map[string]string `pulumi:"namespaceLabels"`
	Project         *string           `pulumi:"project"`
	// The name of the Scope instance.
	Scope string `pulumi:"scope"`
	// Id of the scope
	ScopeId string `pulumi:"scopeId"`
	// The client-provided identifier of the namespace.
	ScopeNamespaceId string                   `pulumi:"scopeNamespaceId"`
	Timeouts         *GkeHubNamespaceTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a GkeHubNamespace resource.
type GkeHubNamespaceArgs struct {
	GkeHubNamespaceId pulumi.StringPtrInput
	// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
	// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
	// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
	NamespaceLabels pulumi.StringMapInput
	Project         pulumi.StringPtrInput
	// The name of the Scope instance.
	Scope pulumi.StringInput
	// Id of the scope
	ScopeId pulumi.StringInput
	// The client-provided identifier of the namespace.
	ScopeNamespaceId pulumi.StringInput
	Timeouts         GkeHubNamespaceTimeoutsPtrInput
}

func (GkeHubNamespaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gkeHubNamespaceArgs)(nil)).Elem()
}

type GkeHubNamespaceInput interface {
	pulumi.Input

	ToGkeHubNamespaceOutput() GkeHubNamespaceOutput
	ToGkeHubNamespaceOutputWithContext(ctx context.Context) GkeHubNamespaceOutput
}

func (*GkeHubNamespace) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubNamespace)(nil)).Elem()
}

func (i *GkeHubNamespace) ToGkeHubNamespaceOutput() GkeHubNamespaceOutput {
	return i.ToGkeHubNamespaceOutputWithContext(context.Background())
}

func (i *GkeHubNamespace) ToGkeHubNamespaceOutputWithContext(ctx context.Context) GkeHubNamespaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GkeHubNamespaceOutput)
}

type GkeHubNamespaceOutput struct{ *pulumi.OutputState }

func (GkeHubNamespaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GkeHubNamespace)(nil)).Elem()
}

func (o GkeHubNamespaceOutput) ToGkeHubNamespaceOutput() GkeHubNamespaceOutput {
	return o
}

func (o GkeHubNamespaceOutput) ToGkeHubNamespaceOutputWithContext(ctx context.Context) GkeHubNamespaceOutput {
	return o
}

// Time the Namespace was created in UTC.
func (o GkeHubNamespaceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Time the Namespace was deleted in UTC.
func (o GkeHubNamespaceOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

func (o GkeHubNamespaceOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o GkeHubNamespaceOutput) GkeHubNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.GkeHubNamespaceId }).(pulumi.StringOutput)
}

// Labels for this Namespace. **Note**: This field is non-authoritative, and will only manage the labels present in your
// configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o GkeHubNamespaceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name for the namespace
func (o GkeHubNamespaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Namespace-level cluster namespace labels. These labels are applied to the related namespace of the member clusters bound
// to the parent Scope. Scope-level labels ('namespace_labels' in the Fleet Scope resource) take precedence over
// Namespace-level labels if they share a key. Keys and values must be Kubernetes-conformant.
func (o GkeHubNamespaceOutput) NamespaceLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringMapOutput { return v.NamespaceLabels }).(pulumi.StringMapOutput)
}

func (o GkeHubNamespaceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the Scope instance.
func (o GkeHubNamespaceOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.Scope }).(pulumi.StringOutput)
}

// Id of the scope
func (o GkeHubNamespaceOutput) ScopeId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.ScopeId }).(pulumi.StringOutput)
}

// The client-provided identifier of the namespace.
func (o GkeHubNamespaceOutput) ScopeNamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.ScopeNamespaceId }).(pulumi.StringOutput)
}

// State of the namespace resource.
func (o GkeHubNamespaceOutput) States() GkeHubNamespaceStateTypeArrayOutput {
	return o.ApplyT(func(v *GkeHubNamespace) GkeHubNamespaceStateTypeArrayOutput { return v.States }).(GkeHubNamespaceStateTypeArrayOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o GkeHubNamespaceOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o GkeHubNamespaceOutput) Timeouts() GkeHubNamespaceTimeoutsPtrOutput {
	return o.ApplyT(func(v *GkeHubNamespace) GkeHubNamespaceTimeoutsPtrOutput { return v.Timeouts }).(GkeHubNamespaceTimeoutsPtrOutput)
}

// Google-generated UUID for this resource.
func (o GkeHubNamespaceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Time the Namespace was updated in UTC.
func (o GkeHubNamespaceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GkeHubNamespace) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GkeHubNamespaceInput)(nil)).Elem(), &GkeHubNamespace{})
	pulumi.RegisterOutputType(GkeHubNamespaceOutput{})
}
