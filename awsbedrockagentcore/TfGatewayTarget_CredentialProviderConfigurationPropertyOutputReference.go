package awsbedrockagentcore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbedrockagentcore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() TfGatewayTarget_ApiKeyPropertyList
	// Experimental.
	ApiKeyInput() interface{}
	// Experimental.
	CallerIamCredentials() TfGatewayTarget_CallerIamCredentialsPropertyList
	// Experimental.
	CallerIamCredentialsInput() interface{}
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
	GatewayIamRole() TfGatewayTarget_GatewayIamRolePropertyList
	// Experimental.
	GatewayIamRoleInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	JwtPassthrough() TfGatewayTarget_JwtPassthroughPropertyList
	// Experimental.
	JwtPassthroughInput() interface{}
	// Experimental.
	Oauth() TfGatewayTarget_OauthPropertyList
	// Experimental.
	OauthInput() interface{}
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
	PutApiKey(value interface{})
	// Experimental.
	PutCallerIamCredentials(value interface{})
	// Experimental.
	PutGatewayIamRole(value interface{})
	// Experimental.
	PutJwtPassthrough(value interface{})
	// Experimental.
	PutOauth(value interface{})
	// Experimental.
	ResetApiKey()
	// Experimental.
	ResetCallerIamCredentials()
	// Experimental.
	ResetGatewayIamRole()
	// Experimental.
	ResetJwtPassthrough()
	// Experimental.
	ResetOauth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference
type jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ApiKey() TfGatewayTarget_ApiKeyPropertyList {
	var returns TfGatewayTarget_ApiKeyPropertyList
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ApiKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CallerIamCredentials() TfGatewayTarget_CallerIamCredentialsPropertyList {
	var returns TfGatewayTarget_CallerIamCredentialsPropertyList
	_jsii_.Get(
		j,
		"callerIamCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CallerIamCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"callerIamCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GatewayIamRole() TfGatewayTarget_GatewayIamRolePropertyList {
	var returns TfGatewayTarget_GatewayIamRolePropertyList
	_jsii_.Get(
		j,
		"gatewayIamRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GatewayIamRoleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"gatewayIamRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) JwtPassthrough() TfGatewayTarget_JwtPassthroughPropertyList {
	var returns TfGatewayTarget_JwtPassthroughPropertyList
	_jsii_.Get(
		j,
		"jwtPassthrough",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) JwtPassthroughInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"jwtPassthroughInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Oauth() TfGatewayTarget_OauthPropertyList {
	var returns TfGatewayTarget_OauthPropertyList
	_jsii_.Get(
		j,
		"oauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) OauthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayTarget_CredentialProviderConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.CredentialProviderConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference_Override(t TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-bedrock-agentcore.TfGatewayTarget.CredentialProviderConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutApiKey(value interface{}) {
	if err := t.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiKey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutCallerIamCredentials(value interface{}) {
	if err := t.validatePutCallerIamCredentialsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCallerIamCredentials",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutGatewayIamRole(value interface{}) {
	if err := t.validatePutGatewayIamRoleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGatewayIamRole",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutJwtPassthrough(value interface{}) {
	if err := t.validatePutJwtPassthroughParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJwtPassthrough",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) PutOauth(value interface{}) {
	if err := t.validatePutOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		t,
		"resetApiKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetCallerIamCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetCallerIamCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetGatewayIamRole() {
	_jsii_.InvokeVoid(
		t,
		"resetGatewayIamRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetJwtPassthrough() {
	_jsii_.InvokeVoid(
		t,
		"resetJwtPassthrough",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ResetOauth() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayTarget_CredentialProviderConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

