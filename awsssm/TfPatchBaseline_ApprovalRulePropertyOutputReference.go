package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsssm/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsssm/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfPatchBaseline_ApprovalRulePropertyOutputReference interface {
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
	PatchFilter() TfPatchBaseline_PatchFilterPropertyList
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

// The jsii proxy struct for TfPatchBaseline_ApprovalRulePropertyOutputReference
type jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ApproveAfterDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ApproveAfterDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"approveAfterDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ApproveUntilDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ApproveUntilDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"approveUntilDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ComplianceLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ComplianceLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"complianceLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) EnableNonSecurity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) EnableNonSecurityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableNonSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) PatchFilter() TfPatchBaseline_PatchFilterPropertyList {
	var returns TfPatchBaseline_PatchFilterPropertyList
	_jsii_.Get(
		j,
		"patchFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) PatchFilterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"patchFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfPatchBaseline_ApprovalRulePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfPatchBaseline_ApprovalRulePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfPatchBaseline_ApprovalRulePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ssm.TfPatchBaseline.ApprovalRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfPatchBaseline_ApprovalRulePropertyOutputReference_Override(t TfPatchBaseline_ApprovalRulePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ssm.TfPatchBaseline.ApprovalRulePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetApproveAfterDays(val *float64) {
	if err := j.validateSetApproveAfterDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveAfterDays",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetApproveUntilDate(val *string) {
	if err := j.validateSetApproveUntilDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"approveUntilDate",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetComplianceLevel(val *string) {
	if err := j.validateSetComplianceLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complianceLevel",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetEnableNonSecurity(val interface{}) {
	if err := j.validateSetEnableNonSecurityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableNonSecurity",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) PutPatchFilter(value interface{}) {
	if err := t.validatePutPatchFilterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPatchFilter",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ResetApproveAfterDays() {
	_jsii_.InvokeVoid(
		t,
		"resetApproveAfterDays",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ResetApproveUntilDate() {
	_jsii_.InvokeVoid(
		t,
		"resetApproveUntilDate",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ResetComplianceLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetComplianceLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ResetEnableNonSecurity() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableNonSecurity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfPatchBaseline_ApprovalRulePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

