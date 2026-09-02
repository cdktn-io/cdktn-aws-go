package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config aws_bedrockagentcore_online_evaluation_config}.
// Experimental.
type TfOnlineEvaluationConfig interface {
	cdktn.TerraformResource
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
	DataSourceConfig() TfOnlineEvaluationConfig_DataSourceConfigPropertyList
	// Experimental.
	DataSourceConfigInput() interface{}
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
	EnableOnCreate() interface{}
	// Experimental.
	SetEnableOnCreate(val interface{})
	// Experimental.
	EnableOnCreateInput() interface{}
	// Experimental.
	EvaluationExecutionRoleArn() *string
	// Experimental.
	SetEvaluationExecutionRoleArn(val *string)
	// Experimental.
	EvaluationExecutionRoleArnInput() *string
	// Experimental.
	Evaluator() TfOnlineEvaluationConfig_EvaluatorPropertyList
	// Experimental.
	EvaluatorInput() interface{}
	// Experimental.
	ExecutionStatus() *string
	// Experimental.
	SetExecutionStatus(val *string)
	// Experimental.
	ExecutionStatusInput() *string
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
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OnlineEvaluationConfigArn() *string
	// Experimental.
	OnlineEvaluationConfigId() *string
	// Experimental.
	OnlineEvaluationConfigName() *string
	// Experimental.
	SetOnlineEvaluationConfigName(val *string)
	// Experimental.
	OnlineEvaluationConfigNameInput() *string
	// Experimental.
	OutputConfig() TfOnlineEvaluationConfig_OutputConfigPropertyList
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
	Rule() TfOnlineEvaluationConfig_RulePropertyList
	// Experimental.
	RuleInput() interface{}
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
	Timeouts() TfOnlineEvaluationConfig_TimeoutsPropertyOutputReference
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
	PutDataSourceConfig(value interface{})
	// Experimental.
	PutEvaluator(value interface{})
	// Experimental.
	PutRule(value interface{})
	// Experimental.
	PutTimeouts(value *TfOnlineEvaluationConfig_TimeoutsProperty)
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
	ResetDataSourceConfig()
	// Experimental.
	ResetDescription()
	// Experimental.
	ResetEvaluator()
	// Experimental.
	ResetExecutionStatus()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRule()
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

// The jsii proxy struct for TfOnlineEvaluationConfig
type jsiiProxy_TfOnlineEvaluationConfig struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) DataSourceConfig() TfOnlineEvaluationConfig_DataSourceConfigPropertyList {
	var returns TfOnlineEvaluationConfig_DataSourceConfigPropertyList
	_jsii_.Get(
		j,
		"dataSourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) DataSourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataSourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) EnableOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) EnableOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) EvaluationExecutionRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationExecutionRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) EvaluationExecutionRoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evaluationExecutionRoleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Evaluator() TfOnlineEvaluationConfig_EvaluatorPropertyList {
	var returns TfOnlineEvaluationConfig_EvaluatorPropertyList
	_jsii_.Get(
		j,
		"evaluator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) EvaluatorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"evaluatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) ExecutionStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) ExecutionStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) OnlineEvaluationConfigArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) OnlineEvaluationConfigId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) OnlineEvaluationConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) OnlineEvaluationConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onlineEvaluationConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) OutputConfig() TfOnlineEvaluationConfig_OutputConfigPropertyList {
	var returns TfOnlineEvaluationConfig_OutputConfigPropertyList
	_jsii_.Get(
		j,
		"outputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Rule() TfOnlineEvaluationConfig_RulePropertyList {
	var returns TfOnlineEvaluationConfig_RulePropertyList
	_jsii_.Get(
		j,
		"rule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) RuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ruleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) Timeouts() TfOnlineEvaluationConfig_TimeoutsPropertyOutputReference {
	var returns TfOnlineEvaluationConfig_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOnlineEvaluationConfig) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config aws_bedrockagentcore_online_evaluation_config} Resource.
// Experimental.
func NewTfOnlineEvaluationConfig(scope constructs.Construct, id *string, config *TfOnlineEvaluationConfigConfig) TfOnlineEvaluationConfig {
	_init_.Initialize()

	if err := validateNewTfOnlineEvaluationConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOnlineEvaluationConfig{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config aws_bedrockagentcore_online_evaluation_config} Resource.
// Experimental.
func NewTfOnlineEvaluationConfig_Override(t TfOnlineEvaluationConfig, scope constructs.Construct, id *string, config *TfOnlineEvaluationConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetEnableOnCreate(val interface{}) {
	if err := j.validateSetEnableOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOnCreate",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetEvaluationExecutionRoleArn(val *string) {
	if err := j.validateSetEvaluationExecutionRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evaluationExecutionRoleArn",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetExecutionStatus(val *string) {
	if err := j.validateSetExecutionStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionStatus",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetOnlineEvaluationConfigName(val *string) {
	if err := j.validateSetOnlineEvaluationConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onlineEvaluationConfigName",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfOnlineEvaluationConfig)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a TfOnlineEvaluationConfig resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfOnlineEvaluationConfig_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfOnlineEvaluationConfig_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
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
func TfOnlineEvaluationConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOnlineEvaluationConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOnlineEvaluationConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOnlineEvaluationConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfOnlineEvaluationConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfOnlineEvaluationConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfOnlineEvaluationConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock-agentcore.TfOnlineEvaluationConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfOnlineEvaluationConfig) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) PutDataSourceConfig(value interface{}) {
	if err := t.validatePutDataSourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataSourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) PutEvaluator(value interface{}) {
	if err := t.validatePutEvaluatorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEvaluator",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) PutRule(value interface{}) {
	if err := t.validatePutRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) PutTimeouts(value *TfOnlineEvaluationConfig_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetDataSourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDataSourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetDescription() {
	_jsii_.InvokeVoid(
		t,
		"resetDescription",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetEvaluator() {
	_jsii_.InvokeVoid(
		t,
		"resetEvaluator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetExecutionStatus() {
	_jsii_.InvokeVoid(
		t,
		"resetExecutionStatus",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetRule() {
	_jsii_.InvokeVoid(
		t,
		"resetRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOnlineEvaluationConfig) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

