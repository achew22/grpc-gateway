/*
 * A Bit of Everything
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Contact: none@example.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package abe

// Nested is nested type.
type ABitOfEverythingNested struct {

	// name is nested field.
	Name string `json:"name,omitempty"`

	Amount int64 `json:"amount,omitempty"`

	Ok *NestedDeepEnum `json:"ok,omitempty"`
}
