package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSpace_SpaceSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AppType() *string
	// Experimental.
	SetAppType(val *string)
	// Experimental.
	AppTypeInput() *string
	// Experimental.
	CodeEditorAppSettings() TfSpace_CodeEditorAppSettingsPropertyOutputReference
	// Experimental.
	CodeEditorAppSettingsInput() *TfSpace_CodeEditorAppSettingsProperty
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
	CustomFileSystem() TfSpace_CustomFileSystemPropertyList
	// Experimental.
	CustomFileSystemInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfSpace_SpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *TfSpace_SpaceSettingsProperty)
	// Experimental.
	JupyterLabAppSettings() TfSpace_JupyterLabAppSettingsPropertyOutputReference
	// Experimental.
	JupyterLabAppSettingsInput() *TfSpace_JupyterLabAppSettingsProperty
	// Experimental.
	JupyterServerAppSettings() TfSpace_JupyterServerAppSettingsPropertyOutputReference
	// Experimental.
	JupyterServerAppSettingsInput() *TfSpace_JupyterServerAppSettingsProperty
	// Experimental.
	KernelGatewayAppSettings() TfSpace_KernelGatewayAppSettingsPropertyOutputReference
	// Experimental.
	KernelGatewayAppSettingsInput() *TfSpace_KernelGatewayAppSettingsProperty
	// Experimental.
	SpaceStorageSettings() TfSpace_SpaceStorageSettingsPropertyOutputReference
	// Experimental.
	SpaceStorageSettingsInput() *TfSpace_SpaceStorageSettingsProperty
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
	PutCodeEditorAppSettings(value *TfSpace_CodeEditorAppSettingsProperty)
	// Experimental.
	PutCustomFileSystem(value interface{})
	// Experimental.
	PutJupyterLabAppSettings(value *TfSpace_JupyterLabAppSettingsProperty)
	// Experimental.
	PutJupyterServerAppSettings(value *TfSpace_JupyterServerAppSettingsProperty)
	// Experimental.
	PutKernelGatewayAppSettings(value *TfSpace_KernelGatewayAppSettingsProperty)
	// Experimental.
	PutSpaceStorageSettings(value *TfSpace_SpaceStorageSettingsProperty)
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

// The jsii proxy struct for TfSpace_SpaceSettingsPropertyOutputReference
type jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) AppType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) AppTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) CodeEditorAppSettings() TfSpace_CodeEditorAppSettingsPropertyOutputReference {
	var returns TfSpace_CodeEditorAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"codeEditorAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) CodeEditorAppSettingsInput() *TfSpace_CodeEditorAppSettingsProperty {
	var returns *TfSpace_CodeEditorAppSettingsProperty
	_jsii_.Get(
		j,
		"codeEditorAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) CustomFileSystem() TfSpace_CustomFileSystemPropertyList {
	var returns TfSpace_CustomFileSystemPropertyList
	_jsii_.Get(
		j,
		"customFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) CustomFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) InternalValue() *TfSpace_SpaceSettingsProperty {
	var returns *TfSpace_SpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) JupyterLabAppSettings() TfSpace_JupyterLabAppSettingsPropertyOutputReference {
	var returns TfSpace_JupyterLabAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterLabAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) JupyterLabAppSettingsInput() *TfSpace_JupyterLabAppSettingsProperty {
	var returns *TfSpace_JupyterLabAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterLabAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) JupyterServerAppSettings() TfSpace_JupyterServerAppSettingsPropertyOutputReference {
	var returns TfSpace_JupyterServerAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"jupyterServerAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) JupyterServerAppSettingsInput() *TfSpace_JupyterServerAppSettingsProperty {
	var returns *TfSpace_JupyterServerAppSettingsProperty
	_jsii_.Get(
		j,
		"jupyterServerAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) KernelGatewayAppSettings() TfSpace_KernelGatewayAppSettingsPropertyOutputReference {
	var returns TfSpace_KernelGatewayAppSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"kernelGatewayAppSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) KernelGatewayAppSettingsInput() *TfSpace_KernelGatewayAppSettingsProperty {
	var returns *TfSpace_KernelGatewayAppSettingsProperty
	_jsii_.Get(
		j,
		"kernelGatewayAppSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) SpaceStorageSettings() TfSpace_SpaceStorageSettingsPropertyOutputReference {
	var returns TfSpace_SpaceStorageSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"spaceStorageSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) SpaceStorageSettingsInput() *TfSpace_SpaceStorageSettingsProperty {
	var returns *TfSpace_SpaceStorageSettingsProperty
	_jsii_.Get(
		j,
		"spaceStorageSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfSpace_SpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfSpace_SpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfSpace_SpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfSpace.SpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfSpace_SpaceSettingsPropertyOutputReference_Override(t TfSpace_SpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfSpace.SpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetAppType(val *string) {
	if err := j.validateSetAppTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appType",
		val,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetInternalValue(val *TfSpace_SpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutCodeEditorAppSettings(value *TfSpace_CodeEditorAppSettingsProperty) {
	if err := t.validatePutCodeEditorAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCodeEditorAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutCustomFileSystem(value interface{}) {
	if err := t.validatePutCustomFileSystemParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomFileSystem",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutJupyterLabAppSettings(value *TfSpace_JupyterLabAppSettingsProperty) {
	if err := t.validatePutJupyterLabAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterLabAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutJupyterServerAppSettings(value *TfSpace_JupyterServerAppSettingsProperty) {
	if err := t.validatePutJupyterServerAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putJupyterServerAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutKernelGatewayAppSettings(value *TfSpace_KernelGatewayAppSettingsProperty) {
	if err := t.validatePutKernelGatewayAppSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKernelGatewayAppSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) PutSpaceStorageSettings(value *TfSpace_SpaceStorageSettingsProperty) {
	if err := t.validatePutSpaceStorageSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSpaceStorageSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetAppType() {
	_jsii_.InvokeVoid(
		t,
		"resetAppType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetCodeEditorAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetCodeEditorAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetCustomFileSystem() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomFileSystem",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetJupyterLabAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterLabAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetJupyterServerAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetJupyterServerAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetKernelGatewayAppSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetKernelGatewayAppSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ResetSpaceStorageSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetSpaceStorageSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfSpace_SpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

