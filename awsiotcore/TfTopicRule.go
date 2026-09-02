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
type TfTopicRule interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CloudwatchAlarm() TfTopicRule_CloudwatchAlarmPropertyList
	// Experimental.
	CloudwatchAlarmInput() interface{}
	// Experimental.
	CloudwatchLogs() TfTopicRule_CloudwatchLogsPropertyList
	// Experimental.
	CloudwatchLogsInput() interface{}
	// Experimental.
	CloudwatchMetric() TfTopicRule_CloudwatchMetricPropertyList
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
	Dynamodb() TfTopicRule_DynamodbPropertyList
	// Experimental.
	DynamodbInput() interface{}
	// Experimental.
	Dynamodbv2() TfTopicRule_Dynamodbv2PropertyList
	// Experimental.
	Dynamodbv2Input() interface{}
	// Experimental.
	Elasticsearch() TfTopicRule_ElasticsearchPropertyList
	// Experimental.
	ElasticsearchInput() interface{}
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	ErrorAction() TfTopicRule_ErrorActionPropertyOutputReference
	// Experimental.
	ErrorActionInput() *TfTopicRule_ErrorActionProperty
	// Experimental.
	Firehose() TfTopicRule_FirehosePropertyList
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
	Http() TfTopicRule_HttpPropertyList
	// Experimental.
	HttpInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	IotAnalytics() TfTopicRule_IotAnalyticsPropertyList
	// Experimental.
	IotAnalyticsInput() interface{}
	// Experimental.
	IotEvents() TfTopicRule_IotEventsPropertyList
	// Experimental.
	IotEventsInput() interface{}
	// Experimental.
	Kafka() TfTopicRule_KafkaPropertyList
	// Experimental.
	KafkaInput() interface{}
	// Experimental.
	Kinesis() TfTopicRule_KinesisPropertyList
	// Experimental.
	KinesisInput() interface{}
	// Experimental.
	Lambda() TfTopicRule_LambdaPropertyList
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
	Republish() TfTopicRule_RepublishPropertyList
	// Experimental.
	RepublishInput() interface{}
	// Experimental.
	S3() TfTopicRule_S3PropertyList
	// Experimental.
	S3Input() interface{}
	// Experimental.
	Sns() TfTopicRule_SnsPropertyList
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
	Sqs() TfTopicRule_SqsPropertyList
	// Experimental.
	SqsInput() interface{}
	// Experimental.
	StepFunctions() TfTopicRule_StepFunctionsPropertyList
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
	Timestream() TfTopicRule_TimestreamPropertyList
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
	PutErrorAction(value *TfTopicRule_ErrorActionProperty)
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

// The jsii proxy struct for TfTopicRule
type jsiiProxy_TfTopicRule struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTopicRule) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchAlarm() TfTopicRule_CloudwatchAlarmPropertyList {
	var returns TfTopicRule_CloudwatchAlarmPropertyList
	_jsii_.Get(
		j,
		"cloudwatchAlarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchAlarmInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchAlarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchLogs() TfTopicRule_CloudwatchLogsPropertyList {
	var returns TfTopicRule_CloudwatchLogsPropertyList
	_jsii_.Get(
		j,
		"cloudwatchLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchLogsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchLogsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchMetric() TfTopicRule_CloudwatchMetricPropertyList {
	var returns TfTopicRule_CloudwatchMetricPropertyList
	_jsii_.Get(
		j,
		"cloudwatchMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) CloudwatchMetricInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cloudwatchMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Dynamodb() TfTopicRule_DynamodbPropertyList {
	var returns TfTopicRule_DynamodbPropertyList
	_jsii_.Get(
		j,
		"dynamodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) DynamodbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Dynamodbv2() TfTopicRule_Dynamodbv2PropertyList {
	var returns TfTopicRule_Dynamodbv2PropertyList
	_jsii_.Get(
		j,
		"dynamodbv2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Dynamodbv2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamodbv2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Elasticsearch() TfTopicRule_ElasticsearchPropertyList {
	var returns TfTopicRule_ElasticsearchPropertyList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) ErrorAction() TfTopicRule_ErrorActionPropertyOutputReference {
	var returns TfTopicRule_ErrorActionPropertyOutputReference
	_jsii_.Get(
		j,
		"errorAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) ErrorActionInput() *TfTopicRule_ErrorActionProperty {
	var returns *TfTopicRule_ErrorActionProperty
	_jsii_.Get(
		j,
		"errorActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Firehose() TfTopicRule_FirehosePropertyList {
	var returns TfTopicRule_FirehosePropertyList
	_jsii_.Get(
		j,
		"firehose",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) FirehoseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firehoseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Http() TfTopicRule_HttpPropertyList {
	var returns TfTopicRule_HttpPropertyList
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) IotAnalytics() TfTopicRule_IotAnalyticsPropertyList {
	var returns TfTopicRule_IotAnalyticsPropertyList
	_jsii_.Get(
		j,
		"iotAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) IotAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) IotEvents() TfTopicRule_IotEventsPropertyList {
	var returns TfTopicRule_IotEventsPropertyList
	_jsii_.Get(
		j,
		"iotEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) IotEventsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iotEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Kafka() TfTopicRule_KafkaPropertyList {
	var returns TfTopicRule_KafkaPropertyList
	_jsii_.Get(
		j,
		"kafka",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) KafkaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kafkaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Kinesis() TfTopicRule_KinesisPropertyList {
	var returns TfTopicRule_KinesisPropertyList
	_jsii_.Get(
		j,
		"kinesis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) KinesisInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kinesisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Lambda() TfTopicRule_LambdaPropertyList {
	var returns TfTopicRule_LambdaPropertyList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Republish() TfTopicRule_RepublishPropertyList {
	var returns TfTopicRule_RepublishPropertyList
	_jsii_.Get(
		j,
		"republish",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) RepublishInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"republishInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) S3() TfTopicRule_S3PropertyList {
	var returns TfTopicRule_S3PropertyList
	_jsii_.Get(
		j,
		"s3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) S3Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"s3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Sns() TfTopicRule_SnsPropertyList {
	var returns TfTopicRule_SnsPropertyList
	_jsii_.Get(
		j,
		"sns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) SnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Sql() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) SqlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) SqlVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) SqlVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Sqs() TfTopicRule_SqsPropertyList {
	var returns TfTopicRule_SqsPropertyList
	_jsii_.Get(
		j,
		"sqs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) SqsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) StepFunctions() TfTopicRule_StepFunctionsPropertyList {
	var returns TfTopicRule_StepFunctionsPropertyList
	_jsii_.Get(
		j,
		"stepFunctions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) StepFunctionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stepFunctionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) Timestream() TfTopicRule_TimestreamPropertyList {
	var returns TfTopicRule_TimestreamPropertyList
	_jsii_.Get(
		j,
		"timestream",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopicRule) TimestreamInput() interface{} {
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
func NewTfTopicRule(scope constructs.Construct, id *string, config *TfTopicRuleConfig) TfTopicRule {
	_init_.Initialize()

	if err := validateNewTfTopicRuleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTopicRule{}

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_topic_rule aws_iot_topic_rule} Resource.
// Experimental.
func NewTfTopicRule_Override(t TfTopicRule, scope constructs.Construct, id *string, config *TfTopicRuleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-iot-core.TfTopicRule",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTopicRule)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetSql(val *string) {
	if err := j.validateSetSqlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sql",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetSqlVersion(val *string) {
	if err := j.validateSetSqlVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlVersion",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTopicRule)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

// Generates CDKTN code for importing a TfTopicRule resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTopicRule_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTopicRule_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.TfTopicRule",
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
func TfTopicRule_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopicRule_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.TfTopicRule",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTopicRule_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopicRule_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.TfTopicRule",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTopicRule_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopicRule_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-iot-core.TfTopicRule",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTopicRule_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-iot-core.TfTopicRule",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTopicRule) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTopicRule) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTopicRule) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTopicRule) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTopicRule) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTopicRule) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTopicRule) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTopicRule) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTopicRule) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTopicRule) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTopicRule) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTopicRule) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopicRule) MarkWriteOnlyAttribute(value interface{}) interface{} {
	if err := t.validateMarkWriteOnlyAttributeParameters(value); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"markWriteOnlyAttribute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTopicRule) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTopicRule) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTopicRule) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTopicRule) PutCloudwatchAlarm(value interface{}) {
	if err := t.validatePutCloudwatchAlarmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchAlarm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutCloudwatchLogs(value interface{}) {
	if err := t.validatePutCloudwatchLogsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLogs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutCloudwatchMetric(value interface{}) {
	if err := t.validatePutCloudwatchMetricParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchMetric",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutDynamodb(value interface{}) {
	if err := t.validatePutDynamodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodb",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutDynamodbv2(value interface{}) {
	if err := t.validatePutDynamodbv2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDynamodbv2",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutElasticsearch(value interface{}) {
	if err := t.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutErrorAction(value *TfTopicRule_ErrorActionProperty) {
	if err := t.validatePutErrorActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putErrorAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutFirehose(value interface{}) {
	if err := t.validatePutFirehoseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFirehose",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutHttp(value interface{}) {
	if err := t.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutIotAnalytics(value interface{}) {
	if err := t.validatePutIotAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIotAnalytics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutIotEvents(value interface{}) {
	if err := t.validatePutIotEventsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIotEvents",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutKafka(value interface{}) {
	if err := t.validatePutKafkaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKafka",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutKinesis(value interface{}) {
	if err := t.validatePutKinesisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesis",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutLambda(value interface{}) {
	if err := t.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLambda",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutRepublish(value interface{}) {
	if err := t.validatePutRepublishParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRepublish",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutS3(value interface{}) {
	if err := t.validatePutS3Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutSns(value interface{}) {
	if err := t.validatePutSnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutSqs(value interface{}) {
	if err := t.validatePutSqsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSqs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutStepFunctions(value interface{}) {
	if err := t.validatePutStepFunctionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStepFunctions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) PutTimestream(value interface{}) {
	if err := t.validatePutTimestreamParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimestream",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTopicRule) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTopicRule) ResetCloudwatchAlarm() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchAlarm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetCloudwatchLogs() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLogs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetCloudwatchMetric() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchMetric",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetDynamodb() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetDynamodbv2() {
	_jsii_.InvokeVoid(
		t,
		"resetDynamodbv2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		t,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetErrorAction() {
	_jsii_.InvokeVoid(
		t,
		"resetErrorAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetFirehose() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehose",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetHttp() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetIotAnalytics() {
	_jsii_.InvokeVoid(
		t,
		"resetIotAnalytics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetIotEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetIotEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetKafka() {
	_jsii_.InvokeVoid(
		t,
		"resetKafka",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetKinesis() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesis",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetLambda() {
	_jsii_.InvokeVoid(
		t,
		"resetLambda",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetRepublish() {
	_jsii_.InvokeVoid(
		t,
		"resetRepublish",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetS3() {
	_jsii_.InvokeVoid(
		t,
		"resetS3",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetSns() {
	_jsii_.InvokeVoid(
		t,
		"resetSns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetSqs() {
	_jsii_.InvokeVoid(
		t,
		"resetSqs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetStepFunctions() {
	_jsii_.InvokeVoid(
		t,
		"resetStepFunctions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) ResetTimestream() {
	_jsii_.InvokeVoid(
		t,
		"resetTimestream",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopicRule) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopicRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		t,
		"with",
		args,
		&returns,
	)

	return returns
}

