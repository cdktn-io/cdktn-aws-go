package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsS3BucketLogging_TargetObjectKeyFormatProperty
	// Experimental.
	SetInternalValue(val *AwsS3BucketLogging_TargetObjectKeyFormatProperty)
	// Experimental.
	PartitionedPrefix() AwsS3BucketLogging_PartitionedPrefixPropertyOutputReference
	// Experimental.
	PartitionedPrefixInput() *AwsS3BucketLogging_PartitionedPrefixProperty
	// Experimental.
	SimplePrefix() AwsS3BucketLogging_SimplePrefixPropertyOutputReference
	// Experimental.
	SimplePrefixInput() *AwsS3BucketLogging_SimplePrefixProperty
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
	PutPartitionedPrefix(value *AwsS3BucketLogging_PartitionedPrefixProperty)
	// Experimental.
	PutSimplePrefix(value *AwsS3BucketLogging_SimplePrefixProperty)
	// Experimental.
	ResetPartitionedPrefix()
	// Experimental.
	ResetSimplePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference
type jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) InternalValue() *AwsS3BucketLogging_TargetObjectKeyFormatProperty {
	var returns *AwsS3BucketLogging_TargetObjectKeyFormatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefix() AwsS3BucketLogging_PartitionedPrefixPropertyOutputReference {
	var returns AwsS3BucketLogging_PartitionedPrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"partitionedPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefixInput() *AwsS3BucketLogging_PartitionedPrefixProperty {
	var returns *AwsS3BucketLogging_PartitionedPrefixProperty
	_jsii_.Get(
		j,
		"partitionedPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefix() AwsS3BucketLogging_SimplePrefixPropertyOutputReference {
	var returns AwsS3BucketLogging_SimplePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"simplePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefixInput() *AwsS3BucketLogging_SimplePrefixProperty {
	var returns *AwsS3BucketLogging_SimplePrefixProperty
	_jsii_.Get(
		j,
		"simplePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference_Override(a AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsS3BucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetInternalValue(val *AwsS3BucketLogging_TargetObjectKeyFormatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutPartitionedPrefix(value *AwsS3BucketLogging_PartitionedPrefixProperty) {
	if err := a.validatePutPartitionedPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPartitionedPrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutSimplePrefix(value *AwsS3BucketLogging_SimplePrefixProperty) {
	if err := a.validatePutSimplePrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSimplePrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetPartitionedPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPartitionedPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetSimplePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetSimplePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsS3BucketLogging_TargetObjectKeyFormatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

