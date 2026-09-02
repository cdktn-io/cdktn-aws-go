package awss3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awss3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awss3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference interface {
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
	InternalValue() *TfBucketLogging_TargetObjectKeyFormatProperty
	// Experimental.
	SetInternalValue(val *TfBucketLogging_TargetObjectKeyFormatProperty)
	// Experimental.
	PartitionedPrefix() TfBucketLogging_PartitionedPrefixPropertyOutputReference
	// Experimental.
	PartitionedPrefixInput() *TfBucketLogging_PartitionedPrefixProperty
	// Experimental.
	SimplePrefix() TfBucketLogging_SimplePrefixPropertyOutputReference
	// Experimental.
	SimplePrefixInput() *TfBucketLogging_SimplePrefixProperty
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
	PutPartitionedPrefix(value *TfBucketLogging_PartitionedPrefixProperty)
	// Experimental.
	PutSimplePrefix(value *TfBucketLogging_SimplePrefixProperty)
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

// The jsii proxy struct for TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference
type jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InternalValue() *TfBucketLogging_TargetObjectKeyFormatProperty {
	var returns *TfBucketLogging_TargetObjectKeyFormatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefix() TfBucketLogging_PartitionedPrefixPropertyOutputReference {
	var returns TfBucketLogging_PartitionedPrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"partitionedPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefixInput() *TfBucketLogging_PartitionedPrefixProperty {
	var returns *TfBucketLogging_PartitionedPrefixProperty
	_jsii_.Get(
		j,
		"partitionedPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefix() TfBucketLogging_SimplePrefixPropertyOutputReference {
	var returns TfBucketLogging_SimplePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"simplePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefixInput() *TfBucketLogging_SimplePrefixProperty {
	var returns *TfBucketLogging_SimplePrefixProperty
	_jsii_.Get(
		j,
		"simplePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBucketLogging_TargetObjectKeyFormatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBucketLogging_TargetObjectKeyFormatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBucketLogging_TargetObjectKeyFormatPropertyOutputReference_Override(t TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.TfBucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetInternalValue(val *TfBucketLogging_TargetObjectKeyFormatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutPartitionedPrefix(value *TfBucketLogging_PartitionedPrefixProperty) {
	if err := t.validatePutPartitionedPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPartitionedPrefix",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutSimplePrefix(value *TfBucketLogging_SimplePrefixProperty) {
	if err := t.validatePutSimplePrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSimplePrefix",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetPartitionedPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPartitionedPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetSimplePrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetSimplePrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

