package sns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sns/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/sns/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic}.
// Experimental.
type AwsTopic interface {
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

// The jsii proxy struct for AwsTopic
type jsiiProxy_AwsTopic struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsTopic) ApplicationFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ApplicationFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ApplicationSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ApplicationSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ApplicationSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ApplicationSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ArchivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ArchivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) BeginningArchiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginningArchiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FifoThroughputScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FifoThroughputScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FirehoseSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) HttpSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) LambdaSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SignatureVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SignatureVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) SqsSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TracingConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTopic) TracingConfigInput() *string {
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
func NewAwsTopic(scope constructs.Construct, id *string, config *AwsTopicConfig) AwsTopic {
	_init_.Initialize()

	if err := validateNewAwsTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTopic{}

	_jsii_.Create(
		"@cdktn/aws-sns.AwsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic} Resource.
// Experimental.
func NewAwsTopic_Override(a AwsTopic, scope constructs.Construct, id *string, config *AwsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sns.AwsTopic",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsTopic)SetApplicationFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetApplicationSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetApplicationSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetApplicationSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetArchivePolicy(val *string) {
	if err := j.validateSetArchivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivePolicy",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetDeliveryPolicy(val *string) {
	if err := j.validateSetDeliveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetFifoThroughputScope(val *string) {
	if err := j.validateSetFifoThroughputScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputScope",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetFifoTopic(val interface{}) {
	if err := j.validateSetFifoTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetFirehoseFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetFirehoseSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetFirehoseSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetFirehoseSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetHttpFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetHttpSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetHttpSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetHttpSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetLambdaFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetLambdaSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetLambdaSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetLambdaSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetSignatureVersion(val *float64) {
	if err := j.validateSetSignatureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureVersion",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetSqsFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetSqsSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetSqsSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetSqsSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsTopic)SetTracingConfig(val *string) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

// Generates CDKTN code for importing a AwsTopic resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsTopic",
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
func AwsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sns.AwsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsTopic) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsTopic) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTopic) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsTopic) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTopic) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsTopic) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsTopic) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsTopic) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsTopic) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsTopic) ResetApplicationFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetApplicationSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetApplicationSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetArchivePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetArchivePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		a,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetFifoThroughputScope() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoThroughputScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetFirehoseFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetFirehoseSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetFirehoseSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetHttpFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetHttpSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetHttpSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetLambdaFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetLambdaSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetLambdaSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetSignatureVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetSignatureVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetSqsFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetSqsSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetSqsSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTopic) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

