package awsappsync

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDatasourceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#api_id TfDatasource#api_id}.
	// Experimental.
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#name TfDatasource#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#type TfDatasource#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#description TfDatasource#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#dynamodb_config TfDatasource#dynamodb_config}
	// Experimental.
	DynamodbConfig *TfDatasource_DynamodbConfigProperty `field:"optional" json:"dynamodbConfig" yaml:"dynamodbConfig"`
	// elasticsearch_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#elasticsearch_config TfDatasource#elasticsearch_config}
	// Experimental.
	ElasticsearchConfig *TfDatasource_ElasticsearchConfigProperty `field:"optional" json:"elasticsearchConfig" yaml:"elasticsearchConfig"`
	// event_bridge_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#event_bridge_config TfDatasource#event_bridge_config}
	// Experimental.
	EventBridgeConfig *TfDatasource_EventBridgeConfigProperty `field:"optional" json:"eventBridgeConfig" yaml:"eventBridgeConfig"`
	// http_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#http_config TfDatasource#http_config}
	// Experimental.
	HttpConfig *TfDatasource_HttpConfigProperty `field:"optional" json:"httpConfig" yaml:"httpConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#id TfDatasource#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#lambda_config TfDatasource#lambda_config}
	// Experimental.
	LambdaConfig *TfDatasource_LambdaConfigProperty `field:"optional" json:"lambdaConfig" yaml:"lambdaConfig"`
	// opensearchservice_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#opensearchservice_config TfDatasource#opensearchservice_config}
	// Experimental.
	OpensearchserviceConfig *TfDatasource_OpensearchserviceConfigProperty `field:"optional" json:"opensearchserviceConfig" yaml:"opensearchserviceConfig"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region TfDatasource#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// relational_database_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#relational_database_config TfDatasource#relational_database_config}
	// Experimental.
	RelationalDatabaseConfig *TfDatasource_RelationalDatabaseConfigProperty `field:"optional" json:"relationalDatabaseConfig" yaml:"relationalDatabaseConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#service_role_arn TfDatasource#service_role_arn}.
	// Experimental.
	ServiceRoleArn *string `field:"optional" json:"serviceRoleArn" yaml:"serviceRoleArn"`
}

