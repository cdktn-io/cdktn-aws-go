package awssns

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssns/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssns/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application aws_sns_platform_application}.
// Experimental.
type TfPlatformApplication interface {
	cdktn.TerraformResource
	// Experimental.
	ApplePlatformBundleId() *string
	// Experimental.
	SetApplePlatformBundleId(val *string)
	// Experimental.
	ApplePlatformBundleIdInput() *string
	// Experimental.
	ApplePlatformTeamId() *string
	// Experimental.
	SetApplePlatformTeamId(val *string)
	// Experimental.
	ApplePlatformTeamIdInput() *string
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
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EventDeliveryFailureTopicArn() *string
	// Experimental.
	SetEventDeliveryFailureTopicArn(val *string)
	// Experimental.
	EventDeliveryFailureTopicArnInput() *string
	// Experimental.
	EventEndpointCreatedTopicArn() *string
	// Experimental.
	SetEventEndpointCreatedTopicArn(val *string)
	// Experimental.
	EventEndpointCreatedTopicArnInput() *string
	// Experimental.
	EventEndpointDeletedTopicArn() *string
	// Experimental.
	SetEventEndpointDeletedTopicArn(val *string)
	// Experimental.
	EventEndpointDeletedTopicArnInput() *string
	// Experimental.
	EventEndpointUpdatedTopicArn() *string
	// Experimental.
	SetEventEndpointUpdatedTopicArn(val *string)
	// Experimental.
	EventEndpointUpdatedTopicArnInput() *string
	// Experimental.
	FailureFeedbackRoleArn() *string
	// Experimental.
	SetFailureFeedbackRoleArn(val *string)
	// Experimental.
	FailureFeedbackRoleArnInput() *string
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
	Platform() *string
	// Experimental.
	SetPlatform(val *string)
	// Experimental.
	PlatformCredential() *string
	// Experimental.
	SetPlatformCredential(val *string)
	// Experimental.
	PlatformCredentialInput() *string
	// Experimental.
	PlatformInput() *string
	// Experimental.
	PlatformPrincipal() *string
	// Experimental.
	SetPlatformPrincipal(val *string)
	// Experimental.
	PlatformPrincipalInput() *string
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
	SuccessFeedbackRoleArn() *string
	// Experimental.
	SetSuccessFeedbackRoleArn(val *string)
	// Experimental.
	SuccessFeedbackRoleArnInput() *string
	// Experimental.
	SuccessFeedbackSampleRate() *string
	// Experimental.
	SetSuccessFeedbackSampleRate(val *string)
	// Experimental.
	SuccessFeedbackSampleRateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetApplePlatformBundleId()
	// Experimental.
	ResetApplePlatformTeamId()
	// Experimental.
	ResetEventDeliveryFailureTopicArn()
	// Experimental.
	ResetEventEndpointCreatedTopicArn()
	// Experimental.
	ResetEventEndpointDeletedTopicArn()
	// Experimental.
	ResetEventEndpointUpdatedTopicArn()
	// Experimental.
	ResetFailureFeedbackRoleArn()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPlatformPrincipal()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSuccessFeedbackRoleArn()
	// Experimental.
	ResetSuccessFeedbackSampleRate()
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

// The jsii proxy struct for TfPlatformApplication
type jsiiProxy_TfPlatformApplication struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfPlatformApplication) ApplePlatformBundleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformBundleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) ApplePlatformBundleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformBundleIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) ApplePlatformTeamId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformTeamId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) ApplePlatformTeamIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applePlatformTeamIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventDeliveryFailureTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventDeliveryFailureTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventDeliveryFailureTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointCreatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointCreatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointCreatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointDeletedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointDeletedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointDeletedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointUpdatedTopicArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) EventEndpointUpdatedTopicArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventEndpointUpdatedTopicArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) FailureFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) FailureFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Platform() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platform",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) PlatformCredential() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredential",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) PlatformCredentialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformCredentialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) PlatformInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) PlatformPrincipal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) PlatformPrincipalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"platformPrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) SuccessFeedbackRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) SuccessFeedbackRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) SuccessFeedbackSampleRate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) SuccessFeedbackSampleRateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successFeedbackSampleRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlatformApplication) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application aws_sns_platform_application} Resource.
// Experimental.
func NewTfPlatformApplication(scope constructs.Construct, id *string, config *TfPlatformApplicationConfig) TfPlatformApplication {
	_init_.Initialize()

	if err := validateNewTfPlatformApplicationParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPlatformApplication{}

	_jsii_.Create(
		"@cdktn/aws-sns.TfPlatformApplication",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sns_platform_application aws_sns_platform_application} Resource.
// Experimental.
func NewTfPlatformApplication_Override(t TfPlatformApplication, scope constructs.Construct, id *string, config *TfPlatformApplicationConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sns.TfPlatformApplication",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetApplePlatformBundleId(val *string) {
	if err := j.validateSetApplePlatformBundleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePlatformBundleId",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetApplePlatformTeamId(val *string) {
	if err := j.validateSetApplePlatformTeamIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applePlatformTeamId",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetEventDeliveryFailureTopicArn(val *string) {
	if err := j.validateSetEventDeliveryFailureTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventDeliveryFailureTopicArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetEventEndpointCreatedTopicArn(val *string) {
	if err := j.validateSetEventEndpointCreatedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointCreatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetEventEndpointDeletedTopicArn(val *string) {
	if err := j.validateSetEventEndpointDeletedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointDeletedTopicArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetEventEndpointUpdatedTopicArn(val *string) {
	if err := j.validateSetEventEndpointUpdatedTopicArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eventEndpointUpdatedTopicArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetFailureFeedbackRoleArn(val *string) {
	if err := j.validateSetFailureFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetPlatform(val *string) {
	if err := j.validateSetPlatformParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platform",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetPlatformCredential(val *string) {
	if err := j.validateSetPlatformCredentialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformCredential",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetPlatformPrincipal(val *string) {
	if err := j.validateSetPlatformPrincipalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformPrincipal",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetSuccessFeedbackRoleArn(val *string) {
	if err := j.validateSetSuccessFeedbackRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successFeedbackRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfPlatformApplication)SetSuccessFeedbackSampleRate(val *string) {
	if err := j.validateSetSuccessFeedbackSampleRateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successFeedbackSampleRate",
		val,
	)
}

// Generates CDKTN code for importing a TfPlatformApplication resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfPlatformApplication_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfPlatformApplication_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfPlatformApplication",
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
func TfPlatformApplication_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPlatformApplication_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfPlatformApplication",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPlatformApplication_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPlatformApplication_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfPlatformApplication",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfPlatformApplication_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfPlatformApplication_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sns.TfPlatformApplication",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfPlatformApplication_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sns.TfPlatformApplication",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfPlatformApplication) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfPlatformApplication) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfPlatformApplication) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPlatformApplication) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlatformApplication) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPlatformApplication) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPlatformApplication) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPlatformApplication) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPlatformApplication) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPlatformApplication) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPlatformApplication) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPlatformApplication) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfPlatformApplication) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlatformApplication) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfPlatformApplication) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPlatformApplication) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfPlatformApplication) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfPlatformApplication) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfPlatformApplication) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetApplePlatformBundleId() {
	_jsii_.InvokeVoid(
		t,
		"resetApplePlatformBundleId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetApplePlatformTeamId() {
	_jsii_.InvokeVoid(
		t,
		"resetApplePlatformTeamId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetEventDeliveryFailureTopicArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEventDeliveryFailureTopicArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetEventEndpointCreatedTopicArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEventEndpointCreatedTopicArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetEventEndpointDeletedTopicArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEventEndpointDeletedTopicArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetEventEndpointUpdatedTopicArn() {
	_jsii_.InvokeVoid(
		t,
		"resetEventEndpointUpdatedTopicArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetFailureFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetPlatformPrincipal() {
	_jsii_.InvokeVoid(
		t,
		"resetPlatformPrincipal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetSuccessFeedbackRoleArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessFeedbackRoleArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) ResetSuccessFeedbackSampleRate() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessFeedbackSampleRate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlatformApplication) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlatformApplication) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

