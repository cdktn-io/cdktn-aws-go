package eventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridge/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsConnection_AuthParametersPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApiKey() AwsConnection_ApiKeyPropertyOutputReference
	// Experimental.
	ApiKeyInput() *AwsConnection_ApiKeyProperty
	// Experimental.
	Basic() AwsConnection_BasicPropertyOutputReference
	// Experimental.
	BasicInput() *AwsConnection_BasicProperty
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
	ConnectivityParameters() AwsConnection_ConnectivityParametersPropertyOutputReference
	// Experimental.
	ConnectivityParametersInput() *AwsConnection_ConnectivityParametersProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsConnection_AuthParametersProperty
	// Experimental.
	SetInternalValue(val *AwsConnection_AuthParametersProperty)
	// Experimental.
	InvocationHttpParameters() AwsConnection_InvocationHttpParametersPropertyOutputReference
	// Experimental.
	InvocationHttpParametersInput() *AwsConnection_InvocationHttpParametersProperty
	// Experimental.
	Oauth() AwsConnection_OauthPropertyOutputReference
	// Experimental.
	OauthInput() *AwsConnection_OauthProperty
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
	PutApiKey(value *AwsConnection_ApiKeyProperty)
	// Experimental.
	PutBasic(value *AwsConnection_BasicProperty)
	// Experimental.
	PutConnectivityParameters(value *AwsConnection_ConnectivityParametersProperty)
	// Experimental.
	PutInvocationHttpParameters(value *AwsConnection_InvocationHttpParametersProperty)
	// Experimental.
	PutOauth(value *AwsConnection_OauthProperty)
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

// The jsii proxy struct for AwsConnection_AuthParametersPropertyOutputReference
type jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ApiKey() AwsConnection_ApiKeyPropertyOutputReference {
	var returns AwsConnection_ApiKeyPropertyOutputReference
	_jsii_.Get(
		j,
		"apiKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ApiKeyInput() *AwsConnection_ApiKeyProperty {
	var returns *AwsConnection_ApiKeyProperty
	_jsii_.Get(
		j,
		"apiKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) Basic() AwsConnection_BasicPropertyOutputReference {
	var returns AwsConnection_BasicPropertyOutputReference
	_jsii_.Get(
		j,
		"basic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) BasicInput() *AwsConnection_BasicProperty {
	var returns *AwsConnection_BasicProperty
	_jsii_.Get(
		j,
		"basicInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ConnectivityParameters() AwsConnection_ConnectivityParametersPropertyOutputReference {
	var returns AwsConnection_ConnectivityParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"connectivityParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ConnectivityParametersInput() *AwsConnection_ConnectivityParametersProperty {
	var returns *AwsConnection_ConnectivityParametersProperty
	_jsii_.Get(
		j,
		"connectivityParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) InternalValue() *AwsConnection_AuthParametersProperty {
	var returns *AwsConnection_AuthParametersProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) InvocationHttpParameters() AwsConnection_InvocationHttpParametersPropertyOutputReference {
	var returns AwsConnection_InvocationHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"invocationHttpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) InvocationHttpParametersInput() *AwsConnection_InvocationHttpParametersProperty {
	var returns *AwsConnection_InvocationHttpParametersProperty
	_jsii_.Get(
		j,
		"invocationHttpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) Oauth() AwsConnection_OauthPropertyOutputReference {
	var returns AwsConnection_OauthPropertyOutputReference
	_jsii_.Get(
		j,
		"oauth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) OauthInput() *AwsConnection_OauthProperty {
	var returns *AwsConnection_OauthProperty
	_jsii_.Get(
		j,
		"oauthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsConnection_AuthParametersPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsConnection_AuthParametersPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsConnection_AuthParametersPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsConnection.AuthParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsConnection_AuthParametersPropertyOutputReference_Override(a AwsConnection_AuthParametersPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.AwsConnection.AuthParametersPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference)SetInternalValue(val *AwsConnection_AuthParametersProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) PutApiKey(value *AwsConnection_ApiKeyProperty) {
	if err := a.validatePutApiKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApiKey",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) PutBasic(value *AwsConnection_BasicProperty) {
	if err := a.validatePutBasicParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBasic",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) PutConnectivityParameters(value *AwsConnection_ConnectivityParametersProperty) {
	if err := a.validatePutConnectivityParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectivityParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) PutInvocationHttpParameters(value *AwsConnection_InvocationHttpParametersProperty) {
	if err := a.validatePutInvocationHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putInvocationHttpParameters",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) PutOauth(value *AwsConnection_OauthProperty) {
	if err := a.validatePutOauthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOauth",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ResetApiKey() {
	_jsii_.InvokeVoid(
		a,
		"resetApiKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ResetBasic() {
	_jsii_.InvokeVoid(
		a,
		"resetBasic",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ResetConnectivityParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectivityParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ResetInvocationHttpParameters() {
	_jsii_.InvokeVoid(
		a,
		"resetInvocationHttpParameters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ResetOauth() {
	_jsii_.InvokeVoid(
		a,
		"resetOauth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsConnection_AuthParametersPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

