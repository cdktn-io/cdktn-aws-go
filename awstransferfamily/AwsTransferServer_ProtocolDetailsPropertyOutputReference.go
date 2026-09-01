package awstransferfamily

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awstransferfamily/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awstransferfamily/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTransferServer_ProtocolDetailsPropertyOutputReference interface {
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
	InternalValue() *AwsTransferServer_ProtocolDetailsProperty
	// Experimental.
	SetInternalValue(val *AwsTransferServer_ProtocolDetailsProperty)
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

// The jsii proxy struct for AwsTransferServer_ProtocolDetailsPropertyOutputReference
type jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) As2Transports() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2Transports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) As2TransportsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"as2TransportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) InternalValue() *AwsTransferServer_ProtocolDetailsProperty {
	var returns *AwsTransferServer_ProtocolDetailsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) PassiveIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) PassiveIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passiveIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) SetStatOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) SetStatOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"setStatOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) TlsSessionResumptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) TlsSessionResumptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSessionResumptionModeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTransferServer_ProtocolDetailsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTransferServer_ProtocolDetailsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTransferServer_ProtocolDetailsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferServer.ProtocolDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTransferServer_ProtocolDetailsPropertyOutputReference_Override(a AwsTransferServer_ProtocolDetailsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-transfer-family.AwsTransferServer.ProtocolDetailsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetAs2Transports(val *[]*string) {
	if err := j.validateSetAs2TransportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"as2Transports",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetInternalValue(val *AwsTransferServer_ProtocolDetailsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetPassiveIp(val *string) {
	if err := j.validateSetPassiveIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passiveIp",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetSetStatOption(val *string) {
	if err := j.validateSetSetStatOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"setStatOption",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference)SetTlsSessionResumptionMode(val *string) {
	if err := j.validateSetTlsSessionResumptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSessionResumptionMode",
		val,
	)
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ResetAs2Transports() {
	_jsii_.InvokeVoid(
		a,
		"resetAs2Transports",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ResetPassiveIp() {
	_jsii_.InvokeVoid(
		a,
		"resetPassiveIp",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ResetSetStatOption() {
	_jsii_.InvokeVoid(
		a,
		"resetSetStatOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ResetTlsSessionResumptionMode() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsSessionResumptionMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTransferServer_ProtocolDetailsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

