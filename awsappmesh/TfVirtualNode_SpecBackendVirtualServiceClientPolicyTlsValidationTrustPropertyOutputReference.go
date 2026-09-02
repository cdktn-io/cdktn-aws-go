package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Acm() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmPropertyOutputReference
	// Experimental.
	AcmInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty
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
	File() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference
	// Experimental.
	FileInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty)
	// Experimental.
	Sds() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsPropertyOutputReference
	// Experimental.
	SdsInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty
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
	PutAcm(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty)
	// Experimental.
	PutFile(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty)
	// Experimental.
	PutSds(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty)
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

// The jsii proxy struct for TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference
type jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) Acm() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmPropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmPropertyOutputReference
	_jsii_.Get(
		j,
		"acm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) AcmInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty
	_jsii_.Get(
		j,
		"acmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) File() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) FileInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) Sds() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsPropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) SdsInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference_Override(t TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference)SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) PutAcm(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty) {
	if err := t.validatePutAcmParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAcm",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) PutFile(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustFileProperty) {
	if err := t.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) PutSds(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustSdsProperty) {
	if err := t.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSds",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ResetAcm() {
	_jsii_.InvokeVoid(
		t,
		"resetAcm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		t,
		"resetFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		t,
		"resetSds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

