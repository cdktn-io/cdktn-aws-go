package appsync


// Experimental.
type AwsApi_LogConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#cloudwatch_logs_role_arn AwsApi#cloudwatch_logs_role_arn}.
	// Experimental.
	CloudwatchLogsRoleArn *string `field:"required" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_api#log_level AwsApi#log_level}.
	// Experimental.
	LogLevel *string `field:"required" json:"logLevel" yaml:"logLevel"`
}

