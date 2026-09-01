package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference interface {
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
	InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigProperty)
	// Experimental.
	MonitoringJobDefinition() AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
	// Experimental.
	MonitoringJobDefinitionInput() *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty
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
	ScheduleConfig() AwsSagemakerMonitoringSchedule_ScheduleConfigPropertyOutputReference
	// Experimental.
	ScheduleConfigInput() *AwsSagemakerMonitoringSchedule_ScheduleConfigProperty
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
	PutMonitoringJobDefinition(value *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty)
	// Experimental.
	PutScheduleConfig(value *AwsSagemakerMonitoringSchedule_ScheduleConfigProperty)
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

// The jsii proxy struct for AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InternalValue() *AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinition() AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionPropertyOutputReference
	_jsii_.Get(
		j,
		"monitoringJobDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionInput() *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty {
	var returns *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty
	_jsii_.Get(
		j,
		"monitoringJobDefinitionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringJobDefinitionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringJobDefinitionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) MonitoringTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"monitoringTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ScheduleConfig() AwsSagemakerMonitoringSchedule_ScheduleConfigPropertyOutputReference {
	var returns AwsSagemakerMonitoringSchedule_ScheduleConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"scheduleConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ScheduleConfigInput() *AwsSagemakerMonitoringSchedule_ScheduleConfigProperty {
	var returns *AwsSagemakerMonitoringSchedule_ScheduleConfigProperty
	_jsii_.Get(
		j,
		"scheduleConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringScheduleConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference_Override(a AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerMonitoringSchedule.MonitoringScheduleConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetMonitoringJobDefinitionName(val *string) {
	if err := j.validateSetMonitoringJobDefinitionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringJobDefinitionName",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetMonitoringType(val *string) {
	if err := j.validateSetMonitoringTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"monitoringType",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) PutMonitoringJobDefinition(value *AwsSagemakerMonitoringSchedule_MonitoringJobDefinitionProperty) {
	if err := a.validatePutMonitoringJobDefinitionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMonitoringJobDefinition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) PutScheduleConfig(value *AwsSagemakerMonitoringSchedule_ScheduleConfigProperty) {
	if err := a.validatePutScheduleConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScheduleConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetMonitoringJobDefinition() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringJobDefinition",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetMonitoringJobDefinitionName() {
	_jsii_.InvokeVoid(
		a,
		"resetMonitoringJobDefinitionName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ResetScheduleConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetScheduleConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerMonitoringSchedule_MonitoringScheduleConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

