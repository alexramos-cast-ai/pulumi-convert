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

type ChronicleRule struct {
	pulumi.CustomResourceState

	// Output only. The run frequencies that are allowed for the rule. Populated in BASIC view and FULL view.
	AllowedRunFrequencies pulumi.StringArrayOutput `pulumi:"allowedRunFrequencies"`
	// Output only. The author of the rule. Extracted from the meta section of text. Populated in BASIC view and FULL view.
	Author          pulumi.StringOutput `pulumi:"author"`
	ChronicleRuleId pulumi.StringOutput `pulumi:"chronicleRuleId"`
	// Output only. A list of a rule's corresponding compilation diagnostic messages such as compilation errors and compilation
	// warnings. Populated in FULL view.
	CompilationDiagnostics ChronicleRuleCompilationDiagnosticArrayOutput `pulumi:"compilationDiagnostics"`
	// Output only. The current compilation state of the rule. Populated in FULL view. Possible values:
	// COMPILATION_STATE_UNSPECIFIED SUCCEEDED FAILED
	CompilationState pulumi.StringOutput `pulumi:"compilationState"`
	// Output only. The timestamp of when the rule was created. Populated in FULL view.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. Resource names of the data tables used in this rule.
	DataTables pulumi.StringArrayOutput `pulumi:"dataTables"`
	// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
	// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
	// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
	// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// Output only. Display name of the rule. Populated in BASIC view and FULL view.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
	// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. Additional metadata specified in the meta section of text. Populated in FULL view.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.
	// Format: projects/{project}/locations/{location}/instances/{instance}/rules/{rule}
	Name pulumi.StringOutput `pulumi:"name"`
	// Output only. Indicate the rule can run in near real time live rule. If this is true, the rule uses the near real time
	// live rule when the run frequency is set to LIVE.
	NearRealTimeLiveRuleEligible pulumi.BoolOutput   `pulumi:"nearRealTimeLiveRuleEligible"`
	Project                      pulumi.StringOutput `pulumi:"project"`
	// Output only. Resource names of the reference lists used in this rule. Populated in FULL view.
	ReferenceLists pulumi.StringArrayOutput `pulumi:"referenceLists"`
	// Output only. The timestamp of when the rule revision was created. Populated in FULL, REVISION_METADATA_ONLY views.
	RevisionCreateTime pulumi.StringOutput `pulumi:"revisionCreateTime"`
	// Output only. The revision ID of the rule. A new revision is created whenever the rule text is changed in any way.
	// Format: v_{10 digits}_{9 digits} Populated in REVISION_METADATA_ONLY view and FULL view.
	RevisionId pulumi.StringOutput `pulumi:"revisionId"`
	// Rule Id is the ID of the Rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
	// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
	// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
	// with both the user's and the rule's scopes. The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	Scope pulumi.StringPtrOutput `pulumi:"scope"`
	// Severity represents the severity level of the rule.
	Severities ChronicleRuleSeverityArrayOutput `pulumi:"severities"`
	// The YARA-L content of the rule. Populated in FULL view.
	Text     pulumi.StringPtrOutput         `pulumi:"text"`
	Timeouts ChronicleRuleTimeoutsPtrOutput `pulumi:"timeouts"`
	// Possible values: RULE_TYPE_UNSPECIFIED SINGLE_EVENT MULTI_EVENT
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewChronicleRule registers a new resource with the given unique name, arguments, and options.
func NewChronicleRule(ctx *pulumi.Context,
	name string, args *ChronicleRuleArgs, opts ...pulumi.ResourceOption) (*ChronicleRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ChronicleRule
	err = ctx.RegisterPackageResource("google-beta:index/chronicleRule:ChronicleRule", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetChronicleRule gets an existing ChronicleRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetChronicleRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ChronicleRuleState, opts ...pulumi.ResourceOption) (*ChronicleRule, error) {
	var resource ChronicleRule
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/chronicleRule:ChronicleRule", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ChronicleRule resources.
type chronicleRuleState struct {
	// Output only. The run frequencies that are allowed for the rule. Populated in BASIC view and FULL view.
	AllowedRunFrequencies []string `pulumi:"allowedRunFrequencies"`
	// Output only. The author of the rule. Extracted from the meta section of text. Populated in BASIC view and FULL view.
	Author          *string `pulumi:"author"`
	ChronicleRuleId *string `pulumi:"chronicleRuleId"`
	// Output only. A list of a rule's corresponding compilation diagnostic messages such as compilation errors and compilation
	// warnings. Populated in FULL view.
	CompilationDiagnostics []ChronicleRuleCompilationDiagnostic `pulumi:"compilationDiagnostics"`
	// Output only. The current compilation state of the rule. Populated in FULL view. Possible values:
	// COMPILATION_STATE_UNSPECIFIED SUCCEEDED FAILED
	CompilationState *string `pulumi:"compilationState"`
	// Output only. The timestamp of when the rule was created. Populated in FULL view.
	CreateTime *string `pulumi:"createTime"`
	// Output only. Resource names of the data tables used in this rule.
	DataTables []string `pulumi:"dataTables"`
	// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
	// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
	// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
	// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// Output only. Display name of the rule. Populated in BASIC view and FULL view.
	DisplayName *string `pulumi:"displayName"`
	// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
	// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
	Etag *string `pulumi:"etag"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance *string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location *string `pulumi:"location"`
	// Output only. Additional metadata specified in the meta section of text. Populated in FULL view.
	Metadata map[string]string `pulumi:"metadata"`
	// Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.
	// Format: projects/{project}/locations/{location}/instances/{instance}/rules/{rule}
	Name *string `pulumi:"name"`
	// Output only. Indicate the rule can run in near real time live rule. If this is true, the rule uses the near real time
	// live rule when the run frequency is set to LIVE.
	NearRealTimeLiveRuleEligible *bool   `pulumi:"nearRealTimeLiveRuleEligible"`
	Project                      *string `pulumi:"project"`
	// Output only. Resource names of the reference lists used in this rule. Populated in FULL view.
	ReferenceLists []string `pulumi:"referenceLists"`
	// Output only. The timestamp of when the rule revision was created. Populated in FULL, REVISION_METADATA_ONLY views.
	RevisionCreateTime *string `pulumi:"revisionCreateTime"`
	// Output only. The revision ID of the rule. A new revision is created whenever the rule text is changed in any way.
	// Format: v_{10 digits}_{9 digits} Populated in REVISION_METADATA_ONLY view and FULL view.
	RevisionId *string `pulumi:"revisionId"`
	// Rule Id is the ID of the Rule.
	RuleId *string `pulumi:"ruleId"`
	// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
	// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
	// with both the user's and the rule's scopes. The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	Scope *string `pulumi:"scope"`
	// Severity represents the severity level of the rule.
	Severities []ChronicleRuleSeverity `pulumi:"severities"`
	// The YARA-L content of the rule. Populated in FULL view.
	Text     *string                `pulumi:"text"`
	Timeouts *ChronicleRuleTimeouts `pulumi:"timeouts"`
	// Possible values: RULE_TYPE_UNSPECIFIED SINGLE_EVENT MULTI_EVENT
	Type *string `pulumi:"type"`
}

type ChronicleRuleState struct {
	// Output only. The run frequencies that are allowed for the rule. Populated in BASIC view and FULL view.
	AllowedRunFrequencies pulumi.StringArrayInput
	// Output only. The author of the rule. Extracted from the meta section of text. Populated in BASIC view and FULL view.
	Author          pulumi.StringPtrInput
	ChronicleRuleId pulumi.StringPtrInput
	// Output only. A list of a rule's corresponding compilation diagnostic messages such as compilation errors and compilation
	// warnings. Populated in FULL view.
	CompilationDiagnostics ChronicleRuleCompilationDiagnosticArrayInput
	// Output only. The current compilation state of the rule. Populated in FULL view. Possible values:
	// COMPILATION_STATE_UNSPECIFIED SUCCEEDED FAILED
	CompilationState pulumi.StringPtrInput
	// Output only. The timestamp of when the rule was created. Populated in FULL view.
	CreateTime pulumi.StringPtrInput
	// Output only. Resource names of the data tables used in this rule.
	DataTables pulumi.StringArrayInput
	// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
	// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
	// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
	// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
	DeletionPolicy pulumi.StringPtrInput
	// Output only. Display name of the rule. Populated in BASIC view and FULL view.
	DisplayName pulumi.StringPtrInput
	// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
	// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
	Etag pulumi.StringPtrInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringPtrInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringPtrInput
	// Output only. Additional metadata specified in the meta section of text. Populated in FULL view.
	Metadata pulumi.StringMapInput
	// Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.
	// Format: projects/{project}/locations/{location}/instances/{instance}/rules/{rule}
	Name pulumi.StringPtrInput
	// Output only. Indicate the rule can run in near real time live rule. If this is true, the rule uses the near real time
	// live rule when the run frequency is set to LIVE.
	NearRealTimeLiveRuleEligible pulumi.BoolPtrInput
	Project                      pulumi.StringPtrInput
	// Output only. Resource names of the reference lists used in this rule. Populated in FULL view.
	ReferenceLists pulumi.StringArrayInput
	// Output only. The timestamp of when the rule revision was created. Populated in FULL, REVISION_METADATA_ONLY views.
	RevisionCreateTime pulumi.StringPtrInput
	// Output only. The revision ID of the rule. A new revision is created whenever the rule text is changed in any way.
	// Format: v_{10 digits}_{9 digits} Populated in REVISION_METADATA_ONLY view and FULL view.
	RevisionId pulumi.StringPtrInput
	// Rule Id is the ID of the Rule.
	RuleId pulumi.StringPtrInput
	// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
	// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
	// with both the user's and the rule's scopes. The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	Scope pulumi.StringPtrInput
	// Severity represents the severity level of the rule.
	Severities ChronicleRuleSeverityArrayInput
	// The YARA-L content of the rule. Populated in FULL view.
	Text     pulumi.StringPtrInput
	Timeouts ChronicleRuleTimeoutsPtrInput
	// Possible values: RULE_TYPE_UNSPECIFIED SINGLE_EVENT MULTI_EVENT
	Type pulumi.StringPtrInput
}

func (ChronicleRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleRuleState)(nil)).Elem()
}

type chronicleRuleArgs struct {
	ChronicleRuleId *string `pulumi:"chronicleRuleId"`
	// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
	// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
	// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
	// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
	// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
	Etag *string `pulumi:"etag"`
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance string `pulumi:"instance"`
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Rule Id is the ID of the Rule.
	RuleId *string `pulumi:"ruleId"`
	// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
	// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
	// with both the user's and the rule's scopes. The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	Scope *string `pulumi:"scope"`
	// The YARA-L content of the rule. Populated in FULL view.
	Text     *string                `pulumi:"text"`
	Timeouts *ChronicleRuleTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ChronicleRule resource.
type ChronicleRuleArgs struct {
	ChronicleRuleId pulumi.StringPtrInput
	// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
	// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
	// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
	// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
	DeletionPolicy pulumi.StringPtrInput
	// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
	// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
	Etag pulumi.StringPtrInput
	// The unique identifier for the Chronicle instance, which is the same as the customer ID.
	Instance pulumi.StringInput
	// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
	// "europe-west2".
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Rule Id is the ID of the Rule.
	RuleId pulumi.StringPtrInput
	// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
	// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
	// with both the user's and the rule's scopes. The scope should be in the format:
	// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
	Scope pulumi.StringPtrInput
	// The YARA-L content of the rule. Populated in FULL view.
	Text     pulumi.StringPtrInput
	Timeouts ChronicleRuleTimeoutsPtrInput
}

func (ChronicleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*chronicleRuleArgs)(nil)).Elem()
}

type ChronicleRuleInput interface {
	pulumi.Input

	ToChronicleRuleOutput() ChronicleRuleOutput
	ToChronicleRuleOutputWithContext(ctx context.Context) ChronicleRuleOutput
}

func (*ChronicleRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleRule)(nil)).Elem()
}

func (i *ChronicleRule) ToChronicleRuleOutput() ChronicleRuleOutput {
	return i.ToChronicleRuleOutputWithContext(context.Background())
}

func (i *ChronicleRule) ToChronicleRuleOutputWithContext(ctx context.Context) ChronicleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ChronicleRuleOutput)
}

type ChronicleRuleOutput struct{ *pulumi.OutputState }

func (ChronicleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChronicleRule)(nil)).Elem()
}

func (o ChronicleRuleOutput) ToChronicleRuleOutput() ChronicleRuleOutput {
	return o
}

func (o ChronicleRuleOutput) ToChronicleRuleOutputWithContext(ctx context.Context) ChronicleRuleOutput {
	return o
}

// Output only. The run frequencies that are allowed for the rule. Populated in BASIC view and FULL view.
func (o ChronicleRuleOutput) AllowedRunFrequencies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringArrayOutput { return v.AllowedRunFrequencies }).(pulumi.StringArrayOutput)
}

// Output only. The author of the rule. Extracted from the meta section of text. Populated in BASIC view and FULL view.
func (o ChronicleRuleOutput) Author() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Author }).(pulumi.StringOutput)
}

func (o ChronicleRuleOutput) ChronicleRuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.ChronicleRuleId }).(pulumi.StringOutput)
}

// Output only. A list of a rule's corresponding compilation diagnostic messages such as compilation errors and compilation
// warnings. Populated in FULL view.
func (o ChronicleRuleOutput) CompilationDiagnostics() ChronicleRuleCompilationDiagnosticArrayOutput {
	return o.ApplyT(func(v *ChronicleRule) ChronicleRuleCompilationDiagnosticArrayOutput { return v.CompilationDiagnostics }).(ChronicleRuleCompilationDiagnosticArrayOutput)
}

// Output only. The current compilation state of the rule. Populated in FULL view. Possible values:
// COMPILATION_STATE_UNSPECIFIED SUCCEEDED FAILED
func (o ChronicleRuleOutput) CompilationState() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.CompilationState }).(pulumi.StringOutput)
}

// Output only. The timestamp of when the rule was created. Populated in FULL view.
func (o ChronicleRuleOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. Resource names of the data tables used in this rule.
func (o ChronicleRuleOutput) DataTables() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringArrayOutput { return v.DataTables }).(pulumi.StringArrayOutput)
}

// Policy to determine if the rule should be deleted forcefully. If deletion_policy = "FORCE", any retrohunts and any
// detections associated with the rule will also be deleted. If deletion_policy = "DEFAULT", the call will only succeed if
// the rule has no associated retrohunts, including completed retrohunts, and no associated detections. Regardless of this
// field's value, the rule deployment associated with this rule will also be deleted. Possible values: DEFAULT, FORCE
func (o ChronicleRuleOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// Output only. Display name of the rule. Populated in BASIC view and FULL view.
func (o ChronicleRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The etag for this rule. If this is provided on update, the request will succeed if and only if it matches the
// server-computed value, and will fail with an ABORTED error otherwise. Populated in BASIC view and FULL view.
func (o ChronicleRuleOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The unique identifier for the Chronicle instance, which is the same as the customer ID.
func (o ChronicleRuleOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// The location of the resource. This is the geographical region where the Chronicle instance resides, such as "us" or
// "europe-west2".
func (o ChronicleRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. Additional metadata specified in the meta section of text. Populated in FULL view.
func (o ChronicleRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Full resource name for the rule. This unique identifier is generated using values provided for the URL parameters.
// Format: projects/{project}/locations/{location}/instances/{instance}/rules/{rule}
func (o ChronicleRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Output only. Indicate the rule can run in near real time live rule. If this is true, the rule uses the near real time
// live rule when the run frequency is set to LIVE.
func (o ChronicleRuleOutput) NearRealTimeLiveRuleEligible() pulumi.BoolOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.BoolOutput { return v.NearRealTimeLiveRuleEligible }).(pulumi.BoolOutput)
}

func (o ChronicleRuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. Resource names of the reference lists used in this rule. Populated in FULL view.
func (o ChronicleRuleOutput) ReferenceLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringArrayOutput { return v.ReferenceLists }).(pulumi.StringArrayOutput)
}

// Output only. The timestamp of when the rule revision was created. Populated in FULL, REVISION_METADATA_ONLY views.
func (o ChronicleRuleOutput) RevisionCreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.RevisionCreateTime }).(pulumi.StringOutput)
}

// Output only. The revision ID of the rule. A new revision is created whenever the rule text is changed in any way.
// Format: v_{10 digits}_{9 digits} Populated in REVISION_METADATA_ONLY view and FULL view.
func (o ChronicleRuleOutput) RevisionId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.RevisionId }).(pulumi.StringOutput)
}

// Rule Id is the ID of the Rule.
func (o ChronicleRuleOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

// Resource name of the DataAccessScope bound to this rule. Populated in BASIC view and FULL view. If reference lists are
// used in the rule, validations will be performed against this scope to ensure that the reference lists are compatible
// with both the user's and the rule's scopes. The scope should be in the format:
// "projects/{project}/locations/{location}/instances/{instance}/dataAccessScopes/{scope}".
func (o ChronicleRuleOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringPtrOutput { return v.Scope }).(pulumi.StringPtrOutput)
}

// Severity represents the severity level of the rule.
func (o ChronicleRuleOutput) Severities() ChronicleRuleSeverityArrayOutput {
	return o.ApplyT(func(v *ChronicleRule) ChronicleRuleSeverityArrayOutput { return v.Severities }).(ChronicleRuleSeverityArrayOutput)
}

// The YARA-L content of the rule. Populated in FULL view.
func (o ChronicleRuleOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringPtrOutput { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ChronicleRuleOutput) Timeouts() ChronicleRuleTimeoutsPtrOutput {
	return o.ApplyT(func(v *ChronicleRule) ChronicleRuleTimeoutsPtrOutput { return v.Timeouts }).(ChronicleRuleTimeoutsPtrOutput)
}

// Possible values: RULE_TYPE_UNSPECIFIED SINGLE_EVENT MULTI_EVENT
func (o ChronicleRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ChronicleRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ChronicleRuleInput)(nil)).Elem(), &ChronicleRule{})
	pulumi.RegisterOutputType(ChronicleRuleOutput{})
}
