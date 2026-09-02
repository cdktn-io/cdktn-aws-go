package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSource_ParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmazonElasticsearch() TfDataSource_AmazonElasticsearchPropertyOutputReference
	// Experimental.
	AmazonElasticsearchInput() *TfDataSource_AmazonElasticsearchProperty
	// Experimental.
	Athena() TfDataSource_AthenaPropertyOutputReference
	// Experimental.
	AthenaInput() *TfDataSource_AthenaProperty
	// Experimental.
	Aurora() TfDataSource_AuroraPropertyOutputReference
	// Experimental.
	AuroraInput() *TfDataSource_AuroraProperty
	// Experimental.
	AuroraPostgresql() TfDataSource_AuroraPostgresqlPropertyOutputReference
	// Experimental.
	AuroraPostgresqlInput() *TfDataSource_AuroraPostgresqlProperty
	// Experimental.
	AwsIotAnalytics() TfDataSource_AwsIotAnalyticsPropertyOutputReference
	// Experimental.
	AwsIotAnalyticsInput() *TfDataSource_AwsIotAnalyticsProperty
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
	Databricks() TfDataSource_DatabricksPropertyOutputReference
	// Experimental.
	DatabricksInput() *TfDataSource_DatabricksProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDataSource_ParametersProperty
	// Experimental.
	SetInternalValue(val *TfDataSource_ParametersProperty)
	// Experimental.
	Jira() TfDataSource_JiraPropertyOutputReference
	// Experimental.
	JiraInput() *TfDataSource_JiraProperty
	// Experimental.
	MariaDb() TfDataSource_MariaDbPropertyOutputReference
	// Experimental.
	MariaDbInput() *TfDataSource_MariaDbProperty
	// Experimental.
	Mysql() TfDataSource_MysqlPropertyOutputReference
	// Experimental.
	MysqlInput() *TfDataSource_MysqlProperty
	// Experimental.
	Oracle() TfDataSource_OraclePropertyOutputReference
	// Experimental.
	OracleInput() *TfDataSource_OracleProperty
	// Experimental.
	Postgresql() TfDataSource_PostgresqlPropertyOutputReference
	// Experimental.
	PostgresqlInput() *TfDataSource_PostgresqlProperty
	// Experimental.
	Presto() TfDataSource_PrestoPropertyOutputReference
	// Experimental.
	PrestoInput() *TfDataSource_PrestoProperty
	// Experimental.
	Rds() TfDataSource_RdsPropertyOutputReference
	// Experimental.
	RdsInput() *TfDataSource_RdsProperty
	// Experimental.
	Redshift() TfDataSource_RedshiftPropertyOutputReference
	// Experimental.
	RedshiftInput() *TfDataSource_RedshiftProperty
	// Experimental.
	S3() TfDataSource_S3PropertyOutputReference
	// Experimental.
	S3Input() *TfDataSource_S3Property
	// Experimental.
	ServiceNow() TfDataSource_ServiceNowPropertyOutputReference
	// Experimental.
	ServiceNowInput() *TfDataSource_ServiceNowProperty
	// Experimental.
	Snowflake() TfDataSource_SnowflakePropertyOutputReference
	// Experimental.
	SnowflakeInput() *TfDataSource_SnowflakeProperty
	// Experimental.
	Spark() TfDataSource_SparkPropertyOutputReference
	// Experimental.
	SparkInput() *TfDataSource_SparkProperty
	// Experimental.
	SqlServer() TfDataSource_SqlServerPropertyOutputReference
	// Experimental.
	SqlServerInput() *TfDataSource_SqlServerProperty
	// Experimental.
	Teradata() TfDataSource_TeradataPropertyOutputReference
	// Experimental.
	TeradataInput() *TfDataSource_TeradataProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Twitter() TfDataSource_TwitterPropertyOutputReference
	// Experimental.
	TwitterInput() *TfDataSource_TwitterProperty
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
	PutAmazonElasticsearch(value *TfDataSource_AmazonElasticsearchProperty)
	// Experimental.
	PutAthena(value *TfDataSource_AthenaProperty)
	// Experimental.
	PutAurora(value *TfDataSource_AuroraProperty)
	// Experimental.
	PutAuroraPostgresql(value *TfDataSource_AuroraPostgresqlProperty)
	// Experimental.
	PutAwsIotAnalytics(value *TfDataSource_AwsIotAnalyticsProperty)
	// Experimental.
	PutDatabricks(value *TfDataSource_DatabricksProperty)
	// Experimental.
	PutJira(value *TfDataSource_JiraProperty)
	// Experimental.
	PutMariaDb(value *TfDataSource_MariaDbProperty)
	// Experimental.
	PutMysql(value *TfDataSource_MysqlProperty)
	// Experimental.
	PutOracle(value *TfDataSource_OracleProperty)
	// Experimental.
	PutPostgresql(value *TfDataSource_PostgresqlProperty)
	// Experimental.
	PutPresto(value *TfDataSource_PrestoProperty)
	// Experimental.
	PutRds(value *TfDataSource_RdsProperty)
	// Experimental.
	PutRedshift(value *TfDataSource_RedshiftProperty)
	// Experimental.
	PutS3(value *TfDataSource_S3Property)
	// Experimental.
	PutServiceNow(value *TfDataSource_ServiceNowProperty)
	// Experimental.
	PutSnowflake(value *TfDataSource_SnowflakeProperty)
	// Experimental.
	PutSpark(value *TfDataSource_SparkProperty)
	// Experimental.
	PutSqlServer(value *TfDataSource_SqlServerProperty)
	// Experimental.
	PutTeradata(value *TfDataSource_TeradataProperty)
	// Experimental.
	PutTwitter(value *TfDataSource_TwitterProperty)
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

// The jsii proxy struct for TfDataSource_ParametersPropertyOutputReference
type jsiiProxy_TfDataSource_ParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AmazonElasticsearch() TfDataSource_AmazonElasticsearchPropertyOutputReference {
	var returns TfDataSource_AmazonElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonElasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AmazonElasticsearchInput() *TfDataSource_AmazonElasticsearchProperty {
	var returns *TfDataSource_AmazonElasticsearchProperty
	_jsii_.Get(
		j,
		"amazonElasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Athena() TfDataSource_AthenaPropertyOutputReference {
	var returns TfDataSource_AthenaPropertyOutputReference
	_jsii_.Get(
		j,
		"athena",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AthenaInput() *TfDataSource_AthenaProperty {
	var returns *TfDataSource_AthenaProperty
	_jsii_.Get(
		j,
		"athenaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Aurora() TfDataSource_AuroraPropertyOutputReference {
	var returns TfDataSource_AuroraPropertyOutputReference
	_jsii_.Get(
		j,
		"aurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AuroraInput() *TfDataSource_AuroraProperty {
	var returns *TfDataSource_AuroraProperty
	_jsii_.Get(
		j,
		"auroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AuroraPostgresql() TfDataSource_AuroraPostgresqlPropertyOutputReference {
	var returns TfDataSource_AuroraPostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"auroraPostgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AuroraPostgresqlInput() *TfDataSource_AuroraPostgresqlProperty {
	var returns *TfDataSource_AuroraPostgresqlProperty
	_jsii_.Get(
		j,
		"auroraPostgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AwsIotAnalytics() TfDataSource_AwsIotAnalyticsPropertyOutputReference {
	var returns TfDataSource_AwsIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"awsIotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) AwsIotAnalyticsInput() *TfDataSource_AwsIotAnalyticsProperty {
	var returns *TfDataSource_AwsIotAnalyticsProperty
	_jsii_.Get(
		j,
		"awsIotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Databricks() TfDataSource_DatabricksPropertyOutputReference {
	var returns TfDataSource_DatabricksPropertyOutputReference
	_jsii_.Get(
		j,
		"databricks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) DatabricksInput() *TfDataSource_DatabricksProperty {
	var returns *TfDataSource_DatabricksProperty
	_jsii_.Get(
		j,
		"databricksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) InternalValue() *TfDataSource_ParametersProperty {
	var returns *TfDataSource_ParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Jira() TfDataSource_JiraPropertyOutputReference {
	var returns TfDataSource_JiraPropertyOutputReference
	_jsii_.Get(
		j,
		"jira",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) JiraInput() *TfDataSource_JiraProperty {
	var returns *TfDataSource_JiraProperty
	_jsii_.Get(
		j,
		"jiraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) MariaDb() TfDataSource_MariaDbPropertyOutputReference {
	var returns TfDataSource_MariaDbPropertyOutputReference
	_jsii_.Get(
		j,
		"mariaDb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) MariaDbInput() *TfDataSource_MariaDbProperty {
	var returns *TfDataSource_MariaDbProperty
	_jsii_.Get(
		j,
		"mariaDbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Mysql() TfDataSource_MysqlPropertyOutputReference {
	var returns TfDataSource_MysqlPropertyOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) MysqlInput() *TfDataSource_MysqlProperty {
	var returns *TfDataSource_MysqlProperty
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Oracle() TfDataSource_OraclePropertyOutputReference {
	var returns TfDataSource_OraclePropertyOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) OracleInput() *TfDataSource_OracleProperty {
	var returns *TfDataSource_OracleProperty
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Postgresql() TfDataSource_PostgresqlPropertyOutputReference {
	var returns TfDataSource_PostgresqlPropertyOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PostgresqlInput() *TfDataSource_PostgresqlProperty {
	var returns *TfDataSource_PostgresqlProperty
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Presto() TfDataSource_PrestoPropertyOutputReference {
	var returns TfDataSource_PrestoPropertyOutputReference
	_jsii_.Get(
		j,
		"presto",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PrestoInput() *TfDataSource_PrestoProperty {
	var returns *TfDataSource_PrestoProperty
	_jsii_.Get(
		j,
		"prestoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Rds() TfDataSource_RdsPropertyOutputReference {
	var returns TfDataSource_RdsPropertyOutputReference
	_jsii_.Get(
		j,
		"rds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) RdsInput() *TfDataSource_RdsProperty {
	var returns *TfDataSource_RdsProperty
	_jsii_.Get(
		j,
		"rdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Redshift() TfDataSource_RedshiftPropertyOutputReference {
	var returns TfDataSource_RedshiftPropertyOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) RedshiftInput() *TfDataSource_RedshiftProperty {
	var returns *TfDataSource_RedshiftProperty
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) S3() TfDataSource_S3PropertyOutputReference {
	var returns TfDataSource_S3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) S3Input() *TfDataSource_S3Property {
	var returns *TfDataSource_S3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ServiceNow() TfDataSource_ServiceNowPropertyOutputReference {
	var returns TfDataSource_ServiceNowPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceNow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ServiceNowInput() *TfDataSource_ServiceNowProperty {
	var returns *TfDataSource_ServiceNowProperty
	_jsii_.Get(
		j,
		"serviceNowInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Snowflake() TfDataSource_SnowflakePropertyOutputReference {
	var returns TfDataSource_SnowflakePropertyOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) SnowflakeInput() *TfDataSource_SnowflakeProperty {
	var returns *TfDataSource_SnowflakeProperty
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Spark() TfDataSource_SparkPropertyOutputReference {
	var returns TfDataSource_SparkPropertyOutputReference
	_jsii_.Get(
		j,
		"spark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) SparkInput() *TfDataSource_SparkProperty {
	var returns *TfDataSource_SparkProperty
	_jsii_.Get(
		j,
		"sparkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) SqlServer() TfDataSource_SqlServerPropertyOutputReference {
	var returns TfDataSource_SqlServerPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlServer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) SqlServerInput() *TfDataSource_SqlServerProperty {
	var returns *TfDataSource_SqlServerProperty
	_jsii_.Get(
		j,
		"sqlServerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Teradata() TfDataSource_TeradataPropertyOutputReference {
	var returns TfDataSource_TeradataPropertyOutputReference
	_jsii_.Get(
		j,
		"teradata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) TeradataInput() *TfDataSource_TeradataProperty {
	var returns *TfDataSource_TeradataProperty
	_jsii_.Get(
		j,
		"teradataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Twitter() TfDataSource_TwitterPropertyOutputReference {
	var returns TfDataSource_TwitterPropertyOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) TwitterInput() *TfDataSource_TwitterProperty {
	var returns *TfDataSource_TwitterProperty
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSource_ParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataSource_ParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSource_ParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSource_ParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSource_ParametersPropertyOutputReference_Override(t TfDataSource_ParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSource.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference)SetInternalValue(val *TfDataSource_ParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSource_ParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutAmazonElasticsearch(value *TfDataSource_AmazonElasticsearchProperty) {
	if err := t.validatePutAmazonElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmazonElasticsearch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutAthena(value *TfDataSource_AthenaProperty) {
	if err := t.validatePutAthenaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAthena",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutAurora(value *TfDataSource_AuroraProperty) {
	if err := t.validatePutAuroraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAurora",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutAuroraPostgresql(value *TfDataSource_AuroraPostgresqlProperty) {
	if err := t.validatePutAuroraPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuroraPostgresql",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutAwsIotAnalytics(value *TfDataSource_AwsIotAnalyticsProperty) {
	if err := t.validatePutAwsIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAwsIotAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutDatabricks(value *TfDataSource_DatabricksProperty) {
	if err := t.validatePutDatabricksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatabricks",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutJira(value *TfDataSource_JiraProperty) {
	if err := t.validatePutJiraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJira",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutMariaDb(value *TfDataSource_MariaDbProperty) {
	if err := t.validatePutMariaDbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMariaDb",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutMysql(value *TfDataSource_MysqlProperty) {
	if err := t.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMysql",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutOracle(value *TfDataSource_OracleProperty) {
	if err := t.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOracle",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutPostgresql(value *TfDataSource_PostgresqlProperty) {
	if err := t.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutPresto(value *TfDataSource_PrestoProperty) {
	if err := t.validatePutPrestoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPresto",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutRds(value *TfDataSource_RdsProperty) {
	if err := t.validatePutRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRds",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutRedshift(value *TfDataSource_RedshiftProperty) {
	if err := t.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedshift",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutS3(value *TfDataSource_S3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutServiceNow(value *TfDataSource_ServiceNowProperty) {
	if err := t.validatePutServiceNowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceNow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutSnowflake(value *TfDataSource_SnowflakeProperty) {
	if err := t.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutSpark(value *TfDataSource_SparkProperty) {
	if err := t.validatePutSparkParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpark",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutSqlServer(value *TfDataSource_SqlServerProperty) {
	if err := t.validatePutSqlServerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqlServer",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutTeradata(value *TfDataSource_TeradataProperty) {
	if err := t.validatePutTeradataParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTeradata",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) PutTwitter(value *TfDataSource_TwitterProperty) {
	if err := t.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTwitter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetAmazonElasticsearch() {
	_jsii_.InvokeVoid(
		t,
		"resetAmazonElasticsearch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetAthena() {
	_jsii_.InvokeVoid(
		t,
		"resetAthena",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetAurora() {
	_jsii_.InvokeVoid(
		t,
		"resetAurora",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetAuroraPostgresql() {
	_jsii_.InvokeVoid(
		t,
		"resetAuroraPostgresql",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetAwsIotAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetAwsIotAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetDatabricks() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabricks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetJira() {
	_jsii_.InvokeVoid(
		t,
		"resetJira",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetMariaDb() {
	_jsii_.InvokeVoid(
		t,
		"resetMariaDb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetMysql() {
	_jsii_.InvokeVoid(
		t,
		"resetMysql",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetOracle() {
	_jsii_.InvokeVoid(
		t,
		"resetOracle",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetPostgresql() {
	_jsii_.InvokeVoid(
		t,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetPresto() {
	_jsii_.InvokeVoid(
		t,
		"resetPresto",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetRds() {
	_jsii_.InvokeVoid(
		t,
		"resetRds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetRedshift() {
	_jsii_.InvokeVoid(
		t,
		"resetRedshift",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetServiceNow() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceNow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetSnowflake() {
	_jsii_.InvokeVoid(
		t,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetSpark() {
	_jsii_.InvokeVoid(
		t,
		"resetSpark",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetSqlServer() {
	_jsii_.InvokeVoid(
		t,
		"resetSqlServer",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetTeradata() {
	_jsii_.InvokeVoid(
		t,
		"resetTeradata",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		t,
		"resetTwitter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSource_ParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

