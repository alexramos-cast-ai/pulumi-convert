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

type DialogflowCxFlow struct {
	pulumi.CustomResourceState

	// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxFlowAdvancedSettingsPtrOutput `pulumi:"advancedSettings"`
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description        pulumi.StringPtrOutput `pulumi:"description"`
	DialogflowCxFlowId pulumi.StringOutput    `pulumi:"dialogflowCxFlowId"`
	// The human-readable name of the flow.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
	// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
	// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
	// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
	// get executed, with the rest being ignored.
	EventHandlers DialogflowCxFlowEventHandlerArrayOutput `pulumi:"eventHandlers"`
	// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
	// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
	// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
	// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
	// compete to control a single Default Start Flow resource in GCP.
	IsDefaultStartFlow pulumi.BoolPtrOutput `pulumi:"isDefaultStartFlow"`
	// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
	// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// The unique identifier of the flow. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings DialogflowCxFlowNluSettingsPtrOutput `pulumi:"nluSettings"`
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent   pulumi.StringPtrOutput            `pulumi:"parent"`
	Timeouts DialogflowCxFlowTimeoutsPtrOutput `pulumi:"timeouts"`
	// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
	// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayOutput `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
	// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
	// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
	// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
	// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes DialogflowCxFlowTransitionRouteArrayOutput `pulumi:"transitionRoutes"`
}

// NewDialogflowCxFlow registers a new resource with the given unique name, arguments, and options.
func NewDialogflowCxFlow(ctx *pulumi.Context,
	name string, args *DialogflowCxFlowArgs, opts ...pulumi.ResourceOption) (*DialogflowCxFlow, error) {
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
	var resource DialogflowCxFlow
	err = ctx.RegisterPackageResource("google:index/dialogflowCxFlow:DialogflowCxFlow", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDialogflowCxFlow gets an existing DialogflowCxFlow resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDialogflowCxFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DialogflowCxFlowState, opts ...pulumi.ResourceOption) (*DialogflowCxFlow, error) {
	var resource DialogflowCxFlow
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dialogflowCxFlow:DialogflowCxFlow", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DialogflowCxFlow resources.
type dialogflowCxFlowState struct {
	// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings *DialogflowCxFlowAdvancedSettings `pulumi:"advancedSettings"`
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description        *string `pulumi:"description"`
	DialogflowCxFlowId *string `pulumi:"dialogflowCxFlowId"`
	// The human-readable name of the flow.
	DisplayName *string `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
	// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
	// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
	// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
	// get executed, with the rest being ignored.
	EventHandlers []DialogflowCxFlowEventHandler `pulumi:"eventHandlers"`
	// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
	// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
	// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
	// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
	// compete to control a single Default Start Flow resource in GCP.
	IsDefaultStartFlow *bool `pulumi:"isDefaultStartFlow"`
	// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
	// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The unique identifier of the flow. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	Name *string `pulumi:"name"`
	// NLU related settings of the flow.
	NluSettings *DialogflowCxFlowNluSettings `pulumi:"nluSettings"`
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent   *string                   `pulumi:"parent"`
	Timeouts *DialogflowCxFlowTimeouts `pulumi:"timeouts"`
	// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
	// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
	// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
	// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
	// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
	// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []DialogflowCxFlowTransitionRoute `pulumi:"transitionRoutes"`
}

type DialogflowCxFlowState struct {
	// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxFlowAdvancedSettingsPtrInput
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description        pulumi.StringPtrInput
	DialogflowCxFlowId pulumi.StringPtrInput
	// The human-readable name of the flow.
	DisplayName pulumi.StringPtrInput
	// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
	// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
	// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
	// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
	// get executed, with the rest being ignored.
	EventHandlers DialogflowCxFlowEventHandlerArrayInput
	// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
	// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
	// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
	// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
	// compete to control a single Default Start Flow resource in GCP.
	IsDefaultStartFlow pulumi.BoolPtrInput
	// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
	// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The unique identifier of the flow. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	Name pulumi.StringPtrInput
	// NLU related settings of the flow.
	NluSettings DialogflowCxFlowNluSettingsPtrInput
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent   pulumi.StringPtrInput
	Timeouts DialogflowCxFlowTimeoutsPtrInput
	// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
	// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
	// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
	// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
	// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
	// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes DialogflowCxFlowTransitionRouteArrayInput
}

func (DialogflowCxFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxFlowState)(nil)).Elem()
}

type dialogflowCxFlowArgs struct {
	// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings *DialogflowCxFlowAdvancedSettings `pulumi:"advancedSettings"`
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description        *string `pulumi:"description"`
	DialogflowCxFlowId *string `pulumi:"dialogflowCxFlowId"`
	// The human-readable name of the flow.
	DisplayName string `pulumi:"displayName"`
	// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
	// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
	// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
	// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
	// get executed, with the rest being ignored.
	EventHandlers []DialogflowCxFlowEventHandler `pulumi:"eventHandlers"`
	// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
	// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
	// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
	// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
	// compete to control a single Default Start Flow resource in GCP.
	IsDefaultStartFlow *bool `pulumi:"isDefaultStartFlow"`
	// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
	// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// NLU related settings of the flow.
	NluSettings *DialogflowCxFlowNluSettings `pulumi:"nluSettings"`
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent   *string                   `pulumi:"parent"`
	Timeouts *DialogflowCxFlowTimeouts `pulumi:"timeouts"`
	// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
	// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups []string `pulumi:"transitionRouteGroups"`
	// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
	// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
	// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
	// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
	// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes []DialogflowCxFlowTransitionRoute `pulumi:"transitionRoutes"`
}

// The set of arguments for constructing a DialogflowCxFlow resource.
type DialogflowCxFlowArgs struct {
	// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxFlowAdvancedSettingsPtrInput
	// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description        pulumi.StringPtrInput
	DialogflowCxFlowId pulumi.StringPtrInput
	// The human-readable name of the flow.
	DisplayName pulumi.StringInput
	// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
	// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
	// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
	// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
	// get executed, with the rest being ignored.
	EventHandlers DialogflowCxFlowEventHandlerArrayInput
	// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
	// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
	// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
	// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
	// compete to control a single Default Start Flow resource in GCP.
	IsDefaultStartFlow pulumi.BoolPtrInput
	// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
	// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
	// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
	// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// NLU related settings of the flow.
	NluSettings DialogflowCxFlowNluSettingsPtrInput
	// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent   pulumi.StringPtrInput
	Timeouts DialogflowCxFlowTimeoutsPtrInput
	// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
	// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
	// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
	TransitionRouteGroups pulumi.StringArrayInput
	// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
	// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
	// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
	// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
	// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
	// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
	TransitionRoutes DialogflowCxFlowTransitionRouteArrayInput
}

func (DialogflowCxFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxFlowArgs)(nil)).Elem()
}

type DialogflowCxFlowInput interface {
	pulumi.Input

	ToDialogflowCxFlowOutput() DialogflowCxFlowOutput
	ToDialogflowCxFlowOutputWithContext(ctx context.Context) DialogflowCxFlowOutput
}

func (*DialogflowCxFlow) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxFlow)(nil)).Elem()
}

func (i *DialogflowCxFlow) ToDialogflowCxFlowOutput() DialogflowCxFlowOutput {
	return i.ToDialogflowCxFlowOutputWithContext(context.Background())
}

func (i *DialogflowCxFlow) ToDialogflowCxFlowOutputWithContext(ctx context.Context) DialogflowCxFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DialogflowCxFlowOutput)
}

type DialogflowCxFlowOutput struct{ *pulumi.OutputState }

func (DialogflowCxFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxFlow)(nil)).Elem()
}

func (o DialogflowCxFlowOutput) ToDialogflowCxFlowOutput() DialogflowCxFlowOutput {
	return o
}

func (o DialogflowCxFlowOutput) ToDialogflowCxFlowOutputWithContext(ctx context.Context) DialogflowCxFlowOutput {
	return o
}

// Hierarchical advanced settings for this flow. The settings exposed at the lower level overrides the settings exposed at
// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
func (o DialogflowCxFlowOutput) AdvancedSettings() DialogflowCxFlowAdvancedSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) DialogflowCxFlowAdvancedSettingsPtrOutput { return v.AdvancedSettings }).(DialogflowCxFlowAdvancedSettingsPtrOutput)
}

// The description of the flow. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o DialogflowCxFlowOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DialogflowCxFlowOutput) DialogflowCxFlowId() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringOutput { return v.DialogflowCxFlowId }).(pulumi.StringOutput)
}

// The human-readable name of the flow.
func (o DialogflowCxFlowOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// A flow's event handlers serve two purposes: They are responsible for handling events (e.g. no match, webhook errors) in
// the flow. They are inherited by every page's [event handlers][Page.event_handlers], which can be used to handle common
// events regardless of the current page. Event handlers defined in the page have higher priority than those defined in the
// flow. Unlike transitionRoutes, these handlers are evaluated on a first-match basis. The first one that matches the event
// get executed, with the rest being ignored.
func (o DialogflowCxFlowOutput) EventHandlers() DialogflowCxFlowEventHandlerArrayOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) DialogflowCxFlowEventHandlerArrayOutput { return v.EventHandlers }).(DialogflowCxFlowEventHandlerArrayOutput)
}

// Marks this as the [Default Start Flow](https://cloud.google.com/dialogflow/cx/docs/concept/flow#start) for an agent.
// When you create an agent, the Default Start Flow is created automatically. The Default Start Flow cannot be deleted;
// deleting the 'google_dialogflow_cx_flow' resource does nothing to the underlying GCP resources. > Avoid having multiple
// 'google_dialogflow_cx_flow' resources linked to the same agent with 'is_default_start_flow = true' because they will
// compete to control a single Default Start Flow resource in GCP.
func (o DialogflowCxFlowOutput) IsDefaultStartFlow() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.BoolPtrOutput { return v.IsDefaultStartFlow }).(pulumi.BoolPtrOutput)
}

// The language of the following fields in flow: Flow.event_handlers.trigger_fulfillment.messages
// Flow.event_handlers.trigger_fulfillment.conditional_cases Flow.transition_routes.trigger_fulfillment.messages
// Flow.transition_routes.trigger_fulfillment.conditional_cases If not specified, the agent's default language is used.
// Many languages are supported. Note: languages must be enabled in the agent before they can be used.
func (o DialogflowCxFlowOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

// The unique identifier of the flow. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
// ID>.
func (o DialogflowCxFlowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NLU related settings of the flow.
func (o DialogflowCxFlowOutput) NluSettings() DialogflowCxFlowNluSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) DialogflowCxFlowNluSettingsPtrOutput { return v.NluSettings }).(DialogflowCxFlowNluSettingsPtrOutput)
}

// The agent to create a flow for. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
func (o DialogflowCxFlowOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

func (o DialogflowCxFlowOutput) Timeouts() DialogflowCxFlowTimeoutsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) DialogflowCxFlowTimeoutsPtrOutput { return v.Timeouts }).(DialogflowCxFlowTimeoutsPtrOutput)
}

// A flow's transition route group serve two purposes: They are responsible for matching the user's first utterances in the
// flow. They are inherited by every page's [transition route groups][Page.transition_route_groups]. Transition route
// groups defined in the page have higher priority than those defined in the flow. Format:projects/<Project
// ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>/transitionRouteGroups/<TransitionRouteGroup ID>.
func (o DialogflowCxFlowOutput) TransitionRouteGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) pulumi.StringArrayOutput { return v.TransitionRouteGroups }).(pulumi.StringArrayOutput)
}

// A flow's transition routes serve two purposes: They are responsible for matching the user's first utterances in the
// flow. They are inherited by every page's [transition routes][Page.transition_routes] and can support use cases such as
// the user saying "help" or "can I talk to a human?", which can be handled in a common way regardless of the current page.
// Transition routes defined in the page have higher priority than those defined in the flow. TransitionRoutes are
// evalauted in the following order: TransitionRoutes with intent specified. TransitionRoutes with only condition
// specified. TransitionRoutes with intent specified are inherited by pages in the flow.
func (o DialogflowCxFlowOutput) TransitionRoutes() DialogflowCxFlowTransitionRouteArrayOutput {
	return o.ApplyT(func(v *DialogflowCxFlow) DialogflowCxFlowTransitionRouteArrayOutput { return v.TransitionRoutes }).(DialogflowCxFlowTransitionRouteArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DialogflowCxFlowInput)(nil)).Elem(), &DialogflowCxFlow{})
	pulumi.RegisterOutputType(DialogflowCxFlowOutput{})
}
