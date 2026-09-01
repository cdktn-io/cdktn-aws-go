package awscloudsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudsearchDomain_IndexFieldPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AnalysisScheme() *string
	// Experimental.
	SetAnalysisScheme(val *string)
	// Experimental.
	AnalysisSchemeInput() *string
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DefaultValue() *string
	// Experimental.
	SetDefaultValue(val *string)
	// Experimental.
	DefaultValueInput() *string
	// Experimental.
	Facet() interface{}
	// Experimental.
	SetFacet(val interface{})
	// Experimental.
	FacetInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	Highlight() interface{}
	// Experimental.
	SetHighlight(val interface{})
	// Experimental.
	HighlightInput() interface{}
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	Return() interface{}
	// Experimental.
	SetReturn(val interface{})
	// Experimental.
	ReturnInput() interface{}
	// Experimental.
	Search() interface{}
	// Experimental.
	SetSearch(val interface{})
	// Experimental.
	SearchInput() interface{}
	// Experimental.
	Sort() interface{}
	// Experimental.
	SetSort(val interface{})
	// Experimental.
	SortInput() interface{}
	// Experimental.
	SourceFields() *string
	// Experimental.
	SetSourceFields(val *string)
	// Experimental.
	SourceFieldsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Type() *string
	// Experimental.
	SetType(val *string)
	// Experimental.
	TypeInput() *string
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
	ResetAnalysisScheme()
	// Experimental.
	ResetDefaultValue()
	// Experimental.
	ResetFacet()
	// Experimental.
	ResetHighlight()
	// Experimental.
	ResetReturn()
	// Experimental.
	ResetSearch()
	// Experimental.
	ResetSort()
	// Experimental.
	ResetSourceFields()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCloudsearchDomain_IndexFieldPropertyOutputReference
type jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) AnalysisScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) AnalysisSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Facet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) FacetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Highlight() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) HighlightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Return() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"return",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ReturnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Search() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) SearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Sort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) SortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) SourceFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) SourceFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCloudsearchDomain_IndexFieldPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsCloudsearchDomain_IndexFieldPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCloudsearchDomain_IndexFieldPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.AwsCloudsearchDomain.IndexFieldPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCloudsearchDomain_IndexFieldPropertyOutputReference_Override(a AwsCloudsearchDomain_IndexFieldPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.AwsCloudsearchDomain.IndexFieldPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetAnalysisScheme(val *string) {
	if err := j.validateSetAnalysisSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisScheme",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetFacet(val interface{}) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetHighlight(val interface{}) {
	if err := j.validateSetHighlightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highlight",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetReturn(val interface{}) {
	if err := j.validateSetReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"return",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetSearch(val interface{}) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetSort(val interface{}) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetSourceFields(val *string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetAnalysisScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetAnalysisScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		a,
		"resetFacet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetHighlight() {
	_jsii_.InvokeVoid(
		a,
		"resetHighlight",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetReturn() {
	_jsii_.InvokeVoid(
		a,
		"resetReturn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		a,
		"resetSearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		a,
		"resetSort",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ResetSourceFields() {
	_jsii_.InvokeVoid(
		a,
		"resetSourceFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCloudsearchDomain_IndexFieldPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

