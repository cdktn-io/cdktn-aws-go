package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference interface {
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
	InternalValue() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	// Experimental.
	SetInternalValue(val *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty)
	// Experimental.
	Match() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference
	// Experimental.
	MatchInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty
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
	PutMatch(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference
type jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InternalValue() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty {
	var returns *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Match() AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference {
	var returns AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) MatchInput() *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty {
	var returns *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference_Override(a AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsVirtualGateway.SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetInternalValue(val *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) PutMatch(value *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty) {
	if err := a.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

