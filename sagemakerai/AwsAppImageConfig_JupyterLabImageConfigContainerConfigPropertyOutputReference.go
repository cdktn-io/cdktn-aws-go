package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference interface {
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
	InternalValue() *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty
	// Experimental.
	SetInternalValue(val *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty)
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

// The jsii proxy struct for AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference
type jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerArguments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerArgumentsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerArgumentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerEntrypoint() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerEntrypointInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"containerEntrypointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ContainerEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"containerEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) InternalValue() *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty {
	var returns *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsAppImageConfig.JupyterLabImageConfigContainerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference_Override(a AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsAppImageConfig.JupyterLabImageConfigContainerConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetContainerArguments(val *[]*string) {
	if err := j.validateSetContainerArgumentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerArguments",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetContainerEntrypoint(val *[]*string) {
	if err := j.validateSetContainerEntrypointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEntrypoint",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetContainerEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetContainerEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetInternalValue(val *AwsAppImageConfig_JupyterLabImageConfigContainerConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ResetContainerArguments() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerArguments",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ResetContainerEntrypoint() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerEntrypoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ResetContainerEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetContainerEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppImageConfig_JupyterLabImageConfigContainerConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

