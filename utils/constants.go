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
  "REACT_VERSION_SCORE": "Score",
  "REACT_VERSION_MAX_SCORE": "MaxScore",
  "REACT_VERSION_MIN_SCORE": "MinScore",
  "REACT_VERSION_WEIGHT": "Weight",
}
