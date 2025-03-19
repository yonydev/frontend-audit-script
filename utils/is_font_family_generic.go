package utils

func IsGenericFontFamily(fontName string) bool {
	genericFontFamilies := map[string]struct{}{
		"serif":         {},
		"sans-serif":    {},
		"monospace":     {},
		"cursive":       {},
		"fantasy":       {},
		"system-ui":     {},
		"ui-serif":      {},
		"ui-sans-serif": {},
		"ui-monospace":  {},
		"ui-rounded":    {},
		"emoji":         {},
		"math":          {},
		"fangsong":      {},
		"inherit":       {},
		"initial":       {},
		"revert":        {},
		"revert-layer":  {},
		"unset":         {},
	}
	_, isGeneric := genericFontFamilies[fontName]
	return isGeneric
}
