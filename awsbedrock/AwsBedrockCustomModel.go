package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrock/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awsbedrock/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model}.
// Experimental.
type AwsBedrockCustomModel interface {
	cdktn.TerraformResource
	// Experimental.
	BaseModelIdentifier() *string
	// Experimental.
	SetBaseModelIdentifier(val *string)
	// Experimental.
	BaseModelIdentifierInput() *string
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
	CustomizationType() *string
	// Experimental.
	SetCustomizationType(val *string)
	// Experimental.
	CustomizationTypeInput() *string
	// Experimental.
	CustomModelArn() *string
	// Experimental.
	CustomModelKmsKeyId() *string
	// Experimental.
	SetCustomModelKmsKeyId(val *string)
	// Experimental.
	CustomModelKmsKeyIdInput() *string
	// Experimental.
	CustomModelName() *string
	// Experimental.
	SetCustomModelName(val *string)
	// Experimental.
	CustomModelNameInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	Hyperparameters() *map[string]*string
	// Experimental.
	SetHyperparameters(val *map[string]*string)
	// Experimental.
	HyperparametersInput() *map[string]*string
	// Experimental.
	Id() *string
	// Experimental.
	JobArn() *string
	// Experimental.
	JobName() *string
	// Experimental.
	SetJobName(val *string)
	// Experimental.
	JobNameInput() *string
	// Experimental.
	JobStatus() *string
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputDataConfig() AwsBedrockCustomModel_OutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
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
	Timeouts() AwsBedrockCustomModel_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrainingDataConfig() AwsBedrockCustomModel_TrainingDataConfigPropertyList
	// Experimental.
	TrainingDataConfigInput() interface{}
	// Experimental.
	TrainingMetrics() AwsBedrockCustomModel_TrainingMetricsPropertyList
	// Experimental.
	ValidationDataConfig() AwsBedrockCustomModel_ValidationDataConfigPropertyList
	// Experimental.
	ValidationDataConfigInput() interface{}
	// Experimental.
	ValidationMetrics() AwsBedrockCustomModel_ValidationMetricsPropertyList
	// Experimental.
	VpcConfig() AwsBedrockCustomModel_VpcConfigPropertyList
	// Experimental.
	VpcConfigInput() interface{}
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
	PutOutputDataConfig(value interface{})
	// Experimental.
	PutTimeouts(value *AwsBedrockCustomModel_TimeoutsProperty)
	// Experimental.
	PutTrainingDataConfig(value interface{})
	// Experimental.
	PutValidationDataConfig(value interface{})
	// Experimental.
	PutVpcConfig(value interface{})
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
	ResetCustomizationType()
	// Experimental.
	ResetCustomModelKmsKeyId()
	// Experimental.
	ResetOutputDataConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTimeouts()
	// Experimental.
	ResetTrainingDataConfig()
	// Experimental.
	ResetValidationDataConfig()
	// Experimental.
	ResetVpcConfig()
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

// The jsii proxy struct for AwsBedrockCustomModel
type jsiiProxy_AwsBedrockCustomModel struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsBedrockCustomModel) BaseModelIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) BaseModelIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseModelIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomizationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomizationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customizationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomModelKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomModelKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomModelName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) CustomModelNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customModelNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Hyperparameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperparameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) HyperparametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperparametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) JobArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) JobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) JobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) JobStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) OutputDataConfig() AwsBedrockCustomModel_OutputDataConfigPropertyList {
	var returns AwsBedrockCustomModel_OutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) Timeouts() AwsBedrockCustomModel_TimeoutsPropertyOutputReference {
	var returns AwsBedrockCustomModel_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TrainingDataConfig() AwsBedrockCustomModel_TrainingDataConfigPropertyList {
	var returns AwsBedrockCustomModel_TrainingDataConfigPropertyList
	_jsii_.Get(
		j,
		"trainingDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TrainingDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trainingDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) TrainingMetrics() AwsBedrockCustomModel_TrainingMetricsPropertyList {
	var returns AwsBedrockCustomModel_TrainingMetricsPropertyList
	_jsii_.Get(
		j,
		"trainingMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) ValidationDataConfig() AwsBedrockCustomModel_ValidationDataConfigPropertyList {
	var returns AwsBedrockCustomModel_ValidationDataConfigPropertyList
	_jsii_.Get(
		j,
		"validationDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) ValidationDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validationDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) ValidationMetrics() AwsBedrockCustomModel_ValidationMetricsPropertyList {
	var returns AwsBedrockCustomModel_ValidationMetricsPropertyList
	_jsii_.Get(
		j,
		"validationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) VpcConfig() AwsBedrockCustomModel_VpcConfigPropertyList {
	var returns AwsBedrockCustomModel_VpcConfigPropertyList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBedrockCustomModel) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model} Resource.
// Experimental.
func NewAwsBedrockCustomModel(scope constructs.Construct, id *string, config *AwsBedrockCustomModelConfig) AwsBedrockCustomModel {
	_init_.Initialize()

	if err := validateNewAwsBedrockCustomModelParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBedrockCustomModel{}

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_custom_model aws_bedrock_custom_model} Resource.
// Experimental.
func NewAwsBedrockCustomModel_Override(a AwsBedrockCustomModel, scope constructs.Construct, id *string, config *AwsBedrockCustomModelConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetBaseModelIdentifier(val *string) {
	if err := j.validateSetBaseModelIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseModelIdentifier",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetCustomizationType(val *string) {
	if err := j.validateSetCustomizationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customizationType",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetCustomModelKmsKeyId(val *string) {
	if err := j.validateSetCustomModelKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customModelKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetCustomModelName(val *string) {
	if err := j.validateSetCustomModelNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customModelName",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetHyperparameters(val *map[string]*string) {
	if err := j.validateSetHyperparametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperparameters",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetJobName(val *string) {
	if err := j.validateSetJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobName",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsBedrockCustomModel)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTN code for importing a AwsBedrockCustomModel resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsBedrockCustomModel_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsBedrockCustomModel_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
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
func AwsBedrockCustomModel_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockCustomModel_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockCustomModel_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockCustomModel_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsBedrockCustomModel_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsBedrockCustomModel_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsBedrockCustomModel_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-bedrock.AwsBedrockCustomModel",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBedrockCustomModel) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBedrockCustomModel) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBedrockCustomModel) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsBedrockCustomModel) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) PutOutputDataConfig(value interface{}) {
	if err := a.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) PutTimeouts(value *AwsBedrockCustomModel_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) PutTrainingDataConfig(value interface{}) {
	if err := a.validatePutTrainingDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrainingDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) PutValidationDataConfig(value interface{}) {
	if err := a.validatePutValidationDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putValidationDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) PutVpcConfig(value interface{}) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetCustomizationType() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomizationType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetCustomModelKmsKeyId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomModelKmsKeyId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetTrainingDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetValidationDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetValidationDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBedrockCustomModel) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBedrockCustomModel) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

