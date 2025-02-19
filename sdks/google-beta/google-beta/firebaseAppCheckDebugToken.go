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

type FirebaseAppCheckDebugToken struct {
	pulumi.CustomResourceState

	// The ID of a [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	// [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	// or [Android
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
	AppId pulumi.StringOutput `pulumi:"appId"`
	// The last segment of the resource name of the debug token.
	DebugTokenId pulumi.StringOutput `pulumi:"debugTokenId"`
	// A human readable display name used to identify this debug token.
	DisplayName                  pulumi.StringOutput                         `pulumi:"displayName"`
	FirebaseAppCheckDebugTokenId pulumi.StringOutput                         `pulumi:"firebaseAppCheckDebugTokenId"`
	Project                      pulumi.StringOutput                         `pulumi:"project"`
	Timeouts                     FirebaseAppCheckDebugTokenTimeoutsPtrOutput `pulumi:"timeouts"`
	// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
	// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
	// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
	// populated in any response.
	Token pulumi.StringOutput `pulumi:"token"`
}

// NewFirebaseAppCheckDebugToken registers a new resource with the given unique name, arguments, and options.
func NewFirebaseAppCheckDebugToken(ctx *pulumi.Context,
	name string, args *FirebaseAppCheckDebugTokenArgs, opts ...pulumi.ResourceOption) (*FirebaseAppCheckDebugToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AppId == nil {
		return nil, errors.New("invalid value for required argument 'AppId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseAppCheckDebugToken
	err = ctx.RegisterPackageResource("google-beta:index/firebaseAppCheckDebugToken:FirebaseAppCheckDebugToken", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseAppCheckDebugToken gets an existing FirebaseAppCheckDebugToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseAppCheckDebugToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseAppCheckDebugTokenState, opts ...pulumi.ResourceOption) (*FirebaseAppCheckDebugToken, error) {
	var resource FirebaseAppCheckDebugToken
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseAppCheckDebugToken:FirebaseAppCheckDebugToken", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseAppCheckDebugToken resources.
type firebaseAppCheckDebugTokenState struct {
	// The ID of a [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	// [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	// or [Android
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
	AppId *string `pulumi:"appId"`
	// The last segment of the resource name of the debug token.
	DebugTokenId *string `pulumi:"debugTokenId"`
	// A human readable display name used to identify this debug token.
	DisplayName                  *string                             `pulumi:"displayName"`
	FirebaseAppCheckDebugTokenId *string                             `pulumi:"firebaseAppCheckDebugTokenId"`
	Project                      *string                             `pulumi:"project"`
	Timeouts                     *FirebaseAppCheckDebugTokenTimeouts `pulumi:"timeouts"`
	// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
	// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
	// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
	// populated in any response.
	Token *string `pulumi:"token"`
}

type FirebaseAppCheckDebugTokenState struct {
	// The ID of a [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	// [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	// or [Android
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
	AppId pulumi.StringPtrInput
	// The last segment of the resource name of the debug token.
	DebugTokenId pulumi.StringPtrInput
	// A human readable display name used to identify this debug token.
	DisplayName                  pulumi.StringPtrInput
	FirebaseAppCheckDebugTokenId pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Timeouts                     FirebaseAppCheckDebugTokenTimeoutsPtrInput
	// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
	// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
	// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
	// populated in any response.
	Token pulumi.StringPtrInput
}

func (FirebaseAppCheckDebugTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckDebugTokenState)(nil)).Elem()
}

type firebaseAppCheckDebugTokenArgs struct {
	// The ID of a [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	// [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	// or [Android
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
	AppId string `pulumi:"appId"`
	// A human readable display name used to identify this debug token.
	DisplayName                  string                              `pulumi:"displayName"`
	FirebaseAppCheckDebugTokenId *string                             `pulumi:"firebaseAppCheckDebugTokenId"`
	Project                      *string                             `pulumi:"project"`
	Timeouts                     *FirebaseAppCheckDebugTokenTimeouts `pulumi:"timeouts"`
	// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
	// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
	// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
	// populated in any response.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a FirebaseAppCheckDebugToken resource.
type FirebaseAppCheckDebugTokenArgs struct {
	// The ID of a [Web
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
	// [Apple
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
	// or [Android
	// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
	AppId pulumi.StringInput
	// A human readable display name used to identify this debug token.
	DisplayName                  pulumi.StringInput
	FirebaseAppCheckDebugTokenId pulumi.StringPtrInput
	Project                      pulumi.StringPtrInput
	Timeouts                     FirebaseAppCheckDebugTokenTimeoutsPtrInput
	// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
	// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
	// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
	// populated in any response.
	Token pulumi.StringInput
}

func (FirebaseAppCheckDebugTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseAppCheckDebugTokenArgs)(nil)).Elem()
}

type FirebaseAppCheckDebugTokenInput interface {
	pulumi.Input

	ToFirebaseAppCheckDebugTokenOutput() FirebaseAppCheckDebugTokenOutput
	ToFirebaseAppCheckDebugTokenOutputWithContext(ctx context.Context) FirebaseAppCheckDebugTokenOutput
}

func (*FirebaseAppCheckDebugToken) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckDebugToken)(nil)).Elem()
}

func (i *FirebaseAppCheckDebugToken) ToFirebaseAppCheckDebugTokenOutput() FirebaseAppCheckDebugTokenOutput {
	return i.ToFirebaseAppCheckDebugTokenOutputWithContext(context.Background())
}

func (i *FirebaseAppCheckDebugToken) ToFirebaseAppCheckDebugTokenOutputWithContext(ctx context.Context) FirebaseAppCheckDebugTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseAppCheckDebugTokenOutput)
}

type FirebaseAppCheckDebugTokenOutput struct{ *pulumi.OutputState }

func (FirebaseAppCheckDebugTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseAppCheckDebugToken)(nil)).Elem()
}

func (o FirebaseAppCheckDebugTokenOutput) ToFirebaseAppCheckDebugTokenOutput() FirebaseAppCheckDebugTokenOutput {
	return o
}

func (o FirebaseAppCheckDebugTokenOutput) ToFirebaseAppCheckDebugTokenOutputWithContext(ctx context.Context) FirebaseAppCheckDebugTokenOutput {
	return o
}

// The ID of a [Web
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.webApps#WebApp.FIELDS.app_id),
// [Apple
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.iosApps#IosApp.FIELDS.app_id),
// or [Android
// App](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects.androidApps#AndroidApp.FIELDS.app_id)
func (o FirebaseAppCheckDebugTokenOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.AppId }).(pulumi.StringOutput)
}

// The last segment of the resource name of the debug token.
func (o FirebaseAppCheckDebugTokenOutput) DebugTokenId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.DebugTokenId }).(pulumi.StringOutput)
}

// A human readable display name used to identify this debug token.
func (o FirebaseAppCheckDebugTokenOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckDebugTokenOutput) FirebaseAppCheckDebugTokenId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.FirebaseAppCheckDebugTokenId }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckDebugTokenOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o FirebaseAppCheckDebugTokenOutput) Timeouts() FirebaseAppCheckDebugTokenTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) FirebaseAppCheckDebugTokenTimeoutsPtrOutput { return v.Timeouts }).(FirebaseAppCheckDebugTokenTimeoutsPtrOutput)
}

// The secret token itself. Must be provided during creation, and must be a UUID4, case insensitive. You may use a method
// of your choice such as random/random_uuid to generate the token. This field is immutable once set, and cannot be
// updated. You can, however, delete this debug token to revoke it. For security reasons, this field will never be
// populated in any response.
func (o FirebaseAppCheckDebugTokenOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseAppCheckDebugToken) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseAppCheckDebugTokenInput)(nil)).Elem(), &FirebaseAppCheckDebugToken{})
	pulumi.RegisterOutputType(FirebaseAppCheckDebugTokenOutput{})
}
