package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference interface {
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
	InternalValue() *TfEndpoint_BlueGreenUpdatePolicyProperty
	// Experimental.
	SetInternalValue(val *TfEndpoint_BlueGreenUpdatePolicyProperty)
	// Experimental.
	MaximumExecutionTimeoutInSeconds() *float64
	// Experimental.
	SetMaximumExecutionTimeoutInSeconds(val *float64)
	// Experimental.
	MaximumExecutionTimeoutInSecondsInput() *float64
	// Experimental.
	TerminationWaitInSeconds() *float64
	// Experimental.
	SetTerminationWaitInSeconds(val *float64)
	// Experimental.
	TerminationWaitInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TrafficRoutingConfiguration() TfEndpoint_TrafficRoutingConfigurationPropertyOutputReference
	// Experimental.
	TrafficRoutingConfigurationInput() *TfEndpoint_TrafficRoutingConfigurationProperty
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
	PutTrafficRoutingConfiguration(value *TfEndpoint_TrafficRoutingConfigurationProperty)
	// Experimental.
	ResetMaximumExecutionTimeoutInSeconds()
	// Experimental.
	ResetTerminationWaitInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
type jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InternalValue() *TfEndpoint_BlueGreenUpdatePolicyProperty {
	var returns *TfEndpoint_BlueGreenUpdatePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerminationWaitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerminationWaitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TrafficRoutingConfiguration() TfEndpoint_TrafficRoutingConfigurationPropertyOutputReference {
	var returns TfEndpoint_TrafficRoutingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"trafficRoutingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TrafficRoutingConfigurationInput() *TfEndpoint_TrafficRoutingConfigurationProperty {
	var returns *TfEndpoint_TrafficRoutingConfigurationProperty
	_jsii_.Get(
		j,
		"trafficRoutingConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfEndpoint_BlueGreenUpdatePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.BlueGreenUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference_Override(t TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfEndpoint.BlueGreenUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetInternalValue(val *TfEndpoint_BlueGreenUpdatePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetMaximumExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetMaximumExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerminationWaitInSeconds(val *float64) {
	if err := j.validateSetTerminationWaitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationWaitInSeconds",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) PutTrafficRoutingConfiguration(value *TfEndpoint_TrafficRoutingConfigurationProperty) {
	if err := t.validatePutTrafficRoutingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTrafficRoutingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ResetMaximumExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ResetTerminationWaitInSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetTerminationWaitInSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

