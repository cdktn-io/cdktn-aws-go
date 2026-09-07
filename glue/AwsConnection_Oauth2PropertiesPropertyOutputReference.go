package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnection_Oauth2PropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationCodeProperties() AwsConnection_AuthorizationCodePropertiesPropertyOutputReference
	// Experimental.
	AuthorizationCodePropertiesInput() *AwsConnection_AuthorizationCodePropertiesProperty
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
	InternalValue() *AwsConnection_Oauth2PropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsConnection_Oauth2PropertiesProperty)
	// Experimental.
	Oauth2ClientApplication() AwsConnection_Oauth2ClientApplicationPropertyOutputReference
	// Experimental.
	Oauth2ClientApplicationInput() *AwsConnection_Oauth2ClientApplicationProperty
	// Experimental.
	Oauth2Credentials() AwsConnection_Oauth2CredentialsPropertyOutputReference
	// Experimental.
	Oauth2CredentialsInput() *AwsConnection_Oauth2CredentialsProperty
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
	PutAuthorizationCodeProperties(value *AwsConnection_AuthorizationCodePropertiesProperty)
	// Experimental.
	PutOauth2ClientApplication(value *AwsConnection_Oauth2ClientApplicationProperty)
	// Experimental.
	PutOauth2Credentials(value *AwsConnection_Oauth2CredentialsProperty)
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

// The jsii proxy struct for AwsConnection_Oauth2PropertiesPropertyOutputReference
type jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) AuthorizationCodeProperties() AwsConnection_AuthorizationCodePropertiesPropertyOutputReference {
	var returns AwsConnection_AuthorizationCodePropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"authorizationCodeProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) AuthorizationCodePropertiesInput() *AwsConnection_AuthorizationCodePropertiesProperty {
	var returns *AwsConnection_AuthorizationCodePropertiesProperty
	_jsii_.Get(
		j,
		"authorizationCodePropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) InternalValue() *AwsConnection_Oauth2PropertiesProperty {
	var returns *AwsConnection_Oauth2PropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2ClientApplication() AwsConnection_Oauth2ClientApplicationPropertyOutputReference {
	var returns AwsConnection_Oauth2ClientApplicationPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2ClientApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2ClientApplicationInput() *AwsConnection_Oauth2ClientApplicationProperty {
	var returns *AwsConnection_Oauth2ClientApplicationProperty
	_jsii_.Get(
		j,
		"oauth2ClientApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2Credentials() AwsConnection_Oauth2CredentialsPropertyOutputReference {
	var returns AwsConnection_Oauth2CredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2Credentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2CredentialsInput() *AwsConnection_Oauth2CredentialsProperty {
	var returns *AwsConnection_Oauth2CredentialsProperty
	_jsii_.Get(
		j,
		"oauth2CredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2GrantType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Oauth2GrantTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oauth2GrantTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TokenUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlParametersMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) TokenUrlParametersMapInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tokenUrlParametersMapInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnection_Oauth2PropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnection_Oauth2PropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnection_Oauth2PropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsConnection.Oauth2PropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnection_Oauth2PropertiesPropertyOutputReference_Override(a AwsConnection_Oauth2PropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsConnection.Oauth2PropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetInternalValue(val *AwsConnection_Oauth2PropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetOauth2GrantType(val *string) {
	if err := j.validateSetOauth2GrantTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oauth2GrantType",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetTokenUrl(val *string) {
	if err := j.validateSetTokenUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrl",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference)SetTokenUrlParametersMap(val *map[string]*string) {
	if err := j.validateSetTokenUrlParametersMapParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenUrlParametersMap",
		val,
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) PutAuthorizationCodeProperties(value *AwsConnection_AuthorizationCodePropertiesProperty) {
	if err := a.validatePutAuthorizationCodePropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthorizationCodeProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) PutOauth2ClientApplication(value *AwsConnection_Oauth2ClientApplicationProperty) {
	if err := a.validatePutOauth2ClientApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2ClientApplication",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) PutOauth2Credentials(value *AwsConnection_Oauth2CredentialsProperty) {
	if err := a.validatePutOauth2CredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2Credentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetAuthorizationCodeProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthorizationCodeProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2ClientApplication() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2ClientApplication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2Credentials() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2Credentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetOauth2GrantType() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2GrantType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetTokenUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ResetTokenUrlParametersMap() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenUrlParametersMap",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnection_Oauth2PropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

