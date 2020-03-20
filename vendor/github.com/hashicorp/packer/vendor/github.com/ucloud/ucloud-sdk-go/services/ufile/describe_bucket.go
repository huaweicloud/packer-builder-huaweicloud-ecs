// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufile

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeBucketRequest is request schema for DescribeBucket action
type DescribeBucketRequest struct {
	request.CommonBase

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 待获取Bucket的名称，若不提供，则获取所有Bucket
	BucketName *string `required:"false"`

	// 获取所有Bucket列表的限制数目，默认为20
	Limit *int `required:"false"`

	// 获取所有Bucket列表的偏移数目，默认为0
	Offset *int `required:"false"`
}

// DescribeBucketResponse is response schema for DescribeBucket action
type DescribeBucketResponse struct {
	response.CommonBase

	// Bucket的描述信息 参数见 UFileBucketSet
	DataSet []UFileBucketSet
}

// NewDescribeBucketRequest will create request of DescribeBucket action.
func (c *UFileClient) NewDescribeBucketRequest() *DescribeBucketRequest {
	req := &DescribeBucketRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeBucket - 获取Bucket的描述信息
func (c *UFileClient) DescribeBucket(req *DescribeBucketRequest) (*DescribeBucketResponse, error) {
	var err error
	var res DescribeBucketResponse

	reqCopier := *req

	// In order to ignore the parameters about Region
	reqCopier.Region = ucloud.String("")

	err = c.Client.InvokeAction("DescribeBucket", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
