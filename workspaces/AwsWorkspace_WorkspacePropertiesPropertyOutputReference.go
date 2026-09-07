package workspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/workspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/workspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsWorkspace_WorkspacePropertiesPropertyOutputReference interface {
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
	// Experimental.
	ComputeTypeName() *string
	// Experimental.
	SetComputeTypeName(val *string)
	// Experimental.
	ComputeTypeNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsWorkspace_WorkspacePropertiesProperty
	// Experimental.
	SetInternalValue(val *AwsWorkspace_WorkspacePropertiesProperty)
	// Experimental.
	RootVolumeSizeGib() *float64
	// Experimental.
	SetRootVolumeSizeGib(val *float64)
	// Experimental.
	RootVolumeSizeGibInput() *float64
	// Experimental.
	RunningMode() *string
	// Experimental.
	SetRunningMode(val *string)
	// Experimental.
	RunningModeAutoStopTimeoutInMinutes() *float64
	// Experimental.
	SetRunningModeAutoStopTimeoutInMinutes(val *float64)
	// Experimental.
	RunningModeAutoStopTimeoutInMinutesInput() *float64
	// Experimental.
	RunningModeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserVolumeSizeGib() *float64
	// Experimental.
	SetUserVolumeSizeGib(val *float64)
	// Experimental.
	UserVolumeSizeGibInput() *float64
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
	ResetComputeTypeName()
	// Experimental.
	ResetRootVolumeSizeGib()
	// Experimental.
	ResetRunningMode()
	// Experimental.
	ResetRunningModeAutoStopTimeoutInMinutes()
	// Experimental.
	ResetUserVolumeSizeGib()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsWorkspace_WorkspacePropertiesPropertyOutputReference
type jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ComputeTypeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ComputeTypeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computeTypeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) InternalValue() *AwsWorkspace_WorkspacePropertiesProperty {
	var returns *AwsWorkspace_WorkspacePropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RootVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RootVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rootVolumeSizeGibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RunningMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RunningModeAutoStopTimeoutInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RunningModeAutoStopTimeoutInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runningModeAutoStopTimeoutInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) RunningModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runningModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) UserVolumeSizeGib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) UserVolumeSizeGibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"userVolumeSizeGibInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsWorkspace_WorkspacePropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsWorkspace_WorkspacePropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsWorkspace_WorkspacePropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspace.WorkspacePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsWorkspace_WorkspacePropertiesPropertyOutputReference_Override(a AwsWorkspace_WorkspacePropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.AwsWorkspace.WorkspacePropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetComputeTypeName(val *string) {
	if err := j.validateSetComputeTypeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computeTypeName",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetInternalValue(val *AwsWorkspace_WorkspacePropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetRootVolumeSizeGib(val *float64) {
	if err := j.validateSetRootVolumeSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootVolumeSizeGib",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetRunningMode(val *string) {
	if err := j.validateSetRunningModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningMode",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetRunningModeAutoStopTimeoutInMinutes(val *float64) {
	if err := j.validateSetRunningModeAutoStopTimeoutInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runningModeAutoStopTimeoutInMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference)SetUserVolumeSizeGib(val *float64) {
	if err := j.validateSetUserVolumeSizeGibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userVolumeSizeGib",
		val,
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ResetComputeTypeName() {
	_jsii_.InvokeVoid(
		a,
		"resetComputeTypeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ResetRootVolumeSizeGib() {
	_jsii_.InvokeVoid(
		a,
		"resetRootVolumeSizeGib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ResetRunningMode() {
	_jsii_.InvokeVoid(
		a,
		"resetRunningMode",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ResetRunningModeAutoStopTimeoutInMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetRunningModeAutoStopTimeoutInMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ResetUserVolumeSizeGib() {
	_jsii_.InvokeVoid(
		a,
		"resetUserVolumeSizeGib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsWorkspace_WorkspacePropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

