package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserPool_SchemaPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AttributeDataType() *string
	// Experimental.
	SetAttributeDataType(val *string)
	// Experimental.
	AttributeDataTypeInput() *string
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
	DeveloperOnlyAttribute() interface{}
	// Experimental.
	SetDeveloperOnlyAttribute(val interface{})
	// Experimental.
	DeveloperOnlyAttributeInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Mutable() interface{}
	// Experimental.
	SetMutable(val interface{})
	// Experimental.
	MutableInput() interface{}
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	NumberAttributeConstraints() TfUserPool_NumberAttributeConstraintsPropertyOutputReference
	// Experimental.
	NumberAttributeConstraintsInput() *TfUserPool_NumberAttributeConstraintsProperty
	// Experimental.
	Required() interface{}
	// Experimental.
	SetRequired(val interface{})
	// Experimental.
	RequiredInput() interface{}
	// Experimental.
	StringAttributeConstraints() TfUserPool_StringAttributeConstraintsPropertyOutputReference
	// Experimental.
	StringAttributeConstraintsInput() *TfUserPool_StringAttributeConstraintsProperty
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
	PutNumberAttributeConstraints(value *TfUserPool_NumberAttributeConstraintsProperty)
	// Experimental.
	PutStringAttributeConstraints(value *TfUserPool_StringAttributeConstraintsProperty)
	// Experimental.
	ResetDeveloperOnlyAttribute()
	// Experimental.
	ResetMutable()
	// Experimental.
	ResetNumberAttributeConstraints()
	// Experimental.
	ResetRequired()
	// Experimental.
	ResetStringAttributeConstraints()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUserPool_SchemaPropertyOutputReference
type jsiiProxy_TfUserPool_SchemaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) AttributeDataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) AttributeDataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) DeveloperOnlyAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) DeveloperOnlyAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) Mutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) MutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) NumberAttributeConstraints() TfUserPool_NumberAttributeConstraintsPropertyOutputReference {
	var returns TfUserPool_NumberAttributeConstraintsPropertyOutputReference
	_jsii_.Get(
		j,
		"numberAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) NumberAttributeConstraintsInput() *TfUserPool_NumberAttributeConstraintsProperty {
	var returns *TfUserPool_NumberAttributeConstraintsProperty
	_jsii_.Get(
		j,
		"numberAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) StringAttributeConstraints() TfUserPool_StringAttributeConstraintsPropertyOutputReference {
	var returns TfUserPool_StringAttributeConstraintsPropertyOutputReference
	_jsii_.Get(
		j,
		"stringAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) StringAttributeConstraintsInput() *TfUserPool_StringAttributeConstraintsProperty {
	var returns *TfUserPool_StringAttributeConstraintsProperty
	_jsii_.Get(
		j,
		"stringAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserPool_SchemaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfUserPool_SchemaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserPool_SchemaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserPool_SchemaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.SchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserPool_SchemaPropertyOutputReference_Override(t TfUserPool_SchemaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.TfUserPool.SchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetAttributeDataType(val *string) {
	if err := j.validateSetAttributeDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeDataType",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetDeveloperOnlyAttribute(val interface{}) {
	if err := j.validateSetDeveloperOnlyAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerOnlyAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetMutable(val interface{}) {
	if err := j.validateSetMutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutable",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserPool_SchemaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) PutNumberAttributeConstraints(value *TfUserPool_NumberAttributeConstraintsProperty) {
	if err := t.validatePutNumberAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNumberAttributeConstraints",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) PutStringAttributeConstraints(value *TfUserPool_StringAttributeConstraintsProperty) {
	if err := t.validatePutStringAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStringAttributeConstraints",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ResetDeveloperOnlyAttribute() {
	_jsii_.InvokeVoid(
		t,
		"resetDeveloperOnlyAttribute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ResetMutable() {
	_jsii_.InvokeVoid(
		t,
		"resetMutable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ResetNumberAttributeConstraints() {
	_jsii_.InvokeVoid(
		t,
		"resetNumberAttributeConstraints",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		t,
		"resetRequired",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ResetStringAttributeConstraints() {
	_jsii_.InvokeVoid(
		t,
		"resetStringAttributeConstraints",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserPool_SchemaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

