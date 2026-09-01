package awss3control

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3control/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3control/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference interface {
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
	DaysAfterInitiation() *float64
	// Experimental.
	SetDaysAfterInitiation(val *float64)
	// Experimental.
	DaysAfterInitiationInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty
	// Experimental.
	SetInternalValue(val *AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference
type jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) DaysAfterInitiation() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysAfterInitiation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) DaysAfterInitiationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daysAfterInitiationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) InternalValue() *AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty {
	var returns *AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlBucketLifecycleConfiguration.AbortIncompleteMultipartUploadPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference_Override(a AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3-control.AwsS3ControlBucketLifecycleConfiguration.AbortIncompleteMultipartUploadPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetDaysAfterInitiation(val *float64) {
	if err := j.validateSetDaysAfterInitiationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysAfterInitiation",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetInternalValue(val *AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3ControlBucketLifecycleConfiguration_AbortIncompleteMultipartUploadPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

