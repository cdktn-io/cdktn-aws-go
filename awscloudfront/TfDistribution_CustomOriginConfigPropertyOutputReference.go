package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudfront/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudfront/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistribution_CustomOriginConfigPropertyOutputReference interface {
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
	HttpPort() *float64
	// Experimental.
	SetHttpPort(val *float64)
	// Experimental.
	HttpPortInput() *float64
	// Experimental.
	HttpsPort() *float64
	// Experimental.
	SetHttpsPort(val *float64)
	// Experimental.
	HttpsPortInput() *float64
	// Experimental.
	InternalValue() *TfDistribution_CustomOriginConfigProperty
	// Experimental.
	SetInternalValue(val *TfDistribution_CustomOriginConfigProperty)
	// Experimental.
	IpAddressType() *string
	// Experimental.
	SetIpAddressType(val *string)
	// Experimental.
	IpAddressTypeInput() *string
	// Experimental.
	OriginKeepaliveTimeout() *float64
	// Experimental.
	SetOriginKeepaliveTimeout(val *float64)
	// Experimental.
	OriginKeepaliveTimeoutInput() *float64
	// Experimental.
	OriginMtlsConfig() TfDistribution_OriginMtlsConfigPropertyOutputReference
	// Experimental.
	OriginMtlsConfigInput() *TfDistribution_OriginMtlsConfigProperty
	// Experimental.
	OriginProtocolPolicy() *string
	// Experimental.
	SetOriginProtocolPolicy(val *string)
	// Experimental.
	OriginProtocolPolicyInput() *string
	// Experimental.
	OriginReadTimeout() *float64
	// Experimental.
	SetOriginReadTimeout(val *float64)
	// Experimental.
	OriginReadTimeoutInput() *float64
	// Experimental.
	OriginSslProtocols() *[]*string
	// Experimental.
	SetOriginSslProtocols(val *[]*string)
	// Experimental.
	OriginSslProtocolsInput() *[]*string
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
	PutOriginMtlsConfig(value *TfDistribution_OriginMtlsConfigProperty)
	// Experimental.
	ResetIpAddressType()
	// Experimental.
	ResetOriginKeepaliveTimeout()
	// Experimental.
	ResetOriginMtlsConfig()
	// Experimental.
	ResetOriginReadTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistribution_CustomOriginConfigPropertyOutputReference
type jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) HttpPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) HttpPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) HttpsPort() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) HttpsPortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"httpsPortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) InternalValue() *TfDistribution_CustomOriginConfigProperty {
	var returns *TfDistribution_CustomOriginConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) IpAddressType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) IpAddressTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginKeepaliveTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginKeepaliveTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originKeepaliveTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginMtlsConfig() TfDistribution_OriginMtlsConfigPropertyOutputReference {
	var returns TfDistribution_OriginMtlsConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"originMtlsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginMtlsConfigInput() *TfDistribution_OriginMtlsConfigProperty {
	var returns *TfDistribution_OriginMtlsConfigProperty
	_jsii_.Get(
		j,
		"originMtlsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginProtocolPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginProtocolPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"originProtocolPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginReadTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginReadTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"originReadTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginSslProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) OriginSslProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"originSslProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistribution_CustomOriginConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistribution_CustomOriginConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistribution_CustomOriginConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.CustomOriginConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistribution_CustomOriginConfigPropertyOutputReference_Override(t TfDistribution_CustomOriginConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudfront.TfDistribution.CustomOriginConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetHttpPort(val *float64) {
	if err := j.validateSetHttpPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpPort",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetHttpsPort(val *float64) {
	if err := j.validateSetHttpsPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsPort",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetInternalValue(val *TfDistribution_CustomOriginConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetIpAddressType(val *string) {
	if err := j.validateSetIpAddressTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipAddressType",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetOriginKeepaliveTimeout(val *float64) {
	if err := j.validateSetOriginKeepaliveTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originKeepaliveTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetOriginProtocolPolicy(val *string) {
	if err := j.validateSetOriginProtocolPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originProtocolPolicy",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetOriginReadTimeout(val *float64) {
	if err := j.validateSetOriginReadTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originReadTimeout",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetOriginSslProtocols(val *[]*string) {
	if err := j.validateSetOriginSslProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"originSslProtocols",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) PutOriginMtlsConfig(value *TfDistribution_OriginMtlsConfigProperty) {
	if err := t.validatePutOriginMtlsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOriginMtlsConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ResetIpAddressType() {
	_jsii_.InvokeVoid(
		t,
		"resetIpAddressType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ResetOriginKeepaliveTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginKeepaliveTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ResetOriginMtlsConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginMtlsConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ResetOriginReadTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetOriginReadTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistribution_CustomOriginConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

