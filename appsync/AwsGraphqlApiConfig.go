package appsync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGraphqlApiConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#authentication_type AwsGraphqlApi#authentication_type}.
	// Experimental.
	AuthenticationType *string `field:"required" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#name AwsGraphqlApi#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// additional_authentication_provider block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#additional_authentication_provider AwsGraphqlApi#additional_authentication_provider}
	// Experimental.
	AdditionalAuthenticationProvider interface{} `field:"optional" json:"additionalAuthenticationProvider" yaml:"additionalAuthenticationProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#api_type AwsGraphqlApi#api_type}.
	// Experimental.
	ApiType *string `field:"optional" json:"apiType" yaml:"apiType"`
	// enhanced_metrics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#enhanced_metrics_config AwsGraphqlApi#enhanced_metrics_config}
	// Experimental.
	EnhancedMetricsConfig *AwsGraphqlApi_EnhancedMetricsConfigProperty `field:"optional" json:"enhancedMetricsConfig" yaml:"enhancedMetricsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#id AwsGraphqlApi#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#introspection_config AwsGraphqlApi#introspection_config}.
	// Experimental.
	IntrospectionConfig *string `field:"optional" json:"introspectionConfig" yaml:"introspectionConfig"`
	// lambda_authorizer_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#lambda_authorizer_config AwsGraphqlApi#lambda_authorizer_config}
	// Experimental.
	LambdaAuthorizerConfig *AwsGraphqlApi_LambdaAuthorizerConfigProperty `field:"optional" json:"lambdaAuthorizerConfig" yaml:"lambdaAuthorizerConfig"`
	// log_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#log_config AwsGraphqlApi#log_config}
	// Experimental.
	LogConfig *AwsGraphqlApi_LogConfigProperty `field:"optional" json:"logConfig" yaml:"logConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#merged_api_execution_role_arn AwsGraphqlApi#merged_api_execution_role_arn}.
	// Experimental.
	MergedApiExecutionRoleArn *string `field:"optional" json:"mergedApiExecutionRoleArn" yaml:"mergedApiExecutionRoleArn"`
	// openid_connect_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#openid_connect_config AwsGraphqlApi#openid_connect_config}
	// Experimental.
	OpenidConnectConfig *AwsGraphqlApi_OpenidConnectConfigProperty `field:"optional" json:"openidConnectConfig" yaml:"openidConnectConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#query_depth_limit AwsGraphqlApi#query_depth_limit}.
	// Experimental.
	QueryDepthLimit *float64 `field:"optional" json:"queryDepthLimit" yaml:"queryDepthLimit"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#region AwsGraphqlApi#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#resolver_count_limit AwsGraphqlApi#resolver_count_limit}.
	// Experimental.
	ResolverCountLimit *float64 `field:"optional" json:"resolverCountLimit" yaml:"resolverCountLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#schema AwsGraphqlApi#schema}.
	// Experimental.
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#tags AwsGraphqlApi#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#tags_all AwsGraphqlApi#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// user_pool_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#user_pool_config AwsGraphqlApi#user_pool_config}
	// Experimental.
	UserPoolConfig *AwsGraphqlApi_UserPoolConfigProperty `field:"optional" json:"userPoolConfig" yaml:"userPoolConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#visibility AwsGraphqlApi#visibility}.
	// Experimental.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_graphql_api#xray_enabled AwsGraphqlApi#xray_enabled}.
	// Experimental.
	XrayEnabled interface{} `field:"optional" json:"xrayEnabled" yaml:"xrayEnabled"`
}

