package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/quicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/quicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSet_DataTransformsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CastColumnTypeOperation() AwsDataSet_CastColumnTypeOperationPropertyOutputReference
	// Experimental.
	CastColumnTypeOperationInput() *AwsDataSet_CastColumnTypeOperationProperty
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
	CreateColumnsOperation() AwsDataSet_CreateColumnsOperationPropertyOutputReference
	// Experimental.
	CreateColumnsOperationInput() *AwsDataSet_CreateColumnsOperationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FilterOperation() AwsDataSet_FilterOperationPropertyOutputReference
	// Experimental.
	FilterOperationInput() *AwsDataSet_FilterOperationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ProjectOperation() AwsDataSet_ProjectOperationPropertyOutputReference
	// Experimental.
	ProjectOperationInput() *AwsDataSet_ProjectOperationProperty
	// Experimental.
	RenameColumnOperation() AwsDataSet_RenameColumnOperationPropertyOutputReference
	// Experimental.
	RenameColumnOperationInput() *AwsDataSet_RenameColumnOperationProperty
	// Experimental.
	TagColumnOperation() AwsDataSet_TagColumnOperationPropertyOutputReference
	// Experimental.
	TagColumnOperationInput() *AwsDataSet_TagColumnOperationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UntagColumnOperation() AwsDataSet_UntagColumnOperationPropertyOutputReference
	// Experimental.
	UntagColumnOperationInput() *AwsDataSet_UntagColumnOperationProperty
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
	PutCastColumnTypeOperation(value *AwsDataSet_CastColumnTypeOperationProperty)
	// Experimental.
	PutCreateColumnsOperation(value *AwsDataSet_CreateColumnsOperationProperty)
	// Experimental.
	PutFilterOperation(value *AwsDataSet_FilterOperationProperty)
	// Experimental.
	PutProjectOperation(value *AwsDataSet_ProjectOperationProperty)
	// Experimental.
	PutRenameColumnOperation(value *AwsDataSet_RenameColumnOperationProperty)
	// Experimental.
	PutTagColumnOperation(value *AwsDataSet_TagColumnOperationProperty)
	// Experimental.
	PutUntagColumnOperation(value *AwsDataSet_UntagColumnOperationProperty)
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

// The jsii proxy struct for AwsDataSet_DataTransformsPropertyOutputReference
type jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperation() AwsDataSet_CastColumnTypeOperationPropertyOutputReference {
	var returns AwsDataSet_CastColumnTypeOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"castColumnTypeOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperationInput() *AwsDataSet_CastColumnTypeOperationProperty {
	var returns *AwsDataSet_CastColumnTypeOperationProperty
	_jsii_.Get(
		j,
		"castColumnTypeOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperation() AwsDataSet_CreateColumnsOperationPropertyOutputReference {
	var returns AwsDataSet_CreateColumnsOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"createColumnsOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperationInput() *AwsDataSet_CreateColumnsOperationProperty {
	var returns *AwsDataSet_CreateColumnsOperationProperty
	_jsii_.Get(
		j,
		"createColumnsOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) FilterOperation() AwsDataSet_FilterOperationPropertyOutputReference {
	var returns AwsDataSet_FilterOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"filterOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) FilterOperationInput() *AwsDataSet_FilterOperationProperty {
	var returns *AwsDataSet_FilterOperationProperty
	_jsii_.Get(
		j,
		"filterOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ProjectOperation() AwsDataSet_ProjectOperationPropertyOutputReference {
	var returns AwsDataSet_ProjectOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"projectOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ProjectOperationInput() *AwsDataSet_ProjectOperationProperty {
	var returns *AwsDataSet_ProjectOperationProperty
	_jsii_.Get(
		j,
		"projectOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) RenameColumnOperation() AwsDataSet_RenameColumnOperationPropertyOutputReference {
	var returns AwsDataSet_RenameColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"renameColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) RenameColumnOperationInput() *AwsDataSet_RenameColumnOperationProperty {
	var returns *AwsDataSet_RenameColumnOperationProperty
	_jsii_.Get(
		j,
		"renameColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) TagColumnOperation() AwsDataSet_TagColumnOperationPropertyOutputReference {
	var returns AwsDataSet_TagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"tagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) TagColumnOperationInput() *AwsDataSet_TagColumnOperationProperty {
	var returns *AwsDataSet_TagColumnOperationProperty
	_jsii_.Get(
		j,
		"tagColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) UntagColumnOperation() AwsDataSet_UntagColumnOperationPropertyOutputReference {
	var returns AwsDataSet_UntagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"untagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) UntagColumnOperationInput() *AwsDataSet_UntagColumnOperationProperty {
	var returns *AwsDataSet_UntagColumnOperationProperty
	_jsii_.Get(
		j,
		"untagColumnOperationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSet_DataTransformsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataSet_DataTransformsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSet_DataTransformsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSet_DataTransformsPropertyOutputReference_Override(a AwsDataSet_DataTransformsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutCastColumnTypeOperation(value *AwsDataSet_CastColumnTypeOperationProperty) {
	if err := a.validatePutCastColumnTypeOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCastColumnTypeOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutCreateColumnsOperation(value *AwsDataSet_CreateColumnsOperationProperty) {
	if err := a.validatePutCreateColumnsOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreateColumnsOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutFilterOperation(value *AwsDataSet_FilterOperationProperty) {
	if err := a.validatePutFilterOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutProjectOperation(value *AwsDataSet_ProjectOperationProperty) {
	if err := a.validatePutProjectOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProjectOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutRenameColumnOperation(value *AwsDataSet_RenameColumnOperationProperty) {
	if err := a.validatePutRenameColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRenameColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutTagColumnOperation(value *AwsDataSet_TagColumnOperationProperty) {
	if err := a.validatePutTagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) PutUntagColumnOperation(value *AwsDataSet_UntagColumnOperationProperty) {
	if err := a.validatePutUntagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUntagColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetCastColumnTypeOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetCastColumnTypeOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetCreateColumnsOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetCreateColumnsOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetFilterOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetProjectOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetProjectOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetRenameColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetRenameColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetTagColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetTagColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ResetUntagColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetUntagColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSet_DataTransformsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

