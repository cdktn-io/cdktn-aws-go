package awskendra

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskendra/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskendra/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKendraIndex_SearchPropertyOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Displayable() interface{}
	// Experimental.
	SetDisplayable(val interface{})
	// Experimental.
	DisplayableInput() interface{}
	// Experimental.
	Facetable() interface{}
	// Experimental.
	SetFacetable(val interface{})
	// Experimental.
	FacetableInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsKendraIndex_SearchProperty
	// Experimental.
	SetInternalValue(val *AwsKendraIndex_SearchProperty)
	// Experimental.
	Searchable() interface{}
	// Experimental.
	SetSearchable(val interface{})
	// Experimental.
	SearchableInput() interface{}
	// Experimental.
	Sortable() interface{}
	// Experimental.
	SetSortable(val interface{})
	// Experimental.
	SortableInput() interface{}
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
	ResetDisplayable()
	// Experimental.
	ResetFacetable()
	// Experimental.
	ResetSearchable()
	// Experimental.
	ResetSortable()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKendraIndex_SearchPropertyOutputReference
type jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Displayable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"displayable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) DisplayableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"displayableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Facetable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) FacetableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"facetableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) InternalValue() *AwsKendraIndex_SearchProperty {
	var returns *AwsKendraIndex_SearchProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Searchable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) SearchableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"searchableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Sortable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) SortableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKendraIndex_SearchPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKendraIndex_SearchPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKendraIndex_SearchPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraIndex.SearchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKendraIndex_SearchPropertyOutputReference_Override(a AwsKendraIndex_SearchPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kendra.AwsKendraIndex.SearchPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetDisplayable(val interface{}) {
	if err := j.validateSetDisplayableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayable",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetFacetable(val interface{}) {
	if err := j.validateSetFacetableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"facetable",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetInternalValue(val *AwsKendraIndex_SearchProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetSearchable(val interface{}) {
	if err := j.validateSetSearchableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"searchable",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetSortable(val interface{}) {
	if err := j.validateSetSortableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortable",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ResetDisplayable() {
	_jsii_.InvokeVoid(
		a,
		"resetDisplayable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ResetFacetable() {
	_jsii_.InvokeVoid(
		a,
		"resetFacetable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ResetSearchable() {
	_jsii_.InvokeVoid(
		a,
		"resetSearchable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ResetSortable() {
	_jsii_.InvokeVoid(
		a,
		"resetSortable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKendraIndex_SearchPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

