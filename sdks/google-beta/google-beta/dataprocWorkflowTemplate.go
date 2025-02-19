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

type DataprocWorkflowTemplate struct {
	pulumi.CustomResourceState

	// Output only. The time template was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
	// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
	// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
	// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
	// on a [managed
	// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	// the cluster is deleted.
	DagTimeout                 pulumi.StringPtrOutput `pulumi:"dagTimeout"`
	DataprocWorkflowTemplateId pulumi.StringOutput    `pulumi:"dataprocWorkflowTemplateId"`
	EffectiveLabels            pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Optional. The encryption configuration for the workflow template.
	EncryptionConfig DataprocWorkflowTemplateEncryptionConfigPtrOutput `pulumi:"encryptionConfig"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs DataprocWorkflowTemplateJobArrayOutput `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
	// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
	// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Output only. The resource name of the workflow template, as described in
	// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
	// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
	// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
	// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters DataprocWorkflowTemplateParameterArrayOutput `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement DataprocWorkflowTemplatePlacementOutput `pulumi:"placement"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput                    `pulumi:"terraformLabels"`
	Timeouts        DataprocWorkflowTemplateTimeoutsPtrOutput `pulumi:"timeouts"`
	// Output only. The time template was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. The current version of this workflow template.
	//
	// Deprecated: Deprecated
	Version pulumi.Float64Output `pulumi:"version"`
}

// NewDataprocWorkflowTemplate registers a new resource with the given unique name, arguments, and options.
func NewDataprocWorkflowTemplate(ctx *pulumi.Context,
	name string, args *DataprocWorkflowTemplateArgs, opts ...pulumi.ResourceOption) (*DataprocWorkflowTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Jobs == nil {
		return nil, errors.New("invalid value for required argument 'Jobs'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Placement == nil {
		return nil, errors.New("invalid value for required argument 'Placement'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataprocWorkflowTemplate
	err = ctx.RegisterPackageResource("google-beta:index/dataprocWorkflowTemplate:DataprocWorkflowTemplate", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataprocWorkflowTemplate gets an existing DataprocWorkflowTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataprocWorkflowTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataprocWorkflowTemplateState, opts ...pulumi.ResourceOption) (*DataprocWorkflowTemplate, error) {
	var resource DataprocWorkflowTemplate
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google-beta:index/dataprocWorkflowTemplate:DataprocWorkflowTemplate", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataprocWorkflowTemplate resources.
type dataprocWorkflowTemplateState struct {
	// Output only. The time template was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
	// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
	// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
	// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
	// on a [managed
	// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	// the cluster is deleted.
	DagTimeout                 *string           `pulumi:"dagTimeout"`
	DataprocWorkflowTemplateId *string           `pulumi:"dataprocWorkflowTemplateId"`
	EffectiveLabels            map[string]string `pulumi:"effectiveLabels"`
	// Optional. The encryption configuration for the workflow template.
	EncryptionConfig *DataprocWorkflowTemplateEncryptionConfig `pulumi:"encryptionConfig"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []DataprocWorkflowTemplateJob `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
	// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
	// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Output only. The resource name of the workflow template, as described in
	// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
	// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
	// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
	// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name *string `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters []DataprocWorkflowTemplateParameter `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement *DataprocWorkflowTemplatePlacement `pulumi:"placement"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string                 `pulumi:"terraformLabels"`
	Timeouts        *DataprocWorkflowTemplateTimeouts `pulumi:"timeouts"`
	// Output only. The time template was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. The current version of this workflow template.
	//
	// Deprecated: Deprecated
	Version *float64 `pulumi:"version"`
}

type DataprocWorkflowTemplateState struct {
	// Output only. The time template was created.
	CreateTime pulumi.StringPtrInput
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
	// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
	// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
	// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
	// on a [managed
	// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	// the cluster is deleted.
	DagTimeout                 pulumi.StringPtrInput
	DataprocWorkflowTemplateId pulumi.StringPtrInput
	EffectiveLabels            pulumi.StringMapInput
	// Optional. The encryption configuration for the workflow template.
	EncryptionConfig DataprocWorkflowTemplateEncryptionConfigPtrInput
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs DataprocWorkflowTemplateJobArrayInput
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
	// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
	// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Output only. The resource name of the workflow template, as described in
	// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
	// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
	// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
	// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name pulumi.StringPtrInput
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters DataprocWorkflowTemplateParameterArrayInput
	// Required. WorkflowTemplate scheduling information.
	Placement DataprocWorkflowTemplatePlacementPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataprocWorkflowTemplateTimeoutsPtrInput
	// Output only. The time template was last updated.
	UpdateTime pulumi.StringPtrInput
	// Output only. The current version of this workflow template.
	//
	// Deprecated: Deprecated
	Version pulumi.Float64PtrInput
}

func (DataprocWorkflowTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocWorkflowTemplateState)(nil)).Elem()
}

type dataprocWorkflowTemplateArgs struct {
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
	// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
	// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
	// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
	// on a [managed
	// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	// the cluster is deleted.
	DagTimeout                 *string `pulumi:"dagTimeout"`
	DataprocWorkflowTemplateId *string `pulumi:"dataprocWorkflowTemplateId"`
	// Optional. The encryption configuration for the workflow template.
	EncryptionConfig *DataprocWorkflowTemplateEncryptionConfig `pulumi:"encryptionConfig"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []DataprocWorkflowTemplateJob `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
	// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
	// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Output only. The resource name of the workflow template, as described in
	// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
	// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
	// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
	// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name *string `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters []DataprocWorkflowTemplateParameter `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement DataprocWorkflowTemplatePlacement `pulumi:"placement"`
	// The project for the resource
	Project  *string                           `pulumi:"project"`
	Timeouts *DataprocWorkflowTemplateTimeouts `pulumi:"timeouts"`
	// Output only. The current version of this workflow template.
	//
	// Deprecated: Deprecated
	Version *float64 `pulumi:"version"`
}

// The set of arguments for constructing a DataprocWorkflowTemplate resource.
type DataprocWorkflowTemplateArgs struct {
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
	// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
	// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
	// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
	// on a [managed
	// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
	// the cluster is deleted.
	DagTimeout                 pulumi.StringPtrInput
	DataprocWorkflowTemplateId pulumi.StringPtrInput
	// Optional. The encryption configuration for the workflow template.
	EncryptionConfig DataprocWorkflowTemplateEncryptionConfigPtrInput
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs DataprocWorkflowTemplateJobArrayInput
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
	// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
	// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
	// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
	// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
	// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// Output only. The resource name of the workflow template, as described in
	// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
	// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
	// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
	// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
	Name pulumi.StringPtrInput
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
	// when the template is instantiated.
	Parameters DataprocWorkflowTemplateParameterArrayInput
	// Required. WorkflowTemplate scheduling information.
	Placement DataprocWorkflowTemplatePlacementInput
	// The project for the resource
	Project  pulumi.StringPtrInput
	Timeouts DataprocWorkflowTemplateTimeoutsPtrInput
	// Output only. The current version of this workflow template.
	//
	// Deprecated: Deprecated
	Version pulumi.Float64PtrInput
}

func (DataprocWorkflowTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataprocWorkflowTemplateArgs)(nil)).Elem()
}

type DataprocWorkflowTemplateInput interface {
	pulumi.Input

	ToDataprocWorkflowTemplateOutput() DataprocWorkflowTemplateOutput
	ToDataprocWorkflowTemplateOutputWithContext(ctx context.Context) DataprocWorkflowTemplateOutput
}

func (*DataprocWorkflowTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocWorkflowTemplate)(nil)).Elem()
}

func (i *DataprocWorkflowTemplate) ToDataprocWorkflowTemplateOutput() DataprocWorkflowTemplateOutput {
	return i.ToDataprocWorkflowTemplateOutputWithContext(context.Background())
}

func (i *DataprocWorkflowTemplate) ToDataprocWorkflowTemplateOutputWithContext(ctx context.Context) DataprocWorkflowTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataprocWorkflowTemplateOutput)
}

type DataprocWorkflowTemplateOutput struct{ *pulumi.OutputState }

func (DataprocWorkflowTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataprocWorkflowTemplate)(nil)).Elem()
}

func (o DataprocWorkflowTemplateOutput) ToDataprocWorkflowTemplateOutput() DataprocWorkflowTemplateOutput {
	return o
}

func (o DataprocWorkflowTemplateOutput) ToDataprocWorkflowTemplateOutputWithContext(ctx context.Context) DataprocWorkflowTemplateOutput {
	return o
}

// Output only. The time template was created.
func (o DataprocWorkflowTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see [JSON representation of
// duration](https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10
// minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at
// the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running
// on a [managed
// cluster](https://www.terraform.io/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
// the cluster is deleted.
func (o DataprocWorkflowTemplateOutput) DagTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringPtrOutput { return v.DagTimeout }).(pulumi.StringPtrOutput)
}

func (o DataprocWorkflowTemplateOutput) DataprocWorkflowTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.DataprocWorkflowTemplateId }).(pulumi.StringOutput)
}

func (o DataprocWorkflowTemplateOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Optional. The encryption configuration for the workflow template.
func (o DataprocWorkflowTemplateOutput) EncryptionConfig() DataprocWorkflowTemplateEncryptionConfigPtrOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) DataprocWorkflowTemplateEncryptionConfigPtrOutput {
		return v.EncryptionConfig
	}).(DataprocWorkflowTemplateEncryptionConfigPtrOutput)
}

// Required. The Directed Acyclic Graph of Jobs to submit.
func (o DataprocWorkflowTemplateOutput) Jobs() DataprocWorkflowTemplateJobArrayOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) DataprocWorkflowTemplateJobArrayOutput { return v.Jobs }).(DataprocWorkflowTemplateJobArrayOutput)
}

// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created
// by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to [RFC
// 1035](https://www.ietf.org/rfc/rfc1035.txt). Label **values** may be empty, but, if present, must contain 1 to 63
// characters, and must conform to [RFC 1035](https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be
// associated with a template. **Note**: This field is non-authoritative, and will only manage the labels present in your
// configuration. Please refer to the field `effective_labels` for all of the labels present on the resource.
func (o DataprocWorkflowTemplateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o DataprocWorkflowTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Output only. The resource name of the workflow template, as described in
// https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of
// the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For
// `projects.locations.workflowTemplates`, the resource name of the template has the following format:
// `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
func (o DataprocWorkflowTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided
// when the template is instantiated.
func (o DataprocWorkflowTemplateOutput) Parameters() DataprocWorkflowTemplateParameterArrayOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) DataprocWorkflowTemplateParameterArrayOutput { return v.Parameters }).(DataprocWorkflowTemplateParameterArrayOutput)
}

// Required. WorkflowTemplate scheduling information.
func (o DataprocWorkflowTemplateOutput) Placement() DataprocWorkflowTemplatePlacementOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) DataprocWorkflowTemplatePlacementOutput { return v.Placement }).(DataprocWorkflowTemplatePlacementOutput)
}

// The project for the resource
func (o DataprocWorkflowTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataprocWorkflowTemplateOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataprocWorkflowTemplateOutput) Timeouts() DataprocWorkflowTemplateTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) DataprocWorkflowTemplateTimeoutsPtrOutput { return v.Timeouts }).(DataprocWorkflowTemplateTimeoutsPtrOutput)
}

// Output only. The time template was last updated.
func (o DataprocWorkflowTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Output only. The current version of this workflow template.
//
// Deprecated: Deprecated
func (o DataprocWorkflowTemplateOutput) Version() pulumi.Float64Output {
	return o.ApplyT(func(v *DataprocWorkflowTemplate) pulumi.Float64Output { return v.Version }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataprocWorkflowTemplateInput)(nil)).Elem(), &DataprocWorkflowTemplate{})
	pulumi.RegisterOutputType(DataprocWorkflowTemplateOutput{})
}
