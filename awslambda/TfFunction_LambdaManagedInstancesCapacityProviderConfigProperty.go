package awslambda


// Experimental.
type TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#capacity_provider_arn TfFunction#capacity_provider_arn}.
	// Experimental.
	CapacityProviderArn *string `field:"required" json:"capacityProviderArn" yaml:"capacityProviderArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#execution_environment_memory_gib_per_vcpu TfFunction#execution_environment_memory_gib_per_vcpu}.
	// Experimental.
	ExecutionEnvironmentMemoryGibPerVcpu *float64 `field:"optional" json:"executionEnvironmentMemoryGibPerVcpu" yaml:"executionEnvironmentMemoryGibPerVcpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#per_execution_environment_max_concurrency TfFunction#per_execution_environment_max_concurrency}.
	// Experimental.
	PerExecutionEnvironmentMaxConcurrency *float64 `field:"optional" json:"perExecutionEnvironmentMaxConcurrency" yaml:"perExecutionEnvironmentMaxConcurrency"`
}

