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

type DataplexTask struct {
	pulumi.CustomResourceState

	// The time when the task was created.
	CreateTime     pulumi.StringOutput `pulumi:"createTime"`
	DataplexTaskId pulumi.StringOutput `pulumi:"dataplexTaskId"`
	// User-provided description of the task.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User friendly display name.
	DisplayName     pulumi.StringPtrOutput `pulumi:"displayName"`
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Configuration for the cluster
	ExecutionSpec DataplexTaskExecutionSpecOutput `pulumi:"executionSpec"`
	// Configuration for the cluster
	ExecutionStatuses DataplexTaskExecutionStatusArrayOutput `pulumi:"executionStatuses"`
	// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The lake in which the task will be created in.
	Lake pulumi.StringPtrOutput `pulumi:"lake"`
	// The location in which the task will be created in.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The relative resource name of the task, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/
	// tasks/{name}.
	Name pulumi.StringOutput `pulumi:"name"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Notebook DataplexTaskNotebookPtrOutput `pulumi:"notebook"`
	Project  pulumi.StringOutput           `pulumi:"project"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Spark DataplexTaskSparkPtrOutput `pulumi:"spark"`
	// Current state of the task.
	State pulumi.StringOutput `pulumi:"state"`
	// The task Id of the task.
	TaskId pulumi.StringPtrOutput `pulumi:"taskId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapOutput        `pulumi:"terraformLabels"`
	Timeouts        DataplexTaskTimeoutsPtrOutput `pulumi:"timeouts"`
	// Configuration for the cluster
	TriggerSpec DataplexTaskTriggerSpecOutput `pulumi:"triggerSpec"`
	// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with
	// the same name.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the task was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDataplexTask registers a new resource with the given unique name, arguments, and options.
func NewDataplexTask(ctx *pulumi.Context,
	name string, args *DataplexTaskArgs, opts ...pulumi.ResourceOption) (*DataplexTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ExecutionSpec == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionSpec'")
	}
	if args.TriggerSpec == nil {
		return nil, errors.New("invalid value for required argument 'TriggerSpec'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var resource DataplexTask
	err = ctx.RegisterPackageResource("google:index/dataplexTask:DataplexTask", name, args, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplexTask gets an existing DataplexTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplexTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplexTaskState, opts ...pulumi.ResourceOption) (*DataplexTask, error) {
	var resource DataplexTask
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	err = ctx.ReadPackageResource("google:index/dataplexTask:DataplexTask", name, id, state, &resource, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplexTask resources.
type dataplexTaskState struct {
	// The time when the task was created.
	CreateTime     *string `pulumi:"createTime"`
	DataplexTaskId *string `pulumi:"dataplexTaskId"`
	// User-provided description of the task.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName     *string           `pulumi:"displayName"`
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Configuration for the cluster
	ExecutionSpec *DataplexTaskExecutionSpec `pulumi:"executionSpec"`
	// Configuration for the cluster
	ExecutionStatuses []DataplexTaskExecutionStatus `pulumi:"executionStatuses"`
	// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake in which the task will be created in.
	Lake *string `pulumi:"lake"`
	// The location in which the task will be created in.
	Location *string `pulumi:"location"`
	// The relative resource name of the task, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/
	// tasks/{name}.
	Name *string `pulumi:"name"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Notebook *DataplexTaskNotebook `pulumi:"notebook"`
	Project  *string               `pulumi:"project"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Spark *DataplexTaskSpark `pulumi:"spark"`
	// Current state of the task.
	State *string `pulumi:"state"`
	// The task Id of the task.
	TaskId *string `pulumi:"taskId"`
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels map[string]string     `pulumi:"terraformLabels"`
	Timeouts        *DataplexTaskTimeouts `pulumi:"timeouts"`
	// Configuration for the cluster
	TriggerSpec *DataplexTaskTriggerSpec `pulumi:"triggerSpec"`
	// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with
	// the same name.
	Uid *string `pulumi:"uid"`
	// The time when the task was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type DataplexTaskState struct {
	// The time when the task was created.
	CreateTime     pulumi.StringPtrInput
	DataplexTaskId pulumi.StringPtrInput
	// User-provided description of the task.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName     pulumi.StringPtrInput
	EffectiveLabels pulumi.StringMapInput
	// Configuration for the cluster
	ExecutionSpec DataplexTaskExecutionSpecPtrInput
	// Configuration for the cluster
	ExecutionStatuses DataplexTaskExecutionStatusArrayInput
	// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake in which the task will be created in.
	Lake pulumi.StringPtrInput
	// The location in which the task will be created in.
	Location pulumi.StringPtrInput
	// The relative resource name of the task, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/
	// tasks/{name}.
	Name pulumi.StringPtrInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Notebook DataplexTaskNotebookPtrInput
	Project  pulumi.StringPtrInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Spark DataplexTaskSparkPtrInput
	// Current state of the task.
	State pulumi.StringPtrInput
	// The task Id of the task.
	TaskId pulumi.StringPtrInput
	// The combination of labels configured directly on the resource and default labels configured on the provider.
	TerraformLabels pulumi.StringMapInput
	Timeouts        DataplexTaskTimeoutsPtrInput
	// Configuration for the cluster
	TriggerSpec DataplexTaskTriggerSpecPtrInput
	// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with
	// the same name.
	Uid pulumi.StringPtrInput
	// The time when the task was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (DataplexTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexTaskState)(nil)).Elem()
}

type dataplexTaskArgs struct {
	DataplexTaskId *string `pulumi:"dataplexTaskId"`
	// User-provided description of the task.
	Description *string `pulumi:"description"`
	// User friendly display name.
	DisplayName *string `pulumi:"displayName"`
	// Configuration for the cluster
	ExecutionSpec DataplexTaskExecutionSpec `pulumi:"executionSpec"`
	// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The lake in which the task will be created in.
	Lake *string `pulumi:"lake"`
	// The location in which the task will be created in.
	Location *string `pulumi:"location"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Notebook *DataplexTaskNotebook `pulumi:"notebook"`
	Project  *string               `pulumi:"project"`
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Spark *DataplexTaskSpark `pulumi:"spark"`
	// The task Id of the task.
	TaskId   *string               `pulumi:"taskId"`
	Timeouts *DataplexTaskTimeouts `pulumi:"timeouts"`
	// Configuration for the cluster
	TriggerSpec DataplexTaskTriggerSpec `pulumi:"triggerSpec"`
}

// The set of arguments for constructing a DataplexTask resource.
type DataplexTaskArgs struct {
	DataplexTaskId pulumi.StringPtrInput
	// User-provided description of the task.
	Description pulumi.StringPtrInput
	// User friendly display name.
	DisplayName pulumi.StringPtrInput
	// Configuration for the cluster
	ExecutionSpec DataplexTaskExecutionSpecInput
	// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
	// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The lake in which the task will be created in.
	Lake pulumi.StringPtrInput
	// The location in which the task will be created in.
	Location pulumi.StringPtrInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Notebook DataplexTaskNotebookPtrInput
	Project  pulumi.StringPtrInput
	// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
	// its memory over time.
	Spark DataplexTaskSparkPtrInput
	// The task Id of the task.
	TaskId   pulumi.StringPtrInput
	Timeouts DataplexTaskTimeoutsPtrInput
	// Configuration for the cluster
	TriggerSpec DataplexTaskTriggerSpecInput
}

func (DataplexTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplexTaskArgs)(nil)).Elem()
}

type DataplexTaskInput interface {
	pulumi.Input

	ToDataplexTaskOutput() DataplexTaskOutput
	ToDataplexTaskOutputWithContext(ctx context.Context) DataplexTaskOutput
}

func (*DataplexTask) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexTask)(nil)).Elem()
}

func (i *DataplexTask) ToDataplexTaskOutput() DataplexTaskOutput {
	return i.ToDataplexTaskOutputWithContext(context.Background())
}

func (i *DataplexTask) ToDataplexTaskOutputWithContext(ctx context.Context) DataplexTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplexTaskOutput)
}

type DataplexTaskOutput struct{ *pulumi.OutputState }

func (DataplexTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplexTask)(nil)).Elem()
}

func (o DataplexTaskOutput) ToDataplexTaskOutput() DataplexTaskOutput {
	return o
}

func (o DataplexTaskOutput) ToDataplexTaskOutputWithContext(ctx context.Context) DataplexTaskOutput {
	return o
}

// The time when the task was created.
func (o DataplexTaskOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

func (o DataplexTaskOutput) DataplexTaskId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.DataplexTaskId }).(pulumi.StringOutput)
}

// User-provided description of the task.
func (o DataplexTaskOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User friendly display name.
func (o DataplexTaskOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o DataplexTaskOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Configuration for the cluster
func (o DataplexTaskOutput) ExecutionSpec() DataplexTaskExecutionSpecOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskExecutionSpecOutput { return v.ExecutionSpec }).(DataplexTaskExecutionSpecOutput)
}

// Configuration for the cluster
func (o DataplexTaskOutput) ExecutionStatuses() DataplexTaskExecutionStatusArrayOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskExecutionStatusArrayOutput { return v.ExecutionStatuses }).(DataplexTaskExecutionStatusArrayOutput)
}

// User-defined labels for the task. **Note**: This field is non-authoritative, and will only manage the labels present in
// your configuration. Please refer to the field 'effective_labels' for all of the labels present on the resource.
func (o DataplexTaskOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The lake in which the task will be created in.
func (o DataplexTaskOutput) Lake() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringPtrOutput { return v.Lake }).(pulumi.StringPtrOutput)
}

// The location in which the task will be created in.
func (o DataplexTaskOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The relative resource name of the task, of the form: projects/{project_number}/locations/{locationId}/lakes/{lakeId}/
// tasks/{name}.
func (o DataplexTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
// its memory over time.
func (o DataplexTaskOutput) Notebook() DataplexTaskNotebookPtrOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskNotebookPtrOutput { return v.Notebook }).(DataplexTaskNotebookPtrOutput)
}

func (o DataplexTaskOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of
// its memory over time.
func (o DataplexTaskOutput) Spark() DataplexTaskSparkPtrOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskSparkPtrOutput { return v.Spark }).(DataplexTaskSparkPtrOutput)
}

// Current state of the task.
func (o DataplexTaskOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The task Id of the task.
func (o DataplexTaskOutput) TaskId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringPtrOutput { return v.TaskId }).(pulumi.StringPtrOutput)
}

// The combination of labels configured directly on the resource and default labels configured on the provider.
func (o DataplexTaskOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringMapOutput { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o DataplexTaskOutput) Timeouts() DataplexTaskTimeoutsPtrOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskTimeoutsPtrOutput { return v.Timeouts }).(DataplexTaskTimeoutsPtrOutput)
}

// Configuration for the cluster
func (o DataplexTaskOutput) TriggerSpec() DataplexTaskTriggerSpecOutput {
	return o.ApplyT(func(v *DataplexTask) DataplexTaskTriggerSpecOutput { return v.TriggerSpec }).(DataplexTaskTriggerSpecOutput)
}

// System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with
// the same name.
func (o DataplexTaskOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the task was last updated.
func (o DataplexTaskOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplexTask) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplexTaskInput)(nil)).Elem(), &DataplexTask{})
	pulumi.RegisterOutputType(DataplexTaskOutput{})
}
