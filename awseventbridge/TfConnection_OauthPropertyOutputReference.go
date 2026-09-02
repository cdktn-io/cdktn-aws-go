package awseventbridge

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridge/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridge/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConnection_OauthPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthorizationEndpoint() *string
	// Experimental.
	SetAuthorizationEndpoint(val *string)
	// Experimental.
	AuthorizationEndpointInput() *string
	// Experimental.
	ClientParameters() TfConnection_ClientParametersPropertyOutputReference
	// Experimental.
	ClientParametersInput() *TfConnection_ClientParametersProperty
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
	HttpMethod() *string
	// Experimental.
	SetHttpMethod(val *string)
	// Experimental.
	HttpMethodInput() *string
	// Experimental.
	InternalValue() *TfConnection_OauthProperty
	// Experimental.
	SetInternalValue(val *TfConnection_OauthProperty)
	// Experimental.
	OauthHttpParameters() TfConnection_OauthHttpParametersPropertyOutputReference
	// Experimental.
	OauthHttpParametersInput() *TfConnection_OauthHttpParametersProperty
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
	PutClientParameters(value *TfConnection_ClientParametersProperty)
	// Experimental.
	PutOauthHttpParameters(value *TfConnection_OauthHttpParametersProperty)
	// Experimental.
	ResetClientParameters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConnection_OauthPropertyOutputReference
type jsiiProxy_TfConnection_OauthPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) AuthorizationEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) AuthorizationEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) ClientParameters() TfConnection_ClientParametersPropertyOutputReference {
	var returns TfConnection_ClientParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"clientParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) ClientParametersInput() *TfConnection_ClientParametersProperty {
	var returns *TfConnection_ClientParametersProperty
	_jsii_.Get(
		j,
		"clientParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) HttpMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) HttpMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) InternalValue() *TfConnection_OauthProperty {
	var returns *TfConnection_OauthProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) OauthHttpParameters() TfConnection_OauthHttpParametersPropertyOutputReference {
	var returns TfConnection_OauthHttpParametersPropertyOutputReference
	_jsii_.Get(
		j,
		"oauthHttpParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) OauthHttpParametersInput() *TfConnection_OauthHttpParametersProperty {
	var returns *TfConnection_OauthHttpParametersProperty
	_jsii_.Get(
		j,
		"oauthHttpParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConnection_OauthPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConnection_OauthPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConnection_OauthPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConnection_OauthPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfConnection.OauthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConnection_OauthPropertyOutputReference_Override(t TfConnection_OauthPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge.TfConnection.OauthPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetAuthorizationEndpoint(val *string) {
	if err := j.validateSetAuthorizationEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationEndpoint",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetHttpMethod(val *string) {
	if err := j.validateSetHttpMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpMethod",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetInternalValue(val *TfConnection_OauthProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConnection_OauthPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) PutClientParameters(value *TfConnection_ClientParametersProperty) {
	if err := t.validatePutClientParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) PutOauthHttpParameters(value *TfConnection_OauthHttpParametersProperty) {
	if err := t.validatePutOauthHttpParametersParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOauthHttpParameters",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) ResetClientParameters() {
	_jsii_.InvokeVoid(
		t,
		"resetClientParameters",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConnection_OauthPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

