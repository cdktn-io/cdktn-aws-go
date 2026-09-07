package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference interface {
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
	InternalValue() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty)
	// Experimental.
	SubjectAlternativeNames() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	// Experimental.
	SubjectAlternativeNamesInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trust() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference
	// Experimental.
	TrustInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
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
	PutSubjectAlternativeNames(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty)
	// Experimental.
	PutTrust(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty)
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

// The jsii proxy struct for AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference
type jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InternalValue() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty {
	var returns *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNames() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference {
	var returns AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNamesInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty {
	var returns *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Trust() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference {
	var returns AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference
	_jsii_.Get(
		j,
		"trust",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TrustInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty {
	var returns *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"trustInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference_Override(a AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetInternalValue(val *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) PutSubjectAlternativeNames(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty) {
	if err := a.validatePutSubjectAlternativeNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubjectAlternativeNames",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) PutTrust(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty) {
	if err := a.validatePutTrustParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrust",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

