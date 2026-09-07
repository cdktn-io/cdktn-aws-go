package lexv2models


// Experimental.
type AwsIntent_BedrockKnowledgeStoreConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bedrock_knowledge_base_arn AwsIntent#bedrock_knowledge_base_arn}.
	// Experimental.
	BedrockKnowledgeBaseArn *string `field:"required" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response AwsIntent#exact_response}.
	// Experimental.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// exact_response_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response_fields AwsIntent#exact_response_fields}
	// Experimental.
	ExactResponseFields interface{} `field:"optional" json:"exactResponseFields" yaml:"exactResponseFields"`
}

