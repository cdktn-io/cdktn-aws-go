package lambda


// Experimental.
type AwsFunction_CapacityProviderConfigProperty struct {
	// lambda_managed_instances_capacity_provider_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#lambda_managed_instances_capacity_provider_config AwsFunction#lambda_managed_instances_capacity_provider_config}
	// Experimental.
	LambdaManagedInstancesCapacityProviderConfig *AwsFunction_LambdaManagedInstancesCapacityProviderConfigProperty `field:"required" json:"lambdaManagedInstancesCapacityProviderConfig" yaml:"lambdaManagedInstancesCapacityProviderConfig"`
}

