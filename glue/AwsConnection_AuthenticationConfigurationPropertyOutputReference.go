package glue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/glue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/glue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnection_AuthenticationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	BasicAuthenticationCredentials() AwsConnection_BasicAuthenticationCredentialsPropertyOutputReference
	// Experimental.
	BasicAuthenticationCredentialsInput() *AwsConnection_BasicAuthenticationCredentialsProperty
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
	InternalValue() *AwsConnection_AuthenticationConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsConnection_AuthenticationConfigurationProperty)
	// Experimental.
	KmsKeyArn() *string
	// Experimental.
	SetKmsKeyArn(val *string)
	// Experimental.
	KmsKeyArnInput() *string
	// Experimental.
	Oauth2Properties() AwsConnection_Oauth2PropertiesPropertyOutputReference
	// Experimental.
	Oauth2PropertiesInput() *AwsConnection_Oauth2PropertiesProperty
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
	PutBasicAuthenticationCredentials(value *AwsConnection_BasicAuthenticationCredentialsProperty)
	// Experimental.
	PutOauth2Properties(value *AwsConnection_Oauth2PropertiesProperty)
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

// The jsii proxy struct for AwsConnection_AuthenticationConfigurationPropertyOutputReference
type jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) BasicAuthenticationCredentials() AwsConnection_BasicAuthenticationCredentialsPropertyOutputReference {
	var returns AwsConnection_BasicAuthenticationCredentialsPropertyOutputReference
	_jsii_.Get(
		j,
		"basicAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) BasicAuthenticationCredentialsInput() *AwsConnection_BasicAuthenticationCredentialsProperty {
	var returns *AwsConnection_BasicAuthenticationCredentialsProperty
	_jsii_.Get(
		j,
		"basicAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) CustomAuthenticationCredentials() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) CustomAuthenticationCredentialsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"customAuthenticationCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) InternalValue() *AwsConnection_AuthenticationConfigurationProperty {
	var returns *AwsConnection_AuthenticationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) KmsKeyArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) KmsKeyArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) Oauth2Properties() AwsConnection_Oauth2PropertiesPropertyOutputReference {
	var returns AwsConnection_Oauth2PropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2Properties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) Oauth2PropertiesInput() *AwsConnection_Oauth2PropertiesProperty {
	var returns *AwsConnection_Oauth2PropertiesProperty
	_jsii_.Get(
		j,
		"oauth2PropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) SecretArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) SecretArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnection_AuthenticationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnection_AuthenticationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnection_AuthenticationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.AwsConnection.AuthenticationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnection_AuthenticationConfigurationPropertyOutputReference_Override(a AwsConnection_AuthenticationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.AwsConnection.AuthenticationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetCustomAuthenticationCredentials(val *map[string]*string) {
	if err := j.validateSetCustomAuthenticationCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customAuthenticationCredentials",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetInternalValue(val *AwsConnection_AuthenticationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetKmsKeyArn(val *string) {
	if err := j.validateSetKmsKeyArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKeyArn",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetSecretArn(val *string) {
	if err := j.validateSetSecretArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretArn",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) PutBasicAuthenticationCredentials(value *AwsConnection_BasicAuthenticationCredentialsProperty) {
	if err := a.validatePutBasicAuthenticationCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasicAuthenticationCredentials",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) PutOauth2Properties(value *AwsConnection_Oauth2PropertiesProperty) {
	if err := a.validatePutOauth2PropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2Properties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ResetBasicAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetBasicAuthenticationCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ResetCustomAuthenticationCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomAuthenticationCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ResetKmsKeyArn() {
	_jsii_.InvokeVoid(
		a,
		"resetKmsKeyArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ResetOauth2Properties() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2Properties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ResetSecretArn() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnection_AuthenticationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

