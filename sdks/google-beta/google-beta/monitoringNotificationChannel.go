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

type MonitoringNotificationChannel struct {
	pulumi.CustomResourceState

	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
	// is limited to 512 Unicode characters.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete                     pulumi.BoolPtrOutput   `pulumi:"forceDelete"`
	Labels                          pulumi.StringMapOutput `pulumi:"labels"`
	MonitoringNotificationChannelId pulumi.StringOutput    `pulumi:"monitoringNotificationChannelId"`
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
	// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
	// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
	// specified in both locations and will cause an error. Changing from one location to a different credential configuration
	// in the config will require an apply to update state.
	SensitiveLabels MonitoringNotificationChannelSensitiveLabelsPtrOutput `pulumi:"sensitiveLabels"`
	Timeouts        MonitoringNotificationChannelTimeoutsPtrOutput        `pulumi:"timeouts"`
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type pulumi.StringOutput `pulumi:"type"`
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
	// begin with a letter.
	UserLabels pulumi.StringMapOutput `pulumi:"userLabels"`
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus pulumi.StringOutput `pulumi:"verificationStatus"`
}

// NewMonitoringNotificationChannel registers a new resource with the given unique name, arguments, and options.
func NewMonitoringNotificationChannel(ctx *pulumi.Context,
	name string, args *MonitoringNotificationChannelArgs, opts ...pulumi.ResourceOption) (*MonitoringNotificationChannel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource MonitoringNotificationChannel
	err = ctx.RegisterPackageResource("google-beta:index/monitoringNotificationChannel:MonitoringNotificationChannel", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMonitoringNotificationChannel gets an existing MonitoringNotificationChannel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMonitoringNotificationChannel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MonitoringNotificationChannelState, opts ...pulumi.ResourceOption) (*MonitoringNotificationChannel, error) {
	var resource MonitoringNotificationChannel
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/monitoringNotificationChannel:MonitoringNotificationChannel", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MonitoringNotificationChannel resources.
type monitoringNotificationChannelState struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `pulumi:"description"`
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
	// is limited to 512 Unicode characters.
	DisplayName *string `pulumi:"displayName"`
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled *bool `pulumi:"enabled"`
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete                     *bool             `pulumi:"forceDelete"`
	Labels                          map[string]string `pulumi:"labels"`
	MonitoringNotificationChannelId *string           `pulumi:"monitoringNotificationChannelId"`
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
	// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
	// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
	// specified in both locations and will cause an error. Changing from one location to a different credential configuration
	// in the config will require an apply to update state.
	SensitiveLabels *MonitoringNotificationChannelSensitiveLabels `pulumi:"sensitiveLabels"`
	Timeouts        *MonitoringNotificationChannelTimeouts        `pulumi:"timeouts"`
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type *string `pulumi:"type"`
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
	// begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus *string `pulumi:"verificationStatus"`
}

type MonitoringNotificationChannelState struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrInput
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
	// is limited to 512 Unicode characters.
	DisplayName pulumi.StringPtrInput
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrInput
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete                     pulumi.BoolPtrInput
	Labels                          pulumi.StringMapInput
	MonitoringNotificationChannelId pulumi.StringPtrInput
	// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
	// [CHANNEL_ID] is automatically assigned by the server on creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
	// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
	// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
	// specified in both locations and will cause an error. Changing from one location to a different credential configuration
	// in the config will require an apply to update state.
	SensitiveLabels MonitoringNotificationChannelSensitiveLabelsPtrInput
	Timeouts        MonitoringNotificationChannelTimeoutsPtrInput
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type pulumi.StringPtrInput
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
	// begin with a letter.
	UserLabels pulumi.StringMapInput
	// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
	// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
	// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
	// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
	// verification or that this specific channel has been exempted from verification because it was created prior to
	// verification being required for channels of this type.This field cannot be modified using a standard
	// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
	VerificationStatus pulumi.StringPtrInput
}

func (MonitoringNotificationChannelState) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringNotificationChannelState)(nil)).Elem()
}

type monitoringNotificationChannelArgs struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description *string `pulumi:"description"`
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
	// is limited to 512 Unicode characters.
	DisplayName *string `pulumi:"displayName"`
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled *bool `pulumi:"enabled"`
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete                     *bool             `pulumi:"forceDelete"`
	Labels                          map[string]string `pulumi:"labels"`
	MonitoringNotificationChannelId *string           `pulumi:"monitoringNotificationChannelId"`
	Project                         *string           `pulumi:"project"`
	// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
	// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
	// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
	// specified in both locations and will cause an error. Changing from one location to a different credential configuration
	// in the config will require an apply to update state.
	SensitiveLabels *MonitoringNotificationChannelSensitiveLabels `pulumi:"sensitiveLabels"`
	Timeouts        *MonitoringNotificationChannelTimeouts        `pulumi:"timeouts"`
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type string `pulumi:"type"`
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
	// begin with a letter.
	UserLabels map[string]string `pulumi:"userLabels"`
}

// The set of arguments for constructing a MonitoringNotificationChannel resource.
type MonitoringNotificationChannelArgs struct {
	// An optional human-readable description of this notification channel. This description may provide additional details,
	// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
	Description pulumi.StringPtrInput
	// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
	// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
	// is limited to 512 Unicode characters.
	DisplayName pulumi.StringPtrInput
	// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
	// notifications to a particular channel without removing the channel from all alerting policies that reference the
	// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
	// same set of alerting policies on the channel at some point in the future.
	Enabled pulumi.BoolPtrInput
	// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
	// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
	// deleted in a delete operation.
	ForceDelete                     pulumi.BoolPtrInput
	Labels                          pulumi.StringMapInput
	MonitoringNotificationChannelId pulumi.StringPtrInput
	Project                         pulumi.StringPtrInput
	// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
	// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
	// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
	// specified in both locations and will cause an error. Changing from one location to a different credential configuration
	// in the config will require an apply to update state.
	SensitiveLabels MonitoringNotificationChannelSensitiveLabelsPtrInput
	Timeouts        MonitoringNotificationChannelTimeoutsPtrInput
	// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
	// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
	// valid values such as "email", "slack", etc...
	Type pulumi.StringInput
	// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
	// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
	// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
	// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
	// begin with a letter.
	UserLabels pulumi.StringMapInput
}

func (MonitoringNotificationChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*monitoringNotificationChannelArgs)(nil)).Elem()
}

type MonitoringNotificationChannelInput interface {
	pulumi.Input

	ToMonitoringNotificationChannelOutput() MonitoringNotificationChannelOutput
	ToMonitoringNotificationChannelOutputWithContext(ctx context.Context) MonitoringNotificationChannelOutput
}

func (*MonitoringNotificationChannel) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringNotificationChannel)(nil)).Elem()
}

func (i *MonitoringNotificationChannel) ToMonitoringNotificationChannelOutput() MonitoringNotificationChannelOutput {
	return i.ToMonitoringNotificationChannelOutputWithContext(context.Background())
}

func (i *MonitoringNotificationChannel) ToMonitoringNotificationChannelOutputWithContext(ctx context.Context) MonitoringNotificationChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonitoringNotificationChannelOutput)
}

type MonitoringNotificationChannelOutput struct{ *pulumi.OutputState }

func (MonitoringNotificationChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonitoringNotificationChannel)(nil)).Elem()
}

func (o MonitoringNotificationChannelOutput) ToMonitoringNotificationChannelOutput() MonitoringNotificationChannelOutput {
	return o
}

func (o MonitoringNotificationChannelOutput) ToMonitoringNotificationChannelOutputWithContext(ctx context.Context) MonitoringNotificationChannelOutput {
	return o
}

// An optional human-readable description of this notification channel. This description may provide additional details,
// beyond the display name, for the channel. This may not exceed 1024 Unicode characters.
func (o MonitoringNotificationChannelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// An optional human-readable name for this notification channel. It is recommended that you specify a non-empty and unique
// name in order to make it easier to identify the channels in your project, though this is not enforced. The display name
// is limited to 512 Unicode characters.
func (o MonitoringNotificationChannelOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Whether notifications are forwarded to the described channel. This makes it possible to disable delivery of
// notifications to a particular channel without removing the channel from all alerting policies that reference the
// channel. This is a more convenient approach when the change is temporary and you want to receive notifications from the
// same set of alerting policies on the channel at some point in the future.
func (o MonitoringNotificationChannelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// If true, the notification channel will be deleted regardless of its use in alert policies (the policies will be updated
// to remove the channel). If false, channels that are still referenced by an existing alerting policy will fail to be
// deleted in a delete operation.
func (o MonitoringNotificationChannelOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

func (o MonitoringNotificationChannelOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o MonitoringNotificationChannelOutput) MonitoringNotificationChannelId() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringOutput { return v.MonitoringNotificationChannelId }).(pulumi.StringOutput)
}

// The full REST resource name for this channel. The syntax is: projects/[PROJECT_ID]/notificationChannels/[CHANNEL_ID] The
// [CHANNEL_ID] is automatically assigned by the server on creation.
func (o MonitoringNotificationChannelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MonitoringNotificationChannelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Different notification type behaviors are configured primarily using the the 'labels' field on this resource. This block
// contains the labels which contain secrets or passwords so that they can be marked sensitive and hidden from plan output.
// The name of the field, eg: password, will be the key in the 'labels' map in the api request. Credentials may not be
// specified in both locations and will cause an error. Changing from one location to a different credential configuration
// in the config will require an apply to update state.
func (o MonitoringNotificationChannelOutput) SensitiveLabels() MonitoringNotificationChannelSensitiveLabelsPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) MonitoringNotificationChannelSensitiveLabelsPtrOutput {
		return v.SensitiveLabels
	}).(MonitoringNotificationChannelSensitiveLabelsPtrOutput)
}

func (o MonitoringNotificationChannelOutput) Timeouts() MonitoringNotificationChannelTimeoutsPtrOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) MonitoringNotificationChannelTimeoutsPtrOutput {
		return v.Timeouts
	}).(MonitoringNotificationChannelTimeoutsPtrOutput)
}

// The type of the notification channel. This field matches the value of the NotificationChannelDescriptor.type field. See
// https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.notificationChannelDescriptors/list to get the list of
// valid values such as "email", "slack", etc...
func (o MonitoringNotificationChannelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// User-supplied key/value data that does not need to conform to the corresponding NotificationChannelDescriptor's schema,
// unlike the labels field. This field is intended to be used for organizing and identifying the NotificationChannel
// objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes,
// whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must
// begin with a letter.
func (o MonitoringNotificationChannelOutput) UserLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringMapOutput { return v.UserLabels }).(pulumi.StringMapOutput)
}

// Indicates whether this channel has been verified or not. On a ListNotificationChannels or GetNotificationChannel
// operation, this field is expected to be populated.If the value is UNVERIFIED, then it indicates that the channel is
// non-functioning (it both requires verification and lacks verification); otherwise, it is assumed that the channel
// works.If the channel is neither VERIFIED nor UNVERIFIED, it implies that the channel is of a type that does not require
// verification or that this specific channel has been exempted from verification because it was created prior to
// verification being required for channels of this type.This field cannot be modified using a standard
// UpdateNotificationChannel operation. To change the value of this field, you must call VerifyNotificationChannel.
func (o MonitoringNotificationChannelOutput) VerificationStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *MonitoringNotificationChannel) pulumi.StringOutput { return v.VerificationStatus }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MonitoringNotificationChannelInput)(nil)).Elem(), &MonitoringNotificationChannel{})
	pulumi.RegisterOutputType(MonitoringNotificationChannelOutput{})
}
