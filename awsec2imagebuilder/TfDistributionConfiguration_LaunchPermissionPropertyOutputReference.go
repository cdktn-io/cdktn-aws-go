package awsec2imagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2imagebuilder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistributionConfiguration_LaunchPermissionPropertyOutputReference interface {
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
	InternalValue() *TfDistributionConfiguration_LaunchPermissionProperty
	// Experimental.
	SetInternalValue(val *TfDistributionConfiguration_LaunchPermissionProperty)
	// Experimental.
	OrganizationalUnitArns() *[]*string
	// Experimental.
	SetOrganizationalUnitArns(val *[]*string)
	// Experimental.
	OrganizationalUnitArnsInput() *[]*string
	// Experimental.
	OrganizationArns() *[]*string
	// Experimental.
	SetOrganizationArns(val *[]*string)
	// Experimental.
	OrganizationArnsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserGroups() *[]*string
	// Experimental.
	SetUserGroups(val *[]*string)
	// Experimental.
	UserGroupsInput() *[]*string
	// Experimental.
	UserIds() *[]*string
	// Experimental.
	SetUserIds(val *[]*string)
	// Experimental.
	UserIdsInput() *[]*string
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
	ResetOrganizationalUnitArns()
	// Experimental.
	ResetOrganizationArns()
	// Experimental.
	ResetUserGroups()
	// Experimental.
	ResetUserIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDistributionConfiguration_LaunchPermissionPropertyOutputReference
type jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) InternalValue() *TfDistributionConfiguration_LaunchPermissionProperty {
	var returns *TfDistributionConfiguration_LaunchPermissionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) OrganizationalUnitArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) OrganizationalUnitArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationalUnitArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) OrganizationArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) OrganizationArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"organizationArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) UserGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) UserGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) UserIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) UserIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"userIdsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDistributionConfiguration_LaunchPermissionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDistributionConfiguration_LaunchPermissionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDistributionConfiguration_LaunchPermissionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.LaunchPermissionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDistributionConfiguration_LaunchPermissionPropertyOutputReference_Override(t TfDistributionConfiguration_LaunchPermissionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2-image-builder.TfDistributionConfiguration.LaunchPermissionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetInternalValue(val *TfDistributionConfiguration_LaunchPermissionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetOrganizationalUnitArns(val *[]*string) {
	if err := j.validateSetOrganizationalUnitArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnitArns",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetOrganizationArns(val *[]*string) {
	if err := j.validateSetOrganizationArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationArns",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetUserGroups(val *[]*string) {
	if err := j.validateSetUserGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userGroups",
		val,
	)
}

func (j *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference)SetUserIds(val *[]*string) {
	if err := j.validateSetUserIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userIds",
		val,
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ResetOrganizationalUnitArns() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationalUnitArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ResetOrganizationArns() {
	_jsii_.InvokeVoid(
		t,
		"resetOrganizationArns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ResetUserGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetUserGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ResetUserIds() {
	_jsii_.InvokeVoid(
		t,
		"resetUserIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDistributionConfiguration_LaunchPermissionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

