package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AlgorithmSpecification() AwsHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyList
	// Experimental.
	AlgorithmSpecificationInput() interface{}
	// Experimental.
	CheckpointConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigPropertyList
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
	HyperParameterRanges() AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyList
	// Experimental.
	HyperParameterRangesInput() interface{}
	// Experimental.
	HyperParameterTuningResourceConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigPropertyList
	// Experimental.
	HyperParameterTuningResourceConfigInput() interface{}
	// Experimental.
	InputDataConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigPropertyList
	// Experimental.
	InputDataConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputDataConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
	// Experimental.
	ResourceConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigPropertyList
	// Experimental.
	ResourceConfigInput() interface{}
	// Experimental.
	RetryStrategy() AwsHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyPropertyList
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
	StoppingCondition() AwsHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionPropertyList
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
	TuningObjective() AwsHyperParameterTuningJob_TrainingJobDefinitionTuningObjectivePropertyList
	// Experimental.
	TuningObjectiveInput() interface{}
	// Experimental.
	VpcConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionVpcConfigPropertyList
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

// The jsii proxy struct for AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference
type jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) AlgorithmSpecification() AwsHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyList
	_jsii_.Get(
		j,
		"algorithmSpecification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) AlgorithmSpecificationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"algorithmSpecificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) CheckpointConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigPropertyList
	_jsii_.Get(
		j,
		"checkpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) CheckpointConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) DefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) DefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableInterContainerTrafficEncryption() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableInterContainerTrafficEncryptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInterContainerTrafficEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableManagedSpotTraining() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableManagedSpotTrainingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableManagedSpotTrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableNetworkIsolation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnableNetworkIsolationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNetworkIsolationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) Environment() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) EnvironmentInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) HyperParameterRanges() AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyList
	_jsii_.Get(
		j,
		"hyperParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) HyperParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) HyperParameterTuningResourceConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigPropertyList
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) HyperParameterTuningResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hyperParameterTuningResourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) InputDataConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigPropertyList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) OutputDataConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResourceConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionResourceConfigPropertyList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) RetryStrategy() AwsHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyPropertyList
	_jsii_.Get(
		j,
		"retryStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) RetryStrategyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"retryStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) StaticHyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) StaticHyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"staticHyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) StoppingCondition() AwsHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) TuningObjective() AwsHyperParameterTuningJob_TrainingJobDefinitionTuningObjectivePropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionTuningObjectivePropertyList
	_jsii_.Get(
		j,
		"tuningObjective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) TuningObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningObjectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) VpcConfig() AwsHyperParameterTuningJob_TrainingJobDefinitionVpcConfigPropertyList {
	var returns AwsHyperParameterTuningJob_TrainingJobDefinitionVpcConfigPropertyList
	_jsii_.Get(
		j,
		"vpcConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) VpcConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcConfigInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TrainingJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference_Override(a AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.TrainingJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetDefinitionName(val *string) {
	if err := j.validateSetDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definitionName",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetEnableInterContainerTrafficEncryption(val interface{}) {
	if err := j.validateSetEnableInterContainerTrafficEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInterContainerTrafficEncryption",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetEnableManagedSpotTraining(val interface{}) {
	if err := j.validateSetEnableManagedSpotTrainingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableManagedSpotTraining",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetEnableNetworkIsolation(val interface{}) {
	if err := j.validateSetEnableNetworkIsolationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNetworkIsolation",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetEnvironment(val *map[string]*string) {
	if err := j.validateSetEnvironmentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetStaticHyperParameters(val *map[string]*string) {
	if err := j.validateSetStaticHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"staticHyperParameters",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutAlgorithmSpecification(value interface{}) {
	if err := a.validatePutAlgorithmSpecificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAlgorithmSpecification",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutCheckpointConfig(value interface{}) {
	if err := a.validatePutCheckpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCheckpointConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutHyperParameterRanges(value interface{}) {
	if err := a.validatePutHyperParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHyperParameterRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutHyperParameterTuningResourceConfig(value interface{}) {
	if err := a.validatePutHyperParameterTuningResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHyperParameterTuningResourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutInputDataConfig(value interface{}) {
	if err := a.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutOutputDataConfig(value interface{}) {
	if err := a.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutResourceConfig(value interface{}) {
	if err := a.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutRetryStrategy(value interface{}) {
	if err := a.validatePutRetryStrategyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryStrategy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := a.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutTuningObjective(value interface{}) {
	if err := a.validatePutTuningObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTuningObjective",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) PutVpcConfig(value interface{}) {
	if err := a.validatePutVpcConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetAlgorithmSpecification() {
	_jsii_.InvokeVoid(
		a,
		"resetAlgorithmSpecification",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetCheckpointConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckpointConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetDefinitionName() {
	_jsii_.InvokeVoid(
		a,
		"resetDefinitionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetEnableInterContainerTrafficEncryption() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableInterContainerTrafficEncryption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetEnableManagedSpotTraining() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableManagedSpotTraining",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetEnableNetworkIsolation() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableNetworkIsolation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetHyperParameterRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetHyperParameterRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetHyperParameterTuningResourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetHyperParameterTuningResourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetRetryStrategy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryStrategy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetStaticHyperParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetStaticHyperParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		a,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetTuningObjective() {
	_jsii_.InvokeVoid(
		a,
		"resetTuningObjective",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ResetVpcConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

