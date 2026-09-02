package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Certificate() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificatePropertyOutputReference
	// Experimental.
	CertificateInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty
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
	Enforce() interface{}
	// Experimental.
	SetEnforce(val interface{})
	// Experimental.
	EnforceInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty)
	// Experimental.
	Ports() *[]*float64
	// Experimental.
	SetPorts(val *[]*float64)
	// Experimental.
	PortsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Validation() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference
	// Experimental.
	ValidationInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty
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
	PutCertificate(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty)
	// Experimental.
	PutValidation(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty)
	// Experimental.
	ResetCertificate()
	// Experimental.
	ResetEnforce()
	// Experimental.
	ResetPorts()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference
type jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Certificate() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificatePropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificatePropertyOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) CertificateInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Enforce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) EnforceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Ports() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) PortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Validation() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ValidationInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty
	_jsii_.Get(
		j,
		"validationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference_Override(t TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetEnforce(val interface{}) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetPorts(val *[]*float64) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) PutCertificate(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateProperty) {
	if err := t.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCertificate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) PutValidation(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty) {
	if err := t.validatePutValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putValidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		t,
		"resetEnforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		t,
		"resetPorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

