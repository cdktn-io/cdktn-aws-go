package ce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCostCategory_RuleRulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() AwsCostCategory_RuleRuleAndPropertyList
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
	CostCategory() AwsCostCategory_RuleRuleCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *AwsCostCategory_RuleRuleCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() AwsCostCategory_RuleRuleDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *AwsCostCategory_RuleRuleDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCostCategory_RuleRuleProperty
	// Experimental.
	SetInternalValue(val *AwsCostCategory_RuleRuleProperty)
	// Experimental.
	Not() AwsCostCategory_RuleRuleNotPropertyOutputReference
	// Experimental.
	NotInput() *AwsCostCategory_RuleRuleNotProperty
	// Experimental.
	Or() AwsCostCategory_RuleRuleOrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() AwsCostCategory_RuleRuleTagsPropertyOutputReference
	// Experimental.
	TagsInput() *AwsCostCategory_RuleRuleTagsProperty
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
	PutCostCategory(value *AwsCostCategory_RuleRuleCostCategoryProperty)
	// Experimental.
	PutDimension(value *AwsCostCategory_RuleRuleDimensionProperty)
	// Experimental.
	PutNot(value *AwsCostCategory_RuleRuleNotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *AwsCostCategory_RuleRuleTagsProperty)
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

// The jsii proxy struct for AwsCostCategory_RuleRulePropertyOutputReference
type jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) And() AwsCostCategory_RuleRuleAndPropertyList {
	var returns AwsCostCategory_RuleRuleAndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) CostCategory() AwsCostCategory_RuleRuleCostCategoryPropertyOutputReference {
	var returns AwsCostCategory_RuleRuleCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) CostCategoryInput() *AwsCostCategory_RuleRuleCostCategoryProperty {
	var returns *AwsCostCategory_RuleRuleCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Dimension() AwsCostCategory_RuleRuleDimensionPropertyOutputReference {
	var returns AwsCostCategory_RuleRuleDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) DimensionInput() *AwsCostCategory_RuleRuleDimensionProperty {
	var returns *AwsCostCategory_RuleRuleDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) InternalValue() *AwsCostCategory_RuleRuleProperty {
	var returns *AwsCostCategory_RuleRuleProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Not() AwsCostCategory_RuleRuleNotPropertyOutputReference {
	var returns AwsCostCategory_RuleRuleNotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) NotInput() *AwsCostCategory_RuleRuleNotProperty {
	var returns *AwsCostCategory_RuleRuleNotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Or() AwsCostCategory_RuleRuleOrPropertyList {
	var returns AwsCostCategory_RuleRuleOrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Tags() AwsCostCategory_RuleRuleTagsPropertyOutputReference {
	var returns AwsCostCategory_RuleRuleTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) TagsInput() *AwsCostCategory_RuleRuleTagsProperty {
	var returns *AwsCostCategory_RuleRuleTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCostCategory_RuleRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCostCategory_RuleRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCostCategory_RuleRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.AwsCostCategory.RuleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCostCategory_RuleRulePropertyOutputReference_Override(a AwsCostCategory_RuleRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.AwsCostCategory.RuleRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference)SetInternalValue(val *AwsCostCategory_RuleRuleProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutAnd(value interface{}) {
	if err := a.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnd",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutCostCategory(value *AwsCostCategory_RuleRuleCostCategoryProperty) {
	if err := a.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutDimension(value *AwsCostCategory_RuleRuleDimensionProperty) {
	if err := a.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDimension",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutNot(value *AwsCostCategory_RuleRuleNotProperty) {
	if err := a.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNot",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutOr(value interface{}) {
	if err := a.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOr",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) PutTags(value *AwsCostCategory_RuleRuleTagsProperty) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		a,
		"resetAnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		a,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		a,
		"resetDimension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		a,
		"resetNot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		a,
		"resetOr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCostCategory_RuleRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

