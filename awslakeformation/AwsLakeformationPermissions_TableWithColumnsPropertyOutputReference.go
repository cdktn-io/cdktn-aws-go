package awslakeformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awslakeformation/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awslakeformation/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CatalogId() *string
	// Experimental.
	SetCatalogId(val *string)
	// Experimental.
	CatalogIdInput() *string
	// Experimental.
	ColumnNames() *[]*string
	// Experimental.
	SetColumnNames(val *[]*string)
	// Experimental.
	ColumnNamesInput() *[]*string
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
	DatabaseName() *string
	// Experimental.
	SetDatabaseName(val *string)
	// Experimental.
	DatabaseNameInput() *string
	// Experimental.
	ExcludedColumnNames() *[]*string
	// Experimental.
	SetExcludedColumnNames(val *[]*string)
	// Experimental.
	ExcludedColumnNamesInput() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsLakeformationPermissions_TableWithColumnsProperty
	// Experimental.
	SetInternalValue(val *AwsLakeformationPermissions_TableWithColumnsProperty)
	// Experimental.
	Name() *string
	// Experimental.
	SetName(val *string)
	// Experimental.
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Wildcard() interface{}
	// Experimental.
	SetWildcard(val interface{})
	// Experimental.
	WildcardInput() interface{}
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
	ResetCatalogId()
	// Experimental.
	ResetColumnNames()
	// Experimental.
	ResetExcludedColumnNames()
	// Experimental.
	ResetWildcard()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference
type jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) CatalogId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) CatalogIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"columnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) DatabaseName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) DatabaseNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"databaseNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ExcludedColumnNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ExcludedColumnNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedColumnNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) InternalValue() *AwsLakeformationPermissions_TableWithColumnsProperty {
	var returns *AwsLakeformationPermissions_TableWithColumnsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) Wildcard() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcard",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) WildcardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"wildcardInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLakeformationPermissions_TableWithColumnsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLakeformationPermissions_TableWithColumnsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-lake-formation.AwsLakeformationPermissions.TableWithColumnsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLakeformationPermissions_TableWithColumnsPropertyOutputReference_Override(a AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-lake-formation.AwsLakeformationPermissions.TableWithColumnsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetCatalogId(val *string) {
	if err := j.validateSetCatalogIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogId",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetColumnNames(val *[]*string) {
	if err := j.validateSetColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnNames",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetDatabaseName(val *string) {
	if err := j.validateSetDatabaseNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"databaseName",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetExcludedColumnNames(val *[]*string) {
	if err := j.validateSetExcludedColumnNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedColumnNames",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetInternalValue(val *AwsLakeformationPermissions_TableWithColumnsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference)SetWildcard(val interface{}) {
	if err := j.validateSetWildcardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wildcard",
		val,
	)
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ResetCatalogId() {
	_jsii_.InvokeVoid(
		a,
		"resetCatalogId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ResetColumnNames() {
	_jsii_.InvokeVoid(
		a,
		"resetColumnNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ResetExcludedColumnNames() {
	_jsii_.InvokeVoid(
		a,
		"resetExcludedColumnNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ResetWildcard() {
	_jsii_.InvokeVoid(
		a,
		"resetWildcard",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLakeformationPermissions_TableWithColumnsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

