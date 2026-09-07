package bedrockagents

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAgentActionGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#action_group_name AwsAgentActionGroup#action_group_name}.
	// Experimental.
	ActionGroupName *string `field:"required" json:"actionGroupName" yaml:"actionGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#agent_id AwsAgentActionGroup#agent_id}.
	// Experimental.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#agent_version AwsAgentActionGroup#agent_version}.
	// Experimental.
	AgentVersion *string `field:"required" json:"agentVersion" yaml:"agentVersion"`
	// action_group_executor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#action_group_executor AwsAgentActionGroup#action_group_executor}
	// Experimental.
	ActionGroupExecutor interface{} `field:"optional" json:"actionGroupExecutor" yaml:"actionGroupExecutor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#action_group_state AwsAgentActionGroup#action_group_state}.
	// Experimental.
	ActionGroupState *string `field:"optional" json:"actionGroupState" yaml:"actionGroupState"`
	// api_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#api_schema AwsAgentActionGroup#api_schema}
	// Experimental.
	ApiSchema interface{} `field:"optional" json:"apiSchema" yaml:"apiSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#description AwsAgentActionGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// function_schema block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#function_schema AwsAgentActionGroup#function_schema}
	// Experimental.
	FunctionSchema interface{} `field:"optional" json:"functionSchema" yaml:"functionSchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#parent_action_group_signature AwsAgentActionGroup#parent_action_group_signature}.
	// Experimental.
	ParentActionGroupSignature *string `field:"optional" json:"parentActionGroupSignature" yaml:"parentActionGroupSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#prepare_agent AwsAgentActionGroup#prepare_agent}.
	// Experimental.
	PrepareAgent interface{} `field:"optional" json:"prepareAgent" yaml:"prepareAgent"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#region AwsAgentActionGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#skip_resource_in_use_check AwsAgentActionGroup#skip_resource_in_use_check}.
	// Experimental.
	SkipResourceInUseCheck interface{} `field:"optional" json:"skipResourceInUseCheck" yaml:"skipResourceInUseCheck"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group#timeouts AwsAgentActionGroup#timeouts}
	// Experimental.
	Timeouts *AwsAgentActionGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

