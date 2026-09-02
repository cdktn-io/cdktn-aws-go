package awswebservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awswebservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfBudget_FilterExpressionNotPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() TfBudget_FilterExpressionNotAndPropertyList
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
	CostCategories() TfBudget_FilterExpressionNotCostCategoriesPropertyOutputReference
	// Experimental.
	CostCategoriesInput() *TfBudget_FilterExpressionNotCostCategoriesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimensions() TfBudget_FilterExpressionNotDimensionsPropertyOutputReference
	// Experimental.
	DimensionsInput() *TfBudget_FilterExpressionNotDimensionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfBudget_FilterExpressionNotProperty
	// Experimental.
	SetInternalValue(val *TfBudget_FilterExpressionNotProperty)
	// Experimental.
	Not() TfBudget_FilterExpressionNotNotPropertyOutputReference
	// Experimental.
	NotInput() *TfBudget_FilterExpressionNotNotProperty
	// Experimental.
	Or() TfBudget_FilterExpressionNotOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() TfBudget_FilterExpressionNotTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfBudget_FilterExpressionNotTagsProperty
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
	PutCostCategories(value *TfBudget_FilterExpressionNotCostCategoriesProperty)
	// Experimental.
	PutDimensions(value *TfBudget_FilterExpressionNotDimensionsProperty)
	// Experimental.
	PutNot(value *TfBudget_FilterExpressionNotNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *TfBudget_FilterExpressionNotTagsProperty)
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

// The jsii proxy struct for TfBudget_FilterExpressionNotPropertyOutputReference
type jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) And() TfBudget_FilterExpressionNotAndPropertyList {
	var returns TfBudget_FilterExpressionNotAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) CostCategories() TfBudget_FilterExpressionNotCostCategoriesPropertyOutputReference {
	var returns TfBudget_FilterExpressionNotCostCategoriesPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) CostCategoriesInput() *TfBudget_FilterExpressionNotCostCategoriesProperty {
	var returns *TfBudget_FilterExpressionNotCostCategoriesProperty
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Dimensions() TfBudget_FilterExpressionNotDimensionsPropertyOutputReference {
	var returns TfBudget_FilterExpressionNotDimensionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) DimensionsInput() *TfBudget_FilterExpressionNotDimensionsProperty {
	var returns *TfBudget_FilterExpressionNotDimensionsProperty
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) InternalValue() *TfBudget_FilterExpressionNotProperty {
	var returns *TfBudget_FilterExpressionNotProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Not() TfBudget_FilterExpressionNotNotPropertyOutputReference {
	var returns TfBudget_FilterExpressionNotNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) NotInput() *TfBudget_FilterExpressionNotNotProperty {
	var returns *TfBudget_FilterExpressionNotNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Or() TfBudget_FilterExpressionNotOrPropertyList {
	var returns TfBudget_FilterExpressionNotOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Tags() TfBudget_FilterExpressionNotTagsPropertyOutputReference {
	var returns TfBudget_FilterExpressionNotTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) TagsInput() *TfBudget_FilterExpressionNotTagsProperty {
	var returns *TfBudget_FilterExpressionNotTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfBudget_FilterExpressionNotPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfBudget_FilterExpressionNotPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfBudget_FilterExpressionNotPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget.FilterExpressionNotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfBudget_FilterExpressionNotPropertyOutputReference_Override(t TfBudget_FilterExpressionNotPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.TfBudget.FilterExpressionNotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference)SetInternalValue(val *TfBudget_FilterExpressionNotProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutAnd(value interface{}) {
	if err := t.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnd",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutCostCategories(value *TfBudget_FilterExpressionNotCostCategoriesProperty) {
	if err := t.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutDimensions(value *TfBudget_FilterExpressionNotDimensionsProperty) {
	if err := t.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimensions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutNot(value *TfBudget_FilterExpressionNotNotProperty) {
	if err := t.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutOr(value interface{}) {
	if err := t.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOr",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) PutTags(value *TfBudget_FilterExpressionNotTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		t,
		"resetAnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		t,
		"resetDimensions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		t,
		"resetNot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		t,
		"resetOr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfBudget_FilterExpressionNotPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

