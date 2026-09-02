package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssns/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssns/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic}.
// Experimental.
type TfTopic interface {
	cdktn.TerraformResource
	// Experimental.
	ApplicationFailureFeedbackRoleArn() *string
	// Experimental.
	SetApplicationFailureFeedbackRoleArn(val *string)
	// Experimental.
	ApplicationFailureFeedbackRoleArnInput() *string
	// Experimental.
	ApplicationSuccessFeedbackRoleArn() *string
	// Experimental.
	SetApplicationSuccessFeedbackRoleArn(val *string)
	// Experimental.
	ApplicationSuccessFeedbackRoleArnInput() *string
	// Experimental.
	ApplicationSuccessFeedbackSampleRate() *float64
	// Experimental.
	SetApplicationSuccessFeedbackSampleRate(val *float64)
	// Experimental.
	ApplicationSuccessFeedbackSampleRateInput() *float64
	// Experimental.
	ArchivePolicy() *string
	// Experimental.
	SetArchivePolicy(val *string)
	// Experimental.
	ArchivePolicyInput() *string
	// Experimental.
	Arn() *string
	// Experimental.
	BeginningArchiveTime() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	ContentBasedDeduplication() interface{}
	// Experimental.
	SetContentBasedDeduplication(val interface{})
	// Experimental.
	ContentBasedDeduplicationInput() interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DeliveryPolicy() *string
	// Experimental.
	SetDeliveryPolicy(val *string)
	// Experimental.
	DeliveryPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	DisplayName() *string
	// Experimental.
	SetDisplayName(val *string)
	// Experimental.
	DisplayNameInput() *string
	// Experimental.
	FifoThroughputScope() *string
	// Experimental.
	SetFifoThroughputScope(val *string)
	// Experimental.
	FifoThroughputScopeInput() *string
	// Experimental.
	FifoTopic() interface{}
	// Experimental.
	SetFifoTopic(val interface{})
	// Experimental.
	FifoTopicInput() interface{}
	// Experimental.
	FirehoseFailureFeedbackRoleArn() *string
	// Experimental.
	SetFirehoseFailureFeedbackRoleArn(val *string)
	// Experimental.
	FirehoseFailureFeedbackRoleArnInput() *string
	// Experimental.
	FirehoseSuccessFeedbackRoleArn() *string
	// Experimental.
	SetFirehoseSuccessFeedbackRoleArn(val *string)
	// Experimental.
	FirehoseSuccessFeedbackRoleArnInput() *string
	// Experimental.
	FirehoseSuccessFeedbackSampleRate() *float64
	// Experimental.
	SetFirehoseSuccessFeedbackSampleRate(val *float64)
	// Experimental.
	FirehoseSuccessFeedbackSampleRateInput() *float64
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HttpFailureFeedbackRoleArn() *string
	// Experimental.
	SetHttpFailureFeedbackRoleArn(val *string)
	// Experimental.
	HttpFailureFeedbackRoleArnInput() *string
	// Experimental.
	HttpSuccessFeedbackRoleArn() *string
	// Experimental.
	SetHttpSuccessFeedbackRoleArn(val *string)
	// Experimental.
	HttpSuccessFeedbackRoleArnInput() *string
	// Experimental.
	HttpSuccessFeedbackSampleRate() *float64
	// Experimental.
	SetHttpSuccessFeedbackSampleRate(val *float64)
	// Experimental.
	HttpSuccessFeedbackSampleRateInput() *float64
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsMasterKeyId() *string
	// Experimental.
	SetKmsMasterKeyId(val *string)
	// Experimental.
	KmsMasterKeyIdInput() *string
	// Experimental.
	LambdaFailureFeedbackRoleArn() *string
	// Experimental.
	SetLambdaFailureFeedbackRoleArn(val *string)
	// Experimental.
	LambdaFailureFeedbackRoleArnInput() *string
	// Experimental.
	LambdaSuccessFeedbackRoleArn() *string
	// Experimental.
	SetLambdaSuccessFeedbackRoleArn(val *string)
	// Experimental.
	LambdaSuccessFeedbackRoleArnInput() *string
	// Experimental.
	LambdaSuccessFeedbackSampleRate() *float64
	// Experimental.
	SetLambdaSuccessFeedbackSampleRate(val *float64)
	// Experimental.
	LambdaSuccessFeedbackSampleRateInput() *float64
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
	// Experimental.
	NamePrefix() *string
	// Experimental.
	SetNamePrefix(val *string)
	// Experimental.
	NamePrefixInput() *string
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	Owner() *string
	// Experimental.
	Policy() *string
	// Experimental.
	SetPolicy(val *string)
	// Experimental.
	PolicyInput() *string
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
	SignatureVersion() *float64
	// Experimental.
	SetSignatureVersion(val *float64)
	// Experimental.
	SignatureVersionInput() *float64
	// Experimental.
	SqsFailureFeedbackRoleArn() *string
	// Experimental.
	SetSqsFailureFeedbackRoleArn(val *string)
	// Experimental.
	SqsFailureFeedbackRoleArnInput() *string
	// Experimental.
	SqsSuccessFeedbackRoleArn() *string
	// Experimental.
	SetSqsSuccessFeedbackRoleArn(val *string)
	// Experimental.
	SqsSuccessFeedbackRoleArnInput() *string
	// Experimental.
	SqsSuccessFeedbackSampleRate() *float64
	// Experimental.
	SetSqsSuccessFeedbackSampleRate(val *float64)
	// Experimental.
	SqsSuccessFeedbackSampleRateInput() *float64
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
	TracingConfig() *string
	// Experimental.
	SetTracingConfig(val *string)
	// Experimental.
	TracingConfigInput() *string
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
	ResetApplicationFailureFeedbackRoleArn()
	// Experimental.
	ResetApplicationSuccessFeedbackRoleArn()
	// Experimental.
	ResetApplicationSuccessFeedbackSampleRate()
	// Experimental.
	ResetArchivePolicy()
	// Experimental.
	ResetContentBasedDeduplication()
	// Experimental.
	ResetDeliveryPolicy()
	// Experimental.
	ResetDisplayName()
	// Experimental.
	ResetFifoThroughputScope()
	// Experimental.
	ResetFifoTopic()
	// Experimental.
	ResetFirehoseFailureFeedbackRoleArn()
	// Experimental.
	ResetFirehoseSuccessFeedbackRoleArn()
	// Experimental.
	ResetFirehoseSuccessFeedbackSampleRate()
	// Experimental.
	ResetHttpFailureFeedbackRoleArn()
	// Experimental.
	ResetHttpSuccessFeedbackRoleArn()
	// Experimental.
	ResetHttpSuccessFeedbackSampleRate()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsMasterKeyId()
	// Experimental.
	ResetLambdaFailureFeedbackRoleArn()
	// Experimental.
	ResetLambdaSuccessFeedbackRoleArn()
	// Experimental.
	ResetLambdaSuccessFeedbackSampleRate()
	// Experimental.
	ResetName()
	// Experimental.
	ResetNamePrefix()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSignatureVersion()
	// Experimental.
	ResetSqsFailureFeedbackRoleArn()
	// Experimental.
	ResetSqsSuccessFeedbackRoleArn()
	// Experimental.
	ResetSqsSuccessFeedbackSampleRate()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTracingConfig()
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

// The jsii proxy struct for TfTopic
type jsiiProxy_TfTopic struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTopic) ApplicationFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ApplicationFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ApplicationSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ApplicationSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ApplicationSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ApplicationSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ArchivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ArchivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) BeginningArchiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginningArchiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FifoThroughputScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FifoThroughputScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FirehoseSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) HttpSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) LambdaSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SignatureVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SignatureVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) SqsSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TracingConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTopic) TracingConfigInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic} Resource.
// Experimental.
func NewTfTopic(scope constructs.Construct, id *string, config *TfTopicConfig) TfTopic {
	_init_.Initialize()

	if err := validateNewTfTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTopic{}

	_jsii_.Create(
		"@cdktn/aws-sns.TfTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic} Resource.
// Experimental.
func NewTfTopic_Override(t TfTopic, scope constructs.Construct, id *string, config *TfTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sns.TfTopic",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTopic)SetApplicationFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetApplicationSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetApplicationSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetApplicationSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetArchivePolicy(val *string) {
	if err := j.validateSetArchivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivePolicy",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetDeliveryPolicy(val *string) {
	if err := j.validateSetDeliveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetFifoThroughputScope(val *string) {
	if err := j.validateSetFifoThroughputScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputScope",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetFifoTopic(val interface{}) {
	if err := j.validateSetFifoTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetFirehoseFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetFirehoseSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetFirehoseSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetFirehoseSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetHttpFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetHttpSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetHttpSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetHttpSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetLambdaFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetLambdaSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetLambdaSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetLambdaSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetSignatureVersion(val *float64) {
	if err := j.validateSetSignatureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureVersion",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetSqsFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetSqsSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetSqsSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetSqsSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_TfTopic)SetTracingConfig(val *string) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

// Generates CDKTN code for importing a TfTopic resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfTopic",
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
func TfTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sns.TfTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTopic) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTopic) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTopic) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTopic) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTopic) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTopic) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTopic) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTopic) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTopic) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTopic) ResetApplicationFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetApplicationFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetApplicationSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetApplicationSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetApplicationSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetApplicationSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetArchivePolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetArchivePolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		t,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		t,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetFifoThroughputScope() {
	_jsii_.InvokeVoid(
		t,
		"resetFifoThroughputScope",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		t,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetFirehoseFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehoseFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetFirehoseSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehoseSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetFirehoseSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetFirehoseSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetHttpFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetHttpSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetHttpSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetLambdaFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetLambdaSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetLambdaSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetLambdaSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetName() {
	_jsii_.InvokeVoid(
		t,
		"resetName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetSignatureVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetSignatureVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetSqsFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetSqsSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetSqsSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetSqsSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetTagsAll() {
	_jsii_.InvokeVoid(
		t,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTopic) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

