package arcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/arcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/arcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference interface {
	cdktn.ComplexObject
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
	Service() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigServicePropertyList
	// Experimental.
	ServiceInput() interface{}
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
	Ungraceful() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigUngracefulPropertyList
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
	PutService(value interface{})
	// Experimental.
	PutUngraceful(value interface{})
	// Experimental.
	ResetService()
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

// The jsii proxy struct for AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference
type jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) CapacityMonitoringApproach() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) CapacityMonitoringApproachInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityMonitoringApproachInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) Service() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigServicePropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigServicePropertyList
	_jsii_.Get(
		j,
		"service",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ServiceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TargetPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TargetPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) Ungraceful() AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigUngracefulPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigUngracefulPropertyList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference_Override(a AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetCapacityMonitoringApproach(val *string) {
	if err := j.validateSetCapacityMonitoringApproachParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityMonitoringApproach",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetTargetPercent(val *float64) {
	if err := j.validateSetTargetPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetPercent",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) PutService(value interface{}) {
	if err := a.validatePutServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putService",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ResetService() {
	_jsii_.InvokeVoid(
		a,
		"resetService",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ResetTargetPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetTargetPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepEcsCapacityIncreaseConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

