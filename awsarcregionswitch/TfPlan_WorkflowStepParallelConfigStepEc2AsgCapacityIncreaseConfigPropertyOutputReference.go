package awsarcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsarcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Asg() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsgPropertyList
	// Experimental.
	AsgInput() interface{}
	// Experimental.
	CapacityMonitoringApproach() *string
	// Experimental.
	SetCapacityMonitoringApproach(val *string)
	// Experimental.
	CapacityMonitoringApproachInput() *string
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
	TargetPercent() *float64
	// Experimental.
	SetTargetPercent(val *float64)
	// Experimental.
	TargetPercentInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimeoutMinutes() *float64
	// Experimental.
	SetTimeoutMinutes(val *float64)
	// Experimental.
	TimeoutMinutesInput() *float64
	// Experimental.
	Ungraceful() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngracefulPropertyList
	// Experimental.
	UngracefulInput() interface{}
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
	PutAsg(value interface{})
	// Experimental.
	PutUngraceful(value interface{})
	// Experimental.
	ResetAsg()
	// Experimental.
	ResetTargetPercent()
	// Experimental.
	ResetTimeoutMinutes()
	// Experimental.
	ResetUngraceful()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference
type jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) Asg() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsgPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigAsgPropertyList
	_jsii_.Get(
		j,
		"asg",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) AsgInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"asgInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) CapacityMonitoringApproach() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) CapacityMonitoringApproachInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TargetPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TargetPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) Ungraceful() TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngracefulPropertyList {
	var returns TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigUngracefulPropertyList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference_Override(t TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.TfPlan.WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetCapacityMonitoringApproach(val *string) {
	if err := j.validateSetCapacityMonitoringApproachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMonitoringApproach",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetTargetPercent(val *float64) {
	if err := j.validateSetTargetPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPercent",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) PutAsg(value interface{}) {
	if err := t.validatePutAsgParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAsg",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) PutUngraceful(value interface{}) {
	if err := t.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ResetAsg() {
	_jsii_.InvokeVoid(
		t,
		"resetAsg",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ResetTargetPercent() {
	_jsii_.InvokeVoid(
		t,
		"resetTargetPercent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		t,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPlan_WorkflowStepParallelConfigStepEc2AsgCapacityIncreaseConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

