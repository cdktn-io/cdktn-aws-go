package awsbedrock


// Experimental.
type AwsBedrockEvaluationJob_HumanWorkflowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#flow_definition_arn AwsBedrockEvaluationJob#flow_definition_arn}.
	// Experimental.
	FlowDefinitionArn *string `field:"required" json:"flowDefinitionArn" yaml:"flowDefinitionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#instructions AwsBedrockEvaluationJob#instructions}.
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
}

