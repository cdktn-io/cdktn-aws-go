package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRoute_SpecPropertyOutputReference interface {
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
	GrpcRoute() TfRoute_GrpcRoutePropertyOutputReference
	// Experimental.
	GrpcRouteInput() *TfRoute_GrpcRouteProperty
	// Experimental.
	Http2Route() TfRoute_Http2RoutePropertyOutputReference
	// Experimental.
	Http2RouteInput() *TfRoute_Http2RouteProperty
	// Experimental.
	HttpRoute() TfRoute_HttpRoutePropertyOutputReference
	// Experimental.
	HttpRouteInput() *TfRoute_HttpRouteProperty
	// Experimental.
	InternalValue() *TfRoute_SpecProperty
	// Experimental.
	SetInternalValue(val *TfRoute_SpecProperty)
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	TcpRoute() TfRoute_TcpRoutePropertyOutputReference
	// Experimental.
	TcpRouteInput() *TfRoute_TcpRouteProperty
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
	PutGrpcRoute(value *TfRoute_GrpcRouteProperty)
	// Experimental.
	PutHttp2Route(value *TfRoute_Http2RouteProperty)
	// Experimental.
	PutHttpRoute(value *TfRoute_HttpRouteProperty)
	// Experimental.
	PutTcpRoute(value *TfRoute_TcpRouteProperty)
	// Experimental.
	ResetGrpcRoute()
	// Experimental.
	ResetHttp2Route()
	// Experimental.
	ResetHttpRoute()
	// Experimental.
	ResetPriority()
	// Experimental.
	ResetTcpRoute()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRoute_SpecPropertyOutputReference
type jsiiProxy_TfRoute_SpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) GrpcRoute() TfRoute_GrpcRoutePropertyOutputReference {
	var returns TfRoute_GrpcRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"grpcRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) GrpcRouteInput() *TfRoute_GrpcRouteProperty {
	var returns *TfRoute_GrpcRouteProperty
	_jsii_.Get(
		j,
		"grpcRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) Http2Route() TfRoute_Http2RoutePropertyOutputReference {
	var returns TfRoute_Http2RoutePropertyOutputReference
	_jsii_.Get(
		j,
		"http2Route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) Http2RouteInput() *TfRoute_Http2RouteProperty {
	var returns *TfRoute_Http2RouteProperty
	_jsii_.Get(
		j,
		"http2RouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) HttpRoute() TfRoute_HttpRoutePropertyOutputReference {
	var returns TfRoute_HttpRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"httpRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) HttpRouteInput() *TfRoute_HttpRouteProperty {
	var returns *TfRoute_HttpRouteProperty
	_jsii_.Get(
		j,
		"httpRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) InternalValue() *TfRoute_SpecProperty {
	var returns *TfRoute_SpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) TcpRoute() TfRoute_TcpRoutePropertyOutputReference {
	var returns TfRoute_TcpRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"tcpRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) TcpRouteInput() *TfRoute_TcpRouteProperty {
	var returns *TfRoute_TcpRouteProperty
	_jsii_.Get(
		j,
		"tcpRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRoute_SpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRoute_SpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRoute_SpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute_SpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRoute_SpecPropertyOutputReference_Override(t TfRoute_SpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetInternalValue(val *TfRoute_SpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) PutGrpcRoute(value *TfRoute_GrpcRouteProperty) {
	if err := t.validatePutGrpcRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGrpcRoute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) PutHttp2Route(value *TfRoute_Http2RouteProperty) {
	if err := t.validatePutHttp2RouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp2Route",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) PutHttpRoute(value *TfRoute_HttpRouteProperty) {
	if err := t.validatePutHttpRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttpRoute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) PutTcpRoute(value *TfRoute_TcpRouteProperty) {
	if err := t.validatePutTcpRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTcpRoute",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ResetGrpcRoute() {
	_jsii_.InvokeVoid(
		t,
		"resetGrpcRoute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ResetHttp2Route() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp2Route",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ResetHttpRoute() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpRoute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		t,
		"resetPriority",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ResetTcpRoute() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpRoute",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRoute_SpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

