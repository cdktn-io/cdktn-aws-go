package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlbListener_DefaultActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticateCognito() TfAlbListener_AuthenticateCognitoPropertyOutputReference
	// Experimental.
	AuthenticateCognitoInput() *TfAlbListener_AuthenticateCognitoProperty
	// Experimental.
	AuthenticateOidc() TfAlbListener_AuthenticateOidcPropertyOutputReference
	// Experimental.
	AuthenticateOidcInput() *TfAlbListener_AuthenticateOidcProperty
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
	FixedResponse() TfAlbListener_FixedResponsePropertyOutputReference
	// Experimental.
	FixedResponseInput() *TfAlbListener_FixedResponseProperty
	// Experimental.
	Forward() TfAlbListener_ForwardPropertyOutputReference
	// Experimental.
	ForwardInput() *TfAlbListener_ForwardProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtValidation() TfAlbListener_JwtValidationPropertyOutputReference
	// Experimental.
	JwtValidationInput() *TfAlbListener_JwtValidationProperty
	// Experimental.
	Order() *float64
	// Experimental.
	SetOrder(val *float64)
	// Experimental.
	OrderInput() *float64
	// Experimental.
	Redirect() TfAlbListener_RedirectPropertyOutputReference
	// Experimental.
	RedirectInput() *TfAlbListener_RedirectProperty
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
	PutAuthenticateCognito(value *TfAlbListener_AuthenticateCognitoProperty)
	// Experimental.
	PutAuthenticateOidc(value *TfAlbListener_AuthenticateOidcProperty)
	// Experimental.
	PutFixedResponse(value *TfAlbListener_FixedResponseProperty)
	// Experimental.
	PutForward(value *TfAlbListener_ForwardProperty)
	// Experimental.
	PutJwtValidation(value *TfAlbListener_JwtValidationProperty)
	// Experimental.
	PutRedirect(value *TfAlbListener_RedirectProperty)
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

// The jsii proxy struct for TfAlbListener_DefaultActionPropertyOutputReference
type jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) AuthenticateCognito() TfAlbListener_AuthenticateCognitoPropertyOutputReference {
	var returns TfAlbListener_AuthenticateCognitoPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) AuthenticateCognitoInput() *TfAlbListener_AuthenticateCognitoProperty {
	var returns *TfAlbListener_AuthenticateCognitoProperty
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) AuthenticateOidc() TfAlbListener_AuthenticateOidcPropertyOutputReference {
	var returns TfAlbListener_AuthenticateOidcPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) AuthenticateOidcInput() *TfAlbListener_AuthenticateOidcProperty {
	var returns *TfAlbListener_AuthenticateOidcProperty
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) FixedResponse() TfAlbListener_FixedResponsePropertyOutputReference {
	var returns TfAlbListener_FixedResponsePropertyOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) FixedResponseInput() *TfAlbListener_FixedResponseProperty {
	var returns *TfAlbListener_FixedResponseProperty
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Forward() TfAlbListener_ForwardPropertyOutputReference {
	var returns TfAlbListener_ForwardPropertyOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ForwardInput() *TfAlbListener_ForwardProperty {
	var returns *TfAlbListener_ForwardProperty
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) JwtValidation() TfAlbListener_JwtValidationPropertyOutputReference {
	var returns TfAlbListener_JwtValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) JwtValidationInput() *TfAlbListener_JwtValidationProperty {
	var returns *TfAlbListener_JwtValidationProperty
	_jsii_.Get(
		j,
		"jwtValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Redirect() TfAlbListener_RedirectPropertyOutputReference {
	var returns TfAlbListener_RedirectPropertyOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) RedirectInput() *TfAlbListener_RedirectProperty {
	var returns *TfAlbListener_RedirectProperty
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAlbListener_DefaultActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAlbListener_DefaultActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAlbListener_DefaultActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlbListener.DefaultActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAlbListener_DefaultActionPropertyOutputReference_Override(t TfAlbListener_DefaultActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.TfAlbListener.DefaultActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutAuthenticateCognito(value *TfAlbListener_AuthenticateCognitoProperty) {
	if err := t.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutAuthenticateOidc(value *TfAlbListener_AuthenticateOidcProperty) {
	if err := t.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutFixedResponse(value *TfAlbListener_FixedResponseProperty) {
	if err := t.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutForward(value *TfAlbListener_ForwardProperty) {
	if err := t.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putForward",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutJwtValidation(value *TfAlbListener_JwtValidationProperty) {
	if err := t.validatePutJwtValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJwtValidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) PutRedirect(value *TfAlbListener_RedirectProperty) {
	if err := t.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRedirect",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		t,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		t,
		"resetForward",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetJwtValidation() {
	_jsii_.InvokeVoid(
		t,
		"resetJwtValidation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		t,
		"resetRedirect",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAlbListener_DefaultActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

