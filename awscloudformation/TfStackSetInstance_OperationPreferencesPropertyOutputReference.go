package awscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStackSetInstance_OperationPreferencesPropertyOutputReference interface {
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
	// Experimental.
	ConcurrencyMode() *string
	// Experimental.
	SetConcurrencyMode(val *string)
	// Experimental.
	ConcurrencyModeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FailureToleranceCount() *float64
	// Experimental.
	SetFailureToleranceCount(val *float64)
	// Experimental.
	FailureToleranceCountInput() *float64
	// Experimental.
	FailureTolerancePercentage() *float64
	// Experimental.
	SetFailureTolerancePercentage(val *float64)
	// Experimental.
	FailureTolerancePercentageInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfStackSetInstance_OperationPreferencesProperty
	// Experimental.
	SetInternalValue(val *TfStackSetInstance_OperationPreferencesProperty)
	// Experimental.
	MaxConcurrentCount() *float64
	// Experimental.
	SetMaxConcurrentCount(val *float64)
	// Experimental.
	MaxConcurrentCountInput() *float64
	// Experimental.
	MaxConcurrentPercentage() *float64
	// Experimental.
	SetMaxConcurrentPercentage(val *float64)
	// Experimental.
	MaxConcurrentPercentageInput() *float64
	// Experimental.
	RegionConcurrencyType() *string
	// Experimental.
	SetRegionConcurrencyType(val *string)
	// Experimental.
	RegionConcurrencyTypeInput() *string
	// Experimental.
	RegionOrder() *[]*string
	// Experimental.
	SetRegionOrder(val *[]*string)
	// Experimental.
	RegionOrderInput() *[]*string
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
	ResetConcurrencyMode()
	// Experimental.
	ResetFailureToleranceCount()
	// Experimental.
	ResetFailureTolerancePercentage()
	// Experimental.
	ResetMaxConcurrentCount()
	// Experimental.
	ResetMaxConcurrentPercentage()
	// Experimental.
	ResetRegionConcurrencyType()
	// Experimental.
	ResetRegionOrder()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStackSetInstance_OperationPreferencesPropertyOutputReference
type jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ConcurrencyMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ConcurrencyModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"concurrencyModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) InternalValue() *TfStackSetInstance_OperationPreferencesProperty {
	var returns *TfStackSetInstance_OperationPreferencesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) MaxConcurrentCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) MaxConcurrentCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) MaxConcurrentPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) MaxConcurrentPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrentPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) RegionConcurrencyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) RegionConcurrencyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionConcurrencyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) RegionOrder() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) RegionOrderInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStackSetInstance_OperationPreferencesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStackSetInstance_OperationPreferencesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStackSetInstance_OperationPreferencesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudformation.TfStackSetInstance.OperationPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStackSetInstance_OperationPreferencesPropertyOutputReference_Override(t TfStackSetInstance_OperationPreferencesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudformation.TfStackSetInstance.OperationPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetConcurrencyMode(val *string) {
	if err := j.validateSetConcurrencyModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"concurrencyMode",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetInternalValue(val *TfStackSetInstance_OperationPreferencesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetMaxConcurrentCount(val *float64) {
	if err := j.validateSetMaxConcurrentCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentCount",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetMaxConcurrentPercentage(val *float64) {
	if err := j.validateSetMaxConcurrentPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrentPercentage",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetRegionConcurrencyType(val *string) {
	if err := j.validateSetRegionConcurrencyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionConcurrencyType",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetRegionOrder(val *[]*string) {
	if err := j.validateSetRegionOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regionOrder",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetConcurrencyMode() {
	_jsii_.InvokeVoid(
		t,
		"resetConcurrencyMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetMaxConcurrentCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrentCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetMaxConcurrentPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrentPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetRegionConcurrencyType() {
	_jsii_.InvokeVoid(
		t,
		"resetRegionConcurrencyType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ResetRegionOrder() {
	_jsii_.InvokeVoid(
		t,
		"resetRegionOrder",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStackSetInstance_OperationPreferencesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

