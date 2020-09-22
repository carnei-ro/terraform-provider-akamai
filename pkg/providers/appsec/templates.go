package appsec

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/jedib0t/go-pretty/v6/table"
)

type OutputTemplates map[string]*OutputTemplate
type OutputTemplate struct {
	TemplateName   string
	TemplateType   string
	TableTitle     string
	TemplateString string
}

func GetTemplate(ots map[string]*OutputTemplate, key string) (*OutputTemplate, error) {
	if f, ok := ots[key]; ok && f != nil {
		fmt.Printf("%s is in the OutputTemplate >> %+v\n", key, f)
		return f, nil
	} else {
		fmt.Printf("%s is NOT in the OutputTemplate!\n", key)
		return nil, fmt.Errorf("Error not found")
	}
}
func RenderTemplates(ots map[string]*OutputTemplate, key string, str interface{}) (string, error) {
	var ostr, tstr bytes.Buffer
	templ, ok := GetTemplate(ots, key)
	if ok == nil {
		var (
			funcs = template.FuncMap{
				"join":  strings.Join,
				"quote": func(in string) string { return fmt.Sprintf("\"%s\"", in) },
			}
		)
		t := template.Must(template.New("").Funcs(funcs).Parse(templ.TemplateString))
		if err := t.Execute(&tstr, str); err != nil {
			return "", nil
		}
		temptype := templ.TemplateType
		if temptype == "TABULAR" {
			tbl := table.NewWriter()
			tbl.SetOutputMirror(&ostr) //os.Stdout)
			tbl.SetTitle(key)
			headers := templ.TableTitle
			headercolumns := strings.Split(headers, "|")
			trhdr := table.Row{}
			for _, header := range headercolumns {
				trhdr = append(trhdr, header)
			}
			tbl.AppendHeader(trhdr)
			ar := strings.Split(tstr.String(), ",")
			for _, recContent := range ar {
				trc := []table.Row{}
				ac := strings.Split(recContent, "|")
				tr := table.Row{}
				for _, colContent := range ac {
					tr = append(tr, colContent)
				}
				trc = append(trc, tr)
				tbl.AppendRows(trc)
			}
			tbl.Render()
		} else {
			return "\n" + tstr.String(), nil
		}
		return "\n" + ostr.String(), nil
	}
	return "", nil
}

//func (ots *OutputTemplates)
func InitTemplates(otm map[string]*OutputTemplate) {
	otm["selectableHosts"] = &OutputTemplate{TemplateName: "selectableHosts", TableTitle: "Hostname", TemplateType: "TABULAR", TemplateString: "{{range .SelectableHosts}}{{.}},{{end}}"}
	otm["selectedHosts"] = &OutputTemplate{TemplateName: "selectedHosts", TableTitle: "Hostnames", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .SelectedHosts}}{{if $index}},{{end}}{{.}}{{end}}"}
	otm["selectedHosts.tf"] = &OutputTemplate{TemplateName: "selectedHosts.tf", TableTitle: "Hostname", TemplateType: "TERRAFORM", TemplateString: "\nresource \"akamai_appsec_selected_hostnames\" \"appsecselectedhostnames\" { \n config_id = {{.ConfigID}}\n version = {{.Version}}\n hostnames = [{{  range $index, $element := .SelectedHosts }}{{if $index}},{{end}}{{quote .}}{{end}}] \n }"}
	otm["ratePolicies"] = &OutputTemplate{TemplateName: "ratePolicies", TableTitle: "ID|PolicyID", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .RatePolicies}}{{if $index}},{{end}}{{.ID}}|{{.Name}}{{end}}"}
	otm["matchTargets"] = &OutputTemplate{TemplateName: "matchTargets", TableTitle: "ID|PolicyID", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .MatchTargets.WebsiteTargets}}{{if $index}},{{end}}{{.ID}}|{{.SecurityPolicy.PolicyID}}{{end}}"}
	otm["reputationProfiles"] = &OutputTemplate{TemplateName: "reputationProfiles", TableTitle: "ID|Name(Title)", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .ReputationProfiles}}{{if $index}},{{end}}{{.ID}}|{{.Name}}{{end}}"}
	otm["customRules"] = &OutputTemplate{TemplateName: "customRules", TableTitle: "ID|Name", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .ReputationProfiles}}{{if $index}},{{end}}{{.ID}}|{{.Name}}{{end}}"}
	otm["rulesets"] = &OutputTemplate{TemplateName: "rulesets", TableTitle: "ID|Name(Title)", TemplateType: "TABULAR", TemplateString: "{{range .Rulesets}}{{range $index, $element := .Rules}}{{if $index}},{{end}}{{.ID}}| {{.Title}}{{end}}{{end}}"}
	otm["securityPolicies"] = &OutputTemplate{TemplateName: "securityPolicies", TableTitle: "ID|Name", TemplateType: "TABULAR", TemplateString: "{{range $index, $element := .SecurityPolicies}}{{if $index}},{{end}}{{.ID}}|{{.Name}}{{end}}"}
	otm["ruleActions"] = &OutputTemplate{TemplateName: "ruleActions", TableTitle: "ID|Action", TemplateType: "TABULAR", TemplateString: "{{range .SecurityPolicies}}{{range $index, $element := .WebApplicationFirewall.RuleActions}}{{if $index}},{{end}}{{.ID}}| {{.Action}}{{end}}{{end}}"}
}