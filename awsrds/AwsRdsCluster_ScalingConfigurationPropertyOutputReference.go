package awsrds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsrds/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsrds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRdsCluster_ScalingConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AutoPause() interface{}
	// Experimental.
	SetAutoPause(val interface{})
	// Experimental.
	AutoPauseInput() interface{}
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
	InternalValue() *AwsRdsCluster_ScalingConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsRdsCluster_ScalingConfigurationProperty)
	// Experimental.
	MaxCapacity() *float64
	// Experimental.
	SetMaxCapacity(val *float64)
	// Experimental.
	MaxCapacityInput() *float64
	// Experimental.
	MinCapacity() *float64
	// Experimental.
	SetMinCapacity(val *float64)
	// Experimental.
	MinCapacityInput() *float64
	// Experimental.
	SecondsBeforeTimeout() *float64
	// Experimental.
	SetSecondsBeforeTimeout(val *float64)
	// Experimental.
	SecondsBeforeTimeoutInput() *float64
	// Experimental.
	SecondsUntilAutoPause() *float64
	// Experimental.
	SetSecondsUntilAutoPause(val *float64)
	// Experimental.
	SecondsUntilAutoPauseInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutAction() *string
	// Experimental.
	SetTimeoutAction(val *string)
	// Experimental.
	TimeoutActionInput() *string
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
	ResetAutoPause()
	// Experimental.
	ResetMaxCapacity()
	// Experimental.
	ResetMinCapacity()
	// Experimental.
	ResetSecondsBeforeTimeout()
	// Experimental.
	ResetSecondsUntilAutoPause()
	// Experimental.
	ResetTimeoutAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRdsCluster_ScalingConfigurationPropertyOutputReference
type jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) AutoPause() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) AutoPauseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoPauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) InternalValue() *AwsRdsCluster_ScalingConfigurationProperty {
	var returns *AwsRdsCluster_ScalingConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) MaxCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) MaxCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) MinCapacity() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) MinCapacityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minCapacityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) SecondsBeforeTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBeforeTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) SecondsBeforeTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsBeforeTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) SecondsUntilAutoPause() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsUntilAutoPause",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) SecondsUntilAutoPauseInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"secondsUntilAutoPauseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) TimeoutAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) TimeoutActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeoutActionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRdsCluster_ScalingConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRdsCluster_ScalingConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRdsCluster_ScalingConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-rds.AwsRdsCluster.ScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRdsCluster_ScalingConfigurationPropertyOutputReference_Override(a AwsRdsCluster_ScalingConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.AwsRdsCluster.ScalingConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetAutoPause(val interface{}) {
	if err := j.validateSetAutoPauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoPause",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetInternalValue(val *AwsRdsCluster_ScalingConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetMaxCapacity(val *float64) {
	if err := j.validateSetMaxCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetMinCapacity(val *float64) {
	if err := j.validateSetMinCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minCapacity",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetSecondsBeforeTimeout(val *float64) {
	if err := j.validateSetSecondsBeforeTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsBeforeTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetSecondsUntilAutoPause(val *float64) {
	if err := j.validateSetSecondsUntilAutoPauseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondsUntilAutoPause",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference)SetTimeoutAction(val *string) {
	if err := j.validateSetTimeoutActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutAction",
		val,
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetAutoPause() {
	_jsii_.InvokeVoid(
		a,
		"resetAutoPause",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetMaxCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetMinCapacity() {
	_jsii_.InvokeVoid(
		a,
		"resetMinCapacity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetSecondsBeforeTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondsBeforeTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetSecondsUntilAutoPause() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondsUntilAutoPause",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ResetTimeoutAction() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRdsCluster_ScalingConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

