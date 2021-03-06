package appsec

import (
	"encoding/json"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/v2/pkg/appsec"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/mock"
)

func TestAccAkamaiAdvancedSettingsLogging_res_basic(t *testing.T) {
	t.Run("match by AdvancedSettingsLogging ID", func(t *testing.T) {
		client := &mockappsec{}

		cu := appsec.UpdateAdvancedSettingsLoggingResponse{}
		expectJSU := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsLogging/AdvancedSettingsLogging.json"))
		json.Unmarshal([]byte(expectJSU), &cu)

		cr := appsec.GetAdvancedSettingsLoggingResponse{}
		expectJS := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsLogging/AdvancedSettingsLogging.json"))
		json.Unmarshal([]byte(expectJS), &cr)

		crd := appsec.RemoveAdvancedSettingsLoggingResponse{}
		expectJSD := compactJSON(loadFixtureBytes("testdata/TestResAdvancedSettingsLogging/AdvancedSettingsLogging.json"))
		json.Unmarshal([]byte(expectJSD), &crd)

		client.On("GetAdvancedSettingsLogging",
			mock.Anything, // ctx is irrelevant for this test
			appsec.GetAdvancedSettingsLoggingRequest{ConfigID: 43253, Version: 7},
		).Return(&cr, nil)

		client.On("UpdateAdvancedSettingsLogging",
			mock.Anything, // ctx is irrelevant for this test
			appsec.UpdateAdvancedSettingsLoggingRequest{ConfigID: 43253, Version: 7, PolicyID: "", JsonPayloadRaw: json.RawMessage{0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x3a, 0x20, 0x74, 0x72, 0x75, 0x65, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x61, 0x6c, 0x6c, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x63, 0x73, 0x64, 0x61, 0x73, 0x64, 0x61, 0x64, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x22, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x7b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x6f, 0x6e, 0x6c, 0x79, 0x22, 0x2c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x22, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x22, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x5d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x7d, 0xa, 0x7d, 0xa}},
		).Return(&cu, nil)

		client.On("RemoveAdvancedSettingsLogging",
			mock.Anything, // ctx is irrelevant for this test
			appsec.RemoveAdvancedSettingsLoggingRequest{ConfigID: 43253, Version: 7, PolicyID: ""},
		).Return(&crd, nil)

		useClient(client, func() {
			resource.Test(t, resource.TestCase{
				IsUnitTest: true,
				Providers:  testAccProviders,
				Steps: []resource.TestStep{
					{
						Config: loadFixtureString("testdata/TestResAdvancedSettingsLogging/match_by_id.tf"),
						Check: resource.ComposeAggregateTestCheckFunc(
							resource.TestCheckResourceAttr("akamai_appsec_advanced_settings_logging.test", "id", "43253:7:"),
						),
						ExpectNonEmptyPlan: true,
					},
				},
			})
		})

		client.AssertExpectations(t)
	})

}
