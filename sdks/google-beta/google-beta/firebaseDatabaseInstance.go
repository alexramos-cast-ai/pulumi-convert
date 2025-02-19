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

type FirebaseDatabaseInstance struct {
	pulumi.CustomResourceState

	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances or
	// https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl pulumi.StringOutput `pulumi:"databaseUrl"`
	// The intended database state. Possible values: ACTIVE, DISABLED.
	DesiredState               pulumi.StringPtrOutput `pulumi:"desiredState"`
	FirebaseDatabaseInstanceId pulumi.StringOutput    `pulumi:"firebaseDatabaseInstanceId"`
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The fully-qualified resource name of the Firebase Realtime Database, in the format:
	// projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID PROJECT_NUMBER: The Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides. Check all [available
	// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region pulumi.StringOutput `pulumi:"region"`
	// The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
	State    pulumi.StringOutput                       `pulumi:"state"`
	Timeouts FirebaseDatabaseInstanceTimeoutsPtrOutput `pulumi:"timeouts"`
	// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
	// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
	// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewFirebaseDatabaseInstance registers a new resource with the given unique name, arguments, and options.
func NewFirebaseDatabaseInstance(ctx *pulumi.Context,
	name string, args *FirebaseDatabaseInstanceArgs, opts ...pulumi.ResourceOption) (*FirebaseDatabaseInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource FirebaseDatabaseInstance
	err = ctx.RegisterPackageResource("google-beta:index/firebaseDatabaseInstance:FirebaseDatabaseInstance", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirebaseDatabaseInstance gets an existing FirebaseDatabaseInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirebaseDatabaseInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirebaseDatabaseInstanceState, opts ...pulumi.ResourceOption) (*FirebaseDatabaseInstance, error) {
	var resource FirebaseDatabaseInstance
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/firebaseDatabaseInstance:FirebaseDatabaseInstance", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirebaseDatabaseInstance resources.
type firebaseDatabaseInstanceState struct {
	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances or
	// https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl *string `pulumi:"databaseUrl"`
	// The intended database state. Possible values: ACTIVE, DISABLED.
	DesiredState               *string `pulumi:"desiredState"`
	FirebaseDatabaseInstanceId *string `pulumi:"firebaseDatabaseInstanceId"`
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	InstanceId *string `pulumi:"instanceId"`
	// The fully-qualified resource name of the Firebase Realtime Database, in the format:
	// projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID PROJECT_NUMBER: The Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides. Check all [available
	// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region *string `pulumi:"region"`
	// The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
	State    *string                           `pulumi:"state"`
	Timeouts *FirebaseDatabaseInstanceTimeouts `pulumi:"timeouts"`
	// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
	// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
	// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	Type *string `pulumi:"type"`
}

type FirebaseDatabaseInstanceState struct {
	// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances or
	// https://{instance-id}.{region}.firebasedatabase.app in other regions.
	DatabaseUrl pulumi.StringPtrInput
	// The intended database state. Possible values: ACTIVE, DISABLED.
	DesiredState               pulumi.StringPtrInput
	FirebaseDatabaseInstanceId pulumi.StringPtrInput
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	InstanceId pulumi.StringPtrInput
	// The fully-qualified resource name of the Firebase Realtime Database, in the format:
	// projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID PROJECT_NUMBER: The Firebase project's
	// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// A reference to the region where the Firebase Realtime database resides. Check all [available
	// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region pulumi.StringPtrInput
	// The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
	State    pulumi.StringPtrInput
	Timeouts FirebaseDatabaseInstanceTimeoutsPtrInput
	// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
	// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
	// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	Type pulumi.StringPtrInput
}

func (FirebaseDatabaseInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseDatabaseInstanceState)(nil)).Elem()
}

type firebaseDatabaseInstanceArgs struct {
	// The intended database state. Possible values: ACTIVE, DISABLED.
	DesiredState               *string `pulumi:"desiredState"`
	FirebaseDatabaseInstanceId *string `pulumi:"firebaseDatabaseInstanceId"`
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	InstanceId string  `pulumi:"instanceId"`
	Project    *string `pulumi:"project"`
	// A reference to the region where the Firebase Realtime database resides. Check all [available
	// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region   string                            `pulumi:"region"`
	Timeouts *FirebaseDatabaseInstanceTimeouts `pulumi:"timeouts"`
	// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
	// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
	// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a FirebaseDatabaseInstance resource.
type FirebaseDatabaseInstanceArgs struct {
	// The intended database state. Possible values: ACTIVE, DISABLED.
	DesiredState               pulumi.StringPtrInput
	FirebaseDatabaseInstanceId pulumi.StringPtrInput
	// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
	InstanceId pulumi.StringInput
	Project    pulumi.StringPtrInput
	// A reference to the region where the Firebase Realtime database resides. Check all [available
	// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
	Region   pulumi.StringInput
	Timeouts FirebaseDatabaseInstanceTimeoutsPtrInput
	// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
	// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
	// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
	Type pulumi.StringPtrInput
}

func (FirebaseDatabaseInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firebaseDatabaseInstanceArgs)(nil)).Elem()
}

type FirebaseDatabaseInstanceInput interface {
	pulumi.Input

	ToFirebaseDatabaseInstanceOutput() FirebaseDatabaseInstanceOutput
	ToFirebaseDatabaseInstanceOutputWithContext(ctx context.Context) FirebaseDatabaseInstanceOutput
}

func (*FirebaseDatabaseInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseDatabaseInstance)(nil)).Elem()
}

func (i *FirebaseDatabaseInstance) ToFirebaseDatabaseInstanceOutput() FirebaseDatabaseInstanceOutput {
	return i.ToFirebaseDatabaseInstanceOutputWithContext(context.Background())
}

func (i *FirebaseDatabaseInstance) ToFirebaseDatabaseInstanceOutputWithContext(ctx context.Context) FirebaseDatabaseInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirebaseDatabaseInstanceOutput)
}

type FirebaseDatabaseInstanceOutput struct{ *pulumi.OutputState }

func (FirebaseDatabaseInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirebaseDatabaseInstance)(nil)).Elem()
}

func (o FirebaseDatabaseInstanceOutput) ToFirebaseDatabaseInstanceOutput() FirebaseDatabaseInstanceOutput {
	return o
}

func (o FirebaseDatabaseInstanceOutput) ToFirebaseDatabaseInstanceOutputWithContext(ctx context.Context) FirebaseDatabaseInstanceOutput {
	return o
}

// The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances or
// https://{instance-id}.{region}.firebasedatabase.app in other regions.
func (o FirebaseDatabaseInstanceOutput) DatabaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.DatabaseUrl }).(pulumi.StringOutput)
}

// The intended database state. Possible values: ACTIVE, DISABLED.
func (o FirebaseDatabaseInstanceOutput) DesiredState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringPtrOutput { return v.DesiredState }).(pulumi.StringPtrOutput)
}

func (o FirebaseDatabaseInstanceOutput) FirebaseDatabaseInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.FirebaseDatabaseInstanceId }).(pulumi.StringOutput)
}

// The globally unique identifier of the Firebase Realtime Database instance. Instance IDs cannot be reused after deletion.
func (o FirebaseDatabaseInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The fully-qualified resource name of the Firebase Realtime Database, in the format:
// projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID PROJECT_NUMBER: The Firebase project's
// ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
// Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510).
func (o FirebaseDatabaseInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FirebaseDatabaseInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A reference to the region where the Firebase Realtime database resides. Check all [available
// regions](https://firebase.google.com/docs/projects/locations#rtdb-locations)
func (o FirebaseDatabaseInstanceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database
func (o FirebaseDatabaseInstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o FirebaseDatabaseInstanceOutput) Timeouts() FirebaseDatabaseInstanceTimeoutsPtrOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) FirebaseDatabaseInstanceTimeoutsPtrOutput { return v.Timeouts }).(FirebaseDatabaseInstanceTimeoutsPtrOutput)
}

// The database type. Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
// Creating user Databases is only available for projects on the Blaze plan. Projects can be upgraded using the Cloud
// Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value:
// "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]
func (o FirebaseDatabaseInstanceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirebaseDatabaseInstance) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirebaseDatabaseInstanceInput)(nil)).Elem(), &FirebaseDatabaseInstance{})
	pulumi.RegisterOutputType(FirebaseDatabaseInstanceOutput{})
}
