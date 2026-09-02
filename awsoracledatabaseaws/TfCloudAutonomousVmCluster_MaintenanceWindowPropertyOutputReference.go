package awsoracledatabaseaws

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsoracledatabaseaws/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference interface {
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
	DaysOfWeek() TfCloudAutonomousVmCluster_DaysOfWeekPropertyList
	// Experimental.
	DaysOfWeekInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HoursOfDay() *[]*float64
	// Experimental.
	SetHoursOfDay(val *[]*float64)
	// Experimental.
	HoursOfDayInput() *[]*float64
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LeadTimeInWeeks() *float64
	// Experimental.
	SetLeadTimeInWeeks(val *float64)
	// Experimental.
	LeadTimeInWeeksInput() *float64
	// Experimental.
	Months() TfCloudAutonomousVmCluster_MonthsPropertyList
	// Experimental.
	MonthsInput() interface{}
	// Experimental.
	Preference() *string
	// Experimental.
	SetPreference(val *string)
	// Experimental.
	PreferenceInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	WeeksOfMonth() *[]*float64
	// Experimental.
	SetWeeksOfMonth(val *[]*float64)
	// Experimental.
	WeeksOfMonthInput() *[]*float64
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
	PutDaysOfWeek(value interface{})
	// Experimental.
	PutMonths(value interface{})
	// Experimental.
	ResetDaysOfWeek()
	// Experimental.
	ResetHoursOfDay()
	// Experimental.
	ResetLeadTimeInWeeks()
	// Experimental.
	ResetMonths()
	// Experimental.
	ResetWeeksOfMonth()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference
type jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) DaysOfWeek() TfCloudAutonomousVmCluster_DaysOfWeekPropertyList {
	var returns TfCloudAutonomousVmCluster_DaysOfWeekPropertyList
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) DaysOfWeekInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) HoursOfDay() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) HoursOfDayInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"hoursOfDayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) LeadTimeInWeeks() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) LeadTimeInWeeksInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leadTimeInWeeksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) Months() TfCloudAutonomousVmCluster_MonthsPropertyList {
	var returns TfCloudAutonomousVmCluster_MonthsPropertyList
	_jsii_.Get(
		j,
		"months",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) MonthsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monthsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) Preference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) PreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) WeeksOfMonth() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) WeeksOfMonthInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"weeksOfMonthInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster.MaintenanceWindowPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference_Override(t TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-oracle-database-aws.TfCloudAutonomousVmCluster.MaintenanceWindowPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetHoursOfDay(val *[]*float64) {
	if err := j.validateSetHoursOfDayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hoursOfDay",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetLeadTimeInWeeks(val *float64) {
	if err := j.validateSetLeadTimeInWeeksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leadTimeInWeeks",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetPreference(val *string) {
	if err := j.validateSetPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preference",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference)SetWeeksOfMonth(val *[]*float64) {
	if err := j.validateSetWeeksOfMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"weeksOfMonth",
		val,
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) PutDaysOfWeek(value interface{}) {
	if err := t.validatePutDaysOfWeekParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDaysOfWeek",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) PutMonths(value interface{}) {
	if err := t.validatePutMonthsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonths",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ResetDaysOfWeek() {
	_jsii_.InvokeVoid(
		t,
		"resetDaysOfWeek",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ResetHoursOfDay() {
	_jsii_.InvokeVoid(
		t,
		"resetHoursOfDay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ResetLeadTimeInWeeks() {
	_jsii_.InvokeVoid(
		t,
		"resetLeadTimeInWeeks",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ResetMonths() {
	_jsii_.InvokeVoid(
		t,
		"resetMonths",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ResetWeeksOfMonth() {
	_jsii_.InvokeVoid(
		t,
		"resetWeeksOfMonth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCloudAutonomousVmCluster_MaintenanceWindowPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

