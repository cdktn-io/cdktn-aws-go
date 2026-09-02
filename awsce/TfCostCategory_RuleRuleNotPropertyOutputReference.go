package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCostCategory_RuleRuleNotPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() TfCostCategory_RuleRuleNotAndPropertyList
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
	CostCategory() TfCostCategory_RuleRuleNotCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *TfCostCategory_RuleRuleNotCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() TfCostCategory_RuleRuleNotDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *TfCostCategory_RuleRuleNotDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCostCategory_RuleRuleNotProperty
	// Experimental.
	SetInternalValue(val *TfCostCategory_RuleRuleNotProperty)
	// Experimental.
	Not() TfCostCategory_RuleRuleNotNotPropertyOutputReference
	// Experimental.
	NotInput() *TfCostCategory_RuleRuleNotNotProperty
	// Experimental.
	Or() TfCostCategory_RuleRuleNotOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() TfCostCategory_RuleRuleNotTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfCostCategory_RuleRuleNotTagsProperty
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
	PutCostCategory(value *TfCostCategory_RuleRuleNotCostCategoryProperty)
	// Experimental.
	PutDimension(value *TfCostCategory_RuleRuleNotDimensionProperty)
	// Experimental.
	PutNot(value *TfCostCategory_RuleRuleNotNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *TfCostCategory_RuleRuleNotTagsProperty)
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

// The jsii proxy struct for TfCostCategory_RuleRuleNotPropertyOutputReference
type jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) And() TfCostCategory_RuleRuleNotAndPropertyList {
	var returns TfCostCategory_RuleRuleNotAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) CostCategory() TfCostCategory_RuleRuleNotCostCategoryPropertyOutputReference {
	var returns TfCostCategory_RuleRuleNotCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) CostCategoryInput() *TfCostCategory_RuleRuleNotCostCategoryProperty {
	var returns *TfCostCategory_RuleRuleNotCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Dimension() TfCostCategory_RuleRuleNotDimensionPropertyOutputReference {
	var returns TfCostCategory_RuleRuleNotDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) DimensionInput() *TfCostCategory_RuleRuleNotDimensionProperty {
	var returns *TfCostCategory_RuleRuleNotDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) InternalValue() *TfCostCategory_RuleRuleNotProperty {
	var returns *TfCostCategory_RuleRuleNotProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Not() TfCostCategory_RuleRuleNotNotPropertyOutputReference {
	var returns TfCostCategory_RuleRuleNotNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) NotInput() *TfCostCategory_RuleRuleNotNotProperty {
	var returns *TfCostCategory_RuleRuleNotNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Or() TfCostCategory_RuleRuleNotOrPropertyList {
	var returns TfCostCategory_RuleRuleNotOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Tags() TfCostCategory_RuleRuleNotTagsPropertyOutputReference {
	var returns TfCostCategory_RuleRuleNotTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) TagsInput() *TfCostCategory_RuleRuleNotTagsProperty {
	var returns *TfCostCategory_RuleRuleNotTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCostCategory_RuleRuleNotPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCostCategory_RuleRuleNotPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCostCategory_RuleRuleNotPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleNotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCostCategory_RuleRuleNotPropertyOutputReference_Override(t TfCostCategory_RuleRuleNotPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleNotPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference)SetInternalValue(val *TfCostCategory_RuleRuleNotProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutAnd(value interface{}) {
	if err := t.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAnd",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutCostCategory(value *TfCostCategory_RuleRuleNotCostCategoryProperty) {
	if err := t.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutDimension(value *TfCostCategory_RuleRuleNotDimensionProperty) {
	if err := t.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimension",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutNot(value *TfCostCategory_RuleRuleNotNotProperty) {
	if err := t.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNot",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutOr(value interface{}) {
	if err := t.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putOr",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) PutTags(value *TfCostCategory_RuleRuleNotTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		t,
		"resetAnd",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		t,
		"resetDimension",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		t,
		"resetNot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		t,
		"resetOr",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleNotPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

