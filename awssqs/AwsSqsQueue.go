package awssqs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssqs/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssqs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue aws_sqs_queue}.
// Experimental.
type AwsSqsQueue interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
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
	DeduplicationScope() *string
	// Experimental.
	SetDeduplicationScope(val *string)
	// Experimental.
	DeduplicationScopeInput() *string
	// Experimental.
	DelaySeconds() *float64
	// Experimental.
	SetDelaySeconds(val *float64)
	// Experimental.
	DelaySecondsInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	FifoQueue() interface{}
	// Experimental.
	SetFifoQueue(val interface{})
	// Experimental.
	FifoQueueInput() interface{}
	// Experimental.
	FifoThroughputLimit() *string
	// Experimental.
	SetFifoThroughputLimit(val *string)
	// Experimental.
	FifoThroughputLimitInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	KmsDataKeyReusePeriodSeconds() *float64
	// Experimental.
	SetKmsDataKeyReusePeriodSeconds(val *float64)
	// Experimental.
	KmsDataKeyReusePeriodSecondsInput() *float64
	// Experimental.
	KmsMasterKeyId() *string
	// Experimental.
	SetKmsMasterKeyId(val *string)
	// Experimental.
	KmsMasterKeyIdInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaxMessageSize() *float64
	// Experimental.
	SetMaxMessageSize(val *float64)
	// Experimental.
	MaxMessageSizeInput() *float64
	// Experimental.
	MessageRetentionSeconds() *float64
	// Experimental.
	SetMessageRetentionSeconds(val *float64)
	// Experimental.
	MessageRetentionSecondsInput() *float64
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
	ReceiveWaitTimeSeconds() *float64
	// Experimental.
	SetReceiveWaitTimeSeconds(val *float64)
	// Experimental.
	ReceiveWaitTimeSecondsInput() *float64
	// Experimental.
	RedriveAllowPolicy() *string
	// Experimental.
	SetRedriveAllowPolicy(val *string)
	// Experimental.
	RedriveAllowPolicyInput() *string
	// Experimental.
	RedrivePolicy() *string
	// Experimental.
	SetRedrivePolicy(val *string)
	// Experimental.
	RedrivePolicyInput() *string
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SqsManagedSseEnabled() interface{}
	// Experimental.
	SetSqsManagedSseEnabled(val interface{})
	// Experimental.
	SqsManagedSseEnabledInput() interface{}
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
	Timeouts() AwsSqsQueue_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Url() *string
	// Experimental.
	VisibilityTimeoutSeconds() *float64
	// Experimental.
	SetVisibilityTimeoutSeconds(val *float64)
	// Experimental.
	VisibilityTimeoutSecondsInput() *float64
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
	PutTimeouts(value *AwsSqsQueue_TimeoutsProperty)
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
	ResetContentBasedDeduplication()
	// Experimental.
	ResetDeduplicationScope()
	// Experimental.
	ResetDelaySeconds()
	// Experimental.
	ResetFifoQueue()
	// Experimental.
	ResetFifoThroughputLimit()
	// Experimental.
	ResetId()
	// Experimental.
	ResetKmsDataKeyReusePeriodSeconds()
	// Experimental.
	ResetKmsMasterKeyId()
	// Experimental.
	ResetMaxMessageSize()
	// Experimental.
	ResetMessageRetentionSeconds()
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
	ResetReceiveWaitTimeSeconds()
	// Experimental.
	ResetRedriveAllowPolicy()
	// Experimental.
	ResetRedrivePolicy()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSqsManagedSseEnabled()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTagsAll()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVisibilityTimeoutSeconds()
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

// The jsii proxy struct for AwsSqsQueue
type jsiiProxy_AwsSqsQueue struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSqsQueue) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ContentBasedDeduplication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ContentBasedDeduplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"contentBasedDeduplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) DeduplicationScope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationScope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) DeduplicationScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deduplicationScopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) DelaySeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delaySeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) DelaySecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"delaySecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) FifoQueue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) FifoQueueInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fifoQueueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) FifoThroughputLimit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) FifoThroughputLimitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fifoThroughputLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) KmsDataKeyReusePeriodSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) KmsDataKeyReusePeriodSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kmsDataKeyReusePeriodSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) KmsMasterKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) KmsMasterKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsMasterKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) MaxMessageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) MaxMessageSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxMessageSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) MessageRetentionSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageRetentionSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) MessageRetentionSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messageRetentionSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) NamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) NamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Policy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) PolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ReceiveWaitTimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveWaitTimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) ReceiveWaitTimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"receiveWaitTimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RedriveAllowPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redriveAllowPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RedriveAllowPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redriveAllowPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RedrivePolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RedrivePolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"redrivePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) SqsManagedSseEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsManagedSseEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) SqsManagedSseEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sqsManagedSseEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TagsAll() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TagsAllInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Timeouts() AwsSqsQueue_TimeoutsPropertyOutputReference {
	var returns AwsSqsQueue_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) VisibilityTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSqsQueue) VisibilityTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"visibilityTimeoutSecondsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue aws_sqs_queue} Resource.
// Experimental.
func NewAwsSqsQueue(scope constructs.Construct, id *string, config *AwsSqsQueueConfig) AwsSqsQueue {
	_init_.Initialize()

	if err := validateNewAwsSqsQueueParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSqsQueue{}

	_jsii_.Create(
		"@cdktn/aws-sqs.AwsSqsQueue",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sqs_queue aws_sqs_queue} Resource.
// Experimental.
func NewAwsSqsQueue_Override(a AwsSqsQueue, scope constructs.Construct, id *string, config *AwsSqsQueueConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sqs.AwsSqsQueue",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetContentBasedDeduplication(val interface{}) {
	if err := j.validateSetContentBasedDeduplicationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentBasedDeduplication",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetDeduplicationScope(val *string) {
	if err := j.validateSetDeduplicationScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deduplicationScope",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetDelaySeconds(val *float64) {
	if err := j.validateSetDelaySecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"delaySeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetFifoQueue(val interface{}) {
	if err := j.validateSetFifoQueueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoQueue",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetFifoThroughputLimit(val *string) {
	if err := j.validateSetFifoThroughputLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fifoThroughputLimit",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetKmsDataKeyReusePeriodSeconds(val *float64) {
	if err := j.validateSetKmsDataKeyReusePeriodSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsDataKeyReusePeriodSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetKmsMasterKeyId(val *string) {
	if err := j.validateSetKmsMasterKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsMasterKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetMaxMessageSize(val *float64) {
	if err := j.validateSetMaxMessageSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxMessageSize",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetMessageRetentionSeconds(val *float64) {
	if err := j.validateSetMessageRetentionSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messageRetentionSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetNamePrefix(val *string) {
	if err := j.validateSetNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namePrefix",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetPolicy(val *string) {
	if err := j.validateSetPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policy",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetReceiveWaitTimeSeconds(val *float64) {
	if err := j.validateSetReceiveWaitTimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"receiveWaitTimeSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetRedriveAllowPolicy(val *string) {
	if err := j.validateSetRedriveAllowPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redriveAllowPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetRedrivePolicy(val *string) {
	if err := j.validateSetRedrivePolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"redrivePolicy",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetSqsManagedSseEnabled(val interface{}) {
	if err := j.validateSetSqsManagedSseEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqsManagedSseEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetTagsAll(val *map[string]*string) {
	if err := j.validateSetTagsAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagsAll",
		val,
	)
}

func (j *jsiiProxy_AwsSqsQueue)SetVisibilityTimeoutSeconds(val *float64) {
	if err := j.validateSetVisibilityTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibilityTimeoutSeconds",
		val,
	)
}

// Generates CDKTN code for importing a AwsSqsQueue resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSqsQueue_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSqsQueue_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sqs.AwsSqsQueue",
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
func AwsSqsQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSqsQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sqs.AwsSqsQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSqsQueue_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSqsQueue_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sqs.AwsSqsQueue",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSqsQueue_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSqsQueue_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sqs.AwsSqsQueue",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSqsQueue_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sqs.AwsSqsQueue",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSqsQueue) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSqsQueue) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSqsQueue) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSqsQueue) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSqsQueue) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSqsQueue) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSqsQueue) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSqsQueue) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSqsQueue) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSqsQueue) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSqsQueue) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSqsQueue) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSqsQueue) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSqsQueue) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSqsQueue) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSqsQueue) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSqsQueue) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSqsQueue) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSqsQueue) PutTimeouts(value *AwsSqsQueue_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSqsQueue) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetContentBasedDeduplication() {
	_jsii_.InvokeVoid(
		a,
		"resetContentBasedDeduplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetDeduplicationScope() {
	_jsii_.InvokeVoid(
		a,
		"resetDeduplicationScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetDelaySeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDelaySeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetFifoQueue() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoQueue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetFifoThroughputLimit() {
	_jsii_.InvokeVoid(
		a,
		"resetFifoThroughputLimit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetKmsDataKeyReusePeriodSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsDataKeyReusePeriodSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetKmsMasterKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsMasterKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetMaxMessageSize() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxMessageSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetMessageRetentionSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMessageRetentionSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetName() {
	_jsii_.InvokeVoid(
		a,
		"resetName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetNamePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetNamePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetReceiveWaitTimeSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetReceiveWaitTimeSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetRedriveAllowPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRedriveAllowPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetRedrivePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRedrivePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetSqsManagedSseEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetSqsManagedSseEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetTagsAll() {
	_jsii_.InvokeVoid(
		a,
		"resetTagsAll",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) ResetVisibilityTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetVisibilityTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSqsQueue) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSqsQueue) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

