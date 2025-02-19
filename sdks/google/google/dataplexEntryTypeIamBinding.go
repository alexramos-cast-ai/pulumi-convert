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

type DataplexEntryTypeIamBinding struct {
	pulumi.CustomResourceState

	Condition                     DataplexEntryTypeIamBindingConditionPtrOutput `pulumi:"condition"`
	DataplexEntryTypeIamBindingId pulumi.StringOutput                           `pulumi:"dataplexEntryTypeIamBindingId"`
	EntryTypeId                   pulumi.StringOutput                           `pulumi:"entryTypeId"`
	Etag                          pulumi.StringOutput                           `pulumi:"etag"`
	Location                      pulumi.StringOutput                           `pulumi:"location"`
	Members                       pulumi.StringArrayOutput                      `pulumi:"members"`
	Project                       pulumi.StringOutput                           `pulumi:"project"`
	Role                          pulumi.StringOutput                           `pulumi:"role"`
}

// NewDataplexEntryTypeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataplexEntryTypeIamBinding(ctx *pulumi.Context,
	name string, args *DataplexEntryTypeIamBindingArgs, opts ...pulumi.ResourceOption) (*DataplexEntryTypeIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EntryTypeId == nil {
		return nil, errors.New("invalid value for required argument 'EntryTypeId'")
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
	var resource DataplexEntryTypeIamBinding
	err = ctx.RegisterPackageResource("google:index/dataplexEntryTypeIamBinding:DataplexEntryTypeIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexEntryTypeIamBinding gets an existing DataplexEntryTypeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexEntryTypeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexEntryTypeIamBindingState, opts ...pulumi.ResourceOption) (*DataplexEntryTypeIamBinding, error) {
	var resource DataplexEntryTypeIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexEntryTypeIamBinding:DataplexEntryTypeIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexEntryTypeIamBinding resources.
type dataplexEntryTypeIamBindingState struct {
	Condition                     *DataplexEntryTypeIamBindingCondition `pulumi:"condition"`
	DataplexEntryTypeIamBindingId *string                               `pulumi:"dataplexEntryTypeIamBindingId"`
	EntryTypeId                   *string                               `pulumi:"entryTypeId"`
	Etag                          *string                               `pulumi:"etag"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	Project                       *string                               `pulumi:"project"`
	Role                          *string                               `pulumi:"role"`
}

type DataplexEntryTypeIamBindingState struct {
	Condition                     DataplexEntryTypeIamBindingConditionPtrInput
	DataplexEntryTypeIamBindingId pulumi.StringPtrInput
	EntryTypeId                   pulumi.StringPtrInput
	Etag                          pulumi.StringPtrInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringPtrInput
}

func (DataplexEntryTypeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryTypeIamBindingState)(nil)).Elem()
}

type dataplexEntryTypeIamBindingArgs struct {
	Condition                     *DataplexEntryTypeIamBindingCondition `pulumi:"condition"`
	DataplexEntryTypeIamBindingId *string                               `pulumi:"dataplexEntryTypeIamBindingId"`
	EntryTypeId                   string                                `pulumi:"entryTypeId"`
	Location                      *string                               `pulumi:"location"`
	Members                       []string                              `pulumi:"members"`
	Project                       *string                               `pulumi:"project"`
	Role                          string                                `pulumi:"role"`
}

// The set of arguments for constructing a DataplexEntryTypeIamBinding resource.
type DataplexEntryTypeIamBindingArgs struct {
	Condition                     DataplexEntryTypeIamBindingConditionPtrInput
	DataplexEntryTypeIamBindingId pulumi.StringPtrInput
	EntryTypeId                   pulumi.StringInput
	Location                      pulumi.StringPtrInput
	Members                       pulumi.StringArrayInput
	Project                       pulumi.StringPtrInput
	Role                          pulumi.StringInput
}

func (DataplexEntryTypeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexEntryTypeIamBindingArgs)(nil)).Elem()
}

type DataplexEntryTypeIamBindingInput interface {
	pulumi.Input

	ToDataplexEntryTypeIamBindingOutput() DataplexEntryTypeIamBindingOutput
	ToDataplexEntryTypeIamBindingOutputWithContext(ctx context.Context) DataplexEntryTypeIamBindingOutput
}

func (*DataplexEntryTypeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryTypeIamBinding)(nil)).Elem()
}

func (i *DataplexEntryTypeIamBinding) ToDataplexEntryTypeIamBindingOutput() DataplexEntryTypeIamBindingOutput {
	return i.ToDataplexEntryTypeIamBindingOutputWithContext(context.Background())
}

func (i *DataplexEntryTypeIamBinding) ToDataplexEntryTypeIamBindingOutputWithContext(ctx context.Context) DataplexEntryTypeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexEntryTypeIamBindingOutput)
}

type DataplexEntryTypeIamBindingOutput struct{ *pulumi.OutputState }

func (DataplexEntryTypeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexEntryTypeIamBinding)(nil)).Elem()
}

func (o DataplexEntryTypeIamBindingOutput) ToDataplexEntryTypeIamBindingOutput() DataplexEntryTypeIamBindingOutput {
	return o
}

func (o DataplexEntryTypeIamBindingOutput) ToDataplexEntryTypeIamBindingOutputWithContext(ctx context.Context) DataplexEntryTypeIamBindingOutput {
	return o
}

func (o DataplexEntryTypeIamBindingOutput) Condition() DataplexEntryTypeIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) DataplexEntryTypeIamBindingConditionPtrOutput { return v.Condition }).(DataplexEntryTypeIamBindingConditionPtrOutput)
}

func (o DataplexEntryTypeIamBindingOutput) DataplexEntryTypeIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.DataplexEntryTypeIamBindingId }).(pulumi.StringOutput)
}

func (o DataplexEntryTypeIamBindingOutput) EntryTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.EntryTypeId }).(pulumi.StringOutput)
}

func (o DataplexEntryTypeIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataplexEntryTypeIamBindingOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o DataplexEntryTypeIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o DataplexEntryTypeIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataplexEntryTypeIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexEntryTypeIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexEntryTypeIamBindingInput)(nil)).Elem(), &DataplexEntryTypeIamBinding{})
	pulumi.RegisterOutputType(DataplexEntryTypeIamBindingOutput{})
}
