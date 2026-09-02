package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnection_AuthenticationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	BasicAuthenticationCredentials() TfConnection_BasicAuthenticationCredentialsPropertyOutputReference
	// Experimental.
	BasicAuthenticationCredentialsInput() *TfConnection_BasicAuthenticationCredentialsProperty
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
	CustomAuthenticationCredentials() *map[string]*string
	// Experimental.
	SetCustomAuthenticationCredentials(val *map[string]*string)
	// Experimental.
	CustomAuthenticationCredentialsInput() *map[string]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConnection_AuthenticationConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfConnection_AuthenticationConfigurationProperty)
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	Oauth2Properties() TfConnection_Oauth2PropertiesPropertyOutputReference
	// Experimental.
	Oauth2PropertiesInput() *TfConnection_Oauth2PropertiesProperty
	// Experimental.
	SecretArn() *string
	// Experimental.
	SetSecretArn(val *string)
	// Experimental.
	SecretArnInput() *string
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
	PutBasicAuthenticationCredentials(value *TfConnection_BasicAuthenticationCredentialsProperty)
	// Experimental.
	PutOauth2Properties(value *TfConnection_Oauth2PropertiesProperty)
	// Experimental.
	ResetBasicAuthenticationCredentials()
	// Experimental.
	ResetCustomAuthenticationCredentials()
	// Experimental.
	ResetKmsKeyArn()
	// Experimental.
	ResetOauth2Properties()
	// Experimental.
	ResetSecretArn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnection_AuthenticationConfigurationPropertyOutputReference
type jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) BasicAuthenticationCredentials() TfConnection_BasicAuthenticationCredentialsPropertyOutputReference {
	var returns TfConnection_BasicAuthenticationCredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"basicAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) BasicAuthenticationCredentialsInput() *TfConnection_BasicAuthenticationCredentialsProperty {
	var returns *TfConnection_BasicAuthenticationCredentialsProperty
	_jsii_.Get(
		j,
		"basicAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) CustomAuthenticationCredentials() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) CustomAuthenticationCredentialsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) InternalValue() *TfConnection_AuthenticationConfigurationProperty {
	var returns *TfConnection_AuthenticationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) Oauth2Properties() TfConnection_Oauth2PropertiesPropertyOutputReference {
	var returns TfConnection_Oauth2PropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) Oauth2PropertiesInput() *TfConnection_Oauth2PropertiesProperty {
	var returns *TfConnection_Oauth2PropertiesProperty
	_jsii_.Get(
		j,
		"oauth2PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnection_AuthenticationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnection_AuthenticationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnection_AuthenticationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfConnection.AuthenticationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnection_AuthenticationConfigurationPropertyOutputReference_Override(t TfConnection_AuthenticationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfConnection.AuthenticationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetCustomAuthenticationCredentials(val *map[string]*string) {
	if err := j.validateSetCustomAuthenticationCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAuthenticationCredentials",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetInternalValue(val *TfConnection_AuthenticationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) PutBasicAuthenticationCredentials(value *TfConnection_BasicAuthenticationCredentialsProperty) {
	if err := t.validatePutBasicAuthenticationCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBasicAuthenticationCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) PutOauth2Properties(value *TfConnection_Oauth2PropertiesProperty) {
	if err := t.validatePutOauth2PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth2Properties",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ResetBasicAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetBasicAuthenticationCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ResetCustomAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomAuthenticationCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		t,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ResetOauth2Properties() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth2Properties",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnection_AuthenticationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

