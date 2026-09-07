package arcregionswitch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/arcregionswitch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/arcregionswitch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Lambda() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaPropertyList
	// Experimental.
	LambdaInput() interface{}
	// Experimental.
	RegionToRun() *string
	// Experimental.
	SetRegionToRun(val *string)
	// Experimental.
	RegionToRunInput() *string
	// Experimental.
	RetryIntervalMinutes() *float64
	// Experimental.
	SetRetryIntervalMinutes(val *float64)
	// Experimental.
	RetryIntervalMinutesInput() *float64
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
	Ungraceful() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList
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
	PutLambda(value interface{})
	// Experimental.
	PutUngraceful(value interface{})
	// Experimental.
	ResetLambda()
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

// The jsii proxy struct for AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference
type jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) Lambda() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigLambdaPropertyList
	_jsii_.Get(
		j,
		"lambda",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) LambdaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lambdaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) RegionToRun() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionToRun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) RegionToRunInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionToRunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) RetryIntervalMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) RetryIntervalMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) TimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) TimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) Ungraceful() AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList {
	var returns AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigUngracefulPropertyList
	_jsii_.Get(
		j,
		"ungraceful",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) UngracefulInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ungracefulInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference_Override(a AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-arc-region-switch.AwsPlan.WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetRegionToRun(val *string) {
	if err := j.validateSetRegionToRunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionToRun",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetRetryIntervalMinutes(val *float64) {
	if err := j.validateSetRetryIntervalMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryIntervalMinutes",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference)SetTimeoutMinutes(val *float64) {
	if err := j.validateSetTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutMinutes",
		val,
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) PutLambda(value interface{}) {
	if err := a.validatePutLambdaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLambda",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) PutUngraceful(value interface{}) {
	if err := a.validatePutUngracefulParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUngraceful",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ResetLambda() {
	_jsii_.InvokeVoid(
		a,
		"resetLambda",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ResetTimeoutMinutes() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutMinutes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ResetUngraceful() {
	_jsii_.InvokeVoid(
		a,
		"resetUngraceful",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPlan_WorkflowStepParallelConfigStepCustomActionLambdaConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

