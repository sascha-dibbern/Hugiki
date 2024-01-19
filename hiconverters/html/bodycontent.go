package html

// ...<body...>xyz</body>... => xyz
func ExtractBodycontent(htmlInput string) string {
	clean_in_start := upWithBodyStartTagRegexp.ReplaceAllString(htmlInput, "")
	clean_also_after_end := fromWithBodyEndTagRegexp.ReplaceAllString(clean_in_start, "")
	return clean_also_after_end
}
