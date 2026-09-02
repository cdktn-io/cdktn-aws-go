package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCostCategory_RuleRulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() TfCostCategory_RuleRuleAndPropertyList
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
	CostCategory() TfCostCategory_RuleRuleCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *TfCostCategory_RuleRuleCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() TfCostCategory_RuleRuleDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *TfCostCategory_RuleRuleDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCostCategory_RuleRuleProperty
	// Experimental.
	SetInternalValue(val *TfCostCategory_RuleRuleProperty)
	// Experimental.
	Not() TfCostCategory_RuleRuleNotPropertyOutputReference
	// Experimental.
	NotInput() *TfCostCategory_RuleRuleNotProperty
	// Experimental.
	Or() TfCostCategory_RuleRuleOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() TfCostCategory_RuleRuleTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfCostCategory_RuleRuleTagsProperty
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
	PutCostCategory(value *TfCostCategory_RuleRuleCostCategoryProperty)
	// Experimental.
	PutDimension(value *TfCostCategory_RuleRuleDimensionProperty)
	// Experimental.
	PutNot(value *TfCostCategory_RuleRuleNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *TfCostCategory_RuleRuleTagsProperty)
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

// The jsii proxy struct for TfCostCategory_RuleRulePropertyOutputReference
type jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) And() TfCostCategory_RuleRuleAndPropertyList {
	var returns TfCostCategory_RuleRuleAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) CostCategory() TfCostCategory_RuleRuleCostCategoryPropertyOutputReference {
	var returns TfCostCategory_RuleRuleCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) CostCategoryInput() *TfCostCategory_RuleRuleCostCategoryProperty {
	var returns *TfCostCategory_RuleRuleCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Dimension() TfCostCategory_RuleRuleDimensionPropertyOutputReference {
	var returns TfCostCategory_RuleRuleDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) DimensionInput() *TfCostCategory_RuleRuleDimensionProperty {
	var returns *TfCostCategory_RuleRuleDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) InternalValue() *TfCostCategory_RuleRuleProperty {
	var returns *TfCostCategory_RuleRuleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Not() TfCostCategory_RuleRuleNotPropertyOutputReference {
	var returns TfCostCategory_RuleRuleNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) NotInput() *TfCostCategory_RuleRuleNotProperty {
	var returns *TfCostCategory_RuleRuleNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Or() TfCostCategory_RuleRuleOrPropertyList {
	var returns TfCostCategory_RuleRuleOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Tags() TfCostCategory_RuleRuleTagsPropertyOutputReference {
	var returns TfCostCategory_RuleRuleTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) TagsInput() *TfCostCategory_RuleRuleTagsProperty {
	var returns *TfCostCategory_RuleRuleTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCostCategory_RuleRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCostCategory_RuleRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCostCategory_RuleRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCostCategory_RuleRulePropertyOutputReference_Override(t TfCostCategory_RuleRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference)SetInternalValue(val *TfCostCategory_RuleRuleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutAnd(value interface{}) {
	if err := t.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnd",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutCostCategory(value *TfCostCategory_RuleRuleCostCategoryProperty) {
	if err := t.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutDimension(value *TfCostCategory_RuleRuleDimensionProperty) {
	if err := t.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimension",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutNot(value *TfCostCategory_RuleRuleNotProperty) {
	if err := t.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutOr(value interface{}) {
	if err := t.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOr",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) PutTags(value *TfCostCategory_RuleRuleTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		t,
		"resetAnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		t,
		"resetDimension",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		t,
		"resetNot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		t,
		"resetOr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

