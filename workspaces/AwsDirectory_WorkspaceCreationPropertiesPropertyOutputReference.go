package workspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/workspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/workspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference interface {
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
	InternalValue() *AwsDirectory_WorkspaceCreationPropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsDirectory_WorkspaceCreationPropertiesProperty)
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

// The jsii proxy struct for AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference
type jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InternalValue() *AwsDirectory_WorkspaceCreationPropertiesProperty {
	var returns *AwsDirectory_WorkspaceCreationPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministratorInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDirectory_WorkspaceCreationPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference_Override(a AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetCustomSecurityGroupId(val *string) {
	if err := j.validateSetCustomSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetDefaultOu(val *string) {
	if err := j.validateSetDefaultOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOu",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableInternetAccess(val interface{}) {
	if err := j.validateSetEnableInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternetAccess",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableMaintenanceMode(val interface{}) {
	if err := j.validateSetEnableMaintenanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMaintenanceMode",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetInternalValue(val *AwsDirectory_WorkspaceCreationPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetUserEnabledAsLocalAdministrator(val interface{}) {
	if err := j.validateSetUserEnabledAsLocalAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEnabledAsLocalAdministrator",
		val,
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetCustomSecurityGroupId() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSecurityGroupId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetDefaultOu() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultOu",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableInternetAccess() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableInternetAccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableMaintenanceMode() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableMaintenanceMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetUserEnabledAsLocalAdministrator() {
	_jsii_.InvokeVoid(
		a,
		"resetUserEnabledAsLocalAdministrator",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

