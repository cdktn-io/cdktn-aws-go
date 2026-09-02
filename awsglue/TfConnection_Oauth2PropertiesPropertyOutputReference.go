package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnection_Oauth2PropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationCodeProperties() TfConnection_AuthorizationCodePropertiesPropertyOutputReference
	// Experimental.
	AuthorizationCodePropertiesInput() *TfConnection_AuthorizationCodePropertiesProperty
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfConnection_Oauth2PropertiesProperty
	// Experimental.
	SetInternalValue(val *TfConnection_Oauth2PropertiesProperty)
	// Experimental.
	Oauth2ClientApplication() TfConnection_Oauth2ClientApplicationPropertyOutputReference
	// Experimental.
	Oauth2ClientApplicationInput() *TfConnection_Oauth2ClientApplicationProperty
	// Experimental.
	Oauth2Credentials() TfConnection_Oauth2CredentialsPropertyOutputReference
	// Experimental.
	Oauth2CredentialsInput() *TfConnection_Oauth2CredentialsProperty
	// Experimental.
	Oauth2GrantType() *string
	// Experimental.
	SetOauth2GrantType(val *string)
	// Experimental.
	Oauth2GrantTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TokenUrl() *string
	// Experimental.
	SetTokenUrl(val *string)
	// Experimental.
	TokenUrlInput() *string
	// Experimental.
	TokenUrlParametersMap() *map[string]*string
	// Experimental.
	SetTokenUrlParametersMap(val *map[string]*string)
	// Experimental.
	TokenUrlParametersMapInput() *map[string]*string
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
	PutAuthorizationCodeProperties(value *TfConnection_AuthorizationCodePropertiesProperty)
	// Experimental.
	PutOauth2ClientApplication(value *TfConnection_Oauth2ClientApplicationProperty)
	// Experimental.
	PutOauth2Credentials(value *TfConnection_Oauth2CredentialsProperty)
	// Experimental.
	ResetAuthorizationCodeProperties()
	// Experimental.
	ResetOauth2ClientApplication()
	// Experimental.
	ResetOauth2Credentials()
	// Experimental.
	ResetOauth2GrantType()
	// Experimental.
	ResetTokenUrl()
	// Experimental.
	ResetTokenUrlParametersMap()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnection_Oauth2PropertiesPropertyOutputReference
type jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) AuthorizationCodeProperties() TfConnection_AuthorizationCodePropertiesPropertyOutputReference {
	var returns TfConnection_AuthorizationCodePropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) AuthorizationCodePropertiesInput() *TfConnection_AuthorizationCodePropertiesProperty {
	var returns *TfConnection_AuthorizationCodePropertiesProperty
	_jsii_.Get(
		j,
		"authorizationCodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) InternalValue() *TfConnection_Oauth2PropertiesProperty {
	var returns *TfConnection_Oauth2PropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2ClientApplication() TfConnection_Oauth2ClientApplicationPropertyOutputReference {
	var returns TfConnection_Oauth2ClientApplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2ClientApplicationInput() *TfConnection_Oauth2ClientApplicationProperty {
	var returns *TfConnection_Oauth2ClientApplicationProperty
	_jsii_.Get(
		j,
		"oauth2ClientApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2Credentials() TfConnection_Oauth2CredentialsPropertyOutputReference {
	var returns TfConnection_Oauth2CredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2Credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2CredentialsInput() *TfConnection_Oauth2CredentialsProperty {
	var returns *TfConnection_Oauth2CredentialsProperty
	_jsii_.Get(
		j,
		"oauth2CredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Oauth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlParametersMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlParametersMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMapInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnection_Oauth2PropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnection_Oauth2PropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnection_Oauth2PropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfConnection.Oauth2PropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnection_Oauth2PropertiesPropertyOutputReference_Override(t TfConnection_Oauth2PropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfConnection.Oauth2PropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetInternalValue(val *TfConnection_Oauth2PropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetOauth2GrantType(val *string) {
	if err := j.validateSetOauth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2GrantType",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (j *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference)SetTokenUrlParametersMap(val *map[string]*string) {
	if err := j.validateSetTokenUrlParametersMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrlParametersMap",
		val,
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) PutAuthorizationCodeProperties(value *TfConnection_AuthorizationCodePropertiesProperty) {
	if err := t.validatePutAuthorizationCodePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAuthorizationCodeProperties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) PutOauth2ClientApplication(value *TfConnection_Oauth2ClientApplicationProperty) {
	if err := t.validatePutOauth2ClientApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth2ClientApplication",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) PutOauth2Credentials(value *TfConnection_Oauth2CredentialsProperty) {
	if err := t.validatePutOauth2CredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth2Credentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetAuthorizationCodeProperties() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthorizationCodeProperties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2ClientApplication() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth2ClientApplication",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2Credentials() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth2Credentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2GrantType() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth2GrantType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ResetTokenUrlParametersMap() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenUrlParametersMap",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnection_Oauth2PropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

