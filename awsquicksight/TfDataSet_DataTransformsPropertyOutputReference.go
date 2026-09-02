package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSet_DataTransformsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CastColumnTypeOperation() TfDataSet_CastColumnTypeOperationPropertyOutputReference
	// Experimental.
	CastColumnTypeOperationInput() *TfDataSet_CastColumnTypeOperationProperty
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
	CreateColumnsOperation() TfDataSet_CreateColumnsOperationPropertyOutputReference
	// Experimental.
	CreateColumnsOperationInput() *TfDataSet_CreateColumnsOperationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FilterOperation() TfDataSet_FilterOperationPropertyOutputReference
	// Experimental.
	FilterOperationInput() *TfDataSet_FilterOperationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ProjectOperation() TfDataSet_ProjectOperationPropertyOutputReference
	// Experimental.
	ProjectOperationInput() *TfDataSet_ProjectOperationProperty
	// Experimental.
	RenameColumnOperation() TfDataSet_RenameColumnOperationPropertyOutputReference
	// Experimental.
	RenameColumnOperationInput() *TfDataSet_RenameColumnOperationProperty
	// Experimental.
	TagColumnOperation() TfDataSet_TagColumnOperationPropertyOutputReference
	// Experimental.
	TagColumnOperationInput() *TfDataSet_TagColumnOperationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UntagColumnOperation() TfDataSet_UntagColumnOperationPropertyOutputReference
	// Experimental.
	UntagColumnOperationInput() *TfDataSet_UntagColumnOperationProperty
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
	PutCastColumnTypeOperation(value *TfDataSet_CastColumnTypeOperationProperty)
	// Experimental.
	PutCreateColumnsOperation(value *TfDataSet_CreateColumnsOperationProperty)
	// Experimental.
	PutFilterOperation(value *TfDataSet_FilterOperationProperty)
	// Experimental.
	PutProjectOperation(value *TfDataSet_ProjectOperationProperty)
	// Experimental.
	PutRenameColumnOperation(value *TfDataSet_RenameColumnOperationProperty)
	// Experimental.
	PutTagColumnOperation(value *TfDataSet_TagColumnOperationProperty)
	// Experimental.
	PutUntagColumnOperation(value *TfDataSet_UntagColumnOperationProperty)
	// Experimental.
	ResetCastColumnTypeOperation()
	// Experimental.
	ResetCreateColumnsOperation()
	// Experimental.
	ResetFilterOperation()
	// Experimental.
	ResetProjectOperation()
	// Experimental.
	ResetRenameColumnOperation()
	// Experimental.
	ResetTagColumnOperation()
	// Experimental.
	ResetUntagColumnOperation()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSet_DataTransformsPropertyOutputReference
type jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperation() TfDataSet_CastColumnTypeOperationPropertyOutputReference {
	var returns TfDataSet_CastColumnTypeOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"castColumnTypeOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperationInput() *TfDataSet_CastColumnTypeOperationProperty {
	var returns *TfDataSet_CastColumnTypeOperationProperty
	_jsii_.Get(
		j,
		"castColumnTypeOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperation() TfDataSet_CreateColumnsOperationPropertyOutputReference {
	var returns TfDataSet_CreateColumnsOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"createColumnsOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperationInput() *TfDataSet_CreateColumnsOperationProperty {
	var returns *TfDataSet_CreateColumnsOperationProperty
	_jsii_.Get(
		j,
		"createColumnsOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) FilterOperation() TfDataSet_FilterOperationPropertyOutputReference {
	var returns TfDataSet_FilterOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"filterOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) FilterOperationInput() *TfDataSet_FilterOperationProperty {
	var returns *TfDataSet_FilterOperationProperty
	_jsii_.Get(
		j,
		"filterOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ProjectOperation() TfDataSet_ProjectOperationPropertyOutputReference {
	var returns TfDataSet_ProjectOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"projectOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ProjectOperationInput() *TfDataSet_ProjectOperationProperty {
	var returns *TfDataSet_ProjectOperationProperty
	_jsii_.Get(
		j,
		"projectOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) RenameColumnOperation() TfDataSet_RenameColumnOperationPropertyOutputReference {
	var returns TfDataSet_RenameColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"renameColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) RenameColumnOperationInput() *TfDataSet_RenameColumnOperationProperty {
	var returns *TfDataSet_RenameColumnOperationProperty
	_jsii_.Get(
		j,
		"renameColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) TagColumnOperation() TfDataSet_TagColumnOperationPropertyOutputReference {
	var returns TfDataSet_TagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"tagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) TagColumnOperationInput() *TfDataSet_TagColumnOperationProperty {
	var returns *TfDataSet_TagColumnOperationProperty
	_jsii_.Get(
		j,
		"tagColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) UntagColumnOperation() TfDataSet_UntagColumnOperationPropertyOutputReference {
	var returns TfDataSet_UntagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"untagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) UntagColumnOperationInput() *TfDataSet_UntagColumnOperationProperty {
	var returns *TfDataSet_UntagColumnOperationProperty
	_jsii_.Get(
		j,
		"untagColumnOperationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSet_DataTransformsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSet_DataTransformsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSet_DataTransformsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSet_DataTransformsPropertyOutputReference_Override(t TfDataSet_DataTransformsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutCastColumnTypeOperation(value *TfDataSet_CastColumnTypeOperationProperty) {
	if err := t.validatePutCastColumnTypeOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCastColumnTypeOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutCreateColumnsOperation(value *TfDataSet_CreateColumnsOperationProperty) {
	if err := t.validatePutCreateColumnsOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCreateColumnsOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutFilterOperation(value *TfDataSet_FilterOperationProperty) {
	if err := t.validatePutFilterOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putFilterOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutProjectOperation(value *TfDataSet_ProjectOperationProperty) {
	if err := t.validatePutProjectOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProjectOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutRenameColumnOperation(value *TfDataSet_RenameColumnOperationProperty) {
	if err := t.validatePutRenameColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRenameColumnOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutTagColumnOperation(value *TfDataSet_TagColumnOperationProperty) {
	if err := t.validatePutTagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTagColumnOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) PutUntagColumnOperation(value *TfDataSet_UntagColumnOperationProperty) {
	if err := t.validatePutUntagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUntagColumnOperation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetCastColumnTypeOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetCastColumnTypeOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetCreateColumnsOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateColumnsOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetFilterOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetFilterOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetProjectOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetProjectOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetRenameColumnOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetRenameColumnOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetTagColumnOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetTagColumnOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ResetUntagColumnOperation() {
	_jsii_.InvokeVoid(
		t,
		"resetUntagColumnOperation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSet_DataTransformsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

