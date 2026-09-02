package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappflow/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappflow/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() TfConnectorProfile_ApiKeyPropertyOutputReference
	// Experimental.
	ApiKeyInput() *TfConnectorProfile_ApiKeyProperty
	// Experimental.
	AuthenticationType() *string
	// Experimental.
	SetAuthenticationType(val *string)
	// Experimental.
	AuthenticationTypeInput() *string
	// Experimental.
	Basic() TfConnectorProfile_BasicPropertyOutputReference
	// Experimental.
	BasicInput() *TfConnectorProfile_BasicProperty
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
	Custom() TfConnectorProfile_CustomPropertyOutputReference
	// Experimental.
	CustomInput() *TfConnectorProfile_CustomProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	// Experimental.
	SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty)
	// Experimental.
	Oauth2() TfConnectorProfile_Oauth2PropertyOutputReference
	// Experimental.
	Oauth2Input() *TfConnectorProfile_Oauth2Property
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
	PutApiKey(value *TfConnectorProfile_ApiKeyProperty)
	// Experimental.
	PutBasic(value *TfConnectorProfile_BasicProperty)
	// Experimental.
	PutCustom(value *TfConnectorProfile_CustomProperty)
	// Experimental.
	PutOauth2(value *TfConnectorProfile_Oauth2Property)
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

// The jsii proxy struct for TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference
type jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ApiKey() TfConnectorProfile_ApiKeyPropertyOutputReference {
	var returns TfConnectorProfile_ApiKeyPropertyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ApiKeyInput() *TfConnectorProfile_ApiKeyProperty {
	var returns *TfConnectorProfile_ApiKeyProperty
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Basic() TfConnectorProfile_BasicPropertyOutputReference {
	var returns TfConnectorProfile_BasicPropertyOutputReference
	_jsii_.Get(
		j,
		"basic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) BasicInput() *TfConnectorProfile_BasicProperty {
	var returns *TfConnectorProfile_BasicProperty
	_jsii_.Get(
		j,
		"basicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Custom() TfConnectorProfile_CustomPropertyOutputReference {
	var returns TfConnectorProfile_CustomPropertyOutputReference
	_jsii_.Get(
		j,
		"custom",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) CustomInput() *TfConnectorProfile_CustomProperty {
	var returns *TfConnectorProfile_CustomProperty
	_jsii_.Get(
		j,
		"customInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InternalValue() *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty {
	var returns *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Oauth2() TfConnectorProfile_Oauth2PropertyOutputReference {
	var returns TfConnectorProfile_Oauth2PropertyOutputReference
	_jsii_.Get(
		j,
		"oauth2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Oauth2Input() *TfConnectorProfile_Oauth2Property {
	var returns *TfConnectorProfile_Oauth2Property
	_jsii_.Get(
		j,
		"oauth2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference_Override(t TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appflow.TfConnectorProfile.ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetInternalValue(val *TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutApiKey(value *TfConnectorProfile_ApiKeyProperty) {
	if err := t.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiKey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutBasic(value *TfConnectorProfile_BasicProperty) {
	if err := t.validatePutBasicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBasic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutCustom(value *TfConnectorProfile_CustomProperty) {
	if err := t.validatePutCustomParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustom",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) PutOauth2(value *TfConnectorProfile_Oauth2Property) {
	if err := t.validatePutOauth2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth2",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		t,
		"resetApiKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetBasic() {
	_jsii_.InvokeVoid(
		t,
		"resetBasic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetCustom() {
	_jsii_.InvokeVoid(
		t,
		"resetCustom",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ResetOauth2() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnectorProfile_ConnectorProfileConfigConnectorProfileCredentialsCustomConnectorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

