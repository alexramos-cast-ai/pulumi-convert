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

type DataprocAutoscalingPolicyIamBinding struct {
	pulumi.CustomResourceState

	Condition                             DataprocAutoscalingPolicyIamBindingConditionPtrOutput `pulumi:"condition"`
	DataprocAutoscalingPolicyIamBindingId pulumi.StringOutput                                   `pulumi:"dataprocAutoscalingPolicyIamBindingId"`
	Etag                                  pulumi.StringOutput                                   `pulumi:"etag"`
	Location                              pulumi.StringOutput                                   `pulumi:"location"`
	Members                               pulumi.StringArrayOutput                              `pulumi:"members"`
	PolicyId                              pulumi.StringOutput                                   `pulumi:"policyId"`
	Project                               pulumi.StringOutput                                   `pulumi:"project"`
	Role                                  pulumi.StringOutput                                   `pulumi:"role"`
}

// NewDataprocAutoscalingPolicyIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataprocAutoscalingPolicyIamBinding(ctx *pulumi.Context,
	name string, args *DataprocAutoscalingPolicyIamBindingArgs, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicyIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.PolicyId == nil {
		return nil, errors.New("invalid value for required argument 'PolicyId'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocAutoscalingPolicyIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/dataprocAutoscalingPolicyIamBinding:DataprocAutoscalingPolicyIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocAutoscalingPolicyIamBinding gets an existing DataprocAutoscalingPolicyIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocAutoscalingPolicyIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocAutoscalingPolicyIamBindingState, opts ...pulumi.ResourceOption) (*DataprocAutoscalingPolicyIamBinding, error) {
	var resource DataprocAutoscalingPolicyIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataprocAutoscalingPolicyIamBinding:DataprocAutoscalingPolicyIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocAutoscalingPolicyIamBinding resources.
type dataprocAutoscalingPolicyIamBindingState struct {
	Condition                             *DataprocAutoscalingPolicyIamBindingCondition `pulumi:"condition"`
	DataprocAutoscalingPolicyIamBindingId *string                                       `pulumi:"dataprocAutoscalingPolicyIamBindingId"`
	Etag                                  *string                                       `pulumi:"etag"`
	Location                              *string                                       `pulumi:"location"`
	Members                               []string                                      `pulumi:"members"`
	PolicyId                              *string                                       `pulumi:"policyId"`
	Project                               *string                                       `pulumi:"project"`
	Role                                  *string                                       `pulumi:"role"`
}

type DataprocAutoscalingPolicyIamBindingState struct {
	Condition                             DataprocAutoscalingPolicyIamBindingConditionPtrInput
	DataprocAutoscalingPolicyIamBindingId pulumi.StringPtrInput
	Etag                                  pulumi.StringPtrInput
	Location                              pulumi.StringPtrInput
	Members                               pulumi.StringArrayInput
	PolicyId                              pulumi.StringPtrInput
	Project                               pulumi.StringPtrInput
	Role                                  pulumi.StringPtrInput
}

func (DataprocAutoscalingPolicyIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyIamBindingState)(nil)).Elem()
}

type dataprocAutoscalingPolicyIamBindingArgs struct {
	Condition                             *DataprocAutoscalingPolicyIamBindingCondition `pulumi:"condition"`
	DataprocAutoscalingPolicyIamBindingId *string                                       `pulumi:"dataprocAutoscalingPolicyIamBindingId"`
	Location                              *string                                       `pulumi:"location"`
	Members                               []string                                      `pulumi:"members"`
	PolicyId                              string                                        `pulumi:"policyId"`
	Project                               *string                                       `pulumi:"project"`
	Role                                  string                                        `pulumi:"role"`
}

// The set of arguments for constructing a DataprocAutoscalingPolicyIamBinding resource.
type DataprocAutoscalingPolicyIamBindingArgs struct {
	Condition                             DataprocAutoscalingPolicyIamBindingConditionPtrInput
	DataprocAutoscalingPolicyIamBindingId pulumi.StringPtrInput
	Location                              pulumi.StringPtrInput
	Members                               pulumi.StringArrayInput
	PolicyId                              pulumi.StringInput
	Project                               pulumi.StringPtrInput
	Role                                  pulumi.StringInput
}

func (DataprocAutoscalingPolicyIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocAutoscalingPolicyIamBindingArgs)(nil)).Elem()
}

type DataprocAutoscalingPolicyIamBindingInput interface {
	pulumi.Input

	ToDataprocAutoscalingPolicyIamBindingOutput() DataprocAutoscalingPolicyIamBindingOutput
	ToDataprocAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamBindingOutput
}

func (*DataprocAutoscalingPolicyIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicyIamBinding)(nil)).Elem()
}

func (i *DataprocAutoscalingPolicyIamBinding) ToDataprocAutoscalingPolicyIamBindingOutput() DataprocAutoscalingPolicyIamBindingOutput {
	return i.ToDataprocAutoscalingPolicyIamBindingOutputWithContext(context.Background())
}

func (i *DataprocAutoscalingPolicyIamBinding) ToDataprocAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocAutoscalingPolicyIamBindingOutput)
}

type DataprocAutoscalingPolicyIamBindingOutput struct{ *pulumi.OutputState }

func (DataprocAutoscalingPolicyIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocAutoscalingPolicyIamBinding)(nil)).Elem()
}

func (o DataprocAutoscalingPolicyIamBindingOutput) ToDataprocAutoscalingPolicyIamBindingOutput() DataprocAutoscalingPolicyIamBindingOutput {
	return o
}

func (o DataprocAutoscalingPolicyIamBindingOutput) ToDataprocAutoscalingPolicyIamBindingOutputWithContext(ctx context.Context) DataprocAutoscalingPolicyIamBindingOutput {
	return o
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Condition() DataprocAutoscalingPolicyIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) DataprocAutoscalingPolicyIamBindingConditionPtrOutput {
		return v.Condition
	}).(DataprocAutoscalingPolicyIamBindingConditionPtrOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) DataprocAutoscalingPolicyIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput {
		return v.DataprocAutoscalingPolicyIamBindingId
	}).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) PolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput { return v.PolicyId }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataprocAutoscalingPolicyIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocAutoscalingPolicyIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocAutoscalingPolicyIamBindingInput)(nil)).Elem(), &DataprocAutoscalingPolicyIamBinding{})
	pulumi.RegisterOutputType(DataprocAutoscalingPolicyIamBindingOutput{})
}
