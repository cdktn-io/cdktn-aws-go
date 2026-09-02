package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfUserProfile_UserSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMountHomeEfs() *string
	// Experimental.
	SetAutoMountHomeEfs(val *string)
	// Experimental.
	AutoMountHomeEfsInput() *string
	// Experimental.
	CanvasAppSettings() TfUserProfile_CanvasAppSettingsPropertyOutputReference
	// Experimental.
	CanvasAppSettingsInput() *TfUserProfile_CanvasAppSettingsProperty
	// Experimental.
	CodeEditorAppSettings() TfUserProfile_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *TfUserProfile_CodeEditorAppSettingsProperty
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
	CustomFileSystemConfig() TfUserProfile_CustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() TfUserProfile_CustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *TfUserProfile_CustomPosixUserConfigProperty
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
	InternalValue() *TfUserProfile_UserSettingsProperty
	// Experimental.
	SetInternalValue(val *TfUserProfile_UserSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() TfUserProfile_JupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *TfUserProfile_JupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() TfUserProfile_JupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *TfUserProfile_JupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *TfUserProfile_KernelGatewayAppSettingsProperty
	// Experimental.
	RSessionAppSettings() TfUserProfile_RSessionAppSettingsPropertyOutputReference
	// Experimental.
	RSessionAppSettingsInput() *TfUserProfile_RSessionAppSettingsProperty
	// Experimental.
	RStudioServerProAppSettings() TfUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProAppSettingsInput() *TfUserProfile_RStudioServerProAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharingSettings() TfUserProfile_SharingSettingsPropertyOutputReference
	// Experimental.
	SharingSettingsInput() *TfUserProfile_SharingSettingsProperty
	// Experimental.
	SpaceStorageSettings() TfUserProfile_SpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *TfUserProfile_SpaceStorageSettingsProperty
	// Experimental.
	StudioWebPortal() *string
	// Experimental.
	SetStudioWebPortal(val *string)
	// Experimental.
	StudioWebPortalInput() *string
	// Experimental.
	StudioWebPortalSettings() TfUserProfile_StudioWebPortalSettingsPropertyOutputReference
	// Experimental.
	StudioWebPortalSettingsInput() *TfUserProfile_StudioWebPortalSettingsProperty
	// Experimental.
	TensorBoardAppSettings() TfUserProfile_TensorBoardAppSettingsPropertyOutputReference
	// Experimental.
	TensorBoardAppSettingsInput() *TfUserProfile_TensorBoardAppSettingsProperty
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
	PutCanvasAppSettings(value *TfUserProfile_CanvasAppSettingsProperty)
	// Experimental.
	PutCodeEditorAppSettings(value *TfUserProfile_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *TfUserProfile_CustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *TfUserProfile_JupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *TfUserProfile_JupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *TfUserProfile_KernelGatewayAppSettingsProperty)
	// Experimental.
	PutRSessionAppSettings(value *TfUserProfile_RSessionAppSettingsProperty)
	// Experimental.
	PutRStudioServerProAppSettings(value *TfUserProfile_RStudioServerProAppSettingsProperty)
	// Experimental.
	PutSharingSettings(value *TfUserProfile_SharingSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *TfUserProfile_SpaceStorageSettingsProperty)
	// Experimental.
	PutStudioWebPortalSettings(value *TfUserProfile_StudioWebPortalSettingsProperty)
	// Experimental.
	PutTensorBoardAppSettings(value *TfUserProfile_TensorBoardAppSettingsProperty)
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

// The jsii proxy struct for TfUserProfile_UserSettingsPropertyOutputReference
type jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettings() TfUserProfile_CanvasAppSettingsPropertyOutputReference {
	var returns TfUserProfile_CanvasAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CanvasAppSettingsInput() *TfUserProfile_CanvasAppSettingsProperty {
	var returns *TfUserProfile_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettings() TfUserProfile_CodeEditorAppSettingsPropertyOutputReference {
	var returns TfUserProfile_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *TfUserProfile_CodeEditorAppSettingsProperty {
	var returns *TfUserProfile_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfig() TfUserProfile_CustomFileSystemConfigPropertyList {
	var returns TfUserProfile_CustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfig() TfUserProfile_CustomPosixUserConfigPropertyOutputReference {
	var returns TfUserProfile_CustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) CustomPosixUserConfigInput() *TfUserProfile_CustomPosixUserConfigProperty {
	var returns *TfUserProfile_CustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) InternalValue() *TfUserProfile_UserSettingsProperty {
	var returns *TfUserProfile_UserSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettings() TfUserProfile_JupyterLabAppSettingsPropertyOutputReference {
	var returns TfUserProfile_JupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *TfUserProfile_JupyterLabAppSettingsProperty {
	var returns *TfUserProfile_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettings() TfUserProfile_JupyterServerAppSettingsPropertyOutputReference {
	var returns TfUserProfile_JupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *TfUserProfile_JupyterServerAppSettingsProperty {
	var returns *TfUserProfile_JupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettings() TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference {
	var returns TfUserProfile_KernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *TfUserProfile_KernelGatewayAppSettingsProperty {
	var returns *TfUserProfile_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettings() TfUserProfile_RSessionAppSettingsPropertyOutputReference {
	var returns TfUserProfile_RSessionAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) RSessionAppSettingsInput() *TfUserProfile_RSessionAppSettingsProperty {
	var returns *TfUserProfile_RSessionAppSettingsProperty
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettings() TfUserProfile_RStudioServerProAppSettingsPropertyOutputReference {
	var returns TfUserProfile_RStudioServerProAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) RStudioServerProAppSettingsInput() *TfUserProfile_RStudioServerProAppSettingsProperty {
	var returns *TfUserProfile_RStudioServerProAppSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SharingSettings() TfUserProfile_SharingSettingsPropertyOutputReference {
	var returns TfUserProfile_SharingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SharingSettingsInput() *TfUserProfile_SharingSettingsProperty {
	var returns *TfUserProfile_SharingSettingsProperty
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettings() TfUserProfile_SpaceStorageSettingsPropertyOutputReference {
	var returns TfUserProfile_SpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) SpaceStorageSettingsInput() *TfUserProfile_SpaceStorageSettingsProperty {
	var returns *TfUserProfile_SpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettings() TfUserProfile_StudioWebPortalSettingsPropertyOutputReference {
	var returns TfUserProfile_StudioWebPortalSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) StudioWebPortalSettingsInput() *TfUserProfile_StudioWebPortalSettingsProperty {
	var returns *TfUserProfile_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettings() TfUserProfile_TensorBoardAppSettingsPropertyOutputReference {
	var returns TfUserProfile_TensorBoardAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) TensorBoardAppSettingsInput() *TfUserProfile_TensorBoardAppSettingsProperty {
	var returns *TfUserProfile_TensorBoardAppSettingsProperty
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfUserProfile_UserSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfUserProfile_UserSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfUserProfile_UserSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfUserProfile_UserSettingsPropertyOutputReference_Override(t TfUserProfile_UserSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfUserProfile.UserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetInternalValue(val *TfUserProfile_UserSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutCanvasAppSettings(value *TfUserProfile_CanvasAppSettingsProperty) {
	if err := t.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *TfUserProfile_CodeEditorAppSettingsProperty) {
	if err := t.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := t.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *TfUserProfile_CustomPosixUserConfigProperty) {
	if err := t.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *TfUserProfile_JupyterLabAppSettingsProperty) {
	if err := t.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *TfUserProfile_JupyterServerAppSettingsProperty) {
	if err := t.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *TfUserProfile_KernelGatewayAppSettingsProperty) {
	if err := t.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutRSessionAppSettings(value *TfUserProfile_RSessionAppSettingsProperty) {
	if err := t.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutRStudioServerProAppSettings(value *TfUserProfile_RStudioServerProAppSettingsProperty) {
	if err := t.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutSharingSettings(value *TfUserProfile_SharingSettingsProperty) {
	if err := t.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutSpaceStorageSettings(value *TfUserProfile_SpaceStorageSettingsProperty) {
	if err := t.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutStudioWebPortalSettings(value *TfUserProfile_StudioWebPortalSettingsProperty) {
	if err := t.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) PutTensorBoardAppSettings(value *TfUserProfile_TensorBoardAppSettingsProperty) {
	if err := t.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		t,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfUserProfile_UserSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

