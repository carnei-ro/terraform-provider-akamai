provider "akamai" {
  edgerc = "~/.edgerc"
}



data "akamai_appsec_selectable_hostnames" "test" {
    config_id = 43253
    version = 7   
}

output "selectablehostnames" {
  value = data.akamai_appsec_selectable_hostnames.test.hostnames
}
