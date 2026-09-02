package awslexmodelbuilding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot aws_lex_bot}.
// Experimental.
type TfBot interface {
	cdktn.TerraformResource
	// Experimental.
	AbortStatement() TfBot_AbortStatementPropertyOutputReference
	// Experimental.
	AbortStatementInput() *TfBot_AbortStatementProperty
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Checksum() *string
	// Experimental.
	ChildDirected() interface{}
	// Experimental.
	SetChildDirected(val interface{})
	// Experimental.
	ChildDirectedInput() interface{}
	// Experimental.
	ClarificationPrompt() TfBot_ClarificationPromptPropertyOutputReference
	// Experimental.
	ClarificationPromptInput() *TfBot_ClarificationPromptProperty
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
	CreatedDate() *string
	// Experimental.
	CreateVersion() interface{}
	// Experimental.
	SetCreateVersion(val interface{})
	// Experimental.
	CreateVersionInput() interface{}
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
	DetectSentiment() interface{}
	// Experimental.
	SetDetectSentiment(val interface{})
	// Experimental.
	DetectSentimentInput() interface{}
	// Experimental.
	EnableModelImprovements() interface{}
	// Experimental.
	SetEnableModelImprovements(val interface{})
	// Experimental.
	EnableModelImprovementsInput() interface{}
	// Experimental.
	FailureReason() *string
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
	IdleSessionTtlInSeconds() *float64
	// Experimental.
	SetIdleSessionTtlInSeconds(val *float64)
	// Experimental.
	IdleSessionTtlInSecondsInput() *float64
	// Experimental.
	Intent() TfBot_IntentPropertyList
	// Experimental.
	IntentInput() interface{}
	// Experimental.
	LastUpdatedDate() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	Locale() *string
	// Experimental.
	SetLocale(val *string)
	// Experimental.
	LocaleInput() *string
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NluIntentConfidenceThreshold() *float64
	// Experimental.
	SetNluIntentConfidenceThreshold(val *float64)
	// Experimental.
	NluIntentConfidenceThresholdInput() *float64
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ProcessBehavior() *string
	// Experimental.
	SetProcessBehavior(val *string)
	// Experimental.
	ProcessBehaviorInput() *string
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
	Status() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfBot_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Version() *string
	// Experimental.
	VoiceId() *string
	// Experimental.
	SetVoiceId(val *string)
	// Experimental.
	VoiceIdInput() *string
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
	PutAbortStatement(value *TfBot_AbortStatementProperty)
	// Experimental.
	PutClarificationPrompt(value *TfBot_ClarificationPromptProperty)
	// Experimental.
	PutIntent(value interface{})
	// Experimental.
	PutTimeouts(value *TfBot_TimeoutsProperty)
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
	ResetClarificationPrompt()
	// Experimental.
	ResetCreateVersion()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDetectSentiment()
	// Experimental.
	ResetEnableModelImprovements()
	// Experimental.
	ResetId()
	// Experimental.
	ResetIdleSessionTtlInSeconds()
	// Experimental.
	ResetLocale()
	// Experimental.
	ResetNluIntentConfidenceThreshold()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProcessBehavior()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetVoiceId()
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

// The jsii proxy struct for TfBot
type jsiiProxy_TfBot struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfBot) AbortStatement() TfBot_AbortStatementPropertyOutputReference {
	var returns TfBot_AbortStatementPropertyOutputReference
	_jsii_.Get(
		j,
		"abortStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) AbortStatementInput() *TfBot_AbortStatementProperty {
	var returns *TfBot_AbortStatementProperty
	_jsii_.Get(
		j,
		"abortStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ChildDirected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ChildDirectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"childDirectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ClarificationPrompt() TfBot_ClarificationPromptPropertyOutputReference {
	var returns TfBot_ClarificationPromptPropertyOutputReference
	_jsii_.Get(
		j,
		"clarificationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ClarificationPromptInput() *TfBot_ClarificationPromptProperty {
	var returns *TfBot_ClarificationPromptProperty
	_jsii_.Get(
		j,
		"clarificationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) DetectSentiment() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentiment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) DetectSentimentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"detectSentimentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) EnableModelImprovements() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) EnableModelImprovementsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableModelImprovementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) FailureReason() *string {
	var returns *string
	_jsii_.Get(
		j,
		"failureReason",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Intent() TfBot_IntentPropertyList {
	var returns TfBot_IntentPropertyList
	_jsii_.Get(
		j,
		"intent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) IntentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"intentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Locale() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locale",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) LocaleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) NluIntentConfidenceThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) NluIntentConfidenceThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"nluIntentConfidenceThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ProcessBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) ProcessBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"processBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Timeouts() TfBot_TimeoutsPropertyOutputReference {
	var returns TfBot_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) VoiceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBot) VoiceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"voiceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot aws_lex_bot} Resource.
// Experimental.
func NewTfBot(scope constructs.Construct, id *string, config *TfBotConfig) TfBot {
	_init_.Initialize()

	if err := validateNewTfBotParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBot{}

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfBot",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_bot aws_lex_bot} Resource.
// Experimental.
func NewTfBot_Override(t TfBot, scope constructs.Construct, id *string, config *TfBotConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfBot",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfBot)SetChildDirected(val interface{}) {
	if err := j.validateSetChildDirectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"childDirected",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetCreateVersion(val interface{}) {
	if err := j.validateSetCreateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetDetectSentiment(val interface{}) {
	if err := j.validateSetDetectSentimentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"detectSentiment",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetEnableModelImprovements(val interface{}) {
	if err := j.validateSetEnableModelImprovementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableModelImprovements",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetLocale(val *string) {
	if err := j.validateSetLocaleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locale",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetNluIntentConfidenceThreshold(val *float64) {
	if err := j.validateSetNluIntentConfidenceThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nluIntentConfidenceThreshold",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetProcessBehavior(val *string) {
	if err := j.validateSetProcessBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processBehavior",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfBot)SetVoiceId(val *string) {
	if err := j.validateSetVoiceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"voiceId",
		val,
	)
}

// Generates CDKTN code for importing a TfBot resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfBot_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfBot_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfBot",
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
func TfBot_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBot_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfBot",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBot_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBot_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfBot",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfBot_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfBot_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfBot",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfBot_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lex-model-building.TfBot",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfBot) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfBot) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfBot) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBot) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBot) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBot) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBot) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBot) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBot) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBot) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBot) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBot) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfBot) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBot) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfBot) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBot) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfBot) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfBot) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfBot) PutAbortStatement(value *TfBot_AbortStatementProperty) {
	if err := t.validatePutAbortStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAbortStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBot) PutClarificationPrompt(value *TfBot_ClarificationPromptProperty) {
	if err := t.validatePutClarificationPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClarificationPrompt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBot) PutIntent(value interface{}) {
	if err := t.validatePutIntentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIntent",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBot) PutTimeouts(value *TfBot_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBot) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfBot) ResetClarificationPrompt() {
	_jsii_.InvokeVoid(
		t,
		"resetClarificationPrompt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetDetectSentiment() {
	_jsii_.InvokeVoid(
		t,
		"resetDetectSentiment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetEnableModelImprovements() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableModelImprovements",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetLocale() {
	_jsii_.InvokeVoid(
		t,
		"resetLocale",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetNluIntentConfidenceThreshold() {
	_jsii_.InvokeVoid(
		t,
		"resetNluIntentConfidenceThreshold",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetProcessBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) ResetVoiceId() {
	_jsii_.InvokeVoid(
		t,
		"resetVoiceId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBot) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBot) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

