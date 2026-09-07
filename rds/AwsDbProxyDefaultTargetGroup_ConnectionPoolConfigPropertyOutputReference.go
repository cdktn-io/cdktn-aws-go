package rds

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/rds/jsii"

	"github.com/cdktn-io/cdktn-aws-go/rds/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference interface {
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
	// Experimental.
	ConnectionBorrowTimeout() *float64
	// Experimental.
	SetConnectionBorrowTimeout(val *float64)
	// Experimental.
	ConnectionBorrowTimeoutInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	// Experimental.
	InitQuery() *string
	// Experimental.
	SetInitQuery(val *string)
	// Experimental.
	InitQueryInput() *string
	// Experimental.
	InternalValue() *AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty
	// Experimental.
	SetInternalValue(val *AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty)
	// Experimental.
	MaxConnectionsPercent() *float64
	// Experimental.
	SetMaxConnectionsPercent(val *float64)
	// Experimental.
	MaxConnectionsPercentInput() *float64
	// Experimental.
	MaxIdleConnectionsPercent() *float64
	// Experimental.
	SetMaxIdleConnectionsPercent(val *float64)
	// Experimental.
	MaxIdleConnectionsPercentInput() *float64
	// Experimental.
	SessionPinningFilters() *[]*string
	// Experimental.
	SetSessionPinningFilters(val *[]*string)
	// Experimental.
	SessionPinningFiltersInput() *[]*string
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
	ResetConnectionBorrowTimeout()
	// Experimental.
	ResetInitQuery()
	// Experimental.
	ResetMaxConnectionsPercent()
	// Experimental.
	ResetMaxIdleConnectionsPercent()
	// Experimental.
	ResetSessionPinningFilters()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference
type jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ConnectionBorrowTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ConnectionBorrowTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionBorrowTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) InitQuery() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQuery",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) InitQueryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"initQueryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) InternalValue() *AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty {
	var returns *AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) MaxConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) MaxConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) MaxIdleConnectionsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) MaxIdleConnectionsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) SessionPinningFilters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFilters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) SessionPinningFiltersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"sessionPinningFiltersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbProxyDefaultTargetGroup.ConnectionPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference_Override(a AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-rds.AwsDbProxyDefaultTargetGroup.ConnectionPoolConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetConnectionBorrowTimeout(val *float64) {
	if err := j.validateSetConnectionBorrowTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionBorrowTimeout",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetInitQuery(val *string) {
	if err := j.validateSetInitQueryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initQuery",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetInternalValue(val *AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetMaxConnectionsPercent(val *float64) {
	if err := j.validateSetMaxConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetMaxIdleConnectionsPercent(val *float64) {
	if err := j.validateSetMaxIdleConnectionsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnectionsPercent",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetSessionPinningFilters(val *[]*string) {
	if err := j.validateSetSessionPinningFiltersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionPinningFilters",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ResetConnectionBorrowTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionBorrowTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ResetInitQuery() {
	_jsii_.InvokeVoid(
		a,
		"resetInitQuery",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ResetMaxConnectionsPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConnectionsPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ResetMaxIdleConnectionsPercent() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxIdleConnectionsPercent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ResetSessionPinningFilters() {
	_jsii_.InvokeVoid(
		a,
		"resetSessionPinningFilters",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDbProxyDefaultTargetGroup_ConnectionPoolConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

