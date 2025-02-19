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

type AccessContextManagerServicePerimeterIngressPolicy struct {
	pulumi.CustomResourceState

	AccessContextManagerServicePerimeterIngressPolicyId pulumi.StringOutput `pulumi:"accessContextManagerServicePerimeterIngressPolicyId"`
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId pulumi.StringOutput `pulumi:"accessPolicyId"`
	// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
	IngressFrom AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrOutput `pulumi:"ingressFrom"`
	// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
	IngressTo AccessContextManagerServicePerimeterIngressPolicyIngressToPtrOutput `pulumi:"ingressTo"`
	// The name of the Service Perimeter to add this resource to.
	Perimeter pulumi.StringOutput                                                `pulumi:"perimeter"`
	Timeouts  AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewAccessContextManagerServicePerimeterIngressPolicy registers a new resource with the given unique name, arguments, and options.
func NewAccessContextManagerServicePerimeterIngressPolicy(ctx *pulumi.Context,
	name string, args *AccessContextManagerServicePerimeterIngressPolicyArgs, opts ...pulumi.ResourceOption) (*AccessContextManagerServicePerimeterIngressPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Perimeter == nil {
		return nil, errors.New("invalid value for required argument 'Perimeter'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource AccessContextManagerServicePerimeterIngressPolicy
	err = ctx.RegisterPackageResource("google-beta:index/accessContextManagerServicePerimeterIngressPolicy:AccessContextManagerServicePerimeterIngressPolicy", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessContextManagerServicePerimeterIngressPolicy gets an existing AccessContextManagerServicePerimeterIngressPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessContextManagerServicePerimeterIngressPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessContextManagerServicePerimeterIngressPolicyState, opts ...pulumi.ResourceOption) (*AccessContextManagerServicePerimeterIngressPolicy, error) {
	var resource AccessContextManagerServicePerimeterIngressPolicy
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/accessContextManagerServicePerimeterIngressPolicy:AccessContextManagerServicePerimeterIngressPolicy", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessContextManagerServicePerimeterIngressPolicy resources.
type accessContextManagerServicePerimeterIngressPolicyState struct {
	AccessContextManagerServicePerimeterIngressPolicyId *string `pulumi:"accessContextManagerServicePerimeterIngressPolicyId"`
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId *string `pulumi:"accessPolicyId"`
	// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
	IngressFrom *AccessContextManagerServicePerimeterIngressPolicyIngressFrom `pulumi:"ingressFrom"`
	// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
	IngressTo *AccessContextManagerServicePerimeterIngressPolicyIngressTo `pulumi:"ingressTo"`
	// The name of the Service Perimeter to add this resource to.
	Perimeter *string                                                    `pulumi:"perimeter"`
	Timeouts  *AccessContextManagerServicePerimeterIngressPolicyTimeouts `pulumi:"timeouts"`
}

type AccessContextManagerServicePerimeterIngressPolicyState struct {
	AccessContextManagerServicePerimeterIngressPolicyId pulumi.StringPtrInput
	// The name of the Access Policy this resource belongs to.
	AccessPolicyId pulumi.StringPtrInput
	// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
	IngressFrom AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrInput
	// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
	IngressTo AccessContextManagerServicePerimeterIngressPolicyIngressToPtrInput
	// The name of the Service Perimeter to add this resource to.
	Perimeter pulumi.StringPtrInput
	Timeouts  AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrInput
}

func (AccessContextManagerServicePerimeterIngressPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerServicePerimeterIngressPolicyState)(nil)).Elem()
}

type accessContextManagerServicePerimeterIngressPolicyArgs struct {
	AccessContextManagerServicePerimeterIngressPolicyId *string `pulumi:"accessContextManagerServicePerimeterIngressPolicyId"`
	// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
	IngressFrom *AccessContextManagerServicePerimeterIngressPolicyIngressFrom `pulumi:"ingressFrom"`
	// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
	IngressTo *AccessContextManagerServicePerimeterIngressPolicyIngressTo `pulumi:"ingressTo"`
	// The name of the Service Perimeter to add this resource to.
	Perimeter string                                                     `pulumi:"perimeter"`
	Timeouts  *AccessContextManagerServicePerimeterIngressPolicyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a AccessContextManagerServicePerimeterIngressPolicy resource.
type AccessContextManagerServicePerimeterIngressPolicyArgs struct {
	AccessContextManagerServicePerimeterIngressPolicyId pulumi.StringPtrInput
	// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
	IngressFrom AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrInput
	// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
	IngressTo AccessContextManagerServicePerimeterIngressPolicyIngressToPtrInput
	// The name of the Service Perimeter to add this resource to.
	Perimeter pulumi.StringInput
	Timeouts  AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrInput
}

func (AccessContextManagerServicePerimeterIngressPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessContextManagerServicePerimeterIngressPolicyArgs)(nil)).Elem()
}

type AccessContextManagerServicePerimeterIngressPolicyInput interface {
	pulumi.Input

	ToAccessContextManagerServicePerimeterIngressPolicyOutput() AccessContextManagerServicePerimeterIngressPolicyOutput
	ToAccessContextManagerServicePerimeterIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerServicePerimeterIngressPolicyOutput
}

func (*AccessContextManagerServicePerimeterIngressPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerServicePerimeterIngressPolicy)(nil)).Elem()
}

func (i *AccessContextManagerServicePerimeterIngressPolicy) ToAccessContextManagerServicePerimeterIngressPolicyOutput() AccessContextManagerServicePerimeterIngressPolicyOutput {
	return i.ToAccessContextManagerServicePerimeterIngressPolicyOutputWithContext(context.Background())
}

func (i *AccessContextManagerServicePerimeterIngressPolicy) ToAccessContextManagerServicePerimeterIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerServicePerimeterIngressPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AccessContextManagerServicePerimeterIngressPolicyOutput)
}

type AccessContextManagerServicePerimeterIngressPolicyOutput struct{ *pulumi.OutputState }

func (AccessContextManagerServicePerimeterIngressPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessContextManagerServicePerimeterIngressPolicy)(nil)).Elem()
}

func (o AccessContextManagerServicePerimeterIngressPolicyOutput) ToAccessContextManagerServicePerimeterIngressPolicyOutput() AccessContextManagerServicePerimeterIngressPolicyOutput {
	return o
}

func (o AccessContextManagerServicePerimeterIngressPolicyOutput) ToAccessContextManagerServicePerimeterIngressPolicyOutputWithContext(ctx context.Context) AccessContextManagerServicePerimeterIngressPolicyOutput {
	return o
}

func (o AccessContextManagerServicePerimeterIngressPolicyOutput) AccessContextManagerServicePerimeterIngressPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) pulumi.StringOutput {
		return v.AccessContextManagerServicePerimeterIngressPolicyId
	}).(pulumi.StringOutput)
}

// The name of the Access Policy this resource belongs to.
func (o AccessContextManagerServicePerimeterIngressPolicyOutput) AccessPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) pulumi.StringOutput {
		return v.AccessPolicyId
	}).(pulumi.StringOutput)
}

// Defines the conditions on the source of a request causing this 'IngressPolicy' to apply.
func (o AccessContextManagerServicePerimeterIngressPolicyOutput) IngressFrom() AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrOutput {
		return v.IngressFrom
	}).(AccessContextManagerServicePerimeterIngressPolicyIngressFromPtrOutput)
}

// Defines the conditions on the 'ApiOperation' and request destination that cause this 'IngressPolicy' to apply.
func (o AccessContextManagerServicePerimeterIngressPolicyOutput) IngressTo() AccessContextManagerServicePerimeterIngressPolicyIngressToPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) AccessContextManagerServicePerimeterIngressPolicyIngressToPtrOutput {
		return v.IngressTo
	}).(AccessContextManagerServicePerimeterIngressPolicyIngressToPtrOutput)
}

// The name of the Service Perimeter to add this resource to.
func (o AccessContextManagerServicePerimeterIngressPolicyOutput) Perimeter() pulumi.StringOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) pulumi.StringOutput { return v.Perimeter }).(pulumi.StringOutput)
}

func (o AccessContextManagerServicePerimeterIngressPolicyOutput) Timeouts() AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrOutput {
	return o.ApplyT(func(v *AccessContextManagerServicePerimeterIngressPolicy) AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrOutput {
		return v.Timeouts
	}).(AccessContextManagerServicePerimeterIngressPolicyTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessContextManagerServicePerimeterIngressPolicyInput)(nil)).Elem(), &AccessContextManagerServicePerimeterIngressPolicy{})
	pulumi.RegisterOutputType(AccessContextManagerServicePerimeterIngressPolicyOutput{})
}
