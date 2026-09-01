package awsmq

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmq/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmq/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMqBroker_LdapServerMetadataPropertyOutputReference interface {
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
	Hosts() *[]*string
	// Experimental.
	SetHosts(val *[]*string)
	// Experimental.
	HostsInput() *[]*string
	// Experimental.
	InternalValue() *AwsMqBroker_LdapServerMetadataProperty
	// Experimental.
	SetInternalValue(val *AwsMqBroker_LdapServerMetadataProperty)
	// Experimental.
	RoleBase() *string
	// Experimental.
	SetRoleBase(val *string)
	// Experimental.
	RoleBaseInput() *string
	// Experimental.
	RoleName() *string
	// Experimental.
	SetRoleName(val *string)
	// Experimental.
	RoleNameInput() *string
	// Experimental.
	RoleSearchMatching() *string
	// Experimental.
	SetRoleSearchMatching(val *string)
	// Experimental.
	RoleSearchMatchingInput() *string
	// Experimental.
	RoleSearchSubtree() interface{}
	// Experimental.
	SetRoleSearchSubtree(val interface{})
	// Experimental.
	RoleSearchSubtreeInput() interface{}
	// Experimental.
	ServiceAccountPassword() *string
	// Experimental.
	SetServiceAccountPassword(val *string)
	// Experimental.
	ServiceAccountPasswordInput() *string
	// Experimental.
	ServiceAccountUsername() *string
	// Experimental.
	SetServiceAccountUsername(val *string)
	// Experimental.
	ServiceAccountUsernameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserBase() *string
	// Experimental.
	SetUserBase(val *string)
	// Experimental.
	UserBaseInput() *string
	// Experimental.
	UserRoleName() *string
	// Experimental.
	SetUserRoleName(val *string)
	// Experimental.
	UserRoleNameInput() *string
	// Experimental.
	UserSearchMatching() *string
	// Experimental.
	SetUserSearchMatching(val *string)
	// Experimental.
	UserSearchMatchingInput() *string
	// Experimental.
	UserSearchSubtree() interface{}
	// Experimental.
	SetUserSearchSubtree(val interface{})
	// Experimental.
	UserSearchSubtreeInput() interface{}
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
	ResetHosts()
	// Experimental.
	ResetRoleBase()
	// Experimental.
	ResetRoleName()
	// Experimental.
	ResetRoleSearchMatching()
	// Experimental.
	ResetRoleSearchSubtree()
	// Experimental.
	ResetServiceAccountPassword()
	// Experimental.
	ResetServiceAccountUsername()
	// Experimental.
	ResetUserBase()
	// Experimental.
	ResetUserRoleName()
	// Experimental.
	ResetUserSearchMatching()
	// Experimental.
	ResetUserSearchSubtree()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsMqBroker_LdapServerMetadataPropertyOutputReference
type jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) Hosts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hosts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) HostsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hostsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) InternalValue() *AwsMqBroker_LdapServerMetadataProperty {
	var returns *AwsMqBroker_LdapServerMetadataProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) RoleSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"roleSearchSubtreeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ServiceAccountPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ServiceAccountPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ServiceAccountUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ServiceAccountUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserBase() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserBaseInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userBaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserSearchMatching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserSearchMatchingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userSearchMatchingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserSearchSubtree() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtree",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) UserSearchSubtreeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userSearchSubtreeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMqBroker_LdapServerMetadataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMqBroker_LdapServerMetadataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMqBroker_LdapServerMetadataPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-mq.AwsMqBroker.LdapServerMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMqBroker_LdapServerMetadataPropertyOutputReference_Override(a AwsMqBroker_LdapServerMetadataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-mq.AwsMqBroker.LdapServerMetadataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetHosts(val *[]*string) {
	if err := j.validateSetHostsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hosts",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetInternalValue(val *AwsMqBroker_LdapServerMetadataProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetRoleBase(val *string) {
	if err := j.validateSetRoleBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleBase",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetRoleSearchMatching(val *string) {
	if err := j.validateSetRoleSearchMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleSearchMatching",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetRoleSearchSubtree(val interface{}) {
	if err := j.validateSetRoleSearchSubtreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleSearchSubtree",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetServiceAccountPassword(val *string) {
	if err := j.validateSetServiceAccountPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountPassword",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetServiceAccountUsername(val *string) {
	if err := j.validateSetServiceAccountUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountUsername",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetUserBase(val *string) {
	if err := j.validateSetUserBaseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userBase",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetUserRoleName(val *string) {
	if err := j.validateSetUserRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userRoleName",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetUserSearchMatching(val *string) {
	if err := j.validateSetUserSearchMatchingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSearchMatching",
		val,
	)
}

func (j *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference)SetUserSearchSubtree(val interface{}) {
	if err := j.validateSetUserSearchSubtreeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userSearchSubtree",
		val,
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetHosts() {
	_jsii_.InvokeVoid(
		a,
		"resetHosts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetRoleBase() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleBase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetRoleName() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetRoleSearchMatching() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleSearchMatching",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetRoleSearchSubtree() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleSearchSubtree",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetServiceAccountPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccountPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetServiceAccountUsername() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccountUsername",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetUserBase() {
	_jsii_.InvokeVoid(
		a,
		"resetUserBase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetUserRoleName() {
	_jsii_.InvokeVoid(
		a,
		"resetUserRoleName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetUserSearchMatching() {
	_jsii_.InvokeVoid(
		a,
		"resetUserSearchMatching",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ResetUserSearchSubtree() {
	_jsii_.InvokeVoid(
		a,
		"resetUserSearchSubtree",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMqBroker_LdapServerMetadataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

