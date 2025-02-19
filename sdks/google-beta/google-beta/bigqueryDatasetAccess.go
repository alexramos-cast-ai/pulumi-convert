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

type BigqueryDatasetAccess struct {
	pulumi.CustomResourceState

	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember        pulumi.BoolOutput   `pulumi:"apiUpdatedMember"`
	BigqueryDatasetAccessId pulumi.StringOutput `pulumi:"bigqueryDatasetAccessId"`
	// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
	Condition BigqueryDatasetAccessConditionPtrOutput `pulumi:"condition"`
	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset BigqueryDatasetAccessDatasetPtrOutput `pulumi:"dataset"`
	// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
	Domain pulumi.StringPtrOutput `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrOutput `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
	// 'allUsers'
	IamMember pulumi.StringPtrOutput `pulumi:"iamMember"`
	Project   pulumi.StringOutput    `pulumi:"project"`
	// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
	// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
	// counterparts, and will show a diff post-create. See [official
	// docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrOutput `pulumi:"role"`
	// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
	// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
	// access to the routine needs to be granted again via an update operation.
	Routine BigqueryDatasetAccessRoutinePtrOutput `pulumi:"routine"`
	// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
	// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup pulumi.StringPtrOutput                 `pulumi:"specialGroup"`
	Timeouts     BigqueryDatasetAccessTimeoutsPtrOutput `pulumi:"timeouts"`
	// An email address of a user to grant access to. For example: fred@example.com
	UserByEmail pulumi.StringPtrOutput `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
	// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
	// the view needs to be granted again via an update operation.
	View BigqueryDatasetAccessViewPtrOutput `pulumi:"view"`
}

// NewBigqueryDatasetAccess registers a new resource with the given unique name, arguments, and options.
func NewBigqueryDatasetAccess(ctx *pulumi.Context,
	name string, args *BigqueryDatasetAccessArgs, opts ...pulumi.ResourceOption) (*BigqueryDatasetAccess, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryDatasetAccess
	err = ctx.RegisterPackageResource("google-beta:index/bigqueryDatasetAccess:BigqueryDatasetAccess", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryDatasetAccess gets an existing BigqueryDatasetAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryDatasetAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryDatasetAccessState, opts ...pulumi.ResourceOption) (*BigqueryDatasetAccess, error) {
	var resource BigqueryDatasetAccess
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/bigqueryDatasetAccess:BigqueryDatasetAccess", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryDatasetAccess resources.
type bigqueryDatasetAccessState struct {
	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember        *bool   `pulumi:"apiUpdatedMember"`
	BigqueryDatasetAccessId *string `pulumi:"bigqueryDatasetAccessId"`
	// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
	Condition *BigqueryDatasetAccessCondition `pulumi:"condition"`
	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset *BigqueryDatasetAccessDataset `pulumi:"dataset"`
	// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetId *string `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
	Domain *string `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail *string `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
	// 'allUsers'
	IamMember *string `pulumi:"iamMember"`
	Project   *string `pulumi:"project"`
	// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
	// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
	// counterparts, and will show a diff post-create. See [official
	// docs](https://cloud.google.com/bigquery/docs/access-control).
	Role *string `pulumi:"role"`
	// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
	// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
	// access to the routine needs to be granted again via an update operation.
	Routine *BigqueryDatasetAccessRoutine `pulumi:"routine"`
	// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
	// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup *string                        `pulumi:"specialGroup"`
	Timeouts     *BigqueryDatasetAccessTimeouts `pulumi:"timeouts"`
	// An email address of a user to grant access to. For example: fred@example.com
	UserByEmail *string `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
	// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
	// the view needs to be granted again via an update operation.
	View *BigqueryDatasetAccessView `pulumi:"view"`
}

type BigqueryDatasetAccessState struct {
	// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
	// stored in state as a different member type
	ApiUpdatedMember        pulumi.BoolPtrInput
	BigqueryDatasetAccessId pulumi.StringPtrInput
	// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
	Condition BigqueryDatasetAccessConditionPtrInput
	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset BigqueryDatasetAccessDatasetPtrInput
	// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetId pulumi.StringPtrInput
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
	Domain pulumi.StringPtrInput
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrInput
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
	// 'allUsers'
	IamMember pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
	// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
	// counterparts, and will show a diff post-create. See [official
	// docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrInput
	// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
	// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
	// access to the routine needs to be granted again via an update operation.
	Routine BigqueryDatasetAccessRoutinePtrInput
	// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
	// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup pulumi.StringPtrInput
	Timeouts     BigqueryDatasetAccessTimeoutsPtrInput
	// An email address of a user to grant access to. For example: fred@example.com
	UserByEmail pulumi.StringPtrInput
	// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
	// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
	// the view needs to be granted again via an update operation.
	View BigqueryDatasetAccessViewPtrInput
}

func (BigqueryDatasetAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatasetAccessState)(nil)).Elem()
}

type bigqueryDatasetAccessArgs struct {
	BigqueryDatasetAccessId *string `pulumi:"bigqueryDatasetAccessId"`
	// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
	Condition *BigqueryDatasetAccessCondition `pulumi:"condition"`
	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset *BigqueryDatasetAccessDataset `pulumi:"dataset"`
	// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetId string `pulumi:"datasetId"`
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
	Domain *string `pulumi:"domain"`
	// An email address of a Google Group to grant access to.
	GroupByEmail *string `pulumi:"groupByEmail"`
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
	// 'allUsers'
	IamMember *string `pulumi:"iamMember"`
	Project   *string `pulumi:"project"`
	// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
	// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
	// counterparts, and will show a diff post-create. See [official
	// docs](https://cloud.google.com/bigquery/docs/access-control).
	Role *string `pulumi:"role"`
	// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
	// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
	// access to the routine needs to be granted again via an update operation.
	Routine *BigqueryDatasetAccessRoutine `pulumi:"routine"`
	// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
	// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup *string                        `pulumi:"specialGroup"`
	Timeouts     *BigqueryDatasetAccessTimeouts `pulumi:"timeouts"`
	// An email address of a user to grant access to. For example: fred@example.com
	UserByEmail *string `pulumi:"userByEmail"`
	// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
	// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
	// the view needs to be granted again via an update operation.
	View *BigqueryDatasetAccessView `pulumi:"view"`
}

// The set of arguments for constructing a BigqueryDatasetAccess resource.
type BigqueryDatasetAccessArgs struct {
	BigqueryDatasetAccessId pulumi.StringPtrInput
	// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
	Condition BigqueryDatasetAccessConditionPtrInput
	// Grants all resources of particular types in a particular dataset read access to the current dataset.
	Dataset BigqueryDatasetAccessDatasetPtrInput
	// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
	// underscores (_). The maximum length is 1,024 characters.
	DatasetId pulumi.StringInput
	// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
	Domain pulumi.StringPtrInput
	// An email address of a Google Group to grant access to.
	GroupByEmail pulumi.StringPtrInput
	// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
	// 'allUsers'
	IamMember pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
	// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
	// counterparts, and will show a diff post-create. See [official
	// docs](https://cloud.google.com/bigquery/docs/access-control).
	Role pulumi.StringPtrInput
	// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
	// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
	// access to the routine needs to be granted again via an update operation.
	Routine BigqueryDatasetAccessRoutinePtrInput
	// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
	// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
	// 'allAuthenticatedUsers': All authenticated BigQuery users.
	SpecialGroup pulumi.StringPtrInput
	Timeouts     BigqueryDatasetAccessTimeoutsPtrInput
	// An email address of a user to grant access to. For example: fred@example.com
	UserByEmail pulumi.StringPtrInput
	// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
	// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
	// the view needs to be granted again via an update operation.
	View BigqueryDatasetAccessViewPtrInput
}

func (BigqueryDatasetAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryDatasetAccessArgs)(nil)).Elem()
}

type BigqueryDatasetAccessInput interface {
	pulumi.Input

	ToBigqueryDatasetAccessOutput() BigqueryDatasetAccessOutput
	ToBigqueryDatasetAccessOutputWithContext(ctx context.Context) BigqueryDatasetAccessOutput
}

func (*BigqueryDatasetAccess) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatasetAccess)(nil)).Elem()
}

func (i *BigqueryDatasetAccess) ToBigqueryDatasetAccessOutput() BigqueryDatasetAccessOutput {
	return i.ToBigqueryDatasetAccessOutputWithContext(context.Background())
}

func (i *BigqueryDatasetAccess) ToBigqueryDatasetAccessOutputWithContext(ctx context.Context) BigqueryDatasetAccessOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryDatasetAccessOutput)
}

type BigqueryDatasetAccessOutput struct{ *pulumi.OutputState }

func (BigqueryDatasetAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryDatasetAccess)(nil)).Elem()
}

func (o BigqueryDatasetAccessOutput) ToBigqueryDatasetAccessOutput() BigqueryDatasetAccessOutput {
	return o
}

func (o BigqueryDatasetAccessOutput) ToBigqueryDatasetAccessOutputWithContext(ctx context.Context) BigqueryDatasetAccessOutput {
	return o
}

// If true, represents that that the iam_member in the config was translated to a different member type by the API, and is
// stored in state as a different member type
func (o BigqueryDatasetAccessOutput) ApiUpdatedMember() pulumi.BoolOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.BoolOutput { return v.ApiUpdatedMember }).(pulumi.BoolOutput)
}

func (o BigqueryDatasetAccessOutput) BigqueryDatasetAccessId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringOutput { return v.BigqueryDatasetAccessId }).(pulumi.StringOutput)
}

// Condition for the binding. If CEL expression in this field is true, this access binding will be considered.
func (o BigqueryDatasetAccessOutput) Condition() BigqueryDatasetAccessConditionPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) BigqueryDatasetAccessConditionPtrOutput { return v.Condition }).(BigqueryDatasetAccessConditionPtrOutput)
}

// Grants all resources of particular types in a particular dataset read access to the current dataset.
func (o BigqueryDatasetAccessOutput) Dataset() BigqueryDatasetAccessDatasetPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) BigqueryDatasetAccessDatasetPtrOutput { return v.Dataset }).(BigqueryDatasetAccessDatasetPtrOutput)
}

// A unique ID for this dataset, without the project name. The ID must contain only letters (a-z, A-Z), numbers (0-9), or
// underscores (_). The maximum length is 1,024 characters.
func (o BigqueryDatasetAccessOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// A domain to grant access to. Any users signed in with the domain specified will be granted the specified access
func (o BigqueryDatasetAccessOutput) Domain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.Domain }).(pulumi.StringPtrOutput)
}

// An email address of a Google Group to grant access to.
func (o BigqueryDatasetAccessOutput) GroupByEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.GroupByEmail }).(pulumi.StringPtrOutput)
}

// Some other type of member that appears in the IAM Policy but isn't a user, group, domain, or special group. For example:
// 'allUsers'
func (o BigqueryDatasetAccessOutput) IamMember() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.IamMember }).(pulumi.StringPtrOutput)
}

func (o BigqueryDatasetAccessOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Describes the rights granted to the user specified by the other member of the access object. Basic, predefined, and
// custom roles are supported. Predefined roles that have equivalent basic roles are swapped by the API to their basic
// counterparts, and will show a diff post-create. See [official
// docs](https://cloud.google.com/bigquery/docs/access-control).
func (o BigqueryDatasetAccessOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.Role }).(pulumi.StringPtrOutput)
}

// A routine from a different dataset to grant access to. Queries executed against that routine will have read access to
// tables in this dataset. The role field is not required when this field is set. If that routine is updated by any user,
// access to the routine needs to be granted again via an update operation.
func (o BigqueryDatasetAccessOutput) Routine() BigqueryDatasetAccessRoutinePtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) BigqueryDatasetAccessRoutinePtrOutput { return v.Routine }).(BigqueryDatasetAccessRoutinePtrOutput)
}

// A special group to grant access to. Possible values include: * 'projectOwners': Owners of the enclosing project. *
// 'projectReaders': Readers of the enclosing project. * 'projectWriters': Writers of the enclosing project. *
// 'allAuthenticatedUsers': All authenticated BigQuery users.
func (o BigqueryDatasetAccessOutput) SpecialGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.SpecialGroup }).(pulumi.StringPtrOutput)
}

func (o BigqueryDatasetAccessOutput) Timeouts() BigqueryDatasetAccessTimeoutsPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) BigqueryDatasetAccessTimeoutsPtrOutput { return v.Timeouts }).(BigqueryDatasetAccessTimeoutsPtrOutput)
}

// An email address of a user to grant access to. For example: fred@example.com
func (o BigqueryDatasetAccessOutput) UserByEmail() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) pulumi.StringPtrOutput { return v.UserByEmail }).(pulumi.StringPtrOutput)
}

// A view from a different dataset to grant access to. Queries executed against that view will have read access to tables
// in this dataset. The role field is not required when this field is set. If that view is updated by any user, access to
// the view needs to be granted again via an update operation.
func (o BigqueryDatasetAccessOutput) View() BigqueryDatasetAccessViewPtrOutput {
	return o.ApplyT(func(v *BigqueryDatasetAccess) BigqueryDatasetAccessViewPtrOutput { return v.View }).(BigqueryDatasetAccessViewPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryDatasetAccessInput)(nil)).Elem(), &BigqueryDatasetAccess{})
	pulumi.RegisterOutputType(BigqueryDatasetAccessOutput{})
}
