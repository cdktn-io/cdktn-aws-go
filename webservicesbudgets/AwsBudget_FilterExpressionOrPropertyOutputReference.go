package webservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudget_FilterExpressionOrPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() AwsBudget_FilterExpressionOrAndPropertyList
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
	CostCategories() AwsBudget_FilterExpressionOrCostCategoriesPropertyOutputReference
	// Experimental.
	CostCategoriesInput() *AwsBudget_FilterExpressionOrCostCategoriesProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimensions() AwsBudget_FilterExpressionOrDimensionsPropertyOutputReference
	// Experimental.
	DimensionsInput() *AwsBudget_FilterExpressionOrDimensionsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Not() AwsBudget_FilterExpressionOrNotPropertyOutputReference
	// Experimental.
	NotInput() *AwsBudget_FilterExpressionOrNotProperty
	// Experimental.
	Or() AwsBudget_FilterExpressionOrOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() AwsBudget_FilterExpressionOrTagsPropertyOutputReference
	// Experimental.
	TagsInput() *AwsBudget_FilterExpressionOrTagsProperty
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
	PutCostCategories(value *AwsBudget_FilterExpressionOrCostCategoriesProperty)
	// Experimental.
	PutDimensions(value *AwsBudget_FilterExpressionOrDimensionsProperty)
	// Experimental.
	PutNot(value *AwsBudget_FilterExpressionOrNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *AwsBudget_FilterExpressionOrTagsProperty)
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

// The jsii proxy struct for AwsBudget_FilterExpressionOrPropertyOutputReference
type jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) And() AwsBudget_FilterExpressionOrAndPropertyList {
	var returns AwsBudget_FilterExpressionOrAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) CostCategories() AwsBudget_FilterExpressionOrCostCategoriesPropertyOutputReference {
	var returns AwsBudget_FilterExpressionOrCostCategoriesPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) CostCategoriesInput() *AwsBudget_FilterExpressionOrCostCategoriesProperty {
	var returns *AwsBudget_FilterExpressionOrCostCategoriesProperty
	_jsii_.Get(
		j,
		"costCategoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Dimensions() AwsBudget_FilterExpressionOrDimensionsPropertyOutputReference {
	var returns AwsBudget_FilterExpressionOrDimensionsPropertyOutputReference
	_jsii_.Get(
		j,
		"dimensions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) DimensionsInput() *AwsBudget_FilterExpressionOrDimensionsProperty {
	var returns *AwsBudget_FilterExpressionOrDimensionsProperty
	_jsii_.Get(
		j,
		"dimensionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Not() AwsBudget_FilterExpressionOrNotPropertyOutputReference {
	var returns AwsBudget_FilterExpressionOrNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) NotInput() *AwsBudget_FilterExpressionOrNotProperty {
	var returns *AwsBudget_FilterExpressionOrNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Or() AwsBudget_FilterExpressionOrOrPropertyList {
	var returns AwsBudget_FilterExpressionOrOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Tags() AwsBudget_FilterExpressionOrTagsPropertyOutputReference {
	var returns AwsBudget_FilterExpressionOrTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) TagsInput() *AwsBudget_FilterExpressionOrTagsProperty {
	var returns *AwsBudget_FilterExpressionOrTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBudget_FilterExpressionOrPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsBudget_FilterExpressionOrPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBudget_FilterExpressionOrPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.FilterExpressionOrPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBudget_FilterExpressionOrPropertyOutputReference_Override(a AwsBudget_FilterExpressionOrPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.FilterExpressionOrPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutAnd(value interface{}) {
	if err := a.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnd",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutCostCategories(value *AwsBudget_FilterExpressionOrCostCategoriesProperty) {
	if err := a.validatePutCostCategoriesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostCategories",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutDimensions(value *AwsBudget_FilterExpressionOrDimensionsProperty) {
	if err := a.validatePutDimensionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDimensions",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutNot(value *AwsBudget_FilterExpressionOrNotProperty) {
	if err := a.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNot",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutOr(value interface{}) {
	if err := a.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOr",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) PutTags(value *AwsBudget_FilterExpressionOrTagsProperty) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		a,
		"resetAnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetCostCategories() {
	_jsii_.InvokeVoid(
		a,
		"resetCostCategories",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetDimensions() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		a,
		"resetNot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		a,
		"resetOr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBudget_FilterExpressionOrPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

