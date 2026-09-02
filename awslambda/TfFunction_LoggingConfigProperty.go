package awslambda


// Experimental.
type TfFunction_LoggingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#log_format TfFunction#log_format}.
	// Experimental.
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#application_log_level TfFunction#application_log_level}.
	// Experimental.
	ApplicationLogLevel *string `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#log_group TfFunction#log_group}.
	// Experimental.
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#system_log_level TfFunction#system_log_level}.
	// Experimental.
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

