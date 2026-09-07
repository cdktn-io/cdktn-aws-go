package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference interface {
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
	ContainerConfig() AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference
	// Experimental.
	ContainerConfigInput() *AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FileSystemConfig() AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference
	// Experimental.
	FileSystemConfigInput() *AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAppImageConfig_CodeEditorAppImageConfigProperty
	// Experimental.
	SetInternalValue(val *AwsAppImageConfig_CodeEditorAppImageConfigProperty)
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
	PutContainerConfig(value *AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty)
	// Experimental.
	PutFileSystemConfig(value *AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty)
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

// The jsii proxy struct for AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference
type jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ContainerConfig() AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference {
	var returns AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"containerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ContainerConfigInput() *AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty {
	var returns *AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	_jsii_.Get(
		j,
		"containerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) FileSystemConfig() AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference {
	var returns AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"fileSystemConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) FileSystemConfigInput() *AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty {
	var returns *AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty
	_jsii_.Get(
		j,
		"fileSystemConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InternalValue() *AwsAppImageConfig_CodeEditorAppImageConfigProperty {
	var returns *AwsAppImageConfig_CodeEditorAppImageConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsAppImageConfig.CodeEditorAppImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference_Override(a AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsAppImageConfig.CodeEditorAppImageConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetInternalValue(val *AwsAppImageConfig_CodeEditorAppImageConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) PutContainerConfig(value *AwsAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty) {
	if err := a.validatePutContainerConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putContainerConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) PutFileSystemConfig(value *AwsAppImageConfig_CodeEditorAppImageConfigFileSystemConfigProperty) {
	if err := a.validatePutFileSystemConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFileSystemConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ResetContainerConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ResetFileSystemConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetFileSystemConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppImageConfig_CodeEditorAppImageConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

