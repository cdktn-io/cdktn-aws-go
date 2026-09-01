package awscloudwatchlogs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awscloudwatchlogs/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Audit() DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditPropertyOutputReference
	// Experimental.
	AuditInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty
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
	Deidentify() DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyPropertyOutputReference
	// Experimental.
	DeidentifyInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty
	// Experimental.
	SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty)
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
	PutAudit(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty)
	// Experimental.
	PutDeidentify(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty)
	// Experimental.
	ResetAudit()
	// Experimental.
	ResetDeidentify()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference
type jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) Audit() DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditPropertyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditPropertyOutputReference
	_jsii_.Get(
		j,
		"audit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) AuditInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty
	_jsii_.Get(
		j,
		"auditInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) Deidentify() DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyPropertyOutputReference {
	var returns DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyPropertyOutputReference
	_jsii_.Get(
		j,
		"deidentify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) DeidentifyInput() *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty
	_jsii_.Get(
		j,
		"deidentifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) InternalValue() *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty {
	var returns *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataAwsCloudwatchLogDataProtectionPolicyDocument.OperationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference_Override(d DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-cloudwatch-logs.DataAwsCloudwatchLogDataProtectionPolicyDocument.OperationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference)SetInternalValue(val *DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) PutAudit(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_AuditProperty) {
	if err := d.validatePutAuditParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAudit",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) PutDeidentify(value *DataAwsCloudwatchLogDataProtectionPolicyDocument_DeidentifyProperty) {
	if err := d.validatePutDeidentifyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putDeidentify",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ResetAudit() {
	_jsii_.InvokeVoid(
		d,
		"resetAudit",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ResetDeidentify() {
	_jsii_.InvokeVoid(
		d,
		"resetDeidentify",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsCloudwatchLogDataProtectionPolicyDocument_OperationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

