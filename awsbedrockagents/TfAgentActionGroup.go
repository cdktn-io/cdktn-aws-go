package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group}.
// Experimental.
type TfAgentActionGroup interface {
	cdktn.TerraformResource
	// Experimental.
	ActionGroupExecutor() TfAgentActionGroup_ActionGroupExecutorPropertyList
	// Experimental.
	ActionGroupExecutorInput() interface{}
	// Experimental.
	ActionGroupId() *string
	// Experimental.
	ActionGroupName() *string
	// Experimental.
	SetActionGroupName(val *string)
	// Experimental.
	ActionGroupNameInput() *string
	// Experimental.
	ActionGroupState() *string
	// Experimental.
	SetActionGroupState(val *string)
	// Experimental.
	ActionGroupStateInput() *string
	// Experimental.
	AgentId() *string
	// Experimental.
	SetAgentId(val *string)
	// Experimental.
	AgentIdInput() *string
	// Experimental.
	AgentVersion() *string
	// Experimental.
	SetAgentVersion(val *string)
	// Experimental.
	AgentVersionInput() *string
	// Experimental.
	ApiSchema() TfAgentActionGroup_ApiSchemaPropertyList
	// Experimental.
	ApiSchemaInput() interface{}
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	FunctionSchema() TfAgentActionGroup_FunctionSchemaPropertyList
	// Experimental.
	FunctionSchemaInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ParentActionGroupSignature() *string
	// Experimental.
	SetParentActionGroupSignature(val *string)
	// Experimental.
	ParentActionGroupSignatureInput() *string
	// Experimental.
	PrepareAgent() interface{}
	// Experimental.
	SetPrepareAgent(val interface{})
	// Experimental.
	PrepareAgentInput() interface{}
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
	SkipResourceInUseCheck() interface{}
	// Experimental.
	SetSkipResourceInUseCheck(val interface{})
	// Experimental.
	SkipResourceInUseCheckInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfAgentActionGroup_TimeoutsPropertyOutputReference
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
	PutActionGroupExecutor(value interface{})
	// Experimental.
	PutApiSchema(value interface{})
	// Experimental.
	PutFunctionSchema(value interface{})
	// Experimental.
	PutTimeouts(value *TfAgentActionGroup_TimeoutsProperty)
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
	ResetActionGroupExecutor()
	// Experimental.
	ResetActionGroupState()
	// Experimental.
	ResetApiSchema()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetFunctionSchema()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetParentActionGroupSignature()
	// Experimental.
	ResetPrepareAgent()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSkipResourceInUseCheck()
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

// The jsii proxy struct for TfAgentActionGroup
type jsiiProxy_TfAgentActionGroup struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupExecutor() TfAgentActionGroup_ActionGroupExecutorPropertyList {
	var returns TfAgentActionGroup_ActionGroupExecutorPropertyList
	_jsii_.Get(
		j,
		"actionGroupExecutor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupExecutorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"actionGroupExecutorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ActionGroupStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"actionGroupStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) AgentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) AgentVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ApiSchema() TfAgentActionGroup_ApiSchemaPropertyList {
	var returns TfAgentActionGroup_ApiSchemaPropertyList
	_jsii_.Get(
		j,
		"apiSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ApiSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) FunctionSchema() TfAgentActionGroup_FunctionSchemaPropertyList {
	var returns TfAgentActionGroup_FunctionSchemaPropertyList
	_jsii_.Get(
		j,
		"functionSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) FunctionSchemaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"functionSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ParentActionGroupSignature() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignature",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) ParentActionGroupSignatureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentActionGroupSignatureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) PrepareAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) PrepareAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) SkipResourceInUseCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) SkipResourceInUseCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) Timeouts() TfAgentActionGroup_TimeoutsPropertyOutputReference {
	var returns TfAgentActionGroup_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAgentActionGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group} Resource.
// Experimental.
func NewTfAgentActionGroup(scope constructs.Construct, id *string, config *TfAgentActionGroupConfig) TfAgentActionGroup {
	_init_.Initialize()

	if err := validateNewTfAgentActionGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAgentActionGroup{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent_action_group aws_bedrockagent_agent_action_group} Resource.
// Experimental.
func NewTfAgentActionGroup_Override(t TfAgentActionGroup, scope constructs.Construct, id *string, config *TfAgentActionGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetActionGroupName(val *string) {
	if err := j.validateSetActionGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupName",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetActionGroupState(val *string) {
	if err := j.validateSetActionGroupStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"actionGroupState",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetAgentId(val *string) {
	if err := j.validateSetAgentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentId",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetAgentVersion(val *string) {
	if err := j.validateSetAgentVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentVersion",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetParentActionGroupSignature(val *string) {
	if err := j.validateSetParentActionGroupSignatureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentActionGroupSignature",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetPrepareAgent(val interface{}) {
	if err := j.validateSetPrepareAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prepareAgent",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfAgentActionGroup)SetSkipResourceInUseCheck(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheck",
		val,
	)
}

// Generates CDKTN code for importing a TfAgentActionGroup resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfAgentActionGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfAgentActionGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
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
func TfAgentActionGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAgentActionGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAgentActionGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAgentActionGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfAgentActionGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfAgentActionGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfAgentActionGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock-agents.TfAgentActionGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAgentActionGroup) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAgentActionGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAgentActionGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAgentActionGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAgentActionGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAgentActionGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAgentActionGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAgentActionGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAgentActionGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAgentActionGroup) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfAgentActionGroup) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) PutActionGroupExecutor(value interface{}) {
	if err := t.validatePutActionGroupExecutorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActionGroupExecutor",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) PutApiSchema(value interface{}) {
	if err := t.validatePutApiSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) PutFunctionSchema(value interface{}) {
	if err := t.validatePutFunctionSchemaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFunctionSchema",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) PutTimeouts(value *TfAgentActionGroup_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetActionGroupExecutor() {
	_jsii_.InvokeVoid(
		t,
		"resetActionGroupExecutor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetActionGroupState() {
	_jsii_.InvokeVoid(
		t,
		"resetActionGroupState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetApiSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetApiSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetFunctionSchema() {
	_jsii_.InvokeVoid(
		t,
		"resetFunctionSchema",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetParentActionGroupSignature() {
	_jsii_.InvokeVoid(
		t,
		"resetParentActionGroupSignature",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetPrepareAgent() {
	_jsii_.InvokeVoid(
		t,
		"resetPrepareAgent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetSkipResourceInUseCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetSkipResourceInUseCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAgentActionGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAgentActionGroup) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

