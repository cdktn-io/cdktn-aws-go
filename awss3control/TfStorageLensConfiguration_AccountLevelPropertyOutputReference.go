package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_AccountLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	// Experimental.
	BucketLevel() TfStorageLensConfiguration_BucketLevelPropertyOutputReference
	// Experimental.
	BucketLevelInput() *TfStorageLensConfiguration_BucketLevelProperty
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
	DetailedStatusCodeMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_AccountLevelProperty)
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
	PutActivityMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutBucketLevel(value *TfStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty)
	// Experimental.
	ResetActivityMetrics()
	// Experimental.
	ResetAdvancedCostOptimizationMetrics()
	// Experimental.
	ResetAdvancedDataProtectionMetrics()
	// Experimental.
	ResetAdvancedPerformanceMetrics()
	// Experimental.
	ResetDetailedStatusCodeMetrics()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStorageLensConfiguration_AccountLevelPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevel() TfStorageLensConfiguration_BucketLevelPropertyOutputReference {
	var returns TfStorageLensConfiguration_BucketLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevelInput() *TfStorageLensConfiguration_BucketLevelProperty {
	var returns *TfStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"bucketLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_AccountLevelProperty {
	var returns *TfStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_AccountLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_AccountLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_AccountLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_AccountLevelPropertyOutputReference_Override(t TfStorageLensConfiguration_AccountLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_AccountLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutActivityMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty) {
	if err := t.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty) {
	if err := t.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty) {
	if err := t.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty) {
	if err := t.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutBucketLevel(value *TfStorageLensConfiguration_BucketLevelProperty) {
	if err := t.validatePutBucketLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putBucketLevel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty) {
	if err := t.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_AccountLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

