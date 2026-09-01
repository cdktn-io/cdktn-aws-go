package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightDashboard_ParametersPropertyOutputReference interface {
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
	DateTimeParameters() AwsQuicksightDashboard_DateTimeParametersPropertyList
	// Experimental.
	DateTimeParametersInput() interface{}
	// Experimental.
	DecimalParameters() AwsQuicksightDashboard_DecimalParametersPropertyList
	// Experimental.
	DecimalParametersInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	IntegerParameters() AwsQuicksightDashboard_IntegerParametersPropertyList
	// Experimental.
	IntegerParametersInput() interface{}
	// Experimental.
	InternalValue() *AwsQuicksightDashboard_ParametersProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightDashboard_ParametersProperty)
	// Experimental.
	StringParameters() AwsQuicksightDashboard_StringParametersPropertyList
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

// The jsii proxy struct for AwsQuicksightDashboard_ParametersPropertyOutputReference
type jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) DateTimeParameters() AwsQuicksightDashboard_DateTimeParametersPropertyList {
	var returns AwsQuicksightDashboard_DateTimeParametersPropertyList
	_jsii_.Get(
		j,
		"dateTimeParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) DateTimeParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dateTimeParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) DecimalParameters() AwsQuicksightDashboard_DecimalParametersPropertyList {
	var returns AwsQuicksightDashboard_DecimalParametersPropertyList
	_jsii_.Get(
		j,
		"decimalParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) DecimalParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decimalParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) IntegerParameters() AwsQuicksightDashboard_IntegerParametersPropertyList {
	var returns AwsQuicksightDashboard_IntegerParametersPropertyList
	_jsii_.Get(
		j,
		"integerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) IntegerParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"integerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) InternalValue() *AwsQuicksightDashboard_ParametersProperty {
	var returns *AwsQuicksightDashboard_ParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) StringParameters() AwsQuicksightDashboard_StringParametersPropertyList {
	var returns AwsQuicksightDashboard_StringParametersPropertyList
	_jsii_.Get(
		j,
		"stringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) StringParametersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightDashboard_ParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightDashboard_ParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightDashboard_ParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDashboard.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightDashboard_ParametersPropertyOutputReference_Override(a AwsQuicksightDashboard_ParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDashboard.ParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference)SetInternalValue(val *AwsQuicksightDashboard_ParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) PutDateTimeParameters(value interface{}) {
	if err := a.validatePutDateTimeParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDateTimeParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) PutDecimalParameters(value interface{}) {
	if err := a.validatePutDecimalParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDecimalParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) PutIntegerParameters(value interface{}) {
	if err := a.validatePutIntegerParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntegerParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) PutStringParameters(value interface{}) {
	if err := a.validatePutStringParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ResetDateTimeParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDateTimeParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ResetDecimalParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetDecimalParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ResetIntegerParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetIntegerParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ResetStringParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetStringParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuicksightDashboard_ParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

