package lexv2models

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/lexv2models/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/lexv2models/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent aws_lexv2models_intent}.
// Experimental.
type AwsIntent interface {
	cdktn.TerraformResource
	// Experimental.
	BotId() *string
	// Experimental.
	SetBotId(val *string)
	// Experimental.
	BotIdInput() *string
	// Experimental.
	BotVersion() *string
	// Experimental.
	SetBotVersion(val *string)
	// Experimental.
	BotVersionInput() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ClosingSetting() AwsIntent_ClosingSettingPropertyList
	// Experimental.
	ClosingSettingInput() interface{}
	// Experimental.
	ConfirmationSetting() AwsIntent_ConfirmationSettingPropertyList
	// Experimental.
	ConfirmationSettingInput() interface{}
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
	CreationDateTime() *string
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
	DialogCodeHook() AwsIntent_DialogCodeHookPropertyList
	// Experimental.
	DialogCodeHookInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FulfillmentCodeHook() AwsIntent_FulfillmentCodeHookPropertyList
	// Experimental.
	FulfillmentCodeHookInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	InitialResponseSetting() AwsIntent_InitialResponseSettingPropertyList
	// Experimental.
	InitialResponseSettingInput() interface{}
	// Experimental.
	InputContext() AwsIntent_InputContextPropertyList
	// Experimental.
	InputContextInput() interface{}
	// Experimental.
	IntentId() *string
	// Experimental.
	KendraConfiguration() AwsIntent_KendraConfigurationPropertyList
	// Experimental.
	KendraConfigurationInput() interface{}
	// Experimental.
	LastUpdatedDateTime() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LocaleId() *string
	// Experimental.
	SetLocaleId(val *string)
	// Experimental.
	LocaleIdInput() *string
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
	OutputContext() AwsIntent_OutputContextPropertyList
	// Experimental.
	OutputContextInput() interface{}
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
	QnaIntentConfiguration() AwsIntent_QnaIntentConfigurationPropertyList
	// Experimental.
	QnaIntentConfigurationInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	SampleUtterance() AwsIntent_SampleUtterancePropertyList
	// Experimental.
	SampleUtteranceInput() interface{}
	// Experimental.
	SlotPriority() AwsIntent_SlotPriorityPropertyList
	// Experimental.
	SlotPriorityInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsIntent_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
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
	PutClosingSetting(value interface{})
	// Experimental.
	PutConfirmationSetting(value interface{})
	// Experimental.
	PutDialogCodeHook(value interface{})
	// Experimental.
	PutFulfillmentCodeHook(value interface{})
	// Experimental.
	PutInitialResponseSetting(value interface{})
	// Experimental.
	PutInputContext(value interface{})
	// Experimental.
	PutKendraConfiguration(value interface{})
	// Experimental.
	PutOutputContext(value interface{})
	// Experimental.
	PutQnaIntentConfiguration(value interface{})
	// Experimental.
	PutSampleUtterance(value interface{})
	// Experimental.
	PutSlotPriority(value interface{})
	// Experimental.
	PutTimeouts(value *AwsIntent_TimeoutsProperty)
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
	ResetClosingSetting()
	// Experimental.
	ResetConfirmationSetting()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetDialogCodeHook()
	// Experimental.
	ResetFulfillmentCodeHook()
	// Experimental.
	ResetInitialResponseSetting()
	// Experimental.
	ResetInputContext()
	// Experimental.
	ResetKendraConfiguration()
	// Experimental.
	ResetOutputContext()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParentIntentSignature()
	// Experimental.
	ResetQnaIntentConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSampleUtterance()
	// Experimental.
	ResetSlotPriority()
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

// The jsii proxy struct for AwsIntent
type jsiiProxy_AwsIntent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsIntent) BotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) BotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) BotVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) BotVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"botVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ClosingSetting() AwsIntent_ClosingSettingPropertyList {
	var returns AwsIntent_ClosingSettingPropertyList
	_jsii_.Get(
		j,
		"closingSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ClosingSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"closingSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ConfirmationSetting() AwsIntent_ConfirmationSettingPropertyList {
	var returns AwsIntent_ConfirmationSettingPropertyList
	_jsii_.Get(
		j,
		"confirmationSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ConfirmationSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confirmationSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) CreationDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"creationDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) DialogCodeHook() AwsIntent_DialogCodeHookPropertyList {
	var returns AwsIntent_DialogCodeHookPropertyList
	_jsii_.Get(
		j,
		"dialogCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) DialogCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dialogCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) FulfillmentCodeHook() AwsIntent_FulfillmentCodeHookPropertyList {
	var returns AwsIntent_FulfillmentCodeHookPropertyList
	_jsii_.Get(
		j,
		"fulfillmentCodeHook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) FulfillmentCodeHookInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fulfillmentCodeHookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) InitialResponseSetting() AwsIntent_InitialResponseSettingPropertyList {
	var returns AwsIntent_InitialResponseSettingPropertyList
	_jsii_.Get(
		j,
		"initialResponseSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) InitialResponseSettingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"initialResponseSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) InputContext() AwsIntent_InputContextPropertyList {
	var returns AwsIntent_InputContextPropertyList
	_jsii_.Get(
		j,
		"inputContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) InputContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) IntentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) KendraConfiguration() AwsIntent_KendraConfigurationPropertyList {
	var returns AwsIntent_KendraConfigurationPropertyList
	_jsii_.Get(
		j,
		"kendraConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) KendraConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kendraConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) LastUpdatedDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdatedDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) LocaleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) LocaleIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) OutputContext() AwsIntent_OutputContextPropertyList {
	var returns AwsIntent_OutputContextPropertyList
	_jsii_.Get(
		j,
		"outputContext",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) OutputContextInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputContextInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ParentIntentSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) ParentIntentSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIntentSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) QnaIntentConfiguration() AwsIntent_QnaIntentConfigurationPropertyList {
	var returns AwsIntent_QnaIntentConfigurationPropertyList
	_jsii_.Get(
		j,
		"qnaIntentConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) QnaIntentConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"qnaIntentConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) SampleUtterance() AwsIntent_SampleUtterancePropertyList {
	var returns AwsIntent_SampleUtterancePropertyList
	_jsii_.Get(
		j,
		"sampleUtterance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) SampleUtteranceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sampleUtteranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) SlotPriority() AwsIntent_SlotPriorityPropertyList {
	var returns AwsIntent_SlotPriorityPropertyList
	_jsii_.Get(
		j,
		"slotPriority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) SlotPriorityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"slotPriorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) Timeouts() AwsIntent_TimeoutsPropertyOutputReference {
	var returns AwsIntent_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsIntent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent aws_lexv2models_intent} Resource.
// Experimental.
func NewAwsIntent(scope constructs.Construct, id *string, config *AwsIntentConfig) AwsIntent {
	_init_.Initialize()

	if err := validateNewAwsIntentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsIntent{}

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent aws_lexv2models_intent} Resource.
// Experimental.
func NewAwsIntent_Override(a AwsIntent, scope constructs.Construct, id *string, config *AwsIntentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsIntent)SetBotId(val *string) {
	if err := j.validateSetBotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botId",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetBotVersion(val *string) {
	if err := j.validateSetBotVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"botVersion",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetLocaleId(val *string) {
	if err := j.validateSetLocaleIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localeId",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetParentIntentSignature(val *string) {
	if err := j.validateSetParentIntentSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentIntentSignature",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsIntent)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

// Generates CDKTN code for importing a AwsIntent resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsIntent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsIntent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-v2-models.AwsIntent",
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
func AwsIntent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIntent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsIntent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIntent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsIntent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsIntent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsIntent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-lex-v2-models.AwsIntent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsIntent) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsIntent) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsIntent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsIntent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsIntent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsIntent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsIntent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsIntent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsIntent) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsIntent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsIntent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsIntent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsIntent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsIntent) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsIntent) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsIntent) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsIntent) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsIntent) PutClosingSetting(value interface{}) {
	if err := a.validatePutClosingSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putClosingSetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutConfirmationSetting(value interface{}) {
	if err := a.validatePutConfirmationSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConfirmationSetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutDialogCodeHook(value interface{}) {
	if err := a.validatePutDialogCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDialogCodeHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutFulfillmentCodeHook(value interface{}) {
	if err := a.validatePutFulfillmentCodeHookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFulfillmentCodeHook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutInitialResponseSetting(value interface{}) {
	if err := a.validatePutInitialResponseSettingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInitialResponseSetting",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutInputContext(value interface{}) {
	if err := a.validatePutInputContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputContext",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutKendraConfiguration(value interface{}) {
	if err := a.validatePutKendraConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKendraConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutOutputContext(value interface{}) {
	if err := a.validatePutOutputContextParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputContext",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutQnaIntentConfiguration(value interface{}) {
	if err := a.validatePutQnaIntentConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putQnaIntentConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutSampleUtterance(value interface{}) {
	if err := a.validatePutSampleUtteranceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSampleUtterance",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutSlotPriority(value interface{}) {
	if err := a.validatePutSlotPriorityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSlotPriority",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) PutTimeouts(value *AwsIntent_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsIntent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsIntent) ResetClosingSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetClosingSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetConfirmationSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetConfirmationSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetDialogCodeHook() {
	_jsii_.InvokeVoid(
		a,
		"resetDialogCodeHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetFulfillmentCodeHook() {
	_jsii_.InvokeVoid(
		a,
		"resetFulfillmentCodeHook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetInitialResponseSetting() {
	_jsii_.InvokeVoid(
		a,
		"resetInitialResponseSetting",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetInputContext() {
	_jsii_.InvokeVoid(
		a,
		"resetInputContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetKendraConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetKendraConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetOutputContext() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputContext",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetParentIntentSignature() {
	_jsii_.InvokeVoid(
		a,
		"resetParentIntentSignature",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetQnaIntentConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetQnaIntentConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetSampleUtterance() {
	_jsii_.InvokeVoid(
		a,
		"resetSampleUtterance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetSlotPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetSlotPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsIntent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsIntent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

