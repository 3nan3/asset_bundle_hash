package pkg

import (
	"fmt"
)

const version = "0.1.0"

func PrintVersion() {
	fmt.Printf("asset_bundle_hash %s\n", version)
}
