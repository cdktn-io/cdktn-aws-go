package awsce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCostCategory_RuleRuleAndAndPropertyOutputReference interface {
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
	// Experimental.
	CostCategory() TfCostCategory_RuleRuleAndAndCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *TfCostCategory_RuleRuleAndAndCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() TfCostCategory_RuleRuleAndAndDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *TfCostCategory_RuleRuleAndAndDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Tags() TfCostCategory_RuleRuleAndAndTagsPropertyOutputReference
	// Experimental.
	TagsInput() *TfCostCategory_RuleRuleAndAndTagsProperty
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
	PutCostCategory(value *TfCostCategory_RuleRuleAndAndCostCategoryProperty)
	// Experimental.
	PutDimension(value *TfCostCategory_RuleRuleAndAndDimensionProperty)
	// Experimental.
	PutTags(value *TfCostCategory_RuleRuleAndAndTagsProperty)
	// Experimental.
	ResetCostCategory()
	// Experimental.
	ResetDimension()
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

// The jsii proxy struct for TfCostCategory_RuleRuleAndAndPropertyOutputReference
type jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) CostCategory() TfCostCategory_RuleRuleAndAndCostCategoryPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndAndCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) CostCategoryInput() *TfCostCategory_RuleRuleAndAndCostCategoryProperty {
	var returns *TfCostCategory_RuleRuleAndAndCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) Dimension() TfCostCategory_RuleRuleAndAndDimensionPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndAndDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) DimensionInput() *TfCostCategory_RuleRuleAndAndDimensionProperty {
	var returns *TfCostCategory_RuleRuleAndAndDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) Tags() TfCostCategory_RuleRuleAndAndTagsPropertyOutputReference {
	var returns TfCostCategory_RuleRuleAndAndTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) TagsInput() *TfCostCategory_RuleRuleAndAndTagsProperty {
	var returns *TfCostCategory_RuleRuleAndAndTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCostCategory_RuleRuleAndAndPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCostCategory_RuleRuleAndAndPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCostCategory_RuleRuleAndAndPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleAndAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCostCategory_RuleRuleAndAndPropertyOutputReference_Override(t TfCostCategory_RuleRuleAndAndPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.TfCostCategory.RuleRuleAndAndPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) PutCostCategory(value *TfCostCategory_RuleRuleAndAndCostCategoryProperty) {
	if err := t.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) PutDimension(value *TfCostCategory_RuleRuleAndAndDimensionProperty) {
	if err := t.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDimension",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) PutTags(value *TfCostCategory_RuleRuleAndAndTagsProperty) {
	if err := t.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTags",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		t,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		t,
		"resetDimension",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		t,
		"resetTags",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCostCategory_RuleRuleAndAndPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

