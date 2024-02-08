package fragment

import (
	"github.com/sascha-dibbern/Hugiki/hiuri"
)

const MenuItem_Start string = "Start"
const MenuItem_Status string = "Status"
const MenuItem_Navigation string = "Files"
const MenuItem_Search string = "Search"
const MenuItem_Git string = "Git"
const MenuItem_Configuration string = "Configuration"
const MenuItem_Help string = "Help"

var MenuOrder = [...]string{MenuItem_Status, MenuItem_Navigation, MenuItem_Search, MenuItem_Git, MenuItem_Configuration, MenuItem_Help}

type MenuItemUri map[string]string

var MenuItemUris = MenuItemUri{
	//MenuItem_Start:         hiuri.UriAction_StartMode,
	MenuItem_Status:        hiuri.UriAction_StatusMode,
	MenuItem_Navigation:    hiuri.UriAction_NavigationMode,
	MenuItem_Search:        hiuri.UriAction_SearchMode,
	MenuItem_Git:           hiuri.UriAction_GitMode,
	MenuItem_Configuration: hiuri.UriAction_ConfigMode,
	MenuItem_Help:          hiuri.UriAction_HelpMode,
}

type MenuItem map[string]string

type MenuItems [6]MenuItem

func BuildMenuState(selected string) MenuItems {
	var menu MenuItems
	for itemindex, itemname := range MenuOrder {
		itemuri := MenuItemUris[itemname]
		if selected == itemname {
			itemuri = ""
		}
		item := MenuItem{
			"Name": itemname,
			"Uri":  itemuri,
		}
		menu[itemindex] = item
	}
	return menu
}

var MenuTemplate = `
{{- define "activemenuitem" -}}
<td {{ if (ne .Uri "") }} hx-get="{{- .Uri -}}" hx-target="#himode" hx-trigger="click"{{ end }}><div class="hi-menuitem">{{- .Name -}}</div></td>
{{- end -}}
{{- define "menu" -}}
<div>
<table class="hi-menu" id="hi-appmenu"><tr>
{{- range . -}}
{{ template "activemenuitem" . }}
{{- end -}}
</tr></table>
</div>
{{- end -}}`
