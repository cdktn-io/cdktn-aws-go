package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bool() AwsSecurityhubConfigurationPolicy_BoolPropertyOutputReference
	// Experimental.
	BoolInput() *AwsSecurityhubConfigurationPolicy_BoolProperty
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
	Double() AwsSecurityhubConfigurationPolicy_DoublePropertyOutputReference
	// Experimental.
	DoubleInput() *AwsSecurityhubConfigurationPolicy_DoubleProperty
	// Experimental.
	Enum() AwsSecurityhubConfigurationPolicy_EnumPropertyOutputReference
	// Experimental.
	EnumInput() *AwsSecurityhubConfigurationPolicy_EnumProperty
	// Experimental.
	EnumList() AwsSecurityhubConfigurationPolicy_EnumListPropertyOutputReference
	// Experimental.
	EnumListInput() *AwsSecurityhubConfigurationPolicy_EnumListProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Int() AwsSecurityhubConfigurationPolicy_IntPropertyOutputReference
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	IntInput() *AwsSecurityhubConfigurationPolicy_IntProperty
	// Experimental.
	IntList() AwsSecurityhubConfigurationPolicy_IntListPropertyOutputReference
	// Experimental.
	IntListInput() *AwsSecurityhubConfigurationPolicy_IntListProperty
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	String() AwsSecurityhubConfigurationPolicy_StringPropertyOutputReference
	// Experimental.
	StringInput() *AwsSecurityhubConfigurationPolicy_StringProperty
	// Experimental.
	StringList() AwsSecurityhubConfigurationPolicy_StringListPropertyOutputReference
	// Experimental.
	StringListInput() *AwsSecurityhubConfigurationPolicy_StringListProperty
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
	PutBool(value *AwsSecurityhubConfigurationPolicy_BoolProperty)
	// Experimental.
	PutDouble(value *AwsSecurityhubConfigurationPolicy_DoubleProperty)
	// Experimental.
	PutEnum(value *AwsSecurityhubConfigurationPolicy_EnumProperty)
	// Experimental.
	PutEnumList(value *AwsSecurityhubConfigurationPolicy_EnumListProperty)
	// Experimental.
	PutInt(value *AwsSecurityhubConfigurationPolicy_IntProperty)
	// Experimental.
	PutIntList(value *AwsSecurityhubConfigurationPolicy_IntListProperty)
	// Experimental.
	PutString(value *AwsSecurityhubConfigurationPolicy_StringProperty)
	// Experimental.
	PutStringList(value *AwsSecurityhubConfigurationPolicy_StringListProperty)
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

// The jsii proxy struct for AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference
type jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Bool() AwsSecurityhubConfigurationPolicy_BoolPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_BoolPropertyOutputReference
	_jsii_.Get(
		j,
		"bool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) BoolInput() *AwsSecurityhubConfigurationPolicy_BoolProperty {
	var returns *AwsSecurityhubConfigurationPolicy_BoolProperty
	_jsii_.Get(
		j,
		"boolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Double() AwsSecurityhubConfigurationPolicy_DoublePropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_DoublePropertyOutputReference
	_jsii_.Get(
		j,
		"double",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) DoubleInput() *AwsSecurityhubConfigurationPolicy_DoubleProperty {
	var returns *AwsSecurityhubConfigurationPolicy_DoubleProperty
	_jsii_.Get(
		j,
		"doubleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Enum() AwsSecurityhubConfigurationPolicy_EnumPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_EnumPropertyOutputReference
	_jsii_.Get(
		j,
		"enum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) EnumInput() *AwsSecurityhubConfigurationPolicy_EnumProperty {
	var returns *AwsSecurityhubConfigurationPolicy_EnumProperty
	_jsii_.Get(
		j,
		"enumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) EnumList() AwsSecurityhubConfigurationPolicy_EnumListPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_EnumListPropertyOutputReference
	_jsii_.Get(
		j,
		"enumList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) EnumListInput() *AwsSecurityhubConfigurationPolicy_EnumListProperty {
	var returns *AwsSecurityhubConfigurationPolicy_EnumListProperty
	_jsii_.Get(
		j,
		"enumListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Int() AwsSecurityhubConfigurationPolicy_IntPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_IntPropertyOutputReference
	_jsii_.Get(
		j,
		"int",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) IntInput() *AwsSecurityhubConfigurationPolicy_IntProperty {
	var returns *AwsSecurityhubConfigurationPolicy_IntProperty
	_jsii_.Get(
		j,
		"intInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) IntList() AwsSecurityhubConfigurationPolicy_IntListPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_IntListPropertyOutputReference
	_jsii_.Get(
		j,
		"intList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) IntListInput() *AwsSecurityhubConfigurationPolicy_IntListProperty {
	var returns *AwsSecurityhubConfigurationPolicy_IntListProperty
	_jsii_.Get(
		j,
		"intListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) String() AwsSecurityhubConfigurationPolicy_StringPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_StringPropertyOutputReference
	_jsii_.Get(
		j,
		"string",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) StringInput() *AwsSecurityhubConfigurationPolicy_StringProperty {
	var returns *AwsSecurityhubConfigurationPolicy_StringProperty
	_jsii_.Get(
		j,
		"stringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) StringList() AwsSecurityhubConfigurationPolicy_StringListPropertyOutputReference {
	var returns AwsSecurityhubConfigurationPolicy_StringListPropertyOutputReference
	_jsii_.Get(
		j,
		"stringList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) StringListInput() *AwsSecurityhubConfigurationPolicy_StringListProperty {
	var returns *AwsSecurityhubConfigurationPolicy_StringListProperty
	_jsii_.Get(
		j,
		"stringListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ValueType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ValueTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"valueTypeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference_Override(a AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubConfigurationPolicy.ParameterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference)SetValueType(val *string) {
	if err := j.validateSetValueTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"valueType",
		val,
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutBool(value *AwsSecurityhubConfigurationPolicy_BoolProperty) {
	if err := a.validatePutBoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutDouble(value *AwsSecurityhubConfigurationPolicy_DoubleProperty) {
	if err := a.validatePutDoubleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDouble",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutEnum(value *AwsSecurityhubConfigurationPolicy_EnumProperty) {
	if err := a.validatePutEnumParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnum",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutEnumList(value *AwsSecurityhubConfigurationPolicy_EnumListProperty) {
	if err := a.validatePutEnumListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnumList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutInt(value *AwsSecurityhubConfigurationPolicy_IntProperty) {
	if err := a.validatePutIntParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInt",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutIntList(value *AwsSecurityhubConfigurationPolicy_IntListProperty) {
	if err := a.validatePutIntListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIntList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutString(value *AwsSecurityhubConfigurationPolicy_StringProperty) {
	if err := a.validatePutStringParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putString",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) PutStringList(value *AwsSecurityhubConfigurationPolicy_StringListProperty) {
	if err := a.validatePutStringListParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringList",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetBool() {
	_jsii_.InvokeVoid(
		a,
		"resetBool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetDouble() {
	_jsii_.InvokeVoid(
		a,
		"resetDouble",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetEnum() {
	_jsii_.InvokeVoid(
		a,
		"resetEnum",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetEnumList() {
	_jsii_.InvokeVoid(
		a,
		"resetEnumList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetInt() {
	_jsii_.InvokeVoid(
		a,
		"resetInt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetIntList() {
	_jsii_.InvokeVoid(
		a,
		"resetIntList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetString() {
	_jsii_.InvokeVoid(
		a,
		"resetString",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ResetStringList() {
	_jsii_.InvokeVoid(
		a,
		"resetStringList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSecurityhubConfigurationPolicy_ParameterPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

