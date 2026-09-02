package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnect/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnect/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUser_PhoneConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AfterContactWorkTimeLimit() *float64
	// Experimental.
	SetAfterContactWorkTimeLimit(val *float64)
	// Experimental.
	AfterContactWorkTimeLimitInput() *float64
	// Experimental.
	AutoAccept() interface{}
	// Experimental.
	SetAutoAccept(val interface{})
	// Experimental.
	AutoAcceptInput() interface{}
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
	DeskPhoneNumber() *string
	// Experimental.
	SetDeskPhoneNumber(val *string)
	// Experimental.
	DeskPhoneNumberInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfUser_PhoneConfigProperty
	// Experimental.
	SetInternalValue(val *TfUser_PhoneConfigProperty)
	// Experimental.
	PhoneType() *string
	// Experimental.
	SetPhoneType(val *string)
	// Experimental.
	PhoneTypeInput() *string
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
	ResetAfterContactWorkTimeLimit()
	// Experimental.
	ResetAutoAccept()
	// Experimental.
	ResetDeskPhoneNumber()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfUser_PhoneConfigPropertyOutputReference
type jsiiProxy_TfUser_PhoneConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) AfterContactWorkTimeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"afterContactWorkTimeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) AfterContactWorkTimeLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"afterContactWorkTimeLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) AutoAccept() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAccept",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) AutoAcceptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoAcceptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) DeskPhoneNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deskPhoneNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) DeskPhoneNumberInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deskPhoneNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) InternalValue() *TfUser_PhoneConfigProperty {
	var returns *TfUser_PhoneConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) PhoneType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) PhoneTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"phoneTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUser_PhoneConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUser_PhoneConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUser_PhoneConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUser_PhoneConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect.TfUser.PhoneConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUser_PhoneConfigPropertyOutputReference_Override(t TfUser_PhoneConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect.TfUser.PhoneConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetAfterContactWorkTimeLimit(val *float64) {
	if err := j.validateSetAfterContactWorkTimeLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"afterContactWorkTimeLimit",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetAutoAccept(val interface{}) {
	if err := j.validateSetAutoAcceptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoAccept",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetDeskPhoneNumber(val *string) {
	if err := j.validateSetDeskPhoneNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deskPhoneNumber",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetInternalValue(val *TfUser_PhoneConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetPhoneType(val *string) {
	if err := j.validateSetPhoneTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"phoneType",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ResetAfterContactWorkTimeLimit() {
	_jsii_.InvokeVoid(
		t,
		"resetAfterContactWorkTimeLimit",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ResetAutoAccept() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoAccept",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ResetDeskPhoneNumber() {
	_jsii_.InvokeVoid(
		t,
		"resetDeskPhoneNumber",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUser_PhoneConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

