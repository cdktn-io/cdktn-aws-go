package awsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsecs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsecs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoScalingMetric() *string
	// Experimental.
	SetAutoScalingMetric(val *string)
	// Experimental.
	AutoScalingMetricInput() *string
	// Experimental.
	AutoScalingTargetValue() *float64
	// Experimental.
	SetAutoScalingTargetValue(val *float64)
	// Experimental.
	AutoScalingTargetValueInput() *float64
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	MaxTaskCount() *float64
	// Experimental.
	SetMaxTaskCount(val *float64)
	// Experimental.
	MaxTaskCountInput() *float64
	// Experimental.
	MinTaskCount() *float64
	// Experimental.
	SetMinTaskCount(val *float64)
	// Experimental.
	MinTaskCountInput() *float64
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
	ResetAutoScalingMetric()
	// Experimental.
	ResetAutoScalingTargetValue()
	// Experimental.
	ResetMaxTaskCount()
	// Experimental.
	ResetMinTaskCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference
type jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) AutoScalingMetric() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) AutoScalingMetricInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoScalingMetricInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) AutoScalingTargetValue() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingTargetValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) AutoScalingTargetValueInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"autoScalingTargetValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) MaxTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) MaxTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) MinTaskCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTaskCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) MinTaskCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minTaskCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEcsExpressGatewayService_ScalingTargetPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsExpressGatewayService.ScalingTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference_Override(a AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ecs.AwsEcsExpressGatewayService.ScalingTargetPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetAutoScalingMetric(val *string) {
	if err := j.validateSetAutoScalingMetricParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingMetric",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetAutoScalingTargetValue(val *float64) {
	if err := j.validateSetAutoScalingTargetValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoScalingTargetValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetMaxTaskCount(val *float64) {
	if err := j.validateSetMaxTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTaskCount",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetMinTaskCount(val *float64) {
	if err := j.validateSetMinTaskCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTaskCount",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ResetAutoScalingMetric() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoScalingMetric",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ResetAutoScalingTargetValue() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoScalingTargetValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ResetMaxTaskCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTaskCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ResetMinTaskCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMinTaskCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEcsExpressGatewayService_ScalingTargetPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

