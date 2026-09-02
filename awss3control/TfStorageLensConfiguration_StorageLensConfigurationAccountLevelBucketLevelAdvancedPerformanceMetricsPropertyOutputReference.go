package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference interface {
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
	// Experimental.
	SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty)
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
	ResetEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference
type jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) InternalValue() *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty {
	var returns *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference_Override(t TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.TfStorageLensConfiguration.StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetInternalValue(val *TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelAdvancedPerformanceMetricsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

