package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_TimeoutPropertyOutputReference interface {
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
	Grpc() TfVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference
	// Experimental.
	GrpcInput() *TfVirtualNode_SpecListenerTimeoutGrpcProperty
	// Experimental.
	Http() TfVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference
	// Experimental.
	Http2() TfVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference
	// Experimental.
	Http2Input() *TfVirtualNode_SpecListenerTimeoutHttp2Property
	// Experimental.
	HttpInput() *TfVirtualNode_SpecListenerTimeoutHttpProperty
	// Experimental.
	InternalValue() *TfVirtualNode_TimeoutProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_TimeoutProperty)
	// Experimental.
	Tcp() TfVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference
	// Experimental.
	TcpInput() *TfVirtualNode_SpecListenerTimeoutTcpProperty
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
	PutGrpc(value *TfVirtualNode_SpecListenerTimeoutGrpcProperty)
	// Experimental.
	PutHttp(value *TfVirtualNode_SpecListenerTimeoutHttpProperty)
	// Experimental.
	PutHttp2(value *TfVirtualNode_SpecListenerTimeoutHttp2Property)
	// Experimental.
	PutTcp(value *TfVirtualNode_SpecListenerTimeoutTcpProperty)
	// Experimental.
	ResetGrpc()
	// Experimental.
	ResetHttp()
	// Experimental.
	ResetHttp2()
	// Experimental.
	ResetTcp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_TimeoutPropertyOutputReference
type jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Grpc() TfVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GrpcInput() *TfVirtualNode_SpecListenerTimeoutGrpcProperty {
	var returns *TfVirtualNode_SpecListenerTimeoutGrpcProperty
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Http() TfVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Http2() TfVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Http2Input() *TfVirtualNode_SpecListenerTimeoutHttp2Property {
	var returns *TfVirtualNode_SpecListenerTimeoutHttp2Property
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) HttpInput() *TfVirtualNode_SpecListenerTimeoutHttpProperty {
	var returns *TfVirtualNode_SpecListenerTimeoutHttpProperty
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) InternalValue() *TfVirtualNode_TimeoutProperty {
	var returns *TfVirtualNode_TimeoutProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Tcp() TfVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference
	_jsii_.Get(
		j,
		"tcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) TcpInput() *TfVirtualNode_SpecListenerTimeoutTcpProperty {
	var returns *TfVirtualNode_SpecListenerTimeoutTcpProperty
	_jsii_.Get(
		j,
		"tcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_TimeoutPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_TimeoutPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_TimeoutPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.TimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_TimeoutPropertyOutputReference_Override(t TfVirtualNode_TimeoutPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.TimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference)SetInternalValue(val *TfVirtualNode_TimeoutProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) PutGrpc(value *TfVirtualNode_SpecListenerTimeoutGrpcProperty) {
	if err := t.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putGrpc",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) PutHttp(value *TfVirtualNode_SpecListenerTimeoutHttpProperty) {
	if err := t.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) PutHttp2(value *TfVirtualNode_SpecListenerTimeoutHttp2Property) {
	if err := t.validatePutHttp2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHttp2",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) PutTcp(value *TfVirtualNode_SpecListenerTimeoutTcpProperty) {
	if err := t.validatePutTcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTcp",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		t,
		"resetGrpc",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		t,
		"resetHttp2",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ResetTcp() {
	_jsii_.InvokeVoid(
		t,
		"resetTcp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_TimeoutPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

