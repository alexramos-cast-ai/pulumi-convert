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

type BigqueryAnalyticsHubListingSubscription struct {
	pulumi.CustomResourceState

	BigqueryAnalyticsHubListingSubscriptionId pulumi.StringOutput `pulumi:"bigqueryAnalyticsHubListingSubscriptionId"`
	// Timestamp when the subscription was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
	// characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringOutput `pulumi:"dataExchangeId"`
	// The destination dataset for this subscription.
	DestinationDataset BigqueryAnalyticsHubListingSubscriptionDestinationDatasetOutput `pulumi:"destinationDataset"`
	// Timestamp when the subscription was last modified.
	LastModifyTime pulumi.StringOutput `pulumi:"lastModifyTime"`
	// Output only. Map of listing resource names to associated linked resource, e.g.
	// projects/123/locations/US/dataExchanges/456/listings/789 > projects/123/datasets/my_dataset
	LinkedDatasetMaps BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMapArrayOutput `pulumi:"linkedDatasetMaps"`
	// Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.
	LinkedResources BigqueryAnalyticsHubListingSubscriptionLinkedResourceArrayOutput `pulumi:"linkedResources"`
	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
	// require URL-escaping, or characters outside of ASCII, spaces.
	ListingId pulumi.StringOutput `pulumi:"listingId"`
	// The name of the location for this subscription.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the subscription. e.g. "projects/myproject/locations/US/subscriptions/123"
	Name pulumi.StringOutput `pulumi:"name"`
	// Display name of the project of this subscription.
	OrganizationDisplayName pulumi.StringOutput `pulumi:"organizationDisplayName"`
	// Organization of the project this subscription belongs to.
	OrganizationId pulumi.StringOutput `pulumi:"organizationId"`
	Project        pulumi.StringOutput `pulumi:"project"`
	// Listing shared asset type.
	ResourceType pulumi.StringOutput `pulumi:"resourceType"`
	// Current state of the subscription.
	State pulumi.StringOutput `pulumi:"state"`
	// Email of the subscriber.
	SubscriberContact pulumi.StringOutput `pulumi:"subscriberContact"`
	// The subscription id used to reference the subscription.
	SubscriptionId pulumi.StringOutput                                      `pulumi:"subscriptionId"`
	Timeouts       BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewBigqueryAnalyticsHubListingSubscription registers a new resource with the given unique name, arguments, and options.
func NewBigqueryAnalyticsHubListingSubscription(ctx *pulumi.Context,
	name string, args *BigqueryAnalyticsHubListingSubscriptionArgs, opts ...pulumi.ResourceOption) (*BigqueryAnalyticsHubListingSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DataExchangeId == nil {
		return nil, errors.New("invalid value for required argument 'DataExchangeId'")
	}
	if args.DestinationDataset == nil {
		return nil, errors.New("invalid value for required argument 'DestinationDataset'")
	}
	if args.ListingId == nil {
		return nil, errors.New("invalid value for required argument 'ListingId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource BigqueryAnalyticsHubListingSubscription
	err = ctx.RegisterPackageResource("google:index/bigqueryAnalyticsHubListingSubscription:BigqueryAnalyticsHubListingSubscription", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBigqueryAnalyticsHubListingSubscription gets an existing BigqueryAnalyticsHubListingSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBigqueryAnalyticsHubListingSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BigqueryAnalyticsHubListingSubscriptionState, opts ...pulumi.ResourceOption) (*BigqueryAnalyticsHubListingSubscription, error) {
	var resource BigqueryAnalyticsHubListingSubscription
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/bigqueryAnalyticsHubListingSubscription:BigqueryAnalyticsHubListingSubscription", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BigqueryAnalyticsHubListingSubscription resources.
type bigqueryAnalyticsHubListingSubscriptionState struct {
	BigqueryAnalyticsHubListingSubscriptionId *string `pulumi:"bigqueryAnalyticsHubListingSubscriptionId"`
	// Timestamp when the subscription was created.
	CreationTime *string `pulumi:"creationTime"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
	// characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId *string `pulumi:"dataExchangeId"`
	// The destination dataset for this subscription.
	DestinationDataset *BigqueryAnalyticsHubListingSubscriptionDestinationDataset `pulumi:"destinationDataset"`
	// Timestamp when the subscription was last modified.
	LastModifyTime *string `pulumi:"lastModifyTime"`
	// Output only. Map of listing resource names to associated linked resource, e.g.
	// projects/123/locations/US/dataExchanges/456/listings/789 > projects/123/datasets/my_dataset
	LinkedDatasetMaps []BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMap `pulumi:"linkedDatasetMaps"`
	// Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.
	LinkedResources []BigqueryAnalyticsHubListingSubscriptionLinkedResource `pulumi:"linkedResources"`
	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
	// require URL-escaping, or characters outside of ASCII, spaces.
	ListingId *string `pulumi:"listingId"`
	// The name of the location for this subscription.
	Location *string `pulumi:"location"`
	// The resource name of the subscription. e.g. "projects/myproject/locations/US/subscriptions/123"
	Name *string `pulumi:"name"`
	// Display name of the project of this subscription.
	OrganizationDisplayName *string `pulumi:"organizationDisplayName"`
	// Organization of the project this subscription belongs to.
	OrganizationId *string `pulumi:"organizationId"`
	Project        *string `pulumi:"project"`
	// Listing shared asset type.
	ResourceType *string `pulumi:"resourceType"`
	// Current state of the subscription.
	State *string `pulumi:"state"`
	// Email of the subscriber.
	SubscriberContact *string `pulumi:"subscriberContact"`
	// The subscription id used to reference the subscription.
	SubscriptionId *string                                          `pulumi:"subscriptionId"`
	Timeouts       *BigqueryAnalyticsHubListingSubscriptionTimeouts `pulumi:"timeouts"`
}

type BigqueryAnalyticsHubListingSubscriptionState struct {
	BigqueryAnalyticsHubListingSubscriptionId pulumi.StringPtrInput
	// Timestamp when the subscription was created.
	CreationTime pulumi.StringPtrInput
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
	// characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringPtrInput
	// The destination dataset for this subscription.
	DestinationDataset BigqueryAnalyticsHubListingSubscriptionDestinationDatasetPtrInput
	// Timestamp when the subscription was last modified.
	LastModifyTime pulumi.StringPtrInput
	// Output only. Map of listing resource names to associated linked resource, e.g.
	// projects/123/locations/US/dataExchanges/456/listings/789 > projects/123/datasets/my_dataset
	LinkedDatasetMaps BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMapArrayInput
	// Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.
	LinkedResources BigqueryAnalyticsHubListingSubscriptionLinkedResourceArrayInput
	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
	// require URL-escaping, or characters outside of ASCII, spaces.
	ListingId pulumi.StringPtrInput
	// The name of the location for this subscription.
	Location pulumi.StringPtrInput
	// The resource name of the subscription. e.g. "projects/myproject/locations/US/subscriptions/123"
	Name pulumi.StringPtrInput
	// Display name of the project of this subscription.
	OrganizationDisplayName pulumi.StringPtrInput
	// Organization of the project this subscription belongs to.
	OrganizationId pulumi.StringPtrInput
	Project        pulumi.StringPtrInput
	// Listing shared asset type.
	ResourceType pulumi.StringPtrInput
	// Current state of the subscription.
	State pulumi.StringPtrInput
	// Email of the subscriber.
	SubscriberContact pulumi.StringPtrInput
	// The subscription id used to reference the subscription.
	SubscriptionId pulumi.StringPtrInput
	Timeouts       BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrInput
}

func (BigqueryAnalyticsHubListingSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryAnalyticsHubListingSubscriptionState)(nil)).Elem()
}

type bigqueryAnalyticsHubListingSubscriptionArgs struct {
	BigqueryAnalyticsHubListingSubscriptionId *string `pulumi:"bigqueryAnalyticsHubListingSubscriptionId"`
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
	// characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId string `pulumi:"dataExchangeId"`
	// The destination dataset for this subscription.
	DestinationDataset BigqueryAnalyticsHubListingSubscriptionDestinationDataset `pulumi:"destinationDataset"`
	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
	// require URL-escaping, or characters outside of ASCII, spaces.
	ListingId string `pulumi:"listingId"`
	// The name of the location for this subscription.
	Location string                                           `pulumi:"location"`
	Project  *string                                          `pulumi:"project"`
	Timeouts *BigqueryAnalyticsHubListingSubscriptionTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a BigqueryAnalyticsHubListingSubscription resource.
type BigqueryAnalyticsHubListingSubscriptionArgs struct {
	BigqueryAnalyticsHubListingSubscriptionId pulumi.StringPtrInput
	// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
	// characters that require URL-escaping, or characters outside of ASCII, spaces.
	DataExchangeId pulumi.StringInput
	// The destination dataset for this subscription.
	DestinationDataset BigqueryAnalyticsHubListingSubscriptionDestinationDatasetInput
	// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
	// require URL-escaping, or characters outside of ASCII, spaces.
	ListingId pulumi.StringInput
	// The name of the location for this subscription.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	Timeouts BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrInput
}

func (BigqueryAnalyticsHubListingSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bigqueryAnalyticsHubListingSubscriptionArgs)(nil)).Elem()
}

type BigqueryAnalyticsHubListingSubscriptionInput interface {
	pulumi.Input

	ToBigqueryAnalyticsHubListingSubscriptionOutput() BigqueryAnalyticsHubListingSubscriptionOutput
	ToBigqueryAnalyticsHubListingSubscriptionOutputWithContext(ctx context.Context) BigqueryAnalyticsHubListingSubscriptionOutput
}

func (*BigqueryAnalyticsHubListingSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryAnalyticsHubListingSubscription)(nil)).Elem()
}

func (i *BigqueryAnalyticsHubListingSubscription) ToBigqueryAnalyticsHubListingSubscriptionOutput() BigqueryAnalyticsHubListingSubscriptionOutput {
	return i.ToBigqueryAnalyticsHubListingSubscriptionOutputWithContext(context.Background())
}

func (i *BigqueryAnalyticsHubListingSubscription) ToBigqueryAnalyticsHubListingSubscriptionOutputWithContext(ctx context.Context) BigqueryAnalyticsHubListingSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BigqueryAnalyticsHubListingSubscriptionOutput)
}

type BigqueryAnalyticsHubListingSubscriptionOutput struct{ *pulumi.OutputState }

func (BigqueryAnalyticsHubListingSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BigqueryAnalyticsHubListingSubscription)(nil)).Elem()
}

func (o BigqueryAnalyticsHubListingSubscriptionOutput) ToBigqueryAnalyticsHubListingSubscriptionOutput() BigqueryAnalyticsHubListingSubscriptionOutput {
	return o
}

func (o BigqueryAnalyticsHubListingSubscriptionOutput) ToBigqueryAnalyticsHubListingSubscriptionOutputWithContext(ctx context.Context) BigqueryAnalyticsHubListingSubscriptionOutput {
	return o
}

func (o BigqueryAnalyticsHubListingSubscriptionOutput) BigqueryAnalyticsHubListingSubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput {
		return v.BigqueryAnalyticsHubListingSubscriptionId
	}).(pulumi.StringOutput)
}

// Timestamp when the subscription was created.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use
// characters that require URL-escaping, or characters outside of ASCII, spaces.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) DataExchangeId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.DataExchangeId }).(pulumi.StringOutput)
}

// The destination dataset for this subscription.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) DestinationDataset() BigqueryAnalyticsHubListingSubscriptionDestinationDatasetOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) BigqueryAnalyticsHubListingSubscriptionDestinationDatasetOutput {
		return v.DestinationDataset
	}).(BigqueryAnalyticsHubListingSubscriptionDestinationDatasetOutput)
}

// Timestamp when the subscription was last modified.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) LastModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.LastModifyTime }).(pulumi.StringOutput)
}

// Output only. Map of listing resource names to associated linked resource, e.g.
// projects/123/locations/US/dataExchanges/456/listings/789 > projects/123/datasets/my_dataset
func (o BigqueryAnalyticsHubListingSubscriptionOutput) LinkedDatasetMaps() BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMapArrayOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMapArrayOutput {
		return v.LinkedDatasetMaps
	}).(BigqueryAnalyticsHubListingSubscriptionLinkedDatasetMapArrayOutput)
}

// Output only. Linked resources created in the subscription. Only contains values if state = STATE_ACTIVE.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) LinkedResources() BigqueryAnalyticsHubListingSubscriptionLinkedResourceArrayOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) BigqueryAnalyticsHubListingSubscriptionLinkedResourceArrayOutput {
		return v.LinkedResources
	}).(BigqueryAnalyticsHubListingSubscriptionLinkedResourceArrayOutput)
}

// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that
// require URL-escaping, or characters outside of ASCII, spaces.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) ListingId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.ListingId }).(pulumi.StringOutput)
}

// The name of the location for this subscription.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the subscription. e.g. "projects/myproject/locations/US/subscriptions/123"
func (o BigqueryAnalyticsHubListingSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Display name of the project of this subscription.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) OrganizationDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.OrganizationDisplayName }).(pulumi.StringOutput)
}

// Organization of the project this subscription belongs to.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) OrganizationId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.OrganizationId }).(pulumi.StringOutput)
}

func (o BigqueryAnalyticsHubListingSubscriptionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Listing shared asset type.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.ResourceType }).(pulumi.StringOutput)
}

// Current state of the subscription.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Email of the subscriber.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) SubscriberContact() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.SubscriberContact }).(pulumi.StringOutput)
}

// The subscription id used to reference the subscription.
func (o BigqueryAnalyticsHubListingSubscriptionOutput) SubscriptionId() pulumi.StringOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) pulumi.StringOutput { return v.SubscriptionId }).(pulumi.StringOutput)
}

func (o BigqueryAnalyticsHubListingSubscriptionOutput) Timeouts() BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrOutput {
	return o.ApplyT(func(v *BigqueryAnalyticsHubListingSubscription) BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrOutput {
		return v.Timeouts
	}).(BigqueryAnalyticsHubListingSubscriptionTimeoutsPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BigqueryAnalyticsHubListingSubscriptionInput)(nil)).Elem(), &BigqueryAnalyticsHubListingSubscription{})
	pulumi.RegisterOutputType(BigqueryAnalyticsHubListingSubscriptionOutput{})
}
