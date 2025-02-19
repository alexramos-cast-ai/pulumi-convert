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

type ApiGatewayApiConfig struct {
	pulumi.CustomResourceState

	// The API to attach the config to.
	Api pulumi.StringOutput `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringOutput `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
	// chosen for the name.
	ApiConfigIdPrefix     pulumi.StringOutput `pulumi:"apiConfigIdPrefix"`
	ApiGatewayApiConfigId pulumi.StringOutput `pulumi:"apiGatewayApiConfigId"`
	// A user-visible name for the API.
	DisplayName     pulumi.StringOutput    `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
	// authentication using the default compute service account
	GatewayConfig ApiGatewayApiConfigGatewayConfigPtrOutput `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices ApiGatewayApiConfigGrpcServiceArrayOutput `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
	// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
	// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
	// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
	// Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs ApiGatewayApiConfigManagedServiceConfigArrayOutput `pulumi:"managedServiceConfigs"`
	// The resource name of the API Config.
	Name pulumi.StringOutput `pulumi:"name"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments ApiGatewayApiConfigOpenapiDocumentArrayOutput `pulumi:"openapiDocuments"`
	Project          pulumi.StringOutput                           `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringOutput `pulumi:"serviceConfigId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput               `pulumi:"terraformLabels"`
	Timeouts        ApiGatewayApiConfigTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApiGatewayApiConfig registers a new resource with the given unique name, arguments, and options.
func NewApiGatewayApiConfig(ctx *pulumi.Context,
	name string, args *ApiGatewayApiConfigArgs, opts ...pulumi.ResourceOption) (*ApiGatewayApiConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Api == nil {
		return nil, errors.New("invalid value for required argument 'Api'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApiGatewayApiConfig
	err = ctx.RegisterPackageResource("google-beta:index/apiGatewayApiConfig:ApiGatewayApiConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApiGatewayApiConfig gets an existing ApiGatewayApiConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApiGatewayApiConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApiGatewayApiConfigState, opts ...pulumi.ResourceOption) (*ApiGatewayApiConfig, error) {
	var resource ApiGatewayApiConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/apiGatewayApiConfig:ApiGatewayApiConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApiGatewayApiConfig resources.
type apiGatewayApiConfigState struct {
	// The API to attach the config to.
	Api *string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
	// chosen for the name.
	ApiConfigIdPrefix     *string `pulumi:"apiConfigIdPrefix"`
	ApiGatewayApiConfigId *string `pulumi:"apiGatewayApiConfigId"`
	// A user-visible name for the API.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
	// authentication using the default compute service account
	GatewayConfig *ApiGatewayApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices []ApiGatewayApiConfigGrpcService `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
	// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
	// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
	// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
	// Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs []ApiGatewayApiConfigManagedServiceConfig `pulumi:"managedServiceConfigs"`
	// The resource name of the API Config.
	Name *string `pulumi:"name"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments []ApiGatewayApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	Project          *string                              `pulumi:"project"`
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId *string `pulumi:"serviceConfigId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string            `pulumi:"terraformLabels"`
	Timeouts        *ApiGatewayApiConfigTimeouts `pulumi:"timeouts"`
}

type ApiGatewayApiConfigState struct {
	// The API to attach the config to.
	Api pulumi.StringPtrInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
	// chosen for the name.
	ApiConfigIdPrefix     pulumi.StringPtrInput
	ApiGatewayApiConfigId pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
	// authentication using the default compute service account
	GatewayConfig ApiGatewayApiConfigGatewayConfigPtrInput
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices ApiGatewayApiConfigGrpcServiceArrayInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
	// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
	// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
	// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
	// Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs ApiGatewayApiConfigManagedServiceConfigArrayInput
	// The resource name of the API Config.
	Name pulumi.StringPtrInput
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments ApiGatewayApiConfigOpenapiDocumentArrayInput
	Project          pulumi.StringPtrInput
	// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
	ServiceConfigId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        ApiGatewayApiConfigTimeoutsPtrInput
}

func (ApiGatewayApiConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayApiConfigState)(nil)).Elem()
}

type apiGatewayApiConfigArgs struct {
	// The API to attach the config to.
	Api string `pulumi:"api"`
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId *string `pulumi:"apiConfigId"`
	// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
	// chosen for the name.
	ApiConfigIdPrefix     *string `pulumi:"apiConfigIdPrefix"`
	ApiGatewayApiConfigId *string `pulumi:"apiGatewayApiConfigId"`
	// A user-visible name for the API.
	DisplayName *string `pulumi:"displayName"`
	// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
	// authentication using the default compute service account
	GatewayConfig *ApiGatewayApiConfigGatewayConfig `pulumi:"gatewayConfig"`
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices []ApiGatewayApiConfigGrpcService `pulumi:"grpcServices"`
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels map[string]string `pulumi:"labels"`
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
	// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
	// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
	// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
	// Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs []ApiGatewayApiConfigManagedServiceConfig `pulumi:"managedServiceConfigs"`
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments []ApiGatewayApiConfigOpenapiDocument `pulumi:"openapiDocuments"`
	Project          *string                              `pulumi:"project"`
	Timeouts         *ApiGatewayApiConfigTimeouts         `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApiGatewayApiConfig resource.
type ApiGatewayApiConfigArgs struct {
	// The API to attach the config to.
	Api pulumi.StringInput
	// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
	ApiConfigId pulumi.StringPtrInput
	// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
	// chosen for the name.
	ApiConfigIdPrefix     pulumi.StringPtrInput
	ApiGatewayApiConfigId pulumi.StringPtrInput
	// A user-visible name for the API.
	DisplayName pulumi.StringPtrInput
	// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
	// authentication using the default compute service account
	GatewayConfig ApiGatewayApiConfigGatewayConfigPtrInput
	// gRPC service definition files. If specified, openapiDocuments must not be included.
	GrpcServices ApiGatewayApiConfigGrpcServiceArrayInput
	// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
	// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
	// resource.
	Labels pulumi.StringMapInput
	// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
	// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
	// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
	// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
	// Singular embedded messages are merged using these rules for nested fields.
	ManagedServiceConfigs ApiGatewayApiConfigManagedServiceConfigArrayInput
	// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
	OpenapiDocuments ApiGatewayApiConfigOpenapiDocumentArrayInput
	Project          pulumi.StringPtrInput
	Timeouts         ApiGatewayApiConfigTimeoutsPtrInput
}

func (ApiGatewayApiConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apiGatewayApiConfigArgs)(nil)).Elem()
}

type ApiGatewayApiConfigInput interface {
	pulumi.Input

	ToApiGatewayApiConfigOutput() ApiGatewayApiConfigOutput
	ToApiGatewayApiConfigOutputWithContext(ctx context.Context) ApiGatewayApiConfigOutput
}

func (*ApiGatewayApiConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayApiConfig)(nil)).Elem()
}

func (i *ApiGatewayApiConfig) ToApiGatewayApiConfigOutput() ApiGatewayApiConfigOutput {
	return i.ToApiGatewayApiConfigOutputWithContext(context.Background())
}

func (i *ApiGatewayApiConfig) ToApiGatewayApiConfigOutputWithContext(ctx context.Context) ApiGatewayApiConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiGatewayApiConfigOutput)
}

type ApiGatewayApiConfigOutput struct{ *pulumi.OutputState }

func (ApiGatewayApiConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiGatewayApiConfig)(nil)).Elem()
}

func (o ApiGatewayApiConfigOutput) ToApiGatewayApiConfigOutput() ApiGatewayApiConfigOutput {
	return o
}

func (o ApiGatewayApiConfigOutput) ToApiGatewayApiConfigOutputWithContext(ctx context.Context) ApiGatewayApiConfigOutput {
	return o
}

// The API to attach the config to.
func (o ApiGatewayApiConfigOutput) Api() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.Api }).(pulumi.StringOutput)
}

// Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
func (o ApiGatewayApiConfigOutput) ApiConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.ApiConfigId }).(pulumi.StringOutput)
}

// Creates a unique name beginning with the specified prefix. If this and api_config_id are unspecified, a random value is
// chosen for the name.
func (o ApiGatewayApiConfigOutput) ApiConfigIdPrefix() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.ApiConfigIdPrefix }).(pulumi.StringOutput)
}

func (o ApiGatewayApiConfigOutput) ApiGatewayApiConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.ApiGatewayApiConfigId }).(pulumi.StringOutput)
}

// A user-visible name for the API.
func (o ApiGatewayApiConfigOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ApiGatewayApiConfigOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Immutable. Gateway specific configuration. If not specified, backend authentication will be set to use OIDC
// authentication using the default compute service account
func (o ApiGatewayApiConfigOutput) GatewayConfig() ApiGatewayApiConfigGatewayConfigPtrOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) ApiGatewayApiConfigGatewayConfigPtrOutput { return v.GatewayConfig }).(ApiGatewayApiConfigGatewayConfigPtrOutput)
}

// gRPC service definition files. If specified, openapiDocuments must not be included.
func (o ApiGatewayApiConfigOutput) GrpcServices() ApiGatewayApiConfigGrpcServiceArrayOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) ApiGatewayApiConfigGrpcServiceArrayOutput { return v.GrpcServices }).(ApiGatewayApiConfigGrpcServiceArrayOutput)
}

// Resource labels to represent user-provided metadata. **Note**: This field is non-authoritative, and will only manage the
// labels present in your configuration. Please refer to the field 'effective_labels' for all of the labels present on the
// resource.
func (o ApiGatewayApiConfigOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See
// https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file
// contents. If multiple files are specified, the files are merged with the following rules: * All singular scalar fields
// are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. *
// Singular embedded messages are merged using these rules for nested fields.
func (o ApiGatewayApiConfigOutput) ManagedServiceConfigs() ApiGatewayApiConfigManagedServiceConfigArrayOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) ApiGatewayApiConfigManagedServiceConfigArrayOutput {
		return v.ManagedServiceConfigs
	}).(ApiGatewayApiConfigManagedServiceConfigArrayOutput)
}

// The resource name of the API Config.
func (o ApiGatewayApiConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
func (o ApiGatewayApiConfigOutput) OpenapiDocuments() ApiGatewayApiConfigOpenapiDocumentArrayOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) ApiGatewayApiConfigOpenapiDocumentArrayOutput { return v.OpenapiDocuments }).(ApiGatewayApiConfigOpenapiDocumentArrayOutput)
}

func (o ApiGatewayApiConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
func (o ApiGatewayApiConfigOutput) ServiceConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringOutput { return v.ServiceConfigId }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o ApiGatewayApiConfigOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o ApiGatewayApiConfigOutput) Timeouts() ApiGatewayApiConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApiGatewayApiConfig) ApiGatewayApiConfigTimeoutsPtrOutput { return v.Timeouts }).(ApiGatewayApiConfigTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApiGatewayApiConfigInput)(nil)).Elem(), &ApiGatewayApiConfig{})
	pulumi.RegisterOutputType(ApiGatewayApiConfigOutput{})
}
