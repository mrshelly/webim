package util

import "testing"

func TestGetCurrentTimestamp(t *testing.T) {
	ts := GetCurrentTimestamp()
	if ts == 0 {
		t.Error()
	}
}

func TestEncryptPassword(t *testing.T) {
	s := EncryptPassword("test")
	if len(s) <= 0 {
		t.Errorf("test %s error", "EncryptPassword")
	}
}

func TestEncryptPasswordWithSalt(t *testing.T) {
	s := EncryptPasswordWithSalt("test", "test")
	if len(s) <= 0 {
		t.Errorf("test %s error", "EncryptPasswordWithSalt")
	}
}
