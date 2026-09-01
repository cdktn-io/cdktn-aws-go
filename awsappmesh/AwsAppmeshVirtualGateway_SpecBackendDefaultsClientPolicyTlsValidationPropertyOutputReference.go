package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference interface {
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
	InternalValue() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty)
	// Experimental.
	SubjectAlternativeNames() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	// Experimental.
	SubjectAlternativeNamesInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Trust() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference
	// Experimental.
	TrustInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
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
	PutSubjectAlternativeNames(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty)
	// Experimental.
	PutTrust(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty)
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

// The jsii proxy struct for AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference
type jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InternalValue() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNames() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference {
	var returns AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
	_jsii_.Get(
		j,
		"subjectAlternativeNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) SubjectAlternativeNamesInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	_jsii_.Get(
		j,
		"subjectAlternativeNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Trust() AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference {
	var returns AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustPropertyOutputReference
	_jsii_.Get(
		j,
		"trust",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) TrustInput() *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty {
	var returns *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty
	_jsii_.Get(
		j,
		"trustInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference_Override(a AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetInternalValue(val *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) PutSubjectAlternativeNames(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty) {
	if err := a.validatePutSubjectAlternativeNamesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSubjectAlternativeNames",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) PutTrust(value *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustProperty) {
	if err := a.validatePutTrustParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrust",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ResetSubjectAlternativeNames() {
	_jsii_.InvokeVoid(
		a,
		"resetSubjectAlternativeNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

