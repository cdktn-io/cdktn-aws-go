package ssoidentitystore

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ssoidentitystore/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ssoidentitystore/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUser_NamePropertyOutputReference interface {
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
	FamilyName() *string
	// Experimental.
	SetFamilyName(val *string)
	// Experimental.
	FamilyNameInput() *string
	// Experimental.
	Formatted() *string
	// Experimental.
	SetFormatted(val *string)
	// Experimental.
	FormattedInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GivenName() *string
	// Experimental.
	SetGivenName(val *string)
	// Experimental.
	GivenNameInput() *string
	// Experimental.
	HonorificPrefix() *string
	// Experimental.
	SetHonorificPrefix(val *string)
	// Experimental.
	HonorificPrefixInput() *string
	// Experimental.
	HonorificSuffix() *string
	// Experimental.
	SetHonorificSuffix(val *string)
	// Experimental.
	HonorificSuffixInput() *string
	// Experimental.
	InternalValue() *AwsUser_NameProperty
	// Experimental.
	SetInternalValue(val *AwsUser_NameProperty)
	// Experimental.
	MiddleName() *string
	// Experimental.
	SetMiddleName(val *string)
	// Experimental.
	MiddleNameInput() *string
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
	ResetFormatted()
	// Experimental.
	ResetHonorificPrefix()
	// Experimental.
	ResetHonorificSuffix()
	// Experimental.
	ResetMiddleName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsUser_NamePropertyOutputReference
type jsiiProxy_AwsUser_NamePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) FamilyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) FamilyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"familyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) Formatted() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) FormattedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formattedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) GivenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) GivenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"givenNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) HonorificPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) HonorificPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) HonorificSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) HonorificSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"honorificSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) InternalValue() *AwsUser_NameProperty {
	var returns *AwsUser_NameProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) MiddleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) MiddleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"middleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUser_NamePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUser_NamePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUser_NamePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUser_NamePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sso-identity-store.AwsUser.NamePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUser_NamePropertyOutputReference_Override(a AwsUser_NamePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sso-identity-store.AwsUser.NamePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetFamilyName(val *string) {
	if err := j.validateSetFamilyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"familyName",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetFormatted(val *string) {
	if err := j.validateSetFormattedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatted",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetGivenName(val *string) {
	if err := j.validateSetGivenNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"givenName",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetHonorificPrefix(val *string) {
	if err := j.validateSetHonorificPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorificPrefix",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetHonorificSuffix(val *string) {
	if err := j.validateSetHonorificSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"honorificSuffix",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetInternalValue(val *AwsUser_NameProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetMiddleName(val *string) {
	if err := j.validateSetMiddleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"middleName",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUser_NamePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ResetFormatted() {
	_jsii_.InvokeVoid(
		a,
		"resetFormatted",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ResetHonorificPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetHonorificPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ResetHonorificSuffix() {
	_jsii_.InvokeVoid(
		a,
		"resetHonorificSuffix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ResetMiddleName() {
	_jsii_.InvokeVoid(
		a,
		"resetMiddleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUser_NamePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

