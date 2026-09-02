package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfListener_DefaultActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticateCognito() TfListener_AuthenticateCognitoPropertyOutputReference
	// Experimental.
	AuthenticateCognitoInput() *TfListener_AuthenticateCognitoProperty
	// Experimental.
	AuthenticateOidc() TfListener_AuthenticateOidcPropertyOutputReference
	// Experimental.
	AuthenticateOidcInput() *TfListener_AuthenticateOidcProperty
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
	FixedResponse() TfListener_FixedResponsePropertyOutputReference
	// Experimental.
	FixedResponseInput() *TfListener_FixedResponseProperty
	// Experimental.
	Forward() TfListener_ForwardPropertyOutputReference
	// Experimental.
	ForwardInput() *TfListener_ForwardProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtValidation() TfListener_JwtValidationPropertyOutputReference
	// Experimental.
	JwtValidationInput() *TfListener_JwtValidationProperty
	// Experimental.
	Order() *float64
	// Experimental.
	SetOrder(val *float64)
	// Experimental.
	OrderInput() *float64
	// Experimental.
	Redirect() TfListener_RedirectPropertyOutputReference
	// Experimental.
	RedirectInput() *TfListener_RedirectProperty
	// Experimental.
	TargetGroupArn() *string
	// Experimental.
	SetTargetGroupArn(val *string)
	// Experimental.
	TargetGroupArnInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutAuthenticateCognito(value *TfListener_AuthenticateCognitoProperty)
	// Experimental.
	PutAuthenticateOidc(value *TfListener_AuthenticateOidcProperty)
	// Experimental.
	PutFixedResponse(value *TfListener_FixedResponseProperty)
	// Experimental.
	PutForward(value *TfListener_ForwardProperty)
	// Experimental.
	PutJwtValidation(value *TfListener_JwtValidationProperty)
	// Experimental.
	PutRedirect(value *TfListener_RedirectProperty)
	// Experimental.
	ResetAuthenticateCognito()
	// Experimental.
	ResetAuthenticateOidc()
	// Experimental.
	ResetFixedResponse()
	// Experimental.
	ResetForward()
	// Experimental.
	ResetJwtValidation()
	// Experimental.
	ResetOrder()
	// Experimental.
	ResetRedirect()
	// Experimental.
	ResetTargetGroupArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfListener_DefaultActionPropertyOutputReference
type jsiiProxy_TfListener_DefaultActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) AuthenticateCognito() TfListener_AuthenticateCognitoPropertyOutputReference {
	var returns TfListener_AuthenticateCognitoPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) AuthenticateCognitoInput() *TfListener_AuthenticateCognitoProperty {
	var returns *TfListener_AuthenticateCognitoProperty
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) AuthenticateOidc() TfListener_AuthenticateOidcPropertyOutputReference {
	var returns TfListener_AuthenticateOidcPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) AuthenticateOidcInput() *TfListener_AuthenticateOidcProperty {
	var returns *TfListener_AuthenticateOidcProperty
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) FixedResponse() TfListener_FixedResponsePropertyOutputReference {
	var returns TfListener_FixedResponsePropertyOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) FixedResponseInput() *TfListener_FixedResponseProperty {
	var returns *TfListener_FixedResponseProperty
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Forward() TfListener_ForwardPropertyOutputReference {
	var returns TfListener_ForwardPropertyOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ForwardInput() *TfListener_ForwardProperty {
	var returns *TfListener_ForwardProperty
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) JwtValidation() TfListener_JwtValidationPropertyOutputReference {
	var returns TfListener_JwtValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) JwtValidationInput() *TfListener_JwtValidationProperty {
	var returns *TfListener_JwtValidationProperty
	_jsii_.Get(
		j,
		"jwtValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Redirect() TfListener_RedirectPropertyOutputReference {
	var returns TfListener_RedirectPropertyOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) RedirectInput() *TfListener_RedirectProperty {
	var returns *TfListener_RedirectProperty
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfListener_DefaultActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfListener_DefaultActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfListener_DefaultActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfListener_DefaultActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfListener.DefaultActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfListener_DefaultActionPropertyOutputReference_Override(t TfListener_DefaultActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfListener.DefaultActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfListener_DefaultActionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutAuthenticateCognito(value *TfListener_AuthenticateCognitoProperty) {
	if err := t.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutAuthenticateOidc(value *TfListener_AuthenticateOidcProperty) {
	if err := t.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutFixedResponse(value *TfListener_FixedResponseProperty) {
	if err := t.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutForward(value *TfListener_ForwardProperty) {
	if err := t.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForward",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutJwtValidation(value *TfListener_JwtValidationProperty) {
	if err := t.validatePutJwtValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJwtValidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) PutRedirect(value *TfListener_RedirectProperty) {
	if err := t.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedirect",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		t,
		"resetForward",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetJwtValidation() {
	_jsii_.InvokeVoid(
		t,
		"resetJwtValidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		t,
		"resetRedirect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfListener_DefaultActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

