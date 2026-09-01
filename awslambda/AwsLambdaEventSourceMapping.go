package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslambda/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslambda/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping}.
// Experimental.
type AwsLambdaEventSourceMapping interface {
	cdktn.TerraformResource
	// Experimental.
	AmazonManagedKafkaEventSourceConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference
	// Experimental.
	AmazonManagedKafkaEventSourceConfigInput() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
	// Experimental.
	Arn() *string
	// Experimental.
	BatchSize() *float64
	// Experimental.
	SetBatchSize(val *float64)
	// Experimental.
	BatchSizeInput() *float64
	// Experimental.
	BisectBatchOnFunctionError() interface{}
	// Experimental.
	SetBisectBatchOnFunctionError(val interface{})
	// Experimental.
	BisectBatchOnFunctionErrorInput() interface{}
	// Experimental.
	CdktfStack() cdktn.TerraformStack
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
	DestinationConfig() AwsLambdaEventSourceMapping_DestinationConfigPropertyOutputReference
	// Experimental.
	DestinationConfigInput() *AwsLambdaEventSourceMapping_DestinationConfigProperty
	// Experimental.
	DocumentDbEventSourceConfig() AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference
	// Experimental.
	DocumentDbEventSourceConfigInput() *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	EventSourceArn() *string
	// Experimental.
	SetEventSourceArn(val *string)
	// Experimental.
	EventSourceArnInput() *string
	// Experimental.
	FilterCriteria() AwsLambdaEventSourceMapping_FilterCriteriaPropertyOutputReference
	// Experimental.
	FilterCriteriaInput() *AwsLambdaEventSourceMapping_FilterCriteriaProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FunctionArn() *string
	// Experimental.
	FunctionName() *string
	// Experimental.
	SetFunctionName(val *string)
	// Experimental.
	FunctionNameInput() *string
	// Experimental.
	FunctionResponseTypes() *[]*string
	// Experimental.
	SetFunctionResponseTypes(val *[]*string)
	// Experimental.
	FunctionResponseTypesInput() *[]*string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	LastModified() *string
	// Experimental.
	LastProcessingResult() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaximumBatchingWindowInSeconds() *float64
	// Experimental.
	SetMaximumBatchingWindowInSeconds(val *float64)
	// Experimental.
	MaximumBatchingWindowInSecondsInput() *float64
	// Experimental.
	MaximumRecordAgeInSeconds() *float64
	// Experimental.
	SetMaximumRecordAgeInSeconds(val *float64)
	// Experimental.
	MaximumRecordAgeInSecondsInput() *float64
	// Experimental.
	MaximumRetryAttempts() *float64
	// Experimental.
	SetMaximumRetryAttempts(val *float64)
	// Experimental.
	MaximumRetryAttemptsInput() *float64
	// Experimental.
	MetricsConfig() AwsLambdaEventSourceMapping_MetricsConfigPropertyOutputReference
	// Experimental.
	MetricsConfigInput() *AwsLambdaEventSourceMapping_MetricsConfigProperty
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ParallelizationFactor() *float64
	// Experimental.
	SetParallelizationFactor(val *float64)
	// Experimental.
	ParallelizationFactorInput() *float64
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	ProvisionedPollerConfig() AwsLambdaEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference
	// Experimental.
	ProvisionedPollerConfigInput() *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	Queues() *[]*string
	// Experimental.
	SetQueues(val *[]*string)
	// Experimental.
	QueuesInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	ScalingConfig() AwsLambdaEventSourceMapping_ScalingConfigPropertyOutputReference
	// Experimental.
	ScalingConfigInput() *AwsLambdaEventSourceMapping_ScalingConfigProperty
	// Experimental.
	SelfManagedEventSource() AwsLambdaEventSourceMapping_SelfManagedEventSourcePropertyOutputReference
	// Experimental.
	SelfManagedEventSourceInput() *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty
	// Experimental.
	SelfManagedKafkaEventSourceConfig() AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference
	// Experimental.
	SelfManagedKafkaEventSourceConfigInput() *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty
	// Experimental.
	SourceAccessConfiguration() AwsLambdaEventSourceMapping_SourceAccessConfigurationPropertyList
	// Experimental.
	SourceAccessConfigurationInput() interface{}
	// Experimental.
	StartingPosition() *string
	// Experimental.
	SetStartingPosition(val *string)
	// Experimental.
	StartingPositionInput() *string
	// Experimental.
	StartingPositionTimestamp() *string
	// Experimental.
	SetStartingPositionTimestamp(val *string)
	// Experimental.
	StartingPositionTimestampInput() *string
	// Experimental.
	State() *string
	// Experimental.
	StateTransitionReason() *string
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
	Timeouts() AwsLambdaEventSourceMapping_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Topics() *[]*string
	// Experimental.
	SetTopics(val *[]*string)
	// Experimental.
	TopicsInput() *[]*string
	// Experimental.
	TumblingWindowInSeconds() *float64
	// Experimental.
	SetTumblingWindowInSeconds(val *float64)
	// Experimental.
	TumblingWindowInSecondsInput() *float64
	// Experimental.
	UseResourceTimeoutForPropagation() interface{}
	// Experimental.
	SetUseResourceTimeoutForPropagation(val interface{})
	// Experimental.
	UseResourceTimeoutForPropagationInput() interface{}
	// Experimental.
	Uuid() *string
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
	PutAmazonManagedKafkaEventSourceConfig(value *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty)
	// Experimental.
	PutDestinationConfig(value *AwsLambdaEventSourceMapping_DestinationConfigProperty)
	// Experimental.
	PutDocumentDbEventSourceConfig(value *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty)
	// Experimental.
	PutFilterCriteria(value *AwsLambdaEventSourceMapping_FilterCriteriaProperty)
	// Experimental.
	PutMetricsConfig(value *AwsLambdaEventSourceMapping_MetricsConfigProperty)
	// Experimental.
	PutProvisionedPollerConfig(value *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty)
	// Experimental.
	PutScalingConfig(value *AwsLambdaEventSourceMapping_ScalingConfigProperty)
	// Experimental.
	PutSelfManagedEventSource(value *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty)
	// Experimental.
	PutSelfManagedKafkaEventSourceConfig(value *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty)
	// Experimental.
	PutSourceAccessConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *AwsLambdaEventSourceMapping_TimeoutsProperty)
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
	ResetAmazonManagedKafkaEventSourceConfig()
	// Experimental.
	ResetBatchSize()
	// Experimental.
	ResetBisectBatchOnFunctionError()
	// Experimental.
	ResetDestinationConfig()
	// Experimental.
	ResetDocumentDbEventSourceConfig()
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetEventSourceArn()
	// Experimental.
	ResetFilterCriteria()
	// Experimental.
	ResetFunctionResponseTypes()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetMaximumBatchingWindowInSeconds()
	// Experimental.
	ResetMaximumRecordAgeInSeconds()
	// Experimental.
	ResetMaximumRetryAttempts()
	// Experimental.
	ResetMetricsConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParallelizationFactor()
	// Experimental.
	ResetProvisionedPollerConfig()
	// Experimental.
	ResetQueues()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetScalingConfig()
	// Experimental.
	ResetSelfManagedEventSource()
	// Experimental.
	ResetSelfManagedKafkaEventSourceConfig()
	// Experimental.
	ResetSourceAccessConfiguration()
	// Experimental.
	ResetStartingPosition()
	// Experimental.
	ResetStartingPositionTimestamp()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTopics()
	// Experimental.
	ResetTumblingWindowInSeconds()
	// Experimental.
	ResetUseResourceTimeoutForPropagation()
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

// The jsii proxy struct for AwsLambdaEventSourceMapping
type jsiiProxy_AwsLambdaEventSourceMapping struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) AmazonManagedKafkaEventSourceConfig() AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) AmazonManagedKafkaEventSourceConfigInput() *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty {
	var returns *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) BisectBatchOnFunctionErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) DestinationConfig() AwsLambdaEventSourceMapping_DestinationConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_DestinationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) DestinationConfigInput() *AwsLambdaEventSourceMapping_DestinationConfigProperty {
	var returns *AwsLambdaEventSourceMapping_DestinationConfigProperty
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) DocumentDbEventSourceConfig() AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"documentDbEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) DocumentDbEventSourceConfigInput() *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty {
	var returns *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty
	_jsii_.Get(
		j,
		"documentDbEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) EventSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FilterCriteria() AwsLambdaEventSourceMapping_FilterCriteriaPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_FilterCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FilterCriteriaInput() *AwsLambdaEventSourceMapping_FilterCriteriaProperty {
	var returns *AwsLambdaEventSourceMapping_FilterCriteriaProperty
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) FunctionResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) LastProcessingResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProcessingResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MetricsConfig() AwsLambdaEventSourceMapping_MetricsConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_MetricsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"metricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) MetricsConfigInput() *AwsLambdaEventSourceMapping_MetricsConfigProperty {
	var returns *AwsLambdaEventSourceMapping_MetricsConfigProperty
	_jsii_.Get(
		j,
		"metricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ProvisionedPollerConfig() AwsLambdaEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"provisionedPollerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ProvisionedPollerConfigInput() *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty {
	var returns *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty
	_jsii_.Get(
		j,
		"provisionedPollerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) QueuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ScalingConfig() AwsLambdaEventSourceMapping_ScalingConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_ScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) ScalingConfigInput() *AwsLambdaEventSourceMapping_ScalingConfigProperty {
	var returns *AwsLambdaEventSourceMapping_ScalingConfigProperty
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SelfManagedEventSource() AwsLambdaEventSourceMapping_SelfManagedEventSourcePropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_SelfManagedEventSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SelfManagedEventSourceInput() *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty {
	var returns *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty
	_jsii_.Get(
		j,
		"selfManagedEventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SelfManagedKafkaEventSourceConfig() AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SelfManagedKafkaEventSourceConfigInput() *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty {
	var returns *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SourceAccessConfiguration() AwsLambdaEventSourceMapping_SourceAccessConfigurationPropertyList {
	var returns AwsLambdaEventSourceMapping_SourceAccessConfigurationPropertyList
	_jsii_.Get(
		j,
		"sourceAccessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) SourceAccessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) StartingPositionTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) StartingPositionTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) StateTransitionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateTransitionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Timeouts() AwsLambdaEventSourceMapping_TimeoutsPropertyOutputReference {
	var returns AwsLambdaEventSourceMapping_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) TumblingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) UseResourceTimeoutForPropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) UseResourceTimeoutForPropagationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
// Experimental.
func NewAwsLambdaEventSourceMapping(scope constructs.Construct, id *string, config *AwsLambdaEventSourceMappingConfig) AwsLambdaEventSourceMapping {
	_init_.Initialize()

	if err := validateNewAwsLambdaEventSourceMappingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLambdaEventSourceMapping{}

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
// Experimental.
func NewAwsLambdaEventSourceMapping_Override(a AwsLambdaEventSourceMapping, scope constructs.Construct, id *string, config *AwsLambdaEventSourceMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetBisectBatchOnFunctionError(val interface{}) {
	if err := j.validateSetBisectBatchOnFunctionErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetEventSourceArn(val *string) {
	if err := j.validateSetEventSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetFunctionResponseTypes(val *[]*string) {
	if err := j.validateSetFunctionResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetQueues(val *[]*string) {
	if err := j.validateSetQueuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetStartingPositionTimestamp(val *string) {
	if err := j.validateSetStartingPositionTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetTopics(val *[]*string) {
	if err := j.validateSetTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetTumblingWindowInSeconds(val *float64) {
	if err := j.validateSetTumblingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsLambdaEventSourceMapping)SetUseResourceTimeoutForPropagation(val interface{}) {
	if err := j.validateSetUseResourceTimeoutForPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourceTimeoutForPropagation",
		val,
	)
}

// Generates CDKTN code for importing a AwsLambdaEventSourceMapping resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsLambdaEventSourceMapping_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsLambdaEventSourceMapping_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
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
func AwsLambdaEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaEventSourceMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLambdaEventSourceMapping_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaEventSourceMapping_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsLambdaEventSourceMapping_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsLambdaEventSourceMapping_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsLambdaEventSourceMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lambda.AwsLambdaEventSourceMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsLambdaEventSourceMapping) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutAmazonManagedKafkaEventSourceConfig(value *AwsLambdaEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty) {
	if err := a.validatePutAmazonManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAmazonManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutDestinationConfig(value *AwsLambdaEventSourceMapping_DestinationConfigProperty) {
	if err := a.validatePutDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutDocumentDbEventSourceConfig(value *AwsLambdaEventSourceMapping_DocumentDbEventSourceConfigProperty) {
	if err := a.validatePutDocumentDbEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDocumentDbEventSourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutFilterCriteria(value *AwsLambdaEventSourceMapping_FilterCriteriaProperty) {
	if err := a.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutMetricsConfig(value *AwsLambdaEventSourceMapping_MetricsConfigProperty) {
	if err := a.validatePutMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutProvisionedPollerConfig(value *AwsLambdaEventSourceMapping_ProvisionedPollerConfigProperty) {
	if err := a.validatePutProvisionedPollerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProvisionedPollerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutScalingConfig(value *AwsLambdaEventSourceMapping_ScalingConfigProperty) {
	if err := a.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutSelfManagedEventSource(value *AwsLambdaEventSourceMapping_SelfManagedEventSourceProperty) {
	if err := a.validatePutSelfManagedEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedEventSource",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutSelfManagedKafkaEventSourceConfig(value *AwsLambdaEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty) {
	if err := a.validatePutSelfManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutSourceAccessConfiguration(value interface{}) {
	if err := a.validatePutSourceAccessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSourceAccessConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) PutTimeouts(value *AwsLambdaEventSourceMapping_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetAmazonManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAmazonManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetBatchSize() {
	_jsii_.InvokeVoid(
		a,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetBisectBatchOnFunctionError() {
	_jsii_.InvokeVoid(
		a,
		"resetBisectBatchOnFunctionError",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetDocumentDbEventSourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDocumentDbEventSourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetEventSourceArn() {
	_jsii_.InvokeVoid(
		a,
		"resetEventSourceArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetFunctionResponseTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetFunctionResponseTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetMetricsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		a,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetProvisionedPollerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetProvisionedPollerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetQueues() {
	_jsii_.InvokeVoid(
		a,
		"resetQueues",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetSelfManagedEventSource() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedEventSource",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetSelfManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetSourceAccessConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceAccessConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetStartingPositionTimestamp() {
	_jsii_.InvokeVoid(
		a,
		"resetStartingPositionTimestamp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetTopics() {
	_jsii_.InvokeVoid(
		a,
		"resetTopics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetTumblingWindowInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTumblingWindowInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ResetUseResourceTimeoutForPropagation() {
	_jsii_.InvokeVoid(
		a,
		"resetUseResourceTimeoutForPropagation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLambdaEventSourceMapping) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

