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
type AwsSnsTopic interface {
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

// The jsii proxy struct for AwsSnsTopic
type jsiiProxy_AwsSnsTopic struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ApplicationSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"applicationSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ArchivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ArchivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"archivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) BeginningArchiveTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"beginningArchiveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) DeliveryPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) DeliveryPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deliveryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FifoThroughputScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FifoThroughputScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FifoTopic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FifoTopicInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoTopicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FirehoseSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firehoseSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) HttpSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) LambdaSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lambdaSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SignatureVersion() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SignatureVersionInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"signatureVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsFailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsFailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsFailureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsSuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsSuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsSuccessFeedbackSampleRate() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) SqsSuccessFeedbackSampleRateInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sqsSuccessFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TracingConfig() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tracingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSnsTopic) TracingConfigInput() *string {
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
func NewAwsSnsTopic(scope constructs.Construct, id *string, config *AwsSnsTopicConfig) AwsSnsTopic {
	_init_.Initialize()

	if err := validateNewAwsSnsTopicParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSnsTopic{}

	_jsii_.Create(
		"@cdktn/aws-sns.AwsSnsTopic",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_topic aws_sns_topic} Resource.
// Experimental.
func NewAwsSnsTopic_Override(a AwsSnsTopic, scope constructs.Construct, id *string, config *AwsSnsTopicConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sns.AwsSnsTopic",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetApplicationFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetApplicationSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetApplicationSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetApplicationSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetApplicationSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetArchivePolicy(val *string) {
	if err := j.validateSetArchivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"archivePolicy",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetDeliveryPolicy(val *string) {
	if err := j.validateSetDeliveryPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deliveryPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetFifoThroughputScope(val *string) {
	if err := j.validateSetFifoThroughputScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputScope",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetFifoTopic(val interface{}) {
	if err := j.validateSetFifoTopicParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoTopic",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetFirehoseFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetFirehoseSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetFirehoseSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetFirehoseSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetFirehoseSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firehoseSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetHttpFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetHttpSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetHttpSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetHttpSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetHttpSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetLambdaFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetLambdaSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetLambdaSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetLambdaSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetLambdaSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lambdaSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetSignatureVersion(val *float64) {
	if err := j.validateSetSignatureVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signatureVersion",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetSqsFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsFailureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetSqsSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSqsSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetSqsSuccessFeedbackSampleRate(val *float64) {
	if err := j.validateSetSqsSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsSuccessFeedbackSampleRate",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsSnsTopic)SetTracingConfig(val *string) {
	if err := j.validateSetTracingConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tracingConfig",
		val,
	)
}

// Generates CDKTN code for importing a AwsSnsTopic resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSnsTopic_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSnsTopic_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsSnsTopic",
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
func AwsSnsTopic_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSnsTopic_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsSnsTopic",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSnsTopic_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSnsTopic_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsSnsTopic",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSnsTopic_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSnsTopic_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.AwsSnsTopic",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSnsTopic_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sns.AwsSnsTopic",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSnsTopic) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSnsTopic) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSnsTopic) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSnsTopic) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSnsTopic) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSnsTopic) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSnsTopic) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSnsTopic) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSnsTopic) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSnsTopic) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSnsTopic) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSnsTopic) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSnsTopic) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSnsTopic) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSnsTopic) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSnsTopic) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSnsTopic) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSnsTopic) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSnsTopic) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetApplicationFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetApplicationSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetApplicationSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetArchivePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetArchivePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		a,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetDeliveryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeliveryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetDisplayName() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetFifoThroughputScope() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoThroughputScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetFifoTopic() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoTopic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetFirehoseFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetFirehoseSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetFirehoseSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetFirehoseSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetHttpFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetHttpSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetHttpSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetLambdaFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetLambdaSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetLambdaSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetLambdaSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetSignatureVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetSignatureVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetSqsFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetSqsSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetSqsSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) ResetTracingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTracingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSnsTopic) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSnsTopic) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

