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

type DialogflowCxTestCase struct {
	pulumi.CustomResourceState

	// When the test was created. A timestamp in RFC3339 text format.
	CreationTime           pulumi.StringOutput `pulumi:"creationTime"`
	DialogflowCxTestCaseId pulumi.StringOutput `pulumi:"dialogflowCxTestCaseId"`
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The latest test result.
	LastTestResults DialogflowCxTestCaseLastTestResultArrayOutput `pulumi:"lastTestResults"`
	// The unique identifier of the test case. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/testCases/<TestCase ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes pulumi.StringPtrOutput `pulumi:"notes"`
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
	// should start with "#" and has a limit of 30 characters
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
	// of agent utterances that should occur when the agent is working properly.
	TestCaseConversationTurns DialogflowCxTestCaseTestCaseConversationTurnArrayOutput `pulumi:"testCaseConversationTurns"`
	// Config for the test case.
	TestConfig DialogflowCxTestCaseTestConfigPtrOutput `pulumi:"testConfig"`
	Timeouts   DialogflowCxTestCaseTimeoutsPtrOutput   `pulumi:"timeouts"`
}

// NewDialogflowCxTestCase registers a new resource with the given unique name, arguments, and options.
func NewDialogflowCxTestCase(ctx *pulumi.Context,
	name string, args *DialogflowCxTestCaseArgs, opts ...pulumi.ResourceOption) (*DialogflowCxTestCase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DialogflowCxTestCase
	err = ctx.RegisterPackageResource("google:index/dialogflowCxTestCase:DialogflowCxTestCase", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDialogflowCxTestCase gets an existing DialogflowCxTestCase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDialogflowCxTestCase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DialogflowCxTestCaseState, opts ...pulumi.ResourceOption) (*DialogflowCxTestCase, error) {
	var resource DialogflowCxTestCase
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dialogflowCxTestCase:DialogflowCxTestCase", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DialogflowCxTestCase resources.
type dialogflowCxTestCaseState struct {
	// When the test was created. A timestamp in RFC3339 text format.
	CreationTime           *string `pulumi:"creationTime"`
	DialogflowCxTestCaseId *string `pulumi:"dialogflowCxTestCaseId"`
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	DisplayName *string `pulumi:"displayName"`
	// The latest test result.
	LastTestResults []DialogflowCxTestCaseLastTestResult `pulumi:"lastTestResults"`
	// The unique identifier of the test case. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/testCases/<TestCase ID>.
	Name *string `pulumi:"name"`
	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes *string `pulumi:"notes"`
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
	// should start with "#" and has a limit of 30 characters
	Tags []string `pulumi:"tags"`
	// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
	// of agent utterances that should occur when the agent is working properly.
	TestCaseConversationTurns []DialogflowCxTestCaseTestCaseConversationTurn `pulumi:"testCaseConversationTurns"`
	// Config for the test case.
	TestConfig *DialogflowCxTestCaseTestConfig `pulumi:"testConfig"`
	Timeouts   *DialogflowCxTestCaseTimeouts   `pulumi:"timeouts"`
}

type DialogflowCxTestCaseState struct {
	// When the test was created. A timestamp in RFC3339 text format.
	CreationTime           pulumi.StringPtrInput
	DialogflowCxTestCaseId pulumi.StringPtrInput
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	DisplayName pulumi.StringPtrInput
	// The latest test result.
	LastTestResults DialogflowCxTestCaseLastTestResultArrayInput
	// The unique identifier of the test case. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/testCases/<TestCase ID>.
	Name pulumi.StringPtrInput
	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes pulumi.StringPtrInput
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
	// should start with "#" and has a limit of 30 characters
	Tags pulumi.StringArrayInput
	// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
	// of agent utterances that should occur when the agent is working properly.
	TestCaseConversationTurns DialogflowCxTestCaseTestCaseConversationTurnArrayInput
	// Config for the test case.
	TestConfig DialogflowCxTestCaseTestConfigPtrInput
	Timeouts   DialogflowCxTestCaseTimeoutsPtrInput
}

func (DialogflowCxTestCaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxTestCaseState)(nil)).Elem()
}

type dialogflowCxTestCaseArgs struct {
	DialogflowCxTestCaseId *string `pulumi:"dialogflowCxTestCaseId"`
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	DisplayName string `pulumi:"displayName"`
	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes *string `pulumi:"notes"`
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
	// should start with "#" and has a limit of 30 characters
	Tags []string `pulumi:"tags"`
	// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
	// of agent utterances that should occur when the agent is working properly.
	TestCaseConversationTurns []DialogflowCxTestCaseTestCaseConversationTurn `pulumi:"testCaseConversationTurns"`
	// Config for the test case.
	TestConfig *DialogflowCxTestCaseTestConfig `pulumi:"testConfig"`
	Timeouts   *DialogflowCxTestCaseTimeouts   `pulumi:"timeouts"`
}

// The set of arguments for constructing a DialogflowCxTestCase resource.
type DialogflowCxTestCaseArgs struct {
	DialogflowCxTestCaseId pulumi.StringPtrInput
	// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
	DisplayName pulumi.StringInput
	// Additional freeform notes about the test case. Limit of 400 characters.
	Notes pulumi.StringPtrInput
	// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
	// should start with "#" and has a limit of 30 characters
	Tags pulumi.StringArrayInput
	// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
	// of agent utterances that should occur when the agent is working properly.
	TestCaseConversationTurns DialogflowCxTestCaseTestCaseConversationTurnArrayInput
	// Config for the test case.
	TestConfig DialogflowCxTestCaseTestConfigPtrInput
	Timeouts   DialogflowCxTestCaseTimeoutsPtrInput
}

func (DialogflowCxTestCaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxTestCaseArgs)(nil)).Elem()
}

type DialogflowCxTestCaseInput interface {
	pulumi.Input

	ToDialogflowCxTestCaseOutput() DialogflowCxTestCaseOutput
	ToDialogflowCxTestCaseOutputWithContext(ctx context.Context) DialogflowCxTestCaseOutput
}

func (*DialogflowCxTestCase) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxTestCase)(nil)).Elem()
}

func (i *DialogflowCxTestCase) ToDialogflowCxTestCaseOutput() DialogflowCxTestCaseOutput {
	return i.ToDialogflowCxTestCaseOutputWithContext(context.Background())
}

func (i *DialogflowCxTestCase) ToDialogflowCxTestCaseOutputWithContext(ctx context.Context) DialogflowCxTestCaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DialogflowCxTestCaseOutput)
}

type DialogflowCxTestCaseOutput struct{ *pulumi.OutputState }

func (DialogflowCxTestCaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxTestCase)(nil)).Elem()
}

func (o DialogflowCxTestCaseOutput) ToDialogflowCxTestCaseOutput() DialogflowCxTestCaseOutput {
	return o
}

func (o DialogflowCxTestCaseOutput) ToDialogflowCxTestCaseOutputWithContext(ctx context.Context) DialogflowCxTestCaseOutput {
	return o
}

// When the test was created. A timestamp in RFC3339 text format.
func (o DialogflowCxTestCaseOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o DialogflowCxTestCaseOutput) DialogflowCxTestCaseId() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringOutput { return v.DialogflowCxTestCaseId }).(pulumi.StringOutput)
}

// The human-readable name of the test case, unique within the agent. Limit of 200 characters.
func (o DialogflowCxTestCaseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The latest test result.
func (o DialogflowCxTestCaseOutput) LastTestResults() DialogflowCxTestCaseLastTestResultArrayOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) DialogflowCxTestCaseLastTestResultArrayOutput { return v.LastTestResults }).(DialogflowCxTestCaseLastTestResultArrayOutput)
}

// The unique identifier of the test case. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
// ID>/testCases/<TestCase ID>.
func (o DialogflowCxTestCaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Additional freeform notes about the test case. Limit of 400 characters.
func (o DialogflowCxTestCaseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

// The agent to create the test case for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
func (o DialogflowCxTestCaseOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Tags are short descriptions that users may apply to test cases for organizational and filtering purposes. Each tag
// should start with "#" and has a limit of 30 characters
func (o DialogflowCxTestCaseOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

// The conversation turns uttered when the test case was created, in chronological order. These include the canonical set
// of agent utterances that should occur when the agent is working properly.
func (o DialogflowCxTestCaseOutput) TestCaseConversationTurns() DialogflowCxTestCaseTestCaseConversationTurnArrayOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) DialogflowCxTestCaseTestCaseConversationTurnArrayOutput {
		return v.TestCaseConversationTurns
	}).(DialogflowCxTestCaseTestCaseConversationTurnArrayOutput)
}

// Config for the test case.
func (o DialogflowCxTestCaseOutput) TestConfig() DialogflowCxTestCaseTestConfigPtrOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) DialogflowCxTestCaseTestConfigPtrOutput { return v.TestConfig }).(DialogflowCxTestCaseTestConfigPtrOutput)
}

func (o DialogflowCxTestCaseOutput) Timeouts() DialogflowCxTestCaseTimeoutsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxTestCase) DialogflowCxTestCaseTimeoutsPtrOutput { return v.Timeouts }).(DialogflowCxTestCaseTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DialogflowCxTestCaseInput)(nil)).Elem(), &DialogflowCxTestCase{})
	pulumi.RegisterOutputType(DialogflowCxTestCaseOutput{})
}
