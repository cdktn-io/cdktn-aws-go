package awselb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselb/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselb/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfListenerRule_ActionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticateCognito() DataTfListenerRule_AuthenticateCognitoPropertyList
	// Experimental.
	AuthenticateCognitoInput() interface{}
	// Experimental.
	AuthenticateOidc() DataTfListenerRule_AuthenticateOidcPropertyList
	// Experimental.
	AuthenticateOidcInput() interface{}
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
	FixedResponse() DataTfListenerRule_FixedResponsePropertyList
	// Experimental.
	FixedResponseInput() interface{}
	// Experimental.
	Forward() DataTfListenerRule_ForwardPropertyList
	// Experimental.
	ForwardInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtValidation() DataTfListenerRule_JwtValidationPropertyList
	// Experimental.
	JwtValidationInput() interface{}
	// Experimental.
	Order() *float64
	// Experimental.
	Redirect() DataTfListenerRule_RedirectPropertyList
	// Experimental.
	RedirectInput() interface{}
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
	PutAuthenticateCognito(value interface{})
	// Experimental.
	PutAuthenticateOidc(value interface{})
	// Experimental.
	PutFixedResponse(value interface{})
	// Experimental.
	PutForward(value interface{})
	// Experimental.
	PutJwtValidation(value interface{})
	// Experimental.
	PutRedirect(value interface{})
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
	ResetRedirect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfListenerRule_ActionPropertyOutputReference
type jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) AuthenticateCognito() DataTfListenerRule_AuthenticateCognitoPropertyList {
	var returns DataTfListenerRule_AuthenticateCognitoPropertyList
	_jsii_.Get(
		j,
		"authenticateCognito",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) AuthenticateCognitoInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticateCognitoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) AuthenticateOidc() DataTfListenerRule_AuthenticateOidcPropertyList {
	var returns DataTfListenerRule_AuthenticateOidcPropertyList
	_jsii_.Get(
		j,
		"authenticateOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) AuthenticateOidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticateOidcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) FixedResponse() DataTfListenerRule_FixedResponsePropertyList {
	var returns DataTfListenerRule_FixedResponsePropertyList
	_jsii_.Get(
		j,
		"fixedResponse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) FixedResponseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fixedResponseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Forward() DataTfListenerRule_ForwardPropertyList {
	var returns DataTfListenerRule_ForwardPropertyList
	_jsii_.Get(
		j,
		"forward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"forwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) JwtValidation() DataTfListenerRule_JwtValidationPropertyList {
	var returns DataTfListenerRule_JwtValidationPropertyList
	_jsii_.Get(
		j,
		"jwtValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) JwtValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Redirect() DataTfListenerRule_RedirectPropertyList {
	var returns DataTfListenerRule_RedirectPropertyList
	_jsii_.Get(
		j,
		"redirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) RedirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfListenerRule_ActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfListenerRule_ActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfListenerRule_ActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elb.DataTfListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfListenerRule_ActionPropertyOutputReference_Override(d DataTfListenerRule_ActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elb.DataTfListenerRule.ActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutAuthenticateCognito(value interface{}) {
	if err := d.validatePutAuthenticateCognitoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthenticateCognito",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutAuthenticateOidc(value interface{}) {
	if err := d.validatePutAuthenticateOidcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAuthenticateOidc",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutFixedResponse(value interface{}) {
	if err := d.validatePutFixedResponseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putFixedResponse",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutForward(value interface{}) {
	if err := d.validatePutForwardParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putForward",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutJwtValidation(value interface{}) {
	if err := d.validatePutJwtValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putJwtValidation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) PutRedirect(value interface{}) {
	if err := d.validatePutRedirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedirect",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetAuthenticateCognito() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticateCognito",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetAuthenticateOidc() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthenticateOidc",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetFixedResponse() {
	_jsii_.InvokeVoid(
		d,
		"resetFixedResponse",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetForward() {
	_jsii_.InvokeVoid(
		d,
		"resetForward",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetJwtValidation() {
	_jsii_.InvokeVoid(
		d,
		"resetJwtValidation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ResetRedirect() {
	_jsii_.InvokeVoid(
		d,
		"resetRedirect",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfListenerRule_ActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

