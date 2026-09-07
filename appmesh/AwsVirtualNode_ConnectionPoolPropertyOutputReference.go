package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualNode_ConnectionPoolPropertyOutputReference interface {
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
	Grpc() AwsVirtualNode_SpecListenerConnectionPoolGrpcPropertyOutputReference
	// Experimental.
	GrpcInput() *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty
	// Experimental.
	Http() AwsVirtualNode_SpecListenerConnectionPoolHttpPropertyList
	// Experimental.
	Http2() AwsVirtualNode_SpecListenerConnectionPoolHttp2PropertyList
	// Experimental.
	Http2Input() interface{}
	// Experimental.
	HttpInput() interface{}
	// Experimental.
	InternalValue() *AwsVirtualNode_ConnectionPoolProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualNode_ConnectionPoolProperty)
	// Experimental.
	Tcp() AwsVirtualNode_SpecListenerConnectionPoolTcpPropertyList
	// Experimental.
	TcpInput() interface{}
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
	PutGrpc(value *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty)
	// Experimental.
	PutHttp(value interface{})
	// Experimental.
	PutHttp2(value interface{})
	// Experimental.
	PutTcp(value interface{})
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

// The jsii proxy struct for AwsVirtualNode_ConnectionPoolPropertyOutputReference
type jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Grpc() AwsVirtualNode_SpecListenerConnectionPoolGrpcPropertyOutputReference {
	var returns AwsVirtualNode_SpecListenerConnectionPoolGrpcPropertyOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GrpcInput() *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty {
	var returns *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Http() AwsVirtualNode_SpecListenerConnectionPoolHttpPropertyList {
	var returns AwsVirtualNode_SpecListenerConnectionPoolHttpPropertyList
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Http2() AwsVirtualNode_SpecListenerConnectionPoolHttp2PropertyList {
	var returns AwsVirtualNode_SpecListenerConnectionPoolHttp2PropertyList
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Http2Input() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) HttpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) InternalValue() *AwsVirtualNode_ConnectionPoolProperty {
	var returns *AwsVirtualNode_ConnectionPoolProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Tcp() AwsVirtualNode_SpecListenerConnectionPoolTcpPropertyList {
	var returns AwsVirtualNode_SpecListenerConnectionPoolTcpPropertyList
	_jsii_.Get(
		j,
		"tcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) TcpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualNode_ConnectionPoolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualNode_ConnectionPoolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualNode_ConnectionPoolPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.ConnectionPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualNode_ConnectionPoolPropertyOutputReference_Override(a AwsVirtualNode_ConnectionPoolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.ConnectionPoolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference)SetInternalValue(val *AwsVirtualNode_ConnectionPoolProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) PutGrpc(value *AwsVirtualNode_SpecListenerConnectionPoolGrpcProperty) {
	if err := a.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) PutHttp(value interface{}) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) PutHttp2(value interface{}) {
	if err := a.validatePutHttp2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) PutTcp(value interface{}) {
	if err := a.validatePutTcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ResetTcp() {
	_jsii_.InvokeVoid(
		a,
		"resetTcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_ConnectionPoolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

