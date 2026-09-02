package awsbedrockagents

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAgentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#agent_name TfAgent#agent_name}.
	// Experimental.
	AgentName *string `field:"required" json:"agentName" yaml:"agentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#agent_resource_role_arn TfAgent#agent_resource_role_arn}.
	// Experimental.
	AgentResourceRoleArn *string `field:"required" json:"agentResourceRoleArn" yaml:"agentResourceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#foundation_model TfAgent#foundation_model}.
	// Experimental.
	FoundationModel *string `field:"required" json:"foundationModel" yaml:"foundationModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#agent_collaboration TfAgent#agent_collaboration}.
	// Experimental.
	AgentCollaboration *string `field:"optional" json:"agentCollaboration" yaml:"agentCollaboration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#customer_encryption_key_arn TfAgent#customer_encryption_key_arn}.
	// Experimental.
	CustomerEncryptionKeyArn *string `field:"optional" json:"customerEncryptionKeyArn" yaml:"customerEncryptionKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#description TfAgent#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#guardrail_configuration TfAgent#guardrail_configuration}.
	// Experimental.
	GuardrailConfiguration interface{} `field:"optional" json:"guardrailConfiguration" yaml:"guardrailConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#idle_session_ttl_in_seconds TfAgent#idle_session_ttl_in_seconds}.
	// Experimental.
	IdleSessionTtlInSeconds *float64 `field:"optional" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#instruction TfAgent#instruction}.
	// Experimental.
	Instruction *string `field:"optional" json:"instruction" yaml:"instruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#memory_configuration TfAgent#memory_configuration}.
	// Experimental.
	MemoryConfiguration interface{} `field:"optional" json:"memoryConfiguration" yaml:"memoryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prepare_agent TfAgent#prepare_agent}.
	// Experimental.
	PrepareAgent interface{} `field:"optional" json:"prepareAgent" yaml:"prepareAgent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#prompt_override_configuration TfAgent#prompt_override_configuration}.
	// Experimental.
	PromptOverrideConfiguration interface{} `field:"optional" json:"promptOverrideConfiguration" yaml:"promptOverrideConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#region TfAgent#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#skip_resource_in_use_check TfAgent#skip_resource_in_use_check}.
	// Experimental.
	SkipResourceInUseCheck interface{} `field:"optional" json:"skipResourceInUseCheck" yaml:"skipResourceInUseCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#tags TfAgent#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#timeouts TfAgent#timeouts}
	// Experimental.
	Timeouts *TfAgent_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

