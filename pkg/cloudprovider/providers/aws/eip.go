package aws

import (
	"fmt"
	aws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"kubeup.com/archon/pkg/cloudprovider"
	"kubeup.com/archon/pkg/cluster"
	"kubeup.com/archon/pkg/util"
)

func (p *awsCloud) EIP() (cloudprovider.EIPInterface, bool) {
	return p, true
}

func (p *awsCloud) EnsureEIP(clusterName string, instance *cluster.Instance) (status *cluster.InstanceStatus, err error) {
	options := cluster.InstanceOptions{}
	eip := EIP{}

	if instance.Labels != nil {
		err = util.MapToStruct(instance.Labels, &options, cluster.AnnotationPrefix)
		if err != nil {
			return
		}
	}

	if options.PreallocatePublicIP && instance.Annotations != nil {
		err = util.MapToStruct(instance.Annotations, &eip, cluster.AnnotationPrefix)
		if err != nil {
			return
		}
	}

	if eip.AllocationID != "" {
		status = &instance.Status
		return
	}

	resp, err := p.ec2.AllocateAddress(&ec2.AllocateAddressInput{
		Domain: aws.String("vpc"),
	})
	if err != nil {
		err = fmt.Errorf("Error allocating EIP: %s", err.Error())
		return
	}

	eip.AllocationID = *resp.AllocationId
	if instance.Annotations == nil {
		instance.Annotations = make(map[string]string)
	}

	err = util.StructToMap(eip, instance.Annotations, AWSAnnotationPrefix)
	if err != nil {
		err = fmt.Errorf("Error allocating EIP: %s", err.Error())
		return
	}

	status.PublicIP = *resp.PublicIp

	return
}

func (p *awsCloud) EnsureEIPDeleted(clusterName string, instance *cluster.Instance) (err error) {
	eip := EIP{}
	err = util.MapToStruct(instance.Annotations, &eip, cluster.AnnotationPrefix)
	if err != nil {
		return
	}

	if eip.AllocationID != "" {
		_, err = p.ec2.ReleaseAddress(&ec2.ReleaseAddressInput{
			AllocationId: aws.String(eip.AllocationID),
		})

		if err != nil {
			err = fmt.Errorf("Error releasing EIP: %s", err.Error())
			return
		}

		eip.AllocationID = ""

		if instance.Annotations == nil {
			instance.Annotations = make(map[string]string)
		}

		err = util.StructToMap(eip, instance.Annotations, AWSAnnotationPrefix)
		if err != nil {
			err = fmt.Errorf("Error allocating EIP: %s", err.Error())
			return
		}
	}

	instance.Status.PublicIP = ""

	return nil
}