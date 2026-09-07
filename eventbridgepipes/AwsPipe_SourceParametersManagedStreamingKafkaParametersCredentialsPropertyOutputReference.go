package eventbridgepipes

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/jsii"

	"github.com/cdktn-io/cdktn-aws-go/eventbridgepipes/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference interface {
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
	InternalValue() *AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	// Experimental.
	SetInternalValue(val *AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty)
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

// The jsii proxy struct for AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference
type jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ClientCertificateTlsAuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateTlsAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InternalValue() *AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty {
	var returns *AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) SaslScram512Auth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512Auth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) SaslScram512AuthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saslScram512AuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference_Override(a AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eventbridge-pipes.AwsPipe.SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetClientCertificateTlsAuth(val *string) {
	if err := j.validateSetClientCertificateTlsAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateTlsAuth",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetInternalValue(val *AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetSaslScram512Auth(val *string) {
	if err := j.validateSetSaslScram512AuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saslScram512Auth",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ResetClientCertificateTlsAuth() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateTlsAuth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ResetSaslScram512Auth() {
	_jsii_.InvokeVoid(
		a,
		"resetSaslScram512Auth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPipe_SourceParametersManagedStreamingKafkaParametersCredentialsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

