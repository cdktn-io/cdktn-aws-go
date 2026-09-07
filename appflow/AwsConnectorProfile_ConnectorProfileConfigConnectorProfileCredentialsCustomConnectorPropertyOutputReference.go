package appflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() AwsConnectorProfile_ApiKeyPropertyOutputReference
	// Experimental.
	ApiKeyInput() *AwsConnectorProfile_ApiKeyProperty
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	Basic() AwsConnectorProfile_BasicPropertyOutputReference
	// Experimental.
	BasicInput() *AwsConnectorProfile_BasicProperty
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
	Custom() AwsConnectorProfile_CustomPropertyOutputReference
	// Experimental.
	CustomInput() *AwsConnectorProfile_CustomProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	// Experimental.
	SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty)
	// Experimental.
	Oauth2() AwsConnectorProfile_Oauth2PropertyOutputReference
	// Experimental.
	Oauth2Input() *AwsConnectorProfile_Oauth2Property
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
	PutApiKey(value *AwsConnectorProfile_ApiKeyProperty)
	// Experimental.
	PutBasic(value *AwsConnectorProfile_BasicProperty)
	// Experimental.
	PutCustom(value *AwsConnectorProfile_CustomProperty)
	// Experimental.
	PutOauth2(value *AwsConnectorProfile_Oauth2Property)
	// Experimental.
	ResetApiKey()
	// Experimental.
	ResetBasic()
	// Experimental.
	ResetCustom()
	// Experimental.
	ResetOauth2()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
type jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ApiKey() AwsConnectorProfile_ApiKeyPropertyOutputReference {
	var returns AwsConnectorProfile_ApiKeyPropertyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ApiKeyInput() *AwsConnectorProfile_ApiKeyProperty {
	var returns *AwsConnectorProfile_ApiKeyProperty
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Basic() AwsConnectorProfile_BasicPropertyOutputReference {
	var returns AwsConnectorProfile_BasicPropertyOutputReference
	_jsii_.Get(
		j,
		"basic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) BasicInput() *AwsConnectorProfile_BasicProperty {
	var returns *AwsConnectorProfile_BasicProperty
	_jsii_.Get(
		j,
		"basicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Custom() AwsConnectorProfile_CustomPropertyOutputReference {
	var returns AwsConnectorProfile_CustomPropertyOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) CustomInput() *AwsConnectorProfile_CustomProperty {
	var returns *AwsConnectorProfile_CustomProperty
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InternalValue() *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty {
	var returns *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Oauth2() AwsConnectorProfile_Oauth2PropertyOutputReference {
	var returns AwsConnectorProfile_Oauth2PropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Oauth2Input() *AwsConnectorProfile_Oauth2Property {
	var returns *AwsConnectorProfile_Oauth2Property
	_jsii_.Get(
		j,
		"oauth2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference_Override(a AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.AwsConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetInternalValue(val *AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutApiKey(value *AwsConnectorProfile_ApiKeyProperty) {
	if err := a.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutBasic(value *AwsConnectorProfile_BasicProperty) {
	if err := a.validatePutBasicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasic",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutCustom(value *AwsConnectorProfile_CustomProperty) {
	if err := a.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustom",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutOauth2(value *AwsConnectorProfile_Oauth2Property) {
	if err := a.validatePutOauth2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetBasic() {
	_jsii_.InvokeVoid(
		a,
		"resetBasic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		a,
		"resetCustom",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetOauth2() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

