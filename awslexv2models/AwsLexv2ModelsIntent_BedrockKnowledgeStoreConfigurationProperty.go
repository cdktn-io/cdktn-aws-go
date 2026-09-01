package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_BedrockKnowledgeStoreConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bedrock_knowledge_base_arn AwsLexv2ModelsIntent#bedrock_knowledge_base_arn}.
	// Experimental.
	BedrockKnowledgeBaseArn *string `field:"required" json:"bedrockKnowledgeBaseArn" yaml:"bedrockKnowledgeBaseArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response AwsLexv2ModelsIntent#exact_response}.
	// Experimental.
	ExactResponse interface{} `field:"optional" json:"exactResponse" yaml:"exactResponse"`
	// exact_response_fields block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#exact_response_fields AwsLexv2ModelsIntent#exact_response_fields}
	// Experimental.
	ExactResponseFields interface{} `field:"optional" json:"exactResponseFields" yaml:"exactResponseFields"`
}

