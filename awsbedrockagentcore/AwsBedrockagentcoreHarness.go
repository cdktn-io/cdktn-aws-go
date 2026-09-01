package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness}.
// Experimental.
type AwsBedrockagentcoreHarness interface {
	cdktn.TerraformResource
	// Experimental.
	AllowedTools() *[]*string
	// Experimental.
	SetAllowedTools(val *[]*string)
	// Experimental.
	AllowedToolsInput() *[]*string
	// Experimental.
	Arn() *string
	// Experimental.
	AuthorizerConfiguration() AwsBedrockagentcoreHarness_AuthorizerConfigurationPropertyList
	// Experimental.
	AuthorizerConfigurationInput() interface{}
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
	Environment() AwsBedrockagentcoreHarness_EnvironmentPropertyList
	// Experimental.
	EnvironmentActual() AwsBedrockagentcoreHarness_EnvironmentActualPropertyList
	// Experimental.
	EnvironmentArtifact() AwsBedrockagentcoreHarness_EnvironmentArtifactPropertyList
	// Experimental.
	EnvironmentArtifactInput() interface{}
	// Experimental.
	EnvironmentInput() interface{}
	// Experimental.
	EnvironmentVariables() *map[string]*string
	// Experimental.
	SetEnvironmentVariables(val *map[string]*string)
	// Experimental.
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	ExecutionRoleArn() *string
	// Experimental.
	SetExecutionRoleArn(val *string)
	// Experimental.
	ExecutionRoleArnInput() *string
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HarnessId() *string
	// Experimental.
	HarnessName() *string
	// Experimental.
	SetHarnessName(val *string)
	// Experimental.
	HarnessNameInput() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MaxIterations() *float64
	// Experimental.
	SetMaxIterations(val *float64)
	// Experimental.
	MaxIterationsInput() *float64
	// Experimental.
	MaxTokens() *float64
	// Experimental.
	SetMaxTokens(val *float64)
	// Experimental.
	MaxTokensInput() *float64
	// Experimental.
	Memory() AwsBedrockagentcoreHarness_MemoryPropertyList
	// Experimental.
	MemoryActual() AwsBedrockagentcoreHarness_MemoryActualPropertyList
	// Experimental.
	MemoryInput() interface{}
	// Experimental.
	Model() AwsBedrockagentcoreHarness_ModelPropertyList
	// Experimental.
	ModelInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
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
	Skill() AwsBedrockagentcoreHarness_SkillPropertyList
	// Experimental.
	SkillInput() interface{}
	// Experimental.
	SystemPrompt() AwsBedrockagentcoreHarness_SystemPromptPropertyList
	// Experimental.
	SystemPromptInput() interface{}
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
	Timeouts() AwsBedrockagentcoreHarness_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutSeconds() *float64
	// Experimental.
	SetTimeoutSeconds(val *float64)
	// Experimental.
	TimeoutSecondsInput() *float64
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	Tool() AwsBedrockagentcoreHarness_ToolPropertyList
	// Experimental.
	ToolInput() interface{}
	// Experimental.
	Truncation() AwsBedrockagentcoreHarness_TruncationPropertyList
	// Experimental.
	TruncationInput() interface{}
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
	PutAuthorizerConfiguration(value interface{})
	// Experimental.
	PutEnvironment(value interface{})
	// Experimental.
	PutEnvironmentArtifact(value interface{})
	// Experimental.
	PutMemory(value interface{})
	// Experimental.
	PutModel(value interface{})
	// Experimental.
	PutSkill(value interface{})
	// Experimental.
	PutSystemPrompt(value interface{})
	// Experimental.
	PutTimeouts(value *AwsBedrockagentcoreHarness_TimeoutsProperty)
	// Experimental.
	PutTool(value interface{})
	// Experimental.
	PutTruncation(value interface{})
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
	ResetAllowedTools()
	// Experimental.
	ResetAuthorizerConfiguration()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetEnvironmentArtifact()
	// Experimental.
	ResetEnvironmentVariables()
	// Experimental.
	ResetMaxIterations()
	// Experimental.
	ResetMaxTokens()
	// Experimental.
	ResetMemory()
	// Experimental.
	ResetModel()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetSkill()
	// Experimental.
	ResetSystemPrompt()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTimeoutSeconds()
	// Experimental.
	ResetTool()
	// Experimental.
	ResetTruncation()
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

// The jsii proxy struct for AwsBedrockagentcoreHarness
type jsiiProxy_AwsBedrockagentcoreHarness struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) AllowedTools() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedTools",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) AllowedToolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedToolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) AuthorizerConfiguration() AwsBedrockagentcoreHarness_AuthorizerConfigurationPropertyList {
	var returns AwsBedrockagentcoreHarness_AuthorizerConfigurationPropertyList
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) AuthorizerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Environment() AwsBedrockagentcoreHarness_EnvironmentPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentPropertyList
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentActual() AwsBedrockagentcoreHarness_EnvironmentActualPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentActualPropertyList
	_jsii_.Get(
		j,
		"environmentActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentArtifact() AwsBedrockagentcoreHarness_EnvironmentArtifactPropertyList {
	var returns AwsBedrockagentcoreHarness_EnvironmentArtifactPropertyList
	_jsii_.Get(
		j,
		"environmentArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) HarnessId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) HarnessName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) HarnessNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"harnessNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MaxIterations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIterations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MaxIterationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIterationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MaxTokens() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokens",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MaxTokensInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTokensInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Memory() AwsBedrockagentcoreHarness_MemoryPropertyList {
	var returns AwsBedrockagentcoreHarness_MemoryPropertyList
	_jsii_.Get(
		j,
		"memory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MemoryActual() AwsBedrockagentcoreHarness_MemoryActualPropertyList {
	var returns AwsBedrockagentcoreHarness_MemoryActualPropertyList
	_jsii_.Get(
		j,
		"memoryActual",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) MemoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"memoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Model() AwsBedrockagentcoreHarness_ModelPropertyList {
	var returns AwsBedrockagentcoreHarness_ModelPropertyList
	_jsii_.Get(
		j,
		"model",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ModelInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Skill() AwsBedrockagentcoreHarness_SkillPropertyList {
	var returns AwsBedrockagentcoreHarness_SkillPropertyList
	_jsii_.Get(
		j,
		"skill",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) SkillInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skillInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) SystemPrompt() AwsBedrockagentcoreHarness_SystemPromptPropertyList {
	var returns AwsBedrockagentcoreHarness_SystemPromptPropertyList
	_jsii_.Get(
		j,
		"systemPrompt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) SystemPromptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemPromptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Timeouts() AwsBedrockagentcoreHarness_TimeoutsPropertyOutputReference {
	var returns AwsBedrockagentcoreHarness_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Tool() AwsBedrockagentcoreHarness_ToolPropertyList {
	var returns AwsBedrockagentcoreHarness_ToolPropertyList
	_jsii_.Get(
		j,
		"tool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) ToolInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"toolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) Truncation() AwsBedrockagentcoreHarness_TruncationPropertyList {
	var returns AwsBedrockagentcoreHarness_TruncationPropertyList
	_jsii_.Get(
		j,
		"truncation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness) TruncationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"truncationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness} Resource.
// Experimental.
func NewAwsBedrockagentcoreHarness(scope constructs.Construct, id *string, config *AwsBedrockagentcoreHarnessConfig) AwsBedrockagentcoreHarness {
	_init_.Initialize()

	if err := validateNewAwsBedrockagentcoreHarnessParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockagentcoreHarness{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_harness aws_bedrockagentcore_harness} Resource.
// Experimental.
func NewAwsBedrockagentcoreHarness_Override(a AwsBedrockagentcoreHarness, scope constructs.Construct, id *string, config *AwsBedrockagentcoreHarnessConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetAllowedTools(val *[]*string) {
	if err := j.validateSetAllowedToolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedTools",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetExecutionRoleArn(val *string) {
	if err := j.validateSetExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRoleArn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetHarnessName(val *string) {
	if err := j.validateSetHarnessNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"harnessName",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetMaxIterations(val *float64) {
	if err := j.validateSetMaxIterationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIterations",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetMaxTokens(val *float64) {
	if err := j.validateSetMaxTokensParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTokens",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockagentcoreHarness)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

// Generates CDKTN code for importing a AwsBedrockagentcoreHarness resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsBedrockagentcoreHarness_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsBedrockagentcoreHarness_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
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
func AwsBedrockagentcoreHarness_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentcoreHarness_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockagentcoreHarness_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentcoreHarness_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockagentcoreHarness_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockagentcoreHarness_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsBedrockagentcoreHarness_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock-agentcore.AwsBedrockagentcoreHarness",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsBedrockagentcoreHarness) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutAuthorizerConfiguration(value interface{}) {
	if err := a.validatePutAuthorizerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthorizerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutEnvironment(value interface{}) {
	if err := a.validatePutEnvironmentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironment",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutEnvironmentArtifact(value interface{}) {
	if err := a.validatePutEnvironmentArtifactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentArtifact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutMemory(value interface{}) {
	if err := a.validatePutMemoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMemory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutModel(value interface{}) {
	if err := a.validatePutModelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putModel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutSkill(value interface{}) {
	if err := a.validatePutSkillParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSkill",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutSystemPrompt(value interface{}) {
	if err := a.validatePutSystemPromptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSystemPrompt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutTimeouts(value *AwsBedrockagentcoreHarness_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutTool(value interface{}) {
	if err := a.validatePutToolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) PutTruncation(value interface{}) {
	if err := a.validatePutTruncationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTruncation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetAllowedTools() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedTools",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetAuthorizerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetEnvironmentArtifact() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentArtifact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetMaxIterations() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIterations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetMaxTokens() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTokens",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetMemory() {
	_jsii_.InvokeVoid(
		a,
		"resetMemory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetModel() {
	_jsii_.InvokeVoid(
		a,
		"resetModel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetSkill() {
	_jsii_.InvokeVoid(
		a,
		"resetSkill",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetSystemPrompt() {
	_jsii_.InvokeVoid(
		a,
		"resetSystemPrompt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetTool() {
	_jsii_.InvokeVoid(
		a,
		"resetTool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ResetTruncation() {
	_jsii_.InvokeVoid(
		a,
		"resetTruncation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockagentcoreHarness) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

