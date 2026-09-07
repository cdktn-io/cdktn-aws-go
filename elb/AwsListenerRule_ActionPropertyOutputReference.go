package elb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsListenerRule_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticateCognito() AwsListenerRule_AuthenticateCognitoPropertyOutputReference
	// Experimental.
	AuthenticateCognitoInput() *AwsListenerRule_AuthenticateCognitoProperty
	// Experimental.
	AuthenticateOidc() AwsListenerRule_AuthenticateOidcPropertyOutputReference
	// Experimental.
	AuthenticateOidcInput() *AwsListenerRule_AuthenticateOidcProperty
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
	FixedResponse() AwsListenerRule_FixedResponsePropertyOutputReference
	// Experimental.
	FixedResponseInput() *AwsListenerRule_FixedResponseProperty
	// Experimental.
	Forward() AwsListenerRule_ForwardPropertyOutputReference
	// Experimental.
	ForwardInput() *AwsListenerRule_ForwardProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtValidation() AwsListenerRule_JwtValidationPropertyOutputReference
	// Experimental.
	JwtValidationInput() *AwsListenerRule_JwtValidationProperty
	// Experimental.
	Order() *float64
	// Experimental.
	SetOrder(val *float64)
	// Experimental.
	OrderInput() *float64
	// Experimental.
	Redirect() AwsListenerRule_RedirectPropertyOutputReference
	// Experimental.
	RedirectInput() *AwsListenerRule_RedirectProperty
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
	PutAuthenticateCognito(value *AwsListenerRule_AuthenticateCognitoProperty)
	// Experimental.
	PutAuthenticateOidc(value *AwsListenerRule_AuthenticateOidcProperty)
	// Experimental.
	PutFixedResponse(value *AwsListenerRule_FixedResponseProperty)
	// Experimental.
	PutForward(value *AwsListenerRule_ForwardProperty)
	// Experimental.
	PutJwtValidation(value *AwsListenerRule_JwtValidationProperty)
	// Experimental.
	PutRedirect(value *AwsListenerRule_RedirectProperty)
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

// The jsii proxy struct for AwsListenerRule_ActionPropertyOutputReference
type jsiiProxy_AwsListenerRule_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) AuthenticateCognito() AwsListenerRule_AuthenticateCognitoPropertyOutputReference {
	var returns AwsListenerRule_AuthenticateCognitoPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) AuthenticateCognitoInput() *AwsListenerRule_AuthenticateCognitoProperty {
	var returns *AwsListenerRule_AuthenticateCognitoProperty
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) AuthenticateOidc() AwsListenerRule_AuthenticateOidcPropertyOutputReference {
	var returns AwsListenerRule_AuthenticateOidcPropertyOutputReference
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) AuthenticateOidcInput() *AwsListenerRule_AuthenticateOidcProperty {
	var returns *AwsListenerRule_AuthenticateOidcProperty
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) FixedResponse() AwsListenerRule_FixedResponsePropertyOutputReference {
	var returns AwsListenerRule_FixedResponsePropertyOutputReference
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) FixedResponseInput() *AwsListenerRule_FixedResponseProperty {
	var returns *AwsListenerRule_FixedResponseProperty
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Forward() AwsListenerRule_ForwardPropertyOutputReference {
	var returns AwsListenerRule_ForwardPropertyOutputReference
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ForwardInput() *AwsListenerRule_ForwardProperty {
	var returns *AwsListenerRule_ForwardProperty
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) JwtValidation() AwsListenerRule_JwtValidationPropertyOutputReference {
	var returns AwsListenerRule_JwtValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"jwtValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) JwtValidationInput() *AwsListenerRule_JwtValidationProperty {
	var returns *AwsListenerRule_JwtValidationProperty
	_jsii_.Get(
		j,
		"jwtValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Redirect() AwsListenerRule_RedirectPropertyOutputReference {
	var returns AwsListenerRule_RedirectPropertyOutputReference
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) RedirectInput() *AwsListenerRule_RedirectProperty {
	var returns *AwsListenerRule_RedirectProperty
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) TargetGroupArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) TargetGroupArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"targetGroupArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsListenerRule_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsListenerRule_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsListenerRule_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsListenerRule_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.AwsListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsListenerRule_ActionPropertyOutputReference_Override(a AwsListenerRule_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.AwsListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetTargetGroupArn(val *string) {
	if err := j.validateSetTargetGroupArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetGroupArn",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutAuthenticateCognito(value *AwsListenerRule_AuthenticateCognitoProperty) {
	if err := a.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutAuthenticateOidc(value *AwsListenerRule_AuthenticateOidcProperty) {
	if err := a.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutFixedResponse(value *AwsListenerRule_FixedResponseProperty) {
	if err := a.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutForward(value *AwsListenerRule_ForwardProperty) {
	if err := a.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putForward",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutJwtValidation(value *AwsListenerRule_JwtValidationProperty) {
	if err := a.validatePutJwtValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJwtValidation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) PutRedirect(value *AwsListenerRule_RedirectProperty) {
	if err := a.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRedirect",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		a,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		a,
		"resetForward",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetJwtValidation() {
	_jsii_.InvokeVoid(
		a,
		"resetJwtValidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		a,
		"resetRedirect",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ResetTargetGroupArn() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetGroupArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsListenerRule_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

