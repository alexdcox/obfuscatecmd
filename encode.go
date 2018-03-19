package main

import (
	"fmt"
	"os/exec"
	"strings"
	"bytes"
)

func main() {
	payload := `echo somesupersecretthing`

	for i := 1; i <= 12; i++ {
		payload = packWithDecoder(obfuscate(payload))
	}

	fmt.Println(payload)
}

func obfuscate(s string) string {
	cmd := exec.Command("base64")
	cmd.Stdin = strings.NewReader(s)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	return out.String()
}

func packWithDecoder(encoded string) string {
	return fmt.Sprintf(
		"a='%s'; %s $a | %s | %s;",
		encoded,
		substringCommand(encoded, "echo"),
		substringCommand(encoded, "base64 -D"),
		substringCommand(encoded, "sh"),
	)
}

func substringCommand(scaffold, blueprint string) string {
	ans := ""

next:
	for _, bc := range blueprint {
		for si, sc := range scaffold {
			if sc == bc {
				ans += fmt.Sprintf("${a:%d:1}", si)
				continue next
			}
		}

		ans += string(bc)
	}

	return ans
}
