package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_DefaultSpaceSettingsPropertyOutputReference interface {
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
	CustomFileSystemConfig() TfDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList
	// Experimental.
	CustomFileSystemConfigInput() interface{}
	// Experimental.
	CustomPosixUserConfig() TfDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference
	// Experimental.
	CustomPosixUserConfigInput() *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty
	// Experimental.
	ExecutionRole() *string
	// Experimental.
	SetExecutionRole(val *string)
	// Experimental.
	ExecutionRoleInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDomain_DefaultSpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDomain_DefaultSpaceSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty
	// Experimental.
	SecurityGroups() *[]*string
	// Experimental.
	SetSecurityGroups(val *[]*string)
	// Experimental.
	SecurityGroupsInput() *[]*string
	// Experimental.
	SpaceStorageSettings() TfDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty
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
	PutCustomPosixUserConfig(value *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty)
	// Experimental.
	PutJupyterLabAppSettings(value *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty)
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

// The jsii proxy struct for TfDomain_DefaultSpaceSettingsPropertyOutputReference
type jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) CustomFileSystemConfig() TfDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList {
	var returns TfDomain_DefaultSpaceSettingsCustomFileSystemConfigPropertyList
	_jsii_.Get(
		j,
		"customFileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) CustomFileSystemConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) CustomPosixUserConfig() TfDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference {
	var returns TfDomain_DefaultSpaceSettingsCustomPosixUserConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"customPosixUserConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) CustomPosixUserConfigInput() *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty {
	var returns *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty
	_jsii_.Get(
		j,
		"customPosixUserConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ExecutionRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ExecutionRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"executionRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) InternalValue() *TfDomain_DefaultSpaceSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterLabAppSettings() TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterServerAppSettings() TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) KernelGatewayAppSettings() TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference {
	var returns TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) SecurityGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) SecurityGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"securityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) SpaceStorageSettings() TfDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference {
	var returns TfDomain_DefaultSpaceSettingsSpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) SpaceStorageSettingsInput() *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty {
	var returns *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_DefaultSpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDomain_DefaultSpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_DefaultSpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_DefaultSpaceSettingsPropertyOutputReference_Override(t TfDomain_DefaultSpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfDomain.DefaultSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetExecutionRole(val *string) {
	if err := j.validateSetExecutionRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"executionRole",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetInternalValue(val *TfDomain_DefaultSpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetSecurityGroups(val *[]*string) {
	if err := j.validateSetSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroups",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutCustomFileSystemConfig(value interface{}) {
	if err := t.validatePutCustomFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutCustomPosixUserConfig(value *TfDomain_DefaultSpaceSettingsCustomPosixUserConfigProperty) {
	if err := t.validatePutCustomPosixUserConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomPosixUserConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *TfDomain_DefaultSpaceSettingsJupyterLabAppSettingsProperty) {
	if err := t.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *TfDomain_DefaultSpaceSettingsJupyterServerAppSettingsProperty) {
	if err := t.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *TfDomain_DefaultSpaceSettingsKernelGatewayAppSettingsProperty) {
	if err := t.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) PutSpaceStorageSettings(value *TfDomain_DefaultSpaceSettingsSpaceStorageSettingsProperty) {
	if err := t.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetCustomFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetCustomPosixUserConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomPosixUserConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetSecurityGroups() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroups",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_DefaultSpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

