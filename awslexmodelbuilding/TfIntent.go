package awslexmodelbuilding

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awslexmodelbuilding/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent aws_lex_intent}.
// Experimental.
type TfIntent interface {
	cdktn.TerraformResource
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	Checksum() *string
	// Experimental.
	ConclusionStatement() TfIntent_ConclusionStatementPropertyOutputReference
	// Experimental.
	ConclusionStatementInput() *TfIntent_ConclusionStatementProperty
	// Experimental.
	ConfirmationPrompt() TfIntent_ConfirmationPromptPropertyOutputReference
	// Experimental.
	ConfirmationPromptInput() *TfIntent_ConfirmationPromptProperty
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
	DialogCodeHook() TfIntent_DialogCodeHookPropertyOutputReference
	// Experimental.
	DialogCodeHookInput() *TfIntent_DialogCodeHookProperty
	// Experimental.
	FollowUpPrompt() TfIntent_FollowUpPromptPropertyOutputReference
	// Experimental.
	FollowUpPromptInput() *TfIntent_FollowUpPromptProperty
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FulfillmentActivity() TfIntent_FulfillmentActivityPropertyOutputReference
	// Experimental.
	FulfillmentActivityInput() *TfIntent_FulfillmentActivityProperty
	// Experimental.
	Id() *string
	// Experimental.
	SetId(val *string)
	// Experimental.
	IdInput() *string
	// Experimental.
	LastUpdatedDate() *string
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
	ParentIntentSignature() *string
	// Experimental.
	SetParentIntentSignature(val *string)
	// Experimental.
	ParentIntentSignatureInput() *string
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
	RejectionStatement() TfIntent_RejectionStatementPropertyOutputReference
	// Experimental.
	RejectionStatementInput() *TfIntent_RejectionStatementProperty
	// Experimental.
	SampleUtterances() *[]*string
	// Experimental.
	SetSampleUtterances(val *[]*string)
	// Experimental.
	SampleUtterancesInput() *[]*string
	// Experimental.
	Slot() TfIntent_SlotPropertyList
	// Experimental.
	SlotInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfIntent_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Version() *string
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
	PutConclusionStatement(value *TfIntent_ConclusionStatementProperty)
	// Experimental.
	PutConfirmationPrompt(value *TfIntent_ConfirmationPromptProperty)
	// Experimental.
	PutDialogCodeHook(value *TfIntent_DialogCodeHookProperty)
	// Experimental.
	PutFollowUpPrompt(value *TfIntent_FollowUpPromptProperty)
	// Experimental.
	PutFulfillmentActivity(value *TfIntent_FulfillmentActivityProperty)
	// Experimental.
	PutRejectionStatement(value *TfIntent_RejectionStatementProperty)
	// Experimental.
	PutSlot(value interface{})
	// Experimental.
	PutTimeouts(value *TfIntent_TimeoutsProperty)
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
	ResetConclusionStatement()
	// Experimental.
	ResetConfirmationPrompt()
	// Experimental.
	ResetCreateVersion()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDialogCodeHook()
	// Experimental.
	ResetFollowUpPrompt()
	// Experimental.
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParentIntentSignature()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRejectionStatement()
	// Experimental.
	ResetSampleUtterances()
	// Experimental.
	ResetSlot()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfIntent
type jsiiProxy_TfIntent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfIntent) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Checksum() *string {
	var returns *string
	_jsii_.Get(
		j,
		"checksum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ConclusionStatement() TfIntent_ConclusionStatementPropertyOutputReference {
	var returns TfIntent_ConclusionStatementPropertyOutputReference
	_jsii_.Get(
		j,
		"conclusionStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ConclusionStatementInput() *TfIntent_ConclusionStatementProperty {
	var returns *TfIntent_ConclusionStatementProperty
	_jsii_.Get(
		j,
		"conclusionStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ConfirmationPrompt() TfIntent_ConfirmationPromptPropertyOutputReference {
	var returns TfIntent_ConfirmationPromptPropertyOutputReference
	_jsii_.Get(
		j,
		"confirmationPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ConfirmationPromptInput() *TfIntent_ConfirmationPromptProperty {
	var returns *TfIntent_ConfirmationPromptProperty
	_jsii_.Get(
		j,
		"confirmationPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) CreatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createdDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) CreateVersion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) CreateVersionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"createVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) DialogCodeHook() TfIntent_DialogCodeHookPropertyOutputReference {
	var returns TfIntent_DialogCodeHookPropertyOutputReference
	_jsii_.Get(
		j,
		"dialogCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) DialogCodeHookInput() *TfIntent_DialogCodeHookProperty {
	var returns *TfIntent_DialogCodeHookProperty
	_jsii_.Get(
		j,
		"dialogCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) FollowUpPrompt() TfIntent_FollowUpPromptPropertyOutputReference {
	var returns TfIntent_FollowUpPromptPropertyOutputReference
	_jsii_.Get(
		j,
		"followUpPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) FollowUpPromptInput() *TfIntent_FollowUpPromptProperty {
	var returns *TfIntent_FollowUpPromptProperty
	_jsii_.Get(
		j,
		"followUpPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) FulfillmentActivity() TfIntent_FulfillmentActivityPropertyOutputReference {
	var returns TfIntent_FulfillmentActivityPropertyOutputReference
	_jsii_.Get(
		j,
		"fulfillmentActivity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) FulfillmentActivityInput() *TfIntent_FulfillmentActivityProperty {
	var returns *TfIntent_FulfillmentActivityProperty
	_jsii_.Get(
		j,
		"fulfillmentActivityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) LastUpdatedDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) ParentIntentSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) RejectionStatement() TfIntent_RejectionStatementPropertyOutputReference {
	var returns TfIntent_RejectionStatementPropertyOutputReference
	_jsii_.Get(
		j,
		"rejectionStatement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) RejectionStatementInput() *TfIntent_RejectionStatementProperty {
	var returns *TfIntent_RejectionStatementProperty
	_jsii_.Get(
		j,
		"rejectionStatementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) SampleUtterances() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) SampleUtterancesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sampleUtterancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Slot() TfIntent_SlotPropertyList {
	var returns TfIntent_SlotPropertyList
	_jsii_.Get(
		j,
		"slot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) SlotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Timeouts() TfIntent_TimeoutsPropertyOutputReference {
	var returns TfIntent_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIntent) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent aws_lex_intent} Resource.
// Experimental.
func NewTfIntent(scope constructs.Construct, id *string, config *TfIntentConfig) TfIntent {
	_init_.Initialize()

	if err := validateNewTfIntentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIntent{}

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lex_intent aws_lex_intent} Resource.
// Experimental.
func NewTfIntent_Override(t TfIntent, scope constructs.Construct, id *string, config *TfIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-model-building.TfIntent",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfIntent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetCreateVersion(val interface{}) {
	if err := j.validateSetCreateVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createVersion",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetParentIntentSignature(val *string) {
	if err := j.validateSetParentIntentSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentIntentSignature",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfIntent)SetSampleUtterances(val *[]*string) {
	if err := j.validateSetSampleUtterancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sampleUtterances",
		val,
	)
}

// Generates CDKTN code for importing a TfIntent resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfIntent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfIntent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfIntent",
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
func TfIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfIntent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfIntent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfIntent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfIntent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfIntent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfIntent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-model-building.TfIntent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lex-model-building.TfIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfIntent) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfIntent) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIntent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIntent) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIntent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfIntent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIntent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfIntent) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfIntent) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfIntent) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfIntent) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfIntent) PutConclusionStatement(value *TfIntent_ConclusionStatementProperty) {
	if err := t.validatePutConclusionStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConclusionStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutConfirmationPrompt(value *TfIntent_ConfirmationPromptProperty) {
	if err := t.validatePutConfirmationPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConfirmationPrompt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutDialogCodeHook(value *TfIntent_DialogCodeHookProperty) {
	if err := t.validatePutDialogCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDialogCodeHook",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutFollowUpPrompt(value *TfIntent_FollowUpPromptProperty) {
	if err := t.validatePutFollowUpPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFollowUpPrompt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutFulfillmentActivity(value *TfIntent_FulfillmentActivityProperty) {
	if err := t.validatePutFulfillmentActivityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFulfillmentActivity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutRejectionStatement(value *TfIntent_RejectionStatementProperty) {
	if err := t.validatePutRejectionStatementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRejectionStatement",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutSlot(value interface{}) {
	if err := t.validatePutSlotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSlot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) PutTimeouts(value *TfIntent_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfIntent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfIntent) ResetConclusionStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetConclusionStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetConfirmationPrompt() {
	_jsii_.InvokeVoid(
		t,
		"resetConfirmationPrompt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetCreateVersion() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateVersion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetDialogCodeHook() {
	_jsii_.InvokeVoid(
		t,
		"resetDialogCodeHook",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetFollowUpPrompt() {
	_jsii_.InvokeVoid(
		t,
		"resetFollowUpPrompt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetParentIntentSignature() {
	_jsii_.InvokeVoid(
		t,
		"resetParentIntentSignature",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetRejectionStatement() {
	_jsii_.InvokeVoid(
		t,
		"resetRejectionStatement",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetSampleUtterances() {
	_jsii_.InvokeVoid(
		t,
		"resetSampleUtterances",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetSlot() {
	_jsii_.InvokeVoid(
		t,
		"resetSlot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIntent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

