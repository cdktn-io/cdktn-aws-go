package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshVirtualNode_TimeoutPropertyOutputReference interface {
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
	Grpc() AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference
	// Experimental.
	GrpcInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty
	// Experimental.
	Http() AwsAppmeshVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference
	// Experimental.
	Http2() AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference
	// Experimental.
	Http2Input() *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property
	// Experimental.
	HttpInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty
	// Experimental.
	InternalValue() *AwsAppmeshVirtualNode_TimeoutProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshVirtualNode_TimeoutProperty)
	// Experimental.
	Tcp() AwsAppmeshVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference
	// Experimental.
	TcpInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty
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
	PutGrpc(value *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty)
	// Experimental.
	PutHttp(value *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty)
	// Experimental.
	PutHttp2(value *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property)
	// Experimental.
	PutTcp(value *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty)
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

// The jsii proxy struct for AwsAppmeshVirtualNode_TimeoutPropertyOutputReference
type jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Grpc() AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference {
	var returns AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcPropertyOutputReference
	_jsii_.Get(
		j,
		"grpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GrpcInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty {
	var returns *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty
	_jsii_.Get(
		j,
		"grpcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Http() AwsAppmeshVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference {
	var returns AwsAppmeshVirtualNode_SpecListenerTimeoutHttpPropertyOutputReference
	_jsii_.Get(
		j,
		"http",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Http2() AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference {
	var returns AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2PropertyOutputReference
	_jsii_.Get(
		j,
		"http2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Http2Input() *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property {
	var returns *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property
	_jsii_.Get(
		j,
		"http2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) HttpInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty {
	var returns *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty
	_jsii_.Get(
		j,
		"httpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) InternalValue() *AwsAppmeshVirtualNode_TimeoutProperty {
	var returns *AwsAppmeshVirtualNode_TimeoutProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Tcp() AwsAppmeshVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference {
	var returns AwsAppmeshVirtualNode_SpecListenerTimeoutTcpPropertyOutputReference
	_jsii_.Get(
		j,
		"tcp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) TcpInput() *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty {
	var returns *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty
	_jsii_.Get(
		j,
		"tcpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshVirtualNode_TimeoutPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshVirtualNode_TimeoutPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshVirtualNode_TimeoutPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualNode.TimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshVirtualNode_TimeoutPropertyOutputReference_Override(a AwsAppmeshVirtualNode_TimeoutPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualNode.TimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference)SetInternalValue(val *AwsAppmeshVirtualNode_TimeoutProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) PutGrpc(value *AwsAppmeshVirtualNode_SpecListenerTimeoutGrpcProperty) {
	if err := a.validatePutGrpcParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGrpc",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) PutHttp(value *AwsAppmeshVirtualNode_SpecListenerTimeoutHttpProperty) {
	if err := a.validatePutHttpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) PutHttp2(value *AwsAppmeshVirtualNode_SpecListenerTimeoutHttp2Property) {
	if err := a.validatePutHttp2Parameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHttp2",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) PutTcp(value *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty) {
	if err := a.validatePutTcpParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTcp",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ResetGrpc() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ResetHttp() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ResetHttp2() {
	_jsii_.InvokeVoid(
		a,
		"resetHttp2",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ResetTcp() {
	_jsii_.InvokeVoid(
		a,
		"resetTcp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualNode_TimeoutPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

