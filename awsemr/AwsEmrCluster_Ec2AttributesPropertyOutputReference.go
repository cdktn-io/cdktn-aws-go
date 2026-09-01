package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrCluster_Ec2AttributesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdditionalMasterSecurityGroups() *string
	// Experimental.
	SetAdditionalMasterSecurityGroups(val *string)
	// Experimental.
	AdditionalMasterSecurityGroupsInput() *string
	// Experimental.
	AdditionalSlaveSecurityGroups() *string
	// Experimental.
	SetAdditionalSlaveSecurityGroups(val *string)
	// Experimental.
	AdditionalSlaveSecurityGroupsInput() *string
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
	EmrManagedMasterSecurityGroup() *string
	// Experimental.
	SetEmrManagedMasterSecurityGroup(val *string)
	// Experimental.
	EmrManagedMasterSecurityGroupInput() *string
	// Experimental.
	EmrManagedSlaveSecurityGroup() *string
	// Experimental.
	SetEmrManagedSlaveSecurityGroup(val *string)
	// Experimental.
	EmrManagedSlaveSecurityGroupInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InstanceProfile() *string
	// Experimental.
	SetInstanceProfile(val *string)
	// Experimental.
	InstanceProfileInput() *string
	// Experimental.
	InternalValue() *AwsEmrCluster_Ec2AttributesProperty
	// Experimental.
	SetInternalValue(val *AwsEmrCluster_Ec2AttributesProperty)
	// Experimental.
	KeyName() *string
	// Experimental.
	SetKeyName(val *string)
	// Experimental.
	KeyNameInput() *string
	// Experimental.
	ServiceAccessSecurityGroup() *string
	// Experimental.
	SetServiceAccessSecurityGroup(val *string)
	// Experimental.
	ServiceAccessSecurityGroupInput() *string
	// Experimental.
	SubnetId() *string
	// Experimental.
	SetSubnetId(val *string)
	// Experimental.
	SubnetIdInput() *string
	// Experimental.
	SubnetIds() *[]*string
	// Experimental.
	SetSubnetIds(val *[]*string)
	// Experimental.
	SubnetIdsInput() *[]*string
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
	ResetAdditionalMasterSecurityGroups()
	// Experimental.
	ResetAdditionalSlaveSecurityGroups()
	// Experimental.
	ResetEmrManagedMasterSecurityGroup()
	// Experimental.
	ResetEmrManagedSlaveSecurityGroup()
	// Experimental.
	ResetKeyName()
	// Experimental.
	ResetServiceAccessSecurityGroup()
	// Experimental.
	ResetSubnetId()
	// Experimental.
	ResetSubnetIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEmrCluster_Ec2AttributesPropertyOutputReference
type jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) AdditionalMasterSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) AdditionalMasterSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalMasterSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) AdditionalSlaveSecurityGroups() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) AdditionalSlaveSecurityGroupsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"additionalSlaveSecurityGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) EmrManagedMasterSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) EmrManagedMasterSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedMasterSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) EmrManagedSlaveSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) EmrManagedSlaveSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emrManagedSlaveSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) InstanceProfile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) InstanceProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceProfileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) InternalValue() *AwsEmrCluster_Ec2AttributesProperty {
	var returns *AwsEmrCluster_Ec2AttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) KeyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) KeyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ServiceAccessSecurityGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ServiceAccessSecurityGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessSecurityGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) SubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) SubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrCluster_Ec2AttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEmrCluster_Ec2AttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrCluster_Ec2AttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.Ec2AttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrCluster_Ec2AttributesPropertyOutputReference_Override(a AwsEmrCluster_Ec2AttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.Ec2AttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetAdditionalMasterSecurityGroups(val *string) {
	if err := j.validateSetAdditionalMasterSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalMasterSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetAdditionalSlaveSecurityGroups(val *string) {
	if err := j.validateSetAdditionalSlaveSecurityGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalSlaveSecurityGroups",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetEmrManagedMasterSecurityGroup(val *string) {
	if err := j.validateSetEmrManagedMasterSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emrManagedMasterSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetEmrManagedSlaveSecurityGroup(val *string) {
	if err := j.validateSetEmrManagedSlaveSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emrManagedSlaveSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetInstanceProfile(val *string) {
	if err := j.validateSetInstanceProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceProfile",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetInternalValue(val *AwsEmrCluster_Ec2AttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetKeyName(val *string) {
	if err := j.validateSetKeyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyName",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetServiceAccessSecurityGroup(val *string) {
	if err := j.validateSetServiceAccessSecurityGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccessSecurityGroup",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetSubnetIds(val *[]*string) {
	if err := j.validateSetSubnetIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetAdditionalMasterSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalMasterSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetAdditionalSlaveSecurityGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalSlaveSecurityGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetEmrManagedMasterSecurityGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrManagedMasterSecurityGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetEmrManagedSlaveSecurityGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetEmrManagedSlaveSecurityGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetKeyName() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetServiceAccessSecurityGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetServiceAccessSecurityGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ResetSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_Ec2AttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

