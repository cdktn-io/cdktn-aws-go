package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmsk/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmsk/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCluster_ClientAuthenticationPropertyOutputReference interface {
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
	InternalValue() *TfCluster_ClientAuthenticationProperty
	// Experimental.
	SetInternalValue(val *TfCluster_ClientAuthenticationProperty)
	// Experimental.
	Sasl() TfCluster_ClientAuthenticationSaslPropertyOutputReference
	// Experimental.
	SaslInput() *TfCluster_ClientAuthenticationSaslProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Tls() TfCluster_TlsPropertyOutputReference
	// Experimental.
	TlsInput() *TfCluster_TlsProperty
	// Experimental.
	Unauthenticated() interface{}
	// Experimental.
	SetUnauthenticated(val interface{})
	// Experimental.
	UnauthenticatedInput() interface{}
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
	PutSasl(value *TfCluster_ClientAuthenticationSaslProperty)
	// Experimental.
	PutTls(value *TfCluster_TlsProperty)
	// Experimental.
	ResetSasl()
	// Experimental.
	ResetTls()
	// Experimental.
	ResetUnauthenticated()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCluster_ClientAuthenticationPropertyOutputReference
type jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) InternalValue() *TfCluster_ClientAuthenticationProperty {
	var returns *TfCluster_ClientAuthenticationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) Sasl() TfCluster_ClientAuthenticationSaslPropertyOutputReference {
	var returns TfCluster_ClientAuthenticationSaslPropertyOutputReference
	_jsii_.Get(
		j,
		"sasl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) SaslInput() *TfCluster_ClientAuthenticationSaslProperty {
	var returns *TfCluster_ClientAuthenticationSaslProperty
	_jsii_.Get(
		j,
		"saslInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) Tls() TfCluster_TlsPropertyOutputReference {
	var returns TfCluster_TlsPropertyOutputReference
	_jsii_.Get(
		j,
		"tls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) TlsInput() *TfCluster_TlsProperty {
	var returns *TfCluster_TlsProperty
	_jsii_.Get(
		j,
		"tlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) Unauthenticated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unauthenticated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) UnauthenticatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unauthenticatedInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCluster_ClientAuthenticationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCluster_ClientAuthenticationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCluster_ClientAuthenticationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster.ClientAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCluster_ClientAuthenticationPropertyOutputReference_Override(t TfCluster_ClientAuthenticationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-msk.TfCluster.ClientAuthenticationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetInternalValue(val *TfCluster_ClientAuthenticationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference)SetUnauthenticated(val interface{}) {
	if err := j.validateSetUnauthenticatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticated",
		val,
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) PutSasl(value *TfCluster_ClientAuthenticationSaslProperty) {
	if err := t.validatePutSaslParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSasl",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) PutTls(value *TfCluster_TlsProperty) {
	if err := t.validatePutTlsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTls",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ResetSasl() {
	_jsii_.InvokeVoid(
		t,
		"resetSasl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ResetTls() {
	_jsii_.InvokeVoid(
		t,
		"resetTls",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ResetUnauthenticated() {
	_jsii_.InvokeVoid(
		t,
		"resetUnauthenticated",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCluster_ClientAuthenticationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

