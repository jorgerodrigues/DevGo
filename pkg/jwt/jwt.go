package jwt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Define the options struct that would be used by the implementation
type DecodeOptions struct {
	JSONOutput bool
}

func DecodeTokenFromArgs(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no token provided")
	}
	token := args[0]

	err := Decode(nil, token, DecodeOptions{JSONOutput: true})
	if err != nil {
		return err
	}
	return nil
}

func CreateToken(claims jwt.MapClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func Decode(output *bytes.Buffer, jwtStr string, opts DecodeOptions) error {
	// Check for empty token
	if jwtStr == "" {
		return fmt.Errorf("empty token")
	}

	var out io.Writer = os.Stdout
	if output != nil {
		out = output
	}

	// Parse the token without validating signature
	token, err := jwt.Parse(jwtStr, func(token *jwt.Token) (interface{}, error) {
		// We're not validating the signature here, just returning nil
		switch token.Method.Alg() {
		case "HS256", "HS384", "HS512":
			// For HMAC algorithms, return a nil []byte
			return []byte(nil), nil
		case "RS256", "RS384", "RS512", "PS256", "PS384", "PS512":
			// We're not actually validating, so just tell the parser we don't have the right key
			return nil, fmt.Errorf("RSA key is not provided")
		case "ES256", "ES384", "ES512":
			// We're not actually validating, so just tell the parser we don't have the right key
			return nil, fmt.Errorf("ECDSA key is not provided")
		case "none":
			// allow parsing of unsigned tokens since it is safe in our context
			return jwt.UnsafeAllowNoneSignatureType, nil
		default:
			return nil, fmt.Errorf("signing method %v not supported", token.Method.Alg())
		}
	})

	// Handle parsing errors, but allow signature validation errors
	var valid bool

	if err != nil {
		if !strings.Contains(err.Error(), "signature is invalid") {
			return fmt.Errorf("invalid token: %w", err)
		}
		valid = false
	} else {
		valid = true
	}

	// Extract token parts
	parts := strings.Split(jwtStr, ".")

	// Check token structure (should have at least 2 parts - header and payload)
	if len(parts) < 2 {
		return fmt.Errorf("invalid token format: token must have at least header and payload")
	}

	// Get claims from the parsed token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("invalid claims format")
	}

	// Check for expiration
	var expired bool
	if exp, ok := claims["exp"]; ok {
		switch v := exp.(type) {
		case float64:
			expired = time.Unix(int64(v), 0).Before(time.Now())
		}
	}

	// Determine signature status
	var signatureStatus string
	if len(parts) < 3 || parts[2] == "" {
		signatureStatus = "NO SIGNATURE"
	} else if valid {
		signatureStatus = "valid (not verified)"
	} else {
		signatureStatus = "invalid"
	}

	// Generate output based on format option
	if opts.JSONOutput == true {
		result := map[string]interface{}{
			"header":    token.Header,
			"payload":   claims,
			"signature": signatureStatus,
			"valid":     valid,
			"expired":   expired,
		}

		jsonBytes, err := json.MarshalIndent(result, "", "  ")

		if err != nil {
			return fmt.Errorf("error marshaling to JSON: %w", err)
		}

		if jsonBytes == nil {
			return fmt.Errorf("error marshaling to JSON: result is nil")
		}

		fmt.Fprintln(out, string(jsonBytes))
	} else {
		// Standard text output
		fmt.Fprintln(out, "Header:")
		headerJSON, _ := json.MarshalIndent(token.Header, "  ", "  ")
		if headerJSON == nil {
			return fmt.Errorf("error marshaling to JSON: header is nil")
		}
		fmt.Fprintf(out, "  %s\n\n", headerJSON)

		fmt.Fprintln(out, "Payload:")
		claimsJSON, _ := json.MarshalIndent(claims, "  ", "  ")
		if claimsJSON == nil {
			return fmt.Errorf("error marshaling to JSON: claims is nil")
		}
		fmt.Fprintf(out, "  %s\n\n", claimsJSON)

		fmt.Fprintln(out, "Signature:", signatureStatus)

		// Display status
		if expired {
			fmt.Fprintln(out, "Status: EXPIRED")
		} else if valid {
			fmt.Fprintln(out, "Status: VALID (signature not verified)")
		} else {
			fmt.Fprintln(out, "Status: INVALID")
		}
	}

	return nil

}
