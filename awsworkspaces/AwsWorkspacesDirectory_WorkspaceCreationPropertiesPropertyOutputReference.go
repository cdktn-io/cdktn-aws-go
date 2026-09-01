package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference interface {
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
	CustomSecurityGroupId() *string
	// Experimental.
	SetCustomSecurityGroupId(val *string)
	// Experimental.
	CustomSecurityGroupIdInput() *string
	// Experimental.
	DefaultOu() *string
	// Experimental.
	SetDefaultOu(val *string)
	// Experimental.
	DefaultOuInput() *string
	// Experimental.
	EnableInternetAccess() interface{}
	// Experimental.
	SetEnableInternetAccess(val interface{})
	// Experimental.
	EnableInternetAccessInput() interface{}
	// Experimental.
	EnableMaintenanceMode() interface{}
	// Experimental.
	SetEnableMaintenanceMode(val interface{})
	// Experimental.
	EnableMaintenanceModeInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsWorkspacesDirectory_WorkspaceCreationPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsWorkspacesDirectory_WorkspaceCreationPropertiesProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserEnabledAsLocalAdministrator() interface{}
	// Experimental.
	SetUserEnabledAsLocalAdministrator(val interface{})
	// Experimental.
	UserEnabledAsLocalAdministratorInput() interface{}
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
	ResetCustomSecurityGroupId()
	// Experimental.
	ResetDefaultOu()
	// Experimental.
	ResetEnableInternetAccess()
	// Experimental.
	ResetEnableMaintenanceMode()
	// Experimental.
	ResetUserEnabledAsLocalAdministrator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference
type jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InternalValue() *AwsWorkspacesDirectory_WorkspaceCreationPropertiesProperty {
	var returns *AwsWorkspacesDirectory_WorkspaceCreationPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministratorInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspacesDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference_Override(a AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspacesDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetCustomSecurityGroupId(val *string) {
	if err := j.validateSetCustomSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetDefaultOu(val *string) {
	if err := j.validateSetDefaultOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOu",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableInternetAccess(val interface{}) {
	if err := j.validateSetEnableInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableMaintenanceMode(val interface{}) {
	if err := j.validateSetEnableMaintenanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMaintenanceMode",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetInternalValue(val *AwsWorkspacesDirectory_WorkspaceCreationPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetUserEnabledAsLocalAdministrator(val interface{}) {
	if err := j.validateSetUserEnabledAsLocalAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEnabledAsLocalAdministrator",
		val,
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetCustomSecurityGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSecurityGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetDefaultOu() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultOu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableMaintenanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMaintenanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetUserEnabledAsLocalAdministrator() {
	_jsii_.InvokeVoid(
		a,
		"resetUserEnabledAsLocalAdministrator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

