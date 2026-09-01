package awscloudwatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Alarm() *string
	// Experimental.
	SetAlarm(val *string)
	// Experimental.
	AlarmInput() *string
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
	ExtensionPeriod() *float64
	// Experimental.
	SetExtensionPeriod(val *float64)
	// Experimental.
	ExtensionPeriodInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty
	// Experimental.
	SetInternalValue(val *AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WaitPeriod() *float64
	// Experimental.
	SetWaitPeriod(val *float64)
	// Experimental.
	WaitPeriodInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference
type jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) Alarm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) AlarmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alarmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ExtensionPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extensionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ExtensionPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"extensionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) InternalValue() *AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty {
	var returns *AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) WaitPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) WaitPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"waitPeriodInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.AwsCloudwatchCompositeAlarm.ActionsSuppressorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference_Override(a AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch.AwsCloudwatchCompositeAlarm.ActionsSuppressorPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetAlarm(val *string) {
	if err := j.validateSetAlarmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alarm",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetExtensionPeriod(val *float64) {
	if err := j.validateSetExtensionPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetInternalValue(val *AwsCloudwatchCompositeAlarm_ActionsSuppressorProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference)SetWaitPeriod(val *float64) {
	if err := j.validateSetWaitPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"waitPeriod",
		val,
	)
}

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudwatchCompositeAlarm_ActionsSuppressorPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

