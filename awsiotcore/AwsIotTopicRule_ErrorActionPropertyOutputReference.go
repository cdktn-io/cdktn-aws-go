package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsIotTopicRule_ErrorActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchAlarm() AwsIotTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	// Experimental.
	CloudwatchAlarmInput() *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty
	// Experimental.
	CloudwatchLogs() AwsIotTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty
	// Experimental.
	CloudwatchMetric() AwsIotTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	// Experimental.
	CloudwatchMetricInput() *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty
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
	Dynamodb() AwsIotTopicRule_ErrorActionDynamodbPropertyOutputReference
	// Experimental.
	DynamodbInput() *AwsIotTopicRule_ErrorActionDynamodbProperty
	// Experimental.
	Dynamodbv2() AwsIotTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	// Experimental.
	Dynamodbv2Input() *AwsIotTopicRule_ErrorActionDynamodbv2Property
	// Experimental.
	Elasticsearch() AwsIotTopicRule_ErrorActionElasticsearchPropertyOutputReference
	// Experimental.
	ElasticsearchInput() *AwsIotTopicRule_ErrorActionElasticsearchProperty
	// Experimental.
	Firehose() AwsIotTopicRule_ErrorActionFirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *AwsIotTopicRule_ErrorActionFirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Http() AwsIotTopicRule_ErrorActionHttpPropertyOutputReference
	// Experimental.
	HttpInput() *AwsIotTopicRule_ErrorActionHttpProperty
	// Experimental.
	InternalValue() *AwsIotTopicRule_ErrorActionProperty
	// Experimental.
	SetInternalValue(val *AwsIotTopicRule_ErrorActionProperty)
	// Experimental.
	IotAnalytics() AwsIotTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	// Experimental.
	IotAnalyticsInput() *AwsIotTopicRule_ErrorActionIotAnalyticsProperty
	// Experimental.
	IotEvents() AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference
	// Experimental.
	IotEventsInput() *AwsIotTopicRule_ErrorActionIotEventsProperty
	// Experimental.
	Kafka() AwsIotTopicRule_ErrorActionKafkaPropertyOutputReference
	// Experimental.
	KafkaInput() *AwsIotTopicRule_ErrorActionKafkaProperty
	// Experimental.
	Kinesis() AwsIotTopicRule_ErrorActionKinesisPropertyOutputReference
	// Experimental.
	KinesisInput() *AwsIotTopicRule_ErrorActionKinesisProperty
	// Experimental.
	Lambda() AwsIotTopicRule_ErrorActionLambdaPropertyOutputReference
	// Experimental.
	LambdaInput() *AwsIotTopicRule_ErrorActionLambdaProperty
	// Experimental.
	Republish() AwsIotTopicRule_ErrorActionRepublishPropertyOutputReference
	// Experimental.
	RepublishInput() *AwsIotTopicRule_ErrorActionRepublishProperty
	// Experimental.
	S3() AwsIotTopicRule_ErrorActionS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsIotTopicRule_ErrorActionS3Property
	// Experimental.
	Sns() AwsIotTopicRule_ErrorActionSnsPropertyOutputReference
	// Experimental.
	SnsInput() *AwsIotTopicRule_ErrorActionSnsProperty
	// Experimental.
	Sqs() AwsIotTopicRule_ErrorActionSqsPropertyOutputReference
	// Experimental.
	SqsInput() *AwsIotTopicRule_ErrorActionSqsProperty
	// Experimental.
	StepFunctions() AwsIotTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	// Experimental.
	StepFunctionsInput() *AwsIotTopicRule_ErrorActionStepFunctionsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timestream() AwsIotTopicRule_ErrorActionTimestreamPropertyOutputReference
	// Experimental.
	TimestreamInput() *AwsIotTopicRule_ErrorActionTimestreamProperty
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
	PutCloudwatchAlarm(value *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty)
	// Experimental.
	PutCloudwatchLogs(value *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty)
	// Experimental.
	PutCloudwatchMetric(value *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty)
	// Experimental.
	PutDynamodb(value *AwsIotTopicRule_ErrorActionDynamodbProperty)
	// Experimental.
	PutDynamodbv2(value *AwsIotTopicRule_ErrorActionDynamodbv2Property)
	// Experimental.
	PutElasticsearch(value *AwsIotTopicRule_ErrorActionElasticsearchProperty)
	// Experimental.
	PutFirehose(value *AwsIotTopicRule_ErrorActionFirehoseProperty)
	// Experimental.
	PutHttp(value *AwsIotTopicRule_ErrorActionHttpProperty)
	// Experimental.
	PutIotAnalytics(value *AwsIotTopicRule_ErrorActionIotAnalyticsProperty)
	// Experimental.
	PutIotEvents(value *AwsIotTopicRule_ErrorActionIotEventsProperty)
	// Experimental.
	PutKafka(value *AwsIotTopicRule_ErrorActionKafkaProperty)
	// Experimental.
	PutKinesis(value *AwsIotTopicRule_ErrorActionKinesisProperty)
	// Experimental.
	PutLambda(value *AwsIotTopicRule_ErrorActionLambdaProperty)
	// Experimental.
	PutRepublish(value *AwsIotTopicRule_ErrorActionRepublishProperty)
	// Experimental.
	PutS3(value *AwsIotTopicRule_ErrorActionS3Property)
	// Experimental.
	PutSns(value *AwsIotTopicRule_ErrorActionSnsProperty)
	// Experimental.
	PutSqs(value *AwsIotTopicRule_ErrorActionSqsProperty)
	// Experimental.
	PutStepFunctions(value *AwsIotTopicRule_ErrorActionStepFunctionsProperty)
	// Experimental.
	PutTimestream(value *AwsIotTopicRule_ErrorActionTimestreamProperty)
	// Experimental.
	ResetCloudwatchAlarm()
	// Experimental.
	ResetCloudwatchLogs()
	// Experimental.
	ResetCloudwatchMetric()
	// Experimental.
	ResetDynamodb()
	// Experimental.
	ResetDynamodbv2()
	// Experimental.
	ResetElasticsearch()
	// Experimental.
	ResetFirehose()
	// Experimental.
	ResetHttp()
	// Experimental.
	ResetIotAnalytics()
	// Experimental.
	ResetIotEvents()
	// Experimental.
	ResetKafka()
	// Experimental.
	ResetKinesis()
	// Experimental.
	ResetLambda()
	// Experimental.
	ResetRepublish()
	// Experimental.
	ResetS3()
	// Experimental.
	ResetSns()
	// Experimental.
	ResetSqs()
	// Experimental.
	ResetStepFunctions()
	// Experimental.
	ResetTimestream()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsIotTopicRule_ErrorActionPropertyOutputReference
type jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarm() AwsIotTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarmInput() *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty {
	var returns *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogs() AwsIotTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogsInput() *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty {
	var returns *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetric() AwsIotTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetricInput() *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty {
	var returns *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Dynamodb() AwsIotTopicRule_ErrorActionDynamodbPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionDynamodbPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) DynamodbInput() *AwsIotTopicRule_ErrorActionDynamodbProperty {
	var returns *AwsIotTopicRule_ErrorActionDynamodbProperty
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2() AwsIotTopicRule_ErrorActionDynamodbv2PropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2Input() *AwsIotTopicRule_ErrorActionDynamodbv2Property {
	var returns *AwsIotTopicRule_ErrorActionDynamodbv2Property
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Elasticsearch() AwsIotTopicRule_ErrorActionElasticsearchPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ElasticsearchInput() *AwsIotTopicRule_ErrorActionElasticsearchProperty {
	var returns *AwsIotTopicRule_ErrorActionElasticsearchProperty
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Firehose() AwsIotTopicRule_ErrorActionFirehosePropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) FirehoseInput() *AwsIotTopicRule_ErrorActionFirehoseProperty {
	var returns *AwsIotTopicRule_ErrorActionFirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Http() AwsIotTopicRule_ErrorActionHttpPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionHttpPropertyOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) HttpInput() *AwsIotTopicRule_ErrorActionHttpProperty {
	var returns *AwsIotTopicRule_ErrorActionHttpProperty
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) InternalValue() *AwsIotTopicRule_ErrorActionProperty {
	var returns *AwsIotTopicRule_ErrorActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) IotAnalytics() AwsIotTopicRule_ErrorActionIotAnalyticsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) IotAnalyticsInput() *AwsIotTopicRule_ErrorActionIotAnalyticsProperty {
	var returns *AwsIotTopicRule_ErrorActionIotAnalyticsProperty
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) IotEvents() AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionIotEventsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) IotEventsInput() *AwsIotTopicRule_ErrorActionIotEventsProperty {
	var returns *AwsIotTopicRule_ErrorActionIotEventsProperty
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Kafka() AwsIotTopicRule_ErrorActionKafkaPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionKafkaPropertyOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) KafkaInput() *AwsIotTopicRule_ErrorActionKafkaProperty {
	var returns *AwsIotTopicRule_ErrorActionKafkaProperty
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Kinesis() AwsIotTopicRule_ErrorActionKinesisPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionKinesisPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) KinesisInput() *AwsIotTopicRule_ErrorActionKinesisProperty {
	var returns *AwsIotTopicRule_ErrorActionKinesisProperty
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Lambda() AwsIotTopicRule_ErrorActionLambdaPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionLambdaPropertyOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) LambdaInput() *AwsIotTopicRule_ErrorActionLambdaProperty {
	var returns *AwsIotTopicRule_ErrorActionLambdaProperty
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Republish() AwsIotTopicRule_ErrorActionRepublishPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionRepublishPropertyOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) RepublishInput() *AwsIotTopicRule_ErrorActionRepublishProperty {
	var returns *AwsIotTopicRule_ErrorActionRepublishProperty
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) S3() AwsIotTopicRule_ErrorActionS3PropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) S3Input() *AwsIotTopicRule_ErrorActionS3Property {
	var returns *AwsIotTopicRule_ErrorActionS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Sns() AwsIotTopicRule_ErrorActionSnsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionSnsPropertyOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) SnsInput() *AwsIotTopicRule_ErrorActionSnsProperty {
	var returns *AwsIotTopicRule_ErrorActionSnsProperty
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Sqs() AwsIotTopicRule_ErrorActionSqsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionSqsPropertyOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) SqsInput() *AwsIotTopicRule_ErrorActionSqsProperty {
	var returns *AwsIotTopicRule_ErrorActionSqsProperty
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) StepFunctions() AwsIotTopicRule_ErrorActionStepFunctionsPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) StepFunctionsInput() *AwsIotTopicRule_ErrorActionStepFunctionsProperty {
	var returns *AwsIotTopicRule_ErrorActionStepFunctionsProperty
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Timestream() AwsIotTopicRule_ErrorActionTimestreamPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionTimestreamPropertyOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) TimestreamInput() *AwsIotTopicRule_ErrorActionTimestreamProperty {
	var returns *AwsIotTopicRule_ErrorActionTimestreamProperty
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsIotTopicRule_ErrorActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsIotTopicRule_ErrorActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsIotTopicRule_ErrorActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsIotTopicRule_ErrorActionPropertyOutputReference_Override(a AwsIotTopicRule_ErrorActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference)SetInternalValue(val *AwsIotTopicRule_ErrorActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchAlarm(value *AwsIotTopicRule_ErrorActionCloudwatchAlarmProperty) {
	if err := a.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchLogs(value *AwsIotTopicRule_ErrorActionCloudwatchLogsProperty) {
	if err := a.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchMetric(value *AwsIotTopicRule_ErrorActionCloudwatchMetricProperty) {
	if err := a.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutDynamodb(value *AwsIotTopicRule_ErrorActionDynamodbProperty) {
	if err := a.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutDynamodbv2(value *AwsIotTopicRule_ErrorActionDynamodbv2Property) {
	if err := a.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutElasticsearch(value *AwsIotTopicRule_ErrorActionElasticsearchProperty) {
	if err := a.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutFirehose(value *AwsIotTopicRule_ErrorActionFirehoseProperty) {
	if err := a.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutHttp(value *AwsIotTopicRule_ErrorActionHttpProperty) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutIotAnalytics(value *AwsIotTopicRule_ErrorActionIotAnalyticsProperty) {
	if err := a.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutIotEvents(value *AwsIotTopicRule_ErrorActionIotEventsProperty) {
	if err := a.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutKafka(value *AwsIotTopicRule_ErrorActionKafkaProperty) {
	if err := a.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKafka",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutKinesis(value *AwsIotTopicRule_ErrorActionKinesisProperty) {
	if err := a.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesis",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutLambda(value *AwsIotTopicRule_ErrorActionLambdaProperty) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutRepublish(value *AwsIotTopicRule_ErrorActionRepublishProperty) {
	if err := a.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRepublish",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutS3(value *AwsIotTopicRule_ErrorActionS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutSns(value *AwsIotTopicRule_ErrorActionSnsProperty) {
	if err := a.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutSqs(value *AwsIotTopicRule_ErrorActionSqsProperty) {
	if err := a.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutStepFunctions(value *AwsIotTopicRule_ErrorActionStepFunctionsProperty) {
	if err := a.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) PutTimestream(value *AwsIotTopicRule_ErrorActionTimestreamProperty) {
	if err := a.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimestream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetDynamodb() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		a,
		"resetKafka",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		a,
		"resetRepublish",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		a,
		"resetSns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		a,
		"resetSqs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsIotTopicRule_ErrorActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

