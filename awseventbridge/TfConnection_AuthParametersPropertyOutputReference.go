package awseventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridge/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnection_AuthParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() TfConnection_ApiKeyPropertyOutputReference
	// Experimental.
	ApiKeyInput() *TfConnection_ApiKeyProperty
	// Experimental.
	Basic() TfConnection_BasicPropertyOutputReference
	// Experimental.
	BasicInput() *TfConnection_BasicProperty
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
	// Experimental.
	ConnectivityParameters() TfConnection_ConnectivityParametersPropertyOutputReference
	// Experimental.
	ConnectivityParametersInput() *TfConnection_ConnectivityParametersProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConnection_AuthParametersProperty
	// Experimental.
	SetInternalValue(val *TfConnection_AuthParametersProperty)
	// Experimental.
	InvocationHttpParameters() TfConnection_InvocationHttpParametersPropertyOutputReference
	// Experimental.
	InvocationHttpParametersInput() *TfConnection_InvocationHttpParametersProperty
	// Experimental.
	Oauth() TfConnection_OauthPropertyOutputReference
	// Experimental.
	OauthInput() *TfConnection_OauthProperty
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
	PutApiKey(value *TfConnection_ApiKeyProperty)
	// Experimental.
	PutBasic(value *TfConnection_BasicProperty)
	// Experimental.
	PutConnectivityParameters(value *TfConnection_ConnectivityParametersProperty)
	// Experimental.
	PutInvocationHttpParameters(value *TfConnection_InvocationHttpParametersProperty)
	// Experimental.
	PutOauth(value *TfConnection_OauthProperty)
	// Experimental.
	ResetApiKey()
	// Experimental.
	ResetBasic()
	// Experimental.
	ResetConnectivityParameters()
	// Experimental.
	ResetInvocationHttpParameters()
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

// The jsii proxy struct for TfConnection_AuthParametersPropertyOutputReference
type jsiiProxy_TfConnection_AuthParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ApiKey() TfConnection_ApiKeyPropertyOutputReference {
	var returns TfConnection_ApiKeyPropertyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ApiKeyInput() *TfConnection_ApiKeyProperty {
	var returns *TfConnection_ApiKeyProperty
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) Basic() TfConnection_BasicPropertyOutputReference {
	var returns TfConnection_BasicPropertyOutputReference
	_jsii_.Get(
		j,
		"basic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) BasicInput() *TfConnection_BasicProperty {
	var returns *TfConnection_BasicProperty
	_jsii_.Get(
		j,
		"basicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ConnectivityParameters() TfConnection_ConnectivityParametersPropertyOutputReference {
	var returns TfConnection_ConnectivityParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"connectivityParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ConnectivityParametersInput() *TfConnection_ConnectivityParametersProperty {
	var returns *TfConnection_ConnectivityParametersProperty
	_jsii_.Get(
		j,
		"connectivityParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) InternalValue() *TfConnection_AuthParametersProperty {
	var returns *TfConnection_AuthParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) InvocationHttpParameters() TfConnection_InvocationHttpParametersPropertyOutputReference {
	var returns TfConnection_InvocationHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"invocationHttpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) InvocationHttpParametersInput() *TfConnection_InvocationHttpParametersProperty {
	var returns *TfConnection_InvocationHttpParametersProperty
	_jsii_.Get(
		j,
		"invocationHttpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) Oauth() TfConnection_OauthPropertyOutputReference {
	var returns TfConnection_OauthPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) OauthInput() *TfConnection_OauthProperty {
	var returns *TfConnection_OauthProperty
	_jsii_.Get(
		j,
		"oauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnection_AuthParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnection_AuthParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnection_AuthParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection_AuthParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfConnection.AuthParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnection_AuthParametersPropertyOutputReference_Override(t TfConnection_AuthParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfConnection.AuthParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference)SetInternalValue(val *TfConnection_AuthParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) PutApiKey(value *TfConnection_ApiKeyProperty) {
	if err := t.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putApiKey",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) PutBasic(value *TfConnection_BasicProperty) {
	if err := t.validatePutBasicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBasic",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) PutConnectivityParameters(value *TfConnection_ConnectivityParametersProperty) {
	if err := t.validatePutConnectivityParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectivityParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) PutInvocationHttpParameters(value *TfConnection_InvocationHttpParametersProperty) {
	if err := t.validatePutInvocationHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putInvocationHttpParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) PutOauth(value *TfConnection_OauthProperty) {
	if err := t.validatePutOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauth",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		t,
		"resetApiKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ResetBasic() {
	_jsii_.InvokeVoid(
		t,
		"resetBasic",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ResetConnectivityParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectivityParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ResetInvocationHttpParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetInvocationHttpParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ResetOauth() {
	_jsii_.InvokeVoid(
		t,
		"resetOauth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnection_AuthParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

