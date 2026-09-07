package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDatasource_DynamodbConfigPropertyOutputReference interface {
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
	DeltaSyncConfig() AwsDatasource_DeltaSyncConfigPropertyOutputReference
	// Experimental.
	DeltaSyncConfigInput() *AwsDatasource_DeltaSyncConfigProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDatasource_DynamodbConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDatasource_DynamodbConfigProperty)
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
	PutDeltaSyncConfig(value *AwsDatasource_DeltaSyncConfigProperty)
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

// The jsii proxy struct for AwsDatasource_DynamodbConfigPropertyOutputReference
type jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) DeltaSyncConfig() AwsDatasource_DeltaSyncConfigPropertyOutputReference {
	var returns AwsDatasource_DeltaSyncConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"deltaSyncConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) DeltaSyncConfigInput() *AwsDatasource_DeltaSyncConfigProperty {
	var returns *AwsDatasource_DeltaSyncConfigProperty
	_jsii_.Get(
		j,
		"deltaSyncConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) InternalValue() *AwsDatasource_DynamodbConfigProperty {
	var returns *AwsDatasource_DynamodbConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) Region() *string {
	var returns *string
	_jsii_.Get(
		j,
		"region",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) RegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"regionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) TableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) TableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) UseCallerCredentials() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentials",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) UseCallerCredentialsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCallerCredentialsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) Versioned() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versioned",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) VersionedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"versionedInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDatasource_DynamodbConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDatasource_DynamodbConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDatasource_DynamodbConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsDatasource.DynamodbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDatasource_DynamodbConfigPropertyOutputReference_Override(a AwsDatasource_DynamodbConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsDatasource.DynamodbConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetInternalValue(val *AwsDatasource_DynamodbConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetRegion(val *string) {
	if err := j.validateSetRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"region",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetTableName(val *string) {
	if err := j.validateSetTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableName",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetUseCallerCredentials(val interface{}) {
	if err := j.validateSetUseCallerCredentialsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCallerCredentials",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference)SetVersioned(val interface{}) {
	if err := j.validateSetVersionedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versioned",
		val,
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) PutDeltaSyncConfig(value *AwsDatasource_DeltaSyncConfigProperty) {
	if err := a.validatePutDeltaSyncConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDeltaSyncConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ResetDeltaSyncConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ResetRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ResetUseCallerCredentials() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCallerCredentials",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ResetVersioned() {
	_jsii_.InvokeVoid(
		a,
		"resetVersioned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDatasource_DynamodbConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

