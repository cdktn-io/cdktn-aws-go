package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference interface {
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
	InternalValue() *TfMonitoringSchedule_MonitoringScheduleConfigProperty
	// Experimental.
	SetInternalValue(val *TfMonitoringSchedule_MonitoringScheduleConfigProperty)
	// Experimental.
	MonitoringJobDefinition() TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
	// Experimental.
	MonitoringJobDefinitionInput() *TfMonitoringSchedule_MonitoringJobDefinitionProperty
	// Experimental.
	MonitoringJobDefinitionName() *string
	// Experimental.
	SetMonitoringJobDefinitionName(val *string)
	// Experimental.
	MonitoringJobDefinitionNameInput() *string
	// Experimental.
	MonitoringType() *string
	// Experimental.
	SetMonitoringType(val *string)
	// Experimental.
	MonitoringTypeInput() *string
	// Experimental.
	ScheduleConfig() TfMonitoringSchedule_ScheduleConfigPropertyOutputReference
	// Experimental.
	ScheduleConfigInput() *TfMonitoringSchedule_ScheduleConfigProperty
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
	PutMonitoringJobDefinition(value *TfMonitoringSchedule_MonitoringJobDefinitionProperty)
	// Experimental.
	PutScheduleConfig(value *TfMonitoringSchedule_ScheduleConfigProperty)
	// Experimental.
	ResetMonitoringJobDefinition()
	// Experimental.
	ResetMonitoringJobDefinitionName()
	// Experimental.
	ResetScheduleConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference
type jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InternalValue() *TfMonitoringSchedule_MonitoringScheduleConfigProperty {
	var returns *TfMonitoringSchedule_MonitoringScheduleConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinition() TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference {
	var returns TfMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringJobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionInput() *TfMonitoringSchedule_MonitoringJobDefinitionProperty {
	var returns *TfMonitoringSchedule_MonitoringJobDefinitionProperty
	_jsii_.Get(
		j,
		"monitoringJobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ScheduleConfig() TfMonitoringSchedule_ScheduleConfigPropertyOutputReference {
	var returns TfMonitoringSchedule_ScheduleConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scheduleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ScheduleConfigInput() *TfMonitoringSchedule_ScheduleConfigProperty {
	var returns *TfMonitoringSchedule_ScheduleConfigProperty
	_jsii_.Get(
		j,
		"scheduleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringScheduleConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference_Override(t TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfMonitoringSchedule.MonitoringScheduleConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetInternalValue(val *TfMonitoringSchedule_MonitoringScheduleConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetMonitoringJobDefinitionName(val *string) {
	if err := j.validateSetMonitoringJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringJobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetMonitoringType(val *string) {
	if err := j.validateSetMonitoringTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringType",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) PutMonitoringJobDefinition(value *TfMonitoringSchedule_MonitoringJobDefinitionProperty) {
	if err := t.validatePutMonitoringJobDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMonitoringJobDefinition",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) PutScheduleConfig(value *TfMonitoringSchedule_ScheduleConfigProperty) {
	if err := t.validatePutScheduleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putScheduleConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetMonitoringJobDefinition() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringJobDefinition",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetMonitoringJobDefinitionName() {
	_jsii_.InvokeVoid(
		t,
		"resetMonitoringJobDefinitionName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetScheduleConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetScheduleConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

