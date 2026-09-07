package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualNode_SpecPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Backend() AwsVirtualNode_BackendPropertyList
	// Experimental.
	BackendDefaults() AwsVirtualNode_BackendDefaultsPropertyOutputReference
	// Experimental.
	BackendDefaultsInput() *AwsVirtualNode_BackendDefaultsProperty
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
	InternalValue() *AwsVirtualNode_SpecProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualNode_SpecProperty)
	// Experimental.
	Listener() AwsVirtualNode_ListenerPropertyList
	// Experimental.
	ListenerInput() interface{}
	// Experimental.
	Logging() AwsVirtualNode_LoggingPropertyOutputReference
	// Experimental.
	LoggingInput() *AwsVirtualNode_LoggingProperty
	// Experimental.
	ServiceDiscovery() AwsVirtualNode_ServiceDiscoveryPropertyOutputReference
	// Experimental.
	ServiceDiscoveryInput() *AwsVirtualNode_ServiceDiscoveryProperty
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
	PutBackendDefaults(value *AwsVirtualNode_BackendDefaultsProperty)
	// Experimental.
	PutListener(value interface{})
	// Experimental.
	PutLogging(value *AwsVirtualNode_LoggingProperty)
	// Experimental.
	PutServiceDiscovery(value *AwsVirtualNode_ServiceDiscoveryProperty)
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

// The jsii proxy struct for AwsVirtualNode_SpecPropertyOutputReference
type jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) Backend() AwsVirtualNode_BackendPropertyList {
	var returns AwsVirtualNode_BackendPropertyList
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) BackendDefaults() AwsVirtualNode_BackendDefaultsPropertyOutputReference {
	var returns AwsVirtualNode_BackendDefaultsPropertyOutputReference
	_jsii_.Get(
		j,
		"backendDefaults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) BackendDefaultsInput() *AwsVirtualNode_BackendDefaultsProperty {
	var returns *AwsVirtualNode_BackendDefaultsProperty
	_jsii_.Get(
		j,
		"backendDefaultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) BackendInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) InternalValue() *AwsVirtualNode_SpecProperty {
	var returns *AwsVirtualNode_SpecProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) Listener() AwsVirtualNode_ListenerPropertyList {
	var returns AwsVirtualNode_ListenerPropertyList
	_jsii_.Get(
		j,
		"listener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"listenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) Logging() AwsVirtualNode_LoggingPropertyOutputReference {
	var returns AwsVirtualNode_LoggingPropertyOutputReference
	_jsii_.Get(
		j,
		"logging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) LoggingInput() *AwsVirtualNode_LoggingProperty {
	var returns *AwsVirtualNode_LoggingProperty
	_jsii_.Get(
		j,
		"loggingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ServiceDiscovery() AwsVirtualNode_ServiceDiscoveryPropertyOutputReference {
	var returns AwsVirtualNode_ServiceDiscoveryPropertyOutputReference
	_jsii_.Get(
		j,
		"serviceDiscovery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ServiceDiscoveryInput() *AwsVirtualNode_ServiceDiscoveryProperty {
	var returns *AwsVirtualNode_ServiceDiscoveryProperty
	_jsii_.Get(
		j,
		"serviceDiscoveryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualNode_SpecPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualNode_SpecPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualNode_SpecPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualNode_SpecPropertyOutputReference_Override(a AwsVirtualNode_SpecPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualNode.SpecPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference)SetInternalValue(val *AwsVirtualNode_SpecProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) PutBackend(value interface{}) {
	if err := a.validatePutBackendParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackend",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) PutBackendDefaults(value *AwsVirtualNode_BackendDefaultsProperty) {
	if err := a.validatePutBackendDefaultsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBackendDefaults",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) PutListener(value interface{}) {
	if err := a.validatePutListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putListener",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) PutLogging(value *AwsVirtualNode_LoggingProperty) {
	if err := a.validatePutLoggingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLogging",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) PutServiceDiscovery(value *AwsVirtualNode_ServiceDiscoveryProperty) {
	if err := a.validatePutServiceDiscoveryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServiceDiscovery",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ResetBackendDefaults() {
	_jsii_.InvokeVoid(
		a,
		"resetBackendDefaults",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ResetListener() {
	_jsii_.InvokeVoid(
		a,
		"resetListener",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ResetLogging() {
	_jsii_.InvokeVoid(
		a,
		"resetLogging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ResetServiceDiscovery() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceDiscovery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualNode_SpecPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

