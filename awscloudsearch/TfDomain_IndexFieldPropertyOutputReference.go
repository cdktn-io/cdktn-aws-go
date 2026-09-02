package awscloudsearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudsearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudsearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomain_IndexFieldPropertyOutputReference interface {
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

// The jsii proxy struct for TfDomain_IndexFieldPropertyOutputReference
type jsiiProxy_TfDomain_IndexFieldPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) AnalysisScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) AnalysisSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"analysisSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) DefaultValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) DefaultValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Facet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) FacetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Highlight() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) HighlightInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"highlightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Return() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"return",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ReturnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"returnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Search() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"search",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) SearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Sort() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sort",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) SortInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) SourceFields() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) SourceFieldsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDomain_IndexFieldPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDomain_IndexFieldPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDomain_IndexFieldPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDomain_IndexFieldPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.TfDomain.IndexFieldPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDomain_IndexFieldPropertyOutputReference_Override(t TfDomain_IndexFieldPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudsearch.TfDomain.IndexFieldPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetAnalysisScheme(val *string) {
	if err := j.validateSetAnalysisSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"analysisScheme",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetDefaultValue(val *string) {
	if err := j.validateSetDefaultValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetFacet(val interface{}) {
	if err := j.validateSetFacetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facet",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetHighlight(val interface{}) {
	if err := j.validateSetHighlightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"highlight",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetReturn(val interface{}) {
	if err := j.validateSetReturnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"return",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetSearch(val interface{}) {
	if err := j.validateSetSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"search",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetSort(val interface{}) {
	if err := j.validateSetSortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sort",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetSourceFields(val *string) {
	if err := j.validateSetSourceFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceFields",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetAnalysisScheme() {
	_jsii_.InvokeVoid(
		t,
		"resetAnalysisScheme",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetDefaultValue() {
	_jsii_.InvokeVoid(
		t,
		"resetDefaultValue",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetFacet() {
	_jsii_.InvokeVoid(
		t,
		"resetFacet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetHighlight() {
	_jsii_.InvokeVoid(
		t,
		"resetHighlight",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetReturn() {
	_jsii_.InvokeVoid(
		t,
		"resetReturn",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetSearch() {
	_jsii_.InvokeVoid(
		t,
		"resetSearch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetSort() {
	_jsii_.InvokeVoid(
		t,
		"resetSort",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ResetSourceFields() {
	_jsii_.InvokeVoid(
		t,
		"resetSourceFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDomain_IndexFieldPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

