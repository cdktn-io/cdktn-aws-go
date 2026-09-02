package awsbedrockagents

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfKnowledgeBaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#name TfKnowledgeBase#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#role_arn TfKnowledgeBase#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#description TfKnowledgeBase#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// knowledge_base_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#knowledge_base_configuration TfKnowledgeBase#knowledge_base_configuration}
	// Experimental.
	KnowledgeBaseConfiguration interface{} `field:"optional" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#region TfKnowledgeBase#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#storage_configuration TfKnowledgeBase#storage_configuration}
	// Experimental.
	StorageConfiguration interface{} `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#tags TfKnowledgeBase#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#timeouts TfKnowledgeBase#timeouts}
	// Experimental.
	Timeouts *TfKnowledgeBase_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

