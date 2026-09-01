package awsbedrockagents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagents/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent}.
// Experimental.
type AwsBedrockagentAgent interface {
	cdktn.TerraformResource
	// Experimental.
	AgentArn() *string
	// Experimental.
	AgentCollaboration() *string
	// Experimental.
	SetAgentCollaboration(val *string)
	// Experimental.
	AgentCollaborationInput() *string
	// Experimental.
	AgentId() *string
	// Experimental.
	AgentName() *string
	// Experimental.
	SetAgentName(val *string)
	// Experimental.
	AgentNameInput() *string
	// Experimental.
	AgentResourceRoleArn() *string
	// Experimental.
	SetAgentResourceRoleArn(val *string)
	// Experimental.
	AgentResourceRoleArnInput() *string
	// Experimental.
	AgentVersion() *string
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
	CustomerEncryptionKeyArn() *string
	// Experimental.
	SetCustomerEncryptionKeyArn(val *string)
	// Experimental.
	CustomerEncryptionKeyArnInput() *string
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
	FoundationModel() *string
	// Experimental.
	SetFoundationModel(val *string)
	// Experimental.
	FoundationModelInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	GuardrailConfiguration() AwsBedrockagentAgent_GuardrailConfigurationPropertyList
	// Experimental.
	GuardrailConfigurationInput() interface{}
	// Experimental.
	Id() *string
	// Experimental.
	IdleSessionTtlInSeconds() *float64
	// Experimental.
	SetIdleSessionTtlInSeconds(val *float64)
	// Experimental.
	IdleSessionTtlInSecondsInput() *float64
	// Experimental.
	Instruction() *string
	// Experimental.
	SetInstruction(val *string)
	// Experimental.
	InstructionInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MemoryConfiguration() AwsBedrockagentAgent_MemoryConfigurationPropertyList
	// Experimental.
	MemoryConfigurationInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	PrepareAgent() interface{}
	// Experimental.
	SetPrepareAgent(val interface{})
	// Experimental.
	PrepareAgentInput() interface{}
	// Experimental.
	PreparedAt() *string
	// Experimental.
	PromptOverrideConfiguration() AwsBedrockagentAgent_PromptOverrideConfigurationPropertyList
	// Experimental.
	PromptOverrideConfigurationInput() interface{}
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
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsBedrockagentAgent_TimeoutsPropertyOutputReference
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
	PutGuardrailConfiguration(value interface{})
	// Experimental.
	PutMemoryConfiguration(value interface{})
	// Experimental.
	PutPromptOverrideConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *AwsBedrockagentAgent_TimeoutsProperty)
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
	ResetAgentCollaboration()
	// Experimental.
	ResetCustomerEncryptionKeyArn()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetGuardrailConfiguration()
	// Experimental.
	ResetIdleSessionTtlInSeconds()
	// Experimental.
	ResetInstruction()
	// Experimental.
	ResetMemoryConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetPrepareAgent()
	// Experimental.
	ResetPromptOverrideConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSkipResourceInUseCheck()
	// Experimental.
	ResetTags()
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

// The jsii proxy struct for AwsBedrockagentAgent
type jsiiProxy_AwsBedrockagentAgent struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentCollaboration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaboration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentCollaborationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentCollaborationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentResourceRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentResourceRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentResourceRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) AgentVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) CustomerEncryptionKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) CustomerEncryptionKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customerEncryptionKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) FoundationModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) FoundationModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"foundationModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) GuardrailConfiguration() AwsBedrockagentAgent_GuardrailConfigurationPropertyList {
	var returns AwsBedrockagentAgent_GuardrailConfigurationPropertyList
	_jsii_.Get(
		j,
		"guardrailConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) GuardrailConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"guardrailConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) IdleSessionTtlInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) IdleSessionTtlInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"idleSessionTtlInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Instruction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instruction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) InstructionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instructionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) MemoryConfiguration() AwsBedrockagentAgent_MemoryConfigurationPropertyList {
	var returns AwsBedrockagentAgent_MemoryConfigurationPropertyList
	_jsii_.Get(
		j,
		"memoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) MemoryConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) PrepareAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) PrepareAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prepareAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) PreparedAt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preparedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) PromptOverrideConfiguration() AwsBedrockagentAgent_PromptOverrideConfigurationPropertyList {
	var returns AwsBedrockagentAgent_PromptOverrideConfigurationPropertyList
	_jsii_.Get(
		j,
		"promptOverrideConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) PromptOverrideConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"promptOverrideConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) SkipResourceInUseCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) SkipResourceInUseCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipResourceInUseCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) Timeouts() AwsBedrockagentAgent_TimeoutsPropertyOutputReference {
	var returns AwsBedrockagentAgent_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentAgent) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent} Resource.
// Experimental.
func NewAwsBedrockagentAgent(scope constructs.Construct, id *string, config *AwsBedrockagentAgentConfig) AwsBedrockagentAgent {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentAgentParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentAgent{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent aws_bedrockagent_agent} Resource.
// Experimental.
func NewAwsBedrockagentAgent_Override(a AwsBedrockagentAgent, scope constructs.Construct, id *string, config *AwsBedrockagentAgentConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetAgentCollaboration(val *string) {
	if err := j.validateSetAgentCollaborationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentCollaboration",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetAgentName(val *string) {
	if err := j.validateSetAgentNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentName",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetAgentResourceRoleArn(val *string) {
	if err := j.validateSetAgentResourceRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentResourceRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetCustomerEncryptionKeyArn(val *string) {
	if err := j.validateSetCustomerEncryptionKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customerEncryptionKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetFoundationModel(val *string) {
	if err := j.validateSetFoundationModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"foundationModel",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetIdleSessionTtlInSeconds(val *float64) {
	if err := j.validateSetIdleSessionTtlInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"idleSessionTtlInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetInstruction(val *string) {
	if err := j.validateSetInstructionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instruction",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetPrepareAgent(val interface{}) {
	if err := j.validateSetPrepareAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prepareAgent",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetSkipResourceInUseCheck(val interface{}) {
	if err := j.validateSetSkipResourceInUseCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipResourceInUseCheck",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentAgent)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a AwsBedrockagentAgent resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsBedrockagentAgent_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsBedrockagentAgent_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
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
func AwsBedrockagentAgent_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentAgent_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockagentAgent_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentAgent_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockagentAgent_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentAgent_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsBedrockagentAgent_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock-agents.AwsBedrockagentAgent",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentAgent) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentAgent) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentAgent) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentAgent) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) PutGuardrailConfiguration(value interface{}) {
	if err := a.validatePutGuardrailConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGuardrailConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) PutMemoryConfiguration(value interface{}) {
	if err := a.validatePutMemoryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMemoryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) PutPromptOverrideConfiguration(value interface{}) {
	if err := a.validatePutPromptOverrideConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPromptOverrideConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) PutTimeouts(value *AwsBedrockagentAgent_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetAgentCollaboration() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentCollaboration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetCustomerEncryptionKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomerEncryptionKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetGuardrailConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetGuardrailConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetIdleSessionTtlInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetIdleSessionTtlInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetInstruction() {
	_jsii_.InvokeVoid(
		a,
		"resetInstruction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetMemoryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetMemoryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetPrepareAgent() {
	_jsii_.InvokeVoid(
		a,
		"resetPrepareAgent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetPromptOverrideConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetPromptOverrideConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetSkipResourceInUseCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipResourceInUseCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentAgent) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentAgent) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

