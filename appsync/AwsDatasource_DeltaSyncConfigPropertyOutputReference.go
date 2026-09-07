package appsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appsync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appsync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDatasource_DeltaSyncConfigPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BaseTableTtl() *float64
	// Experimental.
	SetBaseTableTtl(val *float64)
	// Experimental.
	BaseTableTtlInput() *float64
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
	DeltaSyncTableName() *string
	// Experimental.
	SetDeltaSyncTableName(val *string)
	// Experimental.
	DeltaSyncTableNameInput() *string
	// Experimental.
	DeltaSyncTableTtl() *float64
	// Experimental.
	SetDeltaSyncTableTtl(val *float64)
	// Experimental.
	DeltaSyncTableTtlInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDatasource_DeltaSyncConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDatasource_DeltaSyncConfigProperty)
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
	ResetBaseTableTtl()
	// Experimental.
	ResetDeltaSyncTableTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDatasource_DeltaSyncConfigPropertyOutputReference
type jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) BaseTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) BaseTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"baseTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) DeltaSyncTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) DeltaSyncTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaSyncTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) DeltaSyncTableTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) DeltaSyncTableTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deltaSyncTableTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) InternalValue() *AwsDatasource_DeltaSyncConfigProperty {
	var returns *AwsDatasource_DeltaSyncConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDatasource_DeltaSyncConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDatasource_DeltaSyncConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDatasource_DeltaSyncConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsDatasource.DeltaSyncConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDatasource_DeltaSyncConfigPropertyOutputReference_Override(a AwsDatasource_DeltaSyncConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-appsync.AwsDatasource.DeltaSyncConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetBaseTableTtl(val *float64) {
	if err := j.validateSetBaseTableTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseTableTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetDeltaSyncTableName(val *string) {
	if err := j.validateSetDeltaSyncTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSyncTableName",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetDeltaSyncTableTtl(val *float64) {
	if err := j.validateSetDeltaSyncTableTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaSyncTableTtl",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetInternalValue(val *AwsDatasource_DeltaSyncConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ResetBaseTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetBaseTableTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ResetDeltaSyncTableTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDeltaSyncTableTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDatasource_DeltaSyncConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

