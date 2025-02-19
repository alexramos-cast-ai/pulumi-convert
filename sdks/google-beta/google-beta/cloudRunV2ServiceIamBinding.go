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

type CloudRunV2ServiceIamBinding struct {
	pulumi.CustomResourceState

	CloudRunV2ServiceIamBindingId pulumi.StringOutput                           `pulumi:"cloudRunV2ServiceIamBindingId"`
	Condition                     CloudRunV2ServiceIamBindingConditionPtrOutput `pulumi:"condition"`
	Etag                          pulumi.StringOutput                           `pulumi:"etag"`
	Location                      pulumi.StringOutput                           `pulumi:"location"`
	Members                       pulumi.StringArrayOutput                      `pulumi:"members"`
	Name                          pulumi.StringOutput                           `pulumi:"name"`
	Project                       pulumi.StringOutput                           `pulumi:"project"`
	Role                          pulumi.StringOutput                           `pulumi:"role"`
}

// NewCloudRunV2ServiceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewCloudRunV2ServiceIamBinding(ctx *pulumi.Context,
	name string, args *CloudRunV2ServiceIamBindingArgs, opts ...pulumi.ResourceOption) (*CloudRunV2ServiceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource CloudRunV2ServiceIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/cloudRunV2ServiceIamBinding:CloudRunV2ServiceIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudRunV2ServiceIamBinding gets an existing CloudRunV2ServiceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudRunV2ServiceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudRunV2ServiceIamBindingState, opts ...pulumi.ResourceOption) (*CloudRunV2ServiceIamBinding, error) {
	var resource CloudRunV2ServiceIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/cloudRunV2ServiceIamBinding:CloudRunV2ServiceIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudRunV2ServiceIamBinding resources.
type cloudRunV2ServiceIamBindingState struct {
	CloudRunV2ServiceIamBindingId *string                               `pulumi:"cloudRunV2ServiceIamBindingId"`
	Condition                     *CloudRunV2ServiceIamBindingCondition `pulumi:"condition"`
	Etag                          *string                               `pulumi:"etag"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	Name                          *string                               `pulumi:"name"`
	Project                       *string                               `pulumi:"project"`
	Role                          *string                               `pulumi:"role"`
}

type CloudRunV2ServiceIamBindingState struct {
	CloudRunV2ServiceIamBindingId pulumi.StringPtrInput
	Condition                     CloudRunV2ServiceIamBindingConditionPtrInput
	Etag                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	Name                          pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringPtrInput
}

func (CloudRunV2ServiceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunV2ServiceIamBindingState)(nil)).Elem()
}

type cloudRunV2ServiceIamBindingArgs struct {
	CloudRunV2ServiceIamBindingId *string                               `pulumi:"cloudRunV2ServiceIamBindingId"`
	Condition                     *CloudRunV2ServiceIamBindingCondition `pulumi:"condition"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	Name                          *string                               `pulumi:"name"`
	Project                       *string                               `pulumi:"project"`
	Role                          string                                `pulumi:"role"`
}

// The set of arguments for constructing a CloudRunV2ServiceIamBinding resource.
type CloudRunV2ServiceIamBindingArgs struct {
	CloudRunV2ServiceIamBindingId pulumi.StringPtrInput
	Condition                     CloudRunV2ServiceIamBindingConditionPtrInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	Name                          pulumi.StringPtrInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringInput
}

func (CloudRunV2ServiceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudRunV2ServiceIamBindingArgs)(nil)).Elem()
}

type CloudRunV2ServiceIamBindingInput interface {
	pulumi.Input

	ToCloudRunV2ServiceIamBindingOutput() CloudRunV2ServiceIamBindingOutput
	ToCloudRunV2ServiceIamBindingOutputWithContext(ctx context.Context) CloudRunV2ServiceIamBindingOutput
}

func (*CloudRunV2ServiceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunV2ServiceIamBinding)(nil)).Elem()
}

func (i *CloudRunV2ServiceIamBinding) ToCloudRunV2ServiceIamBindingOutput() CloudRunV2ServiceIamBindingOutput {
	return i.ToCloudRunV2ServiceIamBindingOutputWithContext(context.Background())
}

func (i *CloudRunV2ServiceIamBinding) ToCloudRunV2ServiceIamBindingOutputWithContext(ctx context.Context) CloudRunV2ServiceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudRunV2ServiceIamBindingOutput)
}

type CloudRunV2ServiceIamBindingOutput struct{ *pulumi.OutputState }

func (CloudRunV2ServiceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudRunV2ServiceIamBinding)(nil)).Elem()
}

func (o CloudRunV2ServiceIamBindingOutput) ToCloudRunV2ServiceIamBindingOutput() CloudRunV2ServiceIamBindingOutput {
	return o
}

func (o CloudRunV2ServiceIamBindingOutput) ToCloudRunV2ServiceIamBindingOutputWithContext(ctx context.Context) CloudRunV2ServiceIamBindingOutput {
	return o
}

func (o CloudRunV2ServiceIamBindingOutput) CloudRunV2ServiceIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.CloudRunV2ServiceIamBindingId }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Condition() CloudRunV2ServiceIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) CloudRunV2ServiceIamBindingConditionPtrOutput { return v.Condition }).(CloudRunV2ServiceIamBindingConditionPtrOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o CloudRunV2ServiceIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudRunV2ServiceIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudRunV2ServiceIamBindingInput)(nil)).Elem(), &CloudRunV2ServiceIamBinding{})
	pulumi.RegisterOutputType(CloudRunV2ServiceIamBindingOutput{})
}
