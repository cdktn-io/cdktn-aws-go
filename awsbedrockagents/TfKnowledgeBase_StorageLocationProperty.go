package awsbedrockagents


// Experimental.
type TfKnowledgeBase_StorageLocationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#type TfKnowledgeBase#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// s3_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_knowledge_base#s3_location TfKnowledgeBase#s3_location}
	// Experimental.
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

