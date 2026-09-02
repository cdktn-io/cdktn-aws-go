package awslambda


// Experimental.
type TfFunction_CapacityProviderConfigProperty struct {
	// lambda_managed_instances_capacity_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#lambda_managed_instances_capacity_provider_config TfFunction#lambda_managed_instances_capacity_provider_config}
	// Experimental.
	LambdaManagedInstancesCapacityProviderConfig *TfFunction_LambdaManagedInstancesCapacityProviderConfigProperty `field:"required" json:"lambdaManagedInstancesCapacityProviderConfig" yaml:"lambdaManagedInstancesCapacityProviderConfig"`
}

