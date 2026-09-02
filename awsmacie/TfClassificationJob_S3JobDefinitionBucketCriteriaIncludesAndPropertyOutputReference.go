package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	SimpleCriterion() TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference
	// Experimental.
	SimpleCriterionInput() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty
	// Experimental.
	TagCriterion() TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionPropertyOutputReference
	// Experimental.
	TagCriterionInput() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty
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
	PutSimpleCriterion(value *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty)
	// Experimental.
	PutTagCriterion(value *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty)
	// Experimental.
	ResetSimpleCriterion()
	// Experimental.
	ResetTagCriterion()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference
type jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) SimpleCriterion() TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference {
	var returns TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference
	_jsii_.Get(
		j,
		"simpleCriterion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) SimpleCriterionInput() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty {
	var returns *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty
	_jsii_.Get(
		j,
		"simpleCriterionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) TagCriterion() TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionPropertyOutputReference {
	var returns TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionPropertyOutputReference
	_jsii_.Get(
		j,
		"tagCriterion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) TagCriterionInput() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty {
	var returns *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty
	_jsii_.Get(
		j,
		"tagCriterionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference_Override(t TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) PutSimpleCriterion(value *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty) {
	if err := t.validatePutSimpleCriterionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSimpleCriterion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) PutTagCriterion(value *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndTagCriterionProperty) {
	if err := t.validatePutTagCriterionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagCriterion",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ResetSimpleCriterion() {
	_jsii_.InvokeVoid(
		t,
		"resetSimpleCriterion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ResetTagCriterion() {
	_jsii_.InvokeVoid(
		t,
		"resetTagCriterion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

