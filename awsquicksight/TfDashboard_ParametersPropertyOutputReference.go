package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDashboard_ParametersPropertyOutputReference interface {
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
	DateTimeParameters() TfDashboard_DateTimeParametersPropertyList
	// Experimental.
	DateTimeParametersInput() interface{}
	// Experimental.
	DecimalParameters() TfDashboard_DecimalParametersPropertyList
	// Experimental.
	DecimalParametersInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IntegerParameters() TfDashboard_IntegerParametersPropertyList
	// Experimental.
	IntegerParametersInput() interface{}
	// Experimental.
	InternalValue() *TfDashboard_ParametersProperty
	// Experimental.
	SetInternalValue(val *TfDashboard_ParametersProperty)
	// Experimental.
	StringParameters() TfDashboard_StringParametersPropertyList
	// Experimental.
	StringParametersInput() interface{}
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
	PutDateTimeParameters(value interface{})
	// Experimental.
	PutDecimalParameters(value interface{})
	// Experimental.
	PutIntegerParameters(value interface{})
	// Experimental.
	PutStringParameters(value interface{})
	// Experimental.
	ResetDateTimeParameters()
	// Experimental.
	ResetDecimalParameters()
	// Experimental.
	ResetIntegerParameters()
	// Experimental.
	ResetStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDashboard_ParametersPropertyOutputReference
type jsiiProxy_TfDashboard_ParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) DateTimeParameters() TfDashboard_DateTimeParametersPropertyList {
	var returns TfDashboard_DateTimeParametersPropertyList
	_jsii_.Get(
		j,
		"dateTimeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) DateTimeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) DecimalParameters() TfDashboard_DecimalParametersPropertyList {
	var returns TfDashboard_DecimalParametersPropertyList
	_jsii_.Get(
		j,
		"decimalParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) DecimalParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decimalParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) IntegerParameters() TfDashboard_IntegerParametersPropertyList {
	var returns TfDashboard_IntegerParametersPropertyList
	_jsii_.Get(
		j,
		"integerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) IntegerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) InternalValue() *TfDashboard_ParametersProperty {
	var returns *TfDashboard_ParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) StringParameters() TfDashboard_StringParametersPropertyList {
	var returns TfDashboard_StringParametersPropertyList
	_jsii_.Get(
		j,
		"stringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) StringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDashboard_ParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDashboard_ParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDashboard_ParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDashboard_ParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDashboard.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDashboard_ParametersPropertyOutputReference_Override(t TfDashboard_ParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDashboard.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference)SetInternalValue(val *TfDashboard_ParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_ParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) PutDateTimeParameters(value interface{}) {
	if err := t.validatePutDateTimeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDateTimeParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) PutDecimalParameters(value interface{}) {
	if err := t.validatePutDecimalParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDecimalParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) PutIntegerParameters(value interface{}) {
	if err := t.validatePutIntegerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIntegerParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) PutStringParameters(value interface{}) {
	if err := t.validatePutStringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStringParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ResetDateTimeParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetDateTimeParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ResetDecimalParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetDecimalParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ResetIntegerParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetIntegerParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ResetStringParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetStringParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDashboard_ParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

