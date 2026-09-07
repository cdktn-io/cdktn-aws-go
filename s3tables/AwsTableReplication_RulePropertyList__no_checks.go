//go:build no_runtime_type_checking

package s3tables

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsTableReplication_RulePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsTableReplication_RulePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsTableReplication_RulePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsTableReplication_RulePropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_AwsTableReplication_RulePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsTableReplication_RulePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsTableReplication_RulePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsTableReplication_RulePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

