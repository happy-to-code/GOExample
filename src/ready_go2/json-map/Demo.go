package main

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func main() {
	var s = "{}"
	var s2 = "{\"keySpec\":\"sm2\",\"password\":\"123\",\"pubFormat\":\"sm2_pem\"}"

	valid1 := json.Valid([]byte(s))
	fmt.Println(valid1)

	valid2 := json.Valid([]byte(s2))
	fmt.Println(valid2)

	var a = "did:example:123"
	// func QueryEscape(s string) string
	escape := url.QueryEscape(a)
	fmt.Println(escape)

	unescape, err := url.QueryUnescape(escape)
	if err != nil {
		panic(err)
	}
	fmt.Println(unescape)

	var f float64 = 30000000000
	sprintf := fmt.Sprintf("%f", f)
	fmt.Println(sprintf)
}

// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0x44e75652460e4593b388c3e1fb0808d20acf52c6', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0xca8debf2e3cae385086b92307c184ec077b9648e', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0x3ac60f7bb7a54de281c7b6675d2dcd1825160ea7', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0x2d89fa9137462524716b54d11f393e0d76c17cf1', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0xb669724ecc304926626b24f403f58c9b9354d89f', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0x540fcbaea29ce96525451130996c44f113d9b605', 'standard', '', '', '', '');
// INSERT INTO "public"."casbin_rule_test"("p_type", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '0x2c2d457bc16b12435f5d3b00281e917e5ce4c2cd', 'standard', '', '', '', '');
//
//
//
//
// 0x44e75652460e4593b388c3e1fb0808d20acf52c6
// 0xca8debf2e3cae385086b92307c184ec077b9648e
// 0x3ac60f7bb7a54de281c7b6675d2dcd1825160ea7
// 0x2d89fa9137462524716b54d11f393e0d76c17cf1
// 0xb669724ecc304926626b24f403f58c9b9354d89f
// 0x540fcbaea29ce96525451130996c44f113d9b605
// 0x2c2d457bc16b12435f5d3b00281e917e5ce4c2cd

// 0xea608f7138e9bf0ffc900e62cbbcd45c1ac7df82
// 0x0aa0f81233493a6a112fc0c97234428cb26090ca
