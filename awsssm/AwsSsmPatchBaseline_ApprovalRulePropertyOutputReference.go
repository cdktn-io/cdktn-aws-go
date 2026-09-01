package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApproveAfterDays() *float64
	// Experimental.
	SetApproveAfterDays(val *float64)
	// Experimental.
	ApproveAfterDaysInput() *float64
	// Experimental.
	ApproveUntilDate() *string
	// Experimental.
	SetApproveUntilDate(val *string)
	// Experimental.
	ApproveUntilDateInput() *string
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
	ComplianceLevel() *string
	// Experimental.
	SetComplianceLevel(val *string)
	// Experimental.
	ComplianceLevelInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EnableNonSecurity() interface{}
	// Experimental.
	SetEnableNonSecurity(val interface{})
	// Experimental.
	EnableNonSecurityInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PatchFilter() AwsSsmPatchBaseline_PatchFilterPropertyList
	// Experimental.
	PatchFilterInput() interface{}
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
	PutPatchFilter(value interface{})
	// Experimental.
	ResetApproveAfterDays()
	// Experimental.
	ResetApproveUntilDate()
	// Experimental.
	ResetComplianceLevel()
	// Experimental.
	ResetEnableNonSecurity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference
type jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ApproveAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ApproveAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ApproveUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ApproveUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ComplianceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) EnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) EnableNonSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) PatchFilter() AwsSsmPatchBaseline_PatchFilterPropertyList {
	var returns AwsSsmPatchBaseline_PatchFilterPropertyList
	_jsii_.Get(
		j,
		"patchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) PatchFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patchFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSsmPatchBaseline_ApprovalRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSsmPatchBaseline_ApprovalRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmPatchBaseline.ApprovalRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSsmPatchBaseline_ApprovalRulePropertyOutputReference_Override(a AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.AwsSsmPatchBaseline.ApprovalRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetApproveAfterDays(val *float64) {
	if err := j.validateSetApproveAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveAfterDays",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetApproveUntilDate(val *string) {
	if err := j.validateSetApproveUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveUntilDate",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetComplianceLevel(val *string) {
	if err := j.validateSetComplianceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceLevel",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetEnableNonSecurity(val interface{}) {
	if err := j.validateSetEnableNonSecurityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) PutPatchFilter(value interface{}) {
	if err := a.validatePutPatchFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPatchFilter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ResetApproveAfterDays() {
	_jsii_.InvokeVoid(
		a,
		"resetApproveAfterDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ResetApproveUntilDate() {
	_jsii_.InvokeVoid(
		a,
		"resetApproveUntilDate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ResetComplianceLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetComplianceLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ResetEnableNonSecurity() {
	_jsii_.InvokeVoid(
		a,
		"resetEnableNonSecurity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSsmPatchBaseline_ApprovalRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

