package huaweicloud

import (
	"context"
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/huaweicloud/golangsdk/openstack/imageservice/v2/members"
)

type stepUpdateImageMembers struct{}

func (s *stepUpdateImageMembers) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	imageId := state.Get("image").(string)
	ui := state.Get("ui").(packer.Ui)
	config := state.Get("config").(*Config)

	if len(config.ImageMembers) == 0 {
		return multistep.ActionContinue
	}

	imageClient, err := config.imageV2Client()
	if err != nil {
		err = fmt.Errorf("Error initializing image service client: %s", err)
		state.Put("error", err)
		return multistep.ActionHalt
	}

	for _, member := range config.ImageMembers {
		ui.Say(fmt.Sprintf("Adding member '%s' to image %s", member, imageId))
		r := members.Create(imageClient, imageId, member)
		if _, err = r.Extract(); err != nil {
			err = fmt.Errorf("Error adding member to image: %s", err)
			state.Put("error", err)
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *stepUpdateImageMembers) Cleanup(multistep.StateBag) {
	// No cleanup...
}
