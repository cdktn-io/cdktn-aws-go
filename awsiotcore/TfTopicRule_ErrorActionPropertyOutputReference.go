package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTopicRule_ErrorActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchAlarm() TfTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	// Experimental.
	CloudwatchAlarmInput() *TfTopicRule_ErrorActionCloudwatchAlarmProperty
	// Experimental.
	CloudwatchLogs() TfTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	// Experimental.
	CloudwatchLogsInput() *TfTopicRule_ErrorActionCloudwatchLogsProperty
	// Experimental.
	CloudwatchMetric() TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	// Experimental.
	CloudwatchMetricInput() *TfTopicRule_ErrorActionCloudwatchMetricProperty
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
	Dynamodb() TfTopicRule_ErrorActionDynamodbPropertyOutputReference
	// Experimental.
	DynamodbInput() *TfTopicRule_ErrorActionDynamodbProperty
	// Experimental.
	Dynamodbv2() TfTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	// Experimental.
	Dynamodbv2Input() *TfTopicRule_ErrorActionDynamodbv2Property
	// Experimental.
	Elasticsearch() TfTopicRule_ErrorActionElasticsearchPropertyOutputReference
	// Experimental.
	ElasticsearchInput() *TfTopicRule_ErrorActionElasticsearchProperty
	// Experimental.
	Firehose() TfTopicRule_ErrorActionFirehosePropertyOutputReference
	// Experimental.
	FirehoseInput() *TfTopicRule_ErrorActionFirehoseProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Http() TfTopicRule_ErrorActionHttpPropertyOutputReference
	// Experimental.
	HttpInput() *TfTopicRule_ErrorActionHttpProperty
	// Experimental.
	InternalValue() *TfTopicRule_ErrorActionProperty
	// Experimental.
	SetInternalValue(val *TfTopicRule_ErrorActionProperty)
	// Experimental.
	IotAnalytics() TfTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	// Experimental.
	IotAnalyticsInput() *TfTopicRule_ErrorActionIotAnalyticsProperty
	// Experimental.
	IotEvents() TfTopicRule_ErrorActionIotEventsPropertyOutputReference
	// Experimental.
	IotEventsInput() *TfTopicRule_ErrorActionIotEventsProperty
	// Experimental.
	Kafka() TfTopicRule_ErrorActionKafkaPropertyOutputReference
	// Experimental.
	KafkaInput() *TfTopicRule_ErrorActionKafkaProperty
	// Experimental.
	Kinesis() TfTopicRule_ErrorActionKinesisPropertyOutputReference
	// Experimental.
	KinesisInput() *TfTopicRule_ErrorActionKinesisProperty
	// Experimental.
	Lambda() TfTopicRule_ErrorActionLambdaPropertyOutputReference
	// Experimental.
	LambdaInput() *TfTopicRule_ErrorActionLambdaProperty
	// Experimental.
	Republish() TfTopicRule_ErrorActionRepublishPropertyOutputReference
	// Experimental.
	RepublishInput() *TfTopicRule_ErrorActionRepublishProperty
	// Experimental.
	S3() TfTopicRule_ErrorActionS3PropertyOutputReference
	// Experimental.
	S3Input() *TfTopicRule_ErrorActionS3Property
	// Experimental.
	Sns() TfTopicRule_ErrorActionSnsPropertyOutputReference
	// Experimental.
	SnsInput() *TfTopicRule_ErrorActionSnsProperty
	// Experimental.
	Sqs() TfTopicRule_ErrorActionSqsPropertyOutputReference
	// Experimental.
	SqsInput() *TfTopicRule_ErrorActionSqsProperty
	// Experimental.
	StepFunctions() TfTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	// Experimental.
	StepFunctionsInput() *TfTopicRule_ErrorActionStepFunctionsProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timestream() TfTopicRule_ErrorActionTimestreamPropertyOutputReference
	// Experimental.
	TimestreamInput() *TfTopicRule_ErrorActionTimestreamProperty
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
	PutCloudwatchAlarm(value *TfTopicRule_ErrorActionCloudwatchAlarmProperty)
	// Experimental.
	PutCloudwatchLogs(value *TfTopicRule_ErrorActionCloudwatchLogsProperty)
	// Experimental.
	PutCloudwatchMetric(value *TfTopicRule_ErrorActionCloudwatchMetricProperty)
	// Experimental.
	PutDynamodb(value *TfTopicRule_ErrorActionDynamodbProperty)
	// Experimental.
	PutDynamodbv2(value *TfTopicRule_ErrorActionDynamodbv2Property)
	// Experimental.
	PutElasticsearch(value *TfTopicRule_ErrorActionElasticsearchProperty)
	// Experimental.
	PutFirehose(value *TfTopicRule_ErrorActionFirehoseProperty)
	// Experimental.
	PutHttp(value *TfTopicRule_ErrorActionHttpProperty)
	// Experimental.
	PutIotAnalytics(value *TfTopicRule_ErrorActionIotAnalyticsProperty)
	// Experimental.
	PutIotEvents(value *TfTopicRule_ErrorActionIotEventsProperty)
	// Experimental.
	PutKafka(value *TfTopicRule_ErrorActionKafkaProperty)
	// Experimental.
	PutKinesis(value *TfTopicRule_ErrorActionKinesisProperty)
	// Experimental.
	PutLambda(value *TfTopicRule_ErrorActionLambdaProperty)
	// Experimental.
	PutRepublish(value *TfTopicRule_ErrorActionRepublishProperty)
	// Experimental.
	PutS3(value *TfTopicRule_ErrorActionS3Property)
	// Experimental.
	PutSns(value *TfTopicRule_ErrorActionSnsProperty)
	// Experimental.
	PutSqs(value *TfTopicRule_ErrorActionSqsProperty)
	// Experimental.
	PutStepFunctions(value *TfTopicRule_ErrorActionStepFunctionsProperty)
	// Experimental.
	PutTimestream(value *TfTopicRule_ErrorActionTimestreamProperty)
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

// The jsii proxy struct for TfTopicRule_ErrorActionPropertyOutputReference
type jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarm() TfTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference {
	var returns TfTopicRule_ErrorActionCloudwatchAlarmPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchAlarmInput() *TfTopicRule_ErrorActionCloudwatchAlarmProperty {
	var returns *TfTopicRule_ErrorActionCloudwatchAlarmProperty
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogs() TfTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionCloudwatchLogsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchLogsInput() *TfTopicRule_ErrorActionCloudwatchLogsProperty {
	var returns *TfTopicRule_ErrorActionCloudwatchLogsProperty
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetric() TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference {
	var returns TfTopicRule_ErrorActionCloudwatchMetricPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CloudwatchMetricInput() *TfTopicRule_ErrorActionCloudwatchMetricProperty {
	var returns *TfTopicRule_ErrorActionCloudwatchMetricProperty
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Dynamodb() TfTopicRule_ErrorActionDynamodbPropertyOutputReference {
	var returns TfTopicRule_ErrorActionDynamodbPropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) DynamodbInput() *TfTopicRule_ErrorActionDynamodbProperty {
	var returns *TfTopicRule_ErrorActionDynamodbProperty
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2() TfTopicRule_ErrorActionDynamodbv2PropertyOutputReference {
	var returns TfTopicRule_ErrorActionDynamodbv2PropertyOutputReference
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Dynamodbv2Input() *TfTopicRule_ErrorActionDynamodbv2Property {
	var returns *TfTopicRule_ErrorActionDynamodbv2Property
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Elasticsearch() TfTopicRule_ErrorActionElasticsearchPropertyOutputReference {
	var returns TfTopicRule_ErrorActionElasticsearchPropertyOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ElasticsearchInput() *TfTopicRule_ErrorActionElasticsearchProperty {
	var returns *TfTopicRule_ErrorActionElasticsearchProperty
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Firehose() TfTopicRule_ErrorActionFirehosePropertyOutputReference {
	var returns TfTopicRule_ErrorActionFirehosePropertyOutputReference
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) FirehoseInput() *TfTopicRule_ErrorActionFirehoseProperty {
	var returns *TfTopicRule_ErrorActionFirehoseProperty
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Http() TfTopicRule_ErrorActionHttpPropertyOutputReference {
	var returns TfTopicRule_ErrorActionHttpPropertyOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) HttpInput() *TfTopicRule_ErrorActionHttpProperty {
	var returns *TfTopicRule_ErrorActionHttpProperty
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) InternalValue() *TfTopicRule_ErrorActionProperty {
	var returns *TfTopicRule_ErrorActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) IotAnalytics() TfTopicRule_ErrorActionIotAnalyticsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionIotAnalyticsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) IotAnalyticsInput() *TfTopicRule_ErrorActionIotAnalyticsProperty {
	var returns *TfTopicRule_ErrorActionIotAnalyticsProperty
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) IotEvents() TfTopicRule_ErrorActionIotEventsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionIotEventsPropertyOutputReference
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) IotEventsInput() *TfTopicRule_ErrorActionIotEventsProperty {
	var returns *TfTopicRule_ErrorActionIotEventsProperty
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Kafka() TfTopicRule_ErrorActionKafkaPropertyOutputReference {
	var returns TfTopicRule_ErrorActionKafkaPropertyOutputReference
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) KafkaInput() *TfTopicRule_ErrorActionKafkaProperty {
	var returns *TfTopicRule_ErrorActionKafkaProperty
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Kinesis() TfTopicRule_ErrorActionKinesisPropertyOutputReference {
	var returns TfTopicRule_ErrorActionKinesisPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) KinesisInput() *TfTopicRule_ErrorActionKinesisProperty {
	var returns *TfTopicRule_ErrorActionKinesisProperty
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Lambda() TfTopicRule_ErrorActionLambdaPropertyOutputReference {
	var returns TfTopicRule_ErrorActionLambdaPropertyOutputReference
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) LambdaInput() *TfTopicRule_ErrorActionLambdaProperty {
	var returns *TfTopicRule_ErrorActionLambdaProperty
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Republish() TfTopicRule_ErrorActionRepublishPropertyOutputReference {
	var returns TfTopicRule_ErrorActionRepublishPropertyOutputReference
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) RepublishInput() *TfTopicRule_ErrorActionRepublishProperty {
	var returns *TfTopicRule_ErrorActionRepublishProperty
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) S3() TfTopicRule_ErrorActionS3PropertyOutputReference {
	var returns TfTopicRule_ErrorActionS3PropertyOutputReference
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) S3Input() *TfTopicRule_ErrorActionS3Property {
	var returns *TfTopicRule_ErrorActionS3Property
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Sns() TfTopicRule_ErrorActionSnsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionSnsPropertyOutputReference
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) SnsInput() *TfTopicRule_ErrorActionSnsProperty {
	var returns *TfTopicRule_ErrorActionSnsProperty
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Sqs() TfTopicRule_ErrorActionSqsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionSqsPropertyOutputReference
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) SqsInput() *TfTopicRule_ErrorActionSqsProperty {
	var returns *TfTopicRule_ErrorActionSqsProperty
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) StepFunctions() TfTopicRule_ErrorActionStepFunctionsPropertyOutputReference {
	var returns TfTopicRule_ErrorActionStepFunctionsPropertyOutputReference
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) StepFunctionsInput() *TfTopicRule_ErrorActionStepFunctionsProperty {
	var returns *TfTopicRule_ErrorActionStepFunctionsProperty
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Timestream() TfTopicRule_ErrorActionTimestreamPropertyOutputReference {
	var returns TfTopicRule_ErrorActionTimestreamPropertyOutputReference
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) TimestreamInput() *TfTopicRule_ErrorActionTimestreamProperty {
	var returns *TfTopicRule_ErrorActionTimestreamProperty
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTopicRule_ErrorActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTopicRule_ErrorActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTopicRule_ErrorActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTopicRule_ErrorActionPropertyOutputReference_Override(t TfTopicRule_ErrorActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule.ErrorActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference)SetInternalValue(val *TfTopicRule_ErrorActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchAlarm(value *TfTopicRule_ErrorActionCloudwatchAlarmProperty) {
	if err := t.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchLogs(value *TfTopicRule_ErrorActionCloudwatchLogsProperty) {
	if err := t.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutCloudwatchMetric(value *TfTopicRule_ErrorActionCloudwatchMetricProperty) {
	if err := t.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutDynamodb(value *TfTopicRule_ErrorActionDynamodbProperty) {
	if err := t.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutDynamodbv2(value *TfTopicRule_ErrorActionDynamodbv2Property) {
	if err := t.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutElasticsearch(value *TfTopicRule_ErrorActionElasticsearchProperty) {
	if err := t.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutFirehose(value *TfTopicRule_ErrorActionFirehoseProperty) {
	if err := t.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirehose",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutHttp(value *TfTopicRule_ErrorActionHttpProperty) {
	if err := t.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutIotAnalytics(value *TfTopicRule_ErrorActionIotAnalyticsProperty) {
	if err := t.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutIotEvents(value *TfTopicRule_ErrorActionIotEventsProperty) {
	if err := t.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutKafka(value *TfTopicRule_ErrorActionKafkaProperty) {
	if err := t.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKafka",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutKinesis(value *TfTopicRule_ErrorActionKinesisProperty) {
	if err := t.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesis",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutLambda(value *TfTopicRule_ErrorActionLambdaProperty) {
	if err := t.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambda",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutRepublish(value *TfTopicRule_ErrorActionRepublishProperty) {
	if err := t.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRepublish",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutS3(value *TfTopicRule_ErrorActionS3Property) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutSns(value *TfTopicRule_ErrorActionSnsProperty) {
	if err := t.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutSqs(value *TfTopicRule_ErrorActionSqsProperty) {
	if err := t.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutStepFunctions(value *TfTopicRule_ErrorActionStepFunctionsProperty) {
	if err := t.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) PutTimestream(value *TfTopicRule_ErrorActionTimestreamProperty) {
	if err := t.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimestream",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetDynamodb() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		t,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetFirehose() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehose",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetIotEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetKafka() {
	_jsii_.InvokeVoid(
		t,
		"resetKafka",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetKinesis() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesis",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetRepublish() {
	_jsii_.InvokeVoid(
		t,
		"resetRepublish",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetSns() {
	_jsii_.InvokeVoid(
		t,
		"resetSns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetSqs() {
	_jsii_.InvokeVoid(
		t,
		"resetSqs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		t,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ResetTimestream() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestream",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTopicRule_ErrorActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

