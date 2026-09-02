package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_ListenerPropertyOutputReference interface {
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
	// Experimental.
	ConnectionPool() TfVirtualNode_ConnectionPoolPropertyOutputReference
	// Experimental.
	ConnectionPoolInput() *TfVirtualNode_ConnectionPoolProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	HealthCheck() TfVirtualNode_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *TfVirtualNode_HealthCheckProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutlierDetection() TfVirtualNode_OutlierDetectionPropertyOutputReference
	// Experimental.
	OutlierDetectionInput() *TfVirtualNode_OutlierDetectionProperty
	// Experimental.
	PortMapping() TfVirtualNode_PortMappingPropertyOutputReference
	// Experimental.
	PortMappingInput() *TfVirtualNode_PortMappingProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() TfVirtualNode_TimeoutPropertyOutputReference
	// Experimental.
	TimeoutInput() *TfVirtualNode_TimeoutProperty
	// Experimental.
	Tls() TfVirtualNode_SpecListenerTlsPropertyOutputReference
	// Experimental.
	TlsInput() *TfVirtualNode_SpecListenerTlsProperty
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
	PutConnectionPool(value *TfVirtualNode_ConnectionPoolProperty)
	// Experimental.
	PutHealthCheck(value *TfVirtualNode_HealthCheckProperty)
	// Experimental.
	PutOutlierDetection(value *TfVirtualNode_OutlierDetectionProperty)
	// Experimental.
	PutPortMapping(value *TfVirtualNode_PortMappingProperty)
	// Experimental.
	PutTimeout(value *TfVirtualNode_TimeoutProperty)
	// Experimental.
	PutTls(value *TfVirtualNode_SpecListenerTlsProperty)
	// Experimental.
	ResetConnectionPool()
	// Experimental.
	ResetHealthCheck()
	// Experimental.
	ResetOutlierDetection()
	// Experimental.
	ResetTimeout()
	// Experimental.
	ResetTls()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_ListenerPropertyOutputReference
type jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ConnectionPool() TfVirtualNode_ConnectionPoolPropertyOutputReference {
	var returns TfVirtualNode_ConnectionPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ConnectionPoolInput() *TfVirtualNode_ConnectionPoolProperty {
	var returns *TfVirtualNode_ConnectionPoolProperty
	_jsii_.Get(
		j,
		"connectionPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) HealthCheck() TfVirtualNode_HealthCheckPropertyOutputReference {
	var returns TfVirtualNode_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) HealthCheckInput() *TfVirtualNode_HealthCheckProperty {
	var returns *TfVirtualNode_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) OutlierDetection() TfVirtualNode_OutlierDetectionPropertyOutputReference {
	var returns TfVirtualNode_OutlierDetectionPropertyOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) OutlierDetectionInput() *TfVirtualNode_OutlierDetectionProperty {
	var returns *TfVirtualNode_OutlierDetectionProperty
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PortMapping() TfVirtualNode_PortMappingPropertyOutputReference {
	var returns TfVirtualNode_PortMappingPropertyOutputReference
	_jsii_.Get(
		j,
		"portMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PortMappingInput() *TfVirtualNode_PortMappingProperty {
	var returns *TfVirtualNode_PortMappingProperty
	_jsii_.Get(
		j,
		"portMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) Timeout() TfVirtualNode_TimeoutPropertyOutputReference {
	var returns TfVirtualNode_TimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) TimeoutInput() *TfVirtualNode_TimeoutProperty {
	var returns *TfVirtualNode_TimeoutProperty
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) Tls() TfVirtualNode_SpecListenerTlsPropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTlsPropertyOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) TlsInput() *TfVirtualNode_SpecListenerTlsProperty {
	var returns *TfVirtualNode_SpecListenerTlsProperty
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_ListenerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfVirtualNode_ListenerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_ListenerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.ListenerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_ListenerPropertyOutputReference_Override(t TfVirtualNode_ListenerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.ListenerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutConnectionPool(value *TfVirtualNode_ConnectionPoolProperty) {
	if err := t.validatePutConnectionPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionPool",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutHealthCheck(value *TfVirtualNode_HealthCheckProperty) {
	if err := t.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutOutlierDetection(value *TfVirtualNode_OutlierDetectionProperty) {
	if err := t.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutPortMapping(value *TfVirtualNode_PortMappingProperty) {
	if err := t.validatePutPortMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPortMapping",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutTimeout(value *TfVirtualNode_TimeoutProperty) {
	if err := t.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeout",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) PutTls(value *TfVirtualNode_SpecListenerTlsProperty) {
	if err := t.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTls",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ResetConnectionPool() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionPool",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		t,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		t,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		t,
		"resetTls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_ListenerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

