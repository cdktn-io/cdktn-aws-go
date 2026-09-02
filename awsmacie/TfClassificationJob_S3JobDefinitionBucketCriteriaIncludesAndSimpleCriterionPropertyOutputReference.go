package awsmacie

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsmacie/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsmacie/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Comparator() *string
	// Experimental.
	SetComparator(val *string)
	// Experimental.
	ComparatorInput() *string
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
	InternalValue() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty
	// Experimental.
	SetInternalValue(val *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty)
	// Experimental.
	Key() *string
	// Experimental.
	SetKey(val *string)
	// Experimental.
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Values() *[]*string
	// Experimental.
	SetValues(val *[]*string)
	// Experimental.
	ValuesInput() *[]*string
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
	ResetComparator()
	// Experimental.
	ResetKey()
	// Experimental.
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference
type jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) Comparator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ComparatorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"comparatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) InternalValue() *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty {
	var returns *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference_Override(t TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-macie.TfClassificationJob.S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetComparator(val *string) {
	if err := j.validateSetComparatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"comparator",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetInternalValue(val *TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ResetComparator() {
	_jsii_.InvokeVoid(
		t,
		"resetComparator",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		t,
		"resetKey",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		t,
		"resetValues",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfClassificationJob_S3JobDefinitionBucketCriteriaIncludesAndSimpleCriterionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

