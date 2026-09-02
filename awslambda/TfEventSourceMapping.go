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
type TfEventSourceMapping interface {
	cdktn.TerraformResource
	// Experimental.
	AmazonManagedKafkaEventSourceConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference
	// Experimental.
	AmazonManagedKafkaEventSourceConfigInput() *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
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
	DestinationConfig() TfEventSourceMapping_DestinationConfigPropertyOutputReference
	// Experimental.
	DestinationConfigInput() *TfEventSourceMapping_DestinationConfigProperty
	// Experimental.
	DocumentDbEventSourceConfig() TfEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference
	// Experimental.
	DocumentDbEventSourceConfigInput() *TfEventSourceMapping_DocumentDbEventSourceConfigProperty
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
	FilterCriteria() TfEventSourceMapping_FilterCriteriaPropertyOutputReference
	// Experimental.
	FilterCriteriaInput() *TfEventSourceMapping_FilterCriteriaProperty
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
	MetricsConfig() TfEventSourceMapping_MetricsConfigPropertyOutputReference
	// Experimental.
	MetricsConfigInput() *TfEventSourceMapping_MetricsConfigProperty
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
	ProvisionedPollerConfig() TfEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference
	// Experimental.
	ProvisionedPollerConfigInput() *TfEventSourceMapping_ProvisionedPollerConfigProperty
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
	ScalingConfig() TfEventSourceMapping_ScalingConfigPropertyOutputReference
	// Experimental.
	ScalingConfigInput() *TfEventSourceMapping_ScalingConfigProperty
	// Experimental.
	SelfManagedEventSource() TfEventSourceMapping_SelfManagedEventSourcePropertyOutputReference
	// Experimental.
	SelfManagedEventSourceInput() *TfEventSourceMapping_SelfManagedEventSourceProperty
	// Experimental.
	SelfManagedKafkaEventSourceConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference
	// Experimental.
	SelfManagedKafkaEventSourceConfigInput() *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty
	// Experimental.
	SourceAccessConfiguration() TfEventSourceMapping_SourceAccessConfigurationPropertyList
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
	Timeouts() TfEventSourceMapping_TimeoutsPropertyOutputReference
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
	PutAmazonManagedKafkaEventSourceConfig(value *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty)
	// Experimental.
	PutDestinationConfig(value *TfEventSourceMapping_DestinationConfigProperty)
	// Experimental.
	PutDocumentDbEventSourceConfig(value *TfEventSourceMapping_DocumentDbEventSourceConfigProperty)
	// Experimental.
	PutFilterCriteria(value *TfEventSourceMapping_FilterCriteriaProperty)
	// Experimental.
	PutMetricsConfig(value *TfEventSourceMapping_MetricsConfigProperty)
	// Experimental.
	PutProvisionedPollerConfig(value *TfEventSourceMapping_ProvisionedPollerConfigProperty)
	// Experimental.
	PutScalingConfig(value *TfEventSourceMapping_ScalingConfigProperty)
	// Experimental.
	PutSelfManagedEventSource(value *TfEventSourceMapping_SelfManagedEventSourceProperty)
	// Experimental.
	PutSelfManagedKafkaEventSourceConfig(value *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty)
	// Experimental.
	PutSourceAccessConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *TfEventSourceMapping_TimeoutsProperty)
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

// The jsii proxy struct for TfEventSourceMapping
type jsiiProxy_TfEventSourceMapping struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfEventSourceMapping) AmazonManagedKafkaEventSourceConfig() TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference {
	var returns TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) AmazonManagedKafkaEventSourceConfigInput() *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty {
	var returns *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty
	_jsii_.Get(
		j,
		"amazonManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) BatchSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) BatchSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"batchSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) BisectBatchOnFunctionError() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionError",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) BisectBatchOnFunctionErrorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bisectBatchOnFunctionErrorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) DestinationConfig() TfEventSourceMapping_DestinationConfigPropertyOutputReference {
	var returns TfEventSourceMapping_DestinationConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"destinationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) DestinationConfigInput() *TfEventSourceMapping_DestinationConfigProperty {
	var returns *TfEventSourceMapping_DestinationConfigProperty
	_jsii_.Get(
		j,
		"destinationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) DocumentDbEventSourceConfig() TfEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference {
	var returns TfEventSourceMapping_DocumentDbEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"documentDbEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) DocumentDbEventSourceConfigInput() *TfEventSourceMapping_DocumentDbEventSourceConfigProperty {
	var returns *TfEventSourceMapping_DocumentDbEventSourceConfigProperty
	_jsii_.Get(
		j,
		"documentDbEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) EventSourceArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) EventSourceArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventSourceArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FilterCriteria() TfEventSourceMapping_FilterCriteriaPropertyOutputReference {
	var returns TfEventSourceMapping_FilterCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"filterCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FilterCriteriaInput() *TfEventSourceMapping_FilterCriteriaProperty {
	var returns *TfEventSourceMapping_FilterCriteriaProperty
	_jsii_.Get(
		j,
		"filterCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FunctionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FunctionResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) FunctionResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"functionResponseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) LastModified() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastModified",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) LastProcessingResult() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastProcessingResult",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumBatchingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumBatchingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumBatchingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumRecordAgeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumRecordAgeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRecordAgeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumRetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MaximumRetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MetricsConfig() TfEventSourceMapping_MetricsConfigPropertyOutputReference {
	var returns TfEventSourceMapping_MetricsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"metricsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) MetricsConfigInput() *TfEventSourceMapping_MetricsConfigProperty {
	var returns *TfEventSourceMapping_MetricsConfigProperty
	_jsii_.Get(
		j,
		"metricsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ParallelizationFactor() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ParallelizationFactorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"parallelizationFactorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ProvisionedPollerConfig() TfEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference {
	var returns TfEventSourceMapping_ProvisionedPollerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"provisionedPollerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ProvisionedPollerConfigInput() *TfEventSourceMapping_ProvisionedPollerConfigProperty {
	var returns *TfEventSourceMapping_ProvisionedPollerConfigProperty
	_jsii_.Get(
		j,
		"provisionedPollerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Queues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) QueuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ScalingConfig() TfEventSourceMapping_ScalingConfigPropertyOutputReference {
	var returns TfEventSourceMapping_ScalingConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scalingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) ScalingConfigInput() *TfEventSourceMapping_ScalingConfigProperty {
	var returns *TfEventSourceMapping_ScalingConfigProperty
	_jsii_.Get(
		j,
		"scalingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SelfManagedEventSource() TfEventSourceMapping_SelfManagedEventSourcePropertyOutputReference {
	var returns TfEventSourceMapping_SelfManagedEventSourcePropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedEventSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SelfManagedEventSourceInput() *TfEventSourceMapping_SelfManagedEventSourceProperty {
	var returns *TfEventSourceMapping_SelfManagedEventSourceProperty
	_jsii_.Get(
		j,
		"selfManagedEventSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SelfManagedKafkaEventSourceConfig() TfEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference {
	var returns TfEventSourceMapping_SelfManagedKafkaEventSourceConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SelfManagedKafkaEventSourceConfigInput() *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty {
	var returns *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty
	_jsii_.Get(
		j,
		"selfManagedKafkaEventSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SourceAccessConfiguration() TfEventSourceMapping_SourceAccessConfigurationPropertyList {
	var returns TfEventSourceMapping_SourceAccessConfigurationPropertyList
	_jsii_.Get(
		j,
		"sourceAccessConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) SourceAccessConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sourceAccessConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) StartingPosition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPosition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) StartingPositionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) StartingPositionTimestamp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestamp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) StartingPositionTimestampInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startingPositionTimestampInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) StateTransitionReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateTransitionReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Timeouts() TfEventSourceMapping_TimeoutsPropertyOutputReference {
	var returns TfEventSourceMapping_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TumblingWindowInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) TumblingWindowInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tumblingWindowInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) UseResourceTimeoutForPropagation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) UseResourceTimeoutForPropagationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useResourceTimeoutForPropagationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEventSourceMapping) Uuid() *string {
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
func NewTfEventSourceMapping(scope constructs.Construct, id *string, config *TfEventSourceMappingConfig) TfEventSourceMapping {
	_init_.Initialize()

	if err := validateNewTfEventSourceMappingParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEventSourceMapping{}

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping aws_lambda_event_source_mapping} Resource.
// Experimental.
func NewTfEventSourceMapping_Override(t TfEventSourceMapping, scope constructs.Construct, id *string, config *TfEventSourceMappingConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetBatchSize(val *float64) {
	if err := j.validateSetBatchSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"batchSize",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetBisectBatchOnFunctionError(val interface{}) {
	if err := j.validateSetBisectBatchOnFunctionErrorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bisectBatchOnFunctionError",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetEventSourceArn(val *string) {
	if err := j.validateSetEventSourceArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventSourceArn",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetFunctionName(val *string) {
	if err := j.validateSetFunctionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionName",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetFunctionResponseTypes(val *[]*string) {
	if err := j.validateSetFunctionResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"functionResponseTypes",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetMaximumBatchingWindowInSeconds(val *float64) {
	if err := j.validateSetMaximumBatchingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumBatchingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetMaximumRecordAgeInSeconds(val *float64) {
	if err := j.validateSetMaximumRecordAgeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRecordAgeInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetMaximumRetryAttempts(val *float64) {
	if err := j.validateSetMaximumRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryAttempts",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetParallelizationFactor(val *float64) {
	if err := j.validateSetParallelizationFactorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelizationFactor",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetQueues(val *[]*string) {
	if err := j.validateSetQueuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queues",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetStartingPosition(val *string) {
	if err := j.validateSetStartingPositionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPosition",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetStartingPositionTimestamp(val *string) {
	if err := j.validateSetStartingPositionTimestampParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startingPositionTimestamp",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetTopics(val *[]*string) {
	if err := j.validateSetTopicsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetTumblingWindowInSeconds(val *float64) {
	if err := j.validateSetTumblingWindowInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tumblingWindowInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEventSourceMapping)SetUseResourceTimeoutForPropagation(val interface{}) {
	if err := j.validateSetUseResourceTimeoutForPropagationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useResourceTimeoutForPropagation",
		val,
	)
}

// Generates CDKTN code for importing a TfEventSourceMapping resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfEventSourceMapping_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfEventSourceMapping_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfEventSourceMapping",
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
func TfEventSourceMapping_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEventSourceMapping_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEventSourceMapping_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEventSourceMapping_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfEventSourceMapping_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfEventSourceMapping_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfEventSourceMapping_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lambda.TfEventSourceMapping",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEventSourceMapping) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEventSourceMapping) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEventSourceMapping) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEventSourceMapping) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEventSourceMapping) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEventSourceMapping) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEventSourceMapping) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfEventSourceMapping) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutAmazonManagedKafkaEventSourceConfig(value *TfEventSourceMapping_AmazonManagedKafkaEventSourceConfigProperty) {
	if err := t.validatePutAmazonManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAmazonManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutDestinationConfig(value *TfEventSourceMapping_DestinationConfigProperty) {
	if err := t.validatePutDestinationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDestinationConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutDocumentDbEventSourceConfig(value *TfEventSourceMapping_DocumentDbEventSourceConfigProperty) {
	if err := t.validatePutDocumentDbEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDocumentDbEventSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutFilterCriteria(value *TfEventSourceMapping_FilterCriteriaProperty) {
	if err := t.validatePutFilterCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilterCriteria",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutMetricsConfig(value *TfEventSourceMapping_MetricsConfigProperty) {
	if err := t.validatePutMetricsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMetricsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutProvisionedPollerConfig(value *TfEventSourceMapping_ProvisionedPollerConfigProperty) {
	if err := t.validatePutProvisionedPollerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProvisionedPollerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutScalingConfig(value *TfEventSourceMapping_ScalingConfigProperty) {
	if err := t.validatePutScalingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScalingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutSelfManagedEventSource(value *TfEventSourceMapping_SelfManagedEventSourceProperty) {
	if err := t.validatePutSelfManagedEventSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfManagedEventSource",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutSelfManagedKafkaEventSourceConfig(value *TfEventSourceMapping_SelfManagedKafkaEventSourceConfigProperty) {
	if err := t.validatePutSelfManagedKafkaEventSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfManagedKafkaEventSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutSourceAccessConfiguration(value interface{}) {
	if err := t.validatePutSourceAccessConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSourceAccessConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) PutTimeouts(value *TfEventSourceMapping_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetAmazonManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAmazonManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetBatchSize() {
	_jsii_.InvokeVoid(
		t,
		"resetBatchSize",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetBisectBatchOnFunctionError() {
	_jsii_.InvokeVoid(
		t,
		"resetBisectBatchOnFunctionError",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetDestinationConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDestinationConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetDocumentDbEventSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDocumentDbEventSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetEventSourceArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEventSourceArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetFilterCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetFunctionResponseTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetFunctionResponseTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetMaximumBatchingWindowInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumBatchingWindowInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetMaximumRecordAgeInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRecordAgeInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetMaximumRetryAttempts() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumRetryAttempts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetMetricsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetMetricsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetParallelizationFactor() {
	_jsii_.InvokeVoid(
		t,
		"resetParallelizationFactor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetProvisionedPollerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetProvisionedPollerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetQueues() {
	_jsii_.InvokeVoid(
		t,
		"resetQueues",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetScalingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetScalingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetSelfManagedEventSource() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedEventSource",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetSelfManagedKafkaEventSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedKafkaEventSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetSourceAccessConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceAccessConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetStartingPosition() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPosition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetStartingPositionTimestamp() {
	_jsii_.InvokeVoid(
		t,
		"resetStartingPositionTimestamp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetTopics() {
	_jsii_.InvokeVoid(
		t,
		"resetTopics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetTumblingWindowInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTumblingWindowInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) ResetUseResourceTimeoutForPropagation() {
	_jsii_.InvokeVoid(
		t,
		"resetUseResourceTimeoutForPropagation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEventSourceMapping) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEventSourceMapping) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

