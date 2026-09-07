package sagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/sagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/sagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference interface {
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
	InternalValue() *AwsEndpoint_BlueGreenUpdatePolicyProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_BlueGreenUpdatePolicyProperty)
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
	TrafficRoutingConfiguration() AwsEndpoint_TrafficRoutingConfigurationPropertyOutputReference
	// Experimental.
	TrafficRoutingConfigurationInput() *AwsEndpoint_TrafficRoutingConfigurationProperty
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
	PutTrafficRoutingConfiguration(value *AwsEndpoint_TrafficRoutingConfigurationProperty)
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

// The jsii proxy struct for AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference
type jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InternalValue() *AwsEndpoint_BlueGreenUpdatePolicyProperty {
	var returns *AwsEndpoint_BlueGreenUpdatePolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) MaximumExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerminationWaitInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerminationWaitInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"terminationWaitInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TrafficRoutingConfiguration() AwsEndpoint_TrafficRoutingConfigurationPropertyOutputReference {
	var returns AwsEndpoint_TrafficRoutingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"trafficRoutingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) TrafficRoutingConfigurationInput() *AwsEndpoint_TrafficRoutingConfigurationProperty {
	var returns *AwsEndpoint_TrafficRoutingConfigurationProperty
	_jsii_.Get(
		j,
		"trafficRoutingConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.BlueGreenUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference_Override(a AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsEndpoint.BlueGreenUpdatePolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetInternalValue(val *AwsEndpoint_BlueGreenUpdatePolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetMaximumExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetMaximumExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerminationWaitInSeconds(val *float64) {
	if err := j.validateSetTerminationWaitInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terminationWaitInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) PutTrafficRoutingConfiguration(value *AwsEndpoint_TrafficRoutingConfigurationProperty) {
	if err := a.validatePutTrafficRoutingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTrafficRoutingConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ResetMaximumExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ResetTerminationWaitInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTerminationWaitInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_BlueGreenUpdatePolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

