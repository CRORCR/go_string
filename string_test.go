package go_string

import "testing"

func TestStringClass_DesensitizeMobile(t *testing.T) {
	result := String.DesensitizeMobile(`15212341234`)
	if result != `152****1234` {
		t.Error(`15212341234 error`)
	}

	result1 := String.DesensitizeMobile(`123456`)
	if result1 != `12****56` {
		t.Error(`123456 error`)
	}
}

func TestStringClass_DesensitizeEmail(t *testing.T) {
	result := String.DesensitizeEmail(`hello@qq.com`)
	if result != `hello****@qq.com` {
		t.Error(`hello@qq.com error`)
	}

	result1 := String.DesensitizeEmail(`abc@qq.com`)
	if result1 != `abc****@qq.com` {
		t.Error(`abc@qq.com error`)
	}

	result2 := String.DesensitizeEmail(`a@163.com`)
	if result2 != `a****@163.com` {
		t.Error(`a@163.com error`)
	}
}

func TestStringClass_SpanLeft(t *testing.T) {
	if String.SpanLeft(`1234`, 6, `0`) != `001234` {
		t.Error()
	}
}
