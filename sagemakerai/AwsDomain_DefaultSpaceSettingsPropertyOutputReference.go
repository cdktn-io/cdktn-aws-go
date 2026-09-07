package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDomain_DefaultSpaceSettingsPropertyOutputReference interface {
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
	CustomFileSystemConfig() AwsDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDomain_DefaultSpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsDomain_DefaultSpaceSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SpaceStorageSettings() AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty
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
	PutCustomFileSystemConfig(value interface{})
	// Experimental.
	PutCustomPosixUserConfig(value *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty)
	// Experimental.
	ResetCustomFileSystemConfig()
	// Experimental.
	ResetCustomPosixUserConfig()
	// Experimental.
	ResetJupyterLabAppSettings()
	// Experimental.
	ResetJupyterServerAppSettings()
	// Experimental.
	ResetKernelGatewayAppSettings()
	// Experimental.
	ResetSecurityGroups()
	// Experimental.
	ResetSpaceStorageSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDomain_DefaultSpaceSettingsPropertyOutputReference
type jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) CustomFileSystemConfig() AwsDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList {
	var returns AwsDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) CustomPosixUserConfig() AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) CustomPosixUserConfigInput() *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty {
	var returns *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) InternalValue() *AwsDomain_DefaultSpaceSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterLabAppSettings() AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterServerAppSettings() AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) KernelGatewayAppSettings() AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) SpaceStorageSettings() AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference {
	var returns AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) SpaceStorageSettingsInput() *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty {
	var returns *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDomain_DefaultSpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDomain_DefaultSpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDomain_DefaultSpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDomain_DefaultSpaceSettingsPropertyOutputReference_Override(a AwsDomain_DefaultSpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsDomain.DefaultSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetInternalValue(val *AwsDomain_DefaultSpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := a.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *AwsDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty) {
	if err := a.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *AwsDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty) {
	if err := a.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *AwsDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty) {
	if err := a.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *AwsDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty) {
	if err := a.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) PutSpaceStorageSettings(value *AwsDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty) {
	if err := a.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDomain_DefaultSpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

