package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acm() TfVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference
	// Experimental.
	AcmInput() *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty
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
	File() TfVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference
	// Experimental.
	FileInput() *TfVirtualGateway_SpecListenerTlsCertificateFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVirtualGateway_SpecListenerTlsCertificateProperty
	// Experimental.
	SetInternalValue(val *TfVirtualGateway_SpecListenerTlsCertificateProperty)
	// Experimental.
	Sds() TfVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference
	// Experimental.
	SdsInput() *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty
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
	PutAcm(value *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty)
	// Experimental.
	PutFile(value *TfVirtualGateway_SpecListenerTlsCertificateFileProperty)
	// Experimental.
	PutSds(value *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty)
	// Experimental.
	ResetAcm()
	// Experimental.
	ResetFile()
	// Experimental.
	ResetSds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference
type jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Acm() TfVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference {
	var returns TfVirtualGateway_SpecListenerTlsCertificateAcmPropertyOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) AcmInput() *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty {
	var returns *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) File() TfVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference {
	var returns TfVirtualGateway_SpecListenerTlsCertificateFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) FileInput() *TfVirtualGateway_SpecListenerTlsCertificateFileProperty {
	var returns *TfVirtualGateway_SpecListenerTlsCertificateFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InternalValue() *TfVirtualGateway_SpecListenerTlsCertificateProperty {
	var returns *TfVirtualGateway_SpecListenerTlsCertificateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Sds() TfVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference {
	var returns TfVirtualGateway_SpecListenerTlsCertificateSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) SdsInput() *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty {
	var returns *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecListenerTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference_Override(t TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecListenerTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetInternalValue(val *TfVirtualGateway_SpecListenerTlsCertificateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutAcm(value *TfVirtualGateway_SpecListenerTlsCertificateAcmProperty) {
	if err := t.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAcm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutFile(value *TfVirtualGateway_SpecListenerTlsCertificateFileProperty) {
	if err := t.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) PutSds(value *TfVirtualGateway_SpecListenerTlsCertificateSdsProperty) {
	if err := t.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSds",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		t,
		"resetAcm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		t,
		"resetFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		t,
		"resetSds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecListenerTlsCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

