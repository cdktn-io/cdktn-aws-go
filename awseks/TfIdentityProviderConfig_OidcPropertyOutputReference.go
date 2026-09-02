package awseks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awseks/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awseks/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfIdentityProviderConfig_OidcPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ClientId() *string
	// Experimental.
	SetClientId(val *string)
	// Experimental.
	ClientIdInput() *string
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
	GroupsClaim() *string
	// Experimental.
	SetGroupsClaim(val *string)
	// Experimental.
	GroupsClaimInput() *string
	// Experimental.
	GroupsPrefix() *string
	// Experimental.
	SetGroupsPrefix(val *string)
	// Experimental.
	GroupsPrefixInput() *string
	// Experimental.
	IdentityProviderConfigName() *string
	// Experimental.
	SetIdentityProviderConfigName(val *string)
	// Experimental.
	IdentityProviderConfigNameInput() *string
	// Experimental.
	InternalValue() *TfIdentityProviderConfig_OidcProperty
	// Experimental.
	SetInternalValue(val *TfIdentityProviderConfig_OidcProperty)
	// Experimental.
	IssuerUrl() *string
	// Experimental.
	SetIssuerUrl(val *string)
	// Experimental.
	IssuerUrlInput() *string
	// Experimental.
	RequiredClaims() *map[string]*string
	// Experimental.
	SetRequiredClaims(val *map[string]*string)
	// Experimental.
	RequiredClaimsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UsernameClaim() *string
	// Experimental.
	SetUsernameClaim(val *string)
	// Experimental.
	UsernameClaimInput() *string
	// Experimental.
	UsernamePrefix() *string
	// Experimental.
	SetUsernamePrefix(val *string)
	// Experimental.
	UsernamePrefixInput() *string
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
	ResetGroupsClaim()
	// Experimental.
	ResetGroupsPrefix()
	// Experimental.
	ResetRequiredClaims()
	// Experimental.
	ResetUsernameClaim()
	// Experimental.
	ResetUsernamePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfIdentityProviderConfig_OidcPropertyOutputReference
type jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GroupsClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GroupsClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GroupsPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GroupsPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupsPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) IdentityProviderConfigName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderConfigName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) IdentityProviderConfigNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityProviderConfigNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) InternalValue() *TfIdentityProviderConfig_OidcProperty {
	var returns *TfIdentityProviderConfig_OidcProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) IssuerUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) IssuerUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) RequiredClaims() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requiredClaims",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) RequiredClaimsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"requiredClaimsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) UsernameClaim() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameClaim",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) UsernameClaimInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameClaimInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) UsernamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) UsernamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernamePrefixInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfIdentityProviderConfig_OidcPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfIdentityProviderConfig_OidcPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfIdentityProviderConfig_OidcPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-eks.TfIdentityProviderConfig.OidcPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfIdentityProviderConfig_OidcPropertyOutputReference_Override(t TfIdentityProviderConfig_OidcPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-eks.TfIdentityProviderConfig.OidcPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetGroupsClaim(val *string) {
	if err := j.validateSetGroupsClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsClaim",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetGroupsPrefix(val *string) {
	if err := j.validateSetGroupsPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupsPrefix",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetIdentityProviderConfigName(val *string) {
	if err := j.validateSetIdentityProviderConfigNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityProviderConfigName",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetInternalValue(val *TfIdentityProviderConfig_OidcProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetIssuerUrl(val *string) {
	if err := j.validateSetIssuerUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerUrl",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetRequiredClaims(val *map[string]*string) {
	if err := j.validateSetRequiredClaimsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredClaims",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetUsernameClaim(val *string) {
	if err := j.validateSetUsernameClaimParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameClaim",
		val,
	)
}

func (j *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference)SetUsernamePrefix(val *string) {
	if err := j.validateSetUsernamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernamePrefix",
		val,
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ResetGroupsClaim() {
	_jsii_.InvokeVoid(
		t,
		"resetGroupsClaim",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ResetGroupsPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetGroupsPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ResetRequiredClaims() {
	_jsii_.InvokeVoid(
		t,
		"resetRequiredClaims",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ResetUsernameClaim() {
	_jsii_.InvokeVoid(
		t,
		"resetUsernameClaim",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ResetUsernamePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetUsernamePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfIdentityProviderConfig_OidcPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

