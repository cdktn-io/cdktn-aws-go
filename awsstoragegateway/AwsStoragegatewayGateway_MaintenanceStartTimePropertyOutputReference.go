package awsstoragegateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsstoragegateway/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference interface {
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
	DayOfMonth() *string
	// Experimental.
	SetDayOfMonth(val *string)
	// Experimental.
	DayOfMonthInput() *string
	// Experimental.
	DayOfWeek() *string
	// Experimental.
	SetDayOfWeek(val *string)
	// Experimental.
	DayOfWeekInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	HourOfDay() *float64
	// Experimental.
	SetHourOfDay(val *float64)
	// Experimental.
	HourOfDayInput() *float64
	// Experimental.
	InternalValue() *AwsStoragegatewayGateway_MaintenanceStartTimeProperty
	// Experimental.
	SetInternalValue(val *AwsStoragegatewayGateway_MaintenanceStartTimeProperty)
	// Experimental.
	MinuteOfHour() *float64
	// Experimental.
	SetMinuteOfHour(val *float64)
	// Experimental.
	MinuteOfHourInput() *float64
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
	ResetDayOfMonth()
	// Experimental.
	ResetDayOfWeek()
	// Experimental.
	ResetMinuteOfHour()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference
type jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) DayOfMonth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) DayOfMonthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) DayOfWeek() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) DayOfWeekInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dayOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) HourOfDay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) HourOfDayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"hourOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) InternalValue() *AwsStoragegatewayGateway_MaintenanceStartTimeProperty {
	var returns *AwsStoragegatewayGateway_MaintenanceStartTimeProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) MinuteOfHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minuteOfHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) MinuteOfHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minuteOfHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayGateway.MaintenanceStartTimePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference_Override(a AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-storage-gateway.AwsStoragegatewayGateway.MaintenanceStartTimePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetDayOfMonth(val *string) {
	if err := j.validateSetDayOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfMonth",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetDayOfWeek(val *string) {
	if err := j.validateSetDayOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dayOfWeek",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetHourOfDay(val *float64) {
	if err := j.validateSetHourOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hourOfDay",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetInternalValue(val *AwsStoragegatewayGateway_MaintenanceStartTimeProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetMinuteOfHour(val *float64) {
	if err := j.validateSetMinuteOfHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minuteOfHour",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ResetDayOfMonth() {
	_jsii_.InvokeVoid(
		a,
		"resetDayOfMonth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ResetDayOfWeek() {
	_jsii_.InvokeVoid(
		a,
		"resetDayOfWeek",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ResetMinuteOfHour() {
	_jsii_.InvokeVoid(
		a,
		"resetMinuteOfHour",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStoragegatewayGateway_MaintenanceStartTimePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

