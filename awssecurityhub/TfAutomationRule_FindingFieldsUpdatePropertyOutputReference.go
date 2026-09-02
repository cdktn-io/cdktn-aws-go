package awssecurityhub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssecurityhub/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssecurityhub/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAutomationRule_FindingFieldsUpdatePropertyOutputReference interface {
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
	Note() TfAutomationRule_NotePropertyList
	// Experimental.
	NoteInput() interface{}
	// Experimental.
	RelatedFindings() TfAutomationRule_RelatedFindingsPropertyList
	// Experimental.
	RelatedFindingsInput() interface{}
	// Experimental.
	Severity() TfAutomationRule_SeverityPropertyList
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
	Workflow() TfAutomationRule_WorkflowPropertyList
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

// The jsii proxy struct for TfAutomationRule_FindingFieldsUpdatePropertyOutputReference
type jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Confidence() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ConfidenceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"confidenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Criticality() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) CriticalityInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"criticalityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Note() TfAutomationRule_NotePropertyList {
	var returns TfAutomationRule_NotePropertyList
	_jsii_.Get(
		j,
		"note",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) NoteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) RelatedFindings() TfAutomationRule_RelatedFindingsPropertyList {
	var returns TfAutomationRule_RelatedFindingsPropertyList
	_jsii_.Get(
		j,
		"relatedFindings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) RelatedFindingsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"relatedFindingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Severity() TfAutomationRule_SeverityPropertyList {
	var returns TfAutomationRule_SeverityPropertyList
	_jsii_.Get(
		j,
		"severity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) SeverityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"severityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Types() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"types",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) TypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"typesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) UserDefinedFields() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) UserDefinedFieldsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"userDefinedFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) VerificationState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) VerificationStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"verificationStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Workflow() TfAutomationRule_WorkflowPropertyList {
	var returns TfAutomationRule_WorkflowPropertyList
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) WorkflowInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfAutomationRule_FindingFieldsUpdatePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfAutomationRule_FindingFieldsUpdatePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfAutomationRule_FindingFieldsUpdatePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.FindingFieldsUpdatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfAutomationRule_FindingFieldsUpdatePropertyOutputReference_Override(t TfAutomationRule_FindingFieldsUpdatePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-security-hub.TfAutomationRule.FindingFieldsUpdatePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetConfidence(val *float64) {
	if err := j.validateSetConfidenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidence",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetCriticality(val *float64) {
	if err := j.validateSetCriticalityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"criticality",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetTypes(val *[]*string) {
	if err := j.validateSetTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"types",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetUserDefinedFields(val *map[string]*string) {
	if err := j.validateSetUserDefinedFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userDefinedFields",
		val,
	)
}

func (j *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference)SetVerificationState(val *string) {
	if err := j.validateSetVerificationStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verificationState",
		val,
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutNote(value interface{}) {
	if err := t.validatePutNoteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNote",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutRelatedFindings(value interface{}) {
	if err := t.validatePutRelatedFindingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRelatedFindings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutSeverity(value interface{}) {
	if err := t.validatePutSeverityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSeverity",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) PutWorkflow(value interface{}) {
	if err := t.validatePutWorkflowParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putWorkflow",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetConfidence() {
	_jsii_.InvokeVoid(
		t,
		"resetConfidence",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetCriticality() {
	_jsii_.InvokeVoid(
		t,
		"resetCriticality",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetNote() {
	_jsii_.InvokeVoid(
		t,
		"resetNote",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetRelatedFindings() {
	_jsii_.InvokeVoid(
		t,
		"resetRelatedFindings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetSeverity() {
	_jsii_.InvokeVoid(
		t,
		"resetSeverity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetTypes() {
	_jsii_.InvokeVoid(
		t,
		"resetTypes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetUserDefinedFields() {
	_jsii_.InvokeVoid(
		t,
		"resetUserDefinedFields",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetVerificationState() {
	_jsii_.InvokeVoid(
		t,
		"resetVerificationState",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		t,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfAutomationRule_FindingFieldsUpdatePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

