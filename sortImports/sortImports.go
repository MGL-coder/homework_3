package sortImports

import (
	"os"
	"sort"
	"strings"
)

func Sort(path string) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split( string(input), "\n")
	var imports []string
	insideImportBlock := false
	var importsLineNum int

	for i, line := range lines {
		if insideImportBlock {
			if strings.Contains(line, ")") {
				insideImportBlock = false
			} else {
				imports = append(imports, strings.TrimSpace(line))
			}
		} else if strings.Contains(line, "import (") {
			insideImportBlock = true
			importsLineNum = i + 1
		}
	}

	sort.Strings(imports)

	for i, _ := range imports {
		lines[importsLineNum+i] = "\t" + imports[i]
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(path, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}
