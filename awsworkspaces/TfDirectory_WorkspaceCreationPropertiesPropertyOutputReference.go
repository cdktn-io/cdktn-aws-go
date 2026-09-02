package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference interface {
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
	InternalValue() *TfDirectory_WorkspaceCreationPropertiesProperty
	// Experimental.
	SetInternalValue(val *TfDirectory_WorkspaceCreationPropertiesProperty)
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

// The jsii proxy struct for TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference
type jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) CustomSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOu() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOu",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) DefaultOuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultOuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableInternetAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableInternetAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) EnableMaintenanceModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableMaintenanceModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InternalValue() *TfDirectory_WorkspaceCreationPropertiesProperty {
	var returns *TfDirectory_WorkspaceCreationPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministrator() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministrator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) UserEnabledAsLocalAdministratorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userEnabledAsLocalAdministratorInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDirectory_WorkspaceCreationPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDirectory_WorkspaceCreationPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDirectory_WorkspaceCreationPropertiesPropertyOutputReference_Override(t TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory.WorkspaceCreationPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetCustomSecurityGroupId(val *string) {
	if err := j.validateSetCustomSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetDefaultOu(val *string) {
	if err := j.validateSetDefaultOuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultOu",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableInternetAccess(val interface{}) {
	if err := j.validateSetEnableInternetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableInternetAccess",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetEnableMaintenanceMode(val interface{}) {
	if err := j.validateSetEnableMaintenanceModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableMaintenanceMode",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetInternalValue(val *TfDirectory_WorkspaceCreationPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference)SetUserEnabledAsLocalAdministrator(val interface{}) {
	if err := j.validateSetUserEnabledAsLocalAdministratorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userEnabledAsLocalAdministrator",
		val,
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetCustomSecurityGroupId() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomSecurityGroupId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetDefaultOu() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultOu",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableInternetAccess() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableInternetAccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetEnableMaintenanceMode() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableMaintenanceMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ResetUserEnabledAsLocalAdministrator() {
	_jsii_.InvokeVoid(
		t,
		"resetUserEnabledAsLocalAdministrator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDirectory_WorkspaceCreationPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

