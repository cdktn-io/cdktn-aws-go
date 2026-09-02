package awsquicksight


// Experimental.
type TfDataSource_ParametersProperty struct {
	// amazon_elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#amazon_elasticsearch TfDataSource#amazon_elasticsearch}
	// Experimental.
	AmazonElasticsearch *TfDataSource_AmazonElasticsearchProperty `field:"optional" json:"amazonElasticsearch" yaml:"amazonElasticsearch"`
	// athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#athena TfDataSource#athena}
	// Experimental.
	Athena *TfDataSource_AthenaProperty `field:"optional" json:"athena" yaml:"athena"`
	// aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora TfDataSource#aurora}
	// Experimental.
	Aurora *TfDataSource_AuroraProperty `field:"optional" json:"aurora" yaml:"aurora"`
	// aurora_postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora_postgresql TfDataSource#aurora_postgresql}
	// Experimental.
	AuroraPostgresql *TfDataSource_AuroraPostgresqlProperty `field:"optional" json:"auroraPostgresql" yaml:"auroraPostgresql"`
	// aws_iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aws_iot_analytics TfDataSource#aws_iot_analytics}
	// Experimental.
	AwsIotAnalytics *TfDataSource_AwsIotAnalyticsProperty `field:"optional" json:"awsIotAnalytics" yaml:"awsIotAnalytics"`
	// databricks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#databricks TfDataSource#databricks}
	// Experimental.
	Databricks *TfDataSource_DatabricksProperty `field:"optional" json:"databricks" yaml:"databricks"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#jira TfDataSource#jira}
	// Experimental.
	Jira *TfDataSource_JiraProperty `field:"optional" json:"jira" yaml:"jira"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#maria_db TfDataSource#maria_db}
	// Experimental.
	MariaDb *TfDataSource_MariaDbProperty `field:"optional" json:"mariaDb" yaml:"mariaDb"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#mysql TfDataSource#mysql}
	// Experimental.
	Mysql *TfDataSource_MysqlProperty `field:"optional" json:"mysql" yaml:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#oracle TfDataSource#oracle}
	// Experimental.
	Oracle *TfDataSource_OracleProperty `field:"optional" json:"oracle" yaml:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#postgresql TfDataSource#postgresql}
	// Experimental.
	Postgresql *TfDataSource_PostgresqlProperty `field:"optional" json:"postgresql" yaml:"postgresql"`
	// presto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#presto TfDataSource#presto}
	// Experimental.
	Presto *TfDataSource_PrestoProperty `field:"optional" json:"presto" yaml:"presto"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#rds TfDataSource#rds}
	// Experimental.
	Rds *TfDataSource_RdsProperty `field:"optional" json:"rds" yaml:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#redshift TfDataSource#redshift}
	// Experimental.
	Redshift *TfDataSource_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#s3 TfDataSource#s3}
	// Experimental.
	S3 *TfDataSource_S3Property `field:"optional" json:"s3" yaml:"s3"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#service_now TfDataSource#service_now}
	// Experimental.
	ServiceNow *TfDataSource_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#snowflake TfDataSource#snowflake}
	// Experimental.
	Snowflake *TfDataSource_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#spark TfDataSource#spark}
	// Experimental.
	Spark *TfDataSource_SparkProperty `field:"optional" json:"spark" yaml:"spark"`
	// sql_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#sql_server TfDataSource#sql_server}
	// Experimental.
	SqlServer *TfDataSource_SqlServerProperty `field:"optional" json:"sqlServer" yaml:"sqlServer"`
	// teradata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#teradata TfDataSource#teradata}
	// Experimental.
	Teradata *TfDataSource_TeradataProperty `field:"optional" json:"teradata" yaml:"teradata"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#twitter TfDataSource#twitter}
	// Experimental.
	Twitter *TfDataSource_TwitterProperty `field:"optional" json:"twitter" yaml:"twitter"`
}

