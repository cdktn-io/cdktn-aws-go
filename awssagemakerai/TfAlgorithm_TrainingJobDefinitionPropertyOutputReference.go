package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlgorithm_TrainingJobDefinitionPropertyOutputReference interface {
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
	HyperParameters() *map[string]*string
	// Experimental.
	SetHyperParameters(val *map[string]*string)
	// Experimental.
	HyperParametersInput() *map[string]*string
	// Experimental.
	InputDataConfig() TfAlgorithm_InputDataConfigPropertyList
	// Experimental.
	InputDataConfigInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutputDataConfig() TfAlgorithm_OutputDataConfigPropertyList
	// Experimental.
	OutputDataConfigInput() interface{}
	// Experimental.
	ResourceConfig() TfAlgorithm_ResourceConfigPropertyList
	// Experimental.
	ResourceConfigInput() interface{}
	// Experimental.
	StoppingCondition() TfAlgorithm_StoppingConditionPropertyList
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
	TrainingInputMode() *string
	// Experimental.
	SetTrainingInputMode(val *string)
	// Experimental.
	TrainingInputModeInput() *string
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
	PutInputDataConfig(value interface{})
	// Experimental.
	PutOutputDataConfig(value interface{})
	// Experimental.
	PutResourceConfig(value interface{})
	// Experimental.
	PutStoppingCondition(value interface{})
	// Experimental.
	ResetHyperParameters()
	// Experimental.
	ResetInputDataConfig()
	// Experimental.
	ResetOutputDataConfig()
	// Experimental.
	ResetResourceConfig()
	// Experimental.
	ResetStoppingCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAlgorithm_TrainingJobDefinitionPropertyOutputReference
type jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) HyperParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) HyperParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"hyperParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) InputDataConfig() TfAlgorithm_InputDataConfigPropertyList {
	var returns TfAlgorithm_InputDataConfigPropertyList
	_jsii_.Get(
		j,
		"inputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) InputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"inputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) OutputDataConfig() TfAlgorithm_OutputDataConfigPropertyList {
	var returns TfAlgorithm_OutputDataConfigPropertyList
	_jsii_.Get(
		j,
		"outputDataConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) OutputDataConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"outputDataConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResourceConfig() TfAlgorithm_ResourceConfigPropertyList {
	var returns TfAlgorithm_ResourceConfigPropertyList
	_jsii_.Get(
		j,
		"resourceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResourceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resourceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) StoppingCondition() TfAlgorithm_StoppingConditionPropertyList {
	var returns TfAlgorithm_StoppingConditionPropertyList
	_jsii_.Get(
		j,
		"stoppingCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) StoppingConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stoppingConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) TrainingInputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) TrainingInputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlgorithm_TrainingJobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlgorithm_TrainingJobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlgorithm_TrainingJobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.TrainingJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlgorithm_TrainingJobDefinitionPropertyOutputReference_Override(t TfAlgorithm_TrainingJobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAlgorithm.TrainingJobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetHyperParameters(val *map[string]*string) {
	if err := j.validateSetHyperParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperParameters",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference)SetTrainingInputMode(val *string) {
	if err := j.validateSetTrainingInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingInputMode",
		val,
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) PutInputDataConfig(value interface{}) {
	if err := t.validatePutInputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) PutOutputDataConfig(value interface{}) {
	if err := t.validatePutOutputDataConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutputDataConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) PutResourceConfig(value interface{}) {
	if err := t.validatePutResourceConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putResourceConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) PutStoppingCondition(value interface{}) {
	if err := t.validatePutStoppingConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStoppingCondition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResetHyperParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetHyperParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResetInputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetInputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResetOutputDataConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOutputDataConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResetResourceConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetResourceConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ResetStoppingCondition() {
	_jsii_.InvokeVoid(
		t,
		"resetStoppingCondition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlgorithm_TrainingJobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

