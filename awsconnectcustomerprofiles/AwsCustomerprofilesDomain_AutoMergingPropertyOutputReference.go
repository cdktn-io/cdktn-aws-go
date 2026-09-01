package awsconnectcustomerprofiles

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsconnectcustomerprofiles/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference interface {
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
	ConflictResolution() AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference
	// Experimental.
	ConflictResolutionInput() *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty
	// Experimental.
	Consolidation() AwsCustomerprofilesDomain_ConsolidationPropertyOutputReference
	// Experimental.
	ConsolidationInput() *AwsCustomerprofilesDomain_ConsolidationProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCustomerprofilesDomain_AutoMergingProperty
	// Experimental.
	SetInternalValue(val *AwsCustomerprofilesDomain_AutoMergingProperty)
	// Experimental.
	MinAllowedConfidenceScoreForMerging() *float64
	// Experimental.
	SetMinAllowedConfidenceScoreForMerging(val *float64)
	// Experimental.
	MinAllowedConfidenceScoreForMergingInput() *float64
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
	PutConflictResolution(value *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty)
	// Experimental.
	PutConsolidation(value *AwsCustomerprofilesDomain_ConsolidationProperty)
	// Experimental.
	ResetConflictResolution()
	// Experimental.
	ResetConsolidation()
	// Experimental.
	ResetMinAllowedConfidenceScoreForMerging()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference
type jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ConflictResolution() AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference {
	var returns AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionPropertyOutputReference
	_jsii_.Get(
		j,
		"conflictResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ConflictResolutionInput() *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty {
	var returns *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty
	_jsii_.Get(
		j,
		"conflictResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) Consolidation() AwsCustomerprofilesDomain_ConsolidationPropertyOutputReference {
	var returns AwsCustomerprofilesDomain_ConsolidationPropertyOutputReference
	_jsii_.Get(
		j,
		"consolidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ConsolidationInput() *AwsCustomerprofilesDomain_ConsolidationProperty {
	var returns *AwsCustomerprofilesDomain_ConsolidationProperty
	_jsii_.Get(
		j,
		"consolidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) InternalValue() *AwsCustomerprofilesDomain_AutoMergingProperty {
	var returns *AwsCustomerprofilesDomain_AutoMergingProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) MinAllowedConfidenceScoreForMerging() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMerging",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) MinAllowedConfidenceScoreForMergingInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minAllowedConfidenceScoreForMergingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCustomerprofilesDomain_AutoMergingPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCustomerprofilesDomain_AutoMergingPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesDomain.AutoMergingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCustomerprofilesDomain_AutoMergingPropertyOutputReference_Override(a AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-connect-customer-profiles.AwsCustomerprofilesDomain.AutoMergingPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetInternalValue(val *AwsCustomerprofilesDomain_AutoMergingProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetMinAllowedConfidenceScoreForMerging(val *float64) {
	if err := j.validateSetMinAllowedConfidenceScoreForMergingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minAllowedConfidenceScoreForMerging",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) PutConflictResolution(value *AwsCustomerprofilesDomain_MatchingAutoMergingConflictResolutionProperty) {
	if err := a.validatePutConflictResolutionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConflictResolution",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) PutConsolidation(value *AwsCustomerprofilesDomain_ConsolidationProperty) {
	if err := a.validatePutConsolidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConsolidation",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ResetConflictResolution() {
	_jsii_.InvokeVoid(
		a,
		"resetConflictResolution",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ResetConsolidation() {
	_jsii_.InvokeVoid(
		a,
		"resetConsolidation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ResetMinAllowedConfidenceScoreForMerging() {
	_jsii_.InvokeVoid(
		a,
		"resetMinAllowedConfidenceScoreForMerging",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCustomerprofilesDomain_AutoMergingPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

