package awswebservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBudget_FilterExpressionAndPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() TfBudget_FilterExpressionAndAndPropertyList
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
	// Experimental.
	CostCategories() TfBudget_FilterExpressionAndCostCategoriesPropertyOutputReference
	// Experimental.
	CostCategoriesInput() *TfBudget_FilterExpressionAndCostCategoriesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimensions() TfBudget_FilterExpressionAndDimensionsPropertyOutputReference
	// Experimental.
	DimensionsInput() *TfBudget_FilterExpressionAndDimensionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Not() TfBudget_FilterExpressionAndNotPropertyOutputReference
	// Experimental.
	NotInput() *TfBudget_FilterExpressionAndNotProperty
	// Experimental.
	Or() TfBudget_FilterExpressionAndOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() TfBudget_FilterExpressionAndTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfBudget_FilterExpressionAndTagsProperty
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
	PutCostCategories(value *TfBudget_FilterExpressionAndCostCategoriesProperty)
	// Experimental.
	PutDimensions(value *TfBudget_FilterExpressionAndDimensionsProperty)
	// Experimental.
	PutNot(value *TfBudget_FilterExpressionAndNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *TfBudget_FilterExpressionAndTagsProperty)
	// Experimental.
	ResetAnd()
	// Experimental.
	ResetCostCategories()
	// Experimental.
	ResetDimensions()
	// Experimental.
	ResetNot()
	// Experimental.
	ResetOr()
	// Experimental.
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfBudget_FilterExpressionAndPropertyOutputReference
type jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) And() TfBudget_FilterExpressionAndAndPropertyList {
	var returns TfBudget_FilterExpressionAndAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) CostCategories() TfBudget_FilterExpressionAndCostCategoriesPropertyOutputReference {
	var returns TfBudget_FilterExpressionAndCostCategoriesPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) CostCategoriesInput() *TfBudget_FilterExpressionAndCostCategoriesProperty {
	var returns *TfBudget_FilterExpressionAndCostCategoriesProperty
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Dimensions() TfBudget_FilterExpressionAndDimensionsPropertyOutputReference {
	var returns TfBudget_FilterExpressionAndDimensionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) DimensionsInput() *TfBudget_FilterExpressionAndDimensionsProperty {
	var returns *TfBudget_FilterExpressionAndDimensionsProperty
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Not() TfBudget_FilterExpressionAndNotPropertyOutputReference {
	var returns TfBudget_FilterExpressionAndNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) NotInput() *TfBudget_FilterExpressionAndNotProperty {
	var returns *TfBudget_FilterExpressionAndNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Or() TfBudget_FilterExpressionAndOrPropertyList {
	var returns TfBudget_FilterExpressionAndOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Tags() TfBudget_FilterExpressionAndTagsPropertyOutputReference {
	var returns TfBudget_FilterExpressionAndTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) TagsInput() *TfBudget_FilterExpressionAndTagsProperty {
	var returns *TfBudget_FilterExpressionAndTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBudget_FilterExpressionAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfBudget_FilterExpressionAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBudget_FilterExpressionAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget.FilterExpressionAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBudget_FilterExpressionAndPropertyOutputReference_Override(t TfBudget_FilterExpressionAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget.FilterExpressionAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutAnd(value interface{}) {
	if err := t.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnd",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutCostCategories(value *TfBudget_FilterExpressionAndCostCategoriesProperty) {
	if err := t.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutDimensions(value *TfBudget_FilterExpressionAndDimensionsProperty) {
	if err := t.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimensions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutNot(value *TfBudget_FilterExpressionAndNotProperty) {
	if err := t.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutOr(value interface{}) {
	if err := t.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOr",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) PutTags(value *TfBudget_FilterExpressionAndTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		t,
		"resetAnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		t,
		"resetDimensions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		t,
		"resetNot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		t,
		"resetOr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBudget_FilterExpressionAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

