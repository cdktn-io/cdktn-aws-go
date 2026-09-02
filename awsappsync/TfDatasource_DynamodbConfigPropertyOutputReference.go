package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDatasource_DynamodbConfigPropertyOutputReference interface {
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
	DeltaSyncConfig() TfDatasource_DeltaSyncConfigPropertyOutputReference
	// Experimental.
	DeltaSyncConfigInput() *TfDatasource_DeltaSyncConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDatasource_DynamodbConfigProperty
	// Experimental.
	SetInternalValue(val *TfDatasource_DynamodbConfigProperty)
	// Experimental.
	Region() *string
	// Experimental.
	SetRegion(val *string)
	// Experimental.
	RegionInput() *string
	// Experimental.
	TableName() *string
	// Experimental.
	SetTableName(val *string)
	// Experimental.
	TableNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseCallerCredentials() interface{}
	// Experimental.
	SetUseCallerCredentials(val interface{})
	// Experimental.
	UseCallerCredentialsInput() interface{}
	// Experimental.
	Versioned() interface{}
	// Experimental.
	SetVersioned(val interface{})
	// Experimental.
	VersionedInput() interface{}
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
	PutDeltaSyncConfig(value *TfDatasource_DeltaSyncConfigProperty)
	// Experimental.
	ResetDeltaSyncConfig()
	// Experimental.
	ResetRegion()
	// Experimental.
	ResetUseCallerCredentials()
	// Experimental.
	ResetVersioned()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDatasource_DynamodbConfigPropertyOutputReference
type jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) DeltaSyncConfig() TfDatasource_DeltaSyncConfigPropertyOutputReference {
	var returns TfDatasource_DeltaSyncConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deltaSyncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) DeltaSyncConfigInput() *TfDatasource_DeltaSyncConfigProperty {
	var returns *TfDatasource_DeltaSyncConfigProperty
	_jsii_.Get(
		j,
		"deltaSyncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) InternalValue() *TfDatasource_DynamodbConfigProperty {
	var returns *TfDatasource_DynamodbConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) UseCallerCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) UseCallerCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) Versioned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) VersionedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionedInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDatasource_DynamodbConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDatasource_DynamodbConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDatasource_DynamodbConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.TfDatasource.DynamodbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDatasource_DynamodbConfigPropertyOutputReference_Override(t TfDatasource_DynamodbConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.TfDatasource.DynamodbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetInternalValue(val *TfDatasource_DynamodbConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetUseCallerCredentials(val interface{}) {
	if err := j.validateSetUseCallerCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCallerCredentials",
		val,
	)
}

func (j *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference)SetVersioned(val interface{}) {
	if err := j.validateSetVersionedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versioned",
		val,
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) PutDeltaSyncConfig(value *TfDatasource_DeltaSyncConfigProperty) {
	if err := t.validatePutDeltaSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDeltaSyncConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ResetDeltaSyncConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDeltaSyncConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		t,
		"resetRegion",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ResetUseCallerCredentials() {
	_jsii_.InvokeVoid(
		t,
		"resetUseCallerCredentials",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ResetVersioned() {
	_jsii_.InvokeVoid(
		t,
		"resetVersioned",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDatasource_DynamodbConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

