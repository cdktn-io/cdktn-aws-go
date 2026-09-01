package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CanarySize() AwsSagemakerEndpoint_CanarySizePropertyOutputReference
	// Experimental.
	CanarySizeInput() *AwsSagemakerEndpoint_CanarySizeProperty
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
	InternalValue() *AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty)
	// Experimental.
	LinearStepSize() AwsSagemakerEndpoint_LinearStepSizePropertyOutputReference
	// Experimental.
	LinearStepSizeInput() *AwsSagemakerEndpoint_LinearStepSizeProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
	// Experimental.
	WaitIntervalInSeconds() *float64
	// Experimental.
	SetWaitIntervalInSeconds(val *float64)
	// Experimental.
	WaitIntervalInSecondsInput() *float64
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
	PutCanarySize(value *AwsSagemakerEndpoint_CanarySizeProperty)
	// Experimental.
	PutLinearStepSize(value *AwsSagemakerEndpoint_LinearStepSizeProperty)
	// Experimental.
	ResetCanarySize()
	// Experimental.
	ResetLinearStepSize()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference
type jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) CanarySize() AwsSagemakerEndpoint_CanarySizePropertyOutputReference {
	var returns AwsSagemakerEndpoint_CanarySizePropertyOutputReference
	_jsii_.Get(
		j,
		"canarySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) CanarySizeInput() *AwsSagemakerEndpoint_CanarySizeProperty {
	var returns *AwsSagemakerEndpoint_CanarySizeProperty
	_jsii_.Get(
		j,
		"canarySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) InternalValue() *AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty {
	var returns *AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) LinearStepSize() AwsSagemakerEndpoint_LinearStepSizePropertyOutputReference {
	var returns AwsSagemakerEndpoint_LinearStepSizePropertyOutputReference
	_jsii_.Get(
		j,
		"linearStepSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) LinearStepSizeInput() *AwsSagemakerEndpoint_LinearStepSizeProperty {
	var returns *AwsSagemakerEndpoint_LinearStepSizeProperty
	_jsii_.Get(
		j,
		"linearStepSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) WaitIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) WaitIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitIntervalInSecondsInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpoint.TrafficRoutingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference_Override(a AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerEndpoint.TrafficRoutingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetInternalValue(val *AwsSagemakerEndpoint_TrafficRoutingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference)SetWaitIntervalInSeconds(val *float64) {
	if err := j.validateSetWaitIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitIntervalInSeconds",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) PutCanarySize(value *AwsSagemakerEndpoint_CanarySizeProperty) {
	if err := a.validatePutCanarySizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCanarySize",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) PutLinearStepSize(value *AwsSagemakerEndpoint_LinearStepSizeProperty) {
	if err := a.validatePutLinearStepSizeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLinearStepSize",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ResetCanarySize() {
	_jsii_.InvokeVoid(
		a,
		"resetCanarySize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ResetLinearStepSize() {
	_jsii_.InvokeVoid(
		a,
		"resetLinearStepSize",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerEndpoint_TrafficRoutingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

