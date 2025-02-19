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

type SpannerDatabase struct {
	pulumi.CustomResourceState

	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
	// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
	DatabaseDialect pulumi.StringOutput `pulumi:"databaseDialect"`
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls                 pulumi.StringArrayOutput `pulumi:"ddls"`
	DeletionProtection   pulumi.BoolPtrOutput     `pulumi:"deletionProtection"`
	EnableDropProtection pulumi.BoolPtrOutput     `pulumi:"enableDropProtection"`
	// Encryption configuration for the database
	EncryptionConfig SpannerDatabaseEncryptionConfigPtrOutput `pulumi:"encryptionConfig"`
	// The instance to create the database on.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// '[a-z][-_a-z0-9]*[a-z0-9]'.
	Name              pulumi.StringOutput `pulumi:"name"`
	Project           pulumi.StringOutput `pulumi:"project"`
	SpannerDatabaseId pulumi.StringOutput `pulumi:"spannerDatabaseId"`
	// An explanation of the status of the database.
	State    pulumi.StringOutput              `pulumi:"state"`
	Timeouts SpannerDatabaseTimeoutsPtrOutput `pulumi:"timeouts"`
	// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
	// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
	// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
	// version_retention_period.
	VersionRetentionPeriod pulumi.StringOutput `pulumi:"versionRetentionPeriod"`
}

// NewSpannerDatabase registers a new resource with the given unique name, arguments, and options.
func NewSpannerDatabase(ctx *pulumi.Context,
	name string, args *SpannerDatabaseArgs, opts ...pulumi.ResourceOption) (*SpannerDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource SpannerDatabase
	err = ctx.RegisterPackageResource("google-beta:index/spannerDatabase:SpannerDatabase", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpannerDatabase gets an existing SpannerDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpannerDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpannerDatabaseState, opts ...pulumi.ResourceOption) (*SpannerDatabase, error) {
	var resource SpannerDatabase
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/spannerDatabase:SpannerDatabase", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpannerDatabase resources.
type spannerDatabaseState struct {
	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
	// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
	DatabaseDialect *string `pulumi:"databaseDialect"`
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls                 []string `pulumi:"ddls"`
	DeletionProtection   *bool    `pulumi:"deletionProtection"`
	EnableDropProtection *bool    `pulumi:"enableDropProtection"`
	// Encryption configuration for the database
	EncryptionConfig *SpannerDatabaseEncryptionConfig `pulumi:"encryptionConfig"`
	// The instance to create the database on.
	Instance *string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// '[a-z][-_a-z0-9]*[a-z0-9]'.
	Name              *string `pulumi:"name"`
	Project           *string `pulumi:"project"`
	SpannerDatabaseId *string `pulumi:"spannerDatabaseId"`
	// An explanation of the status of the database.
	State    *string                  `pulumi:"state"`
	Timeouts *SpannerDatabaseTimeouts `pulumi:"timeouts"`
	// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
	// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
	// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
	// version_retention_period.
	VersionRetentionPeriod *string `pulumi:"versionRetentionPeriod"`
}

type SpannerDatabaseState struct {
	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
	// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
	DatabaseDialect pulumi.StringPtrInput
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls                 pulumi.StringArrayInput
	DeletionProtection   pulumi.BoolPtrInput
	EnableDropProtection pulumi.BoolPtrInput
	// Encryption configuration for the database
	EncryptionConfig SpannerDatabaseEncryptionConfigPtrInput
	// The instance to create the database on.
	Instance pulumi.StringPtrInput
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// '[a-z][-_a-z0-9]*[a-z0-9]'.
	Name              pulumi.StringPtrInput
	Project           pulumi.StringPtrInput
	SpannerDatabaseId pulumi.StringPtrInput
	// An explanation of the status of the database.
	State    pulumi.StringPtrInput
	Timeouts SpannerDatabaseTimeoutsPtrInput
	// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
	// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
	// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
	// version_retention_period.
	VersionRetentionPeriod pulumi.StringPtrInput
}

func (SpannerDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*spannerDatabaseState)(nil)).Elem()
}

type spannerDatabaseArgs struct {
	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
	// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
	DatabaseDialect *string `pulumi:"databaseDialect"`
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls                 []string `pulumi:"ddls"`
	DeletionProtection   *bool    `pulumi:"deletionProtection"`
	EnableDropProtection *bool    `pulumi:"enableDropProtection"`
	// Encryption configuration for the database
	EncryptionConfig *SpannerDatabaseEncryptionConfig `pulumi:"encryptionConfig"`
	// The instance to create the database on.
	Instance string `pulumi:"instance"`
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// '[a-z][-_a-z0-9]*[a-z0-9]'.
	Name              *string                  `pulumi:"name"`
	Project           *string                  `pulumi:"project"`
	SpannerDatabaseId *string                  `pulumi:"spannerDatabaseId"`
	Timeouts          *SpannerDatabaseTimeouts `pulumi:"timeouts"`
	// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
	// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
	// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
	// version_retention_period.
	VersionRetentionPeriod *string `pulumi:"versionRetentionPeriod"`
}

// The set of arguments for constructing a SpannerDatabase resource.
type SpannerDatabaseArgs struct {
	// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
	// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
	DatabaseDialect pulumi.StringPtrInput
	// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
	// These statements execute atomically with the creation of the database: if there is an error in any statement, the
	// database is not created.
	Ddls                 pulumi.StringArrayInput
	DeletionProtection   pulumi.BoolPtrInput
	EnableDropProtection pulumi.BoolPtrInput
	// Encryption configuration for the database
	EncryptionConfig SpannerDatabaseEncryptionConfigPtrInput
	// The instance to create the database on.
	Instance pulumi.StringInput
	// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
	// '[a-z][-_a-z0-9]*[a-z0-9]'.
	Name              pulumi.StringPtrInput
	Project           pulumi.StringPtrInput
	SpannerDatabaseId pulumi.StringPtrInput
	Timeouts          SpannerDatabaseTimeoutsPtrInput
	// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
	// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
	// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
	// version_retention_period.
	VersionRetentionPeriod pulumi.StringPtrInput
}

func (SpannerDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spannerDatabaseArgs)(nil)).Elem()
}

type SpannerDatabaseInput interface {
	pulumi.Input

	ToSpannerDatabaseOutput() SpannerDatabaseOutput
	ToSpannerDatabaseOutputWithContext(ctx context.Context) SpannerDatabaseOutput
}

func (*SpannerDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**SpannerDatabase)(nil)).Elem()
}

func (i *SpannerDatabase) ToSpannerDatabaseOutput() SpannerDatabaseOutput {
	return i.ToSpannerDatabaseOutputWithContext(context.Background())
}

func (i *SpannerDatabase) ToSpannerDatabaseOutputWithContext(ctx context.Context) SpannerDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpannerDatabaseOutput)
}

type SpannerDatabaseOutput struct{ *pulumi.OutputState }

func (SpannerDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SpannerDatabase)(nil)).Elem()
}

func (o SpannerDatabaseOutput) ToSpannerDatabaseOutput() SpannerDatabaseOutput {
	return o
}

func (o SpannerDatabaseOutput) ToSpannerDatabaseOutputWithContext(ctx context.Context) SpannerDatabaseOutput {
	return o
}

// The dialect of the Cloud Spanner Database. If it is not provided, "GOOGLE_STANDARD_SQL" will be used. Possible values:
// ["GOOGLE_STANDARD_SQL", "POSTGRESQL"]
func (o SpannerDatabaseOutput) DatabaseDialect() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.DatabaseDialect }).(pulumi.StringOutput)
}

// An optional list of DDL statements to run inside the newly created database. Statements can create tables, indexes, etc.
// These statements execute atomically with the creation of the database: if there is an error in any statement, the
// database is not created.
func (o SpannerDatabaseOutput) Ddls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringArrayOutput { return v.Ddls }).(pulumi.StringArrayOutput)
}

func (o SpannerDatabaseOutput) DeletionProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.BoolPtrOutput { return v.DeletionProtection }).(pulumi.BoolPtrOutput)
}

func (o SpannerDatabaseOutput) EnableDropProtection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.BoolPtrOutput { return v.EnableDropProtection }).(pulumi.BoolPtrOutput)
}

// Encryption configuration for the database
func (o SpannerDatabaseOutput) EncryptionConfig() SpannerDatabaseEncryptionConfigPtrOutput {
	return o.ApplyT(func(v *SpannerDatabase) SpannerDatabaseEncryptionConfigPtrOutput { return v.EncryptionConfig }).(SpannerDatabaseEncryptionConfigPtrOutput)
}

// The instance to create the database on.
func (o SpannerDatabaseOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// A unique identifier for the database, which cannot be changed after the instance is created. Values are of the form
// '[a-z][-_a-z0-9]*[a-z0-9]'.
func (o SpannerDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SpannerDatabaseOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o SpannerDatabaseOutput) SpannerDatabaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.SpannerDatabaseId }).(pulumi.StringOutput)
}

// An explanation of the status of the database.
func (o SpannerDatabaseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o SpannerDatabaseOutput) Timeouts() SpannerDatabaseTimeoutsPtrOutput {
	return o.ApplyT(func(v *SpannerDatabase) SpannerDatabaseTimeoutsPtrOutput { return v.Timeouts }).(SpannerDatabaseTimeoutsPtrOutput)
}

// The retention period for the database. The retention period must be between 1 hour and 7 days, and can be specified in
// days, hours, minutes, or seconds. For example, the values 1d, 24h, 1440m, and 86400s are equivalent. Default value is
// 1h. If this property is used, you must avoid adding new DDL statements to 'ddl' that update the database's
// version_retention_period.
func (o SpannerDatabaseOutput) VersionRetentionPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v *SpannerDatabase) pulumi.StringOutput { return v.VersionRetentionPeriod }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SpannerDatabaseInput)(nil)).Elem(), &SpannerDatabase{})
	pulumi.RegisterOutputType(SpannerDatabaseOutput{})
}
