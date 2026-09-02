package awsapprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsapprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsapprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfService_CodeConfigurationValuesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BuildCommand() *string
	// Experimental.
	SetBuildCommand(val *string)
	// Experimental.
	BuildCommandInput() *string
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
	Fqn() *string
	// Experimental.
	InternalValue() *TfService_CodeConfigurationValuesProperty
	// Experimental.
	SetInternalValue(val *TfService_CodeConfigurationValuesProperty)
	// Experimental.
	Port() *string
	// Experimental.
	SetPort(val *string)
	// Experimental.
	PortInput() *string
	// Experimental.
	Runtime() *string
	// Experimental.
	SetRuntime(val *string)
	// Experimental.
	RuntimeEnvironmentSecrets() *map[string]*string
	// Experimental.
	SetRuntimeEnvironmentSecrets(val *map[string]*string)
	// Experimental.
	RuntimeEnvironmentSecretsInput() *map[string]*string
	// Experimental.
	RuntimeEnvironmentVariables() *map[string]*string
	// Experimental.
	SetRuntimeEnvironmentVariables(val *map[string]*string)
	// Experimental.
	RuntimeEnvironmentVariablesInput() *map[string]*string
	// Experimental.
	RuntimeInput() *string
	// Experimental.
	StartCommand() *string
	// Experimental.
	SetStartCommand(val *string)
	// Experimental.
	StartCommandInput() *string
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
	ResetBuildCommand()
	// Experimental.
	ResetPort()
	// Experimental.
	ResetRuntimeEnvironmentSecrets()
	// Experimental.
	ResetRuntimeEnvironmentVariables()
	// Experimental.
	ResetStartCommand()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfService_CodeConfigurationValuesPropertyOutputReference
type jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) InternalValue() *TfService_CodeConfigurationValuesProperty {
	var returns *TfService_CodeConfigurationValuesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentSecrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentSecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) StartCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) StartCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfService_CodeConfigurationValuesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfService_CodeConfigurationValuesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfService_CodeConfigurationValuesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.CodeConfigurationValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfService_CodeConfigurationValuesPropertyOutputReference_Override(t TfService_CodeConfigurationValuesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.TfService.CodeConfigurationValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetInternalValue(val *TfService_CodeConfigurationValuesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetRuntimeEnvironmentSecrets(val *map[string]*string) {
	if err := j.validateSetRuntimeEnvironmentSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironmentSecrets",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetRuntimeEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetRuntimeEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetStartCommand(val *string) {
	if err := j.validateSetStartCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startCommand",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		t,
		"resetPort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ResetRuntimeEnvironmentSecrets() {
	_jsii_.InvokeVoid(
		t,
		"resetRuntimeEnvironmentSecrets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ResetRuntimeEnvironmentVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetRuntimeEnvironmentVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ResetStartCommand() {
	_jsii_.InvokeVoid(
		t,
		"resetStartCommand",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfService_CodeConfigurationValuesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

