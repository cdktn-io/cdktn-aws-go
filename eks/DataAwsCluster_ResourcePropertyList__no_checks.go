//go:build no_runtime_type_checking

package eks

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsCluster_ResourcePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsCluster_ResourcePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

