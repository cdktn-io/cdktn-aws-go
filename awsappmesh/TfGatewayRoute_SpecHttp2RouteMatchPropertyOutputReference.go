package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Header() TfGatewayRoute_SpecHttp2RouteMatchHeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	Hostname() TfGatewayRoute_SpecHttp2RouteMatchHostnamePropertyOutputReference
	// Experimental.
	HostnameInput() *TfGatewayRoute_SpecHttp2RouteMatchHostnameProperty
	// Experimental.
	InternalValue() *TfGatewayRoute_SpecHttp2RouteMatchProperty
	// Experimental.
	SetInternalValue(val *TfGatewayRoute_SpecHttp2RouteMatchProperty)
	// Experimental.
	Path() TfGatewayRoute_SpecHttp2RouteMatchPathPropertyOutputReference
	// Experimental.
	PathInput() *TfGatewayRoute_SpecHttp2RouteMatchPathProperty
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	QueryParameter() TfGatewayRoute_SpecHttp2RouteMatchQueryParameterPropertyList
	// Experimental.
	QueryParameterInput() interface{}
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
	PutHeader(value interface{})
	// Experimental.
	PutHostname(value *TfGatewayRoute_SpecHttp2RouteMatchHostnameProperty)
	// Experimental.
	PutPath(value *TfGatewayRoute_SpecHttp2RouteMatchPathProperty)
	// Experimental.
	PutQueryParameter(value interface{})
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetHostname()
	// Experimental.
	ResetPath()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetQueryParameter()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference
type jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Header() TfGatewayRoute_SpecHttp2RouteMatchHeaderPropertyList {
	var returns TfGatewayRoute_SpecHttp2RouteMatchHeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Hostname() TfGatewayRoute_SpecHttp2RouteMatchHostnamePropertyOutputReference {
	var returns TfGatewayRoute_SpecHttp2RouteMatchHostnamePropertyOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) HostnameInput() *TfGatewayRoute_SpecHttp2RouteMatchHostnameProperty {
	var returns *TfGatewayRoute_SpecHttp2RouteMatchHostnameProperty
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) InternalValue() *TfGatewayRoute_SpecHttp2RouteMatchProperty {
	var returns *TfGatewayRoute_SpecHttp2RouteMatchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Path() TfGatewayRoute_SpecHttp2RouteMatchPathPropertyOutputReference {
	var returns TfGatewayRoute_SpecHttp2RouteMatchPathPropertyOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PathInput() *TfGatewayRoute_SpecHttp2RouteMatchPathProperty {
	var returns *TfGatewayRoute_SpecHttp2RouteMatchPathProperty
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) QueryParameter() TfGatewayRoute_SpecHttp2RouteMatchQueryParameterPropertyList {
	var returns TfGatewayRoute_SpecHttp2RouteMatchQueryParameterPropertyList
	_jsii_.Get(
		j,
		"queryParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) QueryParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfGatewayRoute.SpecHttp2RouteMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference_Override(t TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfGatewayRoute.SpecHttp2RouteMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetInternalValue(val *TfGatewayRoute_SpecHttp2RouteMatchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PutHeader(value interface{}) {
	if err := t.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PutHostname(value *TfGatewayRoute_SpecHttp2RouteMatchHostnameProperty) {
	if err := t.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHostname",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PutPath(value *TfGatewayRoute_SpecHttp2RouteMatchPathProperty) {
	if err := t.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) PutQueryParameter(value interface{}) {
	if err := t.validatePutQueryParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryParameter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		t,
		"resetHostname",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		t,
		"resetPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetQueryParameter() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryParameter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttp2RouteMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

