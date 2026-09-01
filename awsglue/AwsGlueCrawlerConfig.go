package awsglue

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGlueCrawlerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#database_name AwsGlueCrawler#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#name AwsGlueCrawler#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#role AwsGlueCrawler#role}.
	// Experimental.
	Role *string `field:"required" json:"role" yaml:"role"`
	// catalog_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#catalog_target AwsGlueCrawler#catalog_target}
	// Experimental.
	CatalogTarget interface{} `field:"optional" json:"catalogTarget" yaml:"catalogTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#classifiers AwsGlueCrawler#classifiers}.
	// Experimental.
	Classifiers *[]*string `field:"optional" json:"classifiers" yaml:"classifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#configuration AwsGlueCrawler#configuration}.
	// Experimental.
	Configuration *string `field:"optional" json:"configuration" yaml:"configuration"`
	// delta_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#delta_target AwsGlueCrawler#delta_target}
	// Experimental.
	DeltaTarget interface{} `field:"optional" json:"deltaTarget" yaml:"deltaTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#description AwsGlueCrawler#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// dynamodb_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#dynamodb_target AwsGlueCrawler#dynamodb_target}
	// Experimental.
	DynamodbTarget interface{} `field:"optional" json:"dynamodbTarget" yaml:"dynamodbTarget"`
	// hudi_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#hudi_target AwsGlueCrawler#hudi_target}
	// Experimental.
	HudiTarget interface{} `field:"optional" json:"hudiTarget" yaml:"hudiTarget"`
	// iceberg_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#iceberg_target AwsGlueCrawler#iceberg_target}
	// Experimental.
	IcebergTarget interface{} `field:"optional" json:"icebergTarget" yaml:"icebergTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#id AwsGlueCrawler#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// jdbc_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#jdbc_target AwsGlueCrawler#jdbc_target}
	// Experimental.
	JdbcTarget interface{} `field:"optional" json:"jdbcTarget" yaml:"jdbcTarget"`
	// lake_formation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#lake_formation_configuration AwsGlueCrawler#lake_formation_configuration}
	// Experimental.
	LakeFormationConfiguration *AwsGlueCrawler_LakeFormationConfigurationProperty `field:"optional" json:"lakeFormationConfiguration" yaml:"lakeFormationConfiguration"`
	// lineage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#lineage_configuration AwsGlueCrawler#lineage_configuration}
	// Experimental.
	LineageConfiguration *AwsGlueCrawler_LineageConfigurationProperty `field:"optional" json:"lineageConfiguration" yaml:"lineageConfiguration"`
	// mongodb_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#mongodb_target AwsGlueCrawler#mongodb_target}
	// Experimental.
	MongodbTarget interface{} `field:"optional" json:"mongodbTarget" yaml:"mongodbTarget"`
	// recrawl_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#recrawl_policy AwsGlueCrawler#recrawl_policy}
	// Experimental.
	RecrawlPolicy *AwsGlueCrawler_RecrawlPolicyProperty `field:"optional" json:"recrawlPolicy" yaml:"recrawlPolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#region AwsGlueCrawler#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#s3_target AwsGlueCrawler#s3_target}
	// Experimental.
	S3Target interface{} `field:"optional" json:"s3Target" yaml:"s3Target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#schedule AwsGlueCrawler#schedule}.
	// Experimental.
	Schedule *string `field:"optional" json:"schedule" yaml:"schedule"`
	// schema_change_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#schema_change_policy AwsGlueCrawler#schema_change_policy}
	// Experimental.
	SchemaChangePolicy *AwsGlueCrawler_SchemaChangePolicyProperty `field:"optional" json:"schemaChangePolicy" yaml:"schemaChangePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#security_configuration AwsGlueCrawler#security_configuration}.
	// Experimental.
	SecurityConfiguration *string `field:"optional" json:"securityConfiguration" yaml:"securityConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#table_prefix AwsGlueCrawler#table_prefix}.
	// Experimental.
	TablePrefix *string `field:"optional" json:"tablePrefix" yaml:"tablePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#tags AwsGlueCrawler#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#tags_all AwsGlueCrawler#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

