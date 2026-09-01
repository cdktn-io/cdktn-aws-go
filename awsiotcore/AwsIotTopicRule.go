package awsiotcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsiotcore/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsiotcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule aws_iot_topic_rule}.
// Experimental.
type AwsIotTopicRule interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudwatchAlarm() AwsIotTopicRule_CloudwatchAlarmPropertyList
	// Experimental.
	CloudwatchAlarmInput() interface{}
	// Experimental.
	CloudwatchLogs() AwsIotTopicRule_CloudwatchLogsPropertyList
	// Experimental.
	CloudwatchLogsInput() interface{}
	// Experimental.
	CloudwatchMetric() AwsIotTopicRule_CloudwatchMetricPropertyList
	// Experimental.
	CloudwatchMetricInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	Dynamodb() AwsIotTopicRule_DynamodbPropertyList
	// Experimental.
	DynamodbInput() interface{}
	// Experimental.
	Dynamodbv2() AwsIotTopicRule_Dynamodbv2PropertyList
	// Experimental.
	Dynamodbv2Input() interface{}
	// Experimental.
	Elasticsearch() AwsIotTopicRule_ElasticsearchPropertyList
	// Experimental.
	ElasticsearchInput() interface{}
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	ErrorAction() AwsIotTopicRule_ErrorActionPropertyOutputReference
	// Experimental.
	ErrorActionInput() *AwsIotTopicRule_ErrorActionProperty
	// Experimental.
	Firehose() AwsIotTopicRule_FirehosePropertyList
	// Experimental.
	FirehoseInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Http() AwsIotTopicRule_HttpPropertyList
	// Experimental.
	HttpInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IotAnalytics() AwsIotTopicRule_IotAnalyticsPropertyList
	// Experimental.
	IotAnalyticsInput() interface{}
	// Experimental.
	IotEvents() AwsIotTopicRule_IotEventsPropertyList
	// Experimental.
	IotEventsInput() interface{}
	// Experimental.
	Kafka() AwsIotTopicRule_KafkaPropertyList
	// Experimental.
	KafkaInput() interface{}
	// Experimental.
	Kinesis() AwsIotTopicRule_KinesisPropertyList
	// Experimental.
	KinesisInput() interface{}
	// Experimental.
	Lambda() AwsIotTopicRule_LambdaPropertyList
	// Experimental.
	LambdaInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	Republish() AwsIotTopicRule_RepublishPropertyList
	// Experimental.
	RepublishInput() interface{}
	// Experimental.
	S3() AwsIotTopicRule_S3PropertyList
	// Experimental.
	S3Input() interface{}
	// Experimental.
	Sns() AwsIotTopicRule_SnsPropertyList
	// Experimental.
	SnsInput() interface{}
	// Experimental.
	Sql() *string
	// Experimental.
	SetSql(val *string)
	// Experimental.
	SqlInput() *string
	// Experimental.
	SqlVersion() *string
	// Experimental.
	SetSqlVersion(val *string)
	// Experimental.
	SqlVersionInput() *string
	// Experimental.
	Sqs() AwsIotTopicRule_SqsPropertyList
	// Experimental.
	SqsInput() interface{}
	// Experimental.
	StepFunctions() AwsIotTopicRule_StepFunctionsPropertyList
	// Experimental.
	StepFunctionsInput() interface{}
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() *map[string]*string
	// Experimental.
	SetTagsAll(val *map[string]*string)
	// Experimental.
	TagsAllInput() *map[string]*string
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timestream() AwsIotTopicRule_TimestreamPropertyList
	// Experimental.
	TimestreamInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktn.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Wraps a write-only attribute's already-mapped value so that `ProviderFeature.WRITE_ONLY_ATTRIBUTES` usage is registered at *resolve* time instead of at mutation time (setter/constructor). Called by generated bindings from `synthesizeAttributes()` and `synthesizeHclAttributes()`, e.g. `secret_key_wo: this.markWriteOnlyAttribute(cdktn.stringToTerraform(this._secretKeyWo))`; not intended to be called directly.
	//
	// `undefined` passes through completely unchanged, so the existing
	// undefined-filtering that omits unset attributes from synthesized
	// output (see `resolve()` in `tokens/private/resolve.ts`, and the
	// `value.value !== undefined` filter in generated
	// `synthesizeHclAttributes()`) keeps working untouched. `null` is also
	// passed through unchanged: it already renders as an explicit
	// null-out and must not arm the validation either.
	//
	// Any other value - including one that will itself resolve to nothing
	// (e.g. a `Lazy`/`IResolvable` producer with no value to contribute) -
	// is wrapped in a token whose `resolve()` defers to the real resolver
	// first and registers usage only if what comes back is not
	// `null`/`undefined`; the resolved value is then returned unchanged,
	// so what actually renders is untouched by this wrapper. A producer
	// that resolves to `undefined` therefore neither registers usage nor
	// leaves anything behind in the synthesized attribute - the omission
	// behaves exactly as if the attribute had never been set.
	//
	// Registration goes through `_registerResolveDiscoveredProviderFeatureUsage`
	// rather than `registerProviderFeatureUsage`: usage here is only known at
	// resolve time, and a given element can be resolved across many
	// synthesis passes over its lifetime (repeated `app.synth()` calls,
	// tests reusing a construct tree), so it must represent only the CURRENT
	// pass rather than accumulate forever. Every validation-enabled entry
	// point (`App.synth`; `Testing.synth`/`synthHcl` with validations;
	// `StackSynthesizer.synthesize`) runs a prepare step that deactivates any
	// stale registration and then resolves every element's `toTerraform()`
	// before that same entry point's validations run - see
	// `TerraformStack._runPreparingResolve` - so whatever this closure
	// (re-)registers during that prepare step is always visible to the
	// validation that reads it afterwards, and nothing left over from an
	// earlier pass leaks into the current one.
	// Experimental.
	MarkWriteOnlyAttribute(value interface{}) interface{}
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using its instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Experimental.
	PutCloudwatchAlarm(value interface{})
	// Experimental.
	PutCloudwatchLogs(value interface{})
	// Experimental.
	PutCloudwatchMetric(value interface{})
	// Experimental.
	PutDynamodb(value interface{})
	// Experimental.
	PutDynamodbv2(value interface{})
	// Experimental.
	PutElasticsearch(value interface{})
	// Experimental.
	PutErrorAction(value *AwsIotTopicRule_ErrorActionProperty)
	// Experimental.
	PutFirehose(value interface{})
	// Experimental.
	PutHttp(value interface{})
	// Experimental.
	PutIotAnalytics(value interface{})
	// Experimental.
	PutIotEvents(value interface{})
	// Experimental.
	PutKafka(value interface{})
	// Experimental.
	PutKinesis(value interface{})
	// Experimental.
	PutLambda(value interface{})
	// Experimental.
	PutRepublish(value interface{})
	// Experimental.
	PutS3(value interface{})
	// Experimental.
	PutSns(value interface{})
	// Experimental.
	PutSqs(value interface{})
	// Experimental.
	PutStepFunctions(value interface{})
	// Experimental.
	PutTimestream(value interface{})
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	// Experimental.
	ResetCloudwatchAlarm()
	// Experimental.
	ResetCloudwatchLogs()
	// Experimental.
	ResetCloudwatchMetric()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDynamodb()
	// Experimental.
	ResetDynamodbv2()
	// Experimental.
	ResetElasticsearch()
	// Experimental.
	ResetErrorAction()
	// Experimental.
	ResetFirehose()
	// Experimental.
	ResetHttp()
	// Experimental.
	ResetId()
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
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
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
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimestream()
	// Experimental.
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	// Experimental.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	// Experimental.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for AwsIotTopicRule
type jsiiProxy_AwsIotTopicRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsIotTopicRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchAlarm() AwsIotTopicRule_CloudwatchAlarmPropertyList {
	var returns AwsIotTopicRule_CloudwatchAlarmPropertyList
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchAlarmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchLogs() AwsIotTopicRule_CloudwatchLogsPropertyList {
	var returns AwsIotTopicRule_CloudwatchLogsPropertyList
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchMetric() AwsIotTopicRule_CloudwatchMetricPropertyList {
	var returns AwsIotTopicRule_CloudwatchMetricPropertyList
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) CloudwatchMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Dynamodb() AwsIotTopicRule_DynamodbPropertyList {
	var returns AwsIotTopicRule_DynamodbPropertyList
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) DynamodbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Dynamodbv2() AwsIotTopicRule_Dynamodbv2PropertyList {
	var returns AwsIotTopicRule_Dynamodbv2PropertyList
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Dynamodbv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Elasticsearch() AwsIotTopicRule_ElasticsearchPropertyList {
	var returns AwsIotTopicRule_ElasticsearchPropertyList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) ErrorAction() AwsIotTopicRule_ErrorActionPropertyOutputReference {
	var returns AwsIotTopicRule_ErrorActionPropertyOutputReference
	_jsii_.Get(
		j,
		"errorAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) ErrorActionInput() *AwsIotTopicRule_ErrorActionProperty {
	var returns *AwsIotTopicRule_ErrorActionProperty
	_jsii_.Get(
		j,
		"errorActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Firehose() AwsIotTopicRule_FirehosePropertyList {
	var returns AwsIotTopicRule_FirehosePropertyList
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Http() AwsIotTopicRule_HttpPropertyList {
	var returns AwsIotTopicRule_HttpPropertyList
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) IotAnalytics() AwsIotTopicRule_IotAnalyticsPropertyList {
	var returns AwsIotTopicRule_IotAnalyticsPropertyList
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) IotAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) IotEvents() AwsIotTopicRule_IotEventsPropertyList {
	var returns AwsIotTopicRule_IotEventsPropertyList
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Kafka() AwsIotTopicRule_KafkaPropertyList {
	var returns AwsIotTopicRule_KafkaPropertyList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Kinesis() AwsIotTopicRule_KinesisPropertyList {
	var returns AwsIotTopicRule_KinesisPropertyList
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) KinesisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Lambda() AwsIotTopicRule_LambdaPropertyList {
	var returns AwsIotTopicRule_LambdaPropertyList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Republish() AwsIotTopicRule_RepublishPropertyList {
	var returns AwsIotTopicRule_RepublishPropertyList
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) RepublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) S3() AwsIotTopicRule_S3PropertyList {
	var returns AwsIotTopicRule_S3PropertyList
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Sns() AwsIotTopicRule_SnsPropertyList {
	var returns AwsIotTopicRule_SnsPropertyList
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) SqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) SqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Sqs() AwsIotTopicRule_SqsPropertyList {
	var returns AwsIotTopicRule_SqsPropertyList
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) StepFunctions() AwsIotTopicRule_StepFunctionsPropertyList {
	var returns AwsIotTopicRule_StepFunctionsPropertyList
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) StepFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) Timestream() AwsIotTopicRule_TimestreamPropertyList {
	var returns AwsIotTopicRule_TimestreamPropertyList
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIotTopicRule) TimestreamInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timestreamInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule aws_iot_topic_rule} Resource.
// Experimental.
func NewAwsIotTopicRule(scope constructs.Construct, id *string, config *AwsIotTopicRuleConfig) AwsIotTopicRule {
	_init_.Initialize()

	if err := validateNewAwsIotTopicRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIotTopicRule{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule aws_iot_topic_rule} Resource.
// Experimental.
func NewAwsIotTopicRule_Override(a AwsIotTopicRule, scope constructs.Construct, id *string, config *AwsIotTopicRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetSql(val *string) {
	if err := j.validateSetSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetSqlVersion(val *string) {
	if err := j.validateSetSqlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlVersion",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsIotTopicRule)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a AwsIotTopicRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsIotTopicRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsIotTopicRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Experimental.
func AwsIotTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIotTopicRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsIotTopicRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIotTopicRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsIotTopicRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIotTopicRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsIotTopicRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-iot-core.AwsIotTopicRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIotTopicRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIotTopicRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIotTopicRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIotTopicRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIotTopicRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIotTopicRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIotTopicRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIotTopicRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := a.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutCloudwatchAlarm(value interface{}) {
	if err := a.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutCloudwatchLogs(value interface{}) {
	if err := a.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutCloudwatchMetric(value interface{}) {
	if err := a.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutDynamodb(value interface{}) {
	if err := a.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutDynamodbv2(value interface{}) {
	if err := a.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutElasticsearch(value interface{}) {
	if err := a.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutErrorAction(value *AwsIotTopicRule_ErrorActionProperty) {
	if err := a.validatePutErrorActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putErrorAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutFirehose(value interface{}) {
	if err := a.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFirehose",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutHttp(value interface{}) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutIotAnalytics(value interface{}) {
	if err := a.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutIotEvents(value interface{}) {
	if err := a.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutKafka(value interface{}) {
	if err := a.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKafka",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutKinesis(value interface{}) {
	if err := a.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKinesis",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutLambda(value interface{}) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutRepublish(value interface{}) {
	if err := a.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRepublish",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutS3(value interface{}) {
	if err := a.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutSns(value interface{}) {
	if err := a.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutSqs(value interface{}) {
	if err := a.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutStepFunctions(value interface{}) {
	if err := a.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) PutTimestream(value interface{}) {
	if err := a.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimestream",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetDynamodb() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		a,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		a,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetErrorAction() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetFirehose() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehose",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		a,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetIotEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetKafka() {
	_jsii_.InvokeVoid(
		a,
		"resetKafka",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetKinesis() {
	_jsii_.InvokeVoid(
		a,
		"resetKinesis",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetRepublish() {
	_jsii_.InvokeVoid(
		a,
		"resetRepublish",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetS3() {
	_jsii_.InvokeVoid(
		a,
		"resetS3",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetSns() {
	_jsii_.InvokeVoid(
		a,
		"resetSns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetSqs() {
	_jsii_.InvokeVoid(
		a,
		"resetSqs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		a,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) ResetTimestream() {
	_jsii_.InvokeVoid(
		a,
		"resetTimestream",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIotTopicRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIotTopicRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		a,
		"with",
		args,
		&returns,
	)

	return returns
}

