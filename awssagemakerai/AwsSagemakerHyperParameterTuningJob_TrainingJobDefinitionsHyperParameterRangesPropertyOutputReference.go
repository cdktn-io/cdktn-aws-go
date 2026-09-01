package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoParameters() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersPropertyList
	// Experimental.
	AutoParametersInput() interface{}
	// Experimental.
	CategoricalParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesPropertyList
	// Experimental.
	CategoricalParameterRangesInput() interface{}
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
	// Experimental.
	ContinuousParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesPropertyList
	// Experimental.
	ContinuousParameterRangesInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	IntegerParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesPropertyList
	// Experimental.
	IntegerParameterRangesInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	PutAutoParameters(value interface{})
	// Experimental.
	PutCategoricalParameterRanges(value interface{})
	// Experimental.
	PutContinuousParameterRanges(value interface{})
	// Experimental.
	PutIntegerParameterRanges(value interface{})
	// Experimental.
	ResetAutoParameters()
	// Experimental.
	ResetCategoricalParameterRanges()
	// Experimental.
	ResetContinuousParameterRanges()
	// Experimental.
	ResetIntegerParameterRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference
type jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) AutoParameters() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersPropertyList {
	var returns AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersPropertyList
	_jsii_.Get(
		j,
		"autoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) AutoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) CategoricalParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesPropertyList {
	var returns AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesPropertyList
	_jsii_.Get(
		j,
		"categoricalParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) CategoricalParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoricalParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ContinuousParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesPropertyList {
	var returns AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesPropertyList
	_jsii_.Get(
		j,
		"continuousParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ContinuousParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) IntegerParameterRanges() AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesPropertyList {
	var returns AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesPropertyList
	_jsii_.Get(
		j,
		"integerParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) IntegerParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerHyperParameterTuningJob.TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference_Override(a AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerHyperParameterTuningJob.TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) PutAutoParameters(value interface{}) {
	if err := a.validatePutAutoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutoParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) PutCategoricalParameterRanges(value interface{}) {
	if err := a.validatePutCategoricalParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCategoricalParameterRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) PutContinuousParameterRanges(value interface{}) {
	if err := a.validatePutContinuousParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContinuousParameterRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) PutIntegerParameterRanges(value interface{}) {
	if err := a.validatePutIntegerParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegerParameterRanges",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ResetAutoParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ResetCategoricalParameterRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetCategoricalParameterRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ResetContinuousParameterRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetContinuousParameterRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ResetIntegerParameterRanges() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegerParameterRanges",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

