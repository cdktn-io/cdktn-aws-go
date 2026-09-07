package codedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference interface {
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
	InternalValue() *AwsDeploymentConfig_TrafficRoutingConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDeploymentConfig_TrafficRoutingConfigProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeBasedCanary() AwsDeploymentConfig_TimeBasedCanaryPropertyOutputReference
	// Experimental.
	TimeBasedCanaryInput() *AwsDeploymentConfig_TimeBasedCanaryProperty
	// Experimental.
	TimeBasedLinear() AwsDeploymentConfig_TimeBasedLinearPropertyOutputReference
	// Experimental.
	TimeBasedLinearInput() *AwsDeploymentConfig_TimeBasedLinearProperty
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	PutTimeBasedCanary(value *AwsDeploymentConfig_TimeBasedCanaryProperty)
	// Experimental.
	PutTimeBasedLinear(value *AwsDeploymentConfig_TimeBasedLinearProperty)
	// Experimental.
	ResetTimeBasedCanary()
	// Experimental.
	ResetTimeBasedLinear()
	// Experimental.
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference
type jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InternalValue() *AwsDeploymentConfig_TrafficRoutingConfigProperty {
	var returns *AwsDeploymentConfig_TrafficRoutingConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedCanary() AwsDeploymentConfig_TimeBasedCanaryPropertyOutputReference {
	var returns AwsDeploymentConfig_TimeBasedCanaryPropertyOutputReference
	_jsii_.Get(
		j,
		"timeBasedCanary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedCanaryInput() *AwsDeploymentConfig_TimeBasedCanaryProperty {
	var returns *AwsDeploymentConfig_TimeBasedCanaryProperty
	_jsii_.Get(
		j,
		"timeBasedCanaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedLinear() AwsDeploymentConfig_TimeBasedLinearPropertyOutputReference {
	var returns AwsDeploymentConfig_TimeBasedLinearPropertyOutputReference
	_jsii_.Get(
		j,
		"timeBasedLinear",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TimeBasedLinearInput() *AwsDeploymentConfig_TimeBasedLinearProperty {
	var returns *AwsDeploymentConfig_TimeBasedLinearProperty
	_jsii_.Get(
		j,
		"timeBasedLinearInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentConfig.TrafficRoutingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference_Override(a AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentConfig.TrafficRoutingConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetInternalValue(val *AwsDeploymentConfig_TrafficRoutingConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) PutTimeBasedCanary(value *AwsDeploymentConfig_TimeBasedCanaryProperty) {
	if err := a.validatePutTimeBasedCanaryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeBasedCanary",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) PutTimeBasedLinear(value *AwsDeploymentConfig_TimeBasedLinearProperty) {
	if err := a.validatePutTimeBasedLinearParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeBasedLinear",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetTimeBasedCanary() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeBasedCanary",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetTimeBasedLinear() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeBasedLinear",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		a,
		"resetType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeploymentConfig_TrafficRoutingConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

