package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsservicecatalog/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsservicecatalog/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference interface {
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
	InternalValue() *AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesProperty
	// Experimental.
	SetInternalValue(val *AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesProperty)
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

// The jsii proxy struct for AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference
type jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Accounts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accounts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) AccountsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"accountsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureToleranceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureToleranceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureToleranceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureTolerancePercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) FailureTolerancePercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"failureTolerancePercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InternalValue() *AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesProperty {
	var returns *AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) MaxConcurrencyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConcurrencyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Regions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) RegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"regionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-service-catalog.AwsServicecatalogProvisionedProduct.StackSetProvisioningPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference_Override(a AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-service-catalog.AwsServicecatalogProvisionedProduct.StackSetProvisioningPreferencesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetAccounts(val *[]*string) {
	if err := j.validateSetAccountsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accounts",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetFailureToleranceCount(val *float64) {
	if err := j.validateSetFailureToleranceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureToleranceCount",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetFailureTolerancePercentage(val *float64) {
	if err := j.validateSetFailureTolerancePercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"failureTolerancePercentage",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetInternalValue(val *AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetMaxConcurrencyCount(val *float64) {
	if err := j.validateSetMaxConcurrencyCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyCount",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetMaxConcurrencyPercentage(val *float64) {
	if err := j.validateSetMaxConcurrencyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConcurrencyPercentage",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetRegions(val *[]*string) {
	if err := j.validateSetRegionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"regions",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetAccounts() {
	_jsii_.InvokeVoid(
		a,
		"resetAccounts",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetFailureToleranceCount() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureToleranceCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetFailureTolerancePercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetFailureTolerancePercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetMaxConcurrencyCount() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrencyCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetMaxConcurrencyPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxConcurrencyPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ResetRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsServicecatalogProvisionedProduct_StackSetProvisioningPreferencesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

