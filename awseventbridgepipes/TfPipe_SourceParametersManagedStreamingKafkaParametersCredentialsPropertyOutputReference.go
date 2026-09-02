package awseventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	// Experimental.
	SetInternalValue(val *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty)
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
	ResetClientCertificateTlsAuth()
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

// The jsii proxy struct for TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference
type jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InternalValue() *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty {
	var returns *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) SaslScram512Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) SaslScram512AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference_Override(t TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.TfPipe.SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetClientCertificateTlsAuth(val *string) {
	if err := j.validateSetClientCertificateTlsAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateTlsAuth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetInternalValue(val *TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetSaslScram512Auth(val *string) {
	if err := j.validateSetSaslScram512AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram512Auth",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ResetClientCertificateTlsAuth() {
	_jsii_.InvokeVoid(
		t,
		"resetClientCertificateTlsAuth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ResetSaslScram512Auth() {
	_jsii_.InvokeVoid(
		t,
		"resetSaslScram512Auth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

