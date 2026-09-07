package webservicesbudgets

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/jsii"

	"github.com/cdktn-io/cdktn-aws-go/webservicesbudgets/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsBudget_CostTypesPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	IncludeCredit() interface{}
	// Experimental.
	SetIncludeCredit(val interface{})
	// Experimental.
	IncludeCreditInput() interface{}
	// Experimental.
	IncludeDiscount() interface{}
	// Experimental.
	SetIncludeDiscount(val interface{})
	// Experimental.
	IncludeDiscountInput() interface{}
	// Experimental.
	IncludeOtherSubscription() interface{}
	// Experimental.
	SetIncludeOtherSubscription(val interface{})
	// Experimental.
	IncludeOtherSubscriptionInput() interface{}
	// Experimental.
	IncludeRecurring() interface{}
	// Experimental.
	SetIncludeRecurring(val interface{})
	// Experimental.
	IncludeRecurringInput() interface{}
	// Experimental.
	IncludeRefund() interface{}
	// Experimental.
	SetIncludeRefund(val interface{})
	// Experimental.
	IncludeRefundInput() interface{}
	// Experimental.
	IncludeSubscription() interface{}
	// Experimental.
	SetIncludeSubscription(val interface{})
	// Experimental.
	IncludeSubscriptionInput() interface{}
	// Experimental.
	IncludeSupport() interface{}
	// Experimental.
	SetIncludeSupport(val interface{})
	// Experimental.
	IncludeSupportInput() interface{}
	// Experimental.
	IncludeTax() interface{}
	// Experimental.
	SetIncludeTax(val interface{})
	// Experimental.
	IncludeTaxInput() interface{}
	// Experimental.
	IncludeUpfront() interface{}
	// Experimental.
	SetIncludeUpfront(val interface{})
	// Experimental.
	IncludeUpfrontInput() interface{}
	// Experimental.
	InternalValue() *AwsBudget_CostTypesProperty
	// Experimental.
	SetInternalValue(val *AwsBudget_CostTypesProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UseAmortized() interface{}
	// Experimental.
	SetUseAmortized(val interface{})
	// Experimental.
	UseAmortizedInput() interface{}
	// Experimental.
	UseBlended() interface{}
	// Experimental.
	SetUseBlended(val interface{})
	// Experimental.
	UseBlendedInput() interface{}
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
	ResetIncludeCredit()
	// Experimental.
	ResetIncludeDiscount()
	// Experimental.
	ResetIncludeOtherSubscription()
	// Experimental.
	ResetIncludeRecurring()
	// Experimental.
	ResetIncludeRefund()
	// Experimental.
	ResetIncludeSubscription()
	// Experimental.
	ResetIncludeSupport()
	// Experimental.
	ResetIncludeTax()
	// Experimental.
	ResetIncludeUpfront()
	// Experimental.
	ResetUseAmortized()
	// Experimental.
	ResetUseBlended()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsBudget_CostTypesPropertyOutputReference
type jsiiProxy_AwsBudget_CostTypesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeCredit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCredit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeCreditInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeCreditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeDiscount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDiscount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeDiscountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeDiscountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeOtherSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOtherSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeOtherSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeOtherSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeRecurring() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRecurring",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeRecurringInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRecurringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeRefund() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRefund",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeRefundInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeRefundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeSubscription() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubscription",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeSubscriptionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSubscriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeTax() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeTaxInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeTaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeUpfront() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUpfront",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) IncludeUpfrontInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeUpfrontInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) InternalValue() *AwsBudget_CostTypesProperty {
	var returns *AwsBudget_CostTypesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) UseAmortized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAmortized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) UseAmortizedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAmortizedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) UseBlended() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlended",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) UseBlendedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useBlendedInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsBudget_CostTypesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsBudget_CostTypesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsBudget_CostTypesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsBudget_CostTypesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.CostTypesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsBudget_CostTypesPropertyOutputReference_Override(a AwsBudget_CostTypesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-web-services-budgets.AwsBudget.CostTypesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeCredit(val interface{}) {
	if err := j.validateSetIncludeCreditParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeCredit",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeDiscount(val interface{}) {
	if err := j.validateSetIncludeDiscountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeDiscount",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeOtherSubscription(val interface{}) {
	if err := j.validateSetIncludeOtherSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeOtherSubscription",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeRecurring(val interface{}) {
	if err := j.validateSetIncludeRecurringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRecurring",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeRefund(val interface{}) {
	if err := j.validateSetIncludeRefundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeRefund",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeSubscription(val interface{}) {
	if err := j.validateSetIncludeSubscriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubscription",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeSupport(val interface{}) {
	if err := j.validateSetIncludeSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSupport",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeTax(val interface{}) {
	if err := j.validateSetIncludeTaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeTax",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetIncludeUpfront(val interface{}) {
	if err := j.validateSetIncludeUpfrontParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeUpfront",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetInternalValue(val *AwsBudget_CostTypesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetUseAmortized(val interface{}) {
	if err := j.validateSetUseAmortizedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAmortized",
		val,
	)
}

func (j *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference)SetUseBlended(val interface{}) {
	if err := j.validateSetUseBlendedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useBlended",
		val,
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeCredit() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeCredit",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeDiscount() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeDiscount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeOtherSubscription() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeOtherSubscription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeRecurring() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeRecurring",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeRefund() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeRefund",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeSubscription() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeSubscription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeSupport() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeSupport",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeTax() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeTax",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetIncludeUpfront() {
	_jsii_.InvokeVoid(
		a,
		"resetIncludeUpfront",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetUseAmortized() {
	_jsii_.InvokeVoid(
		a,
		"resetUseAmortized",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ResetUseBlended() {
	_jsii_.InvokeVoid(
		a,
		"resetUseBlended",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsBudget_CostTypesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

