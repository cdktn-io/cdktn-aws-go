package lambda


// Experimental.
type AwsFunctionScalingConfig_FunctionScalingConfigProperty struct {
	// Maximum number of execution environments that can be provisioned for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#max_execution_environments AwsFunctionScalingConfig#max_execution_environments}
	// Experimental.
	MaxExecutionEnvironments *float64 `field:"optional" json:"maxExecutionEnvironments" yaml:"maxExecutionEnvironments"`
	// Minimum number of execution environments to maintain for the function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function_scaling_config#min_execution_environments AwsFunctionScalingConfig#min_execution_environments}
	// Experimental.
	MinExecutionEnvironments *float64 `field:"optional" json:"minExecutionEnvironments" yaml:"minExecutionEnvironments"`
}

