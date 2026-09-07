package macie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/macie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/macie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsClassificationJob_S3JobDefinitionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BucketCriteria() AwsClassificationJob_BucketCriteriaPropertyOutputReference
	// Experimental.
	BucketCriteriaInput() *AwsClassificationJob_BucketCriteriaProperty
	// Experimental.
	BucketDefinitions() AwsClassificationJob_BucketDefinitionsPropertyList
	// Experimental.
	BucketDefinitionsInput() interface{}
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
	InternalValue() *AwsClassificationJob_S3JobDefinitionProperty
	// Experimental.
	SetInternalValue(val *AwsClassificationJob_S3JobDefinitionProperty)
	// Experimental.
	Scoping() AwsClassificationJob_ScopingPropertyOutputReference
	// Experimental.
	ScopingInput() *AwsClassificationJob_ScopingProperty
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
	PutBucketCriteria(value *AwsClassificationJob_BucketCriteriaProperty)
	// Experimental.
	PutBucketDefinitions(value interface{})
	// Experimental.
	PutScoping(value *AwsClassificationJob_ScopingProperty)
	// Experimental.
	ResetBucketCriteria()
	// Experimental.
	ResetBucketDefinitions()
	// Experimental.
	ResetScoping()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsClassificationJob_S3JobDefinitionPropertyOutputReference
type jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) BucketCriteria() AwsClassificationJob_BucketCriteriaPropertyOutputReference {
	var returns AwsClassificationJob_BucketCriteriaPropertyOutputReference
	_jsii_.Get(
		j,
		"bucketCriteria",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) BucketCriteriaInput() *AwsClassificationJob_BucketCriteriaProperty {
	var returns *AwsClassificationJob_BucketCriteriaProperty
	_jsii_.Get(
		j,
		"bucketCriteriaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) BucketDefinitions() AwsClassificationJob_BucketDefinitionsPropertyList {
	var returns AwsClassificationJob_BucketDefinitionsPropertyList
	_jsii_.Get(
		j,
		"bucketDefinitions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) BucketDefinitionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bucketDefinitionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) InternalValue() *AwsClassificationJob_S3JobDefinitionProperty {
	var returns *AwsClassificationJob_S3JobDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) Scoping() AwsClassificationJob_ScopingPropertyOutputReference {
	var returns AwsClassificationJob_ScopingPropertyOutputReference
	_jsii_.Get(
		j,
		"scoping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ScopingInput() *AwsClassificationJob_ScopingProperty {
	var returns *AwsClassificationJob_ScopingProperty
	_jsii_.Get(
		j,
		"scopingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsClassificationJob_S3JobDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsClassificationJob_S3JobDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsClassificationJob_S3JobDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.AwsClassificationJob.S3JobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsClassificationJob_S3JobDefinitionPropertyOutputReference_Override(a AwsClassificationJob_S3JobDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.AwsClassificationJob.S3JobDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference)SetInternalValue(val *AwsClassificationJob_S3JobDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) PutBucketCriteria(value *AwsClassificationJob_BucketCriteriaProperty) {
	if err := a.validatePutBucketCriteriaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBucketCriteria",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) PutBucketDefinitions(value interface{}) {
	if err := a.validatePutBucketDefinitionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putBucketDefinitions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) PutScoping(value *AwsClassificationJob_ScopingProperty) {
	if err := a.validatePutScopingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScoping",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ResetBucketCriteria() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketCriteria",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ResetBucketDefinitions() {
	_jsii_.InvokeVoid(
		a,
		"resetBucketDefinitions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ResetScoping() {
	_jsii_.InvokeVoid(
		a,
		"resetScoping",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsClassificationJob_S3JobDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

