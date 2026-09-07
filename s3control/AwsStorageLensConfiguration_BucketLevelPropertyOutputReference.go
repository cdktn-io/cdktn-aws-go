package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_BucketLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
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
	DetailedStatusCodeMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsStorageLensConfiguration_BucketLevelProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PrefixLevel() AwsStorageLensConfiguration_PrefixLevelPropertyOutputReference
	// Experimental.
	PrefixLevelInput() *AwsStorageLensConfiguration_PrefixLevelProperty
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
	PutActivityMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty)
	// Experimental.
	PutPrefixLevel(value *AwsStorageLensConfiguration_PrefixLevelProperty)
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

// The jsii proxy struct for AwsStorageLensConfiguration_BucketLevelPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_BucketLevelProperty {
	var returns *AwsStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevel() AwsStorageLensConfiguration_PrefixLevelPropertyOutputReference {
	var returns AwsStorageLensConfiguration_PrefixLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevelInput() *AwsStorageLensConfiguration_PrefixLevelProperty {
	var returns *AwsStorageLensConfiguration_PrefixLevelProperty
	_jsii_.Get(
		j,
		"prefixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_BucketLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_BucketLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_BucketLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_BucketLevelPropertyOutputReference_Override(a AwsStorageLensConfiguration_BucketLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_BucketLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutActivityMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty) {
	if err := a.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty) {
	if err := a.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty) {
	if err := a.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty) {
	if err := a.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty) {
	if err := a.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) PutPrefixLevel(value *AwsStorageLensConfiguration_PrefixLevelProperty) {
	if err := a.validatePutPrefixLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefixLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetPrefixLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_BucketLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

