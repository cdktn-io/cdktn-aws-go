package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference interface {
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
	HeaderParameters() *map[string]*string
	// Experimental.
	SetHeaderParameters(val *map[string]*string)
	// Experimental.
	HeaderParametersInput() *map[string]*string
	// Experimental.
	InternalValue() *TfPipe_EnrichmentParametersHttpParametersProperty
	// Experimental.
	SetInternalValue(val *TfPipe_EnrichmentParametersHttpParametersProperty)
	// Experimental.
	PathParameterValues() *[]*string
	// Experimental.
	SetPathParameterValues(val *[]*string)
	// Experimental.
	PathParameterValuesInput() *[]*string
	// Experimental.
	QueryStringParameters() *map[string]*string
	// Experimental.
	SetQueryStringParameters(val *map[string]*string)
	// Experimental.
	QueryStringParametersInput() *map[string]*string
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
	ResetHeaderParameters()
	// Experimental.
	ResetPathParameterValues()
	// Experimental.
	ResetQueryStringParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference
type jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) HeaderParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headerParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) HeaderParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"headerParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) InternalValue() *TfPipe_EnrichmentParametersHttpParametersProperty {
	var returns *TfPipe_EnrichmentParametersHttpParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) PathParameterValues() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathParameterValues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) PathParameterValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pathParameterValuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) QueryStringParameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"queryStringParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) QueryStringParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"queryStringParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_EnrichmentParametersHttpParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_EnrichmentParametersHttpParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.EnrichmentParametersHttpParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_EnrichmentParametersHttpParametersPropertyOutputReference_Override(t TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.EnrichmentParametersHttpParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetHeaderParameters(val *map[string]*string) {
	if err := j.validateSetHeaderParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headerParameters",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetInternalValue(val *TfPipe_EnrichmentParametersHttpParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetPathParameterValues(val *[]*string) {
	if err := j.validateSetPathParameterValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathParameterValues",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetQueryStringParameters(val *map[string]*string) {
	if err := j.validateSetQueryStringParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringParameters",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ResetHeaderParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetHeaderParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ResetPathParameterValues() {
	_jsii_.InvokeVoid(
		t,
		"resetPathParameterValues",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ResetQueryStringParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryStringParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_EnrichmentParametersHttpParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

