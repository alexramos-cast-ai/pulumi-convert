// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package google

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-terraform-provider/sdks/go/google/v6/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCloudbuildTrigger(ctx *pulumi.Context, args *LookupCloudbuildTriggerArgs, opts ...pulumi.InvokeOption) (*LookupCloudbuildTriggerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	ref, err := internal.PkgGetPackageRef(ctx)
	if err != nil {
		return nil, err
	}
	var rv LookupCloudbuildTriggerResult
	err = ctx.InvokePackage("google:index/getCloudbuildTrigger:getCloudbuildTrigger", args, &rv, ref, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCloudbuildTrigger.
type LookupCloudbuildTriggerArgs struct {
	Id        *string `pulumi:"id"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
	TriggerId string  `pulumi:"triggerId"`
}

// A collection of values returned by getCloudbuildTrigger.
type LookupCloudbuildTriggerResult struct {
	ApprovalConfigs               []GetCloudbuildTriggerApprovalConfig               `pulumi:"approvalConfigs"`
	BitbucketServerTriggerConfigs []GetCloudbuildTriggerBitbucketServerTriggerConfig `pulumi:"bitbucketServerTriggerConfigs"`
	Builds                        []GetCloudbuildTriggerBuild                        `pulumi:"builds"`
	CreateTime                    string                                             `pulumi:"createTime"`
	Description                   string                                             `pulumi:"description"`
	Disabled                      bool                                               `pulumi:"disabled"`
	Filename                      string                                             `pulumi:"filename"`
	Filter                        string                                             `pulumi:"filter"`
	GitFileSources                []GetCloudbuildTriggerGitFileSource                `pulumi:"gitFileSources"`
	Githubs                       []GetCloudbuildTriggerGithub                       `pulumi:"githubs"`
	Id                            string                                             `pulumi:"id"`
	IgnoredFiles                  []string                                           `pulumi:"ignoredFiles"`
	IncludeBuildLogs              string                                             `pulumi:"includeBuildLogs"`
	IncludedFiles                 []string                                           `pulumi:"includedFiles"`
	Location                      string                                             `pulumi:"location"`
	Name                          string                                             `pulumi:"name"`
	Project                       *string                                            `pulumi:"project"`
	PubsubConfigs                 []GetCloudbuildTriggerPubsubConfig                 `pulumi:"pubsubConfigs"`
	RepositoryEventConfigs        []GetCloudbuildTriggerRepositoryEventConfig        `pulumi:"repositoryEventConfigs"`
	ServiceAccount                string                                             `pulumi:"serviceAccount"`
	SourceToBuilds                []GetCloudbuildTriggerSourceToBuild                `pulumi:"sourceToBuilds"`
	Substitutions                 map[string]string                                  `pulumi:"substitutions"`
	Tags                          []string                                           `pulumi:"tags"`
	TriggerId                     string                                             `pulumi:"triggerId"`
	TriggerTemplates              []GetCloudbuildTriggerTriggerTemplate              `pulumi:"triggerTemplates"`
	WebhookConfigs                []GetCloudbuildTriggerWebhookConfig                `pulumi:"webhookConfigs"`
}

func LookupCloudbuildTriggerOutput(ctx *pulumi.Context, args LookupCloudbuildTriggerOutputArgs, opts ...pulumi.InvokeOption) LookupCloudbuildTriggerResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCloudbuildTriggerResultOutput, error) {
			args := v.(LookupCloudbuildTriggerArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			ref, err := internal.PkgGetPackageRef(ctx)
			if err != nil {
				return LookupCloudbuildTriggerResultOutput{}, err
			}
			options.PackageRef = ref
			return ctx.InvokeOutput("google:index/getCloudbuildTrigger:getCloudbuildTrigger", args, LookupCloudbuildTriggerResultOutput{}, options).(LookupCloudbuildTriggerResultOutput), nil
		}).(LookupCloudbuildTriggerResultOutput)
}

// A collection of arguments for invoking getCloudbuildTrigger.
type LookupCloudbuildTriggerOutputArgs struct {
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
	TriggerId pulumi.StringInput    `pulumi:"triggerId"`
}

func (LookupCloudbuildTriggerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudbuildTriggerArgs)(nil)).Elem()
}

// A collection of values returned by getCloudbuildTrigger.
type LookupCloudbuildTriggerResultOutput struct{ *pulumi.OutputState }

func (LookupCloudbuildTriggerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCloudbuildTriggerResult)(nil)).Elem()
}

func (o LookupCloudbuildTriggerResultOutput) ToLookupCloudbuildTriggerResultOutput() LookupCloudbuildTriggerResultOutput {
	return o
}

func (o LookupCloudbuildTriggerResultOutput) ToLookupCloudbuildTriggerResultOutputWithContext(ctx context.Context) LookupCloudbuildTriggerResultOutput {
	return o
}

func (o LookupCloudbuildTriggerResultOutput) ApprovalConfigs() GetCloudbuildTriggerApprovalConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerApprovalConfig { return v.ApprovalConfigs }).(GetCloudbuildTriggerApprovalConfigArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) BitbucketServerTriggerConfigs() GetCloudbuildTriggerBitbucketServerTriggerConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerBitbucketServerTriggerConfig {
		return v.BitbucketServerTriggerConfigs
	}).(GetCloudbuildTriggerBitbucketServerTriggerConfigArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Builds() GetCloudbuildTriggerBuildArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerBuild { return v.Builds }).(GetCloudbuildTriggerBuildArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Disabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) bool { return v.Disabled }).(pulumi.BoolOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Filename() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Filename }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Filter() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Filter }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) GitFileSources() GetCloudbuildTriggerGitFileSourceArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerGitFileSource { return v.GitFileSources }).(GetCloudbuildTriggerGitFileSourceArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Githubs() GetCloudbuildTriggerGithubArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerGithub { return v.Githubs }).(GetCloudbuildTriggerGithubArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) IgnoredFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []string { return v.IgnoredFiles }).(pulumi.StringArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) IncludeBuildLogs() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.IncludeBuildLogs }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) IncludedFiles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []string { return v.IncludedFiles }).(pulumi.StringArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Project() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) *string { return v.Project }).(pulumi.StringPtrOutput)
}

func (o LookupCloudbuildTriggerResultOutput) PubsubConfigs() GetCloudbuildTriggerPubsubConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerPubsubConfig { return v.PubsubConfigs }).(GetCloudbuildTriggerPubsubConfigArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) RepositoryEventConfigs() GetCloudbuildTriggerRepositoryEventConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerRepositoryEventConfig {
		return v.RepositoryEventConfigs
	}).(GetCloudbuildTriggerRepositoryEventConfigArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) SourceToBuilds() GetCloudbuildTriggerSourceToBuildArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerSourceToBuild { return v.SourceToBuilds }).(GetCloudbuildTriggerSourceToBuildArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Substitutions() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) map[string]string { return v.Substitutions }).(pulumi.StringMapOutput)
}

func (o LookupCloudbuildTriggerResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) TriggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) string { return v.TriggerId }).(pulumi.StringOutput)
}

func (o LookupCloudbuildTriggerResultOutput) TriggerTemplates() GetCloudbuildTriggerTriggerTemplateArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerTriggerTemplate { return v.TriggerTemplates }).(GetCloudbuildTriggerTriggerTemplateArrayOutput)
}

func (o LookupCloudbuildTriggerResultOutput) WebhookConfigs() GetCloudbuildTriggerWebhookConfigArrayOutput {
	return o.ApplyT(func(v LookupCloudbuildTriggerResult) []GetCloudbuildTriggerWebhookConfig { return v.WebhookConfigs }).(GetCloudbuildTriggerWebhookConfigArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCloudbuildTriggerResultOutput{})
}
