package awsappsync


// Experimental.
type TfGraphqlApi_EnhancedMetricsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#data_source_level_metrics_behavior TfGraphqlApi#data_source_level_metrics_behavior}.
	// Experimental.
	DataSourceLevelMetricsBehavior *string `field:"required" json:"dataSourceLevelMetricsBehavior" yaml:"dataSourceLevelMetricsBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#operation_level_metrics_config TfGraphqlApi#operation_level_metrics_config}.
	// Experimental.
	OperationLevelMetricsConfig *string `field:"required" json:"operationLevelMetricsConfig" yaml:"operationLevelMetricsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#resolver_level_metrics_behavior TfGraphqlApi#resolver_level_metrics_behavior}.
	// Experimental.
	ResolverLevelMetricsBehavior *string `field:"required" json:"resolverLevelMetricsBehavior" yaml:"resolverLevelMetricsBehavior"`
}

