// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudfunctionsFunction(ctx *pulumi.Context, args *LookupCloudfunctionsFunctionArgs, opts ...pulumi.InvokeOption) (*LookupCloudfunctionsFunctionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCloudfunctionsFunctionResult
	err = ctx.InvokePackage("google:index/getCloudfunctionsFunction:getCloudfunctionsFunction", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudfunctionsFunction.
type LookupCloudfunctionsFunctionArgs struct {
	Id      *string `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
}

// A collection of values returned by getCloudfunctionsFunction.
type LookupCloudfunctionsFunctionResult struct {
	AvailableMemoryMb          float64                                              `pulumi:"availableMemoryMb"`
	BuildEnvironmentVariables  map[string]string                                    `pulumi:"buildEnvironmentVariables"`
	BuildServiceAccount        string                                               `pulumi:"buildServiceAccount"`
	BuildWorkerPool            string                                               `pulumi:"buildWorkerPool"`
	Description                string                                               `pulumi:"description"`
	DockerRegistry             string                                               `pulumi:"dockerRegistry"`
	DockerRepository           string                                               `pulumi:"dockerRepository"`
	EffectiveLabels            map[string]string                                    `pulumi:"effectiveLabels"`
	EntryPoint                 string                                               `pulumi:"entryPoint"`
	EnvironmentVariables       map[string]string                                    `pulumi:"environmentVariables"`
	EventTriggers              []GetCloudfunctionsFunctionEventTrigger              `pulumi:"eventTriggers"`
	HttpsTriggerSecurityLevel  string                                               `pulumi:"httpsTriggerSecurityLevel"`
	HttpsTriggerUrl            string                                               `pulumi:"httpsTriggerUrl"`
	Id                         string                                               `pulumi:"id"`
	IngressSettings            string                                               `pulumi:"ingressSettings"`
	KmsKeyName                 string                                               `pulumi:"kmsKeyName"`
	Labels                     map[string]string                                    `pulumi:"labels"`
	MaxInstances               float64                                              `pulumi:"maxInstances"`
	MinInstances               float64                                              `pulumi:"minInstances"`
	Name                       string                                               `pulumi:"name"`
	Project                    *string                                              `pulumi:"project"`
	Region                     *string                                              `pulumi:"region"`
	Runtime                    string                                               `pulumi:"runtime"`
	SecretEnvironmentVariables []GetCloudfunctionsFunctionSecretEnvironmentVariable `pulumi:"secretEnvironmentVariables"`
	SecretVolumes              []GetCloudfunctionsFunctionSecretVolume              `pulumi:"secretVolumes"`
	ServiceAccountEmail        string                                               `pulumi:"serviceAccountEmail"`
	SourceArchiveBucket        string                                               `pulumi:"sourceArchiveBucket"`
	SourceArchiveObject        string                                               `pulumi:"sourceArchiveObject"`
	SourceRepositories         []GetCloudfunctionsFunctionSourceRepository          `pulumi:"sourceRepositories"`
	Status                     string                                               `pulumi:"status"`
	TerraformLabels            map[string]string                                    `pulumi:"terraformLabels"`
	Timeout                    float64                                              `pulumi:"timeout"`
	TriggerHttp                bool                                                 `pulumi:"triggerHttp"`
	VersionId                  string                                               `pulumi:"versionId"`
	VpcConnector               string                                               `pulumi:"vpcConnector"`
	VpcConnectorEgressSettings string                                               `pulumi:"vpcConnectorEgressSettings"`
}

func LookupCloudfunctionsFunctionOutput(ctx *pulumi.Context, args LookupCloudfunctionsFunctionOutputArgs, opts ...pulumi.InvokeOption) LookupCloudfunctionsFunctionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudfunctionsFunctionResultOutput, error) {
			args := v.(LookupCloudfunctionsFunctionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupCloudfunctionsFunctionResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getCloudfunctionsFunction:getCloudfunctionsFunction", args, LookupCloudfunctionsFunctionResultOutput{}, options).(LookupCloudfunctionsFunctionResultOutput), nil
		}).(LookupCloudfunctionsFunctionResultOutput)
}

// A collection of arguments for invoking getCloudfunctionsFunction.
type LookupCloudfunctionsFunctionOutputArgs struct {
	Id      pulumi.StringPtrInput `pulumi:"id"`
	Name    pulumi.StringInput    `pulumi:"name"`
	Project pulumi.StringPtrInput `pulumi:"project"`
	Region  pulumi.StringPtrInput `pulumi:"region"`
}

func (LookupCloudfunctionsFunctionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudfunctionsFunctionArgs)(nil)).Elem()
}

// A collection of values returned by getCloudfunctionsFunction.
type LookupCloudfunctionsFunctionResultOutput struct{ *pulumi.OutputState }

func (LookupCloudfunctionsFunctionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudfunctionsFunctionResult)(nil)).Elem()
}

func (o LookupCloudfunctionsFunctionResultOutput) ToLookupCloudfunctionsFunctionResultOutput() LookupCloudfunctionsFunctionResultOutput {
	return o
}

func (o LookupCloudfunctionsFunctionResultOutput) ToLookupCloudfunctionsFunctionResultOutputWithContext(ctx context.Context) LookupCloudfunctionsFunctionResultOutput {
	return o
}

func (o LookupCloudfunctionsFunctionResultOutput) AvailableMemoryMb() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) float64 { return v.AvailableMemoryMb }).(pulumi.Float64Output)
}

func (o LookupCloudfunctionsFunctionResultOutput) BuildEnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) map[string]string { return v.BuildEnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) BuildServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.BuildServiceAccount }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) BuildWorkerPool() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.BuildWorkerPool }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) DockerRegistry() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.DockerRegistry }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) DockerRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.DockerRepository }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) map[string]string { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) EntryPoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.EntryPoint }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) EnvironmentVariables() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) map[string]string { return v.EnvironmentVariables }).(pulumi.StringMapOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) EventTriggers() GetCloudfunctionsFunctionEventTriggerArrayOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) []GetCloudfunctionsFunctionEventTrigger {
		return v.EventTriggers
	}).(GetCloudfunctionsFunctionEventTriggerArrayOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) HttpsTriggerSecurityLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.HttpsTriggerSecurityLevel }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) HttpsTriggerUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.HttpsTriggerUrl }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) IngressSettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.IngressSettings }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) MaxInstances() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) float64 { return v.MaxInstances }).(pulumi.Float64Output)
}

func (o LookupCloudfunctionsFunctionResultOutput) MinInstances() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) float64 { return v.MinInstances }).(pulumi.Float64Output)
}

func (o LookupCloudfunctionsFunctionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) *string { return v.Region }).(pulumi.StringPtrOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.Runtime }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) SecretEnvironmentVariables() GetCloudfunctionsFunctionSecretEnvironmentVariableArrayOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) []GetCloudfunctionsFunctionSecretEnvironmentVariable {
		return v.SecretEnvironmentVariables
	}).(GetCloudfunctionsFunctionSecretEnvironmentVariableArrayOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) SecretVolumes() GetCloudfunctionsFunctionSecretVolumeArrayOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) []GetCloudfunctionsFunctionSecretVolume {
		return v.SecretVolumes
	}).(GetCloudfunctionsFunctionSecretVolumeArrayOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) ServiceAccountEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.ServiceAccountEmail }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) SourceArchiveBucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.SourceArchiveBucket }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) SourceArchiveObject() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.SourceArchiveObject }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) SourceRepositories() GetCloudfunctionsFunctionSourceRepositoryArrayOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) []GetCloudfunctionsFunctionSourceRepository {
		return v.SourceRepositories
	}).(GetCloudfunctionsFunctionSourceRepositoryArrayOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) TerraformLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) map[string]string { return v.TerraformLabels }).(pulumi.StringMapOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) Timeout() pulumi.Float64Output {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) float64 { return v.Timeout }).(pulumi.Float64Output)
}

func (o LookupCloudfunctionsFunctionResultOutput) TriggerHttp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) bool { return v.TriggerHttp }).(pulumi.BoolOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.VersionId }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) VpcConnector() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.VpcConnector }).(pulumi.StringOutput)
}

func (o LookupCloudfunctionsFunctionResultOutput) VpcConnectorEgressSettings() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudfunctionsFunctionResult) string { return v.VpcConnectorEgressSettings }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudfunctionsFunctionResultOutput{})
}
