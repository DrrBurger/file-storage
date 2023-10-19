package usecases

import (
	"bytes"
	"testing"

	"file-storage/internal/entities"
	"github.com/stretchr/testify/assert"
)

func TestSplitAndCombineFile(t *testing.T) {
	tests := []struct {
		name     string
		file     entities.File
		numParts int
		expected []byte
		hasError bool
	}{
		{
			name:     "Valid Split and Combine",
			file:     entities.File{ID: "testfile", Data: []byte("1234567890")},
			numParts: 2,
			expected: []byte("1234567890"),
		},
		{
			name:     "Split with Zero Parts",
			file:     entities.File{ID: "testfile", Data: []byte("1234567890")},
			numParts: 0,
			hasError: true,
		},
		{
			name:     "Split Empty File",
			file:     entities.File{ID: "testfile", Data: []byte("")},
			numParts: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parts, err := SplitFileIntoParts(tt.file, tt.numParts)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)

			if tt.numParts == -1 { // To handle the reverse order scenario
				reversedParts := []entities.FilePart{parts[1], parts[0]}
				parts = reversedParts
			}

			combined, err := CombineFileFromParts(parts)
			assert.NoError(t, err)

			assert.True(t, bytes.Equal(tt.expected, combined.Data), "Expected and combined data should be the same")
		})
	}
}

func TestCombineParts(t *testing.T) {
	tests := []struct {
		name     string
		parts    []entities.FilePart
		expected []byte
		hasError bool
	}{
		{
			name: "Combine Single Part",
			parts: []entities.FilePart{
				{FileID: "part1", Data: []byte("12345")},
			},
			expected: []byte("12345"),
		},
		{
			name:     "Combine Empty Parts",
			parts:    []entities.FilePart{},
			hasError: true, // We expect an error for this scenario
		},
		{
			name: "Combine Multiple Parts",
			parts: []entities.FilePart{
				{FileID: "part1", Data: []byte("123")},
				{FileID: "part2", Data: []byte("456")},
				{FileID: "part3", Data: []byte("789")},
			},
			expected: []byte("123456789"),
		},
		{
			name: "Combine Out Of Order Parts",
			parts: []entities.FilePart{
				{FileID: "part2", Data: []byte("456")},
				{FileID: "part1", Data: []byte("123")},
				{FileID: "part3", Data: []byte("789")},
			},
			expected: []byte("456123789"),
		},
		{
			name: "Combine With Missing Parts",
			parts: []entities.FilePart{
				{FileID: "part1", Data: []byte("123")},
				{FileID: "part3", Data: []byte("789")},
			},
			expected: []byte("123789"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combined, err := CombineFileFromParts(tt.parts)
			if tt.hasError {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.True(t, bytes.Equal(tt.expected, combined.Data), "Expected and combined data should be the same")
		})
	}
}
