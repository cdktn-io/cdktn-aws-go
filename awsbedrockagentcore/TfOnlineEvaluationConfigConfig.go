package awsbedrockagentcore

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOnlineEvaluationConfigConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#enable_on_create TfOnlineEvaluationConfig#enable_on_create}.
	// Experimental.
	EnableOnCreate interface{} `field:"required" json:"enableOnCreate" yaml:"enableOnCreate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#evaluation_execution_role_arn TfOnlineEvaluationConfig#evaluation_execution_role_arn}.
	// Experimental.
	EvaluationExecutionRoleArn *string `field:"required" json:"evaluationExecutionRoleArn" yaml:"evaluationExecutionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#online_evaluation_config_name TfOnlineEvaluationConfig#online_evaluation_config_name}.
	// Experimental.
	OnlineEvaluationConfigName *string `field:"required" json:"onlineEvaluationConfigName" yaml:"onlineEvaluationConfigName"`
	// data_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#data_source_config TfOnlineEvaluationConfig#data_source_config}
	// Experimental.
	DataSourceConfig interface{} `field:"optional" json:"dataSourceConfig" yaml:"dataSourceConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#description TfOnlineEvaluationConfig#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// evaluator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#evaluator TfOnlineEvaluationConfig#evaluator}
	// Experimental.
	Evaluator interface{} `field:"optional" json:"evaluator" yaml:"evaluator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#execution_status TfOnlineEvaluationConfig#execution_status}.
	// Experimental.
	ExecutionStatus *string `field:"optional" json:"executionStatus" yaml:"executionStatus"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#region TfOnlineEvaluationConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#rule TfOnlineEvaluationConfig#rule}
	// Experimental.
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#tags TfOnlineEvaluationConfig#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#timeouts TfOnlineEvaluationConfig#timeouts}
	// Experimental.
	Timeouts *TfOnlineEvaluationConfig_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

