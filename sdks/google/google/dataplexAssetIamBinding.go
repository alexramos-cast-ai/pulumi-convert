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

type DataplexAssetIamBinding struct {
	pulumi.CustomResourceState

	Asset                     pulumi.StringOutput                       `pulumi:"asset"`
	Condition                 DataplexAssetIamBindingConditionPtrOutput `pulumi:"condition"`
	DataplexAssetIamBindingId pulumi.StringOutput                       `pulumi:"dataplexAssetIamBindingId"`
	DataplexZone              pulumi.StringOutput                       `pulumi:"dataplexZone"`
	Etag                      pulumi.StringOutput                       `pulumi:"etag"`
	Lake                      pulumi.StringOutput                       `pulumi:"lake"`
	Location                  pulumi.StringOutput                       `pulumi:"location"`
	Members                   pulumi.StringArrayOutput                  `pulumi:"members"`
	Project                   pulumi.StringOutput                       `pulumi:"project"`
	Role                      pulumi.StringOutput                       `pulumi:"role"`
}

// NewDataplexAssetIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataplexAssetIamBinding(ctx *pulumi.Context,
	name string, args *DataplexAssetIamBindingArgs, opts ...pulumi.ResourceOption) (*DataplexAssetIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Asset == nil {
		return nil, errors.New("invalid value for required argument 'Asset'")
	}
	if args.DataplexZone == nil {
		return nil, errors.New("invalid value for required argument 'DataplexZone'")
	}
	if args.Lake == nil {
		return nil, errors.New("invalid value for required argument 'Lake'")
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
	var resource DataplexAssetIamBinding
	err = ctx.RegisterPackageResource("google:index/dataplexAssetIamBinding:DataplexAssetIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexAssetIamBinding gets an existing DataplexAssetIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexAssetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexAssetIamBindingState, opts ...pulumi.ResourceOption) (*DataplexAssetIamBinding, error) {
	var resource DataplexAssetIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexAssetIamBinding:DataplexAssetIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexAssetIamBinding resources.
type dataplexAssetIamBindingState struct {
	Asset                     *string                           `pulumi:"asset"`
	Condition                 *DataplexAssetIamBindingCondition `pulumi:"condition"`
	DataplexAssetIamBindingId *string                           `pulumi:"dataplexAssetIamBindingId"`
	DataplexZone              *string                           `pulumi:"dataplexZone"`
	Etag                      *string                           `pulumi:"etag"`
	Lake                      *string                           `pulumi:"lake"`
	Location                  *string                           `pulumi:"location"`
	Members                   []string                          `pulumi:"members"`
	Project                   *string                           `pulumi:"project"`
	Role                      *string                           `pulumi:"role"`
}

type DataplexAssetIamBindingState struct {
	Asset                     pulumi.StringPtrInput
	Condition                 DataplexAssetIamBindingConditionPtrInput
	DataplexAssetIamBindingId pulumi.StringPtrInput
	DataplexZone              pulumi.StringPtrInput
	Etag                      pulumi.StringPtrInput
	Lake                      pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	Members                   pulumi.StringArrayInput
	Project                   pulumi.StringPtrInput
	Role                      pulumi.StringPtrInput
}

func (DataplexAssetIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAssetIamBindingState)(nil)).Elem()
}

type dataplexAssetIamBindingArgs struct {
	Asset                     string                            `pulumi:"asset"`
	Condition                 *DataplexAssetIamBindingCondition `pulumi:"condition"`
	DataplexAssetIamBindingId *string                           `pulumi:"dataplexAssetIamBindingId"`
	DataplexZone              string                            `pulumi:"dataplexZone"`
	Lake                      string                            `pulumi:"lake"`
	Location                  *string                           `pulumi:"location"`
	Members                   []string                          `pulumi:"members"`
	Project                   *string                           `pulumi:"project"`
	Role                      string                            `pulumi:"role"`
}

// The set of arguments for constructing a DataplexAssetIamBinding resource.
type DataplexAssetIamBindingArgs struct {
	Asset                     pulumi.StringInput
	Condition                 DataplexAssetIamBindingConditionPtrInput
	DataplexAssetIamBindingId pulumi.StringPtrInput
	DataplexZone              pulumi.StringInput
	Lake                      pulumi.StringInput
	Location                  pulumi.StringPtrInput
	Members                   pulumi.StringArrayInput
	Project                   pulumi.StringPtrInput
	Role                      pulumi.StringInput
}

func (DataplexAssetIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexAssetIamBindingArgs)(nil)).Elem()
}

type DataplexAssetIamBindingInput interface {
	pulumi.Input

	ToDataplexAssetIamBindingOutput() DataplexAssetIamBindingOutput
	ToDataplexAssetIamBindingOutputWithContext(ctx context.Context) DataplexAssetIamBindingOutput
}

func (*DataplexAssetIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAssetIamBinding)(nil)).Elem()
}

func (i *DataplexAssetIamBinding) ToDataplexAssetIamBindingOutput() DataplexAssetIamBindingOutput {
	return i.ToDataplexAssetIamBindingOutputWithContext(context.Background())
}

func (i *DataplexAssetIamBinding) ToDataplexAssetIamBindingOutputWithContext(ctx context.Context) DataplexAssetIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexAssetIamBindingOutput)
}

type DataplexAssetIamBindingOutput struct{ *pulumi.OutputState }

func (DataplexAssetIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexAssetIamBinding)(nil)).Elem()
}

func (o DataplexAssetIamBindingOutput) ToDataplexAssetIamBindingOutput() DataplexAssetIamBindingOutput {
	return o
}

func (o DataplexAssetIamBindingOutput) ToDataplexAssetIamBindingOutputWithContext(ctx context.Context) DataplexAssetIamBindingOutput {
	return o
}

func (o DataplexAssetIamBindingOutput) Asset() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Asset }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Condition() DataplexAssetIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) DataplexAssetIamBindingConditionPtrOutput { return v.Condition }).(DataplexAssetIamBindingConditionPtrOutput)
}

func (o DataplexAssetIamBindingOutput) DataplexAssetIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.DataplexAssetIamBindingId }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) DataplexZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.DataplexZone }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Lake() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Lake }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o DataplexAssetIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataplexAssetIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexAssetIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexAssetIamBindingInput)(nil)).Elem(), &DataplexAssetIamBinding{})
	pulumi.RegisterOutputType(DataplexAssetIamBindingOutput{})
}
