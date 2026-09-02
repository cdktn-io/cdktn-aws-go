package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsgamelift/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsgamelift/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFleet_RuntimeConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfFleet_RuntimeConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfFleet_RuntimeConfigurationProperty)
	// Experimental.
	MaxConcurrentGameSessionActivations() *float64
	// Experimental.
	SetMaxConcurrentGameSessionActivations(val *float64)
	// Experimental.
	MaxConcurrentGameSessionActivationsInput() *float64
	// Experimental.
	ServerProcess() TfFleet_ServerProcessPropertyList
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

// The jsii proxy struct for TfFleet_RuntimeConfigurationPropertyOutputReference
type jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GameSessionActivationTimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GameSessionActivationTimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"gameSessionActivationTimeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) InternalValue() *TfFleet_RuntimeConfigurationProperty {
	var returns *TfFleet_RuntimeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) MaxConcurrentGameSessionActivations() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) MaxConcurrentGameSessionActivationsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentGameSessionActivationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ServerProcess() TfFleet_ServerProcessPropertyList {
	var returns TfFleet_ServerProcessPropertyList
	_jsii_.Get(
		j,
		"serverProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ServerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFleet_RuntimeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFleet_RuntimeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFleet_RuntimeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-gamelift.TfFleet.RuntimeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFleet_RuntimeConfigurationPropertyOutputReference_Override(t TfFleet_RuntimeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-gamelift.TfFleet.RuntimeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetGameSessionActivationTimeoutSeconds(val *float64) {
	if err := j.validateSetGameSessionActivationTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gameSessionActivationTimeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetInternalValue(val *TfFleet_RuntimeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetMaxConcurrentGameSessionActivations(val *float64) {
	if err := j.validateSetMaxConcurrentGameSessionActivationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentGameSessionActivations",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) PutServerProcess(value interface{}) {
	if err := t.validatePutServerProcessParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putServerProcess",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ResetGameSessionActivationTimeoutSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetGameSessionActivationTimeoutSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ResetMaxConcurrentGameSessionActivations() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrentGameSessionActivations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ResetServerProcess() {
	_jsii_.InvokeVoid(
		t,
		"resetServerProcess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFleet_RuntimeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

