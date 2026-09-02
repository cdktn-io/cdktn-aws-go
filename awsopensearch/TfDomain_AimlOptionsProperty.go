package awsopensearch


// Experimental.
type TfDomain_AimlOptionsProperty struct {
	// natural_language_query_generation_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#natural_language_query_generation_options TfDomain#natural_language_query_generation_options}
	// Experimental.
	NaturalLanguageQueryGenerationOptions *TfDomain_NaturalLanguageQueryGenerationOptionsProperty `field:"optional" json:"naturalLanguageQueryGenerationOptions" yaml:"naturalLanguageQueryGenerationOptions"`
	// s3_vectors_engine block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#s3_vectors_engine TfDomain#s3_vectors_engine}
	// Experimental.
	S3VectorsEngine *TfDomain_S3VectorsEngineProperty `field:"optional" json:"s3VectorsEngine" yaml:"s3VectorsEngine"`
	// serverless_vector_acceleration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/opensearch_domain#serverless_vector_acceleration TfDomain#serverless_vector_acceleration}
	// Experimental.
	ServerlessVectorAcceleration *TfDomain_ServerlessVectorAccelerationProperty `field:"optional" json:"serverlessVectorAcceleration" yaml:"serverlessVectorAcceleration"`
}

