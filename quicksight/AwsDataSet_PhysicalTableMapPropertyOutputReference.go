package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/quicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/quicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDataSet_PhysicalTableMapPropertyOutputReference interface {
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
	CustomSql() AwsDataSet_CustomSqlPropertyOutputReference
	// Experimental.
	CustomSqlInput() *AwsDataSet_CustomSqlProperty
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
	RelationalTable() AwsDataSet_RelationalTablePropertyOutputReference
	// Experimental.
	RelationalTableInput() *AwsDataSet_RelationalTableProperty
	// Experimental.
	S3Source() AwsDataSet_S3SourcePropertyOutputReference
	// Experimental.
	S3SourceInput() *AwsDataSet_S3SourceProperty
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
	PutCustomSql(value *AwsDataSet_CustomSqlProperty)
	// Experimental.
	PutRelationalTable(value *AwsDataSet_RelationalTableProperty)
	// Experimental.
	PutS3Source(value *AwsDataSet_S3SourceProperty)
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

// The jsii proxy struct for AwsDataSet_PhysicalTableMapPropertyOutputReference
type jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) CustomSql() AwsDataSet_CustomSqlPropertyOutputReference {
	var returns AwsDataSet_CustomSqlPropertyOutputReference
	_jsii_.Get(
		j,
		"customSql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) CustomSqlInput() *AwsDataSet_CustomSqlProperty {
	var returns *AwsDataSet_CustomSqlProperty
	_jsii_.Get(
		j,
		"customSqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) PhysicalTableMapId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) PhysicalTableMapIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalTableMapIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) RelationalTable() AwsDataSet_RelationalTablePropertyOutputReference {
	var returns AwsDataSet_RelationalTablePropertyOutputReference
	_jsii_.Get(
		j,
		"relationalTable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) RelationalTableInput() *AwsDataSet_RelationalTableProperty {
	var returns *AwsDataSet_RelationalTableProperty
	_jsii_.Get(
		j,
		"relationalTableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) S3Source() AwsDataSet_S3SourcePropertyOutputReference {
	var returns AwsDataSet_S3SourcePropertyOutputReference
	_jsii_.Get(
		j,
		"s3Source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) S3SourceInput() *AwsDataSet_S3SourceProperty {
	var returns *AwsDataSet_S3SourceProperty
	_jsii_.Get(
		j,
		"s3SourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDataSet_PhysicalTableMapPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsDataSet_PhysicalTableMapPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDataSet_PhysicalTableMapPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSet.PhysicalTableMapPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDataSet_PhysicalTableMapPropertyOutputReference_Override(a AwsDataSet_PhysicalTableMapPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDataSet.PhysicalTableMapPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetPhysicalTableMapId(val *string) {
	if err := j.validateSetPhysicalTableMapIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"physicalTableMapId",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) PutCustomSql(value *AwsDataSet_CustomSqlProperty) {
	if err := a.validatePutCustomSqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCustomSql",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) PutRelationalTable(value *AwsDataSet_RelationalTableProperty) {
	if err := a.validatePutRelationalTableParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelationalTable",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) PutS3Source(value *AwsDataSet_S3SourceProperty) {
	if err := a.validatePutS3SourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Source",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ResetCustomSql() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomSql",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ResetRelationalTable() {
	_jsii_.InvokeVoid(
		a,
		"resetRelationalTable",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ResetS3Source() {
	_jsii_.InvokeVoid(
		a,
		"resetS3Source",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDataSet_PhysicalTableMapPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

