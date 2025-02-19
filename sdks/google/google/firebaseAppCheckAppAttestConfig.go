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

type FirebaseAppCheckAppAttestConfig struct {
	pulumi.CustomResourceState

	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                             pulumi.StringOutput `pulumi:"appId"`
	FirebaseAppCheckAppAttestConfigId pulumi.StringOutput `pulumi:"firebaseAppCheckAppAttestConfigId"`
	// The relative resource name of the App Attest configuration object
	Name     pulumi.StringOutput                              `pulumi:"name"`
	Project  pulumi.StringOutput                              `pulumi:"project"`
	Timeouts FirebaseAppCheckAppAttestConfigTimeoutsPtrOutput `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
	// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
	// fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringOutput `pulumi:"tokenTtl"`
}

// NewFirebaseAppCheckAppAttestConfig registers a new resource with the given unique name, arguments, and options.
func NewFirebaseAppCheckAppAttestConfig(ctx *pulumi.Context,
	name string, args *FirebaseAppCheckAppAttestConfigArgs, opts ...pulumi.ResourceOption) (*FirebaseAppCheckAppAttestConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseAppCheckAppAttestConfig
	err = ctx.RegisterPackageResource("google:index/firebaseAppCheckAppAttestConfig:FirebaseAppCheckAppAttestConfig", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseAppCheckAppAttestConfig gets an existing FirebaseAppCheckAppAttestConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseAppCheckAppAttestConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseAppCheckAppAttestConfigState, opts ...pulumi.ResourceOption) (*FirebaseAppCheckAppAttestConfig, error) {
	var resource FirebaseAppCheckAppAttestConfig
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/firebaseAppCheckAppAttestConfig:FirebaseAppCheckAppAttestConfig", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseAppCheckAppAttestConfig resources.
type firebaseAppCheckAppAttestConfigState struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                             *string `pulumi:"appId"`
	FirebaseAppCheckAppAttestConfigId *string `pulumi:"firebaseAppCheckAppAttestConfigId"`
	// The relative resource name of the App Attest configuration object
	Name     *string                                  `pulumi:"name"`
	Project  *string                                  `pulumi:"project"`
	Timeouts *FirebaseAppCheckAppAttestConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
	// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
	// fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

type FirebaseAppCheckAppAttestConfigState struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                             pulumi.StringPtrInput
	FirebaseAppCheckAppAttestConfigId pulumi.StringPtrInput
	// The relative resource name of the App Attest configuration object
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts FirebaseAppCheckAppAttestConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
	// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
	// fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckAppAttestConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckAppAttestConfigState)(nil)).Elem()
}

type firebaseAppCheckAppAttestConfigArgs struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                             string                                   `pulumi:"appId"`
	FirebaseAppCheckAppAttestConfigId *string                                  `pulumi:"firebaseAppCheckAppAttestConfigId"`
	Project                           *string                                  `pulumi:"project"`
	Timeouts                          *FirebaseAppCheckAppAttestConfigTimeouts `pulumi:"timeouts"`
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
	// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
	// fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl *string `pulumi:"tokenTtl"`
}

// The set of arguments for constructing a FirebaseAppCheckAppAttestConfig resource.
type FirebaseAppCheckAppAttestConfigArgs struct {
	// The ID of an [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
	AppId                             pulumi.StringInput
	FirebaseAppCheckAppAttestConfigId pulumi.StringPtrInput
	Project                           pulumi.StringPtrInput
	Timeouts                          FirebaseAppCheckAppAttestConfigTimeoutsPtrInput
	// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
	// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
	// fractional digits, ending with 's'. Example: "3.5s".
	TokenTtl pulumi.StringPtrInput
}

func (FirebaseAppCheckAppAttestConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckAppAttestConfigArgs)(nil)).Elem()
}

type FirebaseAppCheckAppAttestConfigInput interface {
	pulumi.Input

	ToFirebaseAppCheckAppAttestConfigOutput() FirebaseAppCheckAppAttestConfigOutput
	ToFirebaseAppCheckAppAttestConfigOutputWithContext(ctx context.Context) FirebaseAppCheckAppAttestConfigOutput
}

func (*FirebaseAppCheckAppAttestConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckAppAttestConfig)(nil)).Elem()
}

func (i *FirebaseAppCheckAppAttestConfig) ToFirebaseAppCheckAppAttestConfigOutput() FirebaseAppCheckAppAttestConfigOutput {
	return i.ToFirebaseAppCheckAppAttestConfigOutputWithContext(context.Background())
}

func (i *FirebaseAppCheckAppAttestConfig) ToFirebaseAppCheckAppAttestConfigOutputWithContext(ctx context.Context) FirebaseAppCheckAppAttestConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseAppCheckAppAttestConfigOutput)
}

type FirebaseAppCheckAppAttestConfigOutput struct{ *pulumi.OutputState }

func (FirebaseAppCheckAppAttestConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckAppAttestConfig)(nil)).Elem()
}

func (o FirebaseAppCheckAppAttestConfigOutput) ToFirebaseAppCheckAppAttestConfigOutput() FirebaseAppCheckAppAttestConfigOutput {
	return o
}

func (o FirebaseAppCheckAppAttestConfigOutput) ToFirebaseAppCheckAppAttestConfigOutputWithContext(ctx context.Context) FirebaseAppCheckAppAttestConfigOutput {
	return o
}

// The ID of an [Apple
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id).
func (o FirebaseAppCheckAppAttestConfigOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckAppAttestConfigOutput) FirebaseAppCheckAppAttestConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) pulumi.StringOutput {
		return v.FirebaseAppCheckAppAttestConfigId
	}).(pulumi.StringOutput)
}

// The relative resource name of the App Attest configuration object
func (o FirebaseAppCheckAppAttestConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckAppAttestConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckAppAttestConfigOutput) Timeouts() FirebaseAppCheckAppAttestConfigTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) FirebaseAppCheckAppAttestConfigTimeoutsPtrOutput {
		return v.Timeouts
	}).(FirebaseAppCheckAppAttestConfigTimeoutsPtrOutput)
}

// Specifies the duration for which App Check tokens exchanged from App Attest artifacts will be valid. If unset, a default
// value of 1 hour is assumed. Must be between 30 minutes and 7 days, inclusive. A duration in seconds with up to nine
// fractional digits, ending with 's'. Example: "3.5s".
func (o FirebaseAppCheckAppAttestConfigOutput) TokenTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckAppAttestConfig) pulumi.StringOutput { return v.TokenTtl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseAppCheckAppAttestConfigInput)(nil)).Elem(), &FirebaseAppCheckAppAttestConfig{})
	pulumi.RegisterOutputType(FirebaseAppCheckAppAttestConfigOutput{})
}
