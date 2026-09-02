package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationPolicy_ParameterPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bool() TfConfigurationPolicy_BoolPropertyOutputReference
	// Experimental.
	BoolInput() *TfConfigurationPolicy_BoolProperty
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
	Double() TfConfigurationPolicy_DoublePropertyOutputReference
	// Experimental.
	DoubleInput() *TfConfigurationPolicy_DoubleProperty
	// Experimental.
	Enum() TfConfigurationPolicy_EnumPropertyOutputReference
	// Experimental.
	EnumInput() *TfConfigurationPolicy_EnumProperty
	// Experimental.
	EnumList() TfConfigurationPolicy_EnumListPropertyOutputReference
	// Experimental.
	EnumListInput() *TfConfigurationPolicy_EnumListProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Int() TfConfigurationPolicy_IntPropertyOutputReference
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IntInput() *TfConfigurationPolicy_IntProperty
	// Experimental.
	IntList() TfConfigurationPolicy_IntListPropertyOutputReference
	// Experimental.
	IntListInput() *TfConfigurationPolicy_IntListProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	String() TfConfigurationPolicy_StringPropertyOutputReference
	// Experimental.
	StringInput() *TfConfigurationPolicy_StringProperty
	// Experimental.
	StringList() TfConfigurationPolicy_StringListPropertyOutputReference
	// Experimental.
	StringListInput() *TfConfigurationPolicy_StringListProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ValueType() *string
	// Experimental.
	SetValueType(val *string)
	// Experimental.
	ValueTypeInput() *string
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
	PutBool(value *TfConfigurationPolicy_BoolProperty)
	// Experimental.
	PutDouble(value *TfConfigurationPolicy_DoubleProperty)
	// Experimental.
	PutEnum(value *TfConfigurationPolicy_EnumProperty)
	// Experimental.
	PutEnumList(value *TfConfigurationPolicy_EnumListProperty)
	// Experimental.
	PutInt(value *TfConfigurationPolicy_IntProperty)
	// Experimental.
	PutIntList(value *TfConfigurationPolicy_IntListProperty)
	// Experimental.
	PutString(value *TfConfigurationPolicy_StringProperty)
	// Experimental.
	PutStringList(value *TfConfigurationPolicy_StringListProperty)
	// Experimental.
	ResetBool()
	// Experimental.
	ResetDouble()
	// Experimental.
	ResetEnum()
	// Experimental.
	ResetEnumList()
	// Experimental.
	ResetInt()
	// Experimental.
	ResetIntList()
	// Experimental.
	ResetString()
	// Experimental.
	ResetStringList()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationPolicy_ParameterPropertyOutputReference
type jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Bool() TfConfigurationPolicy_BoolPropertyOutputReference {
	var returns TfConfigurationPolicy_BoolPropertyOutputReference
	_jsii_.Get(
		j,
		"bool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) BoolInput() *TfConfigurationPolicy_BoolProperty {
	var returns *TfConfigurationPolicy_BoolProperty
	_jsii_.Get(
		j,
		"boolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Double() TfConfigurationPolicy_DoublePropertyOutputReference {
	var returns TfConfigurationPolicy_DoublePropertyOutputReference
	_jsii_.Get(
		j,
		"double",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) DoubleInput() *TfConfigurationPolicy_DoubleProperty {
	var returns *TfConfigurationPolicy_DoubleProperty
	_jsii_.Get(
		j,
		"doubleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Enum() TfConfigurationPolicy_EnumPropertyOutputReference {
	var returns TfConfigurationPolicy_EnumPropertyOutputReference
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) EnumInput() *TfConfigurationPolicy_EnumProperty {
	var returns *TfConfigurationPolicy_EnumProperty
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) EnumList() TfConfigurationPolicy_EnumListPropertyOutputReference {
	var returns TfConfigurationPolicy_EnumListPropertyOutputReference
	_jsii_.Get(
		j,
		"enumList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) EnumListInput() *TfConfigurationPolicy_EnumListProperty {
	var returns *TfConfigurationPolicy_EnumListProperty
	_jsii_.Get(
		j,
		"enumListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Int() TfConfigurationPolicy_IntPropertyOutputReference {
	var returns TfConfigurationPolicy_IntPropertyOutputReference
	_jsii_.Get(
		j,
		"int",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) IntInput() *TfConfigurationPolicy_IntProperty {
	var returns *TfConfigurationPolicy_IntProperty
	_jsii_.Get(
		j,
		"intInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) IntList() TfConfigurationPolicy_IntListPropertyOutputReference {
	var returns TfConfigurationPolicy_IntListPropertyOutputReference
	_jsii_.Get(
		j,
		"intList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) IntListInput() *TfConfigurationPolicy_IntListProperty {
	var returns *TfConfigurationPolicy_IntListProperty
	_jsii_.Get(
		j,
		"intListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) String() TfConfigurationPolicy_StringPropertyOutputReference {
	var returns TfConfigurationPolicy_StringPropertyOutputReference
	_jsii_.Get(
		j,
		"string",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) StringInput() *TfConfigurationPolicy_StringProperty {
	var returns *TfConfigurationPolicy_StringProperty
	_jsii_.Get(
		j,
		"stringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) StringList() TfConfigurationPolicy_StringListPropertyOutputReference {
	var returns TfConfigurationPolicy_StringListPropertyOutputReference
	_jsii_.Get(
		j,
		"stringList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) StringListInput() *TfConfigurationPolicy_StringListProperty {
	var returns *TfConfigurationPolicy_StringListProperty
	_jsii_.Get(
		j,
		"stringListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationPolicy_ParameterPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfConfigurationPolicy_ParameterPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationPolicy_ParameterPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationPolicy_ParameterPropertyOutputReference_Override(t TfConfigurationPolicy_ParameterPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference)SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutBool(value *TfConfigurationPolicy_BoolProperty) {
	if err := t.validatePutBoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBool",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutDouble(value *TfConfigurationPolicy_DoubleProperty) {
	if err := t.validatePutDoubleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDouble",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutEnum(value *TfConfigurationPolicy_EnumProperty) {
	if err := t.validatePutEnumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnum",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutEnumList(value *TfConfigurationPolicy_EnumListProperty) {
	if err := t.validatePutEnumListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEnumList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutInt(value *TfConfigurationPolicy_IntProperty) {
	if err := t.validatePutIntParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInt",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutIntList(value *TfConfigurationPolicy_IntListProperty) {
	if err := t.validatePutIntListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putIntList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutString(value *TfConfigurationPolicy_StringProperty) {
	if err := t.validatePutStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putString",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) PutStringList(value *TfConfigurationPolicy_StringListProperty) {
	if err := t.validatePutStringListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStringList",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetBool() {
	_jsii_.InvokeVoid(
		t,
		"resetBool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetDouble() {
	_jsii_.InvokeVoid(
		t,
		"resetDouble",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		t,
		"resetEnum",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetEnumList() {
	_jsii_.InvokeVoid(
		t,
		"resetEnumList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetInt() {
	_jsii_.InvokeVoid(
		t,
		"resetInt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetIntList() {
	_jsii_.InvokeVoid(
		t,
		"resetIntList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetString() {
	_jsii_.InvokeVoid(
		t,
		"resetString",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ResetStringList() {
	_jsii_.InvokeVoid(
		t,
		"resetStringList",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationPolicy_ParameterPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

