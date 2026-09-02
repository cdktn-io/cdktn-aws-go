package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Certificate() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference
	// Experimental.
	CertificateInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty
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
	InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty
	// Experimental.
	SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty)
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
	Validation() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference
	// Experimental.
	ValidationInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
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
	PutCertificate(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty)
	// Experimental.
	PutValidation(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty)
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

// The jsii proxy struct for TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference
type jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Certificate() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference {
	var returns TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) CertificateInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Enforce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) EnforceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Ports() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"ports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) PortsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"portsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Validation() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference {
	var returns TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference
	_jsii_.Get(
		j,
		"validation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ValidationInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
	_jsii_.Get(
		j,
		"validationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference_Override(t TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetEnforce(val interface{}) {
	if err := j.validateSetEnforceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforce",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetPorts(val *[]*float64) {
	if err := j.validateSetPortsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ports",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) PutCertificate(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty) {
	if err := t.validatePutCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCertificate",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) PutValidation(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty) {
	if err := t.validatePutValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putValidation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		t,
		"resetCertificate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ResetEnforce() {
	_jsii_.InvokeVoid(
		t,
		"resetEnforce",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ResetPorts() {
	_jsii_.InvokeVoid(
		t,
		"resetPorts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

