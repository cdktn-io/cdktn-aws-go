package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultUserSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMountHomeEfs() *string
	// Experimental.
	SetAutoMountHomeEfs(val *string)
	// Experimental.
	AutoMountHomeEfsInput() *string
	// Experimental.
	CanvasAppSettings() AwsDomain_CanvasAppSettingsPropertyOutputReference
	// Experimental.
	CanvasAppSettingsInput() *AwsDomain_CanvasAppSettingsProperty
	// Experimental.
	CodeEditorAppSettings() AwsDomain_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *AwsDomain_CodeEditorAppSettingsProperty
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
	CustomFileSystemConfig() AwsDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() AwsDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty
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
	InternalValue() *AwsDomain_DefaultUserSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultUserSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() AwsDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() AwsDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	// Experimental.
	RSessionAppSettings() AwsDomain_RSessionAppSettingsPropertyOutputReference
	// Experimental.
	RSessionAppSettingsInput() *AwsDomain_RSessionAppSettingsProperty
	// Experimental.
	RStudioServerProAppSettings() AwsDomain_RStudioServerProAppSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProAppSettingsInput() *AwsDomain_RStudioServerProAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharingSettings() AwsDomain_SharingSettingsPropertyOutputReference
	// Experimental.
	SharingSettingsInput() *AwsDomain_SharingSettingsProperty
	// Experimental.
	SpaceStorageSettings() AwsDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty
	// Experimental.
	StudioWebPortal() *string
	// Experimental.
	SetStudioWebPortal(val *string)
	// Experimental.
	StudioWebPortalInput() *string
	// Experimental.
	StudioWebPortalSettings() AwsDomain_StudioWebPortalSettingsPropertyOutputReference
	// Experimental.
	StudioWebPortalSettingsInput() *AwsDomain_StudioWebPortalSettingsProperty
	// Experimental.
	TensorBoardAppSettings() AwsDomain_TensorBoardAppSettingsPropertyOutputReference
	// Experimental.
	TensorBoardAppSettingsInput() *AwsDomain_TensorBoardAppSettingsProperty
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
	PutCanvasAppSettings(value *AwsDomain_CanvasAppSettingsProperty)
	// Experimental.
	PutCodeEditorAppSettings(value *AwsDomain_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty)
	// Experimental.
	PutRSessionAppSettings(value *AwsDomain_RSessionAppSettingsProperty)
	// Experimental.
	PutRStudioServerProAppSettings(value *AwsDomain_RStudioServerProAppSettingsProperty)
	// Experimental.
	PutSharingSettings(value *AwsDomain_SharingSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty)
	// Experimental.
	PutStudioWebPortalSettings(value *AwsDomain_StudioWebPortalSettingsProperty)
	// Experimental.
	PutTensorBoardAppSettings(value *AwsDomain_TensorBoardAppSettingsProperty)
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

// The jsii proxy struct for AwsDomain_DefaultUserSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CanvasAppSettings() AwsDomain_CanvasAppSettingsPropertyOutputReference {
	var returns AwsDomain_CanvasAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CanvasAppSettingsInput() *AwsDomain_CanvasAppSettingsProperty {
	var returns *AwsDomain_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CodeEditorAppSettings() AwsDomain_CodeEditorAppSettingsPropertyOutputReference {
	var returns AwsDomain_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *AwsDomain_CodeEditorAppSettingsProperty {
	var returns *AwsDomain_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CustomFileSystemConfig() AwsDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList {
	var returns AwsDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CustomPosixUserConfig() AwsDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference {
	var returns AwsDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) CustomPosixUserConfigInput() *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty {
	var returns *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) InternalValue() *AwsDomain_DefaultUserSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) JupyterLabAppSettings() AwsDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) JupyterServerAppSettings() AwsDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) KernelGatewayAppSettings() AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) RSessionAppSettings() AwsDomain_RSessionAppSettingsPropertyOutputReference {
	var returns AwsDomain_RSessionAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) RSessionAppSettingsInput() *AwsDomain_RSessionAppSettingsProperty {
	var returns *AwsDomain_RSessionAppSettingsProperty
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) RStudioServerProAppSettings() AwsDomain_RStudioServerProAppSettingsPropertyOutputReference {
	var returns AwsDomain_RStudioServerProAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) RStudioServerProAppSettingsInput() *AwsDomain_RStudioServerProAppSettingsProperty {
	var returns *AwsDomain_RStudioServerProAppSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SharingSettings() AwsDomain_SharingSettingsPropertyOutputReference {
	var returns AwsDomain_SharingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SharingSettingsInput() *AwsDomain_SharingSettingsProperty {
	var returns *AwsDomain_SharingSettingsProperty
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SpaceStorageSettings() AwsDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) SpaceStorageSettingsInput() *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty {
	var returns *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalSettings() AwsDomain_StudioWebPortalSettingsPropertyOutputReference {
	var returns AwsDomain_StudioWebPortalSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalSettingsInput() *AwsDomain_StudioWebPortalSettingsProperty {
	var returns *AwsDomain_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) TensorBoardAppSettings() AwsDomain_TensorBoardAppSettingsPropertyOutputReference {
	var returns AwsDomain_TensorBoardAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) TensorBoardAppSettingsInput() *AwsDomain_TensorBoardAppSettingsProperty {
	var returns *AwsDomain_TensorBoardAppSettingsProperty
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultUserSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultUserSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultUserSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultUserSettingsPropertyOutputReference_Override(a AwsDomain_DefaultUserSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultUserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultUserSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutCanvasAppSettings(value *AwsDomain_CanvasAppSettingsProperty) {
	if err := a.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *AwsDomain_CodeEditorAppSettingsProperty) {
	if err := a.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := a.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *AwsDomain_DefaultUserSettingsCustomPosixUserConfigProperty) {
	if err := a.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *AwsDomain_DefaultUserSettingsJupyterLabAppSettingsProperty) {
	if err := a.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *AwsDomain_DefaultUserSettingsJupyterServerAppSettingsProperty) {
	if err := a.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *AwsDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty) {
	if err := a.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutRSessionAppSettings(value *AwsDomain_RSessionAppSettingsProperty) {
	if err := a.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutRStudioServerProAppSettings(value *AwsDomain_RStudioServerProAppSettingsProperty) {
	if err := a.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutSharingSettings(value *AwsDomain_SharingSettingsProperty) {
	if err := a.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutSpaceStorageSettings(value *AwsDomain_DefaultUserSettingsSpaceStorageSettingsProperty) {
	if err := a.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutStudioWebPortalSettings(value *AwsDomain_StudioWebPortalSettingsProperty) {
	if err := a.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) PutTensorBoardAppSettings(value *AwsDomain_TensorBoardAppSettingsProperty) {
	if err := a.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultUserSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

