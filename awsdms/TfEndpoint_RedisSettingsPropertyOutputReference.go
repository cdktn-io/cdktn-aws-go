package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_RedisSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuthPassword() *string
	// Experimental.
	SetAuthPassword(val *string)
	// Experimental.
	AuthPasswordInput() *string
	// Experimental.
	AuthType() *string
	// Experimental.
	SetAuthType(val *string)
	// Experimental.
	AuthTypeInput() *string
	// Experimental.
	AuthUserName() *string
	// Experimental.
	SetAuthUserName(val *string)
	// Experimental.
	AuthUserNameInput() *string
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
	InternalValue() *TfEndpoint_RedisSettingsProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_RedisSettingsProperty)
	// Experimental.
	Port() *float64
	// Experimental.
	SetPort(val *float64)
	// Experimental.
	PortInput() *float64
	// Experimental.
	ServerName() *string
	// Experimental.
	SetServerName(val *string)
	// Experimental.
	ServerNameInput() *string
	// Experimental.
	SslCaCertificateArn() *string
	// Experimental.
	SetSslCaCertificateArn(val *string)
	// Experimental.
	SslCaCertificateArnInput() *string
	// Experimental.
	SslSecurityProtocol() *string
	// Experimental.
	SetSslSecurityProtocol(val *string)
	// Experimental.
	SslSecurityProtocolInput() *string
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
	ResetAuthPassword()
	// Experimental.
	ResetAuthUserName()
	// Experimental.
	ResetSslCaCertificateArn()
	// Experimental.
	ResetSslSecurityProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_RedisSettingsPropertyOutputReference
type jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthUserName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUserName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) AuthUserNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authUserNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) InternalValue() *TfEndpoint_RedisSettingsProperty {
	var returns *TfEndpoint_RedisSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) SslCaCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) SslCaCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslCaCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) SslSecurityProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSecurityProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) SslSecurityProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslSecurityProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_RedisSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_RedisSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_RedisSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.RedisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_RedisSettingsPropertyOutputReference_Override(t TfEndpoint_RedisSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfEndpoint.RedisSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetAuthPassword(val *string) {
	if err := j.validateSetAuthPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authPassword",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetAuthUserName(val *string) {
	if err := j.validateSetAuthUserNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authUserName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetInternalValue(val *TfEndpoint_RedisSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetServerName(val *string) {
	if err := j.validateSetServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverName",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetSslCaCertificateArn(val *string) {
	if err := j.validateSetSslCaCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslCaCertificateArn",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetSslSecurityProtocol(val *string) {
	if err := j.validateSetSslSecurityProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslSecurityProtocol",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ResetAuthPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ResetAuthUserName() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthUserName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ResetSslCaCertificateArn() {
	_jsii_.InvokeVoid(
		t,
		"resetSslCaCertificateArn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ResetSslSecurityProtocol() {
	_jsii_.InvokeVoid(
		t,
		"resetSslSecurityProtocol",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_RedisSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

