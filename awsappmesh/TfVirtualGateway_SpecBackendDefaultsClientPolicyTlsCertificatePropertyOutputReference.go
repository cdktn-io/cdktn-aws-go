package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference interface {
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
	File() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFilePropertyOutputReference
	// Experimental.
	FileInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty
	// Experimental.
	SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty)
	// Experimental.
	Sds() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsPropertyOutputReference
	// Experimental.
	SdsInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty
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
	PutFile(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty)
	// Experimental.
	PutSds(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty)
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

// The jsii proxy struct for TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference
type jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) File() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFilePropertyOutputReference {
	var returns TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) FileInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) Sds() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsPropertyOutputReference {
	var returns TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) SdsInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference_Override(t TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference)SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) PutFile(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty) {
	if err := t.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) PutSds(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateSdsProperty) {
	if err := t.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSds",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		t,
		"resetFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		t,
		"resetSds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

