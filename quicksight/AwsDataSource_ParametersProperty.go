package quicksight


// Experimental.
type AwsDataSource_ParametersProperty struct {
	// amazon_elasticsearch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#amazon_elasticsearch AwsDataSource#amazon_elasticsearch}
	// Experimental.
	AmazonElasticsearch *AwsDataSource_AmazonElasticsearchProperty `field:"optional" json:"amazonElasticsearch" yaml:"amazonElasticsearch"`
	// athena block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#athena AwsDataSource#athena}
	// Experimental.
	Athena *AwsDataSource_AthenaProperty `field:"optional" json:"athena" yaml:"athena"`
	// aurora block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora AwsDataSource#aurora}
	// Experimental.
	Aurora *AwsDataSource_AuroraProperty `field:"optional" json:"aurora" yaml:"aurora"`
	// aurora_postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aurora_postgresql AwsDataSource#aurora_postgresql}
	// Experimental.
	AuroraPostgresql *AwsDataSource_AuroraPostgresqlProperty `field:"optional" json:"auroraPostgresql" yaml:"auroraPostgresql"`
	// aws_iot_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#aws_iot_analytics AwsDataSource#aws_iot_analytics}
	// Experimental.
	AwsIotAnalytics *AwsDataSource_AwsIotAnalyticsProperty `field:"optional" json:"awsIotAnalytics" yaml:"awsIotAnalytics"`
	// databricks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#databricks AwsDataSource#databricks}
	// Experimental.
	Databricks *AwsDataSource_DatabricksProperty `field:"optional" json:"databricks" yaml:"databricks"`
	// jira block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#jira AwsDataSource#jira}
	// Experimental.
	Jira *AwsDataSource_JiraProperty `field:"optional" json:"jira" yaml:"jira"`
	// maria_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#maria_db AwsDataSource#maria_db}
	// Experimental.
	MariaDb *AwsDataSource_MariaDbProperty `field:"optional" json:"mariaDb" yaml:"mariaDb"`
	// mysql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#mysql AwsDataSource#mysql}
	// Experimental.
	Mysql *AwsDataSource_MysqlProperty `field:"optional" json:"mysql" yaml:"mysql"`
	// oracle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#oracle AwsDataSource#oracle}
	// Experimental.
	Oracle *AwsDataSource_OracleProperty `field:"optional" json:"oracle" yaml:"oracle"`
	// postgresql block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#postgresql AwsDataSource#postgresql}
	// Experimental.
	Postgresql *AwsDataSource_PostgresqlProperty `field:"optional" json:"postgresql" yaml:"postgresql"`
	// presto block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#presto AwsDataSource#presto}
	// Experimental.
	Presto *AwsDataSource_PrestoProperty `field:"optional" json:"presto" yaml:"presto"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#rds AwsDataSource#rds}
	// Experimental.
	Rds *AwsDataSource_RdsProperty `field:"optional" json:"rds" yaml:"rds"`
	// redshift block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#redshift AwsDataSource#redshift}
	// Experimental.
	Redshift *AwsDataSource_RedshiftProperty `field:"optional" json:"redshift" yaml:"redshift"`
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#s3 AwsDataSource#s3}
	// Experimental.
	S3 *AwsDataSource_S3Property `field:"optional" json:"s3" yaml:"s3"`
	// service_now block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#service_now AwsDataSource#service_now}
	// Experimental.
	ServiceNow *AwsDataSource_ServiceNowProperty `field:"optional" json:"serviceNow" yaml:"serviceNow"`
	// snowflake block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#snowflake AwsDataSource#snowflake}
	// Experimental.
	Snowflake *AwsDataSource_SnowflakeProperty `field:"optional" json:"snowflake" yaml:"snowflake"`
	// spark block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#spark AwsDataSource#spark}
	// Experimental.
	Spark *AwsDataSource_SparkProperty `field:"optional" json:"spark" yaml:"spark"`
	// sql_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#sql_server AwsDataSource#sql_server}
	// Experimental.
	SqlServer *AwsDataSource_SqlServerProperty `field:"optional" json:"sqlServer" yaml:"sqlServer"`
	// teradata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#teradata AwsDataSource#teradata}
	// Experimental.
	Teradata *AwsDataSource_TeradataProperty `field:"optional" json:"teradata" yaml:"teradata"`
	// twitter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#twitter AwsDataSource#twitter}
	// Experimental.
	Twitter *AwsDataSource_TwitterProperty `field:"optional" json:"twitter" yaml:"twitter"`
}

