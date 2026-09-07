package webservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudget_FilterExpressionAndPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() AwsBudget_FilterExpressionAndAndPropertyList
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
	CostCategories() AwsBudget_FilterExpressionAndCostCategoriesPropertyOutputReference
	// Experimental.
	CostCategoriesInput() *AwsBudget_FilterExpressionAndCostCategoriesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimensions() AwsBudget_FilterExpressionAndDimensionsPropertyOutputReference
	// Experimental.
	DimensionsInput() *AwsBudget_FilterExpressionAndDimensionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Not() AwsBudget_FilterExpressionAndNotPropertyOutputReference
	// Experimental.
	NotInput() *AwsBudget_FilterExpressionAndNotProperty
	// Experimental.
	Or() AwsBudget_FilterExpressionAndOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() AwsBudget_FilterExpressionAndTagsPropertyOutputReference
	// Experimental.
	TagsInput() *AwsBudget_FilterExpressionAndTagsProperty
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
	PutCostCategories(value *AwsBudget_FilterExpressionAndCostCategoriesProperty)
	// Experimental.
	PutDimensions(value *AwsBudget_FilterExpressionAndDimensionsProperty)
	// Experimental.
	PutNot(value *AwsBudget_FilterExpressionAndNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *AwsBudget_FilterExpressionAndTagsProperty)
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

// The jsii proxy struct for AwsBudget_FilterExpressionAndPropertyOutputReference
type jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) And() AwsBudget_FilterExpressionAndAndPropertyList {
	var returns AwsBudget_FilterExpressionAndAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) CostCategories() AwsBudget_FilterExpressionAndCostCategoriesPropertyOutputReference {
	var returns AwsBudget_FilterExpressionAndCostCategoriesPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) CostCategoriesInput() *AwsBudget_FilterExpressionAndCostCategoriesProperty {
	var returns *AwsBudget_FilterExpressionAndCostCategoriesProperty
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Dimensions() AwsBudget_FilterExpressionAndDimensionsPropertyOutputReference {
	var returns AwsBudget_FilterExpressionAndDimensionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) DimensionsInput() *AwsBudget_FilterExpressionAndDimensionsProperty {
	var returns *AwsBudget_FilterExpressionAndDimensionsProperty
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Not() AwsBudget_FilterExpressionAndNotPropertyOutputReference {
	var returns AwsBudget_FilterExpressionAndNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) NotInput() *AwsBudget_FilterExpressionAndNotProperty {
	var returns *AwsBudget_FilterExpressionAndNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Or() AwsBudget_FilterExpressionAndOrPropertyList {
	var returns AwsBudget_FilterExpressionAndOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Tags() AwsBudget_FilterExpressionAndTagsPropertyOutputReference {
	var returns AwsBudget_FilterExpressionAndTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) TagsInput() *AwsBudget_FilterExpressionAndTagsProperty {
	var returns *AwsBudget_FilterExpressionAndTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBudget_FilterExpressionAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBudget_FilterExpressionAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBudget_FilterExpressionAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.FilterExpressionAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBudget_FilterExpressionAndPropertyOutputReference_Override(a AwsBudget_FilterExpressionAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.FilterExpressionAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutAnd(value interface{}) {
	if err := a.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnd",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutCostCategories(value *AwsBudget_FilterExpressionAndCostCategoriesProperty) {
	if err := a.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutDimensions(value *AwsBudget_FilterExpressionAndDimensionsProperty) {
	if err := a.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDimensions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutNot(value *AwsBudget_FilterExpressionAndNotProperty) {
	if err := a.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNot",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutOr(value interface{}) {
	if err := a.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOr",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) PutTags(value *AwsBudget_FilterExpressionAndTagsProperty) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		a,
		"resetAnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		a,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		a,
		"resetNot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		a,
		"resetOr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

