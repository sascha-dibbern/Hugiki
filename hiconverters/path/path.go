package path

import (
	"github.com/sascha-dibbern/Hugiki/hiconfig"
	"github.com/sascha-dibbern/Hugiki/hiconverters"
)

/**/

type hugoToHugikiUriRule struct {
	definition hiconverters.TextConversionRuleDefinition
}

func HugoToHugikiUriRule(matching string, replacement string) hugoToHugikiUriRule {
	// Todo: add check for URI
	return hugoToHugikiUriRule{
		hiconverters.NewTextConversionRuleDefinition(matching, replacement),
	}
}

func (rule hugoToHugikiUriRule) ConvertAll(hugoinput string) string {
	return rule.definition.Matchingregexp.ReplaceAllString(hugoinput, rule.definition.Replacement)
}

/**/

type hugoToHugikiUrlRule struct {
	definition hiconverters.TextConversionRuleDefinition
}

func HugoToHugikiUrlRule(matching_uri string, replacement_uri string) hugoToHugikiUrlRule {
	matching_url := hiconfig.AppConfig().BackendBaseUrl() + matching_uri
	replacement_url := replacement_uri
	return hugoToHugikiUrlRule{
		hiconverters.NewTextConversionRuleDefinition(matching_url, replacement_url),
	}
}

func (rule hugoToHugikiUrlRule) ConvertAll(hugoinput string) string {
	return rule.definition.Matchingregexp.ReplaceAllString(hugoinput, rule.definition.Replacement)
}

/**/

type hugikiToHugoUriRule struct {
	definition hiconverters.TextConversionRuleDefinition
}

func HugikiToHugoUriRule(matching string, replacement string) hugikiToHugoUriRule {
	// Todo: add check for URI
	return hugikiToHugoUriRule{
		hiconverters.NewTextConversionRuleDefinition(matching, replacement),
	}
}

func (rule hugikiToHugoUriRule) ConvertAll(hugoinput string) string {
	return rule.definition.Matchingregexp.ReplaceAllString(hugoinput, rule.definition.Replacement)
}

/**/

type hugikiUriToHugoUrlRule struct {
	definition hiconverters.TextConversionRuleDefinition
}

func HugikiUriToHugoUrlRule(hugiki_uri string, hugo_replacement_uri string) hugikiUriToHugoUrlRule {
	replacement_url := hiconfig.AppConfig().BackendBaseUrl() + "/" + hugo_replacement_uri
	return hugikiUriToHugoUrlRule{
		hiconverters.NewTextConversionRuleDefinition(hugiki_uri, replacement_url),
	}
}

func (rule hugikiUriToHugoUrlRule) ConvertAll(hugoinput string) string {
	return rule.definition.Matchingregexp.ReplaceAllString(hugoinput, rule.definition.Replacement)
}

/**/

type hugikiUriToHugoContentUrlRule struct {
	prerule    hugikiUriToHugoUrlRule
	definition hiconverters.TextConversionRuleDefinition
}

func HugikiUriToHugoContentUrlRule(hugiki_uri string, hugo_replacement_uri string) hugikiUriToHugoContentUrlRule {
	prerule := HugikiUriToHugoUrlRule(hugiki_uri, hugo_replacement_uri)
	return hugikiUriToHugoContentUrlRule{
		prerule:    prerule,
		definition: hiconverters.NewTextConversionRuleDefinition("\\.md", "/"),
	}
}

func (rule hugikiUriToHugoContentUrlRule) ConvertAll(hugoinput string) string {
	uriprefixreplaced := rule.prerule.ConvertAll(hugoinput)
	result := rule.definition.Matchingregexp.ReplaceAllString(uriprefixreplaced, rule.definition.Replacement)
	return result
}
