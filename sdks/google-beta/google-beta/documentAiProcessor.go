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

type DocumentAiProcessor struct {
	pulumi.CustomResourceState

	// The display name. Must be unique.
	DisplayName           pulumi.StringOutput `pulumi:"displayName"`
	DocumentAiProcessorId pulumi.StringOutput `pulumi:"documentAiProcessorId"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the processor.
	Name     pulumi.StringOutput                  `pulumi:"name"`
	Project  pulumi.StringOutput                  `pulumi:"project"`
	Timeouts DocumentAiProcessorTimeoutsPtrOutput `pulumi:"timeouts"`
	// The type of processor. For possible types see the [official
	// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDocumentAiProcessor registers a new resource with the given unique name, arguments, and options.
func NewDocumentAiProcessor(ctx *pulumi.Context,
	name string, args *DocumentAiProcessorArgs, opts ...pulumi.ResourceOption) (*DocumentAiProcessor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DocumentAiProcessor
	err = ctx.RegisterPackageResource("google-beta:index/documentAiProcessor:DocumentAiProcessor", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDocumentAiProcessor gets an existing DocumentAiProcessor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDocumentAiProcessor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DocumentAiProcessorState, opts ...pulumi.ResourceOption) (*DocumentAiProcessor, error) {
	var resource DocumentAiProcessor
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/documentAiProcessor:DocumentAiProcessor", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DocumentAiProcessor resources.
type documentAiProcessorState struct {
	// The display name. Must be unique.
	DisplayName           *string `pulumi:"displayName"`
	DocumentAiProcessorId *string `pulumi:"documentAiProcessorId"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The resource name of the processor.
	Name     *string                      `pulumi:"name"`
	Project  *string                      `pulumi:"project"`
	Timeouts *DocumentAiProcessorTimeouts `pulumi:"timeouts"`
	// The type of processor. For possible types see the [official
	// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type *string `pulumi:"type"`
}

type DocumentAiProcessorState struct {
	// The display name. Must be unique.
	DisplayName           pulumi.StringPtrInput
	DocumentAiProcessorId pulumi.StringPtrInput
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The resource name of the processor.
	Name     pulumi.StringPtrInput
	Project  pulumi.StringPtrInput
	Timeouts DocumentAiProcessorTimeoutsPtrInput
	// The type of processor. For possible types see the [official
	// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringPtrInput
}

func (DocumentAiProcessorState) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorState)(nil)).Elem()
}

type documentAiProcessorArgs struct {
	// The display name. Must be unique.
	DisplayName           string  `pulumi:"displayName"`
	DocumentAiProcessorId *string `pulumi:"documentAiProcessorId"`
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// The location of the resource.
	Location string                       `pulumi:"location"`
	Project  *string                      `pulumi:"project"`
	Timeouts *DocumentAiProcessorTimeouts `pulumi:"timeouts"`
	// The type of processor. For possible types see the [official
	// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a DocumentAiProcessor resource.
type DocumentAiProcessorArgs struct {
	// The display name. Must be unique.
	DisplayName           pulumi.StringInput
	DocumentAiProcessorId pulumi.StringPtrInput
	// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
	KmsKeyName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringInput
	Project  pulumi.StringPtrInput
	Timeouts DocumentAiProcessorTimeoutsPtrInput
	// The type of processor. For possible types see the [official
	// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
	Type pulumi.StringInput
}

func (DocumentAiProcessorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*documentAiProcessorArgs)(nil)).Elem()
}

type DocumentAiProcessorInput interface {
	pulumi.Input

	ToDocumentAiProcessorOutput() DocumentAiProcessorOutput
	ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput
}

func (*DocumentAiProcessor) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessor)(nil)).Elem()
}

func (i *DocumentAiProcessor) ToDocumentAiProcessorOutput() DocumentAiProcessorOutput {
	return i.ToDocumentAiProcessorOutputWithContext(context.Background())
}

func (i *DocumentAiProcessor) ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DocumentAiProcessorOutput)
}

type DocumentAiProcessorOutput struct{ *pulumi.OutputState }

func (DocumentAiProcessorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DocumentAiProcessor)(nil)).Elem()
}

func (o DocumentAiProcessorOutput) ToDocumentAiProcessorOutput() DocumentAiProcessorOutput {
	return o
}

func (o DocumentAiProcessorOutput) ToDocumentAiProcessorOutputWithContext(ctx context.Context) DocumentAiProcessorOutput {
	return o
}

// The display name. Must be unique.
func (o DocumentAiProcessorOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o DocumentAiProcessorOutput) DocumentAiProcessorId() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.DocumentAiProcessorId }).(pulumi.StringOutput)
}

// The KMS key used for encryption/decryption in CMEK scenarios. See https://cloud.google.com/security-key-management.
func (o DocumentAiProcessorOutput) KmsKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringPtrOutput { return v.KmsKeyName }).(pulumi.StringPtrOutput)
}

// The location of the resource.
func (o DocumentAiProcessorOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the processor.
func (o DocumentAiProcessorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DocumentAiProcessorOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

func (o DocumentAiProcessorOutput) Timeouts() DocumentAiProcessorTimeoutsPtrOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) DocumentAiProcessorTimeoutsPtrOutput { return v.Timeouts }).(DocumentAiProcessorTimeoutsPtrOutput)
}

// The type of processor. For possible types see the [official
// list](https://cloud.google.com/document-ai/docs/reference/rest/v1/projects.locations/fetchProcessorTypes#google.cloud.documentai.v1.DocumentProcessorService.FetchProcessorTypes)
func (o DocumentAiProcessorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *DocumentAiProcessor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DocumentAiProcessorInput)(nil)).Elem(), &DocumentAiProcessor{})
	pulumi.RegisterOutputType(DocumentAiProcessorOutput{})
}
