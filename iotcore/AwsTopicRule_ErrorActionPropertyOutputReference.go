package iotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/iotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/iotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTopicRule_ErrorActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchAlarm() AwsTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	// Experimental.
	CloudwatchAlarmInput() *AwsTopicRule_ErrorActionCloudwatchAlarmProperty
	// Experimental.
	CloudwatchLogs() AwsTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *AwsTopicRule_ErrorActionCloudwatchLogsProperty
	// Experimental.
	CloudwatchMetric() AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	// Experimental.
	CloudwatchMetricInput() *AwsTopicRule_ErrorActionCloudwatchMetricProperty
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
	Dynamodb() AwsTopicRule_ErrorActionDynamodbPropertyOutputReference
	// Experimental.
	DynamodbInput() *AwsTopicRule_ErrorActionDynamodbProperty
	// Experimental.
	Dynamodbv2() AwsTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	// Experimental.
	Dynamodbv2Input() *AwsTopicRule_ErrorActionDynamodbv2Property
	// Experimental.
	Elasticsearch() AwsTopicRule_ErrorActionElasticsearchPropertyOutputReference
	// Experimental.
	ElasticsearchInput() *AwsTopicRule_ErrorActionElasticsearchProperty
	// Experimental.
	Firehose() AwsTopicRule_ErrorActionFirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *AwsTopicRule_ErrorActionFirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Http() AwsTopicRule_ErrorActionHttpPropertyOutputReference
	// Experimental.
	HttpInput() *AwsTopicRule_ErrorActionHttpProperty
	// Experimental.
	InternalValue() *AwsTopicRule_ErrorActionProperty
	// Experimental.
	SetInternalValue(val *AwsTopicRule_ErrorActionProperty)
	// Experimental.
	IotAnalytics() AwsTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	// Experimental.
	IotAnalyticsInput() *AwsTopicRule_ErrorActionIotAnalyticsProperty
	// Experimental.
	IotEvents() AwsTopicRule_ErrorActionIotEventsPropertyOutputReference
	// Experimental.
	IotEventsInput() *AwsTopicRule_ErrorActionIotEventsProperty
	// Experimental.
	Kafka() AwsTopicRule_ErrorActionKafkaPropertyOutputReference
	// Experimental.
	KafkaInput() *AwsTopicRule_ErrorActionKafkaProperty
	// Experimental.
	Kinesis() AwsTopicRule_ErrorActionKinesisPropertyOutputReference
	// Experimental.
	KinesisInput() *AwsTopicRule_ErrorActionKinesisProperty
	// Experimental.
	Lambda() AwsTopicRule_ErrorActionLambdaPropertyOutputReference
	// Experimental.
	LambdaInput() *AwsTopicRule_ErrorActionLambdaProperty
	// Experimental.
	Republish() AwsTopicRule_ErrorActionRepublishPropertyOutputReference
	// Experimental.
	RepublishInput() *AwsTopicRule_ErrorActionRepublishProperty
	// Experimental.
	S3() AwsTopicRule_ErrorActionS3PropertyOutputReference
	// Experimental.
	S3Input() *AwsTopicRule_ErrorActionS3Property
	// Experimental.
	Sns() AwsTopicRule_ErrorActionSnsPropertyOutputReference
	// Experimental.
	SnsInput() *AwsTopicRule_ErrorActionSnsProperty
	// Experimental.
	Sqs() AwsTopicRule_ErrorActionSqsPropertyOutputReference
	// Experimental.
	SqsInput() *AwsTopicRule_ErrorActionSqsProperty
	// Experimental.
	StepFunctions() AwsTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	// Experimental.
	StepFunctionsInput() *AwsTopicRule_ErrorActionStepFunctionsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timestream() AwsTopicRule_ErrorActionTimestreamPropertyOutputReference
	// Experimental.
	TimestreamInput() *AwsTopicRule_ErrorActionTimestreamProperty
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
	PutCloudwatchAlarm(value *AwsTopicRule_ErrorActionCloudwatchAlarmProperty)
	// Experimental.
	PutCloudwatchLogs(value *AwsTopicRule_ErrorActionCloudwatchLogsProperty)
	// Experimental.
	PutCloudwatchMetric(value *AwsTopicRule_ErrorActionCloudwatchMetricProperty)
	// Experimental.
	PutDynamodb(value *AwsTopicRule_ErrorActionDynamodbProperty)
	// Experimental.
	PutDynamodbv2(value *AwsTopicRule_ErrorActionDynamodbv2Property)
	// Experimental.
	PutElasticsearch(value *AwsTopicRule_ErrorActionElasticsearchProperty)
	// Experimental.
	PutFirehose(value *AwsTopicRule_ErrorActionFirehoseProperty)
	// Experimental.
	PutHttp(value *AwsTopicRule_ErrorActionHttpProperty)
	// Experimental.
	PutIotAnalytics(value *AwsTopicRule_ErrorActionIotAnalyticsProperty)
	// Experimental.
	PutIotEvents(value *AwsTopicRule_ErrorActionIotEventsProperty)
	// Experimental.
	PutKafka(value *AwsTopicRule_ErrorActionKafkaProperty)
	// Experimental.
	PutKinesis(value *AwsTopicRule_ErrorActionKinesisProperty)
	// Experimental.
	PutLambda(value *AwsTopicRule_ErrorActionLambdaProperty)
	// Experimental.
	PutRepublish(value *AwsTopicRule_ErrorActionRepublishProperty)
	// Experimental.
	PutS3(value *AwsTopicRule_ErrorActionS3Property)
	// Experimental.
	PutSns(value *AwsTopicRule_ErrorActionSnsProperty)
	// Experimental.
	PutSqs(value *AwsTopicRule_ErrorActionSqsProperty)
	// Experimental.
	PutStepFunctions(value *AwsTopicRule_ErrorActionStepFunctionsProperty)
	// Experimental.
	PutTimestream(value *AwsTopicRule_ErrorActionTimestreamProperty)
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

// The jsii proxy struct for AwsTopicRule_ErrorActionPropertyOutputReference
type jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarm() AwsTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarmInput() *AwsTopicRule_ErrorActionCloudwatchAlarmProperty {
	var returns *AwsTopicRule_ErrorActionCloudwatchAlarmProperty
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogs() AwsTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogsInput() *AwsTopicRule_ErrorActionCloudwatchLogsProperty {
	var returns *AwsTopicRule_ErrorActionCloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetric() AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetricInput() *AwsTopicRule_ErrorActionCloudwatchMetricProperty {
	var returns *AwsTopicRule_ErrorActionCloudwatchMetricProperty
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Dynamodb() AwsTopicRule_ErrorActionDynamodbPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionDynamodbPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) DynamodbInput() *AwsTopicRule_ErrorActionDynamodbProperty {
	var returns *AwsTopicRule_ErrorActionDynamodbProperty
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2() AwsTopicRule_ErrorActionDynamodbv2PropertyOutputReference {
	var returns AwsTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2Input() *AwsTopicRule_ErrorActionDynamodbv2Property {
	var returns *AwsTopicRule_ErrorActionDynamodbv2Property
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Elasticsearch() AwsTopicRule_ErrorActionElasticsearchPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ElasticsearchInput() *AwsTopicRule_ErrorActionElasticsearchProperty {
	var returns *AwsTopicRule_ErrorActionElasticsearchProperty
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Firehose() AwsTopicRule_ErrorActionFirehosePropertyOutputReference {
	var returns AwsTopicRule_ErrorActionFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) FirehoseInput() *AwsTopicRule_ErrorActionFirehoseProperty {
	var returns *AwsTopicRule_ErrorActionFirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Http() AwsTopicRule_ErrorActionHttpPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionHttpPropertyOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) HttpInput() *AwsTopicRule_ErrorActionHttpProperty {
	var returns *AwsTopicRule_ErrorActionHttpProperty
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) InternalValue() *AwsTopicRule_ErrorActionProperty {
	var returns *AwsTopicRule_ErrorActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) IotAnalytics() AwsTopicRule_ErrorActionIotAnalyticsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) IotAnalyticsInput() *AwsTopicRule_ErrorActionIotAnalyticsProperty {
	var returns *AwsTopicRule_ErrorActionIotAnalyticsProperty
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) IotEvents() AwsTopicRule_ErrorActionIotEventsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionIotEventsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) IotEventsInput() *AwsTopicRule_ErrorActionIotEventsProperty {
	var returns *AwsTopicRule_ErrorActionIotEventsProperty
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Kafka() AwsTopicRule_ErrorActionKafkaPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionKafkaPropertyOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) KafkaInput() *AwsTopicRule_ErrorActionKafkaProperty {
	var returns *AwsTopicRule_ErrorActionKafkaProperty
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Kinesis() AwsTopicRule_ErrorActionKinesisPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionKinesisPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) KinesisInput() *AwsTopicRule_ErrorActionKinesisProperty {
	var returns *AwsTopicRule_ErrorActionKinesisProperty
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Lambda() AwsTopicRule_ErrorActionLambdaPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionLambdaPropertyOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) LambdaInput() *AwsTopicRule_ErrorActionLambdaProperty {
	var returns *AwsTopicRule_ErrorActionLambdaProperty
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Republish() AwsTopicRule_ErrorActionRepublishPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionRepublishPropertyOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) RepublishInput() *AwsTopicRule_ErrorActionRepublishProperty {
	var returns *AwsTopicRule_ErrorActionRepublishProperty
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) S3() AwsTopicRule_ErrorActionS3PropertyOutputReference {
	var returns AwsTopicRule_ErrorActionS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) S3Input() *AwsTopicRule_ErrorActionS3Property {
	var returns *AwsTopicRule_ErrorActionS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Sns() AwsTopicRule_ErrorActionSnsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionSnsPropertyOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) SnsInput() *AwsTopicRule_ErrorActionSnsProperty {
	var returns *AwsTopicRule_ErrorActionSnsProperty
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Sqs() AwsTopicRule_ErrorActionSqsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionSqsPropertyOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) SqsInput() *AwsTopicRule_ErrorActionSqsProperty {
	var returns *AwsTopicRule_ErrorActionSqsProperty
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) StepFunctions() AwsTopicRule_ErrorActionStepFunctionsPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) StepFunctionsInput() *AwsTopicRule_ErrorActionStepFunctionsProperty {
	var returns *AwsTopicRule_ErrorActionStepFunctionsProperty
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Timestream() AwsTopicRule_ErrorActionTimestreamPropertyOutputReference {
	var returns AwsTopicRule_ErrorActionTimestreamPropertyOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) TimestreamInput() *AwsTopicRule_ErrorActionTimestreamProperty {
	var returns *AwsTopicRule_ErrorActionTimestreamProperty
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTopicRule_ErrorActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTopicRule_ErrorActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTopicRule_ErrorActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTopicRule_ErrorActionPropertyOutputReference_Override(a AwsTopicRule_ErrorActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference)SetInternalValue(val *AwsTopicRule_ErrorActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchAlarm(value *AwsTopicRule_ErrorActionCloudwatchAlarmProperty) {
	if err := a.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchLogs(value *AwsTopicRule_ErrorActionCloudwatchLogsProperty) {
	if err := a.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchMetric(value *AwsTopicRule_ErrorActionCloudwatchMetricProperty) {
	if err := a.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutDynamodb(value *AwsTopicRule_ErrorActionDynamodbProperty) {
	if err := a.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutDynamodbv2(value *AwsTopicRule_ErrorActionDynamodbv2Property) {
	if err := a.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutElasticsearch(value *AwsTopicRule_ErrorActionElasticsearchProperty) {
	if err := a.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutFirehose(value *AwsTopicRule_ErrorActionFirehoseProperty) {
	if err := a.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutHttp(value *AwsTopicRule_ErrorActionHttpProperty) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutIotAnalytics(value *AwsTopicRule_ErrorActionIotAnalyticsProperty) {
	if err := a.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutIotEvents(value *AwsTopicRule_ErrorActionIotEventsProperty) {
	if err := a.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutKafka(value *AwsTopicRule_ErrorActionKafkaProperty) {
	if err := a.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKafka",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutKinesis(value *AwsTopicRule_ErrorActionKinesisProperty) {
	if err := a.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesis",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutLambda(value *AwsTopicRule_ErrorActionLambdaProperty) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutRepublish(value *AwsTopicRule_ErrorActionRepublishProperty) {
	if err := a.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRepublish",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutS3(value *AwsTopicRule_ErrorActionS3Property) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutSns(value *AwsTopicRule_ErrorActionSnsProperty) {
	if err := a.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutSqs(value *AwsTopicRule_ErrorActionSqsProperty) {
	if err := a.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutStepFunctions(value *AwsTopicRule_ErrorActionStepFunctionsProperty) {
	if err := a.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) PutTimestream(value *AwsTopicRule_ErrorActionTimestreamProperty) {
	if err := a.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimestream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetDynamodb() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		a,
		"resetKafka",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		a,
		"resetRepublish",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		a,
		"resetSns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		a,
		"resetSqs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTopicRule_ErrorActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

