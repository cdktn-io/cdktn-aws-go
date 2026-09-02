package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoParameters() TfHyperParameterTuningJob_ConfigParameterRangesAutoParametersPropertyList
	// Experimental.
	AutoParametersInput() interface{}
	// Experimental.
	CategoricalParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesCategoricalParameterRangesPropertyList
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
	ContinuousParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesContinuousParameterRangesPropertyList
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
	IntegerParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesIntegerParameterRangesPropertyList
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

// The jsii proxy struct for TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference
type jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) AutoParameters() TfHyperParameterTuningJob_ConfigParameterRangesAutoParametersPropertyList {
	var returns TfHyperParameterTuningJob_ConfigParameterRangesAutoParametersPropertyList
	_jsii_.Get(
		j,
		"autoParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) AutoParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) CategoricalParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesCategoricalParameterRangesPropertyList {
	var returns TfHyperParameterTuningJob_ConfigParameterRangesCategoricalParameterRangesPropertyList
	_jsii_.Get(
		j,
		"categoricalParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) CategoricalParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"categoricalParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ContinuousParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesContinuousParameterRangesPropertyList {
	var returns TfHyperParameterTuningJob_ConfigParameterRangesContinuousParameterRangesPropertyList
	_jsii_.Get(
		j,
		"continuousParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ContinuousParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"continuousParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) IntegerParameterRanges() TfHyperParameterTuningJob_ConfigParameterRangesIntegerParameterRangesPropertyList {
	var returns TfHyperParameterTuningJob_ConfigParameterRangesIntegerParameterRangesPropertyList
	_jsii_.Get(
		j,
		"integerParameterRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) IntegerParameterRangesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParameterRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfHyperParameterTuningJob_ParameterRangesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfHyperParameterTuningJob_ParameterRangesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.ParameterRangesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfHyperParameterTuningJob_ParameterRangesPropertyOutputReference_Override(t TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfHyperParameterTuningJob.ParameterRangesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) PutAutoParameters(value interface{}) {
	if err := t.validatePutAutoParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutoParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) PutCategoricalParameterRanges(value interface{}) {
	if err := t.validatePutCategoricalParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCategoricalParameterRanges",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) PutContinuousParameterRanges(value interface{}) {
	if err := t.validatePutContinuousParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContinuousParameterRanges",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) PutIntegerParameterRanges(value interface{}) {
	if err := t.validatePutIntegerParameterRangesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIntegerParameterRanges",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ResetAutoParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ResetCategoricalParameterRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetCategoricalParameterRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ResetContinuousParameterRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetContinuousParameterRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ResetIntegerParameterRanges() {
	_jsii_.InvokeVoid(
		t,
		"resetIntegerParameterRanges",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfHyperParameterTuningJob_ParameterRangesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

