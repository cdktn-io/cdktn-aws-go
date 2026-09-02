package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfService_ServicePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClientAlias() TfService_ClientAliasPropertyOutputReference
	// Experimental.
	ClientAliasInput() *TfService_ClientAliasProperty
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
	DiscoveryName() *string
	// Experimental.
	SetDiscoveryName(val *string)
	// Experimental.
	DiscoveryNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IngressPortOverride() *float64
	// Experimental.
	SetIngressPortOverride(val *float64)
	// Experimental.
	IngressPortOverrideInput() *float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PortName() *string
	// Experimental.
	SetPortName(val *string)
	// Experimental.
	PortNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() TfService_TimeoutPropertyOutputReference
	// Experimental.
	TimeoutInput() *TfService_TimeoutProperty
	// Experimental.
	Tls() TfService_TlsPropertyOutputReference
	// Experimental.
	TlsInput() *TfService_TlsProperty
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
	PutClientAlias(value *TfService_ClientAliasProperty)
	// Experimental.
	PutTimeout(value *TfService_TimeoutProperty)
	// Experimental.
	PutTls(value *TfService_TlsProperty)
	// Experimental.
	ResetClientAlias()
	// Experimental.
	ResetDiscoveryName()
	// Experimental.
	ResetIngressPortOverride()
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

// The jsii proxy struct for TfService_ServicePropertyOutputReference
type jsiiProxy_TfService_ServicePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) ClientAlias() TfService_ClientAliasPropertyOutputReference {
	var returns TfService_ClientAliasPropertyOutputReference
	_jsii_.Get(
		j,
		"clientAlias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) ClientAliasInput() *TfService_ClientAliasProperty {
	var returns *TfService_ClientAliasProperty
	_jsii_.Get(
		j,
		"clientAliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) DiscoveryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) DiscoveryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"discoveryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) IngressPortOverride() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPortOverride",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) IngressPortOverrideInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ingressPortOverrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) PortName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) PortNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) Timeout() TfService_TimeoutPropertyOutputReference {
	var returns TfService_TimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) TimeoutInput() *TfService_TimeoutProperty {
	var returns *TfService_TimeoutProperty
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) Tls() TfService_TlsPropertyOutputReference {
	var returns TfService_TlsPropertyOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference) TlsInput() *TfService_TlsProperty {
	var returns *TfService_TlsProperty
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfService_ServicePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfService_ServicePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfService_ServicePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService_ServicePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService.ServicePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfService_ServicePropertyOutputReference_Override(t TfService_ServicePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.TfService.ServicePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetDiscoveryName(val *string) {
	if err := j.validateSetDiscoveryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoveryName",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetIngressPortOverride(val *float64) {
	if err := j.validateSetIngressPortOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ingressPortOverride",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetPortName(val *string) {
	if err := j.validateSetPortNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"portName",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfService_ServicePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) PutClientAlias(value *TfService_ClientAliasProperty) {
	if err := t.validatePutClientAliasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putClientAlias",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) PutTimeout(value *TfService_TimeoutProperty) {
	if err := t.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeout",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) PutTls(value *TfService_TlsProperty) {
	if err := t.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTls",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ResetClientAlias() {
	_jsii_.InvokeVoid(
		t,
		"resetClientAlias",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ResetDiscoveryName() {
	_jsii_.InvokeVoid(
		t,
		"resetDiscoveryName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ResetIngressPortOverride() {
	_jsii_.InvokeVoid(
		t,
		"resetIngressPortOverride",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		t,
		"resetTls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfService_ServicePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

