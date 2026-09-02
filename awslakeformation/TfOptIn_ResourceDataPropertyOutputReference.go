package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOptIn_ResourceDataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Catalog() TfOptIn_CatalogPropertyList
	// Experimental.
	CatalogInput() interface{}
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
	Database() TfOptIn_DatabasePropertyList
	// Experimental.
	DatabaseInput() interface{}
	// Experimental.
	DataCellsFilter() TfOptIn_DataCellsFilterPropertyList
	// Experimental.
	DataCellsFilterInput() interface{}
	// Experimental.
	DataLocation() TfOptIn_DataLocationPropertyList
	// Experimental.
	DataLocationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LfTag() TfOptIn_LfTagPropertyList
	// Experimental.
	LfTagExpression() TfOptIn_LfTagExpressionPropertyList
	// Experimental.
	LfTagExpressionInput() interface{}
	// Experimental.
	LfTagInput() interface{}
	// Experimental.
	LfTagPolicy() TfOptIn_LfTagPolicyPropertyList
	// Experimental.
	LfTagPolicyInput() interface{}
	// Experimental.
	Table() TfOptIn_TablePropertyList
	// Experimental.
	TableInput() interface{}
	// Experimental.
	TableWithColumns() TfOptIn_TableWithColumnsPropertyList
	// Experimental.
	TableWithColumnsInput() interface{}
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
	PutCatalog(value interface{})
	// Experimental.
	PutDatabase(value interface{})
	// Experimental.
	PutDataCellsFilter(value interface{})
	// Experimental.
	PutDataLocation(value interface{})
	// Experimental.
	PutLfTag(value interface{})
	// Experimental.
	PutLfTagExpression(value interface{})
	// Experimental.
	PutLfTagPolicy(value interface{})
	// Experimental.
	PutTable(value interface{})
	// Experimental.
	PutTableWithColumns(value interface{})
	// Experimental.
	ResetCatalog()
	// Experimental.
	ResetDatabase()
	// Experimental.
	ResetDataCellsFilter()
	// Experimental.
	ResetDataLocation()
	// Experimental.
	ResetLfTag()
	// Experimental.
	ResetLfTagExpression()
	// Experimental.
	ResetLfTagPolicy()
	// Experimental.
	ResetTable()
	// Experimental.
	ResetTableWithColumns()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOptIn_ResourceDataPropertyOutputReference
type jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) Catalog() TfOptIn_CatalogPropertyList {
	var returns TfOptIn_CatalogPropertyList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) CatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) Database() TfOptIn_DatabasePropertyList {
	var returns TfOptIn_DatabasePropertyList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) DataCellsFilter() TfOptIn_DataCellsFilterPropertyList {
	var returns TfOptIn_DataCellsFilterPropertyList
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) DataCellsFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) DataLocation() TfOptIn_DataLocationPropertyList {
	var returns TfOptIn_DataLocationPropertyList
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) DataLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTag() TfOptIn_LfTagPropertyList {
	var returns TfOptIn_LfTagPropertyList
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTagExpression() TfOptIn_LfTagExpressionPropertyList {
	var returns TfOptIn_LfTagExpressionPropertyList
	_jsii_.Get(
		j,
		"lfTagExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTagExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTagPolicy() TfOptIn_LfTagPolicyPropertyList {
	var returns TfOptIn_LfTagPolicyPropertyList
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) LfTagPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) Table() TfOptIn_TablePropertyList {
	var returns TfOptIn_TablePropertyList
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) TableWithColumns() TfOptIn_TableWithColumnsPropertyList {
	var returns TfOptIn_TableWithColumnsPropertyList
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) TableWithColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOptIn_ResourceDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfOptIn_ResourceDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOptIn_ResourceDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfOptIn.ResourceDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOptIn_ResourceDataPropertyOutputReference_Override(t TfOptIn_ResourceDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.TfOptIn.ResourceDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutCatalog(value interface{}) {
	if err := t.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCatalog",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutDatabase(value interface{}) {
	if err := t.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDatabase",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutDataCellsFilter(value interface{}) {
	if err := t.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutDataLocation(value interface{}) {
	if err := t.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutLfTag(value interface{}) {
	if err := t.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLfTag",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutLfTagExpression(value interface{}) {
	if err := t.validatePutLfTagExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLfTagExpression",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutLfTagPolicy(value interface{}) {
	if err := t.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutTable(value interface{}) {
	if err := t.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) PutTableWithColumns(value interface{}) {
	if err := t.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalog",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		t,
		"resetDatabase",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetDataLocation() {
	_jsii_.InvokeVoid(
		t,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetLfTag() {
	_jsii_.InvokeVoid(
		t,
		"resetLfTag",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetLfTagExpression() {
	_jsii_.InvokeVoid(
		t,
		"resetLfTagExpression",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		t,
		"resetTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOptIn_ResourceDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

