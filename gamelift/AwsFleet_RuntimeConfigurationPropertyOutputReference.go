package gamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/gamelift/jsii"

	"github.com/cdktn-io/cdktn-aws-go/gamelift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFleet_RuntimeConfigurationPropertyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	GameSessionActivationTimeoutSeconds() *float64
	// Experimental.
	SetGameSessionActivationTimeoutSeconds(val *float64)
	// Experimental.
	GameSessionActivationTimeoutSecondsInput() *float64
	// Experimental.
	InternalValue() *AwsFleet_RuntimeConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFleet_RuntimeConfigurationProperty)
	// Experimental.
	MaxConcurrentGameSessionActivations() *float64
	// Experimental.
	SetMaxConcurrentGameSessionActivations(val *float64)
	// Experimental.
	MaxConcurrentGameSessionActivationsInput() *float64
	// Experimental.
	ServerProcess() AwsFleet_ServerProcessPropertyList
	// Experimental.
	ServerProcessInput() interface{}
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
	PutServerProcess(value interface{})
	// Experimental.
	ResetGameSessionActivationTimeoutSeconds()
	// Experimental.
	ResetMaxConcurrentGameSessionActivations()
	// Experimental.
	ResetServerProcess()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFleet_RuntimeConfigurationPropertyOutputReference
type jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GameSessionActivationTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GameSessionActivationTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) InternalValue() *AwsFleet_RuntimeConfigurationProperty {
	var returns *AwsFleet_RuntimeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) MaxConcurrentGameSessionActivations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) MaxConcurrentGameSessionActivationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ServerProcess() AwsFleet_ServerProcessPropertyList {
	var returns AwsFleet_ServerProcessPropertyList
	_jsii_.Get(
		j,
		"serverProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ServerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFleet_RuntimeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFleet_RuntimeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFleet_RuntimeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-gamelift.AwsFleet.RuntimeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFleet_RuntimeConfigurationPropertyOutputReference_Override(a AwsFleet_RuntimeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-gamelift.AwsFleet.RuntimeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetGameSessionActivationTimeoutSeconds(val *float64) {
	if err := j.validateSetGameSessionActivationTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionActivationTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetInternalValue(val *AwsFleet_RuntimeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetMaxConcurrentGameSessionActivations(val *float64) {
	if err := j.validateSetMaxConcurrentGameSessionActivationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentGameSessionActivations",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) PutServerProcess(value interface{}) {
	if err := a.validatePutServerProcessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerProcess",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ResetGameSessionActivationTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetGameSessionActivationTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ResetMaxConcurrentGameSessionActivations() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrentGameSessionActivations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ResetServerProcess() {
	_jsii_.InvokeVoid(
		a,
		"resetServerProcess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFleet_RuntimeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

