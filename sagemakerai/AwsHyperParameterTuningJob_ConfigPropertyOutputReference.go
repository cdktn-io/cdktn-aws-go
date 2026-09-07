package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsHyperParameterTuningJob_ConfigPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Objective() AwsHyperParameterTuningJob_ObjectivePropertyList
	// Experimental.
	ObjectiveInput() interface{}
	// Experimental.
	ParameterRanges() AwsHyperParameterTuningJob_ParameterRangesPropertyList
	// Experimental.
	ParameterRangesInput() interface{}
	// Experimental.
	RandomSeed() *float64
	// Experimental.
	SetRandomSeed(val *float64)
	// Experimental.
	RandomSeedInput() *float64
	// Experimental.
	ResourceLimits() AwsHyperParameterTuningJob_ResourceLimitsPropertyList
	// Experimental.
	ResourceLimitsInput() interface{}
	// Experimental.
	Strategy() *string
	// Experimental.
	SetStrategy(val *string)
	// Experimental.
	StrategyConfig() AwsHyperParameterTuningJob_StrategyConfigPropertyList
	// Experimental.
	StrategyConfigInput() interface{}
	// Experimental.
	StrategyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrainingJobEarlyStoppingType() *string
	// Experimental.
	SetTrainingJobEarlyStoppingType(val *string)
	// Experimental.
	TrainingJobEarlyStoppingTypeInput() *string
	// Experimental.
	TuningJobCompletionCriteria() AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList
	// Experimental.
	TuningJobCompletionCriteriaInput() interface{}
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
	PutObjective(value interface{})
	// Experimental.
	PutParameterRanges(value interface{})
	// Experimental.
	PutResourceLimits(value interface{})
	// Experimental.
	PutStrategyConfig(value interface{})
	// Experimental.
	PutTuningJobCompletionCriteria(value interface{})
	// Experimental.
	ResetObjective()
	// Experimental.
	ResetParameterRanges()
	// Experimental.
	ResetRandomSeed()
	// Experimental.
	ResetResourceLimits()
	// Experimental.
	ResetStrategyConfig()
	// Experimental.
	ResetTrainingJobEarlyStoppingType()
	// Experimental.
	ResetTuningJobCompletionCriteria()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsHyperParameterTuningJob_ConfigPropertyOutputReference
type jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) Objective() AwsHyperParameterTuningJob_ObjectivePropertyList {
	var returns AwsHyperParameterTuningJob_ObjectivePropertyList
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ParameterRanges() AwsHyperParameterTuningJob_ParameterRangesPropertyList {
	var returns AwsHyperParameterTuningJob_ParameterRangesPropertyList
	_jsii_.Get(
		j,
		"parameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) RandomSeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) RandomSeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResourceLimits() AwsHyperParameterTuningJob_ResourceLimitsPropertyList {
	var returns AwsHyperParameterTuningJob_ResourceLimitsPropertyList
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResourceLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyConfig() AwsHyperParameterTuningJob_StrategyConfigPropertyList {
	var returns AwsHyperParameterTuningJob_StrategyConfigPropertyList
	_jsii_.Get(
		j,
		"strategyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TrainingJobEarlyStoppingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TrainingJobEarlyStoppingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TuningJobCompletionCriteria() AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList {
	var returns AwsHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) TuningJobCompletionCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteriaInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsHyperParameterTuningJob_ConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsHyperParameterTuningJob_ConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsHyperParameterTuningJob_ConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsHyperParameterTuningJob_ConfigPropertyOutputReference_Override(a AwsHyperParameterTuningJob_ConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsHyperParameterTuningJob.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetRandomSeed(val *float64) {
	if err := j.validateSetRandomSeedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomSeed",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference)SetTrainingJobEarlyStoppingType(val *string) {
	if err := j.validateSetTrainingJobEarlyStoppingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobEarlyStoppingType",
		val,
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) PutObjective(value interface{}) {
	if err := a.validatePutObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putObjective",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) PutParameterRanges(value interface{}) {
	if err := a.validatePutParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putParameterRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) PutResourceLimits(value interface{}) {
	if err := a.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) PutStrategyConfig(value interface{}) {
	if err := a.validatePutStrategyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStrategyConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) PutTuningJobCompletionCriteria(value interface{}) {
	if err := a.validatePutTuningJobCompletionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTuningJobCompletionCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetObjective() {
	_jsii_.InvokeVoid(
		a,
		"resetObjective",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetParameterRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetParameterRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetRandomSeed() {
	_jsii_.InvokeVoid(
		a,
		"resetRandomSeed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetStrategyConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetStrategyConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetTrainingJobEarlyStoppingType() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingJobEarlyStoppingType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ResetTuningJobCompletionCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetTuningJobCompletionCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsHyperParameterTuningJob_ConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

