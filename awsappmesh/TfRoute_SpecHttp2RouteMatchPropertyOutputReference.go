package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRoute_SpecHttp2RouteMatchPropertyOutputReference interface {
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
	Header() TfRoute_SpecHttp2RouteMatchHeaderPropertyList
	// Experimental.
	HeaderInput() interface{}
	// Experimental.
	InternalValue() *TfRoute_SpecHttp2RouteMatchProperty
	// Experimental.
	SetInternalValue(val *TfRoute_SpecHttp2RouteMatchProperty)
	// Experimental.
	Method() *string
	// Experimental.
	SetMethod(val *string)
	// Experimental.
	MethodInput() *string
	// Experimental.
	Path() TfRoute_SpecHttp2RouteMatchPathPropertyOutputReference
	// Experimental.
	PathInput() *TfRoute_SpecHttp2RouteMatchPathProperty
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
	QueryParameter() TfRoute_SpecHttp2RouteMatchQueryParameterPropertyList
	// Experimental.
	QueryParameterInput() interface{}
	// Experimental.
	Scheme() *string
	// Experimental.
	SetScheme(val *string)
	// Experimental.
	SchemeInput() *string
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
	PutPath(value *TfRoute_SpecHttp2RouteMatchPathProperty)
	// Experimental.
	PutQueryParameter(value interface{})
	// Experimental.
	ResetHeader()
	// Experimental.
	ResetMethod()
	// Experimental.
	ResetPath()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetQueryParameter()
	// Experimental.
	ResetScheme()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRoute_SpecHttp2RouteMatchPropertyOutputReference
type jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Header() TfRoute_SpecHttp2RouteMatchHeaderPropertyList {
	var returns TfRoute_SpecHttp2RouteMatchHeaderPropertyList
	_jsii_.Get(
		j,
		"header",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) HeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) InternalValue() *TfRoute_SpecHttp2RouteMatchProperty {
	var returns *TfRoute_SpecHttp2RouteMatchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Method() *string {
	var returns *string
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) MethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Path() TfRoute_SpecHttp2RouteMatchPathPropertyOutputReference {
	var returns TfRoute_SpecHttp2RouteMatchPathPropertyOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PathInput() *TfRoute_SpecHttp2RouteMatchPathProperty {
	var returns *TfRoute_SpecHttp2RouteMatchPathProperty
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) QueryParameter() TfRoute_SpecHttp2RouteMatchQueryParameterPropertyList {
	var returns TfRoute_SpecHttp2RouteMatchQueryParameterPropertyList
	_jsii_.Get(
		j,
		"queryParameter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) QueryParameterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryParameterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Scheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) SchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"schemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRoute_SpecHttp2RouteMatchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRoute_SpecHttp2RouteMatchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRoute_SpecHttp2RouteMatchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecHttp2RouteMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRoute_SpecHttp2RouteMatchPropertyOutputReference_Override(t TfRoute_SpecHttp2RouteMatchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecHttp2RouteMatchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetInternalValue(val *TfRoute_SpecHttp2RouteMatchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetMethod(val *string) {
	if err := j.validateSetMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"method",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetScheme(val *string) {
	if err := j.validateSetSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scheme",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PutHeader(value interface{}) {
	if err := t.validatePutHeaderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHeader",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PutPath(value *TfRoute_SpecHttp2RouteMatchPathProperty) {
	if err := t.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) PutQueryParameter(value interface{}) {
	if err := t.validatePutQueryParameterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putQueryParameter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetHeader() {
	_jsii_.InvokeVoid(
		t,
		"resetHeader",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetMethod() {
	_jsii_.InvokeVoid(
		t,
		"resetMethod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		t,
		"resetPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetQueryParameter() {
	_jsii_.InvokeVoid(
		t,
		"resetQueryParameter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ResetScheme() {
	_jsii_.InvokeVoid(
		t,
		"resetScheme",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRoute_SpecHttp2RouteMatchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

