package awsamp

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPrometheusWorkspaceConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#workspace_id AwsPrometheusWorkspaceConfiguration#workspace_id}.
	// Experimental.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// limits_per_label_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#limits_per_label_set AwsPrometheusWorkspaceConfiguration#limits_per_label_set}
	// Experimental.
	LimitsPerLabelSet interface{} `field:"optional" json:"limitsPerLabelSet" yaml:"limitsPerLabelSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#out_of_order_time_window_in_seconds AwsPrometheusWorkspaceConfiguration#out_of_order_time_window_in_seconds}.
	// Experimental.
	OutOfOrderTimeWindowInSeconds *float64 `field:"optional" json:"outOfOrderTimeWindowInSeconds" yaml:"outOfOrderTimeWindowInSeconds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#region AwsPrometheusWorkspaceConfiguration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#retention_period_in_days AwsPrometheusWorkspaceConfiguration#retention_period_in_days}.
	// Experimental.
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#rule_query_offset_in_seconds AwsPrometheusWorkspaceConfiguration#rule_query_offset_in_seconds}.
	// Experimental.
	RuleQueryOffsetInSeconds *float64 `field:"optional" json:"ruleQueryOffsetInSeconds" yaml:"ruleQueryOffsetInSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/prometheus_workspace_configuration#timeouts AwsPrometheusWorkspaceConfiguration#timeouts}
	// Experimental.
	Timeouts *AwsPrometheusWorkspaceConfiguration_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

