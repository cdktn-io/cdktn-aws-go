package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_AttributeTypesSelectorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Address() *[]*string
	// Experimental.
	SetAddress(val *[]*string)
	// Experimental.
	AddressInput() *[]*string
	// Experimental.
	AttributeMatchingModel() *string
	// Experimental.
	SetAttributeMatchingModel(val *string)
	// Experimental.
	AttributeMatchingModelInput() *string
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
	EmailAddress() *[]*string
	// Experimental.
	SetEmailAddress(val *[]*string)
	// Experimental.
	EmailAddressInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_AttributeTypesSelectorProperty
	// Experimental.
	SetInternalValue(val *TfDomain_AttributeTypesSelectorProperty)
	// Experimental.
	PhoneNumber() *[]*string
	// Experimental.
	SetPhoneNumber(val *[]*string)
	// Experimental.
	PhoneNumberInput() *[]*string
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
	ResetAddress()
	// Experimental.
	ResetEmailAddress()
	// Experimental.
	ResetPhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDomain_AttributeTypesSelectorPropertyOutputReference
type jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) Address() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) AddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) AttributeMatchingModel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeMatchingModel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) AttributeMatchingModelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"attributeMatchingModelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) EmailAddress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) EmailAddressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"emailAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) InternalValue() *TfDomain_AttributeTypesSelectorProperty {
	var returns *TfDomain_AttributeTypesSelectorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) PhoneNumber() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) PhoneNumberInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"phoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_AttributeTypesSelectorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_AttributeTypesSelectorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_AttributeTypesSelectorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.AttributeTypesSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_AttributeTypesSelectorPropertyOutputReference_Override(t TfDomain_AttributeTypesSelectorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.TfDomain.AttributeTypesSelectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetAddress(val *[]*string) {
	if err := j.validateSetAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetAttributeMatchingModel(val *string) {
	if err := j.validateSetAttributeMatchingModelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"attributeMatchingModel",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetEmailAddress(val *[]*string) {
	if err := j.validateSetEmailAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emailAddress",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetInternalValue(val *TfDomain_AttributeTypesSelectorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetPhoneNumber(val *[]*string) {
	if err := j.validateSetPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ResetAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ResetEmailAddress() {
	_jsii_.InvokeVoid(
		t,
		"resetEmailAddress",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ResetPhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetPhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_AttributeTypesSelectorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

