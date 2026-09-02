package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference interface {
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
	File() TfVirtualNode_SpecListenerTlsValidationTrustFilePropertyOutputReference
	// Experimental.
	FileInput() *TfVirtualNode_SpecListenerTlsValidationTrustFileProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVirtualNode_SpecListenerTlsValidationTrustProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_SpecListenerTlsValidationTrustProperty)
	// Experimental.
	Sds() TfVirtualNode_SpecListenerTlsValidationTrustSdsPropertyOutputReference
	// Experimental.
	SdsInput() *TfVirtualNode_SpecListenerTlsValidationTrustSdsProperty
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
	PutFile(value *TfVirtualNode_SpecListenerTlsValidationTrustFileProperty)
	// Experimental.
	PutSds(value *TfVirtualNode_SpecListenerTlsValidationTrustSdsProperty)
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

// The jsii proxy struct for TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference
type jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) File() TfVirtualNode_SpecListenerTlsValidationTrustFilePropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTlsValidationTrustFilePropertyOutputReference
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) FileInput() *TfVirtualNode_SpecListenerTlsValidationTrustFileProperty {
	var returns *TfVirtualNode_SpecListenerTlsValidationTrustFileProperty
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) InternalValue() *TfVirtualNode_SpecListenerTlsValidationTrustProperty {
	var returns *TfVirtualNode_SpecListenerTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) Sds() TfVirtualNode_SpecListenerTlsValidationTrustSdsPropertyOutputReference {
	var returns TfVirtualNode_SpecListenerTlsValidationTrustSdsPropertyOutputReference
	_jsii_.Get(
		j,
		"sds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) SdsInput() *TfVirtualNode_SpecListenerTlsValidationTrustSdsProperty {
	var returns *TfVirtualNode_SpecListenerTlsValidationTrustSdsProperty
	_jsii_.Get(
		j,
		"sdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecListenerTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference_Override(t TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecListenerTlsValidationTrustPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference)SetInternalValue(val *TfVirtualNode_SpecListenerTlsValidationTrustProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) PutFile(value *TfVirtualNode_SpecListenerTlsValidationTrustFileProperty) {
	if err := t.validatePutFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFile",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) PutSds(value *TfVirtualNode_SpecListenerTlsValidationTrustSdsProperty) {
	if err := t.validatePutSdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSds",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ResetFile() {
	_jsii_.InvokeVoid(
		t,
		"resetFile",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ResetSds() {
	_jsii_.InvokeVoid(
		t,
		"resetSds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecListenerTlsValidationTrustPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

