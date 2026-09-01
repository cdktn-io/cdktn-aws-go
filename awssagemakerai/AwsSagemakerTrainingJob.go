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
type AwsSagemakerTrainingJob interface {
	cdktn.TerraformResource
	// Experimental.
	AlgorithmSpecification() AwsSagemakerTrainingJob_AlgorithmSpecificationPropertyList
	// Experimental.
	AlgorithmSpecificationInput() interface{}
	// Experimental.
	Arn() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	CheckpointConfig() AwsSagemakerTrainingJob_CheckpointConfigPropertyList
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
	DebugHookConfig() AwsSagemakerTrainingJob_DebugHookConfigPropertyList
	// Experimental.
	DebugHookConfigInput() interface{}
	// Experimental.
	DebugRuleConfigurations() AwsSagemakerTrainingJob_DebugRuleConfigurationsPropertyList
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
	ExperimentConfig() AwsSagemakerTrainingJob_ExperimentConfigPropertyList
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
	InfraCheckConfig() AwsSagemakerTrainingJob_InfraCheckConfigPropertyList
	// Experimental.
	InfraCheckConfigInput() interface{}
	// Experimental.
	InputDataConfig() AwsSagemakerTrainingJob_InputDataConfigPropertyList
	// Experimental.
	InputDataConfigInput() interface{}
	// Experimental.
	Lifecycle() *cdktn.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformResourceLifecycle)
	// Experimental.
	MlflowConfig() AwsSagemakerTrainingJob_MlflowConfigPropertyList
	// Experimental.
	MlflowConfigInput() interface{}
	// Experimental.
	ModelPackageConfig() AwsSagemakerTrainingJob_ModelPackageConfigPropertyList
	// Experimental.
	ModelPackageConfigInput() interface{}
	// The tree node.
	// Experimental.
	Node() constructs.Node
	// Experimental.
	OutputDataConfig() AwsSagemakerTrainingJob_OutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
	// Experimental.
	ProfilerConfig() AwsSagemakerTrainingJob_ProfilerConfigPropertyList
	// Experimental.
	ProfilerConfigInput() interface{}
	// Experimental.
	ProfilerRuleConfigurations() AwsSagemakerTrainingJob_ProfilerRuleConfigurationsPropertyList
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
	RemoteDebugConfig() AwsSagemakerTrainingJob_RemoteDebugConfigPropertyList
	// Experimental.
	RemoteDebugConfigInput() interface{}
	// Experimental.
	ResourceConfig() AwsSagemakerTrainingJob_ResourceConfigPropertyList
	// Experimental.
	ResourceConfigInput() interface{}
	// Experimental.
	RetryStrategy() AwsSagemakerTrainingJob_RetryStrategyPropertyList
	// Experimental.
	RetryStrategyInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	ServerlessJobConfig() AwsSagemakerTrainingJob_ServerlessJobConfigPropertyList
	// Experimental.
	ServerlessJobConfigInput() interface{}
	// Experimental.
	SessionChainingConfig() AwsSagemakerTrainingJob_SessionChainingConfigPropertyList
	// Experimental.
	SessionChainingConfigInput() interface{}
	// Experimental.
	StoppingCondition() AwsSagemakerTrainingJob_StoppingConditionPropertyList
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
	TensorBoardOutputConfig() AwsSagemakerTrainingJob_TensorBoardOutputConfigPropertyList
	// Experimental.
	TensorBoardOutputConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	Timeouts() AwsSagemakerTrainingJob_TimeoutsPropertyOutputReference
	// Experimental.
	TimeoutsInput() interface{}
	// Experimental.
	TrainingJobName() *string
	// Experimental.
	SetTrainingJobName(val *string)
	// Experimental.
	TrainingJobNameInput() *string
	// Experimental.
	VpcConfig() AwsSagemakerTrainingJob_VpcConfigPropertyList
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
	PutTimeouts(value *AwsSagemakerTrainingJob_TimeoutsProperty)
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

// The jsii proxy struct for AwsSagemakerTrainingJob
type jsiiProxy_AwsSagemakerTrainingJob struct {
	internal.Type__cdktnTerraformResource
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) AlgorithmSpecification() AwsSagemakerTrainingJob_AlgorithmSpecificationPropertyList {
	var returns AwsSagemakerTrainingJob_AlgorithmSpecificationPropertyList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) CheckpointConfig() AwsSagemakerTrainingJob_CheckpointConfigPropertyList {
	var returns AwsSagemakerTrainingJob_CheckpointConfigPropertyList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DebugHookConfig() AwsSagemakerTrainingJob_DebugHookConfigPropertyList {
	var returns AwsSagemakerTrainingJob_DebugHookConfigPropertyList
	_jsii_.Get(
		j,
		"debugHookConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DebugHookConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugHookConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DebugRuleConfigurations() AwsSagemakerTrainingJob_DebugRuleConfigurationsPropertyList {
	var returns AwsSagemakerTrainingJob_DebugRuleConfigurationsPropertyList
	_jsii_.Get(
		j,
		"debugRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DebugRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"debugRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DeleteModelPackagesOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DeleteModelPackagesOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteModelPackagesOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DeleteVpcEnisOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DeleteVpcEnisOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteVpcEnisOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ExperimentConfig() AwsSagemakerTrainingJob_ExperimentConfigPropertyList {
	var returns AwsSagemakerTrainingJob_ExperimentConfigPropertyList
	_jsii_.Get(
		j,
		"experimentConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ExperimentConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"experimentConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) HyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) HyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) InfraCheckConfig() AwsSagemakerTrainingJob_InfraCheckConfigPropertyList {
	var returns AwsSagemakerTrainingJob_InfraCheckConfigPropertyList
	_jsii_.Get(
		j,
		"infraCheckConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) InfraCheckConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"infraCheckConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) InputDataConfig() AwsSagemakerTrainingJob_InputDataConfigPropertyList {
	var returns AwsSagemakerTrainingJob_InputDataConfigPropertyList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Lifecycle() *cdktn.TerraformResourceLifecycle {
	var returns *cdktn.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) MlflowConfig() AwsSagemakerTrainingJob_MlflowConfigPropertyList {
	var returns AwsSagemakerTrainingJob_MlflowConfigPropertyList
	_jsii_.Get(
		j,
		"mlflowConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) MlflowConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mlflowConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ModelPackageConfig() AwsSagemakerTrainingJob_ModelPackageConfigPropertyList {
	var returns AwsSagemakerTrainingJob_ModelPackageConfigPropertyList
	_jsii_.Get(
		j,
		"modelPackageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ModelPackageConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modelPackageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) OutputDataConfig() AwsSagemakerTrainingJob_OutputDataConfigPropertyList {
	var returns AwsSagemakerTrainingJob_OutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ProfilerConfig() AwsSagemakerTrainingJob_ProfilerConfigPropertyList {
	var returns AwsSagemakerTrainingJob_ProfilerConfigPropertyList
	_jsii_.Get(
		j,
		"profilerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ProfilerConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ProfilerRuleConfigurations() AwsSagemakerTrainingJob_ProfilerRuleConfigurationsPropertyList {
	var returns AwsSagemakerTrainingJob_ProfilerRuleConfigurationsPropertyList
	_jsii_.Get(
		j,
		"profilerRuleConfigurations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ProfilerRuleConfigurationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"profilerRuleConfigurationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RemoteDebugConfig() AwsSagemakerTrainingJob_RemoteDebugConfigPropertyList {
	var returns AwsSagemakerTrainingJob_RemoteDebugConfigPropertyList
	_jsii_.Get(
		j,
		"remoteDebugConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RemoteDebugConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteDebugConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ResourceConfig() AwsSagemakerTrainingJob_ResourceConfigPropertyList {
	var returns AwsSagemakerTrainingJob_ResourceConfigPropertyList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RetryStrategy() AwsSagemakerTrainingJob_RetryStrategyPropertyList {
	var returns AwsSagemakerTrainingJob_RetryStrategyPropertyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ServerlessJobConfig() AwsSagemakerTrainingJob_ServerlessJobConfigPropertyList {
	var returns AwsSagemakerTrainingJob_ServerlessJobConfigPropertyList
	_jsii_.Get(
		j,
		"serverlessJobConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) ServerlessJobConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverlessJobConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) SessionChainingConfig() AwsSagemakerTrainingJob_SessionChainingConfigPropertyList {
	var returns AwsSagemakerTrainingJob_SessionChainingConfigPropertyList
	_jsii_.Get(
		j,
		"sessionChainingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) SessionChainingConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sessionChainingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) StoppingCondition() AwsSagemakerTrainingJob_StoppingConditionPropertyList {
	var returns AwsSagemakerTrainingJob_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TagsAll() cdktn.StringMap {
	var returns cdktn.StringMap
	_jsii_.Get(
		j,
		"tagsAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TensorBoardOutputConfig() AwsSagemakerTrainingJob_TensorBoardOutputConfigPropertyList {
	var returns AwsSagemakerTrainingJob_TensorBoardOutputConfigPropertyList
	_jsii_.Get(
		j,
		"tensorBoardOutputConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TensorBoardOutputConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tensorBoardOutputConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) Timeouts() AwsSagemakerTrainingJob_TimeoutsPropertyOutputReference {
	var returns AwsSagemakerTrainingJob_TimeoutsPropertyOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TrainingJobName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) TrainingJobNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) VpcConfig() AwsSagemakerTrainingJob_VpcConfigPropertyList {
	var returns AwsSagemakerTrainingJob_VpcConfigPropertyList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerTrainingJob) VpcConfigInput() interface{} {
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
func NewAwsSagemakerTrainingJob(scope constructs.Construct, id *string, config *AwsSagemakerTrainingJobConfig) AwsSagemakerTrainingJob {
	_init_.Initialize()

	if err := validateNewAwsSagemakerTrainingJobParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerTrainingJob{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_training_job aws_sagemaker_training_job} Resource.
// Experimental.
func NewAwsSagemakerTrainingJob_Override(a AwsSagemakerTrainingJob, scope constructs.Construct, id *string, config *AwsSagemakerTrainingJobConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetDeleteModelPackagesOnDestroy(val interface{}) {
	if err := j.validateSetDeleteModelPackagesOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteModelPackagesOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetDeleteVpcEnisOnDestroy(val interface{}) {
	if err := j.validateSetDeleteVpcEnisOnDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deleteVpcEnisOnDestroy",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetHyperParameters(val *map[string]*string) {
	if err := j.validateSetHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperParameters",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetLifecycle(val *cdktn.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerTrainingJob)SetTrainingJobName(val *string) {
	if err := j.validateSetTrainingJobNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobName",
		val,
	)
}

// Generates CDKTN code for importing a AwsSagemakerTrainingJob resource upon running "cdktn plan <stack-name>".
// Experimental.
func AwsSagemakerTrainingJob_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktn.TerraformProvider) cdktn.ImportableResource {
	_init_.Initialize()

	if err := validateAwsSagemakerTrainingJob_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktn.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
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
func AwsSagemakerTrainingJob_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSagemakerTrainingJob_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSagemakerTrainingJob_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSagemakerTrainingJob_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSagemakerTrainingJob_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSagemakerTrainingJob_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSagemakerTrainingJob_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerTrainingJob",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) AddMoveTarget(moveTarget *string) {
	if err := a.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ImportFrom(id *string, provider cdktn.TerraformProvider) {
	if err := a.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) MarkWriteOnlyAttribute(value interface{}) interface{} {
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

func (a *jsiiProxy_AwsSagemakerTrainingJob) MoveFromId(id *string) {
	if err := a.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveFromId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) MoveTo(moveTarget *string, index interface{}) {
	if err := a.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) MoveToId(id *string) {
	if err := a.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"moveToId",
		[]interface{}{id},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutAlgorithmSpecification(value interface{}) {
	if err := a.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutCheckpointConfig(value interface{}) {
	if err := a.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutDebugHookConfig(value interface{}) {
	if err := a.validatePutDebugHookConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDebugHookConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutDebugRuleConfigurations(value interface{}) {
	if err := a.validatePutDebugRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDebugRuleConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutExperimentConfig(value interface{}) {
	if err := a.validatePutExperimentConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExperimentConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutInfraCheckConfig(value interface{}) {
	if err := a.validatePutInfraCheckConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInfraCheckConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutInputDataConfig(value interface{}) {
	if err := a.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutMlflowConfig(value interface{}) {
	if err := a.validatePutMlflowConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMlflowConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutModelPackageConfig(value interface{}) {
	if err := a.validatePutModelPackageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putModelPackageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutOutputDataConfig(value interface{}) {
	if err := a.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutProfilerConfig(value interface{}) {
	if err := a.validatePutProfilerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProfilerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutProfilerRuleConfigurations(value interface{}) {
	if err := a.validatePutProfilerRuleConfigurationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProfilerRuleConfigurations",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutRemoteDebugConfig(value interface{}) {
	if err := a.validatePutRemoteDebugConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRemoteDebugConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutResourceConfig(value interface{}) {
	if err := a.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutRetryStrategy(value interface{}) {
	if err := a.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutServerlessJobConfig(value interface{}) {
	if err := a.validatePutServerlessJobConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerlessJobConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutSessionChainingConfig(value interface{}) {
	if err := a.validatePutSessionChainingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSessionChainingConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutStoppingCondition(value interface{}) {
	if err := a.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutTensorBoardOutputConfig(value interface{}) {
	if err := a.validatePutTensorBoardOutputConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTensorBoardOutputConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutTimeouts(value *AwsSagemakerTrainingJob_TimeoutsProperty) {
	if err := a.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) PutVpcConfig(value interface{}) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := a.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetDebugHookConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDebugHookConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetDebugRuleConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetDebugRuleConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetDeleteModelPackagesOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteModelPackagesOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetDeleteVpcEnisOnDestroy() {
	_jsii_.InvokeVoid(
		a,
		"resetDeleteVpcEnisOnDestroy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetExperimentConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetExperimentConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetHyperParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetHyperParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetInfraCheckConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetInfraCheckConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetMlflowConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetMlflowConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetModelPackageConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetModelPackageConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetProfilerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetProfilerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetProfilerRuleConfigurations() {
	_jsii_.InvokeVoid(
		a,
		"resetProfilerRuleConfigurations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetRemoteDebugConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetRemoteDebugConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetServerlessJobConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetServerlessJobConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetSessionChainingConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionChainingConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetTensorBoardOutputConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetTensorBoardOutputConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetTimeouts() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerTrainingJob) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

