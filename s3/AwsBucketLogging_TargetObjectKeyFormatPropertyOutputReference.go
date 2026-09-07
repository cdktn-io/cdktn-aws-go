package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference interface {
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
	InternalValue() *AwsBucketLogging_TargetObjectKeyFormatProperty
	// Experimental.
	SetInternalValue(val *AwsBucketLogging_TargetObjectKeyFormatProperty)
	// Experimental.
	PartitionedPrefix() AwsBucketLogging_PartitionedPrefixPropertyOutputReference
	// Experimental.
	PartitionedPrefixInput() *AwsBucketLogging_PartitionedPrefixProperty
	// Experimental.
	SimplePrefix() AwsBucketLogging_SimplePrefixPropertyOutputReference
	// Experimental.
	SimplePrefixInput() *AwsBucketLogging_SimplePrefixProperty
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
	PutPartitionedPrefix(value *AwsBucketLogging_PartitionedPrefixProperty)
	// Experimental.
	PutSimplePrefix(value *AwsBucketLogging_SimplePrefixProperty)
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

// The jsii proxy struct for AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference
type jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InternalValue() *AwsBucketLogging_TargetObjectKeyFormatProperty {
	var returns *AwsBucketLogging_TargetObjectKeyFormatProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefix() AwsBucketLogging_PartitionedPrefixPropertyOutputReference {
	var returns AwsBucketLogging_PartitionedPrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"partitionedPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PartitionedPrefixInput() *AwsBucketLogging_PartitionedPrefixProperty {
	var returns *AwsBucketLogging_PartitionedPrefixProperty
	_jsii_.Get(
		j,
		"partitionedPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefix() AwsBucketLogging_SimplePrefixPropertyOutputReference {
	var returns AwsBucketLogging_SimplePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"simplePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) SimplePrefixInput() *AwsBucketLogging_SimplePrefixProperty {
	var returns *AwsBucketLogging_SimplePrefixProperty
	_jsii_.Get(
		j,
		"simplePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucketLogging_TargetObjectKeyFormatPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference_Override(a AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLogging.TargetObjectKeyFormatPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetInternalValue(val *AwsBucketLogging_TargetObjectKeyFormatProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutPartitionedPrefix(value *AwsBucketLogging_PartitionedPrefixProperty) {
	if err := a.validatePutPartitionedPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPartitionedPrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) PutSimplePrefix(value *AwsBucketLogging_SimplePrefixProperty) {
	if err := a.validatePutSimplePrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSimplePrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetPartitionedPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPartitionedPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ResetSimplePrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetSimplePrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucketLogging_TargetObjectKeyFormatPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

