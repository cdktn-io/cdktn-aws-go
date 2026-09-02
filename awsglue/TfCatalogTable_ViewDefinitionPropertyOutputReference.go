package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCatalogTable_ViewDefinitionPropertyOutputReference interface {
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
	Definer() *string
	// Experimental.
	SetDefiner(val *string)
	// Experimental.
	DefinerInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfCatalogTable_ViewDefinitionProperty
	// Experimental.
	SetInternalValue(val *TfCatalogTable_ViewDefinitionProperty)
	// Experimental.
	IsProtected() interface{}
	// Experimental.
	SetIsProtected(val interface{})
	// Experimental.
	IsProtectedInput() interface{}
	// Experimental.
	LastRefreshType() *string
	// Experimental.
	SetLastRefreshType(val *string)
	// Experimental.
	LastRefreshTypeInput() *string
	// Experimental.
	RefreshSeconds() *float64
	// Experimental.
	SetRefreshSeconds(val *float64)
	// Experimental.
	RefreshSecondsInput() *float64
	// Experimental.
	Representations() TfCatalogTable_RepresentationsPropertyList
	// Experimental.
	RepresentationsInput() interface{}
	// Experimental.
	SubObjects() *[]*string
	// Experimental.
	SetSubObjects(val *[]*string)
	// Experimental.
	SubObjectsInput() *[]*string
	// Experimental.
	SubObjectVersionIds() *[]*float64
	// Experimental.
	SetSubObjectVersionIds(val *[]*float64)
	// Experimental.
	SubObjectVersionIdsInput() *[]*float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ViewVersionId() *float64
	// Experimental.
	SetViewVersionId(val *float64)
	// Experimental.
	ViewVersionIdInput() *float64
	// Experimental.
	ViewVersionToken() *string
	// Experimental.
	SetViewVersionToken(val *string)
	// Experimental.
	ViewVersionTokenInput() *string
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
	PutRepresentations(value interface{})
	// Experimental.
	ResetDefiner()
	// Experimental.
	ResetIsProtected()
	// Experimental.
	ResetLastRefreshType()
	// Experimental.
	ResetRefreshSeconds()
	// Experimental.
	ResetRepresentations()
	// Experimental.
	ResetSubObjects()
	// Experimental.
	ResetSubObjectVersionIds()
	// Experimental.
	ResetViewVersionId()
	// Experimental.
	ResetViewVersionToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCatalogTable_ViewDefinitionPropertyOutputReference
type jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) Definer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) DefinerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"definerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) InternalValue() *TfCatalogTable_ViewDefinitionProperty {
	var returns *TfCatalogTable_ViewDefinitionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) IsProtected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) IsProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isProtectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) LastRefreshType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRefreshType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) LastRefreshTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastRefreshTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) RefreshSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) RefreshSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) Representations() TfCatalogTable_RepresentationsPropertyList {
	var returns TfCatalogTable_RepresentationsPropertyList
	_jsii_.Get(
		j,
		"representations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) RepresentationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"representationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) SubObjects() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) SubObjectsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subObjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) SubObjectVersionIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"subObjectVersionIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) SubObjectVersionIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"subObjectVersionIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ViewVersionId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"viewVersionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ViewVersionIdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"viewVersionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ViewVersionToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewVersionToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ViewVersionTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewVersionTokenInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCatalogTable_ViewDefinitionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfCatalogTable_ViewDefinitionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCatalogTable_ViewDefinitionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.ViewDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCatalogTable_ViewDefinitionPropertyOutputReference_Override(t TfCatalogTable_ViewDefinitionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfCatalogTable.ViewDefinitionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetDefiner(val *string) {
	if err := j.validateSetDefinerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"definer",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetInternalValue(val *TfCatalogTable_ViewDefinitionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetIsProtected(val interface{}) {
	if err := j.validateSetIsProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isProtected",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetLastRefreshType(val *string) {
	if err := j.validateSetLastRefreshTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastRefreshType",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetRefreshSeconds(val *float64) {
	if err := j.validateSetRefreshSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshSeconds",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetSubObjects(val *[]*string) {
	if err := j.validateSetSubObjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subObjects",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetSubObjectVersionIds(val *[]*float64) {
	if err := j.validateSetSubObjectVersionIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subObjectVersionIds",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetViewVersionId(val *float64) {
	if err := j.validateSetViewVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewVersionId",
		val,
	)
}

func (j *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference)SetViewVersionToken(val *string) {
	if err := j.validateSetViewVersionTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewVersionToken",
		val,
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) PutRepresentations(value interface{}) {
	if err := t.validatePutRepresentationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRepresentations",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetDefiner() {
	_jsii_.InvokeVoid(
		t,
		"resetDefiner",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetIsProtected() {
	_jsii_.InvokeVoid(
		t,
		"resetIsProtected",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetLastRefreshType() {
	_jsii_.InvokeVoid(
		t,
		"resetLastRefreshType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetRefreshSeconds() {
	_jsii_.InvokeVoid(
		t,
		"resetRefreshSeconds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetRepresentations() {
	_jsii_.InvokeVoid(
		t,
		"resetRepresentations",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetSubObjects() {
	_jsii_.InvokeVoid(
		t,
		"resetSubObjects",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetSubObjectVersionIds() {
	_jsii_.InvokeVoid(
		t,
		"resetSubObjectVersionIds",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetViewVersionId() {
	_jsii_.InvokeVoid(
		t,
		"resetViewVersionId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ResetViewVersionToken() {
	_jsii_.InvokeVoid(
		t,
		"resetViewVersionToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfCatalogTable_ViewDefinitionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

