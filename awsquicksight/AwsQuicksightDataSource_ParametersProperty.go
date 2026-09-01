package awsquicksight


// Experimental.
type AwsQuicksightDataSource_ParametersProperty struct {
	// amazon_elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#amazon_elasticsearch AwsQuicksightDataSource#amazon_elasticsearch}
	// Experimental.
	AmazonElasticsearch *AwsQuicksightDataSource_AmazonElasticsearchProperty `field:"optional" json:"amazonElasticsearch" yaml:"amazonElasticsearch"`
	// athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#athena AwsQuicksightDataSource#athena}
	// Experimental.
	Athena *AwsQuicksightDataSource_AthenaProperty `field:"optional" json:"athena" yaml:"athena"`
	// aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora AwsQuicksightDataSource#aurora}
	// Experimental.
	Aurora *AwsQuicksightDataSource_AuroraProperty `field:"optional" json:"aurora" yaml:"aurora"`
	// aurora_postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora_postgresql AwsQuicksightDataSource#aurora_postgresql}
	// Experimental.
	AuroraPostgresql *AwsQuicksightDataSource_AuroraPostgresqlProperty `field:"optional" json:"auroraPostgresql" yaml:"auroraPostgresql"`
	// aws_iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aws_iot_analytics AwsQuicksightDataSource#aws_iot_analytics}
	// Experimental.
	AwsIotAnalytics *AwsQuicksightDataSource_AwsIotAnalyticsProperty `field:"optional" json:"awsIotAnalytics" yaml:"awsIotAnalytics"`
	// databricks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#databricks AwsQuicksightDataSource#databricks}
	// Experimental.
	Databricks *AwsQuicksightDataSource_DatabricksProperty `field:"optional" json:"databricks" yaml:"databricks"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#jira AwsQuicksightDataSource#jira}
	// Experimental.
	Jira *AwsQuicksightDataSource_JiraProperty `field:"optional" json:"jira" yaml:"jira"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#maria_db AwsQuicksightDataSource#maria_db}
	// Experimental.
	MariaDb *AwsQuicksightDataSource_MariaDbProperty `field:"optional" json:"mariaDb" yaml:"mariaDb"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#mysql AwsQuicksightDataSource#mysql}
	// Experimental.
	Mysql *AwsQuicksightDataSource_MysqlProperty `field:"optional" json:"mysql" yaml:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#oracle AwsQuicksightDataSource#oracle}
	// Experimental.
	Oracle *AwsQuicksightDataSource_OracleProperty `field:"optional" json:"oracle" yaml:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#postgresql AwsQuicksightDataSource#postgresql}
	// Experimental.
	Postgresql *AwsQuicksightDataSource_PostgresqlProperty `field:"optional" json:"postgresql" yaml:"postgresql"`
	// presto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#presto AwsQuicksightDataSource#presto}
	// Experimental.
	Presto *AwsQuicksightDataSource_PrestoProperty `field:"optional" json:"presto" yaml:"presto"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#rds AwsQuicksightDataSource#rds}
	// Experimental.
	Rds *AwsQuicksightDataSource_RdsProperty `field:"optional" json:"rds" yaml:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#redshift AwsQuicksightDataSource#redshift}
	// Experimental.
	Redshift *AwsQuicksightDataSource_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#s3 AwsQuicksightDataSource#s3}
	// Experimental.
	S3 *AwsQuicksightDataSource_S3Property `field:"optional" json:"s3" yaml:"s3"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#service_now AwsQuicksightDataSource#service_now}
	// Experimental.
	ServiceNow *AwsQuicksightDataSource_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#snowflake AwsQuicksightDataSource#snowflake}
	// Experimental.
	Snowflake *AwsQuicksightDataSource_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#spark AwsQuicksightDataSource#spark}
	// Experimental.
	Spark *AwsQuicksightDataSource_SparkProperty `field:"optional" json:"spark" yaml:"spark"`
	// sql_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#sql_server AwsQuicksightDataSource#sql_server}
	// Experimental.
	SqlServer *AwsQuicksightDataSource_SqlServerProperty `field:"optional" json:"sqlServer" yaml:"sqlServer"`
	// teradata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#teradata AwsQuicksightDataSource#teradata}
	// Experimental.
	Teradata *AwsQuicksightDataSource_TeradataProperty `field:"optional" json:"teradata" yaml:"teradata"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#twitter AwsQuicksightDataSource#twitter}
	// Experimental.
	Twitter *AwsQuicksightDataSource_TwitterProperty `field:"optional" json:"twitter" yaml:"twitter"`
}

