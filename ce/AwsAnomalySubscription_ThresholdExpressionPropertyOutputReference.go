package ce

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/ce/jsii"

	"github.com/cdktn-io/cdktn-aws-go/ce/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	And() AwsAnomalySubscription_AndPropertyList
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
	CostCategory() AwsAnomalySubscription_ThresholdExpressionCostCategoryPropertyOutputReference
	// Experimental.
	CostCategoryInput() *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Dimension() AwsAnomalySubscription_ThresholdExpressionDimensionPropertyOutputReference
	// Experimental.
	DimensionInput() *AwsAnomalySubscription_ThresholdExpressionDimensionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsAnomalySubscription_ThresholdExpressionProperty
	// Experimental.
	SetInternalValue(val *AwsAnomalySubscription_ThresholdExpressionProperty)
	// Experimental.
	Not() AwsAnomalySubscription_NotPropertyOutputReference
	// Experimental.
	NotInput() *AwsAnomalySubscription_NotProperty
	// Experimental.
	Or() AwsAnomalySubscription_OrPropertyList
	// Experimental.
	OrInput() interface{}
	// Experimental.
	Tags() AwsAnomalySubscription_ThresholdExpressionTagsPropertyOutputReference
	// Experimental.
	TagsInput() *AwsAnomalySubscription_ThresholdExpressionTagsProperty
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
	PutCostCategory(value *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty)
	// Experimental.
	PutDimension(value *AwsAnomalySubscription_ThresholdExpressionDimensionProperty)
	// Experimental.
	PutNot(value *AwsAnomalySubscription_NotProperty)
	// Experimental.
	PutOr(value interface{})
	// Experimental.
	PutTags(value *AwsAnomalySubscription_ThresholdExpressionTagsProperty)
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

// The jsii proxy struct for AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference
type jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) And() AwsAnomalySubscription_AndPropertyList {
	var returns AwsAnomalySubscription_AndPropertyList
	_jsii_.Get(
		j,
		"and",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) AndInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"andInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) CostCategory() AwsAnomalySubscription_ThresholdExpressionCostCategoryPropertyOutputReference {
	var returns AwsAnomalySubscription_ThresholdExpressionCostCategoryPropertyOutputReference
	_jsii_.Get(
		j,
		"costCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) CostCategoryInput() *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty {
	var returns *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty
	_jsii_.Get(
		j,
		"costCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Dimension() AwsAnomalySubscription_ThresholdExpressionDimensionPropertyOutputReference {
	var returns AwsAnomalySubscription_ThresholdExpressionDimensionPropertyOutputReference
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) DimensionInput() *AwsAnomalySubscription_ThresholdExpressionDimensionProperty {
	var returns *AwsAnomalySubscription_ThresholdExpressionDimensionProperty
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) InternalValue() *AwsAnomalySubscription_ThresholdExpressionProperty {
	var returns *AwsAnomalySubscription_ThresholdExpressionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Not() AwsAnomalySubscription_NotPropertyOutputReference {
	var returns AwsAnomalySubscription_NotPropertyOutputReference
	_jsii_.Get(
		j,
		"not",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) NotInput() *AwsAnomalySubscription_NotProperty {
	var returns *AwsAnomalySubscription_NotProperty
	_jsii_.Get(
		j,
		"notInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Or() AwsAnomalySubscription_OrPropertyList {
	var returns AwsAnomalySubscription_OrPropertyList
	_jsii_.Get(
		j,
		"or",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) OrInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Tags() AwsAnomalySubscription_ThresholdExpressionTagsPropertyOutputReference {
	var returns AwsAnomalySubscription_ThresholdExpressionTagsPropertyOutputReference
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) TagsInput() *AwsAnomalySubscription_ThresholdExpressionTagsProperty {
	var returns *AwsAnomalySubscription_ThresholdExpressionTagsProperty
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAnomalySubscription_ThresholdExpressionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAnomalySubscription_ThresholdExpressionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ce.AwsAnomalySubscription.ThresholdExpressionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAnomalySubscription_ThresholdExpressionPropertyOutputReference_Override(a AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ce.AwsAnomalySubscription.ThresholdExpressionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference)SetInternalValue(val *AwsAnomalySubscription_ThresholdExpressionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutAnd(value interface{}) {
	if err := a.validatePutAndParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAnd",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutCostCategory(value *AwsAnomalySubscription_ThresholdExpressionCostCategoryProperty) {
	if err := a.validatePutCostCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCostCategory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutDimension(value *AwsAnomalySubscription_ThresholdExpressionDimensionProperty) {
	if err := a.validatePutDimensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDimension",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutNot(value *AwsAnomalySubscription_NotProperty) {
	if err := a.validatePutNotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNot",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutOr(value interface{}) {
	if err := a.validatePutOrParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putOr",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) PutTags(value *AwsAnomalySubscription_ThresholdExpressionTagsProperty) {
	if err := a.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTags",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetAnd() {
	_jsii_.InvokeVoid(
		a,
		"resetAnd",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetCostCategory() {
	_jsii_.InvokeVoid(
		a,
		"resetCostCategory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		a,
		"resetDimension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetNot() {
	_jsii_.InvokeVoid(
		a,
		"resetNot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetOr() {
	_jsii_.InvokeVoid(
		a,
		"resetOr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		a,
		"resetTags",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAnomalySubscription_ThresholdExpressionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

