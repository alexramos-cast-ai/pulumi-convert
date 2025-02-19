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

type ServiceDirectoryEndpoint struct {
	pulumi.CustomResourceState

	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrOutput `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
	// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.Float64PtrOutput `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service                    pulumi.StringOutput                       `pulumi:"service"`
	ServiceDirectoryEndpointId pulumi.StringOutput                       `pulumi:"serviceDirectoryEndpointId"`
	Timeouts                   ServiceDirectoryEndpointTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewServiceDirectoryEndpoint registers a new resource with the given unique name, arguments, and options.
func NewServiceDirectoryEndpoint(ctx *pulumi.Context,
	name string, args *ServiceDirectoryEndpointArgs, opts ...pulumi.ResourceOption) (*ServiceDirectoryEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EndpointId == nil {
		return nil, errors.New("invalid value for required argument 'EndpointId'")
	}
	if args.Service == nil {
		return nil, errors.New("invalid value for required argument 'Service'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ServiceDirectoryEndpoint
	err = ctx.RegisterPackageResource("google-beta:index/serviceDirectoryEndpoint:ServiceDirectoryEndpoint", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceDirectoryEndpoint gets an existing ServiceDirectoryEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceDirectoryEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceDirectoryEndpointState, opts ...pulumi.ResourceOption) (*ServiceDirectoryEndpoint, error) {
	var resource ServiceDirectoryEndpoint
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/serviceDirectoryEndpoint:ServiceDirectoryEndpoint", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceDirectoryEndpoint resources.
type serviceDirectoryEndpointState struct {
	// IPv4 or IPv6 address of the endpoint.
	Address *string `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	EndpointId *string `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
	// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name *string `pulumi:"name"`
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network *string `pulumi:"network"`
	// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
	Port *float64 `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service                    *string                           `pulumi:"service"`
	ServiceDirectoryEndpointId *string                           `pulumi:"serviceDirectoryEndpointId"`
	Timeouts                   *ServiceDirectoryEndpointTimeouts `pulumi:"timeouts"`
}

type ServiceDirectoryEndpointState struct {
	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	EndpointId pulumi.StringPtrInput
	// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
	// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
	Name pulumi.StringPtrInput
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network pulumi.StringPtrInput
	// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.Float64PtrInput
	// The resource name of the service that this endpoint provides.
	Service                    pulumi.StringPtrInput
	ServiceDirectoryEndpointId pulumi.StringPtrInput
	Timeouts                   ServiceDirectoryEndpointTimeoutsPtrInput
}

func (ServiceDirectoryEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDirectoryEndpointState)(nil)).Elem()
}

type serviceDirectoryEndpointArgs struct {
	// IPv4 or IPv6 address of the endpoint.
	Address *string `pulumi:"address"`
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	EndpointId string `pulumi:"endpointId"`
	// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
	// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
	Metadata map[string]string `pulumi:"metadata"`
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network *string `pulumi:"network"`
	// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
	Port *float64 `pulumi:"port"`
	// The resource name of the service that this endpoint provides.
	Service                    string                            `pulumi:"service"`
	ServiceDirectoryEndpointId *string                           `pulumi:"serviceDirectoryEndpointId"`
	Timeouts                   *ServiceDirectoryEndpointTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ServiceDirectoryEndpoint resource.
type ServiceDirectoryEndpointArgs struct {
	// IPv4 or IPv6 address of the endpoint.
	Address pulumi.StringPtrInput
	// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
	EndpointId pulumi.StringInput
	// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
	// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
	Metadata pulumi.StringMapInput
	// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
	Network pulumi.StringPtrInput
	// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
	Port pulumi.Float64PtrInput
	// The resource name of the service that this endpoint provides.
	Service                    pulumi.StringInput
	ServiceDirectoryEndpointId pulumi.StringPtrInput
	Timeouts                   ServiceDirectoryEndpointTimeoutsPtrInput
}

func (ServiceDirectoryEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceDirectoryEndpointArgs)(nil)).Elem()
}

type ServiceDirectoryEndpointInput interface {
	pulumi.Input

	ToServiceDirectoryEndpointOutput() ServiceDirectoryEndpointOutput
	ToServiceDirectoryEndpointOutputWithContext(ctx context.Context) ServiceDirectoryEndpointOutput
}

func (*ServiceDirectoryEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDirectoryEndpoint)(nil)).Elem()
}

func (i *ServiceDirectoryEndpoint) ToServiceDirectoryEndpointOutput() ServiceDirectoryEndpointOutput {
	return i.ToServiceDirectoryEndpointOutputWithContext(context.Background())
}

func (i *ServiceDirectoryEndpoint) ToServiceDirectoryEndpointOutputWithContext(ctx context.Context) ServiceDirectoryEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceDirectoryEndpointOutput)
}

type ServiceDirectoryEndpointOutput struct{ *pulumi.OutputState }

func (ServiceDirectoryEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceDirectoryEndpoint)(nil)).Elem()
}

func (o ServiceDirectoryEndpointOutput) ToServiceDirectoryEndpointOutput() ServiceDirectoryEndpointOutput {
	return o
}

func (o ServiceDirectoryEndpointOutput) ToServiceDirectoryEndpointOutputWithContext(ctx context.Context) ServiceDirectoryEndpointOutput {
	return o
}

// IPv4 or IPv6 address of the endpoint.
func (o ServiceDirectoryEndpointOutput) Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringPtrOutput { return v.Address }).(pulumi.StringPtrOutput)
}

// The Resource ID must be 1-63 characters long, including digits, lowercase letters or the hyphen character.
func (o ServiceDirectoryEndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// Metadata for the endpoint. This data can be consumed by service clients. The entire metadata dictionary may contain up
// to 512 characters, spread across all key-value pairs. Metadata that goes beyond any these limits will be rejected.
func (o ServiceDirectoryEndpointOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The resource name for the endpoint in the format 'projects/*/locations/*/namespaces/*/services/*/endpoints/*'.
func (o ServiceDirectoryEndpointOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL to the network, such as projects/PROJECT_NUMBER/locations/global/networks/NETWORK_NAME.
func (o ServiceDirectoryEndpointOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// Port that the endpoint is running on, must be in the range of [0, 65535]. If unspecified, the default is 0.
func (o ServiceDirectoryEndpointOutput) Port() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.Float64PtrOutput { return v.Port }).(pulumi.Float64PtrOutput)
}

// The resource name of the service that this endpoint provides.
func (o ServiceDirectoryEndpointOutput) Service() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringOutput { return v.Service }).(pulumi.StringOutput)
}

func (o ServiceDirectoryEndpointOutput) ServiceDirectoryEndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) pulumi.StringOutput { return v.ServiceDirectoryEndpointId }).(pulumi.StringOutput)
}

func (o ServiceDirectoryEndpointOutput) Timeouts() ServiceDirectoryEndpointTimeoutsPtrOutput {
	return o.ApplyT(func(v *ServiceDirectoryEndpoint) ServiceDirectoryEndpointTimeoutsPtrOutput { return v.Timeouts }).(ServiceDirectoryEndpointTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceDirectoryEndpointInput)(nil)).Elem(), &ServiceDirectoryEndpoint{})
	pulumi.RegisterOutputType(ServiceDirectoryEndpointOutput{})
}
