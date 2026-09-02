package awsworkspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsworkspaces/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsworkspaces/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AccessEndpointConfig() TfDirectory_AccessEndpointConfigPropertyOutputReference
	// Experimental.
	AccessEndpointConfigInput() *TfDirectory_AccessEndpointConfigProperty
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
	DeviceTypeAndroid() *string
	// Experimental.
	SetDeviceTypeAndroid(val *string)
	// Experimental.
	DeviceTypeAndroidInput() *string
	// Experimental.
	DeviceTypeChromeos() *string
	// Experimental.
	SetDeviceTypeChromeos(val *string)
	// Experimental.
	DeviceTypeChromeosInput() *string
	// Experimental.
	DeviceTypeIos() *string
	// Experimental.
	SetDeviceTypeIos(val *string)
	// Experimental.
	DeviceTypeIosInput() *string
	// Experimental.
	DeviceTypeLinux() *string
	// Experimental.
	SetDeviceTypeLinux(val *string)
	// Experimental.
	DeviceTypeLinuxInput() *string
	// Experimental.
	DeviceTypeOsx() *string
	// Experimental.
	SetDeviceTypeOsx(val *string)
	// Experimental.
	DeviceTypeOsxInput() *string
	// Experimental.
	DeviceTypeWeb() *string
	// Experimental.
	SetDeviceTypeWeb(val *string)
	// Experimental.
	DeviceTypeWebInput() *string
	// Experimental.
	DeviceTypeWindows() *string
	// Experimental.
	SetDeviceTypeWindows(val *string)
	// Experimental.
	DeviceTypeWindowsInput() *string
	// Experimental.
	DeviceTypeZeroclient() *string
	// Experimental.
	SetDeviceTypeZeroclient(val *string)
	// Experimental.
	DeviceTypeZeroclientInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDirectory_WorkspaceAccessPropertiesProperty
	// Experimental.
	SetInternalValue(val *TfDirectory_WorkspaceAccessPropertiesProperty)
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
	PutAccessEndpointConfig(value *TfDirectory_AccessEndpointConfigProperty)
	// Experimental.
	ResetAccessEndpointConfig()
	// Experimental.
	ResetDeviceTypeAndroid()
	// Experimental.
	ResetDeviceTypeChromeos()
	// Experimental.
	ResetDeviceTypeIos()
	// Experimental.
	ResetDeviceTypeLinux()
	// Experimental.
	ResetDeviceTypeOsx()
	// Experimental.
	ResetDeviceTypeWeb()
	// Experimental.
	ResetDeviceTypeWindows()
	// Experimental.
	ResetDeviceTypeZeroclient()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference
type jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) AccessEndpointConfig() TfDirectory_AccessEndpointConfigPropertyOutputReference {
	var returns TfDirectory_AccessEndpointConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"accessEndpointConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) AccessEndpointConfigInput() *TfDirectory_AccessEndpointConfigProperty {
	var returns *TfDirectory_AccessEndpointConfigProperty
	_jsii_.Get(
		j,
		"accessEndpointConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeAndroid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeAndroidInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeAndroidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeChromeos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeChromeosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeChromeosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeIos() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIos",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeIosInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeIosInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeLinux() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinux",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeLinuxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeLinuxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeOsx() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsx",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeOsxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeOsxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeWeb() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWeb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeWebInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWebInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeWindows() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeWindowsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeWindowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeZeroclient() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) DeviceTypeZeroclientInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeZeroclientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) InternalValue() *TfDirectory_WorkspaceAccessPropertiesProperty {
	var returns *TfDirectory_WorkspaceAccessPropertiesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDirectory_WorkspaceAccessPropertiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDirectory_WorkspaceAccessPropertiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory.WorkspaceAccessPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDirectory_WorkspaceAccessPropertiesPropertyOutputReference_Override(t TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-workspaces.TfDirectory.WorkspaceAccessPropertiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeAndroid(val *string) {
	if err := j.validateSetDeviceTypeAndroidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeAndroid",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeChromeos(val *string) {
	if err := j.validateSetDeviceTypeChromeosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeChromeos",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeIos(val *string) {
	if err := j.validateSetDeviceTypeIosParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeIos",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeLinux(val *string) {
	if err := j.validateSetDeviceTypeLinuxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeLinux",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeOsx(val *string) {
	if err := j.validateSetDeviceTypeOsxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeOsx",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeWeb(val *string) {
	if err := j.validateSetDeviceTypeWebParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeWeb",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeWindows(val *string) {
	if err := j.validateSetDeviceTypeWindowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeWindows",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetDeviceTypeZeroclient(val *string) {
	if err := j.validateSetDeviceTypeZeroclientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceTypeZeroclient",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetInternalValue(val *TfDirectory_WorkspaceAccessPropertiesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) PutAccessEndpointConfig(value *TfDirectory_AccessEndpointConfigProperty) {
	if err := t.validatePutAccessEndpointConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAccessEndpointConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetAccessEndpointConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetAccessEndpointConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeAndroid() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeAndroid",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeChromeos() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeChromeos",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeIos() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeIos",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeLinux() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeLinux",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeOsx() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeOsx",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeWeb() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeWeb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeWindows() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeWindows",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ResetDeviceTypeZeroclient() {
	_jsii_.InvokeVoid(
		t,
		"resetDeviceTypeZeroclient",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDirectory_WorkspaceAccessPropertiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

