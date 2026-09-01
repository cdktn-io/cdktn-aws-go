package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightDataSource_ParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmazonElasticsearch() AwsQuicksightDataSource_AmazonElasticsearchPropertyOutputReference
	// Experimental.
	AmazonElasticsearchInput() *AwsQuicksightDataSource_AmazonElasticsearchProperty
	// Experimental.
	Athena() AwsQuicksightDataSource_AthenaPropertyOutputReference
	// Experimental.
	AthenaInput() *AwsQuicksightDataSource_AthenaProperty
	// Experimental.
	Aurora() AwsQuicksightDataSource_AuroraPropertyOutputReference
	// Experimental.
	AuroraInput() *AwsQuicksightDataSource_AuroraProperty
	// Experimental.
	AuroraPostgresql() AwsQuicksightDataSource_AuroraPostgresqlPropertyOutputReference
	// Experimental.
	AuroraPostgresqlInput() *AwsQuicksightDataSource_AuroraPostgresqlProperty
	// Experimental.
	AwsIotAnalytics() AwsQuicksightDataSource_AwsIotAnalyticsPropertyOutputReference
	// Experimental.
	AwsIotAnalyticsInput() *AwsQuicksightDataSource_AwsIotAnalyticsProperty
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Databricks() AwsQuicksightDataSource_DatabricksPropertyOutputReference
	// Experimental.
	DatabricksInput() *AwsQuicksightDataSource_DatabricksProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsQuicksightDataSource_ParametersProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightDataSource_ParametersProperty)
	// Experimental.
	Jira() AwsQuicksightDataSource_JiraPropertyOutputReference
	// Experimental.
	JiraInput() *AwsQuicksightDataSource_JiraProperty
	// Experimental.
	MariaDb() AwsQuicksightDataSource_MariaDbPropertyOutputReference
	// Experimental.
	MariaDbInput() *AwsQuicksightDataSource_MariaDbProperty
	// Experimental.
	Mysql() AwsQuicksightDataSource_MysqlPropertyOutputReference
	// Experimental.
	MysqlInput() *AwsQuicksightDataSource_MysqlProperty
	// Experimental.
	Oracle() AwsQuicksightDataSource_OraclePropertyOutputReference
	// Experimental.
	OracleInput() *AwsQuicksightDataSource_OracleProperty
	// Experimental.
	Postgresql() AwsQuicksightDataSource_PostgresqlPropertyOutputReference
	// Experimental.
	PostgresqlInput() *AwsQuicksightDataSource_PostgresqlProperty
	// Experimental.
	Presto() AwsQuicksightDataSource_PrestoPropertyOutputReference
	// Experimental.
	PrestoInput() *AwsQuicksightDataSource_PrestoProperty
	// Experimental.
	Rds() AwsQuicksightDataSource_RdsPropertyOutputReference
	// Experimental.
	RdsInput() *AwsQuicksightDataSource_RdsProperty
	// Experimental.
	Redshift() AwsQuicksightDataSource_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsQuicksightDataSource_RedshiftProperty
	// Experimental.
	S3() AwsQuicksightDataSource_S3PropertyOutputReference
	// Experimental.
	S3Input() *AwsQuicksightDataSource_S3Property
	// Experimental.
	ServiceNow() AwsQuicksightDataSource_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsQuicksightDataSource_ServiceNowProperty
	// Experimental.
	Snowflake() AwsQuicksightDataSource_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsQuicksightDataSource_SnowflakeProperty
	// Experimental.
	Spark() AwsQuicksightDataSource_SparkPropertyOutputReference
	// Experimental.
	SparkInput() *AwsQuicksightDataSource_SparkProperty
	// Experimental.
	SqlServer() AwsQuicksightDataSource_SqlServerPropertyOutputReference
	// Experimental.
	SqlServerInput() *AwsQuicksightDataSource_SqlServerProperty
	// Experimental.
	Teradata() AwsQuicksightDataSource_TeradataPropertyOutputReference
	// Experimental.
	TeradataInput() *AwsQuicksightDataSource_TeradataProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Twitter() AwsQuicksightDataSource_TwitterPropertyOutputReference
	// Experimental.
	TwitterInput() *AwsQuicksightDataSource_TwitterProperty
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAmazonElasticsearch(value *AwsQuicksightDataSource_AmazonElasticsearchProperty)
	// Experimental.
	PutAthena(value *AwsQuicksightDataSource_AthenaProperty)
	// Experimental.
	PutAurora(value *AwsQuicksightDataSource_AuroraProperty)
	// Experimental.
	PutAuroraPostgresql(value *AwsQuicksightDataSource_AuroraPostgresqlProperty)
	// Experimental.
	PutAwsIotAnalytics(value *AwsQuicksightDataSource_AwsIotAnalyticsProperty)
	// Experimental.
	PutDatabricks(value *AwsQuicksightDataSource_DatabricksProperty)
	// Experimental.
	PutJira(value *AwsQuicksightDataSource_JiraProperty)
	// Experimental.
	PutMariaDb(value *AwsQuicksightDataSource_MariaDbProperty)
	// Experimental.
	PutMysql(value *AwsQuicksightDataSource_MysqlProperty)
	// Experimental.
	PutOracle(value *AwsQuicksightDataSource_OracleProperty)
	// Experimental.
	PutPostgresql(value *AwsQuicksightDataSource_PostgresqlProperty)
	// Experimental.
	PutPresto(value *AwsQuicksightDataSource_PrestoProperty)
	// Experimental.
	PutRds(value *AwsQuicksightDataSource_RdsProperty)
	// Experimental.
	PutRedshift(value *AwsQuicksightDataSource_RedshiftProperty)
	// Experimental.
	PutS3(value *AwsQuicksightDataSource_S3Property)
	// Experimental.
	PutServiceNow(value *AwsQuicksightDataSource_ServiceNowProperty)
	// Experimental.
	PutSnowflake(value *AwsQuicksightDataSource_SnowflakeProperty)
	// Experimental.
	PutSpark(value *AwsQuicksightDataSource_SparkProperty)
	// Experimental.
	PutSqlServer(value *AwsQuicksightDataSource_SqlServerProperty)
	// Experimental.
	PutTeradata(value *AwsQuicksightDataSource_TeradataProperty)
	// Experimental.
	PutTwitter(value *AwsQuicksightDataSource_TwitterProperty)
	// Experimental.
	ResetAmazonElasticsearch()
	// Experimental.
	ResetAthena()
	// Experimental.
	ResetAurora()
	// Experimental.
	ResetAuroraPostgresql()
	// Experimental.
	ResetAwsIotAnalytics()
	// Experimental.
	ResetDatabricks()
	// Experimental.
	ResetJira()
	// Experimental.
	ResetMariaDb()
	// Experimental.
	ResetMysql()
	// Experimental.
	ResetOracle()
	// Experimental.
	ResetPostgresql()
	// Experimental.
	ResetPresto()
	// Experimental.
	ResetRds()
	// Experimental.
	ResetRedshift()
	// Experimental.
	ResetS3()
	// Experimental.
	ResetServiceNow()
	// Experimental.
	ResetSnowflake()
	// Experimental.
	ResetSpark()
	// Experimental.
	ResetSqlServer()
	// Experimental.
	ResetTeradata()
	// Experimental.
	ResetTwitter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsQuicksightDataSource_ParametersPropertyOutputReference
type jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AmazonElasticsearch() AwsQuicksightDataSource_AmazonElasticsearchPropertyOutputReference {
	var returns AwsQuicksightDataSource_AmazonElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AmazonElasticsearchInput() *AwsQuicksightDataSource_AmazonElasticsearchProperty {
	var returns *AwsQuicksightDataSource_AmazonElasticsearchProperty
	_jsii_.Get(
		j,
		"amazonElasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Athena() AwsQuicksightDataSource_AthenaPropertyOutputReference {
	var returns AwsQuicksightDataSource_AthenaPropertyOutputReference
	_jsii_.Get(
		j,
		"athena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AthenaInput() *AwsQuicksightDataSource_AthenaProperty {
	var returns *AwsQuicksightDataSource_AthenaProperty
	_jsii_.Get(
		j,
		"athenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Aurora() AwsQuicksightDataSource_AuroraPropertyOutputReference {
	var returns AwsQuicksightDataSource_AuroraPropertyOutputReference
	_jsii_.Get(
		j,
		"aurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AuroraInput() *AwsQuicksightDataSource_AuroraProperty {
	var returns *AwsQuicksightDataSource_AuroraProperty
	_jsii_.Get(
		j,
		"auroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AuroraPostgresql() AwsQuicksightDataSource_AuroraPostgresqlPropertyOutputReference {
	var returns AwsQuicksightDataSource_AuroraPostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"auroraPostgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AuroraPostgresqlInput() *AwsQuicksightDataSource_AuroraPostgresqlProperty {
	var returns *AwsQuicksightDataSource_AuroraPostgresqlProperty
	_jsii_.Get(
		j,
		"auroraPostgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AwsIotAnalytics() AwsQuicksightDataSource_AwsIotAnalyticsPropertyOutputReference {
	var returns AwsQuicksightDataSource_AwsIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"awsIotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) AwsIotAnalyticsInput() *AwsQuicksightDataSource_AwsIotAnalyticsProperty {
	var returns *AwsQuicksightDataSource_AwsIotAnalyticsProperty
	_jsii_.Get(
		j,
		"awsIotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Databricks() AwsQuicksightDataSource_DatabricksPropertyOutputReference {
	var returns AwsQuicksightDataSource_DatabricksPropertyOutputReference
	_jsii_.Get(
		j,
		"databricks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) DatabricksInput() *AwsQuicksightDataSource_DatabricksProperty {
	var returns *AwsQuicksightDataSource_DatabricksProperty
	_jsii_.Get(
		j,
		"databricksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) InternalValue() *AwsQuicksightDataSource_ParametersProperty {
	var returns *AwsQuicksightDataSource_ParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Jira() AwsQuicksightDataSource_JiraPropertyOutputReference {
	var returns AwsQuicksightDataSource_JiraPropertyOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) JiraInput() *AwsQuicksightDataSource_JiraProperty {
	var returns *AwsQuicksightDataSource_JiraProperty
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) MariaDb() AwsQuicksightDataSource_MariaDbPropertyOutputReference {
	var returns AwsQuicksightDataSource_MariaDbPropertyOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) MariaDbInput() *AwsQuicksightDataSource_MariaDbProperty {
	var returns *AwsQuicksightDataSource_MariaDbProperty
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Mysql() AwsQuicksightDataSource_MysqlPropertyOutputReference {
	var returns AwsQuicksightDataSource_MysqlPropertyOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) MysqlInput() *AwsQuicksightDataSource_MysqlProperty {
	var returns *AwsQuicksightDataSource_MysqlProperty
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Oracle() AwsQuicksightDataSource_OraclePropertyOutputReference {
	var returns AwsQuicksightDataSource_OraclePropertyOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) OracleInput() *AwsQuicksightDataSource_OracleProperty {
	var returns *AwsQuicksightDataSource_OracleProperty
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Postgresql() AwsQuicksightDataSource_PostgresqlPropertyOutputReference {
	var returns AwsQuicksightDataSource_PostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PostgresqlInput() *AwsQuicksightDataSource_PostgresqlProperty {
	var returns *AwsQuicksightDataSource_PostgresqlProperty
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Presto() AwsQuicksightDataSource_PrestoPropertyOutputReference {
	var returns AwsQuicksightDataSource_PrestoPropertyOutputReference
	_jsii_.Get(
		j,
		"presto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PrestoInput() *AwsQuicksightDataSource_PrestoProperty {
	var returns *AwsQuicksightDataSource_PrestoProperty
	_jsii_.Get(
		j,
		"prestoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Rds() AwsQuicksightDataSource_RdsPropertyOutputReference {
	var returns AwsQuicksightDataSource_RdsPropertyOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) RdsInput() *AwsQuicksightDataSource_RdsProperty {
	var returns *AwsQuicksightDataSource_RdsProperty
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Redshift() AwsQuicksightDataSource_RedshiftPropertyOutputReference {
	var returns AwsQuicksightDataSource_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) RedshiftInput() *AwsQuicksightDataSource_RedshiftProperty {
	var returns *AwsQuicksightDataSource_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) S3() AwsQuicksightDataSource_S3PropertyOutputReference {
	var returns AwsQuicksightDataSource_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) S3Input() *AwsQuicksightDataSource_S3Property {
	var returns *AwsQuicksightDataSource_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ServiceNow() AwsQuicksightDataSource_ServiceNowPropertyOutputReference {
	var returns AwsQuicksightDataSource_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ServiceNowInput() *AwsQuicksightDataSource_ServiceNowProperty {
	var returns *AwsQuicksightDataSource_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Snowflake() AwsQuicksightDataSource_SnowflakePropertyOutputReference {
	var returns AwsQuicksightDataSource_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) SnowflakeInput() *AwsQuicksightDataSource_SnowflakeProperty {
	var returns *AwsQuicksightDataSource_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Spark() AwsQuicksightDataSource_SparkPropertyOutputReference {
	var returns AwsQuicksightDataSource_SparkPropertyOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) SparkInput() *AwsQuicksightDataSource_SparkProperty {
	var returns *AwsQuicksightDataSource_SparkProperty
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) SqlServer() AwsQuicksightDataSource_SqlServerPropertyOutputReference {
	var returns AwsQuicksightDataSource_SqlServerPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) SqlServerInput() *AwsQuicksightDataSource_SqlServerProperty {
	var returns *AwsQuicksightDataSource_SqlServerProperty
	_jsii_.Get(
		j,
		"sqlServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Teradata() AwsQuicksightDataSource_TeradataPropertyOutputReference {
	var returns AwsQuicksightDataSource_TeradataPropertyOutputReference
	_jsii_.Get(
		j,
		"teradata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) TeradataInput() *AwsQuicksightDataSource_TeradataProperty {
	var returns *AwsQuicksightDataSource_TeradataProperty
	_jsii_.Get(
		j,
		"teradataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Twitter() AwsQuicksightDataSource_TwitterPropertyOutputReference {
	var returns AwsQuicksightDataSource_TwitterPropertyOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) TwitterInput() *AwsQuicksightDataSource_TwitterProperty {
	var returns *AwsQuicksightDataSource_TwitterProperty
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightDataSource_ParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightDataSource_ParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightDataSource_ParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightDataSource_ParametersPropertyOutputReference_Override(a AwsQuicksightDataSource_ParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference)SetInternalValue(val *AwsQuicksightDataSource_ParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutAmazonElasticsearch(value *AwsQuicksightDataSource_AmazonElasticsearchProperty) {
	if err := a.validatePutAmazonElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonElasticsearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutAthena(value *AwsQuicksightDataSource_AthenaProperty) {
	if err := a.validatePutAthenaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAthena",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutAurora(value *AwsQuicksightDataSource_AuroraProperty) {
	if err := a.validatePutAuroraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAurora",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutAuroraPostgresql(value *AwsQuicksightDataSource_AuroraPostgresqlProperty) {
	if err := a.validatePutAuroraPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraPostgresql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutAwsIotAnalytics(value *AwsQuicksightDataSource_AwsIotAnalyticsProperty) {
	if err := a.validatePutAwsIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsIotAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutDatabricks(value *AwsQuicksightDataSource_DatabricksProperty) {
	if err := a.validatePutDatabricksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabricks",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutJira(value *AwsQuicksightDataSource_JiraProperty) {
	if err := a.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJira",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutMariaDb(value *AwsQuicksightDataSource_MariaDbProperty) {
	if err := a.validatePutMariaDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutMysql(value *AwsQuicksightDataSource_MysqlProperty) {
	if err := a.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMysql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutOracle(value *AwsQuicksightDataSource_OracleProperty) {
	if err := a.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOracle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutPostgresql(value *AwsQuicksightDataSource_PostgresqlProperty) {
	if err := a.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutPresto(value *AwsQuicksightDataSource_PrestoProperty) {
	if err := a.validatePutPrestoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPresto",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutRds(value *AwsQuicksightDataSource_RdsProperty) {
	if err := a.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutRedshift(value *AwsQuicksightDataSource_RedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutS3(value *AwsQuicksightDataSource_S3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutServiceNow(value *AwsQuicksightDataSource_ServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutSnowflake(value *AwsQuicksightDataSource_SnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutSpark(value *AwsQuicksightDataSource_SparkProperty) {
	if err := a.validatePutSparkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpark",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutSqlServer(value *AwsQuicksightDataSource_SqlServerProperty) {
	if err := a.validatePutSqlServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutTeradata(value *AwsQuicksightDataSource_TeradataProperty) {
	if err := a.validatePutTeradataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTeradata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) PutTwitter(value *AwsQuicksightDataSource_TwitterProperty) {
	if err := a.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTwitter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetAmazonElasticsearch() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonElasticsearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetAthena() {
	_jsii_.InvokeVoid(
		a,
		"resetAthena",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetAurora() {
	_jsii_.InvokeVoid(
		a,
		"resetAurora",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetAuroraPostgresql() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraPostgresql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetAwsIotAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsIotAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetDatabricks() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabricks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		a,
		"resetJira",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetMariaDb() {
	_jsii_.InvokeVoid(
		a,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetMysql() {
	_jsii_.InvokeVoid(
		a,
		"resetMysql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetOracle() {
	_jsii_.InvokeVoid(
		a,
		"resetOracle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetPostgresql() {
	_jsii_.InvokeVoid(
		a,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetPresto() {
	_jsii_.InvokeVoid(
		a,
		"resetPresto",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetRds() {
	_jsii_.InvokeVoid(
		a,
		"resetRds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetSpark() {
	_jsii_.InvokeVoid(
		a,
		"resetSpark",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetSqlServer() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetTeradata() {
	_jsii_.InvokeVoid(
		a,
		"resetTeradata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		a,
		"resetTwitter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSource_ParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

