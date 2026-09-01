package awscognitoidp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidp/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidp/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCognitoUserPool_SchemaPropertyOutputReference interface {
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
	NumberAttributeConstraints() AwsCognitoUserPool_NumberAttributeConstraintsPropertyOutputReference
	// Experimental.
	NumberAttributeConstraintsInput() *AwsCognitoUserPool_NumberAttributeConstraintsProperty
	// Experimental.
	Required() interface{}
	// Experimental.
	SetRequired(val interface{})
	// Experimental.
	RequiredInput() interface{}
	// Experimental.
	StringAttributeConstraints() AwsCognitoUserPool_StringAttributeConstraintsPropertyOutputReference
	// Experimental.
	StringAttributeConstraintsInput() *AwsCognitoUserPool_StringAttributeConstraintsProperty
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
	PutNumberAttributeConstraints(value *AwsCognitoUserPool_NumberAttributeConstraintsProperty)
	// Experimental.
	PutStringAttributeConstraints(value *AwsCognitoUserPool_StringAttributeConstraintsProperty)
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

// The jsii proxy struct for AwsCognitoUserPool_SchemaPropertyOutputReference
type jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) AttributeDataType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) AttributeDataTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeDataTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) DeveloperOnlyAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) DeveloperOnlyAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"developerOnlyAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) Mutable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) MutableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mutableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) NumberAttributeConstraints() AwsCognitoUserPool_NumberAttributeConstraintsPropertyOutputReference {
	var returns AwsCognitoUserPool_NumberAttributeConstraintsPropertyOutputReference
	_jsii_.Get(
		j,
		"numberAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) NumberAttributeConstraintsInput() *AwsCognitoUserPool_NumberAttributeConstraintsProperty {
	var returns *AwsCognitoUserPool_NumberAttributeConstraintsProperty
	_jsii_.Get(
		j,
		"numberAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) Required() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) RequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) StringAttributeConstraints() AwsCognitoUserPool_StringAttributeConstraintsPropertyOutputReference {
	var returns AwsCognitoUserPool_StringAttributeConstraintsPropertyOutputReference
	_jsii_.Get(
		j,
		"stringAttributeConstraints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) StringAttributeConstraintsInput() *AwsCognitoUserPool_StringAttributeConstraintsProperty {
	var returns *AwsCognitoUserPool_StringAttributeConstraintsProperty
	_jsii_.Get(
		j,
		"stringAttributeConstraintsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCognitoUserPool_SchemaPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCognitoUserPool_SchemaPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCognitoUserPool_SchemaPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.SchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCognitoUserPool_SchemaPropertyOutputReference_Override(a AwsCognitoUserPool_SchemaPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-idp.AwsCognitoUserPool.SchemaPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetAttributeDataType(val *string) {
	if err := j.validateSetAttributeDataTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeDataType",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetDeveloperOnlyAttribute(val interface{}) {
	if err := j.validateSetDeveloperOnlyAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"developerOnlyAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetMutable(val interface{}) {
	if err := j.validateSetMutableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mutable",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetRequired(val interface{}) {
	if err := j.validateSetRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"required",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) PutNumberAttributeConstraints(value *AwsCognitoUserPool_NumberAttributeConstraintsProperty) {
	if err := a.validatePutNumberAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNumberAttributeConstraints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) PutStringAttributeConstraints(value *AwsCognitoUserPool_StringAttributeConstraintsProperty) {
	if err := a.validatePutStringAttributeConstraintsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStringAttributeConstraints",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ResetDeveloperOnlyAttribute() {
	_jsii_.InvokeVoid(
		a,
		"resetDeveloperOnlyAttribute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ResetMutable() {
	_jsii_.InvokeVoid(
		a,
		"resetMutable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ResetNumberAttributeConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetNumberAttributeConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ResetRequired() {
	_jsii_.InvokeVoid(
		a,
		"resetRequired",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ResetStringAttributeConstraints() {
	_jsii_.InvokeVoid(
		a,
		"resetStringAttributeConstraints",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCognitoUserPool_SchemaPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

