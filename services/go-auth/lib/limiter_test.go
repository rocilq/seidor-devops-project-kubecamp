package lib

import (
	"testing"
	"time"
)

func TestValidateIPRequest_NewIP(t *testing.T) {
	rateLimiter := RateLimiterRecord{}

	err := rateLimiter.ValidateIPRequest("192.168.1.1")
	if err != nil {
		t.Fatalf("Expected no error for new IP, but got %v", err)
	}

	record, exists := rateLimiter["192.168.1.1"]
	if !exists || record.count != 1 {
		t.Error("New IP should have been added with count 1")
	}
}

func TestValidateIPRequest_IncrementCount(t *testing.T) {
	rateLimiter := RateLimiterRecord{}
	rateLimiter.ValidateIPRequest("192.168.1.1")
	rateLimiter.ValidateIPRequest("192.168.1.1")

	record, exists := rateLimiter["192.168.1.1"]
	if !exists || record.count != 2 {
		t.Error("IP count should have been incremented to 2")
	}
}

func TestValidateIPRequest_RateLimit(t *testing.T) {
	rateLimiter := RateLimiterRecord{}
	rateLimiter.ValidateIPRequest("192.168.1.1")
	rateLimiter.ValidateIPRequest("192.168.1.1")
	rateLimiter.ValidateIPRequest("192.168.1.1")

	err := rateLimiter.ValidateIPRequest("192.168.1.1")
	if err == nil || err.Error() != "rate limit exceeded" {
		t.Error("Expected 'rate limit exceeded' error")
	}
}

func TestValidateIPRequest_AfterTimePassed(t *testing.T) {
	rateLimiter := RateLimiterRecord{}
	rateLimiter["192.168.1.1"] = UserRecord{
		count:           3,
		accessTimeStamp: time.Now().Add(-2 * time.Minute),
	}

	err := rateLimiter.ValidateIPRequest("192.168.1.1")
	if err != nil {
		t.Fatalf("Expected no error after a minute has passed, but got %v", err)
	}

	record, exists := rateLimiter["192.168.1.1"]
	if !exists || record.count != 1 {
		t.Error("IP count should have been reset to 1 after time passed")
	}
}
