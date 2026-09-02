package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference interface {
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
	ContainerConfig() TfAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference
	// Experimental.
	ContainerConfigInput() *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FileSystemConfig() TfAppImageConfig_JupyterLabImageConfigFileSystemConfigPropertyOutputReference
	// Experimental.
	FileSystemConfigInput() *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfAppImageConfig_JupyterLabImageConfigProperty
	// Experimental.
	SetInternalValue(val *TfAppImageConfig_JupyterLabImageConfigProperty)
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
	PutContainerConfig(value *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty)
	// Experimental.
	PutFileSystemConfig(value *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty)
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

// The jsii proxy struct for TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference
type jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ContainerConfig() TfAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference {
	var returns TfAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"containerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ContainerConfigInput() *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty {
	var returns *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty
	_jsii_.Get(
		j,
		"containerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) FileSystemConfig() TfAppImageConfig_JupyterLabImageConfigFileSystemConfigPropertyOutputReference {
	var returns TfAppImageConfig_JupyterLabImageConfigFileSystemConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) FileSystemConfigInput() *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty {
	var returns *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) InternalValue() *TfAppImageConfig_JupyterLabImageConfigProperty {
	var returns *TfAppImageConfig_JupyterLabImageConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAppImageConfig_JupyterLabImageConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAppImageConfig_JupyterLabImageConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.JupyterLabImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAppImageConfig_JupyterLabImageConfigPropertyOutputReference_Override(t TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.JupyterLabImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference)SetInternalValue(val *TfAppImageConfig_JupyterLabImageConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) PutContainerConfig(value *TfAppImageConfig_JupyterLabImageConfigContainerConfigProperty) {
	if err := t.validatePutContainerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putContainerConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) PutFileSystemConfig(value *TfAppImageConfig_JupyterLabImageConfigFileSystemConfigProperty) {
	if err := t.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ResetContainerConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_JupyterLabImageConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

