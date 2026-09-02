package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference interface {
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
	InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty
	// Experimental.
	SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty)
	// Experimental.
	SubjectAlternativeNames() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	// Experimental.
	SubjectAlternativeNamesInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trust() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference
	// Experimental.
	TrustInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty
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
	PutSubjectAlternativeNames(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty)
	// Experimental.
	PutTrust(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty)
	// Experimental.
	ResetSubjectAlternativeNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference
type jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) InternalValue() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNames() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNamesInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) Trust() TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference {
	var returns TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustPropertyOutputReference
	_jsii_.Get(
		j,
		"trust",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) TrustInput() *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty {
	var returns *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"trustInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference_Override(t TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualNode.SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference)SetInternalValue(val *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) PutSubjectAlternativeNames(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationSubjectAlternativeNamesProperty) {
	if err := t.validatePutSubjectAlternativeNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSubjectAlternativeNames",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) PutTrust(value *TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustProperty) {
	if err := t.validatePutTrustParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrust",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		t,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

