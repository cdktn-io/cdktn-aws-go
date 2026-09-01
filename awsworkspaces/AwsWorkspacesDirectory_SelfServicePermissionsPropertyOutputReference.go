package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ChangeComputeType() interface{}
	// Experimental.
	SetChangeComputeType(val interface{})
	// Experimental.
	ChangeComputeTypeInput() interface{}
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
	IncreaseVolumeSize() interface{}
	// Experimental.
	SetIncreaseVolumeSize(val interface{})
	// Experimental.
	IncreaseVolumeSizeInput() interface{}
	// Experimental.
	InternalValue() *AwsWorkspacesDirectory_SelfServicePermissionsProperty
	// Experimental.
	SetInternalValue(val *AwsWorkspacesDirectory_SelfServicePermissionsProperty)
	// Experimental.
	RebuildWorkspace() interface{}
	// Experimental.
	SetRebuildWorkspace(val interface{})
	// Experimental.
	RebuildWorkspaceInput() interface{}
	// Experimental.
	RestartWorkspace() interface{}
	// Experimental.
	SetRestartWorkspace(val interface{})
	// Experimental.
	RestartWorkspaceInput() interface{}
	// Experimental.
	SwitchRunningMode() interface{}
	// Experimental.
	SetSwitchRunningMode(val interface{})
	// Experimental.
	SwitchRunningModeInput() interface{}
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
	ResetChangeComputeType()
	// Experimental.
	ResetIncreaseVolumeSize()
	// Experimental.
	ResetRebuildWorkspace()
	// Experimental.
	ResetRestartWorkspace()
	// Experimental.
	ResetSwitchRunningMode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference
type jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ChangeComputeType() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeComputeType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ChangeComputeTypeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"changeComputeTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) IncreaseVolumeSize() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseVolumeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) IncreaseVolumeSizeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"increaseVolumeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) InternalValue() *AwsWorkspacesDirectory_SelfServicePermissionsProperty {
	var returns *AwsWorkspacesDirectory_SelfServicePermissionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) RebuildWorkspace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebuildWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) RebuildWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rebuildWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) RestartWorkspace() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartWorkspace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) RestartWorkspaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restartWorkspaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) SwitchRunningMode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchRunningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) SwitchRunningModeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"switchRunningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspacesDirectory.SelfServicePermissionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference_Override(a AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspacesDirectory.SelfServicePermissionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetChangeComputeType(val interface{}) {
	if err := j.validateSetChangeComputeTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"changeComputeType",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetIncreaseVolumeSize(val interface{}) {
	if err := j.validateSetIncreaseVolumeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"increaseVolumeSize",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetInternalValue(val *AwsWorkspacesDirectory_SelfServicePermissionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetRebuildWorkspace(val interface{}) {
	if err := j.validateSetRebuildWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebuildWorkspace",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetRestartWorkspace(val interface{}) {
	if err := j.validateSetRestartWorkspaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartWorkspace",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetSwitchRunningMode(val interface{}) {
	if err := j.validateSetSwitchRunningModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"switchRunningMode",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ResetChangeComputeType() {
	_jsii_.InvokeVoid(
		a,
		"resetChangeComputeType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ResetIncreaseVolumeSize() {
	_jsii_.InvokeVoid(
		a,
		"resetIncreaseVolumeSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ResetRebuildWorkspace() {
	_jsii_.InvokeVoid(
		a,
		"resetRebuildWorkspace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ResetRestartWorkspace() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartWorkspace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ResetSwitchRunningMode() {
	_jsii_.InvokeVoid(
		a,
		"resetSwitchRunningMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkspacesDirectory_SelfServicePermissionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

