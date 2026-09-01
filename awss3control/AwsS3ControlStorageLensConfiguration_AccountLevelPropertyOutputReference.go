package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ActivityMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	// Experimental.
	ActivityMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	// Experimental.
	AdvancedCostOptimizationMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	// Experimental.
	AdvancedCostOptimizationMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	// Experimental.
	AdvancedDataProtectionMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	// Experimental.
	AdvancedDataProtectionMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	// Experimental.
	AdvancedPerformanceMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	// Experimental.
	AdvancedPerformanceMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	// Experimental.
	BucketLevel() AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference
	// Experimental.
	BucketLevelInput() *AwsS3ControlStorageLensConfiguration_BucketLevelProperty
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
	DetailedStatusCodeMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	// Experimental.
	DetailedStatusCodeMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsS3ControlStorageLensConfiguration_AccountLevelProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlStorageLensConfiguration_AccountLevelProperty)
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
	PutActivityMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty)
	// Experimental.
	PutAdvancedCostOptimizationMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty)
	// Experimental.
	PutAdvancedDataProtectionMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty)
	// Experimental.
	PutAdvancedPerformanceMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty)
	// Experimental.
	PutBucketLevel(value *AwsS3ControlStorageLensConfiguration_BucketLevelProperty)
	// Experimental.
	PutDetailedStatusCodeMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty)
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

// The jsii proxy struct for AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference
type jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"activityMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ActivityMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty
	_jsii_.Get(
		j,
		"activityMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedCostOptimizationMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty
	_jsii_.Get(
		j,
		"advancedCostOptimizationMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedDataProtectionMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedDataProtectionMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty
	_jsii_.Get(
		j,
		"advancedDataProtectionMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"advancedPerformanceMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) AdvancedPerformanceMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"advancedPerformanceMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevel() AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_BucketLevelPropertyOutputReference
	_jsii_.Get(
		j,
		"bucketLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) BucketLevelInput() *AwsS3ControlStorageLensConfiguration_BucketLevelProperty {
	var returns *AwsS3ControlStorageLensConfiguration_BucketLevelProperty
	_jsii_.Get(
		j,
		"bucketLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetrics() AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference {
	var returns AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsPropertyOutputReference
	_jsii_.Get(
		j,
		"detailedStatusCodeMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) DetailedStatusCodeMetricsInput() *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty {
	var returns *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"detailedStatusCodeMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) InternalValue() *AwsS3ControlStorageLensConfiguration_AccountLevelProperty {
	var returns *AwsS3ControlStorageLensConfiguration_AccountLevelProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference_Override(a AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlStorageLensConfiguration.AccountLevelPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference)SetInternalValue(val *AwsS3ControlStorageLensConfiguration_AccountLevelProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutActivityMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelActivityMetricsProperty) {
	if err := a.validatePutActivityMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActivityMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedCostOptimizationMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedCostOptimizationMetricsProperty) {
	if err := a.validatePutAdvancedCostOptimizationMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedCostOptimizationMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedDataProtectionMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedDataProtectionMetricsProperty) {
	if err := a.validatePutAdvancedDataProtectionMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedDataProtectionMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutAdvancedPerformanceMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelAdvancedPerformanceMetricsProperty) {
	if err := a.validatePutAdvancedPerformanceMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdvancedPerformanceMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutBucketLevel(value *AwsS3ControlStorageLensConfiguration_BucketLevelProperty) {
	if err := a.validatePutBucketLevelParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBucketLevel",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) PutDetailedStatusCodeMetrics(value *AwsS3ControlStorageLensConfiguration_StorageLensConfigurationAccountLevelDetailedStatusCodeMetricsProperty) {
	if err := a.validatePutDetailedStatusCodeMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDetailedStatusCodeMetrics",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetActivityMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetActivityMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedCostOptimizationMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedCostOptimizationMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedDataProtectionMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedDataProtectionMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetAdvancedPerformanceMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetAdvancedPerformanceMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ResetDetailedStatusCodeMetrics() {
	_jsii_.InvokeVoid(
		a,
		"resetDetailedStatusCodeMetrics",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlStorageLensConfiguration_AccountLevelPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

