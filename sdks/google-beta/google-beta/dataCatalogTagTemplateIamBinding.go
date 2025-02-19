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

type DataCatalogTagTemplateIamBinding struct {
	pulumi.CustomResourceState

	Condition                          DataCatalogTagTemplateIamBindingConditionPtrOutput `pulumi:"condition"`
	DataCatalogTagTemplateIamBindingId pulumi.StringOutput                                `pulumi:"dataCatalogTagTemplateIamBindingId"`
	Etag                               pulumi.StringOutput                                `pulumi:"etag"`
	Members                            pulumi.StringArrayOutput                           `pulumi:"members"`
	Project                            pulumi.StringOutput                                `pulumi:"project"`
	Region                             pulumi.StringOutput                                `pulumi:"region"`
	Role                               pulumi.StringOutput                                `pulumi:"role"`
	TagTemplate                        pulumi.StringOutput                                `pulumi:"tagTemplate"`
}

// NewDataCatalogTagTemplateIamBinding registers a new resource with the given unique name, arguments, and options.
func NewDataCatalogTagTemplateIamBinding(ctx *pulumi.Context,
	name string, args *DataCatalogTagTemplateIamBindingArgs, opts ...pulumi.ResourceOption) (*DataCatalogTagTemplateIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TagTemplate == nil {
		return nil, errors.New("invalid value for required argument 'TagTemplate'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataCatalogTagTemplateIamBinding
	err = ctx.RegisterPackageResource("google-beta:index/dataCatalogTagTemplateIamBinding:DataCatalogTagTemplateIamBinding", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataCatalogTagTemplateIamBinding gets an existing DataCatalogTagTemplateIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataCatalogTagTemplateIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataCatalogTagTemplateIamBindingState, opts ...pulumi.ResourceOption) (*DataCatalogTagTemplateIamBinding, error) {
	var resource DataCatalogTagTemplateIamBinding
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataCatalogTagTemplateIamBinding:DataCatalogTagTemplateIamBinding", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataCatalogTagTemplateIamBinding resources.
type dataCatalogTagTemplateIamBindingState struct {
	Condition                          *DataCatalogTagTemplateIamBindingCondition `pulumi:"condition"`
	DataCatalogTagTemplateIamBindingId *string                                    `pulumi:"dataCatalogTagTemplateIamBindingId"`
	Etag                               *string                                    `pulumi:"etag"`
	Members                            []string                                   `pulumi:"members"`
	Project                            *string                                    `pulumi:"project"`
	Region                             *string                                    `pulumi:"region"`
	Role                               *string                                    `pulumi:"role"`
	TagTemplate                        *string                                    `pulumi:"tagTemplate"`
}

type DataCatalogTagTemplateIamBindingState struct {
	Condition                          DataCatalogTagTemplateIamBindingConditionPtrInput
	DataCatalogTagTemplateIamBindingId pulumi.StringPtrInput
	Etag                               pulumi.StringPtrInput
	Members                            pulumi.StringArrayInput
	Project                            pulumi.StringPtrInput
	Region                             pulumi.StringPtrInput
	Role                               pulumi.StringPtrInput
	TagTemplate                        pulumi.StringPtrInput
}

func (DataCatalogTagTemplateIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTagTemplateIamBindingState)(nil)).Elem()
}

type dataCatalogTagTemplateIamBindingArgs struct {
	Condition                          *DataCatalogTagTemplateIamBindingCondition `pulumi:"condition"`
	DataCatalogTagTemplateIamBindingId *string                                    `pulumi:"dataCatalogTagTemplateIamBindingId"`
	Members                            []string                                   `pulumi:"members"`
	Project                            *string                                    `pulumi:"project"`
	Region                             *string                                    `pulumi:"region"`
	Role                               string                                     `pulumi:"role"`
	TagTemplate                        string                                     `pulumi:"tagTemplate"`
}

// The set of arguments for constructing a DataCatalogTagTemplateIamBinding resource.
type DataCatalogTagTemplateIamBindingArgs struct {
	Condition                          DataCatalogTagTemplateIamBindingConditionPtrInput
	DataCatalogTagTemplateIamBindingId pulumi.StringPtrInput
	Members                            pulumi.StringArrayInput
	Project                            pulumi.StringPtrInput
	Region                             pulumi.StringPtrInput
	Role                               pulumi.StringInput
	TagTemplate                        pulumi.StringInput
}

func (DataCatalogTagTemplateIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataCatalogTagTemplateIamBindingArgs)(nil)).Elem()
}

type DataCatalogTagTemplateIamBindingInput interface {
	pulumi.Input

	ToDataCatalogTagTemplateIamBindingOutput() DataCatalogTagTemplateIamBindingOutput
	ToDataCatalogTagTemplateIamBindingOutputWithContext(ctx context.Context) DataCatalogTagTemplateIamBindingOutput
}

func (*DataCatalogTagTemplateIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTagTemplateIamBinding)(nil)).Elem()
}

func (i *DataCatalogTagTemplateIamBinding) ToDataCatalogTagTemplateIamBindingOutput() DataCatalogTagTemplateIamBindingOutput {
	return i.ToDataCatalogTagTemplateIamBindingOutputWithContext(context.Background())
}

func (i *DataCatalogTagTemplateIamBinding) ToDataCatalogTagTemplateIamBindingOutputWithContext(ctx context.Context) DataCatalogTagTemplateIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCatalogTagTemplateIamBindingOutput)
}

type DataCatalogTagTemplateIamBindingOutput struct{ *pulumi.OutputState }

func (DataCatalogTagTemplateIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCatalogTagTemplateIamBinding)(nil)).Elem()
}

func (o DataCatalogTagTemplateIamBindingOutput) ToDataCatalogTagTemplateIamBindingOutput() DataCatalogTagTemplateIamBindingOutput {
	return o
}

func (o DataCatalogTagTemplateIamBindingOutput) ToDataCatalogTagTemplateIamBindingOutputWithContext(ctx context.Context) DataCatalogTagTemplateIamBindingOutput {
	return o
}

func (o DataCatalogTagTemplateIamBindingOutput) Condition() DataCatalogTagTemplateIamBindingConditionPtrOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) DataCatalogTagTemplateIamBindingConditionPtrOutput {
		return v.Condition
	}).(DataCatalogTagTemplateIamBindingConditionPtrOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) DataCatalogTagTemplateIamBindingId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput {
		return v.DataCatalogTagTemplateIamBindingId
	}).(pulumi.StringOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringArrayOutput { return v.Members }).(pulumi.StringArrayOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput { return v.Role }).(pulumi.StringOutput)
}

func (o DataCatalogTagTemplateIamBindingOutput) TagTemplate() pulumi.StringOutput {
	return o.ApplyT(func(v *DataCatalogTagTemplateIamBinding) pulumi.StringOutput { return v.TagTemplate }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataCatalogTagTemplateIamBindingInput)(nil)).Elem(), &DataCatalogTagTemplateIamBinding{})
	pulumi.RegisterOutputType(DataCatalogTagTemplateIamBindingOutput{})
}
