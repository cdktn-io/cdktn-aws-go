package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/quicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/quicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSource_ParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmazonElasticsearch() AwsDataSource_AmazonElasticsearchPropertyOutputReference
	// Experimental.
	AmazonElasticsearchInput() *AwsDataSource_AmazonElasticsearchProperty
	// Experimental.
	Athena() AwsDataSource_AthenaPropertyOutputReference
	// Experimental.
	AthenaInput() *AwsDataSource_AthenaProperty
	// Experimental.
	Aurora() AwsDataSource_AuroraPropertyOutputReference
	// Experimental.
	AuroraInput() *AwsDataSource_AuroraProperty
	// Experimental.
	AuroraPostgresql() AwsDataSource_AuroraPostgresqlPropertyOutputReference
	// Experimental.
	AuroraPostgresqlInput() *AwsDataSource_AuroraPostgresqlProperty
	// Experimental.
	AwsIotAnalytics() AwsDataSource_AwsIotAnalyticsPropertyOutputReference
	// Experimental.
	AwsIotAnalyticsInput() *AwsDataSource_AwsIotAnalyticsProperty
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
	Databricks() AwsDataSource_DatabricksPropertyOutputReference
	// Experimental.
	DatabricksInput() *AwsDataSource_DatabricksProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDataSource_ParametersProperty
	// Experimental.
	SetInternalValue(val *AwsDataSource_ParametersProperty)
	// Experimental.
	Jira() AwsDataSource_JiraPropertyOutputReference
	// Experimental.
	JiraInput() *AwsDataSource_JiraProperty
	// Experimental.
	MariaDb() AwsDataSource_MariaDbPropertyOutputReference
	// Experimental.
	MariaDbInput() *AwsDataSource_MariaDbProperty
	// Experimental.
	Mysql() AwsDataSource_MysqlPropertyOutputReference
	// Experimental.
	MysqlInput() *AwsDataSource_MysqlProperty
	// Experimental.
	Oracle() AwsDataSource_OraclePropertyOutputReference
	// Experimental.
	OracleInput() *AwsDataSource_OracleProperty
	// Experimental.
	Postgresql() AwsDataSource_PostgresqlPropertyOutputReference
	// Experimental.
	PostgresqlInput() *AwsDataSource_PostgresqlProperty
	// Experimental.
	Presto() AwsDataSource_PrestoPropertyOutputReference
	// Experimental.
	PrestoInput() *AwsDataSource_PrestoProperty
	// Experimental.
	Rds() AwsDataSource_RdsPropertyOutputReference
	// Experimental.
	RdsInput() *AwsDataSource_RdsProperty
	// Experimental.
	Redshift() AwsDataSource_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *AwsDataSource_RedshiftProperty
	// Experimental.
	S3() AwsDataSource_S3PropertyOutputReference
	// Experimental.
	S3Input() *AwsDataSource_S3Property
	// Experimental.
	ServiceNow() AwsDataSource_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *AwsDataSource_ServiceNowProperty
	// Experimental.
	Snowflake() AwsDataSource_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *AwsDataSource_SnowflakeProperty
	// Experimental.
	Spark() AwsDataSource_SparkPropertyOutputReference
	// Experimental.
	SparkInput() *AwsDataSource_SparkProperty
	// Experimental.
	SqlServer() AwsDataSource_SqlServerPropertyOutputReference
	// Experimental.
	SqlServerInput() *AwsDataSource_SqlServerProperty
	// Experimental.
	Teradata() AwsDataSource_TeradataPropertyOutputReference
	// Experimental.
	TeradataInput() *AwsDataSource_TeradataProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Twitter() AwsDataSource_TwitterPropertyOutputReference
	// Experimental.
	TwitterInput() *AwsDataSource_TwitterProperty
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
	PutAmazonElasticsearch(value *AwsDataSource_AmazonElasticsearchProperty)
	// Experimental.
	PutAthena(value *AwsDataSource_AthenaProperty)
	// Experimental.
	PutAurora(value *AwsDataSource_AuroraProperty)
	// Experimental.
	PutAuroraPostgresql(value *AwsDataSource_AuroraPostgresqlProperty)
	// Experimental.
	PutAwsIotAnalytics(value *AwsDataSource_AwsIotAnalyticsProperty)
	// Experimental.
	PutDatabricks(value *AwsDataSource_DatabricksProperty)
	// Experimental.
	PutJira(value *AwsDataSource_JiraProperty)
	// Experimental.
	PutMariaDb(value *AwsDataSource_MariaDbProperty)
	// Experimental.
	PutMysql(value *AwsDataSource_MysqlProperty)
	// Experimental.
	PutOracle(value *AwsDataSource_OracleProperty)
	// Experimental.
	PutPostgresql(value *AwsDataSource_PostgresqlProperty)
	// Experimental.
	PutPresto(value *AwsDataSource_PrestoProperty)
	// Experimental.
	PutRds(value *AwsDataSource_RdsProperty)
	// Experimental.
	PutRedshift(value *AwsDataSource_RedshiftProperty)
	// Experimental.
	PutS3(value *AwsDataSource_S3Property)
	// Experimental.
	PutServiceNow(value *AwsDataSource_ServiceNowProperty)
	// Experimental.
	PutSnowflake(value *AwsDataSource_SnowflakeProperty)
	// Experimental.
	PutSpark(value *AwsDataSource_SparkProperty)
	// Experimental.
	PutSqlServer(value *AwsDataSource_SqlServerProperty)
	// Experimental.
	PutTeradata(value *AwsDataSource_TeradataProperty)
	// Experimental.
	PutTwitter(value *AwsDataSource_TwitterProperty)
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

// The jsii proxy struct for AwsDataSource_ParametersPropertyOutputReference
type jsiiProxy_AwsDataSource_ParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AmazonElasticsearch() AwsDataSource_AmazonElasticsearchPropertyOutputReference {
	var returns AwsDataSource_AmazonElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AmazonElasticsearchInput() *AwsDataSource_AmazonElasticsearchProperty {
	var returns *AwsDataSource_AmazonElasticsearchProperty
	_jsii_.Get(
		j,
		"amazonElasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Athena() AwsDataSource_AthenaPropertyOutputReference {
	var returns AwsDataSource_AthenaPropertyOutputReference
	_jsii_.Get(
		j,
		"athena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AthenaInput() *AwsDataSource_AthenaProperty {
	var returns *AwsDataSource_AthenaProperty
	_jsii_.Get(
		j,
		"athenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Aurora() AwsDataSource_AuroraPropertyOutputReference {
	var returns AwsDataSource_AuroraPropertyOutputReference
	_jsii_.Get(
		j,
		"aurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AuroraInput() *AwsDataSource_AuroraProperty {
	var returns *AwsDataSource_AuroraProperty
	_jsii_.Get(
		j,
		"auroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AuroraPostgresql() AwsDataSource_AuroraPostgresqlPropertyOutputReference {
	var returns AwsDataSource_AuroraPostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"auroraPostgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AuroraPostgresqlInput() *AwsDataSource_AuroraPostgresqlProperty {
	var returns *AwsDataSource_AuroraPostgresqlProperty
	_jsii_.Get(
		j,
		"auroraPostgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AwsIotAnalytics() AwsDataSource_AwsIotAnalyticsPropertyOutputReference {
	var returns AwsDataSource_AwsIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"awsIotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) AwsIotAnalyticsInput() *AwsDataSource_AwsIotAnalyticsProperty {
	var returns *AwsDataSource_AwsIotAnalyticsProperty
	_jsii_.Get(
		j,
		"awsIotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Databricks() AwsDataSource_DatabricksPropertyOutputReference {
	var returns AwsDataSource_DatabricksPropertyOutputReference
	_jsii_.Get(
		j,
		"databricks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) DatabricksInput() *AwsDataSource_DatabricksProperty {
	var returns *AwsDataSource_DatabricksProperty
	_jsii_.Get(
		j,
		"databricksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) InternalValue() *AwsDataSource_ParametersProperty {
	var returns *AwsDataSource_ParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Jira() AwsDataSource_JiraPropertyOutputReference {
	var returns AwsDataSource_JiraPropertyOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) JiraInput() *AwsDataSource_JiraProperty {
	var returns *AwsDataSource_JiraProperty
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) MariaDb() AwsDataSource_MariaDbPropertyOutputReference {
	var returns AwsDataSource_MariaDbPropertyOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) MariaDbInput() *AwsDataSource_MariaDbProperty {
	var returns *AwsDataSource_MariaDbProperty
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Mysql() AwsDataSource_MysqlPropertyOutputReference {
	var returns AwsDataSource_MysqlPropertyOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) MysqlInput() *AwsDataSource_MysqlProperty {
	var returns *AwsDataSource_MysqlProperty
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Oracle() AwsDataSource_OraclePropertyOutputReference {
	var returns AwsDataSource_OraclePropertyOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) OracleInput() *AwsDataSource_OracleProperty {
	var returns *AwsDataSource_OracleProperty
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Postgresql() AwsDataSource_PostgresqlPropertyOutputReference {
	var returns AwsDataSource_PostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PostgresqlInput() *AwsDataSource_PostgresqlProperty {
	var returns *AwsDataSource_PostgresqlProperty
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Presto() AwsDataSource_PrestoPropertyOutputReference {
	var returns AwsDataSource_PrestoPropertyOutputReference
	_jsii_.Get(
		j,
		"presto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PrestoInput() *AwsDataSource_PrestoProperty {
	var returns *AwsDataSource_PrestoProperty
	_jsii_.Get(
		j,
		"prestoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Rds() AwsDataSource_RdsPropertyOutputReference {
	var returns AwsDataSource_RdsPropertyOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) RdsInput() *AwsDataSource_RdsProperty {
	var returns *AwsDataSource_RdsProperty
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Redshift() AwsDataSource_RedshiftPropertyOutputReference {
	var returns AwsDataSource_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) RedshiftInput() *AwsDataSource_RedshiftProperty {
	var returns *AwsDataSource_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) S3() AwsDataSource_S3PropertyOutputReference {
	var returns AwsDataSource_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) S3Input() *AwsDataSource_S3Property {
	var returns *AwsDataSource_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ServiceNow() AwsDataSource_ServiceNowPropertyOutputReference {
	var returns AwsDataSource_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ServiceNowInput() *AwsDataSource_ServiceNowProperty {
	var returns *AwsDataSource_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Snowflake() AwsDataSource_SnowflakePropertyOutputReference {
	var returns AwsDataSource_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) SnowflakeInput() *AwsDataSource_SnowflakeProperty {
	var returns *AwsDataSource_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Spark() AwsDataSource_SparkPropertyOutputReference {
	var returns AwsDataSource_SparkPropertyOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) SparkInput() *AwsDataSource_SparkProperty {
	var returns *AwsDataSource_SparkProperty
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) SqlServer() AwsDataSource_SqlServerPropertyOutputReference {
	var returns AwsDataSource_SqlServerPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) SqlServerInput() *AwsDataSource_SqlServerProperty {
	var returns *AwsDataSource_SqlServerProperty
	_jsii_.Get(
		j,
		"sqlServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Teradata() AwsDataSource_TeradataPropertyOutputReference {
	var returns AwsDataSource_TeradataPropertyOutputReference
	_jsii_.Get(
		j,
		"teradata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) TeradataInput() *AwsDataSource_TeradataProperty {
	var returns *AwsDataSource_TeradataProperty
	_jsii_.Get(
		j,
		"teradataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Twitter() AwsDataSource_TwitterPropertyOutputReference {
	var returns AwsDataSource_TwitterPropertyOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) TwitterInput() *AwsDataSource_TwitterProperty {
	var returns *AwsDataSource_TwitterProperty
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSource_ParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDataSource_ParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSource_ParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSource_ParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSource_ParametersPropertyOutputReference_Override(a AwsDataSource_ParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference)SetInternalValue(val *AwsDataSource_ParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutAmazonElasticsearch(value *AwsDataSource_AmazonElasticsearchProperty) {
	if err := a.validatePutAmazonElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonElasticsearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutAthena(value *AwsDataSource_AthenaProperty) {
	if err := a.validatePutAthenaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAthena",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutAurora(value *AwsDataSource_AuroraProperty) {
	if err := a.validatePutAuroraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAurora",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutAuroraPostgresql(value *AwsDataSource_AuroraPostgresqlProperty) {
	if err := a.validatePutAuroraPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuroraPostgresql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutAwsIotAnalytics(value *AwsDataSource_AwsIotAnalyticsProperty) {
	if err := a.validatePutAwsIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAwsIotAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutDatabricks(value *AwsDataSource_DatabricksProperty) {
	if err := a.validatePutDatabricksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabricks",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutJira(value *AwsDataSource_JiraProperty) {
	if err := a.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJira",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutMariaDb(value *AwsDataSource_MariaDbProperty) {
	if err := a.validatePutMariaDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutMysql(value *AwsDataSource_MysqlProperty) {
	if err := a.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMysql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutOracle(value *AwsDataSource_OracleProperty) {
	if err := a.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOracle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutPostgresql(value *AwsDataSource_PostgresqlProperty) {
	if err := a.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutPresto(value *AwsDataSource_PrestoProperty) {
	if err := a.validatePutPrestoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPresto",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutRds(value *AwsDataSource_RdsProperty) {
	if err := a.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRds",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutRedshift(value *AwsDataSource_RedshiftProperty) {
	if err := a.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedshift",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutS3(value *AwsDataSource_S3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutServiceNow(value *AwsDataSource_ServiceNowProperty) {
	if err := a.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutSnowflake(value *AwsDataSource_SnowflakeProperty) {
	if err := a.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutSpark(value *AwsDataSource_SparkProperty) {
	if err := a.validatePutSparkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpark",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutSqlServer(value *AwsDataSource_SqlServerProperty) {
	if err := a.validatePutSqlServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlServer",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutTeradata(value *AwsDataSource_TeradataProperty) {
	if err := a.validatePutTeradataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTeradata",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) PutTwitter(value *AwsDataSource_TwitterProperty) {
	if err := a.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTwitter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetAmazonElasticsearch() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonElasticsearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetAthena() {
	_jsii_.InvokeVoid(
		a,
		"resetAthena",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetAurora() {
	_jsii_.InvokeVoid(
		a,
		"resetAurora",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetAuroraPostgresql() {
	_jsii_.InvokeVoid(
		a,
		"resetAuroraPostgresql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetAwsIotAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetAwsIotAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetDatabricks() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabricks",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		a,
		"resetJira",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetMariaDb() {
	_jsii_.InvokeVoid(
		a,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetMysql() {
	_jsii_.InvokeVoid(
		a,
		"resetMysql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetOracle() {
	_jsii_.InvokeVoid(
		a,
		"resetOracle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetPostgresql() {
	_jsii_.InvokeVoid(
		a,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetPresto() {
	_jsii_.InvokeVoid(
		a,
		"resetPresto",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetRds() {
	_jsii_.InvokeVoid(
		a,
		"resetRds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		a,
		"resetRedshift",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		a,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetSpark() {
	_jsii_.InvokeVoid(
		a,
		"resetSpark",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetSqlServer() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlServer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetTeradata() {
	_jsii_.InvokeVoid(
		a,
		"resetTeradata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		a,
		"resetTwitter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSource_ParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

