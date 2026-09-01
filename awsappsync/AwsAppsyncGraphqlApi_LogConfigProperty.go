package awsappsync


// Experimental.
type AwsAppsyncGraphqlApi_LogConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#cloudwatch_logs_role_arn AwsAppsyncGraphqlApi#cloudwatch_logs_role_arn}.
	// Experimental.
	CloudwatchLogsRoleArn *string `field:"required" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#field_log_level AwsAppsyncGraphqlApi#field_log_level}.
	// Experimental.
	FieldLogLevel *string `field:"required" json:"fieldLogLevel" yaml:"fieldLogLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#exclude_verbose_content AwsAppsyncGraphqlApi#exclude_verbose_content}.
	// Experimental.
	ExcludeVerboseContent interface{} `field:"optional" json:"excludeVerboseContent" yaml:"excludeVerboseContent"`
}

