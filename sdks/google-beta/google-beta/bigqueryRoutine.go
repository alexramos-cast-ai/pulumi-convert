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

type BigqueryRoutine struct {
	pulumi.CustomResourceState

	// Input/output argument of a function or a stored procedure.
	Arguments         BigqueryRoutineArgumentArrayOutput `pulumi:"arguments"`
	BigqueryRoutineId pulumi.StringOutput                `pulumi:"bigqueryRoutineId"`
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime pulumi.Float64Output `pulumi:"creationTime"`
	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
	// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
	DataGovernanceType pulumi.StringPtrOutput `pulumi:"dataGovernanceType"`
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringOutput `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
	// inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringOutput `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
	// "DETERMINISTIC", "NOT_DETERMINISTIC"]
	DeterminismLevel pulumi.StringPtrOutput `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayOutput `pulumi:"importedLibraries"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
	Language pulumi.StringPtrOutput `pulumi:"language"`
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.Float64Output `pulumi:"lastModifiedTime"`
	Project          pulumi.StringOutput  `pulumi:"project"`
	// Remote function specific options.
	RemoteFunctionOptions BigqueryRoutineRemoteFunctionOptionsPtrOutput `pulumi:"remoteFunctionOptions"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
	// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
	// table result will be cast to match the column types specificed in return table type, at query time.
	ReturnTableType pulumi.StringPtrOutput `pulumi:"returnTableType"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
	// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
	// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
	// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
	// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
	// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
	// by the API.
	ReturnType pulumi.StringPtrOutput `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	RoutineId pulumi.StringOutput `pulumi:"routineId"`
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	RoutineType pulumi.StringOutput `pulumi:"routineType"`
	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	SparkOptions BigqueryRoutineSparkOptionsPtrOutput `pulumi:"sparkOptions"`
	Timeouts     BigqueryRoutineTimeoutsPtrOutput     `pulumi:"timeouts"`
}

// NewBigqueryRoutine registers a new resource with the given unique name, arguments, and options.
func NewBigqueryRoutine(ctx *pulumi.Context,
	name string, args *BigqueryRoutineArgs, opts ...pulumi.ResourceOption) (*BigqueryRoutine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.DefinitionBody == nil {
		return nil, errors.New("invalid value for required argument 'DefinitionBody'")
	}
	if args.RoutineId == nil {
		return nil, errors.New("invalid value for required argument 'RoutineId'")
	}
	if args.RoutineType == nil {
		return nil, errors.New("invalid value for required argument 'RoutineType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryRoutine
	err = ctx.RegisterPackageResource("google-beta:index/bigqueryRoutine:BigqueryRoutine", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryRoutine gets an existing BigqueryRoutine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryRoutine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryRoutineState, opts ...pulumi.ResourceOption) (*BigqueryRoutine, error) {
	var resource BigqueryRoutine
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/bigqueryRoutine:BigqueryRoutine", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryRoutine resources.
type bigqueryRoutineState struct {
	// Input/output argument of a function or a stored procedure.
	Arguments         []BigqueryRoutineArgument `pulumi:"arguments"`
	BigqueryRoutineId *string                   `pulumi:"bigqueryRoutineId"`
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime *float64 `pulumi:"creationTime"`
	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
	// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
	DataGovernanceType *string `pulumi:"dataGovernanceType"`
	// The ID of the dataset containing this routine
	DatasetId *string `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
	// inside (but excluding) the parentheses.
	DefinitionBody *string `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description *string `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
	// "DETERMINISTIC", "NOT_DETERMINISTIC"]
	DeterminismLevel *string `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries []string `pulumi:"importedLibraries"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
	Language *string `pulumi:"language"`
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime *float64 `pulumi:"lastModifiedTime"`
	Project          *string  `pulumi:"project"`
	// Remote function specific options.
	RemoteFunctionOptions *BigqueryRoutineRemoteFunctionOptions `pulumi:"remoteFunctionOptions"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
	// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
	// table result will be cast to match the column types specificed in return table type, at query time.
	ReturnTableType *string `pulumi:"returnTableType"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
	// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
	// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
	// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
	// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
	// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
	// by the API.
	ReturnType *string `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	RoutineId *string `pulumi:"routineId"`
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	RoutineType *string `pulumi:"routineType"`
	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	SparkOptions *BigqueryRoutineSparkOptions `pulumi:"sparkOptions"`
	Timeouts     *BigqueryRoutineTimeouts     `pulumi:"timeouts"`
}

type BigqueryRoutineState struct {
	// Input/output argument of a function or a stored procedure.
	Arguments         BigqueryRoutineArgumentArrayInput
	BigqueryRoutineId pulumi.StringPtrInput
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime pulumi.Float64PtrInput
	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
	// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
	DataGovernanceType pulumi.StringPtrInput
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringPtrInput
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
	// inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringPtrInput
	// The description of the routine if defined.
	Description pulumi.StringPtrInput
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
	// "DETERMINISTIC", "NOT_DETERMINISTIC"]
	DeterminismLevel pulumi.StringPtrInput
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayInput
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
	Language pulumi.StringPtrInput
	// The time when this routine was modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.Float64PtrInput
	Project          pulumi.StringPtrInput
	// Remote function specific options.
	RemoteFunctionOptions BigqueryRoutineRemoteFunctionOptionsPtrInput
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
	// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
	// table result will be cast to match the column types specificed in return table type, at query time.
	ReturnTableType pulumi.StringPtrInput
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
	// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
	// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
	// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
	// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
	// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
	// by the API.
	ReturnType pulumi.StringPtrInput
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	RoutineId pulumi.StringPtrInput
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	RoutineType pulumi.StringPtrInput
	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	SparkOptions BigqueryRoutineSparkOptionsPtrInput
	Timeouts     BigqueryRoutineTimeoutsPtrInput
}

func (BigqueryRoutineState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryRoutineState)(nil)).Elem()
}

type bigqueryRoutineArgs struct {
	// Input/output argument of a function or a stored procedure.
	Arguments         []BigqueryRoutineArgument `pulumi:"arguments"`
	BigqueryRoutineId *string                   `pulumi:"bigqueryRoutineId"`
	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
	// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
	DataGovernanceType *string `pulumi:"dataGovernanceType"`
	// The ID of the dataset containing this routine
	DatasetId string `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
	// inside (but excluding) the parentheses.
	DefinitionBody string `pulumi:"definitionBody"`
	// The description of the routine if defined.
	Description *string `pulumi:"description"`
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
	// "DETERMINISTIC", "NOT_DETERMINISTIC"]
	DeterminismLevel *string `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries []string `pulumi:"importedLibraries"`
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
	Language *string `pulumi:"language"`
	Project  *string `pulumi:"project"`
	// Remote function specific options.
	RemoteFunctionOptions *BigqueryRoutineRemoteFunctionOptions `pulumi:"remoteFunctionOptions"`
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
	// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
	// table result will be cast to match the column types specificed in return table type, at query time.
	ReturnTableType *string `pulumi:"returnTableType"`
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
	// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
	// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
	// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
	// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
	// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
	// by the API.
	ReturnType *string `pulumi:"returnType"`
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	RoutineId string `pulumi:"routineId"`
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	RoutineType string `pulumi:"routineType"`
	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	SparkOptions *BigqueryRoutineSparkOptions `pulumi:"sparkOptions"`
	Timeouts     *BigqueryRoutineTimeouts     `pulumi:"timeouts"`
}

// The set of arguments for constructing a BigqueryRoutine resource.
type BigqueryRoutineArgs struct {
	// Input/output argument of a function or a stored procedure.
	Arguments         BigqueryRoutineArgumentArrayInput
	BigqueryRoutineId pulumi.StringPtrInput
	// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
	// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
	DataGovernanceType pulumi.StringPtrInput
	// The ID of the dataset containing this routine
	DatasetId pulumi.StringInput
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
	// inside (but excluding) the parentheses.
	DefinitionBody pulumi.StringInput
	// The description of the routine if defined.
	Description pulumi.StringPtrInput
	// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
	// "DETERMINISTIC", "NOT_DETERMINISTIC"]
	DeterminismLevel pulumi.StringPtrInput
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayInput
	// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
	Language pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	// Remote function specific options.
	RemoteFunctionOptions BigqueryRoutineRemoteFunctionOptionsPtrInput
	// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
	// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
	// table result will be cast to match the column types specificed in return table type, at query time.
	ReturnTableType pulumi.StringPtrInput
	// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
	// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
	// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
	// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
	// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
	// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
	// by the API.
	ReturnType pulumi.StringPtrInput
	// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
	// length is 256 characters.
	RoutineId pulumi.StringInput
	// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
	RoutineType pulumi.StringInput
	// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
	SparkOptions BigqueryRoutineSparkOptionsPtrInput
	Timeouts     BigqueryRoutineTimeoutsPtrInput
}

func (BigqueryRoutineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryRoutineArgs)(nil)).Elem()
}

type BigqueryRoutineInput interface {
	pulumi.Input

	ToBigqueryRoutineOutput() BigqueryRoutineOutput
	ToBigqueryRoutineOutputWithContext(ctx context.Context) BigqueryRoutineOutput
}

func (*BigqueryRoutine) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryRoutine)(nil)).Elem()
}

func (i *BigqueryRoutine) ToBigqueryRoutineOutput() BigqueryRoutineOutput {
	return i.ToBigqueryRoutineOutputWithContext(context.Background())
}

func (i *BigqueryRoutine) ToBigqueryRoutineOutputWithContext(ctx context.Context) BigqueryRoutineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryRoutineOutput)
}

type BigqueryRoutineOutput struct{ *pulumi.OutputState }

func (BigqueryRoutineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryRoutine)(nil)).Elem()
}

func (o BigqueryRoutineOutput) ToBigqueryRoutineOutput() BigqueryRoutineOutput {
	return o
}

func (o BigqueryRoutineOutput) ToBigqueryRoutineOutputWithContext(ctx context.Context) BigqueryRoutineOutput {
	return o
}

// Input/output argument of a function or a stored procedure.
func (o BigqueryRoutineOutput) Arguments() BigqueryRoutineArgumentArrayOutput {
	return o.ApplyT(func(v *BigqueryRoutine) BigqueryRoutineArgumentArrayOutput { return v.Arguments }).(BigqueryRoutineArgumentArrayOutput)
}

func (o BigqueryRoutineOutput) BigqueryRoutineId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.BigqueryRoutineId }).(pulumi.StringOutput)
}

// The time when this routine was created, in milliseconds since the epoch.
func (o BigqueryRoutineOutput) CreationTime() pulumi.Float64Output {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.Float64Output { return v.CreationTime }).(pulumi.Float64Output)
}

// If set to DATA_MASKING, the function is validated and made available as a masking function. For more information, see
// https://cloud.google.com/bigquery/docs/user-defined-functions#custom-mask Possible values: ["DATA_MASKING"]
func (o BigqueryRoutineOutput) DataGovernanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.DataGovernanceType }).(pulumi.StringPtrOutput)
}

// The ID of the dataset containing this routine
func (o BigqueryRoutineOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring
// inside (but excluding) the parentheses.
func (o BigqueryRoutineOutput) DefinitionBody() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.DefinitionBody }).(pulumi.StringOutput)
}

// The description of the routine if defined.
func (o BigqueryRoutineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED",
// "DETERMINISTIC", "NOT_DETERMINISTIC"]
func (o BigqueryRoutineOutput) DeterminismLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.DeterminismLevel }).(pulumi.StringPtrOutput)
}

// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
func (o BigqueryRoutineOutput) ImportedLibraries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringArrayOutput { return v.ImportedLibraries }).(pulumi.StringArrayOutput)
}

// The language of the routine. Possible values: ["SQL", "JAVASCRIPT", "PYTHON", "JAVA", "SCALA"]
func (o BigqueryRoutineOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.Language }).(pulumi.StringPtrOutput)
}

// The time when this routine was modified, in milliseconds since the epoch.
func (o BigqueryRoutineOutput) LastModifiedTime() pulumi.Float64Output {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.Float64Output { return v.LastModifiedTime }).(pulumi.Float64Output)
}

func (o BigqueryRoutineOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Remote function specific options.
func (o BigqueryRoutineOutput) RemoteFunctionOptions() BigqueryRoutineRemoteFunctionOptionsPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) BigqueryRoutineRemoteFunctionOptionsPtrOutput { return v.RemoteFunctionOptions }).(BigqueryRoutineRemoteFunctionOptionsPtrOutput)
}

// Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from
// definitionBody at query time in each query that references this routine. If present, then the columns in the evaluated
// table result will be cast to match the column types specificed in return table type, at query time.
func (o BigqueryRoutineOutput) ReturnTableType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.ReturnTableType }).(pulumi.StringPtrOutput)
}

// A JSON schema for the return type. Optional if language = "SQL"; required otherwise. If absent, the return type is
// inferred from definitionBody at query time in each query that references this routine. If present, then the evaluated
// result will be cast to the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON string,
// any changes to the string will create a diff, even if the JSON itself hasn't changed. If the API returns a different
// value for the same schema, e.g. it switche d the order of values or replaced STRUCT field type with RECORD field type,
// we currently cannot suppress the recurring diff this causes. As a workaround, we recommend using the schema as returned
// by the API.
func (o BigqueryRoutineOutput) ReturnType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringPtrOutput { return v.ReturnType }).(pulumi.StringPtrOutput)
}

// The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum
// length is 256 characters.
func (o BigqueryRoutineOutput) RoutineId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.RoutineId }).(pulumi.StringOutput)
}

// The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]
func (o BigqueryRoutineOutput) RoutineType() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryRoutine) pulumi.StringOutput { return v.RoutineType }).(pulumi.StringOutput)
}

// Optional. If language is one of "PYTHON", "JAVA", "SCALA", this field stores the options for spark stored procedure.
func (o BigqueryRoutineOutput) SparkOptions() BigqueryRoutineSparkOptionsPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) BigqueryRoutineSparkOptionsPtrOutput { return v.SparkOptions }).(BigqueryRoutineSparkOptionsPtrOutput)
}

func (o BigqueryRoutineOutput) Timeouts() BigqueryRoutineTimeoutsPtrOutput {
	return o.ApplyT(func(v *BigqueryRoutine) BigqueryRoutineTimeoutsPtrOutput { return v.Timeouts }).(BigqueryRoutineTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryRoutineInput)(nil)).Elem(), &BigqueryRoutine{})
	pulumi.RegisterOutputType(BigqueryRoutineOutput{})
}
