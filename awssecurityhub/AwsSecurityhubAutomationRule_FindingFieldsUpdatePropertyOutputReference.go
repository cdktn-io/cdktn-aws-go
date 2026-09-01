package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference interface {
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
	Confidence() *float64
	// Experimental.
	SetConfidence(val *float64)
	// Experimental.
	ConfidenceInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Criticality() *float64
	// Experimental.
	SetCriticality(val *float64)
	// Experimental.
	CriticalityInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	Note() AwsSecurityhubAutomationRule_NotePropertyList
	// Experimental.
	NoteInput() interface{}
	// Experimental.
	RelatedFindings() AwsSecurityhubAutomationRule_RelatedFindingsPropertyList
	// Experimental.
	RelatedFindingsInput() interface{}
	// Experimental.
	Severity() AwsSecurityhubAutomationRule_SeverityPropertyList
	// Experimental.
	SeverityInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Types() *[]*string
	// Experimental.
	SetTypes(val *[]*string)
	// Experimental.
	TypesInput() *[]*string
	// Experimental.
	UserDefinedFields() *map[string]*string
	// Experimental.
	SetUserDefinedFields(val *map[string]*string)
	// Experimental.
	UserDefinedFieldsInput() *map[string]*string
	// Experimental.
	VerificationState() *string
	// Experimental.
	SetVerificationState(val *string)
	// Experimental.
	VerificationStateInput() *string
	// Experimental.
	Workflow() AwsSecurityhubAutomationRule_WorkflowPropertyList
	// Experimental.
	WorkflowInput() interface{}
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
	PutNote(value interface{})
	// Experimental.
	PutRelatedFindings(value interface{})
	// Experimental.
	PutSeverity(value interface{})
	// Experimental.
	PutWorkflow(value interface{})
	// Experimental.
	ResetConfidence()
	// Experimental.
	ResetCriticality()
	// Experimental.
	ResetNote()
	// Experimental.
	ResetRelatedFindings()
	// Experimental.
	ResetSeverity()
	// Experimental.
	ResetTypes()
	// Experimental.
	ResetUserDefinedFields()
	// Experimental.
	ResetVerificationState()
	// Experimental.
	ResetWorkflow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference
type jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Confidence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ConfidenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Criticality() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) CriticalityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Note() AwsSecurityhubAutomationRule_NotePropertyList {
	var returns AwsSecurityhubAutomationRule_NotePropertyList
	_jsii_.Get(
		j,
		"note",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) NoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) RelatedFindings() AwsSecurityhubAutomationRule_RelatedFindingsPropertyList {
	var returns AwsSecurityhubAutomationRule_RelatedFindingsPropertyList
	_jsii_.Get(
		j,
		"relatedFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) RelatedFindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Severity() AwsSecurityhubAutomationRule_SeverityPropertyList {
	var returns AwsSecurityhubAutomationRule_SeverityPropertyList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) UserDefinedFields() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) UserDefinedFieldsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) VerificationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) VerificationStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Workflow() AwsSecurityhubAutomationRule_WorkflowPropertyList {
	var returns AwsSecurityhubAutomationRule_WorkflowPropertyList
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) WorkflowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubAutomationRule.FindingFieldsUpdatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference_Override(a AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.AwsSecurityhubAutomationRule.FindingFieldsUpdatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetConfidence(val *float64) {
	if err := j.validateSetConfidenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetCriticality(val *float64) {
	if err := j.validateSetCriticalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticality",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTypes(val *[]*string) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetUserDefinedFields(val *map[string]*string) {
	if err := j.validateSetUserDefinedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDefinedFields",
		val,
	)
}

func (j *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetVerificationState(val *string) {
	if err := j.validateSetVerificationStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationState",
		val,
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutNote(value interface{}) {
	if err := a.validatePutNoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNote",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutRelatedFindings(value interface{}) {
	if err := a.validatePutRelatedFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRelatedFindings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutSeverity(value interface{}) {
	if err := a.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSeverity",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutWorkflow(value interface{}) {
	if err := a.validatePutWorkflowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putWorkflow",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		a,
		"resetConfidence",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		a,
		"resetCriticality",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetNote() {
	_jsii_.InvokeVoid(
		a,
		"resetNote",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetRelatedFindings() {
	_jsii_.InvokeVoid(
		a,
		"resetRelatedFindings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		a,
		"resetSeverity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetTypes() {
	_jsii_.InvokeVoid(
		a,
		"resetTypes",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		a,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		a,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		a,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSecurityhubAutomationRule_FindingFieldsUpdatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

