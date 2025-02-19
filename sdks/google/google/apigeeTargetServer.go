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

type ApigeeTargetServer struct {
	pulumi.CustomResourceState

	ApigeeTargetServerId pulumi.StringOutput `pulumi:"apigeeTargetServerId"`
	// A human-readable description of this TargetServer.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringOutput `pulumi:"envId"`
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host pulumi.StringOutput `pulumi:"host"`
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
	// more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port pulumi.Float64Output `pulumi:"port"`
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
	// "EXTERNAL_CALLOUT"]
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
	// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	SSlInfo  ApigeeTargetServerSSlInfoPtrOutput  `pulumi:"sSlInfo"`
	Timeouts ApigeeTargetServerTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApigeeTargetServer registers a new resource with the given unique name, arguments, and options.
func NewApigeeTargetServer(ctx *pulumi.Context,
	name string, args *ApigeeTargetServerArgs, opts ...pulumi.ResourceOption) (*ApigeeTargetServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	if args.Host == nil {
		return nil, errors.New("invalid value for required argument 'Host'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeTargetServer
	err = ctx.RegisterPackageResource("google:index/apigeeTargetServer:ApigeeTargetServer", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeTargetServer gets an existing ApigeeTargetServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeTargetServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeTargetServerState, opts ...pulumi.ResourceOption) (*ApigeeTargetServer, error) {
	var resource ApigeeTargetServer
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/apigeeTargetServer:ApigeeTargetServer", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeTargetServer resources.
type apigeeTargetServerState struct {
	ApigeeTargetServerId *string `pulumi:"apigeeTargetServerId"`
	// A human-readable description of this TargetServer.
	Description *string `pulumi:"description"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId *string `pulumi:"envId"`
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host *string `pulumi:"host"`
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
	// more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name *string `pulumi:"name"`
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port *float64 `pulumi:"port"`
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
	// "EXTERNAL_CALLOUT"]
	Protocol *string `pulumi:"protocol"`
	// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
	// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	SSlInfo  *ApigeeTargetServerSSlInfo  `pulumi:"sSlInfo"`
	Timeouts *ApigeeTargetServerTimeouts `pulumi:"timeouts"`
}

type ApigeeTargetServerState struct {
	ApigeeTargetServerId pulumi.StringPtrInput
	// A human-readable description of this TargetServer.
	Description pulumi.StringPtrInput
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringPtrInput
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host pulumi.StringPtrInput
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
	// more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name pulumi.StringPtrInput
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port pulumi.Float64PtrInput
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
	// "EXTERNAL_CALLOUT"]
	Protocol pulumi.StringPtrInput
	// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
	// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	SSlInfo  ApigeeTargetServerSSlInfoPtrInput
	Timeouts ApigeeTargetServerTimeoutsPtrInput
}

func (ApigeeTargetServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeTargetServerState)(nil)).Elem()
}

type apigeeTargetServerArgs struct {
	ApigeeTargetServerId *string `pulumi:"apigeeTargetServerId"`
	// A human-readable description of this TargetServer.
	Description *string `pulumi:"description"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId string `pulumi:"envId"`
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host string `pulumi:"host"`
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
	// more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled *bool `pulumi:"isEnabled"`
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name *string `pulumi:"name"`
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port float64 `pulumi:"port"`
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
	// "EXTERNAL_CALLOUT"]
	Protocol *string `pulumi:"protocol"`
	// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
	// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	SSlInfo  *ApigeeTargetServerSSlInfo  `pulumi:"sSlInfo"`
	Timeouts *ApigeeTargetServerTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApigeeTargetServer resource.
type ApigeeTargetServerArgs struct {
	ApigeeTargetServerId pulumi.StringPtrInput
	// A human-readable description of this TargetServer.
	Description pulumi.StringPtrInput
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringInput
	// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
	Host pulumi.StringInput
	// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
	// more TargetServers need to taken out of rotation periodically. Defaults to true.
	IsEnabled pulumi.BoolPtrInput
	// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
	Name pulumi.StringPtrInput
	// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
	Port pulumi.Float64Input
	// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
	// "EXTERNAL_CALLOUT"]
	Protocol pulumi.StringPtrInput
	// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
	// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
	SSlInfo  ApigeeTargetServerSSlInfoPtrInput
	Timeouts ApigeeTargetServerTimeoutsPtrInput
}

func (ApigeeTargetServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeTargetServerArgs)(nil)).Elem()
}

type ApigeeTargetServerInput interface {
	pulumi.Input

	ToApigeeTargetServerOutput() ApigeeTargetServerOutput
	ToApigeeTargetServerOutputWithContext(ctx context.Context) ApigeeTargetServerOutput
}

func (*ApigeeTargetServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeTargetServer)(nil)).Elem()
}

func (i *ApigeeTargetServer) ToApigeeTargetServerOutput() ApigeeTargetServerOutput {
	return i.ToApigeeTargetServerOutputWithContext(context.Background())
}

func (i *ApigeeTargetServer) ToApigeeTargetServerOutputWithContext(ctx context.Context) ApigeeTargetServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeTargetServerOutput)
}

type ApigeeTargetServerOutput struct{ *pulumi.OutputState }

func (ApigeeTargetServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeTargetServer)(nil)).Elem()
}

func (o ApigeeTargetServerOutput) ToApigeeTargetServerOutput() ApigeeTargetServerOutput {
	return o
}

func (o ApigeeTargetServerOutput) ToApigeeTargetServerOutputWithContext(ctx context.Context) ApigeeTargetServerOutput {
	return o
}

func (o ApigeeTargetServerOutput) ApigeeTargetServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringOutput { return v.ApigeeTargetServerId }).(pulumi.StringOutput)
}

// A human-readable description of this TargetServer.
func (o ApigeeTargetServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Apigee environment group associated with the Apigee environment, in the format
// 'organizations/{{org_name}}/environments/{{env_name}}'.
func (o ApigeeTargetServerOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

// The host name this target connects to. Value must be a valid hostname as described by RFC-1123.
func (o ApigeeTargetServerOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringOutput { return v.Host }).(pulumi.StringOutput)
}

// Enabling/disabling a TargetServer is useful when TargetServers are used in load balancing configurations, and one or
// more TargetServers need to taken out of rotation periodically. Defaults to true.
func (o ApigeeTargetServerOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// The resource id of this reference. Values must match the regular expression [\w\s-.]+.
func (o ApigeeTargetServerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port number this target connects to on the given host. Value must be between 1 and 65535, inclusive.
func (o ApigeeTargetServerOutput) Port() pulumi.Float64Output {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.Float64Output { return v.Port }).(pulumi.Float64Output)
}

// Immutable. The protocol used by this TargetServer. Possible values: ["HTTP", "HTTP2", "GRPC_TARGET", "GRPC",
// "EXTERNAL_CALLOUT"]
func (o ApigeeTargetServerOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// Specifies TLS configuration info for this TargetServer. The JSON name is sSLInfo for legacy/backwards compatibility
// reasons -- Edge originally supported SSL, and the name is still used for TLS configuration.
func (o ApigeeTargetServerOutput) SSlInfo() ApigeeTargetServerSSlInfoPtrOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) ApigeeTargetServerSSlInfoPtrOutput { return v.SSlInfo }).(ApigeeTargetServerSSlInfoPtrOutput)
}

func (o ApigeeTargetServerOutput) Timeouts() ApigeeTargetServerTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeTargetServer) ApigeeTargetServerTimeoutsPtrOutput { return v.Timeouts }).(ApigeeTargetServerTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeTargetServerInput)(nil)).Elem(), &ApigeeTargetServer{})
	pulumi.RegisterOutputType(ApigeeTargetServerOutput{})
}
