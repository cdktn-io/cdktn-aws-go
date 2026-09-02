package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference interface {
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
	ContainerArguments() *[]*string
	// Experimental.
	SetContainerArguments(val *[]*string)
	// Experimental.
	ContainerArgumentsInput() *[]*string
	// Experimental.
	ContainerEntrypoint() *[]*string
	// Experimental.
	SetContainerEntrypoint(val *[]*string)
	// Experimental.
	ContainerEntrypointInput() *[]*string
	// Experimental.
	ContainerEnvironmentVariables() *map[string]*string
	// Experimental.
	SetContainerEnvironmentVariables(val *map[string]*string)
	// Experimental.
	ContainerEnvironmentVariablesInput() *map[string]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	// Experimental.
	SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty)
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
	ResetContainerArguments()
	// Experimental.
	ResetContainerEntrypoint()
	// Experimental.
	ResetContainerEnvironmentVariables()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference
type jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ContainerEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) InternalValue() *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty {
	var returns *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigContainerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference_Override(t TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfAppImageConfig.CodeEditorAppImageConfigContainerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetContainerEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetContainerEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetInternalValue(val *TfAppImageConfig_CodeEditorAppImageConfigContainerConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ResetContainerEnvironmentVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetContainerEnvironmentVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAppImageConfig_CodeEditorAppImageConfigContainerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

