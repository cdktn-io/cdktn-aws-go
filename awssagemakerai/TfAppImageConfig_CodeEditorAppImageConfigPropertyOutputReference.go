package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference interface {
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
	// Experimental.
	ContainerConfig() TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference
	// Experimental.
	ContainerConfigInput() *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FileSystemConfig() TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference
	// Experimental.
	FileSystemConfigInput() *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigProperty
	// Experimental.
	SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigProperty)
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
	PutContainerConfig(value *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty)
	// Experimental.
	PutFileSystemConfig(value *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty)
	// Experimental.
	ResetContainerConfig()
	// Experimental.
	ResetFileSystemConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference
type jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ContainerConfig() TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference {
	var returns TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"containerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ContainerConfigInput() *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty {
	var returns *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	_jsii_.Get(
		j,
		"containerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) FileSystemConfig() TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference {
	var returns TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) FileSystemConfigInput() *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty {
	var returns *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigProperty {
	var returns *TfAppImageConfig_CodeEditorAppImageConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference_Override(t TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) PutContainerConfig(value *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty) {
	if err := t.validatePutContainerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) PutFileSystemConfig(value *TfAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty) {
	if err := t.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ResetContainerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

