package codedeploy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/codedeploy/jsii"

	"github.com/cdktn-io/cdktn-aws-go/codedeploy/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDeploymentConfig_ZonalConfigPropertyOutputReference interface {
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
	FirstZoneMonitorDurationInSeconds() *float64
	// Experimental.
	SetFirstZoneMonitorDurationInSeconds(val *float64)
	// Experimental.
	FirstZoneMonitorDurationInSecondsInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDeploymentConfig_ZonalConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDeploymentConfig_ZonalConfigProperty)
	// Experimental.
	MinimumHealthyHostsPerZone() AwsDeploymentConfig_MinimumHealthyHostsPerZonePropertyOutputReference
	// Experimental.
	MinimumHealthyHostsPerZoneInput() *AwsDeploymentConfig_MinimumHealthyHostsPerZoneProperty
	// Experimental.
	MonitorDurationInSeconds() *float64
	// Experimental.
	SetMonitorDurationInSeconds(val *float64)
	// Experimental.
	MonitorDurationInSecondsInput() *float64
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
	PutMinimumHealthyHostsPerZone(value *AwsDeploymentConfig_MinimumHealthyHostsPerZoneProperty)
	// Experimental.
	ResetFirstZoneMonitorDurationInSeconds()
	// Experimental.
	ResetMinimumHealthyHostsPerZone()
	// Experimental.
	ResetMonitorDurationInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDeploymentConfig_ZonalConfigPropertyOutputReference
type jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) FirstZoneMonitorDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstZoneMonitorDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) FirstZoneMonitorDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"firstZoneMonitorDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) InternalValue() *AwsDeploymentConfig_ZonalConfigProperty {
	var returns *AwsDeploymentConfig_ZonalConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) MinimumHealthyHostsPerZone() AwsDeploymentConfig_MinimumHealthyHostsPerZonePropertyOutputReference {
	var returns AwsDeploymentConfig_MinimumHealthyHostsPerZonePropertyOutputReference
	_jsii_.Get(
		j,
		"minimumHealthyHostsPerZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) MinimumHealthyHostsPerZoneInput() *AwsDeploymentConfig_MinimumHealthyHostsPerZoneProperty {
	var returns *AwsDeploymentConfig_MinimumHealthyHostsPerZoneProperty
	_jsii_.Get(
		j,
		"minimumHealthyHostsPerZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) MonitorDurationInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorDurationInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) MonitorDurationInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"monitorDurationInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDeploymentConfig_ZonalConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDeploymentConfig_ZonalConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDeploymentConfig_ZonalConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentConfig.ZonalConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDeploymentConfig_ZonalConfigPropertyOutputReference_Override(a AwsDeploymentConfig_ZonalConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-codedeploy.AwsDeploymentConfig.ZonalConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetFirstZoneMonitorDurationInSeconds(val *float64) {
	if err := j.validateSetFirstZoneMonitorDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstZoneMonitorDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetInternalValue(val *AwsDeploymentConfig_ZonalConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetMonitorDurationInSeconds(val *float64) {
	if err := j.validateSetMonitorDurationInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitorDurationInSeconds",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) PutMinimumHealthyHostsPerZone(value *AwsDeploymentConfig_MinimumHealthyHostsPerZoneProperty) {
	if err := a.validatePutMinimumHealthyHostsPerZoneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMinimumHealthyHostsPerZone",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ResetFirstZoneMonitorDurationInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetFirstZoneMonitorDurationInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ResetMinimumHealthyHostsPerZone() {
	_jsii_.InvokeVoid(
		a,
		"resetMinimumHealthyHostsPerZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ResetMonitorDurationInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitorDurationInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDeploymentConfig_ZonalConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

