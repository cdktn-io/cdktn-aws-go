package bedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/bedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime}.
// Experimental.
type AwsAgentRuntime interface {
	cdktn.TerraformResource
	// Experimental.
	AgentRuntimeArn() *string
	// Experimental.
	AgentRuntimeArtifact() AwsAgentRuntime_AgentRuntimeArtifactPropertyList
	// Experimental.
	AgentRuntimeArtifactInput() interface{}
	// Experimental.
	AgentRuntimeId() *string
	// Experimental.
	AgentRuntimeName() *string
	// Experimental.
	SetAgentRuntimeName(val *string)
	// Experimental.
	AgentRuntimeNameInput() *string
	// Experimental.
	AgentRuntimeVersion() *string
	// Experimental.
	AuthorizerConfiguration() AwsAgentRuntime_AuthorizerConfigurationPropertyList
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
	Description() *string
	// Experimental.
	SetDescription(val *string)
	// Experimental.
	DescriptionInput() *string
	// Experimental.
	EnvironmentVariables() *map[string]*string
	// Experimental.
	SetEnvironmentVariables(val *map[string]*string)
	// Experimental.
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	FilesystemConfiguration() AwsAgentRuntime_FilesystemConfigurationPropertyList
	// Experimental.
	FilesystemConfigurationInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	LifecycleConfiguration() AwsAgentRuntime_LifecycleConfigurationPropertyList
	// Experimental.
	LifecycleConfigurationInput() interface{}
	// Experimental.
	NetworkConfiguration() AwsAgentRuntime_NetworkConfigurationPropertyList
	// Experimental.
	NetworkConfigurationInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	ProtocolConfiguration() AwsAgentRuntime_ProtocolConfigurationPropertyList
	// Experimental.
	ProtocolConfigurationInput() interface{}
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
	RequestHeaderConfiguration() AwsAgentRuntime_RequestHeaderConfigurationPropertyList
	// Experimental.
	RequestHeaderConfigurationInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
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
	Timeouts() AwsAgentRuntime_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	WorkloadIdentityDetails() AwsAgentRuntime_WorkloadIdentityDetailsPropertyList
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
	PutAgentRuntimeArtifact(value interface{})
	// Experimental.
	PutAuthorizerConfiguration(value interface{})
	// Experimental.
	PutFilesystemConfiguration(value interface{})
	// Experimental.
	PutLifecycleConfiguration(value interface{})
	// Experimental.
	PutNetworkConfiguration(value interface{})
	// Experimental.
	PutProtocolConfiguration(value interface{})
	// Experimental.
	PutRequestHeaderConfiguration(value interface{})
	// Experimental.
	PutTimeouts(value *AwsAgentRuntime_TimeoutsProperty)
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
	ResetAgentRuntimeArtifact()
	// Experimental.
	ResetAuthorizerConfiguration()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEnvironmentVariables()
	// Experimental.
	ResetFilesystemConfiguration()
	// Experimental.
	ResetLifecycleConfiguration()
	// Experimental.
	ResetNetworkConfiguration()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProtocolConfiguration()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRequestHeaderConfiguration()
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

// The jsii proxy struct for AwsAgentRuntime
type jsiiProxy_AwsAgentRuntime struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeArtifact() AwsAgentRuntime_AgentRuntimeArtifactPropertyList {
	var returns AwsAgentRuntime_AgentRuntimeArtifactPropertyList
	_jsii_.Get(
		j,
		"agentRuntimeArtifact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeArtifactInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"agentRuntimeArtifactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AgentRuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agentRuntimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AuthorizerConfiguration() AwsAgentRuntime_AuthorizerConfigurationPropertyList {
	var returns AwsAgentRuntime_AuthorizerConfigurationPropertyList
	_jsii_.Get(
		j,
		"authorizerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) AuthorizerConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authorizerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) FilesystemConfiguration() AwsAgentRuntime_FilesystemConfigurationPropertyList {
	var returns AwsAgentRuntime_FilesystemConfigurationPropertyList
	_jsii_.Get(
		j,
		"filesystemConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) FilesystemConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"filesystemConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) LifecycleConfiguration() AwsAgentRuntime_LifecycleConfigurationPropertyList {
	var returns AwsAgentRuntime_LifecycleConfigurationPropertyList
	_jsii_.Get(
		j,
		"lifecycleConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) LifecycleConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lifecycleConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) NetworkConfiguration() AwsAgentRuntime_NetworkConfigurationPropertyList {
	var returns AwsAgentRuntime_NetworkConfigurationPropertyList
	_jsii_.Get(
		j,
		"networkConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) NetworkConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) ProtocolConfiguration() AwsAgentRuntime_ProtocolConfigurationPropertyList {
	var returns AwsAgentRuntime_ProtocolConfigurationPropertyList
	_jsii_.Get(
		j,
		"protocolConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) ProtocolConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protocolConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RequestHeaderConfiguration() AwsAgentRuntime_RequestHeaderConfigurationPropertyList {
	var returns AwsAgentRuntime_RequestHeaderConfigurationPropertyList
	_jsii_.Get(
		j,
		"requestHeaderConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RequestHeaderConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) Timeouts() AwsAgentRuntime_TimeoutsPropertyOutputReference {
	var returns AwsAgentRuntime_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAgentRuntime) WorkloadIdentityDetails() AwsAgentRuntime_WorkloadIdentityDetailsPropertyList {
	var returns AwsAgentRuntime_WorkloadIdentityDetailsPropertyList
	_jsii_.Get(
		j,
		"workloadIdentityDetails",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime} Resource.
// Experimental.
func NewAwsAgentRuntime(scope constructs.Construct, id *string, config *AwsAgentRuntimeConfig) AwsAgentRuntime {
	_init_.Initialize()

	if err := validateNewAwsAgentRuntimeParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAgentRuntime{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_agent_runtime aws_bedrockagentcore_agent_runtime} Resource.
// Experimental.
func NewAwsAgentRuntime_Override(a AwsAgentRuntime, scope constructs.Construct, id *string, config *AwsAgentRuntimeConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetAgentRuntimeName(val *string) {
	if err := j.validateSetAgentRuntimeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agentRuntimeName",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsAgentRuntime)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a AwsAgentRuntime resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsAgentRuntime_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsAgentRuntime_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
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
func AwsAgentRuntime_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAgentRuntime_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAgentRuntime_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAgentRuntime_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsAgentRuntime_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsAgentRuntime_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAgentRuntime_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock-agentcore.AwsAgentRuntime",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAgentRuntime) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAgentRuntime) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAgentRuntime) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAgentRuntime) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAgentRuntime) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAgentRuntime) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAgentRuntime) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAgentRuntime) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAgentRuntime) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAgentRuntime) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsAgentRuntime) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutAgentRuntimeArtifact(value interface{}) {
	if err := a.validatePutAgentRuntimeArtifactParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAgentRuntimeArtifact",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutAuthorizerConfiguration(value interface{}) {
	if err := a.validatePutAuthorizerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthorizerConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutFilesystemConfiguration(value interface{}) {
	if err := a.validatePutFilesystemConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilesystemConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutLifecycleConfiguration(value interface{}) {
	if err := a.validatePutLifecycleConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLifecycleConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutNetworkConfiguration(value interface{}) {
	if err := a.validatePutNetworkConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNetworkConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutProtocolConfiguration(value interface{}) {
	if err := a.validatePutProtocolConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProtocolConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutRequestHeaderConfiguration(value interface{}) {
	if err := a.validatePutRequestHeaderConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRequestHeaderConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) PutTimeouts(value *AwsAgentRuntime_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetAgentRuntimeArtifact() {
	_jsii_.InvokeVoid(
		a,
		"resetAgentRuntimeArtifact",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetAuthorizerConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizerConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetFilesystemConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFilesystemConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetLifecycleConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetLifecycleConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetNetworkConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetNetworkConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetProtocolConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetProtocolConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetRequestHeaderConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestHeaderConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAgentRuntime) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAgentRuntime) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

