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

type DialogflowCxAgent struct {
	pulumi.CustomResourceState

	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxAgentAdvancedSettingsPtrOutput `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
	// integration.
	AvatarUri pulumi.StringPtrOutput `pulumi:"avatarUri"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringOutput `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description         pulumi.StringPtrOutput `pulumi:"description"`
	DialogflowCxAgentId pulumi.StringOutput    `pulumi:"dialogflowCxAgentId"`
	// The human-readable name of the agent, unique within the location.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrOutput `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Deprecated: Deprecated
	EnableStackdriverLogging pulumi.BoolPtrOutput `pulumi:"enableStackdriverLogging"`
	// Git integration settings for this agent.
	GitIntegrationSettings DialogflowCxAgentGitIntegrationSettingsPtrOutput `pulumi:"gitIntegrationSettings"`
	// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
	// you must configure location settings. This is a one time step but at the moment you can only [configure location
	// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location pulumi.StringOutput `pulumi:"location"`
	// The unique identifier of the agent.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
	// ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrOutput `pulumi:"securitySettings"`
	// Settings related to speech recognition.
	SpeechToTextSettings DialogflowCxAgentSpeechToTextSettingsPtrOutput `pulumi:"speechToTextSettings"`
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only
	// be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	StartFlow pulumi.StringOutput `pulumi:"startFlow"`
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes pulumi.StringArrayOutput `pulumi:"supportedLanguageCodes"`
	// Settings related to speech synthesizing.
	TextToSpeechSettings DialogflowCxAgentTextToSpeechSettingsPtrOutput `pulumi:"textToSpeechSettings"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringOutput                `pulumi:"timeZone"`
	Timeouts DialogflowCxAgentTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewDialogflowCxAgent registers a new resource with the given unique name, arguments, and options.
func NewDialogflowCxAgent(ctx *pulumi.Context,
	name string, args *DialogflowCxAgentArgs, opts ...pulumi.ResourceOption) (*DialogflowCxAgent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DefaultLanguageCode == nil {
		return nil, errors.New("invalid value for required argument 'DefaultLanguageCode'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DialogflowCxAgent
	err = ctx.RegisterPackageResource("google-beta:index/dialogflowCxAgent:DialogflowCxAgent", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDialogflowCxAgent gets an existing DialogflowCxAgent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDialogflowCxAgent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DialogflowCxAgentState, opts ...pulumi.ResourceOption) (*DialogflowCxAgent, error) {
	var resource DialogflowCxAgent
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dialogflowCxAgent:DialogflowCxAgent", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DialogflowCxAgent resources.
type dialogflowCxAgentState struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings *DialogflowCxAgentAdvancedSettings `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
	// integration.
	AvatarUri *string `pulumi:"avatarUri"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode *string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description         *string `pulumi:"description"`
	DialogflowCxAgentId *string `pulumi:"dialogflowCxAgentId"`
	// The human-readable name of the agent, unique within the location.
	DisplayName *string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Deprecated: Deprecated
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Git integration settings for this agent.
	GitIntegrationSettings *DialogflowCxAgentGitIntegrationSettings `pulumi:"gitIntegrationSettings"`
	// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
	// you must configure location settings. This is a one time step but at the moment you can only [configure location
	// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location *string `pulumi:"location"`
	// The unique identifier of the agent.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
	// ID>/securitySettings/<Security Settings ID>.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Settings related to speech recognition.
	SpeechToTextSettings *DialogflowCxAgentSpeechToTextSettings `pulumi:"speechToTextSettings"`
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only
	// be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	StartFlow *string `pulumi:"startFlow"`
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// Settings related to speech synthesizing.
	TextToSpeechSettings *DialogflowCxAgentTextToSpeechSettings `pulumi:"textToSpeechSettings"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone *string                    `pulumi:"timeZone"`
	Timeouts *DialogflowCxAgentTimeouts `pulumi:"timeouts"`
}

type DialogflowCxAgentState struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxAgentAdvancedSettingsPtrInput
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
	// integration.
	AvatarUri pulumi.StringPtrInput
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringPtrInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description         pulumi.StringPtrInput
	DialogflowCxAgentId pulumi.StringPtrInput
	// The human-readable name of the agent, unique within the location.
	DisplayName pulumi.StringPtrInput
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrInput
	// Determines whether this agent should log conversation queries.
	//
	// Deprecated: Deprecated
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Git integration settings for this agent.
	GitIntegrationSettings DialogflowCxAgentGitIntegrationSettingsPtrInput
	// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
	// you must configure location settings. This is a one time step but at the moment you can only [configure location
	// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location pulumi.StringPtrInput
	// The unique identifier of the agent.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
	// ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrInput
	// Settings related to speech recognition.
	SpeechToTextSettings DialogflowCxAgentSpeechToTextSettingsPtrInput
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only
	// be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
	// ID>.
	StartFlow pulumi.StringPtrInput
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes pulumi.StringArrayInput
	// Settings related to speech synthesizing.
	TextToSpeechSettings DialogflowCxAgentTextToSpeechSettingsPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringPtrInput
	Timeouts DialogflowCxAgentTimeoutsPtrInput
}

func (DialogflowCxAgentState) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxAgentState)(nil)).Elem()
}

type dialogflowCxAgentArgs struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings *DialogflowCxAgentAdvancedSettings `pulumi:"advancedSettings"`
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
	// integration.
	AvatarUri *string `pulumi:"avatarUri"`
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode string `pulumi:"defaultLanguageCode"`
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description         *string `pulumi:"description"`
	DialogflowCxAgentId *string `pulumi:"dialogflowCxAgentId"`
	// The human-readable name of the agent, unique within the location.
	DisplayName string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	//
	// Deprecated: Deprecated
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Git integration settings for this agent.
	GitIntegrationSettings *DialogflowCxAgentGitIntegrationSettings `pulumi:"gitIntegrationSettings"`
	// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
	// you must configure location settings. This is a one time step but at the moment you can only [configure location
	// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
	// ID>/securitySettings/<Security Settings ID>.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Settings related to speech recognition.
	SpeechToTextSettings *DialogflowCxAgentSpeechToTextSettings `pulumi:"speechToTextSettings"`
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes []string `pulumi:"supportedLanguageCodes"`
	// Settings related to speech synthesizing.
	TextToSpeechSettings *DialogflowCxAgentTextToSpeechSettings `pulumi:"textToSpeechSettings"`
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone string                     `pulumi:"timeZone"`
	Timeouts *DialogflowCxAgentTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a DialogflowCxAgent resource.
type DialogflowCxAgentArgs struct {
	// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
	// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
	AdvancedSettings DialogflowCxAgentAdvancedSettingsPtrInput
	// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
	// integration.
	AvatarUri pulumi.StringPtrInput
	// The default language of the agent as a language tag. [See Language
	// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
	// codes. This field cannot be updated after creation.
	DefaultLanguageCode pulumi.StringInput
	// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description         pulumi.StringPtrInput
	DialogflowCxAgentId pulumi.StringPtrInput
	// The human-readable name of the agent, unique within the location.
	DisplayName pulumi.StringInput
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrInput
	// Determines whether this agent should log conversation queries.
	//
	// Deprecated: Deprecated
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Git integration settings for this agent.
	GitIntegrationSettings DialogflowCxAgentGitIntegrationSettingsPtrInput
	// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
	// you must configure location settings. This is a one time step but at the moment you can only [configure location
	// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
	// Another options is to use global location so you don't need to manually configure location settings.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
	// ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrInput
	// Settings related to speech recognition.
	SpeechToTextSettings DialogflowCxAgentSpeechToTextSettingsPtrInput
	// The list of all languages supported by this agent (except for the default_language_code).
	SupportedLanguageCodes pulumi.StringArrayInput
	// Settings related to speech synthesizing.
	TextToSpeechSettings DialogflowCxAgentTextToSpeechSettingsPtrInput
	// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
	// Europe/Paris.
	TimeZone pulumi.StringInput
	Timeouts DialogflowCxAgentTimeoutsPtrInput
}

func (DialogflowCxAgentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dialogflowCxAgentArgs)(nil)).Elem()
}

type DialogflowCxAgentInput interface {
	pulumi.Input

	ToDialogflowCxAgentOutput() DialogflowCxAgentOutput
	ToDialogflowCxAgentOutputWithContext(ctx context.Context) DialogflowCxAgentOutput
}

func (*DialogflowCxAgent) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxAgent)(nil)).Elem()
}

func (i *DialogflowCxAgent) ToDialogflowCxAgentOutput() DialogflowCxAgentOutput {
	return i.ToDialogflowCxAgentOutputWithContext(context.Background())
}

func (i *DialogflowCxAgent) ToDialogflowCxAgentOutputWithContext(ctx context.Context) DialogflowCxAgentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DialogflowCxAgentOutput)
}

type DialogflowCxAgentOutput struct{ *pulumi.OutputState }

func (DialogflowCxAgentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DialogflowCxAgent)(nil)).Elem()
}

func (o DialogflowCxAgentOutput) ToDialogflowCxAgentOutput() DialogflowCxAgentOutput {
	return o
}

func (o DialogflowCxAgentOutput) ToDialogflowCxAgentOutputWithContext(ctx context.Context) DialogflowCxAgentOutput {
	return o
}

// Hierarchical advanced settings for this agent. The settings exposed at the lower level overrides the settings exposed at
// the higher level. Hierarchy: Agent->Flow->Page->Fulfillment/Parameter.
func (o DialogflowCxAgentOutput) AdvancedSettings() DialogflowCxAgentAdvancedSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) DialogflowCxAgentAdvancedSettingsPtrOutput { return v.AdvancedSettings }).(DialogflowCxAgentAdvancedSettingsPtrOutput)
}

// The URI of the agent's avatar. Avatars are used throughout the Dialogflow console and in the self-hosted Web Demo
// integration.
func (o DialogflowCxAgentOutput) AvatarUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringPtrOutput { return v.AvatarUri }).(pulumi.StringPtrOutput)
}

// The default language of the agent as a language tag. [See Language
// Support](https://cloud.google.com/dialogflow/cx/docs/reference/language) for a list of the currently supported language
// codes. This field cannot be updated after creation.
func (o DialogflowCxAgentOutput) DefaultLanguageCode() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.DefaultLanguageCode }).(pulumi.StringOutput)
}

// The description of this agent. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o DialogflowCxAgentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DialogflowCxAgentOutput) DialogflowCxAgentId() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.DialogflowCxAgentId }).(pulumi.StringOutput)
}

// The human-readable name of the agent, unique within the location.
func (o DialogflowCxAgentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Indicates if automatic spell correction is enabled in detect intent requests.
func (o DialogflowCxAgentOutput) EnableSpellCorrection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.BoolPtrOutput { return v.EnableSpellCorrection }).(pulumi.BoolPtrOutput)
}

// Determines whether this agent should log conversation queries.
//
// Deprecated: Deprecated
func (o DialogflowCxAgentOutput) EnableStackdriverLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.BoolPtrOutput { return v.EnableStackdriverLogging }).(pulumi.BoolPtrOutput)
}

// Git integration settings for this agent.
func (o DialogflowCxAgentOutput) GitIntegrationSettings() DialogflowCxAgentGitIntegrationSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) DialogflowCxAgentGitIntegrationSettingsPtrOutput {
		return v.GitIntegrationSettings
	}).(DialogflowCxAgentGitIntegrationSettingsPtrOutput)
}

// The name of the location this agent is located in. > **Note:** The first time you are deploying an Agent in your project
// you must configure location settings. This is a one time step but at the moment you can only [configure location
// settings](https://cloud.google.com/dialogflow/cx/docs/concept/region#location-settings) via the Dialogflow CX console.
// Another options is to use global location so you don't need to manually configure location settings.
func (o DialogflowCxAgentOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The unique identifier of the agent.
func (o DialogflowCxAgentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DialogflowCxAgentOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location
// ID>/securitySettings/<Security Settings ID>.
func (o DialogflowCxAgentOutput) SecuritySettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringPtrOutput { return v.SecuritySettings }).(pulumi.StringPtrOutput)
}

// Settings related to speech recognition.
func (o DialogflowCxAgentOutput) SpeechToTextSettings() DialogflowCxAgentSpeechToTextSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) DialogflowCxAgentSpeechToTextSettingsPtrOutput {
		return v.SpeechToTextSettings
	}).(DialogflowCxAgentSpeechToTextSettingsPtrOutput)
}

// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only
// be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow
// ID>.
func (o DialogflowCxAgentOutput) StartFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.StartFlow }).(pulumi.StringOutput)
}

// The list of all languages supported by this agent (except for the default_language_code).
func (o DialogflowCxAgentOutput) SupportedLanguageCodes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringArrayOutput { return v.SupportedLanguageCodes }).(pulumi.StringArrayOutput)
}

// Settings related to speech synthesizing.
func (o DialogflowCxAgentOutput) TextToSpeechSettings() DialogflowCxAgentTextToSpeechSettingsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) DialogflowCxAgentTextToSpeechSettingsPtrOutput {
		return v.TextToSpeechSettings
	}).(DialogflowCxAgentTextToSpeechSettingsPtrOutput)
}

// The time zone of this agent from the [time zone database](https://www.iana.org/time-zones), e.g., America/New_York,
// Europe/Paris.
func (o DialogflowCxAgentOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o DialogflowCxAgentOutput) Timeouts() DialogflowCxAgentTimeoutsPtrOutput {
	return o.ApplyT(func(v *DialogflowCxAgent) DialogflowCxAgentTimeoutsPtrOutput { return v.Timeouts }).(DialogflowCxAgentTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DialogflowCxAgentInput)(nil)).Elem(), &DialogflowCxAgent{})
	pulumi.RegisterOutputType(DialogflowCxAgentOutput{})
}
