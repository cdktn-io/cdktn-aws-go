//go:build no_runtime_type_checking

package awsbackup

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfPlan_RulePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfPlan_RulePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfPlan_RulePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfPlan_RulePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfPlan_RulePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfPlan_RulePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfPlan_RulePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

