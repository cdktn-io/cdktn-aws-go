package s3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference interface {
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
	InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	// Experimental.
	SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty)
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

// The jsii proxy struct for AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference
type jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) InternalValue() *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty {
	var returns *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference_Override(a AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsStorageLensConfiguration.StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetInternalValue(val *AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsStorageLensConfiguration_StorageLensConfigurationAccountLevelBucketLevelDetailedStatusCodeMetricsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

