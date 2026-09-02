package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DefaultUserSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoMountHomeEfs() *string
	// Experimental.
	SetAutoMountHomeEfs(val *string)
	// Experimental.
	AutoMountHomeEfsInput() *string
	// Experimental.
	CanvasAppSettings() TfDomain_CanvasAppSettingsPropertyOutputReference
	// Experimental.
	CanvasAppSettingsInput() *TfDomain_CanvasAppSettingsProperty
	// Experimental.
	CodeEditorAppSettings() TfDomain_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *TfDomain_CodeEditorAppSettingsProperty
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
	CustomFileSystemConfig() TfDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() TfDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *TfDomain_DefaultUserSettingsCustomPosixUserConfigProperty
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
	InternalValue() *TfDomain_DefaultUserSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DefaultUserSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() TfDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	// Experimental.
	RSessionAppSettings() TfDomain_RSessionAppSettingsPropertyOutputReference
	// Experimental.
	RSessionAppSettingsInput() *TfDomain_RSessionAppSettingsProperty
	// Experimental.
	RStudioServerProAppSettings() TfDomain_RStudioServerProAppSettingsPropertyOutputReference
	// Experimental.
	RStudioServerProAppSettingsInput() *TfDomain_RStudioServerProAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SharingSettings() TfDomain_SharingSettingsPropertyOutputReference
	// Experimental.
	SharingSettingsInput() *TfDomain_SharingSettingsProperty
	// Experimental.
	SpaceStorageSettings() TfDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *TfDomain_DefaultUserSettingsSpaceStorageSettingsProperty
	// Experimental.
	StudioWebPortal() *string
	// Experimental.
	SetStudioWebPortal(val *string)
	// Experimental.
	StudioWebPortalInput() *string
	// Experimental.
	StudioWebPortalSettings() TfDomain_StudioWebPortalSettingsPropertyOutputReference
	// Experimental.
	StudioWebPortalSettingsInput() *TfDomain_StudioWebPortalSettingsProperty
	// Experimental.
	TensorBoardAppSettings() TfDomain_TensorBoardAppSettingsPropertyOutputReference
	// Experimental.
	TensorBoardAppSettingsInput() *TfDomain_TensorBoardAppSettingsProperty
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
	PutCanvasAppSettings(value *TfDomain_CanvasAppSettingsProperty)
	// Experimental.
	PutCodeEditorAppSettings(value *TfDomain_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *TfDomain_DefaultUserSettingsCustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty)
	// Experimental.
	PutRSessionAppSettings(value *TfDomain_RSessionAppSettingsProperty)
	// Experimental.
	PutRStudioServerProAppSettings(value *TfDomain_RStudioServerProAppSettingsProperty)
	// Experimental.
	PutSharingSettings(value *TfDomain_SharingSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *TfDomain_DefaultUserSettingsSpaceStorageSettingsProperty)
	// Experimental.
	PutStudioWebPortalSettings(value *TfDomain_StudioWebPortalSettingsProperty)
	// Experimental.
	PutTensorBoardAppSettings(value *TfDomain_TensorBoardAppSettingsProperty)
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

// The jsii proxy struct for TfDomain_DefaultUserSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) AutoMountHomeEfs() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) AutoMountHomeEfsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoMountHomeEfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CanvasAppSettings() TfDomain_CanvasAppSettingsPropertyOutputReference {
	var returns TfDomain_CanvasAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"canvasAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CanvasAppSettingsInput() *TfDomain_CanvasAppSettingsProperty {
	var returns *TfDomain_CanvasAppSettingsProperty
	_jsii_.Get(
		j,
		"canvasAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CodeEditorAppSettings() TfDomain_CodeEditorAppSettingsPropertyOutputReference {
	var returns TfDomain_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *TfDomain_CodeEditorAppSettingsProperty {
	var returns *TfDomain_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CustomFileSystemConfig() TfDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList {
	var returns TfDomain_DefaultUserSettingsCustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CustomPosixUserConfig() TfDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsCustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) CustomPosixUserConfigInput() *TfDomain_DefaultUserSettingsCustomPosixUserConfigProperty {
	var returns *TfDomain_DefaultUserSettingsCustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) DefaultLandingUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) DefaultLandingUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultLandingUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) InternalValue() *TfDomain_DefaultUserSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) JupyterLabAppSettings() TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsJupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) JupyterServerAppSettings() TfDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsJupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) KernelGatewayAppSettings() TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsKernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) RSessionAppSettings() TfDomain_RSessionAppSettingsPropertyOutputReference {
	var returns TfDomain_RSessionAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rSessionAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) RSessionAppSettingsInput() *TfDomain_RSessionAppSettingsProperty {
	var returns *TfDomain_RSessionAppSettingsProperty
	_jsii_.Get(
		j,
		"rSessionAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) RStudioServerProAppSettings() TfDomain_RStudioServerProAppSettingsPropertyOutputReference {
	var returns TfDomain_RStudioServerProAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rStudioServerProAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) RStudioServerProAppSettingsInput() *TfDomain_RStudioServerProAppSettingsProperty {
	var returns *TfDomain_RStudioServerProAppSettingsProperty
	_jsii_.Get(
		j,
		"rStudioServerProAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SharingSettings() TfDomain_SharingSettingsPropertyOutputReference {
	var returns TfDomain_SharingSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"sharingSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SharingSettingsInput() *TfDomain_SharingSettingsProperty {
	var returns *TfDomain_SharingSettingsProperty
	_jsii_.Get(
		j,
		"sharingSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SpaceStorageSettings() TfDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference {
	var returns TfDomain_DefaultUserSettingsSpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) SpaceStorageSettingsInput() *TfDomain_DefaultUserSettingsSpaceStorageSettingsProperty {
	var returns *TfDomain_DefaultUserSettingsSpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortal() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"studioWebPortalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalSettings() TfDomain_StudioWebPortalSettingsPropertyOutputReference {
	var returns TfDomain_StudioWebPortalSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"studioWebPortalSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) StudioWebPortalSettingsInput() *TfDomain_StudioWebPortalSettingsProperty {
	var returns *TfDomain_StudioWebPortalSettingsProperty
	_jsii_.Get(
		j,
		"studioWebPortalSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) TensorBoardAppSettings() TfDomain_TensorBoardAppSettingsPropertyOutputReference {
	var returns TfDomain_TensorBoardAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"tensorBoardAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) TensorBoardAppSettingsInput() *TfDomain_TensorBoardAppSettingsProperty {
	var returns *TfDomain_TensorBoardAppSettingsProperty
	_jsii_.Get(
		j,
		"tensorBoardAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DefaultUserSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DefaultUserSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DefaultUserSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DefaultUserSettingsPropertyOutputReference_Override(t TfDomain_DefaultUserSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultUserSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetAutoMountHomeEfs(val *string) {
	if err := j.validateSetAutoMountHomeEfsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoMountHomeEfs",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetDefaultLandingUri(val *string) {
	if err := j.validateSetDefaultLandingUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLandingUri",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DefaultUserSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetStudioWebPortal(val *string) {
	if err := j.validateSetStudioWebPortalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"studioWebPortal",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutCanvasAppSettings(value *TfDomain_CanvasAppSettingsProperty) {
	if err := t.validatePutCanvasAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCanvasAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *TfDomain_CodeEditorAppSettingsProperty) {
	if err := t.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := t.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *TfDomain_DefaultUserSettingsCustomPosixUserConfigProperty) {
	if err := t.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *TfDomain_DefaultUserSettingsJupyterLabAppSettingsProperty) {
	if err := t.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *TfDomain_DefaultUserSettingsJupyterServerAppSettingsProperty) {
	if err := t.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *TfDomain_DefaultUserSettingsKernelGatewayAppSettingsProperty) {
	if err := t.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutRSessionAppSettings(value *TfDomain_RSessionAppSettingsProperty) {
	if err := t.validatePutRSessionAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRSessionAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutRStudioServerProAppSettings(value *TfDomain_RStudioServerProAppSettingsProperty) {
	if err := t.validatePutRStudioServerProAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRStudioServerProAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutSharingSettings(value *TfDomain_SharingSettingsProperty) {
	if err := t.validatePutSharingSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSharingSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutSpaceStorageSettings(value *TfDomain_DefaultUserSettingsSpaceStorageSettingsProperty) {
	if err := t.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutStudioWebPortalSettings(value *TfDomain_StudioWebPortalSettingsProperty) {
	if err := t.validatePutStudioWebPortalSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putStudioWebPortalSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) PutTensorBoardAppSettings(value *TfDomain_TensorBoardAppSettingsProperty) {
	if err := t.validatePutTensorBoardAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTensorBoardAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetAutoMountHomeEfs() {
	_jsii_.InvokeVoid(
		t,
		"resetAutoMountHomeEfs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetCanvasAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCanvasAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetDefaultLandingUri() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultLandingUri",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetRSessionAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRSessionAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetRStudioServerProAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetRStudioServerProAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetSharingSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSharingSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetStudioWebPortal() {
	_jsii_.InvokeVoid(
		t,
		"resetStudioWebPortal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetStudioWebPortalSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetStudioWebPortalSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ResetTensorBoardAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetTensorBoardAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultUserSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

