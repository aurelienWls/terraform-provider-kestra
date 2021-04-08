package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccResourceFlow(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFlow,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(
						"kestra_flow.foo", "namespace", regexp.MustCompile(".*")),
				),
			},
		},
	})
}

const testAccResourceFlow = `
resource "kestra_flow" "foo" {
  namespace = "io.kestra.tests"
  flow_id = "logs"
  content = <<EOT
{
    "taskDefaults": [
        {
            "type": "io.kestra.core.tasks.debugs.Echo",
            "values": {
                "format": "third {{flow.id}}"
            }
        }
    ],
    "tasks": [
        {
            "format": "first {{task.id}}",
            "id": "t1",
            "level": "TRACE",
            "type": "io.kestra.core.tasks.debugs.Echo"
        },
        {
            "format": "second {{task.type}}",
            "id": "t2",
            "level": "WARN",
            "type": "io.kestra.core.tasks.debugs.Echo"
        },
        {
            "format": "third {{flow.id}}",
            "id": "t3",
            "level": "ERROR",
            "type": "io.kestra.core.tasks.debugs.Echo"
        }
    ]
}
EOT
}
`
