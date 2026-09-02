package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BasicAuth() *string
	// Experimental.
	SetBasicAuth(val *string)
	// Experimental.
	BasicAuthInput() *string
	// Experimental.
	ClientCertificateTlsAuth() *string
	// Experimental.
	SetClientCertificateTlsAuth(val *string)
	// Experimental.
	ClientCertificateTlsAuthInput() *string
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
	InternalValue() *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	// Experimental.
	SetInternalValue(val *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty)
	// Experimental.
	SaslScram256Auth() *string
	// Experimental.
	SetSaslScram256Auth(val *string)
	// Experimental.
	SaslScram256AuthInput() *string
	// Experimental.
	SaslScram512Auth() *string
	// Experimental.
	SetSaslScram512Auth(val *string)
	// Experimental.
	SaslScram512AuthInput() *string
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
	ResetBasicAuth()
	// Experimental.
	ResetClientCertificateTlsAuth()
	// Experimental.
	ResetSaslScram256Auth()
	// Experimental.
	ResetSaslScram512Auth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference
type jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) BasicAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) BasicAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"basicAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) InternalValue() *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty {
	var returns *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) SaslScram256Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram256Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) SaslScram256AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram256AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) SaslScram512Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) SaslScram512AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference_Override(t TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetBasicAuth(val *string) {
	if err := j.validateSetBasicAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"basicAuth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetClientCertificateTlsAuth(val *string) {
	if err := j.validateSetClientCertificateTlsAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateTlsAuth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetInternalValue(val *TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetSaslScram256Auth(val *string) {
	if err := j.validateSetSaslScram256AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram256Auth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetSaslScram512Auth(val *string) {
	if err := j.validateSetSaslScram512AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram512Auth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ResetBasicAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetBasicAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ResetClientCertificateTlsAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetClientCertificateTlsAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ResetSaslScram256Auth() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslScram256Auth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ResetSaslScram512Auth() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslScram512Auth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersSelfManagedKafkaParametersCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

