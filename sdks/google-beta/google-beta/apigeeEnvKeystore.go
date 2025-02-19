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

type ApigeeEnvKeystore struct {
	pulumi.CustomResourceState

	// Aliases in this keystore.
	Aliases             pulumi.StringArrayOutput `pulumi:"aliases"`
	ApigeeEnvKeystoreId pulumi.StringOutput      `pulumi:"apigeeEnvKeystoreId"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringOutput `pulumi:"envId"`
	// The name of the newly created keystore.
	Name     pulumi.StringOutput                `pulumi:"name"`
	Timeouts ApigeeEnvKeystoreTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewApigeeEnvKeystore registers a new resource with the given unique name, arguments, and options.
func NewApigeeEnvKeystore(ctx *pulumi.Context,
	name string, args *ApigeeEnvKeystoreArgs, opts ...pulumi.ResourceOption) (*ApigeeEnvKeystore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EnvId == nil {
		return nil, errors.New("invalid value for required argument 'EnvId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource ApigeeEnvKeystore
	err = ctx.RegisterPackageResource("google-beta:index/apigeeEnvKeystore:ApigeeEnvKeystore", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigeeEnvKeystore gets an existing ApigeeEnvKeystore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigeeEnvKeystore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigeeEnvKeystoreState, opts ...pulumi.ResourceOption) (*ApigeeEnvKeystore, error) {
	var resource ApigeeEnvKeystore
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/apigeeEnvKeystore:ApigeeEnvKeystore", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigeeEnvKeystore resources.
type apigeeEnvKeystoreState struct {
	// Aliases in this keystore.
	Aliases             []string `pulumi:"aliases"`
	ApigeeEnvKeystoreId *string  `pulumi:"apigeeEnvKeystoreId"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId *string `pulumi:"envId"`
	// The name of the newly created keystore.
	Name     *string                    `pulumi:"name"`
	Timeouts *ApigeeEnvKeystoreTimeouts `pulumi:"timeouts"`
}

type ApigeeEnvKeystoreState struct {
	// Aliases in this keystore.
	Aliases             pulumi.StringArrayInput
	ApigeeEnvKeystoreId pulumi.StringPtrInput
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringPtrInput
	// The name of the newly created keystore.
	Name     pulumi.StringPtrInput
	Timeouts ApigeeEnvKeystoreTimeoutsPtrInput
}

func (ApigeeEnvKeystoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvKeystoreState)(nil)).Elem()
}

type apigeeEnvKeystoreArgs struct {
	ApigeeEnvKeystoreId *string `pulumi:"apigeeEnvKeystoreId"`
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId string `pulumi:"envId"`
	// The name of the newly created keystore.
	Name     *string                    `pulumi:"name"`
	Timeouts *ApigeeEnvKeystoreTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a ApigeeEnvKeystore resource.
type ApigeeEnvKeystoreArgs struct {
	ApigeeEnvKeystoreId pulumi.StringPtrInput
	// The Apigee environment group associated with the Apigee environment, in the format
	// 'organizations/{{org_name}}/environments/{{env_name}}'.
	EnvId pulumi.StringInput
	// The name of the newly created keystore.
	Name     pulumi.StringPtrInput
	Timeouts ApigeeEnvKeystoreTimeoutsPtrInput
}

func (ApigeeEnvKeystoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigeeEnvKeystoreArgs)(nil)).Elem()
}

type ApigeeEnvKeystoreInput interface {
	pulumi.Input

	ToApigeeEnvKeystoreOutput() ApigeeEnvKeystoreOutput
	ToApigeeEnvKeystoreOutputWithContext(ctx context.Context) ApigeeEnvKeystoreOutput
}

func (*ApigeeEnvKeystore) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvKeystore)(nil)).Elem()
}

func (i *ApigeeEnvKeystore) ToApigeeEnvKeystoreOutput() ApigeeEnvKeystoreOutput {
	return i.ToApigeeEnvKeystoreOutputWithContext(context.Background())
}

func (i *ApigeeEnvKeystore) ToApigeeEnvKeystoreOutputWithContext(ctx context.Context) ApigeeEnvKeystoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigeeEnvKeystoreOutput)
}

type ApigeeEnvKeystoreOutput struct{ *pulumi.OutputState }

func (ApigeeEnvKeystoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigeeEnvKeystore)(nil)).Elem()
}

func (o ApigeeEnvKeystoreOutput) ToApigeeEnvKeystoreOutput() ApigeeEnvKeystoreOutput {
	return o
}

func (o ApigeeEnvKeystoreOutput) ToApigeeEnvKeystoreOutputWithContext(ctx context.Context) ApigeeEnvKeystoreOutput {
	return o
}

// Aliases in this keystore.
func (o ApigeeEnvKeystoreOutput) Aliases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigeeEnvKeystore) pulumi.StringArrayOutput { return v.Aliases }).(pulumi.StringArrayOutput)
}

func (o ApigeeEnvKeystoreOutput) ApigeeEnvKeystoreId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvKeystore) pulumi.StringOutput { return v.ApigeeEnvKeystoreId }).(pulumi.StringOutput)
}

// The Apigee environment group associated with the Apigee environment, in the format
// 'organizations/{{org_name}}/environments/{{env_name}}'.
func (o ApigeeEnvKeystoreOutput) EnvId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvKeystore) pulumi.StringOutput { return v.EnvId }).(pulumi.StringOutput)
}

// The name of the newly created keystore.
func (o ApigeeEnvKeystoreOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigeeEnvKeystore) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ApigeeEnvKeystoreOutput) Timeouts() ApigeeEnvKeystoreTimeoutsPtrOutput {
	return o.ApplyT(func(v *ApigeeEnvKeystore) ApigeeEnvKeystoreTimeoutsPtrOutput { return v.Timeouts }).(ApigeeEnvKeystoreTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigeeEnvKeystoreInput)(nil)).Elem(), &ApigeeEnvKeystore{})
	pulumi.RegisterOutputType(ApigeeEnvKeystoreOutput{})
}
