package awsbedrockagents

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAgentCollaboratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#agent_id TfAgentCollaborator#agent_id}.
	// Experimental.
	AgentId *string `field:"required" json:"agentId" yaml:"agentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#collaboration_instruction TfAgentCollaborator#collaboration_instruction}.
	// Experimental.
	CollaborationInstruction *string `field:"required" json:"collaborationInstruction" yaml:"collaborationInstruction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#collaborator_name TfAgentCollaborator#collaborator_name}.
	// Experimental.
	CollaboratorName *string `field:"required" json:"collaboratorName" yaml:"collaboratorName"`
	// agent_descriptor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#agent_descriptor TfAgentCollaborator#agent_descriptor}
	// Experimental.
	AgentDescriptor interface{} `field:"optional" json:"agentDescriptor" yaml:"agentDescriptor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#agent_version TfAgentCollaborator#agent_version}.
	// Experimental.
	AgentVersion *string `field:"optional" json:"agentVersion" yaml:"agentVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#prepare_agent TfAgentCollaborator#prepare_agent}.
	// Experimental.
	PrepareAgent interface{} `field:"optional" json:"prepareAgent" yaml:"prepareAgent"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#region TfAgentCollaborator#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#relay_conversation_history TfAgentCollaborator#relay_conversation_history}.
	// Experimental.
	RelayConversationHistory *string `field:"optional" json:"relayConversationHistory" yaml:"relayConversationHistory"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_collaborator#timeouts TfAgentCollaborator#timeouts}
	// Experimental.
	Timeouts *TfAgentCollaborator_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

