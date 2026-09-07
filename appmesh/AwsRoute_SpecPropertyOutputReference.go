package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute_SpecPropertyOutputReference interface {
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
	GrpcRoute() AwsRoute_GrpcRoutePropertyOutputReference
	// Experimental.
	GrpcRouteInput() *AwsRoute_GrpcRouteProperty
	// Experimental.
	Http2Route() AwsRoute_Http2RoutePropertyOutputReference
	// Experimental.
	Http2RouteInput() *AwsRoute_Http2RouteProperty
	// Experimental.
	HttpRoute() AwsRoute_HttpRoutePropertyOutputReference
	// Experimental.
	HttpRouteInput() *AwsRoute_HttpRouteProperty
	// Experimental.
	InternalValue() *AwsRoute_SpecProperty
	// Experimental.
	SetInternalValue(val *AwsRoute_SpecProperty)
	// Experimental.
	Priority() *float64
	// Experimental.
	SetPriority(val *float64)
	// Experimental.
	PriorityInput() *float64
	// Experimental.
	TcpRoute() AwsRoute_TcpRoutePropertyOutputReference
	// Experimental.
	TcpRouteInput() *AwsRoute_TcpRouteProperty
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
	PutGrpcRoute(value *AwsRoute_GrpcRouteProperty)
	// Experimental.
	PutHttp2Route(value *AwsRoute_Http2RouteProperty)
	// Experimental.
	PutHttpRoute(value *AwsRoute_HttpRouteProperty)
	// Experimental.
	PutTcpRoute(value *AwsRoute_TcpRouteProperty)
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

// The jsii proxy struct for AwsRoute_SpecPropertyOutputReference
type jsiiProxy_AwsRoute_SpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GrpcRoute() AwsRoute_GrpcRoutePropertyOutputReference {
	var returns AwsRoute_GrpcRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"grpcRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GrpcRouteInput() *AwsRoute_GrpcRouteProperty {
	var returns *AwsRoute_GrpcRouteProperty
	_jsii_.Get(
		j,
		"grpcRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) Http2Route() AwsRoute_Http2RoutePropertyOutputReference {
	var returns AwsRoute_Http2RoutePropertyOutputReference
	_jsii_.Get(
		j,
		"http2Route",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) Http2RouteInput() *AwsRoute_Http2RouteProperty {
	var returns *AwsRoute_Http2RouteProperty
	_jsii_.Get(
		j,
		"http2RouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) HttpRoute() AwsRoute_HttpRoutePropertyOutputReference {
	var returns AwsRoute_HttpRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"httpRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) HttpRouteInput() *AwsRoute_HttpRouteProperty {
	var returns *AwsRoute_HttpRouteProperty
	_jsii_.Get(
		j,
		"httpRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) InternalValue() *AwsRoute_SpecProperty {
	var returns *AwsRoute_SpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) Priority() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) PriorityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) TcpRoute() AwsRoute_TcpRoutePropertyOutputReference {
	var returns AwsRoute_TcpRoutePropertyOutputReference
	_jsii_.Get(
		j,
		"tcpRoute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) TcpRouteInput() *AwsRoute_TcpRouteProperty {
	var returns *AwsRoute_TcpRouteProperty
	_jsii_.Get(
		j,
		"tcpRouteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute_SpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute_SpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute_SpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute_SpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute_SpecPropertyOutputReference_Override(a AwsRoute_SpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetInternalValue(val *AwsRoute_SpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetPriority(val *float64) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) PutGrpcRoute(value *AwsRoute_GrpcRouteProperty) {
	if err := a.validatePutGrpcRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpcRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) PutHttp2Route(value *AwsRoute_Http2RouteProperty) {
	if err := a.validatePutHttp2RouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp2Route",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) PutHttpRoute(value *AwsRoute_HttpRouteProperty) {
	if err := a.validatePutHttpRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttpRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) PutTcpRoute(value *AwsRoute_TcpRouteProperty) {
	if err := a.validatePutTcpRouteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTcpRoute",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ResetGrpcRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ResetHttp2Route() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2Route",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ResetHttpRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ResetPriority() {
	_jsii_.InvokeVoid(
		a,
		"resetPriority",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ResetTcpRoute() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpRoute",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

