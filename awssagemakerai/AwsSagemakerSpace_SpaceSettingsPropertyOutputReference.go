package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerSpace_SpaceSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppType() *string
	// Experimental.
	SetAppType(val *string)
	// Experimental.
	AppTypeInput() *string
	// Experimental.
	CodeEditorAppSettings() AwsSagemakerSpace_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *AwsSagemakerSpace_CodeEditorAppSettingsProperty
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
	CustomFileSystem() AwsSagemakerSpace_CustomFileSystemPropertyList
	// Experimental.
	CustomFileSystemInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerSpace_SpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerSpace_SpaceSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() AwsSagemakerSpace_JupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *AwsSagemakerSpace_JupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() AwsSagemakerSpace_JupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *AwsSagemakerSpace_JupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() AwsSagemakerSpace_KernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *AwsSagemakerSpace_KernelGatewayAppSettingsProperty
	// Experimental.
	SpaceStorageSettings() AwsSagemakerSpace_SpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *AwsSagemakerSpace_SpaceStorageSettingsProperty
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
	PutCodeEditorAppSettings(value *AwsSagemakerSpace_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystem(value interface{})
	// Experimental.
	PutJupyterLabAppSettings(value *AwsSagemakerSpace_JupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *AwsSagemakerSpace_JupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *AwsSagemakerSpace_KernelGatewayAppSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *AwsSagemakerSpace_SpaceStorageSettingsProperty)
	// Experimental.
	ResetAppType()
	// Experimental.
	ResetCodeEditorAppSettings()
	// Experimental.
	ResetCustomFileSystem()
	// Experimental.
	ResetJupyterLabAppSettings()
	// Experimental.
	ResetJupyterServerAppSettings()
	// Experimental.
	ResetKernelGatewayAppSettings()
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

// The jsii proxy struct for AwsSagemakerSpace_SpaceSettingsPropertyOutputReference
type jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) AppType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) AppTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) CodeEditorAppSettings() AwsSagemakerSpace_CodeEditorAppSettingsPropertyOutputReference {
	var returns AwsSagemakerSpace_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *AwsSagemakerSpace_CodeEditorAppSettingsProperty {
	var returns *AwsSagemakerSpace_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) CustomFileSystem() AwsSagemakerSpace_CustomFileSystemPropertyList {
	var returns AwsSagemakerSpace_CustomFileSystemPropertyList
	_jsii_.Get(
		j,
		"customFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) CustomFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) InternalValue() *AwsSagemakerSpace_SpaceSettingsProperty {
	var returns *AwsSagemakerSpace_SpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) JupyterLabAppSettings() AwsSagemakerSpace_JupyterLabAppSettingsPropertyOutputReference {
	var returns AwsSagemakerSpace_JupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *AwsSagemakerSpace_JupyterLabAppSettingsProperty {
	var returns *AwsSagemakerSpace_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) JupyterServerAppSettings() AwsSagemakerSpace_JupyterServerAppSettingsPropertyOutputReference {
	var returns AwsSagemakerSpace_JupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *AwsSagemakerSpace_JupyterServerAppSettingsProperty {
	var returns *AwsSagemakerSpace_JupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) KernelGatewayAppSettings() AwsSagemakerSpace_KernelGatewayAppSettingsPropertyOutputReference {
	var returns AwsSagemakerSpace_KernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *AwsSagemakerSpace_KernelGatewayAppSettingsProperty {
	var returns *AwsSagemakerSpace_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) SpaceStorageSettings() AwsSagemakerSpace_SpaceStorageSettingsPropertyOutputReference {
	var returns AwsSagemakerSpace_SpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) SpaceStorageSettingsInput() *AwsSagemakerSpace_SpaceStorageSettingsProperty {
	var returns *AwsSagemakerSpace_SpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerSpace_SpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerSpace_SpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerSpace_SpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerSpace.SpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerSpace_SpaceSettingsPropertyOutputReference_Override(a AwsSagemakerSpace_SpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerSpace.SpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetAppType(val *string) {
	if err := j.validateSetAppTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appType",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetInternalValue(val *AwsSagemakerSpace_SpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *AwsSagemakerSpace_CodeEditorAppSettingsProperty) {
	if err := a.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutCustomFileSystem(value interface{}) {
	if err := a.validatePutCustomFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomFileSystem",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *AwsSagemakerSpace_JupyterLabAppSettingsProperty) {
	if err := a.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *AwsSagemakerSpace_JupyterServerAppSettingsProperty) {
	if err := a.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *AwsSagemakerSpace_KernelGatewayAppSettingsProperty) {
	if err := a.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) PutSpaceStorageSettings(value *AwsSagemakerSpace_SpaceStorageSettingsProperty) {
	if err := a.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetAppType() {
	_jsii_.InvokeVoid(
		a,
		"resetAppType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetCustomFileSystem() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomFileSystem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerSpace_SpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

