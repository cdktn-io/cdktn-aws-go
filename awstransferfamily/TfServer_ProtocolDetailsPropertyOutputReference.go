package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServer_ProtocolDetailsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	As2Transports() *[]*string
	// Experimental.
	SetAs2Transports(val *[]*string)
	// Experimental.
	As2TransportsInput() *[]*string
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
	InternalValue() *TfServer_ProtocolDetailsProperty
	// Experimental.
	SetInternalValue(val *TfServer_ProtocolDetailsProperty)
	// Experimental.
	PassiveIp() *string
	// Experimental.
	SetPassiveIp(val *string)
	// Experimental.
	PassiveIpInput() *string
	// Experimental.
	SetStatOption() *string
	// Experimental.
	SetSetStatOption(val *string)
	// Experimental.
	SetStatOptionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TlsSessionResumptionMode() *string
	// Experimental.
	SetTlsSessionResumptionMode(val *string)
	// Experimental.
	TlsSessionResumptionModeInput() *string
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
	ResetAs2Transports()
	// Experimental.
	ResetPassiveIp()
	// Experimental.
	ResetSetStatOption()
	// Experimental.
	ResetTlsSessionResumptionMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfServer_ProtocolDetailsPropertyOutputReference
type jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) As2Transports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2Transports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) As2TransportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2TransportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) InternalValue() *TfServer_ProtocolDetailsProperty {
	var returns *TfServer_ProtocolDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) PassiveIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) PassiveIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) SetStatOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) SetStatOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) TlsSessionResumptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) TlsSessionResumptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfServer_ProtocolDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfServer_ProtocolDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfServer_ProtocolDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfServer.ProtocolDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfServer_ProtocolDetailsPropertyOutputReference_Override(t TfServer_ProtocolDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.TfServer.ProtocolDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetAs2Transports(val *[]*string) {
	if err := j.validateSetAs2TransportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"as2Transports",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetInternalValue(val *TfServer_ProtocolDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetPassiveIp(val *string) {
	if err := j.validateSetPassiveIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveIp",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetSetStatOption(val *string) {
	if err := j.validateSetSetStatOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setStatOption",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference)SetTlsSessionResumptionMode(val *string) {
	if err := j.validateSetTlsSessionResumptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSessionResumptionMode",
		val,
	)
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ResetAs2Transports() {
	_jsii_.InvokeVoid(
		t,
		"resetAs2Transports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ResetPassiveIp() {
	_jsii_.InvokeVoid(
		t,
		"resetPassiveIp",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ResetSetStatOption() {
	_jsii_.InvokeVoid(
		t,
		"resetSetStatOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ResetTlsSessionResumptionMode() {
	_jsii_.InvokeVoid(
		t,
		"resetTlsSessionResumptionMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfServer_ProtocolDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

