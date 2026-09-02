package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsservicecatalog/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsservicecatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Accounts() *[]*string
	// Experimental.
	SetAccounts(val *[]*string)
	// Experimental.
	AccountsInput() *[]*string
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
	FailureToleranceCount() *float64
	// Experimental.
	SetFailureToleranceCount(val *float64)
	// Experimental.
	FailureToleranceCountInput() *float64
	// Experimental.
	FailureTolerancePercentage() *float64
	// Experimental.
	SetFailureTolerancePercentage(val *float64)
	// Experimental.
	FailureTolerancePercentageInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfProvisionedProduct_StackSetProvisioningPreferencesProperty
	// Experimental.
	SetInternalValue(val *TfProvisionedProduct_StackSetProvisioningPreferencesProperty)
	// Experimental.
	MaxConcurrencyCount() *float64
	// Experimental.
	SetMaxConcurrencyCount(val *float64)
	// Experimental.
	MaxConcurrencyCountInput() *float64
	// Experimental.
	MaxConcurrencyPercentage() *float64
	// Experimental.
	SetMaxConcurrencyPercentage(val *float64)
	// Experimental.
	MaxConcurrencyPercentageInput() *float64
	// Experimental.
	Regions() *[]*string
	// Experimental.
	SetRegions(val *[]*string)
	// Experimental.
	RegionsInput() *[]*string
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
	ResetAccounts()
	// Experimental.
	ResetFailureToleranceCount()
	// Experimental.
	ResetFailureTolerancePercentage()
	// Experimental.
	ResetMaxConcurrencyCount()
	// Experimental.
	ResetMaxConcurrencyPercentage()
	// Experimental.
	ResetRegions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference
type jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InternalValue() *TfProvisionedProduct_StackSetProvisioningPreferencesProperty {
	var returns *TfProvisionedProduct_StackSetProvisioningPreferencesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-service-catalog.TfProvisionedProduct.StackSetProvisioningPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference_Override(t TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-service-catalog.TfProvisionedProduct.StackSetProvisioningPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetInternalValue(val *TfProvisionedProduct_StackSetProvisioningPreferencesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetMaxConcurrencyCount(val *float64) {
	if err := j.validateSetMaxConcurrencyCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyCount",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetMaxConcurrencyPercentage(val *float64) {
	if err := j.validateSetMaxConcurrencyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyPercentage",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		t,
		"resetAccounts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetMaxConcurrencyCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrencyCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetMaxConcurrencyPercentage() {
	_jsii_.InvokeVoid(
		t,
		"resetMaxConcurrencyPercentage",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		t,
		"resetRegions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

