package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job}.
// Experimental.
type TfTrainingJob interface {
	cdktn.TerraformResource
	// Experimental.
	AlgorithmSpecification() TfTrainingJob_AlgorithmSpecificationPropertyList
	// Experimental.
	AlgorithmSpecificationInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CheckpointConfig() TfTrainingJob_CheckpointConfigPropertyList
	// Experimental.
	CheckpointConfigInput() interface{}
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
	DebugHookConfig() TfTrainingJob_DebugHookConfigPropertyList
	// Experimental.
	DebugHookConfigInput() interface{}
	// Experimental.
	DebugRuleConfigurations() TfTrainingJob_DebugRuleConfigurationsPropertyList
	// Experimental.
	DebugRuleConfigurationsInput() interface{}
	// Experimental.
	DeleteModelPackagesOnDestroy() interface{}
	// Experimental.
	SetDeleteModelPackagesOnDestroy(val interface{})
	// Experimental.
	DeleteModelPackagesOnDestroyInput() interface{}
	// Experimental.
	DeleteVpcEnisOnDestroy() interface{}
	// Experimental.
	SetDeleteVpcEnisOnDestroy(val interface{})
	// Experimental.
	DeleteVpcEnisOnDestroyInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	EnableInterContainerTrafficEncryption() interface{}
	// Experimental.
	SetEnableInterContainerTrafficEncryption(val interface{})
	// Experimental.
	EnableInterContainerTrafficEncryptionInput() interface{}
	// Experimental.
	EnableManagedSpotTraining() interface{}
	// Experimental.
	SetEnableManagedSpotTraining(val interface{})
	// Experimental.
	EnableManagedSpotTrainingInput() interface{}
	// Experimental.
	EnableNetworkIsolation() interface{}
	// Experimental.
	SetEnableNetworkIsolation(val interface{})
	// Experimental.
	EnableNetworkIsolationInput() interface{}
	// Experimental.
	Environment() *map[string]*string
	// Experimental.
	SetEnvironment(val *map[string]*string)
	// Experimental.
	EnvironmentInput() *map[string]*string
	// Experimental.
	ExperimentConfig() TfTrainingJob_ExperimentConfigPropertyList
	// Experimental.
	ExperimentConfigInput() interface{}
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	HyperParameters() *map[string]*string
	// Experimental.
	SetHyperParameters(val *map[string]*string)
	// Experimental.
	HyperParametersInput() *map[string]*string
	// Experimental.
	InfraCheckConfig() TfTrainingJob_InfraCheckConfigPropertyList
	// Experimental.
	InfraCheckConfigInput() interface{}
	// Experimental.
	InputDataConfig() TfTrainingJob_InputDataConfigPropertyList
	// Experimental.
	InputDataConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MlflowConfig() TfTrainingJob_MlflowConfigPropertyList
	// Experimental.
	MlflowConfigInput() interface{}
	// Experimental.
	ModelPackageConfig() TfTrainingJob_ModelPackageConfigPropertyList
	// Experimental.
	ModelPackageConfigInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputDataConfig() TfTrainingJob_OutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
	// Experimental.
	ProfilerConfig() TfTrainingJob_ProfilerConfigPropertyList
	// Experimental.
	ProfilerConfigInput() interface{}
	// Experimental.
	ProfilerRuleConfigurations() TfTrainingJob_ProfilerRuleConfigurationsPropertyList
	// Experimental.
	ProfilerRuleConfigurationsInput() interface{}
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
	RemoteDebugConfig() TfTrainingJob_RemoteDebugConfigPropertyList
	// Experimental.
	RemoteDebugConfigInput() interface{}
	// Experimental.
	ResourceConfig() TfTrainingJob_ResourceConfigPropertyList
	// Experimental.
	ResourceConfigInput() interface{}
	// Experimental.
	RetryStrategy() TfTrainingJob_RetryStrategyPropertyList
	// Experimental.
	RetryStrategyInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	ServerlessJobConfig() TfTrainingJob_ServerlessJobConfigPropertyList
	// Experimental.
	ServerlessJobConfigInput() interface{}
	// Experimental.
	SessionChainingConfig() TfTrainingJob_SessionChainingConfigPropertyList
	// Experimental.
	SessionChainingConfigInput() interface{}
	// Experimental.
	StoppingCondition() TfTrainingJob_StoppingConditionPropertyList
	// Experimental.
	StoppingConditionInput() interface{}
	// Experimental.
	Tags() *map[string]*string
	// Experimental.
	SetTags(val *map[string]*string)
	// Experimental.
	TagsAll() cdktn.StringMap
	// Experimental.
	TagsInput() *map[string]*string
	// Experimental.
	TensorBoardOutputConfig() TfTrainingJob_TensorBoardOutputConfigPropertyList
	// Experimental.
	TensorBoardOutputConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() TfTrainingJob_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrainingJobName() *string
	// Experimental.
	SetTrainingJobName(val *string)
	// Experimental.
	TrainingJobNameInput() *string
	// Experimental.
	VpcConfig() TfTrainingJob_VpcConfigPropertyList
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
	PutAlgorithmSpecification(value interface{})
	// Experimental.
	PutCheckpointConfig(value interface{})
	// Experimental.
	PutDebugHookConfig(value interface{})
	// Experimental.
	PutDebugRuleConfigurations(value interface{})
	// Experimental.
	PutExperimentConfig(value interface{})
	// Experimental.
	PutInfraCheckConfig(value interface{})
	// Experimental.
	PutInputDataConfig(value interface{})
	// Experimental.
	PutMlflowConfig(value interface{})
	// Experimental.
	PutModelPackageConfig(value interface{})
	// Experimental.
	PutOutputDataConfig(value interface{})
	// Experimental.
	PutProfilerConfig(value interface{})
	// Experimental.
	PutProfilerRuleConfigurations(value interface{})
	// Experimental.
	PutRemoteDebugConfig(value interface{})
	// Experimental.
	PutResourceConfig(value interface{})
	// Experimental.
	PutRetryStrategy(value interface{})
	// Experimental.
	PutServerlessJobConfig(value interface{})
	// Experimental.
	PutSessionChainingConfig(value interface{})
	// Experimental.
	PutStoppingCondition(value interface{})
	// Experimental.
	PutTensorBoardOutputConfig(value interface{})
	// Experimental.
	PutTimeouts(value *TfTrainingJob_TimeoutsProperty)
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
	ResetAlgorithmSpecification()
	// Experimental.
	ResetCheckpointConfig()
	// Experimental.
	ResetDebugHookConfig()
	// Experimental.
	ResetDebugRuleConfigurations()
	// Experimental.
	ResetDeleteModelPackagesOnDestroy()
	// Experimental.
	ResetDeleteVpcEnisOnDestroy()
	// Experimental.
	ResetEnableInterContainerTrafficEncryption()
	// Experimental.
	ResetEnableManagedSpotTraining()
	// Experimental.
	ResetEnableNetworkIsolation()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetExperimentConfig()
	// Experimental.
	ResetHyperParameters()
	// Experimental.
	ResetInfraCheckConfig()
	// Experimental.
	ResetInputDataConfig()
	// Experimental.
	ResetMlflowConfig()
	// Experimental.
	ResetModelPackageConfig()
	// Experimental.
	ResetOutputDataConfig()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	// Experimental.
	ResetProfilerConfig()
	// Experimental.
	ResetProfilerRuleConfigurations()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetRemoteDebugConfig()
	// Experimental.
	ResetResourceConfig()
	// Experimental.
	ResetRetryStrategy()
	// Experimental.
	ResetServerlessJobConfig()
	// Experimental.
	ResetSessionChainingConfig()
	// Experimental.
	ResetStoppingCondition()
	// Experimental.
	ResetTags()
	// Experimental.
	ResetTensorBoardOutputConfig()
	// Experimental.
	ResetTimeouts()
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

// The jsii proxy struct for TfTrainingJob
type jsiiProxy_TfTrainingJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_TfTrainingJob) AlgorithmSpecification() TfTrainingJob_AlgorithmSpecificationPropertyList {
	var returns TfTrainingJob_AlgorithmSpecificationPropertyList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) CheckpointConfig() TfTrainingJob_CheckpointConfigPropertyList {
	var returns TfTrainingJob_CheckpointConfigPropertyList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DebugHookConfig() TfTrainingJob_DebugHookConfigPropertyList {
	var returns TfTrainingJob_DebugHookConfigPropertyList
	_jsii_.Get(
		j,
		"debugHookConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DebugHookConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugHookConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DebugRuleConfigurations() TfTrainingJob_DebugRuleConfigurationsPropertyList {
	var returns TfTrainingJob_DebugRuleConfigurationsPropertyList
	_jsii_.Get(
		j,
		"debugRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DebugRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DeleteModelPackagesOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DeleteModelPackagesOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DeleteVpcEnisOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DeleteVpcEnisOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ExperimentConfig() TfTrainingJob_ExperimentConfigPropertyList {
	var returns TfTrainingJob_ExperimentConfigPropertyList
	_jsii_.Get(
		j,
		"experimentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ExperimentConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) HyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) HyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) InfraCheckConfig() TfTrainingJob_InfraCheckConfigPropertyList {
	var returns TfTrainingJob_InfraCheckConfigPropertyList
	_jsii_.Get(
		j,
		"infraCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) InfraCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infraCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) InputDataConfig() TfTrainingJob_InputDataConfigPropertyList {
	var returns TfTrainingJob_InputDataConfigPropertyList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) MlflowConfig() TfTrainingJob_MlflowConfigPropertyList {
	var returns TfTrainingJob_MlflowConfigPropertyList
	_jsii_.Get(
		j,
		"mlflowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) MlflowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlflowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ModelPackageConfig() TfTrainingJob_ModelPackageConfigPropertyList {
	var returns TfTrainingJob_ModelPackageConfigPropertyList
	_jsii_.Get(
		j,
		"modelPackageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ModelPackageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) OutputDataConfig() TfTrainingJob_OutputDataConfigPropertyList {
	var returns TfTrainingJob_OutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ProfilerConfig() TfTrainingJob_ProfilerConfigPropertyList {
	var returns TfTrainingJob_ProfilerConfigPropertyList
	_jsii_.Get(
		j,
		"profilerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ProfilerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ProfilerRuleConfigurations() TfTrainingJob_ProfilerRuleConfigurationsPropertyList {
	var returns TfTrainingJob_ProfilerRuleConfigurationsPropertyList
	_jsii_.Get(
		j,
		"profilerRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ProfilerRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RemoteDebugConfig() TfTrainingJob_RemoteDebugConfigPropertyList {
	var returns TfTrainingJob_RemoteDebugConfigPropertyList
	_jsii_.Get(
		j,
		"remoteDebugConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RemoteDebugConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebugConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ResourceConfig() TfTrainingJob_ResourceConfigPropertyList {
	var returns TfTrainingJob_ResourceConfigPropertyList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RetryStrategy() TfTrainingJob_RetryStrategyPropertyList {
	var returns TfTrainingJob_RetryStrategyPropertyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ServerlessJobConfig() TfTrainingJob_ServerlessJobConfigPropertyList {
	var returns TfTrainingJob_ServerlessJobConfigPropertyList
	_jsii_.Get(
		j,
		"serverlessJobConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) ServerlessJobConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessJobConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) SessionChainingConfig() TfTrainingJob_SessionChainingConfigPropertyList {
	var returns TfTrainingJob_SessionChainingConfigPropertyList
	_jsii_.Get(
		j,
		"sessionChainingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) SessionChainingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionChainingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) StoppingCondition() TfTrainingJob_StoppingConditionPropertyList {
	var returns TfTrainingJob_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TensorBoardOutputConfig() TfTrainingJob_TensorBoardOutputConfigPropertyList {
	var returns TfTrainingJob_TensorBoardOutputConfigPropertyList
	_jsii_.Get(
		j,
		"tensorBoardOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TensorBoardOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tensorBoardOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) Timeouts() TfTrainingJob_TimeoutsPropertyOutputReference {
	var returns TfTrainingJob_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TrainingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) TrainingJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) VpcConfig() TfTrainingJob_VpcConfigPropertyList {
	var returns TfTrainingJob_VpcConfigPropertyList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTrainingJob) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job} Resource.
// Experimental.
func NewTfTrainingJob(scope constructs.Construct, id *string, config *TfTrainingJobConfig) TfTrainingJob {
	_init_.Initialize()

	if err := validateNewTfTrainingJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTrainingJob{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job} Resource.
// Experimental.
func NewTfTrainingJob_Override(t TfTrainingJob, scope constructs.Construct, id *string, config *TfTrainingJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetDeleteModelPackagesOnDestroy(val interface{}) {
	if err := j.validateSetDeleteModelPackagesOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteModelPackagesOnDestroy",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetDeleteVpcEnisOnDestroy(val interface{}) {
	if err := j.validateSetDeleteVpcEnisOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVpcEnisOnDestroy",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetHyperParameters(val *map[string]*string) {
	if err := j.validateSetHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperParameters",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_TfTrainingJob)SetTrainingJobName(val *string) {
	if err := j.validateSetTrainingJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobName",
		val,
	)
}

// Generates CDKTN code for importing a TfTrainingJob resource upon running "cdktn plan <stack-name>".
// Experimental.
func TfTrainingJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateTfTrainingJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
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
func TfTrainingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTrainingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTrainingJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTrainingJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TfTrainingJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTfTrainingJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TfTrainingJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sagemaker-ai.TfTrainingJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TfTrainingJob) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TfTrainingJob) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TfTrainingJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTrainingJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrainingJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTrainingJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTrainingJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTrainingJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTrainingJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTrainingJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTrainingJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTrainingJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TfTrainingJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTrainingJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (t *jsiiProxy_TfTrainingJob) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTrainingJob) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TfTrainingJob) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TfTrainingJob) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutAlgorithmSpecification(value interface{}) {
	if err := t.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutCheckpointConfig(value interface{}) {
	if err := t.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutDebugHookConfig(value interface{}) {
	if err := t.validatePutDebugHookConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDebugHookConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutDebugRuleConfigurations(value interface{}) {
	if err := t.validatePutDebugRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDebugRuleConfigurations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutExperimentConfig(value interface{}) {
	if err := t.validatePutExperimentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExperimentConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutInfraCheckConfig(value interface{}) {
	if err := t.validatePutInfraCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInfraCheckConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutInputDataConfig(value interface{}) {
	if err := t.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutMlflowConfig(value interface{}) {
	if err := t.validatePutMlflowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMlflowConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutModelPackageConfig(value interface{}) {
	if err := t.validatePutModelPackageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putModelPackageConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutOutputDataConfig(value interface{}) {
	if err := t.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutProfilerConfig(value interface{}) {
	if err := t.validatePutProfilerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProfilerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutProfilerRuleConfigurations(value interface{}) {
	if err := t.validatePutProfilerRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProfilerRuleConfigurations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutRemoteDebugConfig(value interface{}) {
	if err := t.validatePutRemoteDebugConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRemoteDebugConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutResourceConfig(value interface{}) {
	if err := t.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutRetryStrategy(value interface{}) {
	if err := t.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutServerlessJobConfig(value interface{}) {
	if err := t.validatePutServerlessJobConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerlessJobConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutSessionChainingConfig(value interface{}) {
	if err := t.validatePutSessionChainingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSessionChainingConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutStoppingCondition(value interface{}) {
	if err := t.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutTensorBoardOutputConfig(value interface{}) {
	if err := t.validatePutTensorBoardOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTensorBoardOutputConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutTimeouts(value *TfTrainingJob_TimeoutsProperty) {
	if err := t.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) PutVpcConfig(value interface{}) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTrainingJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := t.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetDebugHookConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDebugHookConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetDebugRuleConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetDebugRuleConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetDeleteModelPackagesOnDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteModelPackagesOnDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetDeleteVpcEnisOnDestroy() {
	_jsii_.InvokeVoid(
		t,
		"resetDeleteVpcEnisOnDestroy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetExperimentConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetExperimentConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetHyperParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetHyperParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetInfraCheckConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInfraCheckConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetMlflowConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetMlflowConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetModelPackageConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetModelPackageConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetProfilerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetProfilerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetProfilerRuleConfigurations() {
	_jsii_.InvokeVoid(
		t,
		"resetProfilerRuleConfigurations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetRemoteDebugConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetRemoteDebugConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetServerlessJobConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetServerlessJobConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetSessionChainingConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetSessionChainingConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		t,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetTensorBoardOutputConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetTensorBoardOutputConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTrainingJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTrainingJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

