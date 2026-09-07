package bedrock


// Experimental.
type AwsEvaluationJob_HumanWorkflowConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#flow_definition_arn AwsEvaluationJob#flow_definition_arn}.
	// Experimental.
	FlowDefinitionArn *string `field:"required" json:"flowDefinitionArn" yaml:"flowDefinitionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#instructions AwsEvaluationJob#instructions}.
	// Experimental.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
}

