package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference interface {
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
	InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	// Experimental.
	SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty)
	// Experimental.
	Match() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference
	// Experimental.
	MatchInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty
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
	PutMatch(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
type jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InternalValue() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Match() TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference {
	var returns TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) MatchInput() *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty {
	var returns *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference_Override(t TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetInternalValue(val *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) PutMatch(value *TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty) {
	if err := t.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMatch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

