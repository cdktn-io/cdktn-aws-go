package apprunner

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/apprunner/jsii"

	"github.com/cdktn-io/cdktn-aws-go/apprunner/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsService_CodeConfigurationValuesPropertyOutputReference interface {
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
	InternalValue() *AwsService_CodeConfigurationValuesProperty
	// Experimental.
	SetInternalValue(val *AwsService_CodeConfigurationValuesProperty)
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

// The jsii proxy struct for AwsService_CodeConfigurationValuesPropertyOutputReference
type jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) BuildCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) BuildCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"buildCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) InternalValue() *AwsService_CodeConfigurationValuesProperty {
	var returns *AwsService_CodeConfigurationValuesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) Port() *string {
	var returns *string
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) PortInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) Runtime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentSecrets() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentSecretsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) RuntimeEnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"runtimeEnvironmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) RuntimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) StartCommand() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommand",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) StartCommandInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startCommandInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsService_CodeConfigurationValuesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsService_CodeConfigurationValuesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsService_CodeConfigurationValuesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsService.CodeConfigurationValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsService_CodeConfigurationValuesPropertyOutputReference_Override(a AwsService_CodeConfigurationValuesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-runner.AwsService.CodeConfigurationValuesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetBuildCommand(val *string) {
	if err := j.validateSetBuildCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"buildCommand",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetInternalValue(val *AwsService_CodeConfigurationValuesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetPort(val *string) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetRuntime(val *string) {
	if err := j.validateSetRuntimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtime",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetRuntimeEnvironmentSecrets(val *map[string]*string) {
	if err := j.validateSetRuntimeEnvironmentSecretsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironmentSecrets",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetRuntimeEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetRuntimeEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeEnvironmentVariables",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetStartCommand(val *string) {
	if err := j.validateSetStartCommandParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startCommand",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ResetBuildCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetBuildCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ResetPort() {
	_jsii_.InvokeVoid(
		a,
		"resetPort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ResetRuntimeEnvironmentSecrets() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeEnvironmentSecrets",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ResetRuntimeEnvironmentVariables() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeEnvironmentVariables",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ResetStartCommand() {
	_jsii_.InvokeVoid(
		a,
		"resetStartCommand",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsService_CodeConfigurationValuesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

