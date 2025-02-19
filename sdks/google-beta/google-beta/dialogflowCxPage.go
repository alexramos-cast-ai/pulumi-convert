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

type DialogflowCxPage struct {
	pulumi.CustomResourceState

	// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings   DialogflowCxPageAdvancedSettingsPtrOutput `pulumi:"advancedSettings"`
	DialogflowCxPageId pulumi.StringOutput                       `pulumi:"dialogflowCxPageId"`
	// The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment DialogflowCxPageEntryFulfillmentPtrOutput `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers DialogflowCxPageEventHandlerArrayOutput `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form DialogflowCxPageFormPtrOutput `pulumi:"form"`
	// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
	// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
	// agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// The unique identifier of the page. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>/pages/<Page ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent   pulumi.StringPtrOutput            `pulumi:"parent"`
	Timeouts DialogflowCxPageTimeoutsPtrOutput `pulumi:"timeouts"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
	// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
	// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
	// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
	// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
	// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
	// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
	// defined in the transition route groups with only condition specified.
	TransitionRoutes DialogflowCxPageTransitionRouteArrayOutput `pulumi:"transitionRoutes"`
}

// NewDialogflowCxPage registers a new resource with the given unique name, arguments, and options.
func NewDialogflowCxPage(ctx *pulumi.Context,
	name string, args *DialogflowCxPageArgs, opts ...pulumi.ResourceOption) (*DialogflowCxPage, error) {
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
	var resource DialogflowCxPage
	err = ctx.RegisterPackageResource("google-beta:index/dialogflowCxPage:DialogflowCxPage", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDialogflowCxPage gets an existing DialogflowCxPage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDialogflowCxPage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DialogflowCxPageState, opts ...pulumi.ResourceOption) (*DialogflowCxPage, error) {
	var resource DialogflowCxPage
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dialogflowCxPage:DialogflowCxPage", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DialogflowCxPage resources.
type dialogflowCxPageState struct {
	// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings   *DialogflowCxPageAdvancedSettings `pulumi:"advancedSettings"`
	DialogflowCxPageId *string                           `pulumi:"dialogflowCxPageId"`
	// The human-readable name of the page, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment *DialogflowCxPageEntryFulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []DialogflowCxPageEventHandler `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form *DialogflowCxPageForm `pulumi:"form"`
	// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
	// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
	// agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The unique identifier of the page. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>/pages/<Page ID>.
	Name *string `pulumi:"name"`
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent   *string                   `pulumi:"parent"`
	Timeouts *DialogflowCxPageTimeouts `pulumi:"timeouts"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
	// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
	// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
	// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
	// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
	// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
	// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
	// defined in the transition route groups with only condition specified.
	TransitionRoutes []DialogflowCxPageTransitionRoute `pulumi:"transitionRoutes"`
}

type DialogflowCxPageState struct {
	// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings   DialogflowCxPageAdvancedSettingsPtrInput
	DialogflowCxPageId pulumi.StringPtrInput
	// The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment DialogflowCxPageEntryFulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers DialogflowCxPageEventHandlerArrayInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form DialogflowCxPageFormPtrInput
	// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
	// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
	// agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The unique identifier of the page. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>/pages/<Page ID>.
	Name pulumi.StringPtrInput
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent   pulumi.StringPtrInput
	Timeouts DialogflowCxPageTimeoutsPtrInput
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
	// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
	// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
	// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
	// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
	// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
	// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
	// defined in the transition route groups with only condition specified.
	TransitionRoutes DialogflowCxPageTransitionRouteArrayInput
}

func (DialogflowCxPageState) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxPageState)(nil)).Elem()
}

type dialogflowCxPageArgs struct {
	// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings   *DialogflowCxPageAdvancedSettings `pulumi:"advancedSettings"`
	DialogflowCxPageId *string                           `pulumi:"dialogflowCxPageId"`
	// The human-readable name of the page, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment *DialogflowCxPageEntryFulfillment `pulumi:"entryFulfillment"`
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers []DialogflowCxPageEventHandler `pulumi:"eventHandlers"`
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form *DialogflowCxPageForm `pulumi:"form"`
	// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
	// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
	// agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent   *string                   `pulumi:"parent"`
	Timeouts *DialogflowCxPageTimeouts `pulumi:"timeouts"`
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
	// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
	// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
	// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
	// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
	// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
	// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
	// defined in the transition route groups with only condition specified.
	TransitionRoutes []DialogflowCxPageTransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a DialogflowCxPage resource.
type DialogflowCxPageArgs struct {
	// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings   DialogflowCxPageAdvancedSettingsPtrInput
	DialogflowCxPageId pulumi.StringPtrInput
	// The human-readable name of the page, unique within the agent.
	DisplayName pulumi.StringInput
	// The fulfillment to call when the session is entering the page.
	EntryFulfillment DialogflowCxPageEntryFulfillmentPtrInput
	// Handlers associated with the page to handle events such as webhook errors, no match or no input.
	EventHandlers DialogflowCxPageEventHandlerArrayInput
	// The form associated with the page, used for collecting parameters relevant to the page.
	Form DialogflowCxPageFormPtrInput
	// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
	// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
	// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
	// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
	// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
	// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
	// agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	Parent   pulumi.StringPtrInput
	Timeouts DialogflowCxPageTimeoutsPtrInput
	// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
	// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
	// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
	// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
	// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
	// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
	// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
	// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
	// defined in the transition route groups with only condition specified.
	TransitionRoutes DialogflowCxPageTransitionRouteArrayInput
}

func (DialogflowCxPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxPageArgs)(nil)).Elem()
}

type DialogflowCxPageInput interface {
	pulumi.Input

	ToDialogflowCxPageOutput() DialogflowCxPageOutput
	ToDialogflowCxPageOutputWithContext(ctx context.Context) DialogflowCxPageOutput
}

func (*DialogflowCxPage) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxPage)(nil)).Elem()
}

func (i *DialogflowCxPage) ToDialogflowCxPageOutput() DialogflowCxPageOutput {
	return i.ToDialogflowCxPageOutputWithContext(context.Background())
}

func (i *DialogflowCxPage) ToDialogflowCxPageOutputWithContext(ctx context.Context) DialogflowCxPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DialogflowCxPageOutput)
}

type DialogflowCxPageOutput struct{ *pulumi.OutputState }

func (DialogflowCxPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxPage)(nil)).Elem()
}

func (o DialogflowCxPageOutput) ToDialogflowCxPageOutput() DialogflowCxPageOutput {
	return o
}

func (o DialogflowCxPageOutput) ToDialogflowCxPageOutputWithContext(ctx context.Context) DialogflowCxPageOutput {
	return o
}

// Hierarchical advanced settings for this page. The settings exposed at the lower level overrides the settings exposed at
// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
func (o DialogflowCxPageOutput) AdvancedSettings() DialogflowCxPageAdvancedSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageAdvancedSettingsPtrOutput { return v.AdvancedSettings }).(DialogflowCxPageAdvancedSettingsPtrOutput)
}

func (o DialogflowCxPageOutput) DialogflowCxPageId() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringOutput { return v.DialogflowCxPageId }).(pulumi.StringOutput)
}

// The human-readable name of the page, unique within the agent.
func (o DialogflowCxPageOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The fulfillment to call when the session is entering the page.
func (o DialogflowCxPageOutput) EntryFulfillment() DialogflowCxPageEntryFulfillmentPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageEntryFulfillmentPtrOutput { return v.EntryFulfillment }).(DialogflowCxPageEntryFulfillmentPtrOutput)
}

// Handlers associated with the page to handle events such as webhook errors, no match or no input.
func (o DialogflowCxPageOutput) EventHandlers() DialogflowCxPageEventHandlerArrayOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageEventHandlerArrayOutput { return v.EventHandlers }).(DialogflowCxPageEventHandlerArrayOutput)
}

// The form associated with the page, used for collecting parameters relevant to the page.
func (o DialogflowCxPageOutput) Form() DialogflowCxPageFormPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageFormPtrOutput { return v.Form }).(DialogflowCxPageFormPtrOutput)
}

// The language of the following fields in page: Page.entry_fulfillment.messages Page.entry_fulfillment.conditional_cases
// Page.event_handlers.trigger_fulfillment.messages Page.event_handlers.trigger_fulfillment.conditional_cases
// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.messages
// Page.form.parameters.fill_behavior.initial_prompt_fulfillment.conditional_cases
// Page.form.parameters.fill_behavior.reprompt_event_handlers.messages
// Page.form.parameters.fill_behavior.reprompt_event_handlers.conditional_cases
// Page.transition_routes.trigger_fulfillment.messages Page.transition_routes.trigger_fulfillment.conditional_cases If not
// specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the
// agent before they can be used.
func (o DialogflowCxPageOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

// The unique identifier of the page. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
// ID>/pages/<Page ID>.
func (o DialogflowCxPageOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The flow to create a page for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
func (o DialogflowCxPageOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o DialogflowCxPageOutput) Timeouts() DialogflowCxPageTimeoutsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageTimeoutsPtrOutput { return v.Timeouts }).(DialogflowCxPageTimeoutsPtrOutput)
}

// Ordered list of TransitionRouteGroups associated with the page. Transition route groups must be unique within a page. If
// multiple transition routes within a page scope refer to the same intent, then the precedence order is: page's transition
// route > page's transition route group > flow's transition routes. If multiple transition route groups within a page
// contain the same intent, then the first group in the ordered list takes precedence. Format:projects/<Project
// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
func (o DialogflowCxPageOutput) TransitionRouteGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DialogflowCxPage) pulumi.StringArrayOutput { return v.TransitionRouteGroups }).(pulumi.StringArrayOutput)
}

// A list of transitions for the transition rules of this page. They route the conversation to another page in the same
// flow, or another flow. When we are in a certain page, the TransitionRoutes are evalauted in the following order:
// TransitionRoutes defined in the page with intent specified. TransitionRoutes defined in the transition route groups with
// intent specified. TransitionRoutes defined in flow with intent specified. TransitionRoutes defined in the transition
// route groups with intent specified. TransitionRoutes defined in the page with only condition specified. TransitionRoutes
// defined in the transition route groups with only condition specified.
func (o DialogflowCxPageOutput) TransitionRoutes() DialogflowCxPageTransitionRouteArrayOutput {
	return o.ApplyT(func(v *DialogflowCxPage) DialogflowCxPageTransitionRouteArrayOutput { return v.TransitionRoutes }).(DialogflowCxPageTransitionRouteArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DialogflowCxPageInput)(nil)).Elem(), &DialogflowCxPage{})
	pulumi.RegisterOutputType(DialogflowCxPageOutput{})
}
