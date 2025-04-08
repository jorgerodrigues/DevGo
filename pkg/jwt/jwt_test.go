package jwt

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const secret = "testsecret"
// TestDecode contains all JWT decode test scenarios
func TestDecode(t *testing.T) {
	// Helper function to create test tokens

	// Helper function to check if output contains expected strings
	containsAll := func(t *testing.T, output string, expected []string) {
		for _, str := range expected {
			if !strings.Contains(output, str) {
				t.Errorf("Expected output to contain '%s', but it didn't", str)
			}
		}
	}

	t.Run("valid JWT token", func(t *testing.T) {
		// Create a valid token
		token, _ := CreateToken(jwt.MapClaims{
			"sub":   "1234567890",
			"name":  "John Doe",
			"admin": true,
			"iat":   time.Now().Unix() - 60,          // Issued 1 minute ago
			"exp":   time.Now().Unix() + 60*60,       // Expires in 1 hour
			"nbf":   time.Now().Unix() - 60*60*24*30, // Valid since 30 days ago
		}, secret)


		// Call decode function
		var output bytes.Buffer
		err := Decode(&output, token, DecodeOptions{})
		if err != nil {
			t.Fatalf("Decode returned an error: %v", err)
		}

		// Check output
		outputStr := output.String()
		expectedStrings := []string{
			"Header", "alg", "HS256", "typ", "JWT",
			"Payload", "sub", "1234567890", "name", "John Doe", "admin", "true",
			"iat", "exp", "nbf", "Signature",
		}
		containsAll(t, outputStr, expectedStrings)
	})

	t.Run("expired JWT token", func(t *testing.T) {
		// Create an expired token
		token, _ := CreateToken( jwt.MapClaims{
			"sub":  "1234567890",
			"name": "John Doe",
			"exp":  time.Now().Unix() - 3600, // Expired 1 hour ago
		}, secret)

		// Call decode function
		var output bytes.Buffer
		err := Decode(&output, token, DecodeOptions{})
		if err != nil {
			t.Fatalf("Decode returned an error: %v", err)
		}

		// Check output
		outputStr := output.String()
		if !strings.Contains(outputStr, "EXPIRED") {
      t.Errorf("Expected output to contain 'EXPIRED', but it didn't: %v", outputStr)
		}
	})

	t.Run("invalid JWT format", func(t *testing.T) {
		// Call decode function with invalid token
		var output bytes.Buffer
		err := Decode(&output, "not.a.validtoken", DecodeOptions{})
		
		// Should return an error
		if err == nil {
			t.Fatal("Expected an error for invalid token, but got nil")
		}
		
		// Check error message
		if !strings.Contains(err.Error(), "invalid token") {
			t.Errorf("Expected error message to contain 'invalid token', got: %v", err)
		}
	})

	t.Run("malformed JWT token", func(t *testing.T) {
		// Malformed token (valid format but corrupt data)
		malformedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.incorrect-signature"
		
		// Call decode function
		var output bytes.Buffer
		err := Decode(&output, malformedToken, DecodeOptions{})
		if err != nil {
			t.Fatalf("Decode returned an error for malformed token: %v", err)
		}
		
		// Check output
		outputStr := output.String()
		expectedStrings := []string{"Header", "Payload", "Signature", "invalid"}
		containsAll(t, outputStr, expectedStrings)
	})

	t.Run("token without signature", func(t *testing.T) {
		// Token with no signature
		unsignedToken := "eyJhbGciOiJub25lIiwidHlwIjoiSldUIn0.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIn0."
		
		// Call decode function
		var output bytes.Buffer
		err := Decode(&output, unsignedToken, DecodeOptions{})
		if err != nil {
			t.Fatalf("Decode returned an error: %v", err)
		}
		
		// Check output
		if !strings.Contains(output.String(), "NO SIGNATURE") {
			t.Error("Output does not indicate missing signature")
		}
	})

	t.Run("empty token", func(t *testing.T) {
		// Call decode function with empty token
		var output bytes.Buffer
		err := Decode(&output, "", DecodeOptions{})
		
		// Should return an error
		if err == nil {
			t.Fatal("Expected an error for empty token, but got nil")
		}
		
		// Check error message
		if !strings.Contains(err.Error(), "empty token") {
			t.Errorf("Expected error message to contain 'empty token', got: %v", err)
		}
	})

	t.Run("JSON output", func(t *testing.T) {
		// Create a valid token
		token, _ := CreateToken(jwt.MapClaims{
			"sub":  "1234567890",
			"name": "John Doe",
			"iat":  time.Now().Unix(),
		}, secret)
		
		// Call decode function with JSON output option
		var output bytes.Buffer
		err := Decode(&output, token, DecodeOptions{JSONOutput: true})
		if err != nil {
			t.Fatalf("Decode returned an error: %v", err)
		}
		
		// Verify output is valid JSON
		var result map[string]interface{}
		err = json.Unmarshal(output.Bytes(), &result)
		if err != nil {
			t.Fatalf("Failed to parse JSON output: %v", err)
		}
		
		// Check JSON structure
		requiredFields := []string{"header", "payload", "signature", "valid"}
		for _, field := range requiredFields {
			if _, ok := result[field]; !ok {
				t.Errorf("JSON output missing '%s' field", field)
			}
		}
	})
}

// Table-driven tests for specific parsing scenarios
func TestTokenParsing(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		token         string
		expectedError bool
		errorContains string
	}{
		{
			name:          "valid basic token",
			token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.Rq8IxqeX7eA6GgYxlcHdPFVRNFFZc5rEI3MQTZZbK6A",
			expectedError: false,
		},
		{
			name:          "too few segments",
			token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0",
			expectedError: true,
			errorContains: "invalid token",
		},
		{
			name:          "invalid base64 in header",
			token:         "!invalid!.eyJzdWIiOiIxMjM0NTY3ODkwIn0.Rq8IxqeX7eA6GgYxlcHdPFVRNFFZc5rEI3MQTZZbK6A",
			expectedError: true,
			errorContains: "invalid token",
		},
		{
			name:          "invalid base64 in payload",
			token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.!invalid!.Rq8IxqeX7eA6GgYxlcHdPFVRNFFZc5rEI3MQTZZbK6A",
			expectedError: true,
			errorContains: "invalid token",
		},
		{
			name:          "invalid JSON in header",
			token:         "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCIsaW52YWxpZCB9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.Rq8IxqeX7eA6GgYxlcHdPFVRNFFZc5rEI3MQTZZbK6A",
			expectedError: true,
			errorContains: "invalid token",
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var output bytes.Buffer
			err := Decode(&output, tc.token, DecodeOptions{})

			// Check if error matches expectation
			if tc.expectedError && err == nil {
				t.Errorf("Expected error but got nil")
			}
			if !tc.expectedError && err != nil {
				t.Errorf("Expected no error but got: %v", err)
			}
			if tc.expectedError && err != nil && tc.errorContains != "" {
				if !strings.Contains(err.Error(), tc.errorContains) {
					t.Errorf("Expected error to contain '%s', got: %v", tc.errorContains, err)
				}
			}
		})
	}
}

