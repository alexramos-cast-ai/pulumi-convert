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

type SccV2FolderSccBigQueryExport struct {
	pulumi.CustomResourceState

	// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
	// with a letter, must end with either a letter or a number, and must be 63 characters or less.
	BigQueryExportId pulumi.StringOutput `pulumi:"bigQueryExportId"`
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrOutput `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
	// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	// information on how to write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor pulumi.StringOutput `pulumi:"mostRecentEditor"`
	// The resource name of this export, in the format
	// 'folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in
	// responses, and is ignored when provided in create requests.
	Name pulumi.StringOutput `pulumi:"name"`
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                      pulumi.StringOutput                           `pulumi:"principal"`
	SccV2FolderSccBigQueryExportId pulumi.StringOutput                           `pulumi:"sccV2FolderSccBigQueryExportId"`
	Timeouts                       SccV2FolderSccBigQueryExportTimeoutsPtrOutput `pulumi:"timeouts"`
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewSccV2FolderSccBigQueryExport registers a new resource with the given unique name, arguments, and options.
func NewSccV2FolderSccBigQueryExport(ctx *pulumi.Context,
	name string, args *SccV2FolderSccBigQueryExportArgs, opts ...pulumi.ResourceOption) (*SccV2FolderSccBigQueryExport, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BigQueryExportId == nil {
		return nil, errors.New("invalid value for required argument 'BigQueryExportId'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SccV2FolderSccBigQueryExport
	err = ctx.RegisterPackageResource("google:index/sccV2FolderSccBigQueryExport:SccV2FolderSccBigQueryExport", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSccV2FolderSccBigQueryExport gets an existing SccV2FolderSccBigQueryExport resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSccV2FolderSccBigQueryExport(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SccV2FolderSccBigQueryExportState, opts ...pulumi.ResourceOption) (*SccV2FolderSccBigQueryExport, error) {
	var resource SccV2FolderSccBigQueryExport
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/sccV2FolderSccBigQueryExport:SccV2FolderSccBigQueryExport", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SccV2FolderSccBigQueryExport resources.
type sccV2FolderSccBigQueryExportState struct {
	// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
	// with a letter, must end with either a letter or a number, and must be 63 characters or less.
	BigQueryExportId *string `pulumi:"bigQueryExportId"`
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset *string `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
	// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	// information on how to write a filter.
	Filter *string `pulumi:"filter"`
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	Folder *string `pulumi:"folder"`
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	Location *string `pulumi:"location"`
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor *string `pulumi:"mostRecentEditor"`
	// The resource name of this export, in the format
	// 'folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in
	// responses, and is ignored when provided in create requests.
	Name *string `pulumi:"name"`
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                      *string                               `pulumi:"principal"`
	SccV2FolderSccBigQueryExportId *string                               `pulumi:"sccV2FolderSccBigQueryExportId"`
	Timeouts                       *SccV2FolderSccBigQueryExportTimeouts `pulumi:"timeouts"`
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type SccV2FolderSccBigQueryExportState struct {
	// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
	// with a letter, must end with either a letter or a number, and must be 63 characters or less.
	BigQueryExportId pulumi.StringPtrInput
	// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
	// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
	// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
	// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	// information on how to write a filter.
	Filter pulumi.StringPtrInput
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	Folder pulumi.StringPtrInput
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	Location pulumi.StringPtrInput
	// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
	// provided on export creation or update.
	MostRecentEditor pulumi.StringPtrInput
	// The resource name of this export, in the format
	// 'folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in
	// responses, and is ignored when provided in create requests.
	Name pulumi.StringPtrInput
	// The service account that needs permission to create table and upload data to the BigQuery dataset.
	Principal                      pulumi.StringPtrInput
	SccV2FolderSccBigQueryExportId pulumi.StringPtrInput
	Timeouts                       SccV2FolderSccBigQueryExportTimeoutsPtrInput
	// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
	// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
	// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (SccV2FolderSccBigQueryExportState) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2FolderSccBigQueryExportState)(nil)).Elem()
}

type sccV2FolderSccBigQueryExportArgs struct {
	// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
	// with a letter, must end with either a letter or a number, and must be 63 characters or less.
	BigQueryExportId string `pulumi:"bigQueryExportId"`
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset *string `pulumi:"dataset"`
	// The description of the notification config (max of 1024 characters).
	Description *string `pulumi:"description"`
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
	// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	// information on how to write a filter.
	Filter *string `pulumi:"filter"`
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	Folder string `pulumi:"folder"`
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	Location                       *string                               `pulumi:"location"`
	SccV2FolderSccBigQueryExportId *string                               `pulumi:"sccV2FolderSccBigQueryExportId"`
	Timeouts                       *SccV2FolderSccBigQueryExportTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a SccV2FolderSccBigQueryExport resource.
type SccV2FolderSccBigQueryExportArgs struct {
	// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
	// with a letter, must end with either a letter or a number, and must be 63 characters or less.
	BigQueryExportId pulumi.StringInput
	// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
	// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
	Dataset pulumi.StringPtrInput
	// The description of the notification config (max of 1024 characters).
	Description pulumi.StringPtrInput
	// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
	// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
	// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
	// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
	// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
	// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
	// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
	// information on how to write a filter.
	Filter pulumi.StringPtrInput
	// The folder where Cloud Security Command Center Big Query Export Config lives in.
	Folder pulumi.StringInput
	// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
	Location                       pulumi.StringPtrInput
	SccV2FolderSccBigQueryExportId pulumi.StringPtrInput
	Timeouts                       SccV2FolderSccBigQueryExportTimeoutsPtrInput
}

func (SccV2FolderSccBigQueryExportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*sccV2FolderSccBigQueryExportArgs)(nil)).Elem()
}

type SccV2FolderSccBigQueryExportInput interface {
	pulumi.Input

	ToSccV2FolderSccBigQueryExportOutput() SccV2FolderSccBigQueryExportOutput
	ToSccV2FolderSccBigQueryExportOutputWithContext(ctx context.Context) SccV2FolderSccBigQueryExportOutput
}

func (*SccV2FolderSccBigQueryExport) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2FolderSccBigQueryExport)(nil)).Elem()
}

func (i *SccV2FolderSccBigQueryExport) ToSccV2FolderSccBigQueryExportOutput() SccV2FolderSccBigQueryExportOutput {
	return i.ToSccV2FolderSccBigQueryExportOutputWithContext(context.Background())
}

func (i *SccV2FolderSccBigQueryExport) ToSccV2FolderSccBigQueryExportOutputWithContext(ctx context.Context) SccV2FolderSccBigQueryExportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SccV2FolderSccBigQueryExportOutput)
}

type SccV2FolderSccBigQueryExportOutput struct{ *pulumi.OutputState }

func (SccV2FolderSccBigQueryExportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SccV2FolderSccBigQueryExport)(nil)).Elem()
}

func (o SccV2FolderSccBigQueryExportOutput) ToSccV2FolderSccBigQueryExportOutput() SccV2FolderSccBigQueryExportOutput {
	return o
}

func (o SccV2FolderSccBigQueryExportOutput) ToSccV2FolderSccBigQueryExportOutputWithContext(ctx context.Context) SccV2FolderSccBigQueryExportOutput {
	return o
}

// This must be unique within the organization. It must consist of only lowercase letters, numbers, and hyphens, must start
// with a letter, must end with either a letter or a number, and must be 63 characters or less.
func (o SccV2FolderSccBigQueryExportOutput) BigQueryExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.BigQueryExportId }).(pulumi.StringOutput)
}

// The time at which the BigQuery export was created. This field is set by the server and will be ignored if provided on
// export on creation. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional
// digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccV2FolderSccBigQueryExportOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The dataset to write findings' updates to. Its format is "projects/[projectId]/datasets/[bigquery_dataset_id]". BigQuery
// Dataset unique ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_).
func (o SccV2FolderSccBigQueryExportOutput) Dataset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringPtrOutput { return v.Dataset }).(pulumi.StringPtrOutput)
}

// The description of the notification config (max of 1024 characters).
func (o SccV2FolderSccBigQueryExportOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Expression that defines the filter to apply across create/update events of findings. The expression is a list of zero or
// more restrictions combined via logical operators AND and OR. Parentheses are supported, and OR has higher precedence
// than AND. Restrictions have the form <field> <operator> <value> and may have a - character in front of them to indicate
// negation. The fields map to those defined in the corresponding resource. The supported operators are: * = for all value
// types. * >, <, >=, <= for integer values. * :, meaning substring matching, for strings. The supported value types are: *
// string literals in quotes. * integer literals without quotes. * boolean literals true and false without quotes. See
// [Filtering notifications](https://cloud.google.com/security-command-center/docs/how-to-api-filter-notifications) for
// information on how to write a filter.
func (o SccV2FolderSccBigQueryExportOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

// The folder where Cloud Security Command Center Big Query Export Config lives in.
func (o SccV2FolderSccBigQueryExportOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// The BigQuery export configuration is stored in this location. If not provided, Use global as default.
func (o SccV2FolderSccBigQueryExportOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Email address of the user who last edited the BigQuery export. This field is set by the server and will be ignored if
// provided on export creation or update.
func (o SccV2FolderSccBigQueryExportOutput) MostRecentEditor() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.MostRecentEditor }).(pulumi.StringOutput)
}

// The resource name of this export, in the format
// 'folders/{{folder}}/locations/{{location}}/bigQueryExports/{{big_query_export_id}}'. This field is provided in
// responses, and is ignored when provided in create requests.
func (o SccV2FolderSccBigQueryExportOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The service account that needs permission to create table and upload data to the BigQuery dataset.
func (o SccV2FolderSccBigQueryExportOutput) Principal() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.Principal }).(pulumi.StringOutput)
}

func (o SccV2FolderSccBigQueryExportOutput) SccV2FolderSccBigQueryExportId() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.SccV2FolderSccBigQueryExportId }).(pulumi.StringOutput)
}

func (o SccV2FolderSccBigQueryExportOutput) Timeouts() SccV2FolderSccBigQueryExportTimeoutsPtrOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) SccV2FolderSccBigQueryExportTimeoutsPtrOutput { return v.Timeouts }).(SccV2FolderSccBigQueryExportTimeoutsPtrOutput)
}

// The most recent time at which the BigQuery export was updated. This field is set by the server and will be ignored if
// provided on export creation or update. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
// nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o SccV2FolderSccBigQueryExportOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *SccV2FolderSccBigQueryExport) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SccV2FolderSccBigQueryExportInput)(nil)).Elem(), &SccV2FolderSccBigQueryExport{})
	pulumi.RegisterOutputType(SccV2FolderSccBigQueryExportOutput{})
}
