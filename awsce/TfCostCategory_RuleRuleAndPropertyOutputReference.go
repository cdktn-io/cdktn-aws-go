package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCostCategory_RuleRuleAndPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() TfCostCategory_RuleRuleAndAndPropertyList
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
	CostCategory() TfCostCategory_RuleRuleAndCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *TfCostCategory_RuleRuleAndCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() TfCostCategory_RuleRuleAndDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *TfCostCategory_RuleRuleAndDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Not() TfCostCategory_RuleRuleAndNotPropertyOutputReference
	// Experimental.
	NotInput() *TfCostCategory_RuleRuleAndNotProperty
	// Experimental.
	Or() TfCostCategory_RuleRuleAndOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() TfCostCategory_RuleRuleAndTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfCostCategory_RuleRuleAndTagsProperty
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
	PutCostCategory(value *TfCostCategory_RuleRuleAndCostCategoryProperty)
	// Experimental.
	PutDimension(value *TfCostCategory_RuleRuleAndDimensionProperty)
	// Experimental.
	PutNot(value *TfCostCategory_RuleRuleAndNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *TfCostCategory_RuleRuleAndTagsProperty)
	// Experimental.
	ResetAnd()
	// Experimental.
	ResetCostCategory()
	// Experimental.
	ResetDimension()
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

// The jsii proxy struct for TfCostCategory_RuleRuleAndPropertyOutputReference
type jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) And() TfCostCategory_RuleRuleAndAndPropertyList {
	var returns TfCostCategory_RuleRuleAndAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) CostCategory() TfCostCategory_RuleRuleAndCostCategoryPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) CostCategoryInput() *TfCostCategory_RuleRuleAndCostCategoryProperty {
	var returns *TfCostCategory_RuleRuleAndCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Dimension() TfCostCategory_RuleRuleAndDimensionPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) DimensionInput() *TfCostCategory_RuleRuleAndDimensionProperty {
	var returns *TfCostCategory_RuleRuleAndDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Not() TfCostCategory_RuleRuleAndNotPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) NotInput() *TfCostCategory_RuleRuleAndNotProperty {
	var returns *TfCostCategory_RuleRuleAndNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Or() TfCostCategory_RuleRuleAndOrPropertyList {
	var returns TfCostCategory_RuleRuleAndOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Tags() TfCostCategory_RuleRuleAndTagsPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) TagsInput() *TfCostCategory_RuleRuleAndTagsProperty {
	var returns *TfCostCategory_RuleRuleAndTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCostCategory_RuleRuleAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCostCategory_RuleRuleAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCostCategory_RuleRuleAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCostCategory_RuleRuleAndPropertyOutputReference_Override(t TfCostCategory_RuleRuleAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutAnd(value interface{}) {
	if err := t.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnd",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutCostCategory(value *TfCostCategory_RuleRuleAndCostCategoryProperty) {
	if err := t.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutDimension(value *TfCostCategory_RuleRuleAndDimensionProperty) {
	if err := t.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimension",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutNot(value *TfCostCategory_RuleRuleAndNotProperty) {
	if err := t.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutOr(value interface{}) {
	if err := t.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOr",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) PutTags(value *TfCostCategory_RuleRuleAndTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		t,
		"resetAnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		t,
		"resetDimension",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		t,
		"resetNot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		t,
		"resetOr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

