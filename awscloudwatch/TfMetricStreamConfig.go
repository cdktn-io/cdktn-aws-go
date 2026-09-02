package awscloudwatch

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMetricStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#firehose_arn TfMetricStream#firehose_arn}.
	// Experimental.
	FirehoseArn *string `field:"required" json:"firehoseArn" yaml:"firehoseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#output_format TfMetricStream#output_format}.
	// Experimental.
	OutputFormat *string `field:"required" json:"outputFormat" yaml:"outputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#role_arn TfMetricStream#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// exclude_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#exclude_filter TfMetricStream#exclude_filter}
	// Experimental.
	ExcludeFilter interface{} `field:"optional" json:"excludeFilter" yaml:"excludeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#id TfMetricStream#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// include_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#include_filter TfMetricStream#include_filter}
	// Experimental.
	IncludeFilter interface{} `field:"optional" json:"includeFilter" yaml:"includeFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#include_linked_accounts_metrics TfMetricStream#include_linked_accounts_metrics}.
	// Experimental.
	IncludeLinkedAccountsMetrics interface{} `field:"optional" json:"includeLinkedAccountsMetrics" yaml:"includeLinkedAccountsMetrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#name TfMetricStream#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#name_prefix TfMetricStream#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#region TfMetricStream#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// statistics_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#statistics_configuration TfMetricStream#statistics_configuration}
	// Experimental.
	StatisticsConfiguration interface{} `field:"optional" json:"statisticsConfiguration" yaml:"statisticsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#tags TfMetricStream#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#tags_all TfMetricStream#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_metric_stream#timeouts TfMetricStream#timeouts}
	// Experimental.
	Timeouts *TfMetricStream_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

