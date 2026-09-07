package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsUserProfile_UserSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMountHomeEfs() *string
	// Experimental.
	SetAutoMountHomeEfs(val *string)
	// Experimental.
	AutoMountHomeEfsInput() *string
	// Experimental.
	CanvasAppSettings() AwsUserProfile_CanvasAppSettingsPropertyOutputReference
	// Experimental.
	CanvasAppSettingsInput() *AwsUserProfile_CanvasAppSettingsProperty
	// Experimental.
	CodeEditorAppSettings() AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *AwsUserProfile_CodeEditorAppSettingsProperty
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
	CustomFileSystemConfig() AwsUserProfile_CustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() AwsUserProfile_CustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *AwsUserProfile_CustomPosixUserConfigProperty
	// Experimental.
	DefaultLandingUri() *string
	// Experimental.
	SetDefaultLandingUri(val *string)
	// Experimental.
	DefaultLandingUriInput() *string
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsUserProfile_UserSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsUserProfile_UserSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *AwsUserProfile_JupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() AwsUserProfile_JupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *AwsUserProfile_JupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() AwsUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *AwsUserProfile_KernelGatewayAppSettingsProperty
	// Experimental.
	RSessionAppSettings() AwsUserProfile_RSessionAppSettingsPropertyOutputReference
	// Experimental.
	RSessionAppSettingsInput() *AwsUserProfile_RSessionAppSettingsProperty
	// Experimental.
	RStudioServerProAppSettings() AwsUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProAppSettingsInput() *AwsUserProfile_RStudioServerProAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharingSettings() AwsUserProfile_SharingSettingsPropertyOutputReference
	// Experimental.
	SharingSettingsInput() *AwsUserProfile_SharingSettingsProperty
	// Experimental.
	SpaceStorageSettings() AwsUserProfile_SpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *AwsUserProfile_SpaceStorageSettingsProperty
	// Experimental.
	StudioWebPortal() *string
	// Experimental.
	SetStudioWebPortal(val *string)
	// Experimental.
	StudioWebPortalInput() *string
	// Experimental.
	StudioWebPortalSettings() AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference
	// Experimental.
	StudioWebPortalSettingsInput() *AwsUserProfile_StudioWebPortalSettingsProperty
	// Experimental.
	TensorBoardAppSettings() AwsUserProfile_TensorBoardAppSettingsPropertyOutputReference
	// Experimental.
	TensorBoardAppSettingsInput() *AwsUserProfile_TensorBoardAppSettingsProperty
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
	PutCanvasAppSettings(value *AwsUserProfile_CanvasAppSettingsProperty)
	// Experimental.
	PutCodeEditorAppSettings(value *AwsUserProfile_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *AwsUserProfile_CustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *AwsUserProfile_JupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *AwsUserProfile_JupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *AwsUserProfile_KernelGatewayAppSettingsProperty)
	// Experimental.
	PutRSessionAppSettings(value *AwsUserProfile_RSessionAppSettingsProperty)
	// Experimental.
	PutRStudioServerProAppSettings(value *AwsUserProfile_RStudioServerProAppSettingsProperty)
	// Experimental.
	PutSharingSettings(value *AwsUserProfile_SharingSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *AwsUserProfile_SpaceStorageSettingsProperty)
	// Experimental.
	PutStudioWebPortalSettings(value *AwsUserProfile_StudioWebPortalSettingsProperty)
	// Experimental.
	PutTensorBoardAppSettings(value *AwsUserProfile_TensorBoardAppSettingsProperty)
	// Experimental.
	ResetAutoMountHomeEfs()
	// Experimental.
	ResetCanvasAppSettings()
	// Experimental.
	ResetCodeEditorAppSettings()
	// Experimental.
	ResetCustomFileSystemConfig()
	// Experimental.
	ResetCustomPosixUserConfig()
	// Experimental.
	ResetDefaultLandingUri()
	// Experimental.
	ResetJupyterLabAppSettings()
	// Experimental.
	ResetJupyterServerAppSettings()
	// Experimental.
	ResetKernelGatewayAppSettings()
	// Experimental.
	ResetRSessionAppSettings()
	// Experimental.
	ResetRStudioServerProAppSettings()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSharingSettings()
	// Experimental.
	ResetSpaceStorageSettings()
	// Experimental.
	ResetStudioWebPortal()
	// Experimental.
	ResetStudioWebPortalSettings()
	// Experimental.
	ResetTensorBoardAppSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsUserProfile_UserSettingsPropertyOutputReference
type jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettings() AwsUserProfile_CanvasAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_CanvasAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettingsInput() *AwsUserProfile_CanvasAppSettingsProperty {
	var returns *AwsUserProfile_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettings() AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *AwsUserProfile_CodeEditorAppSettingsProperty {
	var returns *AwsUserProfile_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfig() AwsUserProfile_CustomFileSystemConfigPropertyList {
	var returns AwsUserProfile_CustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfig() AwsUserProfile_CustomPosixUserConfigPropertyOutputReference {
	var returns AwsUserProfile_CustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfigInput() *AwsUserProfile_CustomPosixUserConfigProperty {
	var returns *AwsUserProfile_CustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) InternalValue() *AwsUserProfile_UserSettingsProperty {
	var returns *AwsUserProfile_UserSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettings() AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_JupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *AwsUserProfile_JupyterLabAppSettingsProperty {
	var returns *AwsUserProfile_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettings() AwsUserProfile_JupyterServerAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_JupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *AwsUserProfile_JupyterServerAppSettingsProperty {
	var returns *AwsUserProfile_JupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettings() AwsUserProfile_KernelGatewayAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *AwsUserProfile_KernelGatewayAppSettingsProperty {
	var returns *AwsUserProfile_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettings() AwsUserProfile_RSessionAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_RSessionAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettingsInput() *AwsUserProfile_RSessionAppSettingsProperty {
	var returns *AwsUserProfile_RSessionAppSettingsProperty
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettings() AwsUserProfile_RStudioServerProAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettingsInput() *AwsUserProfile_RStudioServerProAppSettingsProperty {
	var returns *AwsUserProfile_RStudioServerProAppSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SharingSettings() AwsUserProfile_SharingSettingsPropertyOutputReference {
	var returns AwsUserProfile_SharingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SharingSettingsInput() *AwsUserProfile_SharingSettingsProperty {
	var returns *AwsUserProfile_SharingSettingsProperty
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettings() AwsUserProfile_SpaceStorageSettingsPropertyOutputReference {
	var returns AwsUserProfile_SpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettingsInput() *AwsUserProfile_SpaceStorageSettingsProperty {
	var returns *AwsUserProfile_SpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettings() AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference {
	var returns AwsUserProfile_StudioWebPortalSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettingsInput() *AwsUserProfile_StudioWebPortalSettingsProperty {
	var returns *AwsUserProfile_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettings() AwsUserProfile_TensorBoardAppSettingsPropertyOutputReference {
	var returns AwsUserProfile_TensorBoardAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettingsInput() *AwsUserProfile_TensorBoardAppSettingsProperty {
	var returns *AwsUserProfile_TensorBoardAppSettingsProperty
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsUserProfile_UserSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsUserProfile_UserSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsUserProfile_UserSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsUserProfile_UserSettingsPropertyOutputReference_Override(a AwsUserProfile_UserSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetInternalValue(val *AwsUserProfile_UserSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutCanvasAppSettings(value *AwsUserProfile_CanvasAppSettingsProperty) {
	if err := a.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *AwsUserProfile_CodeEditorAppSettingsProperty) {
	if err := a.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := a.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *AwsUserProfile_CustomPosixUserConfigProperty) {
	if err := a.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *AwsUserProfile_JupyterLabAppSettingsProperty) {
	if err := a.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *AwsUserProfile_JupyterServerAppSettingsProperty) {
	if err := a.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *AwsUserProfile_KernelGatewayAppSettingsProperty) {
	if err := a.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutRSessionAppSettings(value *AwsUserProfile_RSessionAppSettingsProperty) {
	if err := a.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutRStudioServerProAppSettings(value *AwsUserProfile_RStudioServerProAppSettingsProperty) {
	if err := a.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutSharingSettings(value *AwsUserProfile_SharingSettingsProperty) {
	if err := a.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutSpaceStorageSettings(value *AwsUserProfile_SpaceStorageSettingsProperty) {
	if err := a.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutStudioWebPortalSettings(value *AwsUserProfile_StudioWebPortalSettingsProperty) {
	if err := a.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) PutTensorBoardAppSettings(value *AwsUserProfile_TensorBoardAppSettingsProperty) {
	if err := a.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsUserProfile_UserSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

