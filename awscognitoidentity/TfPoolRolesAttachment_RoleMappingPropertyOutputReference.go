package awscognitoidentity

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscognitoidentity/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscognitoidentity/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPoolRolesAttachment_RoleMappingPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AmbiguousRoleResolution() *string
	// Experimental.
	SetAmbiguousRoleResolution(val *string)
	// Experimental.
	AmbiguousRoleResolutionInput() *string
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
	IdentityProvider() *string
	// Experimental.
	SetIdentityProvider(val *string)
	// Experimental.
	IdentityProviderInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MappingRule() TfPoolRolesAttachment_MappingRulePropertyList
	// Experimental.
	MappingRuleInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutMappingRule(value interface{})
	// Experimental.
	ResetAmbiguousRoleResolution()
	// Experimental.
	ResetMappingRule()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPoolRolesAttachment_RoleMappingPropertyOutputReference
type jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) AmbiguousRoleResolution() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ambiguousRoleResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) AmbiguousRoleResolutionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ambiguousRoleResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) IdentityProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) IdentityProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) MappingRule() TfPoolRolesAttachment_MappingRulePropertyList {
	var returns TfPoolRolesAttachment_MappingRulePropertyList
	_jsii_.Get(
		j,
		"mappingRule",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) MappingRuleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mappingRuleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPoolRolesAttachment_RoleMappingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPoolRolesAttachment_RoleMappingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPoolRolesAttachment_RoleMappingPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cognito-identity.TfPoolRolesAttachment.RoleMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPoolRolesAttachment_RoleMappingPropertyOutputReference_Override(t TfPoolRolesAttachment_RoleMappingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cognito-identity.TfPoolRolesAttachment.RoleMappingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetAmbiguousRoleResolution(val *string) {
	if err := j.validateSetAmbiguousRoleResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ambiguousRoleResolution",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetIdentityProvider(val *string) {
	if err := j.validateSetIdentityProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProvider",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) PutMappingRule(value interface{}) {
	if err := t.validatePutMappingRuleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMappingRule",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ResetAmbiguousRoleResolution() {
	_jsii_.InvokeVoid(
		t,
		"resetAmbiguousRoleResolution",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ResetMappingRule() {
	_jsii_.InvokeVoid(
		t,
		"resetMappingRule",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPoolRolesAttachment_RoleMappingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

