package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_TransformationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#step_to_apply AwsBedrockagentDataSource#step_to_apply}.
	// Experimental.
	StepToApply *string `field:"required" json:"stepToApply" yaml:"stepToApply"`
	// transformation_function block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#transformation_function AwsBedrockagentDataSource#transformation_function}
	// Experimental.
	TransformationFunction interface{} `field:"optional" json:"transformationFunction" yaml:"transformationFunction"`
}

