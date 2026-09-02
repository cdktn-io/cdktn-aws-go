package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_SpecPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Backend() TfVirtualNode_BackendPropertyList
	// Experimental.
	BackendDefaults() TfVirtualNode_BackendDefaultsPropertyOutputReference
	// Experimental.
	BackendDefaultsInput() *TfVirtualNode_BackendDefaultsProperty
	// Experimental.
	BackendInput() interface{}
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
	InternalValue() *TfVirtualNode_SpecProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_SpecProperty)
	// Experimental.
	Listener() TfVirtualNode_ListenerPropertyList
	// Experimental.
	ListenerInput() interface{}
	// Experimental.
	Logging() TfVirtualNode_LoggingPropertyOutputReference
	// Experimental.
	LoggingInput() *TfVirtualNode_LoggingProperty
	// Experimental.
	ServiceDiscovery() TfVirtualNode_ServiceDiscoveryPropertyOutputReference
	// Experimental.
	ServiceDiscoveryInput() *TfVirtualNode_ServiceDiscoveryProperty
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
	PutBackend(value interface{})
	// Experimental.
	PutBackendDefaults(value *TfVirtualNode_BackendDefaultsProperty)
	// Experimental.
	PutListener(value interface{})
	// Experimental.
	PutLogging(value *TfVirtualNode_LoggingProperty)
	// Experimental.
	PutServiceDiscovery(value *TfVirtualNode_ServiceDiscoveryProperty)
	// Experimental.
	ResetBackend()
	// Experimental.
	ResetBackendDefaults()
	// Experimental.
	ResetListener()
	// Experimental.
	ResetLogging()
	// Experimental.
	ResetServiceDiscovery()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_SpecPropertyOutputReference
type jsiiProxy_TfVirtualNode_SpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) Backend() TfVirtualNode_BackendPropertyList {
	var returns TfVirtualNode_BackendPropertyList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) BackendDefaults() TfVirtualNode_BackendDefaultsPropertyOutputReference {
	var returns TfVirtualNode_BackendDefaultsPropertyOutputReference
	_jsii_.Get(
		j,
		"backendDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) BackendDefaultsInput() *TfVirtualNode_BackendDefaultsProperty {
	var returns *TfVirtualNode_BackendDefaultsProperty
	_jsii_.Get(
		j,
		"backendDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) InternalValue() *TfVirtualNode_SpecProperty {
	var returns *TfVirtualNode_SpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) Listener() TfVirtualNode_ListenerPropertyList {
	var returns TfVirtualNode_ListenerPropertyList
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) Logging() TfVirtualNode_LoggingPropertyOutputReference {
	var returns TfVirtualNode_LoggingPropertyOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) LoggingInput() *TfVirtualNode_LoggingProperty {
	var returns *TfVirtualNode_LoggingProperty
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ServiceDiscovery() TfVirtualNode_ServiceDiscoveryPropertyOutputReference {
	var returns TfVirtualNode_ServiceDiscoveryPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ServiceDiscoveryInput() *TfVirtualNode_ServiceDiscoveryProperty {
	var returns *TfVirtualNode_ServiceDiscoveryProperty
	_jsii_.Get(
		j,
		"serviceDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_SpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_SpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_SpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_SpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_SpecPropertyOutputReference_Override(t TfVirtualNode_SpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference)SetInternalValue(val *TfVirtualNode_SpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) PutBackend(value interface{}) {
	if err := t.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBackend",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) PutBackendDefaults(value *TfVirtualNode_BackendDefaultsProperty) {
	if err := t.validatePutBackendDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBackendDefaults",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) PutListener(value interface{}) {
	if err := t.validatePutListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putListener",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) PutLogging(value *TfVirtualNode_LoggingProperty) {
	if err := t.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLogging",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) PutServiceDiscovery(value *TfVirtualNode_ServiceDiscoveryProperty) {
	if err := t.validatePutServiceDiscoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServiceDiscovery",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ResetBackend() {
	_jsii_.InvokeVoid(
		t,
		"resetBackend",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ResetBackendDefaults() {
	_jsii_.InvokeVoid(
		t,
		"resetBackendDefaults",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ResetListener() {
	_jsii_.InvokeVoid(
		t,
		"resetListener",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		t,
		"resetLogging",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ResetServiceDiscovery() {
	_jsii_.InvokeVoid(
		t,
		"resetServiceDiscovery",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

