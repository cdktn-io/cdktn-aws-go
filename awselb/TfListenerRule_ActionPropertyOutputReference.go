package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfListenerRule_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticateCognito() TfListenerRule_AuthenticateCognitoPropertyOutputReference
	// Experimental.
	AuthenticateCognitoInput() *TfListenerRule_AuthenticateCognitoProperty
	// Experimental.
	AuthenticateOidc() TfListenerRule_AuthenticateOidcPropertyOutputReference
	// Experimental.
	AuthenticateOidcInput() *TfListenerRule_AuthenticateOidcProperty
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
	FixedResponse() TfListenerRule_FixedResponsePropertyOutputReference
	// Experimental.
	FixedResponseInput() *TfListenerRule_FixedResponseProperty
	// Experimental.
	Forward() TfListenerRule_ForwardPropertyOutputReference
	// Experimental.
	ForwardInput() *TfListenerRule_ForwardProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtValidation() TfListenerRule_JwtValidationPropertyOutputReference
	// Experimental.
	JwtValidationInput() *TfListenerRule_JwtValidationProperty
	// Experimental.
	Order() *float64
	// Experimental.
	SetOrder(val *float64)
	// Experimental.
	OrderInput() *float64
	// Experimental.
	Redirect() TfListenerRule_RedirectPropertyOutputReference
	// Experimental.
	RedirectInput() *TfListenerRule_RedirectProperty
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
	PutAuthenticateCognito(value *TfListenerRule_AuthenticateCognitoProperty)
	// Experimental.
	PutAuthenticateOidc(value *TfListenerRule_AuthenticateOidcProperty)
	// Experimental.
	PutFixedResponse(value *TfListenerRule_FixedResponseProperty)
	// Experimental.
	PutForward(value *TfListenerRule_ForwardProperty)
	// Experimental.
	PutJwtValidation(value *TfListenerRule_JwtValidationProperty)
	// Experimental.
	PutRedirect(value *TfListenerRule_RedirectProperty)
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

// The jsii proxy struct for TfListenerRule_ActionPropertyOutputReference
type jsiiProxy_TfListenerRule_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) AuthenticateCognito() TfListenerRule_AuthenticateCognitoPropertyOutputReference {
	var returns TfListenerRule_AuthenticateCognitoPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) AuthenticateCognitoInput() *TfListenerRule_AuthenticateCognitoProperty {
	var returns *TfListenerRule_AuthenticateCognitoProperty
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) AuthenticateOidc() TfListenerRule_AuthenticateOidcPropertyOutputReference {
	var returns TfListenerRule_AuthenticateOidcPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) AuthenticateOidcInput() *TfListenerRule_AuthenticateOidcProperty {
	var returns *TfListenerRule_AuthenticateOidcProperty
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) FixedResponse() TfListenerRule_FixedResponsePropertyOutputReference {
	var returns TfListenerRule_FixedResponsePropertyOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) FixedResponseInput() *TfListenerRule_FixedResponseProperty {
	var returns *TfListenerRule_FixedResponseProperty
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Forward() TfListenerRule_ForwardPropertyOutputReference {
	var returns TfListenerRule_ForwardPropertyOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ForwardInput() *TfListenerRule_ForwardProperty {
	var returns *TfListenerRule_ForwardProperty
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) JwtValidation() TfListenerRule_JwtValidationPropertyOutputReference {
	var returns TfListenerRule_JwtValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) JwtValidationInput() *TfListenerRule_JwtValidationProperty {
	var returns *TfListenerRule_JwtValidationProperty
	_jsii_.Get(
		j,
		"jwtValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Redirect() TfListenerRule_RedirectPropertyOutputReference {
	var returns TfListenerRule_RedirectPropertyOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) RedirectInput() *TfListenerRule_RedirectProperty {
	var returns *TfListenerRule_RedirectProperty
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfListenerRule_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfListenerRule_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfListenerRule_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfListenerRule_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfListenerRule_ActionPropertyOutputReference_Override(t TfListenerRule_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfListenerRule_ActionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutAuthenticateCognito(value *TfListenerRule_AuthenticateCognitoProperty) {
	if err := t.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutAuthenticateOidc(value *TfListenerRule_AuthenticateOidcProperty) {
	if err := t.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutFixedResponse(value *TfListenerRule_FixedResponseProperty) {
	if err := t.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutForward(value *TfListenerRule_ForwardProperty) {
	if err := t.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForward",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutJwtValidation(value *TfListenerRule_JwtValidationProperty) {
	if err := t.validatePutJwtValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJwtValidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) PutRedirect(value *TfListenerRule_RedirectProperty) {
	if err := t.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedirect",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		t,
		"resetForward",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetJwtValidation() {
	_jsii_.InvokeVoid(
		t,
		"resetJwtValidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		t,
		"resetRedirect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfListenerRule_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

