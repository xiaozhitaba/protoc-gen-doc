// AUTOGENERATED CODE. DO NOT EDIT.
package protoc_gen_doc

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
)

var embeddedResources = map[string]string{
	"docbook.tmpl": "H4sIAAAAAAAA/+xZ30/bMBB+719h5XGIhIkhTZNbpMGqaQKEgO3dTa6tNcfObKeAov7vk5M0v1MXWgqbeEGy7/Odfb7vOzfg04eQoQVIRQUfOh/dIwcB90VA+Wzo/LwbH352TkcDTKSmPoPRACGsqWYwupZCC18wdC78OASuiaaCYy+zDhBKEkn4DJA7pgzUcmmWKvANypgLR0niXpEQlsvKWoRwRCQxtnNQvqSRWWYg6XSKKPxfglJklocogyAaDJ0kcccxY1kAJ1tYjXwh+Kwj+ibxzQ7oFLnfiRpTYMEqvnFPJgzQVJIQhg5hrAhchMY+I0pxErZ2URpQ5raxMeNiJkUcIV8wNXQ+VZwjhM1kBL4x3tNAz4fOB8fbGnHknthBx02IngMJqjMIYSnu6zMIYeBaPo7S02IvG3RD7h4jWI+4IBNg6yGVC+0EYq+xR+y1DoL1RASNdZV6r1WD9eAVAqzbN2aU/0bmD/Cysk1KTGXnVZQNsWdgo/X+zAqTLVvcJgeyqj+HKYmZ/kVYDMslyodfUIqumpIEeNATo5Vqk8UUXk9/PdnYywhQUNVL+Vaysuqh4Oi3Bw3cqNzueXoFSkOAyggWyp7sg7Jvg9RFTrYl9leiLIirOJyAfGXud1SZNUevxP+2vzPBNaGc8lnDc2l4usZk12I9XGe7fa7UpM5eR3GwV3vjVE1lkfA43P9rZVdqlybbJnHH+5C4LbXJnO0fkJQs3zuXky2puWeS9dAqH/S1+PrvjaLMzQ+SQwYLYP09G1M+FTIkbC1r3rv6e1d/7+r/TVevcX4T4cnr4xbkgvrP+wKxv37e0csvQc/F2/jE8OKilZ0V2bv+DfyJQWlkl68bUJHgCjaAvrhG5Ve5B4HK89NQk3x2a7la5bTlPpt+qv83/HhpGCrS0vmd9NYnjMjs+Z0WXJ229gdL/3PFRscO+4kNsFt7mwutW8pv3I2k0KKfkat3hdAmgf2As4MDq5MfZEGsoOtHPRe8D9Yotxbr25wvu05aEHXG96Vl1YrSr/YVWlXG685gtMIkzIo6i6KNvJnMbQTMstcLbbG1ydUGU+s87ej4FUoOsFf82+NvAAAA//+VaunsKBkAAA==",
	"html.tmpl": "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",
	"markdown.tmpl": "H4sIAAAAAAAA/+RWQW/bPAy9+1fwS77D2sLJvUh8WLtiGNqiaItdimFVEyYxoEieJQcrIv/3gZJly3bcBih2mi82KYki33uiPIa7XGq5kBwu5aLYotBMp1JEMwaCbXE+0jIbTZMoGo/hkb1wBLmCCyk0Cq2i/T5nYo0wuUo5qrKM9vv/VynHn7QWzucwuWVbLMvoFJ72+8r48Wlcf59EAHWQG1SKrW0cAAC35lqKdbjuquA8XItiaefXUb6IYvvREL81CpVKEcShCmOOO+TQDNt4TcVlGWM9NhD7AfNduugU+X52/n0KTw8LxlkO3xkvEB5fM6Q0lHXGO3LGmpwn0fH01GzXuRDnswwYT9diPsrT9UaPkhmDTY6r+WhMqkgeZTabsmQ2zRKSR72WNp5colrkaUZicp4+yeGuTdW089iFa4g7HDJdweQrU1cp8iUFNGA/wVhYwMA1e0EOBoKVYCIDMT3g3tA2qwdMiB7Fh9hSZpo6wdTyov1CAp194mbbLOz0bhGuhEtcsYJrS2hZQmWeu9nhUCUCm1tbFg0YLe2aRqsNKJ+ZotdtsX3BfAicPkA1PsNANXsfBKuFlbOpk7BUpGLdHXHp/XXU3NDsvzgGFEvYVuqEOE4C0fqe8kHFGqCx97A/Amhb2SDIb0AXgNAvHqlKX/khOQXlDzQ9D8XBbhn9k3p87gnyeVCRDRct6DtqDK6QdwX5lhhvUG/k0mvyHn8VqLRn5R5VJoVCbw+y0iWga3ZtE94FlMBwZ61S6jbYyt3us26+y/moBW8cjR4bqkLcE1GdILr1+rcxQTvJ6LfKY3crNSowcHF25l3f2I7577tXvZGisg5DG2AI/RNxQPourQbaUCj2l88BMYJpAm1XxQOlXJ+KLAvHKPfQdvl7Tw3PnwAAAP//WuevgFwKAAA=",
}

func fetchResource(name string) ([]byte, error) {
	raw, ok := embeddedResources[name]
	if !ok {
		return nil, fmt.Errorf("Could not find resource for '%s'", name)
	}

	compressed, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	buf := bytes.NewBuffer(compressed)
	
	r, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	if _, err := io.Copy(&out, r); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}
