package bedrock


// Experimental.
type AwsEvaluationJob_PrecomputedRagSourceConfigProperty struct {
	// retrieve_and_generate_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieve_and_generate_source_config AwsEvaluationJob#retrieve_and_generate_source_config}
	// Experimental.
	RetrieveAndGenerateSourceConfig interface{} `field:"optional" json:"retrieveAndGenerateSourceConfig" yaml:"retrieveAndGenerateSourceConfig"`
	// retrieve_source_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_evaluation_job#retrieve_source_config AwsEvaluationJob#retrieve_source_config}
	// Experimental.
	RetrieveSourceConfig interface{} `field:"optional" json:"retrieveSourceConfig" yaml:"retrieveSourceConfig"`
}

