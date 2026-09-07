package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_AccountLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	// Experimental.
	BucketLevel() AwsStorageLensConfiguration_BucketLevelPropertyOutputReference
	// Experimental.
	BucketLevelInput() *AwsStorageLensConfiguration_BucketLevelProperty
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
	DetailedStatusCodeMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_AccountLevelProperty)
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
	PutActivityMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutBucketLevel(value *AwsStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty)
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

// The jsii proxy struct for AwsStorageLensConfiguration_AccountLevelPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevel() AwsStorageLensConfiguration_BucketLevelPropertyOutputReference {
	var returns AwsStorageLensConfiguration_BucketLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevelInput() *AwsStorageLensConfiguration_BucketLevelProperty {
	var returns *AwsStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"bucketLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetrics() AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_AccountLevelProperty {
	var returns *AwsStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_AccountLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_AccountLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_AccountLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_AccountLevelPropertyOutputReference_Override(a AwsStorageLensConfiguration_AccountLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_AccountLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutActivityMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty) {
	if err := a.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty) {
	if err := a.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty) {
	if err := a.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty) {
	if err := a.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutBucketLevel(value *AwsStorageLensConfiguration_BucketLevelProperty) {
	if err := a.validatePutBucketLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBucketLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty) {
	if err := a.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_AccountLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

