package opensearch


// Experimental.
type AwsDomain_AimlOptionsProperty struct {
	// natural_language_query_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#natural_language_query_generation_options AwsDomain#natural_language_query_generation_options}
	// Experimental.
	NaturalLanguageQueryGenerationOptions *AwsDomain_NaturalLanguageQueryGenerationOptionsProperty `field:"optional" json:"naturalLanguageQueryGenerationOptions" yaml:"naturalLanguageQueryGenerationOptions"`
	// s3_vectors_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#s3_vectors_engine AwsDomain#s3_vectors_engine}
	// Experimental.
	S3VectorsEngine *AwsDomain_S3VectorsEngineProperty `field:"optional" json:"s3VectorsEngine" yaml:"s3VectorsEngine"`
	// serverless_vector_acceleration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#serverless_vector_acceleration AwsDomain#serverless_vector_acceleration}
	// Experimental.
	ServerlessVectorAcceleration *AwsDomain_ServerlessVectorAccelerationProperty `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

