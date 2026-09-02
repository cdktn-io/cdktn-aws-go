package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsacmpca/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsacmpca/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCertificateAuthority_RevocationConfigurationPropertyOutputReference interface {
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
	CrlConfiguration() TfCertificateAuthority_CrlConfigurationPropertyOutputReference
	// Experimental.
	CrlConfigurationInput() *TfCertificateAuthority_CrlConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCertificateAuthority_RevocationConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfCertificateAuthority_RevocationConfigurationProperty)
	// Experimental.
	OcspConfiguration() TfCertificateAuthority_OcspConfigurationPropertyOutputReference
	// Experimental.
	OcspConfigurationInput() *TfCertificateAuthority_OcspConfigurationProperty
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
	PutCrlConfiguration(value *TfCertificateAuthority_CrlConfigurationProperty)
	// Experimental.
	PutOcspConfiguration(value *TfCertificateAuthority_OcspConfigurationProperty)
	// Experimental.
	ResetCrlConfiguration()
	// Experimental.
	ResetOcspConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCertificateAuthority_RevocationConfigurationPropertyOutputReference
type jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) CrlConfiguration() TfCertificateAuthority_CrlConfigurationPropertyOutputReference {
	var returns TfCertificateAuthority_CrlConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"crlConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) CrlConfigurationInput() *TfCertificateAuthority_CrlConfigurationProperty {
	var returns *TfCertificateAuthority_CrlConfigurationProperty
	_jsii_.Get(
		j,
		"crlConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) InternalValue() *TfCertificateAuthority_RevocationConfigurationProperty {
	var returns *TfCertificateAuthority_RevocationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) OcspConfiguration() TfCertificateAuthority_OcspConfigurationPropertyOutputReference {
	var returns TfCertificateAuthority_OcspConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"ocspConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) OcspConfigurationInput() *TfCertificateAuthority_OcspConfigurationProperty {
	var returns *TfCertificateAuthority_OcspConfigurationProperty
	_jsii_.Get(
		j,
		"ocspConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCertificateAuthority_RevocationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCertificateAuthority_RevocationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCertificateAuthority_RevocationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-acm-pca.TfCertificateAuthority.RevocationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCertificateAuthority_RevocationConfigurationPropertyOutputReference_Override(t TfCertificateAuthority_RevocationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-acm-pca.TfCertificateAuthority.RevocationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference)SetInternalValue(val *TfCertificateAuthority_RevocationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) PutCrlConfiguration(value *TfCertificateAuthority_CrlConfigurationProperty) {
	if err := t.validatePutCrlConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCrlConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) PutOcspConfiguration(value *TfCertificateAuthority_OcspConfigurationProperty) {
	if err := t.validatePutOcspConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOcspConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ResetCrlConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetCrlConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ResetOcspConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetOcspConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCertificateAuthority_RevocationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

