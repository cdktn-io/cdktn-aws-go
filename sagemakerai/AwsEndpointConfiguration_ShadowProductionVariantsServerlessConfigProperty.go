package sagemakerai


// Experimental.
type AwsEndpointConfiguration_ShadowProductionVariantsServerlessConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#max_concurrency AwsEndpointConfiguration#max_concurrency}.
	// Experimental.
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#memory_size_in_mb AwsEndpointConfiguration#memory_size_in_mb}.
	// Experimental.
	MemorySizeInMb *float64 `field:"required" json:"memorySizeInMb" yaml:"memorySizeInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#provisioned_concurrency AwsEndpointConfiguration#provisioned_concurrency}.
	// Experimental.
	ProvisionedConcurrency *float64 `field:"optional" json:"provisionedConcurrency" yaml:"provisionedConcurrency"`
}

