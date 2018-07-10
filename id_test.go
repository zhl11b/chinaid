package chinaid

import (
	"testing"
	"time"
)

// TestIDNum 正确性测试
func TestIDNum(t *testing.T) {
	testID := "420683199006041237"
	detail, err := IDCard(testID).Decode()
	t.Log(detail)
	if err != nil {
		t.Error("Valid id parse failed.")
	}
	if detail.Province != "湖北省" {
		t.Error("Province check failed")
	}
	if detail.City != "襄阳市" {
		t.Error("City check failed")
	}
	birth, err := time.Parse("20060102", "19900604")
	if detail.Birthday != birth {
		t.Error("Birthday check failed")
	}
	if detail.Sex != Male {
		t.Error("Sex check failed")
	}
}
