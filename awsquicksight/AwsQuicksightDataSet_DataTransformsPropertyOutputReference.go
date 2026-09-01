package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightDataSet_DataTransformsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CastColumnTypeOperation() AwsQuicksightDataSet_CastColumnTypeOperationPropertyOutputReference
	// Experimental.
	CastColumnTypeOperationInput() *AwsQuicksightDataSet_CastColumnTypeOperationProperty
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
	CreateColumnsOperation() AwsQuicksightDataSet_CreateColumnsOperationPropertyOutputReference
	// Experimental.
	CreateColumnsOperationInput() *AwsQuicksightDataSet_CreateColumnsOperationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FilterOperation() AwsQuicksightDataSet_FilterOperationPropertyOutputReference
	// Experimental.
	FilterOperationInput() *AwsQuicksightDataSet_FilterOperationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	ProjectOperation() AwsQuicksightDataSet_ProjectOperationPropertyOutputReference
	// Experimental.
	ProjectOperationInput() *AwsQuicksightDataSet_ProjectOperationProperty
	// Experimental.
	RenameColumnOperation() AwsQuicksightDataSet_RenameColumnOperationPropertyOutputReference
	// Experimental.
	RenameColumnOperationInput() *AwsQuicksightDataSet_RenameColumnOperationProperty
	// Experimental.
	TagColumnOperation() AwsQuicksightDataSet_TagColumnOperationPropertyOutputReference
	// Experimental.
	TagColumnOperationInput() *AwsQuicksightDataSet_TagColumnOperationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UntagColumnOperation() AwsQuicksightDataSet_UntagColumnOperationPropertyOutputReference
	// Experimental.
	UntagColumnOperationInput() *AwsQuicksightDataSet_UntagColumnOperationProperty
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
	PutCastColumnTypeOperation(value *AwsQuicksightDataSet_CastColumnTypeOperationProperty)
	// Experimental.
	PutCreateColumnsOperation(value *AwsQuicksightDataSet_CreateColumnsOperationProperty)
	// Experimental.
	PutFilterOperation(value *AwsQuicksightDataSet_FilterOperationProperty)
	// Experimental.
	PutProjectOperation(value *AwsQuicksightDataSet_ProjectOperationProperty)
	// Experimental.
	PutRenameColumnOperation(value *AwsQuicksightDataSet_RenameColumnOperationProperty)
	// Experimental.
	PutTagColumnOperation(value *AwsQuicksightDataSet_TagColumnOperationProperty)
	// Experimental.
	PutUntagColumnOperation(value *AwsQuicksightDataSet_UntagColumnOperationProperty)
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

// The jsii proxy struct for AwsQuicksightDataSet_DataTransformsPropertyOutputReference
type jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperation() AwsQuicksightDataSet_CastColumnTypeOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_CastColumnTypeOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"castColumnTypeOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) CastColumnTypeOperationInput() *AwsQuicksightDataSet_CastColumnTypeOperationProperty {
	var returns *AwsQuicksightDataSet_CastColumnTypeOperationProperty
	_jsii_.Get(
		j,
		"castColumnTypeOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperation() AwsQuicksightDataSet_CreateColumnsOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_CreateColumnsOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"createColumnsOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) CreateColumnsOperationInput() *AwsQuicksightDataSet_CreateColumnsOperationProperty {
	var returns *AwsQuicksightDataSet_CreateColumnsOperationProperty
	_jsii_.Get(
		j,
		"createColumnsOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) FilterOperation() AwsQuicksightDataSet_FilterOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_FilterOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"filterOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) FilterOperationInput() *AwsQuicksightDataSet_FilterOperationProperty {
	var returns *AwsQuicksightDataSet_FilterOperationProperty
	_jsii_.Get(
		j,
		"filterOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ProjectOperation() AwsQuicksightDataSet_ProjectOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_ProjectOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"projectOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ProjectOperationInput() *AwsQuicksightDataSet_ProjectOperationProperty {
	var returns *AwsQuicksightDataSet_ProjectOperationProperty
	_jsii_.Get(
		j,
		"projectOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) RenameColumnOperation() AwsQuicksightDataSet_RenameColumnOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_RenameColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"renameColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) RenameColumnOperationInput() *AwsQuicksightDataSet_RenameColumnOperationProperty {
	var returns *AwsQuicksightDataSet_RenameColumnOperationProperty
	_jsii_.Get(
		j,
		"renameColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) TagColumnOperation() AwsQuicksightDataSet_TagColumnOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_TagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"tagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) TagColumnOperationInput() *AwsQuicksightDataSet_TagColumnOperationProperty {
	var returns *AwsQuicksightDataSet_TagColumnOperationProperty
	_jsii_.Get(
		j,
		"tagColumnOperationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) UntagColumnOperation() AwsQuicksightDataSet_UntagColumnOperationPropertyOutputReference {
	var returns AwsQuicksightDataSet_UntagColumnOperationPropertyOutputReference
	_jsii_.Get(
		j,
		"untagColumnOperation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) UntagColumnOperationInput() *AwsQuicksightDataSet_UntagColumnOperationProperty {
	var returns *AwsQuicksightDataSet_UntagColumnOperationProperty
	_jsii_.Get(
		j,
		"untagColumnOperationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightDataSet_DataTransformsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsQuicksightDataSet_DataTransformsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightDataSet_DataTransformsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightDataSet_DataTransformsPropertyOutputReference_Override(a AwsQuicksightDataSet_DataTransformsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDataSet.DataTransformsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutCastColumnTypeOperation(value *AwsQuicksightDataSet_CastColumnTypeOperationProperty) {
	if err := a.validatePutCastColumnTypeOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCastColumnTypeOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutCreateColumnsOperation(value *AwsQuicksightDataSet_CreateColumnsOperationProperty) {
	if err := a.validatePutCreateColumnsOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCreateColumnsOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutFilterOperation(value *AwsQuicksightDataSet_FilterOperationProperty) {
	if err := a.validatePutFilterOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFilterOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutProjectOperation(value *AwsQuicksightDataSet_ProjectOperationProperty) {
	if err := a.validatePutProjectOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putProjectOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutRenameColumnOperation(value *AwsQuicksightDataSet_RenameColumnOperationProperty) {
	if err := a.validatePutRenameColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRenameColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutTagColumnOperation(value *AwsQuicksightDataSet_TagColumnOperationProperty) {
	if err := a.validatePutTagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTagColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) PutUntagColumnOperation(value *AwsQuicksightDataSet_UntagColumnOperationProperty) {
	if err := a.validatePutUntagColumnOperationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUntagColumnOperation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetCastColumnTypeOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetCastColumnTypeOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetCreateColumnsOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetCreateColumnsOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetFilterOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetFilterOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetProjectOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetProjectOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetRenameColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetRenameColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetTagColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetTagColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ResetUntagColumnOperation() {
	_jsii_.InvokeVoid(
		a,
		"resetUntagColumnOperation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuicksightDataSet_DataTransformsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

