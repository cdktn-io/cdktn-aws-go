package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLakeformationOptIn_ResourceDataPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Catalog() AwsLakeformationOptIn_CatalogPropertyList
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
	Database() AwsLakeformationOptIn_DatabasePropertyList
	// Experimental.
	DatabaseInput() interface{}
	// Experimental.
	DataCellsFilter() AwsLakeformationOptIn_DataCellsFilterPropertyList
	// Experimental.
	DataCellsFilterInput() interface{}
	// Experimental.
	DataLocation() AwsLakeformationOptIn_DataLocationPropertyList
	// Experimental.
	DataLocationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	LfTag() AwsLakeformationOptIn_LfTagPropertyList
	// Experimental.
	LfTagExpression() AwsLakeformationOptIn_LfTagExpressionPropertyList
	// Experimental.
	LfTagExpressionInput() interface{}
	// Experimental.
	LfTagInput() interface{}
	// Experimental.
	LfTagPolicy() AwsLakeformationOptIn_LfTagPolicyPropertyList
	// Experimental.
	LfTagPolicyInput() interface{}
	// Experimental.
	Table() AwsLakeformationOptIn_TablePropertyList
	// Experimental.
	TableInput() interface{}
	// Experimental.
	TableWithColumns() AwsLakeformationOptIn_TableWithColumnsPropertyList
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

// The jsii proxy struct for AwsLakeformationOptIn_ResourceDataPropertyOutputReference
type jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) Catalog() AwsLakeformationOptIn_CatalogPropertyList {
	var returns AwsLakeformationOptIn_CatalogPropertyList
	_jsii_.Get(
		j,
		"catalog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) CatalogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"catalogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) Database() AwsLakeformationOptIn_DatabasePropertyList {
	var returns AwsLakeformationOptIn_DatabasePropertyList
	_jsii_.Get(
		j,
		"database",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) DatabaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"databaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) DataCellsFilter() AwsLakeformationOptIn_DataCellsFilterPropertyList {
	var returns AwsLakeformationOptIn_DataCellsFilterPropertyList
	_jsii_.Get(
		j,
		"dataCellsFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) DataCellsFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataCellsFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) DataLocation() AwsLakeformationOptIn_DataLocationPropertyList {
	var returns AwsLakeformationOptIn_DataLocationPropertyList
	_jsii_.Get(
		j,
		"dataLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) DataLocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTag() AwsLakeformationOptIn_LfTagPropertyList {
	var returns AwsLakeformationOptIn_LfTagPropertyList
	_jsii_.Get(
		j,
		"lfTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTagExpression() AwsLakeformationOptIn_LfTagExpressionPropertyList {
	var returns AwsLakeformationOptIn_LfTagExpressionPropertyList
	_jsii_.Get(
		j,
		"lfTagExpression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTagExpressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagExpressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTagInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTagPolicy() AwsLakeformationOptIn_LfTagPolicyPropertyList {
	var returns AwsLakeformationOptIn_LfTagPolicyPropertyList
	_jsii_.Get(
		j,
		"lfTagPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) LfTagPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lfTagPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) Table() AwsLakeformationOptIn_TablePropertyList {
	var returns AwsLakeformationOptIn_TablePropertyList
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) TableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) TableWithColumns() AwsLakeformationOptIn_TableWithColumnsPropertyList {
	var returns AwsLakeformationOptIn_TableWithColumnsPropertyList
	_jsii_.Get(
		j,
		"tableWithColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) TableWithColumnsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tableWithColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLakeformationOptIn_ResourceDataPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsLakeformationOptIn_ResourceDataPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLakeformationOptIn_ResourceDataPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.AwsLakeformationOptIn.ResourceDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLakeformationOptIn_ResourceDataPropertyOutputReference_Override(a AwsLakeformationOptIn_ResourceDataPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.AwsLakeformationOptIn.ResourceDataPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutCatalog(value interface{}) {
	if err := a.validatePutCatalogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCatalog",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutDatabase(value interface{}) {
	if err := a.validatePutDatabaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDatabase",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutDataCellsFilter(value interface{}) {
	if err := a.validatePutDataCellsFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataCellsFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutDataLocation(value interface{}) {
	if err := a.validatePutDataLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataLocation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutLfTag(value interface{}) {
	if err := a.validatePutLfTagParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLfTag",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutLfTagExpression(value interface{}) {
	if err := a.validatePutLfTagExpressionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLfTagExpression",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutLfTagPolicy(value interface{}) {
	if err := a.validatePutLfTagPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putLfTagPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutTable(value interface{}) {
	if err := a.validatePutTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) PutTableWithColumns(value interface{}) {
	if err := a.validatePutTableWithColumnsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTableWithColumns",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetCatalog() {
	_jsii_.InvokeVoid(
		a,
		"resetCatalog",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetDatabase() {
	_jsii_.InvokeVoid(
		a,
		"resetDatabase",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetDataCellsFilter() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCellsFilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetDataLocation() {
	_jsii_.InvokeVoid(
		a,
		"resetDataLocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetLfTag() {
	_jsii_.InvokeVoid(
		a,
		"resetLfTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetLfTagExpression() {
	_jsii_.InvokeVoid(
		a,
		"resetLfTagExpression",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetLfTagPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetLfTagPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetTable() {
	_jsii_.InvokeVoid(
		a,
		"resetTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ResetTableWithColumns() {
	_jsii_.InvokeVoid(
		a,
		"resetTableWithColumns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLakeformationOptIn_ResourceDataPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

