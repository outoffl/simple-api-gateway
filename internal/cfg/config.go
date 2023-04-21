package cfg

// all this pkg will autoload init func in main call
// @dong.wei
import (
	"fmt"
	"os"
	"path/filepath"
)

const cfgPath = "../conf/"

func readCfgBytes(fileName string) ([]byte, error) {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	cfgFile := fmt.Sprintf("%s/%s%s.yml", dir, cfgPath, fileName)
	cfgBytes, err := os.ReadFile(cfgFile)
	return cfgBytes, err
}
