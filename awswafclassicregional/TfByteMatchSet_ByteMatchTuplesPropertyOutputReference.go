package awswafclassicregional

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswafclassicregional/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswafclassicregional/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfByteMatchSet_ByteMatchTuplesPropertyOutputReference interface {
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
	FieldToMatch() TfByteMatchSet_FieldToMatchPropertyOutputReference
	// Experimental.
	FieldToMatchInput() *TfByteMatchSet_FieldToMatchProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PositionalConstraint() *string
	// Experimental.
	SetPositionalConstraint(val *string)
	// Experimental.
	PositionalConstraintInput() *string
	// Experimental.
	TargetString() *string
	// Experimental.
	SetTargetString(val *string)
	// Experimental.
	TargetStringInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TextTransformation() *string
	// Experimental.
	SetTextTransformation(val *string)
	// Experimental.
	TextTransformationInput() *string
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
	PutFieldToMatch(value *TfByteMatchSet_FieldToMatchProperty)
	// Experimental.
	ResetTargetString()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfByteMatchSet_ByteMatchTuplesPropertyOutputReference
type jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) FieldToMatch() TfByteMatchSet_FieldToMatchPropertyOutputReference {
	var returns TfByteMatchSet_FieldToMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"fieldToMatch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) FieldToMatchInput() *TfByteMatchSet_FieldToMatchProperty {
	var returns *TfByteMatchSet_FieldToMatchProperty
	_jsii_.Get(
		j,
		"fieldToMatchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) PositionalConstraint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) PositionalConstraintInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"positionalConstraintInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TargetString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TargetStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TextTransformation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) TextTransformationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"textTransformationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfByteMatchSet_ByteMatchTuplesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfByteMatchSet_ByteMatchTuplesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfByteMatchSet_ByteMatchTuplesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-waf-classic-regional.TfByteMatchSet.ByteMatchTuplesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfByteMatchSet_ByteMatchTuplesPropertyOutputReference_Override(t TfByteMatchSet_ByteMatchTuplesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-waf-classic-regional.TfByteMatchSet.ByteMatchTuplesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetPositionalConstraint(val *string) {
	if err := j.validateSetPositionalConstraintParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"positionalConstraint",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTargetString(val *string) {
	if err := j.validateSetTargetStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetString",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference)SetTextTransformation(val *string) {
	if err := j.validateSetTextTransformationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"textTransformation",
		val,
	)
}

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) PutFieldToMatch(value *TfByteMatchSet_FieldToMatchProperty) {
	if err := t.validatePutFieldToMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFieldToMatch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) ResetTargetString() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfByteMatchSet_ByteMatchTuplesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

