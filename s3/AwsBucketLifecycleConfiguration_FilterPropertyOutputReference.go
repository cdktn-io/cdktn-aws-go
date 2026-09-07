package s3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/s3/jsii"

	"github.com/cdktn-io/cdktn-aws-go/s3/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBucketLifecycleConfiguration_FilterPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() AwsBucketLifecycleConfiguration_AndPropertyList
	// Experimental.
	AndInput() interface{}
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ObjectSizeGreaterThan() *float64
	// Experimental.
	SetObjectSizeGreaterThan(val *float64)
	// Experimental.
	ObjectSizeGreaterThanInput() *float64
	// Experimental.
	ObjectSizeLessThan() *float64
	// Experimental.
	SetObjectSizeLessThan(val *float64)
	// Experimental.
	ObjectSizeLessThanInput() *float64
	// Experimental.
	Prefix() *string
	// Experimental.
	SetPrefix(val *string)
	// Experimental.
	PrefixInput() *string
	// Experimental.
	Tag() AwsBucketLifecycleConfiguration_TagPropertyList
	// Experimental.
	TagInput() interface{}
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
	PutAnd(value interface{})
	// Experimental.
	PutTag(value interface{})
	// Experimental.
	ResetAnd()
	// Experimental.
	ResetObjectSizeGreaterThan()
	// Experimental.
	ResetObjectSizeLessThan()
	// Experimental.
	ResetPrefix()
	// Experimental.
	ResetTag()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBucketLifecycleConfiguration_FilterPropertyOutputReference
type jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) And() AwsBucketLifecycleConfiguration_AndPropertyList {
	var returns AwsBucketLifecycleConfiguration_AndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ObjectSizeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectSizeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ObjectSizeGreaterThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectSizeGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ObjectSizeLessThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectSizeLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ObjectSizeLessThanInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"objectSizeLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) Prefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) PrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) Tag() AwsBucketLifecycleConfiguration_TagPropertyList {
	var returns AwsBucketLifecycleConfiguration_TagPropertyList
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) TagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBucketLifecycleConfiguration_FilterPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBucketLifecycleConfiguration_FilterPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBucketLifecycleConfiguration_FilterPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLifecycleConfiguration.FilterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBucketLifecycleConfiguration_FilterPropertyOutputReference_Override(a AwsBucketLifecycleConfiguration_FilterPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-s3.AwsBucketLifecycleConfiguration.FilterPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetObjectSizeGreaterThan(val *float64) {
	if err := j.validateSetObjectSizeGreaterThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeGreaterThan",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetObjectSizeLessThan(val *float64) {
	if err := j.validateSetObjectSizeLessThanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"objectSizeLessThan",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetPrefix(val *string) {
	if err := j.validateSetPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"prefix",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) PutAnd(value interface{}) {
	if err := a.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnd",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) PutTag(value interface{}) {
	if err := a.validatePutTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		a,
		"resetAnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ResetObjectSizeGreaterThan() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectSizeGreaterThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ResetObjectSizeLessThan() {
	_jsii_.InvokeVoid(
		a,
		"resetObjectSizeLessThan",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		a,
		"resetTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBucketLifecycleConfiguration_FilterPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

