package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_ConfigPropertyOutputReference interface {
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
	Objective() TfHyperParameterTuningJob_ObjectivePropertyList
	// Experimental.
	ObjectiveInput() interface{}
	// Experimental.
	ParameterRanges() TfHyperParameterTuningJob_ParameterRangesPropertyList
	// Experimental.
	ParameterRangesInput() interface{}
	// Experimental.
	RandomSeed() *float64
	// Experimental.
	SetRandomSeed(val *float64)
	// Experimental.
	RandomSeedInput() *float64
	// Experimental.
	ResourceLimits() TfHyperParameterTuningJob_ResourceLimitsPropertyList
	// Experimental.
	ResourceLimitsInput() interface{}
	// Experimental.
	Strategy() *string
	// Experimental.
	SetStrategy(val *string)
	// Experimental.
	StrategyConfig() TfHyperParameterTuningJob_StrategyConfigPropertyList
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
	TuningJobCompletionCriteria() TfHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList
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

// The jsii proxy struct for TfHyperParameterTuningJob_ConfigPropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) Objective() TfHyperParameterTuningJob_ObjectivePropertyList {
	var returns TfHyperParameterTuningJob_ObjectivePropertyList
	_jsii_.Get(
		j,
		"objective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ObjectiveInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"objectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ParameterRanges() TfHyperParameterTuningJob_ParameterRangesPropertyList {
	var returns TfHyperParameterTuningJob_ParameterRangesPropertyList
	_jsii_.Get(
		j,
		"parameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) RandomSeed() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) RandomSeedInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"randomSeedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResourceLimits() TfHyperParameterTuningJob_ResourceLimitsPropertyList {
	var returns TfHyperParameterTuningJob_ResourceLimitsPropertyList
	_jsii_.Get(
		j,
		"resourceLimits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResourceLimitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceLimitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) Strategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyConfig() TfHyperParameterTuningJob_StrategyConfigPropertyList {
	var returns TfHyperParameterTuningJob_StrategyConfigPropertyList
	_jsii_.Get(
		j,
		"strategyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strategyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) StrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"strategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TrainingJobEarlyStoppingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TrainingJobEarlyStoppingTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingJobEarlyStoppingTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TuningJobCompletionCriteria() TfHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList {
	var returns TfHyperParameterTuningJob_TuningJobCompletionCriteriaPropertyList
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) TuningJobCompletionCriteriaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuningJobCompletionCriteriaInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_ConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_ConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_ConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_ConfigPropertyOutputReference_Override(t TfHyperParameterTuningJob_ConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.ConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetRandomSeed(val *float64) {
	if err := j.validateSetRandomSeedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"randomSeed",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetStrategy(val *string) {
	if err := j.validateSetStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strategy",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference)SetTrainingJobEarlyStoppingType(val *string) {
	if err := j.validateSetTrainingJobEarlyStoppingTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingJobEarlyStoppingType",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) PutObjective(value interface{}) {
	if err := t.validatePutObjectiveParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putObjective",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) PutParameterRanges(value interface{}) {
	if err := t.validatePutParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putParameterRanges",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) PutResourceLimits(value interface{}) {
	if err := t.validatePutResourceLimitsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceLimits",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) PutStrategyConfig(value interface{}) {
	if err := t.validatePutStrategyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStrategyConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) PutTuningJobCompletionCriteria(value interface{}) {
	if err := t.validatePutTuningJobCompletionCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTuningJobCompletionCriteria",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetObjective() {
	_jsii_.InvokeVoid(
		t,
		"resetObjective",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetParameterRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetParameterRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetRandomSeed() {
	_jsii_.InvokeVoid(
		t,
		"resetRandomSeed",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetResourceLimits() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceLimits",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetStrategyConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetStrategyConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetTrainingJobEarlyStoppingType() {
	_jsii_.InvokeVoid(
		t,
		"resetTrainingJobEarlyStoppingType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ResetTuningJobCompletionCriteria() {
	_jsii_.InvokeVoid(
		t,
		"resetTuningJobCompletionCriteria",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

