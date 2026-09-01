package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
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
	DetailedStatusCodeMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsS3ControlStorageLensConfiguration_BucketLevelProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PrefixLevel() AwsS3ControlStorageLensConfiguration_PrefixLevelPropertyOutputReference
	// Experimental.
	PrefixLevelInput() *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty
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
	PutActivityMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty)
	// Experimental.
	PutPrefixLevel(value *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty)
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

// The jsii proxy struct for AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference
type jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ActivityMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) InternalValue() *AwsS3ControlStorageLensConfiguration_BucketLevelProperty {
	var returns *AwsS3ControlStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevel() AwsS3ControlStorageLensConfiguration_PrefixLevelPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_PrefixLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"prefixLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PrefixLevelInput() *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty {
	var returns *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty
	_jsii_.Get(
		j,
		"prefixLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference_Override(a AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.BucketLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference)SetInternalValue(val *AwsS3ControlStorageLensConfiguration_BucketLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutActivityMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelActivityMetricsProperty) {
	if err := a.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedCostOptimizationMetricsProperty) {
	if err := a.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedDataProtectionMetricsProperty) {
	if err := a.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty) {
	if err := a.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty) {
	if err := a.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) PutPrefixLevel(value *AwsS3ControlStorageLensConfiguration_PrefixLevelProperty) {
	if err := a.validatePutPrefixLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefixLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ResetPrefixLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefixLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

