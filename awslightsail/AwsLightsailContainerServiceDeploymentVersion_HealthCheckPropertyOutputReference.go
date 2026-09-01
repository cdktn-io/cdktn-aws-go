package awslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslightsail/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslightsail/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference interface {
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
	HealthyThreshold() *float64
	// Experimental.
	SetHealthyThreshold(val *float64)
	// Experimental.
	HealthyThresholdInput() *float64
	// Experimental.
	InternalValue() *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty
	// Experimental.
	SetInternalValue(val *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty)
	// Experimental.
	IntervalSeconds() *float64
	// Experimental.
	SetIntervalSeconds(val *float64)
	// Experimental.
	IntervalSecondsInput() *float64
	// Experimental.
	Path() *string
	// Experimental.
	SetPath(val *string)
	// Experimental.
	PathInput() *string
	// Experimental.
	SuccessCodes() *string
	// Experimental.
	SetSuccessCodes(val *string)
	// Experimental.
	SuccessCodesInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutSeconds() *float64
	// Experimental.
	SetTimeoutSeconds(val *float64)
	// Experimental.
	TimeoutSecondsInput() *float64
	// Experimental.
	UnhealthyThreshold() *float64
	// Experimental.
	SetUnhealthyThreshold(val *float64)
	// Experimental.
	UnhealthyThresholdInput() *float64
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
	ResetHealthyThreshold()
	// Experimental.
	ResetIntervalSeconds()
	// Experimental.
	ResetPath()
	// Experimental.
	ResetSuccessCodes()
	// Experimental.
	ResetTimeoutSeconds()
	// Experimental.
	ResetUnhealthyThreshold()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference
type jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) HealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) HealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthyThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) InternalValue() *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty {
	var returns *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) IntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) IntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) SuccessCodes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) SuccessCodesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successCodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) UnhealthyThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) UnhealthyThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"unhealthyThresholdInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lightsail.AwsLightsailContainerServiceDeploymentVersion.HealthCheckPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference_Override(a AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lightsail.AwsLightsailContainerServiceDeploymentVersion.HealthCheckPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetHealthyThreshold(val *float64) {
	if err := j.validateSetHealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthyThreshold",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetInternalValue(val *AwsLightsailContainerServiceDeploymentVersion_HealthCheckProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetIntervalSeconds(val *float64) {
	if err := j.validateSetIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetSuccessCodes(val *string) {
	if err := j.validateSetSuccessCodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successCodes",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference)SetUnhealthyThreshold(val *float64) {
	if err := j.validateSetUnhealthyThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unhealthyThreshold",
		val,
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetHealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetHealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetIntervalSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetIntervalSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetSuccessCodes() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessCodes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ResetUnhealthyThreshold() {
	_jsii_.InvokeVoid(
		a,
		"resetUnhealthyThreshold",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLightsailContainerServiceDeploymentVersion_HealthCheckPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

