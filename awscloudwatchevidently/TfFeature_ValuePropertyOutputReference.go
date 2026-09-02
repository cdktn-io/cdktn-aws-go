package awscloudwatchevidently

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchevidently/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchevidently/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFeature_ValuePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BoolValue() *string
	// Experimental.
	SetBoolValue(val *string)
	// Experimental.
	BoolValueInput() *string
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
	DoubleValue() *string
	// Experimental.
	SetDoubleValue(val *string)
	// Experimental.
	DoubleValueInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFeature_ValueProperty
	// Experimental.
	SetInternalValue(val *TfFeature_ValueProperty)
	// Experimental.
	LongValue() *string
	// Experimental.
	SetLongValue(val *string)
	// Experimental.
	LongValueInput() *string
	// Experimental.
	StringValue() *string
	// Experimental.
	SetStringValue(val *string)
	// Experimental.
	StringValueInput() *string
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
	ResetBoolValue()
	// Experimental.
	ResetDoubleValue()
	// Experimental.
	ResetLongValue()
	// Experimental.
	ResetStringValue()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFeature_ValuePropertyOutputReference
type jsiiProxy_TfFeature_ValuePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) BoolValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boolValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) BoolValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boolValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) DoubleValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doubleValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) DoubleValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"doubleValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) InternalValue() *TfFeature_ValueProperty {
	var returns *TfFeature_ValueProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) LongValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) LongValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"longValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) StringValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) StringValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stringValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFeature_ValuePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFeature_ValuePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFeature_ValuePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFeature_ValuePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-evidently.TfFeature.ValuePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFeature_ValuePropertyOutputReference_Override(t TfFeature_ValuePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-evidently.TfFeature.ValuePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetBoolValue(val *string) {
	if err := j.validateSetBoolValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boolValue",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetDoubleValue(val *string) {
	if err := j.validateSetDoubleValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doubleValue",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetInternalValue(val *TfFeature_ValueProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetLongValue(val *string) {
	if err := j.validateSetLongValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"longValue",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetStringValue(val *string) {
	if err := j.validateSetStringValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stringValue",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFeature_ValuePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ResetBoolValue() {
	_jsii_.InvokeVoid(
		t,
		"resetBoolValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ResetDoubleValue() {
	_jsii_.InvokeVoid(
		t,
		"resetDoubleValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ResetLongValue() {
	_jsii_.InvokeVoid(
		t,
		"resetLongValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ResetStringValue() {
	_jsii_.InvokeVoid(
		t,
		"resetStringValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFeature_ValuePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

