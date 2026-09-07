package securityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/securityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/securityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConfigurationPolicy_ParameterPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bool() AwsConfigurationPolicy_BoolPropertyOutputReference
	// Experimental.
	BoolInput() *AwsConfigurationPolicy_BoolProperty
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
	Double() AwsConfigurationPolicy_DoublePropertyOutputReference
	// Experimental.
	DoubleInput() *AwsConfigurationPolicy_DoubleProperty
	// Experimental.
	Enum() AwsConfigurationPolicy_EnumPropertyOutputReference
	// Experimental.
	EnumInput() *AwsConfigurationPolicy_EnumProperty
	// Experimental.
	EnumList() AwsConfigurationPolicy_EnumListPropertyOutputReference
	// Experimental.
	EnumListInput() *AwsConfigurationPolicy_EnumListProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Int() AwsConfigurationPolicy_IntPropertyOutputReference
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IntInput() *AwsConfigurationPolicy_IntProperty
	// Experimental.
	IntList() AwsConfigurationPolicy_IntListPropertyOutputReference
	// Experimental.
	IntListInput() *AwsConfigurationPolicy_IntListProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	String() AwsConfigurationPolicy_StringPropertyOutputReference
	// Experimental.
	StringInput() *AwsConfigurationPolicy_StringProperty
	// Experimental.
	StringList() AwsConfigurationPolicy_StringListPropertyOutputReference
	// Experimental.
	StringListInput() *AwsConfigurationPolicy_StringListProperty
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
	PutBool(value *AwsConfigurationPolicy_BoolProperty)
	// Experimental.
	PutDouble(value *AwsConfigurationPolicy_DoubleProperty)
	// Experimental.
	PutEnum(value *AwsConfigurationPolicy_EnumProperty)
	// Experimental.
	PutEnumList(value *AwsConfigurationPolicy_EnumListProperty)
	// Experimental.
	PutInt(value *AwsConfigurationPolicy_IntProperty)
	// Experimental.
	PutIntList(value *AwsConfigurationPolicy_IntListProperty)
	// Experimental.
	PutString(value *AwsConfigurationPolicy_StringProperty)
	// Experimental.
	PutStringList(value *AwsConfigurationPolicy_StringListProperty)
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

// The jsii proxy struct for AwsConfigurationPolicy_ParameterPropertyOutputReference
type jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Bool() AwsConfigurationPolicy_BoolPropertyOutputReference {
	var returns AwsConfigurationPolicy_BoolPropertyOutputReference
	_jsii_.Get(
		j,
		"bool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) BoolInput() *AwsConfigurationPolicy_BoolProperty {
	var returns *AwsConfigurationPolicy_BoolProperty
	_jsii_.Get(
		j,
		"boolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Double() AwsConfigurationPolicy_DoublePropertyOutputReference {
	var returns AwsConfigurationPolicy_DoublePropertyOutputReference
	_jsii_.Get(
		j,
		"double",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) DoubleInput() *AwsConfigurationPolicy_DoubleProperty {
	var returns *AwsConfigurationPolicy_DoubleProperty
	_jsii_.Get(
		j,
		"doubleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Enum() AwsConfigurationPolicy_EnumPropertyOutputReference {
	var returns AwsConfigurationPolicy_EnumPropertyOutputReference
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) EnumInput() *AwsConfigurationPolicy_EnumProperty {
	var returns *AwsConfigurationPolicy_EnumProperty
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) EnumList() AwsConfigurationPolicy_EnumListPropertyOutputReference {
	var returns AwsConfigurationPolicy_EnumListPropertyOutputReference
	_jsii_.Get(
		j,
		"enumList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) EnumListInput() *AwsConfigurationPolicy_EnumListProperty {
	var returns *AwsConfigurationPolicy_EnumListProperty
	_jsii_.Get(
		j,
		"enumListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Int() AwsConfigurationPolicy_IntPropertyOutputReference {
	var returns AwsConfigurationPolicy_IntPropertyOutputReference
	_jsii_.Get(
		j,
		"int",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) IntInput() *AwsConfigurationPolicy_IntProperty {
	var returns *AwsConfigurationPolicy_IntProperty
	_jsii_.Get(
		j,
		"intInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) IntList() AwsConfigurationPolicy_IntListPropertyOutputReference {
	var returns AwsConfigurationPolicy_IntListPropertyOutputReference
	_jsii_.Get(
		j,
		"intList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) IntListInput() *AwsConfigurationPolicy_IntListProperty {
	var returns *AwsConfigurationPolicy_IntListProperty
	_jsii_.Get(
		j,
		"intListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) String() AwsConfigurationPolicy_StringPropertyOutputReference {
	var returns AwsConfigurationPolicy_StringPropertyOutputReference
	_jsii_.Get(
		j,
		"string",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) StringInput() *AwsConfigurationPolicy_StringProperty {
	var returns *AwsConfigurationPolicy_StringProperty
	_jsii_.Get(
		j,
		"stringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) StringList() AwsConfigurationPolicy_StringListPropertyOutputReference {
	var returns AwsConfigurationPolicy_StringListPropertyOutputReference
	_jsii_.Get(
		j,
		"stringList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) StringListInput() *AwsConfigurationPolicy_StringListProperty {
	var returns *AwsConfigurationPolicy_StringListProperty
	_jsii_.Get(
		j,
		"stringListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConfigurationPolicy_ParameterPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsConfigurationPolicy_ParameterPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConfigurationPolicy_ParameterPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConfigurationPolicy_ParameterPropertyOutputReference_Override(a AwsConfigurationPolicy_ParameterPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference)SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutBool(value *AwsConfigurationPolicy_BoolProperty) {
	if err := a.validatePutBoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutDouble(value *AwsConfigurationPolicy_DoubleProperty) {
	if err := a.validatePutDoubleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDouble",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutEnum(value *AwsConfigurationPolicy_EnumProperty) {
	if err := a.validatePutEnumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnum",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutEnumList(value *AwsConfigurationPolicy_EnumListProperty) {
	if err := a.validatePutEnumListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnumList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutInt(value *AwsConfigurationPolicy_IntProperty) {
	if err := a.validatePutIntParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutIntList(value *AwsConfigurationPolicy_IntListProperty) {
	if err := a.validatePutIntListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutString(value *AwsConfigurationPolicy_StringProperty) {
	if err := a.validatePutStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) PutStringList(value *AwsConfigurationPolicy_StringListProperty) {
	if err := a.validatePutStringListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetBool() {
	_jsii_.InvokeVoid(
		a,
		"resetBool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetDouble() {
	_jsii_.InvokeVoid(
		a,
		"resetDouble",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		a,
		"resetEnum",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetEnumList() {
	_jsii_.InvokeVoid(
		a,
		"resetEnumList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetInt() {
	_jsii_.InvokeVoid(
		a,
		"resetInt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetIntList() {
	_jsii_.InvokeVoid(
		a,
		"resetIntList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetString() {
	_jsii_.InvokeVoid(
		a,
		"resetString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ResetStringList() {
	_jsii_.InvokeVoid(
		a,
		"resetStringList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConfigurationPolicy_ParameterPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

