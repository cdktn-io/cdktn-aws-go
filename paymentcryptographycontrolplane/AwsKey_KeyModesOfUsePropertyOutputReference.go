package paymentcryptographycontrolplane

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/paymentcryptographycontrolplane/jsii"

	"github.com/cdktn-io/cdktn-aws-go/paymentcryptographycontrolplane/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKey_KeyModesOfUsePropertyOutputReference interface {
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
	Decrypt() interface{}
	// Experimental.
	SetDecrypt(val interface{})
	// Experimental.
	DecryptInput() interface{}
	// Experimental.
	DeriveKey() interface{}
	// Experimental.
	SetDeriveKey(val interface{})
	// Experimental.
	DeriveKeyInput() interface{}
	// Experimental.
	Encrypt() interface{}
	// Experimental.
	SetEncrypt(val interface{})
	// Experimental.
	EncryptInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Generate() interface{}
	// Experimental.
	SetGenerate(val interface{})
	// Experimental.
	GenerateInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	NoRestrictions() interface{}
	// Experimental.
	SetNoRestrictions(val interface{})
	// Experimental.
	NoRestrictionsInput() interface{}
	// Experimental.
	Sign() interface{}
	// Experimental.
	SetSign(val interface{})
	// Experimental.
	SignInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Unwrap() interface{}
	// Experimental.
	SetUnwrap(val interface{})
	// Experimental.
	UnwrapInput() interface{}
	// Experimental.
	Verify() interface{}
	// Experimental.
	SetVerify(val interface{})
	// Experimental.
	VerifyInput() interface{}
	// Experimental.
	Wrap() interface{}
	// Experimental.
	SetWrap(val interface{})
	// Experimental.
	WrapInput() interface{}
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
	ResetDecrypt()
	// Experimental.
	ResetDeriveKey()
	// Experimental.
	ResetEncrypt()
	// Experimental.
	ResetGenerate()
	// Experimental.
	ResetNoRestrictions()
	// Experimental.
	ResetSign()
	// Experimental.
	ResetUnwrap()
	// Experimental.
	ResetVerify()
	// Experimental.
	ResetWrap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKey_KeyModesOfUsePropertyOutputReference
type jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Decrypt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decrypt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) DecryptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"decryptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) DeriveKey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deriveKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) DeriveKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deriveKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Encrypt() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) EncryptInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Generate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GenerateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"generateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) NoRestrictions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) NoRestrictionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Sign() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sign",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) SignInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"signInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Unwrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unwrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) UnwrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unwrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Wrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) WrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wrapInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKey_KeyModesOfUsePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsKey_KeyModesOfUsePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKey_KeyModesOfUsePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-payment-cryptography-control-plane.AwsKey.KeyModesOfUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKey_KeyModesOfUsePropertyOutputReference_Override(a AwsKey_KeyModesOfUsePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-payment-cryptography-control-plane.AwsKey.KeyModesOfUsePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetDecrypt(val interface{}) {
	if err := j.validateSetDecryptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"decrypt",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetDeriveKey(val interface{}) {
	if err := j.validateSetDeriveKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deriveKey",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetEncrypt(val interface{}) {
	if err := j.validateSetEncryptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypt",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetGenerate(val interface{}) {
	if err := j.validateSetGenerateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generate",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetNoRestrictions(val interface{}) {
	if err := j.validateSetNoRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noRestrictions",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetSign(val interface{}) {
	if err := j.validateSetSignParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sign",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetUnwrap(val interface{}) {
	if err := j.validateSetUnwrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unwrap",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetVerify(val interface{}) {
	if err := j.validateSetVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (j *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference)SetWrap(val interface{}) {
	if err := j.validateSetWrapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrap",
		val,
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetDecrypt() {
	_jsii_.InvokeVoid(
		a,
		"resetDecrypt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetDeriveKey() {
	_jsii_.InvokeVoid(
		a,
		"resetDeriveKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetEncrypt() {
	_jsii_.InvokeVoid(
		a,
		"resetEncrypt",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetGenerate() {
	_jsii_.InvokeVoid(
		a,
		"resetGenerate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetNoRestrictions() {
	_jsii_.InvokeVoid(
		a,
		"resetNoRestrictions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetSign() {
	_jsii_.InvokeVoid(
		a,
		"resetSign",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetUnwrap() {
	_jsii_.InvokeVoid(
		a,
		"resetUnwrap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetVerify() {
	_jsii_.InvokeVoid(
		a,
		"resetVerify",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ResetWrap() {
	_jsii_.InvokeVoid(
		a,
		"resetWrap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKey_KeyModesOfUsePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

