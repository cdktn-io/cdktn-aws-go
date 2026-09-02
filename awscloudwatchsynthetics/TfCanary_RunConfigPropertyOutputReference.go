package awscloudwatchsynthetics

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchsynthetics/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchsynthetics/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCanary_RunConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActiveTracing() interface{}
	// Experimental.
	SetActiveTracing(val interface{})
	// Experimental.
	ActiveTracingInput() interface{}
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
	EnvironmentVariables() *map[string]*string
	// Experimental.
	SetEnvironmentVariables(val *map[string]*string)
	// Experimental.
	EnvironmentVariablesInput() *map[string]*string
	// Experimental.
	EphemeralStorage() *float64
	// Experimental.
	SetEphemeralStorage(val *float64)
	// Experimental.
	EphemeralStorageInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCanary_RunConfigProperty
	// Experimental.
	SetInternalValue(val *TfCanary_RunConfigProperty)
	// Experimental.
	MemoryInMb() *float64
	// Experimental.
	SetMemoryInMb(val *float64)
	// Experimental.
	MemoryInMbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutInSeconds() *float64
	// Experimental.
	SetTimeoutInSeconds(val *float64)
	// Experimental.
	TimeoutInSecondsInput() *float64
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
	ResetActiveTracing()
	// Experimental.
	ResetEnvironmentVariables()
	// Experimental.
	ResetEphemeralStorage()
	// Experimental.
	ResetMemoryInMb()
	// Experimental.
	ResetTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCanary_RunConfigPropertyOutputReference
type jsiiProxy_TfCanary_RunConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ActiveTracing() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTracing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ActiveTracingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"activeTracingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) EnvironmentVariables() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariables",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) EnvironmentVariablesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"environmentVariablesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) EphemeralStorage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) EphemeralStorageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ephemeralStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) InternalValue() *TfCanary_RunConfigProperty {
	var returns *TfCanary_RunConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) MemoryInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) MemoryInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memoryInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCanary_RunConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCanary_RunConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCanary_RunConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCanary_RunConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary.RunConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCanary_RunConfigPropertyOutputReference_Override(t TfCanary_RunConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-synthetics.TfCanary.RunConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetActiveTracing(val interface{}) {
	if err := j.validateSetActiveTracingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"activeTracing",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetEnvironmentVariables(val *map[string]*string) {
	if err := j.validateSetEnvironmentVariablesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"environmentVariables",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetEphemeralStorage(val *float64) {
	if err := j.validateSetEphemeralStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralStorage",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetInternalValue(val *TfCanary_RunConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetMemoryInMb(val *float64) {
	if err := j.validateSetMemoryInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memoryInMb",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCanary_RunConfigPropertyOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ResetActiveTracing() {
	_jsii_.InvokeVoid(
		t,
		"resetActiveTracing",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ResetEnvironmentVariables() {
	_jsii_.InvokeVoid(
		t,
		"resetEnvironmentVariables",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ResetEphemeralStorage() {
	_jsii_.InvokeVoid(
		t,
		"resetEphemeralStorage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ResetMemoryInMb() {
	_jsii_.InvokeVoid(
		t,
		"resetMemoryInMb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCanary_RunConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

