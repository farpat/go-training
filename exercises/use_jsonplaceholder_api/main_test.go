package use_jsonplaceholder_api

import (
	"testing"
)

func TestRetrieveFirstPost(t *testing.T) {
	// ARRANGE

	// ACT
	result, err := RetrieveFirstPost()

	// ASSERT
	if err != nil {
		t.Errorf("RetrieveFirstPost() = %v; expected nil", err)
	}

	if result.Id == 0 {
		t.Errorf("RetrieveFirstPost() = %v; expected non-zero ID", result)
	}

}

func TestUpdateFirstPost(t *testing.T) {
	// ARRANGE

	// ACT
	err := UpdateFirstPost()

	// ASSERT
	if err != nil {
		t.Errorf("UpdateFirstPost() = %v; expected nil", err)
	}
}
