package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AlgorithmName() *string
	// Experimental.
	SetAlgorithmName(val *string)
	// Experimental.
	AlgorithmNameInput() *string
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
	MetricDefinitions() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationMetricDefinitionsPropertyList
	// Experimental.
	MetricDefinitionsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrainingImage() *string
	// Experimental.
	SetTrainingImage(val *string)
	// Experimental.
	TrainingImageInput() *string
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
	PutMetricDefinitions(value interface{})
	// Experimental.
	ResetAlgorithmName()
	// Experimental.
	ResetMetricDefinitions()
	// Experimental.
	ResetTrainingImage()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference
type jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) AlgorithmName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) AlgorithmNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) MetricDefinitions() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationMetricDefinitionsPropertyList {
	var returns AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationMetricDefinitionsPropertyList
	_jsii_.Get(
		j,
		"metricDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) MetricDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"metricDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TrainingImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TrainingImageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingImageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TrainingInputMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) TrainingInputModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trainingInputModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerHyperParameterTuningJob.TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference_Override(a AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerHyperParameterTuningJob.TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetAlgorithmName(val *string) {
	if err := j.validateSetAlgorithmNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithmName",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetTrainingImage(val *string) {
	if err := j.validateSetTrainingImageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingImage",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference)SetTrainingInputMode(val *string) {
	if err := j.validateSetTrainingInputModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trainingInputMode",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) PutMetricDefinitions(value interface{}) {
	if err := a.validatePutMetricDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMetricDefinitions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ResetAlgorithmName() {
	_jsii_.InvokeVoid(
		a,
		"resetAlgorithmName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ResetMetricDefinitions() {
	_jsii_.InvokeVoid(
		a,
		"resetMetricDefinitions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ResetTrainingImage() {
	_jsii_.InvokeVoid(
		a,
		"resetTrainingImage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

