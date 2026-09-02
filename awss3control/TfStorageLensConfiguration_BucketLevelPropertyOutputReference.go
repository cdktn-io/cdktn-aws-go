package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_BucketLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
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
	DetailedStatusCodeMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfStorageLensConfiguration_BucketLevelProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PrefixLevel() TfStorageLensConfiguration_PrefixLevelPropertyOutputReference
	// Experimental.
	PrefixLevelInput() *TfStorageLensConfiguration_PrefixLevelProperty
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
	PutActivityMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty)
	// Experimental.
	PutPrefixLevel(value *TfStorageLensConfiguration_PrefixLevelProperty)
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
	// Experimental.
	ResetPrefixLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStorageLensConfiguration_BucketLevelPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetrics() TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_BucketLevelProperty {
	var returns *TfStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevel() TfStorageLensConfiguration_PrefixLevelPropertyOutputReference {
	var returns TfStorageLensConfiguration_PrefixLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevelInput() *TfStorageLensConfiguration_PrefixLevelProperty {
	var returns *TfStorageLensConfiguration_PrefixLevelProperty
	_jsii_.Get(
		j,
		"prefixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_BucketLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_BucketLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_BucketLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_BucketLevelPropertyOutputReference_Override(t TfStorageLensConfiguration_BucketLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_BucketLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutActivityMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty) {
	if err := t.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty) {
	if err := t.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty) {
	if err := t.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty) {
	if err := t.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty) {
	if err := t.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) PutPrefixLevel(value *TfStorageLensConfiguration_PrefixLevelProperty) {
	if err := t.validatePutPrefixLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrefixLevel",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		t,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetPrefixLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefixLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_BucketLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

