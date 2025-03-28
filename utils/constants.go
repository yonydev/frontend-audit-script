package utils

var IgnoredDirsAndFiles = map[string]bool{
	"node_modules":      true,
	"__mocks__":         true,
	"__tests__":         true,
	"__test__":          true,
	"tests":             true,
	"test":              true,
	"dist":              true,
	"build":             true,
	"coverage":          true,
	".git":              true,
	".gitignore":        true,
	".vscode":           true,
	".idea":             true,
	".turbo":            true,
	"storybook":         true,
	".storybook":        true,
	"stories":           true,
	".stories":          true,
	"storybook-static":  true,
	".storybook-static": true,
	"package-lock.json": true,
}

var CommonIconLibs = []string{
	"@radix-ui/react-icons",
	"@fortawesome/fontawesome",
	"@heroicons/react",
	"@tabler/icons-react",
	"@material-ui/icons",
	"@iconify/react",
	"@clipmx/cods-icons",
	"@clipmx/clip-icons",
	"boxicons",
	"bootstrap-icons",
	"react-icons",
	"material-icons",
	"feather-icons",
	"ionicons",
	"heroicons",
	"lucide-react",
}

var MuiExtraLibs = []string{"@mui/lab", "@mui/system"}

// Allowed (true) and disallowed (false) styling libraries
var StylingLibs = map[string]bool{
	"jss":               false,
	"react-jss":         false,
	"styled-components": false,
	"@emotion/css":      false,
	"@emotion/react":    true,
	"@emotion/styled":   true,
}

var AssetsExtensions = []string{
	".jpg",
	".jpeg",
	".png",
	".gif",
	".svg",
	".webp",
}

var ReactVersionEnvVars = map[string]string{
	"REACT_VERSION_EVAL_NAME": "Name",
	"REACT_VERSION_SCORE":     "Score",
	"REACT_VERSION_MAX_SCORE": "MaxScore",
	"REACT_VERSION_MIN_SCORE": "MinScore",
	"REACT_VERSION_WEIGHT":    "Weight",
}

var IconLibsEnvVars = map[string]string{
	"ICON_LIBS_EVAL_NAME": "Name",
	"ICON_LIBS_SCORE":     "Score",
	"ICON_LIBS_MAX_SCORE": "MaxScore",
	"ICON_LIBS_MIN_SCORE": "MinScore",
	"ICON_LIBS_WEIGHT":    "Weight",
}

var MuiExtraLibsEnvVars = map[string]string{
	"MUI_EXTRA_LIBS_EVAL_NAME": "Name",
	"MUI_EXTRA_LIBS_SCORE":     "Score",
	"MUI_EXTRA_LIBS_MAX_SCORE": "MaxScore",
	"MUI_EXTRA_LIBS_MIN_SCORE": "MinScore",
	"MUI_EXTRA_LIBS_WEIGHT":    "Weight",
}

var StylingLibsEnvVars = map[string]string{
	"STYLING_LIBS_EVAL_NAME": "Name",
	"STYLING_LIBS_SCORE":     "Score",
	"STYLING_LIBS_MAX_SCORE": "MaxScore",
	"STYLING_LIBS_MIN_SCORE": "MinScore",
	"STYLING_LIBS_WEIGHT":    "Weight",
}

var ThemeProvidersEnvVars = map[string]string{
	"THEME_PROVIDERS_EVAL_NAME": "Name",
	"THEME_PROVIDERS_SCORE":     "Score",
	"THEME_PROVIDERS_MAX_SCORE": "MaxScore",
	"THEME_PROVIDERS_MIN_SCORE": "MinScore",
	"THEME_PROVIDERS_WEIGHT":    "Weight",
}

var WebFontsEnvVars = map[string]string{
	"WEB_FONTS_EVAL_NAME": "Name",
	"WEB_FONTS_SCORE":     "Score",
	"WEB_FONTS_MAX_SCORE": "MaxScore",
	"WEB_FONTS_MIN_SCORE": "MinScore",
	"WEB_FONTS_WEIGHT":    "Weight",
}

var AssetsEnvVars = map[string]string{
	"ASSETS_EVAL_NAME": "Name",
	"ASSETS_SCORE":     "Score",
	"ASSETS_MAX_SCORE": "MaxScore",
	"ASSETS_MIN_SCORE": "MinScore",
	"ASSETS_WEIGHT":    "Weight",
}
