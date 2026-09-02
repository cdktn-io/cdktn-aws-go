package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AlgorithmSpecification() TfHyperParameterTuningJob_TrainingJobDefinitionsAlgorithmSpecificationPropertyList
	// Experimental.
	AlgorithmSpecificationInput() interface{}
	// Experimental.
	CheckpointConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsCheckpointConfigPropertyList
	// Experimental.
	CheckpointConfigInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DefinitionName() *string
	// Experimental.
	SetDefinitionName(val *string)
	// Experimental.
	DefinitionNameInput() *string
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
	Fqn() *string
	// Experimental.
	HyperParameterRanges() TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyList
	// Experimental.
	HyperParameterRangesInput() interface{}
	// Experimental.
	HyperParameterTuningResourceConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigPropertyList
	// Experimental.
	HyperParameterTuningResourceConfigInput() interface{}
	// Experimental.
	InputDataConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyList
	// Experimental.
	InputDataConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputDataConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsOutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
	// Experimental.
	ResourceConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyList
	// Experimental.
	ResourceConfigInput() interface{}
	// Experimental.
	RetryStrategy() TfHyperParameterTuningJob_TrainingJobDefinitionsRetryStrategyPropertyList
	// Experimental.
	RetryStrategyInput() interface{}
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	StaticHyperParameters() *map[string]*string
	// Experimental.
	SetStaticHyperParameters(val *map[string]*string)
	// Experimental.
	StaticHyperParametersInput() *map[string]*string
	// Experimental.
	StoppingCondition() TfHyperParameterTuningJob_TrainingJobDefinitionsStoppingConditionPropertyList
	// Experimental.
	StoppingConditionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TuningObjective() TfHyperParameterTuningJob_TrainingJobDefinitionsTuningObjectivePropertyList
	// Experimental.
	TuningObjectiveInput() interface{}
	// Experimental.
	VpcConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsVpcConfigPropertyList
	// Experimental.
	VpcConfigInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	PutAlgorithmSpecification(value interface{})
	// Experimental.
	PutCheckpointConfig(value interface{})
	// Experimental.
	PutHyperParameterRanges(value interface{})
	// Experimental.
	PutHyperParameterTuningResourceConfig(value interface{})
	// Experimental.
	PutInputDataConfig(value interface{})
	// Experimental.
	PutOutputDataConfig(value interface{})
	// Experimental.
	PutResourceConfig(value interface{})
	// Experimental.
	PutRetryStrategy(value interface{})
	// Experimental.
	PutStoppingCondition(value interface{})
	// Experimental.
	PutTuningObjective(value interface{})
	// Experimental.
	PutVpcConfig(value interface{})
	// Experimental.
	ResetAlgorithmSpecification()
	// Experimental.
	ResetCheckpointConfig()
	// Experimental.
	ResetDefinitionName()
	// Experimental.
	ResetEnableInterContainerTrafficEncryption()
	// Experimental.
	ResetEnableManagedSpotTraining()
	// Experimental.
	ResetEnableNetworkIsolation()
	// Experimental.
	ResetEnvironment()
	// Experimental.
	ResetHyperParameterRanges()
	// Experimental.
	ResetHyperParameterTuningResourceConfig()
	// Experimental.
	ResetInputDataConfig()
	// Experimental.
	ResetOutputDataConfig()
	// Experimental.
	ResetResourceConfig()
	// Experimental.
	ResetRetryStrategy()
	// Experimental.
	ResetStaticHyperParameters()
	// Experimental.
	ResetStoppingCondition()
	// Experimental.
	ResetTuningObjective()
	// Experimental.
	ResetVpcConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) AlgorithmSpecification() TfHyperParameterTuningJob_TrainingJobDefinitionsAlgorithmSpecificationPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsAlgorithmSpecificationPropertyList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) CheckpointConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsCheckpointConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsCheckpointConfigPropertyList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) DefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) DefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) HyperParameterRanges() TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyList
	_jsii_.Get(
		j,
		"hyperParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) HyperParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) HyperParameterTuningResourceConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigPropertyList
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) HyperParameterTuningResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) InputDataConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsInputDataConfigPropertyList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) OutputDataConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsOutputDataConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsOutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResourceConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsResourceConfigPropertyList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) RetryStrategy() TfHyperParameterTuningJob_TrainingJobDefinitionsRetryStrategyPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsRetryStrategyPropertyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) StaticHyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) StaticHyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) StoppingCondition() TfHyperParameterTuningJob_TrainingJobDefinitionsStoppingConditionPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsStoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) TuningObjective() TfHyperParameterTuningJob_TrainingJobDefinitionsTuningObjectivePropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsTuningObjectivePropertyList
	_jsii_.Get(
		j,
		"tuningObjective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) TuningObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningObjectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) VpcConfig() TfHyperParameterTuningJob_TrainingJobDefinitionsVpcConfigPropertyList {
	var returns TfHyperParameterTuningJob_TrainingJobDefinitionsVpcConfigPropertyList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference_Override(t TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.TrainingJobDefinitionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetDefinitionName(val *string) {
	if err := j.validateSetDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionName",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetStaticHyperParameters(val *map[string]*string) {
	if err := j.validateSetStaticHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticHyperParameters",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutAlgorithmSpecification(value interface{}) {
	if err := t.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutCheckpointConfig(value interface{}) {
	if err := t.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutHyperParameterRanges(value interface{}) {
	if err := t.validatePutHyperParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHyperParameterRanges",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutHyperParameterTuningResourceConfig(value interface{}) {
	if err := t.validatePutHyperParameterTuningResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHyperParameterTuningResourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutInputDataConfig(value interface{}) {
	if err := t.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutOutputDataConfig(value interface{}) {
	if err := t.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutResourceConfig(value interface{}) {
	if err := t.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutRetryStrategy(value interface{}) {
	if err := t.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := t.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutTuningObjective(value interface{}) {
	if err := t.validatePutTuningObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTuningObjective",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) PutVpcConfig(value interface{}) {
	if err := t.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		t,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetDefinitionName() {
	_jsii_.InvokeVoid(
		t,
		"resetDefinitionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetHyperParameterRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetHyperParameterRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetHyperParameterTuningResourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetHyperParameterTuningResourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetStaticHyperParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetStaticHyperParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		t,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetTuningObjective() {
	_jsii_.InvokeVoid(
		t,
		"resetTuningObjective",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

