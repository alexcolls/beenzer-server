package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

// SQLFilter sanitizes input to prevent SQL injection
func SQLFilter(input string) string {
	// Remove potentially dangerous characters
	reg := regexp.MustCompile(`[^\w\s@._-]`)
	filtered := reg.ReplaceAllString(input, "")
	
	// Remove SQL keywords
	sqlKeywords := []string{
		"SELECT", "INSERT", "UPDATE", "DELETE", "DROP", "CREATE", "ALTER",
		"EXEC", "EXECUTE", "UNION", "DECLARE", "SCRIPT", "JAVASCRIPT",
	}
	
	upper := strings.ToUpper(filtered)
	for _, keyword := range sqlKeywords {
		upper = strings.ReplaceAll(upper, keyword, "")
	}
	
	// Return original case but filtered
	if len(upper) != len(filtered) {
		return ""
	}
	
	return strings.TrimSpace(filtered)
}

// ConcatPubKeys concatenates two public keys in alphabetical order
func ConcatPubKeys(pubkey1, pubkey2 string) string {
	keys := []string{pubkey1, pubkey2}
	sort.Strings(keys)
	return fmt.Sprintf("_%s_%s_", keys[0], keys[1])
}

// GetTime returns the current time formatted as HH:MM:SS
func GetTime() string {
	return time.Now().Format("15:04:05")
}

// GetDate returns the current date formatted as YYYY-MM-DD
func GetDate() string {
	return time.Now().Format("2006-01-02")
}

// GetTimestamp returns the current Unix timestamp in milliseconds
func GetTimestamp() int64 {
	return time.Now().UnixMilli()
}

// IsValidPubkey checks if a string is a valid Solana public key (base58, 32-44 chars)
func IsValidPubkey(pubkey string) bool {
	if len(pubkey) < 32 || len(pubkey) > 44 {
		return false
	}
	// Base58 alphabet
	validChars := regexp.MustCompile(`^[1-9A-HJ-NP-Za-km-z]+$`)
	return validChars.MatchString(pubkey)
}

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
