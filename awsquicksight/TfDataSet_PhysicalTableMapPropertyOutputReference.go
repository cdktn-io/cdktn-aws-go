package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataSet_PhysicalTableMapPropertyOutputReference interface {
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
	CustomSql() TfDataSet_CustomSqlPropertyOutputReference
	// Experimental.
	CustomSqlInput() *TfDataSet_CustomSqlProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PhysicalTableMapId() *string
	// Experimental.
	SetPhysicalTableMapId(val *string)
	// Experimental.
	PhysicalTableMapIdInput() *string
	// Experimental.
	RelationalTable() TfDataSet_RelationalTablePropertyOutputReference
	// Experimental.
	RelationalTableInput() *TfDataSet_RelationalTableProperty
	// Experimental.
	S3Source() TfDataSet_S3SourcePropertyOutputReference
	// Experimental.
	S3SourceInput() *TfDataSet_S3SourceProperty
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
	PutCustomSql(value *TfDataSet_CustomSqlProperty)
	// Experimental.
	PutRelationalTable(value *TfDataSet_RelationalTableProperty)
	// Experimental.
	PutS3Source(value *TfDataSet_S3SourceProperty)
	// Experimental.
	ResetCustomSql()
	// Experimental.
	ResetRelationalTable()
	// Experimental.
	ResetS3Source()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataSet_PhysicalTableMapPropertyOutputReference
type jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) CustomSql() TfDataSet_CustomSqlPropertyOutputReference {
	var returns TfDataSet_CustomSqlPropertyOutputReference
	_jsii_.Get(
		j,
		"customSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) CustomSqlInput() *TfDataSet_CustomSqlProperty {
	var returns *TfDataSet_CustomSqlProperty
	_jsii_.Get(
		j,
		"customSqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) PhysicalTableMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) PhysicalTableMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) RelationalTable() TfDataSet_RelationalTablePropertyOutputReference {
	var returns TfDataSet_RelationalTablePropertyOutputReference
	_jsii_.Get(
		j,
		"relationalTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) RelationalTableInput() *TfDataSet_RelationalTableProperty {
	var returns *TfDataSet_RelationalTableProperty
	_jsii_.Get(
		j,
		"relationalTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) S3Source() TfDataSet_S3SourcePropertyOutputReference {
	var returns TfDataSet_S3SourcePropertyOutputReference
	_jsii_.Get(
		j,
		"s3Source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) S3SourceInput() *TfDataSet_S3SourceProperty {
	var returns *TfDataSet_S3SourceProperty
	_jsii_.Get(
		j,
		"s3SourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataSet_PhysicalTableMapPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfDataSet_PhysicalTableMapPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataSet_PhysicalTableMapPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.PhysicalTableMapPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataSet_PhysicalTableMapPropertyOutputReference_Override(t TfDataSet_PhysicalTableMapPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDataSet.PhysicalTableMapPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetPhysicalTableMapId(val *string) {
	if err := j.validateSetPhysicalTableMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalTableMapId",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) PutCustomSql(value *TfDataSet_CustomSqlProperty) {
	if err := t.validatePutCustomSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCustomSql",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) PutRelationalTable(value *TfDataSet_RelationalTableProperty) {
	if err := t.validatePutRelationalTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelationalTable",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) PutS3Source(value *TfDataSet_S3SourceProperty) {
	if err := t.validatePutS3SourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Source",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ResetCustomSql() {
	_jsii_.InvokeVoid(
		t,
		"resetCustomSql",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ResetRelationalTable() {
	_jsii_.InvokeVoid(
		t,
		"resetRelationalTable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ResetS3Source() {
	_jsii_.InvokeVoid(
		t,
		"resetS3Source",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataSet_PhysicalTableMapPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

