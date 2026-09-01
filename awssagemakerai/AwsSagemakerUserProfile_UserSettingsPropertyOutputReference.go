package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerUserProfile_UserSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMountHomeEfs() *string
	// Experimental.
	SetAutoMountHomeEfs(val *string)
	// Experimental.
	AutoMountHomeEfsInput() *string
	// Experimental.
	CanvasAppSettings() AwsSagemakerUserProfile_CanvasAppSettingsPropertyOutputReference
	// Experimental.
	CanvasAppSettingsInput() *AwsSagemakerUserProfile_CanvasAppSettingsProperty
	// Experimental.
	CodeEditorAppSettings() AwsSagemakerUserProfile_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *AwsSagemakerUserProfile_CodeEditorAppSettingsProperty
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
	CustomFileSystemConfig() AwsSagemakerUserProfile_CustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() AwsSagemakerUserProfile_CustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *AwsSagemakerUserProfile_CustomPosixUserConfigProperty
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
	InternalValue() *AwsSagemakerUserProfile_UserSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerUserProfile_UserSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() AwsSagemakerUserProfile_JupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *AwsSagemakerUserProfile_JupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() AwsSagemakerUserProfile_JupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *AwsSagemakerUserProfile_JupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() AwsSagemakerUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *AwsSagemakerUserProfile_KernelGatewayAppSettingsProperty
	// Experimental.
	RSessionAppSettings() AwsSagemakerUserProfile_RSessionAppSettingsPropertyOutputReference
	// Experimental.
	RSessionAppSettingsInput() *AwsSagemakerUserProfile_RSessionAppSettingsProperty
	// Experimental.
	RStudioServerProAppSettings() AwsSagemakerUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProAppSettingsInput() *AwsSagemakerUserProfile_RStudioServerProAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharingSettings() AwsSagemakerUserProfile_SharingSettingsPropertyOutputReference
	// Experimental.
	SharingSettingsInput() *AwsSagemakerUserProfile_SharingSettingsProperty
	// Experimental.
	SpaceStorageSettings() AwsSagemakerUserProfile_SpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *AwsSagemakerUserProfile_SpaceStorageSettingsProperty
	// Experimental.
	StudioWebPortal() *string
	// Experimental.
	SetStudioWebPortal(val *string)
	// Experimental.
	StudioWebPortalInput() *string
	// Experimental.
	StudioWebPortalSettings() AwsSagemakerUserProfile_StudioWebPortalSettingsPropertyOutputReference
	// Experimental.
	StudioWebPortalSettingsInput() *AwsSagemakerUserProfile_StudioWebPortalSettingsProperty
	// Experimental.
	TensorBoardAppSettings() AwsSagemakerUserProfile_TensorBoardAppSettingsPropertyOutputReference
	// Experimental.
	TensorBoardAppSettingsInput() *AwsSagemakerUserProfile_TensorBoardAppSettingsProperty
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
	PutCanvasAppSettings(value *AwsSagemakerUserProfile_CanvasAppSettingsProperty)
	// Experimental.
	PutCodeEditorAppSettings(value *AwsSagemakerUserProfile_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *AwsSagemakerUserProfile_CustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *AwsSagemakerUserProfile_JupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *AwsSagemakerUserProfile_JupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *AwsSagemakerUserProfile_KernelGatewayAppSettingsProperty)
	// Experimental.
	PutRSessionAppSettings(value *AwsSagemakerUserProfile_RSessionAppSettingsProperty)
	// Experimental.
	PutRStudioServerProAppSettings(value *AwsSagemakerUserProfile_RStudioServerProAppSettingsProperty)
	// Experimental.
	PutSharingSettings(value *AwsSagemakerUserProfile_SharingSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *AwsSagemakerUserProfile_SpaceStorageSettingsProperty)
	// Experimental.
	PutStudioWebPortalSettings(value *AwsSagemakerUserProfile_StudioWebPortalSettingsProperty)
	// Experimental.
	PutTensorBoardAppSettings(value *AwsSagemakerUserProfile_TensorBoardAppSettingsProperty)
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

// The jsii proxy struct for AwsSagemakerUserProfile_UserSettingsPropertyOutputReference
type jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettings() AwsSagemakerUserProfile_CanvasAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_CanvasAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettingsInput() *AwsSagemakerUserProfile_CanvasAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettings() AwsSagemakerUserProfile_CodeEditorAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *AwsSagemakerUserProfile_CodeEditorAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfig() AwsSagemakerUserProfile_CustomFileSystemConfigPropertyList {
	var returns AwsSagemakerUserProfile_CustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfig() AwsSagemakerUserProfile_CustomPosixUserConfigPropertyOutputReference {
	var returns AwsSagemakerUserProfile_CustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfigInput() *AwsSagemakerUserProfile_CustomPosixUserConfigProperty {
	var returns *AwsSagemakerUserProfile_CustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) InternalValue() *AwsSagemakerUserProfile_UserSettingsProperty {
	var returns *AwsSagemakerUserProfile_UserSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettings() AwsSagemakerUserProfile_JupyterLabAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_JupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *AwsSagemakerUserProfile_JupyterLabAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettings() AwsSagemakerUserProfile_JupyterServerAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_JupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *AwsSagemakerUserProfile_JupyterServerAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_JupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettings() AwsSagemakerUserProfile_KernelGatewayAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *AwsSagemakerUserProfile_KernelGatewayAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettings() AwsSagemakerUserProfile_RSessionAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_RSessionAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettingsInput() *AwsSagemakerUserProfile_RSessionAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_RSessionAppSettingsProperty
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettings() AwsSagemakerUserProfile_RStudioServerProAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettingsInput() *AwsSagemakerUserProfile_RStudioServerProAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_RStudioServerProAppSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SharingSettings() AwsSagemakerUserProfile_SharingSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_SharingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SharingSettingsInput() *AwsSagemakerUserProfile_SharingSettingsProperty {
	var returns *AwsSagemakerUserProfile_SharingSettingsProperty
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettings() AwsSagemakerUserProfile_SpaceStorageSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_SpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettingsInput() *AwsSagemakerUserProfile_SpaceStorageSettingsProperty {
	var returns *AwsSagemakerUserProfile_SpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettings() AwsSagemakerUserProfile_StudioWebPortalSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_StudioWebPortalSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettingsInput() *AwsSagemakerUserProfile_StudioWebPortalSettingsProperty {
	var returns *AwsSagemakerUserProfile_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettings() AwsSagemakerUserProfile_TensorBoardAppSettingsPropertyOutputReference {
	var returns AwsSagemakerUserProfile_TensorBoardAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettingsInput() *AwsSagemakerUserProfile_TensorBoardAppSettingsProperty {
	var returns *AwsSagemakerUserProfile_TensorBoardAppSettingsProperty
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerUserProfile_UserSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerUserProfile_UserSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerUserProfile_UserSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerUserProfile_UserSettingsPropertyOutputReference_Override(a AwsSagemakerUserProfile_UserSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetInternalValue(val *AwsSagemakerUserProfile_UserSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutCanvasAppSettings(value *AwsSagemakerUserProfile_CanvasAppSettingsProperty) {
	if err := a.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *AwsSagemakerUserProfile_CodeEditorAppSettingsProperty) {
	if err := a.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := a.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *AwsSagemakerUserProfile_CustomPosixUserConfigProperty) {
	if err := a.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *AwsSagemakerUserProfile_JupyterLabAppSettingsProperty) {
	if err := a.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *AwsSagemakerUserProfile_JupyterServerAppSettingsProperty) {
	if err := a.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *AwsSagemakerUserProfile_KernelGatewayAppSettingsProperty) {
	if err := a.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutRSessionAppSettings(value *AwsSagemakerUserProfile_RSessionAppSettingsProperty) {
	if err := a.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutRStudioServerProAppSettings(value *AwsSagemakerUserProfile_RStudioServerProAppSettingsProperty) {
	if err := a.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutSharingSettings(value *AwsSagemakerUserProfile_SharingSettingsProperty) {
	if err := a.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutSpaceStorageSettings(value *AwsSagemakerUserProfile_SpaceStorageSettingsProperty) {
	if err := a.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutStudioWebPortalSettings(value *AwsSagemakerUserProfile_StudioWebPortalSettingsProperty) {
	if err := a.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) PutTensorBoardAppSettings(value *AwsSagemakerUserProfile_TensorBoardAppSettingsProperty) {
	if err := a.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerUserProfile_UserSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

