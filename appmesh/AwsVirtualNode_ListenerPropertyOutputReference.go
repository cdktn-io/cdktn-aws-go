package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualNode_ListenerPropertyOutputReference interface {
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
	ConnectionPool() AwsVirtualNode_ConnectionPoolPropertyOutputReference
	// Experimental.
	ConnectionPoolInput() *AwsVirtualNode_ConnectionPoolProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	HealthCheck() AwsVirtualNode_HealthCheckPropertyOutputReference
	// Experimental.
	HealthCheckInput() *AwsVirtualNode_HealthCheckProperty
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	OutlierDetection() AwsVirtualNode_OutlierDetectionPropertyOutputReference
	// Experimental.
	OutlierDetectionInput() *AwsVirtualNode_OutlierDetectionProperty
	// Experimental.
	PortMapping() AwsVirtualNode_PortMappingPropertyOutputReference
	// Experimental.
	PortMappingInput() *AwsVirtualNode_PortMappingProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() AwsVirtualNode_TimeoutPropertyOutputReference
	// Experimental.
	TimeoutInput() *AwsVirtualNode_TimeoutProperty
	// Experimental.
	Tls() AwsVirtualNode_SpecListenerTlsPropertyOutputReference
	// Experimental.
	TlsInput() *AwsVirtualNode_SpecListenerTlsProperty
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
	PutConnectionPool(value *AwsVirtualNode_ConnectionPoolProperty)
	// Experimental.
	PutHealthCheck(value *AwsVirtualNode_HealthCheckProperty)
	// Experimental.
	PutOutlierDetection(value *AwsVirtualNode_OutlierDetectionProperty)
	// Experimental.
	PutPortMapping(value *AwsVirtualNode_PortMappingProperty)
	// Experimental.
	PutTimeout(value *AwsVirtualNode_TimeoutProperty)
	// Experimental.
	PutTls(value *AwsVirtualNode_SpecListenerTlsProperty)
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

// The jsii proxy struct for AwsVirtualNode_ListenerPropertyOutputReference
type jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ConnectionPool() AwsVirtualNode_ConnectionPoolPropertyOutputReference {
	var returns AwsVirtualNode_ConnectionPoolPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionPool",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ConnectionPoolInput() *AwsVirtualNode_ConnectionPoolProperty {
	var returns *AwsVirtualNode_ConnectionPoolProperty
	_jsii_.Get(
		j,
		"connectionPoolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) HealthCheck() AwsVirtualNode_HealthCheckPropertyOutputReference {
	var returns AwsVirtualNode_HealthCheckPropertyOutputReference
	_jsii_.Get(
		j,
		"healthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) HealthCheckInput() *AwsVirtualNode_HealthCheckProperty {
	var returns *AwsVirtualNode_HealthCheckProperty
	_jsii_.Get(
		j,
		"healthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) OutlierDetection() AwsVirtualNode_OutlierDetectionPropertyOutputReference {
	var returns AwsVirtualNode_OutlierDetectionPropertyOutputReference
	_jsii_.Get(
		j,
		"outlierDetection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) OutlierDetectionInput() *AwsVirtualNode_OutlierDetectionProperty {
	var returns *AwsVirtualNode_OutlierDetectionProperty
	_jsii_.Get(
		j,
		"outlierDetectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PortMapping() AwsVirtualNode_PortMappingPropertyOutputReference {
	var returns AwsVirtualNode_PortMappingPropertyOutputReference
	_jsii_.Get(
		j,
		"portMapping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PortMappingInput() *AwsVirtualNode_PortMappingProperty {
	var returns *AwsVirtualNode_PortMappingProperty
	_jsii_.Get(
		j,
		"portMappingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) Timeout() AwsVirtualNode_TimeoutPropertyOutputReference {
	var returns AwsVirtualNode_TimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) TimeoutInput() *AwsVirtualNode_TimeoutProperty {
	var returns *AwsVirtualNode_TimeoutProperty
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) Tls() AwsVirtualNode_SpecListenerTlsPropertyOutputReference {
	var returns AwsVirtualNode_SpecListenerTlsPropertyOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) TlsInput() *AwsVirtualNode_SpecListenerTlsProperty {
	var returns *AwsVirtualNode_SpecListenerTlsProperty
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualNode_ListenerPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsVirtualNode_ListenerPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualNode_ListenerPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.ListenerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualNode_ListenerPropertyOutputReference_Override(a AwsVirtualNode_ListenerPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.ListenerPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutConnectionPool(value *AwsVirtualNode_ConnectionPoolProperty) {
	if err := a.validatePutConnectionPoolParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionPool",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutHealthCheck(value *AwsVirtualNode_HealthCheckProperty) {
	if err := a.validatePutHealthCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHealthCheck",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutOutlierDetection(value *AwsVirtualNode_OutlierDetectionProperty) {
	if err := a.validatePutOutlierDetectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOutlierDetection",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutPortMapping(value *AwsVirtualNode_PortMappingProperty) {
	if err := a.validatePutPortMappingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPortMapping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutTimeout(value *AwsVirtualNode_TimeoutProperty) {
	if err := a.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) PutTls(value *AwsVirtualNode_SpecListenerTlsProperty) {
	if err := a.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTls",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ResetConnectionPool() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionPool",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ResetHealthCheck() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthCheck",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ResetOutlierDetection() {
	_jsii_.InvokeVoid(
		a,
		"resetOutlierDetection",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		a,
		"resetTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_ListenerPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

